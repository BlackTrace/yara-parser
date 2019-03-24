// Code generated by goyacc -p xx -o /grammar/parser.go /grammar/grammar.y. DO NOT EDIT.

//line /grammar/grammar.y:31
package grammar

import __yyfmt__ "fmt"

//line /grammar/grammar.y:31
import (
	"fmt"

	"github.com/Northern-Lights/yara-parser/yara"
)

var ParsedRuleset yara.RuleSet

type regexPair struct {
	text string
	mods yara.StringModifiers
}

//line /grammar/grammar.y:134
type xxSymType struct {
	yys int
	f64 float64
	i64 int64
	s   string
	ss  []string

	pxkw  yara.PrimaryExpressionKeyword
	txt   yara.TextString
	num   yara.Number
	dub   yara.Double
	binex yara.BinaryExpression
	rng   yara.Range
	expr  yara.Expression
	pexpr yara.PrimaryExpression
	rm    yara.RuleModifiers
	m     yara.Metas
	mp    yara.Meta
	mps   yara.Metas
	mod   yara.StringModifiers
	reg   regexPair
	ys    yara.String
	yss   yara.Strings
	yr    yara.Rule
}

const _END_OF_INCLUDED_FILE_ = 57346
const _DOT_DOT_ = 57347
const _RULE_ = 57348
const _PRIVATE_ = 57349
const _GLOBAL_ = 57350
const _META_ = 57351
const _STRINGS_ = 57352
const _CONDITION_ = 57353
const _IDENTIFIER_ = 57354
const _STRING_IDENTIFIER_ = 57355
const _STRING_COUNT_ = 57356
const _STRING_OFFSET_ = 57357
const _STRING_LENGTH_ = 57358
const _STRING_IDENTIFIER_WITH_WILDCARD_ = 57359
const _NUMBER_ = 57360
const _DOUBLE_ = 57361
const _INTEGER_FUNCTION_ = 57362
const _TEXT_STRING_ = 57363
const _HEX_STRING_ = 57364
const _REGEXP_ = 57365
const _ASCII_ = 57366
const _WIDE_ = 57367
const _XOR_ = 57368
const _NOCASE_ = 57369
const _FULLWORD_ = 57370
const _AT_ = 57371
const _FILESIZE_ = 57372
const _ENTRYPOINT_ = 57373
const _ALL_ = 57374
const _ANY_ = 57375
const _IN_ = 57376
const _OF_ = 57377
const _FOR_ = 57378
const _THEM_ = 57379
const _MATCHES_ = 57380
const _CONTAINS_ = 57381
const _IMPORT_ = 57382
const _TRUE_ = 57383
const _FALSE_ = 57384
const _LBRACE_ = 57385
const _RBRACE_ = 57386
const _INCLUDE_ = 57387
const _OR_ = 57388
const _AND_ = 57389
const _EQ_ = 57390
const _NEQ_ = 57391
const _LT_ = 57392
const _LE_ = 57393
const _GT_ = 57394
const _GE_ = 57395
const _SHIFT_LEFT_ = 57396
const _SHIFT_RIGHT_ = 57397
const _NOT_ = 57398
const UNARY_MINUS = 57399

var xxToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"_END_OF_INCLUDED_FILE_",
	"_DOT_DOT_",
	"_RULE_",
	"_PRIVATE_",
	"_GLOBAL_",
	"_META_",
	"_STRINGS_",
	"_CONDITION_",
	"_IDENTIFIER_",
	"_STRING_IDENTIFIER_",
	"_STRING_COUNT_",
	"_STRING_OFFSET_",
	"_STRING_LENGTH_",
	"_STRING_IDENTIFIER_WITH_WILDCARD_",
	"_NUMBER_",
	"_DOUBLE_",
	"_INTEGER_FUNCTION_",
	"_TEXT_STRING_",
	"_HEX_STRING_",
	"_REGEXP_",
	"_ASCII_",
	"_WIDE_",
	"_XOR_",
	"_NOCASE_",
	"_FULLWORD_",
	"_AT_",
	"_FILESIZE_",
	"_ENTRYPOINT_",
	"_ALL_",
	"_ANY_",
	"_IN_",
	"_OF_",
	"_FOR_",
	"_THEM_",
	"_MATCHES_",
	"_CONTAINS_",
	"_IMPORT_",
	"_TRUE_",
	"_FALSE_",
	"_LBRACE_",
	"_RBRACE_",
	"_INCLUDE_",
	"_OR_",
	"_AND_",
	"'|'",
	"'^'",
	"'&'",
	"_EQ_",
	"_NEQ_",
	"_LT_",
	"_LE_",
	"_GT_",
	"_GE_",
	"_SHIFT_LEFT_",
	"_SHIFT_RIGHT_",
	"'+'",
	"'-'",
	"'*'",
	"'\\\\'",
	"'%'",
	"_NOT_",
	"'~'",
	"UNARY_MINUS",
	"':'",
	"'='",
	"'.'",
	"'['",
	"']'",
	"'('",
	"')'",
	"','",
}
var xxStatenames = [...]string{}

const xxEofCode = 1
const xxErrCode = 2
const xxInitialStackSize = 16

//line /grammar/grammar.y:779

//line yacctab:1
var xxExca = [...]int{
	-1, 1,
	1, -1,
	-2, 15,
	-1, 42,
	23, 34,
	-2, 32,
	-1, 52,
	35, 93,
	-2, 79,
	-1, 107,
	35, 93,
	-2, 79,
	-1, 158,
	73, 50,
	74, 50,
	-2, 53,
	-1, 186,
	73, 51,
	74, 51,
	-2, 53,
}

const xxPrivate = 57344

const xxLast = 405

