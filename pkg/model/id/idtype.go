// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package id

// Type is the ID type
type Type int

const (
	// TypeIDECGI is the ID type for ECGI
	TypeIDECGI Type = iota
	// TypeIDUEID is the ID type for UEID
	TypeIDUEID
)

var typeList = []string{
	"TypeIDECGI",
	"TypeIDUEID",
}

// String is the function to return type name
func (t *Type) String() string {
	return typeList[*t]
}
