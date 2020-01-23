package main

import (
	"bufio"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

//获取指定输入流的编码格式
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func htmlanalysis(html string) string {
	r, _ := regexp.Compile("<title>(.*?)</title>")
	title := r.FindStringSubmatch(html)
	if len(title) > 1 {
		return string(title[1])
	}
	return ""
}

func HttpGet(ip string, poc Poc) (*http.Response, error) {
	timeout := time.Duration(6 * time.Second)
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	var url string
	if Port == 80 || Port == 443 {
		url = HttpMethod + "://" + ip + "/" + poc.Rule.Path
	} else {
		url = HttpMethod + "://" + ip + ":" + strconv.Itoa(Port) + "/" + poc.Rule.Path
	}
	resp, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err

	}
	if len(poc.Rule.Header) > 0 {
		for k, v := range poc.Rule.Header {
			resp.Header.Add(k, v)
		}
	} else {
		resp.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_0)")
	}

	resp = resp.WithContext(ctx)

	return http.DefaultClient.Do(resp)
}

func HttpPost(ip string, poc Poc) (*http.Response, error) {
	timeout := time.Duration(6 * time.Second)
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	var url string
	if Port == 80 || Port == 443 {
		url = HttpMethod + "://" + ip
	} else {
		url = HttpMethod + "://" + ip + ":" + strconv.Itoa(Port)
	}
	resp, err := http.NewRequest("POST", url, strings.NewReader(string(poc.Rule.Body)))
	if err != nil {
		return nil, err
	}

	if len(poc.Rule.Header) > 0 {
		for k, v := range poc.Rule.Header {
			resp.Header.Add(k, v)
		}
	} else {
		resp.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_0)")
	}
	resp.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_0)")
	resp = resp.WithContext(ctx)

	return http.DefaultClient.Do(resp)
}

func Request(ip string, pocs []Poc) {
	log.Println(ip + " Being checked")
	for _, poc := range pocs {
		var r *http.Response
		var err error
		if strings.ToUpper(poc.Rule.Method) == "GET" {
			r, err = HttpGet(ip, poc)
		} else {
			r, err = HttpPost(ip, poc)
		}
		if r != nil {
			defer r.Body.Close()
		}
		if err != nil {
			continue
		}
		if r.StatusCode == poc.Rule.Status {
			bodyReader := bufio.NewReader(r.Body)
			e := determineEncoding(bodyReader)
			decodeReader := transform.NewReader(bodyReader, e.NewDecoder())
			all, err := ioutil.ReadAll(decodeReader)
			if err != nil {
				//panic(err)
				return
			} else {
				if strings.Index(string(all), poc.Rule.Contains) != -1 {
					result[ip] = poc.Name

					log.Println(ip + "  Vulnerability detected，The poc:" + poc.Name)

				}
			}
		}

	}

}

func GetResult() {
	if len(result) > 0 {
		fmt.Printf("%d target detected Vulnerability\n", len(result))
		fmt.Println("----------------------------------------------------------------")
		for k, v := range result {
			fmt.Printf("%s    |    %s\n", k, v)
		}
		fmt.Println("----------------------------------------------------------------")
	} else {
		fmt.Println("----------------------------------------------------------------")
		fmt.Println("All target checked. none of them has Vulnerability")
		fmt.Println("----------------------------------------------------------------")
	}
}
