package scanner

import (
	"ghuo/scanner/token"
)

type scanner struct {
	source  string
	tokens  []token.Token
	start   int
	current int
	line    int
}

func (s *scanner) scan() {

}

func (s *scanner) nextch() byte {
	return s.source[s.current]
}

func (s *scanner) isEnd() bool {
	return s.current >= len(s.source)
}
