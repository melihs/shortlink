# shortlink-bot

![Version: 0.3.1](https://img.shields.io/badge/Version-0.3.1-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

Shortlink service for get bot by URL

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
			<td id="deploy--affinity"><a href="./values.yaml#L58">deploy.affinity</a></td>
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
			<td id="deploy--annotations"><a href="./values.yaml#L41">deploy.annotations</a></td>
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
			<td id="deploy--env--MQ_ENABLED"><a href="./values.yaml#L32">deploy.env.MQ_ENABLED</a></td>
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
			<td id="deploy--env--MQ_TYPE"><a href="./values.yaml#L33">deploy.env.MQ_TYPE</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"rabbitmq"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="deploy--env--TELEGRAM_BOT_TOKEN"><a href="./values.yaml#L35">deploy.env.TELEGRAM_BOT_TOKEN</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"YOUR_TELEGRAM_TOKEN"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="deploy--env--TELEGRAM_BOT_USERNAME"><a href="./values.yaml#L36">deploy.env.TELEGRAM_BOT_USERNAME</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"shortlink_my_bot"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="deploy--env--TRACER_URI"><a href="./values.yaml#L34">deploy.env.TRACER_URI</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"grafana-tempo.grafana:4317"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="deploy--image--pullPolicy"><a href="./values.yaml#L49">deploy.image.pullPolicy</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"IfNotPresent"
</pre>
</div>
			</td>
			<td>Global imagePullPolicy Default: 'Always' if image tag is 'latest', else 'IfNotPresent' Ref: http://kubernetes.io/docs/user-guide/images/#pre-pulling-images</td>
		</tr>
		<tr>
			<td id="deploy--image--repository"><a href="./values.yaml#L44">deploy.image.repository</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"registry.gitlab.com/shortlink-org/shortlink/bot"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="deploy--image--tag"><a href="./values.yaml#L45">deploy.image.tag</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"0.16.63"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="deploy--imagePullSecrets"><a href="./values.yaml#L38">deploy.imagePullSecrets</a></td>
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
  "failureThreshold": 1,
  "httpGet": {
    "path": "/live",
    "port": 9090
  },
  "initialDelaySeconds": 5,
  "periodSeconds": 5,
  "successThreshold": 1
}
</pre>
</div>
			</td>
			<td>define a liveness probe that checks every 5 seconds, starting after 5 seconds</td>
		</tr>
		<tr>
			<td id="deploy--nodeSelector"><a href="./values.yaml#L54">deploy.nodeSelector</a></td>
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
			<td>Node labels and tolerations for pod assignment ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#nodeselector ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#taints-and-tolerations-beta-feature</td>
		</tr>
		<tr>
			<td id="deploy--podSecurityContext--fsGroup"><a href="./values.yaml#L94">deploy.podSecurityContext.fsGroup</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
1000
</pre>
</div>
			</td>
			<td>fsGroup is the group ID associated with the container</td>
		</tr>
		<tr>
			<td id="deploy--readinessProbe"><a href="./values.yaml#L71">deploy.readinessProbe</a></td>
			<td>
object
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
{
  "failureThreshold": 30,
  "httpGet": {
    "path": "/ready",
    "port": 9090
  },
  "initialDelaySeconds": 5,
  "periodSeconds": 5,
  "successThreshold": 1
}
</pre>
</div>
			</td>
			<td>define a readiness probe that checks every 5 seconds, starting after 5 seconds</td>
		</tr>
		<tr>
			<td id="deploy--replicaCount"><a href="./values.yaml#L29">deploy.replicaCount</a></td>
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
			<td id="deploy--resources--limits"><a href="./values.yaml#L85">deploy.resources.limits</a></td>
			<td>
object
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
{
  "cpu": "100m",
  "memory": "128Mi"
}
</pre>
</div>
			</td>
			<td>We usually recommend not to specify default resources and to leave this as a conscious choice for the user. This also increases chances charts run on environments with little resources, such as Minikube. If you do want to specify resources, uncomment the following lines, adjust them as necessary, and remove the curly braces after 'resources:'.</td>
		</tr>
		<tr>
			<td id="deploy--resources--requests--cpu"><a href="./values.yaml#L89">deploy.resources.requests.cpu</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"10m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="deploy--resources--requests--memory"><a href="./values.yaml#L90">deploy.resources.requests.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"32Mi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="deploy--securityContext"><a href="./values.yaml#L99">deploy.securityContext</a></td>
			<td>
object
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
{
  "allowPrivilegeEscalation": false,
  "capabilities": {
    "drop": [
      "ALL"
    ]
  },
  "readOnlyRootFilesystem": "true",
  "runAsGroup": 1000,
  "runAsNonRoot": true,
  "runAsUser": 1000
}
</pre>
</div>
			</td>
			<td>Security Context policies for controller pods See https://kubernetes.io/docs/tasks/administer-cluster/sysctl-cluster/ for notes on enabling and using sysctls</td>
		</tr>
		<tr>
			<td id="deploy--tolerations"><a href="./values.yaml#L56">deploy.tolerations</a></td>
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
			<td id="monitoring--enabled"><a href="./values.yaml#L120">monitoring.enabled</a></td>
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
			<td id="podDisruptionBudget--enabled"><a href="./values.yaml#L125">podDisruptionBudget.enabled</a></td>
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
			<td id="secret--enabled"><a href="./values.yaml#L13">secret.enabled</a></td>
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
			<td id="secret--grpcIntermediateCA"><a href="./values.yaml#L22">secret.grpcIntermediateCA</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"-----BEGIN CERTIFICATE-----\nYour CA...\n-----END CERTIFICATE-----\n"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="secret--grpcServerCert"><a href="./values.yaml#L14">secret.grpcServerCert</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"-----BEGIN CERTIFICATE-----\nYour cert...\n-----END CERTIFICATE-----\n"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="secret--grpcServerKey"><a href="./values.yaml#L18">secret.grpcServerKey</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"-----BEGIN EC PRIVATE KEY-----\nYour key...\n-----END EC PRIVATE KEY-----\n"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="service--ports"><a href="./values.yaml#L115">service.ports</a></td>
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
			<td id="service--type"><a href="./values.yaml#L114">service.type</a></td>
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
