package main

import "./db"
import "database/sql"
import _"github.com/go-sql-driver/mysql"
import "github.com/gin-gonic/gin"
import "strconv"
import "fmt"

type DbWorker struct {
	Dsn string 
}
var DB *sql.DB
func main() {
	DB=db.Connect()
	defer DB.Close()
	r:=gin.Default()
	r.GET("/shortUrl/:shortUrl",shortHandler)
	r.GET("/longUrl/:longUrl",longHandler)
	r.Run(":8000")
}

func shortHandler(c *gin.Context) {
	shortResult:=db.GetShortUrl(DB,c.Param("shortUrl"))
	message:="http://139.196.76.36:8000/"+strconv.FormatInt(shortResult,10)
	c.JSON(200,message)
}

func longHandler(c *gin.Context) {
	longUrl,err:=strconv.ParseInt(c.Param("longUrl"),10,64)
	if err==nil{
		fmt.Print("strconv is failed,err=%v",err)
	}
	longResult:=db.GetLongUrl(DB,longUrl)
	c.JSON(200,longResult)
}

