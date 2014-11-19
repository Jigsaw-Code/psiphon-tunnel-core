/*
 * Copyright (c) 2014, Psiphon Inc.
 * All rights reserved.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package main

import (
	"flag"
	psiphon "github.com/Psiphon-Labs/psiphon-tunnel-core/psiphon"
	"log"
)

func main() {
	var configFilename string
	flag.StringVar(&configFilename, "config", "", "configuration file")
	flag.Parse()
	if configFilename == "" {
		log.Fatalf("configuration file is required")
	}
	config, err := psiphon.LoadConfig(configFilename)
	if err != nil {
		log.Fatalf("error loading configuration file: %s", err)
	}
	psiphon.RunForever(config)
}
