// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675680683851/Material.g4 by ANTLR 4.10.1. DO NOT EDIT.

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

type MaterialLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var materiallexerLexerStaticData struct {
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

func materiallexerLexerInit() {
	staticData := &materiallexerLexerStaticData
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
		"", "MATERIAL", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"MATERIAL", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 4, 120, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 3,
		0, 94, 8, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 4, 3, 115, 8, 3,
		11, 3, 12, 3, 116, 1, 3, 1, 3, 0, 0, 4, 1, 1, 3, 2, 5, 3, 7, 4, 1, 0, 1,
		3, 0, 9, 10, 13, 13, 32, 32, 130, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0,
		5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 1, 9, 1, 0, 0, 0, 3, 95, 1, 0, 0, 0, 5,
		99, 1, 0, 0, 0, 7, 114, 1, 0, 0, 0, 9, 10, 3, 5, 2, 0, 10, 93, 3, 3, 1,
		0, 11, 12, 5, 66, 0, 0, 12, 13, 5, 114, 0, 0, 13, 14, 5, 111, 0, 0, 14,
		15, 5, 110, 0, 0, 15, 16, 5, 122, 0, 0, 16, 94, 5, 101, 0, 0, 17, 18, 5,
		67, 0, 0, 18, 19, 5, 101, 0, 0, 19, 20, 5, 114, 0, 0, 20, 21, 5, 97, 0,
		0, 21, 22, 5, 109, 0, 0, 22, 23, 5, 105, 0, 0, 23, 94, 5, 99, 0, 0, 24,
		25, 5, 71, 0, 0, 25, 26, 5, 111, 0, 0, 26, 27, 5, 108, 0, 0, 27, 28, 5,
		100, 0, 0, 28, 29, 5, 47, 0, 0, 29, 30, 5, 83, 0, 0, 30, 31, 5, 116, 0,
		0, 31, 32, 5, 101, 0, 0, 32, 33, 5, 101, 0, 0, 33, 94, 5, 108, 0, 0, 34,
		35, 5, 80, 0, 0, 35, 36, 5, 108, 0, 0, 36, 37, 5, 97, 0, 0, 37, 38, 5,
		116, 0, 0, 38, 39, 5, 105, 0, 0, 39, 40, 5, 110, 0, 0, 40, 41, 5, 117,
		0, 0, 41, 94, 5, 109, 0, 0, 42, 43, 5, 82, 0, 0, 43, 44, 5, 101, 0, 0,
		44, 45, 5, 100, 0, 0, 45, 46, 5, 32, 0, 0, 46, 47, 5, 103, 0, 0, 47, 48,
		5, 111, 0, 0, 48, 49, 5, 108, 0, 0, 49, 94, 5, 100, 0, 0, 50, 51, 5, 82,
		0, 0, 51, 52, 5, 111, 0, 0, 52, 53, 5, 115, 0, 0, 53, 54, 5, 101, 0, 0,
		54, 55, 5, 32, 0, 0, 55, 56, 5, 103, 0, 0, 56, 57, 5, 111, 0, 0, 57, 58,
		5, 108, 0, 0, 58, 94, 5, 100, 0, 0, 59, 60, 5, 83, 0, 0, 60, 61, 5, 105,
		0, 0, 61, 62, 5, 108, 0, 0, 62, 63, 5, 118, 0, 0, 63, 64, 5, 101, 0, 0,
		64, 94, 5, 114, 0, 0, 65, 66, 5, 83, 0, 0, 66, 67, 5, 116, 0, 0, 67, 68,
		5, 101, 0, 0, 68, 69, 5, 101, 0, 0, 69, 94, 5, 108, 0, 0, 70, 71, 5, 87,
		0, 0, 71, 72, 5, 104, 0, 0, 72, 73, 5, 105, 0, 0, 73, 74, 5, 116, 0, 0,
		74, 75, 5, 101, 0, 0, 75, 76, 5, 32, 0, 0, 76, 77, 5, 103, 0, 0, 77, 78,
		5, 111, 0, 0, 78, 79, 5, 108, 0, 0, 79, 94, 5, 100, 0, 0, 80, 81, 5, 89,
		0, 0, 81, 82, 5, 101, 0, 0, 82, 83, 5, 108, 0, 0, 83, 84, 5, 108, 0, 0,
		84, 85, 5, 111, 0, 0, 85, 86, 5, 119, 0, 0, 86, 87, 5, 32, 0, 0, 87, 88,
		5, 103, 0, 0, 88, 89, 5, 111, 0, 0, 89, 90, 5, 108, 0, 0, 90, 94, 5, 100,
		0, 0, 91, 92, 5, 110, 0, 0, 92, 94, 5, 97, 0, 0, 93, 11, 1, 0, 0, 0, 93,
		17, 1, 0, 0, 0, 93, 24, 1, 0, 0, 0, 93, 34, 1, 0, 0, 0, 93, 42, 1, 0, 0,
		0, 93, 50, 1, 0, 0, 0, 93, 59, 1, 0, 0, 0, 93, 65, 1, 0, 0, 0, 93, 70,
		1, 0, 0, 0, 93, 80, 1, 0, 0, 0, 93, 91, 1, 0, 0, 0, 94, 2, 1, 0, 0, 0,
		95, 96, 5, 36, 0, 0, 96, 97, 5, 126, 0, 0, 97, 98, 5, 36, 0, 0, 98, 4,
		1, 0, 0, 0, 99, 100, 5, 32, 0, 0, 100, 101, 5, 33, 0, 0, 101, 102, 5, 36,
		0, 0, 102, 103, 5, 126, 0, 0, 103, 104, 5, 36, 0, 0, 104, 105, 5, 35, 0,
		0, 105, 106, 5, 37, 0, 0, 106, 107, 5, 35, 0, 0, 107, 108, 5, 38, 0, 0,
		108, 109, 5, 36, 0, 0, 109, 110, 5, 38, 0, 0, 110, 111, 5, 33, 0, 0, 111,
		112, 5, 32, 0, 0, 112, 6, 1, 0, 0, 0, 113, 115, 7, 0, 0, 0, 114, 113, 1,
		0, 0, 0, 115, 116, 1, 0, 0, 0, 116, 114, 1, 0, 0, 0, 116, 117, 1, 0, 0,
		0, 117, 118, 1, 0, 0, 0, 118, 119, 6, 3, 0, 0, 119, 8, 1, 0, 0, 0, 3, 0,
		93, 116, 1, 6, 0, 0,
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

// MaterialLexerInit initializes any static state used to implement MaterialLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewMaterialLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func MaterialLexerInit() {
	staticData := &materiallexerLexerStaticData
	staticData.once.Do(materiallexerLexerInit)
}

// NewMaterialLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewMaterialLexer(input antlr.CharStream) *MaterialLexer {
	MaterialLexerInit()
	l := new(MaterialLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &materiallexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Material.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// MaterialLexer tokens.
const (
	MaterialLexerMATERIAL = 1
	MaterialLexerOWNKEY   = 2
	MaterialLexerSPLITKEY = 3
	MaterialLexerWS       = 4
)
