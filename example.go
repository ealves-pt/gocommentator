package main

var UncommentedVar string

// CommentedVar is correctly commented
var CommentedVar string

// Some var
var WronglyCommentedVar string

// TrailedVarCommented should be commented
var TrailedVarCommented, anotherVar, andAnotherOne string

// AnotherTrailedVarCommented should be commented
var AnotherTrailedVarCommented, ButThisOneShouldNotBeHere, andThisOneIsFine string

var (
	// BlockVarCommented should be commented
	BlockVarCommented   string
	BlockVarUncommented string
	// wrongly commented
	BlockVarWronglyCommented string
)

type ignoredUncommented string

type (
	// Test commented
	Test        string
	AnotherTest string
)

type Uncommented string

// some comment
type WronglyCommented string

// Commented is type correctly commented
type Commented string

// The ArticleCommented is type correctly commented
type ArticleCommented string

type X struct{}

// MethodCommented exports some method
func (X) MethodCommented() {}

func (X) MethodUncommented() {}

// Exports some method
func (X) MethodWronglyCommented() {}

// ExportedCommented exports a function that is commented
func ExportedCommented() {
}

// exports a function that is commented
func ExportedWronglyCommented() {}

func ExportedUncommented() {
}
