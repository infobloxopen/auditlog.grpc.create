package auditlogging

import (
	"context"
	"fmt"

	gormResource "github.com/infobloxopen/atlas-app-toolkit/gorm/resource"
)

const (
	separator = "."
)

// Option defines options of audit middleware
type Option func(*options)

// ExclusionList is an array of operations/resources that need to be excluded from auditing
type ExclusionList []string

// ExcludeFunc ...
type ExcludeFunc func(ctx context.Context, exclusionList ExclusionList) bool

type options struct {
	excluderFunc  ExcludeFunc
	appID         string
	exclusionList ExclusionList
}

var defaultOptions = &options{
	excluderFunc:  defaultExcludeFunc,
	appID:         getDefaultAppID(),
	exclusionList: ExclusionList{},
}

func getFullMethod(service, method string) string {
	return fmt.Sprint(service + separator + method)
}

func defaultExcludeFunc(ctx context.Context, exclusionList ExclusionList) bool {
	if len(exclusionList) == 0 {
		return false
	}
	service, method, err := GetReqDetails(ctx)
	if err != nil {
		Logger(ctx).Errorf("couldn't get the req details %v", err)
		return true
	}
	fullMeth := getFullMethod(service, method)
	for _, op := range exclusionList {
		if op != "" && (op == service || op == fullMeth) {
			return true
		}
	}
	return false
}

func getDefaultAppID() string {
	return gormResource.ApplicationName()
}

func evaluateOptions(opts ...Option) *options {
	defaultOpts := &options{}
	*defaultOpts = *defaultOptions
	for _, opt := range opts {
		opt(defaultOpts)
	}
	return defaultOpts
}

// WithExcludeFunc provides a way to audit users to override exclusion logic
func WithExcludeFunc(f ExcludeFunc) Option {
	return func(o *options) {
		o.excluderFunc = f
	}
}

// WithExclusionList takes a list of resource types or actions to exclude
func WithExclusionList(l ExclusionList) Option {
	return func(o *options) {
		o.exclusionList = l
	}
}

// WithAppId provides a option to specify appID explicitely
func WithAppId(appID string) Option {
	return func(o *options) {
		o.appID = appID
	}
}

// NewExclusionList accepts service1 (resourceType), method1, method2 ..
// returns the pairs of the form service.method1, service.method2,
// if methods are not passed, whole service will be part of exclusion List
func NewExclusionList(service string, methods ...string) ExclusionList {
	var exclusionList ExclusionList
	if service == "" {
		return exclusionList
	}
	if len(methods) == 0 {
		return append(exclusionList, service)
	}
	for _, method := range methods {
		exclusionList = append(exclusionList, getFullMethod(service, method))
	}
	return exclusionList
}
