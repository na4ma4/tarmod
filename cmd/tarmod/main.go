package main

import (
	"archive/tar"
	"errors"
	"io"
	"os"

	"github.com/na4ma4/tarmod/internal/mainconfig"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//nolint:gochecknoglobals // cobra uses globals in main
var rootCmd = &cobra.Command{
	Use:  "tarmod <filename>",
	Args: cobra.MinimumNArgs(1),
	RunE: mainCommand,
}

//nolint:gochecknoinits // init is used in main for cobra
func init() {
	cobra.OnInitialize(mainconfig.ConfigInit)

	rootCmd.PersistentFlags().BoolP("debug", "d", false, "Debug output")
	_ = viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
	_ = viper.BindEnv("debug", "DEBUG")
}

// Used to signify an error that should return 1 transient or fixable error.
var errStandard = errors.New("")

func main() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}

func mainCommand(cmd *cobra.Command, args []string) error {
	r, err := os.Open(args[0])
	if err != nil {
		logrus.Errorf("unable to open file: %s", err)

		return errStandard
	}
	defer r.Close()

	tr := tar.NewReader(r)

	for {
		hdr, err := tr.Next()
		if errors.Is(err, io.EOF) {
			break // End of archive
		}

		if err != nil {
			logrus.Errorf("unable to red tar file: %s", err)

			return errStandard
		}

		logrus.Infof("File: %s [%s]", hdr.Name, hdr.ModTime)
	}

	return nil
}
