package main
import (

	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"

)

var DB *sqlx.DB

func initDB()error{
	var err error
	dsn := "gopher:2020O229_@tcp(47.104.241.166:3306)/odin_cmdb?charset=utf8mb4"
	DB, err = sqlx.Open("mysql", dsn)
	if err != nil {
		return err
	}

	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(16)
	return nil
}

type User struct {
	Id int64 `db:"id"`
	Name string `db:"name"`  //如果一个字段中允许有空，则用nullstring
	Age int `db:"age"`
}
func sqlxQueryDate(){
	sqlstr := "select id, name, age from user where id>?"
	var user User
	err := DB.Get(&user, sqlstr, 3 ) //查询一行
	if err != nil{
		fmt.Printf("scan failed, err:%v\n", err)
	}

	fmt.Printf("ID:%d, Name:%s, Age:%d\n", user.Id, user.Name, user.Age)

}

func sqlxQueryMultilRow(){
	sqlstr := "select id, name, age from user where id>?"
	var user []User
	err := DB.Select(&user, sqlstr, 1  ) //查询多行
	if err != nil{
		fmt.Printf("scan failed, err:%v\n", err)
	}

	//fmt.Printf("ID:%d, Name:%v, Age:%d\n", user.Id, user.Name, user.Age)
	//fmt.Println(user)
	for i:=0; i<len(user); i++ {
		fmt.Println(user[i].Id, user[i].Name, user[i].Age)
	}
}

func sqlxupdateData()  {
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

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("failed ,err:%v\n", err)
	}

	fmt.Println("单行查询:")
	sqlxQueryDate()
	fmt.Println("多行查询:")
	sqlxQueryMultilRow()
	fmt.Println("更新查询:")
	sqlxupdateData()
}