var xxAct = [...]int{

	52, 183, 108, 140, 49, 146, 159, 172, 73, 53,
	64, 65, 66, 171, 61, 62, 60, 63, 149, 74,
	70, 201, 202, 192, 193, 204, 58, 59, 71, 72,
	80, 79, 54, 112, 113, 198, 114, 50, 51, 96,
	94, 95, 148, 48, 80, 79, 189, 141, 97, 98,
	89, 90, 91, 92, 93, 102, 68, 208, 107, 105,
	56, 69, 106, 42, 194, 109, 195, 111, 57, 115,
	116, 206, 110, 79, 38, 181, 40, 147, 89, 90,
	91, 92, 93, 122, 123, 124, 125, 126, 127, 128,
	129, 130, 131, 132, 133, 134, 135, 136, 137, 138,
	139, 28, 121, 26, 145, 17, 91, 92, 93, 55,
	151, 152, 153, 39, 155, 18, 81, 82, 104, 158,
	80, 79, 161, 162, 164, 160, 96, 94, 95, 87,
	88, 83, 85, 84, 86, 97, 98, 89, 90, 91,
	92, 93, 163, 97, 98, 89, 90, 91, 92, 93,
	165, 150, 73, 5, 64, 65, 66, 35, 61, 62,
	60, 63, 44, 74, 101, 43, 142, 30, 117, 74,
	58, 59, 71, 72, 99, 118, 143, 186, 77, 100,
	187, 13, 8, 184, 190, 46, 47, 185, 154, 7,
	197, 78, 41, 36, 4, 199, 31, 37, 23, 144,
	68, 203, 20, 205, 45, 69, 14, 207, 33, 182,
	25, 73, 103, 64, 65, 66, 22, 61, 62, 60,
	63, 166, 74, 96, 94, 95, 9, 11, 12, 58,
	59, 196, 97, 98, 89, 90, 91, 92, 93, 120,
	96, 94, 95, 175, 174, 178, 176, 177, 167, 97,
	98, 89, 90, 91, 92, 93, 119, 191, 200, 68,
	188, 180, 157, 156, 69, 150, 96, 94, 95, 67,
	76, 103, 75, 27, 15, 97, 98, 89, 90, 91,
	92, 93, 96, 94, 95, 1, 32, 173, 34, 170,
	24, 97, 98, 89, 90, 91, 92, 93, 96, 94,
	95, 179, 29, 21, 19, 169, 16, 97, 98, 89,
	90, 91, 92, 93, 81, 82, 6, 10, 2, 3,
	0, 168, 0, 0, 96, 94, 95, 87, 88, 83,
	85, 84, 86, 97, 98, 89, 90, 91, 92, 93,
	0, 0, 0, 0, 96, 94, 95, 0, 0, 0,
	0, 0, 0, 97, 98, 89, 90, 91, 92, 93,
	96, 94, 95, 0, 0, 0, 0, 0, 0, 97,
	98, 89, 90, 91, 92, 93, 94, 95, 0, 0,
	0, 0, 0, 0, 97, 98, 89, 90, 91, 92,
	93, 95, 0, 0, 0, 0, 0, 0, 97, 98,
	89, 90, 91, 92, 93,
}
var xxPact = [...]int{

	-1000, 149, -1000, -1000, 161, -1000, 220, 160, -1000, 194,
	-1000, -1000, -1000, -1000, -1000, 38, 72, 190, 207, 186,
	-1000, 200, 36, -1000, -1000, 34, 184, 197, 180, 184,
	-1000, 6, 69, 9, 180, -1000, -5, -1000, 144, -1000,
	-4, -1000, 156, -1000, -1000, 173, -1000, -1000, 74, -1000,
	-1000, -1000, 276, 145, 140, 83, -4, -4, -1000, -1000,
	-7, -1000, -1000, -1000, -1000, 2, -3, -36, 199, 199,
	-1000, -1000, -1000, -1000, -1000, 147, 152, -1000, -1000, -1000,
	-1000, 146, 199, 199, 199, 199, 199, 199, 199, 199,
	199, 199, 199, 199, 199, 199, 199, 199, 199, 199,
	-25, 164, 312, 199, 5, -1000, -55, 78, 74, 199,
	199, 199, 176, 199, -4, -1000, -1000, -1000, -1000, -4,
	-4, -1000, 312, 312, 312, 312, 312, 312, 312, 45,
	45, -1000, -1000, -1000, 341, 86, 327, 19, 19, 312,
	-1000, 199, -1000, 90, 5, 192, -1000, -1000, -1000, -1000,
	-1000, 175, 250, 234, -1000, 218, -60, -67, -1000, 219,
	219, -1000, 26, 296, -1000, 8, 170, -1000, -1000, -1000,
	-1000, -1000, -4, -1000, -1000, -1000, -1000, -1000, -1000, 199,
	-26, -1000, -50, -1000, -1000, -1000, -1000, -9, -1, 199,
	-1000, -37, -1000, 170, -1000, -1000, -52, 296, -4, -1000,
	-47, -1000, 199, -2, -4, 312, -1000, -16, -1000,
}
var xxPgo = [...]int{

	0, 319, 318, 317, 316, 306, 304, 303, 302, 167,
	290, 288, 157, 287, 6, 286, 2, 4, 0, 3,
	285, 274, 273, 272, 270, 269, 263, 262, 20, 109,
	261, 260, 258, 5, 257, 256, 239, 231, 221, 209,
	1,
}
var xxR1 = [...]int{

	0, 20, 20, 20, 20, 20, 1, 21, 22, 2,
	7, 7, 10, 10, 15, 4, 4, 3, 3, 5,
	5, 6, 6, 8, 8, 9, 9, 9, 9, 9,
	11, 11, 23, 12, 24, 12, 12, 14, 14, 13,
	13, 13, 13, 13, 25, 25, 25, 25, 26, 26,
	27, 27, 28, 16, 17, 17, 17, 17, 17, 17,
	17, 17, 30, 32, 17, 34, 17, 17, 17, 35,
	17, 36, 17, 17, 17, 17, 17, 17, 17, 17,
	17, 31, 31, 19, 37, 37, 38, 33, 33, 39,
	39, 40, 40, 29, 29, 29, 18, 18, 18, 18,
	18, 18, 18, 18, 18, 18, 18, 18, 18, 18,
	18, 18, 18, 18, 18, 18, 18, 18, 18, 18,
	18, 18,
}
var xxR2 = [...]int{

	0, 0, 2, 2, 3, 2, 2, 0, 0, 11,
	0, 3, 0, 3, 3, 0, 2, 1, 1, 0,
	2, 1, 2, 1, 2, 3, 3, 4, 3, 3,
	1, 2, 0, 5, 0, 5, 3, 0, 2, 1,
	1, 1, 1, 1, 1, 3, 4, 4, 0, 1,
	1, 3, 1, 1, 1, 1, 3, 3, 1, 3,
	3, 3, 0, 0, 11, 0, 9, 3, 2, 0,
	4, 0, 4, 3, 3, 3, 3, 3, 3, 1,
	3, 3, 1, 5, 1, 3, 0, 4, 1, 1,
	3, 1, 1, 1, 1, 1, 3, 1, 1, 4,
	1, 1, 1, 1, 4, 1, 4, 1, 1, 2,
	3, 3, 3, 3, 3, 3, 3, 3, 2, 3,
	3, 1,
}
var xxChk = [...]int{

	-1000, -20, -2, -1, 45, 4, -4, 40, 21, 6,
	-3, 7, 8, 21, 12, -21, -5, 67, 43, -6,
	12, -7, 9, 12, -10, 10, 67, -22, 67, -8,
	-9, 12, -15, 11, -11, -12, 13, -9, 68, 44,
	67, -12, 68, 21, 18, 60, 41, 42, -16, -17,
	41, 42, -18, 13, 36, -29, 64, 72, 30, 31,
	20, 18, 19, 21, 14, 15, 16, -25, 60, 65,
	-28, 32, 33, 12, 23, -23, -24, 22, 18, 47,
	46, 38, 39, 53, 55, 54, 56, 51, 52, 59,
	60, 61, 62, 63, 49, 50, 48, 57, 58, 29,
	34, -29, -18, 72, 35, -16, -17, -18, -16, 72,
	70, 70, 69, 70, 72, -18, -18, 21, 23, -35,
	-36, -28, -18, -18, -18, -18, -18, -18, -18, -18,
	-18, -18, -18, -18, -18, -18, -18, -18, -18, -18,
	-19, 72, 2, 12, 35, -18, -33, 72, 37, 73,
	73, -18, -18, -18, 12, -18, -26, -27, -17, -14,
	-14, -16, -16, -18, 34, -33, -38, 73, 71, 71,
	71, 73, 74, -13, 25, 24, 27, 28, 26, 5,
	-30, 67, -39, -40, 13, 17, -17, -18, -31, 72,
	-19, -34, 73, 74, 73, 67, -37, -18, 72, -40,
	-32, 73, 74, -16, 72, -18, 73, -16, 73,
}
var xxDef = [...]int{

	1, -2, 2, 3, 0, 5, 0, 0, 4, 0,
	16, 17, 18, 6, 7, 19, 0, 0, 10, 20,
	21, 12, 0, 22, 8, 0, 0, 0, 0, 11,
	23, 0, 0, 0, 13, 30, 0, 24, 0, 9,
	0, 31, -2, 25, 26, 0, 28, 29, 14, 53,
	54, 55, -2, 58, 0, 0, 0, 0, 97, 98,
	0, 100, 101, 102, 103, 105, 107, 108, 0, 0,
	121, 94, 95, 44, 52, 0, 0, 36, 27, 69,
	71, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 93, 0, 0, 68, 53, -2, 0, 0,
	0, 0, 0, 0, 48, 109, 118, 37, 37, 0,
	0, 56, 57, 73, 74, 75, 76, 77, 78, 110,
	111, 112, 113, 114, 115, 116, 117, 119, 120, 59,
	60, 0, 61, 0, 0, 0, 67, 86, 88, 80,
	96, 0, 0, 0, 45, 0, 0, 49, -2, 33,
	35, 70, 72, 0, 62, 0, 0, 99, 104, 106,
	46, 47, 0, 38, 39, 40, 41, 42, 43, 0,
	0, 65, 0, 89, 91, 92, -2, 0, 0, 0,
	82, 0, 87, 0, 83, 63, 0, 84, 0, 90,
	0, 81, 0, 0, 0, 85, 66, 0, 64,
}
var xxTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 63, 50, 3,
	72, 73, 61, 59, 74, 60, 69, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 67, 3,
	3, 68, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 70, 62, 71, 49, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 48, 3, 65,
}
var xxTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 51, 52, 53, 54,
	55, 56, 57, 58, 64, 66,
}
var xxTok3 = [...]int{
	0,
}

var xxErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	xxDebug        = 0
	xxErrorVerbose = false
)

type xxLexer interface {
	Lex(lval *xxSymType) int
	Error(s string)
}

type xxParser interface {
	Parse(xxLexer) int
	Lookahead() int
}

type xxParserImpl struct {
	lval  xxSymType
	stack [xxInitialStackSize]xxSymType
	char  int
}

func (p *xxParserImpl) Lookahead() int {
	return p.char
}

func xxNewParser() xxParser {
	return &xxParserImpl{}
}

const xxFlag = -1000

func xxTokname(c int) string {
	if c >= 1 && c-1 < len(xxToknames) {
		if xxToknames[c-1] != "" {
			return xxToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func xxStatname(s int) string {
	if s >= 0 && s < len(xxStatenames) {
		if xxStatenames[s] != "" {
			return xxStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func xxErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !xxErrorVerbose {
		return "syntax error"
	}

	for _, e := range xxErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + xxTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := xxPact[state]
	for tok := TOKSTART; tok-1 < len(xxToknames); tok++ {
		if n := base + tok; n >= 0 && n < xxLast && xxChk[xxAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if xxDef[state] == -2 {
		i := 0
		for xxExca[i] != -1 || xxExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; xxExca[i] >= 0; i += 2 {
			tok := xxExca[i]
			if tok < TOKSTART || xxExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if xxExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += xxTokname(tok)
	}
	return res
}

func xxlex1(lex xxLexer, lval *xxSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = xxTok1[0]
		goto out
	}
	if char < len(xxTok1) {
		token = xxTok1[char]
		goto out
	}
	if char >= xxPrivate {
		if char < xxPrivate+len(xxTok2) {
			token = xxTok2[char-xxPrivate]
			goto out
		}
	}
	for i := 0; i < len(xxTok3); i += 2 {
		token = xxTok3[i+0]
		if token == char {
			token = xxTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = xxTok2[1] /* unknown char */
	}
	if xxDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", xxTokname(token), uint(char))
	}
	return char, token
}

func xxParse(xxlex xxLexer) int {
	return xxNewParser().Parse(xxlex)
}

func (xxrcvr *xxParserImpl) Parse(xxlex xxLexer) int {
	var xxn int
	var xxVAL xxSymType
	var xxDollar []xxSymType
	_ = xxDollar // silence set and not used
	xxS := xxrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	xxstate := 0
	xxrcvr.char = -1
	xxtoken := -1 // xxrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		xxstate = -1
		xxrcvr.char = -1
		xxtoken = -1
	}()
	xxp := -1
	goto xxstack

ret0:
	return 0

ret1:
	return 1

xxstack:
	/* put a state and value onto the stack */
	if xxDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", xxTokname(xxtoken), xxStatname(xxstate))
	}

	xxp++
	if xxp >= len(xxS) {
		nyys := make([]xxSymType, len(xxS)*2)
		copy(nyys, xxS)
		xxS = nyys
	}
	xxS[xxp] = xxVAL
	xxS[xxp].yys = xxstate

xxnewstate:
	xxn = xxPact[xxstate]
	if xxn <= xxFlag {
		goto xxdefault /* simple state */
	}
	if xxrcvr.char < 0 {
		xxrcvr.char, xxtoken = xxlex1(xxlex, &xxrcvr.lval)
	}
	xxn += xxtoken
	if xxn < 0 || xxn >= xxLast {
		goto xxdefault
	}
	xxn = xxAct[xxn]
	if xxChk[xxn] == xxtoken { /* valid shift */
		xxrcvr.char = -1
		xxtoken = -1
		xxVAL = xxrcvr.lval
		xxstate = xxn
		if Errflag > 0 {
			Errflag--
		}
		goto xxstack
	}

xxdefault:
	/* default state action */
	xxn = xxDef[xxstate]
	if xxn == -2 {
		if xxrcvr.char < 0 {
			xxrcvr.char, xxtoken = xxlex1(xxlex, &xxrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if xxExca[xi+0] == -1 && xxExca[xi+1] == xxstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			xxn = xxExca[xi+0]
			if xxn < 0 || xxn == xxtoken {
				break
			}
		}
		xxn = xxExca[xi+1]
		if xxn < 0 {
			goto ret0
		}
	}
	if xxn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			xxlex.Error(xxErrorMessage(xxstate, xxtoken))
			Nerrs++
			if xxDebug >= 1 {
				__yyfmt__.Printf("%s", xxStatname(xxstate))
				__yyfmt__.Printf(" saw %s\n", xxTokname(xxtoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for xxp >= 0 {
				xxn = xxPact[xxS[xxp].yys] + xxErrCode
				if xxn >= 0 && xxn < xxLast {
					xxstate = xxAct[xxn] /* simulate a shift of "error" */
					if xxChk[xxstate] == xxErrCode {
						goto xxstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if xxDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", xxS[xxp].yys)
				}
				xxp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if xxDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", xxTokname(xxtoken))
			}
			if xxtoken == xxEofCode {
				goto ret1
			}
			xxrcvr.char = -1
			xxtoken = -1
			goto xxnewstate /* try again in the same state */
		}
	}

	/* reduction by production xxn */
	if xxDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", xxn, xxStatname(xxstate))
	}

	xxnt := xxn
	xxpt := xxp
	_ = xxpt // guard against "declared and not used"

	xxp -= xxR2[xxn]
	// xxp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if xxp+1 >= len(xxS) {
		nyys := make([]xxSymType, len(xxS)*2)
		copy(nyys, xxS)
		xxS = nyys
	}
	xxVAL = xxS[xxp+1]

	/* consult goto table to find next state */
	xxn = xxR1[xxn]
	xxg := xxPgo[xxn]
	xxj := xxg + xxS[xxp].yys + 1

	if xxj >= xxLast {
		xxstate = xxAct[xxg]
	} else {
		xxstate = xxAct[xxj]
		if xxChk[xxstate] != -xxn {
			xxstate = xxAct[xxg]
		}
	}
	// dummy call; replaced with literal code
	switch xxnt {

	case 2:
		xxDollar = xxS[xxpt-2 : xxpt+1]
//line /grammar/grammar.y:164
		{
			ParsedRuleset.Rules = append(ParsedRuleset.Rules, xxDollar[2].yr)
		}
	case 3:
		xxDollar = xxS[xxpt-2 : xxpt+1]
//line /grammar/grammar.y:167
		{
			ParsedRuleset.Imports = append(ParsedRuleset.Imports, xxDollar[2].s)
		}
	case 4:
		xxDollar = xxS[xxpt-3 : xxpt+1]
//line /grammar/grammar.y:170
		{
			ParsedRuleset.Includes = append(ParsedRuleset.Includes, xxDollar[3].s)
		}
	case 5:
		xxDollar = xxS[xxpt-2 : xxpt+1]
//line /grammar/grammar.y:173
		{
		}
	case 6:
		xxDollar = xxS[xxpt-2 : xxpt+1]
//line /grammar/grammar.y:179
		{
			xxVAL.s = xxDollar[2].s
		}
	case 7:
		xxDollar = xxS[xxpt-3 : xxpt+1]
//line /grammar/grammar.y:187
		{
			xxVAL.yr.Modifiers = xxDollar[1].rm
			xxVAL.yr.Identifier = xxDollar[3].s

			// Forbid duplicate rules
			for _, r := range ParsedRuleset.Rules {
				if xxDollar[3].s == r.Identifier {
					err := fmt.Errorf(`Duplicate rule "%s"`, xxDollar[3].s)
					panic(err)
				}
			}
		}
	case 8:
		xxDollar = xxS[xxpt-8 : xxpt+1]
//line /grammar/grammar.y:200
		{
			// $4 is the rule created in above action
			xxDollar[4].yr.Tags = xxDollar[5].ss

			// Forbid duplicate tags
			idx := make(map[string]struct{})
			for _, t := range xxDollar[5].ss {
				if _, had := idx[t]; had {
					msg := fmt.Sprintf(`grammar: Rule "%s" has duplicate tag "%s"`,
						xxDollar[4].yr.Identifier,
						t)
					panic(msg)
				}
				idx[t] = struct{}{}
			}

			xxDollar[4].yr.Meta = xxDollar[7].m

			xxDollar[4].yr.Strings = xxDollar[8].yss

			// Forbid duplicate string IDs, except `$` (anonymous)
			idx = make(map[string]struct{})
			for _, s := range xxDollar[8].yss {
				if s.ID == "$" {
					continue
				}
				if _, had := idx[s.ID]; had {
					msg := fmt.Sprintf(
						`grammar: Rule "%s" has duplicated string "%s"`,
						xxDollar[4].yr.Identifier,
						s.ID)
					panic(msg)
				}
				idx[s.ID] = struct{}{}
			}
		}
	case 9:
		xxDollar = xxS[xxpt-11 : xxpt+1]
//line /grammar/grammar.y:237
		{
			xxDollar[4].yr.Condition = xxDollar[10].expr
			xxVAL.yr = xxDollar[4].yr
		}
	case 10:
		xxDollar = xxS[xxpt-0 : xxpt+1]
//line /grammar/grammar.y:246
		{

		}
	case 11:
		xxDollar = xxS[xxpt-3 : xxpt+1]
//line /grammar/grammar.y:250
		{
			xxVAL.m = make(yara.Metas, 0, len(xxDollar[3].mps))
			for _, mpair := range xxDollar[3].mps {
				// YARA is ok with duplicate keys; we follow suit
				xxVAL.m = append(xxVAL.m, mpair)
			}
		}
	case 12:
		xxDollar = xxS[xxpt-0 : xxpt+1]
//line /grammar/grammar.y:262
		{
			xxVAL.yss = yara.Strings{}
		}
	case 13:
		xxDollar = xxS[xxpt-3 : xxpt+1]
//line /grammar/grammar.y:266
		{
			xxVAL.yss = xxDollar[3].yss
		}
	case 14:
		xxDollar = xxS[xxpt-3 : xxpt+1]
//line /grammar/grammar.y:274
		{
			xxVAL.expr = xxDollar[3].expr
		}
	case 15:
		xxDollar = xxS[xxpt-0 : xxpt+1]
//line /grammar/grammar.y:281
		{
			xxVAL.rm = yara.RuleModifiers{}
		}
	case 16:
		xxDollar = xxS[xxpt-2 : xxpt+1]
//line /grammar/grammar.y:282
		{
			xxVAL.rm.Private = xxVAL.rm.Private || xxDollar[2].rm.Private
			xxVAL.rm.Global = xxVAL.rm.Global || xxDollar[2].rm.Global
		}
	case 17:
		xxDollar = xxS[xxpt-1 : xxpt+1]
//line /grammar/grammar.y:290
		{
			xxVAL.rm.Private = true
		}
	case 18:
		xxDollar = xxS[xxpt-1 : xxpt+1]
//line /grammar/grammar.y:291
		{
			xxVAL.rm.Global = true
		}
	case 19:
		xxDollar = xxS[xxpt-0 : xxpt+1]
//line /grammar/grammar.y:297
		{
			xxVAL.ss = []string{}
		}
	case 20:
		xxDollar = xxS[xxpt-2 : xxpt+1]
//line /grammar/grammar.y:301
		{
			xxVAL.ss = xxDollar[2].ss
		}
	case 21:
		xxDollar = xxS[xxpt-1 : xxpt+1]
//line /grammar/grammar.y:309
		{
			xxVAL.ss = []string{xxDollar[1].s}
		}
	case 22:
		xxDollar = xxS[xxpt-2 : xxpt+1]
//line /grammar/grammar.y:313
		{
			xxVAL.ss = append(xxDollar[1].ss, xxDollar[2].s)
		}
	case 23:
		xxDollar = xxS[xxpt-1 : xxpt+1]
//line /grammar/grammar.y:321
		{
			xxVAL.mps = yara.Metas{xxDollar[1].mp}
		}
	case 24:
		xxDollar = xxS[xxpt-2 : xxpt+1]
//line /grammar/grammar.y:322
		{
			xxVAL.mps = append(xxVAL.mps, xxDollar[2].mp)
		}
	case 25:
		xxDollar = xxS[xxpt-3 : xxpt+1]
//line /grammar/grammar.y:328
		{
			xxVAL.mp = yara.Meta{xxDollar[1].s, xxDollar[3].s}
		}
	case 26:
		xxDollar = xxS[xxpt-3 : xxpt+1]
//line /grammar/grammar.y:332
		{
			xxVAL.mp = yara.Meta{xxDollar[1].s, xxDollar[3].i64}
		}
	case 27:
		xxDollar = xxS[xxpt-4 : xxpt+1]
//line /grammar/grammar.y:336
		{
			xxVAL.mp = yara.Meta{xxDollar[1].s, -xxDollar[4].i64}
		}
	case 28:
		xxDollar = xxS[xxpt-3 : xxpt+1]
//line /grammar/grammar.y:340
		{
			xxVAL.mp = yara.Meta{xxDollar[1].s, true}
		}
	case 29:
		xxDollar = xxS[xxpt-3 : xxpt+1]
//line /grammar/grammar.y:344
		{
			xxVAL.mp = yara.Meta{xxDollar[1].s, false}
		}
	case 30:
		xxDollar = xxS[xxpt-1 : xxpt+1]
//line /grammar/grammar.y:351
		{
			xxVAL.yss = yara.Strings{xxDollar[1].ys}
		}
	case 31:
		xxDollar = xxS[xxpt-2 : xxpt+1]
//line /grammar/grammar.y:352
		{
			xxVAL.yss = append(xxDollar[1].yss, xxDollar[2].ys)
		}
	case 32:
		xxDollar = xxS[xxpt-2 : xxpt+1]
//line /grammar/grammar.y:358
		{
			xxVAL.ys.Type = yara.TypeString
			xxVAL.ys.ID = xxDollar[1].s
		}
	case 33:
		xxDollar = xxS[xxpt-5 : xxpt+1]
//line /grammar/grammar.y:363
		{
			xxDollar[3].ys.Text = xxDollar[4].s
			xxDollar[3].ys.Modifiers = xxDollar[5].mod

			xxVAL.ys = xxDollar[3].ys
		}
	case 34:
		xxDollar = xxS[xxpt-2 : xxpt+1]
//line /grammar/grammar.y:370
		{
			xxVAL.ys.Type = yara.TypeRegex
			xxVAL.ys.ID = xxDollar[1].s
		}
	case 35:
		xxDollar = xxS[xxpt-5 : xxpt+1]
//line /grammar/grammar.y:375
		{
			xxDollar[3].ys.Text = xxDollar[4].reg.text

			xxDollar[5].mod.I = xxDollar[4].reg.mods.I
			xxDollar[5].mod.S = xxDollar[4].reg.mods.S

			xxDollar[3].ys.Modifiers = xxDollar[5].mod

			xxVAL.ys = xxDollar[3].ys
		}
	case 36:
		xxDollar = xxS[xxpt-3 : xxpt+1]
//line /grammar/grammar.y:386
		{
			xxVAL.ys.Type = yara.TypeHexString
			xxVAL.ys.ID = xxDollar[1].s
			xxVAL.ys.Text = xxDollar[3].s
		}
	case 37:
		xxDollar = xxS[xxpt-0 : xxpt+1]
//line /grammar/grammar.y:395
		{
			xxVAL.mod = yara.StringModifiers{}
		}
	case 38:
		xxDollar = xxS[xxpt-2 : xxpt+1]
//line /grammar/grammar.y:398
		{
			xxVAL.mod = yara.StringModifiers{
				Wide:     xxDollar[1].mod.Wide || xxDollar[2].mod.Wide,
				ASCII:    xxDollar[1].mod.ASCII || xxDollar[2].mod.ASCII,
				Nocase:   xxDollar[1].mod.Nocase || xxDollar[2].mod.Nocase,
				Fullword: xxDollar[1].mod.Fullword || xxDollar[2].mod.Fullword,
				Xor:      xxDollar[1].mod.Xor || xxDollar[2].mod.Xor,
			}
		}
	case 39:
		xxDollar = xxS[xxpt-1 : xxpt+1]
//line /grammar/grammar.y:411
		{
			xxVAL.mod.Wide = true
		}
	case 40:
		xxDollar = xxS[xxpt-1 : xxpt+1]
//line /grammar/grammar.y:412
		{
			xxVAL.mod.ASCII = true
		}
	case 41:
		xxDollar = xxS[xxpt-1 : xxpt+1]
//line /grammar/grammar.y:413
		{
			xxVAL.mod.Nocase = true
		}
	case 42:
		xxDollar = xxS[xxpt-1 : xxpt+1]
//line /grammar/grammar.y:414
		{
			xxVAL.mod.Fullword = true
		}
	case 43:
		xxDollar = xxS[xxpt-1 : xxpt+1]
//line /grammar/grammar.y:415
		{
			xxVAL.mod.Xor = true
		}
	case 44:
		xxDollar = xxS[xxpt-1 : xxpt+1]
//line /grammar/grammar.y:421
		{

		}
	case 45:
		xxDollar = xxS[xxpt-3 : xxpt+1]
//line /grammar/grammar.y:425
		{

		}
	case 46:
		xxDollar = xxS[xxpt-4 : xxpt+1]
//line /grammar/grammar.y:429
		{

		}
	case 47:
		xxDollar = xxS[xxpt-4 : xxpt+1]
//line /grammar/grammar.y:434
		{

		}
	case 48:
		xxDollar = xxS[xxpt-0 : xxpt+1]
//line /grammar/grammar.y:441
		{
		}
	case 49:
		xxDollar = xxS[xxpt-1 : xxpt+1]
//line /grammar/grammar.y:442
		{
		}
	case 50:
		xxDollar = xxS[xxpt-1 : xxpt+1]
//line /grammar/grammar.y:447
		{

		}
	case 51:
		xxDollar = xxS[xxpt-3 : xxpt+1]
//line /grammar/grammar.y:451
		{

		}
	case 52:
		xxDollar = xxS[xxpt-1 : xxpt+1]
//line /grammar/grammar.y:459
		{

		}
	case 53:
		xxDollar = xxS[xxpt-1 : xxpt+1]
//line /grammar/grammar.y:467
		{
			xxVAL.expr = xxDollar[1].expr
		}
	case 54:
		xxDollar = xxS[xxpt-1 : xxpt+1]
//line /grammar/grammar.y:474
		{
			xxVAL.expr = yara.Boolean(true)
		}
	case 55:
		xxDollar = xxS[xxpt-1 : xxpt+1]
//line /grammar/grammar.y:478
		{
			xxVAL.expr = yara.Boolean(false)
		}
	case 56:
		xxDollar = xxS[xxpt-3 : xxpt+1]
//line /grammar/grammar.y:482
		{

		}
	case 57:
		xxDollar = xxS[xxpt-3 : xxpt+1]
//line /grammar/grammar.y:486
		{

		}
	case 58:
		xxDollar = xxS[xxpt-1 : xxpt+1]
//line /grammar/grammar.y:490
		{
			xxVAL.expr = yara.StringIdentifier{
				Identifier: xxDollar[1].s,
			}
		}
	case 59:
		xxDollar = xxS[xxpt-3 : xxpt+1]
//line /grammar/grammar.y:496
		{
			sid := yara.StringIdentifier{
				Identifier: xxDollar[1].s,
			}
			sid.At(xxDollar[3].pexpr)
			xxVAL.expr = sid
		}
	case 60:
		xxDollar = xxS[xxpt-3 : xxpt+1]
//line /grammar/grammar.y:504
		{
			sid := yara.StringIdentifier{
				Identifier: xxDollar[1].s,
			}
			rangeCopy := xxDollar[3].rng
			sid.In(&rangeCopy)
			xxVAL.expr = sid
		}
	case 61:
		xxDollar = xxS[xxpt-3 : xxpt+1]
//line /grammar/grammar.y:513
		{

		}
	case 62:
		xxDollar = xxS[xxpt-4 : xxpt+1]
//line /grammar/grammar.y:517
		{

		}
	case 63:
		xxDollar = xxS[xxpt-7 : xxpt+1]
//line /grammar/grammar.y:521
		{

		}
	case 64:
		xxDollar = xxS[xxpt-11 : xxpt+1]
//line /grammar/grammar.y:525
		{

		}
	case 65:
		xxDollar = xxS[xxpt-5 : xxpt+1]
//line /grammar/grammar.y:529
		{

		}
	case 66:
		xxDollar = xxS[xxpt-9 : xxpt+1]
//line /grammar/grammar.y:533
		{

		}
	case 67:
		xxDollar = xxS[xxpt-3 : xxpt+1]
//line /grammar/grammar.y:537
		{

		}
	case 68:
		xxDollar = xxS[xxpt-2 : xxpt+1]
//line /grammar/grammar.y:541
		{

		}
	case 69:
		xxDollar = xxS[xxpt-2 : xxpt+1]
//line /grammar/grammar.y:545
		{

		}
	case 70:
		xxDollar = xxS[xxpt-4 : xxpt+1]
//line /grammar/grammar.y:549
		{
			exp1 := xxDollar[1].expr
			exp2 := xxDollar[4].expr
			xxVAL.expr = yara.And(exp1, exp2)
		}
	case 71:
		xxDollar = xxS[xxpt-2 : xxpt+1]
//line /grammar/grammar.y:555
		{

		}
	case 72:
		xxDollar = xxS[xxpt-4 : xxpt+1]
//line /grammar/grammar.y:559
		{
			exp1 := xxDollar[1].expr
			exp2 := xxDollar[4].expr
			xxVAL.expr = yara.Or(exp1, exp2)
		}
	case 73:
		xxDollar = xxS[xxpt-3 : xxpt+1]
//line /grammar/grammar.y:565
		{
			xxVAL.expr = yara.LT(xxDollar[1].pexpr, xxDollar[3].pexpr)
		}
	case 74:
		xxDollar = xxS[xxpt-3 : xxpt+1]
//line /grammar/grammar.y:569
		{
			xxVAL.expr = yara.GT(xxDollar[1].pexpr, xxDollar[3].pexpr)
		}
	case 75:
		xxDollar = xxS[xxpt-3 : xxpt+1]
//line /grammar/grammar.y:573
		{
			xxVAL.expr = yara.LE(xxDollar[1].pexpr, xxDollar[3].pexpr)
		}
	case 76:
		xxDollar = xxS[xxpt-3 : xxpt+1]
//line /grammar/grammar.y:577
		{
			xxVAL.expr = yara.GE(xxDollar[1].pexpr, xxDollar[3].pexpr)
		}
	case 77:
		xxDollar = xxS[xxpt-3 : xxpt+1]
//line /grammar/grammar.y:581
		{
			xxVAL.expr = yara.EQ(xxDollar[1].pexpr, xxDollar[3].pexpr)
		}
	case 78:
		xxDollar = xxS[xxpt-3 : xxpt+1]
//line /grammar/grammar.y:585
		{
			xxVAL.expr = yara.NEQ(xxDollar[1].pexpr, xxDollar[3].pexpr)
		}
	case 79:
		xxDollar = xxS[xxpt-1 : xxpt+1]
//line /grammar/grammar.y:589
		{

		}
	case 80:
		xxDollar = xxS[xxpt-3 : xxpt+1]
//line /grammar/grammar.y:593
		{
			xxVAL.expr = xxDollar[2].expr
		}
	case 81:
		xxDollar = xxS[xxpt-3 : xxpt+1]
//line /grammar/grammar.y:600
		{
		}
	case 82:
		xxDollar = xxS[xxpt-1 : xxpt+1]
//line /grammar/grammar.y:601
		{
		}
	case 83:
		xxDollar = xxS[xxpt-5 : xxpt+1]
//line /grammar/grammar.y:607
		{
			xxVAL.rng = yara.Range{
				Start: xxDollar[2].pexpr,
				End:   xxDollar[4].pexpr,
			}
		}
	case 84:
		xxDollar = xxS[xxpt-1 : xxpt+1]
//line /grammar/grammar.y:618
		{

		}
	case 85:
		xxDollar = xxS[xxpt-3 : xxpt+1]
//line /grammar/grammar.y:622
		{

		}
	case 86:
		xxDollar = xxS[xxpt-1 : xxpt+1]
//line /grammar/grammar.y:630
		{

		}
	case 88:
		xxDollar = xxS[xxpt-1 : xxpt+1]
//line /grammar/grammar.y:635
		{

		}
	case 91:
		xxDollar = xxS[xxpt-1 : xxpt+1]
//line /grammar/grammar.y:649
		{

		}
	case 92:
		xxDollar = xxS[xxpt-1 : xxpt+1]
//line /grammar/grammar.y:653
		{

		}
	case 94:
		xxDollar = xxS[xxpt-1 : xxpt+1]
//line /grammar/grammar.y:662
		{

		}
	case 95:
		xxDollar = xxS[xxpt-1 : xxpt+1]
//line /grammar/grammar.y:666
		{

		}
	case 96:
		xxDollar = xxS[xxpt-3 : xxpt+1]
//line /grammar/grammar.y:674
		{
			xxVAL.pexpr = xxDollar[2].pexpr
		}
	case 97:
		xxDollar = xxS[xxpt-1 : xxpt+1]
//line /grammar/grammar.y:678
		{
			xxVAL.pexpr = yara.Filesize
		}
	case 98:
		xxDollar = xxS[xxpt-1 : xxpt+1]
//line /grammar/grammar.y:682
		{
			xxVAL.pexpr = yara.Entrypoint
		}
	case 99:
		xxDollar = xxS[xxpt-4 : xxpt+1]
//line /grammar/grammar.y:686
		{

		}
	case 100:
		xxDollar = xxS[xxpt-1 : xxpt+1]
//line /grammar/grammar.y:690
		{
			xxVAL.pexpr = yara.Number(xxDollar[1].i64)
		}
	case 101:
		xxDollar = xxS[xxpt-1 : xxpt+1]
//line /grammar/grammar.y:694
		{
			xxVAL.pexpr = yara.Double(xxDollar[1].f64)
		}
	case 102:
		xxDollar = xxS[xxpt-1 : xxpt+1]
//line /grammar/grammar.y:698
		{
			xxVAL.pexpr = yara.TextString(xxDollar[1].s)
		}
	case 103:
		xxDollar = xxS[xxpt-1 : xxpt+1]
//line /grammar/grammar.y:702
		{

		}
	case 104:
		xxDollar = xxS[xxpt-4 : xxpt+1]
//line /grammar/grammar.y:706
		{

		}
	case 105:
		xxDollar = xxS[xxpt-1 : xxpt+1]
//line /grammar/grammar.y:710
		{

		}
	case 106:
		xxDollar = xxS[xxpt-4 : xxpt+1]
//line /grammar/grammar.y:714
		{

		}
	case 107:
		xxDollar = xxS[xxpt-1 : xxpt+1]
//line /grammar/grammar.y:718
		{

		}
	case 108:
		xxDollar = xxS[xxpt-1 : xxpt+1]
//line /grammar/grammar.y:722
		{

		}
	case 109:
		xxDollar = xxS[xxpt-2 : xxpt+1]
//line /grammar/grammar.y:726
		{

		}
	case 110:
		xxDollar = xxS[xxpt-3 : xxpt+1]
//line /grammar/grammar.y:730
		{

		}
	case 111:
		xxDollar = xxS[xxpt-3 : xxpt+1]
//line /grammar/grammar.y:734
		{

		}
	case 112:
		xxDollar = xxS[xxpt-3 : xxpt+1]
//line /grammar/grammar.y:738
		{

		}
	case 113:
		xxDollar = xxS[xxpt-3 : xxpt+1]
//line /grammar/grammar.y:742
		{

		}
	case 114:
		xxDollar = xxS[xxpt-3 : xxpt+1]
//line /grammar/grammar.y:746
		{

		}
	case 115:
		xxDollar = xxS[xxpt-3 : xxpt+1]
//line /grammar/grammar.y:750
		{

		}
	case 116:
		xxDollar = xxS[xxpt-3 : xxpt+1]
//line /grammar/grammar.y:754
		{

		}
	case 117:
		xxDollar = xxS[xxpt-3 : xxpt+1]
//line /grammar/grammar.y:758
		{

		}
	case 118:
		xxDollar = xxS[xxpt-2 : xxpt+1]
//line /grammar/grammar.y:762
		{

		}
	case 119:
		xxDollar = xxS[xxpt-3 : xxpt+1]
//line /grammar/grammar.y:766
		{

		}
	case 120:
		xxDollar = xxS[xxpt-3 : xxpt+1]
//line /grammar/grammar.y:770
		{

		}
	case 121:
		xxDollar = xxS[xxpt-1 : xxpt+1]
//line /grammar/grammar.y:774
		{

		}
	}
	goto xxstack /* stack new state and value */
}
