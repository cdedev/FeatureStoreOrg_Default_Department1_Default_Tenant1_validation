// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672985492567/Vendornoproduits.g4 by ANTLR 4.10.1. DO NOT EDIT.

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

type VendornoproduitsLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var vendornoproduitslexerLexerStaticData struct {
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

func vendornoproduitslexerLexerInit() {
	staticData := &vendornoproduitslexerLexerStaticData
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
		"", "VENDORNOPRODUITS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"VENDORNOPRODUITS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 4, 220, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1,
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
		0, 1, 0, 1, 0, 1, 0, 3, 0, 194, 8, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 3, 4, 3, 215, 8, 3, 11, 3, 12, 3, 216, 1, 3, 1, 3, 0, 0, 4, 1, 1,
		3, 2, 5, 3, 7, 4, 1, 0, 1, 3, 0, 9, 10, 13, 13, 32, 32, 255, 0, 1, 1, 0,
		0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 1, 9, 1, 0,
		0, 0, 3, 195, 1, 0, 0, 0, 5, 199, 1, 0, 0, 0, 7, 214, 1, 0, 0, 0, 9, 10,
		3, 5, 2, 0, 10, 193, 3, 3, 1, 0, 11, 12, 5, 65, 0, 0, 12, 194, 5, 73, 0,
		0, 13, 14, 5, 65, 0, 0, 14, 15, 5, 82, 0, 0, 15, 16, 5, 68, 0, 0, 16, 194,
		5, 79, 0, 0, 17, 18, 5, 65, 0, 0, 18, 19, 5, 83, 0, 0, 19, 20, 5, 65, 0,
		0, 20, 21, 5, 72, 0, 0, 21, 194, 5, 73, 0, 0, 22, 23, 5, 66, 0, 0, 23,
		24, 5, 69, 0, 0, 24, 25, 5, 71, 0, 0, 25, 194, 5, 65, 0, 0, 26, 27, 5,
		66, 0, 0, 27, 28, 5, 79, 0, 0, 28, 29, 5, 82, 0, 0, 29, 30, 5, 68, 0, 0,
		30, 194, 5, 69, 0, 0, 31, 32, 5, 66, 0, 0, 32, 33, 5, 79, 0, 0, 33, 34,
		5, 85, 0, 0, 34, 35, 5, 78, 0, 0, 35, 36, 5, 68, 0, 0, 36, 37, 5, 65, 0,
		0, 37, 38, 5, 82, 0, 0, 38, 194, 5, 89, 0, 0, 39, 40, 5, 66, 0, 0, 40,
		41, 5, 80, 0, 0, 41, 194, 5, 73, 0, 0, 42, 43, 5, 67, 0, 0, 43, 44, 5,
		79, 0, 0, 44, 45, 5, 78, 0, 0, 45, 46, 5, 65, 0, 0, 46, 47, 5, 80, 0, 0,
		47, 48, 5, 82, 0, 0, 48, 49, 5, 79, 0, 0, 49, 50, 5, 76, 0, 0, 50, 194,
		5, 69, 0, 0, 51, 52, 5, 67, 0, 0, 52, 53, 5, 83, 0, 0, 53, 194, 5, 77,
		0, 0, 54, 55, 5, 69, 0, 0, 55, 56, 5, 85, 0, 0, 56, 57, 5, 82, 0, 0, 57,
		58, 5, 79, 0, 0, 58, 59, 5, 76, 0, 0, 59, 60, 5, 73, 0, 0, 60, 61, 5, 86,
		0, 0, 61, 194, 5, 65, 0, 0, 62, 63, 5, 70, 0, 0, 63, 64, 5, 65, 0, 0, 64,
		65, 5, 76, 0, 0, 65, 66, 5, 76, 0, 0, 66, 67, 5, 79, 0, 0, 67, 194, 5,
		84, 0, 0, 68, 69, 5, 70, 0, 0, 69, 70, 5, 68, 0, 0, 70, 194, 5, 76, 0,
		0, 71, 72, 5, 70, 0, 0, 72, 73, 5, 76, 0, 0, 73, 74, 5, 69, 0, 0, 74, 75,
		5, 67, 0, 0, 75, 76, 5, 72, 0, 0, 76, 77, 5, 65, 0, 0, 77, 78, 5, 82, 0,
		0, 78, 194, 5, 68, 0, 0, 79, 80, 5, 70, 0, 0, 80, 194, 5, 78, 0, 0, 81,
		82, 5, 71, 0, 0, 82, 83, 5, 68, 0, 0, 83, 194, 5, 78, 0, 0, 84, 85, 5,
		71, 0, 0, 85, 86, 5, 69, 0, 0, 86, 87, 5, 82, 0, 0, 87, 88, 5, 79, 0, 0,
		88, 89, 5, 76, 0, 0, 89, 90, 5, 83, 0, 0, 90, 91, 5, 84, 0, 0, 91, 92,
		5, 69, 0, 0, 92, 93, 5, 73, 0, 0, 93, 94, 5, 78, 0, 0, 94, 95, 5, 69, 0,
		0, 95, 194, 5, 82, 0, 0, 96, 97, 5, 71, 0, 0, 97, 98, 5, 77, 0, 0, 98,
		99, 5, 68, 0, 0, 99, 194, 5, 70, 0, 0, 100, 101, 5, 72, 0, 0, 101, 102,
		5, 69, 0, 0, 102, 103, 5, 73, 0, 0, 103, 104, 5, 78, 0, 0, 104, 194, 5,
		90, 0, 0, 105, 106, 5, 72, 0, 0, 106, 107, 5, 73, 0, 0, 107, 108, 5, 71,
		0, 0, 108, 109, 5, 72, 0, 0, 109, 110, 5, 70, 0, 0, 110, 111, 5, 79, 0,
		0, 111, 112, 5, 82, 0, 0, 112, 194, 5, 68, 0, 0, 113, 114, 5, 74, 0, 0,
		114, 194, 5, 77, 0, 0, 115, 116, 5, 75, 0, 0, 116, 117, 5, 73, 0, 0, 117,
		118, 5, 76, 0, 0, 118, 119, 5, 67, 0, 0, 119, 120, 5, 79, 0, 0, 120, 121,
		5, 89, 0, 0, 121, 122, 5, 32, 0, 0, 122, 123, 5, 80, 0, 0, 123, 124, 5,
		65, 0, 0, 124, 125, 5, 83, 0, 0, 125, 126, 5, 84, 0, 0, 126, 127, 5, 79,
		0, 0, 127, 128, 5, 82, 0, 0, 128, 129, 5, 65, 0, 0, 129, 194, 5, 76, 0,
		0, 130, 131, 5, 76, 0, 0, 131, 132, 5, 65, 0, 0, 132, 133, 5, 77, 0, 0,
		133, 134, 5, 32, 0, 0, 134, 135, 5, 83, 0, 0, 135, 136, 5, 79, 0, 0, 136,
		137, 5, 79, 0, 0, 137, 138, 5, 78, 0, 0, 138, 139, 5, 32, 0, 0, 139, 140,
		5, 86, 0, 0, 140, 141, 5, 73, 0, 0, 141, 142, 5, 69, 0, 0, 142, 143, 5,
		84, 0, 0, 143, 144, 5, 78, 0, 0, 144, 145, 5, 65, 0, 0, 145, 194, 5, 77,
		0, 0, 146, 147, 5, 76, 0, 0, 147, 194, 5, 66, 0, 0, 148, 149, 5, 76, 0,
		0, 149, 194, 5, 87, 0, 0, 150, 151, 5, 77, 0, 0, 151, 194, 5, 71, 0, 0,
		152, 153, 5, 78, 0, 0, 153, 154, 5, 80, 0, 0, 154, 155, 5, 78, 0, 0, 155,
		194, 5, 90, 0, 0, 156, 157, 5, 78, 0, 0, 157, 158, 5, 85, 0, 0, 158, 159,
		5, 76, 0, 0, 159, 194, 5, 76, 0, 0, 160, 161, 5, 80, 0, 0, 161, 162, 5,
		70, 0, 0, 162, 163, 5, 65, 0, 0, 163, 164, 5, 78, 0, 0, 164, 165, 5, 78,
		0, 0, 165, 166, 5, 69, 0, 0, 166, 194, 5, 82, 0, 0, 167, 168, 5, 80, 0,
		0, 168, 194, 5, 73, 0, 0, 169, 170, 5, 80, 0, 0, 170, 194, 5, 76, 0, 0,
		171, 172, 5, 81, 0, 0, 172, 194, 5, 88, 0, 0, 173, 174, 5, 82, 0, 0, 174,
		175, 5, 65, 0, 0, 175, 176, 5, 70, 0, 0, 176, 177, 5, 65, 0, 0, 177, 178,
		5, 69, 0, 0, 178, 194, 5, 76, 0, 0, 179, 180, 5, 84, 0, 0, 180, 194, 5,
		65, 0, 0, 181, 182, 5, 84, 0, 0, 182, 183, 5, 65, 0, 0, 183, 184, 5, 82,
		0, 0, 184, 194, 5, 65, 0, 0, 185, 186, 5, 84, 0, 0, 186, 187, 5, 85, 0,
		0, 187, 188, 5, 82, 0, 0, 188, 189, 5, 83, 0, 0, 189, 190, 5, 65, 0, 0,
		190, 194, 5, 78, 0, 0, 191, 192, 5, 87, 0, 0, 192, 194, 5, 83, 0, 0, 193,
		11, 1, 0, 0, 0, 193, 13, 1, 0, 0, 0, 193, 17, 1, 0, 0, 0, 193, 22, 1, 0,
		0, 0, 193, 26, 1, 0, 0, 0, 193, 31, 1, 0, 0, 0, 193, 39, 1, 0, 0, 0, 193,
		42, 1, 0, 0, 0, 193, 51, 1, 0, 0, 0, 193, 54, 1, 0, 0, 0, 193, 62, 1, 0,
		0, 0, 193, 68, 1, 0, 0, 0, 193, 71, 1, 0, 0, 0, 193, 79, 1, 0, 0, 0, 193,
		81, 1, 0, 0, 0, 193, 84, 1, 0, 0, 0, 193, 96, 1, 0, 0, 0, 193, 100, 1,
		0, 0, 0, 193, 105, 1, 0, 0, 0, 193, 113, 1, 0, 0, 0, 193, 115, 1, 0, 0,
		0, 193, 130, 1, 0, 0, 0, 193, 146, 1, 0, 0, 0, 193, 148, 1, 0, 0, 0, 193,
		150, 1, 0, 0, 0, 193, 152, 1, 0, 0, 0, 193, 156, 1, 0, 0, 0, 193, 160,
		1, 0, 0, 0, 193, 167, 1, 0, 0, 0, 193, 169, 1, 0, 0, 0, 193, 171, 1, 0,
		0, 0, 193, 173, 1, 0, 0, 0, 193, 179, 1, 0, 0, 0, 193, 181, 1, 0, 0, 0,
		193, 185, 1, 0, 0, 0, 193, 191, 1, 0, 0, 0, 194, 2, 1, 0, 0, 0, 195, 196,
		5, 36, 0, 0, 196, 197, 5, 126, 0, 0, 197, 198, 5, 36, 0, 0, 198, 4, 1,
		0, 0, 0, 199, 200, 5, 32, 0, 0, 200, 201, 5, 33, 0, 0, 201, 202, 5, 36,
		0, 0, 202, 203, 5, 126, 0, 0, 203, 204, 5, 36, 0, 0, 204, 205, 5, 35, 0,
		0, 205, 206, 5, 37, 0, 0, 206, 207, 5, 35, 0, 0, 207, 208, 5, 38, 0, 0,
		208, 209, 5, 36, 0, 0, 209, 210, 5, 38, 0, 0, 210, 211, 5, 33, 0, 0, 211,
		212, 5, 32, 0, 0, 212, 6, 1, 0, 0, 0, 213, 215, 7, 0, 0, 0, 214, 213, 1,
		0, 0, 0, 215, 216, 1, 0, 0, 0, 216, 214, 1, 0, 0, 0, 216, 217, 1, 0, 0,
		0, 217, 218, 1, 0, 0, 0, 218, 219, 6, 3, 0, 0, 219, 8, 1, 0, 0, 0, 3, 0,
		193, 216, 1, 6, 0, 0,
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

// VendornoproduitsLexerInit initializes any static state used to implement VendornoproduitsLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewVendornoproduitsLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func VendornoproduitsLexerInit() {
	staticData := &vendornoproduitslexerLexerStaticData
	staticData.once.Do(vendornoproduitslexerLexerInit)
}

// NewVendornoproduitsLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewVendornoproduitsLexer(input antlr.CharStream) *VendornoproduitsLexer {
	VendornoproduitsLexerInit()
	l := new(VendornoproduitsLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &vendornoproduitslexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Vendornoproduits.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// VendornoproduitsLexer tokens.
const (
	VendornoproduitsLexerVENDORNOPRODUITS = 1
	VendornoproduitsLexerOWNKEY           = 2
	VendornoproduitsLexerSPLITKEY         = 3
	VendornoproduitsLexerWS               = 4
)
