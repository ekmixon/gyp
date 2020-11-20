//line hex/hex_lexer.go:2
//line hex/hex_lexer.l:33
package hex

import (
    "fmt"
    "io"
    "log"
    "os"
    "strconv"

    gyperror "github.com/VirusTotal/gyp/error"
)

type YYcontext struct {}

// YYtype is the structure returned by the lexer every time the scanner asks
// for the next token. If the lexer wants to return an error to the scanner it
// sets the Error field and leaves the Token empty.
type YYtype struct {
  Token int
  Value *hexSymType
  Error gyperror.Error
}

func (s *Scanner) Token(t int) YYtype {
  return YYtype{Token: t}
}

func (s *Scanner) TokenInteger(t int, i int) YYtype {
  return YYtype{Token: t, Value: &hexSymType{integer: i}}
}

func (s *Scanner) TokenByte(t int, value, mask byte) YYtype {
  return YYtype{
    Token: t,
    Value: &hexSymType{
        bm:  byteWithMask{ Mask: byte(mask), Value: byte(value) },
    },
  }
}

func Error(c gyperror.Code, msg string) YYtype {
  return YYtype{Error: gyperror.Error{c, msg, 0,}}
}




//line hex/hex_lexer.go:51

// START OF SKELL ------------------------------------------------------
// A lexical scanner generated by flexgo

type Scanner struct {
	In   io.Reader
	Out  io.Writer
	Lineno int

	Filename      string
	Wrap          func(*Scanner) bool
	IsInteractive func(io.Reader) bool
	Context       YYcontext

	lastAcceptingState   int
	lastAcceptingCpos    int
	debug                bool
	start                int
	stateBuf             []int
	statePtr             int
	fullState            int
	fullMatch            int
	fullLp               int
	lp                   int
	lookingForTrailBegin int
	holdChar             byte
	cBufP                int
	didBufferSwitchOnEof bool
	textPtr              int
	nChars               int
	init                 bool
	moreFlag             bool
    moreLen              int

	// buffer
	inputFile    io.Reader
	chBuf        []byte // input buffer
	bufPos       int    // current position in input buffer
	bufSize      int
	bufNChars    int
	Interactive  bool
	atBol        int // 0 (false) or 1 (true)
	fillBuffer   bool
	bufferStatus int
}

func NewScanner() *Scanner {
	yy := Scanner{
		Lineno: 1,
		In:            os.Stdin,
		Out:           os.Stdout,
		Wrap:          func(yyy *Scanner) bool { return true },
		IsInteractive: func(file io.Reader) bool { return yyInteractiveDefault },
		bufSize:       yyBufSize,
		chBuf:         make([]byte, yyBufSize+2),
		start:         1,
		stateBuf:      make([]int, yyBufSize+2),
		atBol:         1,
		debug:         yyFlexDebug,
		fillBuffer:    true,
	}
	return &yy
}

func (yy *Scanner) NewFile() {
	yy.Restart(yy.In)
}

const yyEndOfBufferChar = 0

const yyBufSize = 32768

const (
	eobActEndOfFile    = 0
	eobActContinueScan = 1
	eobActLastMatch    = 2
)

const (
	yyBufferNew        = 0
	yyBufferNormal     = 1
	yyBufferEofPending = 2
)

// [1.0] the user's section 1 definitions and yytext/yyin/yyout/yy_state_type/yylineno etc. def's & init go here
/* Begin user sect3 */
const yyFlexDebug = false

const yyInteractiveDefault = false
// SKEL ----------------------------------------------------------------

// [1.5] DFA------------------------------------------------------------
// SKEL ----------------------------------------------------------------

// [4.0] data tables for the DFA go here -------------------------------
const yyNumRules = 21
const yyEndOfBuffer = 22
var yyAccept = [39]int16{   0,
        0,    0,    0,    0,    0,    0,   22,   20,   16,   16,
       17,   18,   20,    7,   20,    8,    1,   19,    2,   21,
       15,   14,   14,   11,   12,   13,    0,   10,    3,    4,
        5,    6,   12,    0,    0,   10,    9,    0,
    }

