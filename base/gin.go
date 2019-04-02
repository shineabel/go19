package main

import (
	"github.com/gin-gonic/gin"
	"database/sql"
	"net/http"
	"github.com/go19/pojo"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/api/v1/user")
	v1.POST("/",CreateUser)
	v1.GET("/",GetAll)
	v1.GET("/:id",GetOne)
	v1.PUT("/",UpdateOne)
	v1.DELETE("/:id",DeleteOne)
	r.Run()
}

func CreateUser(c *gin.Context)  {
	id := c.Request.FormValue("id")
	name := c.Request.FormValue("name")


	db, err := sql.Open("mysql","root:shine@/test?charset=utf8")
	checkError(err)

	defer db.Close()

	rs, err := db.Exec("insert into user_info(id,name) values(?,?)",id,name)
	insertId, err := rs.LastInsertId()

	c.JSON(http.StatusOK,gin.H{
		"insertId":insertId,
	})
	
}

func GetAll(c *gin.Context)  {
	db, err := sql.Open("mysql","root:shine@/test?charset=utf8")
	checkError(err)
	defer db.Close()
	rows , err := db.Query("select id,name from user_info")

	persons := make([]pojo.Person,0)
	for rows.Next() {
		var p pojo.Person
		rows.Scan(&p.Id, &p.Name)
		persons = append(persons,p)
	}
	if err = rows.Err(); err != nil {
		fmt.Printf("error:",err)
	}
	c.JSON(http.StatusOK, gin.H{
		"persons":persons,
	})
	
}

func GetOne(c *gin.Context)  {
	id := c.Param("id")
	db , err := sql.Open("mysql","root:shine@/test?charset=utf8")
	checkError(err)

	defer db.Close()

	var (
		p pojo.Person
		result gin.H
	)
	row := db.QueryRow("select id,name from user_info where id = ?",id)
	err = row.Scan(&p.Id,&p.Name)
	if err == nil {
		result = gin.H{
			"result":p,
			"count":1,
		}
	} else {

		result = gin.H{
			"result":nil,
			"count":0,
		}
	}
	c.JSON(http.StatusOK,result)

	
}

func UpdateOne(c *gin.Context)  {
	
}

func DeleteOne(c *gin.Context)  {
	id := c.Param("id")
	db, err :=sql.Open("mysql","root:shine@/test?charset=utf8")
	checkError(err)
	defer db.Close()

	stmt , err2 := db.Prepare("delete from user_info where id = ?")
	checkError(err2)
	res, err := stmt.Exec(id)
	count, err  := res.RowsAffected()
	fmt.Printf("delete count %d:",count)
	if count == 1 {
		c.JSON(http.StatusOK,gin.H{
			"result":"OK",
		})
	} else {
		c.JSON(http.StatusOK,gin.H{
			"result":"Failure",
		})
	}
	
}

func checkError(err error)  {
	if err != nil {
		panic(err)
	}
}