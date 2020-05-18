package main

func main() {
	cmd := parseCmd()

	if cmd.versionFlag {
		println("version 1.0.0")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		newJVM(cmd).start()
	}
}
