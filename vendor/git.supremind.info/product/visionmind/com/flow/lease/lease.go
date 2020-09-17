package lease

const (
	STATUS_STARTING = "starting"
	STATUS_RUNNING  = "running"
	STATUS_WAITING  = "waiting"
	STATUS_STOPPING = "stopping"
	STATUS_STOPPED  = "stopped"
	STATUS_RETRYING = "retrying"
	STATUS_FAILED   = "failed"

	TryStatusKey = "tryStatus"
)
