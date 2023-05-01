// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package db_cluster

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	ec2apitypes "github.com/aws-controllers-k8s/ec2-controller/apis/v1alpha1"
	kmsapitypes "github.com/aws-controllers-k8s/kms-controller/apis/v1alpha1"
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	acktypes "github.com/aws-controllers-k8s/runtime/pkg/types"

	svcapitypes "github.com/aws-controllers-k8s/rds-controller/apis/v1alpha1"
)

// +kubebuilder:rbac:groups=kms.services.k8s.aws,resources=keys,verbs=get;list
// +kubebuilder:rbac:groups=kms.services.k8s.aws,resources=keys/status,verbs=get;list

// +kubebuilder:rbac:groups=kms.services.k8s.aws,resources=keys,verbs=get;list
// +kubebuilder:rbac:groups=kms.services.k8s.aws,resources=keys/status,verbs=get;list

// +kubebuilder:rbac:groups=ec2.services.k8s.aws,resources=securitygroups,verbs=get;list
// +kubebuilder:rbac:groups=ec2.services.k8s.aws,resources=securitygroups/status,verbs=get;list

// ClearResolvedReferences removes any reference values that were made
// concrete in the spec. It returns a copy of the input AWSResource which
// contains the original *Ref values, but none of their respective concrete
// values.
func (rm *resourceManager) ClearResolvedReferences(res acktypes.AWSResource) acktypes.AWSResource {
	ko := rm.concreteResource(res).ko.DeepCopy()

	if ko.Spec.DBClusterParameterGroupRef != nil {
		ko.Spec.DBClusterParameterGroupName = nil
	}

	if ko.Spec.DBSubnetGroupRef != nil {
		ko.Spec.DBSubnetGroupName = nil
	}

	if ko.Spec.KMSKeyRef != nil {
		ko.Spec.KMSKeyID = nil
	}

	if ko.Spec.MasterUserSecretKMSKeyRef != nil {
		ko.Spec.MasterUserSecretKMSKeyID = nil
	}

	if len(ko.Spec.VPCSecurityGroupRefs) > 0 {
		ko.Spec.VPCSecurityGroupIDs = nil
	}

	return &resource{ko}
}

// ResolveReferences finds if there are any Reference field(s) present
// inside AWSResource passed in the parameter and attempts to resolve those
// reference field(s) into their respective target field(s). It returns a
// copy of the input AWSResource with resolved reference(s), a boolean which
// is set to true if the resource contains any references (regardless of if
// they are resolved successfully) and an error if the passed AWSResource's
// reference field(s) could not be resolved.
func (rm *resourceManager) ResolveReferences(
	ctx context.Context,
	apiReader client.Reader,
	res acktypes.AWSResource,
) (acktypes.AWSResource, bool, error) {
	namespace := res.MetaObject().GetNamespace()
	ko := rm.concreteResource(res).ko

	resourceHasReferences := false
	err := validateReferenceFields(ko)
	if fieldHasReferences, err := rm.resolveReferenceForDBClusterParameterGroupName(ctx, apiReader, namespace, ko); err != nil {
		return &resource{ko}, (resourceHasReferences || fieldHasReferences), err
	} else {
		resourceHasReferences = resourceHasReferences || fieldHasReferences
	}

	if fieldHasReferences, err := rm.resolveReferenceForDBSubnetGroupName(ctx, apiReader, namespace, ko); err != nil {
		return &resource{ko}, (resourceHasReferences || fieldHasReferences), err
	} else {
		resourceHasReferences = resourceHasReferences || fieldHasReferences
	}

	if fieldHasReferences, err := rm.resolveReferenceForKMSKeyID(ctx, apiReader, namespace, ko); err != nil {
		return &resource{ko}, (resourceHasReferences || fieldHasReferences), err
	} else {
		resourceHasReferences = resourceHasReferences || fieldHasReferences
	}

	if fieldHasReferences, err := rm.resolveReferenceForMasterUserSecretKMSKeyID(ctx, apiReader, namespace, ko); err != nil {
		return &resource{ko}, (resourceHasReferences || fieldHasReferences), err
	} else {
		resourceHasReferences = resourceHasReferences || fieldHasReferences
	}

	if fieldHasReferences, err := rm.resolveReferenceForVPCSecurityGroupIDs(ctx, apiReader, namespace, ko); err != nil {
		return &resource{ko}, (resourceHasReferences || fieldHasReferences), err
	} else {
		resourceHasReferences = resourceHasReferences || fieldHasReferences
	}

	return &resource{ko}, resourceHasReferences, err
}

