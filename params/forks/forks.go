// Copyright 2023 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package forks

// Fork is a numerical identifier of specific network upgrades (forks).
type Fork int

const (
	Frontier Fork = iota
	FrontierThawing
	Homestead
	DAO
	TangerineWhistle
	SpuriousDragon
	Byzantium
	Constantinople
	Petersburg
	Istanbul
	MuirGlacier
	Berlin
	London
	ArrowGlacier
	GrayGlacier
	Paris
	Shanghai
	Cancun
	Prague
	Osaka
)

// String implements fmt.Stringer.
func (f Fork) String() string {
	s, ok := forkToString[f]
	if !ok {
		return "Unknown fork"
	}
	return s
}

var forkToString = map[Fork]string{
	Frontier:         "Frontier",
	FrontierThawing:  "Frontier Thawing",
	Homestead:        "Homestead",
	DAO:              "DAO",
	TangerineWhistle: "Tangerine Whistle",
	SpuriousDragon:   "Spurious Dragon",
	Byzantium:        "Byzantium",
	Constantinople:   "Constantinople",
	Petersburg:       "Petersburg",
	Istanbul:         "Istanbul",
	MuirGlacier:      "Muir Glacier",
	Berlin:           "Berlin",
	London:           "London",
	ArrowGlacier:     "Arrow Glacier",
	GrayGlacier:      "Gray Glacier",
	Paris:            "Paris",
	Shanghai:         "Shanghai",
	Cancun:           "Cancun",
	Prague:           "Prague",
	Osaka:            "Osaka",
}
