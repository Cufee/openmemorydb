package types

// Response from all handlers
type OperationResult struct {
	Matched int `json:"matched"`
	Created int `json:"created"`
	Updated int `json:"updated"`
	Deleted int `json:"deleted"`

	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}
