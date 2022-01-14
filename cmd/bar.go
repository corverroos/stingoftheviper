package cmd

import (
	"github.com/corverroos/stingoftheviper/bar"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// bar returns the bar cobra sub-command which requires bar config flag.
func newBar() *cobra.Command {
	var conf bar.Config

	cmd := &cobra.Command{
		Use:   "bar",
		Short: "Prints the bar config",
		Long:  "Requires and prints the bar config",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return bar.Run(cmd.OutOrStdout(), conf)
		},
	}

	bindBarFlags(cmd.Flags(), &conf)

	return cmd
}

func bindBarFlags(flags *pflag.FlagSet, config *bar.Config) {
	flags.StringVar(&config.String, "bar_string", "bar default", "bar string field")
	flags.Float64Var(&config.Float, "bar_float", 0.1, "bar float field")
	flags.BoolVar(&config.Bool, "bar_bool", true, "bar bool field")
}
