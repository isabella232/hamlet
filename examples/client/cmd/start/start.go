// Copyright 2019 VMware, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package start

import (
	"github.com/spf13/cobra"
	"github.com/vmware/hamlet/examples/client/pkg/lifecycle"
)

// flagSet represents the flags available with the start subcommand.
type flagSet struct {
	RootCACert         string
	PeerCert           string
	PeerKey            string
	ServerAddr         string
	InsecureSkipVerify bool
}

// NewCommand returns a new Command instance for the start subcommand.
func NewCommand() *cobra.Command {
	flags := &flagSet{}
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "start",
		Short: "Start the client",
		Long:  "Start the client",
		Run: func(cmd *cobra.Command, args []string) {
			lifecycle.Start(flags.RootCACert, flags.PeerCert,
				flags.PeerKey, flags.ServerAddr,
				flags.InsecureSkipVerify)
		},
	}

	cmd.Flags().StringVar(&flags.RootCACert, "root-ca-cert", "", "the root CA certificate path")
	cmd.Flags().StringVar(&flags.PeerCert, "peer-cert", "", "the peer certificate path")
	cmd.Flags().StringVar(&flags.PeerKey, "peer-key", "", "the peer key path")
	cmd.Flags().StringVar(&flags.ServerAddr, "server-addr", "localhost:8000", "the server's address")
	cmd.Flags().BoolVar(&flags.InsecureSkipVerify, "insecure-skip-verify", false, "whether normal verification should be ignored")

	cmd.MarkFlagRequired("root-ca-cert")
	cmd.MarkFlagRequired("peer-cert")
	cmd.MarkFlagRequired("peer-key")
	return cmd
}
