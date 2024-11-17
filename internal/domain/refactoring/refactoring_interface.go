package refactoring

import (
	"context"
)

type Refactoring interface {
	Refactor(ctx context.Context, src []byte, filename string) (string, error)
}
