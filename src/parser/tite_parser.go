// Code generated from tite.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // tite

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type titeParser struct {
	*antlr.BaseParser
}

var TiteParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func titeParserInit() {
	staticData := &TiteParserStaticData
	staticData.LiteralNames = []string{
		"", "':'", "','", "'*'", "'?'", "'||'", "'&&'", "'|'", "'^'", "'&'",
		"'=='", "'!='", "'==='", "'!=='", "'<'", "'>'", "'<='", "'>='", "'<<'",
		"'>>'", "'+'", "'-'", "'..'", "'/'", "'%'", "'//'", "'++'", "'--'",
		"'~'", "'!'", "'**'", "'.'", "'('", "')'", "'$'", "'#'", "'='", "'=>'",
		"'<-'", "'@'", "'*='", "'/='", "'%='", "'//='", "'**='", "'+='", "'-='",
		"'<<='", "'>>='", "'&='", "'^='", "'|='", "'{'", "'}'", "'['", "']'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "IDENTIFIER", "INT", "FLOAT", "CHAR", "STR", "WS",
		"LF",
	}
	staticData.RuleNames = []string{
		"program", "sequence", "declaration", "identifiers", "expression", "type",
		"function", "condition", "otherwise", "disjunction", "conjunction",
		"or", "xor", "and", "equality", "relation", "shift", "additive", "product",
		"factor", "power", "postfix", "primary", "tag", "assign", "literal",
		"block", "array", "input", "delim",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 62, 415, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 1, 0, 1, 0, 1, 1, 1, 1,
		1, 1, 1, 1, 5, 1, 67, 8, 1, 10, 1, 12, 1, 70, 9, 1, 1, 2, 1, 2, 1, 2, 3,
		2, 75, 8, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 5, 3, 82, 8, 3, 10, 3, 12, 3,
		85, 9, 3, 1, 4, 3, 4, 88, 8, 4, 1, 4, 1, 4, 3, 4, 92, 8, 4, 1, 4, 1, 4,
		3, 4, 96, 8, 4, 1, 4, 1, 4, 3, 4, 100, 8, 4, 1, 5, 1, 5, 3, 5, 104, 8,
		5, 1, 6, 1, 6, 3, 6, 108, 8, 6, 1, 7, 1, 7, 1, 7, 3, 7, 113, 8, 7, 1, 7,
		1, 7, 3, 7, 117, 8, 7, 1, 7, 1, 7, 3, 7, 121, 8, 7, 3, 7, 123, 8, 7, 1,
		8, 3, 8, 126, 8, 8, 1, 8, 3, 8, 129, 8, 8, 1, 8, 1, 8, 3, 8, 133, 8, 8,
		1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 142, 8, 9, 1, 9, 1, 9,
		3, 9, 146, 8, 9, 1, 9, 5, 9, 149, 8, 9, 10, 9, 12, 9, 152, 9, 9, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 159, 8, 10, 1, 10, 1, 10, 3, 10, 163,
		8, 10, 1, 10, 5, 10, 166, 8, 10, 10, 10, 12, 10, 169, 9, 10, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 3, 11, 176, 8, 11, 1, 11, 1, 11, 3, 11, 180, 8,
		11, 1, 11, 5, 11, 183, 8, 11, 10, 11, 12, 11, 186, 9, 11, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 3, 12, 193, 8, 12, 1, 12, 1, 12, 3, 12, 197, 8, 12,
		1, 12, 5, 12, 200, 8, 12, 10, 12, 12, 12, 203, 9, 12, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 3, 13, 210, 8, 13, 1, 13, 1, 13, 3, 13, 214, 8, 13, 1,
		13, 5, 13, 217, 8, 13, 10, 13, 12, 13, 220, 9, 13, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 3, 14, 227, 8, 14, 1, 14, 1, 14, 3, 14, 231, 8, 14, 1, 14,
		5, 14, 234, 8, 14, 10, 14, 12, 14, 237, 9, 14, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 3, 15, 244, 8, 15, 1, 15, 1, 15, 3, 15, 248, 8, 15, 1, 15, 5,
		15, 251, 8, 15, 10, 15, 12, 15, 254, 9, 15, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 3, 16, 261, 8, 16, 1, 16, 1, 16, 3, 16, 265, 8, 16, 1, 16, 5, 16,
		268, 8, 16, 10, 16, 12, 16, 271, 9, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 3, 17, 278, 8, 17, 1, 17, 1, 17, 3, 17, 282, 8, 17, 1, 17, 5, 17, 285,
		8, 17, 10, 17, 12, 17, 288, 9, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3,
		18, 295, 8, 18, 1, 18, 1, 18, 3, 18, 299, 8, 18, 1, 18, 5, 18, 302, 8,
		18, 10, 18, 12, 18, 305, 9, 18, 1, 19, 1, 19, 1, 19, 3, 19, 310, 8, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 317, 8, 19, 1, 20, 1, 20, 1,
		20, 3, 20, 322, 8, 20, 1, 20, 1, 20, 3, 20, 326, 8, 20, 1, 20, 1, 20, 3,
		20, 330, 8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 337, 8, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 345, 8, 21, 5, 21, 347, 8,
		21, 10, 21, 12, 21, 350, 9, 21, 1, 22, 3, 22, 353, 8, 22, 1, 22, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 361, 8, 22, 1, 23, 1, 23, 1, 24, 1,
		24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 3, 25, 373, 8, 25, 1, 26,
		1, 26, 3, 26, 377, 8, 26, 1, 26, 3, 26, 380, 8, 26, 1, 26, 3, 26, 383,
		8, 26, 1, 26, 1, 26, 1, 27, 1, 27, 3, 27, 389, 8, 27, 1, 27, 3, 27, 392,
		8, 27, 1, 27, 3, 27, 395, 8, 27, 1, 27, 1, 27, 1, 28, 1, 28, 3, 28, 401,
		8, 28, 1, 28, 3, 28, 404, 8, 28, 1, 28, 3, 28, 407, 8, 28, 1, 28, 1, 28,
		1, 29, 1, 29, 3, 29, 413, 8, 29, 1, 29, 0, 11, 18, 20, 22, 24, 26, 28,
		30, 32, 34, 36, 42, 30, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
		26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 0,
		9, 1, 0, 10, 13, 1, 0, 14, 17, 1, 0, 18, 19, 1, 0, 20, 22, 2, 0, 3, 3,
		23, 25, 2, 0, 20, 21, 26, 29, 2, 0, 31, 31, 34, 35, 1, 0, 36, 51, 2, 0,
		2, 2, 62, 62, 460, 0, 60, 1, 0, 0, 0, 2, 62, 1, 0, 0, 0, 4, 74, 1, 0, 0,
		0, 6, 78, 1, 0, 0, 0, 8, 87, 1, 0, 0, 0, 10, 103, 1, 0, 0, 0, 12, 105,
		1, 0, 0, 0, 14, 122, 1, 0, 0, 0, 16, 125, 1, 0, 0, 0, 18, 136, 1, 0, 0,
		0, 20, 153, 1, 0, 0, 0, 22, 170, 1, 0, 0, 0, 24, 187, 1, 0, 0, 0, 26, 204,
		1, 0, 0, 0, 28, 221, 1, 0, 0, 0, 30, 238, 1, 0, 0, 0, 32, 255, 1, 0, 0,
		0, 34, 272, 1, 0, 0, 0, 36, 289, 1, 0, 0, 0, 38, 316, 1, 0, 0, 0, 40, 329,
		1, 0, 0, 0, 42, 331, 1, 0, 0, 0, 44, 360, 1, 0, 0, 0, 46, 362, 1, 0, 0,
		0, 48, 364, 1, 0, 0, 0, 50, 372, 1, 0, 0, 0, 52, 374, 1, 0, 0, 0, 54, 386,
		1, 0, 0, 0, 56, 398, 1, 0, 0, 0, 58, 410, 1, 0, 0, 0, 60, 61, 3, 2, 1,
		0, 61, 1, 1, 0, 0, 0, 62, 68, 3, 4, 2, 0, 63, 64, 3, 58, 29, 0, 64, 65,
		3, 4, 2, 0, 65, 67, 1, 0, 0, 0, 66, 63, 1, 0, 0, 0, 67, 70, 1, 0, 0, 0,
		68, 66, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 3, 1, 0, 0, 0, 70, 68, 1, 0,
		0, 0, 71, 72, 3, 6, 3, 0, 72, 73, 5, 1, 0, 0, 73, 75, 1, 0, 0, 0, 74, 71,
		1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 76, 1, 0, 0, 0, 76, 77, 3, 8, 4, 0,
		77, 5, 1, 0, 0, 0, 78, 83, 3, 44, 22, 0, 79, 80, 5, 2, 0, 0, 80, 82, 3,
		44, 22, 0, 81, 79, 1, 0, 0, 0, 82, 85, 1, 0, 0, 0, 83, 81, 1, 0, 0, 0,
		83, 84, 1, 0, 0, 0, 84, 7, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 86, 88, 5, 3,
		0, 0, 87, 86, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 99, 1, 0, 0, 0, 89, 100,
		3, 10, 5, 0, 90, 92, 3, 10, 5, 0, 91, 90, 1, 0, 0, 0, 91, 92, 1, 0, 0,
		0, 92, 93, 1, 0, 0, 0, 93, 95, 3, 48, 24, 0, 94, 96, 5, 62, 0, 0, 95, 94,
		1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 98, 3, 8, 4, 0,
		98, 100, 1, 0, 0, 0, 99, 89, 1, 0, 0, 0, 99, 91, 1, 0, 0, 0, 100, 9, 1,
		0, 0, 0, 101, 104, 3, 14, 7, 0, 102, 104, 3, 12, 6, 0, 103, 101, 1, 0,
		0, 0, 103, 102, 1, 0, 0, 0, 104, 11, 1, 0, 0, 0, 105, 107, 3, 56, 28, 0,
		106, 108, 3, 10, 5, 0, 107, 106, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108,
		13, 1, 0, 0, 0, 109, 123, 3, 18, 9, 0, 110, 112, 3, 18, 9, 0, 111, 113,
		5, 62, 0, 0, 112, 111, 1, 0, 0, 0, 112, 113, 1, 0, 0, 0, 113, 114, 1, 0,
		0, 0, 114, 116, 5, 4, 0, 0, 115, 117, 5, 62, 0, 0, 116, 115, 1, 0, 0, 0,
		116, 117, 1, 0, 0, 0, 117, 120, 1, 0, 0, 0, 118, 121, 3, 8, 4, 0, 119,
		121, 3, 16, 8, 0, 120, 118, 1, 0, 0, 0, 120, 119, 1, 0, 0, 0, 121, 123,
		1, 0, 0, 0, 122, 109, 1, 0, 0, 0, 122, 110, 1, 0, 0, 0, 123, 15, 1, 0,
		0, 0, 124, 126, 3, 8, 4, 0, 125, 124, 1, 0, 0, 0, 125, 126, 1, 0, 0, 0,
		126, 128, 1, 0, 0, 0, 127, 129, 5, 62, 0, 0, 128, 127, 1, 0, 0, 0, 128,
		129, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 132, 5, 1, 0, 0, 131, 133,
		5, 62, 0, 0, 132, 131, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 134, 1, 0,
		0, 0, 134, 135, 3, 14, 7, 0, 135, 17, 1, 0, 0, 0, 136, 137, 6, 9, -1, 0,
		137, 138, 3, 20, 10, 0, 138, 150, 1, 0, 0, 0, 139, 141, 10, 1, 0, 0, 140,
		142, 5, 62, 0, 0, 141, 140, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 143,
		1, 0, 0, 0, 143, 145, 5, 5, 0, 0, 144, 146, 5, 62, 0, 0, 145, 144, 1, 0,
		0, 0, 145, 146, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 149, 3, 20, 10,
		0, 148, 139, 1, 0, 0, 0, 149, 152, 1, 0, 0, 0, 150, 148, 1, 0, 0, 0, 150,
		151, 1, 0, 0, 0, 151, 19, 1, 0, 0, 0, 152, 150, 1, 0, 0, 0, 153, 154, 6,
		10, -1, 0, 154, 155, 3, 22, 11, 0, 155, 167, 1, 0, 0, 0, 156, 158, 10,
		1, 0, 0, 157, 159, 5, 62, 0, 0, 158, 157, 1, 0, 0, 0, 158, 159, 1, 0, 0,
		0, 159, 160, 1, 0, 0, 0, 160, 162, 5, 6, 0, 0, 161, 163, 5, 62, 0, 0, 162,
		161, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164, 166,
		3, 22, 11, 0, 165, 156, 1, 0, 0, 0, 166, 169, 1, 0, 0, 0, 167, 165, 1,
		0, 0, 0, 167, 168, 1, 0, 0, 0, 168, 21, 1, 0, 0, 0, 169, 167, 1, 0, 0,
		0, 170, 171, 6, 11, -1, 0, 171, 172, 3, 24, 12, 0, 172, 184, 1, 0, 0, 0,
		173, 175, 10, 1, 0, 0, 174, 176, 5, 62, 0, 0, 175, 174, 1, 0, 0, 0, 175,
		176, 1, 0, 0, 0, 176, 177, 1, 0, 0, 0, 177, 179, 5, 7, 0, 0, 178, 180,
		5, 62, 0, 0, 179, 178, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180, 181, 1, 0,
		0, 0, 181, 183, 3, 24, 12, 0, 182, 173, 1, 0, 0, 0, 183, 186, 1, 0, 0,
		0, 184, 182, 1, 0, 0, 0, 184, 185, 1, 0, 0, 0, 185, 23, 1, 0, 0, 0, 186,
		184, 1, 0, 0, 0, 187, 188, 6, 12, -1, 0, 188, 189, 3, 26, 13, 0, 189, 201,
		1, 0, 0, 0, 190, 192, 10, 1, 0, 0, 191, 193, 5, 62, 0, 0, 192, 191, 1,
		0, 0, 0, 192, 193, 1, 0, 0, 0, 193, 194, 1, 0, 0, 0, 194, 196, 5, 8, 0,
		0, 195, 197, 5, 62, 0, 0, 196, 195, 1, 0, 0, 0, 196, 197, 1, 0, 0, 0, 197,
		198, 1, 0, 0, 0, 198, 200, 3, 26, 13, 0, 199, 190, 1, 0, 0, 0, 200, 203,
		1, 0, 0, 0, 201, 199, 1, 0, 0, 0, 201, 202, 1, 0, 0, 0, 202, 25, 1, 0,
		0, 0, 203, 201, 1, 0, 0, 0, 204, 205, 6, 13, -1, 0, 205, 206, 3, 28, 14,
		0, 206, 218, 1, 0, 0, 0, 207, 209, 10, 1, 0, 0, 208, 210, 5, 62, 0, 0,
		209, 208, 1, 0, 0, 0, 209, 210, 1, 0, 0, 0, 210, 211, 1, 0, 0, 0, 211,
		213, 5, 9, 0, 0, 212, 214, 5, 62, 0, 0, 213, 212, 1, 0, 0, 0, 213, 214,
		1, 0, 0, 0, 214, 215, 1, 0, 0, 0, 215, 217, 3, 28, 14, 0, 216, 207, 1,
		0, 0, 0, 217, 220, 1, 0, 0, 0, 218, 216, 1, 0, 0, 0, 218, 219, 1, 0, 0,
		0, 219, 27, 1, 0, 0, 0, 220, 218, 1, 0, 0, 0, 221, 222, 6, 14, -1, 0, 222,
		223, 3, 30, 15, 0, 223, 235, 1, 0, 0, 0, 224, 226, 10, 1, 0, 0, 225, 227,
		5, 62, 0, 0, 226, 225, 1, 0, 0, 0, 226, 227, 1, 0, 0, 0, 227, 228, 1, 0,
		0, 0, 228, 230, 7, 0, 0, 0, 229, 231, 5, 62, 0, 0, 230, 229, 1, 0, 0, 0,
		230, 231, 1, 0, 0, 0, 231, 232, 1, 0, 0, 0, 232, 234, 3, 30, 15, 0, 233,
		224, 1, 0, 0, 0, 234, 237, 1, 0, 0, 0, 235, 233, 1, 0, 0, 0, 235, 236,
		1, 0, 0, 0, 236, 29, 1, 0, 0, 0, 237, 235, 1, 0, 0, 0, 238, 239, 6, 15,
		-1, 0, 239, 240, 3, 32, 16, 0, 240, 252, 1, 0, 0, 0, 241, 243, 10, 1, 0,
		0, 242, 244, 5, 62, 0, 0, 243, 242, 1, 0, 0, 0, 243, 244, 1, 0, 0, 0, 244,
		245, 1, 0, 0, 0, 245, 247, 7, 1, 0, 0, 246, 248, 5, 62, 0, 0, 247, 246,
		1, 0, 0, 0, 247, 248, 1, 0, 0, 0, 248, 249, 1, 0, 0, 0, 249, 251, 3, 32,
		16, 0, 250, 241, 1, 0, 0, 0, 251, 254, 1, 0, 0, 0, 252, 250, 1, 0, 0, 0,
		252, 253, 1, 0, 0, 0, 253, 31, 1, 0, 0, 0, 254, 252, 1, 0, 0, 0, 255, 256,
		6, 16, -1, 0, 256, 257, 3, 34, 17, 0, 257, 269, 1, 0, 0, 0, 258, 260, 10,
		1, 0, 0, 259, 261, 5, 62, 0, 0, 260, 259, 1, 0, 0, 0, 260, 261, 1, 0, 0,
		0, 261, 262, 1, 0, 0, 0, 262, 264, 7, 2, 0, 0, 263, 265, 5, 62, 0, 0, 264,
		263, 1, 0, 0, 0, 264, 265, 1, 0, 0, 0, 265, 266, 1, 0, 0, 0, 266, 268,
		3, 34, 17, 0, 267, 258, 1, 0, 0, 0, 268, 271, 1, 0, 0, 0, 269, 267, 1,
		0, 0, 0, 269, 270, 1, 0, 0, 0, 270, 33, 1, 0, 0, 0, 271, 269, 1, 0, 0,
		0, 272, 273, 6, 17, -1, 0, 273, 274, 3, 36, 18, 0, 274, 286, 1, 0, 0, 0,
		275, 277, 10, 1, 0, 0, 276, 278, 5, 62, 0, 0, 277, 276, 1, 0, 0, 0, 277,
		278, 1, 0, 0, 0, 278, 279, 1, 0, 0, 0, 279, 281, 7, 3, 0, 0, 280, 282,
		5, 62, 0, 0, 281, 280, 1, 0, 0, 0, 281, 282, 1, 0, 0, 0, 282, 283, 1, 0,
		0, 0, 283, 285, 3, 36, 18, 0, 284, 275, 1, 0, 0, 0, 285, 288, 1, 0, 0,
		0, 286, 284, 1, 0, 0, 0, 286, 287, 1, 0, 0, 0, 287, 35, 1, 0, 0, 0, 288,
		286, 1, 0, 0, 0, 289, 290, 6, 18, -1, 0, 290, 291, 3, 38, 19, 0, 291, 303,
		1, 0, 0, 0, 292, 294, 10, 1, 0, 0, 293, 295, 5, 62, 0, 0, 294, 293, 1,
		0, 0, 0, 294, 295, 1, 0, 0, 0, 295, 296, 1, 0, 0, 0, 296, 298, 7, 4, 0,
		0, 297, 299, 5, 62, 0, 0, 298, 297, 1, 0, 0, 0, 298, 299, 1, 0, 0, 0, 299,
		300, 1, 0, 0, 0, 300, 302, 3, 38, 19, 0, 301, 292, 1, 0, 0, 0, 302, 305,
		1, 0, 0, 0, 303, 301, 1, 0, 0, 0, 303, 304, 1, 0, 0, 0, 304, 37, 1, 0,
		0, 0, 305, 303, 1, 0, 0, 0, 306, 317, 3, 40, 20, 0, 307, 309, 7, 5, 0,
		0, 308, 310, 5, 62, 0, 0, 309, 308, 1, 0, 0, 0, 309, 310, 1, 0, 0, 0, 310,
		311, 1, 0, 0, 0, 311, 317, 3, 38, 19, 0, 312, 313, 5, 7, 0, 0, 313, 314,
		3, 38, 19, 0, 314, 315, 5, 7, 0, 0, 315, 317, 1, 0, 0, 0, 316, 306, 1,
		0, 0, 0, 316, 307, 1, 0, 0, 0, 316, 312, 1, 0, 0, 0, 317, 39, 1, 0, 0,
		0, 318, 330, 3, 42, 21, 0, 319, 321, 3, 42, 21, 0, 320, 322, 5, 62, 0,
		0, 321, 320, 1, 0, 0, 0, 321, 322, 1, 0, 0, 0, 322, 323, 1, 0, 0, 0, 323,
		325, 5, 30, 0, 0, 324, 326, 5, 62, 0, 0, 325, 324, 1, 0, 0, 0, 325, 326,
		1, 0, 0, 0, 326, 327, 1, 0, 0, 0, 327, 328, 3, 38, 19, 0, 328, 330, 1,
		0, 0, 0, 329, 318, 1, 0, 0, 0, 329, 319, 1, 0, 0, 0, 330, 41, 1, 0, 0,
		0, 331, 332, 6, 21, -1, 0, 332, 333, 3, 44, 22, 0, 333, 348, 1, 0, 0, 0,
		334, 336, 10, 1, 0, 0, 335, 337, 5, 62, 0, 0, 336, 335, 1, 0, 0, 0, 336,
		337, 1, 0, 0, 0, 337, 344, 1, 0, 0, 0, 338, 345, 5, 26, 0, 0, 339, 345,
		5, 27, 0, 0, 340, 341, 5, 31, 0, 0, 341, 345, 5, 56, 0, 0, 342, 345, 3,
		54, 27, 0, 343, 345, 3, 56, 28, 0, 344, 338, 1, 0, 0, 0, 344, 339, 1, 0,
		0, 0, 344, 340, 1, 0, 0, 0, 344, 342, 1, 0, 0, 0, 344, 343, 1, 0, 0, 0,
		345, 347, 1, 0, 0, 0, 346, 334, 1, 0, 0, 0, 347, 350, 1, 0, 0, 0, 348,
		346, 1, 0, 0, 0, 348, 349, 1, 0, 0, 0, 349, 43, 1, 0, 0, 0, 350, 348, 1,
		0, 0, 0, 351, 353, 3, 46, 23, 0, 352, 351, 1, 0, 0, 0, 352, 353, 1, 0,
		0, 0, 353, 354, 1, 0, 0, 0, 354, 361, 5, 56, 0, 0, 355, 361, 3, 50, 25,
		0, 356, 357, 5, 32, 0, 0, 357, 358, 3, 8, 4, 0, 358, 359, 5, 33, 0, 0,
		359, 361, 1, 0, 0, 0, 360, 352, 1, 0, 0, 0, 360, 355, 1, 0, 0, 0, 360,
		356, 1, 0, 0, 0, 361, 45, 1, 0, 0, 0, 362, 363, 7, 6, 0, 0, 363, 47, 1,
		0, 0, 0, 364, 365, 7, 7, 0, 0, 365, 49, 1, 0, 0, 0, 366, 373, 5, 57, 0,
		0, 367, 373, 5, 58, 0, 0, 368, 373, 5, 59, 0, 0, 369, 373, 5, 60, 0, 0,
		370, 373, 3, 52, 26, 0, 371, 373, 3, 54, 27, 0, 372, 366, 1, 0, 0, 0, 372,
		367, 1, 0, 0, 0, 372, 368, 1, 0, 0, 0, 372, 369, 1, 0, 0, 0, 372, 370,
		1, 0, 0, 0, 372, 371, 1, 0, 0, 0, 373, 51, 1, 0, 0, 0, 374, 376, 5, 52,
		0, 0, 375, 377, 5, 62, 0, 0, 376, 375, 1, 0, 0, 0, 376, 377, 1, 0, 0, 0,
		377, 379, 1, 0, 0, 0, 378, 380, 3, 2, 1, 0, 379, 378, 1, 0, 0, 0, 379,
		380, 1, 0, 0, 0, 380, 382, 1, 0, 0, 0, 381, 383, 5, 62, 0, 0, 382, 381,
		1, 0, 0, 0, 382, 383, 1, 0, 0, 0, 383, 384, 1, 0, 0, 0, 384, 385, 5, 53,
		0, 0, 385, 53, 1, 0, 0, 0, 386, 388, 5, 54, 0, 0, 387, 389, 5, 62, 0, 0,
		388, 387, 1, 0, 0, 0, 388, 389, 1, 0, 0, 0, 389, 391, 1, 0, 0, 0, 390,
		392, 3, 2, 1, 0, 391, 390, 1, 0, 0, 0, 391, 392, 1, 0, 0, 0, 392, 394,
		1, 0, 0, 0, 393, 395, 5, 62, 0, 0, 394, 393, 1, 0, 0, 0, 394, 395, 1, 0,
		0, 0, 395, 396, 1, 0, 0, 0, 396, 397, 5, 55, 0, 0, 397, 55, 1, 0, 0, 0,
		398, 400, 5, 32, 0, 0, 399, 401, 5, 62, 0, 0, 400, 399, 1, 0, 0, 0, 400,
		401, 1, 0, 0, 0, 401, 403, 1, 0, 0, 0, 402, 404, 3, 2, 1, 0, 403, 402,
		1, 0, 0, 0, 403, 404, 1, 0, 0, 0, 404, 406, 1, 0, 0, 0, 405, 407, 5, 62,
		0, 0, 406, 405, 1, 0, 0, 0, 406, 407, 1, 0, 0, 0, 407, 408, 1, 0, 0, 0,
		408, 409, 5, 33, 0, 0, 409, 57, 1, 0, 0, 0, 410, 412, 7, 8, 0, 0, 411,
		413, 5, 62, 0, 0, 412, 411, 1, 0, 0, 0, 412, 413, 1, 0, 0, 0, 413, 59,
		1, 0, 0, 0, 67, 68, 74, 83, 87, 91, 95, 99, 103, 107, 112, 116, 120, 122,
		125, 128, 132, 141, 145, 150, 158, 162, 167, 175, 179, 184, 192, 196, 201,
		209, 213, 218, 226, 230, 235, 243, 247, 252, 260, 264, 269, 277, 281, 286,
		294, 298, 303, 309, 316, 321, 325, 329, 336, 344, 348, 352, 360, 372, 376,
		379, 382, 388, 391, 394, 400, 403, 406, 412,
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

// titeParserInit initializes any static state used to implement titeParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewtiteParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TiteParserInit() {
	staticData := &TiteParserStaticData
	staticData.once.Do(titeParserInit)
}

// NewtiteParser produces a new parser instance for the optional input antlr.TokenStream.
func NewtiteParser(input antlr.TokenStream) *titeParser {
	TiteParserInit()
	this := new(titeParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &TiteParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "tite.g4"

	return this
}

// titeParser tokens.
const (
	titeParserEOF        = antlr.TokenEOF
	titeParserT__0       = 1
	titeParserT__1       = 2
	titeParserT__2       = 3
	titeParserT__3       = 4
	titeParserT__4       = 5
	titeParserT__5       = 6
	titeParserT__6       = 7
	titeParserT__7       = 8
	titeParserT__8       = 9
	titeParserT__9       = 10
	titeParserT__10      = 11
	titeParserT__11      = 12
	titeParserT__12      = 13
	titeParserT__13      = 14
	titeParserT__14      = 15
	titeParserT__15      = 16
	titeParserT__16      = 17
	titeParserT__17      = 18
	titeParserT__18      = 19
	titeParserT__19      = 20
	titeParserT__20      = 21
	titeParserT__21      = 22
	titeParserT__22      = 23
	titeParserT__23      = 24
	titeParserT__24      = 25
	titeParserT__25      = 26
	titeParserT__26      = 27
	titeParserT__27      = 28
	titeParserT__28      = 29
	titeParserT__29      = 30
	titeParserT__30      = 31
	titeParserT__31      = 32
	titeParserT__32      = 33
	titeParserT__33      = 34
	titeParserT__34      = 35
	titeParserT__35      = 36
	titeParserT__36      = 37
	titeParserT__37      = 38
	titeParserT__38      = 39
	titeParserT__39      = 40
	titeParserT__40      = 41
	titeParserT__41      = 42
	titeParserT__42      = 43
	titeParserT__43      = 44
	titeParserT__44      = 45
	titeParserT__45      = 46
	titeParserT__46      = 47
	titeParserT__47      = 48
	titeParserT__48      = 49
	titeParserT__49      = 50
	titeParserT__50      = 51
	titeParserT__51      = 52
	titeParserT__52      = 53
	titeParserT__53      = 54
	titeParserT__54      = 55
	titeParserIDENTIFIER = 56
	titeParserINT        = 57
	titeParserFLOAT      = 58
	titeParserCHAR       = 59
	titeParserSTR        = 60
	titeParserWS         = 61
	titeParserLF         = 62
)

// titeParser rules.
const (
	titeParserRULE_program     = 0
	titeParserRULE_sequence    = 1
	titeParserRULE_declaration = 2
	titeParserRULE_identifiers = 3
	titeParserRULE_expression  = 4
	titeParserRULE_type        = 5
	titeParserRULE_function    = 6
	titeParserRULE_condition   = 7
	titeParserRULE_otherwise   = 8
	titeParserRULE_disjunction = 9
	titeParserRULE_conjunction = 10
	titeParserRULE_or          = 11
	titeParserRULE_xor         = 12
	titeParserRULE_and         = 13
	titeParserRULE_equality    = 14
	titeParserRULE_relation    = 15
	titeParserRULE_shift       = 16
	titeParserRULE_additive    = 17
	titeParserRULE_product     = 18
	titeParserRULE_factor      = 19
	titeParserRULE_power       = 20
	titeParserRULE_postfix     = 21
	titeParserRULE_primary     = 22
	titeParserRULE_tag         = 23
	titeParserRULE_assign      = 24
	titeParserRULE_literal     = 25
	titeParserRULE_block       = 26
	titeParserRULE_array       = 27
	titeParserRULE_input       = 28
	titeParserRULE_delim       = 29
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Sequence() ISequenceContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = titeParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) Sequence() ISequenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISequenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISequenceContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *titeParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, titeParserRULE_program)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(60)
		p.Sequence()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISequenceContext is an interface to support dynamic dispatch.
type ISequenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDeclaration() []IDeclarationContext
	Declaration(i int) IDeclarationContext
	AllDelim() []IDelimContext
	Delim(i int) IDelimContext

	// IsSequenceContext differentiates from other interfaces.
	IsSequenceContext()
}

type SequenceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySequenceContext() *SequenceContext {
	var p = new(SequenceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_sequence
	return p
}

func InitEmptySequenceContext(p *SequenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_sequence
}

func (*SequenceContext) IsSequenceContext() {}

func NewSequenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SequenceContext {
	var p = new(SequenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = titeParserRULE_sequence

	return p
}

func (s *SequenceContext) GetParser() antlr.Parser { return s.parser }

func (s *SequenceContext) AllDeclaration() []IDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclarationContext); ok {
			tst[i] = t.(IDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *SequenceContext) Declaration(i int) IDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *SequenceContext) AllDelim() []IDelimContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDelimContext); ok {
			len++
		}
	}

	tst := make([]IDelimContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDelimContext); ok {
			tst[i] = t.(IDelimContext)
			i++
		}
	}

	return tst
}

func (s *SequenceContext) Delim(i int) IDelimContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDelimContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDelimContext)
}

func (s *SequenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SequenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SequenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.EnterSequence(s)
	}
}

func (s *SequenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.ExitSequence(s)
	}
}

func (p *titeParser) Sequence() (localctx ISequenceContext) {
	localctx = NewSequenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, titeParserRULE_sequence)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(62)
		p.Declaration()
	}
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(63)
				p.Delim()
			}
			{
				p.SetState(64)
				p.Declaration()
			}

		}
		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	Identifiers() IIdentifiersContext

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_declaration
	return p
}

