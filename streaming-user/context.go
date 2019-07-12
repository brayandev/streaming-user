package user

import "context"

type contextKey string

// ContextKey represents transactions id.
const (
	ContextKeyTransactionID contextKey = "transaction-id"
	ContextKeyInternalID    contextKey = "internal-id"
)

// GetTransactionID get transaction id from context.
func GetTransactionID(ctx context.Context) string {
	if val := ctx.Value(ContextKeyTransactionID); val != nil {
		return val.(string)
	}
	return ""
}

// GetInternalID get internal id from context.
func GetInternalID(ctx context.Context) string {
	if val := ctx.Value(ContextKeyInternalID); val != nil {
		return val.(string)
	}
	return ""
}
