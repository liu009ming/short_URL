package main

import "./db"
import "database/sql"
import _"github.com/go-sql-driver/mysql"
import "github.com/gin-gonic/gin"
import "./bitAlgorithm"
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
	r.GET("/index/:longUrl",indexOfLongUrl)
	r.Run(":8000")
}

func shortHandler(c *gin.Context) {
	shortResult:=db.GetShortUrl(DB,c.Param("shortUrl"))
	message:="http://127.0.0.1:8000/index/"+bitAlgorithm.IntToString(uint64(shortResult))
	c.JSON(200,message)
}

func longHandler(c *gin.Context) {
	longUrl:=bitAlgorithm.StringToInt(c.Param("longUrl"))
	longResult:=db.GetLongUrl(DB,int64(longUrl))
	c.JSON(200,longResult)
}

func indexOfLongUrl(c *gin.Context) {
	longUrl:=bitAlgorithm.StringToInt(c.Param("longUrl"))
	longResult:=db.GetLongUrl(DB,int64(longUrl))
	c.Redirect(301,"http://"+longResult)
}

