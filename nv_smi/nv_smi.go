package nv_smi

import (
	"bartico.com/nvidia-clock-control/exec"

	"fmt"
	"log"
)

func ResetAllLimits() {
	fmt.Println("Resetting all")
	exec.RunWithoutWindow("nvidia-smi", "-rmc")
	exec.RunWithoutWindow("nvidia-smi", "-rgc")
}

func SetCoreLimit(minClock, maxClock int) {
	if (maxClock <= 0) {
		fmt.Println("Resetting core")
		exec.RunWithoutWindow("nvidia-smi", "-rgc")
	} else {
		fmt.Println("Setting core to ", minClock, maxClock)
		exec.RunWithoutWindow("nvidia-smi", "-lgc", fmt.Sprintf("%d,%d", minClock, maxClock))
	}
}

func SetMemoryLimit(minClock, maxClock int) {
	if (maxClock <= 0) {
		fmt.Println("Resetting memory")
		exec.RunWithoutWindow("nvidia-smi", "-rmc")
	} else {
		fmt.Println("Setting memory to ", minClock, maxClock)
		exec.RunWithoutWindow("nvidia-smi", "-lmc", fmt.Sprintf("%d,%d", minClock, maxClock))
	}
}

func QuerySupportedClocks() GPUClocks {
	supportedClocks, err := exec.OutputWithoutWindow("nvidia-smi", "--query-supported-clocks", "mem,gr", "--format=csv,noheader")
	if err != nil {
		log.Fatal(err)
	}
	
	return parseSupportedClocks(string(supportedClocks))
}