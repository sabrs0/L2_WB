package pattern

import (
	"fmt"
	"strings"
)

type WritingState interface {
	write(str string) string
}

type UpperCase struct {
	WritingState
}

func (UC UpperCase) write(str string) string {
	return strings.ToUpper(str)
}

type LowerCase struct {
	WritingState
}

func (LC LowerCase) write(str string) string {
	return strings.ToLower(str)
}

type Editor struct {
	ws WritingState
}

func (e Editor) typing(str string) {
	fmt.Println(e.ws.write(str))
}
func (e *Editor) setWritingState(s WritingState) {
	e.ws = s
}
func statePattern() {
	editor := Editor{}
	editor.setWritingState(UpperCase{})

	editor.typing("hello world")

	editor.setWritingState(LowerCase{})

	editor.typing("SKDJNijfvodkg")
}
