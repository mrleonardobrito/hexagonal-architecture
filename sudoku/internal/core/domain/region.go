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

func (region Region) RenderLine(n_line int) string {
	lineString := ""
	lineString += region[n_line][0] + " " + region[n_line][1] + " " + region[n_line][2] + " | "

	return lineString
}
