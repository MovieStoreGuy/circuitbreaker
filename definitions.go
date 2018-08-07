package circuitbreaker

import (
	"net/http"
)

// Switch is a simple object that allows the user to determine the route.
type Switch interface {
	// When is the generic function that will deterime if the route is accessible at this point in time.
	// The reason for passing a function that returns a bool is so that routes can be reenabled at runtime.
	CloseRouteWhen(func(*http.Route) bool) Switch

	// EnabledRoute is the http route used if the switch is enabled
	EnabledRoute(func(http.ResponseWriter, *http.Request)) Switch

	// DisabledRoute is the route taken if the switch is disabled
	DisabledRoute(func(http.ResponseWriter, *http.Request)) Switch

	// Route is the resulting route taken after all the When clauses are resolved.
	Route() (func (http.ResponseWriter, *http.Request))
}
