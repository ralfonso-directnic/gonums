package gonums

import (
	"embed"
	_ "embed"
	"fmt"
	"regexp"
	"os"
	"time"
	"os/exec"
	"path"
	"text/template"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

//go:embed template/enum.tmpl
var fs embed.FS

type EnumTemplateData struct {
	Package        string
	TypeName       string
	TypeNameLower  string
	TypeNameTitle  string
	TypeNamePlural string
	Enums          []Enum
}

type Enum struct {
	VariableStr      string
	VariableLabel    string
	VariableStrUpper string
	VariableStrLower string
	VariableStrTitle string
	TypeName         string
	TypeNameLower    string
}

type Generator struct {
	outputPath string
	data       []EnumTemplateData
}

func New(cfg Config) *Generator {
	return &Generator{
		data:       parseConfig(cfg),
		outputPath: cfg.OutputPath,
	}
}

func (g *Generator) Generate() error {
	for _, d := range g.data {
		err := generateEnum(d, g.outputPath)
		if err != nil {
			return err
		}
	}
	return nil
}

func generateEnum(etd EnumTemplateData, outPath string) error {
	f, fp, err := mkDirAndGoFiles(outPath, etd.Package, etd.TypeName)
	if err != nil {
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println("Error closing file", err)
		}
	}()

	funcMap := template.FuncMap{
		// The name "title" is what the function will be called in the template text.
		"now": func() string{

			return time.Now().Format(time.RFC3339)
		},
	}

	cmdA := exec.Command("go", "clean")
	err = cmdA.Run()
	if err != nil {
		return err
	}

	t := template.Must(template.New("enum.tmpl").Funcs(funcMap).ParseFS(fs, "template/enum.tmpl"))
	err = t.Execute(f, etd)
	if err != nil {
		return err
	}

	cmd := exec.Command("gofmt", "-w", fp)
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func parseConfig(cfg Config) []EnumTemplateData {
	data := make([]EnumTemplateData, len(cfg.Configs))
	for i, c := range cfg.Configs {
		typeName, packageName, enums := configToVars(c)
		c := cases.Title(language.English, cases.NoLower)
		typeNameTitle := c.String(typeName)
		typeNamePlural := typeNameTitle + "s"
		if typeNameTitle[len(typeNameTitle)-1] == 's' {
			typeNamePlural = typeNameTitle + "es"
		}
		etd := EnumTemplateData{
			Package:        packageName,
			TypeName:       typeNameTitle,
			TypeNameLower:  strings.ToLower(typeName),
			TypeNameTitle:  typeNameTitle,
			TypeNamePlural: typeNamePlural,
			Enums:          enums,
		}
		data[i] = etd
	}
	return data
}

func configToVars(cfg EnumConfig) (string, string, []Enum) {
	typ := cfg.Type
	pkg := cfg.Package
	enumStrs := cfg.Enums
	enums := make([]Enum, len(enumStrs))
	c := cases.Title(language.English, cases.NoLower)
	typeTitle := c.String(typ)
	re := regexp.MustCompile("[^A-Za-z0-9]")
	for i, enumStr := range enumStrs {
		enumLabel := enumStr
		if strings.Contains(enumStr,":") {

			parts := strings.Split(enumStr,":")

			if len(parts)>1{

				enumStr   = parts[0]
				enumLabel = strings.Trim(parts[1]," ")

			}

		}

		op:= re.ReplaceAllString(enumStr, "")

		enums[i] = Enum{
			VariableStr:      enumStr,
			VariableLabel:    enumLabel,
			VariableStrUpper: strings.ToUpper(op),
			VariableStrLower: strings.ToLower(op),
			VariableStrTitle: c.String(op),
			TypeName:         typeTitle,
			TypeNameLower:    strings.ToLower(typ),
		}
	}
	return typ, pkg, enums
}

func mkDirAndGoFiles(outpath, pkg, typ string) (*os.File, string, error) {
	if _, err := os.Stat(outpath); os.IsNotExist(err) {
		err = os.Mkdir(outpath, os.ModePerm)
		if err != nil {
			return nil, "", err
		}
	}
	dir := path.Join(outpath, pkg)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.Mkdir(dir, os.ModePerm)
		if err != nil {
			return nil, "", err
		}
	}
	fName := fmt.Sprintf("%s.go", strings.ToLower(typ))
	fPath := path.Join(dir, fName)
	f, err := os.OpenFile(fPath, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return nil, "", err
	}
	fullPath := path.Join(outpath, pkg, fName)
	return f, fullPath, nil
}
