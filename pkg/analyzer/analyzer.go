package analyzer 

import ( 
	"go/ast" 
	"regexp" 
	"strings" 
	"unicode" 
	"golang.org/x/tools/go/analysis" 
) 

var ( 
	reEnglish = regexp.MustCompile(`^[a-z0-9\s\-_.,()]+$`) 
	
	// Из списка бонусных заданий: кастомные паттерны 
	sensitiveWords = []string{"password", "token", "api_key", "secret"} 
) 

var Analyzer = &analysis.Analyzer{ 
	Name: "loglint", 
	Doc: "checks log messages for style and security", 
	Run: run, 
} 

func run(pass *analysis.Pass) (interface{}, error) { 
	for _, f := range pass.Files { 
		ast.Inspect(f, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr) 
			if !ok { 
				return true 
			} 
			
			// проверяем, что это вызов логгера (slog или zap)
			if !isLogCall(call) { 
				return true 
			} 
			
			// обычно сообщение это первый аргумент
			if len(call.Args) == 0 { 
				return true 
			} 
			
			// пытаемся получить строковый литерал сообщения 
			arg := call.Args[0] 
			lit, ok := arg.(*ast.BasicLit) 
			
			if !ok || lit.Kind != 9 { // 9 = STRING 
				// проверка конкатенации на секреты 
				checkSensitiveInExpr(pass, arg) 
				return true 
			} 
			
			msg := strings.Trim(lit.Value, "`\"") 
			
			// 1. Проверка на строчную букву 
			if len(msg) > 0 && unicode.IsUpper(rune(msg[0])) { 
				pass.Report(analysis.Diagnostic{ 
					Pos: lit.Pos(), 
					Message: "log message should start with a lowercase letter", 
					// Из списка бонусных заданий: авто-исправление 
					SuggestedFixes: []analysis.SuggestedFix{{
						Message: "lower case", 
						TextEdits: []analysis.TextEdit{{ 
							Pos: lit.Pos() + 1, 
							End: lit.Pos() + 2, 
							NewText: []byte(strings.ToLower(string(msg[0]))), 
						}}, 
					}}, 
				}) 
			} 
			
			// 2. Проверка на английский язык 
			if !reEnglish.MatchString(strings.ToLower(msg)) { 
				pass.Reportf(lit.Pos(), "log message should be in English") 
			} 
			
			// 3. Проверка на спецсимволы/эмодзи 
			if strings.ContainsAny(msg, "!?") { 
				pass.Reportf(lit.Pos(), "log message should not contain special characters or emojis") 
			} 
			
			return true 
		}) 
	} 
	
	return nil, nil 
} 

func isLogCall(call *ast.CallExpr) bool { 
	// простоты проверяем имя селектора (Info, Error, Warn) 
	sel, ok := call.Fun.(*ast.SelectorExpr) 
	if !ok { 
		return false 
	} 
	
	methods := map[string]bool{"Info": true, "Error": true, "Warn": true, "Debug": true} 
	return methods[sel.Sel.Name] 
} 

func checkSensitiveInExpr(pass *analysis.Pass, expr ast.Expr) { 
	// рукурсивно обходим выражения на поиск ключевых слов
	ast.Inspect(expr, func(n ast.Node) bool { 
		if lit, ok := n.(*ast.BasicLit); ok { 
			val := strings.ToLower(lit.Value) 
			for _, word := range sensitiveWords { 
				if strings.Contains(val, word) { 
					pass.Reportf(lit.Pos(), "log message contains sensitive data: %s", word) 
				} 
			} 
		} 
		return true 
	}) 
}