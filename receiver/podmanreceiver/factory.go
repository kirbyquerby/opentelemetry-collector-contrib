// Copyright 2020 OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package podmanreceiver

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver/receiverhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/scraperhelper"
)

const (
	typeStr           = "podman_stats"
	defaultAPIVersion = "3.3.1"
)

func NewFactory() component.ReceiverFactory {
	return receiverhelper.NewFactory(
		typeStr,
		createDefaultReceiverConfig,
		receiverhelper.WithMetrics(createMetricsReceiver))
}

func createDefaultConfig() *Config {
	return &Config{
		ScraperControllerSettings: scraperhelper.ScraperControllerSettings{
			ReceiverSettings:   config.NewReceiverSettings(config.NewComponentID(typeStr)),
			CollectionInterval: 10 * time.Second,
		},
		Endpoint:   "unix:///run/podman/podman.sock",
		APIVersion: defaultAPIVersion,
	}
}

func createDefaultReceiverConfig() config.Receiver {
	return createDefaultConfig()
}

func createMetricsReceiver(
	ctx context.Context,
	params component.ReceiverCreateSettings,
	config config.Receiver,
	consumer consumer.Metrics,
) (component.MetricsReceiver, error) {
	podmanConfig := config.(*Config)
	dsr, err := newReceiver(ctx, params.Logger, podmanConfig, consumer, nil)
	if err != nil {
		return nil, err
	}

	return dsr, nil
}
