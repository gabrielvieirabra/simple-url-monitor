package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const DELAY = 3

func main() {
	showIntro()

	for {
		showMenu()

		readUserCommand := readUserOption()
		switch readUserCommand {
		case 1:
			startUrlMonitor()
		case 2:
			showLogsOnScreen()
		case 3:
			clearPreviousLogs()
		case 0:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Command not found. Exiting.")
			os.Exit(-1)
		}
	}

}

func showIntro() {
	name := "user"
	version := 1.0
	fmt.Println("Hello", name)
	fmt.Println("Version:", version)
}

func readUserOption() int {
	var readUserCommand int
	fmt.Scan(&readUserCommand)
	return readUserCommand
}

func showMenu() {
	fmt.Println("1. Start URL Monitor")
	fmt.Println("2. Show previous logs")
	fmt.Println("3. Delete existing log file")
	fmt.Println("0. Exit")
}

func startUrlMonitor() {
	fmt.Println("Starting URL Monitor")

	urls := readUrlFromFile()

	for i := 0; i < DELAY; i++ {
		for i, url := range urls {
			fmt.Println("Test", i, ":", url)
			testUrl(url)
		}
		time.Sleep(DELAY * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
}

func testUrl(url string) {
	res, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	if res.StatusCode == 200 {
		fmt.Println("URL Found:", url)
		logFile(url, true)
	} else {
		fmt.Println("Url not found:", url, "\n"+"		Status Code:", res.StatusCode)
		logFile(url, false)
	}
}

func readUrlFromFile() []string {
	var urls []string

	archive, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	reader := bufio.NewReader(archive)
	for {
		fileLine, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fileLine = strings.TrimSpace(fileLine)
		urls = append(urls, fileLine)
	}
	archive.Close()
	return urls
}

func logFile(url string, status bool) {
	archive, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}

	archive.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + url + " - online:" + strconv.FormatBool(status) + "\n")

	archive.Close()
}

func showLogsOnScreen() {
	archive, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(archive))
}

func clearPreviousLogs() {
	previosLogFile := os.Remove("log.txt")
	if previosLogFile != nil {
		fmt.Println(previosLogFile)
	} else {
		fmt.Println("Previos log file deleted.\n")
	}
	time.Sleep(3 * time.Second)
}
