/*
### Source: https://github.com/unfor19/aws-sdk-golang-examples/blob/master/main.go

### Resource Groups Tagging API
GoLang SDK Docs: https://docs.aws.amazon.com/sdk-for-go/api/service/resourcegroupstaggingapi/
The package resourcegroupstaggingapi provides the client and types for making API requests to AWS Resource Groups Tagging API.
This example was written especially for GoLang newcomers (like me).
Scenario: Getting all resources ARNs which are tagged with "APP_NAME = api-group"
Given: There are two resources that are tagged with "APP_NAME=api-group", both are S3 buckets
###
*/

package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroupstaggingapi"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroupstaggingapi/types"
)

func main() {
	// Using the SDK's default configuration, loading additional config
	// and credentials values from the environment variables, shared
	// credentials, and shared configuration files
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("eu-west-1"),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Variables that will be used for querying
	tagKey := aws.String("APP_NAME")
	tagValues := []string{
		"api-group",
	}

	// Define a tagFilter
	tagFilter := types.TagFilter{}
	tagFilter.Key = tagKey
	tagFilter.Values = tagValues

	// Using the Config value, create the Resource Groups Tagging API client
	svc := resourcegroupstaggingapi.NewFromConfig(cfg)
	params := &resourcegroupstaggingapi.GetResourcesInput{
		TagFilters: []types.TagFilter{
			// {Key: aws.String("APP_NAME"), Values: []string{"group-api"}}, // Inline values
			// {Key: tagKey, Values: tagValues}, // Variables references
			tagFilter, // Variable reference to an object
		},
	}
	
	resp, err := svc.GetResources(context.TODO(), params)
	// Build the request with its input parameters
	if err != nil {
		log.Fatalf("failed to list resources, %v", err)
	}

	// GetResources
	// Returns: GetResourcesOutput { PaginationToken *string `type:"string"` , ResourceTagMappingList []*ResourceTagMapping `type:"list"` }
	// Docs: https://docs.aws.amazon.com/sdk-for-go/api/service/resourcegroupstaggingapi/#GetResourcesOutput
	/*
	The syntax "for _, res" means we ignore the first argument of the response, in this case, ignoring PaginationToken
	You should replace "_" with "pgToken" to get PaginationToken in a variable.
	*/
	for _, res := range resp.ResourceTagMappingList {
		fmt.Println("      Pointer address:", res.ResourceARN)
		fmt.Println(" Value behind pointer:", *res.ResourceARN)
		fmt.Println("                Value:", aws.ToString(res.ResourceARN))
		fmt.Println()
	}

	// Use value
	random_item_index := rand.Intn(len(resp.ResourceTagMappingList))
	s := "      Random Resource: " + *resp.ResourceTagMappingList[random_item_index].ResourceARN
	fmt.Println(s)
}

/*
### Output
      Pointer address: 0xc000393c70
 Value behind pointer: arn:aws:s3:::meir-test-1-api
                Value: arn:aws:s3:::meir-test-1-api
      Pointer address: 0xc000393cd0
 Value behind pointer: arn:aws:s3:::meir-test-2-api
                Value: arn:aws:s3:::meir-test-2-api
      Random Resource: arn:aws:s3:::meir-test-2-api
###
*/
