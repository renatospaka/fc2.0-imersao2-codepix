package main

import (
	"os"

	"github.com/renatospaka/imersao/codepix-go/application/grpc"
	"github.com/renatospaka/imersao/codepix-go/infrastructure/db"
	"github.com/jinzhu/gorm"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
	//50051 é a porta padrão do gRPC


}