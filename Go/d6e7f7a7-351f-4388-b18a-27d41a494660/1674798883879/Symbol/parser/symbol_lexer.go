// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674798883879/Symbol.g4 by ANTLR 4.10.1. DO NOT EDIT.

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

type SymbolLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var symbollexerLexerStaticData struct {
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

func symbollexerLexerInit() {
	staticData := &symbollexerLexerStaticData
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
		"", "SYMBOL", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"SYMBOL", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 4, 289, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1,
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
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 3, 0, 263, 8, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 4, 3, 284,
		8, 3, 11, 3, 12, 3, 285, 1, 3, 1, 3, 0, 0, 4, 1, 1, 3, 2, 5, 3, 7, 4, 1,
		0, 1, 3, 0, 9, 10, 13, 13, 32, 32, 322, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0,
		0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 1, 9, 1, 0, 0, 0, 3, 264, 1, 0,
		0, 0, 5, 268, 1, 0, 0, 0, 7, 283, 1, 0, 0, 0, 9, 10, 3, 5, 2, 0, 10, 262,
		3, 3, 1, 0, 11, 12, 5, 65, 0, 0, 12, 13, 5, 85, 0, 0, 13, 14, 5, 66, 0,
		0, 14, 15, 5, 65, 0, 0, 15, 16, 5, 78, 0, 0, 16, 263, 5, 75, 0, 0, 17,
		18, 5, 65, 0, 0, 18, 19, 5, 88, 0, 0, 19, 20, 5, 73, 0, 0, 20, 21, 5, 83,
		0, 0, 21, 22, 5, 66, 0, 0, 22, 23, 5, 65, 0, 0, 23, 24, 5, 78, 0, 0, 24,
		263, 5, 75, 0, 0, 25, 26, 5, 66, 0, 0, 26, 27, 5, 65, 0, 0, 27, 28, 5,
		78, 0, 0, 28, 29, 5, 68, 0, 0, 29, 30, 5, 72, 0, 0, 30, 31, 5, 65, 0, 0,
		31, 32, 5, 78, 0, 0, 32, 33, 5, 66, 0, 0, 33, 34, 5, 78, 0, 0, 34, 263,
		5, 75, 0, 0, 35, 36, 5, 66, 0, 0, 36, 37, 5, 65, 0, 0, 37, 38, 5, 78, 0,
		0, 38, 39, 5, 75, 0, 0, 39, 40, 5, 66, 0, 0, 40, 41, 5, 65, 0, 0, 41, 42,
		5, 82, 0, 0, 42, 43, 5, 79, 0, 0, 43, 44, 5, 68, 0, 0, 44, 263, 5, 65,
		0, 0, 45, 46, 5, 66, 0, 0, 46, 47, 5, 65, 0, 0, 47, 48, 5, 78, 0, 0, 48,
		49, 5, 75, 0, 0, 49, 50, 5, 73, 0, 0, 50, 51, 5, 78, 0, 0, 51, 52, 5, 68,
		0, 0, 52, 53, 5, 73, 0, 0, 53, 263, 5, 65, 0, 0, 54, 55, 5, 67, 0, 0, 55,
		56, 5, 65, 0, 0, 56, 57, 5, 78, 0, 0, 57, 58, 5, 66, 0, 0, 58, 263, 5,
		75, 0, 0, 59, 60, 5, 67, 0, 0, 60, 61, 5, 69, 0, 0, 61, 62, 5, 78, 0, 0,
		62, 63, 5, 84, 0, 0, 63, 64, 5, 82, 0, 0, 64, 65, 5, 65, 0, 0, 65, 66,
		5, 76, 0, 0, 66, 67, 5, 66, 0, 0, 67, 263, 5, 75, 0, 0, 68, 69, 5, 67,
		0, 0, 69, 70, 5, 83, 0, 0, 70, 71, 5, 66, 0, 0, 71, 72, 5, 66, 0, 0, 72,
		73, 5, 65, 0, 0, 73, 74, 5, 78, 0, 0, 74, 263, 5, 75, 0, 0, 75, 76, 5,
		67, 0, 0, 76, 77, 5, 85, 0, 0, 77, 263, 5, 66, 0, 0, 78, 79, 5, 68, 0,
		0, 79, 80, 5, 67, 0, 0, 80, 81, 5, 66, 0, 0, 81, 82, 5, 66, 0, 0, 82, 83,
		5, 65, 0, 0, 83, 84, 5, 78, 0, 0, 84, 263, 5, 75, 0, 0, 85, 86, 5, 69,
		0, 0, 86, 87, 5, 81, 0, 0, 87, 88, 5, 85, 0, 0, 88, 89, 5, 73, 0, 0, 89,
		90, 5, 84, 0, 0, 90, 91, 5, 65, 0, 0, 91, 92, 5, 83, 0, 0, 92, 93, 5, 66,
		0, 0, 93, 94, 5, 78, 0, 0, 94, 263, 5, 75, 0, 0, 95, 96, 5, 70, 0, 0, 96,
		97, 5, 69, 0, 0, 97, 98, 5, 68, 0, 0, 98, 99, 5, 69, 0, 0, 99, 100, 5,
		82, 0, 0, 100, 101, 5, 65, 0, 0, 101, 102, 5, 76, 0, 0, 102, 103, 5, 66,
		0, 0, 103, 104, 5, 78, 0, 0, 104, 263, 5, 75, 0, 0, 105, 106, 5, 72, 0,
		0, 106, 107, 5, 68, 0, 0, 107, 108, 5, 70, 0, 0, 108, 263, 5, 67, 0, 0,
		109, 110, 5, 73, 0, 0, 110, 111, 5, 67, 0, 0, 111, 112, 5, 73, 0, 0, 112,
		113, 5, 67, 0, 0, 113, 114, 5, 73, 0, 0, 114, 115, 5, 66, 0, 0, 115, 116,
		5, 65, 0, 0, 116, 117, 5, 78, 0, 0, 117, 263, 5, 75, 0, 0, 118, 119, 5,
		73, 0, 0, 119, 120, 5, 68, 0, 0, 120, 121, 5, 66, 0, 0, 121, 263, 5, 73,
		0, 0, 122, 123, 5, 73, 0, 0, 123, 124, 5, 68, 0, 0, 124, 125, 5, 70, 0,
		0, 125, 126, 5, 67, 0, 0, 126, 127, 5, 66, 0, 0, 127, 128, 5, 65, 0, 0,
		128, 129, 5, 78, 0, 0, 129, 263, 5, 75, 0, 0, 130, 131, 5, 73, 0, 0, 131,
		132, 5, 68, 0, 0, 132, 133, 5, 70, 0, 0, 133, 134, 5, 67, 0, 0, 134, 135,
		5, 66, 0, 0, 135, 136, 5, 65, 0, 0, 136, 137, 5, 78, 0, 0, 137, 263, 5,
		75, 0, 0, 138, 139, 5, 73, 0, 0, 139, 140, 5, 68, 0, 0, 140, 141, 5, 70,
		0, 0, 141, 142, 5, 67, 0, 0, 142, 143, 5, 70, 0, 0, 143, 144, 5, 73, 0,
		0, 144, 145, 5, 82, 0, 0, 145, 146, 5, 83, 0, 0, 146, 147, 5, 84, 0, 0,
		147, 263, 5, 66, 0, 0, 148, 149, 5, 73, 0, 0, 149, 150, 5, 78, 0, 0, 150,
		151, 5, 68, 0, 0, 151, 152, 5, 73, 0, 0, 152, 153, 5, 65, 0, 0, 153, 154,
		5, 78, 0, 0, 154, 263, 5, 66, 0, 0, 155, 156, 5, 73, 0, 0, 156, 157, 5,
		78, 0, 0, 157, 158, 5, 68, 0, 0, 158, 159, 5, 85, 0, 0, 159, 160, 5, 83,
		0, 0, 160, 161, 5, 73, 0, 0, 161, 162, 5, 78, 0, 0, 162, 163, 5, 68, 0,
		0, 163, 164, 5, 66, 0, 0, 164, 263, 5, 75, 0, 0, 165, 166, 5, 73, 0, 0,
		166, 167, 5, 79, 0, 0, 167, 263, 5, 66, 0, 0, 168, 169, 5, 75, 0, 0, 169,
		170, 5, 65, 0, 0, 170, 171, 5, 82, 0, 0, 171, 172, 5, 85, 0, 0, 172, 173,
		5, 82, 0, 0, 173, 174, 5, 86, 0, 0, 174, 175, 5, 89, 0, 0, 175, 176, 5,
		83, 0, 0, 176, 177, 5, 89, 0, 0, 177, 263, 5, 65, 0, 0, 178, 179, 5, 75,
		0, 0, 179, 180, 5, 79, 0, 0, 180, 181, 5, 84, 0, 0, 181, 182, 5, 65, 0,
		0, 182, 183, 5, 75, 0, 0, 183, 184, 5, 66, 0, 0, 184, 185, 5, 65, 0, 0,
		185, 186, 5, 78, 0, 0, 186, 263, 5, 75, 0, 0, 187, 188, 5, 77, 0, 0, 188,
		189, 5, 65, 0, 0, 189, 190, 5, 72, 0, 0, 190, 191, 5, 65, 0, 0, 191, 192,
		5, 66, 0, 0, 192, 193, 5, 65, 0, 0, 193, 194, 5, 78, 0, 0, 194, 263, 5,
		75, 0, 0, 195, 196, 5, 80, 0, 0, 196, 197, 5, 78, 0, 0, 197, 263, 5, 66,
		0, 0, 198, 199, 5, 80, 0, 0, 199, 200, 5, 83, 0, 0, 200, 263, 5, 66, 0,
		0, 201, 202, 5, 82, 0, 0, 202, 203, 5, 66, 0, 0, 203, 204, 5, 76, 0, 0,
		204, 205, 5, 66, 0, 0, 205, 206, 5, 65, 0, 0, 206, 207, 5, 78, 0, 0, 207,
		263, 5, 75, 0, 0, 208, 209, 5, 83, 0, 0, 209, 210, 5, 66, 0, 0, 210, 211,
		5, 73, 0, 0, 211, 263, 5, 78, 0, 0, 212, 213, 5, 83, 0, 0, 213, 214, 5,
		79, 0, 0, 214, 215, 5, 85, 0, 0, 215, 216, 5, 84, 0, 0, 216, 217, 5, 72,
		0, 0, 217, 218, 5, 66, 0, 0, 218, 219, 5, 65, 0, 0, 219, 220, 5, 78, 0,
		0, 220, 263, 5, 75, 0, 0, 221, 222, 5, 83, 0, 0, 222, 223, 5, 85, 0, 0,
		223, 224, 5, 82, 0, 0, 224, 225, 5, 89, 0, 0, 225, 226, 5, 79, 0, 0, 226,
		227, 5, 68, 0, 0, 227, 228, 5, 65, 0, 0, 228, 263, 5, 89, 0, 0, 229, 230,
		5, 85, 0, 0, 230, 231, 5, 67, 0, 0, 231, 232, 5, 79, 0, 0, 232, 233, 5,
		66, 0, 0, 233, 234, 5, 65, 0, 0, 234, 235, 5, 78, 0, 0, 235, 263, 5, 75,
		0, 0, 236, 237, 5, 85, 0, 0, 237, 238, 5, 74, 0, 0, 238, 239, 5, 74, 0,
		0, 239, 240, 5, 73, 0, 0, 240, 241, 5, 86, 0, 0, 241, 242, 5, 65, 0, 0,
		242, 243, 5, 78, 0, 0, 243, 244, 5, 83, 0, 0, 244, 245, 5, 70, 0, 0, 245,
		263, 5, 66, 0, 0, 246, 247, 5, 85, 0, 0, 247, 248, 5, 78, 0, 0, 248, 249,
		5, 73, 0, 0, 249, 250, 5, 79, 0, 0, 250, 251, 5, 78, 0, 0, 251, 252, 5,
		66, 0, 0, 252, 253, 5, 65, 0, 0, 253, 254, 5, 78, 0, 0, 254, 263, 5, 75,
		0, 0, 255, 256, 5, 89, 0, 0, 256, 257, 5, 69, 0, 0, 257, 258, 5, 83, 0,
		0, 258, 259, 5, 66, 0, 0, 259, 260, 5, 65, 0, 0, 260, 261, 5, 78, 0, 0,
		261, 263, 5, 75, 0, 0, 262, 11, 1, 0, 0, 0, 262, 17, 1, 0, 0, 0, 262, 25,
		1, 0, 0, 0, 262, 35, 1, 0, 0, 0, 262, 45, 1, 0, 0, 0, 262, 54, 1, 0, 0,
		0, 262, 59, 1, 0, 0, 0, 262, 68, 1, 0, 0, 0, 262, 75, 1, 0, 0, 0, 262,
		78, 1, 0, 0, 0, 262, 85, 1, 0, 0, 0, 262, 95, 1, 0, 0, 0, 262, 105, 1,
		0, 0, 0, 262, 109, 1, 0, 0, 0, 262, 118, 1, 0, 0, 0, 262, 122, 1, 0, 0,
		0, 262, 130, 1, 0, 0, 0, 262, 138, 1, 0, 0, 0, 262, 148, 1, 0, 0, 0, 262,
		155, 1, 0, 0, 0, 262, 165, 1, 0, 0, 0, 262, 168, 1, 0, 0, 0, 262, 178,
		1, 0, 0, 0, 262, 187, 1, 0, 0, 0, 262, 195, 1, 0, 0, 0, 262, 198, 1, 0,
		0, 0, 262, 201, 1, 0, 0, 0, 262, 208, 1, 0, 0, 0, 262, 212, 1, 0, 0, 0,
		262, 221, 1, 0, 0, 0, 262, 229, 1, 0, 0, 0, 262, 236, 1, 0, 0, 0, 262,
		246, 1, 0, 0, 0, 262, 255, 1, 0, 0, 0, 263, 2, 1, 0, 0, 0, 264, 265, 5,
		36, 0, 0, 265, 266, 5, 126, 0, 0, 266, 267, 5, 36, 0, 0, 267, 4, 1, 0,
		0, 0, 268, 269, 5, 32, 0, 0, 269, 270, 5, 33, 0, 0, 270, 271, 5, 36, 0,
		0, 271, 272, 5, 126, 0, 0, 272, 273, 5, 36, 0, 0, 273, 274, 5, 35, 0, 0,
		274, 275, 5, 37, 0, 0, 275, 276, 5, 35, 0, 0, 276, 277, 5, 38, 0, 0, 277,
		278, 5, 36, 0, 0, 278, 279, 5, 38, 0, 0, 279, 280, 5, 33, 0, 0, 280, 281,
		5, 32, 0, 0, 281, 6, 1, 0, 0, 0, 282, 284, 7, 0, 0, 0, 283, 282, 1, 0,
		0, 0, 284, 285, 1, 0, 0, 0, 285, 283, 1, 0, 0, 0, 285, 286, 1, 0, 0, 0,
		286, 287, 1, 0, 0, 0, 287, 288, 6, 3, 0, 0, 288, 8, 1, 0, 0, 0, 3, 0, 262,
		285, 1, 6, 0, 0,
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

// SymbolLexerInit initializes any static state used to implement SymbolLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewSymbolLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func SymbolLexerInit() {
	staticData := &symbollexerLexerStaticData
	staticData.once.Do(symbollexerLexerInit)
}

// NewSymbolLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewSymbolLexer(input antlr.CharStream) *SymbolLexer {
	SymbolLexerInit()
	l := new(SymbolLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &symbollexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Symbol.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SymbolLexer tokens.
const (
	SymbolLexerSYMBOL   = 1
	SymbolLexerOWNKEY   = 2
	SymbolLexerSPLITKEY = 3
	SymbolLexerWS       = 4
)