package resources

//IValidator interface declares the fucntions that a valid resource should implement
type IValidator interface {
	Valid() bool
}

//IResource interface declares the fucntions that a valid resource should implement
type IResource interface {
}

//The phone capabilities
type Capabilities struct {
	SMS   bool
	Voice bool
	MMS   bool
}
