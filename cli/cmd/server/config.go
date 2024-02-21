// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package cmd_server

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/daytonaio/daytona/cli/cmd/output"
	"github.com/daytonaio/daytona/server/config"
	"github.com/daytonaio/daytona/server/frpc"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Output local Daytona Server config",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := config.GetConfig()
		if err != nil {
			log.Fatal(err)
		}

		apiUrl := frpc.GetApiUrl(config)
		output.Output = apiUrl

		fmt.Println(apiUrl)
	},
}
