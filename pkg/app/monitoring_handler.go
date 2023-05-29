package app

import (
	"context"

	"github.com/minguu42/mtasks/pkg/ogen"
)

// GetHealth は GET /health に対応するハンドラ関数
func (h *Handler) GetHealth(_ context.Context) (ogen.GetHealthRes, error) {
	return &ogen.GetHealthOK{}, nil
}
