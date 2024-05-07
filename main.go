package main

import (
	"bartico.com/nvidia-clock-control/icon"
	"bartico.com/nvidia-clock-control/nv_smi"
	"bartico.com/nvidia-clock-control/tray"
	"github.com/getlantern/systray"

	"fmt"
	"time"
)

var currentMemoryLimit = -1
var currentCoreLimit = -1

var mLimits *systray.MenuItem
var mMemoryGroup *systray.MenuItem
var mCoreGroup *systray.MenuItem

func printLimit(limit int) string {
	if limit == -1 {
		return "Unlimited"
	}

	return fmt.Sprintf("%d MHz", limit)
}

func addLimitsInfo() {
	mLimits = systray.AddMenuItem("Limits not available", "Current Limits")
	mLimits.Disable()

	go func() {
		for {
			mLimits.SetTitle(fmt.Sprintf("Memory: %s - Core: %s", printLimit(currentMemoryLimit), printLimit(currentCoreLimit)))	
			time.Sleep(3 * time.Second)
		}
	}()
}
	
var gpuClocks = nv_smi.QuerySupportedClocks()

func addMemControl() {
	mMemoryGroup = systray.AddMenuItem("Memory Clock Limit", "Memory Clock Limit")

	group := new(tray.MenuGroup)

	for _, value := range gpuClocks.Memory() {
		var memoryClock = printLimit(value)
		group.Add(mMemoryGroup.AddSubMenuItemCheckbox(memoryClock, memoryClock, false), func() {
			currentMemoryLimit = value
			nv_smi.SetMemoryLimit(gpuClocks.MinimumMemoryClock, currentMemoryLimit)
		})
	}

	group.Add(mMemoryGroup.AddSubMenuItemCheckbox("Unlimited", "Unlimited", true), func() {
		currentMemoryLimit = -1
		nv_smi.SetMemoryLimit(0, currentMemoryLimit)
	})
}

func addCoreControl() {
	mCoreGroup = systray.AddMenuItem("Core Clock Limit", "Core Clock Limit")

	group := new(tray.MenuGroup)
	
	for _, value := range gpuClocks.Core() {
		var coreClock = printLimit(value)
		group.Add(mCoreGroup.AddSubMenuItemCheckbox(coreClock, coreClock, false), func() {
			currentCoreLimit = value
			nv_smi.SetCoreLimit(gpuClocks.MinimumCoreClock, currentCoreLimit)
		})
	}
	
	group.Add(mCoreGroup.AddSubMenuItemCheckbox("Unlimited", "Unlimited", true), func() {
		currentCoreLimit = -1
		nv_smi.SetCoreLimit(0, currentCoreLimit)
	})
}

func onReady() {
	systray.SetIcon(icon.Data)
	systray.SetTitle("Nvidia Clock Control")

	nv_smi.ResetAllLimits()

	addLimitsInfo()

	systray.AddSeparator()

	addMemControl()
	addCoreControl()

	mQuit := systray.AddMenuItem("Quit", "Quit")
	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
	}()
}

func main() {
	systray.Run(onReady, nil)
}
