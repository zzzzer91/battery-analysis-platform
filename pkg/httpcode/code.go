package httpcode

// informational
const (
	// success
	OK = 200

	// client error
	BadRequest   = 400
	Unauthorized = 401
	Forbidden    = 403
	NotFound     = 404

	// server errors
	InternalServerError = 500
	NotImplemented      = 501
	BadGateway          = 502
	ServiceUnavailable  = 503
)
