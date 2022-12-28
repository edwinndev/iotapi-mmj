package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// @tcp(us-west.connect.psdb.cloud)/iot?tls=true
const DSN = "hd9vbn36x1vrkfl4nwwq:pscale_pw_b79vrXFeLycC49ty6ab1gX7qqQkNzaSKOX7m9lk1Pk5@tcp(us-west.connect.psdb.cloud:3306)/iot?tls=true&parseTime=true"

var Mysql *gorm.DB

func Connect() {
	var err error
	Mysql, err = gorm.Open(mysql.Open(DSN))
	if err != nil {
		log.Fatal("Error al conectar a DB \n" + err.Error())
	}
	log.Println("MySQL connection successful")
}
