/*
Copyright © 2024 Raphael Lüthy <raphael.luethy@fhnw.ch>
*/
package cmd

import (
	"fhnw/iit/croppy/anonymizer"
	"fhnw/iit/croppy/loader"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	topCrop    int
	rightCrop  int
	bottomCrop int
	leftCrop   int
	fileTypes  []string
	path       string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "img-crop",
	Short: "A tool to crop images",
	Long:  `TBD`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Cropping rectangle:\n")
		fmt.Printf("  Top: %dpx\n", topCrop)
		fmt.Printf("  Right: %dpx\n", rightCrop)
		fmt.Printf("  Bottom: %dpx\n", bottomCrop)
		fmt.Printf("  Left: %dpx\n", leftCrop)
		fmt.Printf("  File types: %s\n", fileTypes)
		fmt.Printf("  Path: %s\n", path)
		fileMap, err := loader.LoadFiles(path, fileTypes)
		if err != nil {
			fmt.Printf("Error loading files: %s\n", err)
			os.Exit(1)
		}
		anonymizer.RunAnonymize(fileMap, topCrop, rightCrop, bottomCrop, leftCrop)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntVar(&topCrop, "top", 0, "Distance from the top border to crop")
	rootCmd.Flags().IntVar(&rightCrop, "right", 0, "Distance from the right border to crop")
	rootCmd.Flags().IntVar(&bottomCrop, "bottom", 0, "Distance from the bottom border to crop")
	rootCmd.Flags().IntVar(&leftCrop, "left", 0, "Distance from the left border to crop")
	rootCmd.Flags().StringSliceVar(&fileTypes, "file-types", []string{".jpg", ".png", ".jpeg", ".mp4"}, "File types to crop")
	rootCmd.Flags().StringVar(&path, "path", "./", "Path to crop")
}

func printFileMap(fileMap map[string][]string) {
	for path, files := range fileMap {
		fmt.Printf("Path: %s\n", path)
		for _, file := range files {
			fmt.Printf("  File: %s\n", file)
		}
	}
}
