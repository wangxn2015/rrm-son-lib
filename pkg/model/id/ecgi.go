// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package id

import "fmt"

// NewECGI returns ECGI object
func NewECGI(e uint64) ID {
	id := ECGI(e)
	return &id
}

// ECGI is the type for the ID ECGI
type ECGI uint64

// String returns ECGI string type
func (E *ECGI) String() string {
	return fmt.Sprintf("%d", *E)
}

// GetType returns ID type
func (E *ECGI) GetType() Type {
	return TypeIDECGI
}

// GetID returns ID as interface type
func (E *ECGI) GetID() interface{} {
	return *E
}
