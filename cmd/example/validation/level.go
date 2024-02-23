package validation

/*validation was statically generated 2024-02-23T13:54:17Z by github.com/ralfonso-directnic/gonums, do not edit*/

type Level int

const (
	LevelUnknown Level = iota
	LevelStarter
	LevelPro
)

var (
	strLevelMap = map[Level]string{
		LevelStarter: "Starter",
		LevelPro:     "Pro",
	}

	typeLevelMap = map[string]Level{
		"Starter": LevelStarter,
		"Pro":     LevelPro,
	}
)

func (t Level) String() string {
	return strLevelMap[t]
}

func (t Level) Int() int {
	return int(t)
}

func (t Level) IsValid() bool {
	_, ok := strLevelMap[t]
	return ok
}

func ParseLevel(str string) Level {

	if val, ok := typeLevelMap[str]; ok {

		return val

	}

	return LevelUnknown

}

func LevelStrings() []string {

	return []string{
		"",
		"Starter",
		"Pro",
	}

}

func LevelCollection() []Level {

	return []Level{
		LevelUnknown,
		LevelStarter,
		LevelPro,
	}

}
