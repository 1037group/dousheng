module github.com/1037group/dousheng

go 1.18

replace github.com/apache/thrift => github.com/apache/thrift v0.13.0

replace go.etcd.io/etcd/server/v3 => go.etcd.io/etcd/server/v3 v3.5.7

require (
	github.com/apache/thrift v0.13.0
	github.com/bsm/redislock v0.9.0
	github.com/cloudwego/hertz v0.5.2
	github.com/cloudwego/kitex v0.4.4
	github.com/confluentinc/confluent-kafka-go/v2 v2.0.2
	github.com/disintegration/imaging v1.6.2
	github.com/hertz-contrib/jwt v1.0.2
	github.com/hertz-contrib/obs-opentelemetry/tracing v0.1.1
	github.com/hertz-contrib/pprof v0.1.0
	github.com/kitex-contrib/obs-opentelemetry v0.1.0
	github.com/kitex-contrib/obs-opentelemetry/logging/logrus v0.0.0-20221109071748-a433b0b57972
	github.com/kitex-contrib/registry-etcd v0.0.0-20221223084757-0d49e7162359
	github.com/redis/go-redis/v9 v9.0.2
	github.com/robfig/cron/v3 v3.0.0
	github.com/tencentyun/cos-go-sdk-v5 v0.7.41
	github.com/u2takey/ffmpeg-go v0.4.1
	golang.org/x/crypto v0.6.0
	gorm.io/driver/mysql v1.4.5
	gorm.io/gorm v1.24.4
	gorm.io/plugin/opentelemetry v0.1.0
)

require (
	github.com/ajg/form v1.5.1 // indirect
	github.com/andybalholm/brotli v1.0.4 // indirect
	github.com/aws/aws-sdk-go v1.44.192 // indirect
	github.com/bytedance/go-tagexpr/v2 v2.9.2 // indirect
	github.com/bytedance/gopkg v0.0.0-20220623074550-9d6d3df70991 // indirect
	github.com/bytedance/sonic v1.5.0 // indirect
	github.com/cenkalti/backoff/v4 v4.1.3 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20211019084208-fb5309c8db06 // indirect
	github.com/chenzhuoyu/iasm v0.0.0-20220818063314-28c361dae733 // indirect
	github.com/choleraehyq/pid v0.0.15 // indirect
	github.com/clbanning/mxj v1.8.4 // indirect
	github.com/cloudwego/fastpb v0.0.3 // indirect
	github.com/cloudwego/frugal v0.1.3 // indirect
	github.com/cloudwego/netpoll v0.3.1 // indirect
	github.com/cloudwego/thriftgo v0.2.4 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd/v22 v22.3.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/gavv/httpexpect v2.0.0+incompatible // indirect
	github.com/gavv/httpexpect/v2 v2.12.0 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt/v4 v4.4.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/pprof v0.0.0-20220608213341-c488b8fa1db3 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.7.0 // indirect
	github.com/henrylee2cn/ameda v1.4.10 // indirect
	github.com/henrylee2cn/goutil v0.0.0-20210127050712-89660552f6f8 // indirect
	github.com/imkira/go-interpol v1.1.0 // indirect
	github.com/jhump/protoreflect v1.12.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/klauspost/cpuid/v2 v2.1.0 // indirect
	github.com/mattn/go-sqlite3 v1.14.16 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/moul/http2curl v1.0.0 // indirect
	github.com/mozillazg/go-httpheader v0.3.1 // indirect
	github.com/nyaruka/phonenumbers v1.0.55 // indirect
	github.com/oleiade/lane v1.0.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.9.0 // indirect
	github.com/sanity-io/litter v1.5.5 // indirect
	github.com/savsgio/gotils v0.0.0-20220530130905-52f3993e8d6d // indirect
	github.com/sergi/go-diff v1.3.1 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	github.com/stretchr/testify v1.8.1 // indirect
	github.com/tidwall/gjson v1.14.3 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/u2takey/go-utils v0.3.1 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.44.0 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	github.com/yalp/jsonpath v0.0.0-20180802001716-5cc68e5049a0 // indirect
	github.com/yudai/gojsondiff v1.0.0 // indirect
	github.com/yudai/golcs v0.0.0-20170316035057-ecda9a501e82 // indirect
	go.etcd.io/etcd/api/v3 v3.5.7 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.7 // indirect
	go.etcd.io/etcd/client/v3 v3.5.7 // indirect
	go.etcd.io/etcd/server/v3 v3.5.7 // indirect
	go.opentelemetry.io/contrib/instrumentation/runtime v0.33.0 // indirect
	go.opentelemetry.io/contrib/propagators/b3 v1.9.0 // indirect
	go.opentelemetry.io/contrib/propagators/ot v1.9.0 // indirect
	go.opentelemetry.io/otel v1.9.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/internal/retry v1.9.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric v0.31.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v0.31.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.9.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.9.0 // indirect
	go.opentelemetry.io/otel/metric v0.31.0 // indirect
	go.opentelemetry.io/otel/sdk v1.9.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v0.31.0 // indirect
	go.opentelemetry.io/otel/trace v1.9.0 // indirect
	go.opentelemetry.io/proto/otlp v0.18.0 // indirect
	go.uber.org/atomic v1.8.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.17.0 // indirect
	golang.org/x/arch v0.0.0-20220722155209-00200b7164a7 // indirect
	golang.org/x/image v0.3.0 // indirect
	golang.org/x/net v0.6.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	google.golang.org/genproto v0.0.0-20220503193339-ba3ae3f07e29 // indirect
	google.golang.org/grpc v1.46.2 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/driver/sqlite v1.4.4 // indirect
	moul.io/http2curl/v2 v2.3.0 // indirect
)
