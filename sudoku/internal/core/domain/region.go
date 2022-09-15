package domain

const (
	CELL_EMPTY = "_"
)

type Region [][]string

func NewEmptyRegion(dificulty string) Region {
	region := make([][]string, 3)

	for i := range region {
		region[i] = make([]string, 3)
	}

	for i := range region {
		for j := range region[0] {
			region[i][j] = CELL_EMPTY
		}
	}

	return region
}

func (region Region) ToString() string {
	regionString := ""

	for i := range region {
		for j := range region[i] {
			regionString += region[i][j] + " "
		}
		regionString += "| "
	}

	return regionString
}