// validateReferenceFields validates the reference field and corresponding
// identifier field.
func validateReferenceFields(ko *svcapitypes.DBCluster) error {

	if ko.Spec.DBClusterParameterGroupRef != nil && ko.Spec.DBClusterParameterGroupName != nil {
		return ackerr.ResourceReferenceAndIDNotSupportedFor("DBClusterParameterGroupName", "DBClusterParameterGroupRef")
	}

	if ko.Spec.DBSubnetGroupRef != nil && ko.Spec.DBSubnetGroupName != nil {
		return ackerr.ResourceReferenceAndIDNotSupportedFor("DBSubnetGroupName", "DBSubnetGroupRef")
	}

	if ko.Spec.KMSKeyRef != nil && ko.Spec.KMSKeyID != nil {
		return ackerr.ResourceReferenceAndIDNotSupportedFor("KMSKeyID", "KMSKeyRef")
	}

	if ko.Spec.MasterUserSecretKMSKeyRef != nil && ko.Spec.MasterUserSecretKMSKeyID != nil {
		return ackerr.ResourceReferenceAndIDNotSupportedFor("MasterUserSecretKMSKeyID", "MasterUserSecretKMSKeyRef")
	}

	if len(ko.Spec.VPCSecurityGroupRefs) > 0 && len(ko.Spec.VPCSecurityGroupIDs) > 0 {
		return ackerr.ResourceReferenceAndIDNotSupportedFor("VPCSecurityGroupIDs", "VPCSecurityGroupRefs")
	}
	return nil
}

// resolveReferenceForDBClusterParameterGroupName reads the resource referenced
// from DBClusterParameterGroupRef field and sets the DBClusterParameterGroupName
// from referenced resource. Returns a boolean indicating whether a reference
// contains references, or an error
func (rm *resourceManager) resolveReferenceForDBClusterParameterGroupName(
	ctx context.Context,
	apiReader client.Reader,
	namespace string,
	ko *svcapitypes.DBCluster,
) (hasReferences bool, err error) {
	if ko.Spec.DBClusterParameterGroupRef != nil && ko.Spec.DBClusterParameterGroupRef.From != nil {
		hasReferences = true
		arr := ko.Spec.DBClusterParameterGroupRef.From
		if arr.Name == nil || *arr.Name == "" {
			return hasReferences, fmt.Errorf("provided resource reference is nil or empty: DBClusterParameterGroupRef")
		}
		obj := &svcapitypes.DBClusterParameterGroup{}
		if err := getReferencedResourceState_DBClusterParameterGroup(ctx, apiReader, obj, *arr.Name, namespace); err != nil {
			return hasReferences, err
		}
		ko.Spec.DBClusterParameterGroupName = (*string)(obj.Spec.Name)
	}

	return hasReferences, nil
}

