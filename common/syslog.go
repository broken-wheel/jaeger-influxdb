package common

const (
	// Tags
	AppNameKey  = "appname"
	FacilityKey = "facility"
	HostKey     = "host"
	HostNameKey = "hostname"
	SeverityKey = "severity"

	// Fields
	FacilityCodeKey = "facility_code" // integer
	MessageKey      = "message"       // string
	ProcIDKey       = "procid"        // string
	SeverityCodeKey = "severity_code" // integer
	TimestampKey    = "timestamp"     // integer
	VersionKey      = "version"       // integer
)

var SeverityMap = map[string]int{
	"debug":   7,
	"info":    6,
	"notice":  5,
	"warning": 4,
	"err":     3,
	"crit":    2,
	"alert":   1,
	"emerg":   0,
}
