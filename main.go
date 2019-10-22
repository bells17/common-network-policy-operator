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

package main

import (
	"flag"
	"os"

	"github.com/bells17/common-network-policy-operator/api/v1alpha1"
	"github.com/bells17/common-network-policy-operator/controllers"
	"github.com/bells17/common-network-policy-operator/handlers"
	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/source"
	// +kubebuilder:scaffold:imports
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

func init() {
	_ = clientgoscheme.AddToScheme(scheme)

	_ = v1alpha1.AddToScheme(scheme)
	// +kubebuilder:scaffold:scheme
}

func main() {
	var metricsAddr string
	var enableLeaderElection bool
	flag.StringVar(&metricsAddr, "metrics-addr", ":8080", "The address the metric endpoint binds to.")
	flag.BoolVar(&enableLeaderElection, "enable-leader-election", false,
		"Enable leader election for controller manager. Enabling this will ensure there is only one active controller manager.")
	flag.Parse()

	ctrl.SetLogger(zap.Logger(true))

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:             scheme,
		MetricsBindAddress: metricsAddr,
		LeaderElection:     enableLeaderElection,
		Port:               9443,
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	c, err := controller.New(
		"common-network-policy-controller",
		mgr,
		controller.Options{
			Reconciler: &controllers.CommonNetworkPolicyReconciler{
				Client: mgr.GetClient(),
				Cache:  mgr.GetCache(),
				Log:    ctrl.Log.WithName("controllers").WithName("CommonNetworkPolicy"),
				Scheme: scheme,
			},
		})

	if err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "CommonNetworkPolicy")
		os.Exit(1)
	}

	err = watchClusterEvents(c, ctrl.Log)
	if err != nil {
		setupLog.Error(err, "unable to watch events")
		os.Exit(1)
	}

	// +kubebuilder:scaffold:builder

	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}

func watchClusterEvents(c controller.Controller, log logr.Logger) error {
	err := c.Watch(
		&source.Kind{Type: &v1alpha1.CommonNetworkPolicy{}},
		&handlers.EnqueueRequestsForCommonNetworkPolicyEvent{
			Log: log.WithName("handlers").WithName("CommonNetworkPolicy"),
		},
	)
	if err != nil {
		return err
	}

	err = c.Watch(
		&source.Kind{Type: &networkingv1.NetworkPolicy{}},
		&handlers.EnqueueRequestsForNetworkPolicyEvent{
			Log: log.WithName("handlers").WithName("NetworkPolicy"),
		},
	)
	if err != nil {
		return err
	}

	err = c.Watch(
		&source.Kind{Type: &corev1.Namespace{}},
		&handlers.EnqueueRequestsForNamespaceEvent{
			Log: log.WithName("handlers").WithName("Namespace"),
		},
	)
	if err != nil {
		return err
	}

	return nil
}
