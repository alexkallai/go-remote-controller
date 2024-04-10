package main

import (
	"flag"
	"fmt"
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
	LOGGING                = flag.Bool("log", false, "Turn logging on of off")
	//PASSWORD = flag.String("password", "", "Password for authentication")
)

func handleCliArgs() {
	// Handle CLI args
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

func main() {
	handleCliArgs()
	// Get IP addresses and print the service's availability
	printServerInfo()
	httpserverutils.SetupHttpServerEndpoints(PORT_ARG)

}
