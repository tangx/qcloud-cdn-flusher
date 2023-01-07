package qcdn

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-jarvis/logr"
	cdn "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdn/v20180606"
)

func Do(ctx context.Context, flag *Flag) {
	log := logr.FromContext(ctx)
	log.Info("%+v", flag)

	client := mustClient()
	urls := mustParseURLs(ctx, flag)

	if len(flag.Include) != 0 {
		urls = append(urls, flag.Include...)
	}

	if !flag.Purge && !flag.Push {
		flag.Purge = true
		flag.Push = true
	}

	urls = exclude(urls, flag.Exclude...)

	flush(ctx, client, flag, urls...)
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

func printOutUrls(ctx context.Context, urls ...string) {
	log := logr.FromContext(ctx)
	for _, u := range urls {
		log.Debug(u)
	}
}

func flush(ctx context.Context, client *cdn.Client, flag *Flag, urls ...string) {

	log := logr.FromContext(ctx)

	do := func(targets []string) {
		fmt.Println("")
		printOutUrls(ctx, targets...)
		log.Info("本次处理: %d 个", len(targets))

		if flag.Purge {
			// 刷新 URL
			_ = PurgeSite(client, targets)
		}

		if flag.Push {
			// 预热 URL
			PushSite(client, targets)
		}
	}

	i := 0
	step := 80
	n := len(urls)
	for {
		time.Sleep(1 * time.Second)

		if (i+1)*step >= n-1 {
			ss := urls[i*step:]
			do(ss)

			break
		}

		ss := urls[i*step : (i+1)*step]
		do(ss)

		i += 1
	}
}