// getReferencedResourceState_DBClusterParameterGroup looks up whether a referenced resource
// exists and is in a ACK.ResourceSynced=True state. If the referenced resource does exist and is
// in a Synced state, returns nil, otherwise returns `ackerr.ResourceReferenceTerminalFor` or
// `ResourceReferenceNotSyncedFor` depending on if the resource is in a Terminal state.
func getReferencedResourceState_DBClusterParameterGroup(
	ctx context.Context,
	apiReader client.Reader,
	obj *svcapitypes.DBClusterParameterGroup,
	name string, // the Kubernetes name of the referenced resource
	namespace string, // the Kubernetes namespace of the referenced resource
) error {
	namespacedName := types.NamespacedName{
		Namespace: namespace,
		Name:      name,
	}
	err := apiReader.Get(ctx, namespacedName, obj)
	if err != nil {
		return err
	}
	var refResourceSynced, refResourceTerminal bool
	for _, cond := range obj.Status.Conditions {
		if cond.Type == ackv1alpha1.ConditionTypeResourceSynced &&
			cond.Status == corev1.ConditionTrue {
			refResourceSynced = true
		}
		if cond.Type == ackv1alpha1.ConditionTypeTerminal &&
			cond.Status == corev1.ConditionTrue {
			return ackerr.ResourceReferenceTerminalFor(
				"DBClusterParameterGroup",
				namespace, name)
		}
	}
	if refResourceTerminal {
		return ackerr.ResourceReferenceTerminalFor(
			"DBClusterParameterGroup",
			namespace, name)
	}
	if !refResourceSynced {
		return ackerr.ResourceReferenceNotSyncedFor(
			"DBClusterParameterGroup",
			namespace, name)
	}
	if obj.Spec.Name == nil {
		return ackerr.ResourceReferenceMissingTargetFieldFor(
			"DBClusterParameterGroup",
			namespace, name,
			"Spec.Name")
	}
	return nil
}

// resolveReferenceForDBSubnetGroupName reads the resource referenced
// from DBSubnetGroupRef field and sets the DBSubnetGroupName
// from referenced resource. Returns a boolean indicating whether a reference
// contains references, or an error
func (rm *resourceManager) resolveReferenceForDBSubnetGroupName(
	ctx context.Context,
	apiReader client.Reader,
	namespace string,
	ko *svcapitypes.DBCluster,
) (hasReferences bool, err error) {
	if ko.Spec.DBSubnetGroupRef != nil && ko.Spec.DBSubnetGroupRef.From != nil {
		hasReferences = true
		arr := ko.Spec.DBSubnetGroupRef.From
		if arr.Name == nil || *arr.Name == "" {
			return hasReferences, fmt.Errorf("provided resource reference is nil or empty: DBSubnetGroupRef")
		}
		obj := &svcapitypes.DBSubnetGroup{}
		if err := getReferencedResourceState_DBSubnetGroup(ctx, apiReader, obj, *arr.Name, namespace); err != nil {
			return hasReferences, err
		}
		ko.Spec.DBSubnetGroupName = (*string)(obj.Spec.Name)
	}

	return hasReferences, nil
}

// getReferencedResourceState_DBSubnetGroup looks up whether a referenced resource
// exists and is in a ACK.ResourceSynced=True state. If the referenced resource does exist and is
// in a Synced state, returns nil, otherwise returns `ackerr.ResourceReferenceTerminalFor` or
// `ResourceReferenceNotSyncedFor` depending on if the resource is in a Terminal state.
func getReferencedResourceState_DBSubnetGroup(
	ctx context.Context,
	apiReader client.Reader,
	obj *svcapitypes.DBSubnetGroup,
	name string, // the Kubernetes name of the referenced resource
	namespace string, // the Kubernetes namespace of the referenced resource
) error {
	namespacedName := types.NamespacedName{
		Namespace: namespace,
		Name:      name,
	}
	err := apiReader.Get(ctx, namespacedName, obj)
	if err != nil {
		return err
	}
	var refResourceSynced, refResourceTerminal bool
	for _, cond := range obj.Status.Conditions {
		if cond.Type == ackv1alpha1.ConditionTypeResourceSynced &&
			cond.Status == corev1.ConditionTrue {
			refResourceSynced = true
		}
		if cond.Type == ackv1alpha1.ConditionTypeTerminal &&
			cond.Status == corev1.ConditionTrue {
			return ackerr.ResourceReferenceTerminalFor(
				"DBSubnetGroup",
				namespace, name)
		}
	}
	if refResourceTerminal {
		return ackerr.ResourceReferenceTerminalFor(
			"DBSubnetGroup",
			namespace, name)
	}
	if !refResourceSynced {
		return ackerr.ResourceReferenceNotSyncedFor(
			"DBSubnetGroup",
			namespace, name)
	}
	if obj.Spec.Name == nil {
		return ackerr.ResourceReferenceMissingTargetFieldFor(
			"DBSubnetGroup",
			namespace, name,
			"Spec.Name")
	}
	return nil
}

