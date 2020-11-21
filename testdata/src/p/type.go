package p

// No assertions are done here. You are free to comment whatever.
type typeNotExported string

// TypeCommented this is a valid type comment
type TypeCommented string

// This comment does not comply with the format
type TypeWronglyCommented string // want "comment on exported type TypeWronglyCommented should be of the form "

type TypeUncommented string // want "exported type TypeUncommented should have comment or be unexported"

type (
	// TypeBlockCommented is correctly commented
	TypeBlockCommented string
	// This comment does not comply with the format
	TypeBlockWronglyCommented string // want "comment on exported type TypeBlockWronglyCommented should be of the form "
	TypeBlockUncommented      string // want "exported type TypeBlockUncommented should have comment or be unexported"
)
