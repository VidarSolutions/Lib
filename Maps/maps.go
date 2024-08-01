package maps

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	f "github.com/VidarSolutions/Lib/Files"
)

var (
	m  = NewMap()
	mu sync.Mutex
)

type MapData struct {
	DataRootDir string
	Data        map[string]string
	Description string
	FilePath    string
	Type        string
}

type Map struct {
	store map[string]MapData
}

func NewMap() *Map {
	return &Map{store: make(map[string]MapData)}
}
func (m *Map) Get(key string) (*MapData, bool) {
	data, found := m.store[key]
	return &data, found
}

func (m *Map) Set(key string, data MapData) error {
	// Save data in the map
	m.store[key] = data
	var dirPath, folder, filePath string
	// Get the current working directory
	if data.FilePath != "" {
		if data.DataRootDir !""{
			appRootdir = data.DataRootDir
		}else{
			appRootdir, err := os.Getwd()
			if err != nil {
				return err
			}
		}
		if data.Type !=""{
			folder = data.Type
		}else{
			folder = key
		}

		// Create the file path
		dirPath = filepath.Join(appRootdir, folder)
		filePath = filepath.Join(dirPath, key)
	} else {
		filePath = data.FilePath
	}
	// Create directory if it doesn't exist
	f.EnsureDir(dirPath)

	f.WriteMapData(filePath, data)

	return nil
}