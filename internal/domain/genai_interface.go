package domain

import "context"

//go:generate mockgen -package=domain -source=./genai_interface.go -destination=./genai_mock.go
type GenAI interface {
	StreamQueryResults(ctx context.Context, src []byte, prompt string, ch chan<- string) error
	QueryResuluts(ctx context.Context, src []byte, prompt string, ch chan<- string) error
}
