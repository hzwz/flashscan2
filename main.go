/*
 * @Author: wsm
 * @Date: 2020-01-19 15:26:26
 * @LastEditTime : 2020-02-12 18:18:48
 * @LastEditors  : Please set LastEditors
 * @Description:
 * @FilePath: /flashscan/main.go
 */
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

var (
	h          bool
	HttpMethod string
	IpFile     string
	Port       int
	ThreadNum  int
	PocPath    string
	Output     string
)
var result map[string]string = map[string]string{}

func init() {
	flag.BoolVar(&h, "h", false, "Help")
	flag.StringVar(&HttpMethod, "m", "", "Http method,http or https")
	flag.StringVar(&IpFile, "f", "", "The file of the target")
	flag.IntVar(&Port, "p", 80, "The Port of target")
	flag.IntVar(&ThreadNum, "t", 100, "The num of threads")
	flag.StringVar(&PocPath, "poc", "", "The file of the target")
	flag.StringVar(&Output, "o", "", "The output file path of result")
	flag.Usage = usage
	flag.Parse()
	if h {
		flag.Usage()
		os.Exit(0)
	}
	if flag.NFlag() < 4 {
		usage()
		os.Exit(0)
	}
}

func main() {
	pocs := Getpoc()
	Head()
	p := NewPool(ThreadNum)

	file, err := os.Open(IpFile)
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	rd := bufio.NewReader(file)
	go func() {
		for {
			line, _ := rd.ReadString('\n')
			line = strings.Replace(line, "\n", "", -1)
			if line == "" {
				p.EntryChannel <- NewTask(func() error {
					time.Sleep(10 * time.Second)
					GetResult()
					os.Exit(0)
					return nil
				})
				break
			}
			p.EntryChannel <- NewTask(func() error {
				Request(line, pocs)
				return nil
			})

		}

	}()

	//启动协程池p
	p.Run()

}
