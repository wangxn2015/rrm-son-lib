// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package measurement

import (
	"fmt"
	"github.com/onosproject/rrm-son-lib/pkg/model/device"
	"github.com/onosproject/rrm-son-lib/pkg/model/id"
	"github.com/onosproject/rrm-son-lib/pkg/model/measurement"
	meastype "github.com/onosproject/rrm-son-lib/pkg/model/measurement/type"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestEventA3Handler_WithUEInEventA3(t *testing.T) {
	ueChan := make(chan device.UE)
	eventa3handler := NewMeasEventA3Handler()
	// run event a3 handler
	go eventa3handler.Run()
	//start ue generation
	go UEGeneratorEventA3(ueChan)
	go func() {
		for ue := range ueChan {
			eventa3handler.Chans.InputChan <- ue
		}
	}()

	counter := 0
	for {
		select {
		case <-time.After(20 * time.Second):
			assert.FailNow(t, "Event a3 was not happening")
		case a3event := <-eventa3handler.Chans.OutputChan:
			counter++
			fmt.Println(a3event)
			if counter == 10 {
				return
			}
		}
	}
}

func UEGeneratorEventA3(ch chan device.UE) {
	ueid := id.NewUEID(1, 2, 3)
	ecgi := id.NewECGI(3)
	scell := device.NewCell(ecgi, meastype.A3OffsetRange(1), meastype.HysteresisRange(2), meastype.QOffset3dB, meastype.QOffset4dB, meastype.TTT0ms)
	ecgiCScell1 := id.NewECGI(4)
	ecgiCScell2 := id.NewECGI(5)
	ecgiCScell3 := id.NewECGI(6)
	cscell1 := device.NewCell(ecgiCScell1, meastype.A3OffsetRange(1), meastype.HysteresisRange(2), meastype.QOffset3dB, meastype.QOffset4dB, meastype.TTT0ms)
	cscell2 := device.NewCell(ecgiCScell2, meastype.A3OffsetRange(1), meastype.HysteresisRange(2), meastype.QOffset3dB, meastype.QOffset4dB, meastype.TTT0ms)
	cscell3 := device.NewCell(ecgiCScell3, meastype.A3OffsetRange(1), meastype.HysteresisRange(2), meastype.QOffset3dB, meastype.QOffset4dB, meastype.TTT0ms)
	cscellList := []device.Cell{cscell1, cscell2, cscell3}
	ue := device.NewUE(ueid, scell, cscellList)
	measSCell := measurement.NewMeasEventA3(ecgi, measurement.RSRP(-80))
	measCSCell1 := measurement.NewMeasEventA3(ecgiCScell1, measurement.RSRP(-70))
	measCSCell2 := measurement.NewMeasEventA3(ecgiCScell2, measurement.RSRP(-50))
	measCSCell3 := measurement.NewMeasEventA3(ecgiCScell3, measurement.RSRP(-60))
	ue.GetMeasurements()[measSCell.GetCellID().String()] = measSCell
	ue.GetMeasurements()[measCSCell1.GetCellID().String()] = measCSCell1
	ue.GetMeasurements()[measCSCell2.GetCellID().String()] = measCSCell2
	ue.GetMeasurements()[measCSCell3.GetCellID().String()] = measCSCell3

	for i := 0; i < 10; i++ {
		ch <- ue
		time.Sleep(1 * time.Second)
	}

	fmt.Println("UE generation is done")
}

func TestEventA3Handler_WithoutUEInEventA3(t *testing.T) {
	ueChan := make(chan device.UE)
	eventa3handler := NewMeasEventA3Handler()
	// run event a3 handler
	go eventa3handler.Run()
	//start ue generation
	go UEGeneratorNonEventA3(ueChan)
	go func() {
		for ue := range ueChan {
			eventa3handler.Chans.InputChan <- ue
		}
	}()

	for {
		select {
		case <-time.After(20 * time.Second):
			return
		case a3event := <-eventa3handler.Chans.OutputChan:
			fmt.Println(a3event)
			assert.FailNow(t, "Event a3 should not be arrived")
		}
	}
}

func UEGeneratorNonEventA3(ch chan device.UE) {
	ueid := id.NewUEID(1, 2, 3)
	ecgi := id.NewECGI(3)
	scell := device.NewCell(ecgi, meastype.A3OffsetRange(1), meastype.HysteresisRange(2), meastype.QOffset3dB, meastype.QOffset4dB, meastype.TTT0ms)
	ecgiCScell1 := id.NewECGI(4)
	ecgiCScell2 := id.NewECGI(5)
	ecgiCScell3 := id.NewECGI(6)
	cscell1 := device.NewCell(ecgiCScell1, meastype.A3OffsetRange(1), meastype.HysteresisRange(2), meastype.QOffset3dB, meastype.QOffset4dB, meastype.TTT0ms)
	cscell2 := device.NewCell(ecgiCScell2, meastype.A3OffsetRange(1), meastype.HysteresisRange(2), meastype.QOffset3dB, meastype.QOffset4dB, meastype.TTT0ms)
	cscell3 := device.NewCell(ecgiCScell3, meastype.A3OffsetRange(1), meastype.HysteresisRange(2), meastype.QOffset3dB, meastype.QOffset4dB, meastype.TTT0ms)
	cscellList := []device.Cell{cscell1, cscell2, cscell3}
	ue := device.NewUE(ueid, scell, cscellList)
	measSCell := measurement.NewMeasEventA3(ecgi, measurement.RSRP(-40))
	measCSCell1 := measurement.NewMeasEventA3(ecgiCScell1, measurement.RSRP(-70))
	measCSCell2 := measurement.NewMeasEventA3(ecgiCScell2, measurement.RSRP(-50))
	measCSCell3 := measurement.NewMeasEventA3(ecgiCScell3, measurement.RSRP(-60))
	ue.GetMeasurements()[measSCell.GetCellID().String()] = measSCell
	ue.GetMeasurements()[measCSCell1.GetCellID().String()] = measCSCell1
	ue.GetMeasurements()[measCSCell2.GetCellID().String()] = measCSCell2
	ue.GetMeasurements()[measCSCell3.GetCellID().String()] = measCSCell3

	for i := 0; i < 10; i++ {
		ch <- ue
		time.Sleep(1 * time.Second)
	}

	fmt.Println("UE generation is done")
}
