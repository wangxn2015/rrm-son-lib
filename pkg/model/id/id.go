// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package id

// ID is the interface for ID
type ID interface {
	String() string
	GetType() Type
	GetID() interface{}
}
