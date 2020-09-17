package proto

const (
	EventTypeNonMotor     = "non_motor"
	EventTypeVehicle      = "vehicle"
	EventTypeConstruction = "construction"

	EventTypeNonMotorizedOffset = 2200
	EventTypeConstructionOffset = 2300

	ClassWaimaiOffset = 21000
)

func EventCodeToType(eventType int) string {
	if IsNonMotorEvent(eventType) {
		return EventTypeNonMotor
	}
	if IsVehicleEvent(eventType) {
		return EventTypeVehicle
	}
	if IsConstructionEvent(eventType) {
		return EventTypeConstruction
	}
	return ""
}

func IsNonMotorEvent(eventType int) bool {
	eventNonMotor := eventType - EventTypeNonMotorizedOffset
	return eventNonMotor > 0 && eventNonMotor < 100
}

func IsVehicleEvent(eventType int) bool {
	eventNonMotor := eventType - EventTypeNonMotorizedOffset
	return eventNonMotor > -100 && eventNonMotor < 0
}

func IsConstructionEvent(eventType int) bool {
	eventNonMotor := eventType - EventTypeConstructionOffset
	if eventNonMotor > 0 && eventNonMotor < 100 {
		return true
	}
	eventNonMotor = eventType + EventTypeConstructionOffset
	return eventNonMotor > -100 && eventNonMotor < 0
}
