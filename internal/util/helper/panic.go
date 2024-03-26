package helper

import "log"

func PanicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func PanicErrIfErr(errCondition, err error) {
	if errCondition != nil {
		panic(err)
	}
}

func PanicErrIfNotErr(errCondition, err error) {
	if errCondition != nil {
		panic(err)
	}
}

func LogIfErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
