package agent

// TODO - add nats event topic and benthos ingestor which is dummy and periodically produces a few different events

// One instance per projection
// Subscribes to nats event stream for specified events
// Handles retries and other failures
// Handles offset tracking
// Distributes events to language specific projectors via unified rpc api
// Exposes rpc api used by dashboard (through aperture core which hosts both ui and bff) used for both stats and lifecycle control eg. replay, restart, etc...
// We could build this first and stub / mock language specific projectors

type Config struct {
	ProjectionName     string
	EventSubscriptions []string
}
