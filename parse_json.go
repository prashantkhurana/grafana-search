package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var FinalMap = make(map[string][]string)

func main() {

	files, _ := ioutil.ReadDir("/Users/PKhurana/code/grafana/backup_grafana_db/")

	for _, f := range files {
		fmt.Println(f.Name())
		res2 := map[string]interface{}{}
		str2, _ := ioutil.ReadFile("/Users/PKhurana/code/grafana/backup_grafana_db/" + f.Name())
		json.Unmarshal(str2, &res2)
		dbName := res2["title"]
		parseMap(res2, dbName.(string))
	}

	for k, v := range FinalMap {
		fmt.Println(k, v)
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()

	}

	// name := "bidder-09_22_2017.json"
	// str2, _ := ioutil.ReadFile("/Users/PKhurana/code/grafana/backup_grafana_db/" + name)
	// json.Unmarshal(str2, &res2)

	// dbName := res2["title"]
	// parseMap(res2, dbName.(string))

	// fmt.Println(FinalMap)
}

func parseMap(aMap map[string]interface{}, dbName string) {
	for key, val := range aMap {
		if key == "title" {
			//fmt.Println(key, ":", val)
			FinalMap[dbName] = append(FinalMap[dbName], val.(string))
		}
		switch val.(type) {
		case map[string]interface{}:
			// fmt.Println(key)
			parseMap(val.(map[string]interface{}), dbName)
		case []interface{}:
			parseArray(val.([]interface{}), dbName)
		default:
		}
	}
}

func parseArray(anArray []interface{}, dbName string) {
	for _, val := range anArray {
		switch val.(type) {
		case map[string]interface{}:
			parseMap(val.(map[string]interface{}), dbName)
		case []interface{}:
			parseArray(val.([]interface{}), dbName)
		default:
		}
	}
}
