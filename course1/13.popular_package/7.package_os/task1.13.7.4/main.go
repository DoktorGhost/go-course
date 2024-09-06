package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func ExecBin(binPath string, args ...string) string {
	cmd := exec.Command(binPath, args...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}
	return strings.TrimSpace(string(output))
}

func main() {
	fmt.Println(ExecBin("ls", "-la"))
	fmt.Println(ExecBin("nonexistent-binary"))
}
