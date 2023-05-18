package project

import (
	"database/sql"
	"log"
	"mhygacha/global"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func SQLiteInit() {
	dbFile := "./data.db"
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		log.Println("未检测到数据库文件，开始创建数据库......")
		global.SQLDB, err = sql.Open("sqlite3", dbFile)
		if err != nil {
			log.Fatal(err)
		}
		//defer global.SQLDB.Close()

		_, err = global.SQLDB.Exec(`
		CREATE TABLE IF NOT EXISTS srgachalog (
			Uid TEXT,
			GachaId TEXT,
			GachaType TEXT,
			ItemId TEXT,
			Count TEXT,
			Time TEXT,
			Name TEXT,
			Lang TEXT,
			ItemType TEXT,
			RankType TEXT,
			Id TEXT,
			Authkey TEXT
		);
		`)
		if err != nil {
			log.Printf("发生错误: %q \n", err)
			return
		}
		log.Println("数据库创建完成!")
	} else {
		log.Println("数据库文件存在，开始调用数据库.")
		global.SQLDB, err = sql.Open("sqlite3", dbFile)
		if err != nil {
			log.Fatal(err)
		}
		//defer global.SQLDB.Close()
	}
}
