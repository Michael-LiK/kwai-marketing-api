package adcompass

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/adcompass"
)

// QuotaTending 磁力罗盘对外 quota 腾挪接口
func QuotaTending(clt *core.SDKClient, accessToken string, req *adcompass.QuotaTendingRequest) (string, error) {
	var resp adcompass.QuotaTendingResponse
	if err := clt.Post(accessToken, req, &resp); err != nil {
		return "", err
	}
	return resp.Describe, nil
}
