<div align="center">

# shortlink 

Shortlink is an open-source educational project that provides a pretty user interface and respects GDPR. 

The goal of the project is to demonstrate the practical application of microservices architecture.

[![Artifact Hub](https://img.shields.io/endpoint?url=https://artifacthub.io/badge/repository/shortlink)](https://artifacthub.io/packages/search?repo=shortlink)
[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/shortlink-org/shortlink)](https://pkg.go.dev/mod/github.com/shortlink-org/shortlink)
[![codecov](https://codecov.io/gh/shortlink-org/shortlink/branch/main/graph/badge.svg?token=Wxz5bI4QzF)](https://codecov.io/gh/shortlink-org/shortlink)
[![Go Report Card](https://goreportcard.com/badge/github.com/shortlink-org/shortlink)](https://goreportcard.com/report/github.com/shortlink-org/shortlink)
[![Releases](https://img.shields.io/github/release-pre/shortlink-org/shortlink.svg)](https://github.com/shortlink-org/shortlink/releases)
[![LICENSE](https://img.shields.io/github/license/shortlink-org/shortlink.svg)](https://github.com/shortlink-org/shortlink/blob/main/LICENSE)
[![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/3510/badge)](https://bestpractices.coreinfrastructure.org/projects/3510)
[![StackShare](http://img.shields.io/badge/tech-stack-0690fa.svg?style=flat)](https://stackshare.io/shortlink-org/shortlink)
[![FOSSA Status](https://app.fossa.com/api/projects/custom%2B396%2Fgithub.com%2Fshortlink-org%2Fshortlink.svg?type=shield)](https://app.fossa.com/projects/custom%2B396%2Fgithub.com%2Fshortlink-org%2Fshortlink?ref=badge_shield)
[![DeepSource](https://app.deepsource.com/gh/shortlink-org/shortlink.svg/?label=active+issues&show_trend=true&token=DL-zlqtnyx6CvlHCroG0Jdx5)](https://app.deepsource.com/gh/shortlink-org/shortlink/)

<hr />

<div style="align-items: center; display: flex;">
  <a href="https://www.producthunt.com/posts/shortlink-2?utm_source=badge-featured&utm_medium=badge&utm_souce=badge-shortlink&#0045;2" target="_blank"><img src="https://api.producthunt.com/widgets/embed-image/v1/featured.svg?post_id=374140&theme=light" alt="ShortLink - Get&#0032;ready&#0032;to&#0032;share&#0032;your&#0032;links&#0032;with&#0032;ease&#0033; | Product Hunt" style="width: 250px; height: 54px;" width="250" height="54" /></a>
  <img height="100px" src="https://slsa.dev/images/SLSA-Badge-full-level1.svg" alt="SLSA">
</div>

</div>
<hr />

### High Level Architecture 🚀

The project covers the entire process - from identifying Bounded Contexts to implementing microservices using
cutting-edge technologies and best practices.  
We're constantly researching the best solutions on the market so that we can benefit our
community and solve a problem for millions of people.

![shortlink-architecture](./docs/shortlink-architecture.png)
_Please [star ⭐](https://github.com/shortlink-org/shortlink/stargazers) the repo if you want us to continue developing and improving ShortLink! 😀_


### Architecture decision records (ADR)

An architecture decision record (ADR) is a document that captures an important architecture decision 
made along with its context and consequences.

+ [Docs ADR](https://github.com/joelparkerhenderson/architecture-decision-record)

**Decisions:**
  + [main decisions](./docs/ADR/README.md)
  + [ops decisions](./ops/docs/ADR/README.md)
  + [ui](./ui/nx-monorepo/docs/ADR/README.md)

### Services

<details><summary>DETAILS</summary>
<p>

| Bounded Context       | Service          | Description                                                                       | Language/Framework | Docs                                                                        | Status                                                                                                                                                                  |
|-----------------------|------------------|-----------------------------------------------------------------------------------|--------------------|-----------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| API Boundary          | api-cloudevents  | Internal GateWay                                                                  | Go                 | [docs](./internal/services/api-gateway/README.md)                           | [![App Status](https://argo.shortlink.best/api/badge?name=shortlink-api-cloudevents&revision=true)](https://argo.shortlink.best/applications/shortlink-api-cloudevents) |
| API Boundary          | api-graphql      | Internal GateWay                                                                  | Go                 | [docs](./internal/services/api-gateway/README.md)                           | [![App Status](https://argo.shortlink.best/api/badge?name=shortlink-api-graphql&revision=true)](https://argo.shortlink.best/applications/shortlink-api-graphql)         |
| API Boundary          | api-grpc-web     | Internal GateWay                                                                  | Go                 | [docs](./internal/services/api-gateway/README.md)                           | [![App Status](https://argo.shortlink.best/api/badge?name=shortlink-api-grpc-web&revision=true)](https://argo.shortlink.best/applications/shortlink-api-grpc-web)       |
| API Boundary          | api-http         | Internal GateWay                                                                  | Go                 | [docs](./internal/services/api-gateway/README.md)                           | [![App Status](https://argo.shortlink.best/api/badge?name=shortlink-api-http&revision=true)](https://argo.shortlink.best/applications/shortlink-api-http)               |
| API Boundary          | api-ws           | Websocket service                                                                 | Go                 | [docs](./internal/services/api-gateway/gateways/ws/README.md)               | [![App Status](https://argo.shortlink.best/api/badge?name=shortlink-api-ws&revision=true)](https://argo.shortlink.best/applications/shortlink-api-ws)                   |
| API Boundary          | bff-web          | BFF for web                                                                       | Go                 | [docs](./internal/services/bff-web/README.md)                               | [![App Status](https://argo.shortlink.best/api/badge?name=shortlink-bff-web&revision=true)](https://argo.shortlink.best/applications/shortlink-bff-web)                 |
| Payment Boundary      | billing          | Billing service                                                                   | Go                 | [docs](./internal/services/billing/README.md)                               | [![App Status](https://argo.shortlink.best/api/badge?name=shortlink-billing&revision=true)](https://argo.shortlink.best/applications/shortlink-billing)                 |
| Payment Boundary      | wallet           | Wallet service                                                                    | Go (Solidity)      | [docs](./internal/services/wallet/README.md)                                |                                                                                                                                                                         |
| Notification Boundary | bot              | Telegram bot                                                                      | JAVA               | [docs](./internal/services/bot/README.md)                                   |                                                                                                                                                                         |                                                                    
| Notification Boundary | newsletter       | Newsletter service                                                                | Rust               | [docs](./internal/services/newsletter/README.md)                            | [![App Status](https://argo.shortlink.best/api/badge?name=shortlink-newsletter&revision=true)](https://argo.shortlink.best/applications/shortlink-newsletter)           |                                                              
| Notification Boundary | notify           | Send notify to smtp, slack, telegram                                              | Go                 | [docs](./internal/services/notify/README.md)                                | [![App Status](https://argo.shortlink.best/api/badge?name=shortlink-notify&revision=true)](https://argo.shortlink.best/applications/shortlink-notify)                   |                                                                  
| Chat Boundary         | chat             | Chat service                                                                      | Elixir (Phoenix)   | [docs](./internal/services/chat/README.md)                                  | [![App Status](https://argo.shortlink.best/api/badge?name=shortlink-chat&revision=true)](https://argo.shortlink.best/applications/shortlink-chat)                       |                                                                   
| Integration Boundary  | chrome-extension | Chrome extension                                                                  | JavaScript         | [docs](./internal/extension/chrome-extension/README.md)                     |                                                                                                                                                                         |                                                                         
| Integration Boundary  | ai-plugin        | ChatGPT plugin                                                                    | JSON               | [docs](./ui/nx-monorepo/packages/landing/public/.well-known/ai-plugin.json) |                                                                                                                                                                         |
| ShortDB Boundary      | shortdb          | Custom database                                                                   | Go                 | [docs](./internal/services/shortdb/README.md)                               | [![App Status](https://argo.shortlink.best/api/badge?name=shortldb&revision=true)](https://argo.shortlink.best/applications/shortldb)                                   |                                                                          
| ShortDB Boundary      | shortdb-operator | Kubernetes Operator for [shortdb](./internal/services/shortdb/README.md) database | Go                 | [docs](./internal/services/shortdb-operator/README.md)                      | [![App Status](https://argo.shortlink.best/api/badge?name=shortldb-operator&revision=true)](https://argo.shortlink.best/applications/shortldb-operator)                 |                                                                 
| Platform Boundary     | csi              | CSI example                                                                       | Go                 | [docs](./internal/services/csi/README.md)                                   | [![App Status](https://argo.shortlink.best/api/badge?name=shortlink-csi&revision=true)](https://argo.shortlink.best/applications/shortlink-csi)                         |                                                                     
| Platform Boundary     | logger           | Logger service                                                                    | Go                 | [docs](./internal/services/logger/README.md)                                | [![App Status](https://argo.shortlink.best/api/badge?name=shortlink-logger&revision=true)](https://argo.shortlink.best/applications/shortlink-logger)                   |                                                                  
| Platform Boundary     | shortctl         | Shortlink CLI                                                                     | Go                 | [docs](./internal/services/cli/README.md)                                   |                                                                                                                                                                         |                                                                   
| Link Boundary         | link             | Link service                                                                      | Go                 | [docs](./internal/services/api-gateway/README.md)                           | [![App Status](https://argo.shortlink.best/api/badge?name=shortlink-link&revision=true)](https://argo.shortlink.best/applications/shortlink-link)                       |                                                                    
| Link Boundary         | proxy            | Proxy service for redirect to original URL                                        | TypeScript         | [docs](./internal/services/proxy/README.md)                                 | [![App Status](https://argo.shortlink.best/api/badge?name=shortlink-proxy&revision=true)](https://argo.shortlink.best/applications/shortlink-proxy)                     |                                                                   
| Link Boundary         | metadata         | Parser site by API                                                                | Go                 | [docs](./internal/services/metadata/README.md)                              | [![App Status](https://argo.shortlink.best/api/badge?name=shortlink-metadata&revision=true)](https://argo.shortlink.best/applications/shortlink-metadata)               |                                                                
| Marketing Boundary    | referral         | Referral program                                                                  | Python             | [docs](./internal/services/referral/README.md)                              | [![App Status](https://argo.shortlink.best/api/badge?name=shortlink-referral&revision=true)](https://argo.shortlink.best/applications/shortlink-referral)               |                                                                
| Marketing Boundary    | stats            | Stats server                                                                      | CPP                | [docs](./internal/services/stats/README.md)                                 | [![App Status](https://argo.shortlink.best/api/badge?name=shortlink-stats&revision=true)](https://argo.shortlink.best/applications/shortlink-stats)                     |                                                                   
| Delivery Boundary     | merch            | Merch store                                                                       | Go (Dapr)          | [docs](./internal/services/merch/README.md)                                 | [![App Status](https://argo.shortlink.best/api/badge?name=shortlink-merch&revision=true)](https://argo.shortlink.best/applications/shortlink-merch)                     |                                                                   
| Delivery Boundary     | support          | Support service                                                                   | PHP                | [docs](./internal/services/support/README.md)                               | [![App Status](https://argo.shortlink.best/api/badge?name=shortlink-support&revision=true)](https://argo.shortlink.best/applications/shortlink-support)                 |                                                                 
| Search Boundary       | search           | Search service                                                                    | Coming soon        | [docs](./internal/services/search/README.md)                                | [![App Status](https://argo.shortlink.best/api/badge?name=shortlink-search&revision=true)](https://argo.shortlink.best/applications/shortlink-search)                   |                                                                  

</p>
</details>

### Third-party Service

<details><summary>DETAILS</summary>
<p>

| Service       | Description                                                             | Language/Framework | Docs                                                | Status                                                                                                                                                  |
|---------------|-------------------------------------------------------------------------|--------------------|-----------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|
| ory/kratos    | User management service                                                 | Go                 | [docs](https://www.ory.sh/kratos/docs/)             | [![App Status](https://argo.shortlink.best/api/badge?name=auth&revision=true)](https://argo.shortlink.best/applications/auth)                     |          
| ory/hydra     | OAuth 2.0 Provider                                                      | Go                 | [docs](https://www.ory.sh/keto/docs/)               | [![App Status](https://argo.shortlink.best/api/badge?name=auth&revision=true)](https://argo.shortlink.best/applications/auth)                     |          
| backstage     | Backstage is an open platform for building developer portals.           | TypeScript         | [docs](https://backstage.io/docs/)                  | [![App Status](https://argo.shortlink.best/api/badge?name=backstage&revision=true)](https://argo.shortlink.best/applications/backstage)           |    
| grafana       | Grafana is the open source analytics & monitoring solution for          | More               | [docs](https://grafana.com/docs/)                   | [![App Status](https://argo.shortlink.best/api/badge?name=grafana&revision=true)](https://argo.shortlink.best/applications/grafana)               |       
| cert-manager  | Automatically provision and manage TLS certificates in Kubernetes       | Go                 | [docs](https://cert-manager.io/docs/)               | [![App Status](https://argo.shortlink.best/api/badge?name=cert-manager&revision=true)](https://argo.shortlink.best/applications/cert-manager)     |  
| istio         | Istio is an open platform to connect, manage, and secure microservices. | Go                 | [docs](https://istio.io/latest/docs/)               | [![App Status](https://argo.shortlink.best/api/badge?name=istio&revision=true)](https://argo.shortlink.best/applications/istio)                   |         
| nginx-ingress | Ingress controller for Kubernetes using NGINX                           | Go                 | [docs](https://kubernetes.github.io/ingress-nginx/) | [![App Status](https://argo.shortlink.best/api/badge?name=nginx-ingress&revision=true)](https://argo.shortlink.best/applications/nginx-ingress)   | 
| kafka         | Kafka is used as a message broker                                       | Java               | [docs](https://kafka.apache.org/)                   | [![App Status](https://argo.shortlink.best/api/badge?name=kafka&revision=true)](https://argo.shortlink.best/applications/kafka)                   |         
| keycloak      | Keycloak is an open source identity and access management solution      | Java               | [docs](https://www.keycloak.org/documentation.html) | [![App Status](https://argo.shortlink.best/api/badge?name=keycloak&revision=true)](https://argo.shortlink.best/applications/keycloak)             | 

</p>
</details>

### UI

 - [UI: README](./ui/nx-monorepo/README.md) - UI for ShortLink
 - [MOBILE: README](./ui/mobile/README.md) - Mobile app for ShortLink

### Contributing

 - [Getting Started](./CONTRIBUTING.md#getting-started)

### ChatGPT || OpenAI

This service support [ChatGPT](https://chat.openai.com/chat) as plugin.  
You can use it for ChatGPT by link `https://shortlink.best/.well-known/ai-plugin.json`

You can read official docs [here](https://platform.openai.com/docs/plugins/getting-started/running-a-plugin) for more information.

### License

[![FOSSA Status](https://app.fossa.com/api/projects/custom%2B396%2Fgithub.com%2Fshortlink-org%2Fshortlink.svg?type=large)](https://app.fossa.com/projects/custom%2B396%2Fgithub.com%2Fshortlink-org%2Fshortlink?ref=badge_large)

[mergify]: https://mergify.io

[mergify-status]: https://img.shields.io/endpoint.svg?url=https://dashboard.mergify.io/badges/shortlink-org/shortlink&style=flat
