// MIT LICENSE

package commonnetworkpolicy

import (
	"context"
	"log"
	"reflect"

	commonnetworkpoliciesv1alpha1 "github.com/bells17/common-network-policy-operator/pkg/apis/commonnetworkpolicies/v1alpha1"
	customclientset "github.com/bells17/common-network-policy-operator/pkg/client/clientset/versioned"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new CommonNetworkPolicy Controller and adds it to the Manager with default RBAC. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
// USER ACTION REQUIRED: update cmd/manager/main.go to call this commonnetworkpolicies.Add(mgr) to install this Controller
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	c, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	return &ReconcileCommonNetworkPolicy{
		Client:          mgr.GetClient(),
		scheme:          mgr.GetScheme(),
		clientSet:       newClientSet(c),
		customClientSet: newCustomClientSet(c),
	}
}

// newClientSet returns a new *Kubernetes.Clientset
func newClientSet(c *rest.Config) *kubernetes.Clientset {
	clientSet, err := kubernetes.NewForConfig(c)
	if err != nil {
		log.Fatal(err)
	}

	return clientSet
}

func newCustomClientSet(c *rest.Config) customclientset.Interface {
	cli, err := customclientset.NewForConfig(c)
	if err != nil {
		panic(err.Error())
	}
	return cli
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("commonnetworkpolicy-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to CommonNetworkPolicy
	err = c.Watch(&source.Kind{Type: &commonnetworkpoliciesv1alpha1.CommonNetworkPolicy{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// Uncomment watch a NetworkPolicy created by CommonNetworkPolicy - change this for objects you create
	err = c.Watch(&source.Kind{Type: &networkingv1.NetworkPolicy{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &commonnetworkpoliciesv1alpha1.CommonNetworkPolicy{},
	})
	if err != nil {
		return err
	}

	// watch a Namespace create event
	mapFn := handler.ToRequestsFunc(
		func(a handler.MapObject) []reconcile.Request {
			return []reconcile.Request{
				{NamespacedName: types.NamespacedName{
					Name:      a.Meta.GetName(),
					Namespace: a.Meta.GetNamespace(),
				}},
			}
		})
	// Watch Namespace and trigger Reconciles for objects
	// mapped from the Namespace in the event
	err = c.Watch(
		&source.Kind{Type: &corev1.Namespace{}},
		&handler.EnqueueRequestsFromMapFunc{
			ToRequests: mapFn,
		})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileCommonNetworkPolicy{}

// ReconcileCommonNetworkPolicy reconciles a CommonNetworkPolicy object
type ReconcileCommonNetworkPolicy struct {
	client.Client
	scheme          *runtime.Scheme
	clientSet       kubernetes.Interface
	customClientSet customclientset.Interface
}

// Reconcile reads that state of the cluster for a CommonNetworkPolicy object and makes changes based on the state read
// and what is in the CommonNetworkPolicy.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  The scaffolding writes
// a Deployment as an example
// Automatically generate RBAC rules to allow the Controller to read and write Deployments
// +kubebuilder:rbac:groups=corev1,resources=namespace,verbs=get;list;watch;
// +kubebuilder:rbac:groups=networkingv1,resources=networkpolicy,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=commonnetworkpolicies.bells17.io,resources=commonnetworkpolicies,verbs=get;list;watch;create;update;patch;delete
func (r *ReconcileCommonNetworkPolicy) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	// Update all Common Network Policies
	commonNetworkPolicies, err := r.customClientSet.CommonnetworkpoliciesV1alpha1().CommonNetworkPolicies().List(metav1.ListOptions{})
	if err != nil {
		return reconcile.Result{}, err
	}

	for _, commonNetworkPolicyItem := range commonNetworkPolicies.Items {

		// Fetch the CommonNetworkPolicy instance
		instance := &commonnetworkpoliciesv1alpha1.CommonNetworkPolicy{}
		err := r.Get(context.TODO(), types.NamespacedName{
			Name:      commonNetworkPolicyItem.ObjectMeta.Name,
			Namespace: commonNetworkPolicyItem.ObjectMeta.Namespace,
		}, instance)

		if err != nil {
			if !errors.IsNotFound(err) {
				return reconcile.Result{}, err
			}
		}

		namespaceList, err := r.clientSet.CoreV1().Namespaces().List(metav1.ListOptions{})
		if err != nil {
			return reconcile.Result{}, err
		}

		for _, item := range namespaceList.Items {
			ns := item.ObjectMeta.Name

			isExclude := false
			for _, n := range instance.Spec.ExcludeNamespaces {
				if n == ns {
					isExclude = true
				}
			}
			if isExclude {
				continue
			}

			name := commonNetworkPolicyItem.Name
			if commonNetworkPolicyItem.Spec.NamePrefix != "" {
				name = instance.Spec.NamePrefix + "-" + name
			}

			networkPolicy := &networkingv1.NetworkPolicy{
				ObjectMeta: metav1.ObjectMeta{
					Name:      name,
					Namespace: ns,
				},
				Spec: commonNetworkPolicyItem.Spec.PolicySpec,
			}
			err := r.applyNetworkPolicy(instance, networkPolicy)
			if err != nil {
				return reconcile.Result{}, err
			}

		}
	}

	return reconcile.Result{}, nil
}

// applyNetworkPolicy apply the networkPolicy
func (r *ReconcileCommonNetworkPolicy) applyNetworkPolicy(instance *commonnetworkpoliciesv1alpha1.CommonNetworkPolicy, networkPolicy *networkingv1.NetworkPolicy) error {
	err := controllerutil.SetControllerReference(instance, networkPolicy, r.scheme)
	if err != nil {
		return err
	}

	found := &networkingv1.NetworkPolicy{}
	err = r.Get(context.TODO(), types.NamespacedName{Name: networkPolicy.Name, Namespace: networkPolicy.Namespace}, found)
	if err != nil && errors.IsNotFound(err) {
		log.Printf("Creating Network Policy %s/%s\n", networkPolicy.Namespace, networkPolicy.Name)
		err = r.Create(context.TODO(), networkPolicy)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	if !reflect.DeepEqual(networkPolicy.Spec, found.Spec) {
		log.Printf("Updating Network Policy %s/%s\n", networkPolicy.Namespace, networkPolicy.Name)
		found.Spec = networkPolicy.Spec
		err = r.Update(context.TODO(), found)
		if err != nil {
			return err
		}
	}
	return nil
}
