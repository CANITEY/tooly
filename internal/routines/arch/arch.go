package arch

import (
	"os/exec"
)

func CheckYay() (bool, error) {
	command := exec.Command("yay", "-h")
	_, err := command.Output()
	if err != nil {
		return false, err
	}
	return true, nil
}

func InstallYay() (bool, error) {
	cmd := exec.Command("pacman", "-S", "--needed", "git", "base-devel")
	_, err := cmd.Output()
	if err != nil {
		return false, err
	}

	cmd = exec.Command("git", "clone", "https://github.com/Jguer/yay.git", "/tmp/yay")
	_, err = cmd.Output()
	if err != nil {
		return false, err
	}

	cmd = exec.Command("cd", "/tmp/yay")
	_, err = cmd.Output()
	if err != nil {
		return false, err
	}

	cmd = exec.Command("makepkg", "-si")
	_, err = cmd.Output()
	if err != nil {
		return false, err
	}

	return true, nil
}
