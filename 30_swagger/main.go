package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/yaml.v3"
)

type Request struct {
	Value int    `json:"value"`
	Price int    `json:"price"`
	Desc  string `json:"desc"`
}

func logic() string {
	fmt.Println("logic")
	return "logic result"
}

func execute(value, price int) int {
	return value * price
}

func Execute(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	InputMessage := &Request{}
	if err := json.NewDecoder(r.Body).Decode(&InputMessage); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	result := make(map[string]int)
	result["result"] = execute(InputMessage.Value, InputMessage.Price)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func logicCall(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	result := logic()
	w.Write([]byte(result))
}

func jsonMake(r *Request) map[string]interface{} {
	resp, _ := json.Marshal(r)

	dct := make(map[string]interface{})
	err := json.Unmarshal(resp, &dct)
	if err != nil {
		fmt.Println(err)
	}

	return dct
}

func readYaml(dct map[string]interface{}) map[string]interface{} {
	yamlFile, err := ioutil.ReadFile("openapi/openapi.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	t := make(map[string]interface{})

	err = yaml.Unmarshal(yamlFile, &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	//t["components"].(map[string]interface{})["schemas"].(map[string]interface{})["Request"].(map[string]interface{})["properties"] = dct

	t["paths"].(map[string]interface{})["/execute"].
	(map[string]interface{})["post"].(map[string]interface{})["requestBody"].(map[string]interface{})["content"].
	(map[string]interface{})["application/json"].(map[string]interface{})["schema"].(map[string]interface{})["example"] = dct
	return t
}

func writeYml(dct map[string]interface{}) {
	var b bytes.Buffer
	yamlEncoder := yaml.NewEncoder(&b)
	yamlEncoder.SetIndent(2)
	yamlEncoder.Encode(&dct)
	ioutil.WriteFile("openapi/openapi.yaml", b.Bytes(), 0664)
}

func main() {
	//r := &Request{}
	//result := jsonMake(r)
	//result = readYaml(result)
	//writeYml(result)
	//time.Sleep(time.Second*2)

	router := httprouter.New()
	router.GET("/logic", logicCall)
	router.POST("/execute", Execute)

	// Get spec for swagger-ui
	router.GET("/openapi.yaml", openapiHandler)

	// Static file for swagger-ui
	router.ServeFiles("/swagger/*filepath", NewSwaggerHandler())

	log.Fatal(http.ListenAndServe(":8000", router))
}
