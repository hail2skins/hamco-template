package helpers

import (
	"github.com/newrelic/go-agent/v3/newrelic"
)

func NewRelicApp() (*newrelic.Application, error) {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("Hamco Internet Solutions"),
		newrelic.ConfigLicense("7343fd7813ee61fd40620886e112b50776e2NRAL"),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	return app, err
}
