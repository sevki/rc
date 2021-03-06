// Code generated by goyacc syn.y. DO NOT EDIT.

//line syn.y:13
#include "rc.h"
#include "fns.h"

//line syn.y:16
type yySymType struct{
	yys int
	struct tree *tree;
}

const FOR = 57346
const IN = 57347
const WHILE = 57348
const IF = 57349
const NOT = 57350
const TWIDDLE = 57351
const BANG = 57352
const SUBSHELL = 57353
const SWITCH = 57354
const FN = 57355
const WORD = 57356
const REDIR = 57357
const DUP = 57358
const PIPE = 57359
const SUB = 57360
const SIMPLE = 57361
const ARGLIST = 57362
const WORDS = 57363
const BRACE = 57364
const PAREN = 57365
const PCMD = 57366
const PIPEFD = 57367
const ANDAND = 57368
const OROR = 57369
const COUNT = 57370

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"FOR",
	"IN",
	"WHILE",
	"IF",
	"NOT",
	"TWIDDLE",
	"BANG",
	"SUBSHELL",
	"SWITCH",
	"FN",
	"WORD",
	"REDIR",
	"DUP",
	"PIPE",
	"SUB",
	"SIMPLE",
	"ARGLIST",
	"WORDS",
	"BRACE",
	"PAREN",
	"PCMD",
	"PIPEFD",
	"')'",
	"ANDAND",
	"OROR",
	"'^'",
	"'$'",
	"COUNT",
	"'\"'",
	"'\\n'",
	"';'",
	"'&'",
	"'{'",
	"'}'",
	"'('",
	"'='",
	"'`'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int{
	-1, 0,
	1, 1,
	-2, 18,
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 327

