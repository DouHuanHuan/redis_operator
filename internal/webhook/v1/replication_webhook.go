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

package v1

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"

	k8sioredisv1 "github.com/douzengrui001/redis_operator/api/v1"
)

// nolint:unused
// log is for logging in this package.
var replicationlog = logf.Log.WithName("replication-resource")

// SetupReplicationWebhookWithManager registers the webhook for Replication in the manager.
func SetupReplicationWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).For(&k8sioredisv1.Replication{}).
		WithDefaulter(&ReplicationCustomDefaulter{}).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:path=/mutate-k8s-io-redis-v1-replication,mutating=true,failurePolicy=fail,sideEffects=None,groups=k8s.io.redis,resources=replications,verbs=create;update,versions=v1,name=mreplication-v1.kb.io,admissionReviewVersions=v1

// ReplicationCustomDefaulter struct is responsible for setting default values on the custom resource of the
// Kind Replication when those are created or updated.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as it is used only for temporary operations and does not need to be deeply copied.
type ReplicationCustomDefaulter struct {
	// TODO(user): Add more fields as needed for defaulting
}

var _ webhook.CustomDefaulter = &ReplicationCustomDefaulter{}

// Default implements webhook.CustomDefaulter so a webhook will be registered for the Kind Replication.
func (d *ReplicationCustomDefaulter) Default(ctx context.Context, obj runtime.Object) error {
	replication, ok := obj.(*k8sioredisv1.Replication)

	if !ok {
		return fmt.Errorf("expected an Replication object but got %T", obj)
	}
	replicationlog.Info("Defaulting for Replication", "name", replication.GetName())

	// TODO(user): fill in your defaulting logic.

	return nil
}
