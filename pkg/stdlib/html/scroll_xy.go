package html

import (
	"context"

	"github.com/MontFerret/ferret/pkg/drivers"
	"github.com/MontFerret/ferret/pkg/runtime/core"
	"github.com/MontFerret/ferret/pkg/runtime/values"
	"github.com/MontFerret/ferret/pkg/runtime/values/types"
)

// ScrollXY scrolls by given coordinates.
// @param doc (HTMLDocument) - HTML document.
// @param x (Int|Float) - X coordinate.
// @param y (Int|Float) - Y coordinate.
func ScrollXY(ctx context.Context, args ...core.Value) (core.Value, error) {
	err := core.ValidateArgs(args, 3, 3)

	if err != nil {
		return values.None, err
	}

	doc, err := drivers.ToDocument(args[0])

	if err != nil {
		return values.None, err
	}

	err = core.ValidateType(args[1], types.Int, types.Float)

	if err != nil {
		return values.None, err
	}

	err = core.ValidateType(args[2], types.Int, types.Float)

	if err != nil {
		return values.None, err
	}

	x := values.ToFloat(args[1])
	y := values.ToFloat(args[2])

	return values.None, doc.ScrollByXY(ctx, x, y)
}
