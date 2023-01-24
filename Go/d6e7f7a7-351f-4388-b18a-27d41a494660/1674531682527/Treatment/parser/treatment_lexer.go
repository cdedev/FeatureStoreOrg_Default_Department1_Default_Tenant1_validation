// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674531682527/Treatment.g4 by ANTLR 4.10.1. DO NOT EDIT.

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

type TreatmentLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var treatmentlexerLexerStaticData struct {
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

func treatmentlexerLexerInit() {
	staticData := &treatmentlexerLexerStaticData
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
		"", "TREATMENT", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"TREATMENT", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 4, 63, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 3, 0, 37, 8, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3,
		4, 3, 58, 8, 3, 11, 3, 12, 3, 59, 1, 3, 1, 3, 0, 0, 4, 1, 1, 3, 2, 5, 3,
		7, 4, 1, 0, 1, 3, 0, 9, 10, 13, 13, 32, 32, 65, 0, 1, 1, 0, 0, 0, 0, 3,
		1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 1, 9, 1, 0, 0, 0, 3, 38,
		1, 0, 0, 0, 5, 42, 1, 0, 0, 0, 7, 57, 1, 0, 0, 0, 9, 10, 3, 5, 2, 0, 10,
		36, 3, 3, 1, 0, 11, 12, 5, 112, 0, 0, 12, 13, 5, 108, 0, 0, 13, 14, 5,
		97, 0, 0, 14, 15, 5, 99, 0, 0, 15, 16, 5, 101, 0, 0, 16, 17, 5, 98, 0,
		0, 17, 37, 5, 111, 0, 0, 18, 19, 5, 80, 0, 0, 19, 20, 5, 114, 0, 0, 20,
		21, 5, 111, 0, 0, 21, 22, 5, 103, 0, 0, 22, 23, 5, 97, 0, 0, 23, 24, 5,
		98, 0, 0, 24, 25, 5, 105, 0, 0, 25, 26, 5, 100, 0, 0, 26, 37, 5, 101, 0,
		0, 27, 28, 5, 112, 0, 0, 28, 29, 5, 114, 0, 0, 29, 30, 5, 111, 0, 0, 30,
		31, 5, 103, 0, 0, 31, 32, 5, 97, 0, 0, 32, 33, 5, 98, 0, 0, 33, 34, 5,
		105, 0, 0, 34, 35, 5, 100, 0, 0, 35, 37, 5, 101, 0, 0, 36, 11, 1, 0, 0,
		0, 36, 18, 1, 0, 0, 0, 36, 27, 1, 0, 0, 0, 37, 2, 1, 0, 0, 0, 38, 39, 5,
		36, 0, 0, 39, 40, 5, 126, 0, 0, 40, 41, 5, 36, 0, 0, 41, 4, 1, 0, 0, 0,
		42, 43, 5, 32, 0, 0, 43, 44, 5, 33, 0, 0, 44, 45, 5, 36, 0, 0, 45, 46,
		5, 126, 0, 0, 46, 47, 5, 36, 0, 0, 47, 48, 5, 35, 0, 0, 48, 49, 5, 37,
		0, 0, 49, 50, 5, 35, 0, 0, 50, 51, 5, 38, 0, 0, 51, 52, 5, 36, 0, 0, 52,
		53, 5, 38, 0, 0, 53, 54, 5, 33, 0, 0, 54, 55, 5, 32, 0, 0, 55, 6, 1, 0,
		0, 0, 56, 58, 7, 0, 0, 0, 57, 56, 1, 0, 0, 0, 58, 59, 1, 0, 0, 0, 59, 57,
		1, 0, 0, 0, 59, 60, 1, 0, 0, 0, 60, 61, 1, 0, 0, 0, 61, 62, 6, 3, 0, 0,
		62, 8, 1, 0, 0, 0, 3, 0, 36, 59, 1, 6, 0, 0,
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

// TreatmentLexerInit initializes any static state used to implement TreatmentLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewTreatmentLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func TreatmentLexerInit() {
	staticData := &treatmentlexerLexerStaticData
	staticData.once.Do(treatmentlexerLexerInit)
}

// NewTreatmentLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewTreatmentLexer(input antlr.CharStream) *TreatmentLexer {
	TreatmentLexerInit()
	l := new(TreatmentLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &treatmentlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Treatment.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// TreatmentLexer tokens.
const (
	TreatmentLexerTREATMENT = 1
	TreatmentLexerOWNKEY    = 2
	TreatmentLexerSPLITKEY  = 3
	TreatmentLexerWS        = 4
)
