module github.com/shortlink-org/shortlink/internal/services/csi

go 1.21

require (
	github.com/container-storage-interface/spec v1.9.0
	github.com/golang/glog v1.1.2
	github.com/google/uuid v1.4.0
	github.com/google/wire v0.5.1-0.20220620021424-0675cdc9191c
	github.com/kubernetes-csi/csi-test/v5 v5.1.0
	github.com/spf13/cobra v1.7.0
	github.com/spf13/viper v1.17.0
	github.com/stretchr/testify v1.8.4
	go.opentelemetry.io/otel/trace v1.19.0
	golang.org/x/net v0.17.0
	golang.org/x/sync v0.4.0
	google.golang.org/grpc v1.59.0
	google.golang.org/protobuf v1.31.0
	k8s.io/kubernetes v1.28.3
	k8s.io/mount-utils v0.28.3
	k8s.io/utils v0.0.0-20230726121419-3b25d923346b
)

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-task/slim-sprig v0.0.0-20230315185526-52ccab3ef572 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/pprof v0.0.0-20230510103437-eeec1cb781c3 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mitchellh/mapstructure v1.5.1-0.20220423185008-bf980b35cac4 // indirect
	github.com/moby/sys/mountinfo v0.6.2 // indirect
	github.com/onsi/ginkgo/v2 v2.13.0 // indirect
	github.com/onsi/gomega v1.29.0 // indirect
	github.com/pelletier/go-toml/v2 v2.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/rogpeppe/go-internal v1.11.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/sagikazarmark/locafero v0.3.0 // indirect
	github.com/sagikazarmark/slog-shim v0.1.0 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/afero v1.10.0 // indirect
	github.com/spf13/cast v1.5.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	go.opentelemetry.io/otel v1.19.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/exp v0.0.0-20230905200255-921286631fa9 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	golang.org/x/tools v0.13.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231016165738-49dd2c1f3d0b // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/apimachinery v0.28.3 // indirect
	k8s.io/klog/v2 v2.100.1 // indirect
)

replace (
	k8s.io/api => k8s.io/api v0.28.3
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.28.3
	k8s.io/apimachinery => k8s.io/apimachinery v0.28.3
	k8s.io/apiserver => k8s.io/apiserver v0.28.3
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.28.3
	k8s.io/client-go => k8s.io/client-go v0.28.3
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.28.3
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.28.3
	k8s.io/code-generator => k8s.io/code-generator v0.28.3
	k8s.io/component-base => k8s.io/component-base v0.28.3
	k8s.io/component-helpers => k8s.io/component-helpers v0.28.3
	k8s.io/controller-manager => k8s.io/controller-manager v0.28.3
	k8s.io/cri-api => k8s.io/cri-api v0.28.3
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.28.3
	k8s.io/dynamic-resource-allocation => k8s.io/dynamic-resource-allocation v0.28.3
	k8s.io/kms => k8s.io/kms v0.28.3
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.28.3
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.28.3
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.28.3
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.28.3
	k8s.io/kubectl => k8s.io/kubectl v0.28.3
	k8s.io/kubelet => k8s.io/kubelet v0.28.3
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.28.3
	k8s.io/metrics => k8s.io/metrics v0.28.3
	k8s.io/mount-utils => k8s.io/mount-utils v0.28.3
	k8s.io/pod-security-admission => k8s.io/pod-security-admission v0.28.3
)
