package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

var partern = `(\d+)@qq.com`

func main() {
	url := "https://tieba.baidu.com/p/6051076813?red_tag=1573533731"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	pageBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	pageStr := string(pageBytes)

	re := regexp.MustCompile(partern)
	results := re.FindAllStringSubmatch(pageStr, -1)

	for _, result := range results {
		fmt.Println("email: ", result[0])
		fmt.Println("qq: ", result[1])
	}

}
