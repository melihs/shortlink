# backstage

![Version: 0.3.1](https://img.shields.io/badge/Version-0.3.1-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

Shortlink landing service

**Homepage:** <https://batazor.github.io/shortlink/>

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | <batazor111@gmail.com> | <batazor.ru> |

## Source Code

* <https://github.com/shortlink-org/shortlink>

## Requirements

Kubernetes: `>= 1.28.0 || >= v1.28.0-0`

| Repository | Name | Version |
|------------|------|---------|
| file://../shortlink-template | shortlink-template | 0.6.1 |

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
			<td id="deploy--annotations"><a href="./values.yaml#L50">deploy.annotations</a></td>
			<td>
object
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
{}
</pre>
</div>
			</td>
			<td>Annotations to be added to controller pods</td>
		</tr>
		<tr>
			<td id="deploy--image--pullPolicy"><a href="./values.yaml#L58">deploy.image.pullPolicy</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"Always"
</pre>
</div>
			</td>
			<td>Global imagePullPolicy Default: 'Always' if image tag is 'latest', else 'IfNotPresent' Ref: http://kubernetes.io/docs/user-guide/images/#pre-pulling-images</td>
		</tr>
		<tr>
			<td id="deploy--image--repository"><a href="./values.yaml#L53">deploy.image.repository</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"registry.gitlab.com/shortlink-org/shortlink/backstage"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="deploy--image--tag"><a href="./values.yaml#L54">deploy.image.tag</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"0.16.49"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="deploy--imagePullSecrets"><a href="./values.yaml#L36">deploy.imagePullSecrets</a></td>
			<td>
list
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
[]
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="deploy--livenessProbe"><a href="./values.yaml#L61">deploy.livenessProbe</a></td>
			<td>
object
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
{
  "failureThreshold": 30,
  "httpGet": {
    "path": "/healthcheck",
    "port": 7007
  },
  "httpHeaders": [
    {
      "name": "Host",
      "value": "shortlink.best"
    }
  ],
  "initialDelaySeconds": 300,
  "timeoutSeconds": 60
}
</pre>
</div>
			</td>
			<td>define a liveness probe that checks every 5 seconds, starting after 5 seconds</td>
		</tr>
		<tr>
			<td id="deploy--readinessProbe"><a href="./values.yaml#L73">deploy.readinessProbe</a></td>
			<td>
object
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
{
  "httpGet": {
    "path": "/healthcheck",
    "port": 7007
  },
  "httpHeaders": [
    {
      "name": "Host",
      "value": "shortlink.best"
    }
  ],
  "initialDelaySeconds": 120,
  "timeoutSeconds": 15
}
</pre>
</div>
			</td>
			<td>define a readiness probe that checks every 5 seconds, starting after 5 seconds</td>
		</tr>
		<tr>
			<td id="deploy--replicaCount"><a href="./values.yaml#L34">deploy.replicaCount</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
1
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="deploy--resources--limits--memory"><a href="./values.yaml#L46">deploy.resources.limits.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"248Mi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="deploy--resources--requests"><a href="./values.yaml#L47">deploy.resources.requests</a></td>
			<td>
object
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
{}
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="deploy--securityContext"><a href="./values.yaml#L86">deploy.securityContext</a></td>
			<td>
object
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
{
  "allowPrivilegeEscalation": false,
  "readOnlyRootFilesystem": "false"
}
</pre>
</div>
			</td>
			<td>Security Context policies for controller pods See https://kubernetes.io/docs/tasks/administer-cluster/sysctl-cluster/ for notes on enabling and using sysctls</td>
		</tr>
		<tr>
			<td id="deploy--strategy--rollingUpdate--maxSurge"><a href="./values.yaml#L41">deploy.strategy.rollingUpdate.maxSurge</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
1
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="deploy--strategy--rollingUpdate--maxUnavailable"><a href="./values.yaml#L42">deploy.strategy.rollingUpdate.maxUnavailable</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
0
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="deploy--strategy--type"><a href="./values.yaml#L39">deploy.strategy.type</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"RollingUpdate"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="deploy--volumes[0]--mountPath"><a href="./values.yaml#L92">deploy.volumes[0].mountPath</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"/tmp"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="deploy--volumes[0]--name"><a href="./values.yaml#L91">deploy.volumes[0].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"tmp"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="deploy--volumes[0]--type"><a href="./values.yaml#L93">deploy.volumes[0].type</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"emptyDir"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress--annotations--"cert-manager--io/cluster-issuer""><a href="./values.yaml#L15">ingress.annotations."cert-manager.io/cluster-issuer"</a></td>
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
			<td id="ingress--annotations--"nginx--ingress--kubernetes--io/enable-modsecurity""><a href="./values.yaml#L16">ingress.annotations."nginx.ingress.kubernetes.io/enable-modsecurity"</a></td>
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
			<td id="ingress--annotations--"nginx--ingress--kubernetes--io/enable-opentelemetry""><a href="./values.yaml#L18">ingress.annotations."nginx.ingress.kubernetes.io/enable-opentelemetry"</a></td>
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
			<td id="ingress--annotations--"nginx--ingress--kubernetes--io/enable-owasp-core-rules""><a href="./values.yaml#L17">ingress.annotations."nginx.ingress.kubernetes.io/enable-owasp-core-rules"</a></td>
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
			<td id="ingress--enabled"><a href="./values.yaml#L12">ingress.enabled</a></td>
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
			<td id="ingress--hostname"><a href="./values.yaml#L20">ingress.hostname</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"backstage.shortlink.best"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress--path"><a href="./values.yaml#L25">ingress.path</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"/"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress--service--name"><a href="./values.yaml#L22">ingress.service.name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"backstage"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress--service--port"><a href="./values.yaml#L23">ingress.service.port</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
7007
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress--tls[0]--hosts[0]"><a href="./values.yaml#L30">ingress.tls[0].hosts[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"backstage.shortlink.best"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress--tls[0]--secretName"><a href="./values.yaml#L28">ingress.tls[0].secretName</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"backstage-tls"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress--type"><a href="./values.yaml#L13">ingress.type</a></td>
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
			<td id="monitoring--enabled"><a href="./values.yaml#L105">monitoring.enabled</a></td>
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
			<td id="service--ports[0]--name"><a href="./values.yaml#L98">service.ports[0].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"http"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="service--ports[0]--port"><a href="./values.yaml#L99">service.ports[0].port</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
7007
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="service--ports[0]--protocol"><a href="./values.yaml#L100">service.ports[0].protocol</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"TCP"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="service--ports[0]--public"><a href="./values.yaml#L101">service.ports[0].public</a></td>
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
			<td id="service--type"><a href="./values.yaml#L96">service.type</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"ClusterIP"
</pre>
</div>
			</td>
			<td></td>
		</tr>
	</tbody>
</table>

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.11.0](https://github.com/norwoodj/helm-docs/releases/v1.11.0)
