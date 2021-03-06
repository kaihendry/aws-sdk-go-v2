// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package iotiface provides an interface to enable mocking the AWS IoT service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package iotiface

import (
	"github.com/aws/aws-sdk-go-v2/service/iot"
)

// IoTAPI provides an interface to enable mocking the
// iot.IoT service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS IoT.
//    func myFunc(svc iotiface.IoTAPI) bool {
//        // Make svc.AcceptCertificateTransfer request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := iot.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockIoTClient struct {
//        iotiface.IoTAPI
//    }
//    func (m *mockIoTClient) AcceptCertificateTransfer(input *iot.AcceptCertificateTransferInput) (*iot.AcceptCertificateTransferOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockIoTClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type IoTAPI interface {
	AcceptCertificateTransferRequest(*iot.AcceptCertificateTransferInput) iot.AcceptCertificateTransferRequest

	AddThingToThingGroupRequest(*iot.AddThingToThingGroupInput) iot.AddThingToThingGroupRequest

	AssociateTargetsWithJobRequest(*iot.AssociateTargetsWithJobInput) iot.AssociateTargetsWithJobRequest

	AttachPolicyRequest(*iot.AttachPolicyInput) iot.AttachPolicyRequest

	AttachPrincipalPolicyRequest(*iot.AttachPrincipalPolicyInput) iot.AttachPrincipalPolicyRequest

	AttachThingPrincipalRequest(*iot.AttachThingPrincipalInput) iot.AttachThingPrincipalRequest

	CancelCertificateTransferRequest(*iot.CancelCertificateTransferInput) iot.CancelCertificateTransferRequest

	CancelJobRequest(*iot.CancelJobInput) iot.CancelJobRequest

	ClearDefaultAuthorizerRequest(*iot.ClearDefaultAuthorizerInput) iot.ClearDefaultAuthorizerRequest

	CreateAuthorizerRequest(*iot.CreateAuthorizerInput) iot.CreateAuthorizerRequest

	CreateCertificateFromCsrRequest(*iot.CreateCertificateFromCsrInput) iot.CreateCertificateFromCsrRequest

	CreateJobRequest(*iot.CreateJobInput) iot.CreateJobRequest

	CreateKeysAndCertificateRequest(*iot.CreateKeysAndCertificateInput) iot.CreateKeysAndCertificateRequest

	CreateOTAUpdateRequest(*iot.CreateOTAUpdateInput) iot.CreateOTAUpdateRequest

	CreatePolicyRequest(*iot.CreatePolicyInput) iot.CreatePolicyRequest

	CreatePolicyVersionRequest(*iot.CreatePolicyVersionInput) iot.CreatePolicyVersionRequest

	CreateRoleAliasRequest(*iot.CreateRoleAliasInput) iot.CreateRoleAliasRequest

	CreateStreamRequest(*iot.CreateStreamInput) iot.CreateStreamRequest

	CreateThingRequest(*iot.CreateThingInput) iot.CreateThingRequest

	CreateThingGroupRequest(*iot.CreateThingGroupInput) iot.CreateThingGroupRequest

	CreateThingTypeRequest(*iot.CreateThingTypeInput) iot.CreateThingTypeRequest

	CreateTopicRuleRequest(*iot.CreateTopicRuleInput) iot.CreateTopicRuleRequest

	DeleteAuthorizerRequest(*iot.DeleteAuthorizerInput) iot.DeleteAuthorizerRequest

	DeleteCACertificateRequest(*iot.DeleteCACertificateInput) iot.DeleteCACertificateRequest

	DeleteCertificateRequest(*iot.DeleteCertificateInput) iot.DeleteCertificateRequest

	DeleteOTAUpdateRequest(*iot.DeleteOTAUpdateInput) iot.DeleteOTAUpdateRequest

	DeletePolicyRequest(*iot.DeletePolicyInput) iot.DeletePolicyRequest

	DeletePolicyVersionRequest(*iot.DeletePolicyVersionInput) iot.DeletePolicyVersionRequest

	DeleteRegistrationCodeRequest(*iot.DeleteRegistrationCodeInput) iot.DeleteRegistrationCodeRequest

	DeleteRoleAliasRequest(*iot.DeleteRoleAliasInput) iot.DeleteRoleAliasRequest

	DeleteStreamRequest(*iot.DeleteStreamInput) iot.DeleteStreamRequest

	DeleteThingRequest(*iot.DeleteThingInput) iot.DeleteThingRequest

	DeleteThingGroupRequest(*iot.DeleteThingGroupInput) iot.DeleteThingGroupRequest

	DeleteThingTypeRequest(*iot.DeleteThingTypeInput) iot.DeleteThingTypeRequest

	DeleteTopicRuleRequest(*iot.DeleteTopicRuleInput) iot.DeleteTopicRuleRequest

	DeleteV2LoggingLevelRequest(*iot.DeleteV2LoggingLevelInput) iot.DeleteV2LoggingLevelRequest

	DeprecateThingTypeRequest(*iot.DeprecateThingTypeInput) iot.DeprecateThingTypeRequest

	DescribeAuthorizerRequest(*iot.DescribeAuthorizerInput) iot.DescribeAuthorizerRequest

	DescribeCACertificateRequest(*iot.DescribeCACertificateInput) iot.DescribeCACertificateRequest

	DescribeCertificateRequest(*iot.DescribeCertificateInput) iot.DescribeCertificateRequest

	DescribeDefaultAuthorizerRequest(*iot.DescribeDefaultAuthorizerInput) iot.DescribeDefaultAuthorizerRequest

	DescribeEndpointRequest(*iot.DescribeEndpointInput) iot.DescribeEndpointRequest

	DescribeEventConfigurationsRequest(*iot.DescribeEventConfigurationsInput) iot.DescribeEventConfigurationsRequest

	DescribeIndexRequest(*iot.DescribeIndexInput) iot.DescribeIndexRequest

	DescribeJobRequest(*iot.DescribeJobInput) iot.DescribeJobRequest

	DescribeJobExecutionRequest(*iot.DescribeJobExecutionInput) iot.DescribeJobExecutionRequest

	DescribeRoleAliasRequest(*iot.DescribeRoleAliasInput) iot.DescribeRoleAliasRequest

	DescribeStreamRequest(*iot.DescribeStreamInput) iot.DescribeStreamRequest

	DescribeThingRequest(*iot.DescribeThingInput) iot.DescribeThingRequest

	DescribeThingGroupRequest(*iot.DescribeThingGroupInput) iot.DescribeThingGroupRequest

	DescribeThingRegistrationTaskRequest(*iot.DescribeThingRegistrationTaskInput) iot.DescribeThingRegistrationTaskRequest

	DescribeThingTypeRequest(*iot.DescribeThingTypeInput) iot.DescribeThingTypeRequest

	DetachPolicyRequest(*iot.DetachPolicyInput) iot.DetachPolicyRequest

	DetachPrincipalPolicyRequest(*iot.DetachPrincipalPolicyInput) iot.DetachPrincipalPolicyRequest

	DetachThingPrincipalRequest(*iot.DetachThingPrincipalInput) iot.DetachThingPrincipalRequest

	DisableTopicRuleRequest(*iot.DisableTopicRuleInput) iot.DisableTopicRuleRequest

	EnableTopicRuleRequest(*iot.EnableTopicRuleInput) iot.EnableTopicRuleRequest

	GetEffectivePoliciesRequest(*iot.GetEffectivePoliciesInput) iot.GetEffectivePoliciesRequest

	GetIndexingConfigurationRequest(*iot.GetIndexingConfigurationInput) iot.GetIndexingConfigurationRequest

	GetJobDocumentRequest(*iot.GetJobDocumentInput) iot.GetJobDocumentRequest

	GetLoggingOptionsRequest(*iot.GetLoggingOptionsInput) iot.GetLoggingOptionsRequest

	GetOTAUpdateRequest(*iot.GetOTAUpdateInput) iot.GetOTAUpdateRequest

	GetPolicyRequest(*iot.GetPolicyInput) iot.GetPolicyRequest

	GetPolicyVersionRequest(*iot.GetPolicyVersionInput) iot.GetPolicyVersionRequest

	GetRegistrationCodeRequest(*iot.GetRegistrationCodeInput) iot.GetRegistrationCodeRequest

	GetTopicRuleRequest(*iot.GetTopicRuleInput) iot.GetTopicRuleRequest

	GetV2LoggingOptionsRequest(*iot.GetV2LoggingOptionsInput) iot.GetV2LoggingOptionsRequest

	ListAttachedPoliciesRequest(*iot.ListAttachedPoliciesInput) iot.ListAttachedPoliciesRequest

	ListAuthorizersRequest(*iot.ListAuthorizersInput) iot.ListAuthorizersRequest

	ListCACertificatesRequest(*iot.ListCACertificatesInput) iot.ListCACertificatesRequest

	ListCertificatesRequest(*iot.ListCertificatesInput) iot.ListCertificatesRequest

	ListCertificatesByCARequest(*iot.ListCertificatesByCAInput) iot.ListCertificatesByCARequest

	ListIndicesRequest(*iot.ListIndicesInput) iot.ListIndicesRequest

	ListJobExecutionsForJobRequest(*iot.ListJobExecutionsForJobInput) iot.ListJobExecutionsForJobRequest

	ListJobExecutionsForThingRequest(*iot.ListJobExecutionsForThingInput) iot.ListJobExecutionsForThingRequest

	ListJobsRequest(*iot.ListJobsInput) iot.ListJobsRequest

	ListOTAUpdatesRequest(*iot.ListOTAUpdatesInput) iot.ListOTAUpdatesRequest

	ListOutgoingCertificatesRequest(*iot.ListOutgoingCertificatesInput) iot.ListOutgoingCertificatesRequest

	ListPoliciesRequest(*iot.ListPoliciesInput) iot.ListPoliciesRequest

	ListPolicyPrincipalsRequest(*iot.ListPolicyPrincipalsInput) iot.ListPolicyPrincipalsRequest

	ListPolicyVersionsRequest(*iot.ListPolicyVersionsInput) iot.ListPolicyVersionsRequest

	ListPrincipalPoliciesRequest(*iot.ListPrincipalPoliciesInput) iot.ListPrincipalPoliciesRequest

	ListPrincipalThingsRequest(*iot.ListPrincipalThingsInput) iot.ListPrincipalThingsRequest

	ListRoleAliasesRequest(*iot.ListRoleAliasesInput) iot.ListRoleAliasesRequest

	ListStreamsRequest(*iot.ListStreamsInput) iot.ListStreamsRequest

	ListTargetsForPolicyRequest(*iot.ListTargetsForPolicyInput) iot.ListTargetsForPolicyRequest

	ListThingGroupsRequest(*iot.ListThingGroupsInput) iot.ListThingGroupsRequest

	ListThingGroupsForThingRequest(*iot.ListThingGroupsForThingInput) iot.ListThingGroupsForThingRequest

	ListThingPrincipalsRequest(*iot.ListThingPrincipalsInput) iot.ListThingPrincipalsRequest

	ListThingRegistrationTaskReportsRequest(*iot.ListThingRegistrationTaskReportsInput) iot.ListThingRegistrationTaskReportsRequest

	ListThingRegistrationTasksRequest(*iot.ListThingRegistrationTasksInput) iot.ListThingRegistrationTasksRequest

	ListThingTypesRequest(*iot.ListThingTypesInput) iot.ListThingTypesRequest

	ListThingsRequest(*iot.ListThingsInput) iot.ListThingsRequest

	ListThingsInThingGroupRequest(*iot.ListThingsInThingGroupInput) iot.ListThingsInThingGroupRequest

	ListTopicRulesRequest(*iot.ListTopicRulesInput) iot.ListTopicRulesRequest

	ListV2LoggingLevelsRequest(*iot.ListV2LoggingLevelsInput) iot.ListV2LoggingLevelsRequest

	RegisterCACertificateRequest(*iot.RegisterCACertificateInput) iot.RegisterCACertificateRequest

	RegisterCertificateRequest(*iot.RegisterCertificateInput) iot.RegisterCertificateRequest

	RegisterThingRequest(*iot.RegisterThingInput) iot.RegisterThingRequest

	RejectCertificateTransferRequest(*iot.RejectCertificateTransferInput) iot.RejectCertificateTransferRequest

	RemoveThingFromThingGroupRequest(*iot.RemoveThingFromThingGroupInput) iot.RemoveThingFromThingGroupRequest

	ReplaceTopicRuleRequest(*iot.ReplaceTopicRuleInput) iot.ReplaceTopicRuleRequest

	SearchIndexRequest(*iot.SearchIndexInput) iot.SearchIndexRequest

	SetDefaultAuthorizerRequest(*iot.SetDefaultAuthorizerInput) iot.SetDefaultAuthorizerRequest

	SetDefaultPolicyVersionRequest(*iot.SetDefaultPolicyVersionInput) iot.SetDefaultPolicyVersionRequest

	SetLoggingOptionsRequest(*iot.SetLoggingOptionsInput) iot.SetLoggingOptionsRequest

	SetV2LoggingLevelRequest(*iot.SetV2LoggingLevelInput) iot.SetV2LoggingLevelRequest

	SetV2LoggingOptionsRequest(*iot.SetV2LoggingOptionsInput) iot.SetV2LoggingOptionsRequest

	StartThingRegistrationTaskRequest(*iot.StartThingRegistrationTaskInput) iot.StartThingRegistrationTaskRequest

	StopThingRegistrationTaskRequest(*iot.StopThingRegistrationTaskInput) iot.StopThingRegistrationTaskRequest

	TestAuthorizationRequest(*iot.TestAuthorizationInput) iot.TestAuthorizationRequest

	TestInvokeAuthorizerRequest(*iot.TestInvokeAuthorizerInput) iot.TestInvokeAuthorizerRequest

	TransferCertificateRequest(*iot.TransferCertificateInput) iot.TransferCertificateRequest

	UpdateAuthorizerRequest(*iot.UpdateAuthorizerInput) iot.UpdateAuthorizerRequest

	UpdateCACertificateRequest(*iot.UpdateCACertificateInput) iot.UpdateCACertificateRequest

	UpdateCertificateRequest(*iot.UpdateCertificateInput) iot.UpdateCertificateRequest

	UpdateEventConfigurationsRequest(*iot.UpdateEventConfigurationsInput) iot.UpdateEventConfigurationsRequest

	UpdateIndexingConfigurationRequest(*iot.UpdateIndexingConfigurationInput) iot.UpdateIndexingConfigurationRequest

	UpdateRoleAliasRequest(*iot.UpdateRoleAliasInput) iot.UpdateRoleAliasRequest

	UpdateStreamRequest(*iot.UpdateStreamInput) iot.UpdateStreamRequest

	UpdateThingRequest(*iot.UpdateThingInput) iot.UpdateThingRequest

	UpdateThingGroupRequest(*iot.UpdateThingGroupInput) iot.UpdateThingGroupRequest

	UpdateThingGroupsForThingRequest(*iot.UpdateThingGroupsForThingInput) iot.UpdateThingGroupsForThingRequest
}

var _ IoTAPI = (*iot.IoT)(nil)
