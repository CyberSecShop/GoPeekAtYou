module github.com/cybersecshop/gopeekatyou

go 1.20

require github.com/sirupsen/logrus v1.9.3

require (
	github.com/cybersecshop/gopeekatyou/winmon v0.0.0-20230623132026-ebde6d091124
	golang.org/x/sys v0.9.0
)

require github.com/abdfnx/gosh v0.4.0 // indirect

replace github.com/cybersecshop/gopeekatyou/winmon => ./winmon
