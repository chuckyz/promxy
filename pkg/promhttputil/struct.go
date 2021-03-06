package promhttputil

// These really should be exported from the prom packages.
// This was attempted once already -- https://github.com/prometheus/prometheus/pull/3615
type Status string

const (
	StatusSuccess Status = "success"
	StatusError   Status = "error"
)

type ErrorType string

const (
	ErrorNone     ErrorType = ""
	ErrorTimeout  ErrorType = "timeout"
	ErrorCanceled ErrorType = "canceled"
	ErrorExec     ErrorType = "execution"
	ErrorBadData  ErrorType = "bad_data"
	ErrorInternal ErrorType = "internal"
)
