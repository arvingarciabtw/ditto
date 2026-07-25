/*
Package base provides foundational types and data shared across keyboard
standards: Finger, Key, width constants, key definitions, and shift/altgr maps.
*/
package base

const (
	U0p75 = 2
	U1    = 3
	U1p5  = 4
	U1p75 = 5
	U2    = 6
	U2p5  = 7
	U2p75 = 8
	U3    = 9
	U3p5  = 10
	U3p75 = 11
	U4    = 12
	U4p5  = 13
	U5p75 = 17
	U6    = 18
	U7p5  = 22
	U7p75 = 23
)

const (
	Pinky Finger = iota
	Ring
	Middle
	Index
	Thumb
	Any
)

type Finger int

// TODO: need to refactor key rendering,
// Rightless and Leftless is a bit unclear.
// Perhaps do corners to make it more explicit.

type Key struct {
	Label             string
	Width             int
	Finger            Finger
	Gap               bool
	GapAlt            bool // super hacky, need to change this
	Rightless         bool
	Leftless          bool
	BottomLeftless    bool
	BottomLeftlessAlt bool // super hacky, need to change this
	DivLabel          string
	EvCode            uint16
}
