package utils

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

type LoadInstance struct {
	//LoadId           string // Unique identifier for the load
	RatingLoadID     string
	AccountingLoadID string
	//TariffPlanID     string    // Tariff plan identificator for the data loaded
	LoadTime time.Time // Time of load
}

type CacheFileInfo struct {
	Encoding string
	LoadInfo *LoadInstance
}

func LoadCacheFileInfo(path string) (*CacheFileInfo, error) {
	// open data file
	dataFile, err := os.Open(filepath.Join(path, "cache.info"))
	defer dataFile.Close()
	if err != nil {
		Logger.Err("<cache decoder>: " + err.Error())
		return nil, err
	}

	filesInfo := &CacheFileInfo{}
	dataDecoder := json.NewDecoder(dataFile)
	err = dataDecoder.Decode(filesInfo)
	if err != nil {
		Logger.Err("<cache decoder>: " + err.Error())
		return nil, err
	}
	return filesInfo, nil
}

func SaveCacheFileInfo(path string, cfi *CacheFileInfo) error {
	// open data file
	// create a the path
	if err := os.MkdirAll(path, 0766); err != nil {
		Logger.Err("<cache encoder>:" + err.Error())
		return err
	}

	dataFile, err := os.Create(filepath.Join(path, "cache.info"))
	defer dataFile.Close()
	if err != nil {
		Logger.Err("<cache encoder>:" + err.Error())
		return err
	}

	// serialize the data
	dataEncoder := json.NewEncoder(dataFile)
	if err := dataEncoder.Encode(cfi); err != nil {
		Logger.Err("<cache encoder>:" + err.Error())
		return err
	}
	return nil
}

func CacheFileExists(filePath string) bool {
	if _, err := os.Stat(filePath); err == nil {
		return true
	}
	return false
}