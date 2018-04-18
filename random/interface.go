package random

import "context"

type Interface interface {
	NewFloat64(ctx context.Context, min float64, max float64) (float64, error)
	NewString(ctx context.Context, num int) (string, error)
}
