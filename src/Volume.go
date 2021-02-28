package main

import "fmt"

// VolumeInfo - Volume Information Structure
type VolumeInfo struct {
	ActiveIndex    int
	VolumePerc     int
	MutedState     bool
	BluetoothState bool
}

// Print - Volume Info Method used to output all
//    member variable information
func (v *VolumeInfo) Print() {
	fmt.Printf("Active Index: %d\n", v.ActiveIndex)
	fmt.Printf("Volume: %d%%\n", v.VolumePerc)
	fmt.Printf("Muted: %v\n", v.MutedState)
	fmt.Printf("Bluetooth? %v\n", v.BluetoothState)
}

// Printf - Volume Info Method used to output data
//   in a formatted structure
func (v *VolumeInfo) Printf(iconClr string, strClr string) {
	// SETUP EMOJI VARIABLES
	BlueIcon := ""    // Bluetooth Emoji
	MutedIcon := ""   // Muted Emoji
	Volume0Icon := "" // Low Volume
	Volume1Icon := "" // Medium Volume
	Volume2Icon := "" // High Volume
	VolumeIcon := ""   // Active Volume Emoji

	// FIGURE OUT WHICH EMOJIS TO USE
	if v.MutedState == true { // Muted
		fmt.Printf("%%{F%s}%v\n", iconClr, MutedIcon) // Output right away
		return
	} else if v.BluetoothState == true { // Bluteooth
		VolumeIcon = BlueIcon
	} else if v.VolumePerc > 60 { // Volume 60+
		VolumeIcon = Volume2Icon
	} else if v.VolumePerc > 35 { // Volume 35+
		VolumeIcon = Volume1Icon
	} else { // Volume <35
		VolumeIcon = Volume0Icon
	}

	// OUTPUT FORMATTED OUTPUT
	fmt.Printf("%v %%{F%s}%d%%\n", VolumeIcon, strClr, v.VolumePerc)
}
