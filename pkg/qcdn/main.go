package qcdn

import (
	"context"
	"fmt"
	"strings"

	cdn "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdn/v20180606"
)

func Do(flag *Flag) {

	client := mustClient()
	urls := mustURLs(flag)

	if len(flag.Include) != 0 {
		urls = append(urls, flag.Include...)
	}

	if !flag.Purge && !flag.Push {
		flag.Purge = true
		flag.Push = true
	}

	urls = exclude(urls, "posts")

	fmt.Println(urls)
	if flag.Purge {
		// 刷新 URL
		_ = PurgeSite(client, urls)
	}

	if flag.Push {
		// 预热 URL
		PushSite(client, urls)
	}
}

func isPurgeOK(ctx context.Context, client *cdn.Client, purgeId string) bool {
	select {
	case <-ctx.Done():
		// timeout, treat as OK
		return true
	default:
		ok := IsPurgeTaskDone(client, purgeId)
		if ok {
			return true
		}
		return false
	}
}

func exclude(urls []string, list ...string) []string {
	if len(list) == 0 {
		return urls
	}

	result := []string{}
	for _, u := range urls {
		for _, s := range list {
			if strings.Contains(u, s) {
				continue
			}
			result = append(result, u)
		}

	}

	return result
}
