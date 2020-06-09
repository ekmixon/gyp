/*
Copyright (c) 2007-2013. The YARA Authors. All Rights Reserved.

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

/* Lexical analyzer for YARA */

%top{
package parser

import (
    "fmt"
    "io"
    "log"
    "os"
    "math"
    "strconv"
    "strings"

    "github.com/VirusTotal/gyp/ast"
    "github.com/VirusTotal/gyp/hex"
    gyperror "github.com/VirusTotal/gyp/error"
)

// YYtype is the structure returned by the lexer every time the scanner asks
// for the next token. If the lexer wants to return an error to the scanner it
// sets the Error field and leaves the Token empty.
type YYtype struct {
  Token int
  Error gyperror.Error
}

func Token(t int) YYtype {
  return YYtype{Token: t}
}

func Error(c gyperror.Code, msg string) YYtype {
  return YYtype{Error: gyperror.Error{c, msg, 0,}}
}

// Note that the parser is generated with the yr prefix to avoid name
// collisions with the hex and regexp parsers.
type YYcontext struct {
  lval *yrSymType
}

// The YY_USER_DATA macro is used to define variables inside the Scanner
// Lex() method. These variables manage buffers for gathering groups of
// tokens. Flex collects tokens individually but strings and conditions may
// contain several tokens.
// Two different buffers are necessary because conditions may contain strings.

// This comment applies to the YY_USER_ACTION macro, which is having
// a problem with comments...
// For condition, the colons and whitespace will be collected in the
// prefix and the right brace for the suffix. Use strings.TrimLeft/Right
}

%{

YY_USER_DATA(
  var (
    collectText       bool    // true when reading string/regex/hex text
    text              []byte  // collects text while reading a string/regex/hex
    YYtext            string  // each action sets this to the current yytext
  )
)

YY_USER_ACTION(
  YYtext = string(yytext)

  if collectText {
    text = append(text, yytext...)
  }
)

// Define a constant for end-of-file
const eof = 0
%}

%option reentrant bison-bridge
%option nounput
%option never-interactive
%option yylineno
%option prefix="yara_yy"

%option verbose
%option warn

%x str
%x regexp
%x include
%x comment

digit         [0-9]
letter        [a-zA-Z]
hexdigit      [a-fA-F0-9]
octdigit      [0-7]

%%

".."                    { return Token(_DOT_DOT_);     }
"<"                     { return Token(_LT_);          }
">"                     { return Token(_GT_);          }
"<="                    { return Token(_LE_);          }
">="                    { return Token(_GE_);          }
"=="                    { return Token(_EQ_);          }
"!="                    { return Token(_NEQ_);         }
"<<"                    { return Token(_SHIFT_LEFT_);  }
">>"                    { return Token(_SHIFT_RIGHT_); }
"private"               { return Token(_PRIVATE_);     }
"global"                { return Token(_GLOBAL_);      }
"rule"                  { return Token(_RULE_);        }
"meta"                  { return Token(_META_);        }
"strings"               { return Token(_STRINGS_);     }
"ascii"                 { return Token(_ASCII_);       }
"base64"                { return Token(_BASE64_);      }
"wide"                  { return Token(_WIDE_);        }
"xor"                   { return Token(_XOR_);         }
"fullword"              { return Token(_FULLWORD_);    }
"nocase"                { return Token(_NOCASE_);      }
"condition"             { return Token(_CONDITION_);   }
"true"                  { return Token(_TRUE_);        }
"false"                 { return Token(_FALSE_);       }
"not"                   { return Token(_NOT_);         }
"and"                   { return Token(_AND_);         }
"or"                    { return Token(_OR_);          }
"at"                    { return Token(_AT_);          }
"in"                    { return Token(_IN_);          }
"of"                    { return Token(_OF_);          }
"them"                  { return Token(_THEM_);        }
"for"                   { return Token(_FOR_);         }
"all"                   { return Token(_ALL_);         }
"any"                   { return Token(_ANY_);         }
"entrypoint"            { return Token(_ENTRYPOINT_);  }
"filesize"              { return Token(_FILESIZE_);    }
"matches"               { return Token(_MATCHES_);     }
"contains"              { return Token(_CONTAINS_);    }
"import"                { return Token(_IMPORT_);      }
"include"               { return Token(_INCLUDE_);     }

"/*"                    { BEGIN (comment);       }
<comment>"*/"           { BEGIN (INITIAL);       }
<comment>(.|\n)         { /* skip comments */   }


"//"[^\n]*              { /* skip single-line comments */ }

<<EOF>> { return Token(eof) }


$({letter}|{digit}|_)*"*"  {
  yy.Context.lval.s = string(yytext)
  return Token(_STRING_IDENTIFIER_WITH_WILDCARD_);
}


$({letter}|{digit}|_)*  {
  yy.Context.lval.s = string(yytext)
  return Token(_STRING_IDENTIFIER_);
}


#({letter}|{digit}|_)*  {
  yy.Context.lval.s = string(yytext)
  return Token(_STRING_COUNT_);
}


@({letter}|{digit}|_)*  {
  yy.Context.lval.s = string(yytext)
  return Token(_STRING_OFFSET_);
}


