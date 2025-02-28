//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
)

func InitializeApp() (*App, error) {
	wire.Build(SuperSet)
	return &App{}, nil
}