var yyEc = [256]byte{    0,
        1,    1,    1,    1,    1,    1,    1,    1,    2,    3,
        1,    1,    2,    1,    1,    1,    1,    1,    1,    1,
        1,    1,    1,    1,    1,    1,    1,    1,    1,    1,
        1,    2,    1,    1,    1,    1,    1,    1,    1,    4,
        5,    6,    1,    1,    7,    1,    8,    9,    9,    9,
        9,    9,    9,    9,    9,    9,    9,    1,    1,    1,
        1,    1,   10,    1,   11,   11,   11,   11,   11,   11,
        1,    1,    1,    1,    1,    1,    1,    1,    1,    1,
        1,    1,    1,    1,    1,    1,    1,    1,    1,    1,
       12,    1,   13,    1,    1,    1,   11,   11,   11,   11,

       11,   11,    1,    1,    1,    1,    1,    1,    1,    1,
        1,    1,    1,    1,    1,    1,    1,    1,    1,    1,
        1,    1,   14,   15,   16,    1,    1,    1,    1,    1,
        1,    1,    1,    1,    1,    1,    1,    1,    1,    1,
        1,    1,    1,    1,    1,    1,    1,    1,    1,    1,
        1,    1,    1,    1,    1,    1,    1,    1,    1,    1,
        1,    1,    1,    1,    1,    1,    1,    1,    1,    1,
        1,    1,    1,    1,    1,    1,    1,    1,    1,    1,
        1,    1,    1,    1,    1,    1,    1,    1,    1,    1,
        1,    1,    1,    1,    1,    1,    1,    1,    1,    1,

        1,    1,    1,    1,    1,    1,    1,    1,    1,    1,
        1,    1,    1,    1,    1,    1,    1,    1,    1,    1,
        1,    1,    1,    1,    1,    1,    1,    1,    1,    1,
        1,    1,    1,    1,    1,    1,    1,    1,    1,    1,
        1,    1,    1,    1,    1,    1,    1,    1,    1,    1,
        1,    1,    1,    1,    1,
    }

var yyMeta = [17]byte{    0,
        1,    1,    2,    1,    1,    1,    1,    1,    3,    3,
        3,    1,    1,    1,    1,    1,
    }

var yyBase = [45]uint16{   0,
        0,    0,    0,    0,   15,   27,   39,   53,   53,   53,
       53,   53,   13,   28,   27,   53,   53,   53,   53,   53,
       53,   53,   53,   53,   26,   53,   27,    0,   53,   53,
       53,   53,   23,   25,   17,    0,   21,   53,   40,   43,
       23,   17,   46,   49,
    }

var yyDef = [45]int16{   0,
       38,    1,   39,   39,   40,   40,   38,   38,   38,   38,
       38,   38,   38,   41,   42,   38,   38,   38,   38,   38,
       38,   38,   38,   38,   38,   38,   43,   44,   38,   38,
       38,   38,   38,   43,   43,   44,   43,    0,   38,   38,
       38,   38,   38,   38,
    }

var yyNxt = [70]uint16{   0,
        8,    9,   10,   11,   12,    8,    8,   13,   14,   15,
       14,   16,    8,   17,   18,   19,   22,   23,   27,   31,
       28,   24,   35,   25,   37,   29,   35,   26,   22,   23,
       35,   33,   35,   24,   33,   25,   32,   30,   38,   26,
       20,   20,   20,   21,   21,   21,   34,   34,   34,   36,
       38,   36,    7,   38,   38,   38,   38,   38,   38,   38,
       38,   38,   38,   38,   38,   38,   38,   38,   38,
    }

var yyChk = [70]int16{   0,
        1,    1,    1,    1,    1,    1,    1,    1,    1,    1,
        1,    1,    1,    1,    1,    1,    5,    5,   13,   42,
       13,    5,   35,    5,   35,   41,   37,    5,    6,    6,
       34,   33,   27,    6,   25,    6,   15,   14,    7,    6,
       39,   39,   39,   40,   40,   40,   43,   43,   43,   44,
        0,   44,   38,   38,   38,   38,   38,   38,   38,   38,
       38,   38,   38,   38,   38,   38,   38,   38,   38,
    }

