package aws

import "go.uber.org/fx"

var Module = fx.Provide(
	New,
)