package transfer

var Jdpay =map[string]interface{}{
	// 必须，1~32位，收款人银行卡号或者存折号。
    "card_number": "6228480402564890011",
    // 必须，1~100位，收款人姓名。
    "user_name": "张三",
    // 必须，4位，开户银行编号，详情请参考 企业付款（银行卡）银行编号说明：https://www.pingxx.com/api#%E9%93%B6%E8%A1%8C%E7%BC%96%E5%8F%B7%E8%AF%B4%E6%98%8E。
    "open_bank_code": "0103",
}