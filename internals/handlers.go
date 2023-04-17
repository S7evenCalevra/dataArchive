package internals

import (
	"MessageArchive_P2/db"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type TestMesg struct {
	msisdn      string
	textContext string
}

// ServeHTTP will read the payload that is sent, while confirming if it is a test message or real payload
func (app *application) handler1(rw http.ResponseWriter, r *http.Request) {
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Issue with payload, Please check json format ", http.StatusBadRequest)
		return
	}
	if r.Body != nil {
		defer r.Body.Close()
	}
	f := TestMesg{
		msisdn:      "+99999999999",
		textContext: "This is a test message",
	}
	a := f.msisdn
	b := f.textContext

	switch {

	case r.Method != "POST":
		fmt.Println("A unsupported request was sent")
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	// test case validation from At&t
	case (strings.Contains(string(d), a) && strings.Contains(string(d), b)):
		rw.WriteHeader(http.StatusOK)

	default:
		db.ConnectToDatabase()
		db.InsertToDB(d)
		db.CloseDatabase()
	}
	fmt.Fprint(rw, http.StatusOK)
}
