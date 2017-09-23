package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var FinalMap = make(map[string][]string)

func main() {

	files, _ := ioutil.ReadDir("/Users/PKhurana/code/grafana/backup_grafana_db/")

	for _, f := range files {
		//fmt.Println(f.Name())
		res2 := map[string]interface{}{}
		str2, _ := ioutil.ReadFile("/Users/PKhurana/code/grafana/backup_grafana_db/" + f.Name())
		json.Unmarshal(str2, &res2)
		dbName := res2["title"]
		parseMap(res2, dbName.(string))
	}

	//fmt.Println(string(x))
	// for k, v := range FinalMap {
	// 	fmt.Println(k, v)
	// 	fmt.Println()
	// 	fmt.Println()
	// 	fmt.Println()
	// 	fmt.Println()
	// 	fmt.Println()
	// 	fmt.Println()

	// }

	// name := "bidder-09_22_2017.json"
	// str2, _ := ioutil.ReadFile("/Users/PKhurana/code/grafana/backup_grafana_db/" + name)
	// json.Unmarshal(str2, &res2)

	// dbName := res2["title"]
	// parseMap(res2, dbName.(string))

	// fmt.Println(FinalMap)
	http.HandleFunc("/getGrafanaData", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	x, _ := json.Marshal(FinalMap)
	fmt.Println(string(x))
	fmt.Fprintf(w, string(x))
}

func parseMap(aMap map[string]interface{}, dbName string) {
	for key, val := range aMap {
		if key == "title" || key == "measurement" {
			//fmt.Println(key, ":", val)
			if !strings.Contains(val.(string), "$") {
				FinalMap[dbName] = append(FinalMap[dbName], val.(string))
			}
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
