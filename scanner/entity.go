package scanner

import (
	"os"
	"strconv"
)

var (
	lineno = 1
	tokenBuffer []byte
	lexf *os.File
)

type StateType int

const (
	START StateType = iota
	IN_DECLARE_ASSIGN
	INASSIGN
	INEQ
	INID
	INNUM
	INSTRING
	ENTERING_COMMENT
	IN_COMMENT
	EXITING_COMMENT
	DONE
)

type TokenType int

const (
	BREAK TokenType = iota
	DEFAULT
	FUNC
	INTERFACE
	SELECT
	CASE
	DEFER
	GO
	MAP
	STRUCT
	CHAN
	ELSE
	GOTO
	PACKAGE
	SWITCH
	CONST
	FALLTHROUGH
	IF
	RANGE
	TYPE
	CONTINUE
	FOR
	IMPORT
	RETURN
	VAR

	ID
	NUM
	STRING

	DECLARE_ASSIGN
	ASSIGN
	EQ
	MT
	LT
	PLUS
	MINUS
	TIMES
	OVER
	SEMI
	LPAREN
	RPAREN
	LBRACE
	RBRACE
	LQUOTA
	RQUOTA

	ENDFILE
	ERROR
)

type ReservedWord struct {
	str string
	tok TokenType
}

var (
	reservedWordSet = [...]ReservedWord{ReservedWord{"break", BREAK,},ReservedWord{"default", DEFAULT,},
                                            ReservedWord{"func", FUNC,},ReservedWord{"interface", INTERFACE,},
					    ReservedWord{"select", SELECT,},ReservedWord{"case", CASE,},
					    ReservedWord{"chan", CHAN,},ReservedWord{"const", CONST,},
					    ReservedWord{"continue", CONTINUE,},ReservedWord{"defer", DEFER},
					    ReservedWord{"go", GO,},ReservedWord{"map", MAP,},
					    ReservedWord{"struct", STRUCT,},ReservedWord{"else", ELSE},
					    ReservedWord{"goto", GOTO},ReservedWord{"package", PACKAGE},
					    ReservedWord{"switch", SWITCH,},ReservedWord{"fallthrough", FALLTHROUGH,},
					    ReservedWord{"if", IF,},ReservedWord{"range", RANGE,},
					    ReservedWord{"type", TYPE,},ReservedWord{"for", FOR,},
					    ReservedWord{"import", IMPORT,},ReservedWord{"return", RETURN,},
					    ReservedWord{"var", VAR,},
				    }
)

type LexInfo struct {
	token TokenType
	tokenString string
}

func (l *LexInfo) String() string {
	switch(l.token) {
	case BREAK, DEFAULT, FUNC, INTERFACE,
	     SELECT, CASE, CHAN, CONST, CONTINUE,
	     DEFER, GO, MAP, STRUCT, ELSE, GOTO,
	     PACKAGE, SWITCH, FALLTHROUGH, IF,
	     RANGE, TYPE, FOR, IMPORT,
	     RETURN, VAR:
		return strconv.Itoa(lineno) + " : reserved word: " + l.tokenString

	case DECLARE_ASSIGN:
		return strconv.Itoa(lineno) + " :="

	case ASSIGN:
		return strconv.Itoa(lineno) + " ="

	case EQ:
		return strconv.Itoa(lineno) + " =="

	case LT:
		return strconv.Itoa(lineno) + " <"

	case MT:
		return strconv.Itoa(lineno) + " >"

	case PLUS:
		return strconv.Itoa(lineno) + " +"

	case MINUS:
		return strconv.Itoa(lineno) + " -"

	case TIMES:
		return strconv.Itoa(lineno) + " *"

	case OVER:
		return strconv.Itoa(lineno) + " /"

	case LPAREN:
		return strconv.Itoa(lineno) + " ("

	case RPAREN:
		return strconv.Itoa(lineno) + " )"

	case LBRACE:
		return strconv.Itoa(lineno) + " {"

	case RBRACE:
		return strconv.Itoa(lineno) + " }"

	case RQUOTA:
		return strconv.Itoa(lineno) + " \""

	case LQUOTA:
		return strconv.Itoa(lineno) +  "\""

	case SEMI:
		return strconv.Itoa(lineno) + " ;"

	case ENDFILE:
		return strconv.Itoa(lineno) + " EOF"

	case NUM:
		return strconv.Itoa(lineno) + " NUM, val=" + l.tokenString

	case STRING:
		return strconv.Itoa(lineno) + " STRING, string=" + l.tokenString

	case ID:
		return strconv.Itoa(lineno) + " ID, name=" + l.tokenString

	case ERROR:
		return strconv.Itoa(lineno) + " ERROR:" + l.tokenString

	default:
		return strconv.Itoa(lineno) + " Unknown token:" + strconv.Itoa(int(l.token))

	}
	return ""
}
