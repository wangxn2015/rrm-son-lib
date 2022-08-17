// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package handover

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	measurement2 "github.com/wangxn2015/rrm-son-lib/pkg/measurement"
	"github.com/wangxn2015/rrm-son-lib/pkg/model/device"
	"github.com/wangxn2015/rrm-son-lib/pkg/model/id"
	"github.com/wangxn2015/rrm-son-lib/pkg/model/measurement"
	meastype "github.com/wangxn2015/rrm-son-lib/pkg/model/measurement/type"
	"testing"
	"time"
)

func TestA3HandoverWithEventA3Handler_ZeroTTT(t *testing.T) {
	ueChan := make(chan device.UE)
	eventa3handler := measurement2.NewMeasEventA3Handler()
	handoverhandler := NewA3HandoverHandler()
	// run event a3 handler
	go eventa3handler.Run()
	// start ue generation
	go UEGeneratorZeroTTT(ueChan)
	go func() {
		for ue := range ueChan {
			eventa3handler.Chans.InputChan <- ue
		}
	}()

	// run handover handler
	go handoverhandler.Run()
	// forward event a3 measurement only to handover handler
	go func() {
		for ue := range eventa3handler.Chans.OutputChan {
			handoverhandler.Chans.InputChan <- ue
		}
	}()

	// print out ho decision
	counter := 0
	for {
		select {
		case <-time.After(20 * time.Second):
			assert.FailNow(t, "HO decision messages did not arrive in time or not enough messages arrived")
			return
		case hoDecision := <-handoverhandler.Chans.OutputChan:
			counter++
			fmt.Println(hoDecision)
			if counter == 10 {
				return
			}
		}
	}
}

func UEGeneratorZeroTTT(ch chan device.UE) {
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

func TestA3HandoverWithEventA3Handler_TwoSecTTT(t *testing.T) {
	ueChan := make(chan device.UE)
	eventa3handler := measurement2.NewMeasEventA3Handler()
	handoverhandler := NewA3HandoverHandler()
	// run event a3 handler
	go eventa3handler.Run()
	// start ue generation
	go UEGeneratorTwoSecTTT(ueChan)
	go func() {
		for ue := range ueChan {
			eventa3handler.Chans.InputChan <- ue
		}
	}()

	// run handover handler
	go handoverhandler.Run()
	// forward event a3 measurement only to handover handler
	go func() {
		for ue := range eventa3handler.Chans.OutputChan {
			handoverhandler.Chans.InputChan <- ue
		}
	}()

	// print out ho decision
	counter := 0
	for {
		select {
		case <-time.After(20 * time.Second):
			assert.FailNow(t, "HO decision messages did not arrive in time or not enough messages arrived")
		case hoDecision := <-handoverhandler.Chans.OutputChan:
			counter++
			fmt.Println(hoDecision)
			if counter == 2 {
				return
			}
		}
	}
}

func UEGeneratorTwoSecTTT(ch chan device.UE) {
	ueid := id.NewUEID(1, 2, 3)
	ecgi := id.NewECGI(3)
	scell := device.NewCell(ecgi, meastype.A3OffsetRange(1), meastype.HysteresisRange(2), meastype.QOffset3dB, meastype.QOffset4dB, meastype.TTT2560ms)
	ecgiCScell1 := id.NewECGI(4)
	ecgiCScell2 := id.NewECGI(5)
	ecgiCScell3 := id.NewECGI(6)
	cscell1 := device.NewCell(ecgiCScell1, meastype.A3OffsetRange(1), meastype.HysteresisRange(2), meastype.QOffset3dB, meastype.QOffset4dB, meastype.TTT2560ms)
	cscell2 := device.NewCell(ecgiCScell2, meastype.A3OffsetRange(1), meastype.HysteresisRange(2), meastype.QOffset3dB, meastype.QOffset4dB, meastype.TTT2560ms)
	cscell3 := device.NewCell(ecgiCScell3, meastype.A3OffsetRange(1), meastype.HysteresisRange(2), meastype.QOffset3dB, meastype.QOffset4dB, meastype.TTT2560ms)
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
