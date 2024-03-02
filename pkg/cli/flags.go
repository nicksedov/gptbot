package cli

import (
	"flag"
)

// Command line parameters
var FlagConfig  = flag.String("config", "settings.yaml", "configuration YAML file")