package main

import "fmt"
import "log"
import "net/http"
import "./db"
import "database/sql"
import _"github.com/go-sql-driver/mysql"

type DbWorker struct {
	Dsn string 
}
var DB *sql.DB
func main() {
	DB=db.Connect()
	http.HandleFunc("/",hanler)
	log.Fatal(http.ListenAndServe("localhost:8000",nil))
}

func hanler(w http.ResponseWriter,r *http.Request) {
	fmt.Fprintf(w,"URL.Path = %q\n",r.URL.Path)
	shortUrl:=db.GetShortUrl(DB,r.URL.Path)
	fmt.Fprintf(w,"shortUrl=http://liu123/%d",shortUrl)
}

