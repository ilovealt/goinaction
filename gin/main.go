package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	gindb "github.com/ilovealt/goinaction/gin/database"
)

func main() {
	db, err := sql.Open("postgres", "host='' port=5432 user=''  password=''  dbname='' sslmode=require")
	if err != nil {
		fmt.Println("创建数据库文件失败！ err:", err.Error())
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("连接数据库失败！ err:", err.Error())
		return
	} else {
		fmt.Println("PING数据库成功!")
	}

	sysDictList := gindb.SelectAllDict(db)
	fmt.Println("查询所有数据", sysDictList)

	var sysDictNew gindb.SysDict
	var newId int64 = 11
	sysDictNew.Id = newId
	sysDictNew.Code = "11"
	sysDictNew.Name = "permission"

	_, affectedRows := gindb.InsertDict(db, sysDictNew)
	fmt.Println("新增ID:", newId, " 影响行数:", affectedRows)

	sysDict := gindb.SelectDictById(db, newId)
	fmt.Println("根据ID获取实体:", sysDict)

	affectedRows = gindb.DeleteDict(db, newId)
	fmt.Println("根据ID删除,影响行数:", affectedRows)

	sysDictList = gindb.SelectAllDict(db)
	fmt.Println("查询所有数据", sysDictList)
}
