package tw

// 极简 stub 类型与常量，用于满足 gf 对 tablewriter/tw 的依赖。

// Rendition 表示一套字符绘制方案（这里是占位实现）。
type Rendition struct{}

// Settings 表示绘制配置（占位）。
type Settings struct{}

// Separators 表示表格分隔符（占位）。
type Separators struct{}

// On 是一个占位常量，用于开启某些特性。
const On = 1

// NewSymbolCustom 返回一个占位的 Rendition 实例。
func NewSymbolCustom(_ Settings, _ Separators) Rendition {
	return Rendition{}
}

