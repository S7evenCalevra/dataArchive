// note that database tests are a complex layer of abstraction. Good practice is to test this logic against an actual db.
// That being said, the database layer can be pulled via docker and insert logic can be tested. Since this code does not
// create a database as one was already provided, in running database_test.go it is suggested to pull the sqllite docker
// docker image for testing sql insert + select statements
package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/denisenkom/go-mssqldb"
)

var sqlliteserver string

var testServer = "TESTSRV"
var testUser = "TESTUSR"
var testToken = "TESTTKN"
var testPort = 123
var testDatabase = "TESTDB"

// 'new' exported messageformatter interface that can be mocked
type Messageformatter interface {
	InsertToDB(context.Context, *data) error
}

type messageformatt struct {
	db *sql.DB
}

// Database test models
type data struct {
	ban            string
	messageId      string
	messageContext string
	messageTime    string
	from           string
	to             string
	groupMessage   string
	direction      string
	subject        string
	contentType    string
	textContent    string
	textSize       string
	attachment     string
	name           string
	content        string
	attachSize     string
}

//
func ConnectDB(m *testing.M) (e error) {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", testServer, testUser, testToken, testPort, testDatabase)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
		return
	}
	defer db.Close()
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("DB Connected!")
	return

}

func (n *data) InsertToDB(ctx context.Context, data *Messageformatter) error {
	n.ban = "somevalue1"
	n.messageId = "somevalue2"
	n.messageContext = "somvevalue3"
	n.messageTime = "somevalue4"
	n.from = "somevalue5"
	n.to = "sommeone"
	n.groupMessage = "somevalue6"
	n.direction = "somevalue7"
	n.subject = "somevalue8"
	n.contentType = "somevalue9"
	n.textContent = "somevalue10"
	n.textSize = "somevalue11"
	n.attachment = "somevalue12"
	n.name = "somevalue13"
	n.content = "somevalue14"
	n.attachSize = "somevalue15"

	queryString := fmt.Sprintf("INSERT INTO %s (payload,insertTime) VALUES (@p1, @p2)", stageingtable)
	res, err := db.ExecContext(ctx, queryString, n.ban, n.messageId)
	if err != nil {
		return fmt.Errorf("could not insert row: %w", err)
	}

	if _, err := res.RowsAffected(); err != nil {
		return fmt.Errorf("could not get affected rows: %w", err)
	}

	return nil
}