/* Table of booleans, true if rule could match eol. */
var yyRuleCanMatchEol = [22]int32{   0,
0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 
    0, 0,     };

//line hex/hex_lexer.l:1
/*
Copyright (c) 2013. The YARA Authors. All Rights Reserved.

Redistribution and use in source and binary forms, with or without modification,
are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this
list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice,
this list of conditions and the following disclaimer in the documentation and/or
other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its contributors
may be used to endorse or promote products derived from this software without
specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR
ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON
ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/
/* Lexical analyzer for hex strings */

//line hex/hex_lexer.l:80
// Define a constant for end-of-file
const eof = 0


//line hex/hex_lexer.go:269
// SKEL ----------------------------------------------------------------

const yyInitial  = 0
const comment = 1
const yrange = 2

const yyReadBufSize = 16384

func (yy *Scanner) input(offset, maxRead int) int {

// [5.0] fread()/read() definition of yy_INPUT goes here ---------------
// nothing here, all moved to skeleton
// SKEL ----------------------------------------------------------------

	if yy.Interactive {
		b := make([]byte, 1)
		var n int
		for n = 0; n < maxRead; n++ {
			nn, err := yy.inputFile.Read(b)
			if err != nil && err != io.EOF {
				log.Panicln("Reading 1 byte:", err)
			}
			if nn < 1 {
				break
			}
			yy.chBuf[offset+n] = b[0]
			if b[0] == '\n' {
				n++
				break
			}
		}
		return n
	}

	n, err := yy.inputFile.Read(yy.chBuf[offset:offset+maxRead])
	if err != nil  && err != io.EOF {
		log.Panicf("Reading %d bytes: %v\n", maxRead, err)
	}
	return n
}

/* [6.0] YY_RULE_SETUP definition goes here --------------------------*/

// SKEL ----------------------------------------------------------------

