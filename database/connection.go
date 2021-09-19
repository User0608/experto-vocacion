package database

import (
	"log"
	"sync"

	"github.com/user0608/expertos/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var ones sync.Once
var gdb *gorm.DB

func GetDBConnextion(conf *configs.DBConfigs) *gorm.DB {
	ones.Do(func() {
		schema := schema.NamingStrategy{SingularTable: true}
		db, err := gorm.Open(postgres.Open(conf.GetConnectionString()), &gorm.Config{NamingStrategy: schema})
		if err != nil {
			log.Fatalln("No se pudo connectar la base de datos:", err.Error())
		}
		gdb = db
	})
	return gdb
}
