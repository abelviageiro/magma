/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

package analytics

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"os/user"
	"testing"

	"fbc/lib/go/libgraphql"

	"github.com/stretchr/testify/require"
)

func TestIntegration(t *testing.T) {
	accessToken, ok := os.LookupEnv("INTEGRATION_ACCESS_TOKEN")
	if !ok {
		t.SkipNow()
	}
	u, err := user.Current()
	require.NoError(t, err, "failed getting user")
	c := libgraphql.NewClient(libgraphql.ClientConfig{
		Token:    accessToken,
		Endpoint: fmt.Sprintf("https://graph.%s.sb.expresswifi.com/graphql", u.Username),
		HTTPClient: &http.Client{Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}},
	})

	t.Log("creating session")
	cop := NewCreateSessionOp(&RadiusSession{
		NASIPAddress:         "10.0.0.1",
		NASIdentifier:        "10.0.0.2",
		AcctSessionID:        "10.0.0.3:10.0.0.4",
		CalledStationID:      "10.0.0.5",
		FramedIPAddress:      "10.0.0.6",
		CallingStationID:     "10.0.0.7",
		NormalizedMacAddress: "10.0.0.8",
		RADIUSServerID:       1,
	})
	require.NoError(t, c.Do(cop), "failed creating session")
	t.Logf("radius_session_id:%d\n", cop.Response().FBID)

	t.Logf("updating session with id %d\n", cop.Response().FBID)
	uop := NewUpdateSessionOp(&RadiusSession{
		FBID:                 cop.Response().FBID,
		NASIPAddress:         "10.0.1.1",
		NASIdentifier:        "10.0.1.2",
		AcctSessionID:        "10.0.1.3:10.0.1.4",
		CalledStationID:      "10.0.1.5",
		FramedIPAddress:      "10.0.1.6",
		CallingStationID:     "10.0.1.7",
		NormalizedMacAddress: "10.0.1.8",
		RADIUSServerID:       2,
		UploadBytes:          1234,
		Vendor:               int64(Ruckus),
	})
	require.NoError(t, c.Do(uop), "failed updating session")
}
