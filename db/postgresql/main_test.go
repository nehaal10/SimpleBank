package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
	"github.com/nehaal10/simeplebank/util"
)

var testQuery *Queries

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")

	if err != nil {
		log.Fatal("cannot connect, ", err)
	}

	conn, err := pgx.Connect(context.Background(), config.DBSource)

	if err != nil {
		log.Fatal("cannot connect, ", err)
	}

	testQuery = New(conn)
	os.Exit(m.Run())

}
