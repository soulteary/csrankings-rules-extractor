package extractor

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/soulteary/csrankings-rules-extractor/internal/define"
)

func getRule(r *regexp.Regexp, line string) (string, error) {
	matched := r.FindAllStringSubmatch(line, -1)
	if len(matched) == 1 && len(matched[0]) == 2 {
		return strings.TrimSpace(matched[0][1]), nil
	} else {
		return "", fmt.Errorf("unhandle rule: %s", line)
	}
}

func GetFilters(file string) ([]string, []string, error) {
	buf, err := os.ReadFile(file)
	if err != nil {
		return nil, nil, err
	}

	var REGEXP_INPROCEEDINGS = regexp.MustCompile(`booktitle\s*=\s*"(.+)?"`)
	var REGEXP_ARTICLE = regexp.MustCompile(`journal\s*=\s*"(.+)?"`)

	inproceedings := make([]string, 0)
	articles := make([]string, 0)

	lines := strings.Split(string(buf), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "//inproceedings") {
			if strings.HasPrefix(line, "//inproceedings[booktitle") {
				ret, err := getRule(REGEXP_INPROCEEDINGS, line)
				if err != nil {
					fmt.Println("[unhandle inproceedings]", err)
				} else {
					inproceedings = append(inproceedings, ret)
				}
			} else {
				fmt.Println("[unhandle inproceedings]", line)
			}
		} else if strings.HasPrefix(line, "//article") {
			if strings.HasPrefix(line, "//article[journal") {
				ret, err := getRule(REGEXP_ARTICLE, line)
				if err != nil {
					fmt.Println("[unhandle article]", err)
				} else {
					articles = append(articles, ret)
				}
			} else {
				fmt.Println("[unhandle article]", line)
			}
		} else {
			if define.FLAG_DEBUG {
				fmt.Println("[discord]", line)
			}
		}
	}
	return inproceedings, articles, nil
}
