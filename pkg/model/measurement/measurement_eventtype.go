// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package measurement

const (
	// EventA1 is the Measurement Event A1
	EventA1 MeasEventType = iota
	// EventA2 is the Measurement Event A2
	EventA2
	// EventA3 is the Measurement Event A3
	EventA3
	// EventA4 is the Measurement Event A4
	EventA4
	// EventA5 is the Measurement Event A5
	EventA5
)

// MeasEventType is the type for Measurement Event
type MeasEventType int

var strListMeasEventType = []string{
	"EventA1",
	"EventA2",
	"EventA3",
	"EventA4",
	"EventA5",
}

// String returns value as string type
func (e *MeasEventType) String() string {
	return strListMeasEventType[*e]
}
