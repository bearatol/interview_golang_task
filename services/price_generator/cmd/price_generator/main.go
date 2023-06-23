package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/bearatol/interview_golang_task/pkg/logger"
	pb "github.com/bearatol/interview_golang_task/proto/price_generator"
	"github.com/bearatol/interview_golang_task/sevices/price_generator/internal/config"
	"github.com/bearatol/interview_golang_task/sevices/price_generator/internal/handler"
	"github.com/bearatol/interview_golang_task/sevices/price_generator/internal/pdf"
	"github.com/bearatol/interview_golang_task/sevices/price_generator/internal/service"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.TODO(), syscall.SIGTERM, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	conf, err := config.NewConfig()
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
	pdfRepo, err := pdf.NewPDF(conf)
	if err != nil {
		return err
	}
	serv := service.NewService(pdfRepo)
	handl := handler.NewHandler(ctx, log, serv)

	lis, err := net.Listen("tcp", conf.AppAddr)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPriceGeneratorServer(s, handl)

	log.Infof("Server listening at %v", lis.Addr())
	return s.Serve(lis)
}
