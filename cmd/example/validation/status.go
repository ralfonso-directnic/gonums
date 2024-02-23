package validation

/*validation was statically generated 2024-02-23T13:54:17Z by github.com/ralfonso-directnic/gonums, do not edit*/

type Status int

const (
	StatusUnknown Status = iota
	StatusFailed
	StatusPassed
	StatusSkipped
	StatusScheduled
	StatusRunning
)

var (
	strStatusMap = map[Status]string{
		StatusFailed:    "Failed",
		StatusPassed:    "Passed",
		StatusSkipped:   "Skipped",
		StatusScheduled: "Scheduled",
		StatusRunning:   "Running",
	}

	typeStatusMap = map[string]Status{
		"Failed":    StatusFailed,
		"Passed":    StatusPassed,
		"Skipped":   StatusSkipped,
		"Scheduled": StatusScheduled,
		"Running":   StatusRunning,
	}
)

func (t Status) String() string {
	return strStatusMap[t]
}

func (t Status) Int() int {
	return int(t)
}

func (t Status) IsValid() bool {
	_, ok := strStatusMap[t]
	return ok
}

func ParseStatus(str string) Status {

	if val, ok := typeStatusMap[str]; ok {

		return val

	}

	return StatusUnknown

}

func StatusStrings() []string {

	return []string{
		"",
		"Failed",
		"Passed",
		"Skipped",
		"Scheduled",
		"Running",
	}

}

func StatusCollection() []Status {

	return []Status{
		StatusUnknown,
		StatusFailed,
		StatusPassed,
		StatusSkipped,
		StatusScheduled,
		StatusRunning,
	}

}