func InitEmptyDeclarationContext(p *DeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_declaration
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = titeParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DeclarationContext) Identifiers() IIdentifiersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifiersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifiersContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (p *titeParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, titeParserRULE_declaration)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(74)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(71)
			p.Identifiers()
		}
		{
			p.SetState(72)
			p.Match(titeParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(76)
		p.Expression()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentifiersContext is an interface to support dynamic dispatch.
type IIdentifiersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllPrimary() []IPrimaryContext
	Primary(i int) IPrimaryContext

	// IsIdentifiersContext differentiates from other interfaces.
	IsIdentifiersContext()
}

type IdentifiersContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifiersContext() *IdentifiersContext {
	var p = new(IdentifiersContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_identifiers
	return p
}

func InitEmptyIdentifiersContext(p *IdentifiersContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_identifiers
}

func (*IdentifiersContext) IsIdentifiersContext() {}

func NewIdentifiersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifiersContext {
	var p = new(IdentifiersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = titeParserRULE_identifiers

	return p
}

func (s *IdentifiersContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifiersContext) AllPrimary() []IPrimaryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPrimaryContext); ok {
			len++
		}
	}

	tst := make([]IPrimaryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPrimaryContext); ok {
			tst[i] = t.(IPrimaryContext)
			i++
		}
	}

	return tst
}

