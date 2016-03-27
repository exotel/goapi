package resources

//IValidator interface declares the functions that evry self validating entities should define
type IValidator interface {
	Valid() bool
}

//IResource interface declares the functions that a valid resource should implement
type IResource interface {
}

//Capabilities have the  phone capabilities
type Capabilities struct {
	SMS   bool
	Voice bool
	MMS   bool
}
