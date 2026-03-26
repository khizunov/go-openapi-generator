package codegen

// ParameterIn describes where a parameter lives in the HTTP request.
type ParameterIn string

const (
	ParameterInPath   ParameterIn = "path"
	ParameterInQuery  ParameterIn = "query"
	ParameterInHeader ParameterIn = "header"
	ParameterInCookie ParameterIn = "cookie"
)

// Parameter represents a single operation parameter or request body.
type Parameter struct {
	Name         string
	OriginalName string
	In           ParameterIn
	Type         string
	OpenAPIType  string
	Format       string
	Required     bool
	Description  string
	IsBody       bool
}

// Response represents one response variant of an operation.
type Response struct {
	StatusCode  string
	Type        string
	Description string
	IsArray     bool
}

// Operation represents a single API endpoint method.
type Operation struct {
	Name            string
	OriginalID      string
	Method          string
	Path            string
	PathFmt         string
	PathParamNames  []string
	Tags            []string
	Summary         string
	Description     string
	Parameters      []Parameter
	RequestBody     *Parameter
	Responses       []Response
	SuccessResponse *Response
}
