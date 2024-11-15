package evaluation

import (
	"context"
	"fmt"

	"github.com/kakky/refacgo/internal/application"
	loadfile "github.com/kakky/refacgo/pkg/load_file"
)

type EvaluationWithGenAI struct {
	genAI application.GenAI
}

func NewEvaluationWithGenAI(genAI application.GenAI) *EvaluationWithGenAI {
	return &EvaluationWithGenAI{
		genAI: genAI,
	}
}

func (ev *EvaluationWithGenAI) Evaluate(ctx context.Context, src []byte, filename string, ch chan<- string) error {
	instruction, err := loadfile.LoadInternal("./instruction_text/genai_instruction.txt")
	if err != nil {
		panic(err)
	}
	prompt := fmt.Sprintf("The name of this file is %q.\n\n%v\n\n", filename, string(instruction))
	go func() error {
		err = ev.genAI.Query(ctx, src, prompt, ch)
		if err != nil {
			return err
		}
		return nil
	}()
	return nil
}
