// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674720957369/Benefits.g4 by ANTLR 4.10.1. DO NOT EDIT.

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

type BenefitsLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var benefitslexerLexerStaticData struct {
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

func benefitslexerLexerInit() {
	staticData := &benefitslexerLexerStaticData
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
		"", "BENEFITS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"BENEFITS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 4, 54, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 28, 8, 0, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 3, 4, 3, 49, 8, 3, 11, 3, 12, 3, 50, 1, 3, 1, 3, 0, 0, 4,
		1, 1, 3, 2, 5, 3, 7, 4, 1, 0, 2, 1, 0, 39, 39, 3, 0, 9, 10, 13, 13, 32,
		32, 56, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1,
		0, 0, 0, 1, 9, 1, 0, 0, 0, 3, 29, 1, 0, 0, 0, 5, 33, 1, 0, 0, 0, 7, 48,
		1, 0, 0, 0, 9, 10, 3, 5, 2, 0, 10, 27, 3, 3, 1, 0, 11, 12, 5, 78, 0, 0,
		12, 28, 5, 111, 0, 0, 13, 14, 5, 89, 0, 0, 14, 15, 5, 101, 0, 0, 15, 28,
		5, 115, 0, 0, 16, 17, 5, 68, 0, 0, 17, 18, 5, 111, 0, 0, 18, 19, 5, 110,
		0, 0, 19, 20, 1, 0, 0, 0, 20, 21, 7, 0, 0, 0, 21, 22, 5, 116, 0, 0, 22,
		23, 5, 32, 0, 0, 23, 24, 5, 107, 0, 0, 24, 25, 5, 110, 0, 0, 25, 26, 5,
		111, 0, 0, 26, 28, 5, 119, 0, 0, 27, 11, 1, 0, 0, 0, 27, 13, 1, 0, 0, 0,
		27, 16, 1, 0, 0, 0, 28, 2, 1, 0, 0, 0, 29, 30, 5, 36, 0, 0, 30, 31, 5,
		126, 0, 0, 31, 32, 5, 36, 0, 0, 32, 4, 1, 0, 0, 0, 33, 34, 5, 32, 0, 0,
		34, 35, 5, 33, 0, 0, 35, 36, 5, 36, 0, 0, 36, 37, 5, 126, 0, 0, 37, 38,
		5, 36, 0, 0, 38, 39, 5, 35, 0, 0, 39, 40, 5, 37, 0, 0, 40, 41, 5, 35, 0,
		0, 41, 42, 5, 38, 0, 0, 42, 43, 5, 36, 0, 0, 43, 44, 5, 38, 0, 0, 44, 45,
		5, 33, 0, 0, 45, 46, 5, 32, 0, 0, 46, 6, 1, 0, 0, 0, 47, 49, 7, 1, 0, 0,
		48, 47, 1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 48, 1, 0, 0, 0, 50, 51, 1,
		0, 0, 0, 51, 52, 1, 0, 0, 0, 52, 53, 6, 3, 0, 0, 53, 8, 1, 0, 0, 0, 3,
		0, 27, 50, 1, 6, 0, 0,
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

// BenefitsLexerInit initializes any static state used to implement BenefitsLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewBenefitsLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func BenefitsLexerInit() {
	staticData := &benefitslexerLexerStaticData
	staticData.once.Do(benefitslexerLexerInit)
}

// NewBenefitsLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewBenefitsLexer(input antlr.CharStream) *BenefitsLexer {
	BenefitsLexerInit()
	l := new(BenefitsLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &benefitslexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Benefits.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// BenefitsLexer tokens.
const (
	BenefitsLexerBENEFITS = 1
	BenefitsLexerOWNKEY   = 2
	BenefitsLexerSPLITKEY = 3
	BenefitsLexerWS       = 4
)
