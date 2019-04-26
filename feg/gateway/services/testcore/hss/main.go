/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

// This starts the home subscriber server (hss) service.
package main

import (
	"context"
	"log"

	"magma/feg/cloud/go/protos"
	"magma/feg/gateway/registry"
	"magma/feg/gateway/services/testcore/hss/servicers"
	"magma/feg/gateway/services/testcore/hss/storage"
	"magma/feg/gateway/streamer"
	"magma/orc8r/cloud/go/service"

	"github.com/golang/glog"
)

func main() {
	srv, err := service.NewServiceWithOptions(registry.ModuleName, registry.MOCK_HSS)
	if err != nil {
		log.Fatalf("Error creating hss service: %s", err)
	}
	config, err := servicers.GetHSSConfig()
	if err != nil {
		log.Fatalf("Error getting hss config: %s", err)
	}
	store := storage.NewMemorySubscriberStore()
	servicer, err := servicers.NewHomeSubscriberServer(store, config)
	if err != nil {
		log.Fatalf("Error creating home subscriber server: %s", err)
	}
	protos.RegisterHSSConfiguratorServer(srv.GrpcServer, servicer)

	if config.StreamSubscribers {
		streamerClient := streamer.NewStreamerClient(registry.NewCloudRegistry())
		if err = streamerClient.AddListener(storage.NewSubscriberListener(store)); err != nil {
			glog.Errorf("Failed to start subscriber streaming: %s", err.Error())
		}
	}

	subscribers, err := servicers.GetConfiguredSubscribers()
	if err != nil {
		glog.Errorf("Could not fetch preconfigured subscribers: %s", err)
	} else {
		// Add preconfigured subscribers
		for _, sub := range subscribers {
			_, err = servicer.AddSubscriber(context.Background(), sub)
			if err != nil {
				glog.Errorf("Error adding subscriber: %s", err)
			}
		}
	}
	// Start diameter server
	go func() {
		glog.V(2).Info("Starting home subscriber server")
		err := servicer.Start(make(chan struct{})) // blocks
		glog.Error(err)
	}()

	// Run the service
	err = srv.Run()
	if err != nil {
		log.Fatalf("Error running hss service: %s", err)
	}
}