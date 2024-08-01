package Countries_test

import (
	"fmt"

	c "github.com/Vidarsolutions/Lib/Countries"
)

func Country() {
	c.PrintAll()
	fmt.Printf("Phone Area Code %s is for Country %s", c.CountryPhoneCodes[12].Country, c.CountryPhoneCodes[12].Code)

}