func (s *IdentifiersContext) Primary(i int) IPrimaryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *IdentifiersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifiersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifiersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.EnterIdentifiers(s)
	}
}

func (s *IdentifiersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.ExitIdentifiers(s)
	}
}

func (p *titeParser) Identifiers() (localctx IIdentifiersContext) {
	localctx = NewIdentifiersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, titeParserRULE_identifiers)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)
		p.Primary()
	}
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == titeParserT__1 {
		{
			p.SetState(79)
			p.Match(titeParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(80)
			p.Primary()
		}

		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext
	Assign() IAssignContext
	Expression() IExpressionContext
	LF() antlr.TerminalNode

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = titeParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ExpressionContext) Assign() IAssignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *ExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) LF() antlr.TerminalNode {
	return s.GetToken(titeParserLF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *titeParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, titeParserRULE_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == titeParserT__2 {
		{
			p.SetState(86)
			p.Match(titeParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(89)
			p.Type_()
		}

	case 2:
		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2256303472304455808) != 0 {
			{
				p.SetState(90)
				p.Type_()
			}

		}
		{
			p.SetState(93)
			p.Assign()
		}
		p.SetState(95)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == titeParserLF {
			{
				p.SetState(94)
				p.Match(titeParserLF)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(97)
			p.Expression()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Condition() IConditionContext
	Function() IFunctionContext

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = titeParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) Condition() IConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *TypeContext) Function() IFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.ExitType(s)
	}
}

func (p *titeParser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, titeParserRULE_type)
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(101)
			p.Condition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(102)
			p.Function()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Input() IInputContext
	Type_() ITypeContext

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_function
	return p
}

func InitEmptyFunctionContext(p *FunctionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_function
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = titeParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) Input() IInputContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInputContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInputContext)
}

func (s *FunctionContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.EnterFunction(s)
	}
}

func (s *FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.ExitFunction(s)
	}
}

func (p *titeParser) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, titeParserRULE_function)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		p.Input()
	}
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2256303472304455808) != 0 {
		{
			p.SetState(106)
			p.Type_()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Disjunction() IDisjunctionContext
	Expression() IExpressionContext
	Otherwise() IOtherwiseContext
	AllLF() []antlr.TerminalNode
	LF(i int) antlr.TerminalNode

	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_condition
	return p
}

func InitEmptyConditionContext(p *ConditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_condition
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = titeParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) Disjunction() IDisjunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDisjunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDisjunctionContext)
}

func (s *ConditionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConditionContext) Otherwise() IOtherwiseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOtherwiseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOtherwiseContext)
}

func (s *ConditionContext) AllLF() []antlr.TerminalNode {
	return s.GetTokens(titeParserLF)
}

func (s *ConditionContext) LF(i int) antlr.TerminalNode {
	return s.GetToken(titeParserLF, i)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.EnterCondition(s)
	}
}

func (s *ConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.ExitCondition(s)
	}
}

