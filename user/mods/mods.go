package mods

//TODO: DO NOT DELETE OR MANUALLY EDIT THIS FILE
// This file is automatically generated for mod functionality

import (
	"fmt"
	"time"
)

func Init() {
	startTime := time.Now()
	fmt.Printf("\n[MODLOADER : BEGIN]\n")

	endTime := time.Now()
	fmt.Printf("[MODLOADER : COMPLETE] %s\n\n", endTime.Sub(startTime))
}