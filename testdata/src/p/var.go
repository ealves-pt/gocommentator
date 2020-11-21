package p

var varNotExported string

var (
	varBlockNotExported string
)

// VarInlineCommented this is a valid comment because all the remaining are not exported
var VarInlineCommented, someVar, someOtherVar string

var VarInlineCommented1, VarInlineCommented2, yetAnotherVar string // want "exported var VarInlineCommented2 should have its own declaration"

// VarCommented this is a correctly commented var
var VarCommented string

// This comment does not not comply with the format
var VarWronglyCommented string // want "comment on exported var VarWronglyCommented should be of the form "

var VarUncommented string // want "exported var VarUncommented should have comment or be unexported"

var (
	// VarBlockCommented is correctly commented
	VarBlockCommented string
	// This comment does not comply with the format
	VarBlockWronglyCommented string // want "comment on exported var VarBlockWronglyCommented should be of the form "
	VarBlockUncommented      string // want "exported var VarBlockUncommented should have comment or be unexported"
)

// A block comment applies to all the items inside.
// These items should normally be grouped by their relationship
var (
	VarTopBlockCommented string
)
