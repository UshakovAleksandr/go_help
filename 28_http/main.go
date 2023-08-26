package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetReq() {
	resp, err := http.Get("https://api-blog-flask.herokuapp.com/users/26")
	if err != nil {
		log.Fatal(err)
	}
	// закрываем для того, что бы освободить текущий connect и переиспользовать его. Не открывать новый.
	defer resp.Body.Close() // <-- Зачем?

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	data := make(map[string]interface{})

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data)
}

func PostReq() {
	message := map[string]string{
		"password": "111",
		"username": "ushakov2",
	}

	payload, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("https://api-blog-flask.herokuapp.com/users", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	data := make(map[string]interface{})

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data)
}

func CustomReqGet() {
	request, err := http.NewRequest("GET", "https://api-blog-flask.herokuapp.com/users/26", nil)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)
	fmt.Println(result)
}

func CustomReqGPost() {
	errUnexpectedHTTPStatus := errors.New("not 200")
	errUnexpectedContentType := errors.New("no app/json")

	message := map[string]string{
		"password": "111",
		"username": "ushakov3",
	}

	payload, _ := json.Marshal(&message)

	req, err := http.NewRequest("POST", "https://api-blog-flask.herokuapp.com/users", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println(fmt.Errorf("%w: %s", errUnexpectedHTTPStatus, resp.Status))
	}
	ct := resp.Header.Get("Content-Type")
	if ct != "application/json" {
		fmt.Println(fmt.Errorf("%w: %s", errUnexpectedContentType, ct))
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)
	fmt.Println(result)
}

func fileOpen(path string) []byte {
	data, err := ioutil.ReadFile(fmt.Sprintf("jsons/%s.json", path))
	if err != nil {
		log.Fatalln(err)
	}

	return data
}

func CallWithCert() {
	// Read the key pair to create certificate
	cert, err := tls.LoadX509KeyPair("client_cert.pem", "client_key.pem")
	if err != nil {
		log.Fatal(err)
	}
	// Create a CA certificate pool and add cert.pem to it
	caCert, err := ioutil.ReadFile("certs/ca.crt")
	if err != nil {
		log.Fatal(err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Create a HTTPS client and supply the created CA pool and certificate
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      caCertPool,
				Certificates: []tls.Certificate{cert},
			},
		},
	}

	resp, err := client.Post("https://sm9.sh5.dev-gen2.delta.sbrf.ru/execute", "application/json", bytes.NewBuffer(fileOpen("path")))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	fmt.Println(resp.StatusCode)
	fmt.Println(result)
}

func ReqWithCtx() {
	//client := http.Client{}
	//ctx := context.Background()
	//ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	//defer cancel()
	//req = req.WithContext(ctx)
	//resp, err := client.Do(req)
}

func main() {
	//GetReq()
	//PostReq()
	//CustomReqGet()
	//CustomReqGPost()
}
