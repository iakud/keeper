// Code generated from kds.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type kdsLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var KdsLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func kdslexerLexerInit() {
	staticData := &KdsLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'syntax'", "'import'", "'proto_go_package'", "'weak'", "'public'",
		"'package'", "'option'", "'optional'", "'repeated'", "'oneof'", "'map'",
		"'int32'", "'int64'", "'uint32'", "'uint64'", "'sint32'", "'sint64'",
		"'fixed32'", "'fixed64'", "'sfixed32'", "'sfixed64'", "'bool'", "'string'",
		"'double'", "'float'", "'bytes'", "'timestamp'", "'duration'", "'empty'",
		"'reserved'", "'to'", "'max'", "'enum'", "'entity'", "'component'",
		"'message'", "'service'", "'extend'", "'rpc'", "'stream'", "'returns'",
		"';'", "'='", "'('", "')'", "'['", "']'", "'{'", "'}'", "'<'", "'>'",
		"'.'", "','", "':'", "'+'", "'-'",
	}
	staticData.SymbolicNames = []string{
		"", "SYNTAX", "IMPORT", "PROTO_GO_PACKAGE", "WEAK", "PUBLIC", "PACKAGE",
		"OPTION", "OPTIONAL", "REPEATED", "ONEOF", "MAP", "INT32", "INT64",
		"UINT32", "UINT64", "SINT32", "SINT64", "FIXED32", "FIXED64", "SFIXED32",
		"SFIXED64", "BOOL", "STRING", "DOUBLE", "FLOAT", "BYTES", "TIMESTAMP",
		"DURATION", "EMPTY", "RESERVED", "TO", "MAX", "ENUM", "ENTITY", "COMPONENT",
		"MESSAGE", "SERVICE", "EXTEND", "RPC", "STREAM", "RETURNS", "SEMI",
		"EQ", "LP", "RP", "LB", "RB", "LC", "RC", "LT", "GT", "DOT", "COMMA",
		"COLON", "PLUS", "MINUS", "STR_LIT", "BOOL_LIT", "INT_LIT", "IDENTIFIER",
		"WS", "LINE_COMMENT", "COMMENT",
	}
	staticData.RuleNames = []string{
		"SYNTAX", "IMPORT", "PROTO_GO_PACKAGE", "WEAK", "PUBLIC", "PACKAGE",
		"OPTION", "OPTIONAL", "REPEATED", "ONEOF", "MAP", "INT32", "INT64",
		"UINT32", "UINT64", "SINT32", "SINT64", "FIXED32", "FIXED64", "SFIXED32",
		"SFIXED64", "BOOL", "STRING", "DOUBLE", "FLOAT", "BYTES", "TIMESTAMP",
		"DURATION", "EMPTY", "RESERVED", "TO", "MAX", "ENUM", "ENTITY", "COMPONENT",
		"MESSAGE", "SERVICE", "EXTEND", "RPC", "STREAM", "RETURNS", "SEMI",
		"EQ", "LP", "RP", "LB", "RB", "LC", "RC", "LT", "GT", "DOT", "COMMA",
		"COLON", "PLUS", "MINUS", "STR_LIT", "CHAR_VALUE", "HEX_ESCAPE", "OCT_ESCAPE",
		"CHAR_ESCAPE", "BOOL_LIT", "INT_LIT", "DECIMAL_LIT", "OCTAL_LIT", "HEX_LIT",
		"IDENTIFIER", "LETTER", "DECIMAL_DIGIT", "OCTAL_DIGIT", "HEX_DIGIT",
		"WS", "LINE_COMMENT", "COMMENT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 63, 597, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7,
		62, 2, 63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 67, 7, 67,
		2, 68, 7, 68, 2, 69, 7, 69, 2, 70, 7, 70, 2, 71, 7, 71, 2, 72, 7, 72, 2,
		73, 7, 73, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27,
		1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1,
		29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 31,
		1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1,
		33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34,
		1, 34, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1,
		35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 37,
		1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1,
		39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 40,
		1, 40, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 42, 1, 42, 1, 43, 1, 43, 1,
		44, 1, 44, 1, 45, 1, 45, 1, 46, 1, 46, 1, 47, 1, 47, 1, 48, 1, 48, 1, 49,
		1, 49, 1, 50, 1, 50, 1, 51, 1, 51, 1, 52, 1, 52, 1, 53, 1, 53, 1, 54, 1,
		54, 1, 55, 1, 55, 1, 56, 1, 56, 5, 56, 478, 8, 56, 10, 56, 12, 56, 481,
		9, 56, 1, 56, 1, 56, 1, 56, 5, 56, 486, 8, 56, 10, 56, 12, 56, 489, 9,
		56, 1, 56, 3, 56, 492, 8, 56, 1, 57, 1, 57, 1, 57, 1, 57, 3, 57, 498, 8,
		57, 1, 58, 1, 58, 1, 58, 1, 58, 1, 58, 1, 59, 1, 59, 1, 59, 1, 59, 1, 59,
		1, 60, 1, 60, 1, 60, 1, 61, 1, 61, 1, 61, 1, 61, 1, 61, 1, 61, 1, 61, 1,
		61, 1, 61, 3, 61, 522, 8, 61, 1, 62, 1, 62, 1, 62, 3, 62, 527, 8, 62, 1,
		63, 1, 63, 5, 63, 531, 8, 63, 10, 63, 12, 63, 534, 9, 63, 1, 64, 1, 64,
		5, 64, 538, 8, 64, 10, 64, 12, 64, 541, 9, 64, 1, 65, 1, 65, 1, 65, 4,
		65, 546, 8, 65, 11, 65, 12, 65, 547, 1, 66, 1, 66, 1, 66, 5, 66, 553, 8,
		66, 10, 66, 12, 66, 556, 9, 66, 1, 67, 1, 67, 1, 68, 1, 68, 1, 69, 1, 69,
		1, 70, 1, 70, 1, 71, 4, 71, 567, 8, 71, 11, 71, 12, 71, 568, 1, 71, 1,
		71, 1, 72, 1, 72, 1, 72, 1, 72, 5, 72, 577, 8, 72, 10, 72, 12, 72, 580,
		9, 72, 1, 72, 1, 72, 1, 73, 1, 73, 1, 73, 1, 73, 5, 73, 588, 8, 73, 10,
		73, 12, 73, 591, 9, 73, 1, 73, 1, 73, 1, 73, 1, 73, 1, 73, 3, 479, 487,
		589, 0, 74, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19,
		10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37,
		19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55,
		28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73,
		37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 87, 44, 89, 45, 91,
		46, 93, 47, 95, 48, 97, 49, 99, 50, 101, 51, 103, 52, 105, 53, 107, 54,
		109, 55, 111, 56, 113, 57, 115, 0, 117, 0, 119, 0, 121, 0, 123, 58, 125,
		59, 127, 0, 129, 0, 131, 0, 133, 60, 135, 0, 137, 0, 139, 0, 141, 0, 143,
		61, 145, 62, 147, 63, 1, 0, 10, 3, 0, 0, 0, 10, 10, 92, 92, 2, 0, 88, 88,
		120, 120, 9, 0, 34, 34, 39, 39, 92, 92, 97, 98, 102, 102, 110, 110, 114,
		114, 116, 116, 118, 118, 1, 0, 49, 57, 3, 0, 65, 90, 95, 95, 97, 122, 1,
		0, 48, 57, 1, 0, 48, 55, 3, 0, 48, 57, 65, 70, 97, 102, 3, 0, 9, 10, 12,
		13, 32, 32, 2, 0, 10, 10, 13, 13, 602, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0,
		0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0,
		0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0,
		0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1,
		0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35,
		1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0,
		43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0,
		0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0,
		0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0,
		0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1,
		0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81,
		1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0,
		89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0,
		0, 97, 1, 0, 0, 0, 0, 99, 1, 0, 0, 0, 0, 101, 1, 0, 0, 0, 0, 103, 1, 0,
		0, 0, 0, 105, 1, 0, 0, 0, 0, 107, 1, 0, 0, 0, 0, 109, 1, 0, 0, 0, 0, 111,
		1, 0, 0, 0, 0, 113, 1, 0, 0, 0, 0, 123, 1, 0, 0, 0, 0, 125, 1, 0, 0, 0,
		0, 133, 1, 0, 0, 0, 0, 143, 1, 0, 0, 0, 0, 145, 1, 0, 0, 0, 0, 147, 1,
		0, 0, 0, 1, 149, 1, 0, 0, 0, 3, 156, 1, 0, 0, 0, 5, 163, 1, 0, 0, 0, 7,
		180, 1, 0, 0, 0, 9, 185, 1, 0, 0, 0, 11, 192, 1, 0, 0, 0, 13, 200, 1, 0,
		0, 0, 15, 207, 1, 0, 0, 0, 17, 216, 1, 0, 0, 0, 19, 225, 1, 0, 0, 0, 21,
		231, 1, 0, 0, 0, 23, 235, 1, 0, 0, 0, 25, 241, 1, 0, 0, 0, 27, 247, 1,
		0, 0, 0, 29, 254, 1, 0, 0, 0, 31, 261, 1, 0, 0, 0, 33, 268, 1, 0, 0, 0,
		35, 275, 1, 0, 0, 0, 37, 283, 1, 0, 0, 0, 39, 291, 1, 0, 0, 0, 41, 300,
		1, 0, 0, 0, 43, 309, 1, 0, 0, 0, 45, 314, 1, 0, 0, 0, 47, 321, 1, 0, 0,
		0, 49, 328, 1, 0, 0, 0, 51, 334, 1, 0, 0, 0, 53, 340, 1, 0, 0, 0, 55, 350,
		1, 0, 0, 0, 57, 359, 1, 0, 0, 0, 59, 365, 1, 0, 0, 0, 61, 374, 1, 0, 0,
		0, 63, 377, 1, 0, 0, 0, 65, 381, 1, 0, 0, 0, 67, 386, 1, 0, 0, 0, 69, 393,
		1, 0, 0, 0, 71, 403, 1, 0, 0, 0, 73, 411, 1, 0, 0, 0, 75, 419, 1, 0, 0,
		0, 77, 426, 1, 0, 0, 0, 79, 430, 1, 0, 0, 0, 81, 437, 1, 0, 0, 0, 83, 445,
		1, 0, 0, 0, 85, 447, 1, 0, 0, 0, 87, 449, 1, 0, 0, 0, 89, 451, 1, 0, 0,
		0, 91, 453, 1, 0, 0, 0, 93, 455, 1, 0, 0, 0, 95, 457, 1, 0, 0, 0, 97, 459,
		1, 0, 0, 0, 99, 461, 1, 0, 0, 0, 101, 463, 1, 0, 0, 0, 103, 465, 1, 0,
		0, 0, 105, 467, 1, 0, 0, 0, 107, 469, 1, 0, 0, 0, 109, 471, 1, 0, 0, 0,
		111, 473, 1, 0, 0, 0, 113, 491, 1, 0, 0, 0, 115, 497, 1, 0, 0, 0, 117,
		499, 1, 0, 0, 0, 119, 504, 1, 0, 0, 0, 121, 509, 1, 0, 0, 0, 123, 521,
		1, 0, 0, 0, 125, 526, 1, 0, 0, 0, 127, 528, 1, 0, 0, 0, 129, 535, 1, 0,
		0, 0, 131, 542, 1, 0, 0, 0, 133, 549, 1, 0, 0, 0, 135, 557, 1, 0, 0, 0,
		137, 559, 1, 0, 0, 0, 139, 561, 1, 0, 0, 0, 141, 563, 1, 0, 0, 0, 143,
		566, 1, 0, 0, 0, 145, 572, 1, 0, 0, 0, 147, 583, 1, 0, 0, 0, 149, 150,
		5, 115, 0, 0, 150, 151, 5, 121, 0, 0, 151, 152, 5, 110, 0, 0, 152, 153,
		5, 116, 0, 0, 153, 154, 5, 97, 0, 0, 154, 155, 5, 120, 0, 0, 155, 2, 1,
		0, 0, 0, 156, 157, 5, 105, 0, 0, 157, 158, 5, 109, 0, 0, 158, 159, 5, 112,
		0, 0, 159, 160, 5, 111, 0, 0, 160, 161, 5, 114, 0, 0, 161, 162, 5, 116,
		0, 0, 162, 4, 1, 0, 0, 0, 163, 164, 5, 112, 0, 0, 164, 165, 5, 114, 0,
		0, 165, 166, 5, 111, 0, 0, 166, 167, 5, 116, 0, 0, 167, 168, 5, 111, 0,
		0, 168, 169, 5, 95, 0, 0, 169, 170, 5, 103, 0, 0, 170, 171, 5, 111, 0,
		0, 171, 172, 5, 95, 0, 0, 172, 173, 5, 112, 0, 0, 173, 174, 5, 97, 0, 0,
		174, 175, 5, 99, 0, 0, 175, 176, 5, 107, 0, 0, 176, 177, 5, 97, 0, 0, 177,
		178, 5, 103, 0, 0, 178, 179, 5, 101, 0, 0, 179, 6, 1, 0, 0, 0, 180, 181,
		5, 119, 0, 0, 181, 182, 5, 101, 0, 0, 182, 183, 5, 97, 0, 0, 183, 184,
		5, 107, 0, 0, 184, 8, 1, 0, 0, 0, 185, 186, 5, 112, 0, 0, 186, 187, 5,
		117, 0, 0, 187, 188, 5, 98, 0, 0, 188, 189, 5, 108, 0, 0, 189, 190, 5,
		105, 0, 0, 190, 191, 5, 99, 0, 0, 191, 10, 1, 0, 0, 0, 192, 193, 5, 112,
		0, 0, 193, 194, 5, 97, 0, 0, 194, 195, 5, 99, 0, 0, 195, 196, 5, 107, 0,
		0, 196, 197, 5, 97, 0, 0, 197, 198, 5, 103, 0, 0, 198, 199, 5, 101, 0,
		0, 199, 12, 1, 0, 0, 0, 200, 201, 5, 111, 0, 0, 201, 202, 5, 112, 0, 0,
		202, 203, 5, 116, 0, 0, 203, 204, 5, 105, 0, 0, 204, 205, 5, 111, 0, 0,
		205, 206, 5, 110, 0, 0, 206, 14, 1, 0, 0, 0, 207, 208, 5, 111, 0, 0, 208,
		209, 5, 112, 0, 0, 209, 210, 5, 116, 0, 0, 210, 211, 5, 105, 0, 0, 211,
		212, 5, 111, 0, 0, 212, 213, 5, 110, 0, 0, 213, 214, 5, 97, 0, 0, 214,
		215, 5, 108, 0, 0, 215, 16, 1, 0, 0, 0, 216, 217, 5, 114, 0, 0, 217, 218,
		5, 101, 0, 0, 218, 219, 5, 112, 0, 0, 219, 220, 5, 101, 0, 0, 220, 221,
		5, 97, 0, 0, 221, 222, 5, 116, 0, 0, 222, 223, 5, 101, 0, 0, 223, 224,
		5, 100, 0, 0, 224, 18, 1, 0, 0, 0, 225, 226, 5, 111, 0, 0, 226, 227, 5,
		110, 0, 0, 227, 228, 5, 101, 0, 0, 228, 229, 5, 111, 0, 0, 229, 230, 5,
		102, 0, 0, 230, 20, 1, 0, 0, 0, 231, 232, 5, 109, 0, 0, 232, 233, 5, 97,
		0, 0, 233, 234, 5, 112, 0, 0, 234, 22, 1, 0, 0, 0, 235, 236, 5, 105, 0,
		0, 236, 237, 5, 110, 0, 0, 237, 238, 5, 116, 0, 0, 238, 239, 5, 51, 0,
		0, 239, 240, 5, 50, 0, 0, 240, 24, 1, 0, 0, 0, 241, 242, 5, 105, 0, 0,
		242, 243, 5, 110, 0, 0, 243, 244, 5, 116, 0, 0, 244, 245, 5, 54, 0, 0,
		245, 246, 5, 52, 0, 0, 246, 26, 1, 0, 0, 0, 247, 248, 5, 117, 0, 0, 248,
		249, 5, 105, 0, 0, 249, 250, 5, 110, 0, 0, 250, 251, 5, 116, 0, 0, 251,
		252, 5, 51, 0, 0, 252, 253, 5, 50, 0, 0, 253, 28, 1, 0, 0, 0, 254, 255,
		5, 117, 0, 0, 255, 256, 5, 105, 0, 0, 256, 257, 5, 110, 0, 0, 257, 258,
		5, 116, 0, 0, 258, 259, 5, 54, 0, 0, 259, 260, 5, 52, 0, 0, 260, 30, 1,
		0, 0, 0, 261, 262, 5, 115, 0, 0, 262, 263, 5, 105, 0, 0, 263, 264, 5, 110,
		0, 0, 264, 265, 5, 116, 0, 0, 265, 266, 5, 51, 0, 0, 266, 267, 5, 50, 0,
		0, 267, 32, 1, 0, 0, 0, 268, 269, 5, 115, 0, 0, 269, 270, 5, 105, 0, 0,
		270, 271, 5, 110, 0, 0, 271, 272, 5, 116, 0, 0, 272, 273, 5, 54, 0, 0,
		273, 274, 5, 52, 0, 0, 274, 34, 1, 0, 0, 0, 275, 276, 5, 102, 0, 0, 276,
		277, 5, 105, 0, 0, 277, 278, 5, 120, 0, 0, 278, 279, 5, 101, 0, 0, 279,
		280, 5, 100, 0, 0, 280, 281, 5, 51, 0, 0, 281, 282, 5, 50, 0, 0, 282, 36,
		1, 0, 0, 0, 283, 284, 5, 102, 0, 0, 284, 285, 5, 105, 0, 0, 285, 286, 5,
		120, 0, 0, 286, 287, 5, 101, 0, 0, 287, 288, 5, 100, 0, 0, 288, 289, 5,
		54, 0, 0, 289, 290, 5, 52, 0, 0, 290, 38, 1, 0, 0, 0, 291, 292, 5, 115,
		0, 0, 292, 293, 5, 102, 0, 0, 293, 294, 5, 105, 0, 0, 294, 295, 5, 120,
		0, 0, 295, 296, 5, 101, 0, 0, 296, 297, 5, 100, 0, 0, 297, 298, 5, 51,
		0, 0, 298, 299, 5, 50, 0, 0, 299, 40, 1, 0, 0, 0, 300, 301, 5, 115, 0,
		0, 301, 302, 5, 102, 0, 0, 302, 303, 5, 105, 0, 0, 303, 304, 5, 120, 0,
		0, 304, 305, 5, 101, 0, 0, 305, 306, 5, 100, 0, 0, 306, 307, 5, 54, 0,
		0, 307, 308, 5, 52, 0, 0, 308, 42, 1, 0, 0, 0, 309, 310, 5, 98, 0, 0, 310,
		311, 5, 111, 0, 0, 311, 312, 5, 111, 0, 0, 312, 313, 5, 108, 0, 0, 313,
		44, 1, 0, 0, 0, 314, 315, 5, 115, 0, 0, 315, 316, 5, 116, 0, 0, 316, 317,
		5, 114, 0, 0, 317, 318, 5, 105, 0, 0, 318, 319, 5, 110, 0, 0, 319, 320,
		5, 103, 0, 0, 320, 46, 1, 0, 0, 0, 321, 322, 5, 100, 0, 0, 322, 323, 5,
		111, 0, 0, 323, 324, 5, 117, 0, 0, 324, 325, 5, 98, 0, 0, 325, 326, 5,
		108, 0, 0, 326, 327, 5, 101, 0, 0, 327, 48, 1, 0, 0, 0, 328, 329, 5, 102,
		0, 0, 329, 330, 5, 108, 0, 0, 330, 331, 5, 111, 0, 0, 331, 332, 5, 97,
		0, 0, 332, 333, 5, 116, 0, 0, 333, 50, 1, 0, 0, 0, 334, 335, 5, 98, 0,
		0, 335, 336, 5, 121, 0, 0, 336, 337, 5, 116, 0, 0, 337, 338, 5, 101, 0,
		0, 338, 339, 5, 115, 0, 0, 339, 52, 1, 0, 0, 0, 340, 341, 5, 116, 0, 0,
		341, 342, 5, 105, 0, 0, 342, 343, 5, 109, 0, 0, 343, 344, 5, 101, 0, 0,
		344, 345, 5, 115, 0, 0, 345, 346, 5, 116, 0, 0, 346, 347, 5, 97, 0, 0,
		347, 348, 5, 109, 0, 0, 348, 349, 5, 112, 0, 0, 349, 54, 1, 0, 0, 0, 350,
		351, 5, 100, 0, 0, 351, 352, 5, 117, 0, 0, 352, 353, 5, 114, 0, 0, 353,
		354, 5, 97, 0, 0, 354, 355, 5, 116, 0, 0, 355, 356, 5, 105, 0, 0, 356,
		357, 5, 111, 0, 0, 357, 358, 5, 110, 0, 0, 358, 56, 1, 0, 0, 0, 359, 360,
		5, 101, 0, 0, 360, 361, 5, 109, 0, 0, 361, 362, 5, 112, 0, 0, 362, 363,
		5, 116, 0, 0, 363, 364, 5, 121, 0, 0, 364, 58, 1, 0, 0, 0, 365, 366, 5,
		114, 0, 0, 366, 367, 5, 101, 0, 0, 367, 368, 5, 115, 0, 0, 368, 369, 5,
		101, 0, 0, 369, 370, 5, 114, 0, 0, 370, 371, 5, 118, 0, 0, 371, 372, 5,
		101, 0, 0, 372, 373, 5, 100, 0, 0, 373, 60, 1, 0, 0, 0, 374, 375, 5, 116,
		0, 0, 375, 376, 5, 111, 0, 0, 376, 62, 1, 0, 0, 0, 377, 378, 5, 109, 0,
		0, 378, 379, 5, 97, 0, 0, 379, 380, 5, 120, 0, 0, 380, 64, 1, 0, 0, 0,
		381, 382, 5, 101, 0, 0, 382, 383, 5, 110, 0, 0, 383, 384, 5, 117, 0, 0,
		384, 385, 5, 109, 0, 0, 385, 66, 1, 0, 0, 0, 386, 387, 5, 101, 0, 0, 387,
		388, 5, 110, 0, 0, 388, 389, 5, 116, 0, 0, 389, 390, 5, 105, 0, 0, 390,
		391, 5, 116, 0, 0, 391, 392, 5, 121, 0, 0, 392, 68, 1, 0, 0, 0, 393, 394,
		5, 99, 0, 0, 394, 395, 5, 111, 0, 0, 395, 396, 5, 109, 0, 0, 396, 397,
		5, 112, 0, 0, 397, 398, 5, 111, 0, 0, 398, 399, 5, 110, 0, 0, 399, 400,
		5, 101, 0, 0, 400, 401, 5, 110, 0, 0, 401, 402, 5, 116, 0, 0, 402, 70,
		1, 0, 0, 0, 403, 404, 5, 109, 0, 0, 404, 405, 5, 101, 0, 0, 405, 406, 5,
		115, 0, 0, 406, 407, 5, 115, 0, 0, 407, 408, 5, 97, 0, 0, 408, 409, 5,
		103, 0, 0, 409, 410, 5, 101, 0, 0, 410, 72, 1, 0, 0, 0, 411, 412, 5, 115,
		0, 0, 412, 413, 5, 101, 0, 0, 413, 414, 5, 114, 0, 0, 414, 415, 5, 118,
		0, 0, 415, 416, 5, 105, 0, 0, 416, 417, 5, 99, 0, 0, 417, 418, 5, 101,
		0, 0, 418, 74, 1, 0, 0, 0, 419, 420, 5, 101, 0, 0, 420, 421, 5, 120, 0,
		0, 421, 422, 5, 116, 0, 0, 422, 423, 5, 101, 0, 0, 423, 424, 5, 110, 0,
		0, 424, 425, 5, 100, 0, 0, 425, 76, 1, 0, 0, 0, 426, 427, 5, 114, 0, 0,
		427, 428, 5, 112, 0, 0, 428, 429, 5, 99, 0, 0, 429, 78, 1, 0, 0, 0, 430,
		431, 5, 115, 0, 0, 431, 432, 5, 116, 0, 0, 432, 433, 5, 114, 0, 0, 433,
		434, 5, 101, 0, 0, 434, 435, 5, 97, 0, 0, 435, 436, 5, 109, 0, 0, 436,
		80, 1, 0, 0, 0, 437, 438, 5, 114, 0, 0, 438, 439, 5, 101, 0, 0, 439, 440,
		5, 116, 0, 0, 440, 441, 5, 117, 0, 0, 441, 442, 5, 114, 0, 0, 442, 443,
		5, 110, 0, 0, 443, 444, 5, 115, 0, 0, 444, 82, 1, 0, 0, 0, 445, 446, 5,
		59, 0, 0, 446, 84, 1, 0, 0, 0, 447, 448, 5, 61, 0, 0, 448, 86, 1, 0, 0,
		0, 449, 450, 5, 40, 0, 0, 450, 88, 1, 0, 0, 0, 451, 452, 5, 41, 0, 0, 452,
		90, 1, 0, 0, 0, 453, 454, 5, 91, 0, 0, 454, 92, 1, 0, 0, 0, 455, 456, 5,
		93, 0, 0, 456, 94, 1, 0, 0, 0, 457, 458, 5, 123, 0, 0, 458, 96, 1, 0, 0,
		0, 459, 460, 5, 125, 0, 0, 460, 98, 1, 0, 0, 0, 461, 462, 5, 60, 0, 0,
		462, 100, 1, 0, 0, 0, 463, 464, 5, 62, 0, 0, 464, 102, 1, 0, 0, 0, 465,
		466, 5, 46, 0, 0, 466, 104, 1, 0, 0, 0, 467, 468, 5, 44, 0, 0, 468, 106,
		1, 0, 0, 0, 469, 470, 5, 58, 0, 0, 470, 108, 1, 0, 0, 0, 471, 472, 5, 43,
		0, 0, 472, 110, 1, 0, 0, 0, 473, 474, 5, 45, 0, 0, 474, 112, 1, 0, 0, 0,
		475, 479, 5, 39, 0, 0, 476, 478, 3, 115, 57, 0, 477, 476, 1, 0, 0, 0, 478,
		481, 1, 0, 0, 0, 479, 480, 1, 0, 0, 0, 479, 477, 1, 0, 0, 0, 480, 482,
		1, 0, 0, 0, 481, 479, 1, 0, 0, 0, 482, 492, 5, 39, 0, 0, 483, 487, 5, 34,
		0, 0, 484, 486, 3, 115, 57, 0, 485, 484, 1, 0, 0, 0, 486, 489, 1, 0, 0,
		0, 487, 488, 1, 0, 0, 0, 487, 485, 1, 0, 0, 0, 488, 490, 1, 0, 0, 0, 489,
		487, 1, 0, 0, 0, 490, 492, 5, 34, 0, 0, 491, 475, 1, 0, 0, 0, 491, 483,
		1, 0, 0, 0, 492, 114, 1, 0, 0, 0, 493, 498, 3, 117, 58, 0, 494, 498, 3,
		119, 59, 0, 495, 498, 3, 121, 60, 0, 496, 498, 8, 0, 0, 0, 497, 493, 1,
		0, 0, 0, 497, 494, 1, 0, 0, 0, 497, 495, 1, 0, 0, 0, 497, 496, 1, 0, 0,
		0, 498, 116, 1, 0, 0, 0, 499, 500, 5, 92, 0, 0, 500, 501, 7, 1, 0, 0, 501,
		502, 3, 141, 70, 0, 502, 503, 3, 141, 70, 0, 503, 118, 1, 0, 0, 0, 504,
		505, 5, 92, 0, 0, 505, 506, 3, 139, 69, 0, 506, 507, 3, 139, 69, 0, 507,
		508, 3, 139, 69, 0, 508, 120, 1, 0, 0, 0, 509, 510, 5, 92, 0, 0, 510, 511,
		7, 2, 0, 0, 511, 122, 1, 0, 0, 0, 512, 513, 5, 116, 0, 0, 513, 514, 5,
		114, 0, 0, 514, 515, 5, 117, 0, 0, 515, 522, 5, 101, 0, 0, 516, 517, 5,
		102, 0, 0, 517, 518, 5, 97, 0, 0, 518, 519, 5, 108, 0, 0, 519, 520, 5,
		115, 0, 0, 520, 522, 5, 101, 0, 0, 521, 512, 1, 0, 0, 0, 521, 516, 1, 0,
		0, 0, 522, 124, 1, 0, 0, 0, 523, 527, 3, 127, 63, 0, 524, 527, 3, 129,
		64, 0, 525, 527, 3, 131, 65, 0, 526, 523, 1, 0, 0, 0, 526, 524, 1, 0, 0,
		0, 526, 525, 1, 0, 0, 0, 527, 126, 1, 0, 0, 0, 528, 532, 7, 3, 0, 0, 529,
		531, 3, 137, 68, 0, 530, 529, 1, 0, 0, 0, 531, 534, 1, 0, 0, 0, 532, 530,
		1, 0, 0, 0, 532, 533, 1, 0, 0, 0, 533, 128, 1, 0, 0, 0, 534, 532, 1, 0,
		0, 0, 535, 539, 5, 48, 0, 0, 536, 538, 3, 139, 69, 0, 537, 536, 1, 0, 0,
		0, 538, 541, 1, 0, 0, 0, 539, 537, 1, 0, 0, 0, 539, 540, 1, 0, 0, 0, 540,
		130, 1, 0, 0, 0, 541, 539, 1, 0, 0, 0, 542, 543, 5, 48, 0, 0, 543, 545,
		7, 1, 0, 0, 544, 546, 3, 141, 70, 0, 545, 544, 1, 0, 0, 0, 546, 547, 1,
		0, 0, 0, 547, 545, 1, 0, 0, 0, 547, 548, 1, 0, 0, 0, 548, 132, 1, 0, 0,
		0, 549, 554, 3, 135, 67, 0, 550, 553, 3, 135, 67, 0, 551, 553, 3, 137,
		68, 0, 552, 550, 1, 0, 0, 0, 552, 551, 1, 0, 0, 0, 553, 556, 1, 0, 0, 0,
		554, 552, 1, 0, 0, 0, 554, 555, 1, 0, 0, 0, 555, 134, 1, 0, 0, 0, 556,
		554, 1, 0, 0, 0, 557, 558, 7, 4, 0, 0, 558, 136, 1, 0, 0, 0, 559, 560,
		7, 5, 0, 0, 560, 138, 1, 0, 0, 0, 561, 562, 7, 6, 0, 0, 562, 140, 1, 0,
		0, 0, 563, 564, 7, 7, 0, 0, 564, 142, 1, 0, 0, 0, 565, 567, 7, 8, 0, 0,
		566, 565, 1, 0, 0, 0, 567, 568, 1, 0, 0, 0, 568, 566, 1, 0, 0, 0, 568,
		569, 1, 0, 0, 0, 569, 570, 1, 0, 0, 0, 570, 571, 6, 71, 0, 0, 571, 144,
		1, 0, 0, 0, 572, 573, 5, 47, 0, 0, 573, 574, 5, 47, 0, 0, 574, 578, 1,
		0, 0, 0, 575, 577, 8, 9, 0, 0, 576, 575, 1, 0, 0, 0, 577, 580, 1, 0, 0,
		0, 578, 576, 1, 0, 0, 0, 578, 579, 1, 0, 0, 0, 579, 581, 1, 0, 0, 0, 580,
		578, 1, 0, 0, 0, 581, 582, 6, 72, 1, 0, 582, 146, 1, 0, 0, 0, 583, 584,
		5, 47, 0, 0, 584, 585, 5, 42, 0, 0, 585, 589, 1, 0, 0, 0, 586, 588, 9,
		0, 0, 0, 587, 586, 1, 0, 0, 0, 588, 591, 1, 0, 0, 0, 589, 590, 1, 0, 0,
		0, 589, 587, 1, 0, 0, 0, 590, 592, 1, 0, 0, 0, 591, 589, 1, 0, 0, 0, 592,
		593, 5, 42, 0, 0, 593, 594, 5, 47, 0, 0, 594, 595, 1, 0, 0, 0, 595, 596,
		6, 73, 1, 0, 596, 148, 1, 0, 0, 0, 15, 0, 479, 487, 491, 497, 521, 526,
		532, 539, 547, 552, 554, 568, 578, 589, 2, 6, 0, 0, 0, 1, 0,
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

// kdsLexerInit initializes any static state used to implement kdsLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewkdsLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func KdsLexerInit() {
	staticData := &KdsLexerLexerStaticData
	staticData.once.Do(kdslexerLexerInit)
}

// NewkdsLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewkdsLexer(input antlr.CharStream) *kdsLexer {
	KdsLexerInit()
	l := new(kdsLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &KdsLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "kds.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// kdsLexer tokens.
const (
	kdsLexerSYNTAX           = 1
	kdsLexerIMPORT           = 2
	kdsLexerPROTO_GO_PACKAGE = 3
	kdsLexerWEAK             = 4
	kdsLexerPUBLIC           = 5
	kdsLexerPACKAGE          = 6
	kdsLexerOPTION           = 7
	kdsLexerOPTIONAL         = 8
	kdsLexerREPEATED         = 9
	kdsLexerONEOF            = 10
	kdsLexerMAP              = 11
	kdsLexerINT32            = 12
	kdsLexerINT64            = 13
	kdsLexerUINT32           = 14
	kdsLexerUINT64           = 15
	kdsLexerSINT32           = 16
	kdsLexerSINT64           = 17
	kdsLexerFIXED32          = 18
	kdsLexerFIXED64          = 19
	kdsLexerSFIXED32         = 20
	kdsLexerSFIXED64         = 21
	kdsLexerBOOL             = 22
	kdsLexerSTRING           = 23
	kdsLexerDOUBLE           = 24
	kdsLexerFLOAT            = 25
	kdsLexerBYTES            = 26
	kdsLexerTIMESTAMP        = 27
	kdsLexerDURATION         = 28
	kdsLexerEMPTY            = 29
	kdsLexerRESERVED         = 30
	kdsLexerTO               = 31
	kdsLexerMAX              = 32
	kdsLexerENUM             = 33
	kdsLexerENTITY           = 34
	kdsLexerCOMPONENT        = 35
	kdsLexerMESSAGE          = 36
	kdsLexerSERVICE          = 37
	kdsLexerEXTEND           = 38
	kdsLexerRPC              = 39
	kdsLexerSTREAM           = 40
	kdsLexerRETURNS          = 41
	kdsLexerSEMI             = 42
	kdsLexerEQ               = 43
	kdsLexerLP               = 44
	kdsLexerRP               = 45
	kdsLexerLB               = 46
	kdsLexerRB               = 47
	kdsLexerLC               = 48
	kdsLexerRC               = 49
	kdsLexerLT               = 50
	kdsLexerGT               = 51
	kdsLexerDOT              = 52
	kdsLexerCOMMA            = 53
	kdsLexerCOLON            = 54
	kdsLexerPLUS             = 55
	kdsLexerMINUS            = 56
	kdsLexerSTR_LIT          = 57
	kdsLexerBOOL_LIT         = 58
	kdsLexerINT_LIT          = 59
	kdsLexerIDENTIFIER       = 60
	kdsLexerWS               = 61
	kdsLexerLINE_COMMENT     = 62
	kdsLexerCOMMENT          = 63
)
