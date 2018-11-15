package helper

import "fmt"

// Custom type error
const (
	InvalidRequestError   = "invalid_request_error"
	InvitalidAuth         = "invalid_auth"
	UnknownResource       = "unknown_resource"
	AuthorizationRequired = "authorization_required"
	ConflictResource      = "conflict_resource"
	APIError              = "api_error"
)

// Error model
type Error struct {
	Throwable bool                `json:"-"`
	Message   string              `json:"message"`
	Type      string              `json:"type"`
	Fields    map[string][]string `json:"fields,omitempty"`
}

// ErrorsHTTP is an array of status code
var ErrorsHTTP = map[string]int{
	InvalidRequestError:   400,
	InvitalidAuth:         401,
	AuthorizationRequired: 403,
	UnknownResource:       404,
	ConflictResource:      409,
	APIError:              500,
}

// BuildErrorMessage build the message for http errors
func BuildErrorMessage(code int, resource string) string {
	switch code {
	case 400:
		return fmt.Sprint("Validation Error")
	case 401:
		return fmt.Sprint("You do not have enough rights")
	case 403:
		return fmt.Sprint("Some parameters are invalid")
	case 404:
		return fmt.Sprintf("%s not found", resource)
	case 409:
		return fmt.Sprintf("%v already exist", resource)
	default:
		return fmt.Sprint("unknown error, please contact support")
	}
}
