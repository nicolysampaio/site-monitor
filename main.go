package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitors = 3
const delay = 10

func main() {
	fmt.Println("")
	showIntroduction()

	for {
		showMenu()

		option := scanOption()

		switch option {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Displaying logs...")
		case 0:
			fmt.Println("Exiting program...")
			os.Exit(0)
		default:
			fmt.Println("Invalid command")
			os.Exit(-1)
		}
	}
}

func showIntroduction() {
	name := "Nicoly"
	version := 1.1

	fmt.Println("Hello, Mr/Ms.", name)
	fmt.Println("This program is in version", version)

	fmt.Println("")
}

func showMenu() {
	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Display logs")
	fmt.Println("0 - Exit the program")
}

func scanOption() int {
	var scannedOption int
	fmt.Scan(&scannedOption)

	fmt.Println("The chosen option was", scannedOption)

	fmt.Println("")

	return scannedOption
}

func startMonitoring() {
	fmt.Println("Monitoring...")
	fmt.Println("")

	sites := readSitesFile()

	for i := 0; i < monitors; i++ {
		for i, site := range sites {
			fmt.Println("Testing site", i, ":", site)
			monitorSite(site)
			fmt.Println("")
		}

		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func monitorSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("An error occured:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "loaded successfully!")
	} else {
		fmt.Println("Site:", site, "failed to load. Status Code:", resp.StatusCode)
	}
}

func readSitesFile() []string {
	var sites []string

	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("An error occured:", err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()
	return sites
}
