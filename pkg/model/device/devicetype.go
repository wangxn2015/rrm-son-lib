// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package device

const (
	// DeviceUE is the type for UE
	DeviceUE Type = iota
	// DeviceCell is the type for Cell
	DeviceCell
)

// Type is the type of device
type Type int

var strListDeviceType = []string{
	"DeviceUE",
	"DeviceCell",
}

// String returns value as string type
func (t *Type) String() string {
	return strListDeviceType[*t]
}
