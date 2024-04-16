package logging

import (
	"log"
	"os"
)

func Init() {

	// Open a log file for writing (create if not exists)
	_, err := os.OpenFile("library.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

}
