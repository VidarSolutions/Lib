package Maps_Test

import (
	"fmt"

	m "github.com/VidarSolutions/Lib/Maps"
)

func main() {
	// Create a new MapData instance
	config := m.MapData{
		DataRootDir: "/path/to/config",
		Data: map[string]string{
			"server_address": "127.0.0.1:8080",
			"database_url":   "postgres://user:password@localhost/mydb",
		},
		Description: "Application configuration",
		FilePath:    "/path/to/config/app.json",
		Type:        "configuration",
	}

	// Access and print the configuration data
	fmt.Println("Server Address:", config.Data["server_address"])
	fmt.Println("Database URL:", config.Data["database_url"])
	fmt.Println("Description:", config.Description)
	fmt.Println("File Path:", config.FilePath)
	fmt.Println("Type:", config.Type)
}
