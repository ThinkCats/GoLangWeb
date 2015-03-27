package db

import (
    _ "../../github.com/go-sql-driver/mysql"
    "database/sql"
    "go/ast"
    "fmt"
)

func Conn() sql.DB{
    db,err := sql.Open("mysql","root:12345@tcp(127.0.0.1:3306)/mytest")
    checkErr(err)
    return db
}

func Del(sql string,id int,db sql.DB){
    stmt,err := db.Prepare(sql)
    checkErr(err)
    res,err2 := stmt.Exec(id)
    fmt.Println(res)
    checkErr(err2)
}

func checkErr(err error){
    fmt.Println(err)
    if err != nil{
        log.Fatal(err)
    }
}