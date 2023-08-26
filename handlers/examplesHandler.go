package handlers

import (
	"fmt"
	"net/http"
)

func (h *Handler) ExamplesHandler(w http.ResponseWriter, r *http.Request) {
	examples := `Example 1:
/5/plus/3 -> 8 (5+3)
Example 2:
/3/minus/5 -> -2 (3-5)
Example 3:
/3/into/5/plus/8/into/6 -> 63 (3*5+8*6)
Example 4:
/3/into/lbrack/5/plus/8/rbrack/pow/3 -> 6591 (3*(5+8)**3)
Example 5:
/7/lshift/1 -> 14 (7<<1)
Example 6:
/100/rshift/2 -> 25 (100>>2)
Example 7:
/100/modulo/3 -> 1 (100%3)
`
	fmt.Fprintf(w, examples)
}
