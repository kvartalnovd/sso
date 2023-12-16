package app

import (
	"log/slog"
	"time"

	grpcapp "github.com/kvartalnovd/sso/internal/app/grpc"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {
	// TODO: инициализировать хранилище (storage)

	// TODO: Init auth service (auth)

	grpcApp := grpcapp.New(log, grpcPort)
	return &App{
		GRPCSrv: grpcApp,
	}
}
