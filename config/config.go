package config

const (
	CLIENT_ID     string = ""                                                                                                                                  // 申请的API Key
	CLIENT_SECRET string = ""                                                                                                                                  // 申请的Secret Key
	OAUTH         string = "https://aip.baidubce.com/oauth/2.0/token?grant_type=client_credentials&client_id=" + CLIENT_ID + "&client_secret=" + CLIENT_SECRET // 百度api获取 access_token的地址
	APIPATH       string = "https://aip.baidubce.com/rest/2.0/ocr/v1/accurate_basic"                                                                           // 百度 api 调用的地址
	VERIFYIMGPATH string = "./verifycode.png"                                                                                                                  // 验证码的保存路径

	APPID            string = ""                                                                                                             // 微信 appid
	APPSECRET        string = ""                                                                                                             // 微信 app secret
	WECHATID         string = ""                                                                                                             // 关注工作号获取的 微信id
	TEMPLATEID       string = ""                                                                                                             // 模板 id
	WECHATTOKENURL   string = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + APPID + "&secret=" + APPSECRET // 获取微信 access_token 的地址
	WECHATSENDMSGURL string = "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token="                                        // 调用模板的地址

	HOST           string = "https://yqfkapi.zhxy.net"
	LOGIN          string = HOST + "/api/User/CheckUser"
	CheckIn        string = HOST + "/api/ClockIn/IsClockIn"
	SAVEPATH       string = HOST + "/api/ClockIn/Save"
	VERIFYCODEPATH string = HOST + "/api/common/getverifycode"

	FILENAME string = "./web/uidList" // uidList 的路径

	A3 = "四川省成都市郫都区西华大学第六教学楼" // 打卡地点
)
