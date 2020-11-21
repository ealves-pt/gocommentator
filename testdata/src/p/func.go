package p

// Because the func is not exported any comment is valid
func funcUnexported() {}

// FuncCommented is a valid commented function
func FuncCommented() {}

// This comment does not comply with format
func FuncWronglyCommented() {} // want "comment on exported function FuncWronglyCommented should be of the form "

func FuncUncommented() {} // want "exported function FuncUncommented should have comment or be unexported"
