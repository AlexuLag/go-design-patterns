package main

type Observable struct {
	subs []Observer
}

func (o *Observable) Subscribe(x Observer) {
	o.subs = append(o.subs, x)
}

func (o *Observable) Unsubscribe(x Observer) {
	for i, observer := range o.subs {
		if observer == x {
			// Remove the observer by slicing
			o.subs = append(o.subs[:i], o.subs[i+1:]...)
			return
		}
	}
}

func (o *Observable) Fire(data interface{}) {
	for _, observer := range o.subs {
		observer.Notify(data)
	}
}
