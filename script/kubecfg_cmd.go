// Copyright (c) 2019 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package script

// KubeConfigCommand represents a KUBECONFIG directive in the script
type KubeConfigCommand struct {
	cmd
	kubeCfg string
}

// NewKubeConfigCommand creates a value of type KubeConfigCommand in a script
func NewKubeConfigCommand(index int, args []string) (*KubeConfigCommand, error) {
	cmd := &KubeConfigCommand{cmd: cmd{index: index, name: CmdKubeConfig, args: args}}

	if err := validateCmdArgs(CmdKubeConfig, args); err != nil {
		return nil, err
	}
	cmd.kubeCfg = searchForConfig(args)
	return cmd, nil
}

// Index is the position of the command in the script
func (c *KubeConfigCommand) Index() int {
	return c.cmd.index
}

// Name represents the name of the command
func (c *KubeConfigCommand) Name() string {
	return c.cmd.name
}

// Args returns a slice of raw command arguments
func (c *KubeConfigCommand) Args() []string {
	return c.cmd.args
}

// Config returns the path to the config file
func (c *KubeConfigCommand) Config() string {
	return c.kubeCfg
}

// searchForConfig searches in several places for
// the kubernets config:
// 1. from passed args
// 2. from ENV variable
// 3. from local homedir
func searchForConfig(args []string) string {
	if len(args) > 0 {
		return args[0]
	}
	return Defaults.KubeConfigValue
}
