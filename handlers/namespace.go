package handlers

import (
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/util/workqueue"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var _ handler.EventHandler = (*EnqueueRequestsForNamespaceEvent)(nil)

type EnqueueRequestsForNamespaceEvent struct {
	Log logr.Logger
}

// Create is called in response to an create event - e.g. Pod Creation.
func (h *EnqueueRequestsForNamespaceEvent) Create(e event.CreateEvent, q workqueue.RateLimitingInterface) {
	if e.Meta == nil {
		h.Log.Error(nil, "CreateEvent received with no metadata", "event", e)
		return
	}

	q.Add(reconcile.Request{NamespacedName: types.NamespacedName{
		Name:      "",
		Namespace: e.Meta.GetName(),
	}})
}

// Update is called in response to an update event -  e.g. Pod Updated.
func (h *EnqueueRequestsForNamespaceEvent) Update(e event.UpdateEvent, q workqueue.RateLimitingInterface) {
}

// Delete is called in response to a delete event - e.g. Pod Deleted.
func (h *EnqueueRequestsForNamespaceEvent) Delete(e event.DeleteEvent, q workqueue.RateLimitingInterface) {
}

// Generic is called in response to an event of an unknown type or a synthetic event triggered as a cron or
// external trigger request - e.g. reconcile Autoscaling, or a Webhook.
func (h *EnqueueRequestsForNamespaceEvent) Generic(e event.GenericEvent, queue workqueue.RateLimitingInterface) {
}
