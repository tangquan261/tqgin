package filter

import "tqgin/pkg/filter/sensitive"

var filter *sensitive.Filter

func init() {
	filter = sensitive.New()
	filter.LoadWordDict("./conf/dict.txt")
}

// Filter 过滤敏感词
func Filter(text string) string {
	return filter.Filter(text)
}

// Replace 和谐敏感词
func Replace(text string, repl rune) string {
	return filter.Replace(text, repl)
}

// FindIn 检测敏感词
func FindIn(text string) (bool, string) {
	return filter.FindIn(text)
}

// FindAll 找到所有匹配词
func FindAll(text string) []string {
	return filter.FindAll(text)
}

// Validate 检测字符串是否合法
func Validate(text string) (bool, string) {
	return filter.Validate(text)
}

// RemoveNoise 去除空格等噪音
func RemoveNoise(text string) string {
	return filter.RemoveNoise(text)
}