// The main scanner function which does all the work.
func (yy *Scanner) Lex() YYtype {
	var yyCurrentState int
	var yyBp, yyCp int
	var yyAct int
	var yytext []byte
	var yyleng int
	var yylineno int
	_ = yytext
	_ = yyleng
	_ =  yylineno 

	if !yy.init {
		yy.init = true
		
		if yy.In == nil {
			yy.In = os.Stdin
		}
		if yy.Out == nil {
			yy.Out = os.Stdout
		}
		yy.initBuffer(yy.In)
		yy.loadBufferState()
	}

	yyout := yy.Out
	_ = yyout

// [7.0] user's declarations go here -----------------------------------
//line hex/hex_lexer.l:100


//line hex/hex_lexer.go:347
// SKEL ----------------------------------------------------------------

	for { // loops until end-of-file is reached

// [8.0] yy''more-related code goes here -------------------------------
// SKEL ----------------------------------------------------------------

		yyCp = yy.cBufP

		/* Support of yytext. */
		yy.chBuf[yyCp] = yy.holdChar

		// yyBp points to the position in yy_ch_buf of the start of
		// the current run.
		yyBp = yyCp

// [9.0] code to set up and find next match goes here ------------------
		yyCurrentState = yy.start
yyMatch:
		for {
			yyC := int(yyEc[yy.chBuf[yyCp]])
						if yyAccept[yyCurrentState] != 0 {
				yy.lastAcceptingState = yyCurrentState
				yy.lastAcceptingCpos = yyCp
			}
			for int(yyChk[int(yyBase[yyCurrentState])+yyC]) != yyCurrentState {
				yyCurrentState = int(yyDef[yyCurrentState])
				if yyCurrentState >= 39 {
					yyC = int(yyMeta[yyC])
				}
			}
			yyCurrentState = int(yyNxt[int(yyBase[yyCurrentState])+yyC])
			yyCp++
			if yyCurrentState == 38 {
				break
			}
		}
		yyCp = yy.lastAcceptingCpos
		yyCurrentState = yy.lastAcceptingState
// SKEL ----------------------------------------------------------------

	yyFindAction:

// [10.0] code to find the action number goes here ---------------------
		yyAct = int(yyAccept[yyCurrentState])
// SKEL ----------------------------------------------------------------

		yy.textPtr = yyBp

// [2.0] code to fiddle yytext and yyleng for yy''more() goes here -------
	yyleng = yyCp - yyBp
// SKEL ----------------------------------------------------------------

		yy.holdChar = yy.chBuf[yyCp]
		yy.chBuf[yyCp] = 0

// [3.0] code to copy yytext_ptr to yytext[] goes here, if %array ------
// SKEL ----------------------------------------------------------------

		yy.cBufP = yyCp
		yytext = yy.chBuf[yy.textPtr:yyCp]
 
// [11.0] code for yylineno update goes here ---------------------------

		if yyAct != yyEndOfBuffer && yyRuleCanMatchEol[yyAct] != 0 {
			for yyl := 0; yyl < yyleng; yyl++ {
				if yytext[yyl] == '\n' {
					yy.Lineno++
				}
			}
		}

// SKEL ----------------------------------------------------------------

	doAction: // This label is used only to access EOF actions.

// [12.0] debug code goes here -----------------------------------------
// SKEL ----------------------------------------------------------------

		switch yyAct { // beginning of action switch

// [13.0] actions go here ----------------------------------------------
			case 0: // must back up
			// undo the effects of yy_DO_BEFORE_ACTION
			yy.chBuf[yyCp] = yy.holdChar
			yyCp = yy.lastAcceptingCpos
			yyCurrentState = yy.lastAcceptingState
			goto yyFindAction

case 1:

	yylineno = yy.Lineno
	
//line hex/hex_lexer.l:102
{ return yy.Token(_LBRACE_); }
case 2:

	yylineno = yy.Lineno
	

//line hex/hex_lexer.l:103
{ return yy.Token(_RBRACE_); }
case (yyEndOfBuffer + yyInitial  + 1) :
	fallthrough
case (yyEndOfBuffer + comment + 1) :
	fallthrough
case (yyEndOfBuffer + yrange + 1) :
//line hex/hex_lexer.l:105
{ return yy.Token(eof) }
case 3:

	yylineno = yy.Lineno
	

//line hex/hex_lexer.l:107
{
  val, err := strconv.ParseInt(string(yytext), 16, 16)
  if err != nil {
    // This shouldn't happen.
    panic(fmt.Sprintf("error parsing byte: %s\n", err))
  }
  return yy.TokenByte(_BYTE_, byte(val), byte(0xFF));
}
case 4:

	yylineno = yy.Lineno
	

//line hex/hex_lexer.l:116
{
  yytext[1] = '0'  // Replace ? with 0
  val, err := strconv.ParseInt(string(yytext), 16, 16)
  if err != nil {
    // This shouldn't happen.
    panic(fmt.Sprintf("error parsing byte: %s\n", err))
  }
  return yy.TokenByte(_MASKED_BYTE_, byte(val), byte(0xF0));
}
case 5:

	yylineno = yy.Lineno
	

//line hex/hex_lexer.l:126
{
  yytext[0] = '0'
  val, err := strconv.ParseInt(string(yytext), 16, 16)
  if err != nil {
    // This shouldn't happen.
    panic(fmt.Sprintf("error parsing byte: %s\n", err))
  }
  return yy.TokenByte(_MASKED_BYTE_, byte(val), byte(0x0F));
}
case 6:

	yylineno = yy.Lineno
	

//line hex/hex_lexer.l:136
{
  return yy.TokenByte(_MASKED_BYTE_, byte(0x00), byte(0x00));
}
case 7:

	yylineno = yy.Lineno
	

//line hex/hex_lexer.l:140
{
  return Error(
    gyperror.UnevenNumberOfDigitsError,
    fmt.Sprintf(`uneven number of digits in hex string`))
}
case 8:

	yylineno = yy.Lineno
	

//line hex/hex_lexer.l:146
{
  yy.start = 1 + 2*  (yrange);
  return yy.Token(_LBRACKET_);
}
case 9:
/* rule 9 can match eol */

	yylineno = yy.Lineno
	

//line hex/hex_lexer.l:151
// skip comments
case 10:

	yylineno = yy.Lineno
	

//line hex/hex_lexer.l:153
// skip single-line comments
case 11:

	yylineno = yy.Lineno
	

//line hex/hex_lexer.l:155
{
  return yy.Token(_HYPHEN_);
}
case 12:

	yylineno = yy.Lineno
	

//line hex/hex_lexer.l:159
{
  val, err := strconv.ParseInt(string(yytext), 10, 32)
  if err != nil {
    // This shouldn't happen.
    panic(fmt.Sprintf("error parsing jump limit: %s\n", err))
  }
  return yy.TokenInteger(_NUMBER_, int(val));
}
case 13:

	yylineno = yy.Lineno
	

//line hex/hex_lexer.l:168
{
  yy.start = 1 + 2*  (yyInitial );
  return yy.Token(_RBRACKET_);
}
case 14:
/* rule 14 can match eol */

	yylineno = yy.Lineno
	

//line hex/hex_lexer.l:173
// skip whitespaces
case 15:

	yylineno = yy.Lineno
	

//line hex/hex_lexer.l:175
{
  return Error(
    gyperror.InvalidCharInHexStringError,
    fmt.Sprintf(`invalid character in hex string range: %c (0x%02x)`, yytext[0], yytext[0]))
}
case 16:
/* rule 16 can match eol */

	yylineno = yy.Lineno
	

//line hex/hex_lexer.l:181
// skip whitespaces
case 17:

	yylineno = yy.Lineno
	

//line hex/hex_lexer.l:183
{
  return yy.Token(_LPARENS_)
}
case 18:

	yylineno = yy.Lineno
	

//line hex/hex_lexer.l:187
{
  return yy.Token(_RPARENS_)
}
case 19:

	yylineno = yy.Lineno
	

//line hex/hex_lexer.l:191
{
  return yy.Token(_PIPE_)
}
case 20:

	yylineno = yy.Lineno
	

//line hex/hex_lexer.l:195
{               // reject all other characters
  return Error(
    gyperror.InvalidCharInHexStringError,
    fmt.Sprintf(`invalid character in hex string:  %c (0x%02x)`, yytext[0], yytext[0]))
}
case 21:

	yylineno = yy.Lineno
	

//line hex/hex_lexer.l:201
yyout.Write(yytext) 
//line hex/hex_lexer.go:651
// SKEL ----------------------------------------------------------------

		case yyEndOfBuffer:
			/* Amount of text matched not including the EOB char. */
			yyAmountOfMatchedText := yyCp - yy.textPtr - 1

			/* Undo the effects of yy_DO_BEFORE_ACTION. */
			yy.chBuf[yyCp] = yy.holdChar
			 
			if yy.bufferStatus == yyBufferNew {
				/* We're scanning a new file or input source.  It's
				 * possible that this happened because the user
				 * just pointed yyin at a new source and called
				 * yylex().  If so, then we have to assure
				 * consistency between yy_CURRENT_BUFFER and our
				 * globals.  Here is the right place to do so, because
				 * this is the first action (other than possibly a
				 * back-up) that will match for the new input source.
				 */
				yy.nChars = yy.bufNChars
				yy.inputFile = yy.In
				yy.bufferStatus = yyBufferNormal
			}

			/* Note that here we test for yy_c_buf_p "<=" to the position
			 * of the first EOB in the buffer, since yy_c_buf_p will
			 * already have been incremented past the NUL character
			 * (since all states make transitions on EOB to the
			 * end-of-buffer state).  Contrast this with the test
			 * in input().
			 */
			if yy.cBufP <= yy.nChars {
				/* This was really a NUL. */
				var yyNextState int

				yy.cBufP = yy.textPtr + yyAmountOfMatchedText

				yyCurrentState = yy.getPreviousState()

				/* Okay, we're now positioned to make the NUL
				 * transition.  We couldn't have
				 * yy_get_previous_state() go ahead and do it
				 * for us because it doesn't know how to deal
				 * with the possibility of jamming (and we don't
				 * want to build jamming into it because then it
				 * will run more slowly).
				 */

				yyNextState = yy.tryNulTrans(yyCurrentState)

				yyBp = yy.textPtr + 0 

				if yyNextState != 0 {
					/* Consume the NUL. */
					yy.cBufP++
					yyCp = yy.cBufP
					yyCurrentState = yyNextState
					goto yyMatch
				} else {

// [14.0] code to do back-up for compressed tables and set up yy_cp goes here
				yyCp = yy.lastAcceptingCpos
				yyCurrentState = yy.lastAcceptingState
// SKEL ----------------------------------------------------------------

					goto yyFindAction
				}

			} else {

				switch yy.getNextBuffer() {
				case eobActEndOfFile:
					yy.didBufferSwitchOnEof = false

					if yy.Wrap(yy) {
						// Note: because we've taken care in
						// yy_get_next_buffer() to have set up
						// yytext, we can now set up
						// yy.cBufP so that if some total
						// hoser (like flex itself) wants to
						// call the scanner after we return the
						// yy_NULL, it'll still work - another
						// yy_NULL will get returned.
						yy.cBufP = yy.textPtr + 0 

						yyAct = (yyEndOfBuffer + ((yy.start - 1) / 2)  + 1) 
						goto doAction
					} else {
						if !yy.didBufferSwitchOnEof {
							yy.NewFile()
						}
					}
				case eobActContinueScan:
					yy.cBufP = yy.textPtr + yyAmountOfMatchedText

					yyCurrentState = yy.getPreviousState()

					yyCp = yy.cBufP
					yyBp = yy.textPtr + 0 
					goto yyMatch
				case eobActLastMatch:
					yy.cBufP = yy.nChars

					yyCurrentState = yy.getPreviousState()

					yyCp = yy.cBufP
					yyBp = yy.textPtr + 0 
					goto yyFindAction
				}
			}

		default:
			log.Panicln("fatal flex scanner internal error--no action found:", yyAct)
		} // end of action switch
	} // end of scanning one token
	var yyvalue YYtype
	return yyvalue
} // end of yylex

