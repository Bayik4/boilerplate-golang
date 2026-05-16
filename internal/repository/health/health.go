package health

import "context"

type HealthRepository interface {
	Get(ctx context.Context) (string, error)
}