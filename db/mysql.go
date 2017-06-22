package db

import (
	"github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
	"github.com/petronetto/echo-sample/conf"
)

func Init() *dbr.Session {

	session := getSession()

	return session
}

func getSession() *dbr.Session {

	db, err := dbr.Open("mysql",
		conf.USER+":"+conf.PASSWORD+"@tcp("+conf.HOST+":"+conf.PORT+")/"+conf.DB,
		nil)

	if err != nil {
		logrus.Error(err)
	} else {
		session := db.NewSession(nil)
		return session
	}
	return nil
}
