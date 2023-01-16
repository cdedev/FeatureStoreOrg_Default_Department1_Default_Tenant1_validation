// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673868747376/Cut.g4 by ANTLR 4.10.1. DO NOT EDIT.

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

type CutLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var cutlexerLexerStaticData struct {
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

func cutlexerLexerInit() {
	staticData := &cutlexerLexerStaticData
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
		"", "CUT", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"CUT", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 4, 67, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 41, 8, 0, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 3, 4, 3, 62, 8, 3, 11, 3, 12, 3, 63, 1, 3, 1, 3, 0,
		0, 4, 1, 1, 3, 2, 5, 3, 7, 4, 1, 0, 1, 3, 0, 9, 10, 13, 13, 32, 32, 71,
		0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0,
		1, 9, 1, 0, 0, 0, 3, 42, 1, 0, 0, 0, 5, 46, 1, 0, 0, 0, 7, 61, 1, 0, 0,
		0, 9, 10, 3, 5, 2, 0, 10, 40, 3, 3, 1, 0, 11, 12, 5, 70, 0, 0, 12, 13,
		5, 97, 0, 0, 13, 14, 5, 105, 0, 0, 14, 41, 5, 114, 0, 0, 15, 16, 5, 71,
		0, 0, 16, 17, 5, 111, 0, 0, 17, 18, 5, 111, 0, 0, 18, 41, 5, 100, 0, 0,
		19, 20, 5, 73, 0, 0, 20, 21, 5, 100, 0, 0, 21, 22, 5, 101, 0, 0, 22, 23,
		5, 97, 0, 0, 23, 41, 5, 108, 0, 0, 24, 25, 5, 80, 0, 0, 25, 26, 5, 114,
		0, 0, 26, 27, 5, 101, 0, 0, 27, 28, 5, 109, 0, 0, 28, 29, 5, 105, 0, 0,
		29, 30, 5, 117, 0, 0, 30, 41, 5, 109, 0, 0, 31, 32, 5, 86, 0, 0, 32, 33,
		5, 101, 0, 0, 33, 34, 5, 114, 0, 0, 34, 35, 5, 121, 0, 0, 35, 36, 5, 32,
		0, 0, 36, 37, 5, 71, 0, 0, 37, 38, 5, 111, 0, 0, 38, 39, 5, 111, 0, 0,
		39, 41, 5, 100, 0, 0, 40, 11, 1, 0, 0, 0, 40, 15, 1, 0, 0, 0, 40, 19, 1,
		0, 0, 0, 40, 24, 1, 0, 0, 0, 40, 31, 1, 0, 0, 0, 41, 2, 1, 0, 0, 0, 42,
		43, 5, 36, 0, 0, 43, 44, 5, 126, 0, 0, 44, 45, 5, 36, 0, 0, 45, 4, 1, 0,
		0, 0, 46, 47, 5, 32, 0, 0, 47, 48, 5, 33, 0, 0, 48, 49, 5, 36, 0, 0, 49,
		50, 5, 126, 0, 0, 50, 51, 5, 36, 0, 0, 51, 52, 5, 35, 0, 0, 52, 53, 5,
		37, 0, 0, 53, 54, 5, 35, 0, 0, 54, 55, 5, 38, 0, 0, 55, 56, 5, 36, 0, 0,
		56, 57, 5, 38, 0, 0, 57, 58, 5, 33, 0, 0, 58, 59, 5, 32, 0, 0, 59, 6, 1,
		0, 0, 0, 60, 62, 7, 0, 0, 0, 61, 60, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63,
		61, 1, 0, 0, 0, 63, 64, 1, 0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 66, 6, 3, 0,
		0, 66, 8, 1, 0, 0, 0, 3, 0, 40, 63, 1, 6, 0, 0,
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

// CutLexerInit initializes any static state used to implement CutLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewCutLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func CutLexerInit() {
	staticData := &cutlexerLexerStaticData
	staticData.once.Do(cutlexerLexerInit)
}

// NewCutLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewCutLexer(input antlr.CharStream) *CutLexer {
	CutLexerInit()
	l := new(CutLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &cutlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Cut.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CutLexer tokens.
const (
	CutLexerCUT      = 1
	CutLexerOWNKEY   = 2
	CutLexerSPLITKEY = 3
	CutLexerWS       = 4
)
