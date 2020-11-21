package p

const constNotExported = ""

const (
	constBlockNotExported = ""
)

// ConstCommented this is a correctly commented const
const ConstCommented = ""

// This comment does not not comply with the format
const ConstWronglyCommented = "" // want "comment on exported const ConstWronglyCommented should be of the form "

const ConstUncommented = "" // want "exported const ConstUncommented should have comment or be unexported"

const (
	// ConstBlockCommented is correctly commented
	ConstBlockCommented = ""
	// This comment does not comply with the format
	ConstBlockWronglyCommented = "" // want "comment on exported const ConstBlockWronglyCommented should be of the form "
	ConstBlockUncommented      = "" // want "exported const ConstBlockUncommented should have comment "
)

// A block comment applies to all the items inside.
// These items should normally be grouped by their relationship
const (
	ConstTopBlockCommented = ""
)
