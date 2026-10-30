package main

import (
	"fmt"
	"net/url"
)

func main() {
	// [scheme://][userinfo@]host[:port][/path][?query][#fragment]

	rawUrl := "https://example.com:8080/path?query=param#fragment"

	parseUrl, _ := url.Parse(rawUrl)

	fmt.Println(parseUrl.Scheme)
	fmt.Println(parseUrl.Host)
	fmt.Println(parseUrl.Port())
	fmt.Println(parseUrl.Path)
	fmt.Println(parseUrl.RawQuery)
	fmt.Println(parseUrl.Fragment)

	rawUrl1 := "https://example.com/path?name=Jhon&age=30"

	parseUrl1, _ := url.Parse(rawUrl1)

	queryParams := parseUrl1.Query()

	fmt.Println(queryParams.Get("name"))
	fmt.Println(queryParams["age"])

	baseUrl := &url.URL{
		Scheme: "https",
		Host: "example.com",
		Path: "/path",
	}

	query := baseUrl.Query()
	query.Set("name", "Jhon")
	baseUrl.RawQuery = query.Encode() 

	fmt.Println(baseUrl.String())

	values := url.Values{}
	values.Add("name", "Vaidus")
	values.Add("age", "18")
	encodedQuery := values.Encode()
	fmt.Println(encodedQuery)

	baseUrl2 := "https://example.com/search"
	fullUrl := baseUrl2 + "?" + encodedQuery
	fmt.Println(fullUrl)
}