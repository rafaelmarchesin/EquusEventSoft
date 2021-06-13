package funcs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/guark/guark/app"
)

type Event struct {
	EventName         string `json:"event_name"`
	EventResponsible  string `json:"event_responsible"`
	ResponsibleEmail  string `json:"responsible_email"`
	EventDate         string `json:"event_date"`
	EventTime         string `json:"event_time"`
	EventLocalization string `json:"event_localization"`
}

func convertData(c app.Context) Event {
	return Event{
		EventName:         fmt.Sprintf("%v", c.Params["eventName"]),
		EventResponsible:  fmt.Sprintf("%v", c.Params["eventeResponsible"]),
		ResponsibleEmail:  fmt.Sprintf("%v", c.Params["responsibleEmail"]),
		EventDate:         fmt.Sprintf("%v", c.Params["eventDate"]),
		EventTime:         fmt.Sprintf("%v", c.Params["eventTime"]),
		EventLocalization: fmt.Sprintf("%v", c.Params["eventLocation"]),
	}
}

func CreateEvents(c app.Context) (interface{}, error) {
	eventData := convertData(c)
	pathFile := "./ui/storage"
	//jsonFile, err := os.Open(`book.json`)

	files, _ := ioutil.ReadDir(pathFile)

	fileName := strconv.Itoa(len(files) + 1)

	//if len(files) > 0 {
	//	for _, file := range files {
	//		fmt.Println(file.Name())
	//	}
	//} else {
	//	fmt.Println("---------------------------")
	//}

	// Create JSON file
	file, _ := json.MarshalIndent(eventData, "", " ")
	_ = ioutil.WriteFile(pathFile+"/"+fileName+".json", file, 0644)

	return nil, nil
}
