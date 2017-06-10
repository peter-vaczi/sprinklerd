package cmd

import (
	"log"
	"os"

	"github.com/peter.vaczi/sprinklerd/core"
	"github.com/peter.vaczi/sprinklerd/utils"
	"github.com/spf13/cobra"
)

var setFlagOn bool
var setFlagOff bool
var setFlagPin int = -1

// setCmd represents the add command
var setCmd = &cobra.Command{
	Use:   "set <name> [flags]",
	Short: "Set parameters of a device",
	Long:  `Set parameters of a device`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 1 || (setFlagOn && setFlagOff) {
			cmd.Usage()
			os.Exit(-1)
		}

		var dev core.Device

		err := utils.GetRequest(daemonSocket+"/v1/devices/"+args[0], &dev)
		if err != nil {
			log.Fatal(err)
		}

		log.Print(dev)
		if setFlagPin != -1 {
			dev.Pin = setFlagPin
		}
		if setFlagOn {
			dev.On = true
		}
		if setFlagOff {
			dev.On = false
		}

		log.Print(dev)
		err = utils.PutRequest(daemonSocket+"/v1/devices/"+dev.Name, &dev)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	deviceCmd.AddCommand(setCmd)
	setCmd.PersistentFlags().IntVar(&setFlagPin, "pin", -1, "GPIO pin associated with this device")
	setCmd.PersistentFlags().BoolVar(&setFlagOn, "on", false, "set the device on")
	setCmd.PersistentFlags().BoolVar(&setFlagOff, "off", false, "set the device off")
}
