package mq

import (
	"github.com/google/uuid"
	"strings"
	"time"
)

func createUUID() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}

func getDateByFormat(timeParams ...time.Time) string {
	timeVars := time.Now()
	if len(timeParams) > 0 {
		timeVars = timeParams[0]
	}
	return timeVars.Format("2006-01-02 15:04:05")
}
