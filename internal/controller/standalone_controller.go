/*
Copyright 2025.

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

package controller

import (
	"context"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	k8sioredisv1 "github.com/douzengrui001/redis_operator/api/v1"
)

// StandaloneReconciler reconciles a Standalone object
type StandaloneReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=k8s.io.redis,resources=standalones,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=k8s.io.redis,resources=standalones/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=k8s.io.redis,resources=standalones/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Standalone object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.19.1/pkg/reconcile
func (r *StandaloneReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	l := log.FromContext(ctx)
	standalone := &k8sioredisv1.Standalone{}
	if err := r.Client.Get(ctx, req.NamespacedName, standalone); err != nil {
		l.Error(err, "unable to get standalone")
		return ctrl.Result{}, err
	}

	replicas := standalone.Spec.Replicas
	redisImage := standalone.Spec.Image

	dep := &v1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "standalone-redis",
			Namespace: standalone.Namespace,
		},
		Spec: v1.DeploymentSpec{
			Replicas: &replicas,
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:         "redis",
							Image:        redisImage,
							Env:          standalone.Spec.Env,
							VolumeMounts: standalone.Spec.VolumeMounts,
						},
					},
					Volumes:         []corev1.Volume{},
					SecurityContext: &corev1.PodSecurityContext{},
				},
			},
		},
		Status: v1.DeploymentStatus{},
	}
	if err := r.Client.Create(ctx, dep); err != nil {
		return ctrl.Result{}, err
	}
	l.Info(dep.Namespace)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *StandaloneReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&k8sioredisv1.Standalone{}).
		Named("standalone").
		Complete(r)
}
