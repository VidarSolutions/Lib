package main

import (
	"fmt"

	country "github.com/VidarSolutions/Lib/Countries"
	d "github.com/VidarSolutions/Lib/Dates"
	f "github.com/VidarSolutions/Lib/Files"
	l "github.com/VidarSolutions/Lib/Log"
	m "github.com/VidarSolutions/Lib/Maps"
)

func main() {
	test()

}

func test() {
	//Palce holder for future funcitonality
	//Category := cat.Categories.Category{}
	Country := country.CountryPhoneCodes
	f.EnsureDir("logs")
	currentDate := d.GetDate("yyymmdd")
	AppLog, _ := l.GetLogger("logs/app.log")

	fmt.Printf("Date : %s\n", currentDate)
	AppLog.WriteString("Hello\n")
	md := m.NewMD()
	// Create a new Map and add the MapData to it

	MyMap := md.NewMap("Test")

	fmt.Printf("Area Code for Country %s is %s \n\n", Country[12].Country, Country[12].Code)
	// Print each field of the MapData
	for key, data := range MyMap.Data {
		fmt.Printf("Key: %s\n", key)
		fmt.Printf("DataRootDir: %s\n", data.DataRootDir)
		fmt.Printf("Description: %s\n", data.Description)
		fmt.Printf("FilePath: %s\n", data.FilePath)
		fmt.Printf("Type: %s\n", data.Type)
		for k, v := range data.Data {
			fmt.Printf("Data[%s]: %s\n", k, v)
		}
	}

}
