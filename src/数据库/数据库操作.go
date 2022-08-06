package main

import (
	"fmt"
)

// insert
func insertData() {
	sqlStr := "insert into user_tbl(username, password) values (?,?)"
	ret, err := db.Exec(sqlStr, "张三", "123456")
	if err != nil {
		fmt.Println("insert failed:", err)
		return
	}

	theID, err := ret.LastInsertId() // 新插入数据的ID
	if err != nil {
		fmt.Println("getID failed:", err)
		return
	}
	fmt.Println("theID:", theID)
}

type User struct {
	id       int
	username string
	password string
}

// queryOneRow
func queryOneRow() {
	s := "select * from user_tbl where id = ?"
	var u User
	err = db.QueryRow(s, 1).Scan(&u.id, &u.username, &u.password)
	if err != nil {
		fmt.Println("queryOneRow failed:", err)
		return
	}
	fmt.Println("u", u)
}

// queryRows
func queryRows() {
	s := "select * from user_tbl"
	r, err := db.QueryRow(s)
	if err != nil {
		fmt.Println("queryRows failed:", err)
		return
	}
	defer r.Close() // 释放连接

	for r.Next() {
		var u User
		err := r.Scan(&u.id, &u.username, &u.password)
		if err != nil {
			fmt.Println("queryRows failed:", err)
			return
		}
		fmt.Println("u", u)
	}
}

// update
func update() {
	sql := "update user_tbl set username=?, password=? where id=?"
	ret, err := db.Exec(sql, "name", "112233", "2")
	if err != nil {
		fmt.Println("error failed:", err)
		return
	}
	rows, err := ret.RowsAffected()
	if err != nil {
		fmt.Println("error failed:", err)
		return
	}
	fmt.Println("更新成功,更新的行数", rows)
}

// delete
func delData() {
	sql := "delete from user_tbl where id=?"
	ret, err := db.Exec(sql, "1")
	if err != nil {
		return
	}
	row, err := ret.RowsAffected()
	if err != nil {
		return
	}
	fmt.Println(row)
}
func main() {

}
