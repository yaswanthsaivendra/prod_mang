package database

import (
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB

func InitDB() *gorm.DB {

	db, err := gorm.Open(sqlite.Open("./../gorm.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Sucessfully connected to the db")
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)

	Database = db
	return Database
}

// Init the test db for test cases
func InitTestDB() *gorm.DB {

	test_db, err := gorm.Open(sqlite.Open("./../gorm_test.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Sucessfully connected to the test db")
	}

	sqlDB, err := test_db.DB()
	sqlDB.SetMaxIdleConns(3)

	Database = test_db
	return Database
}

// Delete the database after running testing cases.
func TestDBFree(test_db *gorm.DB) error {
	sqlDB, err := test_db.DB()
	sqlDB.Close()

	err = os.Remove("./../gorm_test.db")
	return err
}

// Using this function to get a connection, you can create your connection pool here.
func GetDB() *gorm.DB {
	return Database
}
