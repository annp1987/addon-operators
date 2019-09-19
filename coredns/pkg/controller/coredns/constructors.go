package coredns

import (
	addonsv1alpha1 "sigs.k8s.io/addon-operators/coredns-operator/pkg/apis/addons/v1alpha1"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const DefaultCorefile = `
.:53 {
	errors
	health
	kubernetes cluster.local in-addr.arpa ip6.arpa {
	pods insecure
	upstream
	fallthrough in-addr.arpa ip6.arpa
	}
	prometheus :9153
	forward . /etc/resolv.conf
	cache 30
	loop
	reload
	loadbalance
}
`

// newConfigMapForCR returns a ConfigMap with the same name/namespace as the cr
func newConfigMapForCR(cr *addonsv1alpha1.CoreDNS) *corev1.ConfigMap {
	return &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name,
			Namespace: cr.Namespace,
		},
		Data: map[string]string{
			"Corefile": cr.Spec.Corefile,
		},
	}
}
