package main

import (
	"fmt"
	"os/exec"
)

func main() {

	cmd := exec.Command("/taste")
	cmd.Env = append(cmd.Env, "LD_PRELOAD=/inject-libc.so")
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
	if err != nil {
		panic(err)
	}

	cmd = exec.Command("/taste")
	output, err = cmd.CombinedOutput()
	fmt.Println(string(output))
	if err != nil {
		panic(err)
	}
}
