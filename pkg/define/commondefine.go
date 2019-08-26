package define

type LoginType int32

const (
	LoginType_MOBILE_LOGIN LoginType = 1 //登录方式 手机登录
	LoginType_QQ_LOGIN     LoginType = 2 //登录方式 QQ登录
	LoginType_WEIXIN_LOGIN LoginType = 3 //登录方式 微信登录
)

type SexType int32

const (
	SexType_Sex_male   SexType = 1 //女
	SexType_Sex_female SexType = 2 //男
)

type PhoneCodeType int32

const (
	PhoneCodeType_Register PhoneCodeType = 1 //这册获取验证码
	PhoneCodeType_Password PhoneCodeType = 2 //修改密码获取验证码
	PhoneCodeType_Bind     PhoneCodeType = 3 //绑定手机获取验证码
)

//发朋友圈类型
type CycleType int32

const (
	CycleType_Common CycleType = 1 //朋友圈文本图片
	CycleType_Audio  CycleType = 2 //朋友圈音频
	CycleType_Video  CycleType = 3 //朋友圈视频
)
