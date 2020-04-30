package controllers

import "os"

func PageWriter(err bool, activity string) {
	file, _ := os.Create("file.txt")

	defer file.Close()

	file.WriteString("")
}
