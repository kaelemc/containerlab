// Copyright 2022 Nokia
// Licensed under the BSD 3-Clause License.
// SPDX-License-Identifier: BSD-3-Clause

package cvsr

import (
	"context"
	_ "embed"
	"fmt"
	"path/filepath"
	"regexp"

	"github.com/charmbracelet/log"
	"github.com/srl-labs/containerlab/nodes"
	"github.com/srl-labs/containerlab/types"
	"github.com/srl-labs/containerlab/utils"
)

var (
	kindNames          = []string{"nokia_cvsr"}
	defaultCredentials = nodes.NewCredentials("admin", "admin")
	// cvsrEnv             = map[string]string{
	// 	"XR_FIRST_BOOT_CONFIG": "/etc/cvsr/first-boot.cfg",
	// 	"XR_MGMT_INTERFACES":   "linux:eth0,xr_name=Mg0/RP0/CPU0/0,chksum,snoop_v4,snoop_v6",
	// 	"XR_EVERY_BOOT_SCRIPT": "/etc/cvsr/mgmt_intf_v6_addr.sh",
	// }
)

const (
	generateable     = true
	generateIfFormat = "eth%d"

	scrapliPlatformName = "nokia_sros"
)

// Register registers the node in the NodeRegistry.
func Register(r *nodes.NodeRegistry) {
	generateNodeAttributes := nodes.NewGenerateNodeAttributes(generateable, generateIfFormat)
	platformAttrs := &nodes.PlatformAttrs{
		ScrapliPlatformName: scrapliPlatformName,
	}

	nrea := nodes.NewNodeRegistryEntryAttributes(defaultCredentials, generateNodeAttributes, platformAttrs)

	r.Register(kindNames, func() nodes.Node {
		return new(cvsr)
	}, nrea)
}

type cvsr struct {
	nodes.DefaultNode
}

func (n *cvsr) Init(cfg *types.NodeConfig, opts ...nodes.NodeOption) error {
	// Init DefaultNode
	n.DefaultNode = *nodes.NewDefaultNode(n)

	n.LicensePolicy = types.LicensePolicyWarn

	n.Cfg = cfg
	for _, o := range opts {
		o(n)
	}

	n.Cfg.Binds = append(n.Cfg.Binds,
		fmt.Sprint(filepath.Join(n.Cfg.LabDir, "sros"), ":/home/sros"),
		fmt.Sprint(filepath.Join(n.Cfg.LabDir, "nokia"), ":/nokia"),
		fmt.Sprint(filepath.Join(n.Cfg.LabDir, "nokia", "license", "license.txt"), ":/nokia/license/license.txt"),
	)

	return nil
}

func (n *cvsr) PreDeploy(ctx context.Context, params *nodes.PreDeployParams) error {

	nodeCfg := n.Config()

	// create dirs for bind mounts
	utils.CreateDirectory(n.Cfg.LabDir, 0777)
	utils.CreateDirectory(filepath.Join(n.Cfg.LabDir, "sros"), 0777)
	utils.CreateDirectory(filepath.Join(n.Cfg.LabDir, "sros", "flash1"), 0777)
	utils.CreateDirectory(filepath.Join(n.Cfg.LabDir, "sros", "flash2"), 0777)
	utils.CreateDirectory(filepath.Join(n.Cfg.LabDir, "sros", "flash3"), 0777)
	utils.CreateDirectory(filepath.Join(n.Cfg.LabDir, "nokia"), 0777)
	utils.CreateDirectory(filepath.Join(n.Cfg.LabDir, "nokia", "license"), 0777)

	dst := filepath.Join(nodeCfg.LabDir, "nokia", "license", "license.txt")

	if nodeCfg.License != "" {
		// copy license file to node specific lab directory
		src := nodeCfg.License
		if err := utils.CopyFile(src, dst, 0644); err != nil {
			return fmt.Errorf("file copy [src %s -> dst %s] failed %v", src, dst, err)
		}
		log.Debugf("CopyFile src %s -> dst %s succeeded", src, dst)
	} else {
		utils.CreateFile(dst, "")
	}

	// utils.CreateDirectory(filepath.Join(n.Cfg.LabDir, "flash3"), 0777)

	// utils.CreateDirectory(filepath.Join(n.Cfg.LabDir, "nokia"), 0777)
	// utils.CreateDirectory(filepath.Join(n.Cfg.LabDir, "nokia", "config"), 0777)
	// utils.CreateDirectory(filepath.Join(n.Cfg.LabDir, "nokia", "license"), 0777)

	_, err := n.LoadOrGenerateCertificate(params.Cert, params.TopologyName)
	if err != nil {
		return nil
	}

	return nil

}

// CheckInterfaceName checks if a name of the interface referenced in the topology file correct.
func (n *cvsr) CheckInterfaceName() error {
	ifRe := regexp.MustCompile(`^1/1/\d+$`)
	for _, e := range n.Endpoints {
		if !ifRe.MatchString(e.GetIfaceName()) {
			return fmt.Errorf("Nokia cvsr interface name %q doesn't match the required pattern. cvsr interfaces should be named as 1/1/X where X is the interface number", e.GetIfaceName())
		}
	}

	return nil
}
