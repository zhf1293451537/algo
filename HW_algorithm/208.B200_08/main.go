package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func getResult(lines []string) int {
	count := 0
	for _, line := range lines {
		str := strings.ReplaceAll(line, "\\\"", "") //	delete /"
		str = strings.ReplaceAll(str, "\\'", "")    //	delete /'
		re := regexp.MustCompile(`'[^']*'`)
		str = re.ReplaceAllString(str, "") //delete "adasda"
		re1 := regexp.MustCompile(`"[^"]*"`)
		str = re1.ReplaceAllString(str, "") //delete 'adasda'
		idx := strings.Index(str, "-")
		if idx != -1 {
			str = str[:idx]
		}
		if strings.Contains(str, ";") {
			temp := strings.Split(str, ";")
			fmt.Println(temp, len(temp))
			for i := range temp {
				if i == len(temp)-1 {
					continue
				}
				if strings.TrimSpace(temp[i]) != "" {
					count++
				}
			}
		}
	}
	if strings.Index(lines[len(lines)-1], ";") == -1 {
		count++
	}
	return count
}
func main() {
	lines := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if text != "" {
			lines = append(lines, text)
		} else {
			fmt.Println(getResult(lines))
			break
		}
	}
}
