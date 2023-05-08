package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
)

var db *sql.DB
var server = os.Getenv("SQL_Hostname")
var port = 1433
var user = os.Getenv("SQL_Username")
var token = os.Getenv("SQL_Token")
var database = os.Getenv("SQL_Database")
var stageingtable = os.Getenv("StageingTable")

var Error = log.New(os.Stdout, "\u001b[31mERROR: \u001b[0m", log.LstdFlags|log.Lshortfile)
var Warning = log.New(os.Stdout, "\u001b[33mWARNING: \u001B[0m", log.LstdFlags|log.Lshortfile)
var Debug = log.New(os.Stdout, "\u001b[36mDEBUG: \u001B[0m", log.LstdFlags|log.Lshortfile)
var Info = log.New(os.Stdout, "\u001b[34mINFO: \u001B[0m", log.LstdFlags|log.Lshortfile)

func ConnectToDatabase() {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", server, user, token, port, database)
	var err error
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		Error.Println("Error creating connection pool to database: ", err.Error())
		return
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		Error.Println(err.Error())
	}

}

// need to have database handler to get select * from table
func GetTableDataSchema(w http.ResponseWriter, r *http.Request) {
	queryString, err := db.Prepare("Select * FROM" + stageingtable)
	if err != nil {
		fmt.Errorf("Please check query and try again")
		return
	}

	defer queryString.Close()

	rows, err := queryString.Query()
	if err != nil {
		fmt.Errorf("Error querying database")
	}

	//did a little something called "ZeNate"
	var arrStr string
	var arr []string

	for rows.Next() {
		rows.Scan(&arrStr)
		json.Unmarshal([]byte(arrStr), &arr)
		//err = rows.Scan(&column1, &column2, &column3)
		if err != nil {
			fmt.Errorf("errors detecting data structure")
			return
		}
	}
	response, err := json.Marshal(arr)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// need func to query returend/selected table and paginate all data to render on frontend

// needs refactoring. not needed in this app
func InsertToDB(d []byte) {
	now := time.Now().UTC().Format(time.RFC3339)
	bodyText := string(d)
	queryString := fmt.Sprintf("INSERT INTO %s (payload,insertTime) VALUES (@p1, @p2)", stageingtable)
	result, err := db.Exec(queryString, bodyText, now)
	if err != nil {
		Error.Println("Error inserting to database: ", err.Error())
		wr := http.ResponseWriter(nil)
		wr.WriteHeader(http.StatusInternalServerError)
		return
	}
	Info.Println(result)
}

func CloseDatabase() {
	defer db.Close()
	Info.Println("DB Connection Closed")
}
