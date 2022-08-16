package fs

import (
	"log"
	"os"
	"io/ioutil"
)

func CreateDir(folderPath string) error {
	if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
		return err
	}

	return nil
}

func ReadFile(fileName string) ([]byte, error){
	f, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("read file %s error: %s", fileName, err.Error())
	}

	return f, nil
}

func WriteFile(fileName string, data []byte) error {
	if err := ioutil.WriteFile(fileName, data, 0644); err != nil {
		log.Fatalf("write file %s error: %s", fileName, err.Error())
	}

	return nil
}