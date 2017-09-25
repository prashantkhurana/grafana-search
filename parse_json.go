package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type panel struct {
	Datasource string
	Title      string
	Id         int
}
type row struct {
	Title  string
	Panels []panel
}
type Dashboard struct {
	Title string
	Tags  []string
	Rows  []row
}

var FinalMap = make(map[string][]string)
var dashboardList []Dashboard

func main() {

	files, _ := ioutil.ReadDir("backup_grafana_db/")

	// dashboard := Dashboard{}
	// str2, _ := ioutil.ReadFile("backup_grafana_db/" + "bidder-09_22_2017.json")
	// json.Unmarshal(str2, &dashboard)
	// ccc, _ := json.Marshal(dashboard)
	// s := string(ccc[:])
	// fmt.Println(s)

	for _, f := range files {
		//fmt.Println(f.Name())
		dashboard := Dashboard{}
		str2, _ := ioutil.ReadFile("backup_grafana_db/" + f.Name())
		json.Unmarshal(str2, &dashboard)
		if !strings.Contains(f.Name(), "monitoring") {
			dashboardList = append(dashboardList, dashboard)
		}
		//ccc, _ := json.Marshal(dashboard)
		//s := string(ccc[:])
		//fmt.Println(s)
		// dbName := res2["title"].(string)
		// res3 := map[string]bool{}
		// parseMap(res2, res3)
		// var keys []string
		// for key, _ := range res3 {
		// 	keys = append(keys, key)
		// }
		// FinalMap[dbName] = keys
	}

	file, err := os.Create("result.txt")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	ccc, _ := json.Marshal(dashboardList)

	err = ioutil.WriteFile("output.json", ccc, 0644)
	// s := string(dashboardList[:])
	//fmt.Fprintf(file, s)

	// countMap := map[string]int{}

	// for _, v := range FinalMap {
	// 	for _, str := range v {
	// 		countMap[str]++
	// 	}
	// }

	// for k, v := range FinalMap {
	// 	for _, str := range v {
	// 		if countMap[str] < 50 && str != "" {
	// 			if _, err := strconv.Atoi(str); err != nil {
	// 				FinalMap2[k] = append(FinalMap2[k], str)
	// 			}
	// 		}
	// 	}
	// }

	// type kv struct {
	// 	Key   string
	// 	Value int
	// }

	// var ss []kv
	// for k, v := range countMap {
	// 	ss = append(ss, kv{k, v})
	// }

	// fmt.Println(len(ss))
	// sort.Slice(ss, func(i, j int) bool {
	// 	return ss[i].Value > ss[j].Value
	// })

	// // for _, kv := range ss {
	// // 	fmt.Println(kv.Key, kv.Value)
	// // }

	// x, _ := json.Marshal(ss)
	// fmt.Println(string(x))

	// //fmt.Println(string(x))
	// // for k, v := range FinalMap {
	// // 	fmt.Println(k, v)
	// // 	fmt.Println()
	// // 	fmt.Println()
	// // 	fmt.Println()
	// // 	fmt.Println()
	// // 	fmt.Println()
	// // 	fmt.Println()

	// // }

	// // name := "bidder-09_22_2017.json"
	// // str2, _ := ioutil.ReadFile("/Users/PKhurana/code/grafana/backup_grafana_db/" + name)
	// // json.Unmarshal(str2, &res2)

	// // dbName := res2["title"]
	// // parseMap(res2, dbName.(string))

	// // fmt.Println(FinalMap)
	http.HandleFunc("/getGrafanaData", handler)
	http.ListenAndServe(":8082", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	x, _ := json.Marshal(dashboardList)
	fmt.Println(string(x))
	fmt.Fprintf(w, "%s", string(x))
}

func parseMap(aMap map[string]interface{}, dbName map[string]bool) {
	for key, val := range aMap {
		if key == "title" || key == "measurement" {
			//fmt.Println(key, ":", val)
			if validString(val.(string)) {

			}
			if !(strings.Contains(val.(string), "$") || strings.Contains(val.(string), "Dashboard Row") || strings.Contains(val.(string), "Row")) {
				dbName[val.(string)] = true
				//FinalMap[dbName] = append(FinalMap[dbName], val.(string))
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

func validString(str string) bool {
	if !(strings.Contains(str, "$") || strings.Contains(str, "Dashboard Row") || strings.Contains(str, "Row")) {
		return true //FinalMap[dbName] = append(FinalMap[dbName], val.(string))
	}
	return false
}

func parseArray(anArray []interface{}, dbName map[string]bool) {
	for _, val := range anArray {
		switch val.(type) {
		case map[string]interface{}:
			parseMap(val.(map[string]interface{}), dbName)
		case []interface{}:
			parseArray(val.([]interface{}), dbName)
		case string:
			dbName[val.(string)] = true
		default:
		}
	}
}
