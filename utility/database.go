// database connection
package utility

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func GetConnection(postgresIp string) *gorm.DB {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		postgresIp, port, user, password, dbname)

	db, err := gorm.Open("postgres", psqlInfo)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")

	}

	log.Println("DB Connection established...")
	return db
}