var yyAct = [...]int{

	66, 3, 64, 40, 5, 3, 12, 65, 70, 39,
	41, 33, 36, 60, 61, 62, 63, 58, 69, 92,
	17, 31, 32, 35, 72, 28, 21, 93, 29, 30,
	77, 76, 78, 79, 80, 88, 45, 45, 45, 40,
	101, 97, 33, 36, 33, 38, 45, 91, 84, 45,
	45, 45, 31, 32, 42, 102, 87, 43, 57, 59,
	81, 72, 89, 33, 45, 86, 109, 71, 45, 90,
	73, 74, 75, 31, 32, 94, 103, 37, 20, 88,
	29, 30, 113, 99, 100, 71, 83, 104, 82, 85,
	2, 45, 105, 68, 4, 34, 45, 45, 4, 1,
	107, 44, 18, 10, 45, 108, 13, 67, 0, 0,
	112, 0, 0, 0, 114, 45, 45, 95, 96, 46,
	47, 48, 49, 50, 51, 52, 53, 54, 55, 25,
	56, 0, 0, 0, 45, 45, 106, 0, 0, 0,
	0, 111, 0, 0, 0, 22, 24, 23, 0, 0,
	0, 0, 0, 27, 0, 26, 46, 47, 48, 49,
	50, 51, 52, 53, 54, 55, 25, 56, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 110, 0,
	0, 0, 22, 24, 23, 0, 0, 0, 0, 0,
	27, 0, 26, 46, 47, 48, 49, 50, 51, 52,
	53, 54, 55, 25, 56, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 98, 0, 0, 0, 22,
	24, 23, 0, 0, 0, 0, 0, 27, 0, 26,
	46, 47, 48, 49, 50, 51, 52, 53, 54, 55,
	25, 56, 46, 47, 48, 49, 50, 51, 52, 53,
	54, 55, 25, 19, 20, 0, 22, 24, 23, 0,
	0, 0, 17, 0, 27, 0, 26, 0, 22, 24,
	23, 0, 0, 0, 0, 0, 27, 7, 26, 8,
	6, 0, 11, 14, 15, 9, 16, 25, 19, 20,
	46, 47, 48, 49, 50, 51, 52, 53, 54, 55,
	25, 56, 0, 22, 24, 23, 0, 0, 0, 17,
	0, 27, 0, 26, 0, 0, 22, 24, 23, 0,
	0, 0, 0, 0, 27, 0, 26,
}
var yyPact = [...]int{

	273, -1000, -8, 46, 273, 62, 1, -28, -35, 286,
	238, 286, 273, 273, 273, 273, -1000, 273, -21, 226,
	-1000, -1000, 286, 286, 286, -1000, -16, -1000, -1000, -1000,
	-1000, 273, 273, 273, -1000, -1000, 62, 286, -1000, -1000,
	273, 286, -1000, 6, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -16, 6, -1000, 6,
	27, 27, 27, 27, 226, -18, -6, 273, -1000, 286,
	286, 6, -1000, 23, -1000, -1000, -1000, 189, 27, 27,
	-1000, -1000, 273, 273, 14, 50, 273, -16, 286, 286,
	-1000, 6, -1000, -1000, -1000, 6, -1000, -1000, -1000, 25,
	25, -1000, -1000, -1000, 25, -1000, -1000, 152, 115, 273,
	-1000, -1000, 25, 273, 25,
}
var yyPgo = [...]int{

	0, 90, 45, 4, 7, 93, 107, 106, 23, 6,
	0, 103, 102, 47, 26, 101, 2, 99, 88, 86,
	82, 66, 65, 56,
}
var yyR1 = [...]int{

	0, 17, 17, 1, 1, 4, 4, 5, 5, 6,
	6, 3, 2, 7, 8, 8, 9, 9, 10, 10,
	18, 10, 19, 10, 20, 10, 21, 10, 22, 10,
	23, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 11, 11, 11, 12, 12, 13, 13,
	13, 14, 14, 14, 14, 14, 14, 14, 14, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 16,
	16,
}
var yyR2 = [...]int{

	0, 0, 2, 1, 2, 1, 2, 2, 2, 1,
	2, 3, 3, 3, 0, 2, 2, 1, 0, 2,
	0, 4, 0, 4, 0, 8, 0, 6, 0, 4,
	0, 4, 1, 3, 3, 3, 3, 2, 2, 2,
	2, 3, 2, 1, 2, 2, 1, 3, 1, 1,
	3, 2, 5, 2, 2, 1, 2, 3, 2, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 0,
	2,
}
var yyChk = [...]int{

	-1000, -17, -1, -10, -5, -3, 7, 4, 6, 12,
	-11, 9, -9, -7, 10, 11, 13, 36, -12, 15,
	16, -14, 30, 32, 31, 14, 40, 38, 33, 34,
	35, 27, 28, 17, -1, -8, -9, 15, -2, 8,
	38, 38, -2, -13, -15, -14, 4, 5, 6, 7,
	8, 9, 10, 11, 12, 13, 15, -13, -9, -13,
	-10, -10, -10, -10, -16, -4, -10, -6, -5, 39,
	29, -13, -3, -13, -13, -13, -3, -16, -10, -10,
	-10, -8, -18, -19, -4, -13, -22, -23, 29, -16,
	-3, -13, 37, 33, -4, -13, -13, 18, 26, -10,
	-10, 26, 5, 26, -10, -3, -13, -16, -16, -21,
	26, 26, -10, -20, -10,
}
var yyDef = [...]int{

	-2, -2, 0, 3, 18, 14, 0, 0, 0, 0,
	32, 0, 18, 18, 18, 18, 69, 18, 43, 0,
	17, 46, 0, 0, 0, 55, 0, 69, 2, 7,
	8, 18, 18, 18, 4, 19, 14, 0, 20, 22,
	18, 0, 28, 30, 48, 49, 59, 60, 61, 62,
	63, 64, 65, 66, 67, 68, 0, 44, 45, 69,
	37, 38, 39, 40, 42, 0, 5, 18, 9, 0,
	0, 16, 58, 51, 53, 54, 56, 0, 34, 35,
	36, 15, 18, 18, 0, 0, 18, 0, 0, 33,
	41, 70, 11, 10, 6, 13, 47, 69, 57, 21,
	23, 12, 69, 26, 29, 31, 50, 0, 0, 18,
	52, 24, 27, 18, 25,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	33, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 32, 3, 30, 3, 35, 3,
	38, 26, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 34,
	3, 39, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 29, 3, 40, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 36, 3, 37,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 27, 28, 31,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{
}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-0:yypt+1]
//line syn.y:24
		{ return 1;}
	case 2:
		yyDollar = yyS[yypt-2:yypt+1]
//line syn.y:25
		{return !compile(yyDollar[1].tree);}
	case 4:
		yyDollar = yyS[yypt-2:yypt+1]
//line syn.y:27
		{yyVAL.tree=tree2(';', yyDollar[1].tree, yyDollar[2].tree);}
	case 6:
		yyDollar = yyS[yypt-2:yypt+1]
//line syn.y:29
		{yyVAL.tree=tree2(';', yyDollar[1].tree, yyDollar[2].tree);}
	case 8:
		yyDollar = yyS[yypt-2:yypt+1]
