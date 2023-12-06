package hwreq

/**
 * 查询单条或删除信息时的通用请求参数
 */
type IDReq struct {
	Id uint `p:"id" v:"required|min:1#请输入唯一ID|请输入唯一ID"`
}

/**
 * 查询列表时的通用请求参数
 */
type LstReq struct {
	Page   uint `p:"page" v:"required|min:1#请输入页码|请输入页码"`
	Limit  uint `p:"limit"`
	Filter []FilterReq
}

type FilterReq struct {
	Type  string //LIKE/LTE/LT/GT/GTE/EQ/NEQ
	Field string
	Value string
}
