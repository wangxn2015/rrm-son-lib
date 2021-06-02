// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package device

import (
	"github.com/onosproject/rrm-son-lib/pkg/model/id"
	meastype "github.com/onosproject/rrm-son-lib/pkg/model/measurement/type"
)

// NewCell returns cell object
func NewCell(ecgi id.ID, a3offset meastype.A3OffsetRange, hyst meastype.HysteresisRange,
	cellOffset meastype.QOffsetRange, freqOffset meastype.QOffsetRange, ttt meastype.TimeToTriggerRange) Cell {
	return &CellImpl{
		ECGI:                 ecgi,
		DeviceType:           DeviceCell,
		A3Offset:             a3offset,
		Hysteresis:           hyst,
		CellIndividualOffset: cellOffset,
		FrequencyOffset:      freqOffset,
		TimeToTrigger:        ttt,
	}
}

// Cell is the interface for Cell device
type Cell interface {
	Device
	GetA3Offset() meastype.A3OffsetRange
	GetHysteresis() meastype.HysteresisRange
	GetCellIndividualOffset() meastype.QOffsetRange
	GetFrequencyOffset() meastype.QOffsetRange
	GetTimeToTrigger() meastype.TimeToTriggerRange
	SetA3Offset(offsetRange meastype.A3OffsetRange)
	SetHysteresis(hysteresisRange meastype.HysteresisRange)
	SetCellIndividualOffset(offsetRange meastype.QOffsetRange)
	SetFrequencyOffset(offsetRange meastype.QOffsetRange)
	SetTimeToTrigger(triggerRange meastype.TimeToTriggerRange)
}

// CellImpl is the struct for Cell implementation
type CellImpl struct {
	ECGI                 id.ID
	DeviceType           Type
	A3Offset             meastype.A3OffsetRange
	Hysteresis           meastype.HysteresisRange
	CellIndividualOffset meastype.QOffsetRange
	FrequencyOffset      meastype.QOffsetRange
	TimeToTrigger        meastype.TimeToTriggerRange
}

// GetID returns ID
func (c *CellImpl) GetID() id.ID {
	return c.ECGI
}

// GetType returns device type
func (c *CellImpl) GetType() Type {
	return c.DeviceType
}

// SetID sets ID
func (c *CellImpl) SetID(i id.ID) {
	c.ECGI = i
}

// SetType sets device type
func (c *CellImpl) SetType(deviceType Type) {
	c.DeviceType = deviceType
}

// GetA3Offset returns a3 offset
func (c *CellImpl) GetA3Offset() meastype.A3OffsetRange {
	return c.A3Offset
}

// GetHysteresis returns hysteresis
func (c *CellImpl) GetHysteresis() meastype.HysteresisRange {
	return c.Hysteresis
}

// GetCellIndividualOffset returns cell individual offset
func (c *CellImpl) GetCellIndividualOffset() meastype.QOffsetRange {
	return c.CellIndividualOffset
}

// GetFrequencyOffset returns frqeuency offset
func (c *CellImpl) GetFrequencyOffset() meastype.QOffsetRange {
	return c.FrequencyOffset
}

// GetTimeToTrigger returns time to trigger
func (c *CellImpl) GetTimeToTrigger() meastype.TimeToTriggerRange {
	return c.TimeToTrigger
}

// SetA3Offset sets a3 offset
func (c *CellImpl) SetA3Offset(offsetRange meastype.A3OffsetRange) {
	c.A3Offset = offsetRange
}

// SetHysteresis sets hysteresis
func (c *CellImpl) SetHysteresis(hysteresisRange meastype.HysteresisRange) {
	c.Hysteresis = hysteresisRange
}

// SetCellIndividualOffset sets cell individual offset
func (c *CellImpl) SetCellIndividualOffset(offsetRange meastype.QOffsetRange) {
	c.CellIndividualOffset = offsetRange
}

// SetFrequencyOffset sets frequency offset
func (c *CellImpl) SetFrequencyOffset(offsetRange meastype.QOffsetRange) {
	c.FrequencyOffset = offsetRange
}

// SetTimeToTrigger sets time to trigger
func (c *CellImpl) SetTimeToTrigger(triggerRange meastype.TimeToTriggerRange) {
	c.TimeToTrigger = triggerRange
}
