package maps

import (
	"encoding/json"
	"os"
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
	Data map[string]MapData
}
func (mData MapData) NewMap(key string) *Map{
	m := NewMap()
	m.Set(key, mData)
	return m
}
func NewMap() *Map {
	return make(map[string]MapData)
}

func NewMD() *MapData{
	m =make(map[string]string)
	m[ID]=JsonStr
	dRoot := os.Getwd()
	Description :="Blank MapData"
	Mdtype := "MapData Replace with JsonStr to your Object or other Data Type"
	filePath	:= MDtype
	return md &MapData{DataRootDir: dRoot, Data: m, Description: Description, FilePath: filePath, Type : MDtype}
}


func NewMapData(MDtype, Description, filePath, droot, ID, JsonStr string) *Map{
	m =make(map[string]string)
	m[ID]=JsonStr
	return md &MapData{DataRootDir: dRoot, Data: m, Description: Description, FilePath: filePath, Type : MDtype}
}

func (m *Map) Get(key string) (*MapData, bool) {
	data, found := m.Data[key]
	return &data, found
}

func (m *Map) Set(key string, data MapData) error {
	// Save data in the map
	m.Data[key] = data
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