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
		"root": {},
	},
))
