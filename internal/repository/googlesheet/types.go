package googlesheet

type ValueRange struct {
	MajorDimension MajorDimension
	Range          string
	Values         [][]interface{}
}

type MajorDimension string

const (
	MajorDimensionRow    MajorDimension = "ROWS"
	MajorDimensionColumn MajorDimension = "COLUMNS"
)
