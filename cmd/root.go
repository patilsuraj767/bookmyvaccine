package cmd

import (
	"fmt"
	"github.com/patilsuraj767/bookmyvaccine/bookmyvaccine"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "bookmyvaccine",
	Short: "bookmyvaccine - command line for continuously checking vaccine availability",
	Run: func(cmd *cobra.Command, args []string) {
		bookmyvaccine.CheckAvailability()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var age_group int
var capacity int
var pincode string
var date string

func init() {
	cobra.OnInitialize()
	rootCmd.PersistentFlags().IntVar(&age_group, "age_group", 18, "Age Group can have two value 18 or 45")
	rootCmd.PersistentFlags().IntVar(&capacity, "capacity", 1, "Number of doses you need")
	rootCmd.PersistentFlags().StringVar(&pincode, "pincode", "", "Pin code of you area (required)")
	rootCmd.MarkPersistentFlagRequired("pincode")
	rootCmd.PersistentFlags().StringVar(&date, "date", "", "Date should be in format DD-MM-YYYY (required)")
	rootCmd.MarkPersistentFlagRequired("date")
	viper.BindPFlag("AGE_GROUP", rootCmd.PersistentFlags().Lookup("age_group"))
	viper.BindPFlag("AVAILABLE_CAPACITY", rootCmd.PersistentFlags().Lookup("capacity"))
	viper.BindPFlag("PIN_CODE", rootCmd.PersistentFlags().Lookup("pincode"))
	viper.BindPFlag("DATE", rootCmd.PersistentFlags().Lookup("date"))
}