func (p *titeParser) Condition() (localctx IConditionContext) {
	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, titeParserRULE_condition)
	var _la int

	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(109)
			p.disjunction(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(110)
			p.disjunction(0)
		}
		p.SetState(112)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == titeParserLF {
			{
				p.SetState(111)
				p.Match(titeParserLF)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(114)
			p.Match(titeParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(116)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(115)
				p.Match(titeParserLF)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(118)
				p.Expression()
			}

		case 2:
			{
				p.SetState(119)
				p.Otherwise()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOtherwiseContext is an interface to support dynamic dispatch.
type IOtherwiseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Condition() IConditionContext
	Expression() IExpressionContext
	AllLF() []antlr.TerminalNode
	LF(i int) antlr.TerminalNode

	// IsOtherwiseContext differentiates from other interfaces.
	IsOtherwiseContext()
}

type OtherwiseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOtherwiseContext() *OtherwiseContext {
	var p = new(OtherwiseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_otherwise
	return p
}

func InitEmptyOtherwiseContext(p *OtherwiseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_otherwise
}

func (*OtherwiseContext) IsOtherwiseContext() {}

func NewOtherwiseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OtherwiseContext {
	var p = new(OtherwiseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = titeParserRULE_otherwise

	return p
}

func (s *OtherwiseContext) GetParser() antlr.Parser { return s.parser }

func (s *OtherwiseContext) Condition() IConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *OtherwiseContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *OtherwiseContext) AllLF() []antlr.TerminalNode {
	return s.GetTokens(titeParserLF)
}

func (s *OtherwiseContext) LF(i int) antlr.TerminalNode {
	return s.GetToken(titeParserLF, i)
}

func (s *OtherwiseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OtherwiseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OtherwiseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.EnterOtherwise(s)
	}
}

func (s *OtherwiseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.ExitOtherwise(s)
	}
}

func (p *titeParser) Otherwise() (localctx IOtherwiseContext) {
	localctx = NewOtherwiseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, titeParserRULE_otherwise)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2260807003212349576) != 0 {
		{
			p.SetState(124)
			p.Expression()
		}

	}
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == titeParserLF {
		{
			p.SetState(127)
			p.Match(titeParserLF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(130)
		p.Match(titeParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == titeParserLF {
		{
			p.SetState(131)
			p.Match(titeParserLF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(134)
		p.Condition()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDisjunctionContext is an interface to support dynamic dispatch.
type IDisjunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Conjunction() IConjunctionContext
	Disjunction() IDisjunctionContext
	AllLF() []antlr.TerminalNode
	LF(i int) antlr.TerminalNode

	// IsDisjunctionContext differentiates from other interfaces.
	IsDisjunctionContext()
}

type DisjunctionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDisjunctionContext() *DisjunctionContext {
	var p = new(DisjunctionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_disjunction
	return p
}

func InitEmptyDisjunctionContext(p *DisjunctionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_disjunction
}

func (*DisjunctionContext) IsDisjunctionContext() {}

func NewDisjunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DisjunctionContext {
	var p = new(DisjunctionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = titeParserRULE_disjunction

	return p
}

func (s *DisjunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *DisjunctionContext) Conjunction() IConjunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConjunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConjunctionContext)
}

func (s *DisjunctionContext) Disjunction() IDisjunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDisjunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDisjunctionContext)
}

func (s *DisjunctionContext) AllLF() []antlr.TerminalNode {
	return s.GetTokens(titeParserLF)
}

func (s *DisjunctionContext) LF(i int) antlr.TerminalNode {
	return s.GetToken(titeParserLF, i)
}

func (s *DisjunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DisjunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DisjunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.EnterDisjunction(s)
	}
}

func (s *DisjunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.ExitDisjunction(s)
	}
}

func (p *titeParser) Disjunction() (localctx IDisjunctionContext) {
	return p.disjunction(0)
}

func (p *titeParser) disjunction(_p int) (localctx IDisjunctionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewDisjunctionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IDisjunctionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 18
	p.EnterRecursionRule(localctx, 18, titeParserRULE_disjunction, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(137)
		p.conjunction(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewDisjunctionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, titeParserRULE_disjunction)
			p.SetState(139)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			p.SetState(141)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == titeParserLF {
				{
					p.SetState(140)
					p.Match(titeParserLF)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(143)
				p.Match(titeParserT__4)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(145)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == titeParserLF {
				{
					p.SetState(144)
					p.Match(titeParserLF)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(147)
				p.conjunction(0)
			}

		}
		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConjunctionContext is an interface to support dynamic dispatch.
type IConjunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Or() IOrContext
	Conjunction() IConjunctionContext
	AllLF() []antlr.TerminalNode
	LF(i int) antlr.TerminalNode

	// IsConjunctionContext differentiates from other interfaces.
	IsConjunctionContext()
}

type ConjunctionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConjunctionContext() *ConjunctionContext {
	var p = new(ConjunctionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_conjunction
	return p
}

func InitEmptyConjunctionContext(p *ConjunctionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_conjunction
}

func (*ConjunctionContext) IsConjunctionContext() {}

func NewConjunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConjunctionContext {
	var p = new(ConjunctionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = titeParserRULE_conjunction

	return p
}

func (s *ConjunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConjunctionContext) Or() IOrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrContext)
}

func (s *ConjunctionContext) Conjunction() IConjunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConjunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConjunctionContext)
}

func (s *ConjunctionContext) AllLF() []antlr.TerminalNode {
	return s.GetTokens(titeParserLF)
}

func (s *ConjunctionContext) LF(i int) antlr.TerminalNode {
	return s.GetToken(titeParserLF, i)
}

func (s *ConjunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConjunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConjunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.EnterConjunction(s)
	}
}

func (s *ConjunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.ExitConjunction(s)
	}
}

func (p *titeParser) Conjunction() (localctx IConjunctionContext) {
	return p.conjunction(0)
}

func (p *titeParser) conjunction(_p int) (localctx IConjunctionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewConjunctionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IConjunctionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 20
	p.EnterRecursionRule(localctx, 20, titeParserRULE_conjunction, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(154)
		p.or(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewConjunctionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, titeParserRULE_conjunction)
			p.SetState(156)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			p.SetState(158)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == titeParserLF {
				{
					p.SetState(157)
					p.Match(titeParserLF)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(160)
				p.Match(titeParserT__5)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(162)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == titeParserLF {
				{
					p.SetState(161)
					p.Match(titeParserLF)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(164)
				p.or(0)
			}

		}
		p.SetState(169)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOrContext is an interface to support dynamic dispatch.
type IOrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Xor() IXorContext
	Or() IOrContext
	AllLF() []antlr.TerminalNode
	LF(i int) antlr.TerminalNode

	// IsOrContext differentiates from other interfaces.
	IsOrContext()
}

type OrContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrContext() *OrContext {
	var p = new(OrContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_or
	return p
}

func InitEmptyOrContext(p *OrContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_or
}

func (*OrContext) IsOrContext() {}

func NewOrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrContext {
	var p = new(OrContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = titeParserRULE_or

	return p
}

func (s *OrContext) GetParser() antlr.Parser { return s.parser }

func (s *OrContext) Xor() IXorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IXorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IXorContext)
}

func (s *OrContext) Or() IOrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrContext)
}

func (s *OrContext) AllLF() []antlr.TerminalNode {
	return s.GetTokens(titeParserLF)
}

func (s *OrContext) LF(i int) antlr.TerminalNode {
	return s.GetToken(titeParserLF, i)
}

func (s *OrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.EnterOr(s)
	}
}

func (s *OrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.ExitOr(s)
	}
}

func (p *titeParser) Or() (localctx IOrContext) {
	return p.or(0)
}

func (p *titeParser) or(_p int) (localctx IOrContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewOrContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IOrContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 22
	p.EnterRecursionRule(localctx, 22, titeParserRULE_or, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(171)
		p.xor(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewOrContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, titeParserRULE_or)
			p.SetState(173)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			p.SetState(175)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == titeParserLF {
				{
					p.SetState(174)
					p.Match(titeParserLF)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(177)
				p.Match(titeParserT__6)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(179)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == titeParserLF {
				{
					p.SetState(178)
					p.Match(titeParserLF)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(181)
				p.xor(0)
			}

		}
		p.SetState(186)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IXorContext is an interface to support dynamic dispatch.
type IXorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	And() IAndContext
	Xor() IXorContext
	AllLF() []antlr.TerminalNode
	LF(i int) antlr.TerminalNode

	// IsXorContext differentiates from other interfaces.
	IsXorContext()
}

type XorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyXorContext() *XorContext {
	var p = new(XorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_xor
	return p
}

func InitEmptyXorContext(p *XorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_xor
}

func (*XorContext) IsXorContext() {}

func NewXorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *XorContext {
	var p = new(XorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = titeParserRULE_xor

	return p
}

func (s *XorContext) GetParser() antlr.Parser { return s.parser }

func (s *XorContext) And() IAndContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAndContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAndContext)
}

func (s *XorContext) Xor() IXorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IXorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IXorContext)
}

func (s *XorContext) AllLF() []antlr.TerminalNode {
	return s.GetTokens(titeParserLF)
}

func (s *XorContext) LF(i int) antlr.TerminalNode {
	return s.GetToken(titeParserLF, i)
}

func (s *XorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *XorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *XorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.EnterXor(s)
	}
}

func (s *XorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.ExitXor(s)
	}
}

func (p *titeParser) Xor() (localctx IXorContext) {
	return p.xor(0)
}

func (p *titeParser) xor(_p int) (localctx IXorContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewXorContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IXorContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 24
	p.EnterRecursionRule(localctx, 24, titeParserRULE_xor, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(188)
		p.and(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewXorContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, titeParserRULE_xor)
			p.SetState(190)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			p.SetState(192)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == titeParserLF {
				{
					p.SetState(191)
					p.Match(titeParserLF)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(194)
				p.Match(titeParserT__7)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(196)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == titeParserLF {
				{
					p.SetState(195)
					p.Match(titeParserLF)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(198)
				p.and(0)
			}

		}
		p.SetState(203)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAndContext is an interface to support dynamic dispatch.
type IAndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Equality() IEqualityContext
	And() IAndContext
	AllLF() []antlr.TerminalNode
	LF(i int) antlr.TerminalNode

	// IsAndContext differentiates from other interfaces.
	IsAndContext()
}

type AndContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAndContext() *AndContext {
	var p = new(AndContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_and
	return p
}

func InitEmptyAndContext(p *AndContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_and
}

func (*AndContext) IsAndContext() {}

func NewAndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AndContext {
	var p = new(AndContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = titeParserRULE_and

	return p
}

func (s *AndContext) GetParser() antlr.Parser { return s.parser }

func (s *AndContext) Equality() IEqualityContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEqualityContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEqualityContext)
}

func (s *AndContext) And() IAndContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAndContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAndContext)
}

func (s *AndContext) AllLF() []antlr.TerminalNode {
	return s.GetTokens(titeParserLF)
}

func (s *AndContext) LF(i int) antlr.TerminalNode {
	return s.GetToken(titeParserLF, i)
}

func (s *AndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.EnterAnd(s)
	}
}

func (s *AndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.ExitAnd(s)
	}
}

