// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675078478535/Geography.g4 by ANTLR 4.10.1. DO NOT EDIT.

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

type GeographyLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var geographylexerLexerStaticData struct {
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

func geographylexerLexerInit() {
	staticData := &geographylexerLexerStaticData
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
		"", "GEOGRAPHY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"GEOGRAPHY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 4, 95, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 69,
		8, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 4, 3, 90, 8, 3, 11, 3,
		12, 3, 91, 1, 3, 1, 3, 0, 0, 4, 1, 1, 3, 2, 5, 3, 7, 4, 1, 0, 1, 3, 0,
		9, 10, 13, 13, 32, 32, 101, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1,
		0, 0, 0, 0, 7, 1, 0, 0, 0, 1, 9, 1, 0, 0, 0, 3, 70, 1, 0, 0, 0, 5, 74,
		1, 0, 0, 0, 7, 89, 1, 0, 0, 0, 9, 10, 3, 5, 2, 0, 10, 68, 3, 3, 1, 0, 11,
		12, 5, 66, 0, 0, 12, 13, 5, 97, 0, 0, 13, 14, 5, 114, 0, 0, 14, 15, 5,
		108, 0, 0, 15, 16, 5, 101, 0, 0, 16, 69, 5, 121, 0, 0, 17, 18, 5, 67, 0,
		0, 18, 19, 5, 111, 0, 0, 19, 20, 5, 97, 0, 0, 20, 21, 5, 114, 0, 0, 21,
		22, 5, 115, 0, 0, 22, 23, 5, 101, 0, 0, 23, 24, 5, 32, 0, 0, 24, 25, 5,
		103, 0, 0, 25, 26, 5, 114, 0, 0, 26, 27, 5, 97, 0, 0, 27, 28, 5, 105, 0,
		0, 28, 29, 5, 110, 0, 0, 29, 69, 5, 115, 0, 0, 30, 31, 5, 67, 0, 0, 31,
		32, 5, 111, 0, 0, 32, 33, 5, 114, 0, 0, 33, 69, 5, 110, 0, 0, 34, 35, 5,
		69, 0, 0, 35, 36, 5, 110, 0, 0, 36, 37, 5, 101, 0, 0, 37, 38, 5, 114, 0,
		0, 38, 39, 5, 103, 0, 0, 39, 40, 5, 121, 0, 0, 40, 41, 5, 32, 0, 0, 41,
		42, 5, 102, 0, 0, 42, 43, 5, 101, 0, 0, 43, 44, 5, 101, 0, 0, 44, 45, 5,
		100, 0, 0, 45, 69, 5, 115, 0, 0, 46, 47, 5, 70, 0, 0, 47, 48, 5, 101, 0,
		0, 48, 49, 5, 101, 0, 0, 49, 50, 5, 100, 0, 0, 50, 51, 5, 32, 0, 0, 51,
		52, 5, 103, 0, 0, 52, 53, 5, 114, 0, 0, 53, 54, 5, 97, 0, 0, 54, 55, 5,
		105, 0, 0, 55, 56, 5, 110, 0, 0, 56, 69, 5, 115, 0, 0, 57, 58, 5, 79, 0,
		0, 58, 59, 5, 97, 0, 0, 59, 60, 5, 116, 0, 0, 60, 69, 5, 115, 0, 0, 61,
		62, 5, 83, 0, 0, 62, 63, 5, 111, 0, 0, 63, 64, 5, 114, 0, 0, 64, 65, 5,
		103, 0, 0, 65, 66, 5, 104, 0, 0, 66, 67, 5, 117, 0, 0, 67, 69, 5, 109,
		0, 0, 68, 11, 1, 0, 0, 0, 68, 17, 1, 0, 0, 0, 68, 30, 1, 0, 0, 0, 68, 34,
		1, 0, 0, 0, 68, 46, 1, 0, 0, 0, 68, 57, 1, 0, 0, 0, 68, 61, 1, 0, 0, 0,
		69, 2, 1, 0, 0, 0, 70, 71, 5, 36, 0, 0, 71, 72, 5, 126, 0, 0, 72, 73, 5,
		36, 0, 0, 73, 4, 1, 0, 0, 0, 74, 75, 5, 32, 0, 0, 75, 76, 5, 33, 0, 0,
		76, 77, 5, 36, 0, 0, 77, 78, 5, 126, 0, 0, 78, 79, 5, 36, 0, 0, 79, 80,
		5, 35, 0, 0, 80, 81, 5, 37, 0, 0, 81, 82, 5, 35, 0, 0, 82, 83, 5, 38, 0,
		0, 83, 84, 5, 36, 0, 0, 84, 85, 5, 38, 0, 0, 85, 86, 5, 33, 0, 0, 86, 87,
		5, 32, 0, 0, 87, 6, 1, 0, 0, 0, 88, 90, 7, 0, 0, 0, 89, 88, 1, 0, 0, 0,
		90, 91, 1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 93, 1,
		0, 0, 0, 93, 94, 6, 3, 0, 0, 94, 8, 1, 0, 0, 0, 3, 0, 68, 91, 1, 6, 0,
		0,
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

// GeographyLexerInit initializes any static state used to implement GeographyLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewGeographyLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func GeographyLexerInit() {
	staticData := &geographylexerLexerStaticData
	staticData.once.Do(geographylexerLexerInit)
}

// NewGeographyLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewGeographyLexer(input antlr.CharStream) *GeographyLexer {
	GeographyLexerInit()
	l := new(GeographyLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &geographylexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Geography.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// GeographyLexer tokens.
const (
	GeographyLexerGEOGRAPHY = 1
	GeographyLexerOWNKEY    = 2
	GeographyLexerSPLITKEY  = 3
	GeographyLexerWS        = 4
)
