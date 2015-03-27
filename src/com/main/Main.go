package main
import (
    "net/http"
    "fmt"
    "strings"
    "log"
    "html/template"
)

func main() {
    http.HandleFunc("/",index)
    http.HandleFunc("/login",login)
    err := http.ListenAndServe(":9090",nil)
    if err != nil {
        log.Fatal("listen and serve ",err)
    }
}

func index(w http.ResponseWriter,r *http.Request){
    r.ParseForm()
    fmt.Println(r.Form)
    fmt.Println(r.URL.Path)
    fmt.Println(r.Form["param"])
    for k,v := range r.Form{
        fmt.Println("key:",k)
        fmt.Println("value:",strings.Join(v,""))
    }
    fmt.Fprintf(w,"hello world")
}

func login(w http.ResponseWriter,r *http.Request){
    fmt.Println("method:",r.Method)
    if r.Method == "GET"{
        t,_ := template.ParseFiles("/home/wang/IdeaProject/GoLangWeb/view/login.gtpl")
        t.Execute(w,nil)
    }else {
        r.ParseForm()
        fmt.Println("username:",r.Form["username"])
        fmt.Println("password:",r.Form["password"])
    }
}
