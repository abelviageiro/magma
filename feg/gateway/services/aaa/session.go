/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

// Package aaa provides Carrier WiFi related services
//
//go:generate protoc -I protos --go_out=plugins=grpc,paths=source_relative:protos protos/context.proto protos/eap.proto
//go:generate protoc -I protos --go_out=plugins=grpc,paths=source_relative:protos protos/accounting.proto
//go:generate protoc -I protos --go_out=plugins=grpc,paths=source_relative:protos protos/authorization.proto
//
package aaa

import (
	"fmt"
	"math/rand"
	"time"

	"magma/feg/gateway/services/aaa/protos"
)

const (
	MinimalSessionTimeout = time.Millisecond * 10
	DefaultSessionTimeout = time.Hour * 6
)

// CreateSessionId creates & returns unique session ID string
func CreateSessionId() string {
	return fmt.Sprintf("%X-%X", time.Now().UnixNano()>>16, rand.Uint32())
}

// Session - struct to save an authenticated session state
type Session interface {
	// Lock - locks the Session's mutex
	Lock()
	// Unlock - unlocks the Session's mutex
	Unlock()
	// GetCtx returns AAA Session Context
	GetCtx() *protos.Context
	// SetCtx sets AAA Session Context - must be called on a Locked session
	SetCtx(ctx *protos.Context)
	// StopTimeout - stops the session's timeout if possible, returns if the timeout was successfully stopped
	StopTimeout() bool
}

// SessionTable - synchronized map of authenticated sessions
type SessionTable interface {
	// AddSession - adds a new session to the table & returns the newly created session pointer.
	// If a session with the same ID already is in the table - returns "Session with SID: XYZ already exist" as well as the
	// existing session.
	AddSession(pc *protos.Context, tout time.Duration, overwrite ...bool) (s Session, err error)
	// GetSession returns session corresponding to the given sid or nil if not found
	GetSession(sid string) (lockedSession Session)
	// RemoveSession - removes the session with the given SID and returns it, returns nil if not found
	RemoveSession(sid string) Session
	// SetTimeout - [Re]sets the session's cleanup timeout to fire after tout duration
	SetTimeout(sid string, tout time.Duration) bool
}
