# tiktok-mall-chart/templates/_helpers.tpl
{{/*
Generate service fullname
*/}}
{{- define "tiktok-mall.fullname" -}}
{{- printf "%s-%s" .Release.Name .Chart.Name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Generate service name for each microservice
*/}}
{{- define "tiktok-mall.serviceName" -}}
{{- printf "%s-%s" .Release.Name .Values.serviceName | trunc 63 | trimSuffix "-" -}}
{{- end -}}
