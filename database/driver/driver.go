package driver

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/boltdb/bolt"
)

type driver struct {
	db *bolt.DB
}

var documentIDFieldName = "_id"
var driverInstance *driver
var databasePath string = "./data/main.db"

func GetDriver() (*driver, error) {
	if driverInstance.db == nil {
		return nil, errors.New("database not initialized")
	}
	return driverInstance, nil
}

func OpenDatabase() {
	if _, err := os.Stat(databasePath); errors.Is(err, os.ErrNotExist) {
		log.Print("database does not exist, creating")
		err := os.WriteFile(databasePath, []byte{}, 0600)
		if err != nil {
			log.Print("failed to create database, please create a file manually")
			log.Fatal(err)
		}
	}

	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open(databasePath, 0600, &bolt.Options{Timeout: 1 * time.Second, ReadOnly: false})
	if err != nil {
		log.Fatal(err)
	}
	var d driver
	d.db = db
	driverInstance = &d
}

func CloseDatabase() {
	driverInstance.db.Close()
}
