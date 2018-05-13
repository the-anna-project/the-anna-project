package nodeidcontext

import (
	"context"
)

type key string

var nodeIDKey key = "nodeid"

func NewContext(ctx context.Context, v string) context.Context {
	return context.WithValue(ctx, nodeIDKey, v)
}

func FromContext(ctx context.Context) (string, bool) {
	v, ok := ctx.Value(nodeIDKey).(string)
	return v, ok
}
