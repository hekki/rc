package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

var binary bool
var hex bool

const bitSize = 64

func init() {
	RootCmd.PersistentFlags().BoolVarP(&binary, "binary", "b", false, "binary input")
	RootCmd.PersistentFlags().BoolVarP(&hex, "hex", "e", false, "hex input")
}

// RootCmd is root command
var RootCmd = &cobra.Command{
	Use:   "rc [num]",
	Short: "Command Line radix converter",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var rowNum int
		var base int

		if binary {
			base = 2
		} else if hex {
			base = 16
		} else {
			base = 10
		}

		rowNum = parseStringNum(args[0], base)
		fmt.Println("Decimal:", toDecimalString(rowNum))
		fmt.Println("Binary:", toBinaryString(rowNum))
		fmt.Println("Hex:", toHexString(rowNum))
	},
}

func toDecimalString(num int) string {
	return fmt.Sprintf("%d", num)
}

func toBinaryString(num int) string {
	return fmt.Sprintf("%b", num)
}

func toHexString(num int) string {
	return fmt.Sprintf("%x", num)
}

func parseStringNum(stringNum string, base int) int {
	num, err := strconv.ParseInt(stringNum, base, bitSize)
	if err != nil {
		log.Fatalf("%s is illegal string.", stringNum)
	}

	return int(num)
}
