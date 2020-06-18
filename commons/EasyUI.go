package commons

type Datagrid struct {
	//当前页显示的数据
	Rows interface{} `json:"rows"`
	//总个数
	Total int `json:"total"`
}
