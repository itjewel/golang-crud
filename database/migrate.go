package database

import (
	"fmt"

	"io/ioutil"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func RunMigration(action string) error {
	suffix := ""
	switch action {
	case "up":
		suffix = ".up.sql"
	case "down":
		suffix = ".down.sql"
	default:
		return fmt.Errorf("invalid action: %s", action)
	}

	files, err := ioutil.ReadDir("migrations")
	if err != nil {
		return err
	}

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), suffix) {
			continue
		}

		path := "migrations/" + file.Name()
		sqlBytes, err := ioutil.ReadFile(path)
		if err != nil {
			log.Println("Cannot read:", path)
			continue
		}

		_, err = DB.Exec(string(sqlBytes))
		if err != nil {
			log.Println("Migration failed:", path, err)
		} else {
			fmt.Println("Executed:", path)
		}
	}

	return nil
}
