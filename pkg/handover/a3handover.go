// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package handover

import (
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/rrm-son-lib/pkg/model/device"
	"github.com/onosproject/rrm-son-lib/pkg/model/measurement"
	"time"
)

var log = logging.GetLogger("rrm-son-lib", "handover", "a3")

// NewA3HandoverHandler returns A3HandoverHandler object
func NewA3HandoverHandler() *A3HandoverHandler {
	return &A3HandoverHandler{
		HandoverMap: make(map[string]A3HandoverTimer),
		Chans: A3HandoverChannel{
			InputChan:  make(chan device.UE),
			OutputChan: make(chan A3HandoverDecision),
		},
	}
}

// A3HandoverHandler is A3 handover handler
type A3HandoverHandler struct {
	HandoverMap map[string]A3HandoverTimer
	Chans       A3HandoverChannel
}

// A3HandoverChannel struct has channels used in A3 handover handler
type A3HandoverChannel struct {
	InputChan  chan device.UE
	OutputChan chan A3HandoverDecision
}

// Run starts A3 handover handler
func (h *A3HandoverHandler) Run() {
	for ue := range h.Chans.InputChan {
		ttt := ue.GetSCell().GetTimeToTrigger()

		if ttt.GetValue().(int) == 0 {
			h.runWithZeroTTT(ue)
		} else {
			h.runWithNonZeroTTT(ue, ttt.GetValue().(int))
		}
	}
}

func (h *A3HandoverHandler) runWithZeroTTT(ue device.UE) {
	hoDecision := NewA3HandoverDecision(ue, h.getTargetCell(ue))
	h.Chans.OutputChan <- hoDecision
	log.Debugf("Handover - %v", hoDecision)
}

func (h *A3HandoverHandler) runWithNonZeroTTT(ue device.UE, ttt int) {
	if _, ok := h.HandoverMap[ue.GetID().String()]; !ok {
		h.HandoverMap[ue.GetID().String()] = NewA3HandoverTimer(ue)
		go func() {
			startTime := time.Now()
			targetCellID := h.getTargetCell(ue).GetID().String()
			for {
				select {
				case <-time.After(time.Duration(ttt) * time.Millisecond):
					// no handover
					return
				case ue := <-h.HandoverMap[ue.GetID().String()].TimerChan:
					tmpTime := time.Now()
					eTime := tmpTime.Sub(startTime)
					tmpTargetCell := h.getTargetCell(ue)
					tmpTargetCellID := tmpTargetCell.GetID().String()

					// target cell changed - reset timer and target cell
					if tmpTargetCellID != targetCellID {
						startTime = time.Now()
						targetCellID = tmpTargetCellID
						continue
					}

					// if still same target cell
					if tmpTargetCellID == targetCellID && eTime.Milliseconds() >= time.Duration(ttt).Milliseconds() {
						// if next report arrives after TTT
						hoDecision := NewA3HandoverDecision(ue, tmpTargetCell)
						h.Chans.OutputChan <- hoDecision
						log.Debugf("Handover - %v", hoDecision)
						return
					}
				}
			}
		}()
	}
}

func (h *A3HandoverHandler) getTargetCell(ue device.UE) device.Cell {
	var targetCell device.Cell
	var bestRSRP measurement.RSRP
	flag := false
	for _, cscell := range ue.GetCSCells() {
		tmpRSRP := ue.GetMeasurements()[cscell.GetID().String()].GetMeasurement().(measurement.RSRP)
		if !flag {
			targetCell = cscell
			bestRSRP = tmpRSRP
			flag = true
			continue
		}

		if tmpRSRP > bestRSRP {
			targetCell = cscell
			bestRSRP = tmpRSRP
		}
	}
	return targetCell
}

// NewA3HandoverTimer returns A3HandoverTimer object
func NewA3HandoverTimer(ue device.UE) A3HandoverTimer {
	return A3HandoverTimer{
		UE:        ue,
		TimerChan: make(chan device.UE),
	}
}

// A3HandoverTimer struct is for A3 handover timer
type A3HandoverTimer struct {
	UE        device.UE
	TimerChan chan device.UE
}

// NewA3HandoverDecision returns A3HandoverDecision object
func NewA3HandoverDecision(ue device.UE, targetCell device.Cell) A3HandoverDecision {
	return A3HandoverDecision{
		UE:          ue,
		ServingCell: ue.GetSCell(),
		TargetCell:  targetCell,
	}
}

// A3HandoverDecision struct has A3 handover decision information
type A3HandoverDecision struct {
	UE          device.UE
	ServingCell device.Cell
	TargetCell  device.Cell
}
