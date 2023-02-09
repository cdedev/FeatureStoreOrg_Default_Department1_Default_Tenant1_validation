// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675920663558/Label.g4 by ANTLR 4.10.1. DO NOT EDIT.

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

type LabelLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var labellexerLexerStaticData struct {
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

func labellexerLexerInit() {
	staticData := &labellexerLexerStaticData
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
		"", "LABEL", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"LABEL", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 4, 101, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 75, 8, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 3, 4, 3, 96, 8, 3, 11, 3, 12, 3, 97, 1, 3, 1, 3, 0, 0, 4, 1, 1,
		3, 2, 5, 3, 7, 4, 1, 0, 2, 1, 0, 39, 39, 3, 0, 9, 10, 13, 13, 32, 32, 116,
		0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0,
		1, 9, 1, 0, 0, 0, 3, 76, 1, 0, 0, 0, 5, 80, 1, 0, 0, 0, 7, 95, 1, 0, 0,
		0, 9, 10, 3, 5, 2, 0, 10, 74, 3, 3, 1, 0, 11, 12, 5, 91, 0, 0, 12, 13,
		7, 0, 0, 0, 13, 14, 5, 65, 0, 0, 14, 15, 7, 0, 0, 0, 15, 75, 5, 93, 0,
		0, 16, 17, 5, 91, 0, 0, 17, 18, 7, 0, 0, 0, 18, 19, 5, 67, 0, 0, 19, 20,
		7, 0, 0, 0, 20, 75, 5, 93, 0, 0, 21, 22, 5, 91, 0, 0, 22, 23, 7, 0, 0,
		0, 23, 24, 5, 68, 0, 0, 24, 25, 7, 0, 0, 0, 25, 75, 5, 93, 0, 0, 26, 27,
		5, 91, 0, 0, 27, 28, 7, 0, 0, 0, 28, 29, 5, 71, 0, 0, 29, 30, 7, 0, 0,
		0, 30, 75, 5, 93, 0, 0, 31, 32, 5, 91, 0, 0, 32, 33, 7, 0, 0, 0, 33, 34,
		5, 72, 0, 0, 34, 35, 7, 0, 0, 0, 35, 75, 5, 93, 0, 0, 36, 37, 5, 91, 0,
		0, 37, 38, 7, 0, 0, 0, 38, 39, 5, 77, 0, 0, 39, 40, 7, 0, 0, 0, 40, 75,
		5, 93, 0, 0, 41, 42, 5, 91, 0, 0, 42, 43, 7, 0, 0, 0, 43, 44, 5, 78, 0,
		0, 44, 45, 7, 0, 0, 0, 45, 75, 5, 93, 0, 0, 46, 47, 5, 91, 0, 0, 47, 48,
		7, 0, 0, 0, 48, 49, 5, 79, 0, 0, 49, 50, 7, 0, 0, 0, 50, 75, 5, 93, 0,
		0, 51, 52, 5, 111, 0, 0, 52, 53, 5, 102, 0, 0, 53, 75, 5, 102, 0, 0, 54,
		55, 5, 111, 0, 0, 55, 75, 5, 110, 0, 0, 56, 57, 5, 99, 0, 0, 57, 58, 5,
		97, 0, 0, 58, 75, 5, 112, 0, 0, 59, 60, 5, 111, 0, 0, 60, 61, 5, 117, 0,
		0, 61, 75, 5, 116, 0, 0, 62, 63, 5, 117, 0, 0, 63, 64, 5, 110, 0, 0, 64,
		75, 5, 98, 0, 0, 65, 66, 5, 99, 0, 0, 66, 67, 5, 50, 0, 0, 67, 75, 5, 53,
		0, 0, 68, 69, 5, 99, 0, 0, 69, 70, 5, 55, 0, 0, 70, 75, 5, 53, 0, 0, 71,
		72, 5, 118, 0, 0, 72, 73, 5, 110, 0, 0, 73, 75, 5, 116, 0, 0, 74, 11, 1,
		0, 0, 0, 74, 16, 1, 0, 0, 0, 74, 21, 1, 0, 0, 0, 74, 26, 1, 0, 0, 0, 74,
		31, 1, 0, 0, 0, 74, 36, 1, 0, 0, 0, 74, 41, 1, 0, 0, 0, 74, 46, 1, 0, 0,
		0, 74, 51, 1, 0, 0, 0, 74, 54, 1, 0, 0, 0, 74, 56, 1, 0, 0, 0, 74, 59,
		1, 0, 0, 0, 74, 62, 1, 0, 0, 0, 74, 65, 1, 0, 0, 0, 74, 68, 1, 0, 0, 0,
		74, 71, 1, 0, 0, 0, 75, 2, 1, 0, 0, 0, 76, 77, 5, 36, 0, 0, 77, 78, 5,
		126, 0, 0, 78, 79, 5, 36, 0, 0, 79, 4, 1, 0, 0, 0, 80, 81, 5, 32, 0, 0,
		81, 82, 5, 33, 0, 0, 82, 83, 5, 36, 0, 0, 83, 84, 5, 126, 0, 0, 84, 85,
		5, 36, 0, 0, 85, 86, 5, 35, 0, 0, 86, 87, 5, 37, 0, 0, 87, 88, 5, 35, 0,
		0, 88, 89, 5, 38, 0, 0, 89, 90, 5, 36, 0, 0, 90, 91, 5, 38, 0, 0, 91, 92,
		5, 33, 0, 0, 92, 93, 5, 32, 0, 0, 93, 6, 1, 0, 0, 0, 94, 96, 7, 1, 0, 0,
		95, 94, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 95, 1, 0, 0, 0, 97, 98, 1,
		0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 100, 6, 3, 0, 0, 100, 8, 1, 0, 0, 0, 3,
		0, 74, 97, 1, 6, 0, 0,
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

// LabelLexerInit initializes any static state used to implement LabelLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewLabelLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func LabelLexerInit() {
	staticData := &labellexerLexerStaticData
	staticData.once.Do(labellexerLexerInit)
}

// NewLabelLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewLabelLexer(input antlr.CharStream) *LabelLexer {
	LabelLexerInit()
	l := new(LabelLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &labellexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Label.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// LabelLexer tokens.
const (
	LabelLexerLABEL    = 1
	LabelLexerOWNKEY   = 2
	LabelLexerSPLITKEY = 3
	LabelLexerWS       = 4
)
