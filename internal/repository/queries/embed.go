package queries

import (
	"embed"
	"regexp"
	"strings"
)

//go:embed knife.sql
var sqlFiles embed.FS

var queryNameRegex = regexp.MustCompile(`(?m)^--\s*name:\s*(\w+)\s*$`)

type Queries struct {
	queries map[string]string
}

func (q *Queries) Get(name string) string {
	return q.queries[name]
}

func MustLoad(filename string) *Queries {
	content, _ := sqlFiles.ReadFile(filename)
	queries := make(map[string]string)
	matches := queryNameRegex.FindAllStringSubmatchIndex(string(content), -1)

	for i, match := range matches {
		name := string(content[match[2]:match[3]])
		queryStart := match[1]
		queryEnd := len(content)
		if i+1 < len(matches) {
			queryEnd = matches[i+1][0]
		}
		queries[name] = strings.TrimSpace(string(content[queryStart:queryEnd]))
	}

	return &Queries{queries: queries}
}
