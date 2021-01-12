package test

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-03-01/network"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2015-11-01/resources"
	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestSpecs(t *testing.T) {
    t.Parallel()

	terraformOptions := &terraform.Options{
		TerraformDir: "spec_test_tf",
		Vars: map[string]interface{}{
			"workload_name": "terratest",
		},
	}

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	resourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")
	virtualNetworkName := terraform.Output(t, terraformOptions, "virtual_network_name")
	subscriptionID := terraform.Output(t, terraformOptions, "subscription_id")

	assert.Equal(t, "terratest-rg", resourceGroupName)

	authorizer, err := azure.NewAuthorizer()
	if err != nil {
		assert.FailNow(t, "Cannot create authorizer")
	}

	// Test Location
	resourceGroupClient := resources.NewGroupsClient(subscriptionID)
	resourceGroupClient.Authorizer = *authorizer
	resourceGroup, err := resourceGroupClient.Get(context.Background(), resourceGroupName)
	if err != nil {
		t.Log(err)
		assert.FailNow(t, "Cannot get resource group")
	}
	assert.Equal(t, "westeurope", *resourceGroup.Location, "Location must be West Europe")

	// Test network CIDR block
	virtualNetworkClient := network.NewVirtualNetworksClient(subscriptionID)
	virtualNetworkClient.Authorizer = *authorizer
	virtualNetwork, err := virtualNetworkClient.Get(context.Background(), resourceGroupName, virtualNetworkName, "")
	if err != nil {
		assert.FailNow(t, "Cannot get network")
	}
	assert.Equal(t, "10.20.0.0/16", (*virtualNetwork.AddressSpace.AddressPrefixes)[0], "Network must be in the 10.20.0.0/16 block")
}
