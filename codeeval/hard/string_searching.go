// Dynamic algorithm using O(NM) time and linear memory.
//
// Dynamic algorithm effectively uses 3 dynamic tables:
// Table 1:
//    O(MN) sized table of boolean values which indicate whether S1[i] == S0[j].
// Tabl 2:
//    O(MN) sized table of boolean values indicating whether or not this node in
//    the table can backtrack to a matching value for S1[0].  Ie. this indicates
//    whether a matching solution is known for the i,j-th entry in the table.
//    Rather than performing a final table-backtrack, we carry forward the
//    knowledge of whether the backtrack would have been successful or not.
// Table 3:
//    O(MN) sized table of boolean values indicating whether this entry in the
//    table is allowed to match 0 characters.  Ie. if one were to perform a back
//    traversal, this would allow a traversal in the "up" direction.
//
// Note, this implementation compresses the 3 tables into a table of struct
// type nodes.
//
// Rather than performaing a table-backtrack, we carry forward the knowledge of
// whether we could have successfully backtracked.
//
// Since no backtracking is actually required, we can process the table 2-rows
// at a time allowing us to use linear memory.  Ie. we don't need any more than
// the current and previous row to compute the Node values for any given row.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Dynamic node.
// PTR - True if there exists a path from this node to the root.
// ZM - True if the character in S0 can be matched zero times.
// Match - True if the character in S1 matches the character in S0.
type Node struct {
	PTR, ZM, Match bool
}

func contains(s0, s1 string) bool {
	row1 := make([]Node, len(s0))
	row2 := make([]Node, len(s0))
	// Used to swap the pointers.
	tmp := make([]Node, len(s0))

	// Initialize first row.
	for j := 0; j < len(s0); j++ {
		if s0[j] == s1[0] {
			row1[j].PTR = true
			row1[j].Match = true
		} else if s1[0] == '\032' {
			row1[j].PTR = true
			row1[j].ZM = true
			row1[j].Match = true
		}
	}

	// Compute the Node values for the next row.
	if len(s1) > 1 {
		for i := 1; i < len(s1); i++ {
			for j := 0; j < len(s0); j++ {
				if s1[i] == s0[j] {
					row2[j].PTR = row1[j].PTR && row1[j].ZM
					if j > 0 {
						row2[j].PTR = row2[j].PTR || row1[j-1].PTR
					}
					row2[j].ZM = false
					row2[j].Match = true
				} else if s1[i] == '\032' {
					row2[j].PTR = row1[j].PTR && row1[j].ZM
					if j > 0 {
						row2[j].PTR = row2[j].PTR || row1[j-1].PTR || row2[j-1].PTR
					}
					row2[j].ZM = true
					row2[j].Match = true
				}
			}
			// Swap the pointers and reset row2.
			tmp = row1
			row1 = row2
			row2 = tmp
			row2 = make([]Node, len(s0))
		}
	}

	// Process last row (saved in row1).
	for j := 0; j < len(s0); j++ {
		if row1[j].PTR {
			return true
		}
	}
	return false
}

func main() {
	fd, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(fd)

	for {
		line, _, _ := reader.ReadLine()
		if line == nil {
			break
		}
		parts := strings.Split(string(line), ",")
		// Rather than write a real lexer, we'll cheat a little and replace * with
		// the unprintable SUB (or substitution) character while temporarily saving
		// \* as the STX/ETX (start-of-text/end-of-text).
		parts[1] = strings.Replace(parts[1], "\\*", "\002\003", -1)
		parts[1] = strings.Replace(parts[1], "*", "\032", -1)
		parts[1] = strings.Replace(parts[1], "\002\003", "*", -1)
		fmt.Println(contains(parts[0], parts[1]))
	}
}
