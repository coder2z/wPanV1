package R

const (
	SUCCESSMSG       = "success"
	FAILMSG          = "fail"
	POLICY_ERROR     = "Policy不存在"
	POLICY_ADD_OK    = "添加成功"
	POLICY_ADD_ERROR = "Policy存在"
	MSG422           = "参数错误"
	MSG401           = "权限不足"
	AUTH_ERROR       = "未登录"
	PASSWORD_T       = "二次密码错误"
	EMAIL_CODE       = "email验证码错误"
	REG_USER_EXIST   = "用户名被注册"
	REG_BCRYPT_ERROR = "加密失败"
	REG_ERROR        = "注册失败"
	REG_OK           = "注册成功"
	SENDCODE_OK      = "发送验证码成功"
	SENDCODE_ERROR   = "发送验证码失败"
	SENDCODE_EXISTS  = "发送频繁，5分钟后再试"
)

const (
	SUCCESS = 1
	FAIL    = 0
)
