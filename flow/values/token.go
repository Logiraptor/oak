package values

import "sync"

type Token struct {
	Name string
	ID   int
}

var tokenLock sync.Mutex
var tokens []Token

func NewToken(name string) Token {
	tokenLock.Lock()
	defer tokenLock.Unlock()

	tokens = append(tokens, Token{
		Name: name,
		ID:   len(tokens),
	})

	return tokens[len(tokens)-1]
}
