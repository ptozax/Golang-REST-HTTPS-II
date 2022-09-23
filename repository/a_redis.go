package repository

import (
	"encoding/json"
	"https/driver"
	"time"
)

//XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX  POST & GET TOKEN   XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
func POST_Redis_String(id string, data string, duration int) error {
	client, err := driver.RedisConnect()
	if err != nil {
		client.Close()
		return err
	}

	durationTime := time.Minute * time.Duration(duration)
	err = client.Set(id, data, durationTime).Err()
	if err != nil {
		client.Close()
		return err
	}
	client.Close()
	return nil
}

func GET_Redis_String(id string) (string, error) {
	client, err := driver.RedisConnect()
	if err != nil {
		client.Close()
		return "", err
	}

	val, err := client.Get(id).Result()
	if err != nil {
		client.Close()
		return "", err
	}

	client.Close()
	return val, nil
}
func DEL_Redis_String(id string) error {
	client, err := driver.RedisConnect()
	if err != nil {
		client.Close()
		return err
	}
	client.Del(id)
	client.Close()
	return nil
}

//XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX  POST & GET JSON   XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX

func POST_Redis_Index_Interface(id string, data interface{}, duration int, redisIndex int) error {

	client, err := driver.RedisConnect_ID(redisIndex)
	if err != nil {
		client.Close()
		return err
	}

	durationTime := time.Minute * time.Duration(duration)
	arrByte, err := json.Marshal(data)
	if err != nil {
		client.Close()
		return err
	}
	client.Set(id, arrByte, durationTime)

	client.Close()
	return nil
}

func GET_Redis_Index_Interface(id string, data interface{}, redisIndex int) error {
	client, err := driver.RedisConnect_ID(redisIndex)
	if err != nil {
		client.Close()
		return err
	}
	stringVal, err := client.Get(id).Result()
	if err != nil {
		client.Close()
		return err
	}

	json.Unmarshal([]byte(stringVal), &data)

	client.Close()
	return nil
}

func DEL_Redis_Index(id string, redisIndex int) error {
	client, err := driver.RedisConnect_ID(redisIndex)
	if err != nil {
		client.Close()
		return err
	}
	client.Del(id)
	client.Close()
	return nil
}
