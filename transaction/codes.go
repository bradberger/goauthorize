package transaction

import (
	"encoding/json"
)

// ResponseCodes is a JSON object pulled directly from Authorize.net which contains
// all possible response codes and their meanings.
var ResponseCodes = `
[	{
		"code" : "I00001",
		"text" : "Successful.",
		"description" : "The request was processed successfully.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "I00002",
		"text" : "The subscription has already been canceled.",
		"description" : "The subscription has already been canceled.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "I00003",
		"text" : "The record has already been deleted.",
		"description" : "The record has already been deleted.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "I00004",
		"text" : "No records found.",
		"description" : "No records have been found that match your query.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "I00005",
		"text" : "The mobile device has been submitted for approval by the account administrator.",
		"description" : "The mobile device was successfully inserted into the database.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "I00006",
		"text" : "The mobile device is approved and ready for use.",
		"description" : "The mobile device was successfully registered and approved by the account administrator.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "I00007",
		"text" : "The Payment Gateway Account service (id&#x3D;8) has already been accepted.",
		"description" : "The Payment Gateway Account service (id&#x3D;8) has already been accepted.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "I00008",
		"text" : "The Payment Gateway Account service (id&#x3D;8) has already been declined.",
		"description" : "The Payment Gateway Account service (id&#x3D;8) has already been declined.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "I00009",
		"text" : "The APIUser already exists.",
		"description" : "The APIUser already exists.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "I00010",
		"text" : "The merchant is activated successfully.",
		"description" : "The merchant is activated successfully.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "I00011",
		"text" : "The merchant is not activated.",
		"description" : "The merchant is not activated.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "I99999",
		"text" : "This feature has not yet been completed. One day it will be but it looks like today is not that day.",
		"description" : "This is a work in progress. This message will not be released to production. It&#x27;s just a dev placeholder.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00001",
		"text" : "An error occurred during processing. Please try again.",
		"description" : "An unexpected system error occurred while processing this request.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00002",
		"text" : "The content-type specified is not supported.",
		"description" : "The only supported content-types are text/xml and application/xml.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00003",
		"text" : "An error occurred while parsing the XML request.",
		"description" : "This is the result of an XML parser error.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00004",
		"text" : "The name of the requested API method is invalid.",
		"description" : "The name of the root node of the XML request is the API method being called. It is not valid.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00005",
		"text" : "The transaction key or API key is invalid or not present.",
		"description" : "User authentication requires a valid value for transaction key or API key.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00006",
		"text" : "The API user name is invalid or not present.",
		"description" : "User authentication requires a valid value for API user name.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00007",
		"text" : "User authentication failed due to invalid authentication values.",
		"description" : "The API user name is invalid and/or the transaction key or API key is invalid.",
		"integration_suggestions": "This error indicates that invalid credentials, the API Login ID or Transaction Key, are being submitted. If you have confirmed that your API login ID and Transaction Key are accurate, you may need to confirm that you are submitting to the correct URL. If you are using a test account, please make sure to post to the sandbox URL. If youâ€™re using a live account, make sure to post to the production URL.",
		"other_suggestions": ""
	},
	{
		"code" : "E00008",
		"text" : "User authentication failed. The account or API user is inactive.",
		"description" : "The payment gateway, reseller, or user account is not currently active.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00009",
		"text" : "The payment gateway account is in Test Mode. The request cannot be processed.",
		"description" : "The requested API method cannot be executed while the payment gateway account is in Test Mode.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00010",
		"text" : "User authentication failed. You do not have the appropriate permissions.",
		"description" : "The user does not have permission to call the API.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00011",
		"text" : "Access denied. You do not have the appropriate permissions.",
		"description" : "The user does not have permission to call the API method.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00012",
		"text" : "A duplicate subscription already exists.",
		"description" : "A duplicate of the subscription was already submitted.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00013",
		"text" : "The field is invalid.",
		"description" : "One of the field values is not valid.",
		"integration_suggestions": "One of the field values is not valid. The response text field should provide you the details of which &amp;quot;field&amp;quot; exactly is invalid so check the response text.",
		"other_suggestions": ""
	},
	{
		"code" : "E00014",
		"text" : "A required field is not present.",
		"description" : "One of the required fields was not present.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00015",
		"text" : "The field length is invalid.",
		"description" : "One of the fields has an invalid length.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00016",
		"text" : "The field type is invalid.",
		"description" : "The field type is not valid.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00017",
		"text" : "The start date cannot occur in the past.",
		"description" : "The subscription start date cannot occur before the subscription submission date.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00018",
		"text" : "The credit card expires before the subscription start date.",
		"description" : "The credit card is not valid as of the start date of the subscription.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00019",
		"text" : "The customer tax id or drivers license information is required.",
		"description" : "The customer tax ID or driver&#x27;s license information (driver&#x27;s license number, driver&#x27;s license state, driver&#x27;s license DOB) is required for the subscription.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00020",
		"text" : "The payment gateway account is not enabled for eCheck.Net subscriptions.",
		"description" : "The payment gateway account is not set up to process eCheck.Net subscriptions.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00021",
		"text" : "The payment gateway account is not enabled for credit card subscriptions.",
		"description" : "The payment gateway account is not set up to process credit card subscriptions.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00022",
		"text" : "The interval length cannot exceed 365 days or 12 months.",
		"description" : "The interval length must be 7 to 365 days or 1 to 12 months.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00023",
		"text" : "The subscription duration cannot exceed three years.",
		"description" : "The number of total occurrences cannot extend the duration of the subscription beyond three years from the start date.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00024",
		"text" : "Trial Occurrences is required when Trial Amount is specified.",
		"description" : "The number of trial occurrences cannot be zero if a valid trial amount is submitted.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00025",
		"text" : "Automated Recurring Billing is not enabled.",
		"description" : "The payment gateway account is not enabled for Automated Recurring Billing.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00026",
		"text" : "Both Trial Amount and Trial Occurrences are required.",
		"description" : "If either a trial amount or number of trial occurrences is specified then values for both must be submitted.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00027",
		"text" : "The transaction was unsuccessful.",
		"description" : "An approval was not returned for the transaction.",
		"integration_suggestions": "",
		"other_suggestions": "For more information, check the errorCode field in the response."
	},
	{
		"code" : "E00028",
		"text" : "Trial Occurrences must be less than Total Occurrences.",
		"description" : "The number of trial occurrences specified must be less than the number of total occurrences specified.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00029",
		"text" : "Payment information is required.",
		"description" : "Payment information is required when creating a subscription or payment profile.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00030",
		"text" : "The payment schedule is required.",
		"description" : "A payment schedule is required when creating a subscription.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00031",
		"text" : "The amount is required.",
		"description" : "The subscription amount is required when creating a subscription.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00032",
		"text" : "The start date is required.",
		"description" : "The subscription start date is required to create a subscription.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00033",
		"text" : "The start date cannot be changed.",
		"description" : "Once a subscription is created the start date cannot be changed.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00034",
		"text" : "The interval information cannot be changed.",
		"description" : "Once a subscription is created the interval cannot be changed.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00035",
		"text" : "The subscription cannot be found.",
		"description" : "The subscription ID for this request is not valid for this merchant.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00036",
		"text" : "The payment type cannot be changed.",
		"description" : "Changing the subscription payment type between credit card and eCheck.Net is not currently supported.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00037",
		"text" : "The subscription cannot be updated.",
		"description" : "Subscriptions that are expired, canceled or terminated cannot be updated.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00038",
		"text" : "The subscription cannot be canceled.",
		"description" : "Subscriptions that are expired or terminated cannot be canceled.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00039",
		"text" : "A duplicate record already exists.",
		"description" : "A duplicate of the customer profile, customer payment profile, or customer address was already submitted.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00040",
		"text" : "The record cannot be found.",
		"description" : "The customer profile ID, payment profile ID, shipping address ID, or transaction ID for this request is not valid for this merchant.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00041",
		"text" : "One or more fields must contain a value.",
		"description" : "All of the fields were empty or missing.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00042",
		"text" : "You cannot add more than {0} payment profiles.",
		"description" : "The maximum number of payment profiles for the customer profile has been reached.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00043",
		"text" : "You cannot add more than {0} shipping addresses.",
		"description" : "The maximum number of shipping addresses for the customer profile has been reached.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00044",
		"text" : "Customer Information Manager is not enabled.",
		"description" : "The payment gateway account is not enabled for Customer Information Manager (CIM).",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00045",
		"text" : "The root node does not reference a valid XML namespace.",
		"description" : "The root node does not reference a valid XML namespace.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00046",
		"text" : "Generic InsertNewMerchant failure.",
		"description" : "Generic InsertNewMerchant failure.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00047",
		"text" : "Merchant Boarding API is not enabled.",
		"description" : "The reseller account is not enabled for Merchant Boarding API.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00048",
		"text" : "At least one payment method must be set in payment types or an echeck service must be provided.",
		"description" : "The merchant account must be set up to accept credit card payments, eCheck payments, or both.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00049",
		"text" : "The operation timed out before it could be completed.",
		"description" : "The database operation timed out before it could be completed.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00050",
		"text" : "Sell Rates cannot be less than Buy Rates",
		"description" : "Cannot set a buyrate to less than the sellrate",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00051",
		"text" : "The original transaction was not issued for this payment profile.",
		"description" : "If customer profile ID, payment profile ID, and shipping address ID are included, they must match the original transaction.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00052",
		"text" : "The maximum number of elements for an array {0} is {1}.",
		"description" : "The maximum number of elements for an array has been reached.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00053",
		"text" : "Server too busy",
		"description" : "The server is currently too busy, please try again later.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00054",
		"text" : "The mobile device is not registered with this merchant account.",
		"description" : "The mobile device identifier is not associated with the merchant account.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00055",
		"text" : "The mobile device has already been registered but is pending approval by the account administrator.",
		"description" : "The mobile device exists but is in a pending status.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00056",
		"text" : "The mobile device has been disabled for use with this account.",
		"description" : "The mobile device exists but has a status of disabled.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00057",
		"text" : "The user does not have permissions to submit requests from a mobile device.",
		"description" : "The user does not have sufficient permissions to use a mobile device with this merchant account.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00058",
		"text" : "The merchant has met or exceeded the number of pending mobile devices permitted for this account.",
		"description" : "The merchant has too many devices in a pending status.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00059",
		"text" : "The authentication type is not allowed for this method call.",
		"description" : "The authentication type is not allowed for the requested method call.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00060",
		"text" : "The transaction type is invalid.",
		"description" : "The transaction type is invalid.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00061",
		"text" : "{0}({1}).",
		"description" : "Could not decrypt DUKPT blobs and returned error.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00062",
		"text" : "Fatal error when calling web service.",
		"description" : "Fatal error when calling web service.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00063",
		"text" : "Calling web service return error.",
		"description" : "Calling web service return error.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00064",
		"text" : "Client authorization denied.",
		"description" : "Client authorization denied.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00065",
		"text" : "Prerequisite failed.",
		"description" : "Prerequisite failed.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00066",
		"text" : "Invalid value.",
		"description" : "Invalid value.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00067",
		"text" : "An error occurred while parsing the XML request.  Too many {0} specified.",
		"description" : "This is the result of an XML parser error.  Too many nodes specified.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00068",
		"text" : "An error occurred while parsing the XML request.  {0} is invalid.",
		"description" : "This is the result of an XML parser error.  The node is invalid.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00069",
		"text" : "The Payment Gateway Account service (id&#x3D;8) has already been accepted.  Decline is not allowed.",
		"description" : "The Payment Gateway Account service (id&#x3D;8) has already been accepted.  Decline is not allowed.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00070",
		"text" : "The Payment Gateway Account service (id&#x3D;8) has already been declined.  Agree is not allowed.",
		"description" : "The Payment Gateway Account service (id&#x3D;8) has already been declined.  Agree is not allowed.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00071",
		"text" : "{0} must contain data.",
		"description" : "All of the fields were empty or missing.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00072",
		"text" : "Node {0} is required.",
		"description" : "Required node missing.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00073",
		"text" : "{0} is invalid.",
		"description" : "One of the field values is not valid.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00074",
		"text" : "This merchant is not associated with this reseller.",
		"description" : "This merchant is not associated with this reseller.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00075",
		"text" : "An error occurred while parsing the XML request.  Missing field(s) {0}.",
		"description" : "This is the result of an XML parser error.  Missing field(s).",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00076",
		"text" : "{0} contains invalid value.",
		"description" : "Invalid value.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00077",
		"text" : "The value of {0} is too long.  The length of value should be {1}",
		"description" : "Value too long.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00078",
		"text" : "Pending Status (not completed).",
		"description" : "Pending Status (not completed).",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00079",
		"text" : "The impersonation login ID is invalid or not present.",
		"description" : "Impersonation partner login ID is invalid or not present.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00080",
		"text" : "The impersonation API Key is invalid or not present.",
		"description" : "Impersonation API Key is invalid or not present.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00081",
		"text" : "Partner account is not authorized to impersonate the login account.",
		"description" : "The partner account is not authorized to impersonate the login account.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00082",
		"text" : "Country for {0} is not valid.",
		"description" : "Country is not valid.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00083",
		"text" : "Bank payment method is not accepted for the selected business country.",
		"description" : "Bank payment method is not accepted for the selected business country.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00084",
		"text" : "Credit card payment method is not accepted for the selected business country.",
		"description" : "Credit card payment method is not accepted for the selected business country.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00085",
		"text" : "State for {0} is not valid.",
		"description" : "State is not valid.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00086",
		"text" : "Merchant has declined authorization to resource.",
		"description" : "Merchant has declined authorization to resource.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00087",
		"text" : "No subscriptions found for the given request.",
		"description" : "There are no subscriptions available for the merchant account for the type of subscriptions requested.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00088",
		"text" : "ProfileIds cannot be sent when requesting CreateProfile.",
		"description" : "CreateProfile and profileIds are mutually exclusive, only one of them can be provided at a time.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00089",
		"text" : "Payment data is required when requesting CreateProfile.",
		"description" : "When requesting CreateProfile payment data cannot be null.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00090",
		"text" : "PaymentProfile cannot be sent with payment data.",
		"description" : "PaymentProfile and PaymentData are mutually exclusive, only one of them can be provided at a time.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00091",
		"text" : "PaymentProfileId cannot be sent with payment data.",
		"description" : "PaymentProfileId and payment data are mutually exclusive, only one of them can be provided at a time.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00092",
		"text" : "ShippingProfileId cannot be sent with ShipTo data.",
		"description" : "ShippingProfileId and ShipToAddress are mutually exclusive, only one of them can be provided at a time.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00093",
		"text" : "PaymentProfile cannot be sent with billing data.",
		"description" : "PaymentProfile and Billing information are mutually exclusive, only one of them can be provided at a time.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00094",
		"text" : "Paging Offset exceeds the maximum allowed value.",
		"description" : "Paging Offset exceeds allowed value. Check and lower the value.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00095",
		"text" : "ShippingProfileId is not provided within Customer Profile.",
		"description" : "When using Customer Profile with Credit Card Info to specify Shipping Profile, Shipping Profile Id must be included.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00096",
		"text" : "Finger Print value is not valid.",
		"description" : "Finger Print value is not valid.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00097",
		"text" : "Finger Print can&#x27;t be generated.",
		"description" : "Finger Print can&#x27;t be generated.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00098",
		"text" : "Customer Profile ID or Shipping Profile ID not found.",
		"description" : "Search for shipping profile using customer profile id and shipping profile id did not find any records.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00099",
		"text" : "Customer profile creation failed. This transaction ID is invalid.",
		"description" : "Customer profile creation failed. This transaction ID is invalid.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00100",
		"text" : "Customer profile creation failed. This transaction type does not support profile creation.",
		"description" : "Customer profile creation failed. This transaction type does not support profile creation.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00101",
		"text" : "Customer profile creation failed.",
		"description" : "Error creating a customer payment profile from transaction.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00102",
		"text" : "Customer Info is missing.",
		"description" : "Error creating a customer profile from transaction.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00103",
		"text" : "Customer profile creation failed. This payment method does not support profile creation.",
		"description" : "Customer profile creation failed. This payment method does not support profile creation.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00104",
		"text" : "Server in maintenance",
		"description" : "The server is in maintenance, so the requested method is unavailable, please try again later.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00105",
		"text" : "The specified payment profile is associated with an active or suspended subscription and cannot be deleted.",
		"description" : "The specified payment profile is associated with an active or suspended subscription and cannot be deleted.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00106",
		"text" : "The specified customer profile is associated with an active or suspended subscription and cannot be deleted.",
		"description" : "The specified customer profile is associated with an active or suspended subscription and cannot be deleted.",
		"integration_suggestions": "",
		"other_suggestions": ""
	},
	{
		"code" : "E00107",
		"text" : "The specified shipping profile is associated with an active or suspended subscription and cannot be deleted.",
		"description" : "The specified shipping profile is associated with an active or suspended subscription and cannot be deleted.",
		"integration_suggestions": "",
		"other_suggestions": ""
	}
,	{
		"code" : "0",
		"text" : "Unknown Error.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "1",
		"text" : "This transaction has been approved.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "2",
		"text" : "This transaction has been declined.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "3",
		"text" : "This transaction has been declined.",
		"integration_suggestions" : "A referral to a voice authorization center was received.  Please call the appropriate number below for a voice authorization.&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;For American Express call: (800) 528-2121&amp;lt;br /&amp;gt;For Diners Club call: (800) 525-9040&amp;lt;br /&amp;gt;For Discover/Novus call: (800) 347-1111&amp;lt;br /&amp;gt;For JCB call : (800) 522-9345&amp;lt;br /&amp;gt;For Visa/MasterCard call: (800) 228-1122&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;Once an authorization is issued, you can then submit the transaction through your Virtual Terminal as a Capture Only transaction.",
		"other_suggestions" : ""
	},
	{
		"code" : "4",
		"text" : "This transaction has been declined.",
		"integration_suggestions" : "",
		"other_suggestions" : "The code returned from the processor indicating that the card used needs to be picked up."
	},
	{
		"code" : "5",
		"text" : "A valid amount is required.",
		"integration_suggestions" : "",
		"other_suggestions" : "The value submitted in the amount field did not pass validation for a number."
	},
	{
		"code" : "6",
		"text" : "The credit card number is invalid.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "7",
		"text" : "Credit card expiration date is invalid.",
		"integration_suggestions" : "",
		"other_suggestions" : "The format of the date submitted was incorrect."
	},
	{
		"code" : "8",
		"text" : "The credit card has expired.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "9",
		"text" : "The ABA code is invalid",
		"integration_suggestions" : "",
		"other_suggestions" : "The value submitted in the x_bank_aba_code field did not pass validation or was not for a valid financial institution."
	},
	{
		"code" : "10",
		"text" : "The account number is invalid",
		"integration_suggestions" : "",
		"other_suggestions" : "The value submitted in the &amp;lt;i&amp;gt;x_bank_acct_num&amp;lt;/i&amp;gt; field did not pass validation."
	},
	{
		"code" : "11",
		"text" : "A duplicate transaction has been submitted.",
		"integration_suggestions" : "The error message &amp;quot;Duplicate Transaction&amp;quot; indicates that a transaction request with the same information has been submitted within two minutes of a previous attempt. Authorize.Net looks for transactions which are likely to be duplicates by matching the data provided with the transaction.&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;The fields that are validated include:&amp;lt;br /&amp;gt;&amp;lt;table cellpadding&amp;#x3D;&amp;quot;0&amp;quot; cellpadding&amp;#x3D;&amp;quot;0&amp;quot;&amp;gt;&amp;lt;tr&amp;gt;&amp;lt;td width&amp;#x3D;&amp;quot;160&amp;quot;&amp;gt;API Login ID&amp;lt;/td&amp;gt;&amp;lt;td&amp;gt;&amp;lt;i&amp;gt;x_login&amp;lt;/i&amp;gt;&amp;lt;/td&amp;gt;&amp;lt;/tr&amp;gt;&amp;lt;tr&amp;gt;&amp;lt;td&amp;gt;Credit Card Number&amp;lt;/td&amp;gt;&amp;lt;td&amp;gt;&amp;lt;i&amp;gt;x_card_num&amp;lt;/i&amp;gt;&amp;lt;/td&amp;gt;&amp;lt;/tr&amp;gt;&amp;lt;tr&amp;gt;&amp;lt;td&amp;gt;Expiration Date&amp;lt;/td&amp;gt;&amp;lt;td&amp;gt;&amp;lt;i&amp;gt;x_exp_date&amp;lt;/i&amp;gt;&amp;lt;/td&amp;gt;&amp;lt;/tr&amp;gt;&amp;lt;tr&amp;gt;&amp;lt;td&amp;gt;Transaction Type&amp;lt;/td&amp;gt;&amp;lt;td&amp;gt;&amp;lt;i&amp;gt;x_type&amp;lt;/i&amp;gt;&amp;lt;/td&amp;gt;&amp;lt;/tr&amp;gt;&amp;lt;tr&amp;gt;&amp;lt;td&amp;gt;Bank Account Number&amp;lt;/td&amp;gt;&amp;lt;td&amp;gt;&amp;lt;i&amp;gt;x_bank_acct_num&amp;lt;/i&amp;gt;&amp;lt;/td&amp;gt;&amp;lt;/tr&amp;gt;&amp;lt;tr&amp;gt;&amp;lt;td&amp;gt;Routing Number&amp;lt;/td&amp;gt;&amp;lt;td&amp;gt;&amp;lt;i&amp;gt;x_bank_aba_code&amp;lt;/td&amp;gt;&amp;lt;/tr&amp;gt;&amp;lt;tr&amp;gt;&amp;lt;td&amp;gt;Purchase Order Number&amp;lt;/td&amp;gt;&amp;lt;td&amp;gt;&amp;lt;i&amp;gt;x_po_num&amp;lt;/i&amp;gt;&amp;lt;/td&amp;gt;&amp;lt;/tr&amp;gt;&amp;lt;tr&amp;gt;&amp;lt;td&amp;gt;Amount&amp;lt;/td&amp;gt;&amp;lt;td&amp;gt;&amp;lt;i&amp;gt;x_amount&amp;lt;/i&amp;gt;&amp;lt;/td&amp;gt;&amp;lt;/tr&amp;gt;&amp;lt;tr&amp;gt;&amp;lt;td&amp;gt;Invoice Number&amp;lt;/td&amp;gt;&amp;lt;td&amp;gt;&amp;lt;i&amp;gt;x_invoice_num&amp;lt;/i&amp;gt;&amp;lt;/td&amp;gt;&amp;lt;/tr&amp;gt;&amp;lt;tr&amp;gt;&amp;lt;td&amp;gt;Customer ID&amp;lt;/td&amp;gt;&amp;lt;td&amp;gt;&amp;lt;i&amp;gt;x_cust_id&amp;lt;/i&amp;gt;&amp;lt;/td&amp;gt;&amp;lt;/tr&amp;gt;&amp;lt;tr&amp;gt;&amp;lt;td&amp;gt;First Name&amp;lt;/td&amp;gt;&amp;lt;td&amp;gt;&amp;lt;i&amp;gt;x_first_name&amp;lt;/i&amp;gt;&amp;lt;/td&amp;gt;&amp;lt;/tr&amp;gt;&amp;lt;tr&amp;gt;&amp;lt;td&amp;gt;Last Name&amp;lt;/td&amp;gt;&amp;lt;td&amp;gt;&amp;lt;i&amp;gt;x_last_name&amp;lt;/i&amp;gt;&amp;lt;/td&amp;gt;&amp;lt;/tr&amp;gt;&amp;lt;tr&amp;gt;&amp;lt;td&amp;gt;Company&amp;lt;/td&amp;gt;&amp;lt;td&amp;gt;&amp;lt;i&amp;gt;x_company&amp;lt;/i&amp;gt;&amp;lt;/td&amp;gt;&amp;lt;/tr&amp;gt;&amp;lt;tr&amp;gt;&amp;lt;td&amp;gt;Address&amp;lt;/td&amp;gt;&amp;lt;td&amp;gt;&amp;lt;i&amp;gt;x_address&amp;lt;/i&amp;gt;&amp;lt;/td&amp;gt;&amp;lt;/tr&amp;gt;&amp;lt;tr&amp;gt;&amp;lt;td&amp;gt;City&amp;lt;/td&amp;gt;&amp;lt;td&amp;gt;&amp;lt;i&amp;gt;x_city&amp;lt;/i&amp;gt;&amp;lt;/td&amp;gt;&amp;lt;/tr&amp;gt;&amp;lt;tr&amp;gt;&amp;lt;td&amp;gt;State&amp;lt;/td&amp;gt;&amp;lt;td&amp;gt;&amp;lt;i&amp;gt;x_state&amp;lt;/i&amp;gt;&amp;lt;/td&amp;gt;&amp;lt;/tr&amp;gt;&amp;lt;tr&amp;gt;&amp;lt;td&amp;gt;Postal Code&amp;lt;/td&amp;gt;&amp;lt;td&amp;gt;&amp;lt;i&amp;gt;x_zip&amp;lt;/i&amp;gt;&amp;lt;/td&amp;gt;&amp;lt;/tr&amp;gt;&amp;lt;tr&amp;gt;&amp;lt;td&amp;gt;Country&amp;lt;/td&amp;gt;&amp;lt;td&amp;gt;&amp;lt;i&amp;gt;x_country&amp;lt;/i&amp;gt;&amp;lt;/td&amp;gt;&amp;lt;/tr&amp;gt;&amp;lt;tr&amp;gt;&amp;lt;td&amp;gt;Duplicate Window&amp;lt;/td&amp;gt;&amp;lt;td&amp;gt;&amp;lt;i&amp;gt;x_duplicate_window&amp;lt;/i&amp;gt;&amp;lt;/td&amp;gt;&amp;lt;/tr&amp;gt;&amp;lt;/table&amp;gt;&amp;lt;br /&amp;gt;If any of the fields change from one transaction to the next, Authorize.Net will not view the transactions as duplicates.&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;The duplicate transaction window will always be two minutes for all transactions submitted through the Virtual Terminal. If you wish to adjust the duplicate transaction window for transactions submitted from your software, such as a website or shopping cart, you may do so by adding the field &amp;lt;i&amp;gt;x_duplicate_window&amp;lt;/i&amp;gt; to your website script.&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;If you are unable to adjust or add this variable to your shopping cart settings, please contact your shopping cart support team for additional assistance in this regard. The variable &amp;lt;i&amp;gt;x_duplicate_window&amp;lt;/i&amp;gt; tells us, &amp;lt;i&amp;gt;in seconds&amp;lt;/i&amp;gt;, how much time we should check for duplicates after a transaction is submitted.&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;The largest value we will accept for &amp;lt;i&amp;gt;x_duplicate_window&amp;lt;/i&amp;gt; is 28800, which equals eight hours. If a value greater than 28800 sent, the payment gateway will default to 28800. If &amp;lt;i&amp;gt;x_duplicate_window&amp;lt;/i&amp;gt; is set to 0 or to a negative number, no duplicate transaction window will be enforced for your software&amp;#x27;s transactions. If no value is sent, the default value of 120 (two minutes) would be used.&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;For example, if you wished to set a duplicate transaction window of five minutes, you would set &amp;lt;i&amp;gt;x_duplicate_window&amp;lt;/i&amp;gt; to 300, as there are 300 seconds in five minutes.&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;&amp;lt;strong&amp;gt;Note&amp;lt;/strong&amp;gt;: By submitting &amp;lt;i&amp;gt;x_duplicate_window&amp;lt;/i&amp;gt; with your online transactions, we will return further details along with this error response, including:&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;&amp;lt;ul class&amp;#x3D;&amp;quot;graydot&amp;quot;&amp;gt;&amp;lt;li&amp;gt;The original transaction ID that was duplicated;&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;AVS and CVV responses for the original transaction;&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;The original authorization code, if the transaction was authorized;&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;The MD5 hash, if a MD5 hash value was generated for the original transaction.&amp;lt;/li&amp;gt;&amp;lt;/ul&amp;gt;&amp;lt;p&amp;gt;If you do not submit the &amp;lt;i&amp;gt;x_duplicate_window&amp;lt;/i&amp;gt; field, we will not return any details from the original transaction, even if you submit a duplicate transaction.",
		"other_suggestions" : "A transaction with identical amount and credit card information was submitted within the previous two minutes."
	},
	{
		"code" : "12",
		"text" : "An authorization code is required but not present.",
		"integration_suggestions" : "",
		"other_suggestions" : "A transaction that required &amp;lt;i&amp;gt;x_auth_code&amp;lt;/i&amp;gt; to be present was submitted without a value."
	},
	{
		"code" : "13",
		"text" : "The merchant login ID or password is invalid or the account is inactive.",
		"integration_suggestions" : "This error indicates you are either posting the incorrect API Login ID within your script, connecting to a server that does not recognize your account, or using an account which is inactive. Please follow these steps to ensure that your software is connecting correctly:&amp;lt;ul class&amp;#x3D;&amp;quot;graydot&amp;quot;&amp;gt;&amp;lt;li&amp;gt;If you are posting your transaction requests to the gateway URLs https://test.authorize.net/gateway/transact.dll or https://certification.authorize.net/gateway/transact.dll and you are using an account given to you by an Authorize.Net Reseller or from Authorize.Net Sales, you may encounter this error. The gateway URLs mentioned above only work with specific test accounts, available upon request by completing the form at &amp;lt;a href&amp;#x3D;&amp;quot;http://developer.authorize.net/testaccount/&amp;quot;&amp;gt;http://developer.authorize.net/testaccount/&amp;lt;/a&amp;gt;&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;Try posting your transaction request to https://secure.authorize.net/gateway/transact.dll instead. If you need to submit a test transaction, you may do so by setting the field &amp;lt;em&amp;gt;x_test_request&amp;lt;/em&amp;gt; to &amp;quot;TRUE&amp;quot;. You may then remove &amp;lt;em&amp;gt;x_test_request&amp;lt;/em&amp;gt; or set it to &amp;quot;FALSE&amp;quot; when you have completed your testing.&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;Note: If you are integrating your website with Advanced Integration Method (AIM) using PHP, please be aware that the gateway URL is listed twice--once in the cURL configuration, and once again in the main body of the code. Using the test gateway URL in the cURL section of the code will create an Error 13 as cURL would connect to the incorrect server.&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;Please check your script and verify you are posting the API Login ID for the account in question. If you are not posting the correct API Login ID, or if you are not sending an API Login ID, please edit the script and confirm that the field &amp;lt;em&amp;gt;x_login&amp;lt;/em&amp;gt; is set to the API Login ID that you may obtain from the Authorize.Net Merchant Interface. Please see the Merchant Integration Guide at &amp;lt;a href&amp;#x3D;&amp;quot;http://www.authorize.net/support/merchant_guide.pdf&amp;quot;&amp;gt;http://www.authorize.net/support/merchant_guide.pdf&amp;lt;/a&amp;gt; for instructions on obtaining the API Login ID. For further information on the API Login ID, Transaction Key, or Test Mode, please check out our online video tutorials at &amp;lt;a href&amp;#x3D;&amp;quot;http://www.authorize.net/videos/&amp;quot;&amp;gt;http://www.authorize.net/videos/&amp;lt;/a&amp;gt;.&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;If you are unable to log into your Authorize.Net Merchant Interface, this could indicate that your account is inactive. If this is the case please contact us directly to reactivate the account. You can reach us toll-free at: (877) 447-3938. We are open Monday through Friday, 5 AM to 5 PM Pacific time (closed major holidays). We will be happy to assist you at that time.&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;If the account is active and the API Login ID is correct, you will want to use the &amp;lt;strong&amp;gt;Data Validation URL&amp;lt;/strong&amp;gt; to see which variables are being passed. For more details, please visit http://developer.authorize.net/tools/datavalidation/&amp;lt;/li&amp;gt;&amp;lt;/ul&amp;gt;",
		"other_suggestions" : ""
	},
	{
		"code" : "14",
		"text" : "The referrer, relay response or receipt link URL is invalid.",
		"integration_suggestions" : "The Invalid Referrer or Relay Response URL message can be the result of a few different situations that could impact SIM  users.&amp;lt;br /&amp;gt;&amp;lt;u&amp;gt;SIM Integration Suggestions&amp;lt;/u&amp;gt;&amp;lt;br /&amp;gt;Error 14 occurs for SIM users when one or more URLs are specified in your Default Response/Receipt URL settings, but we receive a URL with your transaction that does not match any listed here. In this case the transaction would include the field &amp;lt;i&amp;gt;x_receipt_link_url&amp;lt;/i&amp;gt; if you prefer to view our receipt page but provide a link for the customer to use to return to your site. Alternately, the transaction would include the field &amp;lt;i&amp;gt;x_relay_url&amp;lt;/i&amp;gt; to specify which web page on your server should be used as the receipt page your customers see.&amp;lt;br /&amp;gt;By designating a Default Response/Receipt URL, you are telling our system to only return results to one of the listed URLs. If the value of either &amp;lt;i&amp;gt;x_receipt_link_url&amp;lt;/i&amp;gt; or &amp;lt;i&amp;gt;x_relay_url&amp;lt;/i&amp;gt; does not match one of these designated Default Response/Receipt URLs, Error 14 will occur.&amp;lt;br /&amp;gt;While &amp;lt;i&amp;gt;x_receipt_link_url&amp;lt;/i&amp;gt; and &amp;lt;i&amp;gt;x_relay_url&amp;lt;/i&amp;gt; will work without specifying a Default Response/Receipt URL, it is strongly suggested that you set Default Response/Receipt URLs if you are sending either of these fields, to ensure that only these URLs can be used.&amp;lt;br /&amp;gt;To add a valid Response/Receipt URL, please follow these steps:&amp;lt;ul class&amp;#x3D;&amp;quot;graydot&amp;quot;&amp;gt;&amp;lt;li&amp;gt;Login to your Merchant Interface at &amp;lt;a href&amp;#x3D;&amp;quot;https://account.authorize.net/&amp;quot;&amp;gt;https://account.authorize.net/&amp;lt;/a&amp;gt;&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;Click &amp;lt;strong&amp;gt;Settings&amp;lt;/strong&amp;gt; in the main left side menu&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;Click &amp;lt;strong&amp;gt;Response/Receipt URLs&amp;lt;/strong&amp;gt;&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;Click &amp;lt;strong&amp;gt;Add URL&amp;lt;/strong&amp;gt;&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;Enter your &amp;lt;strong&amp;gt;Response URL&amp;lt;/strong&amp;gt;&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;Click &amp;lt;strong&amp;gt;Submit&amp;lt;/strong&amp;gt;&amp;lt;/li&amp;gt;&amp;lt;/ul&amp;gt;",
		"other_suggestions" : "Applicable only to SIM API. The Relay Response or Referrer URL does not match the merchant&amp;#x27;s configured value(s) or is absent."
	},
	{
		"code" : "15",
		"text" : "The transaction ID is invalid or not present.",
		"integration_suggestions" : "",
		"other_suggestions" : "The transaction ID value is non-numeric or was not present for a transaction that requires it (i.e., VOID, PRIOR_AUTH_CAPTURE, and CREDIT)."
	},
	{
		"code" : "16",
		"text" : "The transaction cannot be found.",
		"integration_suggestions" : "This error may be caused by a refund request if the referenced transaction ID (&amp;lt;i&amp;gt;x_trans_id&amp;lt;/i&amp;gt;) was originally processed through a different Authorize.Net account than the one being used for the refund request. Please submit refund transactions using the gateway account that generated the original transaction.&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;The error could also indicate a setup problem with a particular card type. Please contact your Merchant Service Provider (MSP) to check on your payment processing setup and to confirm that there are no issues with the configuration for the card type being submitted in the transaction.&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;Your MSP is the company that manages your merchant account, which is used to receive funds from credit card companies during settlement. The MSP is also responsible for the processor setup which lets Authorize.Net indirectly access your merchant accounts.",
		"other_suggestions" : "The transaction ID sent in was properly formatted but the gateway had no record of the transaction."
	},
	{
		"code" : "17",
		"text" : "The merchant does not accept this type of credit card.",
		"integration_suggestions" : "",
		"other_suggestions" : "The merchant was not configured to accept the credit card submitted in the transaction."
	},
	{
		"code" : "18",
		"text" : "ACH transactions are not accepted by this merchant.",
		"integration_suggestions" : "",
		"other_suggestions" : "The merchant does not accept electronic checks."
	},
	{
		"code" : "19",
		"text" : "An error occurred during processing.  Please try again.",
		"integration_suggestions" : "This error message is displayed when there is a connection issue between Authorize.Net and the credit card processor. It results from Authorize.Net not receiving data in response to the transaction request we sent to the credit card processor. This type of issue is usually fixed quickly and we continue to work towards eliminating these types of connectivity issues. In some cases it may also be due to Internet congestion, and not related to either of our systems.&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;Repeated attempts to process a transaction during this connectivity breakdown may result in multiple authorizations to the credit card. To prevent this from happening, you can use the following test card:&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;Test Visa Card Number: 4012888818888&amp;lt;br /&amp;gt;Expiration Date: 04/10&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;When a decline message appears for this card in Virtual Terminal, rather than the &amp;quot;Try again in 5 minutes&amp;quot; error message, this means the connectivity problem has been resolved and transactions can be processed again normally. ",
		"other_suggestions" : ""
	},
	{
		"code" : "20",
		"text" : "An error occurred during processing.  Please try again.",
		"integration_suggestions" : "This error message is displayed when there is a connection issue between Authorize.Net and the credit card processor. It results from Authorize.Net not receiving data in response to the transaction request we sent to the credit card processor. This type of issue is usually fixed quickly and we continue to work towards eliminating these types of connectivity issues. In some cases it may also be due to Internet congestion, and not related to either of our systems.&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;Repeated attempts to process a transaction during this connectivity breakdown may result in multiple authorizations to the credit card. To prevent this from happening, you can use the following test card:&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;Test Visa Card Number: 4012888818888&amp;lt;br /&amp;gt;Expiration Date: 04/10&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;When a decline message appears for this card in Virtual Terminal, rather than the &amp;quot;Try again in 5 minutes&amp;quot; error message, this means the connectivity problem has been resolved and transactions can be processed again normally. ",
		"other_suggestions" : ""
	},
	{
		"code" : "21",
		"text" : "An error occurred during processing.  Please try again.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "22",
		"text" : "An error occurred during processing.  Please try again.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "23",
		"text" : "An error occurred during processing.  Please try again.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "24",
		"text" : "The Elavon bank number or terminal ID is incorrect. Call Merchant Service Provider.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "25",
		"text" : "An error occurred during processing.  Please try again.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "26",
		"text" : "An error occurred during processing.  Please try again.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "27",
		"text" : "The transaction has been declined because of an AVS mismatch. The address provided does not match billing address of cardholder.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "28",
		"text" : "The merchant does not accept this type of credit card.",
		"integration_suggestions" : "",
		"other_suggestions" : "The Merchant ID at the processor was not configured to accept this card type."
	},
	{
		"code" : "29",
		"text" : "The Paymentech identification numbers are incorrect. Call Merchant Service Provider.",
		"integration_suggestions" : "",
		"other_suggestions" : "Invalid Paymentech client number, merchant number or terminal number."
	},
	{
		"code" : "30",
		"text" : "The configuration with processor is invalid. Call Merchant Service Provider.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "31",
		"text" : "The FDC Merchant ID or Terminal ID is incorrect. Call Merchant Service Provider.",
		"integration_suggestions" : "",
		"other_suggestions" : "The merchant was incorrectly set up at the processor."
	},
	{
		"code" : "32",
		"text" : "The merchant password is invalid or not present.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "33",
		"text" : "%s cannot be left blank.",
		"integration_suggestions" : "The FIELD is set as Required in the Merchant Interface and the FIELD is not being sent to the gateway.  To change the FIELD value to NOT REQUIRED: Login to the Merchant Interface ( &amp;lt;a href&amp;#x3D;&amp;#x27;https://secure.authorize.net&amp;#x27; target&amp;#x3D;&amp;#x27;_blank&amp;#x27;&amp;gt;https://secure.authorize.net&amp;lt;/a&amp;gt; ), click on Settings and Profile -&amp;gt; Payment Form -&amp;gt; Form Fields and uncheck the Required box for the particular FIELD.",
		"other_suggestions" : "The word FIELD will be replaced by an actual field name. This error indicates that a field the merchant specified as required was not filled in."
	},
	{
		"code" : "34",
		"text" : "The VITAL identification numbers are incorrect. Call Merchant Service Provider.",
		"integration_suggestions" : "",
		"other_suggestions" : "The merchant was incorrectly set up at the processor."
	},
	{
		"code" : "35",
		"text" : "An error occurred during processing. Call Merchant Service Provider.",
		"integration_suggestions" : "",
		"other_suggestions" : "The merchant was incorrectly set up at the processor."
	},
	{
		"code" : "36",
		"text" : "The authorization was approved but settlement failed.",
		"integration_suggestions" : "",
		"other_suggestions" : "The customer was approved at the time of authorization, but failed at settlement."
	},
	{
		"code" : "37",
		"text" : "The credit card number is invalid.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "38",
		"text" : "The Global Payment System identification numbers are incorrect. Call Merchant Service Provider.",
		"integration_suggestions" : "The merchant has invalid Global Payment System NDS numbers.",
		"other_suggestions" : "The merchant was incorrectly set up at the processor."
	},
	{
		"code" : "39",
		"text" : "The supplied currency code is either invalid, not supported, not allowed for this merchant or doesnt have an exchange rate.",
		"integration_suggestions" : "The supplied currency code does not match the currency configured for the merchant account, is invalid, is not supported or not allowed for this merchant.",
		"other_suggestions" : ""
	},
	{
		"code" : "40",
		"text" : "This transaction must be encrypted.",
		"integration_suggestions" : "Insecure transaction.",
		"other_suggestions" : ""
	},
	{
		"code" : "41",
		"text" : "This transaction has been declined.",
		"integration_suggestions" : "",
		"other_suggestions" : "Only merchants set up for the FraudScreen.Net service would receive this decline. This code will be returned if a given transactionÃ¢â‚¬â„¢s fraud score is higher than the threshold set by the merchant."
	},
	{
		"code" : "42",
		"text" : "There is missing or invalid information in a required field.",
		"integration_suggestions" : "",
		"other_suggestions" : "This is applicable only to merchants processing through the Wells Fargo SecureSource product who have requirements for transaction submission that are different from merchants not processing through Wells Fargo."
	},
	{
		"code" : "43",
		"text" : "The merchant was incorrectly set up at the processor. Call Merchant Service Provider.",
		"integration_suggestions" : "The merchant has an invalid Terminal ID.",
		"other_suggestions" : "The merchant was incorrectly set up at the processor."
	},
	{
		"code" : "44",
		"text" : "This transaction has been declined.",
		"integration_suggestions" : "Regardless of the Card Code Verification filter settings configured for the payment gateway account in the Merchant Interface, the processor may decline transactions submitted with a card code value that does not match the card code on file for the cardholder at the issuing bank. To avoid unnecessary errors when processing live transactions, only valid card code values should be submitted in the card code field (&amp;lt;i&amp;gt;x_card_code&amp;lt;/i&amp;gt;). If the merchant does not wish to submit card code information, the card code field should not be submitted.",
		"other_suggestions" : "The card code submitted with the transaction did not match the card code on file at the card issuing bank and the transaction was declined."
	},
	{
		"code" : "45",
		"text" : "This transaction has been declined.",
		"integration_suggestions" : "",
		"other_suggestions" : "This error would be returned if the transaction received a code from the processor that matched the rejection criteria set by the merchant for both the AVS and Card Code filters."
	},
	{
		"code" : "46",
		"text" : "Your session has expired or does not exist. You must log in again to continue working.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "47",
		"text" : "The amount requested for settlement cannot be greater than the original amount authorized.",
		"integration_suggestions" : "",
		"other_suggestions" : "This occurs if the merchant tries to capture funds greater than the amount of the original authorization-only transaction."
	},
	{
		"code" : "48",
		"text" : "This processor does not accept partial reversals.",
		"integration_suggestions" : "",
		"other_suggestions" : "The merchant attempted to settle for less than the originally authorized amount."
	},
	{
		"code" : "49",
		"text" : "The transaction amount submitted was greater than the maximum amount allowed.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "50",
		"text" : "This transaction is awaiting settlement and cannot be refunded.",
		"integration_suggestions" : "",
		"other_suggestions" : "Credits or refunds may only be performed against settled transactions. The transaction against which the credit/refund was submitted has not been settled, so a credit cannot be issued."
	},
	{
		"code" : "51",
		"text" : "The sum of all credits against this transaction is greater than the original transaction amount.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "52",
		"text" : "The transaction was authorized but the client could not be notified; it will not be settled.",
		"integration_suggestions" : "When Authorize.Net is responding back to a script on your server, our system waits up to 10 seconds for a response. If we do not get a response in 10 seconds, our server will time out and display an error page. The first thing that you will need to look for is the order that your script executes. It is very important that something is printed to the screen before any other process is started. If your script prints to the screen first, we will recognize that you are receiving the information. The most effective method would be to PRINT the headers, and a line of text such as &amp;quot;Processing, please wait.&amp;quot; &amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;To resolve this issue:&amp;lt;br /&amp;gt;&amp;lt;ul class&amp;#x3D;&amp;quot;graydot&amp;quot;&amp;gt;&amp;lt;li&amp;gt;Check that your script permissions are correct and that it can accept an HTTPS POST.&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;Check that the script is not completing other functions before writing to the screen, such as writing to a database or sending emails.&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;Please check to see if there are different processes that are used in your script for approvals, declines, or errors. Check each process to be sure that they will write to the screen before any other functions.&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;Check if your script is using redirects immediately upon receipt of the response from our servers. Redirects are discouraged because they can potentially interfere with the process.&amp;lt;/li&amp;gt;&amp;lt;/ul&amp;gt;&amp;lt;p&amp;gt;On occasion, timeouts will occur that are outside of the control of your script or our servers. Typical reasons for these timeouts include Internet traffic, your server is overloaded or malfunctioning, or Internet routing issues. Depending upon your server location and what route is used to send data, it is possible that you may occasionally receive the message you are seeing.&amp;lt;/p&amp;gt;",
		"other_suggestions" : ""
	},
	{
		"code" : "53",
		"text" : "The transaction type is invalid for ACH transactions.",
		"integration_suggestions" : "The &amp;lt;i&amp;gt;x_type&amp;lt;/i&amp;gt; value is not valid for ACH transactions.",
		"other_suggestions" : "If &amp;lt;i&amp;gt;x_method&amp;lt;/i&amp;gt; &amp;#x3D; ECHECK, &amp;lt;i&amp;gt;x_type&amp;lt;/i&amp;gt; cannot be set to CAPTURE_ONLY."
	},
	{
		"code" : "54",
		"text" : "The referenced transaction does not meet the criteria for issuing a credit.",
		"integration_suggestions" : "The referenced transaction does not meet the criteria for issuing a credit. It may be unsettled, an invalid type, the wrong currency, an invalid reference transaction ID or settled more than 120 days ago.",
		"other_suggestions" : ""
	},
	{
		"code" : "55",
		"text" : "The sum of credits against the referenced transaction would exceed original debit amount.",
		"integration_suggestions" : "",
		"other_suggestions" : "The transaction is rejected if the sum of this credit and prior credits exceeds the original debit amount."
	},
	{
		"code" : "56",
		"text" : "Credit card transactions are not accepted by this merchant.",
		"integration_suggestions" : "",
		"other_suggestions" : "The merchant processes eCheck.Net transactions only and does not accept credit cards."
	},
	{
		"code" : "57",
		"text" : "An error occurred during processing.  Please try again.",
		"integration_suggestions" : "This error is caused when a transaction is submitted with data that the credit card processor does not recognize or is unable to interpret. In most cases our system will prevent this from happening with front-end safeguards, but since every processor is unique in the way they handle data, some transactions could get through to the processor with invalid or missing data. Examples of these types of discrepancies include placing the incorrect number of characters in the Card Verification Value (Card Code), or sending non-alphanumeric characters in the Zip Code.",
		"other_suggestions" : ""
	},
	{
		"code" : "58",
		"text" : "An error occurred during processing.  Please try again.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "59",
		"text" : "An error occurred during processing.  Please try again.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "60",
		"text" : "An error occurred during processing.  Please try again.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "61",
		"text" : "An error occurred during processing.  Please try again.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "62",
		"text" : "An error occurred during processing.  Please try again.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "63",
		"text" : "An error occurred during processing.  Please try again.",
		"integration_suggestions" : "If you receive an Error 63 repeatedly, please check that the Merchant Business country is set correctly. This is especially pertinent on accounts which use TSYS (formerly Vital or Visanet) as the payment processor, as we have to transmit a number of the Business Information fields on each Transaction attempt. TSYS/Vital/Visanet transactions will fail if all the information is not set correctly.&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;To update the Business Information details for your account:&amp;lt;ul class&amp;#x3D;&amp;quot;graydot&amp;quot;&amp;gt;&amp;lt;li&amp;gt;Login to the Merchant Interface at &amp;lt;a href&amp;#x3D;&amp;quot;https://secure.authorize.net/&amp;quot;&amp;gt;https://secure.authorize.net/&amp;lt;/a&amp;gt;&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;Click &amp;lt;strong&amp;gt;Merchant Profile&amp;lt;/strong&amp;gt; in the main left side menu&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;Click &amp;lt;strong&amp;gt;Edit Business Information&amp;lt;/strong&amp;gt;&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;Update your Address, Phone Number, Fax Number, Products/Services Description, Web Site Address, and Shopping Cart Solution as necessary.&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;Click &amp;lt;strong&amp;gt;Submit&amp;lt;/strong&amp;gt;&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;Click &amp;lt;strong&amp;gt;Continue&amp;lt;/strong&amp;gt; to return to the Merchant Profile page.&amp;lt;/li&amp;gt;&amp;lt;/ul&amp;gt;&amp;lt;p&amp;gt;If the Merchant Profile has your correct address-including the country-then you may need to contact your Merchant Service Provider to confirm that TSYS/Vital/Visanet has the correct address for validation.",
		"other_suggestions" : ""
	},
	{
		"code" : "64",
		"text" : "The referenced transaction was not approved.",
		"integration_suggestions" : "",
		"other_suggestions" : "This error is applicable to Wells Fargo SecureSource merchants only. Credits or refunds cannot be issued against transactions that were not authorized."
	},
	{
		"code" : "65",
		"text" : "This transaction has been declined.",
		"integration_suggestions" : "",
		"other_suggestions" : "The transaction was declined because the merchant configured their account through the Merchant Interface to reject transactions with certain values for a Card Code mismatch."
	},
	{
		"code" : "66",
		"text" : "This transaction cannot be accepted for processing.",
		"integration_suggestions" : "This error can be caused by the following situations:If you are using the SIM connection method, make sure your code is providing values for the SIM required fields listed below:&amp;lt;ul class&amp;#x3D;&amp;quot;graydot&amp;quot;&amp;gt;&amp;lt;li&amp;gt;The sequence number of the transaction (&amp;lt;i&amp;gt;x_fp_sequence&amp;lt;/i&amp;gt;)&amp;lt;/li&amp;gt; &amp;lt;li&amp;gt;The time when the sequence number was generated (&amp;lt;i&amp;gt;x_fp_timestamp&amp;lt;/i&amp;gt;)&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;The Fingerprint Hash (&amp;lt;i&amp;gt;x_fp_hash&amp;lt;/i&amp;gt;)&amp;lt;/li&amp;gt;&amp;lt;/ul&amp;gt;&amp;lt;p&amp;gt;If you are using the Weblink connection method and have not specified any valid browser referrer URLs in your online account settings, that will also cause this particular error. Merchants using Weblink will want to specify URLs in the Referrer URL portion of their account settings for security purposes. Doing so would ensure that only the URLs specified therein can post transactions to your account. This method is applicable because the browser moves from the merchant&amp;#x27;s website to Authorize.Net&amp;#x27;s secure payment form. Typically the URL you see just before our payment form is the referrer URL. If you have multiple sites, there may be more than one referrer URL. You do not have to enter the entire URL if it has variables in it.&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;To implement Valid Browser Referrer URL:&amp;lt;ul class&amp;#x3D;&amp;quot;graydot&amp;quot;&amp;gt;&amp;lt;li&amp;gt;Login to the Merchant Interface at &amp;lt;a href&amp;#x3D;&amp;quot;https://account.authorize.net/&amp;quot;&amp;gt;https://account.authorize.net/&amp;lt;/a&amp;gt;&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;Click &amp;lt;strong&amp;gt;Settings&amp;lt;/strong&amp;gt; in the main left side menu&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;Click &amp;lt;strong&amp;gt;Weblink&amp;lt;/strong&amp;gt;&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;Click &amp;lt;strong&amp;gt;Referrer URLs&amp;lt;/strong&amp;gt;&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;Click &amp;lt;strong&amp;gt;Add URL&amp;lt;/strong&amp;gt;&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;Enter the URL that your site shows before customers are directed to our payment form&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;Click &amp;lt;strong&amp;gt;Submit&amp;lt;/strong&amp;gt;&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;Go to your website and process a test transaction to confirm that the error has stopped&amp;lt;/li&amp;gt;&amp;lt;/ul&amp;gt;&amp;lt;p&amp;gt;&amp;lt;strong&amp;gt;Note&amp;lt;/strong&amp;gt;: Weblink is no longer actively supported and is subject to be discontinued in the near future. We strongly recommend that Weblink merchants upgrade to SIM or AIM. For details about upgrading to AIM or SIM, please read the Weblink Conversion Guide at &amp;lt;a href&amp;#x3D;&amp;quot;http://www.authorize.net/files/conversionguide.pdf&amp;quot; target&amp;#x3D;&amp;quot;_blank&amp;quot;&amp;gt;http://www.authorize.net/files/conversionguide.pdf&amp;lt;/a&amp;gt;.",
		"other_suggestions" : "The transaction did not meet payment gateway security guidelines."
	},
	{
		"code" : "67",
		"text" : "The given transaction type is not supported for this merchant.",
		"integration_suggestions" : "",
		"other_suggestions" : "This error code is applicable to merchants using the Wells Fargo SecureSource product only. This product does not allow transactions of type CAPTURE_ONLY."
	},
	{
		"code" : "68",
		"text" : "The version parameter is invalid",
		"integration_suggestions" : "",
		"other_suggestions" : "The value submitted in &amp;lt;i&amp;gt;x_version&amp;lt;/i&amp;gt; was invalid."
	},
	{
		"code" : "69",
		"text" : "The transaction type is invalid",
		"integration_suggestions" : "",
		"other_suggestions" : "The value submitted in &amp;lt;i&amp;gt;x_type&amp;lt;/i&amp;gt; was invalid."
	},
	{
		"code" : "70",
		"text" : "The transaction method is invalid.",
		"integration_suggestions" : "",
		"other_suggestions" : "The value submitted in &amp;lt;i&amp;gt;x_method&amp;lt;/i&amp;gt; was invalid."
	},
	{
		"code" : "71",
		"text" : "The bank account type is invalid.",
		"integration_suggestions" : "",
		"other_suggestions" : "The value submitted in &amp;lt;i&amp;gt;x_bank_acct_type&amp;lt;/i&amp;gt; was invalid."
	},
	{
		"code" : "72",
		"text" : "The authorization code is invalid.",
		"integration_suggestions" : "",
		"other_suggestions" : "The value submitted in &amp;lt;i&amp;gt;x_auth_code&amp;lt;/i&amp;gt; was more than six characters in length."
	},
	{
		"code" : "73",
		"text" : "The drivers license date of birth is invalid.",
		"integration_suggestions" : "",
		"other_suggestions" : "The format of the value submitted in &amp;lt;i&amp;gt;x_drivers_license_num&amp;lt;/i&amp;gt; was invalid."
	},
	{
		"code" : "74",
		"text" : "The duty amount is invalid.",
		"integration_suggestions" : "",
		"other_suggestions" : "The value submitted in &amp;lt;i&amp;gt;x_duty&amp;lt;/i&amp;gt; failed format validation."
	},
	{
		"code" : "75",
		"text" : "The freight amount is invalid.",
		"integration_suggestions" : "",
		"other_suggestions" : "The value submitted in &amp;lt;i&amp;gt;x_freight&amp;lt;/i&amp;gt; failed format validation."
	},
	{
		"code" : "76",
		"text" : "The tax amount is invalid.",
		"integration_suggestions" : "",
		"other_suggestions" : "The value submitted in &amp;lt;i&amp;gt;x_tax&amp;lt;/i&amp;gt; failed format validation."
	},
	{
		"code" : "77",
		"text" : "The SSN or tax ID is invalid.",
		"integration_suggestions" : "",
		"other_suggestions" : "The value submitted in &amp;lt;i&amp;gt;x_customer_tax_id&amp;lt;/i&amp;gt; failed validation."
	},
	{
		"code" : "78",
		"text" : "The card code is invalid.",
		"integration_suggestions" : "",
		"other_suggestions" : "The value submitted in &amp;lt;i&amp;gt;x_card_code&amp;lt;/i&amp;gt; failed format validation."
	},
	{
		"code" : "79",
		"text" : "The drivers license number is invalid.",
		"integration_suggestions" : "",
		"other_suggestions" : "The value submitted in &amp;lt;i&amp;gt;x_drivers_license_num&amp;lt;/i&amp;gt; failed format validation."
	},
	{
		"code" : "80",
		"text" : "The drivers license state is invalid.",
		"integration_suggestions" : "",
		"other_suggestions" : "The value submitted in &amp;lt;i&amp;gt;x_drivers_license_state&amp;lt;/i&amp;gt; failed format validation."
	},
	{
		"code" : "81",
		"text" : "The requested form type is invalid.",
		"integration_suggestions" : "Invalid value for &amp;lt;i&amp;gt;x_showform&amp;lt;/i&amp;gt;.",
		"other_suggestions" : "The merchant requested an integration method not compatible with the AIM API."
	},
	{
		"code" : "82",
		"text" : "Scripts are only supported in version 2.5.",
		"integration_suggestions" : "",
		"other_suggestions" : "The system no longer supports version 2.5; requests cannot be posted to scripts."
	},
	{
		"code" : "83",
		"text" : "The requested script is either invalid or no longer supported.",
		"integration_suggestions" : "",
		"other_suggestions" : "The system no longer supports version 2.5; requests cannot be posted to scripts."
	},
	{
		"code" : "84",
		"text" : "The device type is invalid or missing.",
		"integration_suggestions" : "deviceType is either missing or invalid. deviceType is required when the retail element is submitted.",
		"other_suggestions" : "Invalid value for &amp;lt;i&amp;gt;x_device_type&amp;lt;/i&amp;gt;."
	},
	{
		"code" : "85",
		"text" : "The market type is invalid",
		"integration_suggestions" : "marketType is either missing or invalid. marketType is required when the retail element is submitted.",
		"other_suggestions" : "Invalid value for &amp;lt;i&amp;gt;x_market_type&amp;lt;/i&amp;gt;."
	},
	{
		"code" : "86",
		"text" : "The Response Format is invalid",
		"integration_suggestions" : "",
		"other_suggestions" : "Invalid value for &amp;lt;i&amp;gt;x_response_format&amp;lt;/i&amp;gt;."
	},
	{
		"code" : "87",
		"text" : "Transactions of this market type cannot be processed on this system.",
		"integration_suggestions" : "This can happen for four reasons:&amp;lt;br /&amp;gt;1) You are attempting to process a type of transaction that your account is not designed for.  An example would be using a card swipe machine to process transactions on an e-commerce or mail order/telephone order (MOTO) account.&amp;lt;br /&amp;gt;2) You are passing an incorrect value to &amp;lt;i&amp;gt;x_market_type&amp;lt;/i&amp;gt;.  If you are using a Retail account you want to pass a &amp;quot;2&amp;quot;. The Implementation Guide for Card Present transactionscan be found at &amp;lt;a href&amp;#x3D;&amp;quot;http://www.authorize.net/support/CP_guide.pdf&amp;quot;&amp;gt;http://www.authorize.net/support/CP_guide.pdf&amp;lt;/a&amp;gt;.&amp;lt;br /&amp;gt;3) Your merchant service provider may be incorrectly setup for this account.&amp;lt;br /&amp;gt;4) The product type (SIC code) is incorrect for this account.",
		"other_suggestions" : "The server is not configured to process the product type derived from the market type of this transaction."
	},
	{
		"code" : "88",
		"text" : "Track1 data is not in a valid format.",
		"integration_suggestions" : "",
		"other_suggestions" : "Invalid value for &amp;lt;i&amp;gt;x_track1&amp;lt;/i&amp;gt;."
	},
	{
		"code" : "89",
		"text" : "Track2 data is not in a valid format.",
		"integration_suggestions" : "",
		"other_suggestions" : "Invalid value for &amp;lt;i&amp;gt;x_track2&amp;lt;/i&amp;gt;."
	},
	{
		"code" : "90",
		"text" : "ACH transactions cannot be accepted by this system.",
		"integration_suggestions" : "ACH transactions cannot be processed using a card present account.",
		"other_suggestions" : ""
	},
	{
		"code" : "91",
		"text" : "Version 2.5 is no longer supported.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "92",
		"text" : "The gateway no longer supports the requested method of integration.",
		"integration_suggestions" : "This error can occur for several possible reasons, depending on which Application Programming Interface (API) your software uses to connect to your Authorize.Net account. Usually it is due to mixing methods between our APIs.&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;&amp;lt;u&amp;gt;Advanced Integration Method (AIM) Suggestions&amp;lt;/u&amp;gt;:&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;1. If you are posting via AIM to our secure site directly from an HTML page you will receive this error. When using AIM to integrate with Authorize.Net, the post must be made from a script running in a secure location on your server, and not through an HTML page.&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;2. Be sure that your account settings are set to allow delimited responses.&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;To verify or edit this you can:&amp;lt;ul class&amp;#x3D;&amp;quot;graydot&amp;quot;&amp;gt;&amp;lt;li&amp;gt;Login to the Merchant Interface at &amp;lt;a href&amp;#x3D;&amp;quot;https://account.authorize.net/&amp;quot;&amp;gt;https://account.authorize.net/&amp;lt;/a&amp;gt;&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;Click on &amp;lt;strong&amp;gt;Settings&amp;lt;/strong&amp;gt;&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;Click on &amp;lt;strong&amp;gt;Direct Response&amp;lt;/strong&amp;gt;&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;Change Delimited Response to &amp;amp;quot;True&amp;amp;quot;&amp;lt;/li&amp;gt;&amp;lt;/ul&amp;gt;&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;3. Confirm that you are posting the variable &amp;lt;i&amp;gt;x_delim_data&amp;lt;/i&amp;gt; in your script, and that it is set to &amp;quot;TRUE&amp;quot;.&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;4. Also confirm that you are sending the variable &amp;lt;i&amp;gt;x_relay_response&amp;lt;/i&amp;gt; with a value equal to &amp;quot;false&amp;quot; in your post to Authorize.Net. If you are not sending this variable and have any default relay or receipt links listed in your &amp;lt;strong&amp;gt;Response/Receipt URL&amp;lt;/strong&amp;gt; settings, you will receive an Error (92).&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;&amp;lt;u&amp;gt;Simple Integration Method (SIM) Suggestions&amp;lt;/u&amp;gt;:&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;1. You will receive this error if there are variables being sent that are not applicable to SIM. Two of the variables that are most commonly incorrect include:&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;&amp;lt;i&amp;gt;x_adc_relay_response&amp;lt;/i&amp;gt; - the variable name should be sent as &amp;lt;i&amp;gt;x_relay_response&amp;lt;/i&amp;gt;&amp;lt;br /&amp;gt;&amp;lt;i&amp;gt;x_adc_relay_url&amp;lt;/i&amp;gt; - the variable name should be sent as &amp;lt;i&amp;gt;x_relay_url&amp;lt;/i&amp;gt;.&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;2. You will receive this error if the proper fingerprint hash variables are not being sent with the transaction request. The variables that need to be included (with appropriate values):&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;&amp;lt;i&amp;gt;x_fp_hash&amp;lt;/i&amp;gt;&amp;lt;br /&amp;gt;&amp;lt;i&amp;gt;x_fp_sequence&amp;lt;/i&amp;gt;&amp;lt;br /&amp;gt;&amp;lt;i&amp;gt;x_fp_timestamp&amp;lt;/i&amp;gt;&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;These variables are used in creating and then passing the fingerprint hash with the transaction request. Authorize.Net then uses the passed variables and the stored transaction key to attempt to create the same fingerprint hash. If the two fingerprints match, we accept the transaction. If not, the transaction request is refused.&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;&amp;lt;u&amp;gt;Weblink Integration Suggestions&amp;lt;/u&amp;gt;:&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;You will receive this error if Weblink has been disabled on your account in error. When logging into your Authorize.net account you may be prompted to disable Weblink. However this should only be disabled if you have already successfully upgraded to SIM or AIM. If you have not successfully upgraded, and have disabled the Weblink connection method then transactions from your website will return Error 92.&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;If you have disabled Weblink in error, please call Merchant Support at 877-447-3938 so we may re-enable the option while you complete your upgrade to SIM or AIM.&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;For documentation on SIM and AIM please visit &amp;lt;a href&amp;#x3D;&amp;quot;http://developer.authorize.net/api/&amp;quot;&amp;gt;http://developer.authorize.net/api/&amp;lt;/a&amp;gt;.",
		"other_suggestions" : ""
	},
	{
		"code" : "93",
		"text" : "A valid country is required.",
		"integration_suggestions" : "",
		"other_suggestions" : "This code is applicable to Wells Fargo SecureSource merchants only. Country is a required field and must contain the value of a supported country."
	},
	{
		"code" : "94",
		"text" : "The shipping state or country is invalid.",
		"integration_suggestions" : "",
		"other_suggestions" : "This code is applicable to Wells Fargo SecureSource merchants only."
	},
	{
		"code" : "95",
		"text" : "A valid state is required.",
		"integration_suggestions" : "",
		"other_suggestions" : "This code is applicable to Wells Fargo SecureSource merchants only."
	},
	{
		"code" : "96",
		"text" : "This country is not authorized for buyers.",
		"integration_suggestions" : "",
		"other_suggestions" : "This code is applicable to Wells Fargo SecureSource merchants only. Country is a required field and must contain the value of a supported country."
	},
	{
		"code" : "97",
		"text" : "This transaction cannot be accepted.",
		"integration_suggestions" : "Please use the &amp;lt;a href&amp;#x3D;&amp;quot;http://developer.authorize.net/api/reference/responseCode97.html&amp;quot;&amp;gt;Response Code 97 Tool&amp;lt;/a&amp;gt;",
		"other_suggestions" : "Applicable only to the SIM API. Fingerprints are only valid for a short period of time. This code indicates that the transaction fingerprint has expired."
	},
	{
		"code" : "98",
		"text" : "This transaction cannot be accepted.",
		"integration_suggestions" : "Part of the security feature for SIM (Server Integration Method) includes the generation of a fingerprint hash. The fingerprint hash is generated by a function in the scripting that uses five parameters to generate the fingerprint hash.&amp;lt;ul class&amp;#x3D;&amp;quot;graydot&amp;quot;&amp;gt;&amp;lt;li&amp;gt;The amount of the transaction;&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;A sequence number&amp;amp;#151;usually an invoice number generated by your scripting, but can be randomly generated as long as it doesn&amp;#x27;t repeat;&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;Your server&amp;#x27;s timestamp, expressed in Greenwich Mean Time (GMT) or Coordinated Universal Time (UTC);&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;Your account&amp;#x27;s Transaction Key; and&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;Your account&amp;#x27;s API Login ID.&amp;lt;/li&amp;gt;&amp;lt;/ul&amp;gt;&amp;lt;p&amp;gt;Any fingerprint hash can only be used once in 24 hours.&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;If a customer fills in incorrect information and the transaction is declined, they cannot click Back and re-enter the information as this will attempt to use the same fingerprint hash in resubmitting the transaction request, which will result in error 98.&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;The customer must be directed back to the page that sets the amount of the transaction request, then re-enter information from that point on.",
		"other_suggestions" : "Applicable only to the SIM API. The transaction fingerprint has already been used."
	},
	{
		"code" : "99",
		"text" : "This transaction cannot be accepted.",
		"integration_suggestions" : "Please use the &amp;lt;a href&amp;#x3D;&amp;quot;http://developer.authorize.net/api/reference/responseCode99.html&amp;quot;&amp;gt;Response Code 99 Tool&amp;lt;/a&amp;gt;.",
		"other_suggestions" : "Applicable only to the SIM API. The server-generated fingerprint does not match the merchant-specified fingerprint in the x_fp_hash field."
	},
	{
		"code" : "100",
		"text" : "The eCheck type parameter is invalid.",
		"integration_suggestions" : "",
		"other_suggestions" : "Applicable only to eCheck.Net. The value specified in the &amp;lt;i&amp;gt;x_echeck_type&amp;lt;/i&amp;gt; field is invalid."
	},
	{
		"code" : "101",
		"text" : "The given name on the account and/or the account type does not match the actual account.",
		"integration_suggestions" : "",
		"other_suggestions" : "Applicable only to eCheck.Net. The specified name on the account and/or the account type do not match the NOC record for this account."
	},
	{
		"code" : "102",
		"text" : "This request cannot be accepted.",
		"integration_suggestions" : "If you are receiving this error, it indicates that either &amp;lt;i&amp;gt;x_password&amp;lt;/i&amp;gt; or &amp;lt;i&amp;gt;x_tran_key&amp;lt;/i&amp;gt; is being submitted with your Weblink request. This represents a security risk as the password or transaction key could be viewed in your source code.&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;We highly recommend that you upgrade your account to one of the newer methods of linking, the Server Integration Method (SIM) or the Advanced Integration Method (AIM).&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;To locate documentation on SIM and AIM please visit &amp;lt;a href&amp;#x3D;&amp;quot;http://developer.authorize.net/guides/&amp;quot;&amp;gt;http://developer.authorize.net/guides/&amp;lt;/a&amp;gt;.",
		"other_suggestions" : "A transaction key was submitted with this WebLink request."
	},
	{
		"code" : "103",
		"text" : "This transaction cannot be accepted.",
		"integration_suggestions" : "This error is generated when your account is in &amp;quot;Password-Required Mode&amp;quot; and you are not sending a valid password, transaction key or hash fingerprint with your transaction request, which is a recommended security measure.&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;Please consider the following details when encountering this error:&amp;lt;ul class&amp;#x3D;&amp;quot;graydot&amp;quot;&amp;gt;&amp;lt;li&amp;gt;If you are using SIM, make sure you are using a valid transaction key to generate and send a fingerprint hash to us along with your transaction request.&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;If you are using AIM please make sure you are posting the gateway defined field &amp;lt;i&amp;gt;x_tran_key&amp;lt;/i&amp;gt;.&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;If you are using a third party shopping cart which uses AIM and are receiving this error, please check with your shopping cart provider to ask if your application can pass the transaction key to the Authorize.Net payment gateway.&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;If you are using an older Authorize.Net account and you submit a password instead of a transaction key with your transactions, you may experience this error. Please ensure that you are posting a transaction key instead.&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;Some shopping carts, for backwards compatibility with older connection methods, may provide the means to submit both a transaction key and a password. You should not use both the transaction key and the password simultaneously; doing so may also result in this error. We recommend using the transaction key instead of the password whenever possible, as transaction keys tend to be more secure than passwords. In such a situation, please leave the password field blank&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;If the shopping cart has a field for the password but no field for the transaction key, please put the transaction key in the password field. Our system will recognize and validate the transaction key properly.&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;Also, while most shopping cart software will have a field for the transaction key, password, or both, some software may not. Please contact your shopping cart provider for details on how to upgrade to a more secure version of your shopping cart software. In an emergency, you may submit transactions without a transaction key, by disabling Password-Required Mode in your account. You may find details in our Password-Required Mode knowledge base document.&amp;lt;/li&amp;gt;&amp;lt;/ul&amp;gt;",
		"other_suggestions" : ""
	},
	{
		"code" : "104",
		"text" : "The transaction is currently under review.",
		"integration_suggestions" : "",
		"other_suggestions" : "Applicable only to eCheck.Net. The value submitted for &amp;lt;i&amp;gt;x_country&amp;lt;/i&amp;gt; failed validation."
	},
	{
		"code" : "105",
		"text" : "The transaction is currently under review.",
		"integration_suggestions" : "",
		"other_suggestions" : "Applicable only to eCheck.Net. The values submitted for &amp;lt;i&amp;gt;x_city&amp;lt;/i&amp;gt; and &amp;lt;i&amp;gt;x_country&amp;lt;/i&amp;gt; failed validation."
	},
	{
		"code" : "106",
		"text" : "The transaction is currently under review.",
		"integration_suggestions" : "",
		"other_suggestions" : "Applicable only to eCheck.Net. The value submitted for &amp;lt;i&amp;gt;x_company&amp;lt;/i&amp;gt; failed validation."
	},
	{
		"code" : "107",
		"text" : "The transaction is currently under review.",
		"integration_suggestions" : "",
		"other_suggestions" : "Applicable only to eCheck.Net. The value submitted for&amp;lt;i&amp;gt; x_bank_acct_name&amp;lt;/i&amp;gt; failed validation."
	},
	{
		"code" : "108",
		"text" : "The transaction is currently under review.",
		"integration_suggestions" : "",
		"other_suggestions" : "Applicable only to eCheck.Net. The values submitted for &amp;lt;i&amp;gt;x_first_name&amp;lt;/i&amp;gt; and &amp;lt;i&amp;gt;x_last_name&amp;lt;/i&amp;gt; failed validation."
	},
	{
		"code" : "109",
		"text" : "The transaction is currently under review.",
		"integration_suggestions" : "",
		"other_suggestions" : "Applicable only to eCheck.Net. The values submitted for &amp;lt;i&amp;gt;x_first_name&amp;lt;/i&amp;gt; and &amp;lt;i&amp;gt;x_last_name&amp;lt;/i&amp;gt; failed validation."
	},
	{
		"code" : "110",
		"text" : "The transaction is currently under review.",
		"integration_suggestions" : "",
		"other_suggestions" : "The value submitted for &amp;lt;i&amp;gt;x_bank_acct_name&amp;lt;/i&amp;gt; does not contain valid characters."
	},
	{
		"code" : "111",
		"text" : "A valid billing country is required.",
		"integration_suggestions" : "",
		"other_suggestions" : "This code is applicable to Wells Fargo SecureSource merchants only."
	},
	{
		"code" : "112",
		"text" : "A valid billing state/province is required.",
		"integration_suggestions" : "",
		"other_suggestions" : "This code is applicable to Wells Fargo SecureSource merchants only."
	},
	{
		"code" : "113",
		"text" : "The commercial card type is invalid.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "114",
		"text" : "The merchant account is in test mode. This automated payment will not be processed.",
		"integration_suggestions" : "The merchant account is in test mode. All automated payments are skipped when in test mode.",
		"other_suggestions" : ""
	},
	{
		"code" : "115",
		"text" : "The merchant account is not active. This automated payment will not be processed.",
		"integration_suggestions" : "The merchant account is not active. All automated payments are skipped when an account is not active.",
		"other_suggestions" : ""
	},
	{
		"code" : "116",
		"text" : "The authentication indicator is invalid.",
		"integration_suggestions" : "",
		"other_suggestions" : "This code is applicable only to merchants that include the &amp;lt;i&amp;gt;x_authentication_indicator&amp;lt;/i&amp;gt; in the transaction request.  The ECI value for a Visa transaction; or the UCAF indicator for a MasterCard transaction submitted in the &amp;lt;i&amp;gt;x_authentication_indicator&amp;lt;/i&amp;gt; field is invalid."
	},
	{
		"code" : "117",
		"text" : "The cardholder authentication value is invalid.",
		"integration_suggestions" : "First, verify that the merchantÃ¢â‚¬â„¢s processor supports Verified by Visa and MasterCard SecureCode authentication values through Authorize.Net.&amp;lt;ul class&amp;#x3D;&amp;quot;graydot&amp;quot;&amp;gt;&amp;lt;li&amp;gt;Chase Paymentech&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;FDMS Nashville (formerly FDC)&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;Global Payments (GPS)&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;TSYS (formerly Vital)&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;Wells Fargo (Verified by Visa only)&amp;lt;/li&amp;gt;&amp;lt;/ul&amp;gt;&amp;lt;p&amp;gt;Also, this error can be received in the event that a special character is included in the cardholder authentication value. To resolve this issue, the special character must be URL encoded.&amp;lt;/p&amp;gt;",
		"other_suggestions" : "This code is applicable only to merchants that include the &amp;lt;i&amp;gt;x_cardholder_authentication_value&amp;lt;/i&amp;gt; in the transaction request. The CAVV for a Visa transaction or the AVV/UCAF for a MasterCard transaction is invalid or contains an invalid character."
	},
	{
		"code" : "118",
		"text" : "The combination of card type, authentication indicator and cardholder authentication value is invalid.",
		"integration_suggestions" : "For example, when the MasterCard value for x_authentication_indicator is 1, the value for x_cardholder_authentication_value must be null. In this scenario, if a value is submitted for x_cardholder_authentication_value, the transaction fails validation and is rejected.&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;More information regarding valid authentication indicator and cardholder authentication value combinations can be found in the Advanced Integration Method (AIM) Developer Guide at &amp;lt;a href&amp;#x3D;&amp;quot;http://www.authorize.net/files/AIM_guide.pdf&amp;quot;&amp;gt;http://www.authorize.net/files/AIM_guide.pdf&amp;lt;/a&amp;gt;.&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;Also, verify that the merchant&amp;#x27;s processor supports Verified by Visa and MasterCard SecureCode authentication values through Authorize.Net. &amp;lt;ul class&amp;#x3D;&amp;quot;graydot&amp;quot;&amp;gt;&amp;lt;li&amp;gt;Chase Paymentech&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;FDMS Nashville (formerly FDC)&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;Global Payments (GPS)&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;TSYS (formerly Vital)&amp;lt;/li&amp;gt;&amp;lt;/ul&amp;gt;",
		"other_suggestions" : "This code is applicable only to merchants that include the x_authentication_indicator and &amp;lt;i&amp;gt;x_authentication_value&amp;lt;/i&amp;gt; in the transaction request. The combination of authentication indicator and cardholder authentication value (Visa or MasterCard) is invalid."
	},
	{
		"code" : "119",
		"text" : "Transactions having cardholder authentication values cannot be marked as recurring.",
		"integration_suggestions" : "Transactions that have Verified by Visa type of input parameters cannot be recurring.",
		"other_suggestions" : "This code is applicable only to merchants that include the &amp;lt;i&amp;gt;x_authentication_indicator&amp;lt;/i&amp;gt; and &amp;lt;i&amp;gt;x_recurring_billing&amp;lt;/i&amp;gt; in the transaction request. Transactions submitted with a value in &amp;lt;i&amp;gt;x_authentication_indicator&amp;lt;/i&amp;gt; AND &amp;lt;i&amp;gt;x_recurring_billing&amp;lt;/i&amp;gt; &amp;#x3D;YES will be rejected."
	},
	{
		"code" : "120",
		"text" : "An error occurred during processing. Please try again.",
		"integration_suggestions" : "",
		"other_suggestions" : "The system-generated void for the original timed-out transaction failed. (The original transaction timed out while waiting for a response from the authorizer.)"
	},
	{
		"code" : "121",
		"text" : "An error occurred during processing. Please try again.",
		"integration_suggestions" : "",
		"other_suggestions" : "The system-generated void for the original errored transaction failed. (The original transaction experienced a database error.)"
	},
	{
		"code" : "122",
		"text" : "An error occurred during processing. Please try again.",
		"integration_suggestions" : "",
		"other_suggestions" : "The system-generated void for the original errored transaction failed. (The original transaction experienced a processing error.)"
	},
	{
		"code" : "123",
		"text" : "This account has not been given the permission(s) required for this request.",
		"integration_suggestions" : "This error indicates that a user&amp;#x27;s personal Login ID is being used to connect a website or billing software to the payment gateway. Personal login IDs may not be used to connect websites to Authorize.Net, for security reasons. For example, if an Account Owner, Account Administrator, Transaction Manager, or Account Analyst login ID is used for website or software implementation, this error will occur.&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;To resolve this issue, the API Login ID and Transaction Key will need to be generated and added to your software&amp;#x27;s configuration, so that the website or software may connect to Authorize.Net properly. The API Login ID can be obtained by following these steps:&amp;lt;ul class&amp;#x3D;&amp;quot;graydot&amp;quot;&amp;gt;&amp;lt;li&amp;gt;Log into the Merchant Interface at &amp;lt;a href&amp;#x3D;&amp;quot;https://account.authorize.net/&amp;quot;&amp;gt;https://account.authorize.net/&amp;lt;/a&amp;gt;&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;Click &amp;lt;strong&amp;gt;Settings&amp;lt;/strong&amp;gt; in the main left side menu&amp;lt;/li&amp;gt; &amp;lt;li&amp;gt;Click &amp;lt;strong&amp;gt;API Login and Transaction Key&amp;lt;/strong&amp;gt;&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;If an API login ID has already been generated, it will be visible here. If an API Login ID needs to be generated, you will want to enter the answer to your &amp;lt;strong&amp;gt;Secret Question&amp;lt;/strong&amp;gt; in order to see the API Login ID.&amp;lt;/li&amp;gt;&amp;lt;/ul&amp;gt;&amp;lt;p&amp;gt;For more information regarding Login IDs, Passwords, API Login IDs, Transaction Keys, and User Administration, please visit our online video tutorials at &amp;lt;a href&amp;#x3D;&amp;quot;http://www.authorize.net/videos/&amp;quot;&amp;gt;http://www.authorize.net/videos/&amp;lt;/a&amp;gt;.",
		"other_suggestions" : "The transaction request must include the API login ID associated with the payment gateway account."
	},
	{
		"code" : "124",
		"text" : "This processor does not accept recurring transactions.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "127",
		"text" : "The transaction resulted in an AVS mismatch. The address provided does not match billing address of cardholder.",
		"integration_suggestions" : "",
		"other_suggestions" : "The system-generated void for the original AVS-rejected transaction failed."
	},
	{
		"code" : "128",
		"text" : "This transaction cannot be processed.",
		"integration_suggestions" : "",
		"other_suggestions" : "The customer&amp;#x27;s financial institution does not currently allow transactions for this account."
	},
	{
		"code" : "130",
		"text" : "This merchant account has been closed.",
		"integration_suggestions" : "",
		"other_suggestions" : "The payment gateway account status is Blacklisted."
	},
	{
		"code" : "131",
		"text" : "This transaction cannot be accepted at this time.",
		"integration_suggestions" : "",
		"other_suggestions" : "The payment gateway account status is Suspended-STA."
	},
	{
		"code" : "132",
		"text" : "This transaction cannot be accepted at this time.",
		"integration_suggestions" : "",
		"other_suggestions" : "The payment gateway account status is Suspended - Blacklist."
	},
	{
		"code" : "141",
		"text" : "This transaction has been declined.",
		"integration_suggestions" : "",
		"other_suggestions" : "The system-generated void for the original FraudScreen-rejected transaction failed."
	},
	{
		"code" : "145",
		"text" : "This transaction has been declined.",
		"integration_suggestions" : "",
		"other_suggestions" : "The system-generated void for the original card code-rejected and AVS-rejected transaction failed."
	},
	{
		"code" : "152",
		"text" : "The transaction was authorized but the client could not be notified; it will not be settled.",
		"integration_suggestions" : "",
		"other_suggestions" : "The system-generated void for the original transaction failed. The response for the original transaction could not be communicated to the client."
	},
	{
		"code" : "153",
		"text" : "There was an error processing the payment data.",
		"integration_suggestions" : "",
		"other_suggestions" : "&amp;lt;ul&amp;gt;   &amp;lt;li&amp;gt;Both opaque parameters must be specified.&amp;lt;/li&amp;gt;  &amp;lt;li&amp;gt;You cannot include card number or expiration date.&amp;lt;/li&amp;gt;  &amp;lt;li&amp;gt;You cannot include track data.&amp;lt;/li&amp;gt;  &amp;lt;li&amp;gt;Must be E-commerce transaction. Confirm that your gateway account is set up as a Card Not Present account.&amp;lt;/li&amp;gt;  &amp;lt;li&amp;gt;The transaction must be authorization or authorize and capture type of transaction.&amp;lt;/li&amp;gt;  &amp;lt;li&amp;gt;You cannot include 3DS data.&amp;lt;/li&amp;gt; &amp;lt;li&amp;gt;You must submit data that can be successfully decrypted.&amp;lt;/li&amp;gt;  &amp;lt;li&amp;gt;Decrypted data must belong to the merchant submitting the request.&amp;lt;/li&amp;gt;  &amp;lt;li&amp;gt;The submitted data must be base64 encoded.&amp;lt;/li&amp;gt;&amp;lt;/ul&amp;gt;"
	},
	{
		"code" : "154",
		"text" : "Processing Apple Payments is not enabled for this merchant account.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "155",
		"text" : "This processor does not support this method of submitting payment data.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "156",
		"text" : "The cryptogram is either invalid or cannot be used in combination with other parameters.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "165",
		"text" : "This transaction has been declined.",
		"integration_suggestions" : "",
		"other_suggestions" : "System void failed. CVV2 Code mismatch based on the CVV response and the merchant settings."
	},
	{
		"code" : "170",
		"text" : "An error occurred during processing. Please contact the merchant.",
		"integration_suggestions" : "",
		"other_suggestions" : "Concord EFS - Provisioning at the processor has not been completed."
	},
	{
		"code" : "171",
		"text" : "An error occurred during processing. Please contact the merchant.",
		"integration_suggestions" : "",
		"other_suggestions" : "Concord EFS - This request is invalid."
	},
	{
		"code" : "172",
		"text" : "An error occurred during processing. Please contact the merchant.",
		"integration_suggestions" : "",
		"other_suggestions" : "Concord EFS - The store ID is invalid."
	},
	{
		"code" : "173",
		"text" : "An error occurred during processing. Please contact the merchant.",
		"integration_suggestions" : "",
		"other_suggestions" : "Concord EFS - The store key is invalid."
	},
	{
		"code" : "174",
		"text" : "The transaction type is invalid. Please contact the merchant.",
		"integration_suggestions" : "",
		"other_suggestions" : "Concord EFS - This transaction type is not accepted by the processor."
	},
	{
		"code" : "175",
		"text" : "This processor does not allow voiding of credits.",
		"integration_suggestions" : "",
		"other_suggestions" : "Concord EFS - This transaction is not allowed. The Concord EFS processing platform does not support voiding credit transactions. Please debit the credit card instead of voiding the credit."
	},
	{
		"code" : "180",
		"text" : "An error occurred during processing.  Please try again.",
		"integration_suggestions" : "There are three different reasons that an Error 180 might occur:&amp;lt;ol&amp;gt;&amp;lt;li&amp;gt;There was an attempt to void a refund on a processor that does not allow that.&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;A merchant on the Concord EFS platform is attempting to pass AMEX CID, when they are not signed up to validate this value with AMEX.&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;Transactions submitted from test credit card numbers (both ours and othersÃ¢â‚¬â„¢) by merchants on the TSYS payment processing platform, will return a response of:  &amp;quot;(180) An error occurred during processing. Please try again. Invalid processor response format,&amp;quot; rather than &amp;quot;(2) Declined.  This transaction has been declined.&amp;quot;&amp;lt;/li&amp;gt;&amp;lt;/ol&amp;gt;&amp;lt;p&amp;gt;Note that the TSYS payment processing platform was formerly known as Vital or Visanet. On TSYS/Vital/Visanet, Error 180 is an valid response indicating that a transaction was submitted and correctly received, but rejected due to using a test card number.  If the processor is incorrectly configured, the response will be something more generic like a response of 30, 34, or 35.",
		"other_suggestions" : "The processor response format is invalid."
	},
	{
		"code" : "181",
		"text" : "An error occurred during processing.  Please try again.",
		"integration_suggestions" : "",
		"other_suggestions" : "The system-generated void for the original invalid transaction failed. (The original transaction included an invalid processor response format.)"
	},
	{
		"code" : "185",
		"text" : "This transaction cannot be processed.",
		"integration_suggestions" : "",
		"other_suggestions" : "Merchant is not configured for VPOS."
	},
	{
		"code" : "191",
		"text" : "This transaction has been declined.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "192",
		"text" : "An error occurred during processing. Please try again.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "193",
		"text" : "The transaction is currently under review.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "195",
		"text" : "One or more of the HTML type configuration fields do not appear to be safe.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "200",
		"text" : "This transaction has been declined",
		"integration_suggestions" : "",
		"other_suggestions" : "This error code applies only to merchants on FDC Omaha. The credit card number is invalid."
	},
	{
		"code" : "201",
		"text" : "This transaction has been declined",
		"integration_suggestions" : "",
		"other_suggestions" : "This error code applies only to merchants on FDC Omaha. The expiration date is invalid."
	},
	{
		"code" : "202",
		"text" : "This transaction has been declined",
		"integration_suggestions" : "",
		"other_suggestions" : "This error code applies only to merchants on FDC Omaha. The transaction type is invalid."
	},
	{
		"code" : "203",
		"text" : "This transaction has been declined",
		"integration_suggestions" : "",
		"other_suggestions" : "This error code applies only to merchants on FDC Omaha. The value submitted in the amount field is invalid."
	},
	{
		"code" : "204",
		"text" : "This transaction has been declined",
		"integration_suggestions" : "",
		"other_suggestions" : "This error code applies only to merchants on FDC Omaha. The department code is invalid."
	},
	{
		"code" : "205",
		"text" : "This transaction has been declined",
		"integration_suggestions" : "",
		"other_suggestions" : "This error code applies only to merchants on FDC Omaha. The value submitted in the merchant number field is invalid."
	},
	{
		"code" : "206",
		"text" : "This transaction has been declined.",
		"integration_suggestions" : "",
		"other_suggestions" : "This error code applies only to merchants on FDC Omaha. The merchant is not on file."
	},
	{
		"code" : "207",
		"text" : "This transaction has been declined",
		"integration_suggestions" : "",
		"other_suggestions" : "This error code applies only to merchants on FDC Omaha. The merchant account is closed."
	},
	{
		"code" : "208",
		"text" : "This transaction has been declined",
		"integration_suggestions" : "",
		"other_suggestions" : "This error code applies only to merchants on FDC Omaha. The merchant is not on file."
	},
	{
		"code" : "209",
		"text" : "This transaction has been declined",
		"integration_suggestions" : "",
		"other_suggestions" : "This error code applies only to merchants on FDC Omaha. Communication with the processor could not be established."
	},
	{
		"code" : "210",
		"text" : "This transaction has been declined",
		"integration_suggestions" : "",
		"other_suggestions" : "This error code applies only to merchants on FDC Omaha. The merchant type is incorrect."
	},
	{
		"code" : "211",
		"text" : "This transaction has been declined",
		"integration_suggestions" : "",
		"other_suggestions" : "This error code applies only to merchants on FDC Omaha. The cardholder is not on file."
	},
	{
		"code" : "212",
		"text" : "This transaction has been declined",
		"integration_suggestions" : "",
		"other_suggestions" : "This error code applies only to merchants on FDC Omaha. The bank configuration is not on file."
	},
	{
		"code" : "213",
		"text" : "This transaction has been declined",
		"integration_suggestions" : "",
		"other_suggestions" : "This error code applies only to merchants on FDC Omaha. The merchant assessment code is incorrect."
	},
	{
		"code" : "214",
		"text" : "This transaction has been declined",
		"integration_suggestions" : "",
		"other_suggestions" : "This error code applies only to merchants on FDC Omaha. This function is currently unavailable."
	},
	{
		"code" : "215",
		"text" : "This transaction has been declined",
		"integration_suggestions" : "",
		"other_suggestions" : "This error code applies only to merchants on FDC Omaha. The encrypted PIN field format is invalid."
	},
	{
		"code" : "216",
		"text" : "This transaction has been declined",
		"integration_suggestions" : "",
		"other_suggestions" : "This error code applies only to merchants on FDC Omaha. The ATM term ID is invalid."
	},
	{
		"code" : "217",
		"text" : "This transaction has been declined",
		"integration_suggestions" : "",
		"other_suggestions" : "This error code applies only to merchants on FDC Omaha. This transaction experienced a general message format problem."
	},
	{
		"code" : "218",
		"text" : "This transaction has been declined",
		"integration_suggestions" : "",
		"other_suggestions" : "This error code applies only to merchants on FDC Omaha. The PIN block format or PIN availability value is invalid."
	},
	{
		"code" : "219",
		"text" : "This transaction has been declined",
		"integration_suggestions" : "",
		"other_suggestions" : "This error code applies only to merchants on FDC Omaha. The ETC void is unmatched."
	},
	{
		"code" : "220",
		"text" : "This transaction has been declined",
		"integration_suggestions" : "",
		"other_suggestions" : "This error code applies only to merchants on FDC Omaha. The primary CPU is not available."
	},
	{
		"code" : "221",
		"text" : "This transaction has been declined",
		"integration_suggestions" : "",
		"other_suggestions" : "This error code applies only to merchants on FDC Omaha. The SE number is invalid."
	},
	{
		"code" : "222",
		"text" : "This transaction has been declined",
		"integration_suggestions" : "",
		"other_suggestions" : "This error code applies only to merchants on FDC Omaha. Duplicate auth request (from INAS)."
	},
	{
		"code" : "223",
		"text" : "This transaction has been declined",
		"integration_suggestions" : "",
		"other_suggestions" : "This error code applies only to merchants on FDC Omaha. This transaction experienced an unspecified error."
	},
	{
		"code" : "224",
		"text" : "This transaction has been declined",
		"integration_suggestions" : "",
		"other_suggestions" : "This error code applies only to merchants on FDC Omaha. Please re-enter the transaction."
	},
	{
		"code" : "225",
		"text" : "This transaction cannot be processed.",
		"integration_suggestions" : "",
		"other_suggestions" : "The transaction has an invalid dynamic currency conversion (DCC) action."
	},
	{
		"code" : "226",
		"text" : "This transaction cannot be processed.",
		"integration_suggestions" : "",
		"other_suggestions" : "Incomplete set of Dynamic Currency Conversion (DCC) parameters."
	},
	{
		"code" : "227",
		"text" : "This transaction cannot be processed.",
		"integration_suggestions" : "",
		"other_suggestions" : "Merchant is not configured for Dynamic Currency Conversion (DCC). "
	},
	{
		"code" : "228",
		"text" : "This transaction cannot be processed.",
		"integration_suggestions" : "",
		"other_suggestions" : "Dynamic Currency Conversion (DCC) is not allowed for this transaction type. "
	},
	{
		"code" : "229",
		"text" : "Conversion rate for this card is available.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "230",
		"text" : "This transaction cannot be processed.",
		"integration_suggestions" : "Transaction could not be found.",
		"other_suggestions" : ""
	},
	{
		"code" : "231",
		"text" : "This transaction cannot be processed.",
		"integration_suggestions" : "",
		"other_suggestions" : "Incoming data is different than the referenced Dynamic Currency Conversion (DCC) transaction."
	},
	{
		"code" : "232",
		"text" : "This transaction cannot be processed.",
		"integration_suggestions" : "",
		"other_suggestions" : "Merchant is not configured for Dynamic Currency Conversion (DCC). "
	},
	{
		"code" : "233",
		"text" : "This transaction cannot be processed.",
		"integration_suggestions" : "",
		"other_suggestions" : "Cannot perform Dynamic Currency Conversion (DCC) action on this transaction."
	},
	{
		"code" : "234",
		"text" : "This transaction cannot be processed.",
		"integration_suggestions" : "",
		"other_suggestions" : "This transaction has been voided. "
	},
	{
		"code" : "235",
		"text" : "This transaction cannot be processed.",
		"integration_suggestions" : "",
		"other_suggestions" : "This transaction has been captured previously. "
	},
	{
		"code" : "236",
		"text" : "This transaction cannot be processed.",
		"integration_suggestions" : "",
		"other_suggestions" : "Dynamic Currency Conversion (DCC) data for the referenced transaction is not available."
	},
	{
		"code" : "237",
		"text" : "This transaction cannot be processed.",
		"integration_suggestions" : "",
		"other_suggestions" : "The referenced transaction has expired."
	},
	{
		"code" : "238",
		"text" : "This transaction cannot be processed.",
		"integration_suggestions" : "",
		"other_suggestions" : "The transaction version does not support Dynamic Currency Conversion (DCC). "
	},
	{
		"code" : "239",
		"text" : "This transaction cannot be processed.",
		"integration_suggestions" : "",
		"other_suggestions" : "The response format for this transaction does not support Dynamic Currency Conversion (DCC)."
	},
	{
		"code" : "240",
		"text" : "This transaction cannot be processed.",
		"integration_suggestions" : "",
		"other_suggestions" : "Currency for Dynamic Currency Conversion (DCC) transactions must be US dollars."
	},
	{
		"code" : "241",
		"text" : "This transaction cannot be processed.",
		"integration_suggestions" : "",
		"other_suggestions" : "Invalid response from card processor. "
	},
	{
		"code" : "242",
		"text" : "This transaction cannot be processed.",
		"integration_suggestions" : "",
		"other_suggestions" : "Recurring billing flag not allowed on Dynamic Currency Conversion (DCC). "
	},
	{
		"code" : "243",
		"text" : "Recurring billing is not allowed for this eCheck.Net type.",
		"integration_suggestions" : "",
		"other_suggestions" : "The combination of values submitted for &amp;lt;i&amp;gt;x_recurring_billing&amp;lt;/i&amp;gt; and &amp;lt;i&amp;gt;x_echeck_type&amp;lt;/i&amp;gt; is not allowed."
	},
	{
		"code" : "244",
		"text" : "This eCheck.Net type is not allowed for this Bank Account Type.",
		"integration_suggestions" : "",
		"other_suggestions" : "The combination of values submitted for &amp;lt;i&amp;gt;x_bank_acct_type&amp;lt;/i&amp;gt; and &amp;lt;i&amp;gt;x_echeck_type&amp;lt;/i&amp;gt; is not allowed."
	},
	{
		"code" : "245",
		"text" : "This eCheck.Net type is not allowed when using the payment gateway hosted payment form.",
		"integration_suggestions" : "",
		"other_suggestions" : "The value submitted for &amp;lt;i&amp;gt;x_echeck_type&amp;lt;/i&amp;gt; is not allowed when using the payment gateway hosted payment form."
	},
	{
		"code" : "246",
		"text" : "This eCheck.Net type is not allowed.",
		"integration_suggestions" : "",
		"other_suggestions" : "The merchantÃ¢â‚¬â„¢s payment gateway account is not enabled to submit the eCheck.Net type."
	},
	{
		"code" : "247",
		"text" : "This eCheck.Net type is not allowed.",
		"integration_suggestions" : "",
		"other_suggestions" : "The combination of values submitted for &amp;lt;i&amp;gt;x_type&amp;lt;/i&amp;gt; and &amp;lt;i&amp;gt;x_echeck_type&amp;lt;/i&amp;gt; is not allowed."
	},
	{
		"code" : "248",
		"text" : "The check number is invalid.",
		"integration_suggestions" : "",
		"other_suggestions" : "Invalid check number. Check number can only consist of letters and numbers and not more than 15 characters."
	},
	{
		"code" : "250",
		"text" : "This transaction has been declined.",
		"integration_suggestions" : "",
		"other_suggestions" : "This transaction was submitted from a blocked IP address."
	},
	{
		"code" : "251",
		"text" : "This transaction has been declined.",
		"integration_suggestions" : "This transaction was submitted from a blocked IP address. ",
		"other_suggestions" : "The transaction was declined as a result of triggering a Fraud Detection Suite filter."
	},
	{
		"code" : "252",
		"text" : "Your order has been received. Thank you for your business!",
		"integration_suggestions" : "",
		"other_suggestions" : "The transaction was accepted, but is being held for merchant review.  The merchant may customize the customer response in the Merchant Interface."
	},
	{
		"code" : "253",
		"text" : "Your order has been received. Thank you for your business!",
		"integration_suggestions" : "",
		"other_suggestions" : "The transaction was accepted and was authorized, but is being held for merchant review.  The merchant may customize the customer response in the Merchant Interface."
	},
	{
		"code" : "254",
		"text" : "This transaction has been declined.",
		"integration_suggestions" : "",
		"other_suggestions" : "The transaction was declined after manual review."
	},
	{
		"code" : "260",
		"text" : "Reversal not supported for this transaction type.",
		"integration_suggestions" : "",
		"other_suggestions" : "Only authorizations and credits can be reversed. "
	},
	{
		"code" : "261",
		"text" : "An error occurred during processing.  Please try again.",
		"integration_suggestions" : "",
		"other_suggestions" : "The transaction experienced an error during sensitive data encryption and was not processed.  Please try again."
	},
	{
		"code" : "262",
		"text" : "Subtotal is invalid.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "265",
		"text" : "This transaction cannot be processed.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "270",
		"text" : "Line item %1 is invalid.",
		"integration_suggestions" : "",
		"other_suggestions" : "A value submitted in x_line_item for the item referenced is invalid."
	},
	{
		"code" : "271",
		"text" : "The number of line items submitted is not allowed. A maximum of %1 line items can be submitted.",
		"integration_suggestions" : "",
		"other_suggestions" : "The number of line items submitted in x_line_item exceeds the allowed maximum of 30."
	},
	{
		"code" : "280",
		"text" : "The auction platform name is invalid.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "281",
		"text" : "The auction platform ID is invalid.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "282",
		"text" : "The auction listing type is invalid.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "283",
		"text" : "The auction listing ID is invalid.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "284",
		"text" : "The auction seller ID is invalid.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "285",
		"text" : "The auction buyer ID is invalid.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "286",
		"text" : "One or more required auction values were not submitted.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "287",
		"text" : "The combination of auction platform ID and auction platform name is invalid.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "288",
		"text" : "This transaction cannot be accepted.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "289",
		"text" : "This processor does not accept zero dollar authorization for this card type.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "290",
		"text" : "There is one or more missing or invalid required fields.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "295",
		"text" : "The amount of this request was only partially approved on the given prepaid card. An additional payment is required to fulfill the balance of this transaction.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "296",
		"text" : "The specified SplitTenderID is invalid.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "297",
		"text" : "Transaction ID and Split Tender ID cannot both be used in the same request.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "298",
		"text" : "This order has already been released or voided therefore new transaction associations cannot be accepted.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "300",
		"text" : "The device ID is invalid.",
		"integration_suggestions" : "",
		"other_suggestions" : "Invalid &amp;lt;i&amp;gt;x_device_id&amp;lt;/i&amp;gt; value. "
	},
	{
		"code" : "301",
		"text" : "The device batch ID is invalid.",
		"integration_suggestions" : "",
		"other_suggestions" : "Invalid &amp;lt;i&amp;gt;x_device_batch_id&amp;lt;/i&amp;gt; value. "
	},
	{
		"code" : "302",
		"text" : "The reversal flag is invalid.",
		"integration_suggestions" : "",
		"other_suggestions" : "Invalid &amp;lt;i&amp;gt;x_reversal&amp;lt;/i&amp;gt; value."
	},
	{
		"code" : "303",
		"text" : "The device batch is full. Please close the batch.",
		"integration_suggestions" : "",
		"other_suggestions" : "The current device batch must be closed manually from the POS device. "
	},
	{
		"code" : "304",
		"text" : "The original transaction is in a closed batch.",
		"integration_suggestions" : "",
		"other_suggestions" : "The original transaction has been settled and cannot be reversed. "
	},
	{
		"code" : "305",
		"text" : "The merchant is configured for auto-close.",
		"integration_suggestions" : "",
		"other_suggestions" : "This merchant is configured for auto-close and cannot manually close batches. "
	},
	{
		"code" : "306",
		"text" : "The batch is already closed.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "307",
		"text" : "The reversal was processed successfully.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "308",
		"text" : "Original transaction for reversal not found.",
		"integration_suggestions" : "",
		"other_suggestions" : "The transaction submitted for reversal was not found. "
	},
	{
		"code" : "309",
		"text" : "The device has been disabled.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "310",
		"text" : "This transaction has already been voided.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "311",
		"text" : "This transaction has already been captured.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "312",
		"text" : "The specified security code was invalid.",
		"integration_suggestions" : "The SIM hosted payment form features a Security Code option (sometimes called CAPTCHA) used to confirm that the payment is being entered by a human being. This feature helps protect your site from automated scripts that may try to test credit card numbers through the payment form.&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;The Security Code works by generating an image that contains random numbers and letters that cannot be read by scripts. The customer is then prompted to enter the letters and numbers exactly as they appear in the image. If the customer enters the correct Security Code, the transaction is accepted as valid.&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;Error 312 indicates that the customer had entered the wrong Security Code. Should this error occur, a new Security Code is generated, and the customer is prompted to try again until they are successful.&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;To turn off the Security Code option:&amp;lt;ul class&amp;#x3D;&amp;quot;graydot&amp;quot;&amp;gt;&amp;lt;li&amp;gt;Login into the Merchant Interface at &amp;lt;a href&amp;#x3D;&amp;quot;https://account.authorize.net/&amp;quot;&amp;gt;https://account.authorize.net/&amp;lt;/a&amp;gt;.&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;Click &amp;lt;strong&amp;gt;Account&amp;lt;/strong&amp;gt; from the main toolbar.&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;Click &amp;lt;strong&amp;gt;Settings&amp;lt;/strong&amp;gt; in the main left side menu.&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;Click &amp;lt;strong&amp;gt;Payment Form&amp;lt;/strong&amp;gt;.&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;Click &amp;lt;strong&amp;gt;Form Fields&amp;lt;/strong&amp;gt;.&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;Deselect the box labeled &amp;quot;&amp;lt;strong&amp;gt;Require the Security Code feature on the Payment Form&amp;lt;/strong&amp;gt;.&amp;quot;&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;Click &amp;lt;strong&amp;gt;Submit&amp;lt;/strong&amp;gt; to save the settings.&amp;lt;/li&amp;gt;&amp;lt;/ul&amp;gt;&amp;lt;p&amp;gt;Note: When using Simple Checkout, the customer is always required to verify a Security Code. Even if the Security Code is disabled from the payment form, the customer is required to verify a Security Code on the Simple Checkout order page.&amp;lt;/p&amp;gt;",
		"other_suggestions" : ""
	},
	{
		"code" : "313",
		"text" : "A new security code was requested.",
		"integration_suggestions" : "The SIM hosted payment form features a Security Code option (sometimes called CAPTCHA) used to confirm that the payment is being entered by a human being. This feature helps protect your site from automated scripts that may try to test credit card numbers through the payment form.&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;The Security Code works by generating an image that contains random numbers and letters that cannot be read by scripts. The customer is then prompted to enter the letters and numbers exactly as they appear in the image. If the customer enters the correct Security Code, the transaction is accepted as valid.&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;If they enter an incorrect value, the customer is prompted to try again until they are successful.&amp;lt;br /&amp;gt;&amp;lt;br /&amp;gt;To turn off the Security Code option:&amp;lt;ul class&amp;#x3D;&amp;quot;graydot&amp;quot;&amp;gt;&amp;lt;li&amp;gt;Log into the Merchant Interface at &amp;lt;a href&amp;#x3D;&amp;quot;https://account.authorize.net/&amp;quot; target&amp;#x3D;&amp;quot;_blank&amp;quot;&amp;gt;https://account.authorize.net/&amp;lt;/a&amp;gt;.&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;Click &amp;lt;strong&amp;gt;Account&amp;lt;/strong&amp;gt; from the main toolbar.&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;Click &amp;lt;strong&amp;gt;Settings&amp;lt;/strong&amp;gt; in the main left side menu.&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;Click &amp;lt;strong&amp;gt;Payment Form&amp;lt;/strong&amp;gt;.&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;Click &amp;lt;strong&amp;gt;Form Fields&amp;lt;/strong&amp;gt;.&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;Deselect the box labeled &amp;quot;&amp;lt;strong&amp;gt;Require the Security Code feature on the Payment Form&amp;lt;/strong&amp;gt;.&amp;quot;&amp;lt;/li&amp;gt;&amp;lt;li&amp;gt;Click &amp;lt;strong&amp;gt;Submit&amp;lt;/strong&amp;gt; to save the settings.&amp;lt;/li&amp;gt;&amp;lt;/ul&amp;gt;&amp;lt;p&amp;gt;Note: When using Simple Checkout, the customer is always required to verify a Security Code. Even if the Security Code is disabled from the payment form, the customer is required to verify a Security Code on the Simple Checkout order page.&amp;lt;/p&amp;gt;",
		"other_suggestions" : ""
	},
	{
		"code" : "314",
		"text" : "This transaction cannot be processed.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "315",
		"text" : "The credit card number is invalid.",
		"integration_suggestions" : "",
		"other_suggestions" : "This is a processor-issued decline."
	},
	{
		"code" : "316",
		"text" : "Credit card expiration date is invalid.",
		"integration_suggestions" : "",
		"other_suggestions" : "This is a processor-issued decline."
	},
	{
		"code" : "317",
		"text" : "The credit card has expired.",
		"integration_suggestions" : "",
		"other_suggestions" : "This is a processor-issued decline."
	},
	{
		"code" : "318",
		"text" : "A duplicate transaction has been submitted.",
		"integration_suggestions" : "",
		"other_suggestions" : "This is a processor-issued decline."
	},
	{
		"code" : "319",
		"text" : "The transaction cannot be found.",
		"integration_suggestions" : "",
		"other_suggestions" : "This is a processor-issued decline."
	},
	{
		"code" : "320",
		"text" : "The device identifier is either not registered or not enabled.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "325",
		"text" : "The request data did not pass the required fields check for this application.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "326",
		"text" : "The request field(s) are either invalid or missing.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "327",
		"text" : "The void request failed. Either the original transaction type does not support void, or the transaction is in the process of being settled.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "328",
		"text" : "A validation error occurred at the processor.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "330",
		"text" : "V.me transactions are not accepted by this merchant.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "331",
		"text" : "The x_call_id value is missing.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "332",
		"text" : "The x_call_id value is not found or invalid.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "333",
		"text" : "A validation error was returned from V.me.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "334",
		"text" : "The V.me transaction is in an invalid state.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "339",
		"text" : "Use x_method to specify method or send only x_call_id or card/account information.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "340",
		"text" : "V.me by Visa does not support voids on captured or credit transactions. Please allow the transaction to settle, then process a refund for the captured transaction.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "341",
		"text" : "The x_discount value is invalid.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "342",
		"text" : "The x_giftwrap value is invalid.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "343",
		"text" : "The x_subtotal value is invalid.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "344",
		"text" : "The x_misc value is invalid.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "350",
		"text" : "Country must be a valid three-character value if specified.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "351",
		"text" : "Employee ID must be 1 to %x characters in length.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "901",
		"text" : "This transaction cannot be accepted at this time due to system maintenance.  Please try again later.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "2000",
		"text" : "Need payer consent.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "2001",
		"text" : "PayPal transactions are not accepted by this merchant.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "2002",
		"text" : "PayPal transactions require x_version of at least 3.1.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "2003",
		"text" : "Request completed successfully",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "2004",
		"text" : "Success  URL is required.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "2005",
		"text" : "Cancel URL is required.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "2006",
		"text" : "Payer ID is required.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "2007",
		"text" : "This processor does not accept zero dollar authorizations.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "2008",
		"text" : "The referenced transaction does not meet the criteria for issuing a Continued Authorization.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "2009",
		"text" : "The referenced transaction does not meet the criteria for issuing a Continued Authorization  w/ Auto Capture.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "2100",
		"text" : "PayPal transactions require valid URL for success_url",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "2101",
		"text" : "PayPal transactions require valid URL for cancel_url",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "2102",
		"text" : "Payment not authorized.  Payment has not been authorized by the user.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "2103",
		"text" : "This transaction has already been authorized.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "2104",
		"text" : "The totals of the cart item amounts do not match order amounts. Be sure the total of the payment detail item parameters add up to the order total.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "2105",
		"text" : "PayPal has rejected the transaction.Invalid Payer ID.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "2106",
		"text" : "PayPal has already captured this transaction.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "2107",
		"text" : "PayPal has rejected the transaction. This Payer ID belongs to a different customer.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "2108",
		"text" : "PayPal has rejected the transaction. x_paypal_hdrimg exceeds maximum allowable length.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	},
	{
		"code" : "2109",
		"text" : "PayPal has rejected the transaction. x_paypal_payflowcolor must be a 6 character hexadecimal value.",
		"integration_suggestions" : "",
		"other_suggestions" : ""
	}
]
`

var (
	// Codes is a slice of ResponseCode which is used to return an appropriate error message.
	Codes []ResponseCode
)

func init() {
	if err := json.Unmarshal([]byte(ResponseCodes), &Codes); err != nil {
		panic(err.Error())
	}
}

// ResponseCode holds information about a given Authorize.net response code.
type ResponseCode struct {
	Code                   string `json:"code"`
	Text                   string `json:"text"`
	IntegrationSuggessions string `json:"integration_suggestions"`
	OtherSuggestions       string `json:"other_suggestions"`
}

func (r *ResponseCode) String() string {
	return r.Text
}

// GetMessage retrieves the error message for the given response code
func GetMessage(code string) string {
	for _, c := range Codes {
		if c.Code == code {
			return c.String()
		}
	}
	return "Uknown code"
}
