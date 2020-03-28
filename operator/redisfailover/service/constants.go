package service

const (
	logNameField      = "redisfailover"
	logNamespaceField = "namespace"
)

// variables refering to the redis exporter port
const (
	exporterPort                  = 9121
	sentinelExporterPort          = 9355
	exporterPortName              = "metrics"
	exporterContainerName         = "metrics"
	sentinelExporterContainerName = "metrics"
	exporterDefaultRequestCPU     = "25m"
	exporterDefaultLimitCPU       = "50m"
	exporterDefaultRequestMemory  = "50Mi"
	exporterDefaultLimitMemory    = "100Mi"
)

const (
	description            = "Manage a Redis Failover deployment"
	bootstrapName          = "redis-bootstrap"
	sentinelName           = "redis-sentinel"
	sentinelRoleName       = "sentinel"
	sentinelConfigFileName = "sentinel.conf"
	redisConfigFileName    = "redis.conf"
	redisName              = "redis"
	redisShutdownName      = "redis-shutdown"
	redisReadinessName     = "redis-readiness"
	redisRoleName          = "redis"
	redisGroupName         = "mymaster"
	appLabel               = "redis-failover"
	hostnameTopologyKey    = "kubernetes.io/hostname"
)
