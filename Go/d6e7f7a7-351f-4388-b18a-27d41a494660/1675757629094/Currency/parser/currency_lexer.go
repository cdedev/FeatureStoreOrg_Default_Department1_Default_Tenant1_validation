// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675757629094/Currency.g4 by ANTLR 4.10.1. DO NOT EDIT.

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

type CurrencyLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var currencylexerLexerStaticData struct {
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

func currencylexerLexerInit() {
	staticData := &currencylexerLexerStaticData
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
		"", "CURRENCY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"CURRENCY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 4, 206, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 3, 0, 180, 8, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 4,
		3, 201, 8, 3, 11, 3, 12, 3, 202, 1, 3, 1, 3, 0, 0, 4, 1, 1, 3, 2, 5, 3,
		7, 4, 1, 0, 1, 3, 0, 9, 10, 13, 13, 32, 32, 261, 0, 1, 1, 0, 0, 0, 0, 3,
		1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 1, 9, 1, 0, 0, 0, 3, 181,
		1, 0, 0, 0, 5, 185, 1, 0, 0, 0, 7, 200, 1, 0, 0, 0, 9, 10, 3, 5, 2, 0,
		10, 179, 3, 3, 1, 0, 11, 12, 5, 69, 0, 0, 12, 13, 5, 85, 0, 0, 13, 180,
		5, 82, 0, 0, 14, 15, 5, 85, 0, 0, 15, 16, 5, 83, 0, 0, 16, 180, 5, 68,
		0, 0, 17, 18, 5, 65, 0, 0, 18, 19, 5, 69, 0, 0, 19, 180, 5, 68, 0, 0, 20,
		21, 5, 65, 0, 0, 21, 22, 5, 82, 0, 0, 22, 180, 5, 83, 0, 0, 23, 24, 5,
		65, 0, 0, 24, 25, 5, 85, 0, 0, 25, 180, 5, 68, 0, 0, 26, 27, 5, 66, 0,
		0, 27, 28, 5, 72, 0, 0, 28, 180, 5, 68, 0, 0, 29, 30, 5, 66, 0, 0, 30,
		31, 5, 82, 0, 0, 31, 180, 5, 76, 0, 0, 32, 33, 5, 67, 0, 0, 33, 34, 5,
		65, 0, 0, 34, 180, 5, 68, 0, 0, 35, 36, 5, 67, 0, 0, 36, 37, 5, 72, 0,
		0, 37, 180, 5, 70, 0, 0, 38, 39, 5, 67, 0, 0, 39, 40, 5, 76, 0, 0, 40,
		180, 5, 80, 0, 0, 41, 42, 5, 67, 0, 0, 42, 43, 5, 78, 0, 0, 43, 180, 5,
		89, 0, 0, 44, 45, 5, 67, 0, 0, 45, 46, 5, 79, 0, 0, 46, 180, 5, 80, 0,
		0, 47, 48, 5, 67, 0, 0, 48, 49, 5, 82, 0, 0, 49, 180, 5, 67, 0, 0, 50,
		51, 5, 67, 0, 0, 51, 52, 5, 90, 0, 0, 52, 180, 5, 75, 0, 0, 53, 54, 5,
		68, 0, 0, 54, 55, 5, 75, 0, 0, 55, 180, 5, 75, 0, 0, 56, 57, 5, 69, 0,
		0, 57, 58, 5, 71, 0, 0, 58, 180, 5, 80, 0, 0, 59, 60, 5, 69, 0, 0, 60,
		61, 5, 85, 0, 0, 61, 180, 5, 82, 0, 0, 62, 63, 5, 71, 0, 0, 63, 64, 5,
		66, 0, 0, 64, 180, 5, 80, 0, 0, 65, 66, 5, 71, 0, 0, 66, 67, 5, 84, 0,
		0, 67, 180, 5, 81, 0, 0, 68, 69, 5, 72, 0, 0, 69, 70, 5, 75, 0, 0, 70,
		180, 5, 68, 0, 0, 71, 72, 5, 72, 0, 0, 72, 73, 5, 78, 0, 0, 73, 180, 5,
		76, 0, 0, 74, 75, 5, 72, 0, 0, 75, 76, 5, 82, 0, 0, 76, 180, 5, 75, 0,
		0, 77, 78, 5, 72, 0, 0, 78, 79, 5, 85, 0, 0, 79, 180, 5, 70, 0, 0, 80,
		81, 5, 73, 0, 0, 81, 82, 5, 68, 0, 0, 82, 180, 5, 82, 0, 0, 83, 84, 5,
		73, 0, 0, 84, 85, 5, 76, 0, 0, 85, 180, 5, 83, 0, 0, 86, 87, 5, 73, 0,
		0, 87, 88, 5, 78, 0, 0, 88, 180, 5, 82, 0, 0, 89, 90, 5, 74, 0, 0, 90,
		91, 5, 79, 0, 0, 91, 180, 5, 68, 0, 0, 92, 93, 5, 74, 0, 0, 93, 94, 5,
		80, 0, 0, 94, 180, 5, 89, 0, 0, 95, 96, 5, 75, 0, 0, 96, 97, 5, 82, 0,
		0, 97, 180, 5, 87, 0, 0, 98, 99, 5, 75, 0, 0, 99, 100, 5, 87, 0, 0, 100,
		180, 5, 68, 0, 0, 101, 102, 5, 76, 0, 0, 102, 103, 5, 66, 0, 0, 103, 180,
		5, 80, 0, 0, 104, 105, 5, 76, 0, 0, 105, 106, 5, 75, 0, 0, 106, 180, 5,
		82, 0, 0, 107, 108, 5, 77, 0, 0, 108, 109, 5, 68, 0, 0, 109, 180, 5, 76,
		0, 0, 110, 111, 5, 77, 0, 0, 111, 112, 5, 88, 0, 0, 112, 180, 5, 78, 0,
		0, 113, 114, 5, 77, 0, 0, 114, 115, 5, 89, 0, 0, 115, 180, 5, 82, 0, 0,
		116, 117, 5, 78, 0, 0, 117, 118, 5, 73, 0, 0, 118, 180, 5, 79, 0, 0, 119,
		120, 5, 78, 0, 0, 120, 121, 5, 90, 0, 0, 121, 180, 5, 68, 0, 0, 122, 123,
		5, 79, 0, 0, 123, 124, 5, 77, 0, 0, 124, 180, 5, 82, 0, 0, 125, 126, 5,
		80, 0, 0, 126, 127, 5, 69, 0, 0, 127, 180, 5, 78, 0, 0, 128, 129, 5, 80,
		0, 0, 129, 130, 5, 72, 0, 0, 130, 180, 5, 80, 0, 0, 131, 132, 5, 80, 0,
		0, 132, 133, 5, 75, 0, 0, 133, 180, 5, 82, 0, 0, 134, 135, 5, 80, 0, 0,
		135, 136, 5, 76, 0, 0, 136, 180, 5, 78, 0, 0, 137, 138, 5, 81, 0, 0, 138,
		139, 5, 65, 0, 0, 139, 180, 5, 82, 0, 0, 140, 141, 5, 82, 0, 0, 141, 142,
		5, 79, 0, 0, 142, 180, 5, 78, 0, 0, 143, 144, 5, 82, 0, 0, 144, 145, 5,
		85, 0, 0, 145, 180, 5, 66, 0, 0, 146, 147, 5, 83, 0, 0, 147, 148, 5, 65,
		0, 0, 148, 180, 5, 82, 0, 0, 149, 150, 5, 83, 0, 0, 150, 151, 5, 69, 0,
		0, 151, 180, 5, 75, 0, 0, 152, 153, 5, 83, 0, 0, 153, 154, 5, 71, 0, 0,
		154, 180, 5, 68, 0, 0, 155, 156, 5, 84, 0, 0, 156, 157, 5, 72, 0, 0, 157,
		180, 5, 66, 0, 0, 158, 159, 5, 84, 0, 0, 159, 160, 5, 82, 0, 0, 160, 180,
		5, 89, 0, 0, 161, 162, 5, 84, 0, 0, 162, 163, 5, 87, 0, 0, 163, 180, 5,
		68, 0, 0, 164, 165, 5, 85, 0, 0, 165, 166, 5, 65, 0, 0, 166, 180, 5, 72,
		0, 0, 167, 168, 5, 85, 0, 0, 168, 169, 5, 89, 0, 0, 169, 180, 5, 85, 0,
		0, 170, 171, 5, 86, 0, 0, 171, 172, 5, 69, 0, 0, 172, 180, 5, 70, 0, 0,
		173, 174, 5, 86, 0, 0, 174, 175, 5, 78, 0, 0, 175, 180, 5, 68, 0, 0, 176,
		177, 5, 90, 0, 0, 177, 178, 5, 65, 0, 0, 178, 180, 5, 82, 0, 0, 179, 11,
		1, 0, 0, 0, 179, 14, 1, 0, 0, 0, 179, 17, 1, 0, 0, 0, 179, 20, 1, 0, 0,
		0, 179, 23, 1, 0, 0, 0, 179, 26, 1, 0, 0, 0, 179, 29, 1, 0, 0, 0, 179,
		32, 1, 0, 0, 0, 179, 35, 1, 0, 0, 0, 179, 38, 1, 0, 0, 0, 179, 41, 1, 0,
		0, 0, 179, 44, 1, 0, 0, 0, 179, 47, 1, 0, 0, 0, 179, 50, 1, 0, 0, 0, 179,
		53, 1, 0, 0, 0, 179, 56, 1, 0, 0, 0, 179, 59, 1, 0, 0, 0, 179, 62, 1, 0,
		0, 0, 179, 65, 1, 0, 0, 0, 179, 68, 1, 0, 0, 0, 179, 71, 1, 0, 0, 0, 179,
		74, 1, 0, 0, 0, 179, 77, 1, 0, 0, 0, 179, 80, 1, 0, 0, 0, 179, 83, 1, 0,
		0, 0, 179, 86, 1, 0, 0, 0, 179, 89, 1, 0, 0, 0, 179, 92, 1, 0, 0, 0, 179,
		95, 1, 0, 0, 0, 179, 98, 1, 0, 0, 0, 179, 101, 1, 0, 0, 0, 179, 104, 1,
		0, 0, 0, 179, 107, 1, 0, 0, 0, 179, 110, 1, 0, 0, 0, 179, 113, 1, 0, 0,
		0, 179, 116, 1, 0, 0, 0, 179, 119, 1, 0, 0, 0, 179, 122, 1, 0, 0, 0, 179,
		125, 1, 0, 0, 0, 179, 128, 1, 0, 0, 0, 179, 131, 1, 0, 0, 0, 179, 134,
		1, 0, 0, 0, 179, 137, 1, 0, 0, 0, 179, 140, 1, 0, 0, 0, 179, 143, 1, 0,
		0, 0, 179, 146, 1, 0, 0, 0, 179, 149, 1, 0, 0, 0, 179, 152, 1, 0, 0, 0,
		179, 155, 1, 0, 0, 0, 179, 158, 1, 0, 0, 0, 179, 161, 1, 0, 0, 0, 179,
		164, 1, 0, 0, 0, 179, 167, 1, 0, 0, 0, 179, 170, 1, 0, 0, 0, 179, 173,
		1, 0, 0, 0, 179, 176, 1, 0, 0, 0, 180, 2, 1, 0, 0, 0, 181, 182, 5, 36,
		0, 0, 182, 183, 5, 126, 0, 0, 183, 184, 5, 36, 0, 0, 184, 4, 1, 0, 0, 0,
		185, 186, 5, 32, 0, 0, 186, 187, 5, 33, 0, 0, 187, 188, 5, 36, 0, 0, 188,
		189, 5, 126, 0, 0, 189, 190, 5, 36, 0, 0, 190, 191, 5, 35, 0, 0, 191, 192,
		5, 37, 0, 0, 192, 193, 5, 35, 0, 0, 193, 194, 5, 38, 0, 0, 194, 195, 5,
		36, 0, 0, 195, 196, 5, 38, 0, 0, 196, 197, 5, 33, 0, 0, 197, 198, 5, 32,
		0, 0, 198, 6, 1, 0, 0, 0, 199, 201, 7, 0, 0, 0, 200, 199, 1, 0, 0, 0, 201,
		202, 1, 0, 0, 0, 202, 200, 1, 0, 0, 0, 202, 203, 1, 0, 0, 0, 203, 204,
		1, 0, 0, 0, 204, 205, 6, 3, 0, 0, 205, 8, 1, 0, 0, 0, 3, 0, 179, 202, 1,
		6, 0, 0,
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

// CurrencyLexerInit initializes any static state used to implement CurrencyLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewCurrencyLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func CurrencyLexerInit() {
	staticData := &currencylexerLexerStaticData
	staticData.once.Do(currencylexerLexerInit)
}

// NewCurrencyLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewCurrencyLexer(input antlr.CharStream) *CurrencyLexer {
	CurrencyLexerInit()
	l := new(CurrencyLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &currencylexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Currency.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CurrencyLexer tokens.
const (
	CurrencyLexerCURRENCY = 1
	CurrencyLexerOWNKEY   = 2
	CurrencyLexerSPLITKEY = 3
	CurrencyLexerWS       = 4
)