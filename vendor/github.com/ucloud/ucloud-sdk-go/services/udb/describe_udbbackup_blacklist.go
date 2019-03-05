//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UDB DescribeUDBBackupBlacklist

package udb

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DescribeUDBBackupBlacklistRequest is request schema for DescribeUDBBackupBlacklist action
type DescribeUDBBackupBlacklistRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"false"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// DB实例Id,该值可以通过DescribeUDBInstance获取
	DBId *string `required:"true"`
}

// DescribeUDBBackupBlacklistResponse is response schema for DescribeUDBBackupBlacklist action
type DescribeUDBBackupBlacklistResponse struct {
	response.CommonBase

	// DB的黑名单列表, db.%为指定库 dbname.tablename为指定表
	Blacklist string
}

// NewDescribeUDBBackupBlacklistRequest will create request of DescribeUDBBackupBlacklist action.
func (c *UDBClient) NewDescribeUDBBackupBlacklistRequest() *DescribeUDBBackupBlacklistRequest {
	req := &DescribeUDBBackupBlacklistRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DescribeUDBBackupBlacklist - 获取UDB实例的备份黑名单
func (c *UDBClient) DescribeUDBBackupBlacklist(req *DescribeUDBBackupBlacklistRequest) (*DescribeUDBBackupBlacklistResponse, error) {
	var err error
	var res DescribeUDBBackupBlacklistResponse

	err = c.client.InvokeAction("DescribeUDBBackupBlacklist", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}