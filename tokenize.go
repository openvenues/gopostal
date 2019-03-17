package postal

/*
#cgo pkg-config: libpostal
#include <libpostal/libpostal.h>
#include <stdlib.h>
*/
import "C"
import (
	"unsafe"
)

type TokenType uint16

const (
	TokenTypeEnd               TokenType = C.LIBPOSTAL_TOKEN_TYPE_END
	TokenTypeWord              TokenType = C.LIBPOSTAL_TOKEN_TYPE_WORD
	TokenTypeAbbreviation      TokenType = C.LIBPOSTAL_TOKEN_TYPE_ABBREVIATION
	TokenTypeIdeographicChar   TokenType = C.LIBPOSTAL_TOKEN_TYPE_IDEOGRAPHIC_CHAR
	TokenTypeHangulSyllable    TokenType = C.LIBPOSTAL_TOKEN_TYPE_HANGUL_SYLLABLE
	TokenTypeAcronym           TokenType = C.LIBPOSTAL_TOKEN_TYPE_ACRONYM
	TokenTypePhrase            TokenType = C.LIBPOSTAL_TOKEN_TYPE_PHRASE
	TokenTypeEmail             TokenType = C.LIBPOSTAL_TOKEN_TYPE_EMAIL
	TokenTypeURL               TokenType = C.LIBPOSTAL_TOKEN_TYPE_URL
	TokenTypeUsPhone           TokenType = C.LIBPOSTAL_TOKEN_TYPE_US_PHONE
	TokenTypeIntlPhone         TokenType = C.LIBPOSTAL_TOKEN_TYPE_INTL_PHONE
	TokenTypeNumeric           TokenType = C.LIBPOSTAL_TOKEN_TYPE_NUMERIC
	TokenTypeOrdinal           TokenType = C.LIBPOSTAL_TOKEN_TYPE_ORDINAL
	TokenTypeRomanNumeral      TokenType = C.LIBPOSTAL_TOKEN_TYPE_ROMAN_NUMERAL
	TokenTypeIdeographicNumber TokenType = C.LIBPOSTAL_TOKEN_TYPE_IDEOGRAPHIC_NUMBER
	TokenTypePeriod            TokenType = C.LIBPOSTAL_TOKEN_TYPE_PERIOD
	TokenTypeExclamation       TokenType = C.LIBPOSTAL_TOKEN_TYPE_EXCLAMATION
	TokenTypeQuestionMark      TokenType = C.LIBPOSTAL_TOKEN_TYPE_QUESTION_MARK
	TokenTypeComma             TokenType = C.LIBPOSTAL_TOKEN_TYPE_COMMA
	TokenTypeColon             TokenType = C.LIBPOSTAL_TOKEN_TYPE_COLON
	TokenTypeSemicolon         TokenType = C.LIBPOSTAL_TOKEN_TYPE_SEMICOLON
	TokenTypePlus              TokenType = C.LIBPOSTAL_TOKEN_TYPE_PLUS
	TokenTypeAmpersand         TokenType = C.LIBPOSTAL_TOKEN_TYPE_AMPERSAND
	TokenTypeAtSign            TokenType = C.LIBPOSTAL_TOKEN_TYPE_AT_SIGN
	TokenTypePound             TokenType = C.LIBPOSTAL_TOKEN_TYPE_POUND
	TokenTypeEllipsis          TokenType = C.LIBPOSTAL_TOKEN_TYPE_ELLIPSIS
	TokenTypeDash              TokenType = C.LIBPOSTAL_TOKEN_TYPE_DASH
	TokenTypeBreakingDash      TokenType = C.LIBPOSTAL_TOKEN_TYPE_BREAKING_DASH
	TokenTypeHyphen            TokenType = C.LIBPOSTAL_TOKEN_TYPE_HYPHEN
	TokenTypePunctOpen         TokenType = C.LIBPOSTAL_TOKEN_TYPE_PUNCT_OPEN
	TokenTypePunctClose        TokenType = C.LIBPOSTAL_TOKEN_TYPE_PUNCT_CLOSE
	TokenTypeDoubleQuote       TokenType = C.LIBPOSTAL_TOKEN_TYPE_DOUBLE_QUOTE
	TokenTypeSingleQuote       TokenType = C.LIBPOSTAL_TOKEN_TYPE_SINGLE_QUOTE
	TokenTypeOpenQuote         TokenType = C.LIBPOSTAL_TOKEN_TYPE_OPEN_QUOTE
	TokenTypeCloseQuote        TokenType = C.LIBPOSTAL_TOKEN_TYPE_CLOSE_QUOTE
	TokenTypeSlash             TokenType = C.LIBPOSTAL_TOKEN_TYPE_SLASH
	TokenTypeBackslash         TokenType = C.LIBPOSTAL_TOKEN_TYPE_BACKSLASH
	TokenTypeGreaterThan       TokenType = C.LIBPOSTAL_TOKEN_TYPE_GREATER_THAN
	TokenTypeLessThan          TokenType = C.LIBPOSTAL_TOKEN_TYPE_LESS_THAN
	TokenTypeOther             TokenType = C.LIBPOSTAL_TOKEN_TYPE_OTHER
	TokenTypeWhitespace        TokenType = C.LIBPOSTAL_TOKEN_TYPE_WHITESPACE
	TokenTypeNewline           TokenType = C.LIBPOSTAL_TOKEN_TYPE_NEWLINE
	TokenTypeInvalidChar       TokenType = C.LIBPOSTAL_TOKEN_TYPE_INVALID_CHAR
)

type Token struct {
	Offset int
	Len    int
	Type   TokenType
}

func Tokenize(input string, whitespace bool) []Token {
	cInput := C.CString(input)
	defer C.free(unsafe.Pointer(cInput))

	cWhitespace := C.bool(whitespace)
	var cNumTokens = C.size_t(0)

	cTokens := C.libpostal_tokenize(cInput, cWhitespace, &cNumTokens)
	numTokens := uint64(cNumTokens)

	cTokensPtr := (*[1 << 28]C.libpostal_token_t)(unsafe.Pointer(cTokens))[:cNumTokens:cNumTokens]

	var tokens []Token
	var i uint64
	for i = 0; i < numTokens; i++ {
		token := Token{
			Offset: int(cTokensPtr[i].offset),
			Len:    int(cTokensPtr[i].len),
			Type:   TokenType(cTokensPtr[i]._type),
		}
		tokens = append(tokens, token)
	}

	return tokens
}
