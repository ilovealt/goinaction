package database

import (
	"database/sql"
	"fmt"
)

type SysDict struct {
	Id   int64
	Code string
	Name string
}

func SelectDictById(db *sql.DB, id int64) (dict SysDict) {
	var stm *sql.Stmt
	stm, err := db.Prepare("select id, code, name from sys_dict where id = $1")
	if err != nil {
		fmt.Println("操作SysDict失败! err:", err.Error())
		return
	}
	defer stm.Close()

	// 赋值
	err = stm.QueryRow(id).Scan(&dict.Id, &dict.Code, &dict.Name)
	if err != nil {
		fmt.Println("赋值SysDict失败! err:", err.Error())
		return
	}

	return
}

func SelectDictByCode(db *sql.DB, code string) (dict SysDict) {
	var stm *sql.Stmt
	stm, err := db.Prepare("select id, code, name from sys_dict where code = $1")
	if err != nil {
		fmt.Println("操作SysDict失败! err:", err.Error())
		return
	}
	defer stm.Close()

	// 赋值
	err = stm.QueryRow(code).Scan(&dict.Id, &dict.Code, &dict.Name)
	if err != nil {
		fmt.Println("赋值SysDict失败! err:", err.Error())
		return
	}

	return

}

func SelectAllDict(db *sql.DB) (dictList []SysDict) {
	var rows *sql.Rows
	rows, err := db.Query("select id, code, name from sys_dict")
	if err != nil {
		fmt.Println("操作SysDict失败! err:", err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		dict := SysDict{}
		err = rows.Scan(&dict.Id, &dict.Code, &dict.Name)
		if err != nil {
			fmt.Println("操作SysDict失败! err:", err.Error())
			return
		}

		dictList = append(dictList, dict)
	}
	return

}

func InsertDict(db *sql.DB, dict SysDict) (insertId int64, affectedRows int64) {
	var stm *sql.Stmt
	stm, err := db.Prepare("INSERT INTO sys_dict (id, code, name) VALUES ($1,$2,$3)")
	if err != nil {
		fmt.Println("操作SysDict失败! err:", err.Error())
		return
	}
	defer stm.Close()

	res, err := stm.Exec(&dict.Id, &dict.Code, &dict.Name)
	if err != nil {
		fmt.Println("赋值SysDict失败! err:", err.Error())
		return
	}

	insertId, _ = res.LastInsertId()
	affectedRows, _ = res.RowsAffected()
	return
}

func DeleteDict(db *sql.DB, id int64) (affectedRows int64) {
	var stm *sql.Stmt
	stm, err := db.Prepare("DELETE FROM sys_dict WHERE id = $1")
	if err != nil {
		fmt.Println("操作SysDict失败! err:", err.Error())
		return
	}
	defer stm.Close()

	res, err := stm.Exec(id)
	if err != nil {
		fmt.Println("赋值SysDict失败! err:", err.Error())
		return
	}

	affectedRows, _ = res.RowsAffected()
	return
}
