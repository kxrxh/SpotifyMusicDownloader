package utils

import "log"

func PanicErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func FatalErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
