// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	alertchannel "github.com/sanmoh-hombal/provider-checkly/internal/controller/cluster/alerts/alertchannel"
	maintenancewindow "github.com/sanmoh-hombal/provider-checkly/internal/controller/cluster/alerts/maintenancewindow"
	check "github.com/sanmoh-hombal/provider-checkly/internal/controller/cluster/checks/check"
	checkgroup "github.com/sanmoh-hombal/provider-checkly/internal/controller/cluster/checks/checkgroup"
	checkgroupv2 "github.com/sanmoh-hombal/provider-checkly/internal/controller/cluster/checks/checkgroupv2"
	dnsmonitor "github.com/sanmoh-hombal/provider-checkly/internal/controller/cluster/checks/dnsmonitor"
	heartbeat "github.com/sanmoh-hombal/provider-checkly/internal/controller/cluster/checks/heartbeat"
	heartbeatmonitor "github.com/sanmoh-hombal/provider-checkly/internal/controller/cluster/checks/heartbeatmonitor"
	icmpmonitor "github.com/sanmoh-hombal/provider-checkly/internal/controller/cluster/checks/icmpmonitor"
	tcpcheck "github.com/sanmoh-hombal/provider-checkly/internal/controller/cluster/checks/tcpcheck"
	tcpmonitor "github.com/sanmoh-hombal/provider-checkly/internal/controller/cluster/checks/tcpmonitor"
	urlmonitor "github.com/sanmoh-hombal/provider-checkly/internal/controller/cluster/checks/urlmonitor"
	clientcertificate "github.com/sanmoh-hombal/provider-checkly/internal/controller/cluster/infra/clientcertificate"
	environmentvariable "github.com/sanmoh-hombal/provider-checkly/internal/controller/cluster/infra/environmentvariable"
	privatelocation "github.com/sanmoh-hombal/provider-checkly/internal/controller/cluster/infra/privatelocation"
	snippet "github.com/sanmoh-hombal/provider-checkly/internal/controller/cluster/infra/snippet"
	checksuite "github.com/sanmoh-hombal/provider-checkly/internal/controller/cluster/playwright/checksuite"
	codebundle "github.com/sanmoh-hombal/provider-checkly/internal/controller/cluster/playwright/codebundle"
	providerconfig "github.com/sanmoh-hombal/provider-checkly/internal/controller/cluster/providerconfig"
	dashboard "github.com/sanmoh-hombal/provider-checkly/internal/controller/cluster/statuspage/dashboard"
	statuspage "github.com/sanmoh-hombal/provider-checkly/internal/controller/cluster/statuspage/statuspage"
	statuspageservice "github.com/sanmoh-hombal/provider-checkly/internal/controller/cluster/statuspage/statuspageservice"
	checktrigger "github.com/sanmoh-hombal/provider-checkly/internal/controller/cluster/trigger/check"
	group "github.com/sanmoh-hombal/provider-checkly/internal/controller/cluster/trigger/group"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		alertchannel.Setup,
		maintenancewindow.Setup,
		check.Setup,
		checkgroup.Setup,
		checkgroupv2.Setup,
		dnsmonitor.Setup,
		heartbeat.Setup,
		heartbeatmonitor.Setup,
		icmpmonitor.Setup,
		tcpcheck.Setup,
		tcpmonitor.Setup,
		urlmonitor.Setup,
		clientcertificate.Setup,
		environmentvariable.Setup,
		privatelocation.Setup,
		snippet.Setup,
		checksuite.Setup,
		codebundle.Setup,
		providerconfig.Setup,
		dashboard.Setup,
		statuspage.Setup,
		statuspageservice.Setup,
		checktrigger.Setup,
		group.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		alertchannel.SetupGated,
		maintenancewindow.SetupGated,
		check.SetupGated,
		checkgroup.SetupGated,
		checkgroupv2.SetupGated,
		dnsmonitor.SetupGated,
		heartbeat.SetupGated,
		heartbeatmonitor.SetupGated,
		icmpmonitor.SetupGated,
		tcpcheck.SetupGated,
		tcpmonitor.SetupGated,
		urlmonitor.SetupGated,
		clientcertificate.SetupGated,
		environmentvariable.SetupGated,
		privatelocation.SetupGated,
		snippet.SetupGated,
		checksuite.SetupGated,
		codebundle.SetupGated,
		providerconfig.SetupGated,
		dashboard.SetupGated,
		statuspage.SetupGated,
		statuspageservice.SetupGated,
		checktrigger.SetupGated,
		group.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
