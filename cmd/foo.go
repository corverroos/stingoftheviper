package cmd

import (
	"github.com/corverroos/stingoftheviper/foo"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// newFoo returns the foo cobra sub-command which requires foo (including foo) config flags.
func newFoo() *cobra.Command {
	var conf foo.Config

	cmd := &cobra.Command{
		Use:   "foo",
		Short: "Prints the foo (incl foo) config",
		Long:  "Requires and prints the foo (incl foo) config",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return foo.Run(cmd.OutOrStdout(), conf)
		},
	}

	bindFooFlags(cmd.Flags(), &conf)

	return cmd
}

func bindFooFlags(flags *pflag.FlagSet, config *foo.Config) {
	flags.StringVar(&config.String, "foo_string", "foo default", "foo string field")
	flags.Float64Var(&config.Float, "foo_float", -0.1, "foo float field")
	flags.BoolVar(&config.Bool, "foo_bool", false, "foo bool field")

	bindBarFlags(flags, &config.Bar)
}
