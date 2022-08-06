package 数据库

import (
	"database/sql"
)

// 定义一个对象db
var db *sql.DB

// DB init
func initDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_db?charset=utf8&parseTime=true"
	// Not sure if the password is correct
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	// test connect DB
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	//err := initDB()
	//if err != nil {
	//	fmt.Println("初始化失败")
	//} else {
	//	fmt.Println("初始化成功")
	//}
}
