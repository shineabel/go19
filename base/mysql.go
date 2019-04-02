package main

import "database/sql"
import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func main() {

	db , err := sql.Open("mysql","root:shine@/test?charset=utf8")
	checkErr(err)
	//stmt, err := db.Prepare("insert into user_info(id,name) values(?,?)")
	//checkErr(err)
	//
	//res, err := stmt.Exec("1","shine")
	//checkErr(err)
	//fmt.Println(res)

	stmt, err := db.Prepare("update user_info set name = ? where id = ?")
	checkErr(err)
	res, err := stmt.Exec("shine34",1)
	 c , err := res.RowsAffected()
	fmt.Printf("update count %d:",c)


	 stmt.Close()
	 db.Close()

}

func checkErr(err error)  {
	if err != nil {
		panic(err)
	}
}