package cmd

import (
	"log"
	"os"

	"github.com/peter-vaczi/sprinkler/core"
	"github.com/peter-vaczi/sprinkler/utils"
	"github.com/spf13/cobra"
)

var addFlagOn bool
var addFlagSwitchOnLow bool
var addFlagPin int

// deviceAddCmd represents the add command
var deviceAddCmd = &cobra.Command{
	Use:   "add <name> [flags]",
	Short: "Add a new device",
	Long:  `Add a new device`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 1 {
			cmd.Usage()
			os.Exit(-1)
		}

		dev := core.Device{Name: args[0], On: addFlagOn, Pin: addFlagPin, SwitchOnLow: addFlagSwitchOnLow}
		err := utils.PostRequest(daemonSocket+"/v1/devices", &dev)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	deviceCmd.AddCommand(deviceAddCmd)
	deviceAddCmd.PersistentFlags().IntVar(&addFlagPin, "pin", 0, "GPIO pin associated with this device")
	deviceAddCmd.PersistentFlags().BoolVar(&addFlagOn, "on", false, "set the device on")
	deviceAddCmd.PersistentFlags().BoolVar(&addFlagSwitchOnLow, "switch-on-low", false, "device is considered 'on' when it's output pin is low")
}
