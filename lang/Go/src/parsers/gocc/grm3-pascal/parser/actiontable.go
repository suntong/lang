// Code generated by gocc; DO NOT EDIT.

package parser

type (
	actionTable [numStates]actionRow
	actionRow   struct {
		canRecover bool
		actions    [numSymbols]action
	}
)

var actionTab = actionTable{
	actionRow{ // S0
		canRecover: false,
		actions: [numSymbols]action{
			nil,      // INVALID
			nil,      // $
			shift(2), // PROGRAM
			nil,      // id
			nil,      // PROCEDURE
			nil,      // ()
			nil,      // BEGIN
			nil,      // END.
			nil,      // ;
			nil,      // empty
			nil,      // FOR
			nil,      // TO
			nil,      // number
			nil,      // DO
			nil,      // :=
			nil,      // string_lit
			nil,      // char_lit
		},
	},
	actionRow{ // S1
		canRecover: false,
		actions: [numSymbols]action{
			nil,          // INVALID
			accept(true), // $
			nil,          // PROGRAM
			nil,          // id
			nil,          // PROCEDURE
			nil,          // ()
			nil,          // BEGIN
			nil,          // END.
			nil,          // ;
			nil,          // empty
			nil,          // FOR
			nil,          // TO
			nil,          // number
			nil,          // DO
			nil,          // :=
			nil,          // string_lit
			nil,          // char_lit
		},
	},
	actionRow{ // S2
		canRecover: false,
		actions: [numSymbols]action{
			nil,      // INVALID
			nil,      // $
			nil,      // PROGRAM
			shift(3), // id
			nil,      // PROCEDURE
			nil,      // ()
			nil,      // BEGIN
			nil,      // END.
			nil,      // ;
			nil,      // empty
			nil,      // FOR
			nil,      // TO
			nil,      // number
			nil,      // DO
			nil,      // :=
			nil,      // string_lit
			nil,      // char_lit
		},
	},
	actionRow{ // S3
		canRecover: false,
		actions: [numSymbols]action{
			nil,      // INVALID
			nil,      // $
			nil,      // PROGRAM
			nil,      // id
			shift(7), // PROCEDURE
			nil,      // ()
			shift(8), // BEGIN
			nil,      // END.
			nil,      // ;
			nil,      // empty
			nil,      // FOR
			nil,      // TO
			nil,      // number
			nil,      // DO
			nil,      // :=
			nil,      // string_lit
			nil,      // char_lit
		},
	},
	actionRow{ // S4
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			reduce(1), // $, reduce: Program
			nil,       // PROGRAM
			nil,       // id
			nil,       // PROCEDURE
			nil,       // ()
			nil,       // BEGIN
			nil,       // END.
			nil,       // ;
			nil,       // empty
			nil,       // FOR
			nil,       // TO
			nil,       // number
			nil,       // DO
			nil,       // :=
			nil,       // string_lit
			nil,       // char_lit
		},
	},
	actionRow{ // S5
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			reduce(2), // $, reduce: ProgramDef
			nil,       // PROGRAM
			nil,       // id
			nil,       // PROCEDURE
			nil,       // ()
			nil,       // BEGIN
			nil,       // END.
			nil,       // ;
			nil,       // empty
			nil,       // FOR
			nil,       // TO
			nil,       // number
			nil,       // DO
			nil,       // :=
			nil,       // string_lit
			nil,       // char_lit
		},
	},
	actionRow{ // S6
		canRecover: false,
		actions: [numSymbols]action{
			nil,      // INVALID
			nil,      // $
			nil,      // PROGRAM
			nil,      // id
			shift(7), // PROCEDURE
			nil,      // ()
			shift(8), // BEGIN
			nil,      // END.
			nil,      // ;
			nil,      // empty
			nil,      // FOR
			nil,      // TO
			nil,      // number
			nil,      // DO
			nil,      // :=
			nil,      // string_lit
			nil,      // char_lit
		},
	},
	actionRow{ // S7
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // PROGRAM
			shift(10), // id
			nil,       // PROCEDURE
			nil,       // ()
			nil,       // BEGIN
			nil,       // END.
			nil,       // ;
			nil,       // empty
			nil,       // FOR
			nil,       // TO
			nil,       // number
			nil,       // DO
			nil,       // :=
			nil,       // string_lit
			nil,       // char_lit
		},
	},
	actionRow{ // S8
		canRecover: false,
		actions: [numSymbols]action{
			nil,        // INVALID
			nil,        // $
			nil,        // PROGRAM
			shift(11),  // id
			nil,        // PROCEDURE
			nil,        // ()
			nil,        // BEGIN
			reduce(10), // END., reduce: ProgramBody
			nil,        // ;
			nil,        // empty
			shift(15),  // FOR
			nil,        // TO
			nil,        // number
			nil,        // DO
			nil,        // :=
			nil,        // string_lit
			nil,        // char_lit
		},
	},
	actionRow{ // S9
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			reduce(3), // $, reduce: ProgramDef
			nil,       // PROGRAM
			nil,       // id
			nil,       // PROCEDURE
			nil,       // ()
			nil,       // BEGIN
			nil,       // END.
			nil,       // ;
			nil,       // empty
			nil,       // FOR
			nil,       // TO
			nil,       // number
			nil,       // DO
			nil,       // :=
			nil,       // string_lit
			nil,       // char_lit
		},
	},
	actionRow{ // S10
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // PROGRAM
			nil,       // id
			nil,       // PROCEDURE
			shift(16), // ()
			nil,       // BEGIN
			nil,       // END.
			nil,       // ;
			nil,       // empty
			nil,       // FOR
			nil,       // TO
			nil,       // number
			nil,       // DO
			nil,       // :=
			nil,       // string_lit
			nil,       // char_lit
		},
	},
	actionRow{ // S11
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // PROGRAM
			nil,       // id
			nil,       // PROCEDURE
			nil,       // ()
			nil,       // BEGIN
			nil,       // END.
			nil,       // ;
			nil,       // empty
			nil,       // FOR
			nil,       // TO
			nil,       // number
			nil,       // DO
			shift(17), // :=
			nil,       // string_lit
			nil,       // char_lit
		},
	},
	actionRow{ // S12
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // PROGRAM
			shift(11), // id
			nil,       // PROCEDURE
			nil,       // ()
			nil,       // BEGIN
			shift(18), // END.
			nil,       // ;
			nil,       // empty
			shift(21), // FOR
			nil,       // TO
			nil,       // number
			nil,       // DO
			nil,       // :=
			nil,       // string_lit
			nil,       // char_lit
		},
	},
	actionRow{ // S13
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // PROGRAM
			nil,       // id
			nil,       // PROCEDURE
			nil,       // ()
			nil,       // BEGIN
			nil,       // END.
			shift(22), // ;
			nil,       // empty
			nil,       // FOR
			nil,       // TO
			nil,       // number
			nil,       // DO
			nil,       // :=
			nil,       // string_lit
			nil,       // char_lit
		},
	},
	actionRow{ // S14
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // PROGRAM
			reduce(8), // id, reduce: ProgramBody
			nil,       // PROCEDURE
			nil,       // ()
			nil,       // BEGIN
			reduce(8), // END., reduce: ProgramBody
			nil,       // ;
			nil,       // empty
			reduce(8), // FOR, reduce: ProgramBody
			nil,       // TO
			nil,       // number
			nil,       // DO
			nil,       // :=
			nil,       // string_lit
			nil,       // char_lit
		},
	},
	actionRow{ // S15
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // PROGRAM
			shift(23), // id
			nil,       // PROCEDURE
			nil,       // ()
			nil,       // BEGIN
			nil,       // END.
			nil,       // ;
			nil,       // empty
			nil,       // FOR
			nil,       // TO
			nil,       // number
			nil,       // DO
			nil,       // :=
			nil,       // string_lit
			nil,       // char_lit
		},
	},
	actionRow{ // S16
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // PROGRAM
			nil,       // id
			nil,       // PROCEDURE
			nil,       // ()
			shift(26), // BEGIN
			nil,       // END.
			nil,       // ;
			nil,       // empty
			nil,       // FOR
			nil,       // TO
			nil,       // number
			nil,       // DO
			nil,       // :=
			nil,       // string_lit
			nil,       // char_lit
		},
	},
	actionRow{ // S17
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // PROGRAM
			shift(27), // id
			nil,       // PROCEDURE
			nil,       // ()
			nil,       // BEGIN
			nil,       // END.
			nil,       // ;
			nil,       // empty
			nil,       // FOR
			nil,       // TO
			shift(28), // number
			nil,       // DO
			nil,       // :=
			shift(30), // string_lit
			shift(31), // char_lit
		},
	},
	actionRow{ // S18
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			reduce(5), // $, reduce: BeginBlock
			nil,       // PROGRAM
			nil,       // id
			nil,       // PROCEDURE
			nil,       // ()
			nil,       // BEGIN
			nil,       // END.
			nil,       // ;
			nil,       // empty
			nil,       // FOR
			nil,       // TO
			nil,       // number
			nil,       // DO
			nil,       // :=
			nil,       // string_lit
			nil,       // char_lit
		},
	},
	actionRow{ // S19
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // PROGRAM
			nil,       // id
			nil,       // PROCEDURE
			nil,       // ()
			nil,       // BEGIN
			nil,       // END.
			shift(32), // ;
			nil,       // empty
			nil,       // FOR
			nil,       // TO
			nil,       // number
			nil,       // DO
			nil,       // :=
			nil,       // string_lit
			nil,       // char_lit
		},
	},
	actionRow{ // S20
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // PROGRAM
			nil,       // id
			nil,       // PROCEDURE
			nil,       // ()
			nil,       // BEGIN
			nil,       // END.
			shift(33), // ;
			nil,       // empty
			nil,       // FOR
			nil,       // TO
			nil,       // number
			nil,       // DO
			nil,       // :=
			nil,       // string_lit
			nil,       // char_lit
		},
	},
	actionRow{ // S21
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // PROGRAM
			shift(23), // id
			nil,       // PROCEDURE
			nil,       // ()
			nil,       // BEGIN
			nil,       // END.
			nil,       // ;
			nil,       // empty
			nil,       // FOR
			nil,       // TO
			nil,       // number
			nil,       // DO
			nil,       // :=
			nil,       // string_lit
			nil,       // char_lit
		},
	},
	actionRow{ // S22
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // PROGRAM
			reduce(6), // id, reduce: ProgramBody
			nil,       // PROCEDURE
			nil,       // ()
			nil,       // BEGIN
			reduce(6), // END., reduce: ProgramBody
			nil,       // ;
			nil,       // empty
			reduce(6), // FOR, reduce: ProgramBody
			nil,       // TO
			nil,       // number
			nil,       // DO
			nil,       // :=
			nil,       // string_lit
			nil,       // char_lit
		},
	},
	actionRow{ // S23
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // PROGRAM
			nil,       // id
			nil,       // PROCEDURE
			nil,       // ()
			nil,       // BEGIN
			nil,       // END.
			nil,       // ;
			nil,       // empty
			nil,       // FOR
			nil,       // TO
			nil,       // number
			nil,       // DO
			shift(35), // :=
			nil,       // string_lit
			nil,       // char_lit
		},
	},
	actionRow{ // S24
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // PROGRAM
			nil,       // id
			nil,       // PROCEDURE
			nil,       // ()
			nil,       // BEGIN
			nil,       // END.
			nil,       // ;
			nil,       // empty
			nil,       // FOR
			shift(36), // TO
			nil,       // number
			nil,       // DO
			nil,       // :=
			nil,       // string_lit
			nil,       // char_lit
		},
	},
	actionRow{ // S25
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // PROGRAM
			nil,       // id
			reduce(4), // PROCEDURE, reduce: Procedure
			nil,       // ()
			reduce(4), // BEGIN, reduce: Procedure
			nil,       // END.
			nil,       // ;
			nil,       // empty
			nil,       // FOR
			nil,       // TO
			nil,       // number
			nil,       // DO
			nil,       // :=
			nil,       // string_lit
			nil,       // char_lit
		},
	},
	actionRow{ // S26
		canRecover: false,
		actions: [numSymbols]action{
			nil,        // INVALID
			nil,        // $
			nil,        // PROGRAM
			shift(11),  // id
			nil,        // PROCEDURE
			nil,        // ()
			nil,        // BEGIN
			reduce(10), // END., reduce: ProgramBody
			nil,        // ;
			nil,        // empty
			shift(15),  // FOR
			nil,        // TO
			nil,        // number
			nil,        // DO
			nil,        // :=
			nil,        // string_lit
			nil,        // char_lit
		},
	},
	actionRow{ // S27
		canRecover: false,
		actions: [numSymbols]action{
			nil,        // INVALID
			nil,        // $
			nil,        // PROGRAM
			nil,        // id
			nil,        // PROCEDURE
			nil,        // ()
			nil,        // BEGIN
			nil,        // END.
			reduce(13), // ;, reduce: RHS
			nil,        // empty
			nil,        // FOR
			nil,        // TO
			nil,        // number
			nil,        // DO
			nil,        // :=
			nil,        // string_lit
			nil,        // char_lit
		},
	},
	actionRow{ // S28
		canRecover: false,
		actions: [numSymbols]action{
			nil,        // INVALID
			nil,        // $
			nil,        // PROGRAM
			nil,        // id
			nil,        // PROCEDURE
			nil,        // ()
			nil,        // BEGIN
			nil,        // END.
			reduce(16), // ;, reduce: RHS
			nil,        // empty
			nil,        // FOR
			nil,        // TO
			nil,        // number
			nil,        // DO
			nil,        // :=
			nil,        // string_lit
			nil,        // char_lit
		},
	},
	actionRow{ // S29
		canRecover: false,
		actions: [numSymbols]action{
			nil,        // INVALID
			nil,        // $
			nil,        // PROGRAM
			nil,        // id
			nil,        // PROCEDURE
			nil,        // ()
			nil,        // BEGIN
			nil,        // END.
			reduce(12), // ;, reduce: Assignment
			nil,        // empty
			nil,        // FOR
			nil,        // TO
			nil,        // number
			nil,        // DO
			nil,        // :=
			nil,        // string_lit
			nil,        // char_lit
		},
	},
	actionRow{ // S30
		canRecover: false,
		actions: [numSymbols]action{
			nil,        // INVALID
			nil,        // $
			nil,        // PROGRAM
			nil,        // id
			nil,        // PROCEDURE
			nil,        // ()
			nil,        // BEGIN
			nil,        // END.
			reduce(14), // ;, reduce: RHS
			nil,        // empty
			nil,        // FOR
			nil,        // TO
			nil,        // number
			nil,        // DO
			nil,        // :=
			nil,        // string_lit
			nil,        // char_lit
		},
	},
	actionRow{ // S31
		canRecover: false,
		actions: [numSymbols]action{
			nil,        // INVALID
			nil,        // $
			nil,        // PROGRAM
			nil,        // id
			nil,        // PROCEDURE
			nil,        // ()
			nil,        // BEGIN
			nil,        // END.
			reduce(15), // ;, reduce: RHS
			nil,        // empty
			nil,        // FOR
			nil,        // TO
			nil,        // number
			nil,        // DO
			nil,        // :=
			nil,        // string_lit
			nil,        // char_lit
		},
	},
	actionRow{ // S32
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // PROGRAM
			reduce(7), // id, reduce: ProgramBody
			nil,       // PROCEDURE
			nil,       // ()
			nil,       // BEGIN
			reduce(7), // END., reduce: ProgramBody
			nil,       // ;
			nil,       // empty
			reduce(7), // FOR, reduce: ProgramBody
			nil,       // TO
			nil,       // number
			nil,       // DO
			nil,       // :=
			nil,       // string_lit
			nil,       // char_lit
		},
	},
	actionRow{ // S33
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // PROGRAM
			reduce(9), // id, reduce: ProgramBody
			nil,       // PROCEDURE
			nil,       // ()
			nil,       // BEGIN
			reduce(9), // END., reduce: ProgramBody
			nil,       // ;
			nil,       // empty
			reduce(9), // FOR, reduce: ProgramBody
			nil,       // TO
			nil,       // number
			nil,       // DO
			nil,       // :=
			nil,       // string_lit
			nil,       // char_lit
		},
	},
	actionRow{ // S34
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // PROGRAM
			nil,       // id
			nil,       // PROCEDURE
			nil,       // ()
			nil,       // BEGIN
			nil,       // END.
			nil,       // ;
			nil,       // empty
			nil,       // FOR
			shift(38), // TO
			nil,       // number
			nil,       // DO
			nil,       // :=
			nil,       // string_lit
			nil,       // char_lit
		},
	},
	actionRow{ // S35
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // PROGRAM
			shift(39), // id
			nil,       // PROCEDURE
			nil,       // ()
			nil,       // BEGIN
			nil,       // END.
			nil,       // ;
			nil,       // empty
			nil,       // FOR
			nil,       // TO
			shift(40), // number
			nil,       // DO
			nil,       // :=
			shift(42), // string_lit
			shift(43), // char_lit
		},
	},
	actionRow{ // S36
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // PROGRAM
			nil,       // id
			nil,       // PROCEDURE
			nil,       // ()
			nil,       // BEGIN
			nil,       // END.
			nil,       // ;
			nil,       // empty
			nil,       // FOR
			nil,       // TO
			shift(44), // number
			nil,       // DO
			nil,       // :=
			nil,       // string_lit
			nil,       // char_lit
		},
	},
	actionRow{ // S37
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // PROGRAM
			shift(11), // id
			nil,       // PROCEDURE
			nil,       // ()
			nil,       // BEGIN
			shift(45), // END.
			nil,       // ;
			nil,       // empty
			shift(21), // FOR
			nil,       // TO
			nil,       // number
			nil,       // DO
			nil,       // :=
			nil,       // string_lit
			nil,       // char_lit
		},
	},
	actionRow{ // S38
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // PROGRAM
			nil,       // id
			nil,       // PROCEDURE
			nil,       // ()
			nil,       // BEGIN
			nil,       // END.
			nil,       // ;
			nil,       // empty
			nil,       // FOR
			nil,       // TO
			shift(46), // number
			nil,       // DO
			nil,       // :=
			nil,       // string_lit
			nil,       // char_lit
		},
	},
	actionRow{ // S39
		canRecover: false,
		actions: [numSymbols]action{
			nil,        // INVALID
			nil,        // $
			nil,        // PROGRAM
			nil,        // id
			nil,        // PROCEDURE
			nil,        // ()
			nil,        // BEGIN
			nil,        // END.
			nil,        // ;
			nil,        // empty
			nil,        // FOR
			reduce(13), // TO, reduce: RHS
			nil,        // number
			nil,        // DO
			nil,        // :=
			nil,        // string_lit
			nil,        // char_lit
		},
	},
	actionRow{ // S40
		canRecover: false,
		actions: [numSymbols]action{
			nil,        // INVALID
			nil,        // $
			nil,        // PROGRAM
			nil,        // id
			nil,        // PROCEDURE
			nil,        // ()
			nil,        // BEGIN
			nil,        // END.
			nil,        // ;
			nil,        // empty
			nil,        // FOR
			reduce(16), // TO, reduce: RHS
			nil,        // number
			nil,        // DO
			nil,        // :=
			nil,        // string_lit
			nil,        // char_lit
		},
	},
	actionRow{ // S41
		canRecover: false,
		actions: [numSymbols]action{
			nil,        // INVALID
			nil,        // $
			nil,        // PROGRAM
			nil,        // id
			nil,        // PROCEDURE
			nil,        // ()
			nil,        // BEGIN
			nil,        // END.
			nil,        // ;
			nil,        // empty
			nil,        // FOR
			reduce(12), // TO, reduce: Assignment
			nil,        // number
			nil,        // DO
			nil,        // :=
			nil,        // string_lit
			nil,        // char_lit
		},
	},
	actionRow{ // S42
		canRecover: false,
		actions: [numSymbols]action{
			nil,        // INVALID
			nil,        // $
			nil,        // PROGRAM
			nil,        // id
			nil,        // PROCEDURE
			nil,        // ()
			nil,        // BEGIN
			nil,        // END.
			nil,        // ;
			nil,        // empty
			nil,        // FOR
			reduce(14), // TO, reduce: RHS
			nil,        // number
			nil,        // DO
			nil,        // :=
			nil,        // string_lit
			nil,        // char_lit
		},
	},
	actionRow{ // S43
		canRecover: false,
		actions: [numSymbols]action{
			nil,        // INVALID
			nil,        // $
			nil,        // PROGRAM
			nil,        // id
			nil,        // PROCEDURE
			nil,        // ()
			nil,        // BEGIN
			nil,        // END.
			nil,        // ;
			nil,        // empty
			nil,        // FOR
			reduce(15), // TO, reduce: RHS
			nil,        // number
			nil,        // DO
			nil,        // :=
			nil,        // string_lit
			nil,        // char_lit
		},
	},
	actionRow{ // S44
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // PROGRAM
			nil,       // id
			nil,       // PROCEDURE
			nil,       // ()
			nil,       // BEGIN
			nil,       // END.
			nil,       // ;
			nil,       // empty
			nil,       // FOR
			nil,       // TO
			nil,       // number
			shift(47), // DO
			nil,       // :=
			nil,       // string_lit
			nil,       // char_lit
		},
	},
	actionRow{ // S45
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // PROGRAM
			nil,       // id
			reduce(5), // PROCEDURE, reduce: BeginBlock
			nil,       // ()
			reduce(5), // BEGIN, reduce: BeginBlock
			nil,       // END.
			nil,       // ;
			nil,       // empty
			nil,       // FOR
			nil,       // TO
			nil,       // number
			nil,       // DO
			nil,       // :=
			nil,       // string_lit
			nil,       // char_lit
		},
	},
	actionRow{ // S46
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // PROGRAM
			nil,       // id
			nil,       // PROCEDURE
			nil,       // ()
			nil,       // BEGIN
			nil,       // END.
			nil,       // ;
			nil,       // empty
			nil,       // FOR
			nil,       // TO
			nil,       // number
			shift(48), // DO
			nil,       // :=
			nil,       // string_lit
			nil,       // char_lit
		},
	},
	actionRow{ // S47
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // PROGRAM
			shift(49), // id
			nil,       // PROCEDURE
			nil,       // ()
			nil,       // BEGIN
			nil,       // END.
			nil,       // ;
			nil,       // empty
			nil,       // FOR
			nil,       // TO
			nil,       // number
			nil,       // DO
			nil,       // :=
			nil,       // string_lit
			nil,       // char_lit
		},
	},
	actionRow{ // S48
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // PROGRAM
			shift(11), // id
			nil,       // PROCEDURE
			nil,       // ()
			nil,       // BEGIN
			nil,       // END.
			nil,       // ;
			nil,       // empty
			nil,       // FOR
			nil,       // TO
			nil,       // number
			nil,       // DO
			nil,       // :=
			nil,       // string_lit
			nil,       // char_lit
		},
	},
	actionRow{ // S49
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // PROGRAM
			nil,       // id
			nil,       // PROCEDURE
			nil,       // ()
			nil,       // BEGIN
			nil,       // END.
			nil,       // ;
			nil,       // empty
			nil,       // FOR
			nil,       // TO
			nil,       // number
			nil,       // DO
			shift(52), // :=
			nil,       // string_lit
			nil,       // char_lit
		},
	},
	actionRow{ // S50
		canRecover: false,
		actions: [numSymbols]action{
			nil,        // INVALID
			nil,        // $
			nil,        // PROGRAM
			reduce(11), // id, reduce: ForLoop
			nil,        // PROCEDURE
			nil,        // ()
			nil,        // BEGIN
			reduce(11), // END., reduce: ForLoop
			nil,        // ;
			nil,        // empty
			reduce(11), // FOR, reduce: ForLoop
			nil,        // TO
			nil,        // number
			nil,        // DO
			nil,        // :=
			nil,        // string_lit
			nil,        // char_lit
		},
	},
	actionRow{ // S51
		canRecover: false,
		actions: [numSymbols]action{
			nil,        // INVALID
			nil,        // $
			nil,        // PROGRAM
			nil,        // id
			nil,        // PROCEDURE
			nil,        // ()
			nil,        // BEGIN
			nil,        // END.
			reduce(11), // ;, reduce: ForLoop
			nil,        // empty
			nil,        // FOR
			nil,        // TO
			nil,        // number
			nil,        // DO
			nil,        // :=
			nil,        // string_lit
			nil,        // char_lit
		},
	},
	actionRow{ // S52
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // PROGRAM
			shift(53), // id
			nil,       // PROCEDURE
			nil,       // ()
			nil,       // BEGIN
			nil,       // END.
			nil,       // ;
			nil,       // empty
			nil,       // FOR
			nil,       // TO
			shift(54), // number
			nil,       // DO
			nil,       // :=
			shift(56), // string_lit
			shift(57), // char_lit
		},
	},
	actionRow{ // S53
		canRecover: false,
		actions: [numSymbols]action{
			nil,        // INVALID
			nil,        // $
			nil,        // PROGRAM
			reduce(13), // id, reduce: RHS
			nil,        // PROCEDURE
			nil,        // ()
			nil,        // BEGIN
			reduce(13), // END., reduce: RHS
			nil,        // ;
			nil,        // empty
			reduce(13), // FOR, reduce: RHS
			nil,        // TO
			nil,        // number
			nil,        // DO
			nil,        // :=
			nil,        // string_lit
			nil,        // char_lit
		},
	},
	actionRow{ // S54
		canRecover: false,
		actions: [numSymbols]action{
			nil,        // INVALID
			nil,        // $
			nil,        // PROGRAM
			reduce(16), // id, reduce: RHS
			nil,        // PROCEDURE
			nil,        // ()
			nil,        // BEGIN
			reduce(16), // END., reduce: RHS
			nil,        // ;
			nil,        // empty
			reduce(16), // FOR, reduce: RHS
			nil,        // TO
			nil,        // number
			nil,        // DO
			nil,        // :=
			nil,        // string_lit
			nil,        // char_lit
		},
	},
	actionRow{ // S55
		canRecover: false,
		actions: [numSymbols]action{
			nil,        // INVALID
			nil,        // $
			nil,        // PROGRAM
			reduce(12), // id, reduce: Assignment
			nil,        // PROCEDURE
			nil,        // ()
			nil,        // BEGIN
			reduce(12), // END., reduce: Assignment
			nil,        // ;
			nil,        // empty
			reduce(12), // FOR, reduce: Assignment
			nil,        // TO
			nil,        // number
			nil,        // DO
			nil,        // :=
			nil,        // string_lit
			nil,        // char_lit
		},
	},
	actionRow{ // S56
		canRecover: false,
		actions: [numSymbols]action{
			nil,        // INVALID
			nil,        // $
			nil,        // PROGRAM
			reduce(14), // id, reduce: RHS
			nil,        // PROCEDURE
			nil,        // ()
			nil,        // BEGIN
			reduce(14), // END., reduce: RHS
			nil,        // ;
			nil,        // empty
			reduce(14), // FOR, reduce: RHS
			nil,        // TO
			nil,        // number
			nil,        // DO
			nil,        // :=
			nil,        // string_lit
			nil,        // char_lit
		},
	},
	actionRow{ // S57
		canRecover: false,
		actions: [numSymbols]action{
			nil,        // INVALID
			nil,        // $
			nil,        // PROGRAM
			reduce(15), // id, reduce: RHS
			nil,        // PROCEDURE
			nil,        // ()
			nil,        // BEGIN
			reduce(15), // END., reduce: RHS
			nil,        // ;
			nil,        // empty
			reduce(15), // FOR, reduce: RHS
			nil,        // TO
			nil,        // number
			nil,        // DO
			nil,        // :=
			nil,        // string_lit
			nil,        // char_lit
		},
	},
}