package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(datastring string) (err error)
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, i := range dataset {
		err := dp.Parse(i)
		if err != nil {
			log.Println(err)
		}
		a, err1 := dp.ActionInfo()
		if err1 != nil {
			log.Println(err)
		}
		fmt.Println(a)
	}
}
