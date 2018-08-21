package cmd

import (
	"fmt"
	"image"
	"os"

	"github.com/sinnott74/hero/internal"
	"github.com/spf13/cobra"
)

var (
	out      string
	bgColor  string
	icons    []string
	height   int
	width    int
	iconsize int
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hero",
	Short: "Hero creates banner images to be displayed on a website",
	RunE: func(cmd *cobra.Command, args []string) error {

		iconDimensions := image.Point{iconsize, iconsize}

		pic := image.NewRGBA(image.Rect(0, 0, width, height))

		// Set the color
		err := internal.SetImageColor(pic, &bgColor)
		if err != nil {
			return err
		}

		// Add the icons
		err = internal.AddIcons(pic, icons, iconDimensions)
		if err != nil {
			return err
		}

		// Create the PNG
		err = internal.CreatePNG(pic, &out)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&out, "output", "o", "./hero.png", "File to output")
	rootCmd.PersistentFlags().StringVarP(&bgColor, "color", "c", "red", "Background color")
	rootCmd.PersistentFlags().StringArrayVarP(&icons, "icons", "i", []string{}, "Icons")
	rootCmd.PersistentFlags().IntVarP(&width, "width", "x", 960, "Max x coordinate of the hero - width of the hero")
	rootCmd.PersistentFlags().IntVarP(&height, "height", "y", 480, "Max y coordinate of the hero - height of the hero")
	rootCmd.PersistentFlags().IntVarP(&iconsize, "iconsize", "s", 300, "Width and height of the incons on the hero")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
