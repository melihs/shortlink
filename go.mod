module github.com/batazor/shortlink

go 1.15

require (
	github.com/Masterminds/squirrel v1.5.0
	github.com/PuerkitoBio/goquery v1.6.0
	github.com/Shopify/sarama v1.27.2
	github.com/cloudevents/sdk-go/v2 v2.3.1
	github.com/container-storage-interface/spec v1.3.0
	github.com/containerd/continuity v0.0.0-20200107194136-26c1120b8d41 // indirect
	github.com/dgraph-io/badger/v2 v2.2007.2
	github.com/dgraph-io/dgo/v2 v2.2.0
	github.com/getsentry/sentry-go v0.9.0
	github.com/go-chi/chi v4.1.2+incompatible
	github.com/go-chi/cors v1.1.1
	github.com/go-chi/render v1.0.1
	github.com/go-kit/kit v0.10.0
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/go-telegram-bot-api/telegram-bot-api v4.6.4+incompatible
	github.com/gocql/gocql v0.0.0-20201209090715-f485b5f9159c
	github.com/golang-migrate/migrate/v4 v4.14.1
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.4.3
	github.com/google/wire v0.4.0
	github.com/gorilla/mux v1.8.0
	github.com/graph-gophers/graphql-go v0.0.0-20201113091052-beb923fada29
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.0.1
	github.com/heptiolabs/healthcheck v0.0.0-20180807145615-6ff867650f40
	github.com/jackc/pgx/v4 v4.10.0
	github.com/jmoiron/sqlx v1.2.0
	github.com/kubernetes-csi/csi-test/v4 v4.0.2
	github.com/lib/pq v1.9.0
	github.com/markbates/pkger v0.17.1
	github.com/mattn/go-sqlite3 v1.14.5
	github.com/nats-io/nats.go v1.10.0
	github.com/onsi/ginkgo v1.14.1 // indirect
	github.com/onsi/gomega v1.10.2 // indirect
	github.com/opentracing-contrib/go-grpc v0.0.0-20200813121455-4a6760c71486
	github.com/opentracing-contrib/go-stdlib v1.0.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/ory/dockertest/v3 v3.6.2
	github.com/pborman/uuid v1.2.1
	github.com/prometheus/client_golang v1.8.0
	github.com/scylladb/gocqlx/v2 v2.3.0
	github.com/sirupsen/logrus v1.7.0
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.1
	github.com/streadway/amqp v1.0.0
	github.com/stretchr/testify v1.6.1
	github.com/syndtr/goleveldb v1.0.0
	github.com/technoweenie/multipartstreamer v1.0.1 // indirect
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	github.com/uber/jaeger-lib v2.2.0+incompatible // indirect
	go.mongodb.org/mongo-driver v1.4.4
	go.uber.org/atomic v1.7.0
	go.uber.org/automaxprocs v1.3.0
	go.uber.org/goleak v1.1.10
	go.uber.org/zap v1.16.0
	golang.org/x/crypto v0.0.0-20201208171446-5f87f3452ae9 // indirect
	golang.org/x/mod v0.4.0 // indirect
	golang.org/x/net v0.0.0-20201209123823-ac852fbbde11
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a
	golang.org/x/tools v0.0.0-20201124202034-299f270db459 // indirect
	google.golang.org/api v0.30.0
	google.golang.org/genproto v0.0.0-20201214200347-8c77b98c765d
	google.golang.org/grpc v1.34.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.0.1
	google.golang.org/protobuf v1.25.0
	gopkg.in/DATA-DOG/go-sqlmock.v1 v1.3.0 // indirect
	gopkg.in/rethinkdb/rethinkdb-go.v6 v6.2.1
	gopkg.in/yaml.v2 v2.4.0 // indirect
	k8s.io/kubernetes v1.20.0
	k8s.io/mount-utils v0.20.0
	k8s.io/utils v0.0.0-20201110183641-67b214c5f920
)

replace (
	k8s.io/api => k8s.io/api v0.20.0-beta.0
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.20.0-beta.0
	k8s.io/apimachinery => k8s.io/apimachinery v0.20.0-beta.0
	k8s.io/apiserver => k8s.io/apiserver v0.20.0-beta.0
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.20.0-beta.0
	k8s.io/client-go => k8s.io/client-go v0.20.0-beta.0
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.20.0-beta.0
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.20.0-beta.0
	k8s.io/code-generator => k8s.io/code-generator v0.20.0-beta.0
	k8s.io/component-base => k8s.io/component-base v0.20.0-beta.0
	k8s.io/component-helpers => k8s.io/component-helpers v0.20.0-beta.0
	k8s.io/controller-manager => k8s.io/controller-manager v0.20.0-beta.0
	k8s.io/cri-api => k8s.io/cri-api v0.20.0-beta.0
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.20.0-beta.0
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.20.0-beta.0
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.20.0-beta.0
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.20.0-beta.0
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.20.0-beta.0
	k8s.io/kubectl => k8s.io/kubectl v0.20.0-beta.0
	k8s.io/kubelet => k8s.io/kubelet v0.20.0-beta.0
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.20.0-beta.0
	k8s.io/metrics => k8s.io/metrics v0.20.0-beta.0
	k8s.io/mount-utils => k8s.io/mount-utils v0.20.0-beta.0
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.20.0-beta.0
)
