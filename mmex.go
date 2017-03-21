package main

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "github.com/ant0ine/go-json-rest/rest"
    "log"
    "net/http"
)

func GetAccount(w rest.ResponseWriter, req *rest.Request) {
    user_id := req.PathParam("user_id")

    db, err := sql.Open("sqlite3", user_id + ".db")
    if err != nil {
        log.Fatal(err)
    }

    defer db.Close()
}

func main () {

    api := rest.NewApi()
    api.Use(rest.DefaultDevStack...)

    router, err := rest.MakeRouter(
        rest.Get("/account/:id", GetAccount),
    )

    if err != nil {
        log.Fatal(err)
    }

    api.SetApp(router)
    log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}
