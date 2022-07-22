package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/checkpoint-restore/go-criu/v5/crit"
	"github.com/spf13/cobra"
)

// The crit service used to invoke all commands
var c crit.CritSvc

// All members needed for crit struct
var inputFilePath, outputFilePath, inputDirPath string
var pretty, noPayload bool

// The `crit` command
var rootCmd = &cobra.Command{
	Use:   "crit",
	Short: "CRIU Image Tool(CRIT) to manipulate CRIU image files",
	Long: `CRIU Image Tool (CRIT) is a command line tool to investigate
binary image files generated by CRIU and view them in JSON.
This is a Go implementation of the original Python app.
Find the complete documentation is at https://criu.org/CRIT`,
}

// The `crit decode` command
var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Convert binary image to JSON",
	Long: `Convert the input binary image to JSON and write it to a file.
If no output file is provided, the JSON is printed to stdout.`,
	Run: func(cmd *cobra.Command, args []string) {
		c = crit.NewCli(inputFilePath, outputFilePath,
			inputDirPath, pretty, noPayload)
		img, err := c.Decode()
		if err != nil {
			log.Fatal(err)
		}

		var jsonData []byte
		if pretty {
			jsonData, err = json.MarshalIndent(img, "", "    ")
		} else {
			jsonData, err = json.Marshal(img)
		}
		if err != nil {
			log.Fatal(errors.New(fmt.Sprint("Error processing data into JSON: ", err)))
		}
		// If no output file, print to stdout
		if outputFilePath == "" {
			fmt.Println(string(jsonData))
			return
		}
		// Write to output file
		jsonFile, err := os.Create(outputFilePath)
		if err != nil {
			log.Fatal(errors.New(fmt.Sprint("Error opening destination file: ", err)))
		}
		defer jsonFile.Close()
		_, err = jsonFile.Write(jsonData)
		if err != nil {
			log.Fatal(errors.New(fmt.Sprint("Error writing JSON data: ", err)))
		}
	},
}

// The `crit encode` command
var encodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "Convert JSON to binary image file",
	Long:  "Convert the input JSON to a CRIU image file.",
	Run: func(cmd *cobra.Command, args []string) {
		c = crit.NewCli(inputFilePath, outputFilePath,
			inputDirPath, pretty, noPayload)
		// Convert JSON to Go struct
		img, err := c.Parse()
		if err != nil {
			log.Fatal(err)
		}
		// Write Go struct to binary image file
		if err := c.Encode(img); err != nil {
			log.Fatal(err)
		}
	},
}

// The `crit show` command
var showCmd = &cobra.Command{
	Use:   "show INPATH",
	Short: "Convert binary image to human-readable JSON",
	Long:  "Convert the input binary image to human-readable JSON and print to stdout",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		inputFilePath = args[0]
		pretty = true
		c = crit.NewCli(inputFilePath, outputFilePath,
			inputDirPath, pretty, noPayload)
		img, err := c.Decode()
		if err != nil {
			log.Fatal(err)
		}

		jsonData, err := json.MarshalIndent(img, "", "    ")
		if err != nil {
			log.Fatal(errors.New(fmt.Sprint("Error processing data into JSON: ", err)))
		}
		fmt.Println(string(jsonData))
	},
}

// The `crit info` command
var infoCmd = &cobra.Command{
	Use:   "info INPATH",
	Short: "Show information about the image file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		inputFilePath = args[0]
		c = crit.NewCli(inputFilePath, outputFilePath,
			inputDirPath, pretty, noPayload)
		img, err := c.Info()
		if err != nil {
			log.Fatal(err)
		}

		jsonData, err := json.MarshalIndent(img, "", "    ")
		if err != nil {
			log.Fatal(errors.New(fmt.Sprint("Error processing data into JSON: ", err)))
		}
		fmt.Println(string(jsonData))
	},
}

// The `crit x` command
var xCmd = &cobra.Command{
	Use:   "x DIR {ps|fds|mems|rss}",
	Short: "Explore the image directory",
	Long:  "Explore the image directory with one of (ps, fds, mems, rss) options",
	// Exactly two arguments are required:
	// * Path of the input directory
	// * Explore type
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		inputDirPath = args[0]
		// We can use an empty interface to hold the
		// returned object since we don't really care
		// about the data itself, as long as we can
		// marshal it into JSON and display it.
		var xData interface{}
		var err error

		c = crit.NewCli(inputFilePath, outputFilePath,
			inputDirPath, pretty, noPayload)
		// Switch the explore type and call the handler.
		switch args[1] {
		case "ps":
			xData, err = c.ExplorePs()
		case "fds":
			xData, err = c.ExploreFds()
		case "mems":
			xData, err = c.ExploreMems()
		case "rss":
			xData, err = c.ExploreRss()
		default:
			err = errors.New("Error exploring directory: Invalid explore type")
		}
		if err != nil {
			log.Fatal(err)
		}

		jsonData, err := json.MarshalIndent(xData, "", "    ")
		if err != nil {
			log.Fatal(errors.New(fmt.Sprint("Error processing data into JSON: ", err)))
		}
		fmt.Println(string(jsonData))
	},
}

// Add all commands to the root command and configure flags
func init() {
	// Disable completion generation
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	// Decode options
	decodeCmd.Flags().StringVarP(&inputFilePath, "input", "i", "",
		"Path to the binary image file")
	decodeCmd.Flags().StringVarP(&outputFilePath, "output", "o", "",
		"Path to the destination JSON file")
	decodeCmd.Flags().BoolVar(&pretty, "pretty", false,
		"Provide indented and multi-line JSON output")
	rootCmd.AddCommand(decodeCmd)
	// Encode options
	encodeCmd.Flags().StringVarP(&inputFilePath, "input", "i", "",
		"Path to the JSON file")
	encodeCmd.Flags().StringVarP(&outputFilePath, "output", "o", "",
		"Path to the destination image file")
	rootCmd.AddCommand(encodeCmd)
	// Show options
	showCmd.Flags().BoolVar(&noPayload, "nopl", false,
		"Do not show payload contents")
	rootCmd.AddCommand(showCmd)
	// Info and X commands
	rootCmd.AddCommand(infoCmd)
	rootCmd.AddCommand(xCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
