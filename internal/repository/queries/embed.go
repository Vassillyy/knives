package queries

import (
	"embed"
	"fmt"
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
	query, ok := q.queries[name]
	if !ok {
		panic(fmt.Sprintf("query %q not found", name))
	}
	return query
}

func MustLoad(filename string) *Queries {
	q, err := load(filename)
	if err != nil {
		panic(err)
	}
	return q
}

func load(filename string) (*Queries, error) {
	content, err := sqlFiles.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("read sql file %s: %w", filename, err)
	}

	queries, err := parse(string(content))
	if err != nil {
		return nil, fmt.Errorf("parse sql file %s: %w", filename, err)
	}

	return &Queries{queries: queries}, nil
}

func parse(content string) (map[string]string, error) {
	queries := make(map[string]string)
	matches := queryNameRegex.FindAllStringSubmatchIndex(content, -1)
	if len(matches) == 0 {
		return nil, fmt.Errorf("no queries found")
	}

	for i, match := range matches {
		nameStart, nameEnd := match[2], match[3]
		name := content[nameStart:nameEnd]

		queryStart := match[1]
		var queryEnd int
		if i+1 < len(matches) {
			queryEnd = matches[i+1][0]
		} else {
			queryEnd = len(content)
		}

		query := strings.TrimSpace(content[queryStart:queryEnd])
		queries[name] = query
	}
	return queries, nil
}
