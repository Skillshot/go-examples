package server

//go:generate mockgen -destination=../mocks/logger_mock.go -package=mocks . Logger
type Logger interface {
	Info(message string)
	Error(errMsg string)
}
