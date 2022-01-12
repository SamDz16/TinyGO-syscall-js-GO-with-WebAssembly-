package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"syscall/js"
)

// func add() js.Func {
// 	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
// 		if len(args) == 2 {
// 			return args[0].Int() + args[1].Int()
// 		}
// 		return -1
// 	})
// }

// func capitalized() js.Func {
// 	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
// 		if len(args) == 1 {
// 			return strings.ToUpper(args[0].String())
// 		}
// 		return -1
// 	})
// }

func fetchData() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// if len(args) == 2 {
			// endpoint := args[0].String()
			// query := args[1].String()

			endpoint := "https://dbpedia.org/sparql"
			query := "select distinct ?Concept where {[] a ?Concept} LIMIT 100"

			data := url.Values{}
			data.Set("query", query)

			client := &http.Client{}
			r, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode())) // URL-encoded payload
			if err != nil {
				log.Fatal(err)
			}
			r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

			res, err := client.Do(r)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(res.Status)
			defer res.Body.Close()
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(string(body))
			return string(body)
		// }
		// return "This function requiers two arguments"
	})
}

// func fetchData() js.Func {
// 	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
// 		if len(args) == 2 {
// 			endpoint := args[0].String()
// 			query := args[1].String()

// 			data := url.Values{}
// 			data.Set("query", query)

// 			client := &http.Client{}
// 			r, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode())) // URL-encoded payload
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
// 			r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

// 			res, err := client.Do(r)
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			log.Println(res.Status)
// 			defer res.Body.Close()
// 			body, err := ioutil.ReadAll(res.Body)
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			log.Println(string(body))
// 			return string(body)
// 		}
// 		return "This function requiers two arguments"
// 	})
// }

func main() {
	fmt.Printf("Hello Web Assembly from Go!\n")

	// js.Global().Set("add", add())
	// js.Global().Set("capitalized", capitalized())
	js.Global().Set("fetchData", fetchData())
	// js.Global().Set("anotherFunc", anotherFunc())

	<-make(chan struct{}, 0)
}