//line syn.y:31
		{yyVAL.tree=tree1('&', yyDollar[1].tree);}
	case 11:
		yyDollar = yyS[yypt-3:yypt+1]
//line syn.y:34
		{yyVAL.tree=tree1(BRACE, yyDollar[2].tree);}
	case 12:
		yyDollar = yyS[yypt-3:yypt+1]
//line syn.y:35
		{yyVAL.tree=tree1(PCMD, yyDollar[2].tree);}
	case 13:
		yyDollar = yyS[yypt-3:yypt+1]
//line syn.y:36
		{yyVAL.tree=tree2('=', yyDollar[1].tree, yyDollar[3].tree);}
	case 14:
		yyDollar = yyS[yypt-0:yypt+1]
//line syn.y:37
		{yyVAL.tree=0;}
	case 15:
		yyDollar = yyS[yypt-2:yypt+1]
//line syn.y:38
		{yyVAL.tree=mung2(yyDollar[1].tree, yyDollar[1].tree->child[0], yyDollar[2].tree);}
	case 16:
		yyDollar = yyS[yypt-2:yypt+1]
//line syn.y:39
		{yyVAL.tree=mung1(yyDollar[1].tree, yyDollar[1].tree->rtype==HERE?heredoc(yyDollar[2].tree):yyDollar[2].tree);}
	case 18:
		yyDollar = yyS[yypt-0:yypt+1]
//line syn.y:41
		{yyVAL.tree=0;}
	case 19:
		yyDollar = yyS[yypt-2:yypt+1]
//line syn.y:42
		{yyVAL.tree=epimung(yyDollar[1].tree, yyDollar[2].tree);}
	case 20:
		yyDollar = yyS[yypt-2:yypt+1]
//line syn.y:43
		{skipnl();}
	case 21:
		yyDollar = yyS[yypt-4:yypt+1]
//line syn.y:44
		{yyVAL.tree=mung2(yyDollar[1].tree, yyDollar[2].tree, yyDollar[4].tree);}
	case 22:
		yyDollar = yyS[yypt-2:yypt+1]
//line syn.y:45
		{skipnl();}
	case 23:
		yyDollar = yyS[yypt-4:yypt+1]
//line syn.y:45
		{yyVAL.tree=mung1(yyDollar[2].tree, yyDollar[4].tree);}
	case 24:
		yyDollar = yyS[yypt-6:yypt+1]
//line syn.y:46
		{skipnl();}
	case 25:
		yyDollar = yyS[yypt-8:yypt+1]
//line syn.y:55
		{yyVAL.tree=mung3(yyDollar[1].tree, yyDollar[3].tree, yyDollar[5].tree ? yyDollar[5].tree : tree1(PAREN, yyDollar[5].tree), yyDollar[8].tree);}
	case 26:
		yyDollar = yyS[yypt-4:yypt+1]
//line syn.y:56
		{skipnl();}
	case 27:
		yyDollar = yyS[yypt-6:yypt+1]
//line syn.y:57
		{yyVAL.tree=mung3(yyDollar[1].tree, yyDollar[3].tree, (struct tree *)0, yyDollar[6].tree);}
	case 28:
		yyDollar = yyS[yypt-2:yypt+1]
//line syn.y:58
		{skipnl();}
	case 29:
		yyDollar = yyS[yypt-4:yypt+1]
//line syn.y:59
		{yyVAL.tree=mung2(yyDollar[1].tree, yyDollar[2].tree, yyDollar[4].tree);}
	case 30:
		yyDollar = yyS[yypt-2:yypt+1]
//line syn.y:60
		{skipnl();}
	case 31:
		yyDollar = yyS[yypt-4:yypt+1]
//line syn.y:61
		{yyVAL.tree=tree2(SWITCH, yyDollar[2].tree, yyDollar[4].tree);}
	case 32:
		yyDollar = yyS[yypt-1:yypt+1]
//line syn.y:62
		{yyVAL.tree=simplemung(yyDollar[1].tree);}
	case 33:
		yyDollar = yyS[yypt-3:yypt+1]
//line syn.y:63
		{yyVAL.tree=mung2(yyDollar[1].tree, yyDollar[2].tree, yyDollar[3].tree);}
	case 34:
		yyDollar = yyS[yypt-3:yypt+1]
//line syn.y:64
		{yyVAL.tree=tree2(ANDAND, yyDollar[1].tree, yyDollar[3].tree);}
	case 35:
		yyDollar = yyS[yypt-3:yypt+1]