func (p *titeParser) And() (localctx IAndContext) {
	return p.and(0)
}

func (p *titeParser) and(_p int) (localctx IAndContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewAndContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IAndContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 26
	p.EnterRecursionRule(localctx, 26, titeParserRULE_and, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(205)
		p.equality(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewAndContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, titeParserRULE_and)
			p.SetState(207)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			p.SetState(209)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == titeParserLF {
				{
					p.SetState(208)
					p.Match(titeParserLF)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(211)
				p.Match(titeParserT__8)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(213)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == titeParserLF {
				{
					p.SetState(212)
					p.Match(titeParserLF)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(215)
				p.equality(0)
			}

		}
		p.SetState(220)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEqualityContext is an interface to support dynamic dispatch.
type IEqualityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Relation() IRelationContext
	Equality() IEqualityContext
	AllLF() []antlr.TerminalNode
	LF(i int) antlr.TerminalNode

	// IsEqualityContext differentiates from other interfaces.
	IsEqualityContext()
}

type EqualityContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqualityContext() *EqualityContext {
	var p = new(EqualityContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_equality
	return p
}

func InitEmptyEqualityContext(p *EqualityContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_equality
}

func (*EqualityContext) IsEqualityContext() {}

func NewEqualityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualityContext {
	var p = new(EqualityContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = titeParserRULE_equality

	return p
}

func (s *EqualityContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualityContext) Relation() IRelationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationContext)
}

func (s *EqualityContext) Equality() IEqualityContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEqualityContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEqualityContext)
}

func (s *EqualityContext) AllLF() []antlr.TerminalNode {
	return s.GetTokens(titeParserLF)
}

func (s *EqualityContext) LF(i int) antlr.TerminalNode {
	return s.GetToken(titeParserLF, i)
}

func (s *EqualityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.EnterEquality(s)
	}
}

func (s *EqualityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.ExitEquality(s)
	}
}

func (p *titeParser) Equality() (localctx IEqualityContext) {
	return p.equality(0)
}

func (p *titeParser) equality(_p int) (localctx IEqualityContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewEqualityContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IEqualityContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 28
	p.EnterRecursionRule(localctx, 28, titeParserRULE_equality, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(222)
		p.relation(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(235)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewEqualityContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, titeParserRULE_equality)
			p.SetState(224)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			p.SetState(226)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == titeParserLF {
				{
					p.SetState(225)
					p.Match(titeParserLF)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(228)
				_la = p.GetTokenStream().LA(1)

				if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&15360) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			p.SetState(230)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == titeParserLF {
				{
					p.SetState(229)
					p.Match(titeParserLF)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(232)
				p.relation(0)
			}

		}
		p.SetState(237)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelationContext is an interface to support dynamic dispatch.
type IRelationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Shift() IShiftContext
	Relation() IRelationContext
	AllLF() []antlr.TerminalNode
	LF(i int) antlr.TerminalNode

	// IsRelationContext differentiates from other interfaces.
	IsRelationContext()
}

type RelationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationContext() *RelationContext {
	var p = new(RelationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_relation
	return p
}

func InitEmptyRelationContext(p *RelationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_relation
}

func (*RelationContext) IsRelationContext() {}

func NewRelationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationContext {
	var p = new(RelationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = titeParserRULE_relation

	return p
}

func (s *RelationContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationContext) Shift() IShiftContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IShiftContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IShiftContext)
}

func (s *RelationContext) Relation() IRelationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationContext)
}

func (s *RelationContext) AllLF() []antlr.TerminalNode {
	return s.GetTokens(titeParserLF)
}

func (s *RelationContext) LF(i int) antlr.TerminalNode {
	return s.GetToken(titeParserLF, i)
}

func (s *RelationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.EnterRelation(s)
	}
}

func (s *RelationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.ExitRelation(s)
	}
}

func (p *titeParser) Relation() (localctx IRelationContext) {
	return p.relation(0)
}

func (p *titeParser) relation(_p int) (localctx IRelationContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewRelationContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IRelationContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 30
	p.EnterRecursionRule(localctx, 30, titeParserRULE_relation, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(239)
		p.shift(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(252)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewRelationContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, titeParserRULE_relation)
			p.SetState(241)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			p.SetState(243)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == titeParserLF {
				{
					p.SetState(242)
					p.Match(titeParserLF)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(245)
				_la = p.GetTokenStream().LA(1)

				if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&245760) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			p.SetState(247)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == titeParserLF {
				{
					p.SetState(246)
					p.Match(titeParserLF)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(249)
				p.shift(0)
			}

		}
		p.SetState(254)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IShiftContext is an interface to support dynamic dispatch.
type IShiftContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Additive() IAdditiveContext
	Shift() IShiftContext
	AllLF() []antlr.TerminalNode
	LF(i int) antlr.TerminalNode

	// IsShiftContext differentiates from other interfaces.
	IsShiftContext()
}

type ShiftContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShiftContext() *ShiftContext {
	var p = new(ShiftContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_shift
	return p
}

func InitEmptyShiftContext(p *ShiftContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_shift
}

func (*ShiftContext) IsShiftContext() {}

func NewShiftContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShiftContext {
	var p = new(ShiftContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = titeParserRULE_shift

	return p
}

func (s *ShiftContext) GetParser() antlr.Parser { return s.parser }

func (s *ShiftContext) Additive() IAdditiveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdditiveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAdditiveContext)
}

func (s *ShiftContext) Shift() IShiftContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IShiftContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IShiftContext)
}

func (s *ShiftContext) AllLF() []antlr.TerminalNode {
	return s.GetTokens(titeParserLF)
}

func (s *ShiftContext) LF(i int) antlr.TerminalNode {
	return s.GetToken(titeParserLF, i)
}

func (s *ShiftContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShiftContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ShiftContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.EnterShift(s)
	}
}

func (s *ShiftContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.ExitShift(s)
	}
}

func (p *titeParser) Shift() (localctx IShiftContext) {
	return p.shift(0)
}

func (p *titeParser) shift(_p int) (localctx IShiftContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewShiftContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IShiftContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 32
	p.EnterRecursionRule(localctx, 32, titeParserRULE_shift, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(256)
		p.additive(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(269)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewShiftContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, titeParserRULE_shift)
			p.SetState(258)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			p.SetState(260)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == titeParserLF {
				{
					p.SetState(259)
					p.Match(titeParserLF)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(262)
				_la = p.GetTokenStream().LA(1)

				if !(_la == titeParserT__17 || _la == titeParserT__18) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			p.SetState(264)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == titeParserLF {
				{
					p.SetState(263)
					p.Match(titeParserLF)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(266)
				p.additive(0)
			}

		}
		p.SetState(271)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAdditiveContext is an interface to support dynamic dispatch.
type IAdditiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Product() IProductContext
	Additive() IAdditiveContext
	AllLF() []antlr.TerminalNode
	LF(i int) antlr.TerminalNode

	// IsAdditiveContext differentiates from other interfaces.
	IsAdditiveContext()
}

type AdditiveContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdditiveContext() *AdditiveContext {
	var p = new(AdditiveContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_additive
	return p
}

func InitEmptyAdditiveContext(p *AdditiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_additive
}

func (*AdditiveContext) IsAdditiveContext() {}

func NewAdditiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdditiveContext {
	var p = new(AdditiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = titeParserRULE_additive

	return p
}

func (s *AdditiveContext) GetParser() antlr.Parser { return s.parser }

func (s *AdditiveContext) Product() IProductContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProductContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProductContext)
}

func (s *AdditiveContext) Additive() IAdditiveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdditiveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAdditiveContext)
}

func (s *AdditiveContext) AllLF() []antlr.TerminalNode {
	return s.GetTokens(titeParserLF)
}

func (s *AdditiveContext) LF(i int) antlr.TerminalNode {
	return s.GetToken(titeParserLF, i)
}

func (s *AdditiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AdditiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.EnterAdditive(s)
	}
}

func (s *AdditiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.ExitAdditive(s)
	}
}

func (p *titeParser) Additive() (localctx IAdditiveContext) {
	return p.additive(0)
}

func (p *titeParser) additive(_p int) (localctx IAdditiveContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewAdditiveContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IAdditiveContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 34
	p.EnterRecursionRule(localctx, 34, titeParserRULE_additive, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(273)
		p.product(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(286)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewAdditiveContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, titeParserRULE_additive)
			p.SetState(275)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			p.SetState(277)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == titeParserLF {
				{
					p.SetState(276)
					p.Match(titeParserLF)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(279)
				_la = p.GetTokenStream().LA(1)

				if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7340032) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			p.SetState(281)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == titeParserLF {
				{
					p.SetState(280)
					p.Match(titeParserLF)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(283)
				p.product(0)
			}

		}
		p.SetState(288)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IProductContext is an interface to support dynamic dispatch.
type IProductContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Factor() IFactorContext
	Product() IProductContext
	AllLF() []antlr.TerminalNode
	LF(i int) antlr.TerminalNode

	// IsProductContext differentiates from other interfaces.
	IsProductContext()
}

type ProductContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProductContext() *ProductContext {
	var p = new(ProductContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_product
	return p
}

func InitEmptyProductContext(p *ProductContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_product
}

func (*ProductContext) IsProductContext() {}

func NewProductContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProductContext {
	var p = new(ProductContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = titeParserRULE_product

	return p
}

func (s *ProductContext) GetParser() antlr.Parser { return s.parser }

func (s *ProductContext) Factor() IFactorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFactorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *ProductContext) Product() IProductContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProductContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProductContext)
}

func (s *ProductContext) AllLF() []antlr.TerminalNode {
	return s.GetTokens(titeParserLF)
}

func (s *ProductContext) LF(i int) antlr.TerminalNode {
	return s.GetToken(titeParserLF, i)
}

func (s *ProductContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProductContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProductContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.EnterProduct(s)
	}
}

