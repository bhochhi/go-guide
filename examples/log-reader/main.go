package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("myapp.log")
	if err != nil {
		log.Fatal("____error here.")
	}
	defer f.Close()
	r := bufio.NewReader(f)

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			log.Fatal("Error reading...")
			break
		}
		if strings.Contains(line, "ERROR") {
			fmt.Println(line)
		}

	}

}
