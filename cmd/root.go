package cmd

import (
	"github.com/spf13/cobra"

	"mahuike/cmd/aws"
	"mahuike/cmd/gcp"
	"mahuike/logger"

	log "github.com/sirupsen/logrus"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "mahuike CLOUD SERVICE [OPTIONS]",
		Short: "A Cloud Native Event injector",
	}
)

var l = logger.GetLogger().WithFields(log.Fields{"pkg": "main"})

// func parseArgs() {
// 	flag.StringVar(&c.ServiceId, "f", "", "Function to invoke")
// 	flag.StringVar(&c.Region, "r", "", "Region where the resource is located")
// 	flag.StringVar(&c.Service, "s", "aws-lambda", "Service to call")

// 	flag.Parse()
// }

// Execute executes the root command.
func Execute() error {
	// parseArgs()
	// l.Info("Start")

	// scanner := bufio.NewScanner(os.Stdin)
	// for scanner.Scan() {
	// 	c.Data = scanner.Text()
	// }

	// if scanner.Err() != nil {
	// 	l.Error(scanner.Err().Error())
	// }

	// inv := invoker.NewInvoker(&c)
	// inv.Run(c)
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(aws.AwsCmd)
	rootCmd.AddCommand(gcp.GcpCmd)
	// cobra.OnInitialize(initConfig)

	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	//rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")
	//rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
	// rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	// viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	// viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	// viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	// viper.SetDefault("license", "apache")

	// rootCmd.AddCommand(addCmd)
	// rootCmd.AddCommand(initCmd)
}

// func initConfig() {
// 	if cfgFile != "" {
// 		// Use config file from the flag.
// 		viper.SetConfigFile(cfgFile)
// 	} else {
// 		// Find home directory.
// 		home, err := os.UserHomeDir()
// 		cobra.CheckErr(err)

// 		// Search config in home directory with name ".cobra" (without extension).
// 		viper.AddConfigPath(home)
// 		viper.SetConfigType("yaml")
// 		viper.SetConfigName(".cobra")
// 	}

// 	viper.AutomaticEnv()

// 	if err := viper.ReadInConfig(); err == nil {
// 		fmt.Println("Using config file:", viper.ConfigFileUsed())
// 	}
// }
