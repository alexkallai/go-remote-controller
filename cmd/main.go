package main

import (
	"flag"
	"fmt"
	"go-remote-controller/cmd/cmdpatchutil"
	"go-remote-controller/cmd/httpserverutils"
	"go-remote-controller/cmd/qrcodeutils"
	"io"
	"log"
	"os"
)

// create command line flags with fallback values
var (
	PORT_ARG               = flag.Int("port", 5000, "Port to serve on")
	SENSITIVITY_MULTIPLIER = flag.Float64("mousesensitivity", 2.0, "The ratio of mouse movement on screen / touch movement on phone screen [pixel/pixel]")
	HELP_ARG               = flag.Bool("h", false, "Print only the CLI arguments")
	LOGGING                = flag.Bool("log", false, "Turn logging on or off")
	QEMODE                 = flag.Bool("quickeditmode", false, "Don't disable Quick Edit mode for the cmd window")
	//PASSWORD = flag.String("password", "", "Password for authentication")
)

// Handle CLI args
func handleCliArgs() {
	flag.Parse()
	fmt.Println("CLI arguments:")
	flag.PrintDefaults()
	httpserverutils.SetSensitivityFromCli(*SENSITIVITY_MULTIPLIER)
	// TODO: invert logging var check for prod build
	if *LOGGING {
		log.SetOutput(io.Discard)
	}
	if *HELP_ARG {
		fmt.Println("Exiting...")
		os.Exit(0)
	}
}

// Get IP addresses and print the service's availability
func printServerInfo() {
	ipAddresses := qrcodeutils.GetLocalIpAdresses()
	port_string := fmt.Sprintf(":%d", *PORT_ARG)
	log.Println("\n Serving on the following interfaces:")
	for _, ipAddress := range ipAddresses {
		full_address := "http://" + ipAddress + port_string
		fmt.Println("\n" + full_address)
		qrcodeutils.PrintQrCode(full_address)
	}
}

// Windows CMD has an issue where it may hang the execution when
// there is an interaction, so this solves that
func handleQuickEditMode() {
	if !*QEMODE {
		log.Println("Disabling Quick Edit mode for cmd")
		err := cmdpatchutil.DisableQuickEditMode()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}
}

func startHttpServer() {
	httpserverutils.SetupHttpServerEndpoints(PORT_ARG)
}

func main() {
	handleCliArgs()
	handleQuickEditMode()
	printServerInfo()
	startHttpServer()
}
