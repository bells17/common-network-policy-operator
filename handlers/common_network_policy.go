package handlers

import (
	"reflect"

	"github.com/bells17/common-network-policy-operator/api/v1alpha1"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/util/workqueue"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var _ handler.EventHandler = (*EnqueueRequestsForCommonNetworkPolicyEvent)(nil)

type EnqueueRequestsForCommonNetworkPolicyEvent struct {
	Log logr.Logger
}

// Create is called in response to an create event - e.g. Pod Creation.
func (h *EnqueueRequestsForCommonNetworkPolicyEvent) Create(e event.CreateEvent, q workqueue.RateLimitingInterface) {
	if e.Meta == nil {
		h.Log.Error(nil, "CreateEvent received with no metadata", "event", e)
		return
	}

	q.Add(reconcile.Request{NamespacedName: types.NamespacedName{
		Name:      e.Meta.GetName(),
		Namespace: e.Meta.GetNamespace(),
	}})
}

// Update is called in response to an update event -  e.g. Pod Updated.
func (h *EnqueueRequestsForCommonNetworkPolicyEvent) Update(e event.UpdateEvent, q workqueue.RateLimitingInterface) {
	if e.MetaOld == nil {
		h.Log.Error(nil, "UpdateEvent received with no metadata", "event", e)
		return
	}

	if e.MetaNew == nil {
		h.Log.Error(nil, "UpdateEvent received with no metadata", "event", e)
		return
	}

	epOld := e.ObjectOld.(*v1alpha1.CommonNetworkPolicy)
	epNew := e.ObjectNew.(*v1alpha1.CommonNetworkPolicy)
	if !reflect.DeepEqual(epOld.Spec, epNew.Spec) {
		q.Add(reconcile.Request{NamespacedName: types.NamespacedName{
			Name:      e.MetaNew.GetName(),
			Namespace: e.MetaNew.GetNamespace(),
		}})
	}
}

// Delete is called in response to a delete event - e.g. Pod Deleted.
func (h *EnqueueRequestsForCommonNetworkPolicyEvent) Delete(e event.DeleteEvent, q workqueue.RateLimitingInterface) {
	if e.Meta == nil {
		h.Log.Error(nil, "DeleteEvent received with no metadata", "event", e)
		return
	}

	q.Add(reconcile.Request{NamespacedName: types.NamespacedName{
		Name:      e.Meta.GetName(),
		Namespace: e.Meta.GetNamespace(),
	}})
}

// Generic is called in response to an event of an unknown type or a synthetic event triggered as a cron or
// external trigger request - e.g. reconcile Autoscaling, or a Webhook.
func (h *EnqueueRequestsForCommonNetworkPolicyEvent) Generic(e event.GenericEvent, queue workqueue.RateLimitingInterface) {
}
