package commands

import (
	crypto "github.com/authit/crypto/commands"
	frontend "github.com/authit/frontend/commands"
)

var ConfigOptions = map[string]interface{}{
	"crypto":  crypto.ConfigOptions,
	"frontend":  frontend.ConfigOptions,
}
