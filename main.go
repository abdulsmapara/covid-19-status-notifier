package main

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"strings"
	"encoding/json"
	"io/ioutil"
	"os"
	"bytes"
	"net/http"
	"github.com/olekukonko/tablewriter"
	"log"

)



func getData() ([]string, [][]string, []string){
	resp, err := soup.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	doc := soup.HTMLParse(resp)
	tableData := doc.FindAll("table")
	header := []string{"S.No.","State/UT","In","Fr","Cd","Dt"}
	final_row := []string{}
	row := [][]string{}

	for _, children := range(tableData[7].Children()) {	
		// Headers here
		for _, headChildren := range(children.Children()) {
			count_data := 0

			data := []string{}
			if headChildren.NodeValue == "tr" {
				for _, rowChildren := range(headChildren.Children()) {
					if rowChildren.NodeValue == "td" {
						count_data += 1
						data = append(data,strings.Replace(rowChildren.FullText(),"\n","",2))
					}
				}
				if count_data == 5 {
					for _,val := range data { 
						final_row = append(final_row, strings.Replace(val," ","",100))
					}
				} else if count_data > 0 {
					row = append(row, data)
					
				}
			}
		}

	}
	return header, row, final_row
}
func main() {


	logFile, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer logFile.Close()
    log.SetOutput(logFile)
    log.Println("Starting the program")
	
	defer func() {
		if r := recover(); r != nil {
			log.Fatal(r)
		}
	}()
	// Fetch data
	header, row, final_row := getData()
	
	log.Println("Data fetched from URL")

	// Prepare data
	rowData := make(map[string][]string)
	for _, rowVal := range row {
		rowData[rowVal[1]] = rowVal[2:]
	}

	log.Println("Fetched data has been prepared")


	// read previous data
	jsonFile, err := os.Open("data.json")
		
	if err != nil {
		log.Fatal(err)
	}

	defer jsonFile.Close()

    byteValue, err := ioutil.ReadAll(jsonFile)

    if err != nil {
    	log.Fatal(err)
    }

    var pastData map[string][]string
    json.Unmarshal(byteValue, &pastData)

    changed := false
    info := ""
    for state, values := range rowData {
    	if oldValue, exists := pastData[state]; exists {
    		if !testEquality(oldValue,values) {
    			changed = true
    			info += "State/UT " + state + " changed from " + strings.Join(oldValue,",") + " to " + strings.Join(values,",") + "\n"
    		}
    	} else {
    		changed = true	
    		info += "State/UT " + state + " added -> " + strings.Join(values,",") + "\n"
    	}
    }
    if changed {
    	log.Println("The data fetched has changed")
    	log.Println("INFO: " + info)
    	file, _ := json.Marshal(rowData)
    	_ = ioutil.WriteFile("data.json", file, 0644)	
	    url := WEBHOOK_URL
	    
	    buf := new(bytes.Buffer)
	    table := tablewriter.NewWriter(buf)
	    
	    table.SetHeader(header)

	    for _, v := range row {
	        table.Append(v)
	    }
		final_row = append([]string{""}, final_row...)    
		table.Append(final_row)
	    
	    table.SetAlignment(tablewriter.ALIGN_LEFT)

	    table.Render()

	    data := map[string]string{
	    	"text" : info + "```" + buf.String() + "```",
	    	"username": "virus-tracker",
	    }

	    js, _ := json.Marshal(data)
	    
	    payload := strings.NewReader(string(js))

		req, _ := http.NewRequest("POST", url, payload)

		req.Header.Add("content-type", "application/json")
		req.Header.Add("cache-control", "no-cache")

		res, err := http.DefaultClient.Do(req)

		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

		fmt.Println(string(body))
    } else {
    	log.Println("The data fetched has not changed")
    }


}