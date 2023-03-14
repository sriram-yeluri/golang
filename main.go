package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Simple_Api_Call(){
	resp, err := http.Get("http://example.com/")
	if err != nil{
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body , _ := ioutil.ReadAll(resp.Body)
	log.Print(string(body))
}

func Api_Call_With_Transport() {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("http://example.com/")

	if err != nil {
		log.Print(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	log.Print(string(body))
}

func JSONPlaceholder() {
	url := "https://jsonplaceholder.typicode.com/users"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Print(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Print(err)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

func main() {
	fmt.Println("Learning golang programming ..... ")
	// GetVariables()
	// var d = data{"MyName"}
	// d2 := data{"MyNewName"}

	// d.GetData()
	// d2.GetData()

	// var s1 = student{1, "firstStudent"}
	// s2 := student{2, "secondStudent"}

	// s1.DisplayStudent()
	// s2.DisplayStudent()

	// DisplayMaps()
	// m := SetMap()
	// for k, v := range m {
	// 	fmt.Println(k, ":", v)
	// }

	Simple_Api_Call()
	Api_Call_With_Transport()
	// JSONPlaceholder()

}
