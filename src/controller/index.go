package controller

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Work struct {
	Id   int
	Name string
}

func IndexGET(c *gin.Context) {
	data := gin.H{}
	data["title"] = "Title"
	data["content"] = "This is a content"

	db, _ := sql.Open("mysql", "root:root@/master")
	defer db.Close()

	work := Work{}
	db.QueryRow("SELECT id, name FROM works where id = ?", 1).Scan(&work.Id, &work.Name)

	data["content"] = work.Name
	fmt.Println(work)
	c.HTML(http.StatusOK, "index.tmpl", data)
}
