package qcdn

import (
	"fmt"
	"strings"

	cdn "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdn/v20180606"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
)

type PurgeResp struct {
	Response struct {
		TaskId    string `json:"TaskId"`
		RequestId string `json:"RequestId"`
	} `json:"Response"`
}

// PurgeSite 刷新缓存， 强制缓存过气
// https://cloud.tencent.com/document/api/228/37870
func PurgeSite(client *cdn.Client, urls []string) string {

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := cdn.NewPurgeUrlsCacheRequest()
	request.Urls = common.StringPtrs(urls)
	// 返回的resp是一个PurgeUrlsCacheResponse的实例，与请求对象对应
	response, err := client.PurgeUrlsCache(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return ""
	}
	if err != nil {
		panic(err)
	}

	// 输出json格式的字符串回包
	fmt.Println("刷新 URL:", response.ToJsonString())

	taskid := response.Response.TaskId
	return *taskid
}

func IsPurgeTaskDone(client *cdn.Client, id string) bool {
	request := cdn.NewDescribePurgeTasksRequest()
	request.TaskId = &id

	response, err := client.DescribePurgeTasks(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return false
	}
	if err != nil {
		panic(err)
	}
	// 输出json格式的字符串回包
	fmt.Printf("查询刷新结果: %s\n", response.ToJsonString())

	// response.Response.PurgeLogs
	if len(response.Response.PurgeLogs) == 0 {
		return false
	}

	for _, task := range response.Response.PurgeLogs {
		if *task.TaskId != id {
			continue
		}

		status := *task.Status
		if strings.ToLower(status) == "done" {
			return true
		} else {
			return false
		}
	}

	return false
}
