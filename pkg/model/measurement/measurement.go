// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package measurement

import "github.com/onosproject/rrm-son-lib/pkg/model/id"

// Measurement is the interface for Measurement
type Measurement interface {
	GetMeasurementEventType() MeasEventType
	GetCellID() id.ID
	GetMeasurement() interface{}
	SetMeasurementEventType(eventType MeasEventType)
	SetCellID(id.ID)
	SetMeasurement(interface{})
}
