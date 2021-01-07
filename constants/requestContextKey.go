package constants

type requestContextKeys struct {
	AuthorizedUser Constant
}

// RequestContextKey is a context item added to a request in router middleware
var RequestContextKey = &requestContextKeys{
	AuthorizedUser: "authorizedUser",
}
