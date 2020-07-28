package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB //连接池对象
var err error

const (
	USER_NAME = "root"
	USER_PWD  = "mytest"
	HOST      = "172.24.0.4"
	PORT      = "3306"
	DATABASE  = "figureprint_prod"
	CHARSET   = "utf8"
)

func connMysql() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", USER_NAME, USER_PWD, HOST, PORT, DATABASE, CHARSET)
	Db, err = sql.Open("mysql", dsn)
	//defer Db.Close()
	if err != nil {
		log.Println("Mysql Connect Error")
		log.Printf("连接配置不正确:%s", err.Error())
		return
	}

	// 最大连接数
	Db.SetMaxOpenConns(100)
	// 闲置连接数
	Db.SetMaxIdleConns(20)
	// 最大连接周期
	Db.SetConnMaxLifetime(100 * time.Second)

	if err = Db.Ping(); nil != err {
		log.Println("Mysql Connect Error: " + err.Error())
	}
	return
}

type weapon struct {
	id         int
	weaponName string
	modeId     int
}

func queryList() {
	sql := `SELECT id,weapon_name,weapon_mode_id FROM weapons limit 10`
	rows, err := Db.Query(sql)
	if err != nil {
		log.Printf("exec %s query failed ,err :%v\n", sql, err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var weapon1 weapon
		err := rows.Scan(&weapon1.id, &weapon1.weaponName, &weapon1.modeId)
		if err != nil {
			log.Printf("scan failed ,err: %v \n", err)
		}
		fmt.Printf("weapon:%#v\n", weapon1)
		//rows.Scan()

	}

}

func main() {
	fmt.Println("Start mysql---")
	log.Println("go")
	connMysql()
	queryList()
}
