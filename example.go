package main

import "log/slog"

func main() {
	slog.Info("Starting server")
	slog.Error("Ошибка")
	slog.Warn("failed!")
}