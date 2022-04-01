package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/hashiserver/dummy-image-generator/controller"
	"github.com/hashiserver/dummy-image-generator/lib"
	"github.com/hashiserver/dummy-image-generator/model"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dummy",
	Short: "dummy image generate",
	Long:  `Generate a dummy image. Specify size. Specify color. Text specification.`,
	Run: func(cmd *cobra.Command, args []string) {
		var option model.Option

		// flag読み取り・validation
		option.Width, _ = cmd.Flags().GetInt("width")
		if option.Width == 0 {
			log.Println("-w is required. Please enter a width")
			os.Exit(1)
		}

		option.Height, _ = cmd.Flags().GetInt("height")
		if option.Height == 0 {
			log.Println("-t is required. Please enter a height")
			os.Exit(1)
		}

		option.Extension, _ = cmd.Flags().GetString("extension")
		if !lib.ValidateExtension(option.Extension) {
			log.Println("Only 'jpg', 'jpeg', 'png', and 'gif' are allowed with the -e option")
			os.Exit(1)
		}

		// filePathの生成
		filePath := "./dummy_" + strconv.Itoa(option.Width) + "_" + strconv.Itoa(option.Height) + "." + option.Extension

		err := controller.CreateImageFile(filePath, option)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
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
	rootCmd.Flags().StringP("extension", "e", "png", "Image Extension")
}
