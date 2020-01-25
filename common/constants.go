package common

// FlyingDeviceType describes a flying device type
type FlyingDeviceType uint

const (
	// Paraglider flyingDevice
	Paraglider FlyingDeviceType = 1
	// SpeedRider flyingDevice
	SpeedRider FlyingDeviceType = 2
	// HangGlider flyingDevice
	HangGlider FlyingDeviceType = 3
	// HotAirBaloon flyingDevice
	HotAirBaloon FlyingDeviceType = 4
)

// SessionParameters is used to define consts
type SessionParameters string

const (
	SessionParamUserID SessionParameters = "userid"
	SessionParamRoles  SessionParameters = "roles"
)
