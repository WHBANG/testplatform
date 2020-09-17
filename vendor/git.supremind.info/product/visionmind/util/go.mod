module git.supremind.info/product/visionmind/util

go 1.13

require (
	git.supremind.info/product/visionmind/com v0.0.0-20200806025148-9b55ab598d44
	github.com/facebookgo/clock v0.0.0-20150410010913-600d898af40a // indirect
	github.com/facebookgo/ensure v0.0.0-20200202191622-63f1cf65ac4c // indirect
	github.com/facebookgo/freeport v0.0.0-20150612182905-d4adf43b75b9 // indirect
	github.com/facebookgo/httpdown v0.0.0-20180706035922-5979d39b15c2 // indirect
	github.com/facebookgo/stack v0.0.0-20160209184415-751773369052 // indirect
	github.com/facebookgo/stats v0.0.0-20151006221625-1b76add642e4 // indirect
	github.com/facebookgo/subset v0.0.0-20200203212716-c811ad88dec4 // indirect
	github.com/ghodss/yaml v1.0.0
	github.com/hashicorp/consul/api v1.5.0
	github.com/jlaffaye/ftp v0.0.0-20200730135723-c2ee4fa2503b
	github.com/kavu/go_reuseport v1.5.0 // indirect
	github.com/labstack/echo/v4 v4.1.6
	github.com/prometheus/client_golang v1.7.1
	github.com/qiniu v0.0.0-00010101000000-000000000000
	github.com/satori/go.uuid v1.2.0
	github.com/stretchr/testify v1.6.1
	golang.org/x/text v0.3.3
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
	qbox.us v0.0.0-00010101000000-000000000000
	qiniupkg.com v0.0.0-00010101000000-000000000000 // indirect
	qiniupkg.com/x v0.0.0-00010101000000-000000000000
	qiniupkg.com/x/rollog.v1 v0.0.0-00010101000000-000000000000 // indirect
)

replace (
	github.com/qiniu => git.supremind.info/product/visionmind/com/dep/github.com/qiniu v0.0.0-20200805051246-1b53f0dc1990
	github.com/qiniu/db/mgoutil.v3 => git.supremind.info/product/visionmind/com/dep/github.com/qiniu/db/mgoutil.v3 v0.0.0-20200805051246-1b53f0dc1990
	qbox.us => git.supremind.info/product/visionmind/com/dep/qbox.us v0.0.0-20200805051246-1b53f0dc1990
	qiniupkg.com => git.supremind.info/product/visionmind/com/dep/qiniupkg.com v0.0.0-20200805051246-1b53f0dc1990
	qiniupkg.com/x => git.supremind.info/product/visionmind/com/dep/qiniupkg.com/x v0.0.0-20200805051246-1b53f0dc1990
	qiniupkg.com/x/rollog.v1 => git.supremind.info/product/visionmind/com/dep/qiniupkg.com/x/rollog.v1 v0.0.0-20200805051246-1b53f0dc1990
)
