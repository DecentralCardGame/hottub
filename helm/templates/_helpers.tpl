{{/* vim: set filetype=mustache: */}}
{{/*
Expand the name of the chart.
*/}}
{{- define "hottub.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "hottub.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "hottub.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "hottub.labels" -}}
helm.sh/chart: {{ include "hottub.chart" . }}
{{ include "hottub.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "hottub.selectorLabels" -}}
app.kubernetes.io/name: {{ include "hottub.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "hottub.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "hottub.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}

{{/*
Create a default fully qualified postgresql name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
*/}}
{{- define "hottub.postgresql.fullname" -}}
{{- printf "%s-%s" .Release.Name "postgresql" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a default fully qualified name for the secret that contains the database password.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
*/}}
{{- define "hottub.databaseSecret.fullname" -}}
{{- if .Values.postgresql.enabled -}}
{{- include "hottub.postgresql.fullname" . }}
{{- else -}}
{{- include "hottub.fullname" . }}
{{- end -}}
{{- end -}}

{{/*
Set the proper name for the secretKeyRef key that contains the database password.
*/}}
{{- define "hottub.databaseSecret.key" -}}
{{- if .Values.postgresql.enabled -}}
postgres-password
{{- else -}}
db-password
{{- end -}}
{{- end -}}

{{/*
Set the proper database host. If postgresql is installed as part of this chart, use the default service name,
else use user-provided host
*/}}
{{- define "strapi.database.host" }}
{{- if .Values.postgresql.enabled -}}
{{- include "hottub.postgresql.fullname" . }}
{{- else -}}
{{- .Values.database.host | quote }}
{{- end -}}
{{- end -}}

{{/*
Set the proper database port. If postgresql is installed as part of this chart, use the default postgresql port,
else use user-provided port
*/}}
{{- define "hottub.database.port" }}
{{- if .Values.postgresql.enabled -}}
{{- default "5432" ( .Values.postgresql.service.port | quote ) }}
{{- else -}}
{{- .Values.database.port | quote }}
{{- end -}}
{{- end -}}
