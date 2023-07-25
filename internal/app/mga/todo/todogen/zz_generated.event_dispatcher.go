// +build !ignore_autogenerated

// Code generated by mga tool. DO NOT EDIT.

package todogen

import (
	"context"
	"emperror.dev/errors"
	"github.com/sagikazarmark/modern-go-application/internal/app/mga/todo"
)

// EventBus is a generic event bus.
type EventBus interface {
	// Publish sends an event to the underlying message bus.
	Publish(ctx context.Context, event interface{}) error
}

// EventDispatcher dispatches events through the underlying generic event bus.
type EventDispatcher struct {
	bus EventBus
}

// NewEventDispatcher returns a new EventDispatcher instance.
func NewEventDispatcher(bus EventBus) EventDispatcher {
	return EventDispatcher{bus: bus}
}

// MarkedAsComplete dispatches a(n) MarkedAsComplete event.
func (d EventDispatcher) MarkedAsComplete(ctx context.Context, event todo.MarkedAsComplete) error {
	err := d.bus.Publish(ctx, event)
	if err != nil {
		return errors.WithDetails(errors.WithMessage(err, "failed to dispatch event"), "event", "MarkedAsComplete")
	}

	return nil
}
