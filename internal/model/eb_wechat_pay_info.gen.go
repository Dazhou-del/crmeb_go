// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameEbWechatPayInfo = "eb_wechat_pay_info"

// EbWechatPayInfo 微信订单表
type EbWechatPayInfo struct {
	ID             int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	AppID          string `gorm:"column:app_id;comment:公众号唯一标识" json:"app_id"`                                                                                       // 公众号唯一标识
	MchID          string `gorm:"column:mch_id;comment:商户号" json:"mch_id"`                                                                                           // 商户号
	DeviceInfo     string `gorm:"column:device_info;comment:设备号,PC网页或公众号内支付可以传-WEB" json:"device_info"`                                                              // 设备号,PC网页或公众号内支付可以传-WEB
	OpenID         string `gorm:"column:open_id;comment:用户的唯一标识" json:"open_id"`                                                                                     // 用户的唯一标识
	NonceStr       string `gorm:"column:nonce_str;comment:随机字符串" json:"nonce_str"`                                                                                   // 随机字符串
	Sign           string `gorm:"column:sign;comment:签名" json:"sign"`                                                                                                // 签名
	SignType       string `gorm:"column:sign_type;default:MD5;comment:签名类型，默认为MD5，支持HMAC-SHA256和MD5" json:"sign_type"`                                               // 签名类型，默认为MD5，支持HMAC-SHA256和MD5
	Body           string `gorm:"column:body;comment:商品描述" json:"body"`                                                                                              // 商品描述
	Detail         string `gorm:"column:detail;comment:商品详细描述，对于使用单品优惠的商户，该字段必须按照规范上传" json:"detail"`                                                                // 商品详细描述，对于使用单品优惠的商户，该字段必须按照规范上传
	Attach         string `gorm:"column:attach;comment:附加数据，在查询API和支付通知中原样返回，可作为自定义参数使用" json:"attach"`                                                              // 附加数据，在查询API和支付通知中原样返回，可作为自定义参数使用
	OutTradeNo     string `gorm:"column:out_trade_no;comment:商户订单号,要求32个字符内" json:"out_trade_no"`                                                                    // 商户订单号,要求32个字符内
	FeeType        string `gorm:"column:fee_type;comment:标价币种：CNY：人民币 GBP：英镑 HKD：港币 USD：美元 JPY：日元 CAD：加拿大元 AUD：澳大利亚元 EUR：欧元 NZD：新西兰元 KRW：韩元 THB：泰铢" json:"fee_type"` // 标价币种：CNY：人民币 GBP：英镑 HKD：港币 USD：美元 JPY：日元 CAD：加拿大元 AUD：澳大利亚元 EUR：欧元 NZD：新西兰元 KRW：韩元 THB：泰铢
	TotalFee       int32  `gorm:"column:total_fee;comment:标价金额" json:"total_fee"`                                                                                    // 标价金额
	SpbillCreateIP string `gorm:"column:spbill_create_ip;comment:终端IP" json:"spbill_create_ip"`                                                                      // 终端IP
	TimeStart      string `gorm:"column:time_start;comment:交易起始时间" json:"time_start"`                                                                                // 交易起始时间
	TimeExpire     string `gorm:"column:time_expire;comment:交易结束时间" json:"time_expire"`                                                                              // 交易结束时间
	NotifyURL      string `gorm:"column:notify_url;comment:通知地址" json:"notify_url"`                                                                                  // 通知地址
	TradeType      string `gorm:"column:trade_type;comment:交易类型,取值为：JSAPI，NATIVE，APP等" json:"trade_type"`                                                            // 交易类型,取值为：JSAPI，NATIVE，APP等
	ProductID      string `gorm:"column:product_id;comment:商品ID" json:"product_id"`                                                                                  // 商品ID
	SceneInfo      string `gorm:"column:scene_info;comment:场景信息" json:"scene_info"`                                                                                  // 场景信息
	ErrCode        string `gorm:"column:err_code;comment:错误代码" json:"err_code"`                                                                                      // 错误代码
	PrepayID       string `gorm:"column:prepay_id;comment:预支付交易会话标识" json:"prepay_id"`                                                                               // 预支付交易会话标识
	CodeURL        string `gorm:"column:code_url;comment:二维码链接" json:"code_url"`                                                                                     // 二维码链接
	IsSubscribe    string `gorm:"column:is_subscribe;comment:是否关注公众账号" json:"is_subscribe"`                                                                          // 是否关注公众账号
	TradeState     string `gorm:"column:trade_state;comment:交易状态" json:"trade_state"`                                                                                // 交易状态
	BankType       string `gorm:"column:bank_type;comment:付款银行" json:"bank_type"`                                                                                    // 付款银行
	CashFee        int32  `gorm:"column:cash_fee;comment:现金支付金额" json:"cash_fee"`                                                                                    // 现金支付金额
	CouponFee      int32  `gorm:"column:coupon_fee;comment:代金券金额" json:"coupon_fee"`                                                                                 // 代金券金额
	TransactionID  string `gorm:"column:transaction_id;comment:微信支付订单号" json:"transaction_id"`                                                                       // 微信支付订单号
	TimeEnd        string `gorm:"column:time_end;comment:支付完成时间" json:"time_end"`                                                                                    // 支付完成时间
	TradeStateDesc string `gorm:"column:trade_state_desc;comment:交易状态描述" json:"trade_state_desc"`                                                                    // 交易状态描述
}

// TableName EbWechatPayInfo's table name
func (*EbWechatPayInfo) TableName() string {
	return TableNameEbWechatPayInfo
}
