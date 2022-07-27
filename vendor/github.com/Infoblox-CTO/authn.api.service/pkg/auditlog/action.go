package auditlog

import (
	"net/http"
)

type Action int

const (
	ActionUnknown Action = iota
	ActionUpdate
	ActionCreate
	ActionDelete
)

func ActionFromHTTP(method string) Action {

	var ret Action

	switch method {
	case http.MethodPut:
		ret = ActionUpdate
	case http.MethodPost:
		ret = ActionCreate
	case http.MethodDelete:
		ret = ActionDelete
	}

	return ret
}

func (a Action) String() string {

	var ret string

	switch a {
	case ActionUpdate:
		ret = "Update"
	case ActionCreate:
		ret = "Create"
	case ActionDelete:
		ret = "Delete"
	default:
		ret = "Unknown"
	}

	return ret
}

func (a Action) Status() string {

	var ret string

	switch a {
	case ActionUpdate:
		ret = "Updated"
	case ActionCreate:
		ret = "Created"
	case ActionDelete:
		ret = "Deleted"
	default:
		ret = "Unknown"
	}

	return ret
}

func (a Action) MessageFormat() string {

	var ret string

	switch a {
	case ActionUpdate:
		ret = `%s with name "%s" has been updated`
	case ActionCreate:
		ret = `%s with name "%s" has been created`
	case ActionDelete:
		ret = `%s with name "%s" has been deleted`
	default:
		ret = `unknown action for resource %s with name "%s"`
	}

	return ret
}
