package nextmail

import (
	"math"
	"math/big"
	"strings"

	"nextmail.sh/internal/protocol"
)

type ImageCrop struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

var imageCropUnits = new(big.Int).Exp(big.NewInt(10), big.NewInt(protocol.ImageCropDecimals), nil)

func (c ImageCrop) text() string {
	sides := []float64{c.X, c.Y, c.Width, c.Height}
	texts := make([]string, len(sides))

	for index, side := range sides {
		texts[index] = imageCropSide(side)
	}

	return strings.Join(texts, protocol.ImageCropSeparator)
}

func imageCropSide(value float64) string {
	size := math.Abs(value)

	if math.IsNaN(value) || size >= fixedNotationCeil {
		return jsNumber(value)
	}

	exact := new(big.Rat).SetFloat64(size)
	units := new(big.Int).Mul(exact.Num(), imageCropUnits)
	units.Lsh(units, 1).Add(units, exact.Denom())
	units.Quo(units, new(big.Int).Lsh(exact.Denom(), 1))
	rounded, _ := new(big.Rat).SetFrac(units, imageCropUnits).Float64()

	return jsNumber(math.Copysign(rounded, value))
}
