package helpers

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

func CheckFlags() {
	devFlag := flag.Bool("dev", false, "Run the server in development mode")
	buildFlag := flag.Bool("build", false, "Build the client for production")
	flag.Parse()

	if *buildFlag {
		err := BuildClient()
		if err != nil {
			fmt.Println("error building client:", err)
		}
		return
	}

	if *devFlag {
		fmt.Println("starting server in development mode")
		go RunDevServer()
	} else {
		fmt.Println("starting server in production mode")
		fs := http.FileServer(http.Dir("../client/dist"))
		http.Handle("/", fs)
		fmt.Println("production server running on http://localhost:8080")
	}
}

func RunDevServer() error {
	cmd := exec.Command("npm", "run", "dev")
	cmd.Stdout = nil
	cmd.Stderr = nil

	fmt.Println("development server running on http://localhost:5173 (Vite)")
	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("failed to start command: %w", err)
	}

	err = cmd.Wait()
	if err != nil {
		return fmt.Errorf("command finished with error: %w", err)
	}

	return nil
}

func BuildClient() error {
	cmd := exec.Command("npm", "run", "build")
	cmd.Stdout = nil
	cmd.Stderr = nil

	fmt.Println("building client for production")
	err := cmd.Start()
	if err != nil {
		fmt.Println("failed to start command:", err)
		os.Exit(0)
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Println("command finished with error:", err)
		os.Exit(0)
	}

	fmt.Println("client built, run with no flags to start server in production mode")
	os.Exit(0)
	return nil
}
