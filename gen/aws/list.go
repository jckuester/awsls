// +build codegen

package aws

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"

	"github.com/apex/log"
	"github.com/aws/aws-sdk-go-v2/private/model/api"
	"github.com/jckuester/awsls/gen/util"
)

type GeneratedResourceInfo struct {
	Type         string
	Tags         bool
	CreationTime bool
	Owner        bool
}

func GenerateListFunctions(outputPath string, resourceServices map[string]string, resourceIDs map[string]string,
	resourceTypesWithTags []string, apis api.APIs) (map[string]string, map[string][]GeneratedResourceInfo) {
	listFunctionNames := map[string]string{}
	genResourceInfo := map[string][]GeneratedResourceInfo{}

	for _, service := range ServicesCoveredByTerraform(resourceServices) {
		fmt.Printf("\nservice: %s\n---\n", service)

		var genResourceInfoPerService []GeneratedResourceInfo

		for rType, rService := range resourceServices {
			if rService != service {
				continue
			}

			_, ok := ExcludedResourceTypes[rType]
			if ok {
				log.WithField("resource", rType).Info("exclude")
				continue
			}

			service, ok := resourceServices[rType]
			if !ok {
				log.WithField("resource", rType).Warnf("service not found")
				continue
			}

			var op Operation
			var outputFieldName string

			listOpCandidates := GetListOperationCandidates(rType, service, apis)
			if len(listOpCandidates) == 0 {
				log.WithField("resource", rType).Warnf("no list operation candidate found")
				continue
			}

			var listOpCandidateWithFoundOutputField []string
			for _, listOpCandidate := range listOpCandidates {
				outputFieldCandidates := GetOutputFieldCandidates(rType, listOpCandidate)
				if len(outputFieldCandidates) == 0 {
					continue
				}

				if len(outputFieldCandidates) > 1 {
					log.WithFields(log.Fields{
						"resource":   rType,
						"operation":  listOpCandidate.ExportedName,
						"candidates": outputFieldCandidates,
					}).Infof("multiple output field candidates")
					continue
				}

				listOpCandidateWithFoundOutputField = append(listOpCandidateWithFoundOutputField, listOpCandidate.ExportedName)
				op = listOpCandidate
				outputFieldName = outputFieldCandidates[0]
				op.OutputListName = outputFieldName
			}

			if len(listOpCandidateWithFoundOutputField) == 0 {
				log.WithField("resource", rType).Warnf("no list operation candidate with struct found")
				continue
			}

			if len(listOpCandidateWithFoundOutputField) > 1 {
				log.WithFields(log.Fields{
					"resource":   rType,
					"candidates": listOpCandidateWithFoundOutputField}).Warnf("multiple list operation candidates found")
				continue
			}

			if len(op.InputRef.Shape.Required) > 0 {
				log.WithField("resource", rType).Warnf("required input fields: %s", op.InputRef.Shape.Required)
				continue
			}

			outputField := op.OutputRef.Shape.MemberRefs[outputFieldName]

			resourceID, ok := ManualMatchedResourceID[rType]
			if !ok {
				resourceID, ok = resourceIDs[rType]
				if !ok {
					log.WithField("resource", rType).Warn("no ID found")
					continue
				}

				if resourceID == "NAME_PLACEHOLDER" {
					resourceIDCandidates := GetResourceIDNameCandidates(outputField)
					if len(resourceIDCandidates) > 1 {
						log.WithFields(log.Fields{
							"resource":   rType,
							"candidates": resourceIDCandidates,
						}).Warnf("found multiple name field ID candidates as resource ID for NAME_PLACEHOLDER")

						continue
					}

					if len(resourceIDCandidates) == 0 {
						log.WithFields(log.Fields{
							"resource": rType,
						}).Warnf("found no name field candidates as resource ID for NAME_PLACEHOLDER")

						continue
					}

					resourceID = resourceIDCandidates[0]
				}
			}

			op.TerraformType = rType
			op.ResourceID = resourceID
			op.OpName = TypeToOpName(rType)

			listFunctionNames[rType] = TypeToOpName(rType)

			genInfo := GeneratedResourceInfo{
				Type: rType,
			}

			op.Inputs = Inputs[rType]

			op.GetTagsGoCode = GetTagsGoCode(outputField)

			/*
				if op.GetTagsGoCode != "" {
					genInfo.Tags = true
				}
			*/

			for _, rWithTags := range resourceTypesWithTags {
				if rWithTags == rType {
					genInfo.Tags = true
				}
			}

			getTagCode, imports := GetCreationTimeGoCode(outputField)

			op.GetCreationTimeGoCode = getTagCode
			op.Imports = imports

			if op.GetCreationTimeGoCode != "" {
				genInfo.CreationTime = true
			}

			op.GetOwnerGoCode = GetOwnerGoCode(outputField)

			if op.GetOwnerGoCode != "" {
				genInfo.Owner = true
			}

			for k, _ := range op.InputRef.Shape.MemberRefs {
				if strings.Contains(strings.ToLower(k), "owner") {
					log.Infof("input; found owner field for %s: %s", rType, k)
				}
			}

			genResourceInfoPerService = append(genResourceInfoPerService, genInfo)

			err := writeListFunction(outputPath, &op, rType)
			if err != nil {
				log.Fatal(err.Error())
			}
		}

		if len(genResourceInfoPerService) > 0 {
			sort.Slice(genResourceInfoPerService, func(i, j int) bool {
				return genResourceInfoPerService[i].Type < genResourceInfoPerService[j].Type
			})

			genResourceInfo[service] = genResourceInfoPerService
		}
	}

	return listFunctionNames, genResourceInfo
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

// GetOutputFieldCandidates gets the output field that contains a list of resources the given resource type
// (e.g., field name LogGroups of type []LogGroup in output DescribeLogGroupsOutput)
//
// Note: if there is a manual match entry, this will be returned.
func GetOutputFieldCandidates(resourceType string, op Operation) []string {
	_, ok := ManualMatchedOutputFields[resourceType]
	if ok {
		return []string{ManualMatchedOutputFields[resourceType]}
	}

	var outputFieldCandidates []string

	for fieldName, v := range op.OutputRef.Shape.MemberRefs {
		if v.Shape.Type == "list" {
			if v.Shape.MemberRef.Shape.Type == "structure" {
				outputFieldCandidates = append(outputFieldCandidates, fieldName)
			}
		}
	}

	return outputFieldCandidates
}

func GetTagsGoCode(outputField *api.ShapeRef) string {
	for k, v := range outputField.Shape.MemberRef.Shape.MemberRefs {
		if k == "Tags" {
			if v.Shape.Type == "list" {
				return `tags := map[string]string{}
						for _, t := range r.Tags {
							tags[*t.Key] = *t.Value
						}`
			}

			if v.Shape.Type == "map" {
				return `tags := map[string]string{}
						for k, v := range r.Tags {
							tags[k] = v
						}`
			}
		}

		if strings.Contains(k, "Tag") {
			log.Infof("tags: %s %s", k, v.Shape.Type)
		}
	}

	return ""
}

func GetCreationTimeGoCode(outputField *api.ShapeRef) (string, []string) {
	creationTimeFieldNames := []string{
		"LaunchTime",
		"CreateTime",
		"CreateDate",
		"CreatedTime",
		"CreationDate",
		"CreationTime",
		"CreationTimestamp",
		"StartTime",
		"InstanceCreateTime",
	}

	for k, v := range outputField.Shape.MemberRef.Shape.MemberRefs {
		for _, name := range creationTimeFieldNames {
			if k == name {
				if v.Shape.Type == "string" {
					return `t, err := time.Parse("2006-01-02T15:04:05.000Z0700", *r.` + k + `)
							if err != nil {
								return nil, err
							}`, []string{"time"}
				}

				if v.Shape.Type == "timestamp" {
					return `t := ` + fmt.Sprintf("*r.%s", k), []string{}
				}

				if v.Shape.Type == "long" {
					return fmt.Sprintf("t := time.Unix(0, *r.%s * 1000000).UTC()", k), []string{"time"}
				}

				log.Warnf("uncovered creation time type: %s", v.Shape.Type)
			}
		}
	}

	return "", []string{}
}

func GetOwnerGoCode(outputField *api.ShapeRef) string {
	for k, _ := range outputField.Shape.MemberRef.Shape.MemberRefs {
		if k == "OwnerId" {
			return `if *r.OwnerId != client.AccountID {
						continue
					}`
		}
		if strings.Contains(strings.ToLower(k), "owner") {
			log.Infof("output; found owner field: %s", k)
		}
	}

	return ""
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

// GetListOperationCandidates gets the possible list operation
//
// Note:
//  * The list operation can be a Get, List or Describe function
//  * If there is a manual match entry, this will be returned.
func GetListOperationCandidates(resourceType, service string, apis api.APIs) []Operation {
	manualMatchedOp, ok := ManualMatchedListOps[resourceType]
	if ok {
		for _, op := range operationsOfService(apis, service, "") {
			if op.ExportedName == manualMatchedOp {
				return []Operation{{Operation: *op}}
			}
		}
	}

	var result []Operation

	prefixes := []string{"Describe", "Get", "List"}

	for _, prefix := range prefixes {
		operations := operationsOfService(apis, service, prefix)

		matchingOp := exactMatch(resourceType, operations, prefix)
		if matchingOp != nil {
			result = append(result, Operation{Operation: *matchingOp})
		}
	}

	return result
}

func exactMatch(terraformType string, operations []*api.Operation, opPrefix string) *api.Operation {
	plurals := pluralizeListFunctionCandidateNames(terraformType)

	for _, t := range plurals {
		for _, op := range operations {
			opNoPrefix := strings.ToLower(strings.TrimPrefix(op.ExportedName, opPrefix))

			if t == opNoPrefix {
				return op
			}
		}
	}

	return nil
}

func pluralizeListFunctionCandidateNames(terraformType string) []string {
	tNoPrefix := strings.TrimPrefix(terraformType, "aws_")
	tSplit := strings.Split(tNoPrefix, "_")

	var tCombined []string
	for i := 0; i < len(tSplit); i++ {
		tCombined = append(tCombined, strings.Join(tSplit[i:], ""))
	}

	var plural []string
	for _, c := range tCombined {
		plural = append(plural, []string{c + "s", c + "es"}...)
		if strings.HasSuffix(c, "y") {
			plural = append(plural, strings.TrimSuffix(c, "y")+"ies")
		}
	}
	return plural
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

// TypeToOpName generates a name for the list function based on the resource type.
func TypeToOpName(terraformType string) string {
	split := strings.Split(strings.TrimPrefix(terraformType, "aws_"), "_")
	capitalized := strings.Title(strings.Join(split, " "))

	return strings.Join(strings.Split(capitalized, " "), "")
}

func writeListFunction(outputPath string, op *Operation, terraformType string) error {
	err := util.WriteGoFile(
		filepath.Join(outputPath, terraformType+".go"),
		util.CodeLayout,
		"",
		"aws",
		op.GoCode(),
	)

	if err != nil {
		return fmt.Errorf("failed to write %s Go Code to file, %v", op.ExportedName, err)
	}

	return nil
}
