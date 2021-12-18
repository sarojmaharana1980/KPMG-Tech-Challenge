package main

import "encoding/json"
import "fmt"
import "log"
import "bufio"
import "net/http"
import "os"
import "flag"
import "strings"

type MetaData interface{}

// Function retriveData retrives data from passed url
// Arguements : url
// Return type : Array of Strings 

func retriveData(url string) []string {
	response, error := http.Get(url)
	if error != nil {
		log.Fatalf(format:"erroror: %v\n", error)
		os.Exit(1)
	}
	defer response.Body.Close()
	reader := bufio.NewReader(response.Body)
	data := make([]string, 0)
	for {
		line, error := reader.ReadString('\n')
		data = append(data, strings.TrimRight(line, "\n"))
		if error != nil {
			break
		}
	}
	return data
}



// Function getMetaData retrives meteadata from passed url
// Arguements : url
// Return type : Metadata 

func getMetaData(url string) MetaData {
	metaData := make(map[string]interface{})
	data := retriveData(url)
	for _, line := range data {
		switch {
		default:
			metaData[line] = retriveData(url + line)[0]
		case line == "metrics/":
			break
		case strings.HasSuffix(line, "/"):
			metaData[line[:len(line)-1]] = getMetaData(url + line)
		case strings.HasSuffix(url, "public-keys/"):
			keyId := strings.SplitN(line, "=", 2)[0]
			metaData[line] = retriveData(url + keyId + "/openssh-key")[0]
		}
	}
	return metaData
}

//main method started

func main() {
	var info map[string]interface{}
	flag.Parse()
	url := "http://192.168.1.100/latest/meta-data/"
	
	//get metadata using getMetaData method
	
	metaData := getMetaData(url)
	cont, error := json.Marshal(metaData)
	if error != nil {
		log.Fatalf(format:"erroror: %v\n", error)
		os.Exit(1)
	}
	if len(os.Args) == 1 {
	    // write the output
		os.Stdout.Write(cont)
		fmt.Println()
	}
	if len(os.Args) == 2 {
	    // Started formatiing json format
		json.Unmarshal([]byte(cont),&info)
		gettingArgs := os.Args[1] 
		fmt.Println(info[gettingArgs])
	}

}