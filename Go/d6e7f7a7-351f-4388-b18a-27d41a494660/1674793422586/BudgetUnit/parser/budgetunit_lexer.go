// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674793422586/BudgetUnit.g4 by ANTLR 4.10.1. DO NOT EDIT.

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

type BudgetUnitLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var budgetunitlexerLexerStaticData struct {
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

func budgetunitlexerLexerInit() {
	staticData := &budgetunitlexerLexerStaticData
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
		"", "BUDGETUNIT", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"BUDGETUNIT", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 4, 57, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 31, 8, 0, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 4, 3, 52, 8, 3, 11, 3, 12, 3, 53, 1,
		3, 1, 3, 0, 0, 4, 1, 1, 3, 2, 5, 3, 7, 4, 1, 0, 1, 3, 0, 9, 10, 13, 13,
		32, 32, 59, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7,
		1, 0, 0, 0, 1, 9, 1, 0, 0, 0, 3, 32, 1, 0, 0, 0, 5, 36, 1, 0, 0, 0, 7,
		51, 1, 0, 0, 0, 9, 10, 3, 5, 2, 0, 10, 30, 3, 3, 1, 0, 11, 12, 5, 65, 0,
		0, 12, 13, 5, 110, 0, 0, 13, 14, 5, 110, 0, 0, 14, 15, 5, 117, 0, 0, 15,
		16, 5, 97, 0, 0, 16, 31, 5, 108, 0, 0, 17, 18, 5, 72, 0, 0, 18, 19, 5,
		111, 0, 0, 19, 20, 5, 117, 0, 0, 20, 21, 5, 114, 0, 0, 21, 22, 5, 108,
		0, 0, 22, 31, 5, 121, 0, 0, 23, 24, 5, 77, 0, 0, 24, 25, 5, 111, 0, 0,
		25, 26, 5, 110, 0, 0, 26, 27, 5, 116, 0, 0, 27, 28, 5, 104, 0, 0, 28, 29,
		5, 108, 0, 0, 29, 31, 5, 121, 0, 0, 30, 11, 1, 0, 0, 0, 30, 17, 1, 0, 0,
		0, 30, 23, 1, 0, 0, 0, 31, 2, 1, 0, 0, 0, 32, 33, 5, 36, 0, 0, 33, 34,
		5, 126, 0, 0, 34, 35, 5, 36, 0, 0, 35, 4, 1, 0, 0, 0, 36, 37, 5, 32, 0,
		0, 37, 38, 5, 33, 0, 0, 38, 39, 5, 36, 0, 0, 39, 40, 5, 126, 0, 0, 40,
		41, 5, 36, 0, 0, 41, 42, 5, 35, 0, 0, 42, 43, 5, 37, 0, 0, 43, 44, 5, 35,
		0, 0, 44, 45, 5, 38, 0, 0, 45, 46, 5, 36, 0, 0, 46, 47, 5, 38, 0, 0, 47,
		48, 5, 33, 0, 0, 48, 49, 5, 32, 0, 0, 49, 6, 1, 0, 0, 0, 50, 52, 7, 0,
		0, 0, 51, 50, 1, 0, 0, 0, 52, 53, 1, 0, 0, 0, 53, 51, 1, 0, 0, 0, 53, 54,
		1, 0, 0, 0, 54, 55, 1, 0, 0, 0, 55, 56, 6, 3, 0, 0, 56, 8, 1, 0, 0, 0,
		3, 0, 30, 53, 1, 6, 0, 0,
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

// BudgetUnitLexerInit initializes any static state used to implement BudgetUnitLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewBudgetUnitLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func BudgetUnitLexerInit() {
	staticData := &budgetunitlexerLexerStaticData
	staticData.once.Do(budgetunitlexerLexerInit)
}

// NewBudgetUnitLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewBudgetUnitLexer(input antlr.CharStream) *BudgetUnitLexer {
	BudgetUnitLexerInit()
	l := new(BudgetUnitLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &budgetunitlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "BudgetUnit.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// BudgetUnitLexer tokens.
const (
	BudgetUnitLexerBUDGETUNIT = 1
	BudgetUnitLexerOWNKEY     = 2
	BudgetUnitLexerSPLITKEY   = 3
	BudgetUnitLexerWS         = 4
)
