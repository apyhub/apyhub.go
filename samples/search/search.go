package search

import (
	"fmt"
	"log"

	apyhub "github.com/apyhub/apyhub.go"
)

func init() {
	apyhub.InitApyHub("Enter your apyhub token")
}

func Search() {
	matched, err := apyhub.FuzzySearch("tc", "tic tac tok", false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(matched)
}
