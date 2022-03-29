package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/hashiserver/dummy-image-generator/utils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dummy",
	Short: "dummy image generate",
	Long:  `Generate a dummy image. Specify size. Specify color. Text specification.`,
	Run: func(cmd *cobra.Command, args []string) {
		// flag読み取り
		width, _ := cmd.Flags().GetInt("width")
		if width == 0 {
			log.Println("-w is required. Please enter a width")
			os.Exit(1)
		}

		height, _ := cmd.Flags().GetInt("height")
		if height == 0 {
			log.Println("-t is required. Please enter a height")
			os.Exit(1)
		}

		extension, _ := cmd.Flags().GetString("extension")
		if !utils.IsExtension(extension) {
			log.Println("Only 'jpg', 'jpeg', 'png', and 'gif' are allowed with the -e option")
			os.Exit(1)
		}

		utils.CreateImageFile(width, height, extension)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntP("width", "w", 0, "Image Width")
	rootCmd.Flags().IntP("height", "t", 0, "Image Height")
	rootCmd.Flags().StringP("extension", "e", "", "Image Extension")
}
