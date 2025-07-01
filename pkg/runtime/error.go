package runtime

import (
	"fmt"
	"strings"
)

type MultiError []error

func (me MultiError) Error() string {
	var sb strings.Builder
	sb.WriteString("multiple errors occurred:\n")
	for i, err := range me {
		sb.WriteString(fmt.Sprintf("\t%d: %s\n", i+1, err.Error()))
	}
	return sb.String()
}