//line syn.y:65
		{yyVAL.tree=tree2(OROR, yyDollar[1].tree, yyDollar[3].tree);}
	case 36:
		yyDollar = yyS[yypt-3:yypt+1]
//line syn.y:66
		{yyVAL.tree=mung2(yyDollar[2].tree, yyDollar[1].tree, yyDollar[3].tree);}
	case 37:
		yyDollar = yyS[yypt-2:yypt+1]
//line syn.y:67
		{yyVAL.tree=mung2(yyDollar[1].tree, yyDollar[1].tree->child[0], yyDollar[2].tree);}
	case 38:
		yyDollar = yyS[yypt-2:yypt+1]
//line syn.y:68
		{yyVAL.tree=mung3(yyDollar[1].tree, yyDollar[1].tree->child[0], yyDollar[1].tree->child[1], yyDollar[2].tree);}
	case 39:
		yyDollar = yyS[yypt-2:yypt+1]
//line syn.y:69
		{yyVAL.tree=mung1(yyDollar[1].tree, yyDollar[2].tree);}
	case 40:
		yyDollar = yyS[yypt-2:yypt+1]
//line syn.y:70
		{yyVAL.tree=mung1(yyDollar[1].tree, yyDollar[2].tree);}
	case 41:
		yyDollar = yyS[yypt-3:yypt+1]
//line syn.y:71
		{yyVAL.tree=tree2(FN, yyDollar[2].tree, yyDollar[3].tree);}
	case 42:
		yyDollar = yyS[yypt-2:yypt+1]
//line syn.y:72
		{yyVAL.tree=tree1(FN, yyDollar[2].tree);}
	case 44:
		yyDollar = yyS[yypt-2:yypt+1]
//line syn.y:74
		{yyVAL.tree=tree2(ARGLIST, yyDollar[1].tree, yyDollar[2].tree);}
	case 45:
		yyDollar = yyS[yypt-2:yypt+1]
//line syn.y:75
		{yyVAL.tree=tree2(ARGLIST, yyDollar[1].tree, yyDollar[2].tree);}
	case 47:
		yyDollar = yyS[yypt-3:yypt+1]
//line syn.y:77
		{yyVAL.tree=tree2('^', yyDollar[1].tree, yyDollar[3].tree);}
	case 48:
		yyDollar = yyS[yypt-1:yypt+1]
//line syn.y:78
		{lastword=1; yyDollar[1].tree->type=WORD;}
	case 50:
		yyDollar = yyS[yypt-3:yypt+1]
//line syn.y:80
		{yyVAL.tree=tree2('^', yyDollar[1].tree, yyDollar[3].tree);}
	case 51:
		yyDollar = yyS[yypt-2:yypt+1]
//line syn.y:81
		{yyVAL.tree=tree1('$', yyDollar[2].tree);}
	case 52:
		yyDollar = yyS[yypt-5:yypt+1]
//line syn.y:82
		{yyVAL.tree=tree2(SUB, yyDollar[2].tree, yyDollar[4].tree);}
	case 53:
		yyDollar = yyS[yypt-2:yypt+1]
//line syn.y:83
		{yyVAL.tree=tree1('"', yyDollar[2].tree);}
	case 54:
		yyDollar = yyS[yypt-2:yypt+1]
//line syn.y:84
		{yyVAL.tree=tree1(COUNT, yyDollar[2].tree);}
	case 56:
		yyDollar = yyS[yypt-2:yypt+1]
//line syn.y:86
		{yyVAL.tree=tree1('`', yyDollar[2].tree);}
	case 57:
		yyDollar = yyS[yypt-3:yypt+1]
//line syn.y:87
		{yyVAL.tree=tree1(PAREN, yyDollar[2].tree);}
	case 58:
		yyDollar = yyS[yypt-2:yypt+1]
//line syn.y:88
		{yyVAL.tree=mung1(yyDollar[1].tree, yyDollar[2].tree); yyVAL.tree->type=PIPEFD;}
	case 69:
		yyDollar = yyS[yypt-0:yypt+1]
//line syn.y:90
		{yyVAL.tree=(struct tree*)0;}
	case 70:
		yyDollar = yyS[yypt-2:yypt+1]
//line syn.y:91
		{yyVAL.tree=tree2(WORDS, yyDollar[1].tree, yyDollar[2].tree);}
	}
	goto yystack /* stack new state and value */
}
