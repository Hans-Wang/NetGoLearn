package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func initDB()error {
	var err error
	DB, err = sql.Open("mysql", "gopher:2020O229_@tcp(47.104.241.166:3306)/odin_cmdb?charset=utf8mb4")

	if err != nil {
		fmt.Println("连接失败", err)
	}

	return nil
}

func testQueryDate()  {
	//insertStr := "insert into user (id,name,age) values(?,?,?)"
	//DB.QueryRow(insertStr, 2,"bin",20 )
	//DB.QueryRow(insertStr, 3,40 )


	sqlstr := "select id, name, age from user where id>?"
	rows := DB.QueryRow(sqlstr, 1 ) //查询一行
	//DB.Query() //查询多行
	var user User

	err := rows.Scan(&user.Id, &user.Name,&user.Age)
	if err != nil{
		fmt.Printf("scan failed, err:%v\n", err)
	}

	fmt.Printf("ID:%d, Name:%s, Age:%d\n", user.Id, user.Name, user.Age)
}

func testQueryMultilRow(){
	sqlstr := "select id, name, age from user where id>?"
	rows, err := DB.Query(sqlstr, 0 ) //查询多行
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	if err != nil{
		fmt.Printf("scan failed, err:%v\n", err)
	}
	var user User
	for rows.Next(){
		err := rows.Scan(&user.Id, &user.Name,&user.Age)
		if err != nil{
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}

		fmt.Printf("ID:%d, Name:%v, Age:%d\n", user.Id, user.Name, user.Age)
	}
}

type User struct {
	Id int64 `db:"id"`
	Name sql.NullString `db:"name"`  //如果一个字段中允许有空，则用nullstring
	Age int `db:"age"`
}

func testInsertData(){
	sqlstr:= "insert into user(name, age)values(?,?)"
	result, err := DB.Exec(sqlstr, "apache",11)

	if err != nil{
		fmt.Printf("Insert data failed, err:%v\n", err)
		return
	}


	id, err := result.LastInsertId()
	if err!= nil {
		fmt.Printf("get last insert id failed, error:%s\n", err)
	}
	fmt.Println("id",id)
}

func testupdateData()  {
	sqlstr := "update user set name=?,age=? where id=?"
	result, err := DB.Exec(sqlstr, "python",33,4)

	if err != nil{
		fmt.Printf("Update data failed, err:%v\n", err)
		return
	}
	n, err := result.RowsAffected()

	if err != nil{
		fmt.Printf("get  RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("更新影响了:%d行\n",n )
}

func testdeleteData()  {
	sqlstr := "delete from user where id=?"
	result, err := DB.Exec(sqlstr, 2)

	if err != nil{
		fmt.Printf("delete data failed, err:%v\n", err)
		return
	}
	n, err := result.RowsAffected()

	if err != nil{
		fmt.Printf("get  RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("删除了:%d行\n",n )
}

func transUpdate(){
	conn, err := DB.Begin()

	if err != nil {
		if conn != nil {
			conn.Rollback()
		}
		fmt.Printf("begin failed, err %v\n", err)
		return
	}

	sqlstr := "update user set age=? where id=?"
	_, err = conn.Exec(sqlstr, 30, 1)
	if err != nil {
		conn.Rollback()
		fmt.Printf("exec sql:%s failed, err:%v\n", sqlstr, err)
		return
	}
	sqlstr2 := "update user set age=? where id=?"
	_, err = conn.Exec(sqlstr2, 33,4)
	if err != nil {
		conn.Rollback()
		fmt.Printf("exec sql:%s failed, err:%v\n", sqlstr2, err)
		return
	}

	err = conn.Commit()
	if err != nil {
		conn.Rollback()
		fmt.Printf("commit failed, err:%v\n", err)
		return
	}

	fmt.Println("提交成功。")
}


func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init db failed, err :%v\n", err)
		return
	}
	fmt.Println("单行查询:")
	testQueryDate()
	fmt.Println("多行查询:")
	testQueryMultilRow()
	fmt.Println("插入(insert):")
	//testInsertData()
	fmt.Println("更新(update):")
	//testupdateData()

	fmt.Println("删除(delete):")
	//testdeleteData()

	fmt.Println("事物:")
	transUpdate()
}
