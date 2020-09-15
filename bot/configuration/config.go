package configuration

import "github.com/hashicorp/hcl/v2/hclsimple"

// BaritoneConfig is the configuration structure for the bot
type BaritoneConfig struct {
	Token string `hcl:"token"`
}

// Config is the instance of BaritoneConfig with the loaded config
var Config BaritoneConfig

func init() {
	hclsimple.DecodeFile("conf.hcl", nil, &Config)
}
