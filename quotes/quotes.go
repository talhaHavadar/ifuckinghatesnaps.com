// Package quotes holds the published quotes. Each file is NNNN.txt, where
// NNNN is the quote's number in its URL. Wrap the part to highlight in **.
package quotes

import "embed"

//go:embed *.txt
var FS embed.FS
