package q

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Quiki lexer.
var Quiki = internal.Register(MustNewLexer(
	&Config{
		Name:      "quiki",
		Aliases:   []string{"qm"},
		Filenames: []string{"*.qm", "*.quiki", ".page"},
		DotAll:    true,
	},
	Rules{

		// rules at root level
		"root": {

			// start a comment
			{`/\*`, CommentMultiline, Push("nested-comment")},

			// start a map{} or map-based block
			{`(map|infobox|infosec|history|style|imagebox|image|model|toc|gallery)(\s*)(\{)`, ByGroups(NameFunction, Text, Punctuation), Push("map")},

			// start a list{} or list-based block
			{`(list|numlist)(\s*)(\{)`, ByGroups(NameFunction, Text, Punctuation), Push("list")},

			// start a generic block
			{`([\w\-\$\.]+)(\s*)(\{)`, ByGroups(NameFunction, Text, Punctuation), Push("block")},
		},

		// rules inside comments
		"nested-comment": {
			{`\*/`, CommentMultiline, Pop(1)}, // decrease comment level
			{`/\*`, CommentMultiline, Push()}, // increase comment level
			{`[^*/]+`, CommentMultiline, nil}, // comment content
			{`[*/]`, CommentMultiline, nil},   // comment content
		},

		// rules inside any block
		"block": {

			/* nested block */
			{`([\w\-\$\.]+)(\s*)(\{)`, ByGroups(NameFunction, Text, Punctuation), Push("block")},

			/* brace escape */
			{`\{`, Punctuation, Push("block")},

			/* exit block OR brace escape */
			{`\}`, Punctuation, Pop(1)},
		},

		// rules inside a map{} or map-based block
		"map": {
			Include("block"),
			{`:`, Punctuation, Push("value")},
		},

		// rules inside a list{} or list-based block
		"list": {
			Include("block"),
			Default(Push("value")), // enter value
		},

		// rules inside a map or list value
		"value": {
			{`(?<!\\);`, Punctuation, Pop(1)},           // exit value with ;
			{`(?<!\\)}`, Punctuation, Pop(2)},           // exit value AND block with }
			{`\[`, Punctuation, Push("formatted-text")}, // enter formatted text with [
			{`.+?`, String, nil},                        // content in the value
		},

		// rules inside a text formatting token
		"formatted-text": {
			{`\[`, Punctuation, Push()}, // increase format level with [
			{`\]`, Punctuation, Pop(1)}, // decrease format level with ]
		},
	},
))
