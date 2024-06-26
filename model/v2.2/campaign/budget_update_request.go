package campaign

import "encoding/json"

type BudgetUpdateRequest struct {
	//advertiser_id	long	必填	广告主 ID	在获取 access_token 的时候返回
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	//unit_ids	long	必填	广告计划 ID
	CampaignId uint64 `json:"campaign_id"`
	//广告组 bid_type 为 CPC 和 eCPC 时：不得低于 0.2 元，不得高于 100 元，单位：厘；广告组 bid_type 为 OCPC 时：行为出价不得低于 1 元；激活出价不得低于 5 元（白名单用户不得低于 2 元），单位：厘
	DayBudget uint64 `json:"day_budget"`
}

// Url implement PostRequest interface
func (r BudgetUpdateRequest) Url() string {
	return "gw/dsp/campaign/update"
}

// Encode implement PostRequest interface
func (r BudgetUpdateRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
