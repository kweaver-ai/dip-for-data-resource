package main

import (
	"bufio"
	"bytes"
	//"encoding/base64"
	"fmt"
	"github.com/samber/lo"
	//"io"
	"encoding/json"
	"net/http"
	"strings"
)

//	func main() {
//		req, _ := http.NewRequest("GET", "http://127.0.0.1:8181/api/internal/af-sailor-service/v1/assistant/sse", nil)
//		req.Header.Set("Accept", "text/event-stream")
//		//req.Header.Set("Cache-Control", "no-cache")
//		//req.Header.Set("Connection", "keep-alive")
//		//req.Header.Set("Access-Control-Allow-Origin", "*")
//		client := &http.Client{}
//
//		res, _ := client.Do(req)
//		defer res.Body.Close()
//
//		reader := bufio.NewReader(res.Body)
//
//		for {
//			line, _ := reader.ReadString('\n')
//			if line == "\n" || line == "\r\n" {
//				// 一个完整的事件读取完成
//				continue
//
//			}
//
//			//fmt.Println("hello")
//
//			fields := strings.SplitN(line, ":", 2)
//			if len(fields) < 2 {
//				continue
//			}
//			//fmt.Println(fields[0])
//
//			switch fields[0] {
//			case "event":
//				fmt.Printf("event: %s\n", fields[1])
//			case "data":
//				fmt.Printf("data: %s\n", fields[1])
//			case "id":
//				fmt.Printf("id: %s\n", fields[1])
//			case "retry":
//				fmt.Printf("retry: %s\n", fields[1])
//			}
//		}
//	}
type BB struct {
	a int
	b int
}

func main() {
	bodyParam := map[string]any{
		"name": "lili",
		"age":  123,
		"sex":  "1234",
	}
	//var bodyStr string
	var body = bytes.NewReader(lo.T2(json.Marshal(bodyParam)).A)
	//bodyStr, ok := bodyParam.(string)
	//if ok {
	//	if len(bodyStr) > 0 {
	//		body = strings.NewReader(bodyStr)
	//	}
	//} else if body, ok = bodyParam.(io.Reader); !ok {
	//	return nil, errors.New("invalid req body param")
	//}

	req, _ := http.NewRequest("POST", "http://127.0.0.1:9999/api/internal/af-sailor-service/v1/assistant/sse", body)
	req.Header.Set("Accept", "text/event-stream")
	//req.Header.Set("Cache-Control", "no-cache")
	//req.Header.Set("Connection", "keep-alive")
	//req.Header.Set("Access-Control-Allow-Origin", "*")
	client := &http.Client{}

	res, _ := client.Do(req)
	defer res.Body.Close()

	reader := bufio.NewReader(res.Body)

	for {
		line, _ := reader.ReadString('\n')
		if line == "\n" || line == "\r\n" {
			// 一个完整的事件读取完成
			continue
		}

		//fmt.Println("hello")

		fields := strings.SplitN(line, ":", 2)
		if len(fields) < 2 {
			continue
		}
		//fmt.Println(fields[0])

		switch fields[0] {
		case "event":
			fmt.Printf("event: %s\n", fields[1])
		case "data":
			fmt.Printf("data: %s\n", fields[1])
		case "id":
			fmt.Printf("id: %s\n", fields[1])
		case "retry":
			fmt.Printf("retry: %s\n", fields[1])
		}
	}
}
