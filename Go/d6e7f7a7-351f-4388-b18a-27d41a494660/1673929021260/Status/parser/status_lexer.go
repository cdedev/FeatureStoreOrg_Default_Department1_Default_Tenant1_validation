// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673929021260/Status.g4 by ANTLR 4.10.1. DO NOT EDIT.

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

type StatusLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var statuslexerLexerStaticData struct {
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

func statuslexerLexerInit() {
	staticData := &statuslexerLexerStaticData
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
		"", "STATUS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"STATUS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 4, 217, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1,
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
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 3, 0, 191, 8, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 4, 3, 212,
		8, 3, 11, 3, 12, 3, 213, 1, 3, 1, 3, 0, 0, 4, 1, 1, 3, 2, 5, 3, 7, 4, 1,
		0, 1, 3, 0, 9, 10, 13, 13, 32, 32, 233, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0,
		0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 1, 9, 1, 0, 0, 0, 3, 192, 1, 0,
		0, 0, 5, 196, 1, 0, 0, 0, 7, 211, 1, 0, 0, 0, 9, 10, 3, 5, 2, 0, 10, 190,
		3, 3, 1, 0, 11, 12, 5, 97, 0, 0, 12, 13, 5, 99, 0, 0, 13, 14, 5, 116, 0,
		0, 14, 15, 5, 105, 0, 0, 15, 16, 5, 118, 0, 0, 16, 191, 5, 101, 0, 0, 17,
		18, 5, 100, 0, 0, 18, 19, 5, 101, 0, 0, 19, 20, 5, 101, 0, 0, 20, 21, 5,
		112, 0, 0, 21, 22, 5, 32, 0, 0, 22, 23, 5, 115, 0, 0, 23, 24, 5, 112, 0,
		0, 24, 25, 5, 97, 0, 0, 25, 26, 5, 99, 0, 0, 26, 191, 5, 101, 0, 0, 27,
		28, 5, 101, 0, 0, 28, 29, 5, 110, 0, 0, 29, 30, 5, 103, 0, 0, 30, 31, 5,
		105, 0, 0, 31, 32, 5, 110, 0, 0, 32, 33, 5, 101, 0, 0, 33, 34, 5, 101,
		0, 0, 34, 35, 5, 114, 0, 0, 35, 36, 5, 105, 0, 0, 36, 37, 5, 110, 0, 0,
		37, 191, 5, 103, 0, 0, 38, 39, 5, 102, 0, 0, 39, 40, 5, 97, 0, 0, 40, 41,
		5, 105, 0, 0, 41, 42, 5, 108, 0, 0, 42, 43, 5, 117, 0, 0, 43, 44, 5, 114,
		0, 0, 44, 191, 5, 101, 0, 0, 45, 46, 5, 102, 0, 0, 46, 47, 5, 111, 0, 0,
		47, 48, 5, 114, 0, 0, 48, 49, 5, 32, 0, 0, 49, 50, 5, 111, 0, 0, 50, 51,
		5, 114, 0, 0, 51, 52, 5, 98, 0, 0, 52, 53, 5, 99, 0, 0, 53, 54, 5, 111,
		0, 0, 54, 55, 5, 109, 0, 0, 55, 191, 5, 109, 0, 0, 56, 57, 5, 105, 0, 0,
		57, 58, 5, 100, 0, 0, 58, 59, 5, 32, 0, 0, 59, 60, 5, 114, 0, 0, 60, 61,
		5, 101, 0, 0, 61, 62, 5, 115, 0, 0, 62, 63, 5, 101, 0, 0, 63, 64, 5, 114,
		0, 0, 64, 65, 5, 118, 0, 0, 65, 66, 5, 101, 0, 0, 66, 191, 5, 100, 0, 0,
		67, 68, 5, 105, 0, 0, 68, 69, 5, 110, 0, 0, 69, 70, 5, 97, 0, 0, 70, 71,
		5, 99, 0, 0, 71, 72, 5, 116, 0, 0, 72, 73, 5, 105, 0, 0, 73, 74, 5, 118,
		0, 0, 74, 191, 5, 101, 0, 0, 75, 76, 5, 108, 0, 0, 76, 77, 5, 97, 0, 0,
		77, 78, 5, 117, 0, 0, 78, 79, 5, 110, 0, 0, 79, 80, 5, 99, 0, 0, 80, 81,
		5, 104, 0, 0, 81, 82, 5, 32, 0, 0, 82, 83, 5, 102, 0, 0, 83, 84, 5, 97,
		0, 0, 84, 85, 5, 105, 0, 0, 85, 86, 5, 108, 0, 0, 86, 87, 5, 101, 0, 0,
		87, 191, 5, 100, 0, 0, 88, 89, 5, 110, 0, 0, 89, 90, 5, 111, 0, 0, 90,
		91, 5, 110, 0, 0, 91, 92, 5, 45, 0, 0, 92, 93, 5, 97, 0, 0, 93, 94, 5,
		109, 0, 0, 94, 95, 5, 97, 0, 0, 95, 96, 5, 116, 0, 0, 96, 97, 5, 101, 0,
		0, 97, 98, 5, 117, 0, 0, 98, 191, 5, 114, 0, 0, 99, 100, 5, 110, 0, 0,
		100, 101, 5, 111, 0, 0, 101, 102, 5, 110, 0, 0, 102, 103, 5, 45, 0, 0,
		103, 104, 5, 111, 0, 0, 104, 105, 5, 112, 0, 0, 105, 106, 5, 101, 0, 0,
		106, 107, 5, 114, 0, 0, 107, 108, 5, 97, 0, 0, 108, 109, 5, 116, 0, 0,
		109, 110, 5, 105, 0, 0, 110, 111, 5, 111, 0, 0, 111, 112, 5, 110, 0, 0,
		112, 113, 5, 97, 0, 0, 113, 191, 5, 108, 0, 0, 114, 115, 5, 111, 0, 0,
		115, 116, 5, 112, 0, 0, 116, 117, 5, 101, 0, 0, 117, 118, 5, 114, 0, 0,
		118, 119, 5, 97, 0, 0, 119, 120, 5, 116, 0, 0, 120, 121, 5, 105, 0, 0,
		121, 122, 5, 111, 0, 0, 122, 123, 5, 110, 0, 0, 123, 124, 5, 97, 0, 0,
		124, 191, 5, 108, 0, 0, 125, 126, 5, 112, 0, 0, 126, 127, 5, 114, 0, 0,
		127, 128, 5, 111, 0, 0, 128, 129, 5, 106, 0, 0, 129, 130, 5, 101, 0, 0,
		130, 131, 5, 99, 0, 0, 131, 132, 5, 116, 0, 0, 132, 133, 5, 32, 0, 0, 133,
		134, 5, 99, 0, 0, 134, 135, 5, 97, 0, 0, 135, 136, 5, 110, 0, 0, 136, 137,
		5, 99, 0, 0, 137, 138, 5, 101, 0, 0, 138, 139, 5, 108, 0, 0, 139, 140,
		5, 101, 0, 0, 140, 191, 5, 100, 0, 0, 141, 142, 5, 114, 0, 0, 142, 143,
		5, 101, 0, 0, 143, 144, 5, 45, 0, 0, 144, 145, 5, 101, 0, 0, 145, 146,
		5, 110, 0, 0, 146, 147, 5, 116, 0, 0, 147, 148, 5, 101, 0, 0, 148, 149,
		5, 114, 0, 0, 149, 150, 5, 101, 0, 0, 150, 191, 5, 100, 0, 0, 151, 152,
		5, 116, 0, 0, 152, 153, 5, 111, 0, 0, 153, 154, 5, 32, 0, 0, 154, 155,
		5, 98, 0, 0, 155, 156, 5, 101, 0, 0, 156, 157, 5, 32, 0, 0, 157, 158, 5,
		108, 0, 0, 158, 159, 5, 97, 0, 0, 159, 160, 5, 117, 0, 0, 160, 161, 5,
		110, 0, 0, 161, 162, 5, 99, 0, 0, 162, 163, 5, 104, 0, 0, 163, 164, 5,
		101, 0, 0, 164, 191, 5, 100, 0, 0, 165, 166, 5, 117, 0, 0, 166, 167, 5,
		110, 0, 0, 167, 168, 5, 107, 0, 0, 168, 169, 5, 110, 0, 0, 169, 170, 5,
		111, 0, 0, 170, 171, 5, 119, 0, 0, 171, 191, 5, 110, 0, 0, 172, 173, 5,
		117, 0, 0, 173, 174, 5, 110, 0, 0, 174, 175, 5, 107, 0, 0, 175, 176, 5,
		110, 0, 0, 176, 177, 5, 111, 0, 0, 177, 178, 5, 119, 0, 0, 178, 191, 5,
		110, 0, 0, 179, 180, 5, 119, 0, 0, 180, 181, 5, 101, 0, 0, 181, 182, 5,
		97, 0, 0, 182, 183, 5, 116, 0, 0, 183, 184, 5, 104, 0, 0, 184, 185, 5,
		101, 0, 0, 185, 186, 5, 114, 0, 0, 186, 187, 5, 32, 0, 0, 187, 188, 5,
		115, 0, 0, 188, 189, 5, 97, 0, 0, 189, 191, 5, 116, 0, 0, 190, 11, 1, 0,
		0, 0, 190, 17, 1, 0, 0, 0, 190, 27, 1, 0, 0, 0, 190, 38, 1, 0, 0, 0, 190,
		45, 1, 0, 0, 0, 190, 56, 1, 0, 0, 0, 190, 67, 1, 0, 0, 0, 190, 75, 1, 0,
		0, 0, 190, 88, 1, 0, 0, 0, 190, 99, 1, 0, 0, 0, 190, 114, 1, 0, 0, 0, 190,
		125, 1, 0, 0, 0, 190, 141, 1, 0, 0, 0, 190, 151, 1, 0, 0, 0, 190, 165,
		1, 0, 0, 0, 190, 172, 1, 0, 0, 0, 190, 179, 1, 0, 0, 0, 191, 2, 1, 0, 0,
		0, 192, 193, 5, 36, 0, 0, 193, 194, 5, 126, 0, 0, 194, 195, 5, 36, 0, 0,
		195, 4, 1, 0, 0, 0, 196, 197, 5, 32, 0, 0, 197, 198, 5, 33, 0, 0, 198,
		199, 5, 36, 0, 0, 199, 200, 5, 126, 0, 0, 200, 201, 5, 36, 0, 0, 201, 202,
		5, 35, 0, 0, 202, 203, 5, 37, 0, 0, 203, 204, 5, 35, 0, 0, 204, 205, 5,
		38, 0, 0, 205, 206, 5, 36, 0, 0, 206, 207, 5, 38, 0, 0, 207, 208, 5, 33,
		0, 0, 208, 209, 5, 32, 0, 0, 209, 6, 1, 0, 0, 0, 210, 212, 7, 0, 0, 0,
		211, 210, 1, 0, 0, 0, 212, 213, 1, 0, 0, 0, 213, 211, 1, 0, 0, 0, 213,
		214, 1, 0, 0, 0, 214, 215, 1, 0, 0, 0, 215, 216, 6, 3, 0, 0, 216, 8, 1,
		0, 0, 0, 3, 0, 190, 213, 1, 6, 0, 0,
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

// StatusLexerInit initializes any static state used to implement StatusLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewStatusLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func StatusLexerInit() {
	staticData := &statuslexerLexerStaticData
	staticData.once.Do(statuslexerLexerInit)
}

// NewStatusLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewStatusLexer(input antlr.CharStream) *StatusLexer {
	StatusLexerInit()
	l := new(StatusLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &statuslexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Status.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// StatusLexer tokens.
const (
	StatusLexerSTATUS   = 1
	StatusLexerOWNKEY   = 2
	StatusLexerSPLITKEY = 3
	StatusLexerWS       = 4
)
