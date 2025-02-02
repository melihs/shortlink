# kyverno

![Version: 0.2.1](https://img.shields.io/badge/Version-0.2.1-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | <batazor111@gmail.com> | <batazor.ru> |

## Requirements

Kubernetes: `>= 1.28.0 || >= v1.28.0-0`

| Repository | Name | Version |
|------------|------|---------|
| https://kyverno.github.io/kyverno | kyverno | 3.0.5 |
| https://kyverno.github.io/kyverno | kyverno-policies | 3.0.4 |
| https://kyverno.github.io/policy-reporter | policy-reporter | 2.21.1 |

## Values

<table height="400px" >
	<thead>
		<th>Key</th>
		<th>Type</th>
		<th>Default</th>
		<th>Description</th>
	</thead>
	<tbody>
		<tr>
			<td id="kyverno-policies--background"><a href="./values.yaml#L78">kyverno-policies.background</a></td>
			<td>
bool
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
false
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kyverno-policies--enabled"><a href="./values.yaml#L74">kyverno-policies.enabled</a></td>
			<td>
bool
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
true
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kyverno-policies--failurePolicy"><a href="./values.yaml#L85">kyverno-policies.failurePolicy</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"Ignore"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kyverno-policies--podSecuritySeverity"><a href="./values.yaml#L76">kyverno-policies.podSecuritySeverity</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"low"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kyverno-policies--validationFailureActionByPolicy--disallow-capabilities-strict"><a href="./values.yaml#L81">kyverno-policies.validationFailureActionByPolicy.disallow-capabilities-strict</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"audit"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kyverno-policies--validationFailureActionByPolicy--disallow-host-path"><a href="./values.yaml#L82">kyverno-policies.validationFailureActionByPolicy.disallow-host-path</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"audit"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kyverno-policies--validationFailureActionByPolicy--disallow-host-ports"><a href="./values.yaml#L83">kyverno-policies.validationFailureActionByPolicy.disallow-host-ports</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"audit"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kyverno--admissionController--hostNetwork"><a href="./values.yaml#L9">kyverno.admissionController.hostNetwork</a></td>
			<td>
bool
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
false
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kyverno--admissionController--serviceMonitor--additionalLabels--release"><a href="./values.yaml#L20">kyverno.admissionController.serviceMonitor.additionalLabels.release</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"prometheus-operator"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kyverno--admissionController--serviceMonitor--enabled"><a href="./values.yaml#L17">kyverno.admissionController.serviceMonitor.enabled</a></td>
			<td>
bool
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
true
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kyverno--admissionController--tracing--address"><a href="./values.yaml#L13">kyverno.admissionController.tracing.address</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"grafana-tempo.grafana"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kyverno--admissionController--tracing--enabled"><a href="./values.yaml#L12">kyverno.admissionController.tracing.enabled</a></td>
			<td>
bool
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
true
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kyverno--admissionController--tracing--port"><a href="./values.yaml#L14">kyverno.admissionController.tracing.port</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
4317
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kyverno--backgroundController--enabled"><a href="./values.yaml#L40">kyverno.backgroundController.enabled</a></td>
			<td>
bool
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
true
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kyverno--backgroundController--serviceMonitor--additionalLabels--release"><a href="./values.yaml#L51">kyverno.backgroundController.serviceMonitor.additionalLabels.release</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"prometheus-operator"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kyverno--backgroundController--serviceMonitor--enabled"><a href="./values.yaml#L48">kyverno.backgroundController.serviceMonitor.enabled</a></td>
			<td>
bool
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
true
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kyverno--backgroundController--tracing--address"><a href="./values.yaml#L44">kyverno.backgroundController.tracing.address</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"grafana-tempo.grafana"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kyverno--backgroundController--tracing--enabled"><a href="./values.yaml#L43">kyverno.backgroundController.tracing.enabled</a></td>
			<td>
bool
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
true
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kyverno--backgroundController--tracing--port"><a href="./values.yaml#L45">kyverno.backgroundController.tracing.port</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
4317
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kyverno--cleanupController--enabled"><a href="./values.yaml#L54">kyverno.cleanupController.enabled</a></td>
			<td>
bool
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
true
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kyverno--cleanupController--logging--format"><a href="./values.yaml#L66">kyverno.cleanupController.logging.format</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"json"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kyverno--cleanupController--networkPolicy--enabled"><a href="./values.yaml#L57">kyverno.cleanupController.networkPolicy.enabled</a></td>
			<td>
bool
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
true
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kyverno--cleanupController--serviceMonitor--additionalLabels--release"><a href="./values.yaml#L63">kyverno.cleanupController.serviceMonitor.additionalLabels.release</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"prometheus-operator"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kyverno--cleanupController--serviceMonitor--enabled"><a href="./values.yaml#L60">kyverno.cleanupController.serviceMonitor.enabled</a></td>
			<td>
bool
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
true
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kyverno--cleanupController--tracing--address"><a href="./values.yaml#L70">kyverno.cleanupController.tracing.address</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"grafana-tempo.grafana"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kyverno--cleanupController--tracing--enabled"><a href="./values.yaml#L69">kyverno.cleanupController.tracing.enabled</a></td>
			<td>
bool
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
true
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kyverno--cleanupController--tracing--port"><a href="./values.yaml#L71">kyverno.cleanupController.tracing.port</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
4317
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kyverno--enabled"><a href="./values.yaml#L6">kyverno.enabled</a></td>
			<td>
bool
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
true
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kyverno--reportsController--enabled"><a href="./values.yaml#L26">kyverno.reportsController.enabled</a></td>
			<td>
bool
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
true
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kyverno--reportsController--serviceMonitor--additionalLabels--release"><a href="./values.yaml#L37">kyverno.reportsController.serviceMonitor.additionalLabels.release</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"prometheus-operator"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kyverno--reportsController--serviceMonitor--enabled"><a href="./values.yaml#L34">kyverno.reportsController.serviceMonitor.enabled</a></td>
			<td>
bool
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
true
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kyverno--reportsController--tracing--address"><a href="./values.yaml#L30">kyverno.reportsController.tracing.address</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"grafana-tempo.grafana"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kyverno--reportsController--tracing--enabled"><a href="./values.yaml#L29">kyverno.reportsController.tracing.enabled</a></td>
			<td>
bool
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
true
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kyverno--reportsController--tracing--port"><a href="./values.yaml#L31">kyverno.reportsController.tracing.port</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
4317
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kyverno--webhooksCleanup--enabled"><a href="./values.yaml#L23">kyverno.webhooksCleanup.enabled</a></td>
			<td>
bool
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
false
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="policy-reporter--enabled"><a href="./values.yaml#L88">policy-reporter.enabled</a></td>
			<td>
bool
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
true
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="policy-reporter--global--plugins--kyverno"><a href="./values.yaml#L149">policy-reporter.global.plugins.kyverno</a></td>
			<td>
bool
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
true
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="policy-reporter--grafana--folder--annotation"><a href="./values.yaml#L144">policy-reporter.grafana.folder.annotation</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"grafana_dashboard_folder"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="policy-reporter--grafana--folder--name"><a href="./values.yaml#L145">policy-reporter.grafana.folder.name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"Security"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="policy-reporter--kyvernoPlugin--enabled"><a href="./values.yaml#L137">policy-reporter.kyvernoPlugin.enabled</a></td>
			<td>
bool
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
true
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="policy-reporter--metrics--enabled"><a href="./values.yaml#L105">policy-reporter.metrics.enabled</a></td>
			<td>
bool
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
true
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="policy-reporter--monitoring--enabled"><a href="./values.yaml#L140">policy-reporter.monitoring.enabled</a></td>
			<td>
bool
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
true
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="policy-reporter--networkPolicy--enabled"><a href="./values.yaml#L99">policy-reporter.networkPolicy.enabled</a></td>
			<td>
bool
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
false
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="policy-reporter--resources--limits--cpu"><a href="./values.yaml#L92">policy-reporter.resources.limits.cpu</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"100m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="policy-reporter--resources--limits--memory"><a href="./values.yaml#L93">policy-reporter.resources.limits.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"128Mi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="policy-reporter--resources--requests--cpu"><a href="./values.yaml#L95">policy-reporter.resources.requests.cpu</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"15m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="policy-reporter--resources--requests--memory"><a href="./values.yaml#L96">policy-reporter.resources.requests.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"75Mi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="policy-reporter--rest--enabled"><a href="./values.yaml#L102">policy-reporter.rest.enabled</a></td>
			<td>
bool
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
true
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="policy-reporter--target--loki--host"><a href="./values.yaml#L153">policy-reporter.target.loki.host</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"http://grafana-loki.grafana:3100"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="policy-reporter--target--loki--minimumPriority"><a href="./values.yaml#L154">policy-reporter.target.loki.minimumPriority</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"warning"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="policy-reporter--target--loki--skipExistingOnStartup"><a href="./values.yaml#L155">policy-reporter.target.loki.skipExistingOnStartup</a></td>
			<td>
bool
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
true
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="policy-reporter--target--loki--sources[0]"><a href="./values.yaml#L157">policy-reporter.target.loki.sources[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"kyverno"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="policy-reporter--ui--enabled"><a href="./values.yaml#L108">policy-reporter.ui.enabled</a></td>
			<td>
bool
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
true
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="policy-reporter--ui--ingress--annotations--"cert-manager--io/cluster-issuer""><a href="./values.yaml#L118">policy-reporter.ui.ingress.annotations."cert-manager.io/cluster-issuer"</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"cert-manager-production"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="policy-reporter--ui--ingress--annotations--"nginx--ingress--kubernetes--io/enable-modsecurity""><a href="./values.yaml#L119">policy-reporter.ui.ingress.annotations."nginx.ingress.kubernetes.io/enable-modsecurity"</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"false"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="policy-reporter--ui--ingress--annotations--"nginx--ingress--kubernetes--io/enable-opentelemetry""><a href="./values.yaml#L121">policy-reporter.ui.ingress.annotations."nginx.ingress.kubernetes.io/enable-opentelemetry"</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"true"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="policy-reporter--ui--ingress--annotations--"nginx--ingress--kubernetes--io/enable-owasp-core-rules""><a href="./values.yaml#L120">policy-reporter.ui.ingress.annotations."nginx.ingress.kubernetes.io/enable-owasp-core-rules"</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"true"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="policy-reporter--ui--ingress--annotations--"nginx--ingress--kubernetes--io/rewrite-target""><a href="./values.yaml#L122">policy-reporter.ui.ingress.annotations."nginx.ingress.kubernetes.io/rewrite-target"</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"/$1"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="policy-reporter--ui--ingress--annotations--"nginx--ingress--kubernetes--io/use-regex""><a href="./values.yaml#L123">policy-reporter.ui.ingress.annotations."nginx.ingress.kubernetes.io/use-regex"</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"true"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="policy-reporter--ui--ingress--className"><a href="./values.yaml#L115">policy-reporter.ui.ingress.className</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"nginx"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="policy-reporter--ui--ingress--enabled"><a href="./values.yaml#L114">policy-reporter.ui.ingress.enabled</a></td>
			<td>
bool
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
true
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="policy-reporter--ui--ingress--hosts[0]--host"><a href="./values.yaml#L126">policy-reporter.ui.ingress.hosts[0].host</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"shortlink.best"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="policy-reporter--ui--ingress--hosts[0]--paths[0]--path"><a href="./values.yaml#L128">policy-reporter.ui.ingress.hosts[0].paths[0].path</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"/kyverno/?(.*)"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="policy-reporter--ui--ingress--hosts[0]--paths[0]--pathType"><a href="./values.yaml#L129">policy-reporter.ui.ingress.hosts[0].paths[0].pathType</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"Prefix"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="policy-reporter--ui--ingress--tls[0]--hosts[0]"><a href="./values.yaml#L134">policy-reporter.ui.ingress.tls[0].hosts[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"shortlink.best"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="policy-reporter--ui--ingress--tls[0]--secretName"><a href="./values.yaml#L132">policy-reporter.ui.ingress.tls[0].secretName</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"shortlink-ingress-tls"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="policy-reporter--ui--plugins--kyverno"><a href="./values.yaml#L111">policy-reporter.ui.plugins.kyverno</a></td>
			<td>
bool
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
true
</pre>
</div>
			</td>
			<td></td>
		</tr>
	</tbody>
</table>

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.11.0](https://github.com/norwoodj/helm-docs/releases/v1.11.0)
