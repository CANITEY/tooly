package dependencies

import (
	"os/exec"
)

type msg struct {
	Pass string
	Err error
}

// function checks if a dependency is installed or not
// @param: dep -> dependency to check
// @output: string -> the version of the dependency
// @output: error -> indicator in case the dependency isn't installed
func checkDependency(dep string) (string, error) {
	output, err := exec.Command(dep, "--version").Output()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

// function checks if a dependency is installed or not
// @param: deps -> dependencies to check
// @output: map of all depenecies and the corresponding message
func CheckDependencies(deps ...string) (map[string]msg) {
	var output = make(map[string]msg)
	for _, dep := range deps {
		if out, err := checkDependency(dep); err != nil {
			output[dep] = msg{
				Err: err,
			}
		} else {
			output[dep] = msg{
				Pass: out,
			}
		}
	}
	return output
}
