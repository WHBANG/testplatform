module git.supremind.info/product/visionmind/sdk/vmr/go_sdk

go 1.13

require (
	git.supremind.info/product/visionmind/com v0.0.0-20200910060653-92ec7cf7cb38
	git.supremind.info/product/visionmind/util v0.0.0-20200821022334-6e587766c8ac
	github.com/bitly/go-simplejson v0.5.0
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869 // indirect
	github.com/jarcoal/httpmock v1.0.5
	github.com/qiniu/db/mgoutil.v3 v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.6.1
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
	qiniupkg.com/x v0.0.0-00010101000000-000000000000
)

replace (
	github.com/qiniu => git.supremind.info/product/visionmind/com/dep/github.com/qiniu v0.0.0-20200805051246-1b53f0dc1990
	github.com/qiniu/db/mgoutil.v3 => git.supremind.info/product/visionmind/com/dep/github.com/qiniu/db/mgoutil.v3 v0.0.0-20200805051246-1b53f0dc1990
	qbox.us => git.supremind.info/product/visionmind/com/dep/qbox.us v0.0.0-20200805051246-1b53f0dc1990
	qiniupkg.com => git.supremind.info/product/visionmind/com/dep/qiniupkg.com v0.0.0-20200805051246-1b53f0dc1990
	qiniupkg.com/x => git.supremind.info/product/visionmind/com/dep/qiniupkg.com/x v0.0.0-20200805051246-1b53f0dc1990
	qiniupkg.com/x/rollog.v1 => git.supremind.info/product/visionmind/com/dep/qiniupkg.com/x/rollog.v1 v0.0.0-20200805051246-1b53f0dc1990
)
