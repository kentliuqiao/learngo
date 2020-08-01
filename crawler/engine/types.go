package engine

// ParserFunc ...
type ParserFunc func(contents []byte, url string) ParseResult

// Request ...
type Request struct {
	URL        string
	ParserFunc ParserFunc
}

// ParseResult ...
type ParseResult struct {
	Requests []Request
	Items    []Item
}

// Item ...
type Item struct {
	URL  string
	Type string
	ID   string

	Payload interface{}
}

// NilParser do nothing
func NilParser([]byte) ParseResult {
	return ParseResult{}
}
