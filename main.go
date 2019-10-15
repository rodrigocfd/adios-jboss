package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	outLines := runCmd("ps aux | grep jboss")
	pids := findJbossPids(outLines)
	killPids(pids)
}

func runCmd(cmd string) []string {
	outBin, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Printf("Error executing ps aux:\n%s\n", err)
	}
	return strings.Split(string(outBin), "\n")
}

func findJbossPids(psLines []string) []string {
	pids := []string{}
	for _, psLine := range psLines {
		if strings.Contains(psLine, "/bin/java") {
			tokens := strings.Fields(psLine)
			pids = append(pids, tokens[1])
		}
	}
	return pids
}

func killPids(pids []string) {
	if len(pids) == 0 {
		fmt.Println("Esta bosta não está rodando.")
	} else {
		for _, pid := range pids {
			killStr := fmt.Sprintf("kill -9 %s", pid)
			fmt.Printf("Tocando o foda-se: %s\n", killStr)
			runCmd(killStr)
		}
	}
}
