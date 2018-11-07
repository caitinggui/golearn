package main

import (
	"log"
	"os"

	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func Database(db *sql.DB) gin.HandlerFunc {
	log.Println("do something once")
	return func(c *gin.Context) {
		c.Set("SqlDb", db)
		c.Next()
	}
}

func MyHandle(c *gin.Context) {
	db := c.MustGet("SqlDb").(*sql.DB)
	sqlStmt := `create table foo (id integer not null primary key, name text);`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		c.String(500, "create table error")
		return
	}

	sqlInsert := "insert into foo(id, name) values(1, 'test')"
	_, err = db.Exec(sqlInsert)
	if err != nil {
		log.Fatal(err)
		c.String(500, "insert data error")
	}

	row, _ := db.Query("select id, name from foo limit 1")
	var id int
	var name string
	for row.Next() {
		row.Scan(&id, &name)

	}

	c.JSON(200, gin.H{"id": id, "name": name})
}

// 程序的入口
// 这边可以用来建立路由
func main() {
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	os.Remove("./foo.db")
	log.Println("remove foo.db")

	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		log.Fatal(err)

	}
	defer db.Close()
	r.Use(Database(db))
	r.GET("/", MyHandle)

	r.Run() // listen and serve on 0.0.0.0:8080

}
