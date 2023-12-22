{{ extends "partial_base.go.tpl" }}

{{ define "content" }}
    {{ template "component/projections/projection_item.go.tpl" .Data.Projection }}
{{ end }}