// resolveReferenceForKMSKeyID reads the resource referenced
// from KMSKeyRef field and sets the KMSKeyID
// from referenced resource. Returns a boolean indicating whether a reference
// contains references, or an error
func (rm *resourceManager) resolveReferenceForKMSKeyID(
	ctx context.Context,
	apiReader client.Reader,
	namespace string,
	ko *svcapitypes.DBCluster,
) (hasReferences bool, err error) {
	if ko.Spec.KMSKeyRef != nil && ko.Spec.KMSKeyRef.From != nil {
		hasReferences = true
		arr := ko.Spec.KMSKeyRef.From
		if arr.Name == nil || *arr.Name == "" {
			return hasReferences, fmt.Errorf("provided resource reference is nil or empty: KMSKeyRef")
		}
		obj := &kmsapitypes.Key{}
		if err := getReferencedResourceState_Key(ctx, apiReader, obj, *arr.Name, namespace); err != nil {
			return hasReferences, err
		}
		ko.Spec.KMSKeyID = (*string)(obj.Status.ACKResourceMetadata.ARN)
	}

	return hasReferences, nil
}

// getReferencedResourceState_Key looks up whether a referenced resource
// exists and is in a ACK.ResourceSynced=True state. If the referenced resource does exist and is
// in a Synced state, returns nil, otherwise returns `ackerr.ResourceReferenceTerminalFor` or
// `ResourceReferenceNotSyncedFor` depending on if the resource is in a Terminal state.
func getReferencedResourceState_Key(
	ctx context.Context,
	apiReader client.Reader,
	obj *kmsapitypes.Key,
	name string, // the Kubernetes name of the referenced resource
	namespace string, // the Kubernetes namespace of the referenced resource
) error {
	namespacedName := types.NamespacedName{
		Namespace: namespace,
		Name:      name,
	}
	err := apiReader.Get(ctx, namespacedName, obj)
	if err != nil {
		return err
	}
	var refResourceSynced, refResourceTerminal bool
	for _, cond := range obj.Status.Conditions {
		if cond.Type == ackv1alpha1.ConditionTypeResourceSynced &&
			cond.Status == corev1.ConditionTrue {
			refResourceSynced = true
		}
		if cond.Type == ackv1alpha1.ConditionTypeTerminal &&
			cond.Status == corev1.ConditionTrue {
			return ackerr.ResourceReferenceTerminalFor(
				"Key",
				namespace, name)
		}
	}
	if refResourceTerminal {
		return ackerr.ResourceReferenceTerminalFor(
			"Key",
			namespace, name)
	}
	if !refResourceSynced {
		return ackerr.ResourceReferenceNotSyncedFor(
			"Key",
			namespace, name)
	}
	if obj.Status.ACKResourceMetadata == nil || obj.Status.ACKResourceMetadata.ARN == nil {
		return ackerr.ResourceReferenceMissingTargetFieldFor(
			"Key",
			namespace, name,
			"Status.ACKResourceMetadata.ARN")
	}
	return nil
}

