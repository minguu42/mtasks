package app

import (
	"context"

	"github.com/minguu42/mtasks/app/ogen"
)

// この2つの変数の値はビルド時に埋め込む
var (
	version  string
	revision string
)

// GetHealth は GET /health に対応するハンドラ関数
func (h *handler) GetHealth(_ context.Context) (ogen.GetHealthRes, error) {
	return &ogen.GetHealthOK{
		Version:  version,
		Revision: revision,
	}, nil
}
