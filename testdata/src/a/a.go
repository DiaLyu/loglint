package a 

import ( 
	"log/slog" 
	"go.uber.org/zap" 
) 

func main() { 
	// Правило 1: Строчная буква
	slog.Info("Starting")
	
	// Правило 2: Только английский
	slog.Error("ошибка")
	
	// Правило 3: Спецсимволы и эмодзи
	slog.Warn("failed!")
	
	// Правило 4: Чувствительные данные 
	password := "123" 
	slog.Info("user password: " + password)
	
	// Правильный лог 
	slog.Info("server started") 
}