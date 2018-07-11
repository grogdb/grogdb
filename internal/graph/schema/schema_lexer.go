// Code generated from internal/graph/schema/Schema.g4 by ANTLR 4.7.1. DO NOT EDIT.

package schema

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 22, 185,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6,
	3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14,
	3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3,
	18, 3, 18, 7, 18, 126, 10, 18, 12, 18, 14, 18, 129, 11, 18, 3, 19, 3, 19,
	3, 19, 7, 19, 134, 10, 19, 12, 19, 14, 19, 137, 11, 19, 3, 20, 3, 20, 3,
	21, 3, 21, 3, 22, 6, 22, 144, 10, 22, 13, 22, 14, 22, 145, 3, 22, 3, 22,
	3, 23, 5, 23, 151, 10, 23, 3, 23, 3, 23, 6, 23, 155, 10, 23, 13, 23, 14,
	23, 156, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 7, 24, 165, 10, 24,
	12, 24, 14, 24, 168, 11, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 25,
	3, 25, 3, 25, 3, 25, 7, 25, 179, 10, 25, 12, 25, 14, 25, 182, 11, 25, 3,
	25, 3, 25, 3, 166, 2, 26, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9,
	17, 10, 19, 11, 21, 12, 23, 13, 25, 2, 27, 2, 29, 2, 31, 2, 33, 14, 35,
	15, 37, 16, 39, 17, 41, 18, 43, 19, 45, 20, 47, 21, 49, 22, 3, 2, 8, 3,
	2, 67, 92, 4, 2, 97, 97, 99, 124, 5, 2, 67, 92, 97, 97, 99, 124, 3, 2,
	50, 59, 5, 2, 11, 11, 14, 14, 34, 34, 4, 2, 12, 12, 15, 15, 2, 190, 2,
	3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2,
	11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2,
	2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 33, 3, 2, 2,
	2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2,
	2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3,
	2, 2, 2, 3, 51, 3, 2, 2, 2, 5, 58, 3, 2, 2, 2, 7, 60, 3, 2, 2, 2, 9, 62,
	3, 2, 2, 2, 11, 67, 3, 2, 2, 2, 13, 70, 3, 2, 2, 2, 15, 74, 3, 2, 2, 2,
	17, 80, 3, 2, 2, 2, 19, 87, 3, 2, 2, 2, 21, 92, 3, 2, 2, 2, 23, 101, 3,
	2, 2, 2, 25, 110, 3, 2, 2, 2, 27, 112, 3, 2, 2, 2, 29, 114, 3, 2, 2, 2,
	31, 116, 3, 2, 2, 2, 33, 118, 3, 2, 2, 2, 35, 122, 3, 2, 2, 2, 37, 130,
	3, 2, 2, 2, 39, 138, 3, 2, 2, 2, 41, 140, 3, 2, 2, 2, 43, 143, 3, 2, 2,
	2, 45, 154, 3, 2, 2, 2, 47, 160, 3, 2, 2, 2, 49, 174, 3, 2, 2, 2, 51, 52,
	7, 117, 2, 2, 52, 53, 7, 101, 2, 2, 53, 54, 7, 106, 2, 2, 54, 55, 7, 103,
	2, 2, 55, 56, 7, 111, 2, 2, 56, 57, 7, 99, 2, 2, 57, 4, 3, 2, 2, 2, 58,
	59, 7, 125, 2, 2, 59, 6, 3, 2, 2, 2, 60, 61, 7, 127, 2, 2, 61, 8, 3, 2,
	2, 2, 62, 63, 7, 118, 2, 2, 63, 64, 7, 123, 2, 2, 64, 65, 7, 114, 2, 2,
	65, 66, 7, 103, 2, 2, 66, 10, 3, 2, 2, 2, 67, 68, 7, 75, 2, 2, 68, 69,
	7, 70, 2, 2, 69, 12, 3, 2, 2, 2, 70, 71, 7, 107, 2, 2, 71, 72, 7, 112,
	2, 2, 72, 73, 7, 118, 2, 2, 73, 14, 3, 2, 2, 2, 74, 75, 7, 104, 2, 2, 75,
	76, 7, 110, 2, 2, 76, 77, 7, 113, 2, 2, 77, 78, 7, 99, 2, 2, 78, 79, 7,
	118, 2, 2, 79, 16, 3, 2, 2, 2, 80, 81, 7, 117, 2, 2, 81, 82, 7, 118, 2,
	2, 82, 83, 7, 116, 2, 2, 83, 84, 7, 107, 2, 2, 84, 85, 7, 112, 2, 2, 85,
	86, 7, 105, 2, 2, 86, 18, 3, 2, 2, 2, 87, 88, 7, 100, 2, 2, 88, 89, 7,
	113, 2, 2, 89, 90, 7, 113, 2, 2, 90, 91, 7, 110, 2, 2, 91, 20, 3, 2, 2,
	2, 92, 93, 7, 102, 2, 2, 93, 94, 7, 99, 2, 2, 94, 95, 7, 118, 2, 2, 95,
	96, 7, 103, 2, 2, 96, 97, 7, 118, 2, 2, 97, 98, 7, 107, 2, 2, 98, 99, 7,
	111, 2, 2, 99, 100, 7, 103, 2, 2, 100, 22, 3, 2, 2, 2, 101, 102, 7, 105,
	2, 2, 102, 103, 7, 103, 2, 2, 103, 104, 7, 113, 2, 2, 104, 105, 7, 114,
	2, 2, 105, 106, 7, 113, 2, 2, 106, 107, 7, 107, 2, 2, 107, 108, 7, 112,
	2, 2, 108, 109, 7, 118, 2, 2, 109, 24, 3, 2, 2, 2, 110, 111, 9, 2, 2, 2,
	111, 26, 3, 2, 2, 2, 112, 113, 9, 3, 2, 2, 113, 28, 3, 2, 2, 2, 114, 115,
	9, 4, 2, 2, 115, 30, 3, 2, 2, 2, 116, 117, 9, 5, 2, 2, 117, 32, 3, 2, 2,
	2, 118, 119, 5, 39, 20, 2, 119, 120, 5, 35, 18, 2, 120, 121, 5, 41, 21,
	2, 121, 34, 3, 2, 2, 2, 122, 127, 5, 25, 13, 2, 123, 126, 5, 29, 15, 2,
	124, 126, 5, 31, 16, 2, 125, 123, 3, 2, 2, 2, 125, 124, 3, 2, 2, 2, 126,
	129, 3, 2, 2, 2, 127, 125, 3, 2, 2, 2, 127, 128, 3, 2, 2, 2, 128, 36, 3,
	2, 2, 2, 129, 127, 3, 2, 2, 2, 130, 135, 5, 27, 14, 2, 131, 134, 5, 29,
	15, 2, 132, 134, 5, 31, 16, 2, 133, 131, 3, 2, 2, 2, 133, 132, 3, 2, 2,
	2, 134, 137, 3, 2, 2, 2, 135, 133, 3, 2, 2, 2, 135, 136, 3, 2, 2, 2, 136,
	38, 3, 2, 2, 2, 137, 135, 3, 2, 2, 2, 138, 139, 7, 93, 2, 2, 139, 40, 3,
	2, 2, 2, 140, 141, 7, 95, 2, 2, 141, 42, 3, 2, 2, 2, 142, 144, 9, 6, 2,
	2, 143, 142, 3, 2, 2, 2, 144, 145, 3, 2, 2, 2, 145, 143, 3, 2, 2, 2, 145,
	146, 3, 2, 2, 2, 146, 147, 3, 2, 2, 2, 147, 148, 8, 22, 2, 2, 148, 44,
	3, 2, 2, 2, 149, 151, 7, 15, 2, 2, 150, 149, 3, 2, 2, 2, 150, 151, 3, 2,
	2, 2, 151, 152, 3, 2, 2, 2, 152, 155, 7, 12, 2, 2, 153, 155, 7, 15, 2,
	2, 154, 150, 3, 2, 2, 2, 154, 153, 3, 2, 2, 2, 155, 156, 3, 2, 2, 2, 156,
	154, 3, 2, 2, 2, 156, 157, 3, 2, 2, 2, 157, 158, 3, 2, 2, 2, 158, 159,
	8, 23, 2, 2, 159, 46, 3, 2, 2, 2, 160, 161, 7, 49, 2, 2, 161, 162, 7, 44,
	2, 2, 162, 166, 3, 2, 2, 2, 163, 165, 11, 2, 2, 2, 164, 163, 3, 2, 2, 2,
	165, 168, 3, 2, 2, 2, 166, 167, 3, 2, 2, 2, 166, 164, 3, 2, 2, 2, 167,
	169, 3, 2, 2, 2, 168, 166, 3, 2, 2, 2, 169, 170, 7, 44, 2, 2, 170, 171,
	7, 49, 2, 2, 171, 172, 3, 2, 2, 2, 172, 173, 8, 24, 2, 2, 173, 48, 3, 2,
	2, 2, 174, 175, 7, 49, 2, 2, 175, 176, 7, 49, 2, 2, 176, 180, 3, 2, 2,
	2, 177, 179, 10, 7, 2, 2, 178, 177, 3, 2, 2, 2, 179, 182, 3, 2, 2, 2, 180,
	178, 3, 2, 2, 2, 180, 181, 3, 2, 2, 2, 181, 183, 3, 2, 2, 2, 182, 180,
	3, 2, 2, 2, 183, 184, 8, 25, 2, 2, 184, 50, 3, 2, 2, 2, 13, 2, 125, 127,
	133, 135, 145, 150, 154, 156, 166, 180, 3, 8, 2, 2,
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
	"", "'schema'", "'{'", "'}'", "'type'", "'ID'", "'int'", "'float'", "'string'",
	"'bool'", "'datetime'", "'geopoint'", "", "", "", "'['", "']'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "TypeIdentArray", "TypeIdent",
	"FieldIdent", "LBRACK", "RBRACK", "WS", "NEWLINE", "COMMENT", "LINE_COMMENT",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "ULetter", "LLetter", "Letter", "Digit", "TypeIdentArray",
	"TypeIdent", "FieldIdent", "LBRACK", "RBRACK", "WS", "NEWLINE", "COMMENT",
	"LINE_COMMENT",
}

type SchemaLexer struct {
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

func NewSchemaLexer(input antlr.CharStream) *SchemaLexer {

	l := new(SchemaLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Schema.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SchemaLexer tokens.
const (
	SchemaLexerT__0           = 1
	SchemaLexerT__1           = 2
	SchemaLexerT__2           = 3
	SchemaLexerT__3           = 4
	SchemaLexerT__4           = 5
	SchemaLexerT__5           = 6
	SchemaLexerT__6           = 7
	SchemaLexerT__7           = 8
	SchemaLexerT__8           = 9
	SchemaLexerT__9           = 10
	SchemaLexerT__10          = 11
	SchemaLexerTypeIdentArray = 12
	SchemaLexerTypeIdent      = 13
	SchemaLexerFieldIdent     = 14
	SchemaLexerLBRACK         = 15
	SchemaLexerRBRACK         = 16
	SchemaLexerWS             = 17
	SchemaLexerNEWLINE        = 18
	SchemaLexerCOMMENT        = 19
	SchemaLexerLINE_COMMENT   = 20
)