/* yy_get_next_buffer - try to read in a new buffer
 *
 * Returns a code representing an action:
 *	EOB_ACT_LAST_MATCH -
 *	EOB_ACT_CONTINUE_SCAN - continue scanning from current position
 *	EOB_ACT_END_OF_FILE - end of file
 */
func (yy *Scanner) getNextBuffer() int {

	var numberToMove int
	var retval int

	if yy.cBufP > yy.nChars+1 {
		log.Panic("fatal flex scanner internal error--end of buffer missed")
	}

	if !yy.fillBuffer {
		// Don't try to fill the buffer, so this is an EOF.
		if yy.cBufP-yy.textPtr-0  == 1 {
			// We matched a single character, the EOB, so
			// treat this as a final EOF.
			return eobActEndOfFile
		} else {
			// We matched some text prior to the EOB, first
			// process it.
			return eobActLastMatch
		}
	}

	// Try to read more data.

	// First move last chars to start of buffer.
	numberToMove = yy.cBufP - yy.textPtr - 1

	copy(yy.chBuf, yy.chBuf[yy.textPtr:yy.textPtr+numberToMove])

	if yy.bufferStatus == yyBufferEofPending {
		// don't do the read, it's not guaranteed to return an EOF,
		// just force an EOF
		yy.nChars = 0
		yy.bufNChars = 0
	} else {
		numToRead := yy.bufSize - numberToMove - 1

		for numToRead <= 0 {
			// Not enough room in the buffer - grow it.

			yyCBufPOffset := yy.cBufP

			new_size := yy.bufSize * 2

			if new_size <= 0 {
				yy.bufSize += yy.bufSize / 8
			} else {
				yy.bufSize *= 2
			}

			// Include room in for 2 EOB chars.
			bb := make([]byte, yy.bufSize+2-len(yy.chBuf))
			yy.chBuf = append(yy.chBuf, bb...)

			yy.cBufP = yyCBufPOffset

			numToRead = yy.bufSize - numberToMove - 1

		}

		if numToRead > yyReadBufSize {
			numToRead = yyReadBufSize
		}

		// Read in more data.
		yy.nChars = yy.input(numberToMove, numToRead)
		yy.bufNChars = yy.nChars
	}

	if yy.nChars == 0 {
		if numberToMove == 0  {
			retval = eobActEndOfFile
			yy.Restart(yy.In)
		} else {
			retval = eobActLastMatch
			yy.bufferStatus = yyBufferEofPending
		}
	} else {
		retval = eobActContinueScan
	}

	if yy.nChars+numberToMove > yy.bufSize {
		// Extend the array by 50%, plus the number we really need. *
		newSize := yy.nChars + numberToMove + (yy.nChars >> 1)
		if leng := len(yy.chBuf); leng < newSize {
			chBuf := make([]byte, newSize-leng)
			yy.chBuf = append(yy.chBuf, chBuf...)
		}
	}

	yy.nChars += numberToMove
	//yy.bufNChars += numberToMove // TODO: missing in C skel, bug?
	yy.chBuf[yy.nChars] = yyEndOfBufferChar
	yy.chBuf[yy.nChars+1] = yyEndOfBufferChar

	yy.textPtr = 0

	return retval
}

