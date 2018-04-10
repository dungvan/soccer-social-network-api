package infrastructure

import (
	"github.com/jinzhu/gorm"
	// blank import.
	_ "github.com/lib/pq"
)

// SQL struct.
type SQL struct {
	DB *gorm.DB
}

type dbInfo struct {
	dbms    string
	host    string
	user    string
	pass    string
	name    string
	logmode bool
}

// NewSQL returns new SQL.
func NewSQL() *SQL {
	v := dbInfo{
		dbms:    GetConfigString("database.dbms"),
		host:    GetConfigString("database.host"),
		user:    GetConfigString("database.user"),
		pass:    GetConfigString("database.pass"),
		name:    GetConfigString("database.name"),
		logmode: GetConfigBool("database.logmode"),
	}
	connect := "host=" + v.host + " user=" + v.user + " dbname=" + v.name + " sslmode=disable password=" + v.pass
	db, err := gorm.Open(v.dbms, connect)
	db.LogMode(v.logmode)
	if err != nil {
		panic(err)
	}
	return &SQL{db}
}
