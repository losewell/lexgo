package scanner

import (
	"os"
	"log"
	"fmt"
        "io"
	"unicode"
	"strings"
	"path/filepath"
)

func ungetNextChar(f *os.File) (err error) {
	_, err = f.Seek(-1, 1)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func reservedWalk(s string) TokenType {
	for _, word := range reservedWordSet {
		if s == word.str {
			return word.tok
		}
	}
	return ID
}

func getNextToken(f *os.File, fpath, dirname string) error {
	var state StateType
	var currentToken TokenType
	var save bool
	hasEscapeChar := false

	b := make([]byte, 1)
	for {
		save = true
		_, err := f.Read(b)
		if err == io.EOF {
			state = DONE
			currentToken = ENDFILE
			break
		}
		c := b[0]

		switch state {
		case START:
			if unicode.IsDigit(rune(c)) {
				state = INNUM
			} else if unicode.IsLetter(rune(c)) {
				state = INID
			} else if unicode.IsSpace(rune(c)) || c == ',' || c == '.'{
				save = false
				if c == '\n' {
					lineno++
				}
			} else if c == ':' {
				state = IN_DECLARE_ASSIGN
			} else if c == '=' {
				state = INASSIGN
			} else if c == '"' {
				save = false
				state = INSTRING
			}else if c == '/' {
				save = false
				state = ENTERING_COMMENT
			} else {
				state = DONE
				switch c {
				case '>':
					currentToken = MT
				case '<':
					currentToken = LT
				case '+':
					currentToken = PLUS
				case '-':
					currentToken = MINUS
				case '*':
					currentToken = TIMES
				case '/':
					currentToken = OVER
				case '(':
					currentToken = LPAREN
				case ')':
					currentToken = RPAREN
				case '{':
					currentToken = LBRACE
				case '}':
					currentToken = RBRACE
//				case '\"':
//					currentToken = LQUOTA
//				case '\"':
//					currentToken = RQUOTA
				case ';':
					currentToken = SEMI
				default:
					currentToken = ERROR
				}
			}

		case IN_DECLARE_ASSIGN:
			state = DONE
			if c == '=' {
				currentToken = DECLARE_ASSIGN
			} else {
				ungetNextChar(f)
				save = false
				currentToken = ERROR

			}

		case INASSIGN:
			state = DONE
			if c == '=' {
				currentToken = EQ
			} else {
				ungetNextChar(f)
				save = false
				currentToken = ASSIGN
			}

		case ENTERING_COMMENT:
			save = false
			if c == '*' {
				state = IN_COMMENT
			} else {
				ungetNextChar(f)
				save = false
				state = DONE
				currentToken = OVER
			}

		case IN_COMMENT:
			save = false
			if c == '*' {
				state = EXITING_COMMENT
			}

		case EXITING_COMMENT:
			if c == '/' {
				state = START
			} else {
				state = IN_COMMENT
			}

		case INNUM:
			if !unicode.IsDigit(rune(c)) {
				ungetNextChar(f)
				currentToken = NUM
				state = DONE
				save = false
			}

		case INSTRING:
			if c == '\\' && !hasEscapeChar {
				save = false
				hasEscapeChar = true
			} else if c == '"' && !hasEscapeChar {
				currentToken = STRING
				state = DONE
				save = false
			}

		case INID:
			if !unicode.IsLetter(rune(c)) {
				ungetNextChar(f)
				currentToken = ID
				state = DONE
				save = false
			}

		}
		if save {
			tokenBuffer = append(tokenBuffer, c)
		}
		if state == DONE {
			if currentToken == ID {
				currentToken = reservedWalk(string(tokenBuffer[:]))
			}
			fmt.Println("token:", currentToken, "tokenString:", tokenBuffer)
			generateLexFile(currentToken, tokenBuffer, fpath, dirname)
			state = START
			tokenBuffer = nil
		}
	}
	return nil
}

func SourcefileWalk(fpath, dirname string) (err error) {
	f, err := os.Open(fpath)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	getNextToken(f, fpath, dirname)
	return
}

func generateLexFile(tokenType TokenType, token []byte, fpath, dirname string) {
	var err error
	lastDot := strings.LastIndex(fpath, ".")
	lexFileName := fmt.Sprintf("%s.lex", fpath[:lastDot])
	absFileName := filepath.Join(dirname, lexFileName)

	lexf, err = os.OpenFile(absFileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = lexf.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	l := &LexInfo{tokenType, string(token[:])}
	fmt.Println("l:", l)
	_, err = fmt.Fprintln(lexf, l.String())
	if err != nil {
		log.Fatal(err)
	}
	return
}
