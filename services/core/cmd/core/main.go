package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bearatol/interview_golang_task/pkg/logger"
	"github.com/bearatol/interview_golang_task/sevices/core/internal/config"
	"github.com/bearatol/interview_golang_task/sevices/core/internal/grpc_conn/price_generator"
	"github.com/bearatol/interview_golang_task/sevices/core/internal/handler"
	"github.com/bearatol/interview_golang_task/sevices/core/internal/repository"
	"github.com/bearatol/interview_golang_task/sevices/core/internal/repository/postgres"
	"github.com/bearatol/interview_golang_task/sevices/core/internal/service"
	"go.uber.org/zap"
)

const (
	Local      = "--local"
	LocalShort = "-l"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.TODO(), syscall.SIGTERM, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	var local bool
	for _, arg := range os.Args {
		if arg == Local || arg == LocalShort {
			local = true
		}
	}

	conf, err := config.NewConfig(local)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	zapLogger, err := logger.NewLogger()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	log := logger.NewSugarLogger(zapLogger)

	log.Error(run(ctx, conf, log))

	<-ctx.Done()
}

func run(ctx context.Context, conf *config.Config, log *zap.SugaredLogger) error {
	db, err := postgres.NewPGConn(conf.DB)
	if err != nil {
		return err
	}

	repo := repository.NewRepository(db)

	priceGen, err := price_generator.NewConn(conf.PriceGenAdd)
	if err != nil {
		return err
	}
	defer priceGen.Conn.Close()

	serv := service.NewService(repo, priceGen)
	handl := handler.NewHandler(ctx, log, serv)
	router := handl.SetupRouter()

	log.Infof("Starting server on %s", conf.AppAddr)
	return router.Run(conf.AppAddr)
}
