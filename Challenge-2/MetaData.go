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

// main method started

func main() {
	var info map[string]interface{}
	flag.Parse()
	url := "http://192.168.1.100/latest/meta-data/"
	
	//get metadata using fetchMetaData method
	
	metaData := fetchMetaData(url)
	cont, err := json.Marshal(metaData)
	if err != nil {
		log.Fatalf(format:"Error: %v\n", err)
		os.Exit(1)
	}
	if len(os.Args) == 1 {
		os.Stdout.Write(cont)
		fmt.Println()
	}
	if len(os.Args) == 2 {
		json.Unmarshal([]byte(cont),&info)
		gettingArgs := os.Args[1] 
		fmt.Println(info[gettingArgs])
	}

}