
// generated by gocc; DO NOT EDIT.

package parser

type(
	actionTable [numStates]actionRow
	actionRow struct {
		canRecover bool
		actions [numSymbols]action
	}
)

var actionTab = actionTable{
	actionRow{ // S0
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(1),		/* $, reduce: Pipeline */
			nil,		/* empty */
			nil,		/* ; */
			shift(4),		/* component */
			nil,		/* id */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* , */
			shift(5),		/* connect */
			nil,		/* dot */
			nil,		/* to */
			nil,		/* string_lit */
			nil,		/* int_lit */
			nil,		/* true */
			nil,		/* false */
			
		},

	},
	actionRow{ // S1
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			accept(true),		/* $ */
			nil,		/* empty */
			nil,		/* ; */
			nil,		/* component */
			nil,		/* id */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* , */
			nil,		/* connect */
			nil,		/* dot */
			nil,		/* to */
			nil,		/* string_lit */
			nil,		/* int_lit */
			nil,		/* true */
			nil,		/* false */
			
		},

	},
	actionRow{ // S2
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			shift(6),		/* ; */
			nil,		/* component */
			nil,		/* id */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* , */
			nil,		/* connect */
			nil,		/* dot */
			nil,		/* to */
			nil,		/* string_lit */
			nil,		/* int_lit */
			nil,		/* true */
			nil,		/* false */
			
		},

	},
	actionRow{ // S3
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			shift(7),		/* ; */
			nil,		/* component */
			nil,		/* id */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* , */
			nil,		/* connect */
			nil,		/* dot */
			nil,		/* to */
			nil,		/* string_lit */
			nil,		/* int_lit */
			nil,		/* true */
			nil,		/* false */
			
		},

	},
	actionRow{ // S4
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* ; */
			nil,		/* component */
			shift(8),		/* id */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* , */
			nil,		/* connect */
			nil,		/* dot */
			nil,		/* to */
			nil,		/* string_lit */
			nil,		/* int_lit */
			nil,		/* true */
			nil,		/* false */
			
		},

	},
	actionRow{ // S5
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* ; */
			nil,		/* component */
			shift(9),		/* id */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* , */
			nil,		/* connect */
			nil,		/* dot */
			nil,		/* to */
			nil,		/* string_lit */
			nil,		/* int_lit */
			nil,		/* true */
			nil,		/* false */
			
		},

	},
	actionRow{ // S6
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(1),		/* $, reduce: Pipeline */
			nil,		/* empty */
			nil,		/* ; */
			shift(4),		/* component */
			nil,		/* id */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* , */
			shift(5),		/* connect */
			nil,		/* dot */
			nil,		/* to */
			nil,		/* string_lit */
			nil,		/* int_lit */
			nil,		/* true */
			nil,		/* false */
			
		},

	},
	actionRow{ // S7
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(1),		/* $, reduce: Pipeline */
			nil,		/* empty */
			nil,		/* ; */
			shift(4),		/* component */
			nil,		/* id */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* , */
			shift(5),		/* connect */
			nil,		/* dot */
			nil,		/* to */
			nil,		/* string_lit */
			nil,		/* int_lit */
			nil,		/* true */
			nil,		/* false */
			
		},

	},
	actionRow{ // S8
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* ; */
			nil,		/* component */
			shift(12),		/* id */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* , */
			nil,		/* connect */
			nil,		/* dot */
			nil,		/* to */
			nil,		/* string_lit */
			nil,		/* int_lit */
			nil,		/* true */
			nil,		/* false */
			
		},

	},
	actionRow{ // S9
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* ; */
			nil,		/* component */
			nil,		/* id */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* , */
			nil,		/* connect */
			shift(13),		/* dot */
			nil,		/* to */
			nil,		/* string_lit */
			nil,		/* int_lit */
			nil,		/* true */
			nil,		/* false */
			
		},

	},
	actionRow{ // S10
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(2),		/* $, reduce: Pipeline */
			nil,		/* empty */
			nil,		/* ; */
			nil,		/* component */
			nil,		/* id */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* , */
			nil,		/* connect */
			nil,		/* dot */
			nil,		/* to */
			nil,		/* string_lit */
			nil,		/* int_lit */
			nil,		/* true */
			nil,		/* false */
			
		},

	},
	actionRow{ // S11
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(3),		/* $, reduce: Pipeline */
			nil,		/* empty */
			nil,		/* ; */
			nil,		/* component */
			nil,		/* id */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* , */
			nil,		/* connect */
			nil,		/* dot */
			nil,		/* to */
			nil,		/* string_lit */
			nil,		/* int_lit */
			nil,		/* true */
			nil,		/* false */
			
		},

	},
	actionRow{ // S12
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* ; */
			nil,		/* component */
			nil,		/* id */
			shift(14),		/* ( */
			nil,		/* ) */
			nil,		/* , */
			nil,		/* connect */
			nil,		/* dot */
			nil,		/* to */
			nil,		/* string_lit */
			nil,		/* int_lit */
			nil,		/* true */
			nil,		/* false */
			
		},

	},
	actionRow{ // S13
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* ; */
			nil,		/* component */
			shift(15),		/* id */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* , */
			nil,		/* connect */
			nil,		/* dot */
			nil,		/* to */
			nil,		/* string_lit */
			nil,		/* int_lit */
			nil,		/* true */
			nil,		/* false */
			
		},

	},
	actionRow{ // S14
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* ; */
			nil,		/* component */
			nil,		/* id */
			nil,		/* ( */
			reduce(5),		/* ), reduce: PossiblyEmptyArgList */
			nil,		/* , */
			nil,		/* connect */
			nil,		/* dot */
			nil,		/* to */
			shift(22),		/* string_lit */
			shift(23),		/* int_lit */
			shift(24),		/* true */
			shift(25),		/* false */
			
		},

	},
	actionRow{ // S15
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* ; */
			nil,		/* component */
			nil,		/* id */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* , */
			nil,		/* connect */
			nil,		/* dot */
			shift(26),		/* to */
			nil,		/* string_lit */
			nil,		/* int_lit */
			nil,		/* true */
			nil,		/* false */
			
		},

	},
	actionRow{ // S16
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* ; */
			nil,		/* component */
			nil,		/* id */
			nil,		/* ( */
			shift(27),		/* ) */
			nil,		/* , */
			nil,		/* connect */
			nil,		/* dot */
			nil,		/* to */
			nil,		/* string_lit */
			nil,		/* int_lit */
			nil,		/* true */
			nil,		/* false */
			
		},

	},
	actionRow{ // S17
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* ; */
			nil,		/* component */
			nil,		/* id */
			nil,		/* ( */
			reduce(6),		/* ), reduce: PossiblyEmptyArgList */
			nil,		/* , */
			nil,		/* connect */
			nil,		/* dot */
			nil,		/* to */
			nil,		/* string_lit */
			nil,		/* int_lit */
			nil,		/* true */
			nil,		/* false */
			
		},

	},
	actionRow{ // S18
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* ; */
			nil,		/* component */
			nil,		/* id */
			nil,		/* ( */
			reduce(7),		/* ), reduce: ArgList */
			shift(28),		/* , */
			nil,		/* connect */
			nil,		/* dot */
			nil,		/* to */
			nil,		/* string_lit */
			nil,		/* int_lit */
			nil,		/* true */
			nil,		/* false */
			
		},

	},
	actionRow{ // S19
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* ; */
			nil,		/* component */
			nil,		/* id */
			nil,		/* ( */
			reduce(10),		/* ), reduce: Value */
			reduce(10),		/* ,, reduce: Value */
			nil,		/* connect */
			nil,		/* dot */
			nil,		/* to */
			nil,		/* string_lit */
			nil,		/* int_lit */
			nil,		/* true */
			nil,		/* false */
			
		},

	},
	actionRow{ // S20
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* ; */
			nil,		/* component */
			nil,		/* id */
			nil,		/* ( */
			reduce(11),		/* ), reduce: Value */
			reduce(11),		/* ,, reduce: Value */
			nil,		/* connect */
			nil,		/* dot */
			nil,		/* to */
			nil,		/* string_lit */
			nil,		/* int_lit */
			nil,		/* true */
			nil,		/* false */
			
		},

	},
	actionRow{ // S21
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* ; */
			nil,		/* component */
			nil,		/* id */
			nil,		/* ( */
			reduce(12),		/* ), reduce: Value */
			reduce(12),		/* ,, reduce: Value */
			nil,		/* connect */
			nil,		/* dot */
			nil,		/* to */
			nil,		/* string_lit */
			nil,		/* int_lit */
			nil,		/* true */
			nil,		/* false */
			
		},

	},
	actionRow{ // S22
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* ; */
			nil,		/* component */
			nil,		/* id */
			nil,		/* ( */
			reduce(13),		/* ), reduce: String */
			reduce(13),		/* ,, reduce: String */
			nil,		/* connect */
			nil,		/* dot */
			nil,		/* to */
			nil,		/* string_lit */
			nil,		/* int_lit */
			nil,		/* true */
			nil,		/* false */
			
		},

	},
	actionRow{ // S23
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* ; */
			nil,		/* component */
			nil,		/* id */
			nil,		/* ( */
			reduce(14),		/* ), reduce: Int */
			reduce(14),		/* ,, reduce: Int */
			nil,		/* connect */
			nil,		/* dot */
			nil,		/* to */
			nil,		/* string_lit */
			nil,		/* int_lit */
			nil,		/* true */
			nil,		/* false */
			
		},

	},
	actionRow{ // S24
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* ; */
			nil,		/* component */
			nil,		/* id */
			nil,		/* ( */
			reduce(15),		/* ), reduce: Bool */
			reduce(15),		/* ,, reduce: Bool */
			nil,		/* connect */
			nil,		/* dot */
			nil,		/* to */
			nil,		/* string_lit */
			nil,		/* int_lit */
			nil,		/* true */
			nil,		/* false */
			
		},

	},
	actionRow{ // S25
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* ; */
			nil,		/* component */
			nil,		/* id */
			nil,		/* ( */
			reduce(16),		/* ), reduce: Bool */
			reduce(16),		/* ,, reduce: Bool */
			nil,		/* connect */
			nil,		/* dot */
			nil,		/* to */
			nil,		/* string_lit */
			nil,		/* int_lit */
			nil,		/* true */
			nil,		/* false */
			
		},

	},
	actionRow{ // S26
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* ; */
			nil,		/* component */
			shift(29),		/* id */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* , */
			nil,		/* connect */
			nil,		/* dot */
			nil,		/* to */
			nil,		/* string_lit */
			nil,		/* int_lit */
			nil,		/* true */
			nil,		/* false */
			
		},

	},
	actionRow{ // S27
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			reduce(4),		/* ;, reduce: Component */
			nil,		/* component */
			nil,		/* id */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* , */
			nil,		/* connect */
			nil,		/* dot */
			nil,		/* to */
			nil,		/* string_lit */
			nil,		/* int_lit */
			nil,		/* true */
			nil,		/* false */
			
		},

	},
	actionRow{ // S28
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* ; */
			nil,		/* component */
			nil,		/* id */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* , */
			nil,		/* connect */
			nil,		/* dot */
			nil,		/* to */
			shift(22),		/* string_lit */
			shift(23),		/* int_lit */
			shift(24),		/* true */
			shift(25),		/* false */
			
		},

	},
	actionRow{ // S29
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* ; */
			nil,		/* component */
			nil,		/* id */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* , */
			nil,		/* connect */
			shift(31),		/* dot */
			nil,		/* to */
			nil,		/* string_lit */
			nil,		/* int_lit */
			nil,		/* true */
			nil,		/* false */
			
		},

	},
	actionRow{ // S30
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* ; */
			nil,		/* component */
			nil,		/* id */
			nil,		/* ( */
			reduce(8),		/* ), reduce: ArgList */
			nil,		/* , */
			nil,		/* connect */
			nil,		/* dot */
			nil,		/* to */
			nil,		/* string_lit */
			nil,		/* int_lit */
			nil,		/* true */
			nil,		/* false */
			
		},

	},
	actionRow{ // S31
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			nil,		/* ; */
			nil,		/* component */
			shift(32),		/* id */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* , */
			nil,		/* connect */
			nil,		/* dot */
			nil,		/* to */
			nil,		/* string_lit */
			nil,		/* int_lit */
			nil,		/* true */
			nil,		/* false */
			
		},

	},
	actionRow{ // S32
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* empty */
			reduce(9),		/* ;, reduce: Connection */
			nil,		/* component */
			nil,		/* id */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* , */
			nil,		/* connect */
			nil,		/* dot */
			nil,		/* to */
			nil,		/* string_lit */
			nil,		/* int_lit */
			nil,		/* true */
			nil,		/* false */
			
		},

	},
	
}

