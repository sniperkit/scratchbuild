package arch

import (
	"github.com/ryanuber/go-glob"
)

// KnownArch represents a known architecture.
type KnownArch int

const (
	NONE KnownArch = iota
	// AMD64 architecture.
	AMD64
	// ARMv6 architecture.
	ARMV6
	// ARMv7 architecture.
	ARMV7
	// ARMv8 64-bit ARM architecture.
	ARMV8
)

// DefaultArch is the architecture most Docker images are compatible with on default.
var DefaultArch KnownArch = AMD64

// KnownArchNames are the string regex representations of KnownArch.
var KnownArchNames = map[string]KnownArch{
	"arm":     ARMV6,
	"armv6l":  ARMV6,
	"armv7l":  ARMV7,
	"aarch64": ARMV8,
	"armv8*":  ARMV8,

	"x86_64": AMD64,
	"amd64":  AMD64,
	"i386":   AMD64, // iffy
}

// KnownArchCompat contains known compatibilities between architectures.
var KnownArchCompat = map[KnownArch][]KnownArch{
	// ARMV8 can use all ARM images.
	ARMV8: {ARMV6, ARMV7},
	// ARMV7 can use ARMV6 images.
	ARMV7: {ARMV6},
}

// ParseArch attempts to determine which arch the (uname -m) output represents.
func ParseArch(arch string) (KnownArch, bool) {
	for archp, archv := range KnownArchNames {
		if !glob.Glob(archp, arch) {
			continue
		}
		return archv, true
	}
	return AMD64, false
}
