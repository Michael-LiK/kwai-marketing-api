package app

type ListRequest struct {
	//advertiser_id	Long		必填	广告主id
	AdvertiserId int64 `json:"advertiser_id"`
	//list_type	Integer		可选	列表类型	不传-全部, 1-我创建的, 2-共享给我的
	ListType int `json:"list_type"`
	//platform	String	"ios"	可选	android或ios
	Platform string `json:"platform"`
	//app_ids	Long[]		可选	批量应用id查询	最多支持查询100个
	AppIds []int64 `json:"app_ids"`
	//key_word	String		可选	关键词	支持应用ID或应用名称搜索
	KeyWord string `json:"key_word"`
	//page	Integer		可选	当前页	页码，默认1
	Page int `json:"page"`
	//page_size	Integer		可选	分页大小	个数，默认10
	PageSize int `json:"page_size"`
}

func (r ListRequest) Url() string {
	return "gw/dsp/appcenter/app/release/list"
}

// Encode implement PostRequest interface
func (r ListRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}