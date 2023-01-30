package data

import (
	"fmt"
	"log"

	apyhub "github.com/apyhub/apyhub.go"
)

func init() {
	apyhub.InitApyHub("APT0rzEvAp76Izffo3ek7Zg7FpkK6csxzJ5puXCwWAw3lg0cFtj")
}

func Timezone() {

	timezone, err := apyhub.Timezones()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("timezone:%+v", timezone)
}
