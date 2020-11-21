package analyzer_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ealves-pt/gocommentator/analyzer"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAll(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get wd: %s", err)
	}

	testdata := filepath.Join(filepath.Dir(wd), "testdata")
	analysistest.Run(t, testdata, analyzer.Analyzer, "p")
}

// import (
// 	"go/ast"
// 	"go/parser"
// 	"go/token"
// 	"go/types"
// 	"testing"
// )

// func TestExportedType(t *testing.T) {
// 	tests := []struct {
// 		typString string
// 		exp       bool
// 	}{
// 		{"int", true},
// 		{"string", false},
// 		{"T", true},
// 		{"t", false},
// 		{"*T", true},
// 		{"*t", false},
// 		{"map[int]complex128", true},
// 	}
//
// 	for _, test := range tests {
// 		src := `package foo; type T int; type t int; type string struct{}`
// 		fset := token.NewFileSet()
// 		file, err := parser.ParseFile(fset, "foo.go", src, 0)
// 		if err != nil {
// 			t.Fatalf("Parsing %q: %v", src, err)
// 		}
//
// 		config := &types.Config{}
// 		pkg, err := config.Check(file.Name.Name, fset, []*ast.File{file}, nil)
// 		if err != nil {
// 			t.Fatalf("Type checking %q: %v", src, err)
// 		}
// 	}
// }
