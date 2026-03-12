package a 

import ( 
	"log/slog" 
	"go.uber.org/zap" 
) 

func main() { 
	// Правило 1: Строчная буква
	slog.Info("Starting") // want "log message should start with a lowercase letter"
	
	// Правило 2: Только английский
	slog.Error("ошибка") // want "log message should start with a lowercase letter" "log message should be in English"
	
	// Правило 3: Спецсимволы и эмодзи
	slog.Warn("failed!") // want "log message should be in English" "log message should not contain special characters or emojis"
	
	// Правило 4: Чувствительные данные 
	password := "123" 
	slog.Info("user password: " + password) // want "log message contains sensitive data: password"

	// zap пример
	zap.L().Info("Starting zap") // want "log message should start with a lowercase letter"
	
	// Правильный лог 
	slog.Info("server started") 
}