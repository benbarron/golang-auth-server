package services

import (
	"log"
	"os"
	"time"
)

type LoggingService struct {
	Context 	string
}

func NewLoggingService(context string) *LoggingService {
	return &LoggingService{
		Context: context,
	}
}

func (l *LoggingService) Log(message string) {
	if os.Getenv("SERVER_DEBUG") == "true" {
		log.Println(time.Now().UTC(), "  ", "[", l.Context, "]", "  ", message)
	}
}