package port

// Deps is everything the user use cases need from outside the core.
// Each use case takes only the fields it uses.
type Deps struct {
	Users Repository
}
