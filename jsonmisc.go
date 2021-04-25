package jsonmisc

import (
	unicodemisc "github.com/djeeno/go-unicode-misc"
)

func AppendQuote(dst []byte, s string) []byte {
	if dst == nil {
		dst = make([]byte, 0, len(s))
	}

	for i := 0; i < len(s); i++ {
		// https://tools.ietf.org/html/rfc8259#section-7
		// MUST be escaped: quotation mark, reverse solidus, and the control characters (U+0000 through U+001F).
		if s[i] != '"' && s[i] != '\\' && 0x1F < s[i] {
			dst = append(dst, s[i])
			continue
		}

		switch s[i] {
		case '"', '\\':
			dst = append(dst, '\\', s[i])
		case '\b' /* 0x08 */ :
			dst = append(dst, '\\', 'b')
		case '\f' /* 0x0C */ :
			dst = append(dst, '\\', 'f')
		case '\n' /* 0x0A */ :
			dst = append(dst, '\\', 'n')
		case '\r' /* 0x0D */ :
			dst = append(dst, '\\', 'r')
		case '\t' /* 0x09 */ :
			dst = append(dst, '\\', 't')
		default:
			dst = append(dst, unicodemisc.UnicodeEscapeSequence(s[i])...)
		}
	}

	return dst
}
