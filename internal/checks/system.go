package checks

import (
	"fmt"
	"os/exec"
	"strings"
)


func IsRoot() (bool, error) {
	command := exec.Command("whoami")
	output, err := command.Output()
	if err != nil {
		return false, err
	}

	out := string(output)

	out = strings.TrimSpace(out)

	if out == "root" {
		return true, nil
	}

	return false, nil
}

func GetBaseDistro() (string, error) {
	command := exec.Command("cat", "/etc/os-release")
	output, err := command.Output()
	if err != nil {
		return "", err
	}
	lines := strings.Fields(string(output))
	var fields = map[string]string{}
	for _, line := range lines {
		entry := strings.Split(line, "=")
		fields[entry[0]] = entry[1]
	}

	if val, ok := fields["ID_LIKE"]; ok {
		return val, nil
	} else {
		return "", fmt.Errorf("Report to author with your problem @canitey in discord")
	}
}
