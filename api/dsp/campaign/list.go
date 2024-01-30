package campaign

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dsp/campaign"
)

// List 获取广告计划信息
func List(clt *core.SDKClient, accessToken string, req *campaign.ListRequest) (*campaign.ListResponse, error) {
	var resp campaign.ListResponse
	if err := clt.Post(accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
