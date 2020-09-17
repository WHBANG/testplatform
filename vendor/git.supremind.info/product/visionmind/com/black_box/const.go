package black_box

const (
	ProductJiaotong = "jiaotong"
	ProductVMR      = "vmr"
	ProductCA       = "ca"

	ServiceConsole    = "console"
	ServiceFlow       = "flow"
	ServiceFileserver = "fileserver"
	ServiceDevice     = "device"
	ServiceStream     = "stream"

	EventLevelInfo  = "INFO"
	EventLevelError = "ERROR"

	EventTypeConsoleGetMsg        = "GetMessage"
	EventTypeConsoleFileter       = "Filtered"
	EventTypeConsoleEnterDatabase = "EnterDatabase"
	EventTypeConsoleMark          = "Mark"
	EventTypeConsoleDownloadVideo = "DownloadVideo"

	EventTypeTaskCreate = "TaskCreate"
	EventTypeTaskUpdate = "TaskUpdate"
	EventTypeTaskDelete = "TaskDelete"
	EventTypeTaskStart  = "TaskStart"
	EventTypeTaskStop   = "TaskStop"
	EventTypeTaskStatus = "TaskStatus"

	EventTypeDevicePlayURL     = "DevicePlayURL"
	EventTypeDeviceDownloadURL = "DeviceDownloadURL"
	EventTypeDeviceSnap        = "DeviceSnap"

	EventTypeFileserverClean = "FileserverClean"

	WebhookNotificationStatusFiring   = "firing"
	WebhookNotificationStatusResolved = "resolved"
)
