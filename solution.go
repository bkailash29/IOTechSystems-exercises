package main

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type Device struct {
	Name      string `json:"Name"`
	Type      string `json:"Type"`
	Info      string `json:"Info"`
	Value     string `json:"value"`
	Timestamp string `json:"timestamp"`
}

type Output struct {
	TotalValue int      `json:"TotalValue"`
	UUIDs      []string `json:"UUIDs"`
}

func main() {
	// Step 1: Parse the data from data.json
	file, err := ioutil.ReadFile("exercise-02/data/data.json")
	if err != nil {
		panic(err)
	}

	var devices []Device
	err = json.Unmarshal(file, &devices)
	if err != nil {
		panic(err)
	}

	// Step 2: Discard devices with timestamps before the current time
	currentTime := time.Now().Unix()
	var filteredDevices []Device
	for _, device := range devices {
		timestamp, err := strconv.ParseInt(device.Timestamp, 10, 64)
		if err != nil {
			panic(err)
		}

		if timestamp >= currentTime {
			filteredDevices = append(filteredDevices, device)
		}
	}

	// Step 3: Get the total of all value entries and parse the uuid
	totalValue := 0
	var uuidList []string
	for _, device := range filteredDevices {
		value, err := base64.StdEncoding.DecodeString(device.Value)
		if err != nil {
			panic(err)
		}

		valueInt, err := strconv.Atoi(string(value))
		if err != nil {
			panic(err)
		}

		totalValue += valueInt

		uuid := strings.Split(device.Info, "uuid:")[1]
		uuid = strings.Split(uuid, ",")[0]
		uuidList = append(uuidList, uuid)
	}

	// Step 4: Output the values total and the list of uuids
	output := Output{
		TotalValue: totalValue,
		UUIDs:      uuidList,
	}

	// Step 5: Write the data to a file in JSON format
	outputJSON, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("output.json", outputJSON, 0644)
	if err != nil {
		panic(err)
	}
}
