package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type selectObject struct {
	Params []string
}

type target struct {
	Measurement string
	Select      [][]selectObject
}

type yaxes struct {
	Label string
}

type xaxis struct {
	Name string
}
type seriesOverride struct {
	alias string
}
type panels struct {
	Datasource      string
	Title           string
	Targets         []target
	Yaxes           []yaxes
	Xaxis           xaxis
	SeriesOverrides []seriesOverride
}

type row struct {
	Title  string
	Panels []panels
}

type dashboardResponse struct {
	Title    string
	Timezone string
	Tags     []string
	Rows     []row
}

func main() {

	res := dashboardResponse{}
	str, err := ioutil.ReadFile("/Users/PKhurana/code/grafana/backup_grafana_db/bidder-09_22_2017.json")
	check(err)
	//fmt.Print(string(str2))

	json.Unmarshal(str, &res)
	//fmt.Printf("%+v\n", res)

	res2 := map[string]interface{}{}

	str2, _ := ioutil.ReadFile("/Users/PKhurana/code/grafana/backup_grafana_db/bidder-09_22_2017.json")
	json.Unmarshal(str2, &res2)
	//fmt.Printf("%+v\n", res2)

	parseMap(res2)

}

func parseMap(aMap map[string]interface{}) {
	for key, val := range aMap {
		if key == "title" {
			fmt.Println(key, ":", val)
		}
		switch val.(type) {
		case map[string]interface{}:
			// fmt.Println(key)
			parseMap(val.(map[string]interface{}))
		case []interface{}:
			parseArray(val.([]interface{}))
		default:
		}
	}
}

func parseArray(anArray []interface{}) {
	for _, val := range anArray {
		switch val.(type) {
		case map[string]interface{}:
			parseMap(val.(map[string]interface{}))
		case []interface{}:
			parseArray(val.([]interface{}))
		default:
		}
	}
}
