package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)
func getRightPort(apiPort int, apiUrl string) string {
	var key string
	 for i := 3000; i < 4001 ; i++ {
		url := fmt.Sprintf("%s:%d",apiUrl, i)
		response, err := http.Get(url)
		if err == nil {
			defer response.Body.Close()
			body, err := io.ReadAll(response.Body)
			if err != nil {
				fmt.Println(err)
			}
			str:= string(body)
			splitStr:= strings.Split(str, ": ")
			key = splitStr[1]
			break;
		}else{
			fmt.Println(err)
		}
	}
	return key
}

func postURL(apiUrl string, secretKey string, apiPort2 int, filepath string)  string {
	var urlKey string
	url := fmt.Sprintf("%s:%d",apiUrl,apiPort2)
	var jsonStr = []byte(`{}`)
	resp, err := http.Post(url+"/?secretKey="+secretKey, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Printf("Error : %v\n", err)
	}
	if resp.Body == nil {
		fmt.Println("resp body nil")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
	str:= string(body)
	splitStr:= strings.Split(str, ": ")
	urlKey = splitStr[1]
	fileContent(apiUrl)
	return urlKey

}
func fileContent(apiUrl string) {
	file, err := os.Open("finalResult.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	content, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	port:= content[0]
	finalKey := content[2]

	lastPort, err:= strconv.Atoi(port[0])
	if err != nil {
		fmt.Println("Error during conversion")
		return
	}
	finalStep(apiUrl, lastPort, finalKey)

}
func finalStep(apiUrl string, lastPort int , lastKey []string){
	url := fmt.Sprintf("%s:%d",apiUrl, lastPort)
	fmt.Println(url)
	var jsonStr = []byte(`{}`)
	resp, err := http.Post(url+"/?finalKey="+ lastKey[0], "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Printf("Error : %v\n", err)
	}
	if resp.Body == nil {
		fmt.Println("resp body nil")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}

