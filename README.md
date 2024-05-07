# nvidia-clock-control

This software provides a simple system tray app that allows to select and limit NVIDIA GPU's Core and Memory clock speeds, keeping temperatures within desired ranges.  

## Yet another clock control tool?
Despite existing methods to limit the NVIDIA GPU power on laptops, they may be grumpy and activate (or not) at their own will -- I am talking to you, [WhisperMode](https://www.nvidia.com/en-us/geforce/technologies/whisper-mode/)! -- or be a drastic [FPS Limiter](https://steamcommunity.com/sharedfiles/filedetails/?id=2950123288), which limits way more than just power consumption. 

This is a lightweight, faster and easier-to-use tool to achieve the same. 
👋 Hey there laptop gamer looking for a quieter late-night gaming session, this is for you!

## Required Software
* GPU Clocks are controled using [_nvidia-smi_](https://developer.nvidia.com/system-management-interface).

## Disclaimer
Use at your risk! Although _nvidia-smi_ ships with Nvidia drivers, there are no documented side-effects of using it to limit GPU clocks. Threat this software as an overclocking tool!

## Supported Hardware
* GPUs supported are restricted to ones _nvidia-smi_ supports, as per current _nvidia-smi_ documentation, Volta architecture or superior.

## Build
To build without console window, use:
go build -ldflags "-H=windowsgui"

## TO-DO
- [ ] Unit Tests
- [ ] Better UX for the lengthy Core Clock list
- [ ] Support for Profiles
- [ ] Multi-GPU support