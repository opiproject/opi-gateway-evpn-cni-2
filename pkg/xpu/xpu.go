// SPDX-License-Identifier: Apache-2.0
// Copyright (C) 2023 Nordix Foundation.
// Copyright (c) 2024 Ericsson AB.

// Package xpu handles the configuration of IPU/DPU cards
package xpu

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	godpunet "github.com/opiproject/godpu/network"
	evpngwtypes "github.com/opiproject/opi-gateway-evpn-cni/pkg/types"

	xpuMgr "github.com/opiproject/opi-api/network/evpn-gw/v1alpha1/gen/go"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateBridgePort creates a bridge port
func CreateBridgePort(conf *evpngwtypes.NetConf, mac string) error {
	var typeOfPort string
	var logicalBridges []string

	// Get a Client
	client, err := godpunet.NewBridgePort(conf.OpiEvpnBridgeConn)
	if err != nil {
		return fmt.Errorf("CreateBridgePort: Error occurred while initializing client connection:  %q", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	if conf.LogicalBridge != "" {
		typeOfPort = "access"
		logicalBridges = []string{conf.LogicalBridge}
	} else {
		typeOfPort = "trunk"
		if len(conf.LogicalBridges) > 0 {
			logicalBridges = conf.LogicalBridges
		}
	}

	// grpc call to create the bridge port
	bridgePort, err := client.CreateBridgePort(ctx, "", mac, typeOfPort, logicalBridges)
	if err != nil {
		return fmt.Errorf("CreateBridgePort: Error occurred while creating Bridge Port: %q", err)
	}

	// storing the name of the created bridge port to the netconf object for caching purposes
	// The format of the name is //network.opiproject.org/ports/94c353e4-ff01-41b4-ac20-285a45688e40
	// That means that we need to do split in  order to take the last part of it. Otherwise the Get
	// request will fail.
	nameSlice := strings.Split(bridgePort.GetName(), "/")
	conf.BridgePortName = nameSlice[len(nameSlice)-1]

	// retrying until the status of the created bridge port is up
	for i := 1; i <= 10; i++ {
		bridgePort, err := client.GetBridgePort(ctx, conf.BridgePortName)
		if err != nil {
			return fmt.Errorf("CreateBridgePort: Error occurred while getting Bridge Port %s : %q", conf.BridgePortName, err)
		}

		if bridgePort.GetStatus().GetOperStatus() == xpuMgr.BPOperStatus_BP_OPER_STATUS_UP {
			break
		}

		if i == 10 {
			return errors.New("CreateBridgePort: The status of created BridgePort is not UP")
		}

		time.Sleep(1 * time.Second)
	}

	return nil
}

// DeleteBridgePort deletes a bridge port
func DeleteBridgePort(conf *evpngwtypes.NetConf) error {
	// Check if the BridgePortName exists in the NetConf object.
	// If it doesn't exist then we simply return nil as there is no point to continue
	// as we need the BridgePortName for the BridgePort delete process to execute.
	// The reason that we do not return error is because we want to give the chance
	// to the delete process to continue with the rest of the tasks
	// (e.g. ReleaseVFs, ResetVFs, etc...) so there is no leftovers in the system.
	if conf.BridgePortName == "" {
		return nil
	}

	// Get a Client
	client, err := godpunet.NewBridgePort(conf.OpiEvpnBridgeConn)
	if err != nil {
		return fmt.Errorf("DeleteBridgePort: Error occurred while initializing client connection:  %q", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	// Delete Bridge Port. If Bridge Port is missing then we consider the
	// deletion request successful as our system is idempotent.
	_, err = client.DeleteBridgePort(ctx, conf.BridgePortName, true)
	if err != nil {
		return fmt.Errorf("DeleteBridgePort: Error occurred while deleting Bridge Port %s : %q", conf.BridgePortName, err)
	}

	// Checking if the bridge port is actually deleted
	for i := 1; i <= 10; i++ {
		_, err := client.GetBridgePort(ctx, conf.BridgePortName)
		if err != nil {
			if status.Code(err) == codes.NotFound {
				break
			}
			return fmt.Errorf("DeleteBridgePort: Error occurred while getting Bridge Port %s : %q", conf.BridgePortName, err)
		}
		if i == 10 {
			return fmt.Errorf("DeleteBridgePort: Bridge Port %s has not been deleted successfully", conf.BridgePortName)
		}
		time.Sleep(1 * time.Second)
	}

	return nil
}
