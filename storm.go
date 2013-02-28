package main

import "flag"
import "fmt"
import "net/http"
import "runtime"

var (
	client = flag.Int("c", 10, "Clients")
	seconds = flag.Int("t", 60, "Seconds")
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

func main() {
	flag.Parse()
	fmt.Println(*client, "clients, run", *seconds, "seconds")
	runtime.GOMAXPROCS(*client)
	success, fail := 0, 0
	c := make(chan bool)
	for i := 0; i < 3; i++ {
		go fetch("http://www.baidu.com/", c)
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
