package cmd

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

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
		rowNum, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("%s is not a decimal number.", args[0])
		}

		binaryNum := fmt.Sprintf("Binary: %b", rowNum)
		fmt.Println(binaryNum)

		hexNum := fmt.Sprintf("Hex: %x", rowNum)
		fmt.Println(hexNum)
	},
}
