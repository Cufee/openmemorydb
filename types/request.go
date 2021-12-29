package types

type Document map[string]interface{}

func (o Document) IsValid() bool {
	return o != nil
}

type Filter map[string]interface{}

// Base information for all handlers, this should not be used directly
type OperationRequestBase struct {
	Database   string `json:"database"`
	Collection string `json:"collection"`
	Operation  string `json:"operation"` // Used to validate the request
}

// Reqeust information for create handler
const CreateOperationName = "create"

type OperationRequestCreateOne struct {
	OperationRequestBase
	Document Document `json:"document"`
}
type OperationRequestCreateMany struct {
	OperationRequestBase
	Documents []Document `json:"documents"`
}

// Request information for read handlers
const ReadOperationName = "read"

type OperationRequestReadOne struct {
	OperationRequestBase
	Filter Filter `json:"filter"`
}
type OperationRequestReadMany struct {
	OperationRequestBase
	Filter Filter `json:"filter"`
	Limit  int    `json:"limit"`
}

// Request information for update handlers
const UpdateOperationName = "update"

type UpdateOptions struct {
	Filter Filter   `json:"filter"`
	Update Document `json:"update"`
	Upsert bool     `json:"upsert"`
}

func (o *UpdateOptions) IsValid() bool {
	return o.Filter != nil && o.Update.IsValid()
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
const DeleteOperationName = "delete"

type DeleteOptions struct {
	Filter    Filter `json:"filter"`
	AcceptNil bool   `json:"acceptNil"`
}

func (o DeleteOptions) IsValid() bool {
	return !(o.Filter == nil && !o.AcceptNil)
}

type OperationRequestDeleteOne struct {
	DeleteOptions
	OperationRequestBase
}
type OperationRequestDeleteMany struct {
	DeleteOptions
	OperationRequestBase
	Limit int `json:"limit"`
}
