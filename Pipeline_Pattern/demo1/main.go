package main

import (
	"fmt"
	"strings"

	"github.com/izniburak/pipeline-go"
)

type UpperCasePipe struct{}

func (u *UpperCasePipe) Handle(value pipeline.PipeValue, next pipeline.PipeNext) pipeline.PipeValue {
	// get value
	text := value.(string)
	capitalized := strings.ToUpper(text)
	return next(capitalized)
}

type TrimSpacePipe struct{}

func (t *TrimSpacePipe) Handle(value pipeline.PipeValue, next pipeline.PipeNext) pipeline.PipeValue {
	// get value
	text := value.(string)
	trimmed := strings.Trim(text, " ")
	return next(trimmed)
}

func main() {
	text := "   buki.dev   "

	pipes := []pipeline.PipeInterface{
		new(UpperCasePipe),
		new(TrimSpacePipe),
	}
	result := pipeline.Send(text).Through(pipes).ThenReturn()

	fmt.Println(result) // BUKI.DEV
}
