// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/damoun/uptrends/v4/client/account"
	"github.com/damoun/uptrends/v4/client/alert"
	"github.com/damoun/uptrends/v4/client/alert_definition"
	"github.com/damoun/uptrends/v4/client/checkpoint"
	"github.com/damoun/uptrends/v4/client/dashboard"
	"github.com/damoun/uptrends/v4/client/integration"
	"github.com/damoun/uptrends/v4/client/miscellaneous"
	"github.com/damoun/uptrends/v4/client/monitor"
	"github.com/damoun/uptrends/v4/client/monitor_check"
	"github.com/damoun/uptrends/v4/client/monitor_group"
	"github.com/damoun/uptrends/v4/client/operator"
	"github.com/damoun/uptrends/v4/client/operator_group"
	"github.com/damoun/uptrends/v4/client/public_status_page"
	"github.com/damoun/uptrends/v4/client/register"
	"github.com/damoun/uptrends/v4/client/scheduled_report"
	"github.com/damoun/uptrends/v4/client/sla"
	"github.com/damoun/uptrends/v4/client/statistics"
	"github.com/damoun/uptrends/v4/client/status"
	"github.com/damoun/uptrends/v4/client/vault"
)

// Default uptrends API v4 HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "api.uptrends.com"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/v4"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new uptrends API v4 HTTP client.
func NewHTTPClient(formats strfmt.Registry) *UptrendsAPIV4 {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new uptrends API v4 HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *UptrendsAPIV4 {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new uptrends API v4 client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *UptrendsAPIV4 {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(UptrendsAPIV4)
	cli.Transport = transport
	cli.Account = account.New(transport, formats)
	cli.Alert = alert.New(transport, formats)
	cli.AlertDefinition = alert_definition.New(transport, formats)
	cli.Checkpoint = checkpoint.New(transport, formats)
	cli.Dashboard = dashboard.New(transport, formats)
	cli.Integration = integration.New(transport, formats)
	cli.Miscellaneous = miscellaneous.New(transport, formats)
	cli.Monitor = monitor.New(transport, formats)
	cli.MonitorCheck = monitor_check.New(transport, formats)
	cli.MonitorGroup = monitor_group.New(transport, formats)
	cli.Operator = operator.New(transport, formats)
	cli.OperatorGroup = operator_group.New(transport, formats)
	cli.PublicStatusPage = public_status_page.New(transport, formats)
	cli.Register = register.New(transport, formats)
	cli.ScheduledReport = scheduled_report.New(transport, formats)
	cli.SLA = sla.New(transport, formats)
	cli.Statistics = statistics.New(transport, formats)
	cli.Status = status.New(transport, formats)
	cli.Vault = vault.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// UptrendsAPIV4 is a client for uptrends API v4
type UptrendsAPIV4 struct {
	Account account.ClientService

	Alert alert.ClientService

	AlertDefinition alert_definition.ClientService

	Checkpoint checkpoint.ClientService

	Dashboard dashboard.ClientService

	Integration integration.ClientService

	Miscellaneous miscellaneous.ClientService

	Monitor monitor.ClientService

	MonitorCheck monitor_check.ClientService

	MonitorGroup monitor_group.ClientService

	Operator operator.ClientService

	OperatorGroup operator_group.ClientService

	PublicStatusPage public_status_page.ClientService

	Register register.ClientService

	ScheduledReport scheduled_report.ClientService

	SLA sla.ClientService

	Statistics statistics.ClientService

	Status status.ClientService

	Vault vault.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *UptrendsAPIV4) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Account.SetTransport(transport)
	c.Alert.SetTransport(transport)
	c.AlertDefinition.SetTransport(transport)
	c.Checkpoint.SetTransport(transport)
	c.Dashboard.SetTransport(transport)
	c.Integration.SetTransport(transport)
	c.Miscellaneous.SetTransport(transport)
	c.Monitor.SetTransport(transport)
	c.MonitorCheck.SetTransport(transport)
	c.MonitorGroup.SetTransport(transport)
	c.Operator.SetTransport(transport)
	c.OperatorGroup.SetTransport(transport)
	c.PublicStatusPage.SetTransport(transport)
	c.Register.SetTransport(transport)
	c.ScheduledReport.SetTransport(transport)
	c.SLA.SetTransport(transport)
	c.Statistics.SetTransport(transport)
	c.Status.SetTransport(transport)
	c.Vault.SetTransport(transport)
}
