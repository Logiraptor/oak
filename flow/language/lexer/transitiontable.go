
// generated by gocc; DO NOT EDIT.

package lexer



/*
Let s be the current state
Let r be the current input rune
transitionTable[s](r) returns the next state.
*/
type TransitionTable [NumStates] func(rune) int

var TransTab = TransitionTable{
	
		// S0
		func(r rune) int {
			switch {
			case r == 9 : // ['\t','\t']
				return 1
			case r == 10 : // ['\n','\n']
				return 1
			case r == 13 : // ['\r','\r']
				return 1
			case r == 32 : // [' ',' ']
				return 1
			case r == 34 : // ['"','"']
				return 2
			case r == 40 : // ['(','(']
				return 3
			case r == 41 : // [')',')']
				return 4
			case r == 44 : // [',',',']
				return 5
			case r == 46 : // ['.','.']
				return 6
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case r == 59 : // [';',';']
				return 8
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case 97 <= r && r <= 98 : // ['a','b']
				return 9
			case r == 99 : // ['c','c']
				return 10
			case 100 <= r && r <= 101 : // ['d','e']
				return 9
			case r == 102 : // ['f','f']
				return 11
			case 103 <= r && r <= 114 : // ['g','r']
				return 9
			case r == 115 : // ['s','s']
				return 12
			case r == 116 : // ['t','t']
				return 13
			case 117 <= r && r <= 122 : // ['u','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S1
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S2
		func(r rune) int {
			switch {
			case r == 34 : // ['"','"']
				return 14
			
			
			default:
				return 2
			}
			
		},
	
		// S3
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S4
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S5
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S6
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S7
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			
			
			
			}
			return NoState
			
		},
	
		// S8
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S9
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 15
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case 97 <= r && r <= 122 : // ['a','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S10
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 15
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case 97 <= r && r <= 110 : // ['a','n']
				return 16
			case r == 111 : // ['o','o']
				return 17
			case 112 <= r && r <= 122 : // ['p','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S11
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 15
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case r == 97 : // ['a','a']
				return 18
			case 98 <= r && r <= 113 : // ['b','q']
				return 16
			case r == 114 : // ['r','r']
				return 19
			case 115 <= r && r <= 122 : // ['s','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S12
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 15
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case 97 <= r && r <= 100 : // ['a','d']
				return 16
			case r == 101 : // ['e','e']
				return 20
			case 102 <= r && r <= 122 : // ['f','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S13
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 15
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case 97 <= r && r <= 110 : // ['a','n']
				return 16
			case r == 111 : // ['o','o']
				return 21
			case 112 <= r && r <= 113 : // ['p','q']
				return 16
			case r == 114 : // ['r','r']
				return 22
			case 115 <= r && r <= 122 : // ['s','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S14
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S15
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 15
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case 97 <= r && r <= 122 : // ['a','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S16
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 15
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case 97 <= r && r <= 122 : // ['a','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S17
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 15
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case 97 <= r && r <= 108 : // ['a','l']
				return 16
			case r == 109 : // ['m','m']
				return 23
			case r == 110 : // ['n','n']
				return 24
			case 111 <= r && r <= 122 : // ['o','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S18
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 15
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case 97 <= r && r <= 107 : // ['a','k']
				return 16
			case r == 108 : // ['l','l']
				return 25
			case 109 <= r && r <= 122 : // ['m','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S19
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 15
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case 97 <= r && r <= 110 : // ['a','n']
				return 16
			case r == 111 : // ['o','o']
				return 26
			case 112 <= r && r <= 122 : // ['p','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S20
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 15
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case 97 <= r && r <= 109 : // ['a','m']
				return 16
			case r == 110 : // ['n','n']
				return 27
			case 111 <= r && r <= 122 : // ['o','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S21
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 15
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case 97 <= r && r <= 122 : // ['a','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S22
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 15
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case 97 <= r && r <= 116 : // ['a','t']
				return 16
			case r == 117 : // ['u','u']
				return 28
			case 118 <= r && r <= 122 : // ['v','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S23
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 15
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case 97 <= r && r <= 111 : // ['a','o']
				return 16
			case r == 112 : // ['p','p']
				return 29
			case 113 <= r && r <= 122 : // ['q','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S24
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 15
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case 97 <= r && r <= 109 : // ['a','m']
				return 16
			case r == 110 : // ['n','n']
				return 30
			case 111 <= r && r <= 122 : // ['o','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S25
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 15
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case 97 <= r && r <= 114 : // ['a','r']
				return 16
			case r == 115 : // ['s','s']
				return 31
			case 116 <= r && r <= 122 : // ['t','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S26
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 15
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case 97 <= r && r <= 109 : // ['a','m']
				return 16
			case r == 110 : // ['n','n']
				return 32
			case 111 <= r && r <= 122 : // ['o','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S27
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 15
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case 97 <= r && r <= 99 : // ['a','c']
				return 16
			case r == 100 : // ['d','d']
				return 33
			case 101 <= r && r <= 122 : // ['e','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S28
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 15
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case 97 <= r && r <= 100 : // ['a','d']
				return 16
			case r == 101 : // ['e','e']
				return 34
			case 102 <= r && r <= 122 : // ['f','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S29
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 15
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case 97 <= r && r <= 110 : // ['a','n']
				return 16
			case r == 111 : // ['o','o']
				return 35
			case 112 <= r && r <= 122 : // ['p','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S30
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 15
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case 97 <= r && r <= 100 : // ['a','d']
				return 16
			case r == 101 : // ['e','e']
				return 36
			case 102 <= r && r <= 122 : // ['f','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S31
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 15
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case 97 <= r && r <= 100 : // ['a','d']
				return 16
			case r == 101 : // ['e','e']
				return 37
			case 102 <= r && r <= 122 : // ['f','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S32
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 15
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case 97 <= r && r <= 115 : // ['a','s']
				return 16
			case r == 116 : // ['t','t']
				return 38
			case 117 <= r && r <= 122 : // ['u','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S33
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 15
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case 97 <= r && r <= 122 : // ['a','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S34
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 15
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case 97 <= r && r <= 122 : // ['a','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S35
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 15
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case 97 <= r && r <= 109 : // ['a','m']
				return 16
			case r == 110 : // ['n','n']
				return 39
			case 111 <= r && r <= 122 : // ['o','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S36
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 15
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case 97 <= r && r <= 98 : // ['a','b']
				return 16
			case r == 99 : // ['c','c']
				return 40
			case 100 <= r && r <= 122 : // ['d','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S37
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 15
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case 97 <= r && r <= 122 : // ['a','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S38
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 15
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case 97 <= r && r <= 100 : // ['a','d']
				return 16
			case r == 101 : // ['e','e']
				return 41
			case 102 <= r && r <= 122 : // ['f','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S39
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 15
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case 97 <= r && r <= 100 : // ['a','d']
				return 16
			case r == 101 : // ['e','e']
				return 42
			case 102 <= r && r <= 122 : // ['f','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S40
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 15
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case 97 <= r && r <= 115 : // ['a','s']
				return 16
			case r == 116 : // ['t','t']
				return 43
			case 117 <= r && r <= 122 : // ['u','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S41
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 15
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case 97 <= r && r <= 109 : // ['a','m']
				return 16
			case r == 110 : // ['n','n']
				return 44
			case 111 <= r && r <= 122 : // ['o','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S42
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 15
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case 97 <= r && r <= 109 : // ['a','m']
				return 16
			case r == 110 : // ['n','n']
				return 45
			case 111 <= r && r <= 122 : // ['o','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S43
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 15
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case 97 <= r && r <= 122 : // ['a','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S44
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 15
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case 97 <= r && r <= 99 : // ['a','c']
				return 16
			case r == 100 : // ['d','d']
				return 46
			case 101 <= r && r <= 122 : // ['e','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S45
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 15
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case 97 <= r && r <= 115 : // ['a','s']
				return 16
			case r == 116 : // ['t','t']
				return 47
			case 117 <= r && r <= 122 : // ['u','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S46
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 15
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case 97 <= r && r <= 122 : // ['a','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S47
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 15
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case 97 <= r && r <= 122 : // ['a','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
}
