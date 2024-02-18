package main

import (
	"fmt"
	dbConfig "github.com/NickNaskida/Gorder/internal/db"
	"github.com/NickNaskida/Gorder/internal/models"
	interfaces "github.com/NickNaskida/Gorder/pkg/v1"
	handler "github.com/NickNaskida/Gorder/pkg/v1/handler/grpc"
	repo "github.com/NickNaskida/Gorder/pkg/v1/reposiroty"
	"github.com/NickNaskida/Gorder/pkg/v1/usecase"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
	"net"
)

func main() {
	// connect to the db
	db := dbConfig.DbConn()
	migrations(db)

	// add a listener address
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("ERROR STARTING THE SERVER : %v", err)
	}

	// start the grpc server
	grpcServer := grpc.NewServer()

	orderUseCase := initOrderServer(db)
	handler.NewServer(grpcServer, orderUseCase)

	// start serving to the address
	log.Fatal(grpcServer.Serve(lis))
}

func initOrderServer(db *gorm.DB) interfaces.UseCaseInterface {
	orderRepo := repo.New(db)
	return usecase.New(orderRepo)
}

func migrations(db *gorm.DB) {
	err := db.AutoMigrate(&models.Order{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Migrated")
	}
}
