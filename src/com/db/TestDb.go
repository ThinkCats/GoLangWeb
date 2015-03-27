package main
import (
    _ "../../github.com/go-sql-driver/mysql"
    "database/sql"
    "fmt"
    "log"
)

type user struct {
    name string
    age int
    descrip string
}

func main() {
    db,err := sql.Open("mysql","root:12345@tcp(127.0.0.1:3306)/mytest")
    checkErr(err)
    fmt.Println("=== connected to mysql ===")
    rows,err := db.Query("select * from user")
    checkErr(err)
    ppp := user{}
    for rows.Next(){
        err = rows.Scan(&ppp.name,&ppp.age,&ppp.descrip)
        //checkErr(err)
    }
    fmt.Println("name:",ppp.name)
    fmt.Println("age:",ppp.age)
    fmt.Println("descrip:",ppp.descrip)
}

func checkErr(err error){
    fmt.Println(err)
    if err != nil{
        log.Fatal(err)
    }
}