// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package device

import "github.com/onosproject/rrm-son-lib/pkg/model/id"

// Device is the interface for Device
type Device interface {
	GetID() id.ID
	GetType() Type
	SetID(id.ID)
	SetType(Type)
}
