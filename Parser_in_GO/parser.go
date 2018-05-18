package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

type as_befehl struct {
	opcode string
	exp1   int
	exp2   int
}

var (
	i             int
	befehl        string
	parsed_befehl = make([]string, 100)
	err           error
)

func main() {
	// Parse to list...
	test := strings.Split(befehl, ",")
	fmt.Println(test)

	// Ließ ganzes FileSync-settings: installing git-gui (9/15)

	/*
		data, err := ioutil.ReadFile("original.s")
		check(err)
		fmt.Println(string(data))
	*/

	// Öffne File
	/*
		f, err := os.Open("/tmp/dat")
		check(err)



		// Schließe es wieder
		f.Close()
	*/

	var bla = make([]string, 100)
	bla = readFileWithReadLine("original.s")
	//fmt.Println(bla)

	for i = 0; i < len(bla); i++ {
		fmt.Println(bla[i])
	}

}

func readFileWithReadLine(fn string) []string {

	// Open File
	file, err := os.Open(fn)
	defer file.Close()

	if err != nil {
		check(err)
	}

	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)

	for {
		var buffer bytes.Buffer

		var l []byte
		var isPrefix bool
		for {
			l, isPrefix, err = reader.ReadLine()
			buffer.Write(l)

			// If we've reached the end of the line, stop reading.
			if !isPrefix {
				break
			}

			// If we're just at the EOF, break
			if err != nil {
				break
			}
		}

		if err == io.EOF {
			break
		}

		line := buffer.String()

		parsed_befehl = append(parsed_befehl, line)

	}
	// Wenn er nicht bis zum Dateiende gelaufen ist...
	if err != io.EOF {
		fmt.Printf(" > Failed!: %v\n", err)
	}

	return parsed_befehl
}

// Allgemeiner Errorhandler
func check(e error) {
	if e != nil {
		panic(e)
	}
}

//func tokenizer(string_array []string)
