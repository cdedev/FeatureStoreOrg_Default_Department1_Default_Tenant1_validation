// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673235590013/Axons.g4 by ANTLR 4.10.1. DO NOT EDIT.

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

type AxonsLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var axonslexerLexerStaticData struct {
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

func axonslexerLexerInit() {
	staticData := &axonslexerLexerStaticData
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
		"", "AXONS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"AXONS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 4, 40, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		3, 4, 3, 35, 8, 3, 11, 3, 12, 3, 36, 1, 3, 1, 3, 0, 0, 4, 1, 1, 3, 2, 5,
		3, 7, 4, 1, 0, 2, 1, 0, 48, 49, 3, 0, 9, 10, 13, 13, 32, 32, 40, 0, 1,
		1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 1, 9,
		1, 0, 0, 0, 3, 15, 1, 0, 0, 0, 5, 19, 1, 0, 0, 0, 7, 34, 1, 0, 0, 0, 9,
		10, 3, 5, 2, 0, 10, 11, 3, 3, 1, 0, 11, 12, 7, 0, 0, 0, 12, 13, 7, 0, 0,
		0, 13, 14, 7, 0, 0, 0, 14, 2, 1, 0, 0, 0, 15, 16, 5, 36, 0, 0, 16, 17,
		5, 126, 0, 0, 17, 18, 5, 36, 0, 0, 18, 4, 1, 0, 0, 0, 19, 20, 5, 32, 0,
		0, 20, 21, 5, 33, 0, 0, 21, 22, 5, 36, 0, 0, 22, 23, 5, 126, 0, 0, 23,
		24, 5, 36, 0, 0, 24, 25, 5, 35, 0, 0, 25, 26, 5, 37, 0, 0, 26, 27, 5, 35,
		0, 0, 27, 28, 5, 38, 0, 0, 28, 29, 5, 36, 0, 0, 29, 30, 5, 38, 0, 0, 30,
		31, 5, 33, 0, 0, 31, 32, 5, 32, 0, 0, 32, 6, 1, 0, 0, 0, 33, 35, 7, 1,
		0, 0, 34, 33, 1, 0, 0, 0, 35, 36, 1, 0, 0, 0, 36, 34, 1, 0, 0, 0, 36, 37,
		1, 0, 0, 0, 37, 38, 1, 0, 0, 0, 38, 39, 6, 3, 0, 0, 39, 8, 1, 0, 0, 0,
		2, 0, 36, 1, 6, 0, 0,
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

// AxonsLexerInit initializes any static state used to implement AxonsLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewAxonsLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func AxonsLexerInit() {
	staticData := &axonslexerLexerStaticData
	staticData.once.Do(axonslexerLexerInit)
}

// NewAxonsLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewAxonsLexer(input antlr.CharStream) *AxonsLexer {
	AxonsLexerInit()
	l := new(AxonsLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &axonslexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Axons.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// AxonsLexer tokens.
const (
	AxonsLexerAXONS    = 1
	AxonsLexerOWNKEY   = 2
	AxonsLexerSPLITKEY = 3
	AxonsLexerWS       = 4
)
