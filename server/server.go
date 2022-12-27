package server

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"math"
	"net"
	"net/http"
	"reservation/config"
	"reservation/github.com/reservation/api"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)
type apiServiceImpl struct {
	api.UnimplementedReservationServiceServer
	config config.Config
}

type Validator interface {
	Validate() error
}
func InitLogger() {
	// 通过 zap.NewProduction() 创建一个 logger
	logger, _ = zap.NewProduction()
}

var logger *zap.Logger
func RunGRPCServer(config config.Config) error {
	InitLogger()
	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if p, ok := req.(Validator); ok {
			if err := p.Validate(); err != nil {
				// 参数没有通过验证时;返回一个错误不继续向下执行
				return nil, status.Error(codes.InvalidArgument, err.Error())
			}
		}
		// 通过验证后继续向下执行
		return handler(ctx, req)
	}

	// Init listener
	grpcPort := config.GRPCPort
	listener, err := net.Listen("tcp", grpcPort)
	if err != nil {
		logger.Error("Run(). Failed to listen " + grpcPort + err.Error())
		return err
	}
	defer func() {
		if err := listener.Close(); err != nil {
			logger.Error("Run(). Failed to close " + grpcPort + err.Error())
		}
	}()
	options := []grpc.ServerOption{
		grpc.MaxRecvMsgSize(math.MaxUint32),
		grpc.UnaryInterceptor(
			interceptor,
		),
	}
	resServer, err := NewReservationServer(config)
	if err != nil {
		logger.Error("Run(). Create res server failed:" + err.Error())
		return err
	}
	s := grpc.NewServer(options...)
	api.RegisterReservationServiceServer(s, resServer)
	// Run grpc server
	if err := s.Serve(listener); err != nil {
		log.Fatal("Run(). Failed to serve: " + err.Error())
		return err
	}
	return nil
}

func RunHTTPServer(config *config.Config) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	conn, err := grpc.DialContext(
		ctx,
		"localhost"+config.GRPCPort,
		grpc.WithBlock(),
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(math.MaxUint32)),
	)
	if err != nil {
		log.Fatal("start http server failed" + err.Error())
	}
	gateway := runtime.NewServeMux()
	err = api.RegisterReservationServiceHandler(context.Background(), gateway, conn)
	if err != nil {
		log.Fatal("start http server failed, " + err.Error())
	}
	gwServer := &http.Server{
		Addr:    config.HTTPPort,
		Handler: gateway,
	}
	err = gwServer.ListenAndServe()
	if err != nil {
		log.Fatal("RunHttpServer().start HTTP server failed. " + err.Error())
		return err
	}
	return nil
}


// NewReservationServer creates an instance of DataManagerServiceServer
func NewReservationServer(config config.Config) (api.ReservationServiceServer, error) {
	svc := &apiServiceImpl{
		config: config,
	}
	return svc, nil
}