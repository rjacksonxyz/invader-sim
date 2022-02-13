package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	filepath            string
	numAliens, numSteps uint64
	rootCmd             = &cobra.Command{
		Use:  "invader-sim",
		Long: "Invader Sim: Alien Invasion Simulator\nSpec info available at: https://github.com/zkmiyavi/invader-sim/prompt/prompt.txt",
		RunE: func(cmd *cobra.Command, args []string) error {
			return simulate()
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func init() {
	//parse command line arguments
	rootCmd.Flags().StringVarP(&filepath, "filepath", "f", "map.txt", "filepath to .txt file containing map simulator will use")
	rootCmd.Flags().Uint64VarP(&numAliens, "aliens", "a", 10, "number of alien invaders")
	rootCmd.Flags().Uint64VarP(&numSteps, "steps", "s", 10000, "maximum possible number of steps by aliens")
}

func simulate() error {
	//implement simulation
	return nil
}
