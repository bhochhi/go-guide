package main

import (

	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	path := flag.String("path", "myapp.log", "The path to the log that should be analyzed")
	level := flag.String("level", "ERROR", "Log level to search for. Options are DEBUG, INFO, ERROR, and CRITICAL")
	f, err := os.Open(*path)
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
		if strings.Contains(line, *level) {
			fmt.Println(line)
		}

	}

}
