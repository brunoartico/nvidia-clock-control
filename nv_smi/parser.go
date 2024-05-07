package nv_smi

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
)

func mhzToInt(freq string) int {
	clock, err := strconv.Atoi(strings.TrimSpace(freq[:len(freq)-4]))
	if err != nil {
		log.Fatal(err)
	}

	return clock
}
 
func parseSupportedClocks(response string) GPUClocks {
	memClocks := mapset.NewSet[int]()
	coreClocks := mapset.NewSet[int]()

	minMemClock := int(^uint(0) >> 1)
	minCoreClock := minMemClock

	r := csv.NewReader(strings.NewReader(response))

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		memClock := mhzToInt(record[0])
		coreClock := mhzToInt(record[1])

		memClocks.Add(memClock)
		coreClocks.Add(coreClock)

		if minMemClock > memClock {
			minMemClock = memClock
		}

		if minCoreClock > coreClock {
			minCoreClock = coreClock
		}

	}

	return GPUClocks{memClocks, minMemClock, coreClocks, minCoreClock}
}
