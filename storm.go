package main

import (
    "flag"
    "fmt"
    "io/ioutil"
    "net/http"
    "runtime"
    "strings"
    "time"
)

const VERSION = "0.1.0"

var (
	help = flag.Bool("h", false, "Helps")
	client = flag.Int("c", 10, "Clients")
	seconds = flag.Int64("t", 60, "Seconds")
	url = flag.String("url", "", "URL")
	file = flag.String("f", "", "URL list file")
	urls []string
)

func fetch(url string, c chan bool) {
	status := false
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	if http.StatusOK == resp.StatusCode {
		status = true
	}

	c <- status
}

func readLines(filename string)(lines []string, err error) {
    bytes, err := ioutil.ReadFile(filename)
    if err != nil {
        return
    }
    
    for _, line := range strings.Split(string(bytes), "\n") {
        if line != "" {             
            lines = append(lines, line)
        }
    }
    return
}

func usage() {
	fmt.Println("Storm version:", VERSION)

	fmt.Println("Usage: storm [-c clients] [-t Seconds] [-url url]")
	fmt.Println("")
	fmt.Println("Options:")
	fmt.Println("-h\t:", "this help")
	fmt.Println("-c\t:", "how many clients you want to simulate")
	fmt.Println("-t\t:", "how long you want to bench")
	fmt.Println("-url\t:", "bench url")
	fmt.Println("-f\t:", "url list file, one url per line.")
}

func main() {
	flag.Parse()
	if *help {
		usage()
		return
	}
	filename := *file
	if (*url == "" && filename == "") {
	    fmt.Println("Please input url.")
	    return
	}
	
	if filename != "" {
	    lines, err := readLines(filename)
        if err != nil {
            fmt.Println(err)
            return
        }
        urls = lines
	}
	fmt.Println(*client, "clients, run", *seconds, "seconds")
	runtime.GOMAXPROCS(*client)
	
	success, fail := 0, 0
	c := make(chan bool)
	
	st := time.Now().Unix()
	for (st + *seconds > time.Now().Unix()) {
	    if filename != "" {
	        for _, uri := range urls {
                go fetch(uri, c)
            }
	    } else {
	        go fetch(*url, c)
	    }	    
		
		status := <- c
		if status == true {
			success += 1
		}else{
			fail += 1
		}
	}

	total := success + fail
	fmt.Println("Total:", total)
	fmt.Println("Success:", success)
	fmt.Println("Fail:", fail)
}
