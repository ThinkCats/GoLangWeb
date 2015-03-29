package main
import (
    _ "../../github.com/go-sql-driver/mysql"
    "database/sql"
    "log"
)

type user struct {
    name string
    age int
    descrip string
}

func main() {
    db := Conn()
    list := []user{}
    list = Query("select * from user",db)
    for _,u := range list{
        log.Println(u.name)
        log.Println(u.age)
        log.Println(u.descrip)
    }
    log.Println("===== del test ===")
    i,_ := Del("delete from user where name = ?","wang",db)
    log.Println("del complete ",i)

    log.Println("----- begin insert ----")
    //uu :=[]string{"mm","20","niciaa"}
    //j := Save("insert into user values (?,?,?)",db,"ww","12","hii")
    //log.Println("insert complete ",j)
}

func Conn() *sql.DB{
    db,err := sql.Open("mysql","root:12345@tcp(127.0.0.1:3306)/mytest")
    checkErr(err)
    return db
}

func Query(sql string,db *sql.DB) []user{
    u := user{}
    items :=[]user{}
    rows,err := db.Query(sql)
    checkErr(err)
    for rows.Next(){
         rows.Scan(&u.name,&u.age,&u.descrip)
        items = append(items,u)
    }
    return items
}

func Del(sql string,args  interface{},db *sql.DB) (int64,error){
    log.Println("== del begin ==")
    stmt,err := db.Prepare(sql)
    checkErr(err)
    res,err2 := stmt.Exec(args)
    checkErr(err2)
    return res.RowsAffected()
}

func Save(sql string,db *sql.DB,args ...interface{})int64{
    stmt,_ := db.Prepare(sql)
    res,_ := stmt.Exec(args)
    k,_ := res.RowsAffected()
    return k
}

func checkErr(err error){
    //fmt.Println(err)
    if err != nil{
        log.Fatal(err)
    }
}