// Binary mktyping generates TypingActionBuilder.
package main

import (
	"flag"
	"reflect"

	"github.com/go-faster/errors"

	"github.com/KoNekoD/td/tdp"
	"github.com/KoNekoD/td/telegram/message/internal/mkrun"
	"github.com/KoNekoD/td/tg"
)

// Field represents type field.
type Field struct {
	// Name is Go name of field.
	Name string
	// Type is Go type of field.
	Type string
}

// Type represents generated type.
type Type struct {
	// Name is Go name of type.
	Name string
	// Fields is slice of type fields.
	Fields []Field
	// SchemaType is related schema type.
	SchemaType tdp.Type
}

const rawTemplate = `// Code generated by mktyping, DO NOT EDIT.
package {{ $.PackageName }}

import (
	"context"

	"github.com/KoNekoD/td/tg"
)

var (
	_ = tg.Invoker(nil)
	_ = context.Context(nil)
)

{{- /*gotype: github.com/KoNekoD/td/telegram/message/internal/mkrun.Config*/ -}}
{{- range $typ := $.Data }}
{{ $helperName := trimSuffix (trimPrefix $typ.Name "SendMessage") "Action" -}}
// {{ $helperName }} sends {{ $typ.Name }}.
func (b *TypingActionBuilder) {{ $helperName }}(ctx context.Context,
{{- range $f := $typ.Fields }}{{ lowerFirst $f.Name }} {{ $f.Type }},{{ end }}) error {
	return b.send(ctx, &tg.{{ $typ.Name }}{
		{{- range $f := $typ.Fields }}
		{{ $f.Name }}: {{ lowerFirst $f.Name }},
		{{- end }}
	})
}
{{- end }}
`

var (
	constructors = tg.ClassConstructorsMap()
	create       = tg.TypesConstructorMap()
)

type generator struct{}

func (g generator) Name() string {
	return "mktyping"
}

func (g generator) Flags(set *flag.FlagSet) {}

func (g generator) Template() string {
	return rawTemplate
}

func (g generator) Data() (interface{}, error) {
	var types []Type
	for _, typeID := range constructors[tg.SendMessageActionClassName] {
		v, ok := create[typeID]().(tdp.Object)
		if !ok {
			return nil, errors.Errorf("bad type %#x", typeID)
		}
		schemaType := v.TypeInfo()
		tv := reflect.TypeOf(v).Elem()

		var fields []Field
		for _, field := range schemaType.Fields {
			rf, ok := tv.FieldByName(field.Name)
			if !ok {
				return nil, errors.Errorf(
					"field of %q type %q not found",
					field.Name, schemaType.Name,
				)
			}
			fields = append(fields, Field{
				Name: field.Name,
				Type: rf.Type.String(),
			})
		}
		types = append(types, Type{
			Name:       tv.Name(),
			Fields:     fields,
			SchemaType: v.TypeInfo(),
		})
	}

	return types, nil
}

func main() {
	mkrun.Main(generator{})
}
