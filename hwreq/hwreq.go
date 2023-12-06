package hwreq

/**
 * 当前登录用户信息
 */
type CurUser struct {
	Id      uint   // 用户ID
	Phone   string // 手机号
	Role    uint   // 0-门店，1-商户，2-运营端
	ShopId  uint   // 商铺id
	StoreId uint   // 门店id
}

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
