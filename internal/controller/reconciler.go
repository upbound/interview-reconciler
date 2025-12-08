package controller

import (
	"context"

	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/upbound/interview-reconciler/api/v1alpha1"
)

// Reconciler reconciles a PostgresConnection object.
type Reconciler struct {
	client.Client
	Log    logging.Logger
	Scheme *runtime.Scheme
}

// Reconcile implements the reconciliation loop for PostgresConnection
// resources.
func (r *Reconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := r.Log.WithValues("postgresconnection", req.NamespacedName)
	log.Debug("Reconciling PostgresConnection")

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *Reconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.PostgresConnection{}).
		Complete(r)
}
