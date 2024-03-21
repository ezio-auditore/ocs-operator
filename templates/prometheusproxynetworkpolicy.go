package templates

import (
	corev1 "k8s.io/api/core/v1"
	netv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

var (
	prometheusProxyPort     = intstr.FromInt(KubeRBACProxyPortNumber)
	prometheusProxyProtocol = corev1.ProtocolTCP
)

var PrometheusProxyNetworkPolicyTemplate = netv1.NetworkPolicy{
	Spec: netv1.NetworkPolicySpec{
		Ingress: []netv1.NetworkPolicyIngressRule{
			{
				Ports: []netv1.NetworkPolicyPort{
					{
						Port:     &prometheusProxyPort,
						Protocol: &prometheusProxyProtocol,
					},
				},
			},
		},
		PolicyTypes: []netv1.PolicyType{
			netv1.PolicyTypeIngress,
		},
		PodSelector: metav1.LabelSelector{
			MatchExpressions: []metav1.LabelSelectorRequirement{
				{
					Key:      "prometheus",
					Operator: metav1.LabelSelectorOpIn,
					Values: []string{
						"odf-prometheus",
					},
				},
			},
		},
	},
}