// resolveReferenceForMasterUserSecretKMSKeyID reads the resource referenced
// from MasterUserSecretKMSKeyRef field and sets the MasterUserSecretKMSKeyID
// from referenced resource. Returns a boolean indicating whether a reference
// contains references, or an error
func (rm *resourceManager) resolveReferenceForMasterUserSecretKMSKeyID(
	ctx context.Context,
	apiReader client.Reader,
	namespace string,
	ko *svcapitypes.DBCluster,
) (hasReferences bool, err error) {
	if ko.Spec.MasterUserSecretKMSKeyRef != nil && ko.Spec.MasterUserSecretKMSKeyRef.From != nil {
		hasReferences = true
		arr := ko.Spec.MasterUserSecretKMSKeyRef.From
		if arr.Name == nil || *arr.Name == "" {
			return hasReferences, fmt.Errorf("provided resource reference is nil or empty: MasterUserSecretKMSKeyRef")
		}
		obj := &kmsapitypes.Key{}
		if err := getReferencedResourceState_Key(ctx, apiReader, obj, *arr.Name, namespace); err != nil {
			return hasReferences, err
		}
		ko.Spec.MasterUserSecretKMSKeyID = (*string)(obj.Status.ACKResourceMetadata.ARN)
	}

	return hasReferences, nil
}

// resolveReferenceForVPCSecurityGroupIDs reads the resource referenced
// from VPCSecurityGroupRefs field and sets the VPCSecurityGroupIDs
// from referenced resource. Returns a boolean indicating whether a reference
// contains references, or an error
func (rm *resourceManager) resolveReferenceForVPCSecurityGroupIDs(
	ctx context.Context,
	apiReader client.Reader,
	namespace string,
	ko *svcapitypes.DBCluster,
) (hasReferences bool, err error) {
	for _, f0iter := range ko.Spec.VPCSecurityGroupRefs {
		if f0iter != nil && f0iter.From != nil {
			hasReferences = true
			arr := f0iter.From
			if arr.Name == nil || *arr.Name == "" {
				return hasReferences, fmt.Errorf("provided resource reference is nil or empty: VPCSecurityGroupRefs")
			}
			obj := &ec2apitypes.SecurityGroup{}
			if err := getReferencedResourceState_SecurityGroup(ctx, apiReader, obj, *arr.Name, namespace); err != nil {
				return hasReferences, err
			}
			if ko.Spec.VPCSecurityGroupIDs == nil {
				ko.Spec.VPCSecurityGroupIDs = make([]*string, 0, 1)
			}
			ko.Spec.VPCSecurityGroupIDs = append(ko.Spec.VPCSecurityGroupIDs, (*string)(obj.Status.ID))
		}
	}

	return hasReferences, nil
}

// getReferencedResourceState_SecurityGroup looks up whether a referenced resource
// exists and is in a ACK.ResourceSynced=True state. If the referenced resource does exist and is
// in a Synced state, returns nil, otherwise returns `ackerr.ResourceReferenceTerminalFor` or
// `ResourceReferenceNotSyncedFor` depending on if the resource is in a Terminal state.
func getReferencedResourceState_SecurityGroup(
	ctx context.Context,
	apiReader client.Reader,
	obj *ec2apitypes.SecurityGroup,
	name string, // the Kubernetes name of the referenced resource
	namespace string, // the Kubernetes namespace of the referenced resource
) error {
	namespacedName := types.NamespacedName{
		Namespace: namespace,
		Name:      name,
	}
	err := apiReader.Get(ctx, namespacedName, obj)
	if err != nil {
		return err
	}
	var refResourceSynced, refResourceTerminal bool
	for _, cond := range obj.Status.Conditions {
		if cond.Type == ackv1alpha1.ConditionTypeResourceSynced &&
			cond.Status == corev1.ConditionTrue {
			refResourceSynced = true
		}
		if cond.Type == ackv1alpha1.ConditionTypeTerminal &&
			cond.Status == corev1.ConditionTrue {
			return ackerr.ResourceReferenceTerminalFor(
				"SecurityGroup",
				namespace, name)
		}
	}
	if refResourceTerminal {
		return ackerr.ResourceReferenceTerminalFor(
			"SecurityGroup",
			namespace, name)
	}
	if !refResourceSynced {
		return ackerr.ResourceReferenceNotSyncedFor(
			"SecurityGroup",
			namespace, name)
	}
	if obj.Status.ID == nil {
		return ackerr.ResourceReferenceMissingTargetFieldFor(
			"SecurityGroup",
			namespace, name,
			"Status.ID")
	}
	return nil
}
