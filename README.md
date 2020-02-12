# Flashscan 
## About 

Flash scan is a pocscanner of Websites. 

-----
## Installation
go is required. 

**Required Dependencies Install**
```bash
root@kali# go build
```
```bash
root@kali# ./flashscan -h
```
-----
### __Help Section__
```
Options:
  -f string
        The file of the target
  -h    Help
  -m string
        Http method,http or https
  -o string
        The output file path of result
  -p int
        The Port of target (default 80)
  -poc string
        The poc file which need to load
  -t int
        The num of threads (default 100)
```
----- 
### __Example Section__
#### *Command line examples*
```bash
root@kali# ./flashscan -m http -p 80 -f 20191031.txt -poc phpstudy -t 10 -o result.txt
```
-----