!({letter}|{digit}|_)*  {
  yy.Context.lval.s = string(yytext)
  return Token(_STRING_LENGTH_);
}


u?int(8|16|32)(be)? {
  yy.Context.lval.s = string(yytext)
  return Token(_INTEGER_FUNCTION_);
}


({letter}|_)({letter}|{digit}|_)*  {
  yy.Context.lval.s = string(yytext)
  return Token(_IDENTIFIER_);
}


{digit}+(MB|KB){0,1}  {
  var err error
  s := strings.TrimRight(YYtext, "MKB")
  yy.Context.lval.i64, err = strconv.ParseInt(s, 10, 64)
  if err != nil {
    return Error(
      gyperror.NumberConversionError,
      fmt.Sprintf("%s", err))
  }

  if strings.HasSuffix(YYtext, "KB") {
      if yy.Context.lval.i64 > math.MaxInt64 / 1024 {
        return Error(
          gyperror.IntegerOverflowError,
          fmt.Sprintf("Found %s; Max: %d", YYtext, math.MaxInt64))
      } else {
        yy.Context.lval.i64 *= 1024
      }
  } else if strings.HasSuffix(YYtext, "MB") {
      if yy.Context.lval.i64 > math.MaxInt64 / 1048576 {
        return Error(
          gyperror.IntegerOverflowError,
          fmt.Sprintf("Found %s; Max: %d", YYtext, math.MaxInt64))
      } else {
          yy.Context.lval.i64 *= 1048576
      }
  }
  return Token(_NUMBER_);
}

{digit}+"."{digit}+  {
  return Token(_DOUBLE_);
}

0x{hexdigit}+  {
  var err error
  yy.Context.lval.i64, err = strconv.ParseInt(YYtext, 0, 64)
  if err != nil {
    return Error(
      gyperror.NumberConversionError,
      fmt.Sprintf("%s", err))
  }

  return Token(_NUMBER_);
}

0o{octdigit}+  {
  var err error
  s := strings.TrimLeft(YYtext, "0o")
  yy.Context.lval.i64, err = strconv.ParseInt(s, 8, 64)
  if err != nil {
    return Error(
      gyperror.NumberConversionError,
      fmt.Sprintf("%s", err))
  }
  return Token(_NUMBER_);
}


<str>\"   {     /* saw closing quote - all done */

  // NOTE: text will end with `"` char
  collectText = false
  yy.Context.lval.s = strings.TrimSuffix(string(text), `"`)

  BEGIN (INITIAL);

  return Token(_TEXT_STRING_);
}


<str>\\t   {
}


<str>\\n   {
}


<str>\\\"   {
}


<str>\\\\   {
}


<str>\\x{hexdigit}{2} {
}


<str>[^\\\n\"]+   { }


<str>\n  {
  return Error(
    gyperror.UnterminatedStringError,
    "unterminate string")
}


<str>\\(.|\n) {
  return Error(
    gyperror.IllegalEscapeSequenceError,
    "illegal escape sequence")
}


<regexp>\/i?s?  {

  collectText = false

  txt := strings.TrimSuffix(string(text), YYtext)

  var mods ast.RegexpModifiers
  for _, c := range YYtext {
      switch c {
      case 'i':
          mods = mods | ast.RegexpCaseInsensitive
      case 's':
          mods = mods | ast.RegexpDotAll
      case '/':
          // Ignore
      default:
          // Should be impossible
          return Error(
            gyperror.InvalidRegexModifierError,
            fmt.Sprintf(`invalid regexp modifier "%c"`, c))
      }
  }

  yy.Context.lval.reg = &ast.LiteralRegexp {
      Value: txt,
      Modifiers: mods,
  }

  BEGIN (INITIAL);
  return Token(_REGEXP_);
}


<regexp>\\\/  {
}


<regexp>\\. {
}


<regexp>[^/\n\\]+ { }


<regexp>\n  {
  return Error(
    gyperror.UnterminatedRegexError,
    "unterminated regexp")
}


\"  {
  collectText = true
  text = []byte{}
  BEGIN (str);
}


"/"  {
  collectText = true
  text = []byte{}
  BEGIN (regexp);
}


\{(({hexdigit}|[ \-|\?\[\]\(\)\n\r\t]|\/\*[^*]*\*\/)+|\/\/.*)+\}  {
  // Match hex-digits with whitespace or comments. The latter are stripped
  // out by hex_lexer.l

  // NOTE: The above comment may not apply. We plan to not use hex_lexer.l

  // No need to collect like str and regexp start conditions
  hexReader := strings.NewReader(YYtext)

  hexTokens, err := hex.Parse(hexReader)
  if err != nil {
    return YYtype{Error: err.(gyperror.Error)}
  }
  yy.Context.lval.hexTokens = hexTokens

  return Token(_HEX_STRING_);
}


[ \t\r\n]   /* skip whitespace */

.   {

  r := int(yytext[0])

  if r >= 32 && r < 127 {
    return Token(r)
  }

  return Error(
    gyperror.NonAsciiByteError,
    fmt.Sprintf(`non-ascii character "\x%02x"`, r))
}

%%