// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674799725772/OperatingSystem.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type OperatingSystemLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var operatingsystemlexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func operatingsystemlexerLexerInit() {
	staticData := &operatingsystemlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "OPERATINGSYSTEM", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"OPERATINGSYSTEM", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 4, 48, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 3,
		0, 22, 8, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 4, 3, 43, 8, 3, 11,
		3, 12, 3, 44, 1, 3, 1, 3, 0, 0, 4, 1, 1, 3, 2, 5, 3, 7, 4, 1, 0, 1, 3,
		0, 9, 10, 13, 13, 32, 32, 49, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5,
		1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 1, 9, 1, 0, 0, 0, 3, 23, 1, 0, 0, 0, 5, 27,
		1, 0, 0, 0, 7, 42, 1, 0, 0, 0, 9, 10, 3, 5, 2, 0, 10, 21, 3, 3, 1, 0, 11,
		12, 5, 65, 0, 0, 12, 13, 5, 110, 0, 0, 13, 14, 5, 100, 0, 0, 14, 15, 5,
		114, 0, 0, 15, 16, 5, 111, 0, 0, 16, 17, 5, 105, 0, 0, 17, 22, 5, 100,
		0, 0, 18, 19, 5, 105, 0, 0, 19, 20, 5, 79, 0, 0, 20, 22, 5, 83, 0, 0, 21,
		11, 1, 0, 0, 0, 21, 18, 1, 0, 0, 0, 22, 2, 1, 0, 0, 0, 23, 24, 5, 36, 0,
		0, 24, 25, 5, 126, 0, 0, 25, 26, 5, 36, 0, 0, 26, 4, 1, 0, 0, 0, 27, 28,
		5, 32, 0, 0, 28, 29, 5, 33, 0, 0, 29, 30, 5, 36, 0, 0, 30, 31, 5, 126,
		0, 0, 31, 32, 5, 36, 0, 0, 32, 33, 5, 35, 0, 0, 33, 34, 5, 37, 0, 0, 34,
		35, 5, 35, 0, 0, 35, 36, 5, 38, 0, 0, 36, 37, 5, 36, 0, 0, 37, 38, 5, 38,
		0, 0, 38, 39, 5, 33, 0, 0, 39, 40, 5, 32, 0, 0, 40, 6, 1, 0, 0, 0, 41,
		43, 7, 0, 0, 0, 42, 41, 1, 0, 0, 0, 43, 44, 1, 0, 0, 0, 44, 42, 1, 0, 0,
		0, 44, 45, 1, 0, 0, 0, 45, 46, 1, 0, 0, 0, 46, 47, 6, 3, 0, 0, 47, 8, 1,
		0, 0, 0, 3, 0, 21, 44, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// OperatingSystemLexerInit initializes any static state used to implement OperatingSystemLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewOperatingSystemLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func OperatingSystemLexerInit() {
	staticData := &operatingsystemlexerLexerStaticData
	staticData.once.Do(operatingsystemlexerLexerInit)
}

// NewOperatingSystemLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewOperatingSystemLexer(input antlr.CharStream) *OperatingSystemLexer {
	OperatingSystemLexerInit()
	l := new(OperatingSystemLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &operatingsystemlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "OperatingSystem.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// OperatingSystemLexer tokens.
const (
	OperatingSystemLexerOPERATINGSYSTEM = 1
	OperatingSystemLexerOWNKEY          = 2
	OperatingSystemLexerSPLITKEY        = 3
	OperatingSystemLexerWS              = 4
)
