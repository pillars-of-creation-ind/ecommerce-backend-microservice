package main

import (
	"account/account"
	"account/config"
	"account/database"
	pb "account/generated"
	"account/server"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"os"
	"os/exec"
)

func main() {
	productionMode := flag.Bool("production", false, "enable production mode")
	flag.Parse()
	cfgName := func() string {
		if *productionMode {
			return "prod"
		}
		return "dev"
	}()
	cfg, err := config.LoadConfig(cfgName)
	if err != nil {
		log.Fatal(err)
	}
	if !*productionMode {
		if err := os.RemoveAll("./generated"); err != nil {
			log.Fatal(err)
		}
		if err := os.Mkdir("./generated", os.ModePerm); err != nil {
			log.Fatal(err)
		}
		if err := exec.Command("protoc", "--go_out=generated", "--go-grpc_out=generated", "--proto_path=proto", "proto/*.proto").Run(); err != nil {
			log.Fatal(err)
		}
	}
	db := database.Connect(context.Background(), fmt.Sprintf("mongodb://%s:%s", cfg.MongoHostname, cfg.MongoPort))
	defer db.Disconnect()
	accountRepository := account.NewRepository(db.Collection(cfg.DatabaseName, cfg.CollectionItems))
	accountUseCase := account.NewUseCase(accountRepository)
	accountService := account.NewService(accountUseCase, account.NewMapper())
	authInterceptor := server.NewInterceptor("Authorization", func(ctx context.Context, header string) error {
		if header != cfg.ApiKey {
			return status.Errorf(codes.Unauthenticated, "Invalid API key")
		}
		return nil
	})
	server.Server{Address: cfg.ServerAddress}.Launch(func(server *grpc.Server) {
		pb.RegisterAccountServiceServer(server, accountService)
	}, authInterceptor)
}
