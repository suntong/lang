// Code generated by gocc; DO NOT EDIT.

package lexer

/*
Let s be the current state
Let r be the current input rune
transitionTable[s](r) returns the next state.
*/
type TransitionTable [NumStates]func(rune) int

var TransTab = TransitionTable{
	// S0
	func(r rune) int {
		switch {
		case r == 9: // ['\t','\t']
			return 1
		case r == 10: // ['\n','\n']
			return 1
		case r == 11: // ['\v','\v']
			return 1
		case r == 12: // ['\f','\f']
			return 1
		case r == 13: // ['\r','\r']
			return 1
		case r == 32: // [' ',' ']
			return 1
		case r == 33: // ['!','!']
			return 2
		case r == 34: // ['"','"']
			return 3
		case 35 <= r && r <= 38: // ['#','&']
			return 4
		case r == 39: // [''',''']
			return 5
		case r == 40: // ['(','(']
			return 6
		case r == 41: // [')',')']
			return 7
		case 42 <= r && r <= 44: // ['*',',']
			return 4
		case r == 45: // ['-','-']
			return 8
		case r == 46: // ['.','.']
			return 9
		case r == 47: // ['/','/']
			return 10
		case 48 <= r && r <= 57: // ['0','9']
			return 11
		case r == 58: // [':',':']
			return 12
		case r == 59: // [';',';']
			return 13
		case 60 <= r && r <= 64: // ['<','@']
			return 4
		case 65 <= r && r <= 90: // ['A','Z']
			return 14
		case r == 91: // ['[','[']
			return 15
		case r == 92: // ['\','\']
			return 4
		case r == 93: // [']',']']
			return 16
		case r == 94: // ['^','^']
			return 4
		case r == 95: // ['_','_']
			return 14
		case 97 <= r && r <= 122: // ['a','z']
			return 14
		case r == 123: // ['{','{']
			return 17
		case r == 124: // ['|','|']
			return 18
		case r == 125: // ['}','}']
			return 19
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
		case 65 <= r && r <= 90: // ['A','Z']
			return 20
		case r == 95: // ['_','_']
			return 20
		case 97 <= r && r <= 122: // ['a','z']
			return 20
		}
		return NoState
	},
	// S3
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 21
		case 35 <= r && r <= 47: // ['#','/']
			return 21
		case 48 <= r && r <= 57: // ['0','9']
			return 22
		case 58 <= r && r <= 64: // [':','@']
			return 21
		case 65 <= r && r <= 90: // ['A','Z']
			return 23
		case 91 <= r && r <= 94: // ['[','^']
			return 21
		case r == 95: // ['_','_']
			return 23
		case 97 <= r && r <= 122: // ['a','z']
			return 23
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
		case r == 9: // ['\t','\t']
			return 24
		case r == 10: // ['\n','\n']
			return 24
		case r == 11: // ['\v','\v']
			return 24
		case r == 12: // ['\f','\f']
			return 24
		case r == 13: // ['\r','\r']
			return 24
		case r == 32: // [' ',' ']
			return 24
		case r == 33: // ['!','!']
			return 25
		case 35 <= r && r <= 47: // ['#','/']
			return 25
		case 48 <= r && r <= 57: // ['0','9']
			return 26
		case 58 <= r && r <= 64: // [':','@']
			return 25
		case 65 <= r && r <= 90: // ['A','Z']
			return 27
		case 91 <= r && r <= 94: // ['[','^']
			return 25
		case r == 95: // ['_','_']
			return 27
		case 97 <= r && r <= 122: // ['a','z']
			return 27
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
		}
		return NoState
	},
	// S10
	func(r rune) int {
		switch {
		case r == 42: // ['*','*']
			return 28
		case r == 47: // ['/','/']
			return 29
		}
		return NoState
	},
	// S11
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S12
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S13
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S14
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 30
		case 65 <= r && r <= 90: // ['A','Z']
			return 31
		case r == 95: // ['_','_']
			return 31
		case 97 <= r && r <= 122: // ['a','z']
			return 31
		}
		return NoState
	},
	// S15
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S16
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S17
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S18
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S19
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S20
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 30
		case 65 <= r && r <= 90: // ['A','Z']
			return 31
		case r == 95: // ['_','_']
			return 31
		case 97 <= r && r <= 122: // ['a','z']
			return 31
		}
		return NoState
	},
	// S21
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 21
		case r == 34: // ['"','"']
			return 32
		case 35 <= r && r <= 47: // ['#','/']
			return 21
		case 48 <= r && r <= 57: // ['0','9']
			return 22
		case 58 <= r && r <= 64: // [':','@']
			return 21
		case 65 <= r && r <= 90: // ['A','Z']
			return 23
		case 91 <= r && r <= 94: // ['[','^']
			return 21
		case r == 95: // ['_','_']
			return 23
		case 97 <= r && r <= 122: // ['a','z']
			return 23
		}
		return NoState
	},
	// S22
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 21
		case r == 34: // ['"','"']
			return 32
		case 35 <= r && r <= 47: // ['#','/']
			return 21
		case 48 <= r && r <= 57: // ['0','9']
			return 22
		case 58 <= r && r <= 64: // [':','@']
			return 21
		case 65 <= r && r <= 90: // ['A','Z']
			return 23
		case 91 <= r && r <= 94: // ['[','^']
			return 21
		case r == 95: // ['_','_']
			return 23
		case 97 <= r && r <= 122: // ['a','z']
			return 23
		}
		return NoState
	},
	// S23
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 21
		case r == 34: // ['"','"']
			return 32
		case 35 <= r && r <= 47: // ['#','/']
			return 21
		case 48 <= r && r <= 57: // ['0','9']
			return 22
		case 58 <= r && r <= 64: // [':','@']
			return 21
		case 65 <= r && r <= 90: // ['A','Z']
			return 23
		case 91 <= r && r <= 94: // ['[','^']
			return 21
		case r == 95: // ['_','_']
			return 23
		case 97 <= r && r <= 122: // ['a','z']
			return 23
		}
		return NoState
	},
	// S24
	func(r rune) int {
		switch {
		case r == 39: // [''',''']
			return 32
		}
		return NoState
	},
	// S25
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 25
		case 35 <= r && r <= 38: // ['#','&']
			return 25
		case r == 39: // [''',''']
			return 33
		case 40 <= r && r <= 47: // ['(','/']
			return 25
		case 48 <= r && r <= 57: // ['0','9']
			return 26
		case 58 <= r && r <= 64: // [':','@']
			return 25
		case 65 <= r && r <= 90: // ['A','Z']
			return 27
		case 91 <= r && r <= 94: // ['[','^']
			return 25
		case r == 95: // ['_','_']
			return 27
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S26
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 25
		case 35 <= r && r <= 38: // ['#','&']
			return 25
		case r == 39: // [''',''']
			return 33
		case 40 <= r && r <= 47: // ['(','/']
			return 25
		case 48 <= r && r <= 57: // ['0','9']
			return 26
		case 58 <= r && r <= 64: // [':','@']
			return 25
		case 65 <= r && r <= 90: // ['A','Z']
			return 27
		case 91 <= r && r <= 94: // ['[','^']
			return 25
		case r == 95: // ['_','_']
			return 27
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S27
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 25
		case 35 <= r && r <= 38: // ['#','&']
			return 25
		case r == 39: // [''',''']
			return 33
		case 40 <= r && r <= 47: // ['(','/']
			return 25
		case 48 <= r && r <= 57: // ['0','9']
			return 26
		case 58 <= r && r <= 64: // [':','@']
			return 25
		case 65 <= r && r <= 90: // ['A','Z']
			return 27
		case 91 <= r && r <= 94: // ['[','^']
			return 25
		case r == 95: // ['_','_']
			return 27
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S28
	func(r rune) int {
		switch {
		case r == 42: // ['*','*']
			return 34
		default:
			return 28
		}
	},
	// S29
	func(r rune) int {
		switch {
		case r == 10: // ['\n','\n']
			return 35
		default:
			return 29
		}
	},
	// S30
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 30
		case 65 <= r && r <= 90: // ['A','Z']
			return 31
		case r == 95: // ['_','_']
			return 31
		case 97 <= r && r <= 122: // ['a','z']
			return 31
		}
		return NoState
	},
	// S31
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 30
		case 65 <= r && r <= 90: // ['A','Z']
			return 31
		case r == 95: // ['_','_']
			return 31
		case 97 <= r && r <= 122: // ['a','z']
			return 31
		}
		return NoState
	},
	// S32
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S33
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 25
		case 35 <= r && r <= 38: // ['#','&']
			return 25
		case r == 39: // [''',''']
			return 33
		case 40 <= r && r <= 47: // ['(','/']
			return 25
		case 48 <= r && r <= 57: // ['0','9']
			return 26
		case 58 <= r && r <= 64: // [':','@']
			return 25
		case 65 <= r && r <= 90: // ['A','Z']
			return 27
		case 91 <= r && r <= 94: // ['[','^']
			return 25
		case r == 95: // ['_','_']
			return 27
		case 97 <= r && r <= 122: // ['a','z']
			return 27
		}
		return NoState
	},
	// S34
	func(r rune) int {
		switch {
		case r == 42: // ['*','*']
			return 34
		case r == 47: // ['/','/']
			return 36
		default:
			return 28
		}
	},
	// S35
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S36
	func(r rune) int {
		switch {
		}
		return NoState
	},
}
