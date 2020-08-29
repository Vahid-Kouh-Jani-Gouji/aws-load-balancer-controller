package nlb

import "sigs.k8s.io/aws-alb-ingress-controller/pkg/model/core"

type resourceVisitor struct {
	resources []core.Resource
}

var _ core.ResourceVisitor = &resourceVisitor{}

func (r *resourceVisitor) Visit(res core.Resource) error {
	r.resources = append(r.resources, res)
	return nil
}
