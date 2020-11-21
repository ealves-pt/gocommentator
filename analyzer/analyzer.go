// Copyright (c) 2013, Google Inc.
// All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file or at
// https://developers.google.com/open-source/licenses/bsd.

package analyzer

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "gocommentator",
	Doc:      "Complains when exported names are missing comments or if they are not of the right form.",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

var commonMethods = map[string]bool{
	"Error":     true,
	"Read":      true,
	"ServeHTTP": true,
	"String":    true,
	"Write":     true,
	"Unwrap":    true,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.GenDecl)(nil),
		(*ast.FuncDecl)(nil),
		(*ast.TypeSpec)(nil),
		(*ast.ValueSpec)(nil),
	}

	// last GenDecl entered
	var lastGen *ast.GenDecl

	// Set of GenDecls that have already had missing comments flagged.
	genDeclMissingComments := make(map[*ast.GenDecl]bool)

	inspector.Preorder(
		nodeFilter,
		func(node ast.Node) {
			var err error
			switch v := node.(type) {
			case *ast.GenDecl:
				if v.Tok != token.IMPORT {
					lastGen = v
				}
				return
			case *ast.FuncDecl:
				err = lintFuncDoc(v)
			case *ast.TypeSpec:
				doc := v.Doc
				if doc == nil {
					doc = lastGen.Doc
				}
				err = lintTypeDoc(v, doc)
			case *ast.ValueSpec:
				err = lintValueSpecDoc(v, lastGen, genDeclMissingComments)
			}

			if err != nil {
				pass.Reportf(node.Pos(), err.Error())
			}
		},
	)

	return nil, nil
}

// lintFuncDoc examines doc comments on functions and methods.
// It complains if they are missing, or not of the right form.
// It has specific exclusions for well-known methods (see commonMethods above).
func lintFuncDoc(fn *ast.FuncDecl) error {
	if !ast.IsExported(fn.Name.Name) {
		return nil
	}

	kind := "function"
	name := fn.Name.Name
	if fn.Recv != nil && len(fn.Recv.List) > 0 {
		// method
		kind = "method"
		recv := receiverType(fn)

		if !ast.IsExported(recv) {
			return nil
		}

		if commonMethods[name] {
			return nil
		}

		name = recv + "." + name
	}

	if fn.Doc == nil {
		return fmt.Errorf(
			"exported %s %s should have comment or be unexported",
			kind,
			name,
		)
	}

	s := fn.Doc.Text()
	prefix := fn.Name.Name + " "
	if !strings.HasPrefix(s, prefix) {
		return fmt.Errorf(
			`comment on exported %s %s should be of the form "%s..."`,
			kind,
			name,
			prefix,
		)
	}

	return nil
}

// lintTypeDoc examines the doc comment on a type.
// It complains if they are missing from an exported type,
// or if they are not of the standard form.
func lintTypeDoc(t *ast.TypeSpec, doc *ast.CommentGroup) error {
	if !ast.IsExported(t.Name.Name) {
		return nil
	}

	if doc == nil {
		return fmt.Errorf(
			"exported type %v should have comment or be unexported",
			t.Name,
		)
	}

	s := doc.Text()
	articles := [3]string{"A", "An", "The"}
	for _, a := range articles {
		if strings.HasPrefix(s, a+" ") {
			s = s[len(a)+1:]
			break
		}
	}
	if !strings.HasPrefix(s, t.Name.Name+" ") {
		return fmt.Errorf(
			`comment on exported type %v should be of the form "%v ..." (with optional leading article)`,
			t.Name,
			t.Name,
		)
	}

	return nil
}

// lintValueSpecDoc examines package-global variables and constants.
// It complains if they are not individually declared,
// or if they are not suitably documented in the right form (unless they are in a block that is commented).
func lintValueSpecDoc(vs *ast.ValueSpec, gd *ast.GenDecl, genDeclMissingComments map[*ast.GenDecl]bool) error {
	kind := "var"
	if gd.Tok == token.CONST {
		kind = "const"
	}

	if len(vs.Names) > 1 {
		// Check that none are exported except for the first.
		for _, n := range vs.Names[1:] {
			if ast.IsExported(n.Name) {
				return fmt.Errorf(
					"exported %s %s should have its own declaration",
					kind,
					n.Name,
				)
			}
		}
	}

	// Only one name.
	name := vs.Names[0].Name
	if !ast.IsExported(name) {
		return nil
	}

	if vs.Doc == nil && gd.Doc == nil {
		if genDeclMissingComments[gd] {
			return nil
		}

		block := ""
		if kind == "const" && gd.Lparen.IsValid() {
			block = " (or a comment on this block)"
		}

		genDeclMissingComments[gd] = true

		return fmt.Errorf(
			"exported %s %s should have comment%s or be unexported",
			kind,
			name,
			block,
		)
	}

	// If this GenDecl has parens and a comment, we don't check its comment form.
	if gd.Lparen.IsValid() && gd.Doc != nil {
		return nil
	}

	// The relevant text to check will be on either vs.Doc or gd.Doc.
	// Use vs.Doc preferentially.
	doc := vs.Doc
	if doc == nil {
		doc = gd.Doc
	}
	prefix := name + " "
	if !strings.HasPrefix(doc.Text(), prefix) {
		return fmt.Errorf(
			`comment on exported %s %s should be of the form "%s..."`,
			kind,
			name,
			prefix,
		)
	}

	return nil
}

// receiverType returns the named type of the method receiver, sans "*",
// or "invalid-type" if fn.Recv is ill formed.
func receiverType(fn *ast.FuncDecl) string {
	switch e := fn.Recv.List[0].Type.(type) {
	case *ast.Ident:
		return e.Name
	case *ast.StarExpr:
		if id, ok := e.X.(*ast.Ident); ok {
			return id.Name
		}
	}
	// The parser accepts much more than just the legal forms.
	return "invalid-type"
}
