package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var finalCurl []string
var urlPathParams string
var host string

func errCheck(msg string, e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", msg, e)
	} else {
		fmt.Fprintln(os.Stderr, msg)
	}
	os.Exit(1)
}

func verbPathVersion(line string) {
	finalCurl = append(finalCurl, "curl")

	re := regexp.MustCompile(`([^\s]+)`)
	matches := re.FindAllString(line, 3)

	if len(matches) < 1 {
		errCheck("[ERROR] Make sure it's a valid raw HTTP request!", nil)
	}

	switch matches[2] {
	case "HTTP/1.0":
		finalCurl = append(finalCurl, "--http1.0")
	case "HTTP/1.1":
		finalCurl = append(finalCurl, "--http1.1")
	case "HTTP/2":
		finalCurl = append(finalCurl, "--http2")
	case "HTTP/3":
		finalCurl = append(finalCurl, "--http3")
	}

	finalCurl = append(finalCurl, "-X", matches[0])
	urlPathParams = matches[1]

	if strings.HasPrefix(matches[1], "http://") || strings.HasPrefix(matches[1], "https://") {
		finalCurl = append(finalCurl, urlPathParams)
	}
}

func headerBody(lines []string) {
	var count int

	// headers
	headerRegex := regexp.MustCompile(`([a-zA-Z0-9-]*[a-zA-Z0-9])?:\s*(.*)`)
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			break
		}
		matches := headerRegex.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			count++
			if len(match) > 2 {
				match[2] = strings.ReplaceAll(match[2], "\\", "\\\\")
				match[2] = strings.ReplaceAll(match[2], "\"", "\\\"")

				wrapped := fmt.Sprintf("-H '%s: %s'", match[1], match[2])
				finalCurl = append(finalCurl, wrapped)

				if match[1] == "Host" {
					host = strings.TrimSuffix(match[2], "/")
				}
			}
		}
	}

	// request body
	if !(count+1 > len(lines)) {
		body := strings.Join(lines[count+1:], "\n")
		body = strings.ReplaceAll(body, "\\", "\\\\")
		body = strings.ReplaceAll(body, "\"", "\\\"")
		
		wrapped := fmt.Sprintf("-d '%s'", body)
		finalCurl = append(finalCurl, wrapped)
	}

	if !(strings.HasPrefix(urlPathParams, "http://") || strings.HasPrefix(urlPathParams, "https://")) {
		urlPathParams = "https://" + host + urlPathParams
		wrapped := fmt.Sprintf("$'%s'", urlPathParams) // can skip this
		finalCurl = append(finalCurl, wrapped)
	}
}

func readInput() []string {
	var lines []string

	// Check if there's input from pipe
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			errCheck("[ERROR] Failed to read from stdin", err)
		}
		return lines
	}

	// If no pipe input, try to read from file argument
	if len(os.Args) > 1 {
		filename := os.Args[1]
		fileinfo, err := os.Stat(filename)
		if err != nil {
			errCheck("[ERROR] File does not exist", err)
		}
		if fileinfo.IsDir() {
			errCheck("[ERROR] Directory path, not a file", nil)
		}

		file, err := os.Open(filename)
		if err != nil {
			errCheck("[ERROR] Failed to open file", err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			errCheck("[ERROR] Failed to read file", err)
		}
		return lines
	}

	fmt.Fprintln(os.Stderr, "Usage: raw2curl <file-path> or pipe input via stdin")
	os.Exit(1)
	return nil
}

func main() {
	lines := readInput()

	if len(lines) < 2 {
		errCheck("[ERROR] Input must contain at least a request line and headers", nil)
	}

	verbPathVersion(lines[0])
	headerBody(lines[1:])
	fmt.Println(strings.Join(finalCurl, " "))
}
