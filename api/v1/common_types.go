package v1

import corev1 "k8s.io/api/core/v1"

type MetricsSettings struct {
	Enabled       bool            `json:"enabled"`
	Image         string          `json:"image"`
	PullPolicy    string          `json:"pullPolicy"`
	Port          int32           `json:"port"`
	Env           []corev1.EnvVar `json:"env,omitempty"`
	Command       []string        `json:"command,omitempty"`
	OtherSettings `json:",inline"`
}
type AuthSettings struct {
	TLSCRTPath string                      `json:"crt,omitempty"`
	TLSKeyPath string                      `json:"key,omitempty"`
	TLSCAPath  string                      `json:"ca,omitempty"`
	SecretRef  corev1.LocalObjectReference `json:"secretRef,omitempty"`
}
type Volume struct {
	Name         string              `json:"name"`
	VolumeSource corev1.VolumeSource `json:"volumeSource,omitempty"`
}

type OtherSettings struct {
	DNSPolicy                 corev1.DNSPolicy                  `json:"dnsPolicy,omitempty" `
	SchedulerName             string                            `json:"schedulerName,omitempty"`
	RestartPolicy             corev1.RestartPolicy              `json:"restartPolicy,omitempty"`
	Tolerations               []corev1.Toleration               `json:"tolerations,omitempty"`
	NodeSelector              *corev1.NodeSelector              `json:"nodeSelector,omitempty"`
	Affinity                  *corev1.Affinity                  `json:"affinity,omitempty"`
	SecurityContext           *corev1.PodSecurityContext        `json:"securityContext,omitempty"`
	ReadinessProbe            *corev1.Probe                     `json:"readinessProbe,omitempty"`
	LivenessProbe             *corev1.Probe                     `json:"livenessProbe,omitempty" `
	Resources                 *corev1.ResourceClaim             `json:"resources"`
	TopologySpreadConstraints []corev1.TopologySpreadConstraint `json:"topologySpreadConstraints,omitempty"`
}