func (s *ProductContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.ExitProduct(s)
	}
}

func (p *titeParser) Product() (localctx IProductContext) {
	return p.product(0)
}

func (p *titeParser) product(_p int) (localctx IProductContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewProductContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IProductContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 36
	p.EnterRecursionRule(localctx, 36, titeParserRULE_product, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(290)
		p.Factor()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(303)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 45, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewProductContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, titeParserRULE_product)
			p.SetState(292)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			p.SetState(294)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == titeParserLF {
				{
					p.SetState(293)
					p.Match(titeParserLF)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(296)
				_la = p.GetTokenStream().LA(1)

				if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&58720264) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			p.SetState(298)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == titeParserLF {
				{
					p.SetState(297)
					p.Match(titeParserLF)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(300)
				p.Factor()
			}

		}
		p.SetState(305)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 45, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFactorContext is an interface to support dynamic dispatch.
type IFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Power() IPowerContext
	Factor() IFactorContext
	LF() antlr.TerminalNode

	// IsFactorContext differentiates from other interfaces.
	IsFactorContext()
}

type FactorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFactorContext() *FactorContext {
	var p = new(FactorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_factor
	return p
}

func InitEmptyFactorContext(p *FactorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_factor
}

func (*FactorContext) IsFactorContext() {}

func NewFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactorContext {
	var p = new(FactorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = titeParserRULE_factor

	return p
}

func (s *FactorContext) GetParser() antlr.Parser { return s.parser }

func (s *FactorContext) Power() IPowerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPowerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPowerContext)
}

func (s *FactorContext) Factor() IFactorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFactorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *FactorContext) LF() antlr.TerminalNode {
	return s.GetToken(titeParserLF, 0)
}

func (s *FactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FactorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.EnterFactor(s)
	}
}

func (s *FactorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.ExitFactor(s)
	}
}

func (p *titeParser) Factor() (localctx IFactorContext) {
	localctx = NewFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, titeParserRULE_factor)
	var _la int

	p.SetState(316)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case titeParserT__30, titeParserT__31, titeParserT__33, titeParserT__34, titeParserT__51, titeParserT__53, titeParserIDENTIFIER, titeParserINT, titeParserFLOAT, titeParserCHAR, titeParserSTR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(306)
			p.Power()
		}

	case titeParserT__19, titeParserT__20, titeParserT__25, titeParserT__26, titeParserT__27, titeParserT__28:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(307)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1009778688) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(309)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == titeParserLF {
			{
				p.SetState(308)
				p.Match(titeParserLF)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(311)
			p.Factor()
		}

	case titeParserT__6:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(312)
			p.Match(titeParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(313)
			p.Factor()
		}
		{
			p.SetState(314)
			p.Match(titeParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPowerContext is an interface to support dynamic dispatch.
type IPowerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Postfix() IPostfixContext
	Factor() IFactorContext
	AllLF() []antlr.TerminalNode
	LF(i int) antlr.TerminalNode

	// IsPowerContext differentiates from other interfaces.
	IsPowerContext()
}

type PowerContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPowerContext() *PowerContext {
	var p = new(PowerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_power
	return p
}

func InitEmptyPowerContext(p *PowerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_power
}

func (*PowerContext) IsPowerContext() {}

func NewPowerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PowerContext {
	var p = new(PowerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = titeParserRULE_power

	return p
}

func (s *PowerContext) GetParser() antlr.Parser { return s.parser }

func (s *PowerContext) Postfix() IPostfixContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPostfixContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPostfixContext)
}

func (s *PowerContext) Factor() IFactorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFactorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *PowerContext) AllLF() []antlr.TerminalNode {
	return s.GetTokens(titeParserLF)
}

func (s *PowerContext) LF(i int) antlr.TerminalNode {
	return s.GetToken(titeParserLF, i)
}

func (s *PowerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PowerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.EnterPower(s)
	}
}

func (s *PowerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.ExitPower(s)
	}
}

func (p *titeParser) Power() (localctx IPowerContext) {
	localctx = NewPowerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, titeParserRULE_power)
	var _la int

	p.SetState(329)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 50, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(318)
			p.postfix(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(319)
			p.postfix(0)
		}
		p.SetState(321)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == titeParserLF {
			{
				p.SetState(320)
				p.Match(titeParserLF)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(323)
			p.Match(titeParserT__29)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(325)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == titeParserLF {
			{
				p.SetState(324)
				p.Match(titeParserLF)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(327)
			p.Factor()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPostfixContext is an interface to support dynamic dispatch.
type IPostfixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Primary() IPrimaryContext
	Postfix() IPostfixContext
	IDENTIFIER() antlr.TerminalNode
	Array() IArrayContext
	Input() IInputContext
	LF() antlr.TerminalNode

	// IsPostfixContext differentiates from other interfaces.
	IsPostfixContext()
}

type PostfixContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPostfixContext() *PostfixContext {
	var p = new(PostfixContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_postfix
	return p
}

func InitEmptyPostfixContext(p *PostfixContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_postfix
}

func (*PostfixContext) IsPostfixContext() {}

func NewPostfixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PostfixContext {
	var p = new(PostfixContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = titeParserRULE_postfix

	return p
}

func (s *PostfixContext) GetParser() antlr.Parser { return s.parser }

func (s *PostfixContext) Primary() IPrimaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *PostfixContext) Postfix() IPostfixContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPostfixContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPostfixContext)
}

func (s *PostfixContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(titeParserIDENTIFIER, 0)
}

func (s *PostfixContext) Array() IArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *PostfixContext) Input() IInputContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInputContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInputContext)
}

func (s *PostfixContext) LF() antlr.TerminalNode {
	return s.GetToken(titeParserLF, 0)
}

func (s *PostfixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostfixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PostfixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.EnterPostfix(s)
	}
}

func (s *PostfixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.ExitPostfix(s)
	}
}

func (p *titeParser) Postfix() (localctx IPostfixContext) {
	return p.postfix(0)
}

func (p *titeParser) postfix(_p int) (localctx IPostfixContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewPostfixContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPostfixContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 42
	p.EnterRecursionRule(localctx, 42, titeParserRULE_postfix, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(332)
		p.Primary()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(348)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 53, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewPostfixContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, titeParserRULE_postfix)
			p.SetState(334)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			p.SetState(336)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == titeParserLF {
				{
					p.SetState(335)
					p.Match(titeParserLF)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(344)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case titeParserT__25:
				{
					p.SetState(338)
					p.Match(titeParserT__25)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case titeParserT__26:
				{
					p.SetState(339)
					p.Match(titeParserT__26)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case titeParserT__30:
				{
					p.SetState(340)
					p.Match(titeParserT__30)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(341)
					p.Match(titeParserIDENTIFIER)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case titeParserT__53:
				{
					p.SetState(342)
					p.Array()
				}

			case titeParserT__31:
				{
					p.SetState(343)
					p.Input()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

		}
		p.SetState(350)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 53, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimaryContext is an interface to support dynamic dispatch.
type IPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	Tag() ITagContext
	Literal() ILiteralContext
	Expression() IExpressionContext

	// IsPrimaryContext differentiates from other interfaces.
	IsPrimaryContext()
}

type PrimaryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryContext() *PrimaryContext {
	var p = new(PrimaryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_primary
	return p
}

func InitEmptyPrimaryContext(p *PrimaryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_primary
}

func (*PrimaryContext) IsPrimaryContext() {}

func NewPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryContext {
	var p = new(PrimaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = titeParserRULE_primary

	return p
}

func (s *PrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(titeParserIDENTIFIER, 0)
}

func (s *PrimaryContext) Tag() ITagContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITagContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITagContext)
}

func (s *PrimaryContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *PrimaryContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.EnterPrimary(s)
	}
}

func (s *PrimaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.ExitPrimary(s)
	}
}

func (p *titeParser) Primary() (localctx IPrimaryContext) {
	localctx = NewPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, titeParserRULE_primary)
	var _la int

	p.SetState(360)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case titeParserT__30, titeParserT__33, titeParserT__34, titeParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(352)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&53687091200) != 0 {
			{
				p.SetState(351)
				p.Tag()
			}

		}
		{
			p.SetState(354)
			p.Match(titeParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case titeParserT__51, titeParserT__53, titeParserINT, titeParserFLOAT, titeParserCHAR, titeParserSTR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(355)
			p.Literal()
		}

	case titeParserT__31:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(356)
			p.Match(titeParserT__31)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(357)
			p.Expression()
		}
		{
			p.SetState(358)
			p.Match(titeParserT__32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITagContext is an interface to support dynamic dispatch.
type ITagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTagContext differentiates from other interfaces.
	IsTagContext()
}

type TagContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTagContext() *TagContext {
	var p = new(TagContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_tag
	return p
}

func InitEmptyTagContext(p *TagContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_tag
}

func (*TagContext) IsTagContext() {}

func NewTagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TagContext {
	var p = new(TagContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = titeParserRULE_tag

	return p
}

func (s *TagContext) GetParser() antlr.Parser { return s.parser }
func (s *TagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.EnterTag(s)
	}
}

func (s *TagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.ExitTag(s)
	}
}

func (p *titeParser) Tag() (localctx ITagContext) {
	localctx = NewTagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, titeParserRULE_tag)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(362)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&53687091200) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignContext is an interface to support dynamic dispatch.
type IAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAssignContext differentiates from other interfaces.
	IsAssignContext()
}

type AssignContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignContext() *AssignContext {
	var p = new(AssignContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_assign
	return p
}

func InitEmptyAssignContext(p *AssignContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_assign
}

func (*AssignContext) IsAssignContext() {}

func NewAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignContext {
	var p = new(AssignContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = titeParserRULE_assign

	return p
}

func (s *AssignContext) GetParser() antlr.Parser { return s.parser }
func (s *AssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.EnterAssign(s)
	}
}

func (s *AssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.ExitAssign(s)
	}
}

func (p *titeParser) Assign() (localctx IAssignContext) {
	localctx = NewAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, titeParserRULE_assign)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(364)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4503530907893760) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	CHAR() antlr.TerminalNode
	STR() antlr.TerminalNode
	Block() IBlockContext
	Array() IArrayContext

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = titeParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) INT() antlr.TerminalNode {
	return s.GetToken(titeParserINT, 0)
}

