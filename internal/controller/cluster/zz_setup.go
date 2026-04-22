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
	playwrightcodebundle "github.com/sanmoh-hombal/provider-checkly/internal/controller/cluster/checks/playwrightcodebundle"
	tcpcheck "github.com/sanmoh-hombal/provider-checkly/internal/controller/cluster/checks/tcpcheck"
	tcpmonitor "github.com/sanmoh-hombal/provider-checkly/internal/controller/cluster/checks/tcpmonitor"
	urlmonitor "github.com/sanmoh-hombal/provider-checkly/internal/controller/cluster/checks/urlmonitor"
	clientcertificate "github.com/sanmoh-hombal/provider-checkly/internal/controller/cluster/infra/clientcertificate"
	environmentvariable "github.com/sanmoh-hombal/provider-checkly/internal/controller/cluster/infra/environmentvariable"
	privatelocation "github.com/sanmoh-hombal/provider-checkly/internal/controller/cluster/infra/privatelocation"
	snippet "github.com/sanmoh-hombal/provider-checkly/internal/controller/cluster/infra/snippet"
	checksuite "github.com/sanmoh-hombal/provider-checkly/internal/controller/cluster/playwright/checksuite"
	providerconfig "github.com/sanmoh-hombal/provider-checkly/internal/controller/cluster/providerconfig"
	dashboard "github.com/sanmoh-hombal/provider-checkly/internal/controller/cluster/statuspage/dashboard"
	statuspage "github.com/sanmoh-hombal/provider-checkly/internal/controller/cluster/statuspage/statuspage"
	statuspageservice "github.com/sanmoh-hombal/provider-checkly/internal/controller/cluster/statuspage/statuspageservice"
	triggercheck "github.com/sanmoh-hombal/provider-checkly/internal/controller/cluster/triggers/triggercheck"
	triggergroup "github.com/sanmoh-hombal/provider-checkly/internal/controller/cluster/triggers/triggergroup"
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
		playwrightcodebundle.Setup,
		tcpcheck.Setup,
		tcpmonitor.Setup,
		urlmonitor.Setup,
		clientcertificate.Setup,
		environmentvariable.Setup,
		privatelocation.Setup,
		snippet.Setup,
		checksuite.Setup,
		providerconfig.Setup,
		dashboard.Setup,
		statuspage.Setup,
		statuspageservice.Setup,
		triggercheck.Setup,
		triggergroup.Setup,
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
		playwrightcodebundle.SetupGated,
		tcpcheck.SetupGated,
		tcpmonitor.SetupGated,
		urlmonitor.SetupGated,
		clientcertificate.SetupGated,
		environmentvariable.SetupGated,
		privatelocation.SetupGated,
		snippet.SetupGated,
		checksuite.SetupGated,
		providerconfig.SetupGated,
		dashboard.SetupGated,
		statuspage.SetupGated,
		statuspageservice.SetupGated,
		triggercheck.SetupGated,
		triggergroup.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
