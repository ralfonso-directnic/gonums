# Gonums

gonums is a package to help you generate go type safe enums

# Usage

See example main.go, this can easily be turned into a command to generate before building

### Example
Defining the list of enums their type and package in a JSON list is the configuration format.  This allows generating many enums in one shot.
input.json:
```json
{
  "enums": [
    {
      "package": "validation",
      "type": "status",
      "values": [
        "Failed",
        "Passed",
        "Skipped",
        "Scheduled",
        "Running"
      ]
    }
  ]
}
```

```golang
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

```
## Features


#### Extendable
The enums can have additional functionality added by just adding another file in the same package with the extra functionality defined there.  Having the extra functionality in another file will allow the generation and regeneration of the enums to not affect this extra functionality.

#### Safety

The above `Validation Status` can be found in the examples directory along with another file extending the behaviour of the `Status` enum and the `config.json` that was used to generate.

#### Credit

This package was inspired by https://github.com/zarldev/goenums, it just had features I didn't want and didn't support multiple enums in the same package