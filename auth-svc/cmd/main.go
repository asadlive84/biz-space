package main

import (
	"fmt"
	"log"
	"net"

	"github.com/asadlive84/bizspace/auth-svc/internal/config"
	"github.com/asadlive84/bizspace/auth-svc/internal/db"
	q "github.com/asadlive84/bizspace/auth-svc/internal/query"
	"github.com/asadlive84/bizspace/auth-svc/services"
	pb "github.com/asadlive84/bizspace/proto/auth/pb"
	"github.com/jmoiron/sqlx"

	"google.golang.org/grpc"
)

func main() {
	configDir := "../../auth-svc/internal/config"
	configMigrationsDir := "../../auth-svc/internal/migrations"
	c, err := config.LoadConfig(configDir)

	if err != nil {
		log.Default().Printf("error is %+v", err)
		log.Fatalln("Failed to getting config", err)
	}

	dbConfig, err := db.NewDatabaseConfig(db.DBConfig{
		POSTGRES_USER:     c.POSTGRES_USER,
		POSTGRES_PASSWORD: c.POSTGRES_PASSWORD,
		POSTGRES_DB:       c.POSTGRES_DB,
		POSTGRES_PORT:     c.POSTGRES_PORT,
		POSTGRES_HOST:     c.POSTGRES_HOST,
		PORT:              c.PORT,
		JWT_SECRET_KEY:    c.JWT_SECRET_KEY,
	})

	if err != nil {
		log.Fatalln("Failed to database NewDatabaseConfig:", err)
	}

	_, err = db.DbInit(dbConfig, configMigrationsDir)

	if err != nil {
		log.Fatalln("Failed to database migration:", err)
	}

	hdb, err := sqlx.Connect("postgres", dbConfig)

	if err != nil {
		log.Fatalln("Failed to database sqlx open up:", err)
	}

	lis, err := net.Listen("tcp", c.PORT)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Auth Svc on", c.PORT)

	queryDb := q.QueryInit{
		Db: hdb,
	}

	s := services.Server{
		// H: h,
		Q: &queryDb,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, &s)

	fmt.Println("=============>Auth Server is running<===============")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}

}
