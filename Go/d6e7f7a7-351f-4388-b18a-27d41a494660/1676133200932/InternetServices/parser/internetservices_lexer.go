// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676133200932/InternetServices.g4 by ANTLR 4.10.1. DO NOT EDIT.

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

type InternetServicesLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var internetserviceslexerLexerStaticData struct {
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

func internetserviceslexerLexerInit() {
	staticData := &internetserviceslexerLexerStaticData
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
		"", "INTERNETSERVICES", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"INTERNETSERVICES", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 4, 48, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 3,
		0, 22, 8, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 4, 3, 43, 8, 3, 11,
		3, 12, 3, 44, 1, 3, 1, 3, 0, 0, 4, 1, 1, 3, 2, 5, 3, 7, 4, 1, 0, 1, 3,
		0, 9, 10, 13, 13, 32, 32, 51, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5,
		1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 1, 9, 1, 0, 0, 0, 3, 23, 1, 0, 0, 0, 5, 27,
		1, 0, 0, 0, 7, 42, 1, 0, 0, 0, 9, 10, 3, 5, 2, 0, 10, 21, 3, 3, 1, 0, 11,
		12, 5, 51, 0, 0, 12, 22, 5, 71, 0, 0, 13, 14, 5, 52, 0, 0, 14, 22, 5, 71,
		0, 0, 15, 16, 5, 78, 0, 0, 16, 17, 5, 111, 0, 0, 17, 18, 5, 110, 0, 0,
		18, 22, 5, 101, 0, 0, 19, 20, 5, 110, 0, 0, 20, 22, 5, 97, 0, 0, 21, 11,
		1, 0, 0, 0, 21, 13, 1, 0, 0, 0, 21, 15, 1, 0, 0, 0, 21, 19, 1, 0, 0, 0,
		22, 2, 1, 0, 0, 0, 23, 24, 5, 36, 0, 0, 24, 25, 5, 126, 0, 0, 25, 26, 5,
		36, 0, 0, 26, 4, 1, 0, 0, 0, 27, 28, 5, 32, 0, 0, 28, 29, 5, 33, 0, 0,
		29, 30, 5, 36, 0, 0, 30, 31, 5, 126, 0, 0, 31, 32, 5, 36, 0, 0, 32, 33,
		5, 35, 0, 0, 33, 34, 5, 37, 0, 0, 34, 35, 5, 35, 0, 0, 35, 36, 5, 38, 0,
		0, 36, 37, 5, 36, 0, 0, 37, 38, 5, 38, 0, 0, 38, 39, 5, 33, 0, 0, 39, 40,
		5, 32, 0, 0, 40, 6, 1, 0, 0, 0, 41, 43, 7, 0, 0, 0, 42, 41, 1, 0, 0, 0,
		43, 44, 1, 0, 0, 0, 44, 42, 1, 0, 0, 0, 44, 45, 1, 0, 0, 0, 45, 46, 1,
		0, 0, 0, 46, 47, 6, 3, 0, 0, 47, 8, 1, 0, 0, 0, 3, 0, 21, 44, 1, 6, 0,
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

// InternetServicesLexerInit initializes any static state used to implement InternetServicesLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewInternetServicesLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func InternetServicesLexerInit() {
	staticData := &internetserviceslexerLexerStaticData
	staticData.once.Do(internetserviceslexerLexerInit)
}

// NewInternetServicesLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewInternetServicesLexer(input antlr.CharStream) *InternetServicesLexer {
	InternetServicesLexerInit()
	l := new(InternetServicesLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &internetserviceslexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "InternetServices.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// InternetServicesLexer tokens.
const (
	InternetServicesLexerINTERNETSERVICES = 1
	InternetServicesLexerOWNKEY           = 2
	InternetServicesLexerSPLITKEY         = 3
	InternetServicesLexerWS               = 4
)
