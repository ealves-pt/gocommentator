package p

// ExportedMethodCommented is a valid comment on an exported method for an exported type
func (TypeStructCommented) ExportedMethodCommented() {}

// This comment does not comply with the format
func (TypeStructCommented) ExportedMethodWronglyCommented() {} // want "comment on exported method TypeStructCommented.ExportedMethodWronglyCommented should be of the form "

func (TypeStructCommented) ExportedMethodUncommented() {} // want "exported method TypeStructCommented.ExportedMethodUncommented should have comment or be unexported"

// Because the method is not exported any comment is valid
func (TypeStructCommented) unexportedMethod() {}

// Because the receiver type is not exported any comment is valid
func (typeStructNotExported) ExportedMethodCommented() {}

// Because the receiver type is not exported any comment is valid
func (typeStructNotExported) ExportedMethodWronglyCommented() {}

// Because the receiver type is not exported any comment is valid
func (typeStructNotExported) ExportedMethodUncommented() {}

// Because the receiver type is not exported any comment is valid
func (typeStructNotExported) unexportedMethod() {}
