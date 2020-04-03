package dao

import (

	sql "github.com/jmoiron/sqlx"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitMysql()(err error){
	fmt.Println("InitMysql....")
	if db == nil {
		db, err = sql.Connect("mysql", "gopher:2020O229_@tcp(47.104.241.166:3306)/go-blog?charset=utf8mb4")
		if err != nil {
			return
		}
	}

	return
}


func ModifyDB(sql string, args ... interface{})(int64, error){
	result, err := db.Exec(sql, args...)
	if err!= nil{
		fmt.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return  count, nil
}

func QueryRowDB(dest interface{},sql string, args ...interface{}) error{
	return db.Get(dest,sql, args ...)
}
