package main

 import (
	"fmt"
	"time"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
 )

 // Activities table SQL :
 // id          bigint(20)  AUTO_INCREMENT
 // username    varchar(50)
 // created_on  timestamp
 // action      char(1)
 // description varchar(300)
 // visibility  char(1)


 // NOTE : In the struct, CreatedOn will be translated into created_on in sql level
type Activities struct {
        Id          int       `sql:"AUTO_INCREMENT"`
        Username    string    `sql:"varchar(50);unique"`
        CreatedOn   time.Time `sql:"timestamp"`
        Action      string    `sql:"type:char(1)"`
        Description string    `sql:"type:varchar(300)"`
        Visibility  string    `sql:"type:char(1)"`
}

var db *gorm.DB
var err error

func injection(db gorm.DB){
	log.Print("Injecting data")
	for i :=0; i<10; i++{

		activity := Activities{
			Username:    "testuser"+strconv.Itoa(i),
			CreatedOn:  time.Now().Local(),
			Description: "Testing",
			Visibility:  "S",
		}
		db.Create(&activity)
		db.Last(&activity)
		log.Print(activity)
	}
}


func main() {

	dsn := "host=localhost user=admin password=password dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
			fmt.Println(err)
	}

	db.AutoMigrate(&Activities{})

	// Injection data
	injection(*db)
}