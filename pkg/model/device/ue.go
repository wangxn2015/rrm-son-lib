// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package device

import (
	"github.com/onosproject/rrm-son-lib/pkg/model/id"
	"github.com/onosproject/rrm-son-lib/pkg/model/measurement"
)

// NewUE returns UE object
func NewUE(ueid id.ID, scell Cell, cscells []Cell) UE {
	return &UEImpl{
		UEID:         ueid,
		DeviceType:   DeviceUE,
		SCell:        scell,
		CSCells:      cscells,
		Measurements: make(map[string]measurement.Measurement),
	}
}

// UE is the interface for UE device
type UE interface {
	Device
	GetSCell() Cell
	GetCSCells() []Cell
	GetMeasurements() map[string]measurement.Measurement
	SetScell(Cell)
	SetCSCells([]Cell)
	SetMeasurements(map[string]measurement.Measurement)
}

// UEImpl is the struct for UE implementation
type UEImpl struct {
	UEID         id.ID
	DeviceType   Type
	SCell        Cell
	CSCells      []Cell
	Measurements map[string]measurement.Measurement
}

// GetID returns ID
func (u *UEImpl) GetID() id.ID {
	return u.UEID
}

// GetType returns device type
func (u *UEImpl) GetType() Type {
	return u.DeviceType
}

// SetID sets ID
func (u *UEImpl) SetID(i id.ID) {
	u.UEID = i
}

// SetType sets device type
func (u *UEImpl) SetType(t Type) {
	u.DeviceType = t
}

// GetSCell returns serving cell
func (u *UEImpl) GetSCell() Cell {
	return u.SCell
}

// GetCSCells returns the list of candidate serving cells
func (u *UEImpl) GetCSCells() []Cell {
	return u.CSCells
}

// GetMeasurements returns measurement map
func (u *UEImpl) GetMeasurements() map[string]measurement.Measurement {
	return u.Measurements
}

// SetScell sets serving cell
func (u *UEImpl) SetScell(cell Cell) {
	u.SCell = cell
}

// SetCSCells sets the list of candidate serving cells
func (u *UEImpl) SetCSCells(cells []Cell) {
	u.CSCells = cells
}

// SetMeasurements sets the measurement map
func (u *UEImpl) SetMeasurements(m map[string]measurement.Measurement) {
	u.Measurements = m
}
