package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
	"github.com/nehaal10/simeplebank/api"
	db "github.com/nehaal10/simeplebank/db/postgresql"
	"github.com/nehaal10/simeplebank/util"
)

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("cannot connect, ", err)
	}

	conn, err := pgx.Connect(context.Background(), config.DBSource)

	if err != nil {
		log.Fatal("cannot connect, ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server")
	}
}
