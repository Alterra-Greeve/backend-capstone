package helper

import "log"

func Logger(message string) {
	log.Println("[LOG_DE_START] ", message, " [LOG_DE_END]")
}
