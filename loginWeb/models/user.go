package models

import (
	"fmt"
	"netgo/loginWeb/dao"
)

type User struct {
	Id int
	Username string
	Password string
	Email string
	Status int
	CreateTime string
}

//插入数据
func InsertUser(user *User) (int64, error){
	fmt.Println(user.Username,user.Password, user.Email,user.Status,user.CreateTime)

	fmt.Println("SQL:")
	fmt.Printf("insert into z_users(name,password,email,status,created_at)values(%s, %s, %s, %v, %s)",
		user.Username,user.Password, user.Email,user.Status,user.CreateTime)
	return dao.ModifyDB("insert into z_users(name,password,email,status,created_at)values(?,?,?,?,?)",
		user.Username,user.Password, user.Email,user.Status,user.CreateTime)
}

//根据用户名查询id
func QuerUserWithUsername(username string) int{
	var user User
	//row := dao.QueryRowDB("select id from z_users where name=?", username)

	//id := 0
	//row.Scan(&id)
	//return id
	err := dao.QueryRowDB(&user, "select id from z_users where name=?", username)
	if err != nil {
		return 0
	}
	return user.Id
}

func QuerUserWithParam(username, password string)int{
	fmt.Printf("user:%v, password:%v\n", username, password)

	//row := dao.QueryRowDB("select id from z_users where name=? and password=?", username,password)
	//id := 0
	//row.Scan(&id)
	//return id
	var user User
	err := dao.QueryRowDB(&user, "select id from z_users where name=? and password=?", username,password)
	if err != nil {
		return 0
	}
	return user.Id
}