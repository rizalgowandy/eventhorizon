// Copyright (c) 2014 - Max Ekman <max@looplab.se>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package domain

import (
	"github.com/looplab/eventhorizon"
)

const (
	CreateInviteCommand  eventhorizon.CommandType = "CreateInvite"
	AcceptInviteCommand                           = "AcceptInvite"
	DeclineInviteCommand                          = "DeclineInvite"
)

// CreateInvite is a command for creating invites.
type CreateInvite struct {
	InvitationID eventhorizon.UUID
	Name         string
	Age          int `eh:"optional"`
}

func (c *CreateInvite) AggregateID() eventhorizon.UUID            { return c.InvitationID }
func (c *CreateInvite) AggregateType() eventhorizon.AggregateType { return InvitationAggregateType }
func (c *CreateInvite) CommandType() eventhorizon.CommandType     { return CreateInviteCommand }

// AcceptInvite is a command for accepting invites.
type AcceptInvite struct {
	InvitationID eventhorizon.UUID
}

func (c *AcceptInvite) AggregateID() eventhorizon.UUID            { return c.InvitationID }
func (c *AcceptInvite) AggregateType() eventhorizon.AggregateType { return InvitationAggregateType }
func (c *AcceptInvite) CommandType() eventhorizon.CommandType     { return AcceptInviteCommand }

// DeclineInvite is a command for declining invites.
type DeclineInvite struct {
	InvitationID eventhorizon.UUID
}

func (c *DeclineInvite) AggregateID() eventhorizon.UUID            { return c.InvitationID }
func (c *DeclineInvite) AggregateType() eventhorizon.AggregateType { return InvitationAggregateType }
func (c *DeclineInvite) CommandType() eventhorizon.CommandType     { return DeclineInviteCommand }