// Code generated from internal/parser/query/Query.g4 by ANTLR 4.7.1. DO NOT EDIT.

package query

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 21, 176,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3,
	4, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3,
	12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 17, 3, 17, 3, 17, 7, 17, 117, 10, 17, 12, 17, 14, 17, 120, 11, 17, 3,
	18, 3, 18, 3, 18, 7, 18, 125, 10, 18, 12, 18, 14, 18, 128, 11, 18, 3, 19,
	3, 19, 3, 20, 3, 20, 3, 21, 6, 21, 135, 10, 21, 13, 21, 14, 21, 136, 3,
	21, 3, 21, 3, 22, 5, 22, 142, 10, 22, 3, 22, 3, 22, 6, 22, 146, 10, 22,
	13, 22, 14, 22, 147, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 7, 23, 156,
	10, 23, 12, 23, 14, 23, 159, 11, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23,
	3, 24, 3, 24, 3, 24, 3, 24, 7, 24, 170, 10, 24, 12, 24, 14, 24, 173, 11,
	24, 3, 24, 3, 24, 3, 157, 2, 25, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8,
	15, 9, 17, 10, 19, 11, 21, 12, 23, 2, 25, 2, 27, 2, 29, 2, 31, 13, 33,
	14, 35, 15, 37, 16, 39, 17, 41, 18, 43, 19, 45, 20, 47, 21, 3, 2, 8, 3,
	2, 67, 92, 3, 2, 99, 124, 5, 2, 67, 92, 97, 97, 99, 124, 3, 2, 50, 59,
	5, 2, 11, 11, 14, 14, 34, 34, 4, 2, 12, 12, 15, 15, 2, 181, 2, 3, 3, 2,
	2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2,
	2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3,
	2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35,
	3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2,
	43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 3, 49, 3, 2, 2, 2,
	5, 54, 3, 2, 2, 2, 7, 56, 3, 2, 2, 2, 9, 58, 3, 2, 2, 2, 11, 61, 3, 2,
	2, 2, 13, 65, 3, 2, 2, 2, 15, 71, 3, 2, 2, 2, 17, 78, 3, 2, 2, 2, 19, 83,
	3, 2, 2, 2, 21, 92, 3, 2, 2, 2, 23, 101, 3, 2, 2, 2, 25, 103, 3, 2, 2,
	2, 27, 105, 3, 2, 2, 2, 29, 107, 3, 2, 2, 2, 31, 109, 3, 2, 2, 2, 33, 113,
	3, 2, 2, 2, 35, 121, 3, 2, 2, 2, 37, 129, 3, 2, 2, 2, 39, 131, 3, 2, 2,
	2, 41, 134, 3, 2, 2, 2, 43, 145, 3, 2, 2, 2, 45, 151, 3, 2, 2, 2, 47, 165,
	3, 2, 2, 2, 49, 50, 7, 118, 2, 2, 50, 51, 7, 123, 2, 2, 51, 52, 7, 114,
	2, 2, 52, 53, 7, 103, 2, 2, 53, 4, 3, 2, 2, 2, 54, 55, 7, 125, 2, 2, 55,
	6, 3, 2, 2, 2, 56, 57, 7, 127, 2, 2, 57, 8, 3, 2, 2, 2, 58, 59, 7, 75,
	2, 2, 59, 60, 7, 70, 2, 2, 60, 10, 3, 2, 2, 2, 61, 62, 7, 107, 2, 2, 62,
	63, 7, 112, 2, 2, 63, 64, 7, 118, 2, 2, 64, 12, 3, 2, 2, 2, 65, 66, 7,
	104, 2, 2, 66, 67, 7, 110, 2, 2, 67, 68, 7, 113, 2, 2, 68, 69, 7, 99, 2,
	2, 69, 70, 7, 118, 2, 2, 70, 14, 3, 2, 2, 2, 71, 72, 7, 117, 2, 2, 72,
	73, 7, 118, 2, 2, 73, 74, 7, 116, 2, 2, 74, 75, 7, 107, 2, 2, 75, 76, 7,
	112, 2, 2, 76, 77, 7, 105, 2, 2, 77, 16, 3, 2, 2, 2, 78, 79, 7, 100, 2,
	2, 79, 80, 7, 113, 2, 2, 80, 81, 7, 113, 2, 2, 81, 82, 7, 110, 2, 2, 82,
	18, 3, 2, 2, 2, 83, 84, 7, 102, 2, 2, 84, 85, 7, 99, 2, 2, 85, 86, 7, 118,
	2, 2, 86, 87, 7, 103, 2, 2, 87, 88, 7, 118, 2, 2, 88, 89, 7, 107, 2, 2,
	89, 90, 7, 111, 2, 2, 90, 91, 7, 103, 2, 2, 91, 20, 3, 2, 2, 2, 92, 93,
	7, 105, 2, 2, 93, 94, 7, 103, 2, 2, 94, 95, 7, 113, 2, 2, 95, 96, 7, 114,
	2, 2, 96, 97, 7, 113, 2, 2, 97, 98, 7, 107, 2, 2, 98, 99, 7, 112, 2, 2,
	99, 100, 7, 118, 2, 2, 100, 22, 3, 2, 2, 2, 101, 102, 9, 2, 2, 2, 102,
	24, 3, 2, 2, 2, 103, 104, 9, 3, 2, 2, 104, 26, 3, 2, 2, 2, 105, 106, 9,
	4, 2, 2, 106, 28, 3, 2, 2, 2, 107, 108, 9, 5, 2, 2, 108, 30, 3, 2, 2, 2,
	109, 110, 5, 37, 19, 2, 110, 111, 5, 33, 17, 2, 111, 112, 5, 39, 20, 2,
	112, 32, 3, 2, 2, 2, 113, 118, 5, 23, 12, 2, 114, 117, 5, 27, 14, 2, 115,
	117, 5, 29, 15, 2, 116, 114, 3, 2, 2, 2, 116, 115, 3, 2, 2, 2, 117, 120,
	3, 2, 2, 2, 118, 116, 3, 2, 2, 2, 118, 119, 3, 2, 2, 2, 119, 34, 3, 2,
	2, 2, 120, 118, 3, 2, 2, 2, 121, 126, 5, 25, 13, 2, 122, 125, 5, 27, 14,
	2, 123, 125, 5, 29, 15, 2, 124, 122, 3, 2, 2, 2, 124, 123, 3, 2, 2, 2,
	125, 128, 3, 2, 2, 2, 126, 124, 3, 2, 2, 2, 126, 127, 3, 2, 2, 2, 127,
	36, 3, 2, 2, 2, 128, 126, 3, 2, 2, 2, 129, 130, 7, 93, 2, 2, 130, 38, 3,
	2, 2, 2, 131, 132, 7, 95, 2, 2, 132, 40, 3, 2, 2, 2, 133, 135, 9, 6, 2,
	2, 134, 133, 3, 2, 2, 2, 135, 136, 3, 2, 2, 2, 136, 134, 3, 2, 2, 2, 136,
	137, 3, 2, 2, 2, 137, 138, 3, 2, 2, 2, 138, 139, 8, 21, 2, 2, 139, 42,
	3, 2, 2, 2, 140, 142, 7, 15, 2, 2, 141, 140, 3, 2, 2, 2, 141, 142, 3, 2,
	2, 2, 142, 143, 3, 2, 2, 2, 143, 146, 7, 12, 2, 2, 144, 146, 7, 15, 2,
	2, 145, 141, 3, 2, 2, 2, 145, 144, 3, 2, 2, 2, 146, 147, 3, 2, 2, 2, 147,
	145, 3, 2, 2, 2, 147, 148, 3, 2, 2, 2, 148, 149, 3, 2, 2, 2, 149, 150,
	8, 22, 2, 2, 150, 44, 3, 2, 2, 2, 151, 152, 7, 49, 2, 2, 152, 153, 7, 44,
	2, 2, 153, 157, 3, 2, 2, 2, 154, 156, 11, 2, 2, 2, 155, 154, 3, 2, 2, 2,
	156, 159, 3, 2, 2, 2, 157, 158, 3, 2, 2, 2, 157, 155, 3, 2, 2, 2, 158,
	160, 3, 2, 2, 2, 159, 157, 3, 2, 2, 2, 160, 161, 7, 44, 2, 2, 161, 162,
	7, 49, 2, 2, 162, 163, 3, 2, 2, 2, 163, 164, 8, 23, 2, 2, 164, 46, 3, 2,
	2, 2, 165, 166, 7, 49, 2, 2, 166, 167, 7, 49, 2, 2, 167, 171, 3, 2, 2,
	2, 168, 170, 10, 7, 2, 2, 169, 168, 3, 2, 2, 2, 170, 173, 3, 2, 2, 2, 171,
	169, 3, 2, 2, 2, 171, 172, 3, 2, 2, 2, 172, 174, 3, 2, 2, 2, 173, 171,
	3, 2, 2, 2, 174, 175, 8, 24, 2, 2, 175, 48, 3, 2, 2, 2, 13, 2, 116, 118,
	124, 126, 136, 141, 145, 147, 157, 171, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'type'", "'{'", "'}'", "'ID'", "'int'", "'float'", "'string'", "'bool'",
	"'datetime'", "'geopoint'", "", "", "", "'['", "']'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "TypeIdentArray", "TypeIdent",
	"FieldIdent", "LBRACK", "RBRACK", "WS", "NEWLINE", "COMMENT", "LINE_COMMENT",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "ULetter", "LLetter", "Letter", "Digit", "TypeIdentArray", "TypeIdent",
	"FieldIdent", "LBRACK", "RBRACK", "WS", "NEWLINE", "COMMENT", "LINE_COMMENT",
}

type QueryLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewQueryLexer(input antlr.CharStream) *QueryLexer {

	l := new(QueryLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Query.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// QueryLexer tokens.
const (
	QueryLexerT__0           = 1
	QueryLexerT__1           = 2
	QueryLexerT__2           = 3
	QueryLexerT__3           = 4
	QueryLexerT__4           = 5
	QueryLexerT__5           = 6
	QueryLexerT__6           = 7
	QueryLexerT__7           = 8
	QueryLexerT__8           = 9
	QueryLexerT__9           = 10
	QueryLexerTypeIdentArray = 11
	QueryLexerTypeIdent      = 12
	QueryLexerFieldIdent     = 13
	QueryLexerLBRACK         = 14
	QueryLexerRBRACK         = 15
	QueryLexerWS             = 16
	QueryLexerNEWLINE        = 17
	QueryLexerCOMMENT        = 18
	QueryLexerLINE_COMMENT   = 19
)