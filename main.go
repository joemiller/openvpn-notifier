package main

// note: a more feature-filled pushover lib, if needed in the future: https://github.com/gregdel/pushover

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/bdenning/pushover"
	"github.com/hpcloud/tail"
)

var (
	// May  4 18:46:11 gw openvpn[24037]: 16.17.4.3:28867 [joe-iphone] Peer Connection Initiated with [AF_INET]16.17.4.3:28867
	r = regexp.MustCompile(`^\w+\s+\d+ \d+:\d+:\d+ \w+ openvpn\[\d+\]: ([\d\.]+):\d+ \[([\w\-_\.]+)\] Peer Connection Initiated`)

	user    string
	token   string
	logFile string
)

func main() {
	user = os.Getenv("PUSHOVER_USER")
	if user == "" {
		log.Fatal("Missing environment var PUSHOVER_USER")
	}
	token = os.Getenv("PUSHOVER_TOKEN")
	if token == "" {
		log.Fatal("Missing environment var PUSHOVER_TOKEN")
	}
	logFile = os.Getenv("OPENVPN_LOGFILE")
	if logFile == "" {
		logFile = "/var/log/messages"
	}

	config := tail.Config{
		Follow:    true,
		ReOpen:    true,
		MustExist: true,
		Location:  &tail.SeekInfo{0, os.SEEK_END},
		Poll:      true,
	}
	t, err := tail.TailFile(logFile, config)
	if err != nil {
		log.Fatal(err)
	}

	for line := range t.Lines {
		matches := r.FindStringSubmatch(line.Text)
		if matches != nil {
			msg := fmt.Sprintf("VPN client connected: %s (%s)", matches[2], matches[1])
			log.Println(msg)
			notify(msg)
		}
	}
}

func notify(message string) {
	m := pushover.NewMessage(token, user)
	_, err := m.Push(message)
	if err != nil {
		log.Printf("Error sending notification: %s", err)
	}
}
