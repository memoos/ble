package dev

import (
	"github.com/memoos/ble"
	"github.com/memoos/ble/linux"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return linux.NewDevice(opts...)
}
