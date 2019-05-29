package utils

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/user/contactsApi/models"
	"io"
	"log"
	"net/http"
	"os"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{} {"status" : status, "message" : message}
}

func Respond (w http.ResponseWriter, data map[string] interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func ReadStatesFromCsv() []models.State {
	csvFile, error := os.Open("data/states.csv")
	if error != nil {
		fmt.Println(error)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	fmt.Print(csvFile)
	var states []models.State
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			fmt.Println(error)
			log.Fatal(error)
		}
		states = append(states, models.State{
			Abbreviation: line[0],
			Name: line[1],
		})
	}
	statesJson, _ := json.Marshal(states)
	fmt.Println(string(statesJson))
	return states
}