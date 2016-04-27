package flow

import (
	"fmt"

	"go/types"

	"github.com/dustin/go-humanize"
)

func cardinalityMismatchError(source, dest ComponentID, sourceSig, destSig *types.Tuple) error {
	return fmt.Errorf(`
As I infer the types of values flowing through your program, I see a mismatch in this connection.

	%[1]s -> %[2]s

There are %[3]d results coming from %[1]s:

	%[4]s

But %[2]s is expecting %[5]d argument[s]:

	%[6]s

HINT: These should have identical length and types.
`, source, dest, sourceSig.Len(), sourceSig, destSig.Len(), destSig)
}

func typeMismatchError(source, dest ComponentID, argIndex int, sourceType, endType types.Type) error {
	return fmt.Errorf(`
As I infer the types of values flowing through your program, I see a mismatch in this connection.

	%[1]s -> %[2]s

The %[3]s result of %[1]s has type:

	%[4]s

But the %[3]s argument of %[2]s has type:

	%[5]s

HINT: These should have identical types.
`, source, dest, humanize.Ordinal(argIndex+1), sourceType, endType)
}
