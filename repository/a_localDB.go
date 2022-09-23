package repository

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func GET_localDB(filename string, data interface{}) error {
	osFile, err := os.Open("./localDB/" + filename + ".json") // open file
	defer osFile.Close()                                      // close file for can read or edit
	byteItem, _ := ioutil.ReadAll(osFile)                     // os.file to []byte
	json.Unmarshal(byteItem, data)                            // []byte to struct
	if err != nil {
		return err
	}
	return nil
}

func POST_localDB(filename string, data interface{}) error {
	byteDataTosave, _ := json.MarshalIndent(data, "", " ")
	err := os.WriteFile("./localDB/"+filename+".json", byteDataTosave, 0755)
	if err != nil {
		return err
	}
	return nil
}
