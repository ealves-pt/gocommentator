package p

// TypeCommented this is a valid type comment
type TypeCommented string

type TypeUncommented string // want "exported type TypeUncommented should have comment or be unexported"

// This comment does not comply with the format
type TypeWronglyCommented string // want "comment on exported type TypeWronglyCommented should be of the form "

// Any comment here is valid because it's not exported
type typeNotExported string

type (
	// TypeBlockCommented is correctly commented
	TypeBlockCommented   string
	TypeBlockUncommented string // want "exported type TypeBlockUncommented should have comment or be unexported"
	// This comment does not comply with the format
	TypeBlockWronglyCommented string // want "comment on exported type TypeBlockWronglyCommented should be of the form "
	// Any comment here is valid because it's not exported
	typeBlockNotExported string
)

// TypeStructCommented this is a valid type comment
type TypeStructCommented struct{}

type TypeStructUncommented struct{} // want "exported type TypeStructUncommented should have comment or be unexported"

// This comment does not comply with the format
type TypeScructWronglyCommented struct{} // want "comment on exported type TypeScructWronglyCommented should be of the form "

// Any comment here is valid because it's not exported
type typeStructNotExported struct{}
