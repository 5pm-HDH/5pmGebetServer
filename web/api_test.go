package web

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"log"
	"strings"
	"testing"
	"time"
)

var testDatabase *sql.DB

const keyTestMaster = "test_master_key"
const keyTestView = "test_view_key"
const keyTestPray = "test_pray_key"

func TestMain(m *testing.M) {
	setup()
	m.Run()
	teardown()
}

func setup() {
	db, _ := sql.Open("sqlite3", "file::memory:?cache=shared")
	file, _ := ioutil.ReadFile("../test/ddl_test.sql")

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		_, err := db.Exec(request)
		if err != nil {
			log.Print(err)
		}
	}

	testDatabase = db

	SetDatabase(db)
	p := Prayer{
		Id:       "9997",
		Content:  "Test_Prayer",
		Public:   true,
		Approved: false,
		Date:     time.Now(),
	}
	_ = DatabaseInsertPrayer(p)

	p.Id = "9998"
	p.Approved = true
	_ = DatabaseInsertPrayer(p)

	p.Id = "9999"
	p.Public = false
	_ = DatabaseInsertPrayer(p)

	fmt.Printf("\033[1;36m%s\033[0m", "> Setup completed\n")
}

func teardown() {
	// Do something here.

	fmt.Printf("\033[1;36m%s\033[0m", "> Teardown completed")
	fmt.Printf("\n")
}
