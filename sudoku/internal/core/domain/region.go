package domain

const (
	CELL_EMPTY = "_"
)

type Region []string

func NewEmptyRegion(dificulty string) Region {
	region := make([]string, 9)

	for i := range region {
		region[i] = CELL_EMPTY
	}

	return region
}

func (region Region) RenderLine(n_line int) string {
	lineString := ""

	pos := 3 * n_line
	lineString += region[pos] + " " + region[pos+1] + " " + region[pos+2] + " | "

	return lineString
}
