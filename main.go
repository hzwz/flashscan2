/*
 * @Author: your name
 * @Date: 2020-01-19 15:26:26
 * @LastEditTime : 2020-01-24 23:08:54
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
func usage() {
	fmt.Fprintf(os.Stderr, `quantum version: quantum/1.0.0
Usage: flashscan -m http -p 80 -f ip.txt -poc dedecms -t 100
Options:
`)
	flag.PrintDefaults()
}

func main() {
	pocs := Getpoc()
	Head()
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
