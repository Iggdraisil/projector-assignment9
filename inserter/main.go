package main

import (
	//  "bytes"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-faker/faker/v4"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	} else {
		fmt.Println("successful...")
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/mydb?charset=utf8")
	checkErr(err)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)
	err = db.Ping()
	checkErr(err)

	router := gin.New()
	router.Use(gin.Recovery())
	router.POST("/person", func(c *gin.Context) {
		address := faker.GetRealAddress()
		_, err := db.Exec("insert into user (`name`, `birthdate`, `email`, `state`, `city`) values (?,?,?,?,?)",
			faker.Name(), faker.Date(), faker.Email(), address.State, address.City,
		)
		if err != nil {
			c.JSON(500, gin.H{
				"user":  nil,
				"count": 0,
			})
			fmt.Println(err)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"count": 1,
			})
			//lastId, exc := result.LastInsertId()
			//if exc != nil {
			//	fmt.Printf("err=%s \n", exc)
			//} else {
			//	fmt.Printf("last_id=%d \n", lastId)
			//}
		}
	})
	err = router.Run(":3000")
	if err != nil {
		panic(err)
	}
}
