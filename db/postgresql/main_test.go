package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

const (
	dbSource = "postgresql://root:secret@localhost:5431/simple_bank?sslmode=disable"
)

var testQuery *Queries

func TestMain(m *testing.M) {
	conn, err := pgx.Connect(context.Background(), dbSource)

	if err != nil {
		log.Fatal("cannot connect, ", err)
	}

	testQuery = New(conn)
	os.Exit(m.Run())

}
