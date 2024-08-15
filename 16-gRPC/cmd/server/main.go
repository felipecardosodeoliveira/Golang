package main

import (
	"database/sql"
	"net"

	"google.golang.org/grpc/reflection"

	"github.com/felipecardosodeoliveira/Golang/16-gRPC/internal/database"
	"github.com/felipecardosodeoliveira/Golang/16-gRPC/internal/pb"
	"github.com/felipecardosodeoliveira/Golang/16-gRPC/internal/service"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
)

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDB := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDB)
	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
