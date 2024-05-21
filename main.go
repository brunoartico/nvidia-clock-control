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
var memoryGroup *tray.MenuGroup
var coreGroup *tray.MenuGroup

var gpuClocks = nv_smi.QuerySupportedClocks()

func printLimit(limit int) string {
	if limit == -1 {
		return "Unlimited"
	}

	return fmt.Sprintf("%d MHz", limit)
}

func addLimitsInfo() {
	mLimits := systray.AddMenuItem("Limits not available", "Current Limits")
	mLimits.Disable()

	go func() {
		for {
			mLimits.SetTitle(fmt.Sprintf("Memory: %s - Core: %s", printLimit(currentMemoryLimit), printLimit(currentCoreLimit)))	
			time.Sleep(3 * time.Second)
		}
	}()
}

func addMemControl() {
	menuItemMemoryGroup := systray.AddMenuItem("Memory Clock Limit", "Memory Clock Limit")

	memoryGroup = new(tray.MenuGroup)

	for _, value := range gpuClocks.Memory() {
		var memoryClock = printLimit(value)
		memoryGroup.Add(menuItemMemoryGroup.AddSubMenuItemCheckbox(memoryClock, memoryClock, false), value, func() {
			currentMemoryLimit = value
			nv_smi.SetMemoryLimit(gpuClocks.MinimumMemoryClock, currentMemoryLimit)
		})
	}
}

func addCoreControl() {
	menuItemCoreGroup := systray.AddMenuItem("Core Clock Limit", "Core Clock Limit")

	coreGroup = new(tray.MenuGroup)
	
	coreClocks := gpuClocks.Core()

	for _, value := range coreClocks {
		var coreClock = printLimit(value)
		coreGroup.Add(menuItemCoreGroup.AddSubMenuItemCheckbox(coreClock, coreClock, false), value, func() {
			currentCoreLimit = value
			nv_smi.SetCoreLimit(gpuClocks.MinimumCoreClock, currentCoreLimit)
		})
	}
}

func addResetControls() {
	resetAll := systray.AddMenuItem("Reset All Limits", "Reset All Limits")
	go func() {
		for {
			<- resetAll.ClickedCh
			currentMemoryLimit = -1
			currentCoreLimit = -1
			nv_smi.ResetAllLimits()
			memoryGroup.UncheckAll()
			coreGroup.UncheckAll()
		}
	}()
}

func hideSomeCoreLimits(coreStepping int) {
	allItems := coreGroup.GetAll()
	prev := allItems[0].GetValue()
	for _, groupItem := range allItems {
		groupItem.Show()
		value := groupItem.GetValue()
		if prev - value < coreStepping  {
			groupItem.Hide()
		} else {
			prev = value
		}
	}
}

func addSettingsControl() {
	settingsMenu := systray.AddMenuItem("Settings", "Settings")

	hideCoreLimitSubItem := settingsMenu.AddSubMenuItem("Hide Core Limits", "Hide Core Limits")

	hideCoreLimitsMenuGroup := new(tray.MenuGroup)
	
	for _, value := range [6]int{250, 200, 150, 100, 50, 30} {
		displayValue := fmt.Sprintf("%d MHz Steps", value)
		hideCoreLimitsMenuGroup.Add(hideCoreLimitSubItem.AddSubMenuItemCheckbox(displayValue, displayValue, false), value, func() {
			hideSomeCoreLimits(value)
		})
	}
	
	showAllCoreLimits := settingsMenu.AddSubMenuItem("Show All Core Limits", "Show All Core Limits")
	go func() {
		for {
			<-showAllCoreLimits.ClickedCh
			hideCoreLimitsMenuGroup.UncheckAll()
			coreGroup.ShowAll()
		}
	}()
}

func onReady() {
	systray.SetIcon(icon.Data)
	systray.SetTitle("Nvidia Clock Control")

	nv_smi.ResetAllLimits()

	addLimitsInfo()

	systray.AddSeparator()

	addMemControl()
	addCoreControl()

	systray.AddSeparator()

	addResetControls()

	systray.AddSeparator()
	
	addSettingsControl()

	systray.AddSeparator()

	mQuit := systray.AddMenuItem("Quit", "Quit")
	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
	}()
}

func main() {
	systray.Run(onReady, nil)
}
