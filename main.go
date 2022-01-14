package main

import (
	_ "github.com/corverroos/stingoftheviper/bar"
	"github.com/corverroos/stingoftheviper/cmd"
	_ "github.com/corverroos/stingoftheviper/foo"
	"github.com/spf13/cobra"
)

func main() {
	cobra.CheckErr(cmd.New().Execute())
}
