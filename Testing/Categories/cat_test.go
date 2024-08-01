package Categories_Test

import (
	"fmt"

	t "github.com/VidarSolutions/Lib/Testing"
)

func TestCat() {
	// Example usage
	cat, err := NewCategory("DragonSites", "test")
	results(cat, err)
	cat, err = NewCategory("DragonSites", "test")
	results(cat, err)
	cat, err = NewCategory("DragonSites", "test")
	results(cat, err)

}

func results(cat Category, err error) {
	if err != nil {
		fmt.Println(t.CheckError(err))
	} else {
		fmt.Printf(" Name : %s Id %", cat.Name, cat.ID)
	}
}
