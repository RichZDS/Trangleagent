package tablewriter

// 这是对上游 tablewriter 的一个极简 shim，只实现 gf 在开发工具中用到的
// 少量符号，实际运行时不会影响你的 HTTP API 逻辑。

type Table struct{}

// Option 是配置 Table 的函数式选项类型。
type Option func(*Table)

// NewTable 创建一个空表实例。
func NewTable(_ interface{}, _ ...Option) *Table {
	return &Table{}
}

// WithRenderer 是一个空实现的配置项，占位以满足 gf 的调用。
func WithRenderer(_ interface{}) Option {
	return func(t *Table) {}
}

