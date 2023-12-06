package hwresp

/**
 * 新增操作时的返回
 */
type IdResp struct {
	Id int `json:"id"`
}

/**
 * 查询列表时的返回
 */
type LstResp struct {
	CurPage   int
	Limit     int
	Total     int
	PageTotal int
	Rows      []interface{}
}
