package main

import (
	"fmt"
	"net/http"
	"time"
	"io/ioutil"
	"regexp"
	"log"
	"net/smtp"
)

func send(body string) {
	from := "name@yandex.ru"
	pass := "pass"
	to := "name@gmail.com"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: My new Ip\n\n" +
		body

	err := smtp.SendMail("smtp.yandex.ru:587",
		smtp.PlainAuth("", from, pass, "smtp.yandex.ru"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
	
	log.Print("sent")
}

func main() {
	var myNewIp string = "192.168.0.1"
	var myOldIp string = "192.168.1.1"

	for {

		// Make HTTP request
		resp, err := http.Get("http://yandex.ru/internet")
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		// Read response data in to memory
	    body, err := ioutil.ReadAll(resp.Body)
	    if err != nil {
	        log.Fatal("Error reading HTTP body. ", err)
    	}
    	//convert byte body to string
		pageContent := string(body)
		//fmt.Println(pageContent)

    	// Create a regular expression to find ipV4
		re := regexp.MustCompile(`((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)`)
		myNewIp = re.FindString(pageContent)


		// send my ip to email
		if myNewIp != myOldIp && myNewIp !="" {
			fmt.Println("my ip = ", myNewIp)
			send(myNewIp)
			myOldIp = myNewIp
		}

		//sleep 300 second
		time.Sleep(300 * time.Second)

	}
}
