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
	corev1 "k8s.io/api/core/v1"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	k8sioredisv1 "github.com/douzengrui001/redis_operator/api/v1"
)

// nolint:unused
// log is for logging in this package.
var standalonelog = logf.Log.WithName("standalone-resource")

// SetupStandaloneWebhookWithManager registers the webhook for Standalone in the manager.
func SetupStandaloneWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).For(&k8sioredisv1.Standalone{}).
		WithValidator(&StandaloneCustomValidator{}).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// NOTE: The 'path' attribute must follow a specific pattern and should not be modified directly here.
// Modifying the path for an invalid path can cause API server errors; failing to locate the webhook.
// +kubebuilder:webhook:path=/validate-k8s-io-redis-v1-standalone,mutating=false,failurePolicy=fail,sideEffects=None,groups=k8s.io.redis,resources=standalones,verbs=create;update,versions=v1,name=vstandalone-v1.kb.io,admissionReviewVersions=v1

// StandaloneCustomValidator struct is responsible for validating the Standalone resource
// when it is created, updated, or deleted.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as this struct is used only for temporary operations and does not need to be deeply copied.
type StandaloneCustomValidator struct {
	//TODO(user): Add more fields as needed for validation
}

var _ webhook.CustomValidator = &StandaloneCustomValidator{}

// ValidateCreate implements webhook.CustomValidator so a webhook will be registered for the type Standalone.
func (v *StandaloneCustomValidator) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	standalone, ok := obj.(*k8sioredisv1.Standalone)
	if !ok {
		return nil, fmt.Errorf("expected a Standalone object but got %T", obj)
	}
	standalonelog.Info("Validation for Standalone upon creation", "name", standalone.GetName())

	if standalone.GetNamespace() == "" {
		standalone.Namespace = "redis-operator"
	}

	if standalone.Spec.RestartPolicy == "" {
		standalone.Spec.RestartPolicy = corev1.RestartPolicyAlways
	}

	if standalone.Spec.SchedulerName == "" {
		standalone.Spec.SchedulerName = "default-scheduler"
	}

	if standalone.Spec.DNSPolicy == "" {
		standalone.Spec.DNSPolicy = corev1.DNSDefault
	}

	return nil, nil
}

// ValidateUpdate implements webhook.CustomValidator so a webhook will be registered for the type Standalone.
func (v *StandaloneCustomValidator) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) (admission.Warnings, error) {
	standalone, ok := newObj.(*k8sioredisv1.Standalone)
	if !ok {
		return nil, fmt.Errorf("expected a Standalone object for the newObj but got %T", newObj)
	}
	standalonelog.Info("Validation for Standalone upon update", "name", standalone.GetName())

	// TODO(user): fill in your validation logic upon object update.

	return nil, nil
}

// ValidateDelete implements webhook.CustomValidator so a webhook will be registered for the type Standalone.
func (v *StandaloneCustomValidator) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	standalone, ok := obj.(*k8sioredisv1.Standalone)
	if !ok {
		return nil, fmt.Errorf("expected a Standalone object but got %T", obj)
	}
	standalonelog.Info("Validation for Standalone upon deletion", "name", standalone.GetName())

	// TODO(user): fill in your validation logic upon object deletion.

	return nil, nil
}
