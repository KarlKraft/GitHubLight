//
//  Settings.go
//  LightServer
//
//  Created by Karl Kraft on 12/30/2023
//  Copyright 2023-2024 Karl Kraft. All rights reserved
//

package arduino

import (
	"github.com/BurntSushi/toml"
	"log"
)

type Settings struct {
	BoxIP   string
	BoxPort uint16
}

func ReadSettings(path string) *Settings {
	s := &Settings{}
	meta, err := toml.DecodeFile(path, s)
	if err != nil {
		log.Fatalf("Unable to read file %s (%v)", path, err)
	}
	if !meta.IsDefined("BoxIP") {
		log.Fatalf("Specify BoxIP in configuration (e.g. BoxIP=\"192.168.1.59\"")
	}
	if !meta.IsDefined("BoxPort") {
		log.Fatalf("Specify BoxPort in configuration (e.g. BoxPort=49581")
	}
	return s
}
