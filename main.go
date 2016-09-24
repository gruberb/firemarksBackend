package main

import(
    "fmt"
    "log"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "net/http"
    "github.com/labstack/echo"
    "github.com/labstack/echo/engine/standard"
)

type Link struct {
    Url string
    Name string
}

func main() {

    // Write test data to the MongoDB
    session, err := mgo.Dial("localhost/firemarks")
    if err != nil {
        panic(err)
    }
    defer session.Close()

    // Optional. Switch the session to a monotonic behavior.
    session.SetMode(mgo.Monotonic, true)

    c := session.DB("firemarks").C("links")
    err = c.Insert(&Link{"https://github.com", "GitHub"})
    if err != nil {
        log.Fatal(err)
    }

    result := Link{}
    err = c.Find(bson.M{"name": "GitHub"}).One(&result)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("URL:", result.Url)

    // Start the web server
    e := echo.New()
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, fmt.Sprintf("%v", result))
    })
    e.Run(standard.New(":3000"))
}
