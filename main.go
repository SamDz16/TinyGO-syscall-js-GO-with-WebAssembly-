package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"syscall/js"
)

func add() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 2 {
			return args[0].Int() + args[1].Int()
		}
		return -1
	})
}

func capitalized() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 1 {
			return strings.ToUpper(args[0].String())
		}
		return -1
	})
}

func anotherFunc() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return "hhh"
	})
}

func fetchData() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 2 {
			APIUrl := args[0].String()
			query := args[1].String()

			data := url.Values{
				"query": {query},
			}

			u, _ := url.ParseRequestURI(APIUrl)
			urlStr := u.String()

			client := &http.Client{}
			r, _ := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(data.Encode()))
			r.Header.Add("Content-type", "application/x-www-form-urlencoded")

			resp, _ := client.Do(r)

			// Read the response body
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatalln(err)
			}

			// defer resp.Body.Close()

			// Convert body to type string
			return string(body)
		}
		return "This function requiers two arguments"
	})
}

func main() {
	fmt.Printf("Hello Web Assembly from Go!\n")

	js.Global().Set("add", add())
	js.Global().Set("capitalized", capitalized())
	// js.Global().Set("fetchData", fetchData())
	js.Global().Set("anotherFunc", anotherFunc())

	<-make(chan struct{}, 0)
}
