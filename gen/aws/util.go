//go:build codegen

package aws

import (
	"fmt"
	"sort"
	"strings"

	"github.com/aws/aws-sdk-go/private/model/api"

	"github.com/apex/log"
)

// ListFunctionName generates a name for the list function of the resource type.
func (r ResourceType) ListFunctionName() string {
	split := strings.Split(strings.TrimPrefix(r.Name, "aws_"), "_")
	capitalize := strings.Title(strings.Join(split, " "))

	return "List" + strings.Join(strings.Split(capitalize, " "), "")
}

func findResourceID(rType string, resourceIDs map[string]string, outputField *api.ShapeRef) (string, error) {
	resourceID, ok := ManualMatchedResourceID[rType]
	if ok {
		return resourceID, nil
	}

	resourceID, ok = resourceIDs[rType]
	if !ok {
		return "", fmt.Errorf("no resource ID found")
	}

	if resourceID == "NAME_PLACEHOLDER" {
		resourceIDCandidates := GetResourceIDNameCandidates(outputField)
		if len(resourceIDCandidates) > 1 {
			return "", fmt.Errorf("found multiple name field ID candidates as resource ID for NAME_PLACEHOLDER")

		}

		if len(resourceIDCandidates) == 0 {
			return "", fmt.Errorf("found no name field candidates as resource ID for NAME_PLACEHOLDER")
		}

		resourceID = resourceIDCandidates[0]
	}

	return resourceID, nil
}

func GetResourceIDNameCandidates(v *api.ShapeRef) []string {
	var result []string

	for k, _ := range v.Shape.MemberRef.Shape.MemberRefs {
		if k == "Name" {
			return []string{k}
		}

		if strings.Contains(strings.ToLower(k), "name") {
			result = append(result, k)
		}
	}

	return result
}

// GetOutputFieldCandidates gets the output field that contains a list of resources of the given resource type
// (e.g., field name LogGroups of type []LogGroup in DescribeLogGroupsOutput)
//
// Note: if there is a manual match entry, this will be returned.
func GetOutputFieldCandidates(resourceType string, op ListOperation, shapeType string) []string {
	_, ok := ManualMatchedOutputFields[resourceType]
	if ok {
		return []string{ManualMatchedOutputFields[resourceType]}
	}

	var outputFieldCandidates []string

	for fieldName, v := range op.OutputRef.Shape.MemberRefs {
		if v.Shape.Type == "list" {
			if v.Shape.MemberRef.Shape.Type == shapeType {
				outputFieldCandidates = append(outputFieldCandidates, fieldName)
			}
		}
	}

	return outputFieldCandidates
}

func findOutputField(rType string, listOpCandidates []ListOperation, shapeType string) (string, ListOperation, error) {
	var listOpCandidatesWithFoundOutputField []string
	var outputFieldName string
	var op ListOperation

	for _, listOpCandidate := range listOpCandidates {
		outputFieldCandidates := GetOutputFieldCandidates(rType, listOpCandidate, shapeType)
		if len(outputFieldCandidates) == 0 {
			continue
		}

		if len(outputFieldCandidates) > 1 {
			log.WithFields(log.Fields{
				"resource":   rType,
				"operation":  listOpCandidate.ExportedName,
				"candidates": outputFieldCandidates,
			}).Warnf("multiple output field candidates")
			continue
		}

		listOpCandidatesWithFoundOutputField = append(listOpCandidatesWithFoundOutputField, listOpCandidate.ExportedName)
		op = listOpCandidate
		outputFieldName = outputFieldCandidates[0]
		op.OutputListName = outputFieldName
	}

	if len(listOpCandidatesWithFoundOutputField) == 0 {
		return "", op, fmt.Errorf("no list operation candidate with struct found")
	}

	if len(listOpCandidatesWithFoundOutputField) > 1 {
		return "", op, fmt.Errorf("multiple list operation candidates found: %s", listOpCandidatesWithFoundOutputField)
	}

	return outputFieldName, op, nil
}

func Operations(apis api.APIs, prefixes []string) []string {
	var result []string

	for _, a := range apis {
		for _, v := range a.Operations {
			for _, prefix := range prefixes {
				if strings.HasPrefix(v.Name, prefix) && !strings.Contains(v.Name, "Tags") {
					log.Debugf("%s", v.Name)
					result = append(result, v.Name)
				}
			}
		}
	}
	return result
}

type ListOperationCandidates struct {
	List      *api.Operation
	Get       *api.Operation
	Describes *api.Operation
}

// FindListOperationCandidates returns all list operation candidates for a resource type.
// A list operation is a candidate, if
//  * it's name starts with Get, List or Describe
//  * the operation belongs to the same service as the resource type
//  * the name of the operation is a plural of the resource type name
// Note: If there is a manual match entry, this will be returned.
func FindListOperationCandidates(resourceType, service string, apis api.APIs) []ListOperation {
	manualMatchedOp, ok := ManualMatchedListOps[resourceType]
	if ok {
		for _, op := range operationsOfService(apis, service, "") {
			if op.ExportedName == manualMatchedOp {
				return []ListOperation{{Operation: *op}}
			}
		}
	}

	var result []ListOperation

	prefixes := []string{"Describe", "Get", "List"}
	var ops []string

	for _, prefix := range prefixes {
		operations := operationsOfService(apis, service, prefix)

		for _, op := range operations {
			ops = append(ops, op.ExportedName)
		}
		sort.Strings(ops)

		matchingOp := exactMatch(resourceType, operations, prefix)
		if matchingOp != nil {
			result = append(result, ListOperation{Operation: *matchingOp})
		}
	}

	if len(result) == 0 {
		log.Debugf("list operations: %s", ops)
	}

	return result
}

func exactMatch(rType string, operations []*api.Operation, opPrefix string) *api.Operation {
	plurals := pluralizeListFunctionCandidateNames(rType)

	for _, plural := range plurals {
		for _, op := range operations {
			opWithoutPrefix := strings.ToLower(strings.TrimPrefix(op.ExportedName, opPrefix))

			if plural == opWithoutPrefix {
				return op
			}
		}
	}

	return nil
}

func pluralizeListFunctionCandidateNames(rType string) []string {
	rTypeWithoutPrefix := strings.TrimPrefix(rType, "aws_")
	tSplit := strings.Split(rTypeWithoutPrefix, "_")

	var tCombined []string
	for i := 0; i < len(tSplit); i++ {
		tCombined = append(tCombined, strings.Join(tSplit[i:], ""))
	}

	var result []string
	for _, c := range tCombined {
		result = append(result, []string{c + "s", c + "es"}...)
		if strings.HasSuffix(c, "y") {
			result = append(result, strings.TrimSuffix(c, "y")+"ies")
		}
	}

	return result
}

// operationsOfService returns the operations with a given prefix that belong to a service.
func operationsOfService(apis api.APIs, service, opPrefix string) []*api.Operation {
	var result []*api.Operation

	for _, api := range apis {
		if service != api.PackageName() {
			continue
		}

		for _, op := range api.Operations {
			if strings.HasPrefix(op.ExportedName, opPrefix) {
				result = append(result, op)
			}
		}
	}

	return result
}
