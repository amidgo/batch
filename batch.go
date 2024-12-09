package batch

import (
	"fmt"
	"strings"
)

func Query(template string, rowsCount, argsPerRow uint64) string {
	if rowsCount == 0 || argsPerRow == 0 {
		return "()"
	}

	buf := strings.Builder{}

	for i := range rowsCount {
		buf.WriteString("(")
		index := i * argsPerRow

		for i := range argsPerRow {
			buf.WriteString(fmt.Sprintf("$%d", index+i+1))

			if i < argsPerRow-1 {
				buf.WriteByte(',')
			}
		}

		buf.WriteString(")")

		if i < rowsCount-1 {
			buf.WriteByte(',')
		}
	}

	return fmt.Sprintf(template, buf.String())
}
