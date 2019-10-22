/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"reflect"

	"github.com/bells17/common-network-policy-operator/api/v1alpha1"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// CommonNetworkPolicyReconciler reconciles a CommonNetworkPolicy object
type CommonNetworkPolicyReconciler struct {
	client.Client
	Cache  cache.Cache
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups="",resources=namespaces,verbs=get;list;watch;
// +kubebuilder:rbac:groups=networking.k8s.io,resources=networkpolicies,verbs=*
// +kubebuilder:rbac:groups=commonnetworkpolicies.bells17.io,resources=commonnetworkpolicies,verbs=*
// +kubebuilder:rbac:groups=commonnetworkpolicies.bells17.io,resources=commonnetworkpolicies/finalizers,verbs=*

func (r *CommonNetworkPolicyReconciler) Reconcile(req reconcile.Request) (reconcile.Result, error) {
	r.Log.Info("Reconcile starting", "Namespace", req.Namespace, "Name", req.Name)

	// Fetch target CommonNetworkPolicy list
	var err error
	commonNetworkPolicies := &v1alpha1.CommonNetworkPolicyList{}
	if req.Name != "" {
		commonNetworkPolicy := v1alpha1.CommonNetworkPolicy{}
		err := r.Cache.Get(context.TODO(), types.NamespacedName{
			Name:      req.Name,
			Namespace: "",
		}, &commonNetworkPolicy)

		if err != nil {
			if !apierrors.IsNotFound(err) {
				// Maybe only delete CommonNetworkPolicy event
				r.Log.Info("Target CommonNetworkPolicy is already deleted", "Name", req.Name)
				return reconcile.Result{}, nil
			}

			r.Log.Error(err, "Fetch CommonNetworkPolicy Error", "Name", req.Name)
			return reconcile.Result{}, err
		}

		commonNetworkPolicies.Items = append(commonNetworkPolicies.Items, commonNetworkPolicy)
	} else {
		err := r.Cache.List(context.TODO(), commonNetworkPolicies)

		if err != nil {
			if !apierrors.IsNotFound(err) {
				return reconcile.Result{}, nil
			}

			r.Log.Error(err, "Fetch CommonNetworkPolicy List Error", "Name", req.Name)
			return reconcile.Result{}, err
		}
	}

	// Ensure Network Policies
	for _, commonNetworkPolicy := range commonNetworkPolicies.Items {
		// Fetch target Namespace list
		namespaceList := &corev1.NamespaceList{}
		if req.Namespace != "" {
			namespace := corev1.Namespace{}
			err := r.Cache.Get(context.TODO(), types.NamespacedName{
				Name:      req.Namespace,
				Namespace: "",
			}, &namespace)

			if err != nil {
				if !apierrors.IsNotFound(err) {
					r.Log.Info("Target Namespace is already deleted", "Name", req.Namespace)
					return reconcile.Result{}, nil
				}

				r.Log.Error(err, "Fetch Namespace Error", "Name", req.Namespace)
				return reconcile.Result{}, err
			}

			namespaceList.Items = append(namespaceList.Items, namespace)
		} else {
			err = r.Cache.List(context.TODO(), namespaceList)
			if err != nil {
				r.Log.Error(err, "Fetch Namespace List Error", "Name", req.Namespace)
				return reconcile.Result{}, err
			}
		}

		for _, namespace := range namespaceList.Items {
			// Check what is exclude
			isExclude := false
			for _, n := range commonNetworkPolicy.Spec.ExcludeNamespaces {
				if n == namespace.ObjectMeta.GetName() {
					isExclude = true
					continue
				}
			}
			if isExclude {
				continue
			}

			name := commonNetworkPolicy.Name
			if commonNetworkPolicy.Spec.NamePrefix != "" {
				name = commonNetworkPolicy.Spec.NamePrefix + "-" + name
			}

			networkPolicy := &networkingv1.NetworkPolicy{
				ObjectMeta: metav1.ObjectMeta{
					Name:      name,
					Namespace: namespace.ObjectMeta.GetName(),
				},
				Spec: commonNetworkPolicy.Spec.PolicySpec,
			}

			err := r.applyNetworkPolicy(commonNetworkPolicy, networkPolicy)
			if err != nil {
				r.Log.Error(err, "applyNetworkPolicy error")
				return reconcile.Result{}, err
			}
		}
	}

	r.Log.Info("Reconcile complete", "Namespace", req.Namespace, "Name", req.Name)
	return reconcile.Result{}, nil
}

func (r *CommonNetworkPolicyReconciler) applyNetworkPolicy(
	cnp v1alpha1.CommonNetworkPolicy,
	np *networkingv1.NetworkPolicy,
) error {
	err := controllerutil.SetControllerReference(&cnp, np, r.Scheme)
	if err != nil {
		return errors.Wrap(err, "SetControllerReference failed")
	}

	found := &networkingv1.NetworkPolicy{}
	err = r.Get(context.TODO(), types.NamespacedName{Name: np.Name, Namespace: np.Namespace}, found)
	if err != nil {
		if !apierrors.IsNotFound(err) {
			return errors.Wrap(err, "Fetch NetworkPolicy Error")
		}

		// Create a new NetworkPolicy when not created it yet
		r.Log.Info("Creating Network Policy", "Namespace", np.Namespace, "Name", np.Name)
		err = r.Create(context.TODO(), np)
		if err != nil {
			return errors.Wrap(err, "Create NetworkPolicy failed")
		}
	}

	if !reflect.DeepEqual(np.Spec, found.Spec) {
		r.Log.Info("Updating Network Policy", "Namespace", np.Namespace, "Name", np.Name)
		found.Spec = np.Spec
		err = r.Update(context.TODO(), found)
		if err != nil {
			return errors.Wrap(err, "Update NetworkPolicy failed")
		}
	}
	return nil
}
