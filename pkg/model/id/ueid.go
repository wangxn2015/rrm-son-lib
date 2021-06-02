// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package id

import "fmt"

// IMSI is the type for UE's ID
type IMSI uint64

// CRNTI is the other type for UE's ID
type CRNTI uint32

// NewUEID returns UEID object
func NewUEID(imsi uint64, crnti uint32, ecgi uint64) ID {
	return &UEID{
		IMSI:  IMSI(imsi),
		CRNTI: CRNTI(crnti),
		ECGI:  ECGI(ecgi),
	}
}

// UEID is the struct for UE's ID
type UEID struct {
	IMSI  IMSI
	CRNTI CRNTI
	ECGI  ECGI
}

// String returns UEID string type
func (U *UEID) String() string {
	return fmt.Sprintf("%d", *U)
}

// GetType returns ID type
func (U *UEID) GetType() Type {
	return TypeIDUEID
}

// GetID returns ID as interface type
func (U *UEID) GetID() interface{} {
	return *U
}
