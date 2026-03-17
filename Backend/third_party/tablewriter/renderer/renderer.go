package renderer

// 极简 renderer 占位实现，用于满足 gf 对 tablewriter/renderer 的依赖。

type Blueprint struct{}

// NewBlueprint 返回一个空的 Blueprint，占位使用。
func NewBlueprint() *Blueprint {
	return &Blueprint{}
}

