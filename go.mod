module github.com/batazor/shortlink

go 1.15

require (
	github.com/Masterminds/squirrel v1.4.0
	github.com/Microsoft/go-winio v0.4.15-0.20190919025122-fc70bd9a86b5 // indirect
	github.com/PuerkitoBio/goquery v1.5.1
	github.com/Shopify/sarama v1.27.0
	github.com/cloudevents/sdk-go/v2 v2.3.1
	github.com/container-storage-interface/spec v1.3.0
	github.com/containerd/continuity v0.0.0-20200107194136-26c1120b8d41 // indirect
	github.com/dgraph-io/badger/v2 v2.2007.2
	github.com/dgraph-io/dgo/v2 v2.2.0
	github.com/docker/docker v1.4.2-0.20200309214505-aa6a9891b09c // indirect
	github.com/getsentry/sentry-go v0.7.0
	github.com/go-chi/chi v4.1.2+incompatible
	github.com/go-chi/cors v1.1.1
	github.com/go-chi/render v1.0.1
	github.com/go-kit/kit v0.10.0
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/go-telegram-bot-api/telegram-bot-api v4.6.4+incompatible
	github.com/gocql/gocql v0.0.0-20200926162733-393f0c961220
	github.com/golang-migrate/migrate/v4 v4.13.0
	github.com/golang/protobuf v1.4.2
	github.com/google/wire v0.4.0
	github.com/gorilla/mux v1.8.0
	github.com/graph-gophers/graphql-go v0.0.0-20201003130358-c5bdf3b1108e
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.0.0-beta.5
	github.com/heptiolabs/healthcheck v0.0.0-20180807145615-6ff867650f40
	github.com/jackc/pgx/v4 v4.9.0
	github.com/jmoiron/sqlx v1.2.0
	github.com/kubernetes-csi/csi-test/v4 v4.0.1
	github.com/lib/pq v1.8.0
	github.com/markbates/pkger v0.17.1
	github.com/mattn/go-sqlite3 v1.14.4
	github.com/matttproud/golang_protobuf_extensions v1.0.2-0.20181231171920-c182affec369 // indirect
	github.com/nats-io/nats.go v1.10.0
	github.com/onsi/ginkgo v1.14.1 // indirect
	github.com/onsi/gomega v1.10.2 // indirect
	github.com/opencontainers/runc v1.0.0-rc91.0.20200707015106-819fcc687efb // indirect
	github.com/opentracing-contrib/go-grpc v0.0.0-20200813121455-4a6760c71486
	github.com/opentracing-contrib/go-stdlib v1.0.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/ory/dockertest/v3 v3.6.0
	github.com/pborman/uuid v1.2.0
	github.com/pelletier/go-toml v1.6.0 // indirect
	github.com/prometheus/client_golang v1.7.1
	github.com/scylladb/gocqlx/v2 v2.1.0
	github.com/sirupsen/logrus v1.7.0
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/cobra v1.0.0
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.7.1
	github.com/streadway/amqp v1.0.0
	github.com/stretchr/testify v1.6.1
	github.com/syndtr/goleveldb v1.0.0
	github.com/technoweenie/multipartstreamer v1.0.1 // indirect
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	github.com/uber/jaeger-lib v2.2.0+incompatible // indirect
	go.mongodb.org/mongo-driver v1.4.1
	go.uber.org/atomic v1.7.0
	go.uber.org/automaxprocs v1.3.0
	go.uber.org/goleak v1.1.10
	go.uber.org/zap v1.16.0
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a // indirect
	golang.org/x/sync v0.0.0-20200625203802-6e8e738ad208
	golang.org/x/sys v0.0.0-20200831180312-196b9ba8737a // indirect
	golang.org/x/tools v0.0.0-20200918232735-d647fc253266 // indirect
	google.golang.org/genproto v0.0.0-20201006033701-bcad7cf615f2
	google.golang.org/grpc v1.33.0-dev
	google.golang.org/protobuf v1.25.0
	gopkg.in/DATA-DOG/go-sqlmock.v1 v1.3.0 // indirect
	gopkg.in/ini.v1 v1.51.1 // indirect
	gopkg.in/rethinkdb/rethinkdb-go.v6 v6.2.1
	gotest.tools v2.2.0+incompatible // indirect
)

replace (
	github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.5
	k8s.io/api => k8s.io/api v0.18.3
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.18.3
	k8s.io/apimachinery => k8s.io/apimachinery v0.18.3
	k8s.io/apiserver => k8s.io/apiserver v0.18.3
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.18.3
	k8s.io/client-go => k8s.io/client-go v0.18.3
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.18.3
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.18.3
	k8s.io/code-generator => k8s.io/code-generator v0.18.3
	k8s.io/component-base => k8s.io/component-base v0.18.3
	k8s.io/cri-api => k8s.io/cri-api v0.18.3
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.18.3
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.18.3
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.18.3
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.18.3
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.18.3
	k8s.io/kubectl => k8s.io/kubectl v0.18.3
	k8s.io/kubelet => k8s.io/kubelet v0.18.3
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.18.3
	k8s.io/metrics => k8s.io/metrics v0.18.3
	k8s.io/mount-utils => k8s.io/mount-utils v0.20.0-alpha.3
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.18.3
)
