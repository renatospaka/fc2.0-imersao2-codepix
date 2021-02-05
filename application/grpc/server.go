package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/jinzhu/gorm"
	"github.com/renatospaka/imersao/codepix-go/application/grpc/pb"
	"github.com/renatospaka/imersao/codepix-go/application/usecase"
	"github.com/renatospaka/imersao/codepix-go/infrastructure/repository"
	"google.golang.org/grpc"

	//"github.com/ktr0731/evans/usecase"
	"google.golang.org/grpc/reflection"
)

func StartGrpcServer(database *gorm.DB, port int) {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	// registrando o serviço no gRPC (muito foda isso, não entendi picas)
	pixRepository := repository.PixKeyRepositoryDb{Db: database}
	pixUseCase := usecase.PixUseCase{PixKeyRepository: pixRepository}
	pixGrpcService := NewPixGrpcService(pixUseCase)
	pb.RegisterPixServiceServer(grpcServer, pixGrpcService)

	address := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("Cannot start gRPC server", err)
	}

	log.Printf("gRPC server has been started on port: %d", port)

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("Cannot start gRPC server", err)
	}
}
