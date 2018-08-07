package circuitbreaker

import (
  "net/http"
)

struct blad struct {
  conditions [](func(*http.Request)bool)
  enabledRoute func(http.ResponseWriter, *http.Request)
  disabledRoute func(http.ResponseWriter, *http.Request)
}

// NewDefaultSwitch returns a black switch object that can be used with any expressions that result in a bool
func NewDefaultSwitch() Switch {
  return &blad{}
}

func (b *blad) When(f func(*http.Request)bool) Switch {
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

func (b *blad) Route() (func(http.ResponseWriter, *http.Request)) {
  // I don't like this panic but I need a better way to handle it.
  // Potentially might be worth to default to dropping the connection if the route isn't defined
  if b.enabledRoute == nil {
    panic("No enabled route set")
  }
  if b.disabledRoute == nil {
    panic("No disabled route set")
  }
  return func(w http.ResponseWriter,r *http.Request) {
    for _, cond := range b.conditions {
      if !cond(r) {
        b.disabledRoute(w, r)
      }
    }
    b.enabledRoute(w, r)
  }
}
