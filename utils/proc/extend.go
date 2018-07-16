package proc

import (
	"os"
	"fmt"
	"strconv"
)

// return all pids in procfs
func AllPids() []string {
	p, err := os.Open(procfs)
	if err != nil {
		fmt.Println("can not open", procfs)
		return nil
	}

	defer p.Close()

	var files []string
	files, err = p.Readdirnames(0)
	if err != nil {
		fmt.Printf("Readdirnames %v failed\n", procfs)
		return nil
	}

	var pids []string
	for _, v := range(files) {
		if pid, _ := strconv.Atoi(v); pid != 0 {
			pids = append(pids, v)
		}
	}
	return pids
}