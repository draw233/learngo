package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"jsondemo/model"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var (
	urlx = "https://api.juejin.cn/interact_api/v1/pin_tab_lead?aid=2608&uuid=7347226625274201636&spider=0"
)

func main() {
	fmt.Println("1")
	HttpPostFormDemo()
}

func HttpPostFormDemo() {
	var param = url.Values{}
	param.Set("name", "你")
	param.Set("password", "pass")
	xxx := param.Encode()
	fmt.Println(xxx)
	resp, err := http.Post("http://127.0.0.1:4523/m1/5825518-5510942-default/demo/login?apifoxApiId=258519976", "application/x-www-form-urlencoded", strings.NewReader(xxx))
	if err != nil {
		log.Fatal("xx")
	}
	defer resp.Body.Close()

	xx, _ := io.ReadAll(resp.Body)
	fmt.Println(string(xx))
}

func HttpPostJsonDemo() {

	resp, err := http.Post("http://127.0.0.1:4523/m1/5825518-5510942-default/demo/login?apifoxApiId=258520620", "application/json", bytes.NewReader([]byte(`{"name":"a","passwor":"xx"}`)))
	if err != nil {
		log.Fatal("xx")
	}
	defer resp.Body.Close()

	xx, _ := io.ReadAll(resp.Body)
	fmt.Println(string(xx))
}

func HttpGetDemo() {
	req, err := http.NewRequest(http.MethodGet, urlx, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("h1", "v1")
	var data model.Resp

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	// parse to json
	json.NewDecoder(resp.Body).Decode(&data)
	fmt.Println(data)
}

func chanDemo1() {
	ch := make(chan string, 2)
	go send(ch)
	for {
		select {
		case data, open := <-ch:
			if open {
				fmt.Println(data)
			} else {
				return
			}
		}
	}
}

func send(ch chan string) {
	for i := 0; i < 10; i++ {
		ch <- strconv.Itoa(i)
	}
	// close chan
	close(ch)
}

func chanDemo2() {
	ch := make(chan string, 2)
	go send(ch)
	for x := range ch {
		fmt.Println(x)
	}
}

func timeClickDemo() {
	context, cc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cc()
	newFunction(context)
	<-context.Done()
	fmt.Println("time out ")

}

func newFunction(context context.Context) {
	ticker := time.NewTicker(time.Second)
	now := time.Now()
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			fmt.Println("do work ", time.Since(now))
		case <-context.Done():
			fmt.Println("done")
			return
		}
	}
}

func DefaultValue() {
	var i int
	fmt.Println(i)
	var str string
	fmt.Println(str)
	var arr []int
	fmt.Println(arr)
	var s struct{}
	fmt.Println(s)
	var b bool
	fmt.Println(b)
	var it interface{}
	fmt.Println(it)

}
