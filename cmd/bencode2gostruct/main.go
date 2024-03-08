package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"
	"text/template"

	"github.com/jackpal/bencode-go"
	"golang.org/x/tools/imports"
)

const structTemplate = `
package {{ .Package }}

type {{ .Model.Name }} struct {
{{- range .Model.Fields }}
    {{ .Key }}		{{ .Type }} ` + "`bencode:\"{{ .BencodeKey }}\"`" + `
{{- end }}
}
`

type GoFile struct {
	Package string
	Model   *Model
}

type Model struct {
	Name   string
	Fields []*Field
}

type Field struct {
	Key        string
	BencodeKey string
	Type       string
}

func main() {
	inFile := flag.String("in", "", "Bencoded file to generate model from")
	outFile := flag.String("out", "", "Go model file to output")
	packageName := flag.String("package", "", "The `package` name for the Go file")
	structName := flag.String("model-name", "", "Model name for the struct")
	flag.Parse()

	if *inFile == "" || *outFile == "" || *packageName == "" || *structName == "" {
		flag.Usage()
		os.Exit(1)
	}

	fields, err := parseBencodeFile(*inFile)
	if err != nil {
		panic(err)
	}

	goFile := &GoFile{
		Package: *packageName,
		Model: &Model{
			Name:   *structName,
			Fields: fields,
		},
	}

	var rendered []byte
	if rendered, err = renderTemplate(goFile); err != nil {
		panic(err)
	}

	if err = outputGoFile(*outFile, rendered); err != nil {
		panic(err)
	}

}

func goTypeForBencodeType(t reflect.Type) string {
	switch t.Kind() {
	case reflect.Int, reflect.Int64:
		return "int64"
	case reflect.String:
		return "string"
	case reflect.Slice:
		return "[]interface{}"
	case reflect.Map:
		return "map[string]interface{}"
	default:
		return "interface{}"
	}
}

func outputGoFile(filename string, data []byte) error {
	var fmtOut []byte
	var err error
	fmtOut, err = imports.Process(filename, data, nil)

	fmt.Println(string(fmtOut))

	var oFile *os.File
	oFile, err = os.Create(filename)
	if err != nil {
		return err
	}
	defer oFile.Close()
	_, err = oFile.Write(fmtOut)
	return err
}

func parseBencodeFile(filename string) ([]*Field, error) {
	var data interface{}
	file, err := os.Open(filename)
	if err != nil {
		return []*Field{}, err
	}
	defer file.Close()

	data, err = bencode.Decode(file)
	if err != nil {
		return []*Field{}, err
	}

	decodedMap, ok := data.(map[string]interface{})
	if !ok {
		return []*Field{}, fmt.Errorf("expected top-level dictionary in Bencode data")
	}

	keys := make([]string, 0, len(decodedMap))
	for key := range decodedMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	fields := make([]*Field, 0, len(decodedMap))
	for _, key := range keys {
		value := decodedMap[key]
		fieldType := reflect.TypeOf(value)
		goType := goTypeForBencodeType(fieldType)
		fields = append(fields, &Field{Key: toPascalCase(key), BencodeKey: key, Type: goType})
	}
	return fields, nil
}

func renderTemplate(data any) ([]byte, error) {
	tmpl, err := template.New("struct").Parse(structTemplate)
	if err != nil {
		return []byte{}, err
	}
	var out bytes.Buffer

	err = tmpl.Execute(&out, data)
	return out.Bytes(), err
}

func toPascalCase(input string) string {
	words := strings.FieldsFunc(input, func(r rune) bool {
		return r == '_' || r == '-'
	})

	var pascalCase string
	for _, word := range words {
		pascalCase += strings.ToUpper(string(word[0])) + word[1:]
	}

	return pascalCase
}
