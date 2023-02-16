// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676522166037/Sick.g4 by ANTLR 4.10.1. DO NOT EDIT.

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

type SickLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var sicklexerLexerStaticData struct {
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

func sicklexerLexerInit() {
	staticData := &sicklexerLexerStaticData
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
		"", "SICK", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"SICK", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 4, 38, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 4, 3, 33,
		8, 3, 11, 3, 12, 3, 34, 1, 3, 1, 3, 0, 0, 4, 1, 1, 3, 2, 5, 3, 7, 4, 1,
		0, 2, 2, 0, 102, 102, 116, 116, 3, 0, 9, 10, 13, 13, 32, 32, 38, 0, 1,
		1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 1, 9,
		1, 0, 0, 0, 3, 13, 1, 0, 0, 0, 5, 17, 1, 0, 0, 0, 7, 32, 1, 0, 0, 0, 9,
		10, 3, 5, 2, 0, 10, 11, 3, 3, 1, 0, 11, 12, 7, 0, 0, 0, 12, 2, 1, 0, 0,
		0, 13, 14, 5, 36, 0, 0, 14, 15, 5, 126, 0, 0, 15, 16, 5, 36, 0, 0, 16,
		4, 1, 0, 0, 0, 17, 18, 5, 32, 0, 0, 18, 19, 5, 33, 0, 0, 19, 20, 5, 36,
		0, 0, 20, 21, 5, 126, 0, 0, 21, 22, 5, 36, 0, 0, 22, 23, 5, 35, 0, 0, 23,
		24, 5, 37, 0, 0, 24, 25, 5, 35, 0, 0, 25, 26, 5, 38, 0, 0, 26, 27, 5, 36,
		0, 0, 27, 28, 5, 38, 0, 0, 28, 29, 5, 33, 0, 0, 29, 30, 5, 32, 0, 0, 30,
		6, 1, 0, 0, 0, 31, 33, 7, 1, 0, 0, 32, 31, 1, 0, 0, 0, 33, 34, 1, 0, 0,
		0, 34, 32, 1, 0, 0, 0, 34, 35, 1, 0, 0, 0, 35, 36, 1, 0, 0, 0, 36, 37,
		6, 3, 0, 0, 37, 8, 1, 0, 0, 0, 2, 0, 34, 1, 6, 0, 0,
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

// SickLexerInit initializes any static state used to implement SickLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewSickLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func SickLexerInit() {
	staticData := &sicklexerLexerStaticData
	staticData.once.Do(sicklexerLexerInit)
}

// NewSickLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewSickLexer(input antlr.CharStream) *SickLexer {
	SickLexerInit()
	l := new(SickLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &sicklexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Sick.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SickLexer tokens.
const (
	SickLexerSICK     = 1
	SickLexerOWNKEY   = 2
	SickLexerSPLITKEY = 3
	SickLexerWS       = 4
)
