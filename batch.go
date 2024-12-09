package batch

import (
	"fmt"
	"strings"
)

func Query(template string, rowsCount, argsPerRow uint64) string {
	if rowsCount == 0 || argsPerRow == 0 {
		return fmt.Sprintf(template, "()")
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

type MapFunc func(argNumber uint64) string

func QueryMap(template string, rowsCount, argsPerRow uint64, argMapper map[uint64]MapFunc) string {
	if rowsCount == 0 || argsPerRow == 0 {
		return fmt.Sprintf(template, "()")
	}

	if argMapper == nil {
		argMapper = make(map[uint64]MapFunc)
	}

	buf := strings.Builder{}

	for i := range rowsCount {
		buf.WriteString("(")
		index := i * argsPerRow

		for i := range argsPerRow {
			argNumber := index + i + 1
			switch f := argMapper[i]; f {
			case nil:
				buf.WriteString(fmt.Sprintf("$%d", argNumber))
			default:
				buf.WriteString(f(argNumber))
			}

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
