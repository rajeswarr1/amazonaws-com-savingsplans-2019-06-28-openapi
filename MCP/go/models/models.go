package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// DeleteQueuedSavingsPlanResponse represents the DeleteQueuedSavingsPlanResponse schema from the OpenAPI specification
type DeleteQueuedSavingsPlanResponse struct {
}

// SavingsPlanOfferingRateFilterElement represents the SavingsPlanOfferingRateFilterElement schema from the OpenAPI specification
type SavingsPlanOfferingRateFilterElement struct {
	Name interface{} `json:"name,omitempty"`
	Values interface{} `json:"values,omitempty"`
}

// ParentSavingsPlanOffering represents the ParentSavingsPlanOffering schema from the OpenAPI specification
type ParentSavingsPlanOffering struct {
	Durationseconds interface{} `json:"durationSeconds,omitempty"`
	Offeringid interface{} `json:"offeringId,omitempty"`
	Paymentoption interface{} `json:"paymentOption,omitempty"`
	Plandescription interface{} `json:"planDescription,omitempty"`
	Plantype interface{} `json:"planType,omitempty"`
	Currency interface{} `json:"currency,omitempty"`
}

// DescribeSavingsPlansOfferingsRequest represents the DescribeSavingsPlansOfferingsRequest schema from the OpenAPI specification
type DescribeSavingsPlansOfferingsRequest struct {
	Offeringids interface{} `json:"offeringIds,omitempty"`
	Plantypes interface{} `json:"planTypes,omitempty"`
	Servicecodes interface{} `json:"serviceCodes,omitempty"`
	Durations interface{} `json:"durations,omitempty"`
	Maxresults interface{} `json:"maxResults,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Producttype interface{} `json:"productType,omitempty"`
	Usagetypes interface{} `json:"usageTypes,omitempty"`
	Filters interface{} `json:"filters,omitempty"`
	Operations interface{} `json:"operations,omitempty"`
	Paymentoptions interface{} `json:"paymentOptions,omitempty"`
	Descriptions interface{} `json:"descriptions,omitempty"`
	Currencies interface{} `json:"currencies,omitempty"`
}

// SavingsPlanRateFilter represents the SavingsPlanRateFilter schema from the OpenAPI specification
type SavingsPlanRateFilter struct {
	Name interface{} `json:"name,omitempty"`
	Values interface{} `json:"values,omitempty"`
}

// TagMap represents the TagMap schema from the OpenAPI specification
type TagMap struct {
}

// SavingsPlanOfferingProperty represents the SavingsPlanOfferingProperty schema from the OpenAPI specification
type SavingsPlanOfferingProperty struct {
	Name interface{} `json:"name,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

// ListTagsForResourceResponse represents the ListTagsForResourceResponse schema from the OpenAPI specification
type ListTagsForResourceResponse struct {
	Tags interface{} `json:"tags,omitempty"`
}

// SavingsPlanRateProperty represents the SavingsPlanRateProperty schema from the OpenAPI specification
type SavingsPlanRateProperty struct {
	Value interface{} `json:"value,omitempty"`
	Name interface{} `json:"name,omitempty"`
}

// TagResourceResponse represents the TagResourceResponse schema from the OpenAPI specification
type TagResourceResponse struct {
}

// CreateSavingsPlanRequest represents the CreateSavingsPlanRequest schema from the OpenAPI specification
type CreateSavingsPlanRequest struct {
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Commitment interface{} `json:"commitment"`
	Purchasetime interface{} `json:"purchaseTime,omitempty"`
	Savingsplanofferingid interface{} `json:"savingsPlanOfferingId"`
	Tags interface{} `json:"tags,omitempty"`
	Upfrontpaymentamount interface{} `json:"upfrontPaymentAmount,omitempty"`
}

// DescribeSavingsPlanRatesRequest represents the DescribeSavingsPlanRatesRequest schema from the OpenAPI specification
type DescribeSavingsPlanRatesRequest struct {
	Filters interface{} `json:"filters,omitempty"`
	Maxresults interface{} `json:"maxResults,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Savingsplanid interface{} `json:"savingsPlanId"`
}

// UntagResourceResponse represents the UntagResourceResponse schema from the OpenAPI specification
type UntagResourceResponse struct {
}

// DescribeSavingsPlanRatesResponse represents the DescribeSavingsPlanRatesResponse schema from the OpenAPI specification
type DescribeSavingsPlanRatesResponse struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Savingsplanid interface{} `json:"savingsPlanId,omitempty"`
	Searchresults interface{} `json:"searchResults,omitempty"`
}

// DescribeSavingsPlansOfferingsResponse represents the DescribeSavingsPlansOfferingsResponse schema from the OpenAPI specification
type DescribeSavingsPlansOfferingsResponse struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Searchresults interface{} `json:"searchResults,omitempty"`
}

// SavingsPlanOfferingRate represents the SavingsPlanOfferingRate schema from the OpenAPI specification
type SavingsPlanOfferingRate struct {
	Properties interface{} `json:"properties,omitempty"` // The properties.
	Rate interface{} `json:"rate,omitempty"`
	Savingsplanoffering interface{} `json:"savingsPlanOffering,omitempty"`
	Servicecode interface{} `json:"serviceCode,omitempty"`
	Unit interface{} `json:"unit,omitempty"`
	Usagetype interface{} `json:"usageType,omitempty"`
	Operation interface{} `json:"operation,omitempty"`
	Producttype interface{} `json:"productType,omitempty"`
}

// DeleteQueuedSavingsPlanRequest represents the DeleteQueuedSavingsPlanRequest schema from the OpenAPI specification
type DeleteQueuedSavingsPlanRequest struct {
	Savingsplanid interface{} `json:"savingsPlanId"`
}

// DescribeSavingsPlansRequest represents the DescribeSavingsPlansRequest schema from the OpenAPI specification
type DescribeSavingsPlansRequest struct {
	Filters interface{} `json:"filters,omitempty"`
	Maxresults interface{} `json:"maxResults,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Savingsplanarns interface{} `json:"savingsPlanArns,omitempty"`
	Savingsplanids interface{} `json:"savingsPlanIds,omitempty"`
	States interface{} `json:"states,omitempty"`
}

// SavingsPlan represents the SavingsPlan schema from the OpenAPI specification
type SavingsPlan struct {
	Description interface{} `json:"description,omitempty"`
	Ec2instancefamily interface{} `json:"ec2InstanceFamily,omitempty"`
	Producttypes interface{} `json:"productTypes,omitempty"`
	Termdurationinseconds interface{} `json:"termDurationInSeconds,omitempty"`
	State interface{} `json:"state,omitempty"`
	Upfrontpaymentamount interface{} `json:"upfrontPaymentAmount,omitempty"`
	Savingsplanarn interface{} `json:"savingsPlanArn,omitempty"`
	Commitment interface{} `json:"commitment,omitempty"`
	Offeringid interface{} `json:"offeringId,omitempty"`
	Region interface{} `json:"region,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Paymentoption interface{} `json:"paymentOption,omitempty"`
	Start interface{} `json:"start,omitempty"`
	Savingsplantype interface{} `json:"savingsPlanType,omitempty"`
	End interface{} `json:"end,omitempty"`
	Currency interface{} `json:"currency,omitempty"`
	Recurringpaymentamount interface{} `json:"recurringPaymentAmount,omitempty"`
	Savingsplanid interface{} `json:"savingsPlanId,omitempty"`
}

// DescribeSavingsPlansOfferingRatesResponse represents the DescribeSavingsPlansOfferingRatesResponse schema from the OpenAPI specification
type DescribeSavingsPlansOfferingRatesResponse struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Searchresults interface{} `json:"searchResults,omitempty"`
}

// SavingsPlanOfferingFilterElement represents the SavingsPlanOfferingFilterElement schema from the OpenAPI specification
type SavingsPlanOfferingFilterElement struct {
	Name interface{} `json:"name,omitempty"`
	Values interface{} `json:"values,omitempty"`
}

// SavingsPlanOffering represents the SavingsPlanOffering schema from the OpenAPI specification
type SavingsPlanOffering struct {
	Description interface{} `json:"description,omitempty"`
	Paymentoption interface{} `json:"paymentOption,omitempty"`
	Producttypes interface{} `json:"productTypes,omitempty"`
	Offeringid interface{} `json:"offeringId,omitempty"`
	Operation interface{} `json:"operation,omitempty"`
	Properties interface{} `json:"properties,omitempty"` // The properties.
	Usagetype interface{} `json:"usageType,omitempty"`
	Currency interface{} `json:"currency,omitempty"`
	Plantype interface{} `json:"planType,omitempty"`
	Servicecode interface{} `json:"serviceCode,omitempty"`
	Durationseconds interface{} `json:"durationSeconds,omitempty"`
}

// TagResourceRequest represents the TagResourceRequest schema from the OpenAPI specification
type TagResourceRequest struct {
	Resourcearn interface{} `json:"resourceArn"`
	Tags interface{} `json:"tags"`
}

// UntagResourceRequest represents the UntagResourceRequest schema from the OpenAPI specification
type UntagResourceRequest struct {
	Resourcearn interface{} `json:"resourceArn"`
	Tagkeys interface{} `json:"tagKeys"`
}

// ListTagsForResourceRequest represents the ListTagsForResourceRequest schema from the OpenAPI specification
type ListTagsForResourceRequest struct {
	Resourcearn interface{} `json:"resourceArn"`
}

// SavingsPlanFilter represents the SavingsPlanFilter schema from the OpenAPI specification
type SavingsPlanFilter struct {
	Name interface{} `json:"name,omitempty"`
	Values interface{} `json:"values,omitempty"`
}

// SavingsPlanOfferingRateProperty represents the SavingsPlanOfferingRateProperty schema from the OpenAPI specification
type SavingsPlanOfferingRateProperty struct {
	Name interface{} `json:"name,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

// SavingsPlanRate represents the SavingsPlanRate schema from the OpenAPI specification
type SavingsPlanRate struct {
	Unit interface{} `json:"unit,omitempty"`
	Usagetype interface{} `json:"usageType,omitempty"`
	Currency interface{} `json:"currency,omitempty"`
	Operation interface{} `json:"operation,omitempty"`
	Producttype interface{} `json:"productType,omitempty"`
	Properties interface{} `json:"properties,omitempty"` // The properties.
	Rate interface{} `json:"rate,omitempty"`
	Servicecode interface{} `json:"serviceCode,omitempty"`
}

// DescribeSavingsPlansResponse represents the DescribeSavingsPlansResponse schema from the OpenAPI specification
type DescribeSavingsPlansResponse struct {
	Savingsplans interface{} `json:"savingsPlans,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// DescribeSavingsPlansOfferingRatesRequest represents the DescribeSavingsPlansOfferingRatesRequest schema from the OpenAPI specification
type DescribeSavingsPlansOfferingRatesRequest struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Operations interface{} `json:"operations,omitempty"`
	Products interface{} `json:"products,omitempty"`
	Savingsplanpaymentoptions interface{} `json:"savingsPlanPaymentOptions,omitempty"`
	Usagetypes interface{} `json:"usageTypes,omitempty"`
	Maxresults interface{} `json:"maxResults,omitempty"`
	Savingsplanofferingids interface{} `json:"savingsPlanOfferingIds,omitempty"`
	Filters interface{} `json:"filters,omitempty"`
	Savingsplantypes interface{} `json:"savingsPlanTypes,omitempty"`
	Servicecodes interface{} `json:"serviceCodes,omitempty"`
}

// CreateSavingsPlanResponse represents the CreateSavingsPlanResponse schema from the OpenAPI specification
type CreateSavingsPlanResponse struct {
	Savingsplanid interface{} `json:"savingsPlanId,omitempty"`
}
