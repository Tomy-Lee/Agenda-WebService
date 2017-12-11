package entities

import (

	// mysql driver
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
	// gorm sqlite3 driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type serv interface {
	load()
}

var (
	gormDb *gorm.DB
	servs  []serv
)

func addServ(s serv) {
	servs = append(servs, s)
}

// Init ..
func Init(dbPath string) {
	db, err := gorm.Open("sqlite3", dbPath)
	checkErr(err)
	gormDb = db

	var err2 interface{}
	finished := make(chan bool)
	for _, s := range servs {
		go func(s serv) {
			defer func() {
				if e := recover(); e != nil {
					err2 = e
				}
				finished <- true
			}()
			s.load()
		}(s)
	}
	// wait for all servs to finish loading
	for _ = range servs {
		<-finished
	}
	if err2 != nil {
		log.Fatal(err2)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
