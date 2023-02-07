package generate

import (
	"fmt"
	"log"
	"os"

	apyhub "github.com/apyhub/apyhub.go"
	h "github.com/apyhub/apyhub.go/helper"
)

func init() {
	apyhub.InitApyHub("Enter your apyhub token")
}

func Ical() {
	event := h.CalenderEvent{
		Summary:         "Final Call",
		Description:     "Casting for the New James Bond Movie",
		OrganizerEmail:  "johndoe@apyhub.com",
		AttendeesEmails: []string{"mark@apyhub.com"},
		Location:        "US",
		TimeZone:        "Europe/Helsinki",
		StartTime:       "0:00",
		EndTime:         "0:00",
		MeetingDate:     "30-11-2022",
		Reccuring:       true,
		Recurrence: h.Reccurent{
			Frequency: "WEEKLY",
			Count:     2,
		},
	}
	byt, err := apyhub.IcalGeneratorAsFile(event)
	if err != nil {
		log.Fatal(err)
	}

	// write bytes in specific file
	if err := os.WriteFile("./test_file/ical.ics", byt, 0666); err != nil {
		log.Fatal(err)
	}

	url, err := apyhub.IcalGeneratorAsURL(event)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)
}
