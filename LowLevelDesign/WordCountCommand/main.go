package main

import (
	"fmt"
	"os"
	"strings"
)

type Configuration struct {
	ByteCountNeeded bool
	LineCountNeeded bool
	WordCountNeeded bool
	CharCountNeeded bool
}

func (c *Configuration) SetDefault() {
	c.ByteCountNeeded = true
	c.LineCountNeeded = true
	c.WordCountNeeded = true
}

func main() {
	args := os.Args[1:]
	config := &Configuration{}
	filePath := ""
	if strings.HasPrefix(args[0], "-") {
		for _, c := range args[0][1:] {
			if c == 'l' {
				config.LineCountNeeded = true
			}
			if c == 'c' {
				config.ByteCountNeeded = true
			}
			if c == 'w' {
				config.WordCountNeeded = true
			}
			if c == 'm' {
				config.CharCountNeeded = true
			}
		}
		if len(args) < 2 {
			fmt.Println("file name missing")
			os.Exit(1)
		} else {
			filePath = args[2]
		}
	} else {
		if len(args) < 1 {
			fmt.Println("file name missing")
			os.Exit(1)
		}
		config.SetDefault()
		filePath = args[0]
	}

	f, _ := os.Open(filePath)
	buffer := make([]byte, 100)
	ans := make([]int, 0)
	Bytes := 0
	LC := 0
	WC := 0
	CC := 0
	for true {

		if n, _ := f.Read(buffer); n > 0 {
			s := string(buffer)
			Bytes += n
			charEncountered := false
			for i := 0; i < len(s); i += 1 {
				switch s[i] {
				case '\n':
					LC += 1
					charEncountered = false
				case ' ':
					if charEncountered {
						WC += 1
					}
					charEncountered = false
				default:
					charEncountered = true
					CC += 1
				}
			}
		} else if n == 0 {
			if config.LineCountNeeded {
				ans = append(ans, LC)
			}
			if config.WordCountNeeded {
				ans = append(ans, WC)
			}
			if config.ByteCountNeeded {
				ans = append(ans, Bytes)
			}
			if config.CharCountNeeded {
				ans = append(ans, CC)
			}
			break
		}
	}
	fmt.Println(ans)
	for _, a := range ans {
		fmt.Printf("%d ", a)
	}
	fmt.Printf("%s", filePath)
}

func calculate() {}
