package main

import (
	"fmt"
	"database/sql"
	"github.com/pkg/errors"
 )
 
 type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
 
 ~
 var Db *sql.DB
 
 func init() {
	var err error
	Db, err = sql.Open("mysql", "root:ttalbe@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
	   panic(err)
	}
 }
 
 func QueryUser() (User, error) {
	var user User
	row := Db.QueryRow("SELECT id, name, age FROM user")
	err := row.Scan(&user.ID, &user.Name)
	if err != nil{
	   return user,errors.Wrap(err,"dao#select user err")
	}
	return user,nil
 }
 
 func main()  {
	user ,err :=QueryUser()
	if err != nil{
	  fmt.Printf("query user err : %+v",err)
	  return~
	}
	fmt.Println("query user : ",user)
 }