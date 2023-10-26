package db

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"

	_ "github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const dbDriver = "postgres"

var testQueries *Queries

var testDB *sql.DB

func TestMain(m *testing.M) {

	if err := godotenv.Load("../../.env"); err != nil {
		fmt.Println("No env found!!")
	}
	ip := os.Getenv("IP_PG")
	fmt.Println(ip)
	dbSource := fmt.Sprintf("postgresql://root:secret@%v:5432/simple_bank?sslmode=disable", ip)
	fmt.Println(dbSource)
	//main entry point for all tests
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}
