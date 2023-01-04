package cmd

import (
	"context"
	"os"

	"github.com/go-jarvis/cobrautils"
	"github.com/spf13/cobra"
	"github.com/tangx/qcloud-cdn-flusher/pkg/qcdn"
	"gopkg.in/yaml.v3"
)

var rootCmd = &cobra.Command{
	Use:   "flusher",
	Short: "刷新并预热",
	Run: func(cmd *cobra.Command, args []string) {
		execute()
	},
}

var flag = &qcdn.Flag{
	SiteMap:       "./docs/sitemap.xml",
	LastModInDays: 7,
}

func init() {
	cobrautils.BindFlags(rootCmd, flag)
}

func Execute() error {
	return rootCmd.Execute()
}

func execute() {
	if flag.Config != "" {
		data, err := os.ReadFile(flag.Config)
		if err != nil {
			panic(err)
		}

		err = yaml.Unmarshal(data, flag)
		if err != nil {
			panic(err)
		}
	}

	ctx := context.Background()
	qcdn.Do(ctx, flag)
}
