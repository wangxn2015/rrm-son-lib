// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package measurement

import (
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/rrm-son-lib/pkg/model/device"
	"github.com/onosproject/rrm-son-lib/pkg/model/measurement"
	meastype "github.com/onosproject/rrm-son-lib/pkg/model/measurement/type"
)

var log = logging.GetLogger("rrm-son-lib", "measurement", "eventa3")

// A3Status is the status if the UE is in the event a3 status
type A3Status bool

// NewMeasEventA3Handler returns MeasEventA3Handler object
func NewMeasEventA3Handler() *MeasEventA3Handler {
	return &MeasEventA3Handler{
		EventMap: make(map[string]MeasEventA3Obj),
		Chans: MeasEventA3Channel{
			InputChan:  make(chan device.UE),
			OutputChan: make(chan device.UE),
		},
	}
}

// MeasEventA3Handler is Event a3 handler
type MeasEventA3Handler struct {
	Chans    MeasEventA3Channel
	EventMap map[string]MeasEventA3Obj
}

// MeasEventA3Channel has all channels used in Event A3 handler
type MeasEventA3Channel struct {
	InputChan  chan device.UE
	OutputChan chan device.UE
}

// Run starts handler
func (h *MeasEventA3Handler) Run() {
	for ue := range h.Chans.InputChan {
		h.updateA3EventMap(ue)
		obj := h.EventMap[ue.GetID().String()]
		if obj.isInA3Event() {
			log.Debugf("UE %v is in A3 event - report through output channel: UE info - %v", ue.GetID().String(), ue)
			h.Chans.OutputChan <- ue
		}
	}
}

func (h *MeasEventA3Handler) updateA3EventMap(ue device.UE) {
	if _, ok := h.EventMap[ue.GetID().String()]; !ok {
		h.EventMap[ue.GetID().String()] = NewMeasEventA3Obj(ue)
	}

	obj := h.EventMap[ue.GetID().String()]
	log.Debugf("A3 Object [id: %v]: %v", ue.GetID().String(), ue)

	// update A3StatusMap
	tmpA3StatusMap := obj.A3StatusMap
	for _, cscell := range ue.GetCSCells() {
		if _, ok := tmpA3StatusMap[cscell.GetID().String()]; !ok {
			obj.A3StatusMap[cscell.GetID().String()] = h.isEnteringA3Condition(ue, cscell)
			continue
		}

		if tmpA3StatusMap[cscell.GetID().String()] {
			obj.A3StatusMap[cscell.GetID().String()] = !h.isLeavingA3Condition(ue, cscell)
		} else {
			obj.A3StatusMap[cscell.GetID().String()] = h.isEnteringA3Condition(ue, cscell)
		}
	}
}

func (h *MeasEventA3Handler) isEnteringA3Condition(ue device.UE, cell device.Cell) A3Status {
	mp := ue.GetMeasurements()[ue.GetSCell().GetID().String()].GetMeasurement().(measurement.RSRP)
	ofp := ue.GetSCell().GetFrequencyOffset()
	ocp := ue.GetSCell().GetCellIndividualOffset()
	a3offset := ue.GetSCell().GetA3Offset()
	hyst := ue.GetSCell().GetHysteresis()

	mn := ue.GetMeasurements()[cell.GetID().String()].GetMeasurement().(measurement.RSRP)
	ofn := cell.GetFrequencyOffset()
	ocn := cell.GetCellIndividualOffset()

	return float64(mn)+float64(ofn.GetValue().(int))+float64(ocn.GetValue().(int))-float64(hyst.GetValue().(meastype.HysteresisRange)) >
		float64(mp)+float64(ofp.GetValue().(int))+float64(ocp.GetValue().(int))+float64(a3offset.GetValue().(meastype.A3OffsetRange))
}

func (h *MeasEventA3Handler) isLeavingA3Condition(ue device.UE, cell device.Cell) A3Status {
	mp := ue.GetMeasurements()[ue.GetSCell().GetID().String()].GetMeasurement().(measurement.RSRP)
	ofp := ue.GetSCell().GetFrequencyOffset()
	ocp := ue.GetSCell().GetCellIndividualOffset()
	a3offset := ue.GetSCell().GetA3Offset()
	hyst := ue.GetSCell().GetHysteresis()

	mn := ue.GetMeasurements()[cell.GetID().String()].GetMeasurement().(measurement.RSRP)
	ofn := cell.GetFrequencyOffset()
	ocn := cell.GetCellIndividualOffset()

	return float64(mn)+float64(ofn.GetValue().(int))+float64(ocn.GetValue().(int))+float64(hyst.GetValue().(meastype.HysteresisRange)) <
		float64(mp)+float64(ofp.GetValue().(int))+float64(ocp.GetValue().(int))+float64(a3offset.GetValue().(meastype.A3OffsetRange))
}

// NewMeasEventA3Obj returns MeasEventA3Obj
func NewMeasEventA3Obj(ue device.UE) MeasEventA3Obj {
	a3StatusMap := make(map[string]A3Status)
	// init a3status map
	for _, cell := range ue.GetCSCells() {
		a3StatusMap[cell.GetID().String()] = false
	}
	return MeasEventA3Obj{
		UE:          ue,
		A3StatusMap: a3StatusMap,
	}
}

// MeasEventA3Obj is the struct for Event A3 record
type MeasEventA3Obj struct {
	UE          device.UE
	A3StatusMap map[string]A3Status
}

func (o *MeasEventA3Obj) isInA3Event() A3Status {
	for _, v := range o.A3StatusMap {
		if v {
			return true
		}
	}
	return false
}
