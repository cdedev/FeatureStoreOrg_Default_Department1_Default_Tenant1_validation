// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672985218491/VendorCode.g4 by ANTLR 4.10.1. DO NOT EDIT.

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

type VendorCodeLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var vendorcodelexerLexerStaticData struct {
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

func vendorcodelexerLexerInit() {
	staticData := &vendorcodelexerLexerStaticData
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
		"", "VENDORCODE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"VENDORCODE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 4, 49, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 20, 8, 0,
		10, 0, 12, 0, 23, 9, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 4, 3,
		44, 8, 3, 11, 3, 12, 3, 45, 1, 3, 1, 3, 0, 0, 4, 1, 1, 3, 2, 5, 3, 7, 4,
		1, 0, 2, 2, 0, 45, 45, 48, 57, 3, 0, 9, 10, 13, 13, 32, 32, 50, 0, 1, 1,
		0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 1, 9, 1,
		0, 0, 0, 3, 24, 1, 0, 0, 0, 5, 28, 1, 0, 0, 0, 7, 43, 1, 0, 0, 0, 9, 10,
		3, 5, 2, 0, 10, 11, 3, 3, 1, 0, 11, 12, 5, 86, 0, 0, 12, 13, 5, 69, 0,
		0, 13, 14, 5, 78, 0, 0, 14, 15, 5, 68, 0, 0, 15, 16, 5, 79, 0, 0, 16, 17,
		5, 82, 0, 0, 17, 21, 1, 0, 0, 0, 18, 20, 7, 0, 0, 0, 19, 18, 1, 0, 0, 0,
		20, 23, 1, 0, 0, 0, 21, 19, 1, 0, 0, 0, 21, 22, 1, 0, 0, 0, 22, 2, 1, 0,
		0, 0, 23, 21, 1, 0, 0, 0, 24, 25, 5, 36, 0, 0, 25, 26, 5, 126, 0, 0, 26,
		27, 5, 36, 0, 0, 27, 4, 1, 0, 0, 0, 28, 29, 5, 32, 0, 0, 29, 30, 5, 33,
		0, 0, 30, 31, 5, 36, 0, 0, 31, 32, 5, 126, 0, 0, 32, 33, 5, 36, 0, 0, 33,
		34, 5, 35, 0, 0, 34, 35, 5, 37, 0, 0, 35, 36, 5, 35, 0, 0, 36, 37, 5, 38,
		0, 0, 37, 38, 5, 36, 0, 0, 38, 39, 5, 38, 0, 0, 39, 40, 5, 33, 0, 0, 40,
		41, 5, 32, 0, 0, 41, 6, 1, 0, 0, 0, 42, 44, 7, 1, 0, 0, 43, 42, 1, 0, 0,
		0, 44, 45, 1, 0, 0, 0, 45, 43, 1, 0, 0, 0, 45, 46, 1, 0, 0, 0, 46, 47,
		1, 0, 0, 0, 47, 48, 6, 3, 0, 0, 48, 8, 1, 0, 0, 0, 3, 0, 21, 45, 1, 6,
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

// VendorCodeLexerInit initializes any static state used to implement VendorCodeLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewVendorCodeLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func VendorCodeLexerInit() {
	staticData := &vendorcodelexerLexerStaticData
	staticData.once.Do(vendorcodelexerLexerInit)
}

// NewVendorCodeLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewVendorCodeLexer(input antlr.CharStream) *VendorCodeLexer {
	VendorCodeLexerInit()
	l := new(VendorCodeLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &vendorcodelexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "VendorCode.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// VendorCodeLexer tokens.
const (
	VendorCodeLexerVENDORCODE = 1
	VendorCodeLexerOWNKEY     = 2
	VendorCodeLexerSPLITKEY   = 3
	VendorCodeLexerWS         = 4
)