func (s *LiteralContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(titeParserFLOAT, 0)
}

func (s *LiteralContext) CHAR() antlr.TerminalNode {
	return s.GetToken(titeParserCHAR, 0)
}

func (s *LiteralContext) STR() antlr.TerminalNode {
	return s.GetToken(titeParserSTR, 0)
}

func (s *LiteralContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *LiteralContext) Array() IArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (p *titeParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, titeParserRULE_literal)
	p.SetState(372)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case titeParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(366)
			p.Match(titeParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case titeParserFLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(367)
			p.Match(titeParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case titeParserCHAR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(368)
			p.Match(titeParserCHAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case titeParserSTR:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(369)
			p.Match(titeParserSTR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case titeParserT__51:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(370)
			p.Block()
		}

	case titeParserT__53:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(371)
			p.Array()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLF() []antlr.TerminalNode
	LF(i int) antlr.TerminalNode
	Sequence() ISequenceContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = titeParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllLF() []antlr.TerminalNode {
	return s.GetTokens(titeParserLF)
}

func (s *BlockContext) LF(i int) antlr.TerminalNode {
	return s.GetToken(titeParserLF, i)
}

func (s *BlockContext) Sequence() ISequenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISequenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISequenceContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *titeParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, titeParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(374)
		p.Match(titeParserT__51)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(376)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 57, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(375)
			p.Match(titeParserLF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(379)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2260807003212349576) != 0 {
		{
			p.SetState(378)
			p.Sequence()
		}

	}
	p.SetState(382)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == titeParserLF {
		{
			p.SetState(381)
			p.Match(titeParserLF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(384)
		p.Match(titeParserT__52)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArrayContext is an interface to support dynamic dispatch.
type IArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLF() []antlr.TerminalNode
	LF(i int) antlr.TerminalNode
	Sequence() ISequenceContext

	// IsArrayContext differentiates from other interfaces.
	IsArrayContext()
}

type ArrayContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayContext() *ArrayContext {
	var p = new(ArrayContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_array
	return p
}

func InitEmptyArrayContext(p *ArrayContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_array
}

func (*ArrayContext) IsArrayContext() {}

func NewArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayContext {
	var p = new(ArrayContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = titeParserRULE_array

	return p
}

func (s *ArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayContext) AllLF() []antlr.TerminalNode {
	return s.GetTokens(titeParserLF)
}

func (s *ArrayContext) LF(i int) antlr.TerminalNode {
	return s.GetToken(titeParserLF, i)
}

func (s *ArrayContext) Sequence() ISequenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISequenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISequenceContext)
}

func (s *ArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.EnterArray(s)
	}
}

func (s *ArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.ExitArray(s)
	}
}

func (p *titeParser) Array() (localctx IArrayContext) {
	localctx = NewArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, titeParserRULE_array)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(386)
		p.Match(titeParserT__53)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(388)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 60, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(387)
			p.Match(titeParserLF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(391)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2260807003212349576) != 0 {
		{
			p.SetState(390)
			p.Sequence()
		}

	}
	p.SetState(394)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == titeParserLF {
		{
			p.SetState(393)
			p.Match(titeParserLF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(396)
		p.Match(titeParserT__54)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInputContext is an interface to support dynamic dispatch.
type IInputContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLF() []antlr.TerminalNode
	LF(i int) antlr.TerminalNode
	Sequence() ISequenceContext

	// IsInputContext differentiates from other interfaces.
	IsInputContext()
}

type InputContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInputContext() *InputContext {
	var p = new(InputContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_input
	return p
}

func InitEmptyInputContext(p *InputContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_input
}

func (*InputContext) IsInputContext() {}

func NewInputContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputContext {
	var p = new(InputContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = titeParserRULE_input

	return p
}

func (s *InputContext) GetParser() antlr.Parser { return s.parser }

func (s *InputContext) AllLF() []antlr.TerminalNode {
	return s.GetTokens(titeParserLF)
}

func (s *InputContext) LF(i int) antlr.TerminalNode {
	return s.GetToken(titeParserLF, i)
}

func (s *InputContext) Sequence() ISequenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISequenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISequenceContext)
}

func (s *InputContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InputContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.EnterInput(s)
	}
}

func (s *InputContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.ExitInput(s)
	}
}

func (p *titeParser) Input() (localctx IInputContext) {
	localctx = NewInputContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, titeParserRULE_input)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(398)
		p.Match(titeParserT__31)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(400)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 63, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(399)
			p.Match(titeParserLF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(403)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2260807003212349576) != 0 {
		{
			p.SetState(402)
			p.Sequence()
		}

	}
	p.SetState(406)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == titeParserLF {
		{
			p.SetState(405)
			p.Match(titeParserLF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(408)
		p.Match(titeParserT__32)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDelimContext is an interface to support dynamic dispatch.
type IDelimContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLF() []antlr.TerminalNode
	LF(i int) antlr.TerminalNode

	// IsDelimContext differentiates from other interfaces.
	IsDelimContext()
}

type DelimContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDelimContext() *DelimContext {
	var p = new(DelimContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_delim
	return p
}

func InitEmptyDelimContext(p *DelimContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = titeParserRULE_delim
}

func (*DelimContext) IsDelimContext() {}

func NewDelimContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DelimContext {
	var p = new(DelimContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = titeParserRULE_delim

	return p
}

func (s *DelimContext) GetParser() antlr.Parser { return s.parser }

func (s *DelimContext) AllLF() []antlr.TerminalNode {
	return s.GetTokens(titeParserLF)
}

func (s *DelimContext) LF(i int) antlr.TerminalNode {
	return s.GetToken(titeParserLF, i)
}

func (s *DelimContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DelimContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DelimContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.EnterDelim(s)
	}
}

func (s *DelimContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(titeListener); ok {
		listenerT.ExitDelim(s)
	}
}

func (p *titeParser) Delim() (localctx IDelimContext) {
	localctx = NewDelimContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, titeParserRULE_delim)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(410)
		_la = p.GetTokenStream().LA(1)

		if !(_la == titeParserT__1 || _la == titeParserLF) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(412)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == titeParserLF {
		{
			p.SetState(411)
			p.Match(titeParserLF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *titeParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 9:
		var t *DisjunctionContext = nil
		if localctx != nil {
			t = localctx.(*DisjunctionContext)
		}
		return p.Disjunction_Sempred(t, predIndex)

	case 10:
		var t *ConjunctionContext = nil
		if localctx != nil {
			t = localctx.(*ConjunctionContext)
		}
		return p.Conjunction_Sempred(t, predIndex)

	case 11:
		var t *OrContext = nil
		if localctx != nil {
			t = localctx.(*OrContext)
		}
		return p.Or_Sempred(t, predIndex)

	case 12:
		var t *XorContext = nil
		if localctx != nil {
			t = localctx.(*XorContext)
		}
		return p.Xor_Sempred(t, predIndex)

	case 13:
		var t *AndContext = nil
		if localctx != nil {
			t = localctx.(*AndContext)
		}
		return p.And_Sempred(t, predIndex)

	case 14:
		var t *EqualityContext = nil
		if localctx != nil {
			t = localctx.(*EqualityContext)
		}
		return p.Equality_Sempred(t, predIndex)

	case 15:
		var t *RelationContext = nil
		if localctx != nil {
			t = localctx.(*RelationContext)
		}
		return p.Relation_Sempred(t, predIndex)

	case 16:
		var t *ShiftContext = nil
		if localctx != nil {
			t = localctx.(*ShiftContext)
		}
		return p.Shift_Sempred(t, predIndex)

	case 17:
		var t *AdditiveContext = nil
		if localctx != nil {
			t = localctx.(*AdditiveContext)
		}
		return p.Additive_Sempred(t, predIndex)

	case 18:
		var t *ProductContext = nil
		if localctx != nil {
			t = localctx.(*ProductContext)
		}
		return p.Product_Sempred(t, predIndex)

	case 21:
		var t *PostfixContext = nil
		if localctx != nil {
			t = localctx.(*PostfixContext)
		}
		return p.Postfix_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *titeParser) Disjunction_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *titeParser) Conjunction_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *titeParser) Or_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *titeParser) Xor_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *titeParser) And_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *titeParser) Equality_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 5:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *titeParser) Relation_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 6:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *titeParser) Shift_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 7:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *titeParser) Additive_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 8:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *titeParser) Product_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 9:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *titeParser) Postfix_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 10:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
