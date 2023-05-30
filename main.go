package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type Device struct {
	Name    string   `json:"Name"`
	Type    string   `json:"Type"`
	Info    string   `json:"Info"`
	Sensors []Sensor `json:"Sensors"`
}

type Sensor struct {
	Name    string `json:"Name"`
	Payload int    `json:"Payload"`
}

type ReformattedDevice struct {
	Name       string `json:"Name"`
	UUID       string `json:"UUID"`
	PayloadSum int    `json:"PayloadSum"`
}

type ReformattedData struct {
	Devices []ReformattedDevice `json:"Devices"`
}

func main() {
	// Read the data from the devices.json file
	data, err := ioutil.ReadFile("exercise-01/data/devices.json")
	if err != nil {
		fmt.Printf("Failed to read file: %v", err)
		return
	}

	// Parse the JSON data into a slice of devices
	var devices []Device
	err = json.Unmarshal(data, &devices)
	if err != nil {
		fmt.Printf("Failed to parse JSON: %v", err)
		return
	}

	// Extract UUID from the Info field and calculate sum of sensor payloads
	reformattedDevices := make([]ReformattedDevice, len(devices))
	for i, device := range devices {
		uuid := extractUUID(device.Info)
		payloadSum := calculatePayloadSum(device.Sensors)
		reformattedDevices[i] = ReformattedDevice{
			Name:       device.Name,
			UUID:       uuid,
			PayloadSum: payloadSum,
		}
	}

	// Sort the reformatted devices by Name (ascending)
	sort.Slice(reformattedDevices, func(i, j int) bool {
		return reformattedDevices[i].Name < reformattedDevices[j].Name
	})

	// Create the reformatted data
	reformattedData := ReformattedData{
		Devices: reformattedDevices,
	}

	// Convert the reformatted data to JSON format
	outputJSON, err := json.MarshalIndent(reformattedData, "", "    ")
	if err != nil {
		fmt.Printf("Failed to convert to JSON: %v", err)
		return
	}

	// Write the reformatted data to a new file
	err = ioutil.WriteFile("exercise-01/output.json", outputJSON, 0644)
	if err != nil {
		fmt.Printf("Failed to write file: %v", err)
		return
	}

	fmt.Println("Reformatted data has been written to output.json")
}

func extractUUID(info string) string {
	// Extract the UUID from the Info field by splitting the string
	// and getting the last element
	infoSplit := strings.Split(info, " ")
	uuid := infoSplit[len(infoSplit)-1]
	return strings.TrimPrefix(uuid, "uuid:")
}

func calculatePayloadSum(sensors []Sensor) int {
	// Calculate the sum of sensor payloads
	sum := 0
	for _, sensor := range sensors {
		sum += sensor.Payload
	}
	return sum
}