/* yy_get_previous_state - get the state just before the EOB char was reached */
func (yy *Scanner) getPreviousState() int {

	var yyCurrentState int
	var yyCp int

// [15.0] code to get the start state into yy_current_state goes here --
	yyCurrentState = yy.start
// SKEL ----------------------------------------------------------------

	for yyCp = yy.textPtr + 0 ; yyCp < yy.cBufP; yyCp++ {

// [16.0] code to find the next state goes here ------------------------
		yyC := yyIfElse(yy.chBuf[yyCp] != 0, int(yyEc[yy.chBuf[yyCp]]), 1)
				if yyAccept[yyCurrentState] != 0 {
			yy.lastAcceptingState = yyCurrentState
			yy.lastAcceptingCpos = yyCp
		}
		for int(yyChk[int(yyBase[yyCurrentState])+yyC]) != yyCurrentState {
			yyCurrentState = int(yyDef[yyCurrentState])
			if yyCurrentState >= 39 {
				yyC = int(yyMeta[yyC])
			}
		}
		yyCurrentState = int(yyNxt[int(yyBase[yyCurrentState])+yyC])
// SKEL ----------------------------------------------------------------

	}
	return yyCurrentState
}

/* yy_try_NUL_trans - try to make a transition on the NUL character
 *
 * synopsis
 *      next_state = yy_try_NUL_trans( current_state );
 */
