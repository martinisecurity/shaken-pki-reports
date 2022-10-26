package lint

// LintStatus is an enum returned by lints inside of a LintResult.
type LintStatus int

// Known LintStatus values
const (
	// Unused / unset LintStatus
	Reserved LintStatus = 0
	// Not Applicable
	NA     LintStatus = 1
	Pass   LintStatus = 3
	Notice LintStatus = 4
	Warn   LintStatus = 5
	Error  LintStatus = 6
	Fatal  LintStatus = 7
)

// String returns the canonical representation of a LintStatus as a string.
func (e LintStatus) String() string {
	switch e {
	case Reserved:
		return "reserved"
	case NA:
		return "not applicable"
	case Pass:
		return "pass"
	case Notice:
		return "info"
	case Warn:
		return "warn"
	case Error:
		return "error"
	case Fatal:
		return "fatal"
	default:
		return ""
	}
}
