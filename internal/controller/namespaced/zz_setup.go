// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	alertchannel "github.com/sanmoh-hombal/provider-checkly/internal/controller/namespaced/alerts/alertchannel"
	maintenancewindow "github.com/sanmoh-hombal/provider-checkly/internal/controller/namespaced/alerts/maintenancewindow"
	check "github.com/sanmoh-hombal/provider-checkly/internal/controller/namespaced/checks/check"
	checkgroup "github.com/sanmoh-hombal/provider-checkly/internal/controller/namespaced/checks/checkgroup"
	checkgroupv2 "github.com/sanmoh-hombal/provider-checkly/internal/controller/namespaced/checks/checkgroupv2"
	dnsmonitor "github.com/sanmoh-hombal/provider-checkly/internal/controller/namespaced/checks/dnsmonitor"
	heartbeat "github.com/sanmoh-hombal/provider-checkly/internal/controller/namespaced/checks/heartbeat"
	heartbeatmonitor "github.com/sanmoh-hombal/provider-checkly/internal/controller/namespaced/checks/heartbeatmonitor"
	icmpmonitor "github.com/sanmoh-hombal/provider-checkly/internal/controller/namespaced/checks/icmpmonitor"
	tcpcheck "github.com/sanmoh-hombal/provider-checkly/internal/controller/namespaced/checks/tcpcheck"
	tcpmonitor "github.com/sanmoh-hombal/provider-checkly/internal/controller/namespaced/checks/tcpmonitor"
	urlmonitor "github.com/sanmoh-hombal/provider-checkly/internal/controller/namespaced/checks/urlmonitor"
	certificate "github.com/sanmoh-hombal/provider-checkly/internal/controller/namespaced/client/certificate"
	environmentvariable "github.com/sanmoh-hombal/provider-checkly/internal/controller/namespaced/infra/environmentvariable"
	snippet "github.com/sanmoh-hombal/provider-checkly/internal/controller/namespaced/infra/snippet"
	checksuite "github.com/sanmoh-hombal/provider-checkly/internal/controller/namespaced/playwright/checksuite"
	codebundle "github.com/sanmoh-hombal/provider-checkly/internal/controller/namespaced/playwright/codebundle"
	location "github.com/sanmoh-hombal/provider-checkly/internal/controller/namespaced/private/location"
	providerconfig "github.com/sanmoh-hombal/provider-checkly/internal/controller/namespaced/providerconfig"
	page "github.com/sanmoh-hombal/provider-checkly/internal/controller/namespaced/status/page"
	pageservice "github.com/sanmoh-hombal/provider-checkly/internal/controller/namespaced/status/pageservice"
	dashboard "github.com/sanmoh-hombal/provider-checkly/internal/controller/namespaced/statuspage/dashboard"
	checktrigger "github.com/sanmoh-hombal/provider-checkly/internal/controller/namespaced/trigger/check"
	group "github.com/sanmoh-hombal/provider-checkly/internal/controller/namespaced/trigger/group"
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
		certificate.Setup,
		environmentvariable.Setup,
		snippet.Setup,
		checksuite.Setup,
		codebundle.Setup,
		location.Setup,
		providerconfig.Setup,
		page.Setup,
		pageservice.Setup,
		dashboard.Setup,
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
		certificate.SetupGated,
		environmentvariable.SetupGated,
		snippet.SetupGated,
		checksuite.SetupGated,
		codebundle.SetupGated,
		location.SetupGated,
		providerconfig.SetupGated,
		page.SetupGated,
		pageservice.SetupGated,
		dashboard.SetupGated,
		checktrigger.SetupGated,
		group.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
