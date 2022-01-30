package request

// GetMessages request object.
type GetMessages struct {
	Offset uint64 `form:"offset,omitempty"`
}
