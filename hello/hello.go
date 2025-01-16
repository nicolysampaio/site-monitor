package main

import (
	"fmt"
	"net/http"
	"os"
)

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

	sites := []string{"https://www.caelum.com.br", "https://httpbin.org/status/500", "https://httpbin.org/status/200", "https://www.alura.com.br"}

	for i, site := range sites {
		fmt.Println("Testing site", i, ":", site)
		monitorSite(site)
		fmt.Println("")
	}

	fmt.Println("")
}

func monitorSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "loaded with success!")
	} else {
		fmt.Println("Site:", site, "failed to load. Status Code:", resp.StatusCode)
	}
}
