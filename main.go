package main

import (
	"collectlinks/collect"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup       // wait group creation for text processing
	var wgResult sync.WaitGroup // wait group creation for results writing to new file
	res := make(chan string)    // chanel for goroutines
	var result []string

	fileContent, err := os.ReadFile("repitations") // reading the input file
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var x, y int
	y = 2                     // number of text parts
	x = len(fileContent) / y  // length of text parts
	m := make(map[int][]byte) // text as a map of bytes
	for i := 0; i <= len(fileContent); {
		if i == x*y {
			m[i] = fileContent[i:len(fileContent)]
			break
		}
		m[i] = fileContent[i : i+x]
		i = i + x
	}
	c := collect.New()
	wg.Add(1)
	go func() {
		defer wg.Done()
		res <- c.FindLinks(string(m[n*0]))
		time.Sleep(time.Second * 2)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		res <- c.FindLinks(string(m[n*1]))
		time.Sleep(time.Second * 2)
	}()

	wgResult.Add(1)
	go func() {
		defer wgResult.Done()
		for i := range res {
			result = append(result, i)
		}
	}()
	wg.Wait()
	close(res)
	wgResult.Wait()

	data := strings.Join(result, " ")
	collect.CreateFileandWriteData("file_with_links", data)
	collect.OpenAndPrintFileData("file_with_links")

}