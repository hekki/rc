package cmd

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

var binary bool
var hex bool

func init() {
	RootCmd.PersistentFlags().BoolVarP(&binary, "binary", "b", false, "binary flag")
	RootCmd.PersistentFlags().BoolVarP(&hex, "hex", "e", false, "hex flag")
}

// RootCmd is root command
var RootCmd = &cobra.Command{
	Use:   "rc [num]",
	Short: "Command Line radix converter",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires argument")
		}
		if len(args) > 1 {
			return errors.New("Too many arguments")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		var rowNum int

		if binary {
			bRowNum, err := strconv.ParseInt(args[0], 2, 64)
			if err != nil {
				log.Fatalf("%s is not a binary number.", args[0])
			}

			rowNum = int(bRowNum)
		} else if hex {
			hRowNum, err := strconv.ParseInt(args[0], 16, 64)
			if err != nil {
				log.Fatalf("%s is not a hex number.", args[0])
			}

			rowNum = int(hRowNum)
		} else {
			dRowNum, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatalf("%s is not a decimal number.", args[0])
			}

			rowNum = dRowNum
		}

		formatPrint(rowNum)
	},
}

func formatPrint(d int) {
	decimalNum := fmt.Sprintf("Decimal: %d", d)
	fmt.Println(decimalNum)

	binaryNum := fmt.Sprintf("Binary: %b", d)
	fmt.Println(binaryNum)

	hexNum := fmt.Sprintf("Hex: %x", d)
	fmt.Println(hexNum)
}
