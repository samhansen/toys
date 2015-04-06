// Dynamic algorithm using O(NM) time and memory.
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
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Dynamic node.
// PTR - (path to root) True if there exists a back-path allowing this node to
//    be matched.
// ZM - (zero-match) True if the character in S1 can be matched zero times.
// Match - True if the character in S1 matches the character in S0.
type Node struct {
	PTR, ZM, Match bool
}

func contains(s0, s1 string) bool {
	m := make([][]Node, len(s1))
	for i := range m {
		m[i] = make([]Node, len(s0))
	}

	// Initialize first row.
	for j := 0; j < len(s0); j++ {
		if s0[j] == s1[0] {
			m[0][j].PTR = true
			m[0][j].Match = true
		} else if s1[0] == '\032' {
			m[0][j].PTR = true
			m[0][j].ZM = true
			m[0][j].Match = true
		}
	}

	// Fill table.
	if len(s1) > 1 {
		for i := 1; i < len(s1); i++ {
			for j := 0; j < len(s0); j++ {
				if s1[i] == s0[j] {
					m[i][j].PTR = m[i-1][j].PTR && m[i-1][j].ZM
					if j > 0 {
						m[i][j].PTR = m[i][j].PTR || m[i-1][j-1].PTR
					}
					m[i][j].ZM = false
					m[i][j].Match = true
				} else if s1[i] == '\032' {
					m[i][j].PTR = m[i-1][j].PTR && m[i-1][j].ZM
					if j > 0 {
						m[i][j].PTR = m[i][j].PTR || m[i-1][j-1].PTR || m[i][j-1].PTR
					}
					m[i][j].ZM = true
					m[i][j].Match = true
				}
			}
		}
	}

	// Process last row.
	i := len(s1) - 1
	for j := 0; j < len(s0); j++ {
		if m[i][j].PTR {
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
