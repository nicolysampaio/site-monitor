package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
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

	return scannedOption
}

func startMonitoring() {
	fmt.Println("Monitoring...")
	site := "https://httpbin.org/status/404"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "loaded with success!")
	} else {
		fmt.Println("Site:", site, "failed to load. Status Code:", resp.StatusCode)
	}
}
