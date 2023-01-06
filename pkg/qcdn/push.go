package qcdn

import (
	"fmt"

	cdn "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdn/v20180606"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
)

// PushSite 预热网站
// https://cloud.tencent.com/document/api/228/37869
func PushSite(client *cdn.Client, urls []string) {

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := cdn.NewPushUrlsCacheRequest()
	request.Urls = common.StringPtrs(urls)

	// 返回的resp是一个PushUrlsCacheResponse的实例，与请求对象对应
	response, err := client.PushUrlsCache(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	// 输出json格式的字符串回包
	fmt.Println("预热 URL:", response.ToJsonString())
}
