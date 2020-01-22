/*
 * @Author: your name
 * @Date: 2020-01-19 15:26:26
 * @LastEditTime : 2020-01-22 23:51:17
 * @LastEditors  : Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /flashscan/main.go
 */
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	h          bool
	HttpMethod string
	IpFile     string
	Port       int
	ThreadNum  int
	PocPath    string
)
var result map[string]string = map[string]string{}

func init() {
	flag.BoolVar(&h, "h", false, "help")
	flag.StringVar(&HttpMethod, "m", "", "http method,http or https")
	flag.StringVar(&IpFile, "f", "", "the file of the target")
	flag.IntVar(&Port, "p", 80, "the Port of target")
	flag.IntVar(&ThreadNum, "t", 100, "the num of threads")
	flag.StringVar(&PocPath, "poc", "", "the file of the target")
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
func usage() {
	fmt.Fprintf(os.Stderr, `quantum version: quantum/1.0.0
Usage: flashscan -m http -p 80 -f ip.txt -poc dedecms -t 100
Options:
`)
	flag.PrintDefaults()
}

func main() {
	pocs := Getpoc()

	fmt.Println(`
	.d888 888                   888                                          
	d88P"  888                   888                                          
	888    888                   888                                          
	888888 888  8888b.  .d8888b  88888b.  .d8888b   .d8888b  8888b.  88888b.  
	888    888     "88b 88K      888 "88b 88K      d88P"        "88b 888 "88b 
	888    888 .d888888 "Y8888b. 888  888 "Y8888b. 888      .d888888 888  888 
	888    888 888  888      X88 888  888      X88 Y88b.    888  888 888  888 
	888    888 "Y888888  88888P' 888  888  88888P'  "Y8888P "Y888888 888  888 																							  
	`)

	p := NewPool(ThreadNum)

	file, err := os.Open("20191031.txt")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	rd := bufio.NewReader(file)
	go func() {
		for {
			line, err := rd.ReadString('\n')
			if err != nil {
				//break
			}
			if io.EOF == err {
				p.EntryChannel <- NewTask(func() error {
					GetResult()
					os.Exit(1)
					return nil
				})

			}
			p.EntryChannel <- NewTask(func() error {
				Request(line[:len(line)-1], pocs)
				return nil
			})
		}
	}()

	//启动协程池p
	p.Run()

}
