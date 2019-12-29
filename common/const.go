package common

const (
	ZERO = iota
	ONE
	TWO
)

const (
	_         = iota
	UNDELETED // 未删除
	DELETED   // 已删除
)

// business_kind
/**
1.系统操作;2.错误日志;3.用户登录;4.全站浏览;5.文章浏览;6.文章赞;7.文章踩;8.文章评论;9.评论赞;10.评论踩;11.评论回复;12.回复赞;13.回复踩
*/
const (
	_ = iota
	BUSINESS_SYSTEM_OP
	BUSINESS_ERROR_LOG
	BUSINESS_USER_LOGIN
	BUSINESS_GLOBAL_VISIT
	BUSINESS_ARTICLE_VISIT
	BUSINESS_ARTICLE_LIKE
	BUSINESS_ARTICLE_UNLIKE

	BUSINESS_ARTICLE_COMMENT
	BUSINESS_COMMENT_LIKE
	BUSINESS_COMMENT_UNLIKE

	BUSINESS_COMMENT_REPLY
	BUSINESS_REPLY_LIKE
	BUSINESS_REPLY_UNLIKE
)

// article kind
const (
	ARTICLE_NORMAL = iota + 1
	ARTICLE_ABOUT
	ARTICLE_NOTE
)

const (
	ADMIN = "admin"
	COLOR = "#000000"
	ITV   = 3600 * 24 // 有效时间
)

// quanta key
const (
	KEY_1 = "admin.pass"
	KEY_2 = "cmnt.open"
	KEY_3 = "cmnt.check"
)

var business = map[int]string{
	0:                       "未知",
	BUSINESS_SYSTEM_OP:      "系统操作",
	BUSINESS_ERROR_LOG:      "错误日志",
	BUSINESS_USER_LOGIN:     "用户登录",
	BUSINESS_GLOBAL_VISIT:   "全站浏览",
	BUSINESS_ARTICLE_VISIT:  "文章浏览",
	BUSINESS_ARTICLE_LIKE:   "文章踩",
	BUSINESS_ARTICLE_UNLIKE: "文章评论",

	BUSINESS_ARTICLE_COMMENT: "评论赞",
	BUSINESS_COMMENT_LIKE:    "评论踩",
	BUSINESS_COMMENT_UNLIKE:  "评论回复",

	BUSINESS_COMMENT_REPLY: "评论回复",
	BUSINESS_REPLY_LIKE:    "回复赞",
	BUSINESS_REPLY_UNLIKE:  "回复踩",
}
