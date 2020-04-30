package main

import (
	"Crawler/controllers"
	"fmt"
	"os"
)

func main() {

	Twitter := os.Args[1]
	if Twitter == "" {
		fmt.Println("Please set a the Twitter id. Twitter id can be foudn at the end of the url like : twitter.com/(id)")
	}
	controllers.Crawler(Twitter)

}
