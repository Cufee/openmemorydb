package types

type Document map[string]interface{}
type Filter map[string]interface{}

// Base information for all handlers, this should not be used directly
type OperationRequestBase struct {
	Database   string `json:"database"`
	Collection string `json:"collection"`
	Operation  string `json:"operation"` // Not user, wii be used by the handler to check if the opertions match the handler
}

// Reqeust information for create handler
type OperationRequestCreateOne struct {
	OperationRequestBase
	Document Document `json:"document"`
}
type OperationRequestCreateMany struct {
	OperationRequestBase
	Documents []Document `json:"documents"`
}

// Request information for read handlers
type OperationRequestReadOne struct {
	OperationRequestBase
	Filter Filter `json:"filter"`
}

// Request information for read handlers
type OperationRequestReadMany struct {
	OperationRequestBase
	Filter Filter `json:"filter"`
	Limit  int    `json:"limit"`
}

// Request information for update handlers
type UpdateOptions struct {
	Filter Filter   `json:"filter"`
	Update Document `json:"update"`
	Upsert bool     `json:"upsert"`
}
type OperationRequestUpdateOne struct {
	OperationRequestBase
	Update UpdateOptions `json:"updateOption"`
}
type OperationRequestUpdateMany struct {
	OperationRequestBase
	Updates []UpdateOptions `json:"updateOptions"`
}

// Request information for delete handlers
type OperationRequestDeleteOne struct {
	OperationRequestBase
	Filter Filter `json:"filter"`
}
type OperationRequestDeleteMany struct {
	OperationRequestBase
	Filter Filter `json:"filter"`
	Limit  int    `json:"limit"`
}
