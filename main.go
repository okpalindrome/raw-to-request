package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var final_curl []string
var url_path_params string
var host string

func err_check(e error) {
	if e != nil {
		fmt.Print(e)
		panic(e)
	}
}

func verb_path_version(line string) {
	final_curl = append(final_curl, "curl")

	re := regexp.MustCompile(`([^\s]+)`)
	matches := re.FindAllString(line, 3)

	switch matches[2] {
	case "HTTP/1.0":
		final_curl = append(final_curl, "--http1.0")
	case "HTTP/1.1":
		final_curl = append(final_curl, "--http1.1")
	case "HTTP/2":
		final_curl = append(final_curl, "--http2")
	case "HTTP/3":
		final_curl = append(final_curl, "--http3")
	}

	final_curl = append(final_curl, "-X", matches[0])

	url_path_params = matches[1]

	if strings.HasPrefix(matches[1], "http://") || strings.HasPrefix(matches[1], "https://") {
		final_curl = append(final_curl, url_path_params)
	}
}

func header_body(lines []string) {

	var count int

	// headers
	header_regex := regexp.MustCompile(`([a-zA-Z0-9-]*[a-zA-Z0-9])?:\s*(.*)`)
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			break
		}
		matches := header_regex.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			count++
			if len(match) > 2 {

				match[2] = strings.ReplaceAll(match[2], "\\", "\\\\") // Replace \ with \\
				match[2] = strings.ReplaceAll(match[2], "\"", "\\\"") // Replace " with \"

				wrapped := fmt.Sprintf("$'%s: %s'", match[1], match[2])

				final_curl = append(final_curl, "-H", wrapped)

				if match[1] == "Host" {
					host = strings.TrimSuffix(match[2], "/")
					// host = match[2]
				}
			}
		}

	}

	// request body
	if count+1 > len(lines) {
		// Do nothing
	} else {
		body := strings.Join(lines[count+1:], "\n")

		body = strings.ReplaceAll(body, "\\", "\\\\") // Replace \ with \\
		body = strings.ReplaceAll(body, "\"", "\\\"") // Replace " with \"

		wrapped_body := fmt.Sprintf("$'%s'", body)

		final_curl = append(final_curl, "-b", wrapped_body)
	}

	if !(strings.HasPrefix(url_path_params, "http://") || strings.HasPrefix(url_path_params, "https://")) {
		url_path_params = "https://" + host + url_path_params
		wrapped := fmt.Sprintf("$'%s'", url_path_params)
		final_curl = append(final_curl, wrapped)
	}

}

func main() {

	filename := flag.String("file", "", "Raw request file to parse.")

	flag.Usage = func() {
		fmt.Println("Usage of raw2curl:")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *filename == "" {
		flag.Usage()
		os.Exit(0)
	}

	readFile, err := os.Open(*filename)
	err_check(err)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	defer readFile.Close()

	verb_path_version(fileLines[0])

	header_body(fileLines[1:])

	fmt.Print(strings.Join(final_curl, " "))
}
