package circuitbreaker

import (
	"net/http"
)

type blad struct {
	conditions    [](func(*http.Request) bool)
	enabledRoute  func(http.ResponseWriter, *http.Request)
	disabledRoute func(http.ResponseWriter, *http.Request)
}

func dropConnection(w http.ResponseWriter, r *http.Request) {
	// Need to figure out how best to drop the connection without leaking information
}

// NewDefaultSwitch returns a black switch object that can be used with any expressions that result in a bool
func NewDefaultSwitch() Switch {
	return &blad{
		enabledRoute:  dropConnection,
		disabledRoute: dropConnection,
	}
}

func (b *blad) CloseRouteWhen(f func(*http.Request) bool) Switch {
	b.conditions = append(b.conditions, f)
	return b
}

func (b *blad) EnabledRoute(f func(http.ResponseWriter, *http.Request)) Switch {
	b.enabledRoute = f
	return b
}

func (b *blad) DisabledRoute(f func(http.ResponseWriter, *http.Request)) Switch {
	b.disabledRoute = f
	return b
}

func (b *blad) Route() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, cond := range b.conditions {
			if !cond(r) {
				b.disabledRoute(w, r)
			}
		}
		b.enabledRoute(w, r)
	}
}
