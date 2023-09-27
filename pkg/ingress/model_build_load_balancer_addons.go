package ingress

import (
	"context"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/util/sets"
	"sigs.k8s.io/aws-load-balancer-controller/pkg/annotations"
	"sigs.k8s.io/aws-load-balancer-controller/pkg/model/core"
	gamodel "sigs.k8s.io/aws-load-balancer-controller/pkg/model/globalaccelerator"
	shieldmodel "sigs.k8s.io/aws-load-balancer-controller/pkg/model/shield"
	wafregionalmodel "sigs.k8s.io/aws-load-balancer-controller/pkg/model/wafregional"
	wafv2model "sigs.k8s.io/aws-load-balancer-controller/pkg/model/wafv2"
)

func (t *defaultModelBuildTask) buildLoadBalancerAddOns(ctx context.Context, lbARN core.StringToken) error {

	if _, err := t.buildWAFv2WebACLAssociation(ctx, lbARN); err != nil {
		return err
	}
	if _, err := t.buildWAFRegionalWebACLAssociation(ctx, lbARN); err != nil {
		return err
	}
	if _, err := t.buildShieldProtection(ctx, lbARN); err != nil {
		return err
	}
	if _, err := t.buildGAEndpoint(ctx, lbARN); err != nil {
		return err
	}

	return nil
}

func (t *defaultModelBuildTask) buildWAFv2WebACLAssociation(_ context.Context, lbARN core.StringToken) (*wafv2model.WebACLAssociation, error) {
	explicitWebACLARNs := sets.NewString()
	for _, member := range t.ingGroup.Members {
		rawWebACLARN := ""
		if exists := t.annotationParser.ParseStringAnnotation(annotations.IngressSuffixWAFv2ACLARN, &rawWebACLARN, member.Ing.Annotations); exists {
			explicitWebACLARNs.Insert(rawWebACLARN)
		}
	}
	if len(explicitWebACLARNs) == 0 {
		return nil, nil
	}
	if len(explicitWebACLARNs) > 1 {
		return nil, errors.Errorf("conflicting WAFv2 WebACL ARNs: %v", explicitWebACLARNs.List())
	}
	webACLARN, _ := explicitWebACLARNs.PopAny()
	if webACLARN != "" {
		association := wafv2model.NewWebACLAssociation(t.stack, resourceIDLoadBalancer, wafv2model.WebACLAssociationSpec{
			WebACLARN:   webACLARN,
			ResourceARN: lbARN,
		})
		return association, nil
	}
	return nil, nil
}

func (t *defaultModelBuildTask) buildWAFRegionalWebACLAssociation(_ context.Context, lbARN core.StringToken) (*wafregionalmodel.WebACLAssociation, error) {
	explicitWebACLIDs := sets.NewString()
	for _, member := range t.ingGroup.Members {
		rawWebACLARN := ""
		if exists := t.annotationParser.ParseStringAnnotation(annotations.IngressSuffixWAFACLID, &rawWebACLARN, member.Ing.Annotations); exists {
			explicitWebACLIDs.Insert(rawWebACLARN)
		} else if exists := t.annotationParser.ParseStringAnnotation(annotations.IngressSuffixWebACLID, &rawWebACLARN, member.Ing.Annotations); exists {
			explicitWebACLIDs.Insert(rawWebACLARN)
		}
	}
	if len(explicitWebACLIDs) == 0 {
		return nil, nil
	}
	if len(explicitWebACLIDs) > 1 {
		return nil, errors.Errorf("conflicting WAFRegional WebACL IDs: %v", explicitWebACLIDs.List())
	}
	webACLID, _ := explicitWebACLIDs.PopAny()
	if webACLID != "" {
		association := wafregionalmodel.NewWebACLAssociation(t.stack, resourceIDLoadBalancer, wafregionalmodel.WebACLAssociationSpec{
			WebACLID:    webACLID,
			ResourceARN: lbARN,
		})
		return association, nil
	}
	return nil, nil
}

func (t *defaultModelBuildTask) buildShieldProtection(_ context.Context, lbARN core.StringToken) (*shieldmodel.Protection, error) {
	explicitEnableProtections := make(map[bool]struct{})
	for _, member := range t.ingGroup.Members {
		rawEnableProtection := false
		exists, err := t.annotationParser.ParseBoolAnnotation(annotations.IngressSuffixShieldAdvancedProtection, &rawEnableProtection, member.Ing.Annotations)
		if err != nil {
			return nil, err
		}
		if exists {
			explicitEnableProtections[rawEnableProtection] = struct{}{}
		}
	}
	if len(explicitEnableProtections) == 0 {
		return nil, nil
	}
	if len(explicitEnableProtections) > 1 {
		return nil, errors.New("conflicting enable shield advanced protection")
	}
	if _, enableProtection := explicitEnableProtections[true]; enableProtection {
		protection := shieldmodel.NewProtection(t.stack, resourceIDLoadBalancer, shieldmodel.ProtectionSpec{
			ResourceARN: lbARN,
		})
		return protection, nil
	}
	return nil, nil
}

func (t *defaultModelBuildTask) buildGAEndpoint(_ context.Context, lbARN core.StringToken) (*gamodel.Endpoint, error) {
	explicitEPGARNs := sets.NewString()
	epCreateByARN := make(map[string]string)
	rawEPGARN := ""

	for _, member := range t.ingGroup.Members {
		// Unfortunately we can't support deletion of an endpoint just by removing
		// the `ga-epg-arn` annotation, because cleaning up afterwards would require
		// us to scan all accelerators and endpointgroups for the desired endpoint.
		// To still enable a way to delete the endpoint, users can set `ga-epg-create=false`
		epCreate := "false"
		if exists := t.annotationParser.ParseStringAnnotation(annotations.IngressSuffixGAEndpointGroup, &rawEPGARN, member.Ing.Annotations); exists {
			explicitEPGARNs.Insert(rawEPGARN)
			t.annotationParser.ParseStringAnnotation(annotations.IngressSuffixGAEndpointCreate, &epCreate, member.Ing.Annotations)
			epCreateByARN[rawEPGARN] = epCreate
		}
	}
	if len(explicitEPGARNs) > 1 {
		return nil, errors.Errorf("conflicting Global Accelerator EndpointGroup ARNs: %v", explicitEPGARNs.List())
	}
	if len(explicitEPGARNs) == 1 {
		epgARN, _ := explicitEPGARNs.PopAny()
		if epgARN != "" {
			create := false
			if epCreateByARN[epgARN] == "true" {
				create = true
			}
			endpoint := gamodel.NewEndpoint(t.stack, resourceIDLoadBalancer, gamodel.EndpointSpec{
				EndpointGroupARN: epgARN,
				ResourceARN:      lbARN,
				Create:           create,
			})
			return endpoint, nil
		}
	}

	return nil, nil
}