func (yy *Scanner) tryNulTrans(yyCurrentState int) int {

	var yyIsJam bool
	var yyCp int
	_ = yyCp

// [17.0] code to find the next state, and perhaps do backing up, goes here
	yyCp = yy.cBufP

	yyC := 1
		if yyAccept[yyCurrentState] != 0 {
		yy.lastAcceptingState = yyCurrentState
		yy.lastAcceptingCpos = yyCp
	}
	for int(yyChk[int(yyBase[yyCurrentState])+yyC]) != yyCurrentState {
		yyCurrentState = int(yyDef[yyCurrentState])
		if yyCurrentState >= 39 {
			yyC = int(yyMeta[yyC])
		}
	}
	yyCurrentState = int(yyNxt[int(yyBase[yyCurrentState])+yyC])
	if yyCurrentState == 38 {
		yyIsJam = true
	}
// SKEL ----------------------------------------------------------------

	if yyIsJam {
		return 0
	}
	return yyCurrentState
}

func (yy *Scanner) Input() (byte, error) {

	yy.chBuf[yy.cBufP] = yy.holdChar

	if yy.chBuf[yy.cBufP] == yyEndOfBufferChar {
		// yy_c_buf_p now points to the character we want to return.
		// If this occurs *before* the EOB characters, then it's a
		// valid NUL; if not, then we've hit the end of the buffer.
		if yy.cBufP < yy.nChars {
			// This was really a NUL.
			yy.chBuf[yy.cBufP] = 0
		} else {
			// need more input
			offset := yy.cBufP - yy.textPtr
			yy.cBufP++

			switch yy.getNextBuffer() {
			case eobActLastMatch:
					/* This happens because yy_g_n_b()
					 * sees that we've accumulated a
					 * token and flags that we need to
					 * try matching the token before
					 * proceeding.  But for input(),
					 * there's no matching to consider.
					 * So convert the EOB_ACT_LAST_MATCH
					 * to EOB_ACT_END_OF_FILE.
					 */

					/* Reset buffer status. */
				yy.Restart(yy.In)

				fallthrough

			case eobActEndOfFile:
				if yy.Wrap(yy) {
					return 0, io.EOF
				}

				if !yy.didBufferSwitchOnEof {
					yy.Restart(yy.In)
				}

				return yy.Input()

			case eobActContinueScan:
				yy.cBufP = yy.textPtr + offset
			}
		}
	}

	c := yy.chBuf[yy.cBufP]
	yy.chBuf[yy.cBufP] = 0	// preserve yytext
	yy.cBufP++
	yy.holdChar = yy.chBuf[yy.cBufP]

// [19.0] update BOL and yylineno --------------------------------------
	if c == '\n' {
		yy.Lineno++
	}
// SKEL ----------------------------------------------------------------

return c, nil
}

