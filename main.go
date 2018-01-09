package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/golang/glog"
)

func NewRootCmd() *cobra.Command {
	var (
		enableAnalytics = true
	)
	var rootCmd = &cobra.Command{
		Short:             `Log Demo`,
		DisableAutoGenTag: true,
		PersistentPreRun: func(c *cobra.Command, args []string) {
			c.Flags().VisitAll(func(flag *pflag.Flag) {
				log.Printf("FLAG: --%s=%q", flag.Name, flag.Value)
			})
		},
	}
	rootCmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)
	// ref: https://github.com/kubernetes/kubernetes/issues/17162#issuecomment-225596212
	flag.CommandLine.Parse([]string{})
	rootCmd.PersistentFlags().BoolVar(&enableAnalytics, "analytics", enableAnalytics, "Send analytical events to Google Analytics")

	rootCmd.AddCommand(NewCmdCheck())
	return rootCmd
}

func NewCmdCheck() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "check",
		Short:             "Check restic backup",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("log.Println_____")
			log.Println("node.Name")

			// ---------------------------------

			fmt.Println("glog.Infoln_____")
			glog.Infoln("node.Name")

			fmt.Println("glog.Warningln_____")
			glog.Warningln("node.Name")

			fmt.Println("glog.Errorln_____")
			glog.Errorln("node.Name")

			// ---------------------------------

			fmt.Println("glog.V(0).Infoln_____")
			glog.V(0).Infoln("node.Name")

			fmt.Println("glog.V(1).Infoln_____")
			glog.V(1).Infoln("node.Name")

			fmt.Println("glog.V(2).Infoln_____")
			glog.V(2).Infoln("node.Name")

			fmt.Println("glog.V(3).Infoln_____")
			glog.V(3).Infoln("node.Name")

			fmt.Println("glog.V(4).Infoln_____")
			glog.V(4).Infoln("node.Name")
		},
	}
	return cmd
}

func main() {
	if err := NewRootCmd().Execute(); err != nil {
		log.Fatalln("Error in Stash Main:", err)
	}
}
