// Copyright 2023 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"context"
	"flag"
	"github.com/gardener/etcd-wrapper/internal/types"

	"github.com/gardener/etcd-wrapper/internal/app"
	"go.uber.org/zap"
)

var (
	EtcdCmd = Command{
		Name:      "start-etcd",
		ShortDesc: "Starts the etcd-wrapper application by initializing and starting an embedded etcd",
		UsageLine: "",
		LongDesc: `Initializes the etcd data directory by coordinating with a backup-sidecar container
and starts an embedded etcd.

Flags:
	--tls-enabled
		Enables TLS for the application (enabled by default)
	--sidecar-host-port string
		Host address and port of the backup restore sidecar with which this container will interact during initialization. Should be of the format <host>:<port> and must not include the protocol.
	--sidecar-ca-cert-bundle-path string
		Path of CA cert bundle (This will be used when TLS is enabled via tls-enabled flag.`,
		AddFlags: AddEtcdFlags,
		Run:      InitAndStartEtcd,
	}
	sidecarConfig = types.SidecarConfig{}
)

// AddEtcdFlags adds flags from the parsed flagset into application structs
func AddEtcdFlags(fs *flag.FlagSet) {
	fs.BoolVar(&sidecarConfig.TLSEnabled, "tls-enabled", types.DefaultTLSEnabled, "Enables TLS for the application")
	fs.StringVar(&sidecarConfig.HostPort, "sidecar-base-address", types.DefaultSideCarHostPort, "Base address of the backup restore sidecar")
	sidecarConfig.CaCertBundlePath = fs.String("sidecar-ca-cert-bundle-path", "", "File path of CA cert bundle") //TODO @aaronfern: define a reasonable default
}

// InitAndStartEtcd sets up and starts an embedded etcd
func InitAndStartEtcd(ctx context.Context, logger *zap.Logger) error {
	etcdApp, err := app.NewApplication(ctx, &sidecarConfig, logger)
	if err != nil {
		return err
	}
	if err := etcdApp.Setup(); err != nil {
		return err
	}
	if err := etcdApp.Start(); err != nil {
		return err
	}
	return nil
}