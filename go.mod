module git.supremind.info/testplatform

go 1.14

require (
	git.supremind.info/product/visionmind/com v0.0.0-20200910060653-92ec7cf7cb38
	git.supremind.info/product/visionmind/sdk/vmr v0.0.0-20200910063422-31685e83acf1 // indirect
	git.supremind.info/product/visionmind/sdk/vmr/go_sdk v0.0.0-20200910063422-31685e83acf1
	git.supremind.info/product/visionmind/util v0.0.0-20200821022334-6e587766c8ac // indirect
	git.supremind.info/products/atom/apiserver/client v1.1.0
	git.supremind.info/products/atom/com v1.3.0 // indirect
	git.supremind.info/products/atom/proto/go/api v1.0.22
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/bndr/gojenkins v1.0.1
	github.com/gin-gonic/gin v1.6.3
	github.com/globalsign/mgo v0.0.0-20190517090918-73267e130ca1 // indirect
	github.com/go-playground/validator/v10 v10.3.0 // indirect
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.6.0
	github.com/smartystreets/assertions v1.0.1 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/swaggo/files v0.0.0-20190704085106-630677cd5c14
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.5.1
	golang.org/x/tools v0.0.0-20200324035526-dbf25ea225ce // indirect
	google.golang.org/grpc v1.31.0 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
	gopkg.in/yaml.v2 v2.3.0 // indirect
	honnef.co/go/tools v0.0.1-2020.1.3 // indirect
)

replace (
	github.com/qiniu => git.supremind.info/product/visionmind/com/dep/github.com/qiniu v0.0.0-20200805051246-1b53f0dc1990
	github.com/qiniu/db/mgoutil.v3 => git.supremind.info/product/visionmind/com/dep/github.com/qiniu/db/mgoutil.v3 v0.0.0-20200805051246-1b53f0dc1990
	qbox.us => git.supremind.info/product/visionmind/com/dep/qbox.us v0.0.0-20200805051246-1b53f0dc1990
	qiniupkg.com => git.supremind.info/product/visionmind/com/dep/qiniupkg.com v0.0.0-20200805051246-1b53f0dc1990
	qiniupkg.com/x => git.supremind.info/product/visionmind/com/dep/qiniupkg.com/x v0.0.0-20200805051246-1b53f0dc1990
	qiniupkg.com/x/rollog.v1 => git.supremind.info/product/visionmind/com/dep/qiniupkg.com/x/rollog.v1 v0.0.0-20200805051246-1b53f0dc1990
)
