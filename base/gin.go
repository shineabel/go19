package main

import (
	"github.com/gin-gonic/gin"
	"database/sql"
	"net/http"
	"github.com/go19/pojo"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"os"
	"io"
	"log"
	"time"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/api/v1/user")
	v1.POST("/",CreateUser)
	v1.GET("/",GetAll)
	v1.GET("/:id",GetOne)
	v1.PUT("/",UpdateOne)
	v1.DELETE("/:id",DeleteOne)

	v1.POST("/upload",upload)

	r2 := r.Group("/v2")
	r2.GET("/test2",redirect)


	//r2.Use(middleware())
	//
	//{
	//
	//	r2.GET("/test3", func(c *gin.Context) {
	//		tag := c.MustGet("tag")
	//		chainId := c.MustGet("chainId")
	//		c.JSON(http.StatusOK, gin.H{
	//			"tag":     tag,
	//			"chainId": chainId,
	//		})
	//	})
	//}
	r2.GET("/auth", func(c *gin.Context) {
		cookie := &http.Cookie{
			Name:"session_id",
			Value:"123",
			Path:"/",
			HttpOnly:true,
		}
		http.SetCookie(c.Writer,cookie)
		c.JSON(http.StatusOK,"LoginSuccessful")

	})

	r3 := r.Group("/v3")
	r3.GET("/sync", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		fmt.Println(" sync done, ",c.Request.URL.Path)
	})

	r3.GET("/async", func(c *gin.Context) {

		cc := c.Copy()
		go func(tc *gin.Context) {
			time.Sleep(5 * time.Second)
			fmt.Println("async done.",tc.Request.URL.Path)
		}(cc)
	})

	r2.GET("/home",AuthMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"ok":"OK",
		})
	})
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
	c.XML(http.StatusOK,result)

	
}

func UpdateOne(c *gin.Context)  {
	id := c.Request.FormValue("id")
	name := c.Request.FormValue("name")

	db, err:= sql.Open("mysql","root:shine@/test?charset=utf8")
	checkError(err)
	defer db.Close()

	rs, err := db.Exec("update user_info set name = ? where id = ?",name, id)

	count, err := rs.RowsAffected()
	c.JSON(http.StatusOK,gin.H{
		"count":count,
	})
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

func upload(c *gin.Context)  {

	file, header, err :=c.Request.FormFile("upload")
	if err != nil {
		log.Fatal("upload error:",err)
		c.String(http.StatusBadRequest,"BadRequest")
		return
	}
	out, err := os.Create(header.Filename)
	defer out.Close()
	_, err2 := io.Copy(out,file)
	checkError(err2)
	c.String(http.StatusOK,"upload successful")
}

func redirect(c *gin.Context)  {

	c.Redirect(http.StatusTemporaryRedirect,"http://www.baidu.com")
}

func AuthMiddleware()  gin.HandlerFunc {

	return func(c *gin.Context) {
		//fmt.Println("-------before-----")
		//		//c.Set("tag","shine")
		//		//c.Set("chainId","rainbow-dev")
		//		//c.Next()
		//		//fmt.Println("------after -------")

		if cookie, err := c.Request.Cookie("session_id"); err == nil {
			v := cookie.Value
			if v == "123" {
				c.Next()
				return
			}
		}


		c.JSON(http.StatusUnauthorized, gin.H{
			"error":"Unauthorized",
		})
		c.Abort()
		return

	}
}

func checkError(err error)  {
	if err != nil {
		panic(err)
	}
}