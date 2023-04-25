package linuxsystem

import (
	"os/exec"
	"strings"
)

type Software struct {
	Name    string
	Version string
}

// return a list a software based on snap package on the machine
func ListSnapPackages() ([]Software, error) {
	cmd := exec.Command("snap", "list")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(output), "\n")[1:]
	packages := make([]Software, 0, len(lines))
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) >= 1 {
			packages = append(packages, Software{Name: fields[0], Version: fields[1]})
		}
	}

	return packages, nil
}
