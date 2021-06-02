// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jsonconv

import (
	"unicode/utf8"
)

var hex = "0123456789abcdef"

// AppendQuote appends the double-quoted JSON string literal representing
// s, to dest and returns the extended buffer
func AppendQuote(dest []byte, s string) []byte {
	return appendQuote(dest, s, false)
}

// AppendQuoteEscapeHTML behaves as AppendQuote but also escapes
// <, >, and &
func AppendQuoteEscapeHTML(dest []byte, s string) []byte {
	return appendQuote(dest, s, true)
}

// appendQuote is a modified version of encoding/json/encodeState.string
func appendQuote(dest []byte, s string, escapeHTML bool) []byte {
	dest = append(dest, '"')
	start := 0
	for i := 0; i < len(s); {
		if b := s[i]; b < utf8.RuneSelf {
			if htmlSafeSet[b] || (!escapeHTML && safeSet[b]) {
				i++
				continue
			}
			if start < i {
				dest = append(dest, s[start:i]...)
			}
			dest = append(dest, '\\')
			switch b {
			case '\\', '"':
				dest = append(dest, b)
			case '\n':
				dest = append(dest, 'n')
			case '\r':
				dest = append(dest, 'r')
			case '\t':
				dest = append(dest, 't')
			default:
				// This encodes bytes < 0x20 except for \t, \n and \r.
				// If escapeHTML is set, it also escapes <, >, and &
				// because they can lead to security holes when
				// user-controlled strings are rendered into JSON
				// and served to some browsers.
				dest = append(dest, []byte(`u00`)...)
				dest = append(dest, hex[b>>4])
				dest = append(dest, hex[b&0xF])
			}
			i++
			start = i
			continue
		}
		c, size := utf8.DecodeRuneInString(s[i:])
		if c == utf8.RuneError && size == 1 {
			if start < i {
				dest = append(dest, s[start:i]...)
			}
			dest = append(dest, []byte(`\ufffd`)...)
			i += size
			start = i
			continue
		}
		// U+2028 is LINE SEPARATOR.
		// U+2029 is PARAGRAPH SEPARATOR.
		// They are both technically valid characters in JSON strings,
		// but don't work in JSONP, which has to be evaluated as JavaScript,
		// and can lead to security holes there. It is valid JSON to
		// escape them, so we do so unconditionally.
		// See http://timelessrepo.com/json-isnt-a-javascript-subset for discussion.
		if c == '\u2028' || c == '\u2029' {
			if start < i {
				dest = append(dest, s[start:i]...)
			}
			dest = append(dest, []byte(`\u202`)...)
			dest = append(dest, hex[c&0xF])
			i += size
			start = i
			continue
		}
		i += size
	}
	if start < len(s) {
		dest = append(dest, s[start:]...)
	}
	dest = append(dest, '"')
	return dest
}

// AppendQuoteBytes appends the double-quoted JSON string literal
// representing s, to dest and returns the extended buffer
func AppendQuoteBytes(dest, s []byte) []byte {
	return appendQuoteBytes(dest, s, false)
}

// AppendQuoteBytesEscapeHTML behaves as AppendQuoteBytes but also
// escapes <, >, and &
func AppendQuoteBytesEscapeHTML(dest, s []byte) []byte {
	return appendQuoteBytes(dest, s, true)
}

// appendQuoteBytes is a modified version of encoding/json/encodeState.stringBytes
func appendQuoteBytes(dest, s []byte, escapeHTML bool) []byte {
	dest = append(dest, '"')
	start := 0
	for i := 0; i < len(s); {
		if b := s[i]; b < utf8.RuneSelf {
			if htmlSafeSet[b] || (!escapeHTML && safeSet[b]) {
				i++
				continue
			}
			if start < i {
				dest = append(dest, s[start:i]...)
			}
			dest = append(dest, '\\')
			switch b {
			case '\\', '"':
				dest = append(dest, b)
			case '\n':
				dest = append(dest, 'n')
			case '\r':
				dest = append(dest, 'r')
			case '\t':
				dest = append(dest, 't')
			default:
				// This encodes bytes < 0x20 except for \t, \n and \r.
				// If escapeHTML is set, it also escapes <, >, and &
				// because they can lead to security holes when
				// user-controlled strings are rendered into JSON
				// and served to some browsers.
				dest = append(dest, []byte(`u00`)...)
				dest = append(dest, hex[b>>4])
				dest = append(dest, hex[b&0xF])
			}
			i++
			start = i
			continue
		}
		c, size := utf8.DecodeRune(s[i:])
		if c == utf8.RuneError && size == 1 {
			if start < i {
				dest = append(dest, s[start:i]...)
			}
			dest = append(dest, []byte(`\ufffd`)...)
			i += size
			start = i
			continue
		}
		// U+2028 is LINE SEPARATOR.
		// U+2029 is PARAGRAPH SEPARATOR.
		// They are both technically valid characters in JSON strings,
		// but don't work in JSONP, which has to be evaluated as JavaScript,
		// and can lead to security holes there. It is valid JSON to
		// escape them, so we do so unconditionally.
		// See http://timelessrepo.com/json-isnt-a-javascript-subset for discussion.
		if c == '\u2028' || c == '\u2029' {
			if start < i {
				dest = append(dest, s[start:i]...)
			}
			dest = append(dest, []byte(`\u202`)...)
			dest = append(dest, hex[c&0xF])
			i += size
			start = i
			continue
		}
		i += size
	}
	if start < len(s) {
		dest = append(dest, s[start:]...)
	}
	dest = append(dest, '"')
	return dest
}

// Quote returns the double-quoted JSON string literal representing s
func Quote(s string) []byte {
	return appendQuote(nil, s, false)
}

// QuoteEscapeHTML behaves as Quote, but also escapes <, >, and &
func QuoteEscapeHTML(s string) []byte {
	return appendQuote(nil, s, true)
}

// QuoteBytes returns the double-quoted JSON string literal
// representing s
func QuoteBytes(s []byte) []byte {
	return appendQuoteBytes(nil, s, false)
}

// QuoteBytesEscapeHTML behaves as QuoteBytes, but also escapes <, >, and &
func QuoteBytesEscapeHTML(s []byte) []byte {
	return appendQuoteBytes(nil, s, true)
}