/*
Copyright 2023.

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
	"github.com/kubergate/rudder/internal/datastore"
	"github.com/kubergate/rudder/pkg/logger"
	k8serror "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/source"
	gwapiv1b1 "sigs.k8s.io/gateway-api/apis/v1beta1"
)

// HttpRouteReconciler reconciles a HttpRoute object
type HttpRouteReconciler struct {
	Client           client.Client
	Manager          manager.Manager
	HttpRouteKVStore *datastore.HttpRouteKVStore
}

func NewHTTPRouteController(mgr manager.Manager, httpRouteKVStore *datastore.HttpRouteKVStore) error {
	httpRouteReconciler := &HttpRouteReconciler{
		Client:           mgr.GetClient(),
		Manager:          mgr,
		HttpRouteKVStore: httpRouteKVStore,
	}
	c, err := controller.New("HTTPRoute", mgr, controller.Options{Reconciler: httpRouteReconciler, MaxConcurrentReconciles: 1})
	if err != nil {
		return err
	}
	if err := c.Watch(source.Kind(mgr.GetCache(), &gwapiv1b1.HTTPRoute{}), &handler.EnqueueRequestForObject{}); err != nil {
		return err
	}
	return nil
}

//+kubebuilder:rbac:groups=rudder.kubergate.io,resources=httproutes,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=rudder.kubergate.io,resources=httproutes/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=rudder.kubergate.io,resources=httproutes/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the HttpRoute object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.13.0/pkg/reconcile
func (r *HttpRouteReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	var httpRoute gwapiv1b1.HTTPRoute
	if err := r.Client.Get(ctx, req.NamespacedName, &httpRoute); err != nil {
		found := r.HttpRouteKVStore.Contains(types.NamespacedName{Name: req.Name, Namespace: req.Namespace})
		if found && k8serror.IsNotFound(err) {
			// TODO: Handle HttpRoute deletion using finalizers if required
			err := r.HttpRouteKVStore.Delete(types.NamespacedName{Name: req.Name, Namespace: req.Namespace})
			if err != nil {
				return ctrl.Result{}, err
			}
			logger.LoggerRudder.Sugar().Infof("HttpRoute: %v deleted", req.NamespacedName)
			return ctrl.Result{}, nil
		}
		// Ignore other errors
		return ctrl.Result{}, nil
	}

	httpRouteInKVStore, ok := r.HttpRouteKVStore.Get(types.NamespacedName{Name: httpRoute.Name, Namespace: httpRoute.Namespace})
	if !ok {
		// TODO: Handle new HttpRoute logic
		_, err := r.HttpRouteKVStore.Add(types.NamespacedName{Name: httpRoute.Name, Namespace: httpRoute.Namespace}, httpRoute)
		if err != nil {
			return ctrl.Result{}, err
		}
		logger.LoggerRudder.Sugar().Infof("HttpRoute added to HttpRouteKVStore: %v", httpRoute.Name)
		return ctrl.Result{}, nil
	}
	if httpRoute.ObjectMeta.Generation > httpRouteInKVStore.ObjectMeta.Generation {
		// TODO: Handle HttpRoute update logic
		_, err := r.HttpRouteKVStore.Add(types.NamespacedName{Name: httpRoute.Name, Namespace: httpRoute.Namespace}, httpRoute)
		if err != nil {
			return ctrl.Result{}, err
		}
		logger.LoggerRudder.Sugar().Infof("HttpRoute updated in HttpRouteKVStore: %v", httpRoute.Name)
		return ctrl.Result{}, nil
	}
	return ctrl.Result{}, nil
}
