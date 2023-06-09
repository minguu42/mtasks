package handler

import (
	"context"

	"github.com/minguu42/mtasks/gen/ogen"
)

// version と revision の値はビルド時に埋め込まれる
var (
	version  string
	revision string
)

// GetHealth は GET /health に対応するハンドラ関数
func (h *Handler) GetHealth(_ context.Context) (*ogen.GetHealthOK, error) {
	return &ogen.GetHealthOK{
		Version:  version,
		Revision: revision,
	}, nil
}