/** Immediately switch to a different input stream.
 * @param input_file A readable stream.
 *
 * @note This function does not reset the start condition to @c yyInitial  .
 */
func (yy *Scanner) Restart(input_file io.Reader) {
	yy.initBuffer(input_file)
	yy.loadBufferState()
}

func (yy *Scanner) loadBufferState() {
	yy.nChars = yy.bufNChars
	yy.cBufP = yy.bufPos
	yy.textPtr = yy.cBufP
	yy.In = yy.inputFile
	yy.holdChar = yy.chBuf[yy.cBufP]
}

/* Initializes or reinitializes a buffer.
 * This function is sometimes called more than once on the same buffer,
 * such as during a yyrestart() or at EOF.
 */
func (yy *Scanner) initBuffer(file io.Reader) {

	yy.flushBuffer()

	yy.inputFile = file

	yy.fillBuffer = true

	yy.Interactive = yy.IsInteractive(file)

}

/** Discard all buffered characters. On the next scan, YY_INPUT will be called.
 * @param b the buffer state to be flushed, usually @c YY_CURRENT_BUFFER.
 *
 */
func (yy *Scanner) flushBuffer() {

	yy.bufNChars = 0

	/* We always need two end-of-buffer characters.  The first causes
	 * a transition to the end-of-buffer state.  The second causes
	 * a jam in that state.
	 */
	yy.chBuf[0] = yyEndOfBufferChar
	yy.chBuf[1] = yyEndOfBufferChar

	yy.bufPos = 0

	yy.atBol = 1
	yy.bufferStatus = yyBufferNew

	yy.loadBufferState()
}

func yyIfElse(b bool, i1, i2 int) int {
	if b {
		return i1
	}
	return i2
}

func YYmain(filenames ...string) (interface{}, error) {

	var errval error

	yy := NewScanner()

	yy.Filename = "<stdin>"

	if len(filenames) > 0 {
		yy.Filename = filenames[0]
		yy.In, errval = os.Open(yy.Filename)
		if errval != nil {
			return nil, errval
		}
		yy.Wrap = func(yyy *Scanner) bool {
			if len(filenames) == 0 {
				// should not happen
				return true
			}
			yyy.In.(*os.File).Close()
			filenames = filenames[1:]
			if len(filenames) == 0 {
				return true
			}
			yyy.Filename = filenames[0]
			yyy.In, errval = os.Open(yyy.Filename)
			if errval != nil {
				return true
			}
			return false
		}
	}

	return yy.Lex(), errval

}

// END OF SKELL --------------------------------------------------------
//line hex/hex_lexer.l:201



