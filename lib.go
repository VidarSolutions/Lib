package Lib

import (
	country "Country"

	d "github.com/VidarSoltuions/Lib/Dates"
	f "github.com/VidarSoltuions/Lib/Files"
	cat "github.com/VidarSolutions/Lib/Categories"
	country "github.com/VidarSolutions/Lib/Countries"
	l "github.com/VidarSolutions/Lib/Logs"
)

func main() {
	test()

}

func test() {
	//Palce holder for future funcitonality
	Category := cat.Categories.Category{}
	Country := country.CountryPhoneCodes
	f.EnsureDir("logs")
	currentDate := d.GetDate("yyymmdd")
	AppLog, _ := l.GetLogger("logs/app.log")
	md := NewMD()
	MyMap := md.NewMap("Test")

}
