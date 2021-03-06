package circuitbreaker

import (
	"net/http"
)

type bland struct {
	conditions    [](func(*http.Request) bool)
	enabledRoute  func(http.ResponseWriter, *http.Request)
	disabledRoute func(http.ResponseWriter, *http.Request)
}

func dropConnection(w http.ResponseWriter, r *http.Request) {
	r.Body.Close()
}

// NewDefaultSwitch returns a black switch object that can be used with any expressions that result in a bool
func NewDefaultSwitch() Switch {
	return &bland{
		enabledRoute:  dropConnection,
		disabledRoute: dropConnection,
	}
}

func (b *bland) Open(f func(*http.Request) bool) Switch {
	b.conditions = append(b.conditions, f)
	return b
}

func (b *bland) EnabledRoute(f func(http.ResponseWriter, *http.Request)) Switch {
	b.enabledRoute = f
	return b
}

func (b *bland) DisabledRoute(f func(http.ResponseWriter, *http.Request)) Switch {
	b.disabledRoute = f
	return b
}

func (b *bland) Route() func(http.ResponseWriter, *http.Request) {
	// Potential spot for extra functional fun
	return func(w http.ResponseWriter, r *http.Request) {
		for _, cond := range b.conditions {
			if !cond(r) {
				b.disabledRoute(w, r)
				return
			}
		}
		b.enabledRoute(w, r)
	}
}
