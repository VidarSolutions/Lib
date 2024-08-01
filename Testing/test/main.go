package main

import (
	"fmt"

	cat "github.com/VidarSolutions/Lib/Categories"
	country "github.com/VidarSolutions/Lib/Countries"
	d "github.com/VidarSolutions/Lib/Dates"
	f "github.com/VidarSolutions/Lib/Files"
	l "github.com/VidarSolutions/Lib/Log"
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
	md := NewMD()
	MyMap := md.NewMap("Test")

	fmt.Printf("Area Code for Country % is %s", Country[12].Country,[12].Code)
	fmt.Printf("My Map is blank Map Description is %s ", &MyMap["test".Description])

}
