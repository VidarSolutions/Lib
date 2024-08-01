package Files

import (
	"encoding/json"
	"fmt"
	"os"

	d "github.com/VidarSolutions/Lib/Dates"
	l "github.com/VidarSolutions/Lib/Log"
)

func EnsureDir(dirName string) (l.LogEntry, error) {
	currenDate = d.GetDate()
	mess := fmt.Sprintf("Check If Directory %s Exists \n", dirName)
	msg := l.LogEntry{Message: mess, Date: currentDate, LogLvl: l.INFO}

	if _, err := os.Stat(dirName); os.IsNotExist(err) {

		err := os.MkdirAll(dirName, 0755)
		if err != nil {
			msg.Message = fmt.Sprintf("Failed to create directory:%s %v \n", dirName, err)
		} else {
			msg.Message = Sprintf("Directory %s Did not Exist Creating Directory \n", dirName)
		}
	} else {
		msg.Message = Sprintf("Directory %s Found \n", dirName)
	}
	return msg, err
}

func WriteMApData(filePath string, data interface{}) error {
	// Convert data to JSON
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	// Write JSON data to the file
	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return err
	}
}
