package jsonconv

import (
	"bytes"
	"testing"
)

// Considering the original code comes from the go stdlib, these tests
// seem unnecessary. But light alterations have been made, so why not
// make sure?

type quoteTest struct {
	desc string
	input []byte
	expected []byte
	escapeHTML bool
}

var quoteTests = []quoteTest{
	{
		desc: "empty string",
		input: []byte(``),
		expected: []byte(`""`),
	},
	{
		desc: "ascii",
		input: []byte(`foo bar`),
		expected: []byte(`"foo bar"`),
	},
	{
		desc: "non ascii",
		input: []byte(`hello 😎`),
		expected: []byte(`"hello 😎"`),
	},
	{
		desc: "requires escaping #1",
		input: []byte(`hello\"there"`),
		expected: []byte(`"hello\\\"there\""`),
	},
	{
		desc: "requires escaping #2",
		input: []byte{'*', 0, '\u0009', '\u000d', '\u000a', '\f', '\f', '\f', '\f'},
		expected: []byte(`"*\u0000\t\r\n\u000c\u000c\u000c\u000c"`),
	},
	{
		desc: "requires escaping #3",
		input: []byte(`{"JSON":"ro cks!"}`),
		expected: []byte(`"{\"JSON\":\"ro\u2028cks!\"}"`),
	},
	{
		desc: "escape html",
		input: []byte(`<a href="https://www.example.com/?foo=1&bar=2">example</a>`),
		expected: []byte(`"\u003ca href=\"https://www.example.com/?foo=1\u0026bar=2\"\u003eexample\u003c/a\u003e"`),
		escapeHTML: true,
	},
	{
		desc: "invalid utf8",
		input: []byte{'h', 'e', 'l', 'l', 'o', '\xc0', '\xFF', '!'},
		expected: []byte(`"hello\ufffd\ufffd!"`),
	},
}

func TestQuote(t *testing.T) {
	for i, test := range quoteTests {
		var output []byte
		if test.escapeHTML {
			output = QuoteEscapeHTML(string(test.input))
		} else {
			output = Quote(string(test.input))
		}
		if !bytes.Equal(test.expected, output) {
			if test.desc != "" {
				t.Errorf("test \"%s\" failed: unexpected %s", test.desc, string(output))
			} else {
				t.Errorf("test idx %d failed: unexpected %s", i, string(output))
			}
		}
	}
}

func TestQuoteBytes(t *testing.T) {
	for i, test := range quoteTests {
		var output []byte
		if test.escapeHTML {
			output = QuoteBytesEscapeHTML(test.input)
		} else {
			output = QuoteBytes(test.input)
		}
		if !bytes.Equal(test.expected, output) {
			if test.desc != "" {
				t.Errorf("test \"%s\" failed: unexpected %s", test.desc, string(output))
			} else {
				t.Errorf("test idx %d failed: unexpected %s", i, string(output))
			}
		}
	}
}
