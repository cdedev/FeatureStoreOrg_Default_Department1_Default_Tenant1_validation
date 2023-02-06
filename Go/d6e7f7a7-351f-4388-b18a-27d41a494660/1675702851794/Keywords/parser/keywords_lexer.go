// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675702851794/Keywords.g4 by ANTLR 4.10.1. DO NOT EDIT.

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

type KeywordsLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var keywordslexerLexerStaticData struct {
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

func keywordslexerLexerInit() {
	staticData := &keywordslexerLexerStaticData
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
		"", "KEYWORDS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"KEYWORDS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 4, 148, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 3, 0, 122, 8, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 3, 4, 3, 143, 8, 3, 11, 3, 12, 3, 144, 1, 3, 1, 3, 0, 0, 4, 1, 1,
		3, 2, 5, 3, 7, 4, 1, 0, 1, 3, 0, 9, 10, 13, 13, 32, 32, 155, 0, 1, 1, 0,
		0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 1, 9, 1, 0,
		0, 0, 3, 123, 1, 0, 0, 0, 5, 127, 1, 0, 0, 0, 7, 142, 1, 0, 0, 0, 9, 10,
		3, 5, 2, 0, 10, 121, 3, 3, 1, 0, 11, 12, 5, 55, 0, 0, 12, 13, 5, 50, 0,
		0, 13, 14, 5, 48, 0, 0, 14, 15, 5, 66, 0, 0, 15, 16, 5, 72, 0, 0, 16, 17,
		5, 65, 0, 0, 17, 18, 5, 50, 0, 0, 18, 19, 5, 50, 0, 0, 19, 20, 5, 71, 0,
		0, 20, 21, 5, 82, 0, 0, 21, 22, 5, 48, 0, 0, 22, 23, 5, 48, 0, 0, 23, 24,
		5, 51, 0, 0, 24, 25, 5, 50, 0, 0, 25, 122, 5, 48, 0, 0, 26, 27, 5, 70,
		0, 0, 27, 28, 5, 111, 0, 0, 28, 29, 5, 111, 0, 0, 29, 30, 5, 100, 0, 0,
		30, 31, 5, 32, 0, 0, 31, 32, 5, 97, 0, 0, 32, 33, 5, 105, 0, 0, 33, 34,
		5, 100, 0, 0, 34, 35, 5, 44, 0, 0, 35, 36, 5, 77, 0, 0, 36, 37, 5, 117,
		0, 0, 37, 38, 5, 108, 0, 0, 38, 39, 5, 116, 0, 0, 39, 40, 5, 105, 0, 0,
		40, 41, 5, 121, 0, 0, 41, 42, 5, 101, 0, 0, 42, 43, 5, 97, 0, 0, 43, 122,
		5, 114, 0, 0, 44, 45, 5, 77, 0, 0, 45, 46, 5, 117, 0, 0, 46, 47, 5, 108,
		0, 0, 47, 48, 5, 116, 0, 0, 48, 49, 5, 105, 0, 0, 49, 50, 5, 121, 0, 0,
		50, 51, 5, 101, 0, 0, 51, 52, 5, 97, 0, 0, 52, 122, 5, 114, 0, 0, 53, 54,
		5, 77, 0, 0, 54, 55, 5, 117, 0, 0, 55, 56, 5, 108, 0, 0, 56, 57, 5, 116,
		0, 0, 57, 58, 5, 105, 0, 0, 58, 59, 5, 121, 0, 0, 59, 60, 5, 101, 0, 0,
		60, 61, 5, 97, 0, 0, 61, 62, 5, 114, 0, 0, 62, 63, 5, 44, 0, 0, 63, 64,
		5, 67, 0, 0, 64, 65, 5, 79, 0, 0, 65, 66, 5, 86, 0, 0, 66, 67, 5, 73, 0,
		0, 67, 68, 5, 68, 0, 0, 68, 69, 5, 45, 0, 0, 69, 70, 5, 49, 0, 0, 70, 71,
		5, 57, 0, 0, 71, 72, 5, 32, 0, 0, 72, 73, 5, 97, 0, 0, 73, 74, 5, 115,
		0, 0, 74, 75, 5, 115, 0, 0, 75, 76, 5, 105, 0, 0, 76, 77, 5, 115, 0, 0,
		77, 78, 5, 116, 0, 0, 78, 79, 5, 97, 0, 0, 79, 80, 5, 110, 0, 0, 80, 81,
		5, 99, 0, 0, 81, 122, 5, 101, 0, 0, 82, 83, 5, 82, 0, 0, 83, 84, 5, 97,
		0, 0, 84, 85, 5, 112, 0, 0, 85, 86, 5, 105, 0, 0, 86, 87, 5, 100, 0, 0,
		87, 88, 5, 32, 0, 0, 88, 89, 5, 82, 0, 0, 89, 90, 5, 101, 0, 0, 90, 91,
		5, 115, 0, 0, 91, 92, 5, 112, 0, 0, 92, 93, 5, 111, 0, 0, 93, 94, 5, 110,
		0, 0, 94, 95, 5, 115, 0, 0, 95, 122, 5, 101, 0, 0, 96, 97, 5, 82, 0, 0,
		97, 98, 5, 97, 0, 0, 98, 99, 5, 112, 0, 0, 99, 100, 5, 105, 0, 0, 100,
		101, 5, 100, 0, 0, 101, 102, 5, 32, 0, 0, 102, 103, 5, 114, 0, 0, 103,
		104, 5, 101, 0, 0, 104, 105, 5, 115, 0, 0, 105, 106, 5, 112, 0, 0, 106,
		107, 5, 111, 0, 0, 107, 108, 5, 110, 0, 0, 108, 109, 5, 115, 0, 0, 109,
		122, 5, 101, 0, 0, 110, 111, 5, 85, 0, 0, 111, 112, 5, 83, 0, 0, 112, 113,
		5, 65, 0, 0, 113, 114, 5, 73, 0, 0, 114, 115, 5, 68, 0, 0, 115, 116, 5,
		47, 0, 0, 116, 117, 5, 66, 0, 0, 117, 118, 5, 72, 0, 0, 118, 122, 5, 65,
		0, 0, 119, 120, 5, 110, 0, 0, 120, 122, 5, 97, 0, 0, 121, 11, 1, 0, 0,
		0, 121, 26, 1, 0, 0, 0, 121, 44, 1, 0, 0, 0, 121, 53, 1, 0, 0, 0, 121,
		82, 1, 0, 0, 0, 121, 96, 1, 0, 0, 0, 121, 110, 1, 0, 0, 0, 121, 119, 1,
		0, 0, 0, 122, 2, 1, 0, 0, 0, 123, 124, 5, 36, 0, 0, 124, 125, 5, 126, 0,
		0, 125, 126, 5, 36, 0, 0, 126, 4, 1, 0, 0, 0, 127, 128, 5, 32, 0, 0, 128,
		129, 5, 33, 0, 0, 129, 130, 5, 36, 0, 0, 130, 131, 5, 126, 0, 0, 131, 132,
		5, 36, 0, 0, 132, 133, 5, 35, 0, 0, 133, 134, 5, 37, 0, 0, 134, 135, 5,
		35, 0, 0, 135, 136, 5, 38, 0, 0, 136, 137, 5, 36, 0, 0, 137, 138, 5, 38,
		0, 0, 138, 139, 5, 33, 0, 0, 139, 140, 5, 32, 0, 0, 140, 6, 1, 0, 0, 0,
		141, 143, 7, 0, 0, 0, 142, 141, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144,
		142, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146, 147,
		6, 3, 0, 0, 147, 8, 1, 0, 0, 0, 3, 0, 121, 144, 1, 6, 0, 0,
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

// KeywordsLexerInit initializes any static state used to implement KeywordsLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewKeywordsLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func KeywordsLexerInit() {
	staticData := &keywordslexerLexerStaticData
	staticData.once.Do(keywordslexerLexerInit)
}

// NewKeywordsLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewKeywordsLexer(input antlr.CharStream) *KeywordsLexer {
	KeywordsLexerInit()
	l := new(KeywordsLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &keywordslexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Keywords.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// KeywordsLexer tokens.
const (
	KeywordsLexerKEYWORDS = 1
	KeywordsLexerOWNKEY   = 2
	KeywordsLexerSPLITKEY = 3
	KeywordsLexerWS       = 4
)
