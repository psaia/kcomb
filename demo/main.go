package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"runtime"
	"strings"
	"text/template"
	"time"

	"github.com/psaia/kcomb"
)

// Excuse the lack of error handling. The purpose of this is to make a quick demo to better
// demonstrate the library's stream API and efficiency.

const missionStatement = "the {{.Fruit}} fell in love with the {{.Vegetable}}."

func main() {
	col1Data, _ := ioutil.ReadFile("./col-1.txt")
	col2Data, _ := ioutil.ReadFile("./col-2.txt")
	col1DataSlice := strings.Split(string(col1Data), "\n")
	col2DataSlice := strings.Split(string(col2Data), "\n")

	// Trim the empty string on the end.
	col1DataSlice = col1DataSlice[0 : len(col1DataSlice)-1]
	col2DataSlice = col2DataSlice[0 : len(col2DataSlice)-1]

	var col1, col2 []kcomb.Datum

	for _, item := range col1DataSlice {
		col1 = append(col1, kcomb.Datum{Value: item})
	}

	for _, item := range col2DataSlice {
		col2 = append(col2, kcomb.Datum{Value: item})
	}

	done := make(chan struct{})
	stream := compileTpl(done, kcomb.CombineGenerator(done, []kcomb.Set{col1, col2}))
	i := 0

	for v := range stream {
		log.Println(v)
		if i%1000 == 0 {
			printMemUsage()
		}
		time.Sleep(time.Millisecond * 10) // For effect.
		i++
	}
}

// A simple data generator for compiling a template in the stream.
func compileTpl(
	done <-chan struct{},
	valueStream <-chan kcomb.Set,
) <-chan string {
	stream := make(chan string)
	tmpl, _ := template.New("str").Parse(missionStatement)

	go func() {
		defer close(stream)

		for v := range valueStream {
			data := struct {
				Fruit     string
				Vegetable string
			}{
				Fruit:     v[0].Value.(string),
				Vegetable: v[1].Value.(string),
			}
			var tpl bytes.Buffer

			if err := tmpl.Execute(&tpl, data); err != nil {
				panic(err)
			}
			select {
			case <-done:
				return
			case stream <- tpl.String():
			}
		}
	}()
	return stream
}

func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("-------> Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
