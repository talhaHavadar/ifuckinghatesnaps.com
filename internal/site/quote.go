package site

import (
	"errors"
	"fmt"
	"html"
	"html/template"
	"io/fs"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

const highlightMarker = "**"

var (
	quoteFileName  = regexp.MustCompile(`^([0-9]+)\.txt$`)
	paragraphBreak = regexp.MustCompile(`\n[ \t]*\n`)
)

type quote struct {
	ID         int
	Paragraphs []template.HTML
}

type collection struct {
	byID map[int]quote
	ids  []int
}

func loadQuotes(fsys fs.FS) (collection, error) {
	entries, err := fs.ReadDir(fsys, ".")
	if err != nil {
		return collection{}, fmt.Errorf("listing quotes: %w", err)
	}

	c := collection{byID: map[int]quote{}}
	for _, entry := range entries {
		name := entry.Name()
		id, err := parseQuoteFileName(name)
		if err != nil {
			return collection{}, err
		}
		if _, dup := c.byID[id]; dup {
			return collection{}, fmt.Errorf("%s: quote %d is defined twice", name, id)
		}
		raw, err := fs.ReadFile(fsys, name)
		if err != nil {
			return collection{}, fmt.Errorf("reading %s: %w", name, err)
		}
		paragraphs, err := renderQuote(string(raw))
		if err != nil {
			return collection{}, fmt.Errorf("%s: %w", name, err)
		}
		c.byID[id] = quote{ID: id, Paragraphs: paragraphs}
		c.ids = append(c.ids, id)
	}
	if len(c.ids) == 0 {
		return collection{}, errors.New("no quotes found")
	}
	slices.Sort(c.ids)
	return c, nil
}

func parseQuoteFileName(name string) (int, error) {
	m := quoteFileName.FindStringSubmatch(name)
	if m == nil {
		return 0, fmt.Errorf("%s: quote files must be named <number>.txt", name)
	}
	id, err := strconv.Atoi(m[1])
	if err != nil || id < 1 {
		return 0, fmt.Errorf("%s: quote number must be a positive integer", name)
	}
	return id, nil
}

func renderQuote(text string) ([]template.HTML, error) {
	text = strings.TrimSpace(strings.ReplaceAll(text, "\r\n", "\n"))
	if text == "" {
		return nil, errors.New("quote is empty")
	}

	var paragraphs []template.HTML
	for _, p := range paragraphBreak.Split(text, -1) {
		rendered, err := renderParagraph(strings.TrimSpace(p))
		if err != nil {
			return nil, err
		}
		paragraphs = append(paragraphs, rendered)
	}
	return paragraphs, nil
}

func renderParagraph(p string) (template.HTML, error) {
	parts := strings.Split(p, highlightMarker)
	if len(parts)%2 == 0 {
		return "", fmt.Errorf("unclosed %s in paragraph %q", highlightMarker, p)
	}

	var b strings.Builder
	for i, part := range parts {
		escaped := html.EscapeString(part)
		if i%2 == 0 {
			b.WriteString(escaped)
			continue
		}
		if strings.TrimSpace(part) == "" {
			return "", fmt.Errorf("empty highlight in paragraph %q", p)
		}
		b.WriteString("<mark>" + escaped + "</mark>")
	}
	// Every piece of quote text went through html.EscapeString above.
	return template.HTML(b.String()), nil
}

func (c collection) latest() quote {
	return c.byID[c.ids[len(c.ids)-1]]
}

func (c collection) previous(id int) int {
	i, _ := slices.BinarySearch(c.ids, id)
	if i == 0 {
		return c.ids[len(c.ids)-1]
	}
	return c.ids[i-1]
}
