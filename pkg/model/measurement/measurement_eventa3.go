// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package measurement

import "github.com/wangxn2015/rrm-son-lib/pkg/model/id"

// RSRP is the type for RSRP
type RSRP float64

// NewMeasEventA3 returns MeasurementEventA3 object
func NewMeasEventA3(ecgi id.ID, rsrp RSRP) Measurement {
	return &MeasEventA3{
		ECGI:          ecgi,
		RSRP:          rsrp,
		MeasEventType: EventA3,
	}
}

// MeasEventA3 is the struct for measurement event A3
type MeasEventA3 struct {
	ECGI          id.ID
	RSRP          RSRP
	MeasEventType MeasEventType
}

// GetMeasurementEventType returns measurement event a3
func (m *MeasEventA3) GetMeasurementEventType() MeasEventType {
	return m.MeasEventType
}

// GetCellID returns cell ID
func (m *MeasEventA3) GetCellID() id.ID {
	return m.ECGI
}

// GetMeasurement returns RSRP
func (m *MeasEventA3) GetMeasurement() interface{} {
	return m.RSRP
}

// SetMeasurementEventType sets measurement event type
func (m *MeasEventA3) SetMeasurementEventType(eventType MeasEventType) {
	m.MeasEventType = eventType
}

// SetCellID sets cell ID
func (m *MeasEventA3) SetCellID(i id.ID) {
	m.ECGI = i
}

// SetMeasurement sets measurement
func (m *MeasEventA3) SetMeasurement(i interface{}) {
	m.RSRP = i.(RSRP)
}
