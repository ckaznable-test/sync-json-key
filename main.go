package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type JsonFile struct {
	path string
	data map[string]interface{}
}

func main() {
	data, err := getAllJsonFile("./test")
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, fileData := range data {
		data := changeKeyWithJson(fileData.data, "test", "test4")

		updatedData, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			continue
		}

		fmt.Printf("%v %v\n", fileData.path, string(updatedData))
	}
}

func getAllJsonFile(dir string) ([]JsonFile, error) {
	var jsonFiles []JsonFile

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 檢查是否為 JSON 檔案
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".json") {
			// 讀取JSON檔案
			file, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			// 解析JSON
			var data map[string]interface{}
			err = json.Unmarshal(file, &data)
			if err != nil {
				return err
			}

			jsonFiles = append(jsonFiles, JsonFile{path, data})
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return jsonFiles, nil
}

func changeKeyWithJson(obj map[string]interface{}, oKey string, tKey string) map[string]interface{} {
	tmp := obj[oKey]
	delete(obj, oKey)
	obj[tKey] = tmp
	return obj
}
