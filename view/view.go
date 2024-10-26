package view

import (
	"context"
)

type Screen interface {
	Render(ctx context.Context)
}
