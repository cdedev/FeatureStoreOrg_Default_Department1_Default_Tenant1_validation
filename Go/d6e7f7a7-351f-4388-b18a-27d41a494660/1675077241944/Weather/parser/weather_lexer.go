// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675077241944/Weather.g4 by ANTLR 4.10.1. DO NOT EDIT.

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

type WeatherLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var weatherlexerLexerStaticData struct {
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

func weatherlexerLexerInit() {
	staticData := &weatherlexerLexerStaticData
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
		"", "WEATHER", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"WEATHER", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 4, 96, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1,
		0, 1, 0, 1, 0, 5, 0, 13, 8, 0, 10, 0, 12, 0, 16, 9, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 3, 0, 70, 8, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 4, 3, 91,
		8, 3, 11, 3, 12, 3, 92, 1, 3, 1, 3, 0, 0, 4, 1, 1, 3, 2, 5, 3, 7, 4, 1,
		0, 2, 2, 0, 46, 46, 48, 57, 3, 0, 9, 10, 13, 13, 32, 32, 103, 0, 1, 1,
		0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 1, 9, 1,
		0, 0, 0, 3, 71, 1, 0, 0, 0, 5, 75, 1, 0, 0, 0, 7, 90, 1, 0, 0, 0, 9, 10,
		3, 5, 2, 0, 10, 14, 3, 3, 1, 0, 11, 13, 7, 0, 0, 0, 12, 11, 1, 0, 0, 0,
		13, 16, 1, 0, 0, 0, 14, 12, 1, 0, 0, 0, 14, 15, 1, 0, 0, 0, 15, 69, 1,
		0, 0, 0, 16, 14, 1, 0, 0, 0, 17, 18, 5, 67, 0, 0, 18, 19, 5, 108, 0, 0,
		19, 20, 5, 101, 0, 0, 20, 21, 5, 97, 0, 0, 21, 70, 5, 114, 0, 0, 22, 23,
		5, 67, 0, 0, 23, 24, 5, 108, 0, 0, 24, 25, 5, 111, 0, 0, 25, 26, 5, 117,
		0, 0, 26, 27, 5, 100, 0, 0, 27, 70, 5, 121, 0, 0, 28, 29, 5, 70, 0, 0,
		29, 30, 5, 111, 0, 0, 30, 31, 5, 103, 0, 0, 31, 32, 5, 103, 0, 0, 32, 70,
		5, 121, 0, 0, 33, 34, 5, 82, 0, 0, 34, 35, 5, 97, 0, 0, 35, 36, 5, 105,
		0, 0, 36, 37, 5, 110, 0, 0, 37, 70, 5, 121, 0, 0, 38, 39, 5, 83, 0, 0,
		39, 40, 5, 104, 0, 0, 40, 41, 5, 111, 0, 0, 41, 42, 5, 119, 0, 0, 42, 43,
		5, 101, 0, 0, 43, 44, 5, 114, 0, 0, 44, 70, 5, 121, 0, 0, 45, 46, 5, 83,
		0, 0, 46, 47, 5, 117, 0, 0, 47, 48, 5, 110, 0, 0, 48, 49, 5, 110, 0, 0,
		49, 70, 5, 121, 0, 0, 50, 51, 5, 86, 0, 0, 51, 52, 5, 97, 0, 0, 52, 53,
		5, 108, 0, 0, 53, 54, 5, 117, 0, 0, 54, 55, 5, 101, 0, 0, 55, 56, 5, 32,
		0, 0, 56, 57, 5, 110, 0, 0, 57, 58, 5, 111, 0, 0, 58, 59, 5, 116, 0, 0,
		59, 60, 5, 32, 0, 0, 60, 61, 5, 97, 0, 0, 61, 62, 5, 118, 0, 0, 62, 63,
		5, 97, 0, 0, 63, 64, 5, 105, 0, 0, 64, 65, 5, 108, 0, 0, 65, 66, 5, 97,
		0, 0, 66, 67, 5, 98, 0, 0, 67, 68, 5, 108, 0, 0, 68, 70, 5, 101, 0, 0,
		69, 17, 1, 0, 0, 0, 69, 22, 1, 0, 0, 0, 69, 28, 1, 0, 0, 0, 69, 33, 1,
		0, 0, 0, 69, 38, 1, 0, 0, 0, 69, 45, 1, 0, 0, 0, 69, 50, 1, 0, 0, 0, 70,
		2, 1, 0, 0, 0, 71, 72, 5, 36, 0, 0, 72, 73, 5, 126, 0, 0, 73, 74, 5, 36,
		0, 0, 74, 4, 1, 0, 0, 0, 75, 76, 5, 32, 0, 0, 76, 77, 5, 33, 0, 0, 77,
		78, 5, 36, 0, 0, 78, 79, 5, 126, 0, 0, 79, 80, 5, 36, 0, 0, 80, 81, 5,
		35, 0, 0, 81, 82, 5, 37, 0, 0, 82, 83, 5, 35, 0, 0, 83, 84, 5, 38, 0, 0,
		84, 85, 5, 36, 0, 0, 85, 86, 5, 38, 0, 0, 86, 87, 5, 33, 0, 0, 87, 88,
		5, 32, 0, 0, 88, 6, 1, 0, 0, 0, 89, 91, 7, 1, 0, 0, 90, 89, 1, 0, 0, 0,
		91, 92, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 94, 1,
		0, 0, 0, 94, 95, 6, 3, 0, 0, 95, 8, 1, 0, 0, 0, 4, 0, 14, 69, 92, 1, 6,
		0, 0,
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

// WeatherLexerInit initializes any static state used to implement WeatherLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewWeatherLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func WeatherLexerInit() {
	staticData := &weatherlexerLexerStaticData
	staticData.once.Do(weatherlexerLexerInit)
}

// NewWeatherLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewWeatherLexer(input antlr.CharStream) *WeatherLexer {
	WeatherLexerInit()
	l := new(WeatherLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &weatherlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Weather.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// WeatherLexer tokens.
const (
	WeatherLexerWEATHER  = 1
	WeatherLexerOWNKEY   = 2
	WeatherLexerSPLITKEY = 3
	WeatherLexerWS       = 4
)
