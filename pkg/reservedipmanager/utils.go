// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package reservedipmanager

import (
	"net"

	spiderpoolip "github.com/spidernet-io/spiderpool/pkg/ip"
	spiderpoolv1 "github.com/spidernet-io/spiderpool/pkg/k8s/apis/spiderpool.spidernet.io/v1"
	spiderpooltypes "github.com/spidernet-io/spiderpool/pkg/types"
)

func GetReservedIPsByIPVersion(version spiderpooltypes.IPVersion, rIPList *spiderpoolv1.SpiderReservedIPList) ([]net.IP, error) {
	var ips []net.IP
	for _, r := range rIPList.Items {
		if *r.Spec.IPVersion != version {
			continue
		}

		rIPs, err := spiderpoolip.ParseIPRanges(version, r.Spec.IPs)
		if err != nil {
			return nil, err
		}
		ips = append(ips, rIPs...)
	}

	return ips, nil
}
