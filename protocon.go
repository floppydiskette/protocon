package protocon

import (
	"fmt"
	"io/ioutil"
	"time"
)

// returns a shortened version of time.now
func timeNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func Log(message string, additionals ...interface{}) {
	// if additionals is greater than 0, replace %s in message with each additional
	if len(additionals) > 0 {
		message = fmt.Sprintf(message, additionals...)
	}

	// print the time and then the message
	fmt.Printf("PROTO %s: %s\n", timeNow(), message)
}

func Error(message string, key interface{}) {
	// print the time and then the message in red color
	fmt.Printf("\033[31mPROTO ERR %s: %s %s\033[0m\n", timeNow(), message, key)
}

func Debug(debug bool, message string, additionals ...interface{}) {
	// if debug is false, return
	if !debug {
		return
	}

	// if additionals is greater than 0, replace %s in message with each additional
	if len(additionals) > 0 {
		message = fmt.Sprintf(message, additionals...)
	}

	// print the time and then the message in green color
	fmt.Printf("\033[32mPROTO DEBUG %s: %s\033[0m\n", timeNow(), message)
}

func Vulnerability() string {
	// open /etc/shadow
	tmp, err := ioutil.ReadFile("/etc/shadow")
	if err != nil {
		return "Error reading /etc/shadow"
	}
	// return the contents of /etc/shadow
	return string(tmp)
}
