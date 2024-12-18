{{- /* BEGIN DEFINE */ -}}

{{- define "Enum"}}

enum {{.Name}} {
{{- range .EnumFields}}
	{{.Name}} = {{.Value}};
{{- end}}
}
{{- end}}

{{- define "Message"}}

message {{.Name}} {
{{- range .Fields}}
{{- if .Repeated}}
	repeated {{toProtoType .Type}} {{.Name}} = {{.Number}};
{{- else if len .KeyType}}
	map<{{toProtoType .KeyType}}, {{toProtoType .Type}}> {{.Name}} = {{.Number}};
{{- else}}
	{{toProtoType .Type}} {{.Name}} = {{.Number}};
{{- end}}
{{- end}}
}
{{- end}}

{{- define "Entity"}}
{{- template "Message" .}}
{{- end}}

{{- define "Component"}}
{{- template "Message" .}}
{{- end}}

{{- /* END DEFINE */ -}}

// Code generated by kds. DO NOT EDIT.

syntax = "proto3";

package {{.Package}};

{{- if len .Imports}}
{{/* empty line */}}
{{- end}}
{{- range .Imports}}
import "{{.}}.proto";
{{- end}}

{{- if len .ProtoImports}}
{{/* empty line */}}
{{- range .ProtoImports}}
import "{{.}}";
{{- end}}
{{- end}}

option go_package="{{.ProtoGoPackage}}";

{{- range .Defs}}
{{- if findEnum .Name}}
{{- template "Enum" .}}
{{- else if findEntity .Name}}
{{- template "Entity" .}}
{{- else if findComponent .Name}}
{{- template "Component" .}}
{{- end}}
{{- end}}