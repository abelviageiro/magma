/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

package view_factory_test

import (
	"encoding/json"
	"os"
	"testing"

	"magma/orc8r/cloud/go/orc8r"
	"magma/orc8r/cloud/go/protos"
	"magma/orc8r/cloud/go/serde"
	checkinti "magma/orc8r/cloud/go/services/checkind/test_init"
	"magma/orc8r/cloud/go/services/checkind/test_utils"
	"magma/orc8r/cloud/go/services/config"
	configti "magma/orc8r/cloud/go/services/config/test_init"
	"magma/orc8r/cloud/go/services/magmad"
	storagetu "magma/orc8r/cloud/go/services/magmad/obsidian/handlers/test_utils"
	"magma/orc8r/cloud/go/services/magmad/obsidian/handlers/view_factory"
	magmadprotos "magma/orc8r/cloud/go/services/magmad/protos"
	magmadti "magma/orc8r/cloud/go/services/magmad/test_init"

	"github.com/stretchr/testify/assert"
)

func TestFullGatewayViewFactoryLegacyImpl_GetGatewayViewsForNetwork(t *testing.T) {
	os.Setenv(orc8r.UseConfiguratorEnv, "0")
	// Test setup
	magmadti.StartTestService(t)
	configti.StartTestService(t)
	checkinti.StartTestService(t)

	serde.UnregisterSerdesForDomain(t, config.SerdeDomain)
	err := serde.RegisterSerdes(storagetu.NewConfig1Manager(), storagetu.NewConfig2Manager())
	assert.NoError(t, err)

	// Setup fixture data
	networkID, err := magmad.RegisterNetwork(&magmadprotos.MagmadNetworkRecord{Name: "foobar"}, "xservice1")
	assert.NoError(t, err)

	// Register gateways
	record1 := &magmadprotos.AccessGatewayRecord{
		HwId: &protos.AccessGatewayID{Id: "hw1"},
	}
	record2 := &magmadprotos.AccessGatewayRecord{
		HwId: &protos.AccessGatewayID{Id: "hw2"},
	}
	_, err = magmad.RegisterGatewayWithId(networkID, record1, "gw1")
	assert.NoError(t, err)
	_, err = magmad.RegisterGatewayWithId(networkID, record2, "gw2")
	assert.NoError(t, err)

	// gw1 has cfg1 and cfg2, gw2 only has cfg2
	err = config.CreateConfig(
		networkID,
		storagetu.NewConfig1Manager().GetType(),
		"gw1",
		cfg1,
	)
	assert.NoError(t, err)
	err = config.CreateConfig(
		networkID,
		storagetu.NewConfig2Manager().GetType(),
		"gw1",
		cfg2,
	)
	assert.NoError(t, err)
	err = config.CreateConfig(
		networkID,
		storagetu.NewConfig2Manager().GetType(),
		"gw2",
		cfg2,
	)
	assert.NoError(t, err)

	// gw1 has status, gw2 does not
	checkinReq := &protos.CheckinRequest{
		GatewayId:       "hw1",
		MagmaPkgVersion: "1.2.3",
		Status: &protos.ServiceStatus{
			Meta: map[string]string{
				"hello": "world",
			},
		},
		SystemStatus: &protos.SystemStatus{
			CpuUser:   31498,
			CpuSystem: 8361,
			CpuIdle:   1869111,
			MemTotal:  1016084,
			MemUsed:   54416,
			MemFree:   412772,
		},
	}
	test_utils.Checkin(t, checkinReq)

	fact := &view_factory.FullGatewayViewFactoryLegacyImpl{}
	actual, err := fact.GetGatewayViewsForNetwork(networkID)
	assert.NoError(t, err)
	// Wipe out timestamps from status so we can compare the structs
	for _, state := range actual {
		if state.LegacyStatus != nil {
			state.LegacyStatus.CertExpirationTime = 0
			state.LegacyStatus.Time = 0
		}
	}

	expected := map[string]*view_factory.GatewayState{
		"gw1": {
			GatewayID: "gw1",
			Config: map[string]interface{}{
				storagetu.NewConfig1Manager().GetType(): cfg1,
				storagetu.NewConfig2Manager().GetType(): cfg2,
			},
			LegacyRecord: record1,
			LegacyStatus: &protos.GatewayStatus{
				Checkin:            checkinReq,
				Time:               0,
				CertExpirationTime: 0,
			},
		},
		"gw2": {
			GatewayID: "gw2",
			Config: map[string]interface{}{
				storagetu.NewConfig2Manager().GetType(): cfg2,
			},
			LegacyRecord: record2,
		},
	}

	marshaledExpected, err := json.Marshal(expected)
	assert.NoError(t, err)
	marshaledActual, err := json.Marshal(actual)
	assert.NoError(t, err)
	assert.Equal(t, marshaledExpected, marshaledActual)
}
