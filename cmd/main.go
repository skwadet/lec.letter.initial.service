package main

import (
	"context"
	"fmt"
	"github.com/ecosafety/lec.letter.initial.service/internal/gateway/document_generator"

	"github.com/ecosafety/lec.letter.initial.service/configs"
	"github.com/ecosafety/lec.letter.initial.service/internal/app"

	pb "github.com/ecosafety/lec.letter.initial.service/internal/pb/letter_initial_api"

	"github.com/ecosafety/lec.letter.initial.service/internal/service"

	"github.com/ecosafety/lec.letter.initial.service/internal/storage/memory/letter_memory"
	"github.com/ecosafety/lec.letter.initial.service/internal/storage/memory/objective_memory"
	"github.com/ecosafety/lec.letter.initial.service/internal/storage/memory/purpose_memory"

	"github.com/ecosafety/lec.letter.initial.service/internal/storage/pg_letter"
	"github.com/ecosafety/lec.letter.initial.service/internal/storage/pg_objective"
	"github.com/ecosafety/lec.letter.initial.service/internal/storage/pg_purpose"

	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// Create listener
	listener, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatalln("can't create the listener due to:", err)
	}
	log.Println("created listener on port :80")

	ctx := context.Background()

	// Init grpc server
	grpcServer := grpc.NewServer()
	c, configErr := configs.LoadConfig()

	if configErr != nil {
		log.Fatalln("cannot read config due to:", configErr)
	}
	log.Println("loaded config")

	// DB Init
	log.Println("init database")
	dbPool, dbErr := pgxpool.Connect(
		ctx,
		fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", c.DbUser, c.DbPassword, c.DbHost, c.DbPort, c.DbName),
	)
	if dbErr != nil {
		log.Fatalln("cannot connect to db due to:", dbErr)
	}

	// Init storages
	pgLetterStorage := pg_letter.NewStorage(dbPool)
	pgObjectiveStorage := pg_objective.NewStorage(dbPool)
	pgPurposeStorage := pg_purpose.NewStorage(dbPool)

	generatorC, connErr := document_generator.New(ctx, c.GeneratorHost)
	if connErr != nil {
		log.Fatalln("can't create the connect with module service due to:", connErr.Error())
	}

	// Init Service
	newService := service.NewService(
		ctx,
		pgLetterStorage,
		letter_memory.NewCache(),
		pgObjectiveStorage,
		objective_memory.NewCache(),
		pgPurposeStorage,
		purpose_memory.NewCache(),
		generatorC,
	)

	// Init Application
	newApp := app.NewLecLetterInitialService(newService)

	// Register Service
	pb.RegisterLetterInitialServiceServer(grpcServer, newApp)

	// Start serving
	log.Println("starting server")
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalln("can't create the server", err)
	}
}
