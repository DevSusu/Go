# Developing a RESTful API with Go and Gin

## Start

```
> mkdir web-service-gin && cd web-service-gin
> go mod init github.com/DevSusu/Go/golang.org/web-service-gin
go: creating new go.mod: module github.com/DevSusu/Go/golang.org/web-service-gin

> vim main.go
```

```go
package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

type album struct {
    ID      string  `json:"id"`
    Title   string  `json:"title"`
    Artist  string  `json:"artist"`
    Price   float64 `json:"price"`
}

var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumByID(c *gin.Context) {
    id := c.Param("id")

    for _, al := range albums {
        if al.ID == id {
            c.IndentedJSON(http.StatusOK, al)
            return
        }
    }

    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func postAlbums(c *gin.Context) {
    var newAlbum album

    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)
    router.GET("/albums/:id", getAlbumByID)
    router.POST("/albums", postAlbums)

    router.Run("localhost:8080")
}
```

```
▶ curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
HTTP/1.1 201 Created
Content-Type: application/json; charset=utf-8
Date: Wed, 04 Aug 2021 08:47:36 GMT
Content-Length: 116

{
    "id": "4",
    "title": "The Modern Sound of Betty Carter",
    "artist": "Betty Carter",
    "price": 49.99
}%

Workspace/Go/golang.org  main ✗ system                                       1h56m ◒
▶ curl http://localhost:8080/albumscurl http://localhost:8080/albums \
    --header "Content-Type: application/json" \
    --request "GET"

Workspace/Go/golang.org  main ✗ system                                      1h57m ◒  ⍉
▶ curl http://localhost:8080/albums \
    --header "Content-Type: application/json" \
    --request "GET"
[
    {
        "id": "1",
        "title": "Blue Train",
        "artist": "John Coltrane",
        "price": 56.99
    },
    {
        "id": "2",
        "title": "Jeru",
        "artist": "Gerry Mulligan",
        "price": 17.99
    },
    {
        "id": "3",
        "title": "Sarah Vaughan and Clifford Brown",
        "artist": "Sarah Vaughan",
        "price": 39.99
    },
    {
        "id": "4",
        "title": "The Modern Sound of Betty Carter",
        "artist": "Betty Carter",
        "price": 49.99
    }
]%

Workspace/Go/golang.org  main ✗ system                                       1h59m ◒
▶ curl http://localhost:8080/albums/1 \
    --header "Content-Type: application/json" \
    --request "GET"
{
    "id": "1",
    "title": "Blue Train",
    "artist": "John Coltrane",
    "price": 56.99
}

Workspace/Go/golang.org  main ✗ system                                        2h0m ◒
▶ curl http://localhost:8080/albums/2 \
    --header "Content-Type: application/json" \
    --request "GET"
{
    "id": "2",
    "title": "Jeru",
    "artist": "Gerry Mulligan",
    "price": 17.99
}
```
