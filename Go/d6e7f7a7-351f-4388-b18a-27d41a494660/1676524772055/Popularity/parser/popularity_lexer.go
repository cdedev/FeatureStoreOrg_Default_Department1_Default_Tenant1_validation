// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676524772055/Popularity.g4 by ANTLR 4.10.1. DO NOT EDIT.

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

type PopularityLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var popularitylexerLexerStaticData struct {
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

func popularitylexerLexerInit() {
	staticData := &popularitylexerLexerStaticData
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
		"", "POPULARITY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"POPULARITY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 4, 42, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1,
		0, 1, 0, 1, 0, 5, 0, 13, 8, 0, 10, 0, 12, 0, 16, 9, 0, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 3, 4, 3, 37, 8, 3, 11, 3, 12, 3, 38, 1, 3, 1, 3, 0,
		0, 4, 1, 1, 3, 2, 5, 3, 7, 4, 1, 0, 2, 1, 0, 48, 57, 3, 0, 9, 10, 13, 13,
		32, 32, 43, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7,
		1, 0, 0, 0, 1, 9, 1, 0, 0, 0, 3, 17, 1, 0, 0, 0, 5, 21, 1, 0, 0, 0, 7,
		36, 1, 0, 0, 0, 9, 10, 3, 5, 2, 0, 10, 14, 3, 3, 1, 0, 11, 13, 7, 0, 0,
		0, 12, 11, 1, 0, 0, 0, 13, 16, 1, 0, 0, 0, 14, 12, 1, 0, 0, 0, 14, 15,
		1, 0, 0, 0, 15, 2, 1, 0, 0, 0, 16, 14, 1, 0, 0, 0, 17, 18, 5, 36, 0, 0,
		18, 19, 5, 126, 0, 0, 19, 20, 5, 36, 0, 0, 20, 4, 1, 0, 0, 0, 21, 22, 5,
		32, 0, 0, 22, 23, 5, 33, 0, 0, 23, 24, 5, 36, 0, 0, 24, 25, 5, 126, 0,
		0, 25, 26, 5, 36, 0, 0, 26, 27, 5, 35, 0, 0, 27, 28, 5, 37, 0, 0, 28, 29,
		5, 35, 0, 0, 29, 30, 5, 38, 0, 0, 30, 31, 5, 36, 0, 0, 31, 32, 5, 38, 0,
		0, 32, 33, 5, 33, 0, 0, 33, 34, 5, 32, 0, 0, 34, 6, 1, 0, 0, 0, 35, 37,
		7, 1, 0, 0, 36, 35, 1, 0, 0, 0, 37, 38, 1, 0, 0, 0, 38, 36, 1, 0, 0, 0,
		38, 39, 1, 0, 0, 0, 39, 40, 1, 0, 0, 0, 40, 41, 6, 3, 0, 0, 41, 8, 1, 0,
		0, 0, 3, 0, 14, 38, 1, 6, 0, 0,
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

// PopularityLexerInit initializes any static state used to implement PopularityLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewPopularityLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func PopularityLexerInit() {
	staticData := &popularitylexerLexerStaticData
	staticData.once.Do(popularitylexerLexerInit)
}

// NewPopularityLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewPopularityLexer(input antlr.CharStream) *PopularityLexer {
	PopularityLexerInit()
	l := new(PopularityLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &popularitylexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Popularity.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// PopularityLexer tokens.
const (
	PopularityLexerPOPULARITY = 1
	PopularityLexerOWNKEY     = 2
	PopularityLexerSPLITKEY   = 3
	PopularityLexerWS         = 4
)
