package dmp

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dmp"
)

// PopulationUpload 人群包上传接口
func PopulationUploadV2(ctx context.Context, clt *core.SDKClient, accessToken string, req *dmp.PopulationUploadRequestV2) (*dmp.PopulationV2, error) {
	var resp dmp.PopulationV2
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
