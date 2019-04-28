package db

import (
	"time"
	"fmt"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

type DbWorker struct{
	Dsn string
}

func Connect()(*sql.DB) {
	dw := DbWorker{
		Dsn : "root:123456@tcp(127.0.0.1:3306)/short_url",
	} 
	db,err:=sql.Open("mysql",dw.Dsn)
	if err!=nil{
		fmt.Printf("Open mysql failed ,err %v\n",err)
		return nil
	}
	db.SetConnMaxLifetime(100*time.Second)
	db.SetMaxOpenConns(100)
	return db
}

func QueryByShortUrl(db *sql.DB,index int64 ) (string){
	longUrl := ""
	row:=db.QueryRow("select longUrl from short_url where shortUrl=?",index)
	err := row.Scan(&longUrl)
	if err!=nil{
		fmt.Printf("query failed,err:%v",err)
		return "" 
	}
	return longUrl
}

func QueryByLongUrl(db *sql.DB,longUrl string)(int64) {
	var shortUrl int64 = -1
	row:=db.QueryRow("select shortUrl from short_url where longUrl=?",longUrl)
	err:=row.Scan(&shortUrl)
	if err!=nil{
		fmt.Printf("query failed,err=%v",err)
		return -1
	}
	return shortUrl
}

func Insert(db *sql.DB,longUrl string)(int64) {
	result,err:=db.Exec("insert into short_url(longUrl) values(?)",longUrl)
	if err!=nil{
		fmt.Printf("insert failed,err%v",err)
		return -1
	}
	lastInsertId,err:=result.LastInsertId()
	if err!=nil{
		fmt.Printf("get lastId faild,err=%v",err)
		return -1
	}
	return lastInsertId
}

func GetShortUrl(db *sql.DB,longUrl string)(int64) {
	shortUrl:=QueryByLongUrl(db,longUrl)
	if shortUrl==-1{
		shortUrl=Insert(db,longUrl)
	}
	return shortUrl
}