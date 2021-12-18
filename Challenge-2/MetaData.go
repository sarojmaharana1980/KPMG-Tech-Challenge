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
	cont, err := json.Marshal(metaData)
	if err != nil {
		log.Fatalf(format:"Error: %v\n", err)
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