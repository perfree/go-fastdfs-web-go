package commons

// Pager 分页结果集
type Pager struct {
	State int
	Msg   string
	Total int64
	Data  interface{}
}
