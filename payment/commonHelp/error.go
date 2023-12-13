package commonhelp

import "log"

func FailCriticalOnErr(err error, msg string) {
	if err != nil {
		log.Fatalf(msg+" : %v", err)
	}
}
