package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"regexp"
	"strings"
)

func parseFile(inputFile string) ([]Line, error) {
	var lastPerson string
	var lastTime string
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	results := []Line{}

	// this regex "kinda" looks like it'll just match headers
	// `Person Name  <min>:<sec>  `
	re := regexp.MustCompile(`.*:[0-9][0-9]  $`)

	for scanner.Scan() {
		txt := scanner.Text()
		// process one line at a time
		if re.MatchString(txt) {
			person, err := getNameFromTitle(re.FindString(txt))
			if err != nil {
				return results, err
			}
			time, err := getTimeFromTitle(re.FindString(txt))
			if err != nil {
				return results, err
			}
			lastPerson = person
			lastTime = time
		} else if len(txt) > 0 {
			x := Line{
				Person: lastPerson,
				Time:   lastTime,
				Text:   txt,
			}
			if !strings.Contains(txt, "Transcribed by https://otter.ai") {
				results = append(results, x)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return results, err
	}
	return results, err
}

func getTimeFromTitle(input string) (string, error) {
	x := strings.Split(input, "  ")
	if len(x) != 3 {
		return "", errors.New("couldn't determine time")
	}
	return x[1], nil
}
func getNameFromTitle(input string) (string, error) {
	x := strings.Split(input, "  ")
	if len(x) != 3 {
		return "", errors.New("couldn't determine username")
	}
	return x[0], nil
}

type Line struct {
	Person string
	Time   string
	Text   string
}

// type ResultantLines struct {
// 	Row Line
// }
