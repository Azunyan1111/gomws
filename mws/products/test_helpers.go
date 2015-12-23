package products

import (
	"strconv"

	"github.com/smartystreets/goconvey/convey"
)

func decimalWithUnitsAsserter(dwn DecimalWithUnits, unit string, value float64) {
	convey.Convey("Units is "+unit, func() {
		convey.So(dwn.Units, convey.ShouldEqual, unit)
	})

	convey.Convey("Value is "+strconv.FormatFloat(value, 'f', 2, 64), func() {
		convey.So(dwn.Value, convey.ShouldEqual, value)
	})
}

func dimensionsAsserter(dim DimensionType, lengthUnit, weightUnit string, expectValue map[string]float64) {
	convey.Convey("Height", func() {
		decimalWithUnitsAsserter(dim.Height, lengthUnit, expectValue["Height"])
	})

	convey.Convey("Length", func() {
		decimalWithUnitsAsserter(dim.Length, lengthUnit, expectValue["Length"])
	})

	convey.Convey("Width", func() {
		decimalWithUnitsAsserter(dim.Width, lengthUnit, expectValue["Width"])
	})

	convey.Convey("Weight", func() {
		decimalWithUnitsAsserter(dim.Weight, weightUnit, expectValue["Weight"])
	})
}
