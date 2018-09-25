// Package jobs provides access to the Cloud Talent Solution API.
//
// See https://cloud.google.com/talent-solution/job-search/docs/
//
// Usage example:
//
//   import "google.golang.org/api/jobs/v2"
//   ...
//   jobsService, err := jobs.New(oauthHttpClient)
package jobs // import "google.golang.org/api/jobs/v2"

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	context "golang.org/x/net/context"
	ctxhttp "golang.org/x/net/context/ctxhttp"
	gensupport "google.golang.org/api/gensupport"
	googleapi "google.golang.org/api/googleapi"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = ctxhttp.Do

const apiId = "jobs:v2"
const apiName = "jobs"
const apiVersion = "v2"
const basePath = "https://jobs.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

	// Manage job postings
	JobsScope = "https://www.googleapis.com/auth/jobs"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Companies = NewCompaniesService(s)
	s.Jobs = NewJobsService(s)
	s.V2 = NewV2Service(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Companies *CompaniesService

	Jobs *JobsService

	V2 *V2Service
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewCompaniesService(s *Service) *CompaniesService {
	rs := &CompaniesService{s: s}
	rs.Jobs = NewCompaniesJobsService(s)
	return rs
}

type CompaniesService struct {
	s *Service

	Jobs *CompaniesJobsService
}

func NewCompaniesJobsService(s *Service) *CompaniesJobsService {
	rs := &CompaniesJobsService{s: s}
	return rs
}

type CompaniesJobsService struct {
	s *Service
}

func NewJobsService(s *Service) *JobsService {
	rs := &JobsService{s: s}
	return rs
}

type JobsService struct {
	s *Service
}

func NewV2Service(s *Service) *V2Service {
	rs := &V2Service{s: s}
	return rs
}

type V2Service struct {
	s *Service
}

// BatchDeleteJobsRequest: Input only.
//
// Batch delete jobs request.
type BatchDeleteJobsRequest struct {
	// Filter: Required.
	//
	// The filter string specifies the jobs to be deleted.
	//
	// Supported operator: =, AND
	//
	// The fields eligible for filtering are:
	//
	// * `companyName` (Required)
	// * `requisitionId` (Required)
	//
	// Sample Query: companyName = "companies/123" AND requisitionId =
	// "req-1"
	Filter string `json:"filter,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Filter") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Filter") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *BatchDeleteJobsRequest) MarshalJSON() ([]byte, error) {
	type NoMethod BatchDeleteJobsRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// BucketRange: Represents starting and ending value of a range in
// double.
type BucketRange struct {
	// From: Starting value of the bucket range.
	From float64 `json:"from,omitempty"`

	// To: Ending value of the bucket range.
	To float64 `json:"to,omitempty"`

	// ForceSendFields is a list of field names (e.g. "From") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "From") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *BucketRange) MarshalJSON() ([]byte, error) {
	type NoMethod BucketRange
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *BucketRange) UnmarshalJSON(data []byte) error {
	type NoMethod BucketRange
	var s1 struct {
		From gensupport.JSONFloat64 `json:"from"`
		To   gensupport.JSONFloat64 `json:"to"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.From = float64(s1.From)
	s.To = float64(s1.To)
	return nil
}

// BucketizedCount: Represents count of jobs within one bucket.
type BucketizedCount struct {
	// Count: Number of jobs whose numeric field value fall into `range`.
	Count int64 `json:"count,omitempty"`

	// Range: Bucket range on which histogram was performed for the numeric
	// field,
	// that is, the count represents number of jobs in this range.
	Range *BucketRange `json:"range,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Count") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Count") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *BucketizedCount) MarshalJSON() ([]byte, error) {
	type NoMethod BucketizedCount
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CommuteInfo: Output only.
//
// Commute details related to this job.
type CommuteInfo struct {
	// JobLocation: Location used as the destination in the commute
	// calculation.
	JobLocation *JobLocation `json:"jobLocation,omitempty"`

	// TravelDuration: The number of seconds required to travel to the job
	// location from the query
	// location. A duration of 0 seconds indicates that the job is
	// not
	// reachable within the requested duration, but was returned as part of
	// an
	// expanded query.
	TravelDuration string `json:"travelDuration,omitempty"`

	// ForceSendFields is a list of field names (e.g. "JobLocation") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "JobLocation") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CommuteInfo) MarshalJSON() ([]byte, error) {
	type NoMethod CommuteInfo
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CommutePreference: Input only.
//
// Parameters needed for commute search.
type CommutePreference struct {
	// AllowNonStreetLevelAddress: Optional.
	// If `true`, jobs without street level addresses may also be
	// returned.
	// For city level addresses, the city center is used. For state and
	// coarser
	// level addresses, text matching is used.
	// If this field is set to `false` or is not specified, only jobs that
	// include
	// street level addresses will be returned by commute search.
	AllowNonStreetLevelAddress bool `json:"allowNonStreetLevelAddress,omitempty"`

	// DepartureHourLocal: Optional.
	//
	// The departure hour to use to calculate traffic impact. Accepts
	// an
	// integer between 0 and 23, representing the hour in the time zone of
	// the
	// start_location. Must not be present if road_traffic is specified.
	DepartureHourLocal int64 `json:"departureHourLocal,omitempty"`

	// Method: Required.
	//
	// The method of transportation for which to calculate the commute time.
	//
	// Possible values:
	//   "COMMUTE_METHOD_UNSPECIFIED" - Commute method is not specified.
	//   "DRIVING" - Commute time is calculated based on driving time.
	//   "TRANSIT" - Commute time is calculated based on public transit
	// including bus, metro,
	// subway, etc.
	Method string `json:"method,omitempty"`

	// RoadTraffic: Optional.
	//
	// Specifies the traffic density to use when caculating commute
	// time.
	// Must not be present if departure_hour_local is specified.
	//
	// Possible values:
	//   "ROAD_TRAFFIC_UNSPECIFIED" - Road traffic situation is not
	// specified.
	//   "TRAFFIC_FREE" - Optimal commute time without considering any
	// traffic impact.
	//   "BUSY_HOUR" - Commute time calculation takes in account the peak
	// traffic impact.
	RoadTraffic string `json:"roadTraffic,omitempty"`

	// StartLocation: Required.
	//
	// The latitude and longitude of the location from which to calculate
	// the
	// commute time.
	StartLocation *LatLng `json:"startLocation,omitempty"`

	// TravelTime: Required.
	//
	// The maximum travel time in seconds. The maximum allowed value is
	// `3600s`
	// (one hour). Format is `123s`.
	TravelTime string `json:"travelTime,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "AllowNonStreetLevelAddress") to unconditionally include in API
	// requests. By default, fields with empty values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g.
	// "AllowNonStreetLevelAddress") to include in API requests with the
	// JSON null value. By default, fields with empty values are omitted
	// from API requests. However, any field with an empty value appearing
	// in NullFields will be sent to the server as null. It is an error if a
	// field in this list has a non-empty value. This may be used to include
	// null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CommutePreference) MarshalJSON() ([]byte, error) {
	type NoMethod CommutePreference
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Company: A Company resource represents a company in the service. A
// company is the
// entity that owns job listings, that is, the hiring entity responsible
// for
// employing applicants for the job position.
type Company struct {
	// CareerPageLink: Optional.
	//
	// The URL to employer's career site or careers page on the employer's
	// web
	// site.
	CareerPageLink string `json:"careerPageLink,omitempty"`

	// CompanyInfoSources: Optional.
	//
	// Identifiers external to the application that help to further
	// identify
	// the employer.
	CompanyInfoSources []*CompanyInfoSource `json:"companyInfoSources,omitempty"`

	// CompanySize: Optional.
	//
	// The employer's company size.
	//
	// Possible values:
	//   "COMPANY_SIZE_UNSPECIFIED" - Default value if the size is not
	// specified.
	//   "MINI" - The company has less than 50 employees.
	//   "SMALL" - The company has between 50 and 99 employees.
	//   "SMEDIUM" - The company has between 100 and 499 employees.
	//   "MEDIUM" - The company has between 500 and 999 employees.
	//   "BIG" - The company has between 1,000 and 4,999 employees.
	//   "BIGGER" - The company has between 5,000 and 9,999 employees.
	//   "GIANT" - The company has 10,000 or more employees.
	CompanySize string `json:"companySize,omitempty"`

	// DisableLocationOptimization: Deprecated. Do not use this
	// field.
	//
	// Optional.
	//
	// This field is no longer used. Any value set to it is ignored.
	DisableLocationOptimization bool `json:"disableLocationOptimization,omitempty"`

	// DisplayName: Required.
	//
	// The name of the employer to be displayed with the job,
	// for example, "Google, LLC.".
	DisplayName string `json:"displayName,omitempty"`

	// DistributorBillingCompanyId: Optional.
	//
	// The unique company identifier provided by the client to identify
	// an
	// employer for billing purposes. Recommended practice is to use
	// the distributor_company_id.
	//
	// Defaults to same value as distributor_company_id when a value
	// is not provided.
	DistributorBillingCompanyId string `json:"distributorBillingCompanyId,omitempty"`

	// DistributorCompanyId: Required.
	//
	// A client's company identifier, used to uniquely identify the
	// company. If an employer has a subsidiary or sub-brand, such as
	// "Alphabet"
	// and "Google", which the client wishes to use as the company displayed
	// on
	// the job. Best practice is to create a distinct company identifier for
	// each
	// distinct brand displayed.
	//
	// The maximum number of allowed characters is 255.
	DistributorCompanyId string `json:"distributorCompanyId,omitempty"`

	// EeoText: Optional.
	//
	// Equal Employment Opportunity legal disclaimer text to be
	// associated with all jobs, and typically to be displayed in
	// all
	// roles.
	//
	// The maximum number of allowed characters is 500.
	EeoText string `json:"eeoText,omitempty"`

	// HiringAgency: Optional.
	//
	// Set to true if it is the hiring agency that post jobs for
	// other
	// employers.
	//
	// Defaults to false if not provided.
	HiringAgency bool `json:"hiringAgency,omitempty"`

	// HqLocation: Optional.
	//
	// The street address of the company's main headquarters, which may
	// be
	// different from the job location. The service attempts
	// to geolocate the provided address, and populates a more
	// specific
	// location wherever possible in structured_company_hq_location.
	HqLocation string `json:"hqLocation,omitempty"`

	// ImageUrl: Optional.
	//
	// A URL that hosts the employer's company logo. If provided,
	// the logo image should be squared at 80x80 pixels.
	//
	// The url must be a Google Photos or Google Album url.
	// Only images in these Google sub-domains are accepted.
	ImageUrl string `json:"imageUrl,omitempty"`

	// KeywordSearchableCustomAttributes: Optional.
	//
	// A list of keys of filterable Job.custom_attributes,
	// whose
	// corresponding `string_values` are used in keyword search. Jobs
	// with
	// `string_values` under these specified field keys are returned if
	// any
	// of the values matches the search keyword. Custom field values
	// with
	// parenthesis, brackets and special symbols might not be properly
	// searchable,
	// and those keyword queries need to be surrounded by quotes.
	KeywordSearchableCustomAttributes []string `json:"keywordSearchableCustomAttributes,omitempty"`

	// KeywordSearchableCustomFields: Deprecated. Use
	// keyword_searchable_custom_attributes instead.
	//
	// Optional.
	//
	// A list of filterable custom fields that should be used in
	// keyword
	// search. The jobs of this company are returned if any of these
	// custom
	// fields matches the search keyword. Custom field values with
	// parenthesis,
	// brackets and special symbols might not be properly searchable, and
	// those
	// keyword queries need to be surrounded by quotes.
	KeywordSearchableCustomFields []int64 `json:"keywordSearchableCustomFields,omitempty"`

	// Name: Required during company update.
	//
	// The resource name for a company. This is generated by the service
	// when a
	// company is created, for
	// example,
	// "companies/0000aaaa-1111-bbbb-2222-cccc3333dddd".
	Name string `json:"name,omitempty"`

	// StructuredCompanyHqLocation: Output only.
	//
	// A structured headquarters location of the company,
	// resolved from hq_location if possible.
	StructuredCompanyHqLocation *JobLocation `json:"structuredCompanyHqLocation,omitempty"`

	// Suspended: Output only.
	//
	// Indicates whether a company is flagged to be suspended from
	// public
	// availability by the service when job content appears
	// suspicious,
	// abusive, or spammy.
	Suspended bool `json:"suspended,omitempty"`

	// Title: Deprecated. Use display_name instead.
	//
	// Required.
	//
	// The name of the employer to be displayed with the job,
	// for example, "Google, LLC.".
	Title string `json:"title,omitempty"`

	// Website: Optional.
	//
	// The URL representing the company's primary web site or home
	// page,
	// such as, "www.google.com".
	Website string `json:"website,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "CareerPageLink") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CareerPageLink") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *Company) MarshalJSON() ([]byte, error) {
	type NoMethod Company
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CompanyInfoSource: A resource that represents an external  Google
// identifier for a company,
// for example, a Google+ business page or a Google Maps business
// page.
// For unsupported types, use `unknown_type_id`.
type CompanyInfoSource struct {
	// FreebaseMid: Optional.
	//
	// The Google's Knowledge Graph value for the employer's company.
	FreebaseMid string `json:"freebaseMid,omitempty"`

	// GplusId: Optional.
	//
	// The numeric identifier for the employer's Google+ business page.
	GplusId string `json:"gplusId,omitempty"`

	// MapsCid: Optional.
	//
	// The numeric identifier for the employer's headquarters on Google
	// Maps,
	// namely, the Google Maps CID (cell id).
	MapsCid string `json:"mapsCid,omitempty"`

	// UnknownTypeId: Optional.
	//
	// A Google identifier that does not match any of the other types.
	UnknownTypeId string `json:"unknownTypeId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "FreebaseMid") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "FreebaseMid") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CompanyInfoSource) MarshalJSON() ([]byte, error) {
	type NoMethod CompanyInfoSource
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CompensationEntry: A compensation entry that represents one component
// of compensation, such
// as base pay, bonus, or other compensation type.
//
// Annualization: One compensation entry can be annualized if
// - it contains valid amount or range.
// - and its expected_units_per_year is set or can be derived.
// Its annualized range is determined as (amount or range)
// times
// expected_units_per_year.
type CompensationEntry struct {
	// Amount: Optional.
	//
	// Compensation amount.
	Amount *Money `json:"amount,omitempty"`

	// Description: Optional.
	//
	// Compensation description.  For example, could
	// indicate equity terms or provide additional context to an
	// estimated
	// bonus.
	Description string `json:"description,omitempty"`

	// ExpectedUnitsPerYear: Optional.
	//
	// Expected number of units paid each year. If not specified,
	// when
	// Job.employment_types is FULLTIME, a default value is inferred
	// based on unit. Default values:
	// - HOURLY: 2080
	// - DAILY: 260
	// - WEEKLY: 52
	// - MONTHLY: 12
	// - ANNUAL: 1
	ExpectedUnitsPerYear float64 `json:"expectedUnitsPerYear,omitempty"`

	// Range: Optional.
	//
	// Compensation range.
	Range *CompensationRange `json:"range,omitempty"`

	// Type: Required.
	//
	// Compensation type.
	//
	// Possible values:
	//   "COMPENSATION_TYPE_UNSPECIFIED" - Default value. Equivalent to
	// OTHER_COMPENSATION_TYPE.
	//   "BASE" - Base compensation: Refers to the fixed amount of money
	// paid to an
	// employee by an employer in return for work performed. Base
	// compensation
	// does not include benefits, bonuses or any other potential
	// compensation
	// from an employer.
	//   "BONUS" - Bonus.
	//   "SIGNING_BONUS" - Signing bonus.
	//   "EQUITY" - Equity.
	//   "PROFIT_SHARING" - Profit sharing.
	//   "COMMISSIONS" - Commission.
	//   "TIPS" - Tips.
	//   "OTHER_COMPENSATION_TYPE" - Other compensation type.
	Type string `json:"type,omitempty"`

	// Unit: Optional.
	//
	// Frequency of the specified amount.
	//
	// Default is CompensationUnit.OTHER_COMPENSATION_UNIT.
	//
	// Possible values:
	//   "COMPENSATION_UNIT_UNSPECIFIED" - Default value. Equivalent to
	// OTHER_COMPENSATION_UNIT.
	//   "HOURLY" - Hourly.
	//   "DAILY" - Daily.
	//   "WEEKLY" - Weekly
	//   "MONTHLY" - Monthly.
	//   "YEARLY" - Yearly.
	//   "ONE_TIME" - One time.
	//   "OTHER_COMPENSATION_UNIT" - Other compensation units.
	Unit string `json:"unit,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Amount") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Amount") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CompensationEntry) MarshalJSON() ([]byte, error) {
	type NoMethod CompensationEntry
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *CompensationEntry) UnmarshalJSON(data []byte) error {
	type NoMethod CompensationEntry
	var s1 struct {
		ExpectedUnitsPerYear gensupport.JSONFloat64 `json:"expectedUnitsPerYear"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.ExpectedUnitsPerYear = float64(s1.ExpectedUnitsPerYear)
	return nil
}

// CompensationFilter: Input only.
//
// Filter on job compensation type and amount.
type CompensationFilter struct {
	// IncludeJobsWithUnspecifiedCompensationRange: Optional.
	//
	// Whether to include jobs whose compensation range is unspecified.
	IncludeJobsWithUnspecifiedCompensationRange bool `json:"includeJobsWithUnspecifiedCompensationRange,omitempty"`

	// Range: Optional.
	//
	// Compensation range.
	Range *CompensationRange `json:"range,omitempty"`

	// Type: Required.
	//
	// Type of filter.
	//
	// Possible values:
	//   "FILTER_TYPE_UNSPECIFIED" - Filter type unspecified. Position
	// holder, INVALID, should never be used.
	//   "UNIT_ONLY" - Filter by `base compensation entry's` unit. A job is
	// a match if and
	// only if the job contains a base CompensationEntry and the
	// base
	// CompensationEntry's unit matches provided units.
	// Populate one or more units.
	//
	// See CompensationInfo.CompensationEntry for definition of
	// base compensation entry.
	//   "UNIT_AND_AMOUNT" - Filter by `base compensation entry's` unit and
	// amount / range. A job
	// is a match if and only if the job contains a base CompensationEntry,
	// and
	// the base entry's unit matches provided compensation_units and
	// amount
	// or range overlaps with provided compensation_range.
	//
	// See CompensationInfo.CompensationEntry for definition of
	// base compensation entry.
	//
	// Set exactly one units and populate range.
	//   "ANNUALIZED_BASE_AMOUNT" - Filter by annualized base compensation
	// amount and `base compensation
	// entry's` unit. Populate range and zero or more units.
	//   "ANNUALIZED_TOTAL_AMOUNT" - Filter by annualized total compensation
	// amount and `base compensation
	// entry's` unit . Populate range and zero or more units.
	Type string `json:"type,omitempty"`

	// Units: Required.
	//
	// Specify desired `base compensation
	// entry's`
	// CompensationInfo.CompensationUnit.
	//
	// Possible values:
	//   "COMPENSATION_UNIT_UNSPECIFIED" - Default value. Equivalent to
	// OTHER_COMPENSATION_UNIT.
	//   "HOURLY" - Hourly.
	//   "DAILY" - Daily.
	//   "WEEKLY" - Weekly
	//   "MONTHLY" - Monthly.
	//   "YEARLY" - Yearly.
	//   "ONE_TIME" - One time.
	//   "OTHER_COMPENSATION_UNIT" - Other compensation units.
	Units []string `json:"units,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "IncludeJobsWithUnspecifiedCompensationRange") to unconditionally
	// include in API requests. By default, fields with empty values are
	// omitted from API requests. However, any non-pointer, non-interface
	// field appearing in ForceSendFields will be sent to the server
	// regardless of whether the field is empty or not. This may be used to
	// include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g.
	// "IncludeJobsWithUnspecifiedCompensationRange") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CompensationFilter) MarshalJSON() ([]byte, error) {
	type NoMethod CompensationFilter
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CompensationHistogramRequest: Input only.
//
// Compensation based histogram request.
type CompensationHistogramRequest struct {
	// BucketingOption: Required.
	//
	// Numeric histogram options, like buckets, whether include min or max
	// value.
	BucketingOption *NumericBucketingOption `json:"bucketingOption,omitempty"`

	// Type: Required.
	//
	// Type of the request, representing which field the histogramming
	// should be
	// performed over. A single request can only specify one histogram of
	// each
	// `CompensationHistogramRequestType`.
	//
	// Possible values:
	//   "COMPENSATION_HISTOGRAM_REQUEST_TYPE_UNSPECIFIED" - Default value.
	// Invalid.
	//   "BASE" - Histogram by job's base compensation. See
	// CompensationEntry for
	// definition of base compensation.
	//   "ANNUALIZED_BASE" - Histogram by job's annualized base
	// compensation. See CompensationEntry
	// for definition of annualized base compensation.
	//   "ANNUALIZED_TOTAL" - Histogram by job's annualized total
	// compensation. See CompensationEntry
	// for definition of annualized total compensation.
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BucketingOption") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "BucketingOption") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *CompensationHistogramRequest) MarshalJSON() ([]byte, error) {
	type NoMethod CompensationHistogramRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CompensationHistogramResult: Output only.
//
// Compensation based histogram result.
type CompensationHistogramResult struct {
	// Result: Histogram result.
	Result *NumericBucketingResult `json:"result,omitempty"`

	// Type: Type of the request, corresponding
	// to
	// CompensationHistogramRequest.type.
	//
	// Possible values:
	//   "COMPENSATION_HISTOGRAM_REQUEST_TYPE_UNSPECIFIED" - Default value.
	// Invalid.
	//   "BASE" - Histogram by job's base compensation. See
	// CompensationEntry for
	// definition of base compensation.
	//   "ANNUALIZED_BASE" - Histogram by job's annualized base
	// compensation. See CompensationEntry
	// for definition of annualized base compensation.
	//   "ANNUALIZED_TOTAL" - Histogram by job's annualized total
	// compensation. See CompensationEntry
	// for definition of annualized total compensation.
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Result") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Result") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CompensationHistogramResult) MarshalJSON() ([]byte, error) {
	type NoMethod CompensationHistogramResult
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CompensationInfo: Job compensation details.
type CompensationInfo struct {
	// Amount: Deprecated. Use entries instead.
	//
	// Optional.
	//
	// The amount of compensation or pay for the job.
	// As an alternative, compensation_amount_min
	// and
	// compensation_amount_max may be used to define a range
	// of
	// compensation.
	Amount *Money `json:"amount,omitempty"`

	// AnnualizedBaseCompensationRange: Output only.
	//
	// Annualized base compensation range. Computed as
	// base compensation entry's CompensationEntry.compensation
	// times
	// CompensationEntry.expected_units_per_year.
	//
	// See CompensationEntry for explanation on compensation annualization.
	AnnualizedBaseCompensationRange *CompensationRange `json:"annualizedBaseCompensationRange,omitempty"`

	// AnnualizedTotalCompensationRange: Output only.
	//
	// Annualized total compensation range. Computed as
	// all compensation entries' CompensationEntry.compensation
	// times
	// CompensationEntry.expected_units_per_year.
	//
	// See CompensationEntry for explanation on compensation annualization.
	AnnualizedTotalCompensationRange *CompensationRange `json:"annualizedTotalCompensationRange,omitempty"`

	// Entries: Optional.
	//
	// Job compensation information.
	//
	// At most one entry can be of
	// type
	// CompensationInfo.CompensationType.BASE, which is
	// referred as ** base compensation entry ** for the job.
	Entries []*CompensationEntry `json:"entries,omitempty"`

	// Max: Deprecated. Use entries instead.
	//
	// Optional.
	//
	// An upper bound on a range for compensation or pay for the job.
	// The currency type is specified in compensation_amount.
	Max *Money `json:"max,omitempty"`

	// Min: Deprecated. Use entries instead.
	//
	// Optional.
	//
	// A lower bound on a range for compensation or pay for the job.
	// The currency type is specified in compensation_amount.
	Min *Money `json:"min,omitempty"`

	// Type: Deprecated. Use entries instead.
	//
	// Optional.
	//
	// Type of job compensation.
	//
	// Possible values:
	//   "JOB_COMPENSATION_TYPE_UNSPECIFIED" - The default value if the type
	// is not specified.
	//   "HOURLY" - The job compensation is quoted by the number of hours
	// worked.
	//   "SALARY" - The job compensation is quoted on an annual basis.
	//   "PER_PROJECT" - The job compensation is quoted by project
	// completion.
	//   "COMMISSION" - The job compensation is quoted based solely on
	// commission.
	//   "OTHER_TYPE" - The job compensation is not quoted according to the
	// listed compensation
	// options.
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Amount") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Amount") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CompensationInfo) MarshalJSON() ([]byte, error) {
	type NoMethod CompensationInfo
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CompensationRange: Compensation range.
type CompensationRange struct {
	// Max: Required.
	//
	// The maximum amount of compensation.
	Max *Money `json:"max,omitempty"`

	// Min: Required.
	//
	// The minimum amount of compensation.
	Min *Money `json:"min,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Max") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Max") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CompensationRange) MarshalJSON() ([]byte, error) {
	type NoMethod CompensationRange
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CompleteQueryResponse: Output only.
//
// Response of auto-complete query.
type CompleteQueryResponse struct {
	// CompletionResults: Results of the matching job/company candidates.
	CompletionResults []*CompletionResult `json:"completionResults,omitempty"`

	// Metadata: Additional information for the API invocation, such as the
	// request
	// tracking id.
	Metadata *ResponseMetadata `json:"metadata,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "CompletionResults")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CompletionResults") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *CompleteQueryResponse) MarshalJSON() ([]byte, error) {
	type NoMethod CompleteQueryResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CompletionResult: Output only.
//
// Resource that represents completion results.
type CompletionResult struct {
	// ImageUrl: The URL for the company logo if `type=COMPANY_NAME`.
	ImageUrl string `json:"imageUrl,omitempty"`

	// Suggestion: The suggestion for the query.
	Suggestion string `json:"suggestion,omitempty"`

	// Type: The completion topic.
	//
	// Possible values:
	//   "COMPLETION_TYPE_UNSPECIFIED" - Default value.
	//   "JOB_TITLE" - Only suggest job titles.
	//   "COMPANY_NAME" - Only suggest company names.
	//   "COMBINED" - Suggest both job titles and company names.
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ImageUrl") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ImageUrl") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CompletionResult) MarshalJSON() ([]byte, error) {
	type NoMethod CompletionResult
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CreateJobRequest: Input only.
//
// Create job request.
type CreateJobRequest struct {
	// DisableStreetAddressResolution: Deprecated. Please use
	// processing_options. This flag is ignored if
	// processing_options is set.
	//
	// Optional.
	//
	// If set to `true`, the service does not attempt to resolve a
	// more precise address for the job.
	DisableStreetAddressResolution bool `json:"disableStreetAddressResolution,omitempty"`

	// Job: Required.
	//
	// The Job to be created.
	Job *Job `json:"job,omitempty"`

	// ProcessingOptions: Optional.
	//
	// Options for job processing.
	ProcessingOptions *JobProcessingOptions `json:"processingOptions,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "DisableStreetAddressResolution") to unconditionally include in API
	// requests. By default, fields with empty values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g.
	// "DisableStreetAddressResolution") to include in API requests with the
	// JSON null value. By default, fields with empty values are omitted
	// from API requests. However, any field with an empty value appearing
	// in NullFields will be sent to the server as null. It is an error if a
	// field in this list has a non-empty value. This may be used to include
	// null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CreateJobRequest) MarshalJSON() ([]byte, error) {
	type NoMethod CreateJobRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CustomAttribute: Custom attribute values that are either filterable
// or non-filterable.
type CustomAttribute struct {
	// Filterable: Optional.
	//
	// If the `filterable` flag is true, custom field values are
	// searchable.
	// If false, values are not searchable.
	//
	// Default is false.
	Filterable bool `json:"filterable,omitempty"`

	// LongValue: Optional but at least one of string_values or long_value
	// must
	// be specified.
	//
	// This field is used to perform number range search.
	// (`EQ`, `GT`, `GE`, `LE`, `LT`) over filterable `long_value`.
	// For
	// `long_value`, a value between Long.MIN and Long.MIN is allowed.
	LongValue int64 `json:"longValue,omitempty,string"`

	// StringValues: Optional but at least one of string_values or
	// long_value must
	// be specified.
	//
	// This field is used to perform a string match (`CASE_SENSITIVE_MATCH`
	// or
	// `CASE_INSENSITIVE_MATCH`) search.
	// For filterable `string_values`, a maximum total number of 200
	// values
	// is allowed, with each `string_value` has a byte size of no more
	// than
	// 255B. For unfilterable `string_values`, the maximum total byte size
	// of
	// unfilterable `string_values` is 50KB.
	//
	// Empty strings are not allowed.
	StringValues *StringValues `json:"stringValues,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Filterable") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Filterable") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CustomAttribute) MarshalJSON() ([]byte, error) {
	type NoMethod CustomAttribute
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CustomAttributeHistogramRequest: Custom attributes histogram request.
// An error will be thrown if neither
// string_value_histogram or long_value_histogram_bucketing_option
// has
// been defined.
type CustomAttributeHistogramRequest struct {
	// Key: Required.
	//
	// Specifies the custom field key to perform a histogram on. If
	// specified
	// without `long_value_histogram_bucketing_option`, histogram on string
	// values
	// of the given `key` is triggered, otherwise histogram is performed on
	// long
	// values.
	Key string `json:"key,omitempty"`

	// LongValueHistogramBucketingOption: Optional.
	//
	// Specifies buckets used to perform a range histogram on
	// Job's
	// filterable long custom field values, or min/max value requirements.
	LongValueHistogramBucketingOption *NumericBucketingOption `json:"longValueHistogramBucketingOption,omitempty"`

	// StringValueHistogram: Optional. If set to true, the response will
	// include the histogram value for
	// each key as a string.
	StringValueHistogram bool `json:"stringValueHistogram,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Key") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Key") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CustomAttributeHistogramRequest) MarshalJSON() ([]byte, error) {
	type NoMethod CustomAttributeHistogramRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CustomAttributeHistogramResult: Output only.
//
// Custom attribute histogram result.
type CustomAttributeHistogramResult struct {
	// Key: Stores the key of custom attribute the histogram is performed
	// on.
	Key string `json:"key,omitempty"`

	// LongValueHistogramResult: Stores bucketed histogram counting result
	// or min/max values for
	// custom attribute long values associated with `key`.
	LongValueHistogramResult *NumericBucketingResult `json:"longValueHistogramResult,omitempty"`

	// StringValueHistogramResult: Stores a map from the values of string
	// custom field associated
	// with `key` to the number of jobs with that value in this histogram
	// result.
	StringValueHistogramResult map[string]int64 `json:"stringValueHistogramResult,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Key") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Key") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CustomAttributeHistogramResult) MarshalJSON() ([]byte, error) {
	type NoMethod CustomAttributeHistogramResult
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CustomField: Resource that represents the custom data not captured by
// the standard fields.
type CustomField struct {
	// Values: Optional.
	//
	// The values of the custom data.
	Values []string `json:"values,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Values") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Values") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CustomField) MarshalJSON() ([]byte, error) {
	type NoMethod CustomField
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CustomFieldFilter: Input only.
//
// Custom field filter of the search.
type CustomFieldFilter struct {
	// Queries: Required.
	//
	// The query strings for the filter.
	Queries []string `json:"queries,omitempty"`

	// Type: Optional.
	//
	// The type of filter.
	// Defaults to FilterType.OR.
	//
	// Possible values:
	//   "FILTER_TYPE_UNSPECIFIED" - Default value.
	//   "OR" - Search for a match with any query.
	//   "AND" - Search for a match with all queries.
	//   "NOT" - Negate the set of filter values for the search.
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Queries") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Queries") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CustomFieldFilter) MarshalJSON() ([]byte, error) {
	type NoMethod CustomFieldFilter
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Date: Represents a whole calendar date, e.g. date of birth. The time
// of day and
// time zone are either specified elsewhere or are not significant. The
// date
// is relative to the Proleptic Gregorian Calendar. The day may be 0
// to
// represent a year and month where the day is not significant, e.g.
// credit card
// expiration date. The year may be 0 to represent a month and day
// independent
// of year, e.g. anniversary date. Related types are
// google.type.TimeOfDay
// and `google.protobuf.Timestamp`.
type Date struct {
	// Day: Day of month. Must be from 1 to 31 and valid for the year and
	// month, or 0
	// if specifying a year/month where the day is not significant.
	Day int64 `json:"day,omitempty"`

	// Month: Month of year. Must be from 1 to 12, or 0 if specifying a date
	// without a
	// month.
	Month int64 `json:"month,omitempty"`

	// Year: Year of date. Must be from 1 to 9999, or 0 if specifying a date
	// without
	// a year.
	Year int64 `json:"year,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Day") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Day") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Date) MarshalJSON() ([]byte, error) {
	type NoMethod Date
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DeleteJobsByFilterRequest: Deprecated. Use BatchDeleteJobsRequest
// instead.
//
// Input only.
//
// Delete job by filter request.
//
// The job typically becomes unsearchable within 10 seconds, but it may
// take
// up to 5 minutes.
type DeleteJobsByFilterRequest struct {
	// DisableFastProcess: Optional.
	//
	// If set to true, this call waits for all processing steps to
	// complete
	// before the job is cleaned up. Otherwise, the call returns while
	// some
	// steps are still taking place asynchronously, hence faster.
	DisableFastProcess bool `json:"disableFastProcess,omitempty"`

	// Filter: Required.
	//
	// Restrictions on the scope of the delete request.
	Filter *Filter `json:"filter,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DisableFastProcess")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DisableFastProcess") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *DeleteJobsByFilterRequest) MarshalJSON() ([]byte, error) {
	type NoMethod DeleteJobsByFilterRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DeviceInfo: Input only.
//
// Device information collected from the job seeker, candidate, or
// other entity conducting the job search. Providing this information
// improves
// the quality of the search results across devices.
type DeviceInfo struct {
	// DeviceType: Optional.
	//
	// Type of the device.
	//
	// Possible values:
	//   "DEVICE_TYPE_UNSPECIFIED" - The device type isn't specified.
	//   "WEB" - A desktop web browser, such as, Chrome, Firefox, Safari, or
	// Internet
	// Explorer)
	//   "MOBILE_WEB" - A mobile device web browser, such as a phone or
	// tablet with a Chrome
	// browser.
	//   "ANDROID" - An Android device native application.
	//   "IOS" - An iOS device native application.
	//   "BOT" - A bot, as opposed to a device operated by human beings,
	// such as a web crawler.
	//   "OTHER" - Other devices types.
	DeviceType string `json:"deviceType,omitempty"`

	// Id: Optional.
	//
	// A device-specific ID. The ID must be a unique identifier that
	// distinguishes
	// the device from other devices.
	Id string `json:"id,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DeviceType") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DeviceType") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DeviceInfo) MarshalJSON() ([]byte, error) {
	type NoMethod DeviceInfo
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Empty: A generic empty message that you can re-use to avoid defining
// duplicated
// empty messages in your APIs. A typical example is to use it as the
// request
// or the response type of an API method. For instance:
//
//     service Foo {
//       rpc Bar(google.protobuf.Empty) returns
// (google.protobuf.Empty);
//     }
//
// The JSON representation for `Empty` is empty JSON object `{}`.
type Empty struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
}

// ExtendedCompensationFilter: Deprecated. Always use
// CompensationFilter.
//
// Input only.
//
// Filter on job compensation type and amount.
type ExtendedCompensationFilter struct {
	// CompensationRange: Optional.
	//
	// Compensation range.
	CompensationRange *ExtendedCompensationInfoCompensationRange `json:"compensationRange,omitempty"`

	// CompensationUnits: Required.
	//
	// Specify desired `base compensation
	// entry's`
	// ExtendedCompensationInfo.CompensationUnit.
	//
	// Possible values:
	//   "EXTENDED_COMPENSATION_UNIT_UNSPECIFIED" - Default value.
	// Equivalent to OTHER_COMPENSATION_UNIT.
	//   "HOURLY" - Hourly.
	//   "DAILY" - Daily.
	//   "WEEKLY" - Weekly
	//   "MONTHLY" - Monthly.
	//   "YEARLY" - Yearly.
	//   "ONE_TIME" - One time.
	//   "OTHER_COMPENSATION_UNIT" - Other compensation units.
	CompensationUnits []string `json:"compensationUnits,omitempty"`

	// Currency: Optional.
	//
	// Specify currency in 3-letter
	// [ISO 4217](https://www.iso.org/iso-4217-currency-codes.html) format.
	// If
	// unspecified, jobs are returned regardless of currency.
	Currency string `json:"currency,omitempty"`

	// IncludeJobWithUnspecifiedCompensationRange: Optional.
	//
	// Whether to include jobs whose compensation range is unspecified.
	IncludeJobWithUnspecifiedCompensationRange bool `json:"includeJobWithUnspecifiedCompensationRange,omitempty"`

	// Type: Required.
	//
	// Type of filter.
	//
	// Possible values:
	//   "FILTER_TYPE_UNSPECIFIED" - Filter type unspecified. Position
	// holder, INVALID, should never be used.
	//   "UNIT_ONLY" - Filter by `base compensation entry's` unit. A job is
	// a match if and
	// only if the job contains a base CompensationEntry and the
	// base
	// CompensationEntry's unit matches provided
	// compensation_units.
	// Populate one or more compensation_units.
	//
	// See ExtendedCompensationInfo.CompensationEntry for definition of
	// base compensation entry.
	//   "UNIT_AND_AMOUNT" - Filter by `base compensation entry's` unit and
	// amount / range. A job
	// is a match if and only if the job contains a base CompensationEntry,
	// and
	// the base entry's unit matches provided compensation_units and
	// amount
	// or range overlaps with provided compensation_range.
	//
	// See ExtendedCompensationInfo.CompensationEntry for definition of
	// base compensation entry.
	//
	// Set exactly one
	// compensation_units and populate
	// compensation_range.
	//   "ANNUALIZED_BASE_AMOUNT" - Filter by annualized base compensation
	// amount and `base compensation
	// entry's` unit. Populate compensation_range and zero or
	// more
	// compensation_units.
	//   "ANNUALIZED_TOTAL_AMOUNT" - Filter by annualized total compensation
	// amount and `base compensation
	// entry's` unit . Populate compensation_range and zero or
	// more
	// compensation_units.
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CompensationRange")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CompensationRange") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ExtendedCompensationFilter) MarshalJSON() ([]byte, error) {
	type NoMethod ExtendedCompensationFilter
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ExtendedCompensationInfo: Deprecated. Use
// CompensationInfo.
//
// Describes job compensation.
type ExtendedCompensationInfo struct {
	// AnnualizedBaseCompensationRange: Output only.
	//
	// Annualized base compensation range.
	AnnualizedBaseCompensationRange *ExtendedCompensationInfoCompensationRange `json:"annualizedBaseCompensationRange,omitempty"`

	// AnnualizedBaseCompensationUnspecified: Output only.
	//
	// Indicates annualized base compensation range cannot be derived, due
	// to
	// the job's base compensation entry cannot be annualized.
	// See CompensationEntry for explanation on annualization and
	// base
	// compensation entry.
	AnnualizedBaseCompensationUnspecified bool `json:"annualizedBaseCompensationUnspecified,omitempty"`

	// AnnualizedTotalCompensationRange: Output only.
	//
	// Annualized total compensation range.
	AnnualizedTotalCompensationRange *ExtendedCompensationInfoCompensationRange `json:"annualizedTotalCompensationRange,omitempty"`

	// AnnualizedTotalCompensationUnspecified: Output only.
	//
	// Indicates annualized total compensation range cannot be derived, due
	// to
	// the job's all CompensationEntry cannot be annualized.
	// See CompensationEntry for explanation on annualization and
	// base
	// compensation entry.
	AnnualizedTotalCompensationUnspecified bool `json:"annualizedTotalCompensationUnspecified,omitempty"`

	// Currency: Optional.
	//
	// A 3-letter [ISO
	// 4217](https://www.iso.org/iso-4217-currency-codes.html)
	// currency code.
	Currency string `json:"currency,omitempty"`

	// Entries: Optional.
	//
	// Job compensation information.
	//
	// At most one entry can be of
	// type
	// ExtendedCompensationInfo.CompensationType.BASE, which is
	// referred as ** base compensation entry ** for the job.
	Entries []*ExtendedCompensationInfoCompensationEntry `json:"entries,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "AnnualizedBaseCompensationRange") to unconditionally include in API
	// requests. By default, fields with empty values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g.
	// "AnnualizedBaseCompensationRange") to include in API requests with
	// the JSON null value. By default, fields with empty values are omitted
	// from API requests. However, any field with an empty value appearing
	// in NullFields will be sent to the server as null. It is an error if a
	// field in this list has a non-empty value. This may be used to include
	// null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ExtendedCompensationInfo) MarshalJSON() ([]byte, error) {
	type NoMethod ExtendedCompensationInfo
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ExtendedCompensationInfoCompensationEntry: Deprecated. See
// CompensationInfo.
//
// A compensation entry that represents one component of compensation,
// such
// as base pay, bonus, or other compensation type.
//
// Annualization: One compensation entry can be annualized if
// - it contains valid amount or range.
// - and its expected_units_per_year is set or can be derived.
// Its annualized range is determined as (amount or range)
// times
// expected_units_per_year.
type ExtendedCompensationInfoCompensationEntry struct {
	// Amount: Optional.
	//
	// Monetary amount.
	Amount *ExtendedCompensationInfoDecimal `json:"amount,omitempty"`

	// Description: Optional.
	//
	// Compensation description.
	Description string `json:"description,omitempty"`

	// ExpectedUnitsPerYear: Optional.
	//
	// Expected number of units paid each year. If not specified,
	// when
	// Job.employment_types is FULLTIME, a default value is inferred
	// based on unit. Default values:
	// - HOURLY: 2080
	// - DAILY: 260
	// - WEEKLY: 52
	// - MONTHLY: 12
	// - ANNUAL: 1
	ExpectedUnitsPerYear *ExtendedCompensationInfoDecimal `json:"expectedUnitsPerYear,omitempty"`

	// Range: Optional.
	//
	// Compensation range.
	Range *ExtendedCompensationInfoCompensationRange `json:"range,omitempty"`

	// Type: Required.
	//
	// Compensation type.
	//
	// Possible values:
	//   "EXTENDED_COMPENSATION_TYPE_UNSPECIFIED" - Default value.
	// Equivalent to OTHER_COMPENSATION_TYPE.
	//   "BASE" - Base compensation: Refers to the fixed amount of money
	// paid to an
	// employee by an employer in return for work performed. Base
	// compensation
	// does not include benefits, bonuses or any other potential
	// compensation
	// from an employer.
	//   "BONUS" - Bonus.
	//   "SIGNING_BONUS" - Signing bonus.
	//   "EQUITY" - Equity.
	//   "PROFIT_SHARING" - Profit sharing.
	//   "COMMISSIONS" - Commission.
	//   "TIPS" - Tips.
	//   "OTHER_COMPENSATION_TYPE" - Other compensation type.
	Type string `json:"type,omitempty"`

	// Unit: Optional.
	//
	// Frequency of the specified amount.
	//
	// Default is CompensationUnit.OTHER_COMPENSATION_UNIT.
	//
	// Possible values:
	//   "EXTENDED_COMPENSATION_UNIT_UNSPECIFIED" - Default value.
	// Equivalent to OTHER_COMPENSATION_UNIT.
	//   "HOURLY" - Hourly.
	//   "DAILY" - Daily.
	//   "WEEKLY" - Weekly
	//   "MONTHLY" - Monthly.
	//   "YEARLY" - Yearly.
	//   "ONE_TIME" - One time.
	//   "OTHER_COMPENSATION_UNIT" - Other compensation units.
	Unit string `json:"unit,omitempty"`

	// Unspecified: Optional.
	//
	// Indicates compensation amount and range are unset.
	Unspecified bool `json:"unspecified,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Amount") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Amount") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ExtendedCompensationInfoCompensationEntry) MarshalJSON() ([]byte, error) {
	type NoMethod ExtendedCompensationInfoCompensationEntry
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ExtendedCompensationInfoCompensationRange: Deprecated. See
// CompensationInfo.
//
// Compensation range.
type ExtendedCompensationInfoCompensationRange struct {
	// Max: Required.
	//
	// Maximum value.
	Max *ExtendedCompensationInfoDecimal `json:"max,omitempty"`

	// Min: Required.
	//
	// Minimum value.
	Min *ExtendedCompensationInfoDecimal `json:"min,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Max") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Max") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ExtendedCompensationInfoCompensationRange) MarshalJSON() ([]byte, error) {
	type NoMethod ExtendedCompensationInfoCompensationRange
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ExtendedCompensationInfoDecimal: Deprecated. See
// CompensationInfo.
//
// Decimal number.
type ExtendedCompensationInfoDecimal struct {
	// Micros: Micro (10^-6) units.
	// The value must be between -999,999 and +999,999 inclusive.
	// If `units` is positive, `micros` must be positive or zero.
	// If `units` is zero, `micros` can be positive, zero, or negative.
	// If `units` is negative, `micros` must be negative or zero.
	// For example -1.75 is represented as `units`=-1 and `micros`=-750,000.
	Micros int64 `json:"micros,omitempty"`

	// Units: Whole units.
	Units int64 `json:"units,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "Micros") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Micros") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ExtendedCompensationInfoDecimal) MarshalJSON() ([]byte, error) {
	type NoMethod ExtendedCompensationInfoDecimal
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Filter: Deprecated. Use BatchDeleteJobsRequest instead.
//
// Input only.
//
// Filter for jobs to be deleted.
type Filter struct {
	// RequisitionId: Required.
	//
	// The requisition ID (or posting ID) assigned by the client to identify
	// a
	// job. This is intended for client identification and tracking
	// of
	// listings.
	// name takes precedence over this field
	// The maximum number of allowed characters is 225.
	RequisitionId string `json:"requisitionId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "RequisitionId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "RequisitionId") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Filter) MarshalJSON() ([]byte, error) {
	type NoMethod Filter
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GetHistogramRequest: Deprecated. Use
// SearchJobsRequest.histogram_facets instead to make
// a single call with both search and histogram.
//
// Input only.
//
// A request for the `GetHistogram` method.
type GetHistogramRequest struct {
	// AllowBroadening: Optional.
	//
	// Controls whether to broaden the search to avoid too few results for
	// a
	// given query in instances where a search has sparse results. Results
	// from a
	// broadened query is a superset of the results from the original
	// query.
	//
	// Defaults to false.
	AllowBroadening bool `json:"allowBroadening,omitempty"`

	// Filters: Deprecated. Use query instead.
	//
	// Optional.
	//
	// Restrictions on the scope of the histogram.
	Filters *JobFilters `json:"filters,omitempty"`

	// Query: Optional.
	//
	// Query used to search against jobs, such as keyword, location filters,
	// etc.
	Query *JobQuery `json:"query,omitempty"`

	// RequestMetadata: Meta information, such as `user_id`, collected from
	// the job searcher or
	// other entity conducting the job search, which is used to improve the
	// search
	// quality of the service. Users determine identifier values, which must
	// be
	// unique and consist.
	RequestMetadata *RequestMetadata `json:"requestMetadata,omitempty"`

	// SearchTypes: Required.
	//
	// A list of facets that specify the histogram data to be
	// calculated
	// against and returned.
	//
	// Histogram response times can be slow, and counts
	// can be approximations. This call may be temporarily or permanently
	// removed
	// prior to the production release of Cloud Talent Solution.
	//
	// Possible values:
	//   "JOB_FIELD_UNSPECIFIED" - The default value if search type is not
	// specified.
	//   "COMPANY_ID" - Filter by the company id field.
	//   "EMPLOYMENT_TYPE" - Filter by the employment type field, such as
	// `FULL_TIME` or `PART_TIME`.
	//   "COMPANY_SIZE" - Filter by the company size type field, such as
	// `BIG`, `SMALL` or `BIGGER`.
	//   "DATE_PUBLISHED" - Filter by the date published field. Values are
	// stringified
	// with TimeRange, for example, TimeRange.PAST_MONTH.
	//   "CUSTOM_FIELD_1" - Filter by custom field 1.
	//   "CUSTOM_FIELD_2" - Filter by custom field 2.
	//   "CUSTOM_FIELD_3" - Filter by custom field 3.
	//   "CUSTOM_FIELD_4" - Filter by custom field 4.
	//   "CUSTOM_FIELD_5" - Filter by custom field 5.
	//   "CUSTOM_FIELD_6" - Filter by custom field 6.
	//   "CUSTOM_FIELD_7" - Filter by custom field 7.
	//   "CUSTOM_FIELD_8" - Filter by custom field 8.
	//   "CUSTOM_FIELD_9" - Filter by custom field 9.
	//   "CUSTOM_FIELD_10" - Filter by custom field 10.
	//   "CUSTOM_FIELD_11" - Filter by custom field 11.
	//   "CUSTOM_FIELD_12" - Filter by custom field 12.
	//   "CUSTOM_FIELD_13" - Filter by custom field 13.
	//   "CUSTOM_FIELD_14" - Filter by custom field 14.
	//   "CUSTOM_FIELD_15" - Filter by custom field 15.
	//   "CUSTOM_FIELD_16" - Filter by custom field 16.
	//   "CUSTOM_FIELD_17" - Filter by custom field 17.
	//   "CUSTOM_FIELD_18" - Filter by custom field 18.
	//   "CUSTOM_FIELD_19" - Filter by custom field 19.
	//   "CUSTOM_FIELD_20" - Filter by custom field 20.
	//   "EDUCATION_LEVEL" - Filter by the required education level of the
	// job.
	//   "EXPERIENCE_LEVEL" - Filter by the required experience level of the
	// job.
	//   "ADMIN1" - Filter by Admin1, which is a global placeholder
	// for
	// referring to state, province, or the particular term a country uses
	// to
	// define the geographic structure below the country level.
	// Examples include states codes such as "CA", "IL", "NY",
	// and
	// provinces, such as "BC".
	//   "COUNTRY" - Filter by the country code of job, such as US, JP, FR.
	//   "CITY" - Filter by the "city name", "Admin1 code", for
	// example,
	// "Mountain View, CA" or "New York, NY".
	//   "LOCALE" - Filter by the locale field of a job, such as "en-US",
	// "fr-FR".
	//
	// This is the BCP-47 language code, such as "en-US" or "sr-Latn".
	// For more information, see
	// [Tags for Identifying Languages](https://tools.ietf.org/html/bcp47).
	//   "LANGUAGE" - Filter by the language code portion of the locale
	// field, such as "en" or
	// "fr".
	//   "CATEGORY" - Filter by the Category.
	//   "CITY_COORDINATE" - Filter by the city center GPS coordinate
	// (latitude and longitude), for
	// example, 37.4038522,-122.0987765. Since the coordinates of a city
	// center
	// can change, clients may need to refresh them periodically.
	//   "ADMIN1_COUNTRY" - A combination of state or province code with a
	// country code. This field
	// differs from `JOB_ADMIN1`, which can be used in multiple countries.
	//   "COMPANY_TITLE" - Deprecated. Use COMPANY_DISPLAY_NAME
	// instead.
	//
	// Company display name.
	//   "COMPANY_DISPLAY_NAME" - Company display name.
	//   "BASE_COMPENSATION_UNIT" - Base compensation unit.
	SearchTypes []string `json:"searchTypes,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AllowBroadening") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AllowBroadening") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GetHistogramRequest) MarshalJSON() ([]byte, error) {
	type NoMethod GetHistogramRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GetHistogramResponse: Deprecated. Use
// SearchJobsRequest.histogram_facets instead to make
// a single call with both search and histogram.
//
// Output only.
//
// The response of the GetHistogram method.
type GetHistogramResponse struct {
	// Metadata: Additional information for the API invocation, such as the
	// request
	// tracking id.
	Metadata *ResponseMetadata `json:"metadata,omitempty"`

	// Results: The Histogram results.
	Results []*HistogramResult `json:"results,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Metadata") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Metadata") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GetHistogramResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GetHistogramResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// HistogramFacets: Input only.
//
// Histogram facets to be specified in SearchJobsRequest.
type HistogramFacets struct {
	// CompensationHistogramFacets: Optional.
	//
	// Specifies compensation field-based histogram requests.
	// Duplicate values of CompensationHistogramRequest.type are not
	// allowed.
	CompensationHistogramFacets []*CompensationHistogramRequest `json:"compensationHistogramFacets,omitempty"`

	// CustomAttributeHistogramFacets: Optional.
	//
	// Specifies the custom attributes histogram requests.
	// Duplicate values of CustomAttributeHistogramRequest.key are not
	// allowed.
	CustomAttributeHistogramFacets []*CustomAttributeHistogramRequest `json:"customAttributeHistogramFacets,omitempty"`

	// SimpleHistogramFacets: Optional. Specifies the simple type of
	// histogram facets, for example,
	// `COMPANY_SIZE`, `EMPLOYMENT_TYPE` etc. This field is equivalent
	// to
	// GetHistogramRequest.
	//
	// Possible values:
	//   "JOB_FIELD_UNSPECIFIED" - The default value if search type is not
	// specified.
	//   "COMPANY_ID" - Filter by the company id field.
	//   "EMPLOYMENT_TYPE" - Filter by the employment type field, such as
	// `FULL_TIME` or `PART_TIME`.
	//   "COMPANY_SIZE" - Filter by the company size type field, such as
	// `BIG`, `SMALL` or `BIGGER`.
	//   "DATE_PUBLISHED" - Filter by the date published field. Values are
	// stringified
	// with TimeRange, for example, TimeRange.PAST_MONTH.
	//   "CUSTOM_FIELD_1" - Filter by custom field 1.
	//   "CUSTOM_FIELD_2" - Filter by custom field 2.
	//   "CUSTOM_FIELD_3" - Filter by custom field 3.
	//   "CUSTOM_FIELD_4" - Filter by custom field 4.
	//   "CUSTOM_FIELD_5" - Filter by custom field 5.
	//   "CUSTOM_FIELD_6" - Filter by custom field 6.
	//   "CUSTOM_FIELD_7" - Filter by custom field 7.
	//   "CUSTOM_FIELD_8" - Filter by custom field 8.
	//   "CUSTOM_FIELD_9" - Filter by custom field 9.
	//   "CUSTOM_FIELD_10" - Filter by custom field 10.
	//   "CUSTOM_FIELD_11" - Filter by custom field 11.
	//   "CUSTOM_FIELD_12" - Filter by custom field 12.
	//   "CUSTOM_FIELD_13" - Filter by custom field 13.
	//   "CUSTOM_FIELD_14" - Filter by custom field 14.
	//   "CUSTOM_FIELD_15" - Filter by custom field 15.
	//   "CUSTOM_FIELD_16" - Filter by custom field 16.
	//   "CUSTOM_FIELD_17" - Filter by custom field 17.
	//   "CUSTOM_FIELD_18" - Filter by custom field 18.
	//   "CUSTOM_FIELD_19" - Filter by custom field 19.
	//   "CUSTOM_FIELD_20" - Filter by custom field 20.
	//   "EDUCATION_LEVEL" - Filter by the required education level of the
	// job.
	//   "EXPERIENCE_LEVEL" - Filter by the required experience level of the
	// job.
	//   "ADMIN1" - Filter by Admin1, which is a global placeholder
	// for
	// referring to state, province, or the particular term a country uses
	// to
	// define the geographic structure below the country level.
	// Examples include states codes such as "CA", "IL", "NY",
	// and
	// provinces, such as "BC".
	//   "COUNTRY" - Filter by the country code of job, such as US, JP, FR.
	//   "CITY" - Filter by the "city name", "Admin1 code", for
	// example,
	// "Mountain View, CA" or "New York, NY".
	//   "LOCALE" - Filter by the locale field of a job, such as "en-US",
	// "fr-FR".
	//
	// This is the BCP-47 language code, such as "en-US" or "sr-Latn".
	// For more information, see
	// [Tags for Identifying Languages](https://tools.ietf.org/html/bcp47).
	//   "LANGUAGE" - Filter by the language code portion of the locale
	// field, such as "en" or
	// "fr".
	//   "CATEGORY" - Filter by the Category.
	//   "CITY_COORDINATE" - Filter by the city center GPS coordinate
	// (latitude and longitude), for
	// example, 37.4038522,-122.0987765. Since the coordinates of a city
	// center
	// can change, clients may need to refresh them periodically.
	//   "ADMIN1_COUNTRY" - A combination of state or province code with a
	// country code. This field
	// differs from `JOB_ADMIN1`, which can be used in multiple countries.
	//   "COMPANY_TITLE" - Deprecated. Use COMPANY_DISPLAY_NAME
	// instead.
	//
	// Company display name.
	//   "COMPANY_DISPLAY_NAME" - Company display name.
	//   "BASE_COMPENSATION_UNIT" - Base compensation unit.
	SimpleHistogramFacets []string `json:"simpleHistogramFacets,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "CompensationHistogramFacets") to unconditionally include in API
	// requests. By default, fields with empty values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g.
	// "CompensationHistogramFacets") to include in API requests with the
	// JSON null value. By default, fields with empty values are omitted
	// from API requests. However, any field with an empty value appearing
	// in NullFields will be sent to the server as null. It is an error if a
	// field in this list has a non-empty value. This may be used to include
	// null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *HistogramFacets) MarshalJSON() ([]byte, error) {
	type NoMethod HistogramFacets
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// HistogramResult: Output only.
//
// Result of a histogram call. The response contains the histogram map
// for the
// search type specified by HistogramResult.field.
// The response is a map of each filter value to the corresponding count
// of
// jobs for that filter.
type HistogramResult struct {
	// SearchType: The Histogram search filters.
	//
	// Possible values:
	//   "JOB_FIELD_UNSPECIFIED" - The default value if search type is not
	// specified.
	//   "COMPANY_ID" - Filter by the company id field.
	//   "EMPLOYMENT_TYPE" - Filter by the employment type field, such as
	// `FULL_TIME` or `PART_TIME`.
	//   "COMPANY_SIZE" - Filter by the company size type field, such as
	// `BIG`, `SMALL` or `BIGGER`.
	//   "DATE_PUBLISHED" - Filter by the date published field. Values are
	// stringified
	// with TimeRange, for example, TimeRange.PAST_MONTH.
	//   "CUSTOM_FIELD_1" - Filter by custom field 1.
	//   "CUSTOM_FIELD_2" - Filter by custom field 2.
	//   "CUSTOM_FIELD_3" - Filter by custom field 3.
	//   "CUSTOM_FIELD_4" - Filter by custom field 4.
	//   "CUSTOM_FIELD_5" - Filter by custom field 5.
	//   "CUSTOM_FIELD_6" - Filter by custom field 6.
	//   "CUSTOM_FIELD_7" - Filter by custom field 7.
	//   "CUSTOM_FIELD_8" - Filter by custom field 8.
	//   "CUSTOM_FIELD_9" - Filter by custom field 9.
	//   "CUSTOM_FIELD_10" - Filter by custom field 10.
	//   "CUSTOM_FIELD_11" - Filter by custom field 11.
	//   "CUSTOM_FIELD_12" - Filter by custom field 12.
	//   "CUSTOM_FIELD_13" - Filter by custom field 13.
	//   "CUSTOM_FIELD_14" - Filter by custom field 14.
	//   "CUSTOM_FIELD_15" - Filter by custom field 15.
	//   "CUSTOM_FIELD_16" - Filter by custom field 16.
	//   "CUSTOM_FIELD_17" - Filter by custom field 17.
	//   "CUSTOM_FIELD_18" - Filter by custom field 18.
	//   "CUSTOM_FIELD_19" - Filter by custom field 19.
	//   "CUSTOM_FIELD_20" - Filter by custom field 20.
	//   "EDUCATION_LEVEL" - Filter by the required education level of the
	// job.
	//   "EXPERIENCE_LEVEL" - Filter by the required experience level of the
	// job.
	//   "ADMIN1" - Filter by Admin1, which is a global placeholder
	// for
	// referring to state, province, or the particular term a country uses
	// to
	// define the geographic structure below the country level.
	// Examples include states codes such as "CA", "IL", "NY",
	// and
	// provinces, such as "BC".
	//   "COUNTRY" - Filter by the country code of job, such as US, JP, FR.
	//   "CITY" - Filter by the "city name", "Admin1 code", for
	// example,
	// "Mountain View, CA" or "New York, NY".
	//   "LOCALE" - Filter by the locale field of a job, such as "en-US",
	// "fr-FR".
	//
	// This is the BCP-47 language code, such as "en-US" or "sr-Latn".
	// For more information, see
	// [Tags for Identifying Languages](https://tools.ietf.org/html/bcp47).
	//   "LANGUAGE" - Filter by the language code portion of the locale
	// field, such as "en" or
	// "fr".
	//   "CATEGORY" - Filter by the Category.
	//   "CITY_COORDINATE" - Filter by the city center GPS coordinate
	// (latitude and longitude), for
	// example, 37.4038522,-122.0987765. Since the coordinates of a city
	// center
	// can change, clients may need to refresh them periodically.
	//   "ADMIN1_COUNTRY" - A combination of state or province code with a
	// country code. This field
	// differs from `JOB_ADMIN1`, which can be used in multiple countries.
	//   "COMPANY_TITLE" - Deprecated. Use COMPANY_DISPLAY_NAME
	// instead.
	//
	// Company display name.
	//   "COMPANY_DISPLAY_NAME" - Company display name.
	//   "BASE_COMPENSATION_UNIT" - Base compensation unit.
	SearchType string `json:"searchType,omitempty"`

	// Values: A map from the values of field to the number of jobs with
	// that value
	// in this search result.
	//
	// Key: search type (filter names, such as the companyName).
	//
	// Values: the count of jobs that match the filter for this search.
	Values map[string]int64 `json:"values,omitempty"`

	// ForceSendFields is a list of field names (e.g. "SearchType") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "SearchType") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *HistogramResult) MarshalJSON() ([]byte, error) {
	type NoMethod HistogramResult
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// HistogramResults: Output only.
//
// Histogram results that matches HistogramFacets specified
// in
// SearchJobsRequest.
type HistogramResults struct {
	// CompensationHistogramResults: Specifies compensation field-based
	// histogram results that
	// matches
	// HistogramFacets.compensation_histogram_requests.
	CompensationHistogramResults []*CompensationHistogramResult `json:"compensationHistogramResults,omitempty"`

	// CustomAttributeHistogramResults: Specifies histogram results for
	// custom attributes that
	// matches HistogramFacets.custom_attribute_histogram_facets.
	CustomAttributeHistogramResults []*CustomAttributeHistogramResult `json:"customAttributeHistogramResults,omitempty"`

	// SimpleHistogramResults: Specifies histogram results that
	// matches
	// HistogramFacets.simple_histogram_facets.
	SimpleHistogramResults []*HistogramResult `json:"simpleHistogramResults,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "CompensationHistogramResults") to unconditionally include in API
	// requests. By default, fields with empty values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g.
	// "CompensationHistogramResults") to include in API requests with the
	// JSON null value. By default, fields with empty values are omitted
	// from API requests. However, any field with an empty value appearing
	// in NullFields will be sent to the server as null. It is an error if a
	// field in this list has a non-empty value. This may be used to include
	// null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *HistogramResults) MarshalJSON() ([]byte, error) {
	type NoMethod HistogramResults
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Job: A Job resource represents a job posting (also referred to as a
// "job listing"
// or "job requisition"). A job belongs to a Company, which is the
// hiring
// entity responsible for the job.
type Job struct {
	// ApplicationEmailList: Optional but at least one of
	// application_urls,
	// application_email_list or application_instruction must
	// be
	// specified.
	//
	// Use this field to specify email address(es) to which resumes
	// or
	// applications can be sent.
	//
	// The maximum number of allowed characters is 255.
	ApplicationEmailList []string `json:"applicationEmailList,omitempty"`

	// ApplicationInstruction: Optional but at least one of
	// application_urls,
	// application_email_list or application_instruction must
	// be
	// specified.
	//
	// Use this field to provide instructions, such as "Mail your
	// application
	// to ...", that a candidate can follow to apply for the job.
	//
	// This field accepts and sanitizes HTML input, and also accepts
	// bold, italic, ordered list, and unordered list markup tags.
	//
	// The maximum number of allowed characters is 3,000.
	ApplicationInstruction string `json:"applicationInstruction,omitempty"`

	// ApplicationUrls: Optional but at least one of
	// application_urls,
	// application_email_list or application_instruction must
	// be
	// specified.
	//
	// Use this URL field to direct an applicant to a website, for example
	// to
	// link to an online application form.
	//
	// The maximum number of allowed characters is 2,000.
	ApplicationUrls []string `json:"applicationUrls,omitempty"`

	// Benefits: Optional.
	//
	// The benefits included with the job.
	//
	// Possible values:
	//   "JOB_BENEFIT_TYPE_UNSPECIFIED" - Default value if the type is not
	// specified.
	//   "CHILD_CARE" - The job includes access to programs that support
	// child care, such
	// as daycare.
	//   "DENTAL" - The job includes dental services that are covered by a
	// dental
	// insurance plan.
	//   "DOMESTIC_PARTNER" - The job offers specific benefits to domestic
	// partners.
	//   "FLEXIBLE_HOURS" - The job allows for a flexible work schedule.
	//   "MEDICAL" - The job includes health services that are covered by a
	// medical
	// insurance plan.
	//   "LIFE_INSURANCE" - The job includes a life insurance plan provided
	// by the employer or
	// available for purchase by the employee.
	//   "PARENTAL_LEAVE" - The job allows for a leave of absence to a
	// parent to care for a newborn
	// child.
	//   "RETIREMENT_PLAN" - The job includes a workplace retirement plan
	// provided by the
	// employer or available for purchase by the employee.
	//   "SICK_DAYS" - The job allows for paid time off due to illness.
	//   "TELECOMMUTE" - Deprecated. Set Region.TELECOMMUTE instead.
	//
	// The job allows telecommuting (working remotely).
	//   "VACATION" - The job includes paid time off for vacation.
	//   "VISION" - The job includes vision services that are covered by a
	// vision
	// insurance plan.
	Benefits []string `json:"benefits,omitempty"`

	// CompanyDisplayName: Output only.
	//
	// The name of the company listing the job.
	CompanyDisplayName string `json:"companyDisplayName,omitempty"`

	// CompanyName: Optional but one of company_name or
	// distributor_company_id must be
	// provided.
	//
	// The resource name of the company listing the job, such
	// as
	// /companies/foo. This field takes precedence over
	// the
	// distributor-assigned company identifier, distributor_company_id.
	CompanyName string `json:"companyName,omitempty"`

	// CompanyTitle: Deprecated. Use company_display_name instead.
	//
	// Output only.
	//
	// The name of the company listing the job.
	CompanyTitle string `json:"companyTitle,omitempty"`

	// CompensationInfo: Optional.
	//
	// Job compensation information.
	CompensationInfo *CompensationInfo `json:"compensationInfo,omitempty"`

	// CreateTime: Output only.
	//
	// The timestamp when this job was created.
	CreateTime string `json:"createTime,omitempty"`

	// CustomAttributes: Optional.
	//
	// A map of fields to hold both filterable and non-filterable custom
	// job
	// attributes that are not covered by the provided structured
	// fields.
	//
	// This field is a more general combination of the deprecated
	// id-based
	// filterable_custom_fields and
	// string-based
	// non_filterable_custom_fields.
	//
	// The keys of the map are strings up to 64 bytes and must match
	// the
	// pattern: a-zA-Z*.
	//
	// At most 100 filterable and at most 100 unfilterable keys are
	// supported.
	// For filterable `string_values`, across all keys at most 200 values
	// are
	// allowed, with each string no more than 255 characters. For
	// unfilterable
	// `string_values`, the maximum total size of `string_values` across all
	// keys
	// is 50KB.
	CustomAttributes map[string]CustomAttribute `json:"customAttributes,omitempty"`

	// Department: Optional.
	//
	// The department or functional area within the company with the
	// open
	// position.
	//
	// The maximum number of allowed characters is 255.
	Department string `json:"department,omitempty"`

	// Description: Required.
	//
	// The description of the job, which typically includes a
	// multi-paragraph
	// description of the company and related information. Separate fields
	// are
	// provided on the job object for responsibilities,
	// qualifications, and other job characteristics. Use of
	// these separate job fields is recommended.
	//
	// This field accepts and sanitizes HTML input, and also accepts
	// bold, italic, ordered list, and unordered list markup tags.
	//
	// The maximum number of allowed characters is 100,000.
	Description string `json:"description,omitempty"`

	// DistributorCompanyId: Optional but one of company_name or
	// distributor_company_id must be
	// provided.
	//
	// A unique company identifier used by job distributors to identify
	// an
	// employer's company entity. company_name takes precedence over
	// this field, and is the recommended field to use to identify
	// companies.
	//
	// The maximum number of allowed characters is 255.
	DistributorCompanyId string `json:"distributorCompanyId,omitempty"`

	// EducationLevels: Optional.
	//
	// The desired education level for the job, such as
	// "Bachelors", "Masters", "Doctorate".
	//
	// Possible values:
	//   "EDUCATION_LEVEL_UNSPECIFIED" - The default value if the level is
	// not specified.
	//   "HIGH_SCHOOL" - A High School diploma is required for the position.
	//   "ASSOCIATE" - An Associate degree is required for the position.
	//   "BACHELORS" - A Bachelors degree is required for the position.
	//   "MASTERS" - A Masters degree is required for the position.
	//   "DOCTORATE" - A Doctorate degree is required for the position.
	//   "NO_DEGREE_REQUIRED" - No formal education is required for the
	// position.
	EducationLevels []string `json:"educationLevels,omitempty"`

	// EmploymentTypes: Optional.
	//
	// The employment type(s) of a job, for example,
	// full time or
	// part time.
	//
	// Possible values:
	//   "EMPLOYMENT_TYPE_UNSPECIFIED" - The default value if the employment
	// type is not specified.
	//   "FULL_TIME" - The job requires working a number of hours that
	// constitute full
	// time employment, typically 40 or more hours per week.
	//   "PART_TIME" - The job entails working fewer hours than a full time
	// job,
	// typically less than 40 hours a week.
	//   "CONTRACTOR" - The job is offered as a contracted, as opposed to a
	// salaried employee,
	// position.
	//   "TEMPORARY" - The job is offered as a temporary employment
	// opportunity, usually
	// a short-term engagement.
	//   "INTERN" - The job is a fixed-term opportunity for students or
	// entry-level job seekers
	// to obtain on-the-job training, typically offered as a summer
	// position.
	//   "VOLUNTEER" - The is an opportunity for an individual to volunteer,
	// where there is no
	// expectation of compensation for the provided services.
	//   "PER_DIEM" - The job requires an employee to work on an as-needed
	// basis with a
	// flexible schedule.
	//   "CONTRACT_TO_HIRE" - The job is offered as a contracted position
	// with the understanding
	// that it is converted into a full-time position at the end of
	// the
	// contract. Jobs of this type are also returned by a search
	// for
	// EmploymentType.CONTRACTOR jobs.
	//   "FLY_IN_FLY_OUT" - The job involves employing people in remote
	// areas and flying them
	// temporarily to the work site instead of relocating employees and
	// their
	// families permanently.
	//   "OTHER" - The job does not fit any of the other listed types.
	EmploymentTypes []string `json:"employmentTypes,omitempty"`

	// EndDate: Optional.
	//
	// The end date of the job in UTC time zone. Typically this field
	// is used for contracting engagements.
	// Dates prior to 1970/1/1 and invalid date formats are ignored.
	EndDate *Date `json:"endDate,omitempty"`

	// ExpireTime: Optional but strongly recommended for the best
	// service
	// experience.
	//
	// The expiration timestamp of the job. After this timestamp, the
	// job is marked as expired, and it no longer appears in search results.
	// The
	// expired job can't be deleted or listed by the DeleteJob and
	// ListJobs APIs, but it can be retrieved with the GetJob API or
	// updated with the UpdateJob API. An expired job can be updated
	// and
	// opened again by using a future expiration timestamp. Updating an
	// expired job fails if there is another
	// existing open job with same requisition_id, company_name
	// and
	// language_code.
	//
	// The expired jobs are retained in our system for 90 days. However,
	// the
	// overall expired job count cannot exceed 3 times the maximum of open
	// jobs
	// count over the past week, otherwise jobs with earlier expire time
	// are
	// cleaned first. Expired jobs are no longer accessible after they are
	// cleaned
	// out.
	// The format of this field is RFC 3339 date strings.
	// Example:
	// 2000-01-01T00:00:00.999999999Z
	// See
	// [https://www.ietf.org/rfc/
	// rfc3339.txt](https://www.ietf.org/rfc/rfc3339.txt).
	//
	// A valid date range is between 1970-01-01T00:00:00.0Z
	// and
	// 2100-12-31T23:59:59.999Z. Invalid dates are ignored and treated as
	// expire
	// time not provided.
	//
	// If this value is not provided at the time of job creation or is
	// invalid, the job posting
	// expires after 30 days from the job's creation time. For example, if
	// the
	// job was created on 2017/01/01 13:00AM UTC with an unspecified
	// expiration
	// date, the job expires after 2017/01/31 13:00AM UTC.
	//
	// If this value is not provided but expiry_date is, expiry_date
	// is
	// used.
	//
	// If this value is not provided on job update, it depends on the field
	// masks
	// set by UpdateJobRequest.update_job_fields. If the field masks
	// include
	// expiry_time, or the masks are empty meaning that every field
	// is
	// updated, the job posting expires after 30 days from the job's
	// last
	// update time. Otherwise the expiration date isn't updated.
	ExpireTime string `json:"expireTime,omitempty"`

	// ExpiryDate: Deprecated. Use expire_time instead.
	//
	// Optional but strongly recommended to be provided for the best
	// service
	// experience.
	//
	// The expiration date of the job in UTC time. After 12 am on this date,
	// the
	// job is marked as expired, and it no longer appears in search
	// results.
	// The expired job can't be deleted or listed by the DeleteJob
	// and
	// ListJobs APIs, but it can be retrieved with the GetJob API or
	// updated with the UpdateJob API. An expired job can be updated
	// and
	// opened again by using a future expiration date. It can also remain
	// expired.
	// Updating an expired job to be open fails if there is another existing
	// open
	// job with same requisition_id, company_name and language_code.
	//
	// The expired jobs are retained in our system for 90 days. However,
	// the
	// overall expired job count cannot exceed 3 times the maximum of open
	// jobs
	// count over the past week, otherwise jobs with earlier expire time
	// are
	// removed first. Expired jobs are no longer accessible after they are
	// cleaned
	// out.
	//
	// A valid date range is between 1970/1/1 and 2100/12/31. Invalid dates
	// are
	// ignored and treated as expiry date not provided.
	//
	// If this value is not provided on job creation or is invalid, the
	// job
	// posting expires after 30 days from the job's creation time. For
	// example, if
	// the job was created on 2017/01/01 13:00AM UTC with an
	// unspecified
	// expiration date, the job expires after 2017/01/31 13:00AM UTC.
	//
	// If this value is not provided on job update, it depends on the field
	// masks
	// set by UpdateJobRequest.update_job_fields. If the field masks
	// include
	// expiry_date, or the masks are empty meaning that every field
	// is
	// updated, the job expires after 30 days from the job's last update
	// time.
	// Otherwise the expiration date isn't updated.
	ExpiryDate *Date `json:"expiryDate,omitempty"`

	// ExtendedCompensationInfo: Deprecated. Always use
	// compensation_info.
	//
	// Optional.
	//
	// Job compensation information.
	//
	// This field replaces compensation_info.
	ExtendedCompensationInfo *ExtendedCompensationInfo `json:"extendedCompensationInfo,omitempty"`

	// FilterableCustomFields: Deprecated. Use custom_attributes
	// instead.
	//
	// Optional.
	//
	// A map of fields to hold filterable custom job attributes not captured
	// by
	// the standard fields such as job_title, company_name, or
	// level. These custom fields store arbitrary
	// string values, and can be used for purposes not covered by
	// the structured fields. For the best search experience, use of
	// the
	// structured rather than custom fields is recommended.
	//
	// Data stored in these custom fields fields are indexed and
	// searched against by keyword searches
	// (see
	// SearchJobsRequest.custom_field_filters][]).
	//
	// The map key must be a number between 1-20. If an invalid key
	// is
	// provided on job create or update, an error is returned.
	FilterableCustomFields map[string]CustomField `json:"filterableCustomFields,omitempty"`

	// Incentives: Optional.
	//
	// A description of bonus, commission, and other compensation
	// incentives associated with the job not including salary or pay.
	//
	// The maximum number of allowed characters is 10,000.
	Incentives string `json:"incentives,omitempty"`

	// JobLocations: Output only.
	//
	// Structured locations of the job, resolved from locations.
	JobLocations []*JobLocation `json:"jobLocations,omitempty"`

	// JobTitle: Required.
	//
	// The title of the job, such as "Software Engineer"
	//
	// The maximum number of allowed characters is 500.
	JobTitle string `json:"jobTitle,omitempty"`

	// LanguageCode: Optional.
	//
	// The language of the posting. This field is distinct from
	// any requirements for fluency that are associated with the
	// job.
	//
	// Language codes must be in BCP-47 format, such as "en-US" or
	// "sr-Latn".
	// For more information, see
	// [Tags for Identifying Languages](https://tools.ietf.org/html/bcp47){:
	// class="external" target="_blank" }.
	//
	// The default value is `en-US`.
	LanguageCode string `json:"languageCode,omitempty"`

	// Level: Optional.
	//
	// The experience level associated with the job, such as "Entry Level".
	//
	// Possible values:
	//   "JOB_LEVEL_UNSPECIFIED" - The default value if the level is not
	// specified.
	//   "ENTRY_LEVEL" - Entry-level individual contributors, typically with
	// less than 2 years of
	// experience in a similar role. Includes interns.
	//   "EXPERIENCED" - Experienced individual contributors, typically with
	// 2+ years of
	// experience in a similar role.
	//   "MANAGER" - Entry- to mid-level managers responsible for managing a
	// team of people.
	//   "DIRECTOR" - Senior-level managers responsible for managing teams
	// of managers.
	//   "EXECUTIVE" - Executive-level managers and above, including C-level
	// positions.
	Level string `json:"level,omitempty"`

	// Locations: Optional but strongly recommended for the best service
	// experience.
	//
	// Location(s) where the emploeyer is looking to hire for this job
	// posting.
	//
	// Specifying the full street address(es) of the hiring location
	// enables
	// better API results, especially job searches by commute time.
	//
	// At most 50 locations are allowed for best search performance. If a
	// job has
	// more locations, it is suggested to split it into multiple jobs with
	// unique
	// requisition_ids (e.g. 'ReqA' becomes 'ReqA-1', 'ReqA-2', etc.)
	// as
	// multiple jobs with the same requisition_id, company_name
	// and
	// language_code are not allowed. If the original requisition_id must
	// be preserved, a custom field should be used for storage. It is
	// also
	// suggested to group the locations that close to each other in the same
	// job
	// for better search experience.
	//
	// The maximum number of allowed characters is 500.
	Locations []string `json:"locations,omitempty"`

	// Name: Required during job update.
	//
	// Resource name assigned to a job by the API, for example, "/jobs/foo".
	// Use
	// of this field in job queries and API calls is preferred over the use
	// of
	// requisition_id since this value is unique.
	Name string `json:"name,omitempty"`

	// PromotionValue: Optional.
	//
	// A promotion value of the job, as determined by the client.
	// The value determines the sort order of the jobs returned when
	// searching for
	// jobs using the featured jobs search call, with higher promotional
	// values
	// being returned first and ties being resolved by relevance sort. Only
	// the
	// jobs with a promotionValue >0 are returned in a
	// FEATURED_JOB_SEARCH.
	//
	// Default value is 0, and negative values are treated as 0.
	PromotionValue int64 `json:"promotionValue,omitempty"`

	// PublishDate: Optional.
	//
	// The date this job was most recently published in UTC format. The
	// default
	// value is the time the request arrives at the server.
	PublishDate *Date `json:"publishDate,omitempty"`

	// Qualifications: Optional.
	//
	// A description of the qualifications required to perform the
	// job. The use of this field is recommended
	// as an alternative to using the more general description field.
	//
	// This field accepts and sanitizes HTML input, and also accepts
	// bold, italic, ordered list, and unordered list markup tags.
	//
	// The maximum number of allowed characters is 10,000.
	Qualifications string `json:"qualifications,omitempty"`

	// ReferenceUrl: Output only.
	//
	// The URL of a web page that displays job details.
	ReferenceUrl string `json:"referenceUrl,omitempty"`

	// Region: Optional.
	//
	// The job Region (for example, state, country) throughout which the
	// job
	// is available. If this field is set, a
	// LocationFilter in a search query within the job region
	// finds this job if an exact location match is not specified.
	// If this field is set, setting job locations
	// to the same location level as this field is strongly recommended.
	//
	// Possible values:
	//   "REGION_UNSPECIFIED" - If the region is unspecified, the job is
	// only returned if it
	// matches the LocationFilter.
	//   "STATE_WIDE" - In additiona to exact location matching, job is
	// returned when the
	// LocationFilter in search query is in the same state as this job.
	// For example, if a `STATE_WIDE` job is posted in "CA, USA", it
	// is
	// returned if LocationFilter has "Mountain View".
	//   "NATION_WIDE" - In addition to exact location matching, job is
	// returned when
	// LocationFilter in search query is in the same country as this
	// job.
	// For example, if a `NATION_WIDE` job is posted in "USA", it
	// is
	// returned if LocationFilter has 'Mountain View'.
	//   "TELECOMMUTE" - Job allows employees to work remotely
	// (telecommute).
	// If locations are provided with this value, the job is
	// considered as having a location, but telecommuting is allowed.
	Region string `json:"region,omitempty"`

	// RequisitionId: Required.
	//
	// The requisition ID, also referred to as the posting ID, assigned by
	// the
	// client to identify a job. This field is intended to be used by
	// clients
	// for client identification and tracking of listings. A job is not
	// allowed
	// to be created if there is another job with the same
	// requisition_id,
	// company_name and language_code.
	//
	// The maximum number of allowed characters is 255.
	RequisitionId string `json:"requisitionId,omitempty"`

	// Responsibilities: Optional.
	//
	// A description of job responsibilities. The use of this field
	// is
	// recommended as an alternative to using the more general
	// description
	// field.
	//
	// This field accepts and sanitizes HTML input, and also accepts
	// bold, italic, ordered list, and unordered list markup tags.
	//
	// The maximum number of allowed characters is 10,000.
	Responsibilities string `json:"responsibilities,omitempty"`

	// StartDate: Optional.
	//
	// The start date of the job in UTC time zone. Typically this field
	// is used for contracting engagements.
	// Dates prior to 1970/1/1 and invalid date formats are ignored.
	StartDate *Date `json:"startDate,omitempty"`

	// UnindexedCustomFields: Deprecated. Use custom_attributes
	// instead.
	//
	// Optional.
	//
	// A map of fields to hold non-filterable custom job attributes, similar
	// to
	// filterable_custom_fields. These fields are distinct in that the
	// data
	// in these fields are not indexed. Therefore, the client cannot
	// search
	// against them, nor can the client use them to list jobs.
	//
	// The key of the map can be any valid string.
	UnindexedCustomFields map[string]CustomField `json:"unindexedCustomFields,omitempty"`

	// UpdateTime: Output only.
	//
	// The timestamp when this job was last updated.
	UpdateTime string `json:"updateTime,omitempty"`

	// Visibility: Optional.
	//
	// The visibility of the job.
	// Defaults to JobVisibility.PRIVATE if not specified.
	// Currently only JobVisibility.PRIVATE is supported.
	//
	// Possible values:
	//   "JOB_VISIBILITY_UNSPECIFIED" - Default value.
	//   "PRIVATE" - The Job is only visible to the owner.
	//   "GOOGLE" - The Job is visible to the owner and may be visible to
	// other applications
	// and processes at Google.
	//
	// Not yet supported. Use PRIVATE.
	//   "PUBLIC" - The Job is visible to the owner and may be visible to
	// all other API
	// clients.
	//
	// Not yet supported. Use PRIVATE.
	Visibility string `json:"visibility,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g.
	// "ApplicationEmailList") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ApplicationEmailList") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *Job) MarshalJSON() ([]byte, error) {
	type NoMethod Job
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// JobFilters: Input only.
//
// Deprecated. Use JobQuery instead.
//
// The filters required to perform a search query or histogram.
type JobFilters struct {
	// Categories: Optional.
	//
	// The category filter specifies the categories of jobs to search
	// against.
	// See Category for more information.
	//
	// If a value is not specified, jobs from any category are searched
	// against.
	//
	// If multiple values are specified, jobs from any of the
	// specified
	// categories are searched against.
	//
	// Possible values:
	//   "JOB_CATEGORY_UNSPECIFIED" - The default value if the category is
	// not specified.
	//   "ACCOUNTING_AND_FINANCE" - An accounting and finance job, such as
	// an Accountant.
	//   "ADMINISTRATIVE_AND_OFFICE" - And administrative and office job,
	// such as an Administrative Assistant.
	//   "ADVERTISING_AND_MARKETING" - An advertising and marketing job,
	// such as Marketing Manager.
	//   "ANIMAL_CARE" - An animal care job, such as Veterinarian.
	//   "ART_FASHION_AND_DESIGN" - An art, fashion, or design job, such as
	// Designer.
	//   "BUSINESS_OPERATIONS" - A business operations job, such as Business
	// Operations Manager.
	//   "CLEANING_AND_FACILITIES" - A cleaning and facilities job, such as
	// Custodial Staff.
	//   "COMPUTER_AND_IT" - A computer and IT job, such as Systems
	// Administrator.
	//   "CONSTRUCTION" - A construction job, such as General Laborer.
	//   "CUSTOMER_SERVICE" - A customer service job, such s Cashier.
	//   "EDUCATION" - An education job, such as School Teacher.
	//   "ENTERTAINMENT_AND_TRAVEL" - An entertainment and travel job, such
	// as Flight Attendant.
	//   "FARMING_AND_OUTDOORS" - A farming or outdoor job, such as Park
	// Ranger.
	//   "HEALTHCARE" - A healthcare job, such as Registered Nurse.
	//   "HUMAN_RESOURCES" - A human resources job, such as Human Resources
	// Director.
	//   "INSTALLATION_MAINTENANCE_AND_REPAIR" - An installation,
	// maintenance, or repair job, such as Electrician.
	//   "LEGAL" - A legal job, such as Law Clerk.
	//   "MANAGEMENT" - A management job, often used in conjunction with
	// another category,
	// such as Store Manager.
	//   "MANUFACTURING_AND_WAREHOUSE" - A manufacturing or warehouse job,
	// such as Assembly Technician.
	//   "MEDIA_COMMUNICATIONS_AND_WRITING" - A media, communications, or
	// writing job, such as Media Relations.
	//   "OIL_GAS_AND_MINING" - An oil, gas or mining job, such as Offshore
	// Driller.
	//   "PERSONAL_CARE_AND_SERVICES" - A personal care and services job,
	// such as Hair Stylist.
	//   "PROTECTIVE_SERVICES" - A protective services job, such as Security
	// Guard.
	//   "REAL_ESTATE" - A real estate job, such as Buyer's Agent.
	//   "RESTAURANT_AND_HOSPITALITY" - A restaurant and hospitality job,
	// such as Restaurant Server.
	//   "SALES_AND_RETAIL" - A sales and/or retail job, such Sales
	// Associate.
	//   "SCIENCE_AND_ENGINEERING" - A science and engineering job, such as
	// Lab Technician.
	//   "SOCIAL_SERVICES_AND_NON_PROFIT" - A social services or non-profit
	// job, such as Case Worker.
	//   "SPORTS_FITNESS_AND_RECREATION" - A sports, fitness, or recreation
	// job, such as Personal Trainer.
	//   "TRANSPORTATION_AND_LOGISTICS" - A transportation or logistics job,
	// such as Truck Driver.
	Categories []string `json:"categories,omitempty"`

	// CommuteFilter: Optional.
	//
	//  Allows filtering jobs by commute time with different travel methods
	// (e.g.
	//  driving or public transit). Note: this only works with COMMUTE
	//  MODE. When specified, [JobFilters.location_filters] will be
	//  ignored.
	//
	//  Currently we do not support sorting by commute time.
	CommuteFilter *CommutePreference `json:"commuteFilter,omitempty"`

	// CompanyNames: Optional.
	//
	// The company names filter specifies the company entities to
	// search
	// against.
	//
	// If a value is not specified, jobs are searched for against all
	// companies.
	//
	// If multiple values are specified, jobs are searched against
	// the
	// specified companies.
	//
	// At most 20 company filters are allowed.
	CompanyNames []string `json:"companyNames,omitempty"`

	// CompanyTitles: Optional.
	//
	// This filter specifies the exact company titles
	// of jobs to search against.
	//
	// If a value is not specified, jobs within the search results can
	// be
	// associated with any company.
	//
	// If multiple values are specified, jobs within the search results may
	// be
	// associated with any of the specified companies.
	//
	// At most 20 company title filters are allowed.
	CompanyTitles []string `json:"companyTitles,omitempty"`

	// CompensationFilter: Optional.
	//
	// This search filter is applied only to
	// Job.compensation_info. For example, if the filter is specified
	// as "Hourly job with per-hour compensation > $15", only jobs that
	// meet
	// this criteria are searched. If a filter is not defined, all open
	// jobs
	// are searched.
	CompensationFilter *CompensationFilter `json:"compensationFilter,omitempty"`

	// CustomAttributeFilter: Optional.
	//
	// This filter specifies a structured syntax to match against
	// the
	// Job.custom_attributes that are marked as `filterable`.
	//
	// The syntax for this expression is a subset of Google SQL
	// syntax.
	//
	// Supported operators are: =, !=, <, <=, >, >= where the left of the
	// operator
	// is a custom field key and the right of the operator is a number or
	// string
	// (surrounded by quotes) value.
	//
	// Supported functions are LOWER(<field_name>) to
	// perform case insensitive match and EMPTY(<field_name>) to filter on
	// the
	// existence of a key.
	//
	// Boolean expressions (AND/OR/NOT) are supported up to 3 levels
	// of
	// nesting (For example, "((A AND B AND C) OR NOT D) AND E"), and there
	// can
	// be a maximum of 50 comparisons/functions in the expression. The
	// expression
	// must be < 2000 characters in length.
	//
	// Sample Query:
	// (key1 = "TEST" OR LOWER(key1)="test" OR NOT EMPTY(key1)) AND key2 >
	// 100
	CustomAttributeFilter string `json:"customAttributeFilter,omitempty"`

	// CustomFieldFilters: Deprecated. Use custom_attribute_filter
	// instead.
	//
	// Optional.
	//
	// This filter specifies searching against
	// custom field values. See Job.filterable_custom_fields for
	// information.
	// The key value specifies a number between 1-20 (the service
	// supports 20 custom fields) corresponding to the desired custom field
	// map
	// value. If an invalid key is provided or specified together
	// with
	// custom_attribute_filter, an error is thrown.
	CustomFieldFilters map[string]CustomFieldFilter `json:"customFieldFilters,omitempty"`

	// DisableSpellCheck: Optional.
	//
	// This flag controls the spell-check feature. If false, the
	// service attempts to correct a misspelled query,
	// for example, "enginee" is corrected to "engineer".
	//
	// Defaults to false: a spell check is performed.
	DisableSpellCheck bool `json:"disableSpellCheck,omitempty"`

	// EmploymentTypes: Optional.
	//
	// The employment type filter specifies the employment type of jobs
	// to
	// search against, such as EmploymentType.FULL_TIME.
	//
	// If a value is not specified, jobs in the search results include
	// any
	// employment type.
	//
	// If multiple values are specified, jobs in the search results include
	// any
	// of the specified employment types.
	//
	// Possible values:
	//   "EMPLOYMENT_TYPE_UNSPECIFIED" - The default value if the employment
	// type is not specified.
	//   "FULL_TIME" - The job requires working a number of hours that
	// constitute full
	// time employment, typically 40 or more hours per week.
	//   "PART_TIME" - The job entails working fewer hours than a full time
	// job,
	// typically less than 40 hours a week.
	//   "CONTRACTOR" - The job is offered as a contracted, as opposed to a
	// salaried employee,
	// position.
	//   "TEMPORARY" - The job is offered as a temporary employment
	// opportunity, usually
	// a short-term engagement.
	//   "INTERN" - The job is a fixed-term opportunity for students or
	// entry-level job seekers
	// to obtain on-the-job training, typically offered as a summer
	// position.
	//   "VOLUNTEER" - The is an opportunity for an individual to volunteer,
	// where there is no
	// expectation of compensation for the provided services.
	//   "PER_DIEM" - The job requires an employee to work on an as-needed
	// basis with a
	// flexible schedule.
	//   "CONTRACT_TO_HIRE" - The job is offered as a contracted position
	// with the understanding
	// that it is converted into a full-time position at the end of
	// the
	// contract. Jobs of this type are also returned by a search
	// for
	// EmploymentType.CONTRACTOR jobs.
	//   "FLY_IN_FLY_OUT" - The job involves employing people in remote
	// areas and flying them
	// temporarily to the work site instead of relocating employees and
	// their
	// families permanently.
	//   "OTHER" - The job does not fit any of the other listed types.
	EmploymentTypes []string `json:"employmentTypes,omitempty"`

	// ExtendedCompensationFilter: Deprecated. Always use
	// compensation_filter.
	//
	// Optional.
	//
	// This search filter is applied only to
	// Job.extended_compensation_info. For example, if the filter is
	// specified
	// as "Hourly job with per-hour compensation > $15", only jobs that
	// meet
	// these criteria are searched. If a filter is not defined, all open
	// jobs
	// are searched.
	ExtendedCompensationFilter *ExtendedCompensationFilter `json:"extendedCompensationFilter,omitempty"`

	// LanguageCodes: Optional.
	//
	// This filter specifies the locale of jobs to search against,
	// for example, "en-US".
	//
	// If a value is not specified, the search results may contain jobs in
	// any
	// locale.
	//
	//
	// Language codes should be in BCP-47 format, for example, "en-US" or
	// "sr-Latn".
	// For more information, see
	// [Tags for Identifying
	// Languages](https://tools.ietf.org/html/bcp47).
	//
	// At most 10 language code filters are allowed.
	LanguageCodes []string `json:"languageCodes,omitempty"`

	// LocationFilters: Optional.
	//
	// The location filter specifies geo-regions containing the jobs
	// to
	// search against. See LocationFilter for more information.
	//
	// If a location value is not specified, jobs are be retrieved
	// from all locations.
	//
	// If multiple values are specified, jobs are retrieved from any of
	// the
	// specified locations, and, if different values are specified
	// for the LocationFilter.distance_in_miles parameter, the
	// maximum
	// provided distance is used for all locations.
	//
	// At most 5 location filters are allowed.
	LocationFilters []*LocationFilter `json:"locationFilters,omitempty"`

	// PublishDateRange: Optional.
	//
	// Jobs published within a range specified by this filter are
	// searched
	// against, for example, DateRange.PAST_MONTH. If a value is
	// not
	// specified, all open jobs are searched against regardless of the
	// date they were published.
	//
	// Possible values:
	//   "DATE_RANGE_UNSPECIFIED" - Default value: Filtering on time is not
	// performed.
	//   "PAST_24_HOURS" - The past 24 hours
	//   "PAST_WEEK" - The past week (7 days)
	//   "PAST_MONTH" - The past month (30 days)
	//   "PAST_YEAR" - The past year (365 days)
	//   "PAST_3_DAYS" - The past 3 days
	PublishDateRange string `json:"publishDateRange,omitempty"`

	// Query: Optional.
	//
	// The query filter contains the keywords that match against the
	// job
	// title, description, and location fields.
	//
	// The maximum query size is 255 bytes/characters.
	Query string `json:"query,omitempty"`

	// TenantJobOnly: Deprecated. Do not use this field.
	//
	// This flag controls whether the job search should be restricted to
	// jobs
	// owned by the current user.
	//
	// Defaults to false where all jobs accessible to the
	// user are searched against.
	TenantJobOnly bool `json:"tenantJobOnly,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Categories") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Categories") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *JobFilters) MarshalJSON() ([]byte, error) {
	type NoMethod JobFilters
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// JobLocation: Output only.
//
// A resource that represents a location with full
// geographic
// information.
type JobLocation struct {
	// LatLng: An object representing a latitude/longitude pair.
	LatLng *LatLng `json:"latLng,omitempty"`

	// LocationType: The type of a location, which corresponds to the
	// address lines field of
	// PostalAddress. For example, "Downtown, Atlanta, GA, USA" has a type
	// of
	// LocationType#NEIGHBORHOOD, and "Kansas City, KS, USA" has a type
	// of
	// LocationType#LOCALITY.
	//
	// Possible values:
	//   "LOCATION_TYPE_UNSPECIFIED" - Default value if the type is not
	// specified.
	//   "COUNTRY" - A country level location.
	//   "ADMINISTRATIVE_AREA" - A state or equivalent level location.
	//   "SUB_ADMINISTRATIVE_AREA" - A county or equivalent level location.
	//   "LOCALITY" - A city or equivalent level location.
	//   "POSTAL_CODE" - A postal code level location.
	//   "SUB_LOCALITY" - A sublocality is a subdivision of a locality, for
	// example a city borough,
	// ward, or arrondissement. Sublocalities are usually recognized by a
	// local
	// political authority. For example, Manhattan and Brooklyn are
	// recognized
	// as boroughs by the City of New York, and are therefore modeled
	// as
	// sublocalities.
	//   "SUB_LOCALITY_1" - A district or equivalent level location.
	//   "SUB_LOCALITY_2" - A smaller district or equivalent level display.
	//   "NEIGHBORHOOD" - A neighborhood level location.
	//   "STREET_ADDRESS" - A street address level location.
	LocationType string `json:"locationType,omitempty"`

	// PostalAddress: Postal address of the location that includes human
	// readable information,
	// such as postal delivery and payments addresses. Given a postal
	// address,
	// a postal service can deliver items to a premises, P.O. Box, or
	// other
	// delivery location.
	PostalAddress *PostalAddress `json:"postalAddress,omitempty"`

	// RadiusMeters: Radius in meters of the job location. This value is
	// derived from the
	// location bounding box in which a circle with the specified
	// radius
	// centered from LatLng coves the area associated with the job
	// location.
	// For example, currently, "Mountain View, CA, USA" has a radius
	// of
	// 7885.79 meters.
	RadiusMeters float64 `json:"radiusMeters,omitempty"`

	// ForceSendFields is a list of field names (e.g. "LatLng") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "LatLng") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *JobLocation) MarshalJSON() ([]byte, error) {
	type NoMethod JobLocation
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *JobLocation) UnmarshalJSON(data []byte) error {
	type NoMethod JobLocation
	var s1 struct {
		RadiusMeters gensupport.JSONFloat64 `json:"radiusMeters"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.RadiusMeters = float64(s1.RadiusMeters)
	return nil
}

// JobProcessingOptions: Input only.
//
// Options for job processing.
type JobProcessingOptions struct {
	// DisableStreetAddressResolution: Optional.
	//
	// If set to `true`, the service does not attempt to resolve a
	// more precise address for the job.
	DisableStreetAddressResolution bool `json:"disableStreetAddressResolution,omitempty"`

	// HtmlSanitization: Optional.
	//
	// Option for job HTML content sanitization. Applied fields are:
	//
	// * description
	// * applicationInstruction
	// * incentives
	// * qualifications
	// * responsibilities
	//
	// HTML tags in these fields may be stripped if sanitiazation is not
	// disabled.
	//
	// Defaults to HtmlSanitization.SIMPLE_FORMATTING_ONLY.
	//
	// Possible values:
	//   "HTML_SANITIZATION_UNSPECIFIED" - Default value.
	//   "HTML_SANITIZATION_DISABLED" - Disables sanitization on HTML input.
	//   "SIMPLE_FORMATTING_ONLY" - Sanitizes HTML input, only accepts bold,
	// italic, ordered list, and
	// unordered list markup tags.
	HtmlSanitization string `json:"htmlSanitization,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "DisableStreetAddressResolution") to unconditionally include in API
	// requests. By default, fields with empty values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g.
	// "DisableStreetAddressResolution") to include in API requests with the
	// JSON null value. By default, fields with empty values are omitted
	// from API requests. However, any field with an empty value appearing
	// in NullFields will be sent to the server as null. It is an error if a
	// field in this list has a non-empty value. This may be used to include
	// null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *JobProcessingOptions) MarshalJSON() ([]byte, error) {
	type NoMethod JobProcessingOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// JobQuery: Input only.
//
// The query required to perform a search query or histogram.
type JobQuery struct {
	// Categories: Optional.
	//
	// The category filter specifies the categories of jobs to search
	// against.
	// See Category for more information.
	//
	// If a value is not specified, jobs from any category are searched
	// against.
	//
	// If multiple values are specified, jobs from any of the
	// specified
	// categories are searched against.
	//
	// Possible values:
	//   "JOB_CATEGORY_UNSPECIFIED" - The default value if the category is
	// not specified.
	//   "ACCOUNTING_AND_FINANCE" - An accounting and finance job, such as
	// an Accountant.
	//   "ADMINISTRATIVE_AND_OFFICE" - And administrative and office job,
	// such as an Administrative Assistant.
	//   "ADVERTISING_AND_MARKETING" - An advertising and marketing job,
	// such as Marketing Manager.
	//   "ANIMAL_CARE" - An animal care job, such as Veterinarian.
	//   "ART_FASHION_AND_DESIGN" - An art, fashion, or design job, such as
	// Designer.
	//   "BUSINESS_OPERATIONS" - A business operations job, such as Business
	// Operations Manager.
	//   "CLEANING_AND_FACILITIES" - A cleaning and facilities job, such as
	// Custodial Staff.
	//   "COMPUTER_AND_IT" - A computer and IT job, such as Systems
	// Administrator.
	//   "CONSTRUCTION" - A construction job, such as General Laborer.
	//   "CUSTOMER_SERVICE" - A customer service job, such s Cashier.
	//   "EDUCATION" - An education job, such as School Teacher.
	//   "ENTERTAINMENT_AND_TRAVEL" - An entertainment and travel job, such
	// as Flight Attendant.
	//   "FARMING_AND_OUTDOORS" - A farming or outdoor job, such as Park
	// Ranger.
	//   "HEALTHCARE" - A healthcare job, such as Registered Nurse.
	//   "HUMAN_RESOURCES" - A human resources job, such as Human Resources
	// Director.
	//   "INSTALLATION_MAINTENANCE_AND_REPAIR" - An installation,
	// maintenance, or repair job, such as Electrician.
	//   "LEGAL" - A legal job, such as Law Clerk.
	//   "MANAGEMENT" - A management job, often used in conjunction with
	// another category,
	// such as Store Manager.
	//   "MANUFACTURING_AND_WAREHOUSE" - A manufacturing or warehouse job,
	// such as Assembly Technician.
	//   "MEDIA_COMMUNICATIONS_AND_WRITING" - A media, communications, or
	// writing job, such as Media Relations.
	//   "OIL_GAS_AND_MINING" - An oil, gas or mining job, such as Offshore
	// Driller.
	//   "PERSONAL_CARE_AND_SERVICES" - A personal care and services job,
	// such as Hair Stylist.
	//   "PROTECTIVE_SERVICES" - A protective services job, such as Security
	// Guard.
	//   "REAL_ESTATE" - A real estate job, such as Buyer's Agent.
	//   "RESTAURANT_AND_HOSPITALITY" - A restaurant and hospitality job,
	// such as Restaurant Server.
	//   "SALES_AND_RETAIL" - A sales and/or retail job, such Sales
	// Associate.
	//   "SCIENCE_AND_ENGINEERING" - A science and engineering job, such as
	// Lab Technician.
	//   "SOCIAL_SERVICES_AND_NON_PROFIT" - A social services or non-profit
	// job, such as Case Worker.
	//   "SPORTS_FITNESS_AND_RECREATION" - A sports, fitness, or recreation
	// job, such as Personal Trainer.
	//   "TRANSPORTATION_AND_LOGISTICS" - A transportation or logistics job,
	// such as Truck Driver.
	Categories []string `json:"categories,omitempty"`

	// CommuteFilter: Optional.
	//
	//  Allows filtering jobs by commute time with different travel methods
	// (for
	//  example, driving or public transit). Note: This only works with
	// COMMUTE
	//  MODE. When specified, [JobQuery.location_filters] is
	//  ignored.
	//
	//  Currently we don't support sorting by commute time.
	CommuteFilter *CommutePreference `json:"commuteFilter,omitempty"`

	// CompanyDisplayNames: Optional.
	//
	// This filter specifies the exact company display
	// name of the jobs to search against.
	//
	// If a value isn't specified, jobs within the search results
	// are
	// associated with any company.
	//
	// If multiple values are specified, jobs within the search results may
	// be
	// associated with any of the specified companies.
	//
	// At most 20 company display name filters are allowed.
	CompanyDisplayNames []string `json:"companyDisplayNames,omitempty"`

	// CompanyNames: Optional.
	//
	// This filter specifies the company entities to search against.
	//
	// If a value isn't specified, jobs are searched for against
	// all
	// companies.
	//
	// If multiple values are specified, jobs are searched against
	// the
	// companies specified.
	//
	// At most 20 company filters are allowed.
	CompanyNames []string `json:"companyNames,omitempty"`

	// CompensationFilter: Optional.
	//
	// This search filter is applied only to
	// Job.compensation_info. For example, if the filter is specified
	// as "Hourly job with per-hour compensation > $15", only jobs
	// meeting
	// these criteria are searched. If a filter isn't defined, all open
	// jobs
	// are searched.
	CompensationFilter *CompensationFilter `json:"compensationFilter,omitempty"`

	// CustomAttributeFilter: Optional.
	//
	// This filter specifies a structured syntax to match against
	// the
	// Job.custom_attributes marked as `filterable`.
	//
	// The syntax for this expression is a subset of Google SQL
	// syntax.
	//
	// Supported operators are: =, !=, <, <=, >, >= where the left of the
	// operator
	// is a custom field key and the right of the operator is a number or
	// string
	// (surrounded by quotes) value.
	//
	// Supported functions are LOWER(<field_name>) to
	// perform case insensitive match and EMPTY(<field_name>) to filter on
	// the
	// existence of a key.
	//
	// Boolean expressions (AND/OR/NOT) are supported up to 3 levels
	// of
	// nesting (for example, "((A AND B AND C) OR NOT D) AND E"), a maximum
	// of 50
	// comparisons/functions are allowed in the expression. The
	// expression
	// must be < 2000 characters in length.
	//
	// Sample Query:
	// (key1 = "TEST" OR LOWER(key1)="test" OR NOT EMPTY(key1)) AND key2 >
	// 100
	CustomAttributeFilter string `json:"customAttributeFilter,omitempty"`

	// DisableSpellCheck: Optional.
	//
	// This flag controls the spell-check feature. If false, the
	// service attempts to correct a misspelled query,
	// for example, "enginee" is corrected to "engineer".
	//
	// Defaults to false: a spell check is performed.
	DisableSpellCheck bool `json:"disableSpellCheck,omitempty"`

	// EmploymentTypes: Optional.
	//
	// The employment type filter specifies the employment type of jobs
	// to
	// search against, such as EmploymentType.FULL_TIME.
	//
	// If a value is not specified, jobs in the search results include
	// any
	// employment type.
	//
	// If multiple values are specified, jobs in the search results
	// include
	// any of the specified employment types.
	//
	// Possible values:
	//   "EMPLOYMENT_TYPE_UNSPECIFIED" - The default value if the employment
	// type is not specified.
	//   "FULL_TIME" - The job requires working a number of hours that
	// constitute full
	// time employment, typically 40 or more hours per week.
	//   "PART_TIME" - The job entails working fewer hours than a full time
	// job,
	// typically less than 40 hours a week.
	//   "CONTRACTOR" - The job is offered as a contracted, as opposed to a
	// salaried employee,
	// position.
	//   "TEMPORARY" - The job is offered as a temporary employment
	// opportunity, usually
	// a short-term engagement.
	//   "INTERN" - The job is a fixed-term opportunity for students or
	// entry-level job seekers
	// to obtain on-the-job training, typically offered as a summer
	// position.
	//   "VOLUNTEER" - The is an opportunity for an individual to volunteer,
	// where there is no
	// expectation of compensation for the provided services.
	//   "PER_DIEM" - The job requires an employee to work on an as-needed
	// basis with a
	// flexible schedule.
	//   "CONTRACT_TO_HIRE" - The job is offered as a contracted position
	// with the understanding
	// that it is converted into a full-time position at the end of
	// the
	// contract. Jobs of this type are also returned by a search
	// for
	// EmploymentType.CONTRACTOR jobs.
	//   "FLY_IN_FLY_OUT" - The job involves employing people in remote
	// areas and flying them
	// temporarily to the work site instead of relocating employees and
	// their
	// families permanently.
	//   "OTHER" - The job does not fit any of the other listed types.
	EmploymentTypes []string `json:"employmentTypes,omitempty"`

	// LanguageCodes: Optional.
	//
	// This filter specifies the locale of jobs to search against,
	// for example, "en-US".
	//
	// If a value isn't specified, the search results can contain jobs in
	// any
	// locale.
	//
	//
	// Language codes should be in BCP-47 format, such as "en-US" or
	// "sr-Latn".
	// For more information, see
	// [Tags for Identifying
	// Languages](https://tools.ietf.org/html/bcp47).
	//
	// At most 10 language code filters are allowed.
	LanguageCodes []string `json:"languageCodes,omitempty"`

	// LocationFilters: Optional.
	//
	// The location filter specifies geo-regions containing the jobs
	// to
	// search against. See LocationFilter for more information.
	//
	// If a location value isn't specified, jobs fitting the other
	// search
	// criteria are retrieved regardless of where they're located.
	//
	// If multiple values are specified, jobs are retrieved from any of
	// the
	// specified locations, and, if different values are specified
	// for the LocationFilter.distance_in_miles parameter, the
	// maximum
	// provided distance is used for all locations.
	//
	// At most 5 location filters are allowed.
	LocationFilters []*LocationFilter `json:"locationFilters,omitempty"`

	// PublishDateRange: Optional.
	//
	// Jobs published within a range specified by this filter are
	// searched
	// against, for example, DateRange.PAST_MONTH. If a value
	// isn't
	// specified, all open jobs are searched against regardless of
	// their
	// published date.
	//
	// Possible values:
	//   "DATE_RANGE_UNSPECIFIED" - Default value: Filtering on time is not
	// performed.
	//   "PAST_24_HOURS" - The past 24 hours
	//   "PAST_WEEK" - The past week (7 days)
	//   "PAST_MONTH" - The past month (30 days)
	//   "PAST_YEAR" - The past year (365 days)
	//   "PAST_3_DAYS" - The past 3 days
	PublishDateRange string `json:"publishDateRange,omitempty"`

	// Query: Optional.
	//
	// The query string that matches against the job title, description,
	// and
	// location fields.
	//
	// The maximum query size is 255 bytes.
	Query string `json:"query,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Categories") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Categories") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *JobQuery) MarshalJSON() ([]byte, error) {
	type NoMethod JobQuery
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// LatLng: An object representing a latitude/longitude pair. This is
// expressed as a pair
// of doubles representing degrees latitude and degrees longitude.
// Unless
// specified otherwise, this must conform to the
// <a
// href="http://www.unoosa.org/pdf/icg/2012/template/WGS_84.pdf">WGS84
// st
// andard</a>. Values must be within normalized ranges.
type LatLng struct {
	// Latitude: The latitude in degrees. It must be in the range [-90.0,
	// +90.0].
	Latitude float64 `json:"latitude,omitempty"`

	// Longitude: The longitude in degrees. It must be in the range [-180.0,
	// +180.0].
	Longitude float64 `json:"longitude,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Latitude") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Latitude") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *LatLng) MarshalJSON() ([]byte, error) {
	type NoMethod LatLng
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *LatLng) UnmarshalJSON(data []byte) error {
	type NoMethod LatLng
	var s1 struct {
		Latitude  gensupport.JSONFloat64 `json:"latitude"`
		Longitude gensupport.JSONFloat64 `json:"longitude"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Latitude = float64(s1.Latitude)
	s.Longitude = float64(s1.Longitude)
	return nil
}

// ListCompaniesResponse: Output only.
//
// The List companies response object.
type ListCompaniesResponse struct {
	// Companies: Companies for the current client.
	Companies []*Company `json:"companies,omitempty"`

	// Metadata: Additional information for the API invocation, such as the
	// request
	// tracking id.
	Metadata *ResponseMetadata `json:"metadata,omitempty"`

	// NextPageToken: A token to retrieve the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Companies") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Companies") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListCompaniesResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListCompaniesResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListCompanyJobsResponse: Deprecated. Use ListJobsResponse
// instead.
//
// Output only.
//
// The List jobs response object.
type ListCompanyJobsResponse struct {
	// Jobs: The Jobs for a given company.
	//
	// The maximum number of items returned is based on the limit
	// field
	// provided in the request.
	Jobs []*Job `json:"jobs,omitempty"`

	// Metadata: Additional information for the API invocation, such as the
	// request
	// tracking id.
	Metadata *ResponseMetadata `json:"metadata,omitempty"`

	// NextPageToken: A token to retrieve the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// TotalSize: The total number of open jobs. The result will be
	// empty if ListCompanyJobsRequest.include_jobs_count is not enabled
	// or if no open jobs are available.
	TotalSize int64 `json:"totalSize,omitempty,string"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Jobs") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Jobs") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListCompanyJobsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListCompanyJobsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListJobsResponse: Output only.
//
// List jobs response.
type ListJobsResponse struct {
	// Jobs: The Jobs for a given company.
	//
	// The maximum number of items returned is based on the limit
	// field
	// provided in the request.
	Jobs []*Job `json:"jobs,omitempty"`

	// Metadata: Additional information for the API invocation, such as the
	// request
	// tracking id.
	Metadata *ResponseMetadata `json:"metadata,omitempty"`

	// NextPageToken: A token to retrieve the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Jobs") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Jobs") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListJobsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListJobsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// LocationFilter: Input only.
//
// Geographic region of the search.
type LocationFilter struct {
	// DistanceInMiles: Optional.
	//
	//
	// The distance_in_miles is applied when the location being searched for
	// is
	// identified as a city or smaller. When the location being searched for
	// is a
	// state or larger, this field is ignored.
	DistanceInMiles float64 `json:"distanceInMiles,omitempty"`

	// IsTelecommute: Optional.
	//
	// Allows the client to return jobs without a
	// set location, specifically, telecommuting jobs (telecomuting is
	// considered
	// by the service as a special location.
	// Job.allow_telecommute indicates if a job permits telecommuting.
	// If this field is true, telecommuting jobs are searched, and
	// name and lat_lng are
	// ignored.
	// This filter can be used by itself to search exclusively for
	// telecommuting
	// jobs, or it can be combined with another location
	// filter to search for a combination of job locations,
	// such as "Mountain View" or "telecommuting" jobs. However, when used
	// in
	// combination with other location filters, telecommuting jobs can
	// be
	// treated as less relevant than other jobs in the search response.
	IsTelecommute bool `json:"isTelecommute,omitempty"`

	// LatLng: Optional.
	//
	// The latitude and longitude of the geographic center from which
	// to
	// search. This field is ignored if `location_name` is provided.
	LatLng *LatLng `json:"latLng,omitempty"`

	// Name: Optional.
	//
	// The address name, such as "Mountain View" or "Bay Area".
	Name string `json:"name,omitempty"`

	// RegionCode: Optional.
	//
	// CLDR region code of the country/region of the address. This will be
	// used
	// to address ambiguity of the user-input location, e.g.
	// "Liverpool"
	// against "Liverpool, NY, US" or "Liverpool, UK".
	//
	// Set this field if all the jobs to search against are from a same
	// region,
	// or jobs are world-wide but the job seeker is from a specific
	// region.
	//
	// See http://cldr.unicode.org/
	// and
	// http://www.unicode.org/cldr/charts/30/supplemental/territory_infor
	// mation.html
	// for details. Example: "CH" for Switzerland.
	RegionCode string `json:"regionCode,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DistanceInMiles") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DistanceInMiles") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *LocationFilter) MarshalJSON() ([]byte, error) {
	type NoMethod LocationFilter
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *LocationFilter) UnmarshalJSON(data []byte) error {
	type NoMethod LocationFilter
	var s1 struct {
		DistanceInMiles gensupport.JSONFloat64 `json:"distanceInMiles"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.DistanceInMiles = float64(s1.DistanceInMiles)
	return nil
}

// MatchingJob: Output only.
//
// Job entry with metadata inside SearchJobsResponse.
type MatchingJob struct {
	// CommuteInfo: Commute information which is generated based on
	// specified
	//  CommutePreference.
	CommuteInfo *CommuteInfo `json:"commuteInfo,omitempty"`

	// Job: Job resource that matches the specified SearchJobsRequest.
	Job *Job `json:"job,omitempty"`

	// JobSummary: A summary of the job with core information that's
	// displayed on the search
	// results listing page.
	JobSummary string `json:"jobSummary,omitempty"`

	// JobTitleSnippet: Contains snippets of text from the Job.job_title
	// field most
	// closely matching a search query's keywords, if available. The
	// matching query
	// keywords are enclosed in HTML bold tags.
	JobTitleSnippet string `json:"jobTitleSnippet,omitempty"`

	// SearchTextSnippet: Contains snippets of text from the Job.description
	// and similar
	// fields that most closely match a search query's keywords, if
	// available.
	// All HTML tags in the original fields are stripped when returned in
	// this
	// field, and matching query keywords are enclosed in HTML bold tags.
	SearchTextSnippet string `json:"searchTextSnippet,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CommuteInfo") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CommuteInfo") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *MatchingJob) MarshalJSON() ([]byte, error) {
	type NoMethod MatchingJob
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Money: Represents an amount of money with its currency type.
type Money struct {
	// CurrencyCode: The 3-letter currency code defined in ISO 4217.
	CurrencyCode string `json:"currencyCode,omitempty"`

	// Nanos: Number of nano (10^-9) units of the amount.
	// The value must be between -999,999,999 and +999,999,999 inclusive.
	// If `units` is positive, `nanos` must be positive or zero.
	// If `units` is zero, `nanos` can be positive, zero, or negative.
	// If `units` is negative, `nanos` must be negative or zero.
	// For example $-1.75 is represented as `units`=-1 and
	// `nanos`=-750,000,000.
	Nanos int64 `json:"nanos,omitempty"`

	// Units: The whole units of the amount.
	// For example if `currencyCode` is "USD", then 1 unit is one US
	// dollar.
	Units int64 `json:"units,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "CurrencyCode") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CurrencyCode") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Money) MarshalJSON() ([]byte, error) {
	type NoMethod Money
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// NumericBucketingOption: Input only.
//
// Use this field to specify bucketing option for the histogram search
// response.
type NumericBucketingOption struct {
	// BucketBounds: Required.
	//
	// Two adjacent values form a histogram bucket. Values should be
	// in
	// ascending order. For example, if [5, 10, 15] are provided, four
	// buckets are
	// created: (-inf, 5), 5, 10), [10, 15), [15, inf). At most
	// 20
	// [buckets_bound is supported.
	BucketBounds []float64 `json:"bucketBounds,omitempty"`

	// RequiresMinMax: Optional.
	//
	// If set to true, the histogram result includes minimum/maximum
	// value of the numeric field.
	RequiresMinMax bool `json:"requiresMinMax,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BucketBounds") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "BucketBounds") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *NumericBucketingOption) MarshalJSON() ([]byte, error) {
	type NoMethod NumericBucketingOption
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// NumericBucketingResult: Output only.
//
// Custom numeric bucketing result.
type NumericBucketingResult struct {
	// Counts: Count within each bucket. Its size is the length
	// of
	// NumericBucketingOption.bucket_bounds plus 1.
	Counts []*BucketizedCount `json:"counts,omitempty"`

	// MaxValue: Stores the maximum value of the numeric field. Will be
	// populated only if
	// [NumericBucketingOption.requires_min_max] is set to true.
	MaxValue float64 `json:"maxValue,omitempty"`

	// MinValue: Stores the minimum value of the numeric field. Will be
	// populated only if
	// [NumericBucketingOption.requires_min_max] is set to true.
	MinValue float64 `json:"minValue,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Counts") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Counts") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *NumericBucketingResult) MarshalJSON() ([]byte, error) {
	type NoMethod NumericBucketingResult
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *NumericBucketingResult) UnmarshalJSON(data []byte) error {
	type NoMethod NumericBucketingResult
	var s1 struct {
		MaxValue gensupport.JSONFloat64 `json:"maxValue"`
		MinValue gensupport.JSONFloat64 `json:"minValue"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.MaxValue = float64(s1.MaxValue)
	s.MinValue = float64(s1.MinValue)
	return nil
}

// PostalAddress: Represents a postal address, e.g. for postal delivery
// or payments addresses.
// Given a postal address, a postal service can deliver items to a
// premise, P.O.
// Box or similar.
// It is not intended to model geographical locations (roads,
// towns,
// mountains).
//
// In typical usage an address would be created via user input or from
// importing
// existing data, depending on the type of process.
//
// Advice on address input / editing:
//  - Use an i18n-ready address widget such as
//    https://github.com/googlei18n/libaddressinput)
// - Users should not be presented with UI elements for input or editing
// of
//   fields outside countries where that field is used.
//
// For more guidance on how to use this schema, please
// see:
// https://support.google.com/business/answer/6397478
type PostalAddress struct {
	// AddressLines: Unstructured address lines describing the lower levels
	// of an address.
	//
	// Because values in address_lines do not have type information and
	// may
	// sometimes contain multiple values in a single field (e.g.
	// "Austin, TX"), it is important that the line order is clear. The
	// order of
	// address lines should be "envelope order" for the country/region of
	// the
	// address. In places where this can vary (e.g. Japan), address_language
	// is
	// used to make it explicit (e.g. "ja" for large-to-small ordering
	// and
	// "ja-Latn" or "en" for small-to-large). This way, the most specific
	// line of
	// an address can be selected based on the language.
	//
	// The minimum permitted structural representation of an address
	// consists
	// of a region_code with all remaining information placed in
	// the
	// address_lines. It would be possible to format such an address
	// very
	// approximately without geocoding, but no semantic reasoning could
	// be
	// made about any of the address components until it was at
	// least
	// partially resolved.
	//
	// Creating an address only containing a region_code and address_lines,
	// and
	// then geocoding is the recommended way to handle completely
	// unstructured
	// addresses (as opposed to guessing which parts of the address should
	// be
	// localities or administrative areas).
	AddressLines []string `json:"addressLines,omitempty"`

	// AdministrativeArea: Optional. Highest administrative subdivision
	// which is used for postal
	// addresses of a country or region.
	// For example, this can be a state, a province, an oblast, or a
	// prefecture.
	// Specifically, for Spain this is the province and not the
	// autonomous
	// community (e.g. "Barcelona" and not "Catalonia").
	// Many countries don't use an administrative area in postal addresses.
	// E.g.
	// in Switzerland this should be left unpopulated.
	AdministrativeArea string `json:"administrativeArea,omitempty"`

	// LanguageCode: Optional. BCP-47 language code of the contents of this
	// address (if
	// known). This is often the UI language of the input form or is
	// expected
	// to match one of the languages used in the address' country/region, or
	// their
	// transliterated equivalents.
	// This can affect formatting in certain countries, but is not
	// critical
	// to the correctness of the data and will never affect any validation
	// or
	// other non-formatting related operations.
	//
	// If this value is not known, it should be omitted (rather than
	// specifying a
	// possibly incorrect default).
	//
	// Examples: "zh-Hant", "ja", "ja-Latn", "en".
	LanguageCode string `json:"languageCode,omitempty"`

	// Locality: Optional. Generally refers to the city/town portion of the
	// address.
	// Examples: US city, IT comune, UK post town.
	// In regions of the world where localities are not well defined or do
	// not fit
	// into this structure well, leave locality empty and use address_lines.
	Locality string `json:"locality,omitempty"`

	// Organization: Optional. The name of the organization at the address.
	Organization string `json:"organization,omitempty"`

	// PostalCode: Optional. Postal code of the address. Not all countries
	// use or require
	// postal codes to be present, but where they are used, they may
	// trigger
	// additional validation with other parts of the address (e.g.
	// state/zip
	// validation in the U.S.A.).
	PostalCode string `json:"postalCode,omitempty"`

	// Recipients: Optional. The recipient at the address.
	// This field may, under certain circumstances, contain multiline
	// information.
	// For example, it might contain "care of" information.
	Recipients []string `json:"recipients,omitempty"`

	// RegionCode: Required. CLDR region code of the country/region of the
	// address. This
	// is never inferred and it is up to the user to ensure the value
	// is
	// correct. See http://cldr.unicode.org/
	// and
	// http://www.unicode.org/cldr/charts/30/supplemental/territory_infor
	// mation.html
	// for details. Example: "CH" for Switzerland.
	RegionCode string `json:"regionCode,omitempty"`

	// Revision: The schema revision of the `PostalAddress`. This must be
	// set to 0, which is
	// the latest revision.
	//
	// All new revisions **must** be backward compatible with old revisions.
	Revision int64 `json:"revision,omitempty"`

	// SortingCode: Optional. Additional, country-specific, sorting code.
	// This is not used
	// in most regions. Where it is used, the value is either a string
	// like
	// "CEDEX", optionally followed by a number (e.g. "CEDEX 7"), or just a
	// number
	// alone, representing the "sector code" (Jamaica), "delivery area
	// indicator"
	// (Malawi) or "post office indicator" (e.g. Côte d'Ivoire).
	SortingCode string `json:"sortingCode,omitempty"`

	// Sublocality: Optional. Sublocality of the address.
	// For example, this can be neighborhoods, boroughs, districts.
	Sublocality string `json:"sublocality,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AddressLines") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AddressLines") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PostalAddress) MarshalJSON() ([]byte, error) {
	type NoMethod PostalAddress
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// RequestMetadata: Input only.
//
// Meta information related to the job searcher or entity
// conducting the job search. This information is used to improve
// the
// performance of the service.
type RequestMetadata struct {
	// DeviceInfo: Optional.
	//
	// The type of device used by the job seeker at the time of the call to
	// the
	// service.
	DeviceInfo *DeviceInfo `json:"deviceInfo,omitempty"`

	// Domain: Required.
	//
	// The client-defined scope or source of the service call, which
	// typically
	// is the domain on
	// which the service has been implemented and is currently being
	// run.
	//
	// For example, if the service is being run by client <em>Foo,
	// Inc.</em>, on
	// job board www.foo.com and career site www.bar.com, then this field
	// is
	// set to "foo.com" for use on the job board, and "bar.com" for use on
	// the
	// career site.
	//
	// If this field is not available for some reason, send "UNKNOWN". Note
	// that
	// any improvements to the service model for a particular tenant site
	// rely on
	// this field being set correctly to some domain.
	Domain string `json:"domain,omitempty"`

	// SessionId: Required.
	//
	// A unique session identification string. A session is defined as
	// the
	// duration of an end user's interaction with the service over a
	// period.
	// Obfuscate this field for privacy concerns before
	// providing it to the API.
	//
	// If this field is not available for some reason, please send
	// "UNKNOWN". Note
	// that any improvements to the service model for a particular tenant
	// site,
	// rely on this field being set correctly to some unique session_id.
	SessionId string `json:"sessionId,omitempty"`

	// UserId: Required.
	//
	// A unique user identification string, as determined by the client.
	// The
	// client is responsible for ensuring client-level uniqueness of this
	// value
	// in order to have the strongest positive impact on search
	// quality.
	// Obfuscate this field for privacy concerns before
	// providing it to the service.
	//
	// If this field is not available for some reason, please send
	// "UNKNOWN". Note
	// that any improvements to the service model for a particular tenant
	// site,
	// rely on this field being set correctly to some unique user_id.
	UserId string `json:"userId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DeviceInfo") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DeviceInfo") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *RequestMetadata) MarshalJSON() ([]byte, error) {
	type NoMethod RequestMetadata
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ResponseMetadata: Output only.
//
// Additional information returned to client, such as
// debugging
// information.
type ResponseMetadata struct {
	// ExperimentIdList: Identifiers for the versions of the search
	// algorithm used during
	// this API invocation if multiple algorithms are used.
	// The default value is empty.
	// For search response only.
	ExperimentIdList []int64 `json:"experimentIdList,omitempty"`

	// Mode: For search response only. Indicates the mode of a performed
	// search.
	//
	// Possible values:
	//   "SEARCH_MODE_UNSPECIFIED" - The mode of the search method isn't
	// specified.
	//   "JOB_SEARCH" - The job search doesn't include support for featured
	// jobs.
	//   "FEATURED_JOB_SEARCH" - The job search matches only against
	// featured jobs (jobs with a
	// promotionValue > 0). This method doesn't return any jobs having
	// a
	// promotionValue <= 0. The search results order is determined by
	// the
	// promotionValue (jobs with a higher promotionValue are returned higher
	// up in
	// the search results), with relevance being used as a tiebreaker.
	//   "EMAIL_ALERT_SEARCH" - Deprecated. Please use the
	// SearchJobsForAlert API.
	//
	// The job search matches against jobs suited to email notifications.
	Mode string `json:"mode,omitempty"`

	// RequestId: A unique id associated with this call.
	// This id is logged for tracking purposes.
	RequestId string `json:"requestId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ExperimentIdList") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ExperimentIdList") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ResponseMetadata) MarshalJSON() ([]byte, error) {
	type NoMethod ResponseMetadata
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SearchJobsRequest: Input only.
//
// The Request body of the `SearchJobs` call.
type SearchJobsRequest struct {
	// DisableRelevanceThresholding: Deprecated. Any value provided in this
	// field is ignored.
	//
	// Optional.
	//
	// Controls whether to disable relevance thresholding.
	// Relevance
	// thresholding removes jobs that have low relevance in search
	// results,
	// for example, removing "Assistant to the CEO" positions from the
	// search
	// results of a search for "CEO".
	//
	// Disabling relevance thresholding improves the accuracy of
	// subsequent
	// search requests.
	//
	// Defaults to false.
	DisableRelevanceThresholding bool `json:"disableRelevanceThresholding,omitempty"`

	// EnableBroadening: Optional.
	//
	// Controls whether to broaden the search when it produces sparse
	// results.
	// Broadened queries append results to the end of the matching
	// results
	// list.
	//
	// Defaults to false.
	EnableBroadening bool `json:"enableBroadening,omitempty"`

	// EnablePreciseResultSize: Optional.
	//
	// Controls if the search job request requires the return of a
	// precise
	// count of the first 300 results. Setting this to `true`
	// ensures
	// consistency in the number of results per page. Best practice is to
	// set this
	// value to true if a client allows users to jump directly to
	// a
	// non-sequential search results page.
	//
	// Enabling this flag may adversely impact performance.
	//
	// Defaults to false.
	EnablePreciseResultSize bool `json:"enablePreciseResultSize,omitempty"`

	// Filters: Deprecated. Use query instead.
	//
	// Optional.
	//
	// Restrictions on the scope of the search request, such as filtering
	// by location.
	Filters *JobFilters `json:"filters,omitempty"`

	// HistogramFacets: Optional.
	//
	// Restrictions on what fields to perform histogram on, such
	// as
	// `COMPANY_SIZE` etc.
	HistogramFacets *HistogramFacets `json:"histogramFacets,omitempty"`

	// JobView: Optional.
	//
	// The number of job attributes returned for jobs in the
	// search response. Defaults to JobView.SMALL if no value is specified.
	//
	// Possible values:
	//   "JOB_VIEW_UNSPECIFIED" - Default value.
	//   "SMALL" - A small view of the job, with the following attributes in
	// the search results:
	// Job.name, Job.requisition_id, Job.job_title,
	// Job.company_name, Job.job_locations,
	// Job.description,
	// Job.visibility.
	// Note: Job.description is deprecated. It is scheduled to be
	// removed
	// from MatchingJob.Job objects in the SearchJobsResponse results
	// on 12/31/2018.
	//   "MINIMAL" - A minimal view of the job, with the following
	// attributes in the search
	// results: Job.name, Job.requisition_id,
	// Job.job_title,
	// Job.company_name, Job.job_locations.
	//   "FULL" - All available attributes are included in the search
	// results.
	// Note: [Job.description, Job.responsibilities,
	// Job.qualifications and Job.incentives are deprecated.
	// These fields are scheduled to be removed from MatchingJob.Job
	// objects
	// in the SearchJobsResponse results on 12/31/2018.
	// See the alternative MatchingJob.search_text_snippet
	// and
	// MatchingJob.job_summary fields.
	JobView string `json:"jobView,omitempty"`

	// Mode: Required.
	//
	// Mode of a search.
	//
	// Possible values:
	//   "SEARCH_MODE_UNSPECIFIED" - The mode of the search method isn't
	// specified.
	//   "JOB_SEARCH" - The job search doesn't include support for featured
	// jobs.
	//   "FEATURED_JOB_SEARCH" - The job search matches only against
	// featured jobs (jobs with a
	// promotionValue > 0). This method doesn't return any jobs having
	// a
	// promotionValue <= 0. The search results order is determined by
	// the
	// promotionValue (jobs with a higher promotionValue are returned higher
	// up in
	// the search results), with relevance being used as a tiebreaker.
	//   "EMAIL_ALERT_SEARCH" - Deprecated. Please use the
	// SearchJobsForAlert API.
	//
	// The job search matches against jobs suited to email notifications.
	Mode string `json:"mode,omitempty"`

	// Offset: Optional.
	//
	// An integer that specifies the current offset (that is, starting
	// result location, amongst the jobs deemed by the API as relevant)
	// in
	// search results. This field is only considered if page_token is
	// unset.
	//
	// For example, 0 means to  return results starting from the first
	// matching
	// job, and 10 means to return from the 11th job. This can be used
	// for
	// pagination, (for example, pageSize = 10 and offset = 10 means to
	// return
	// from the second page).
	Offset int64 `json:"offset,omitempty"`

	// OrderBy: Deprecated. Use sort_by instead.
	//
	// Optional.
	//
	// The criteria determining how search results are sorted.
	// Defaults to SortBy.RELEVANCE_DESC if no value is specified.
	//
	// Possible values:
	//   "SORT_BY_UNSPECIFIED" - Default value.
	//   "RELEVANCE_DESC" - By descending relevance, as determined by the
	// API algorithms.
	//
	// Relevance thresholding of query results is only available for queries
	// if
	// RELEVANCE_DESC sort ordering is specified.
	//   "PUBLISHED_DATE_DESC" - Sort by published date descending.
	//   "UPDATED_DATE_DESC" - Sort by updated date descending.
	//   "TITLE" - Sort by job title ascending.
	//   "TITLE_DESC" - Sort by job title descending.
	//   "ANNUALIZED_BASE_COMPENSATION" - Sort by job annualized base
	// compensation in ascending order.
	// If job's annualized base compensation is unspecified, they are put
	// at
	// the end of search result.
	//   "ANNUALIZED_TOTAL_COMPENSATION" - Sort by job annualized total
	// compensation in ascending order.
	// If job's annualized total compensation is unspecified, they are put
	// at
	// the end of search result.
	//   "ANNUALIZED_BASE_COMPENSATION_DESC" - Sort by job annualized base
	// compensation in descending order.
	// If job's annualized base compensation is unspecified, they are put
	// at
	// the end of search result.
	//   "ANNUALIZED_TOTAL_COMPENSATION_DESC" - Sort by job annualized total
	// compensation in descending order.
	// If job's annualized total compensation is unspecified, they are put
	// at
	// the end of search result.
	OrderBy string `json:"orderBy,omitempty"`

	// PageSize: Optional.
	//
	// A limit on the number of jobs returned in the search
	// results.
	// Increasing this value above the default value of 10 can increase
	// search
	// response time. The value can be between 1 and 100.
	PageSize int64 `json:"pageSize,omitempty"`

	// PageToken: Optional.
	//
	// The token specifying the current offset within
	// search results. See SearchJobsResponse.next_page_token for
	// an explanation of how to obtain the next set of query results.
	PageToken string `json:"pageToken,omitempty"`

	// Query: Optional.
	//
	// Query used to search against jobs, such as keyword, location filters,
	// etc.
	Query *JobQuery `json:"query,omitempty"`

	// RequestMetadata: Required.
	//
	// The meta information collected about the job searcher, used to
	// improve the
	// search quality of the service.. The identifiers, (such as `user_id`)
	// are
	// provided by users, and must be unique and consistent.
	RequestMetadata *RequestMetadata `json:"requestMetadata,omitempty"`

	// SortBy: Optional.
	//
	// The criteria determining how search results are sorted.
	// Defaults to SortBy.RELEVANCE_DESC if no value is specified.
	//
	// Possible values:
	//   "SORT_BY_UNSPECIFIED" - Default value.
	//   "RELEVANCE_DESC" - By descending relevance, as determined by the
	// API algorithms.
	//
	// Relevance thresholding of query results is only available for queries
	// if
	// RELEVANCE_DESC sort ordering is specified.
	//   "PUBLISHED_DATE_DESC" - Sort by published date descending.
	//   "UPDATED_DATE_DESC" - Sort by updated date descending.
	//   "TITLE" - Sort by job title ascending.
	//   "TITLE_DESC" - Sort by job title descending.
	//   "ANNUALIZED_BASE_COMPENSATION" - Sort by job annualized base
	// compensation in ascending order.
	// If job's annualized base compensation is unspecified, they are put
	// at
	// the end of search result.
	//   "ANNUALIZED_TOTAL_COMPENSATION" - Sort by job annualized total
	// compensation in ascending order.
	// If job's annualized total compensation is unspecified, they are put
	// at
	// the end of search result.
	//   "ANNUALIZED_BASE_COMPENSATION_DESC" - Sort by job annualized base
	// compensation in descending order.
	// If job's annualized base compensation is unspecified, they are put
	// at
	// the end of search result.
	//   "ANNUALIZED_TOTAL_COMPENSATION_DESC" - Sort by job annualized total
	// compensation in descending order.
	// If job's annualized total compensation is unspecified, they are put
	// at
	// the end of search result.
	SortBy string `json:"sortBy,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "DisableRelevanceThresholding") to unconditionally include in API
	// requests. By default, fields with empty values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g.
	// "DisableRelevanceThresholding") to include in API requests with the
	// JSON null value. By default, fields with empty values are omitted
	// from API requests. However, any field with an empty value appearing
	// in NullFields will be sent to the server as null. It is an error if a
	// field in this list has a non-empty value. This may be used to include
	// null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *SearchJobsRequest) MarshalJSON() ([]byte, error) {
	type NoMethod SearchJobsRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SearchJobsResponse: Output only.
//
// Response for SearchJob method.
type SearchJobsResponse struct {
	// AppliedCommuteFilter: The commute filter the service applied to the
	// specified query. This
	// information is only available when query has a valid
	// CommutePreference.
	AppliedCommuteFilter *CommutePreference `json:"appliedCommuteFilter,omitempty"`

	// AppliedJobLocationFilters: The location filters that the service
	// applied to the specified query. If
	// any filters are lat-lng based, the JobLocation.location_type
	// is
	// JobLocation.LocationType#LOCATION_TYPE_UNSPECIFIED.
	AppliedJobLocationFilters []*JobLocation `json:"appliedJobLocationFilters,omitempty"`

	// EstimatedTotalSize: An estimation of the number of jobs that match
	// the specified query.
	//
	// This number is not guaranteed to be accurate. For accurate
	// results,
	// seenenable_precise_result_size.
	EstimatedTotalSize int64 `json:"estimatedTotalSize,omitempty,string"`

	// HistogramResults: The histogram results that match with
	// specified
	// SearchJobsRequest.HistogramFacets.
	HistogramResults *HistogramResults `json:"histogramResults,omitempty"`

	// JobView: Corresponds to SearchJobsRequest.job_view.
	//
	// Possible values:
	//   "JOB_VIEW_UNSPECIFIED" - Default value.
	//   "SMALL" - A small view of the job, with the following attributes in
	// the search results:
	// Job.name, Job.requisition_id, Job.job_title,
	// Job.company_name, Job.job_locations,
	// Job.description,
	// Job.visibility.
	// Note: Job.description is deprecated. It is scheduled to be
	// removed
	// from MatchingJob.Job objects in the SearchJobsResponse results
	// on 12/31/2018.
	//   "MINIMAL" - A minimal view of the job, with the following
	// attributes in the search
	// results: Job.name, Job.requisition_id,
	// Job.job_title,
	// Job.company_name, Job.job_locations.
	//   "FULL" - All available attributes are included in the search
	// results.
	// Note: [Job.description, Job.responsibilities,
	// Job.qualifications and Job.incentives are deprecated.
	// These fields are scheduled to be removed from MatchingJob.Job
	// objects
	// in the SearchJobsResponse results on 12/31/2018.
	// See the alternative MatchingJob.search_text_snippet
	// and
	// MatchingJob.job_summary fields.
	JobView string `json:"jobView,omitempty"`

	// MatchingJobs: The Job entities that match the specified
	// SearchJobsRequest.
	MatchingJobs []*MatchingJob `json:"matchingJobs,omitempty"`

	// Metadata: Additional information for the API invocation, such as the
	// request
	// tracking id.
	Metadata *ResponseMetadata `json:"metadata,omitempty"`

	// NextPageToken: The token that specifies the starting position of the
	// next page of results.
	// This field is empty if there are no more results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// NumJobsFromBroadenedQuery: If query broadening is enabled, we may
	// append additional results from the
	// broadened query. This number indicates how many of the jobs returned
	// in the
	// jobs field are from the broadened query. These results are always at
	// the
	// end of the jobs list. In particular, a value of 0 means all the jobs
	// in the
	// jobs list are from the original (without broadening) query. If
	// this
	// field is non-zero, subsequent requests with offset after this result
	// set
	// should contain all broadened results.
	NumJobsFromBroadenedQuery int64 `json:"numJobsFromBroadenedQuery,omitempty"`

	// SpellResult: The spell checking result, and correction.
	SpellResult *SpellingCorrection `json:"spellResult,omitempty"`

	// TotalSize: The precise result count, which is available only if the
	// client set
	// enable_precise_result_size to `true` or if the response
	// is the last page of results. Otherwise, the value will be `-1`.
	TotalSize int64 `json:"totalSize,omitempty,string"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g.
	// "AppliedCommuteFilter") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AppliedCommuteFilter") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *SearchJobsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod SearchJobsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SpellingCorrection: Output only.
//
// Spell check result.
type SpellingCorrection struct {
	// Corrected: Indicates if the query was corrected by the spell checker.
	Corrected bool `json:"corrected,omitempty"`

	// CorrectedText: Correction output consisting of the corrected keyword
	// string.
	CorrectedText string `json:"correctedText,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Corrected") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Corrected") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *SpellingCorrection) MarshalJSON() ([]byte, error) {
	type NoMethod SpellingCorrection
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// StringValues: Represents array of string values.
type StringValues struct {
	// Values: Required.
	//
	// String values.
	Values []string `json:"values,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Values") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Values") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *StringValues) MarshalJSON() ([]byte, error) {
	type NoMethod StringValues
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// UpdateJobRequest: Input only.
//
// Update job request.
type UpdateJobRequest struct {
	// DisableStreetAddressResolution: Deprecated. Please use
	// processing_options. This flag is ignored if
	// processing_options is set.
	//
	// Optional.
	//
	// If set to `true`, the service does not attempt resolve a more
	// precise
	// address for the job.
	DisableStreetAddressResolution bool `json:"disableStreetAddressResolution,omitempty"`

	// Job: Required.
	//
	// The Job to be updated.
	Job *Job `json:"job,omitempty"`

	// ProcessingOptions: Optional.
	//
	// Options for job
	// processing.
	//
	// UpdateJobRequest.disable_street_address_resolution is ignored if
	// this
	// flag is set.
	ProcessingOptions *JobProcessingOptions `json:"processingOptions,omitempty"`

	// UpdateJobFields: Optional but strongly recommended to be provided for
	// the best service
	// experience.
	//
	// If update_job_fields is provided, only the specified fields in
	// job are updated. Otherwise all the fields are updated.
	//
	// A field mask to restrict the fields that are updated. Valid values
	// are:
	//
	// * jobTitle
	// * employmentTypes
	// * description
	// * applicationUrls
	// * applicationEmailList
	// * applicationInstruction
	// * responsibilities
	// * qualifications
	// * educationLevels
	// * level
	// * department
	// * startDate
	// * endDate
	// * compensationInfo
	// * incentives
	// * languageCode
	// * benefits
	// * expireTime
	// * customAttributes
	// * visibility
	// * publishDate
	// * promotionValue
	// * locations
	// * region
	// * expiryDate (deprecated)
	// * filterableCustomFields (deprecated)
	// * unindexedCustomFields (deprecated)
	UpdateJobFields string `json:"updateJobFields,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "DisableStreetAddressResolution") to unconditionally include in API
	// requests. By default, fields with empty values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g.
	// "DisableStreetAddressResolution") to include in API requests with the
	// JSON null value. By default, fields with empty values are omitted
	// from API requests. However, any field with an empty value appearing
	// in NullFields will be sent to the server as null. It is an error if a
	// field in this list has a non-empty value. This may be used to include
	// null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *UpdateJobRequest) MarshalJSON() ([]byte, error) {
	type NoMethod UpdateJobRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "jobs.companies.create":

type CompaniesCreateCall struct {
	s          *Service
	company    *Company
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Create: Creates a new company entity.
func (r *CompaniesService) Create(company *Company) *CompaniesCreateCall {
	c := &CompaniesCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.company = company
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CompaniesCreateCall) Fields(s ...googleapi.Field) *CompaniesCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *CompaniesCreateCall) Context(ctx context.Context) *CompaniesCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *CompaniesCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *CompaniesCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.company)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/companies")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "jobs.companies.create" call.
// Exactly one of *Company or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Company.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *CompaniesCreateCall) Do(opts ...googleapi.CallOption) (*Company, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Company{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a new company entity.",
	//   "flatPath": "v2/companies",
	//   "httpMethod": "POST",
	//   "id": "jobs.companies.create",
	//   "parameterOrder": [],
	//   "parameters": {},
	//   "path": "v2/companies",
	//   "request": {
	//     "$ref": "Company"
	//   },
	//   "response": {
	//     "$ref": "Company"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/jobs"
	//   ]
	// }

}

// method id "jobs.companies.delete":

type CompaniesDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes the specified company.
func (r *CompaniesService) Delete(name string) *CompaniesDeleteCall {
	c := &CompaniesDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CompaniesDeleteCall) Fields(s ...googleapi.Field) *CompaniesDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *CompaniesDeleteCall) Context(ctx context.Context) *CompaniesDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *CompaniesDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *CompaniesDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "jobs.companies.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *CompaniesDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes the specified company.",
	//   "flatPath": "v2/companies/{companiesId}",
	//   "httpMethod": "DELETE",
	//   "id": "jobs.companies.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required.\n\nThe resource name of the company to be deleted,\nsuch as, \"companies/0000aaaa-1111-bbbb-2222-cccc3333dddd\".",
	//       "location": "path",
	//       "pattern": "^companies/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/{+name}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/jobs"
	//   ]
	// }

}

// method id "jobs.companies.get":

type CompaniesGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Retrieves the specified company.
func (r *CompaniesService) Get(name string) *CompaniesGetCall {
	c := &CompaniesGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CompaniesGetCall) Fields(s ...googleapi.Field) *CompaniesGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *CompaniesGetCall) IfNoneMatch(entityTag string) *CompaniesGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *CompaniesGetCall) Context(ctx context.Context) *CompaniesGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *CompaniesGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *CompaniesGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "jobs.companies.get" call.
// Exactly one of *Company or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Company.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *CompaniesGetCall) Do(opts ...googleapi.CallOption) (*Company, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Company{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the specified company.",
	//   "flatPath": "v2/companies/{companiesId}",
	//   "httpMethod": "GET",
	//   "id": "jobs.companies.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required.\n\nResource name of the company to retrieve,\nsuch as \"companies/0000aaaa-1111-bbbb-2222-cccc3333dddd\".",
	//       "location": "path",
	//       "pattern": "^companies/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/{+name}",
	//   "response": {
	//     "$ref": "Company"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/jobs"
	//   ]
	// }

}

// method id "jobs.companies.list":

type CompaniesListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists all companies associated with a Cloud Talent Solution
// account.
func (r *CompaniesService) List() *CompaniesListCall {
	c := &CompaniesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// MustHaveOpenJobs sets the optional parameter "mustHaveOpenJobs": Set
// to true if the companies request must have open jobs.
//
// Defaults to false.
//
// If true, at most page_size of companies are fetched, among which
// only those with open jobs are returned.
func (c *CompaniesListCall) MustHaveOpenJobs(mustHaveOpenJobs bool) *CompaniesListCall {
	c.urlParams_.Set("mustHaveOpenJobs", fmt.Sprint(mustHaveOpenJobs))
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of companies to be returned, at most 100.
// Default is 100 if a non-positive number is provided.
func (c *CompaniesListCall) PageSize(pageSize int64) *CompaniesListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The starting
// indicator from which to return results.
func (c *CompaniesListCall) PageToken(pageToken string) *CompaniesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CompaniesListCall) Fields(s ...googleapi.Field) *CompaniesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *CompaniesListCall) IfNoneMatch(entityTag string) *CompaniesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *CompaniesListCall) Context(ctx context.Context) *CompaniesListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *CompaniesListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *CompaniesListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/companies")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "jobs.companies.list" call.
// Exactly one of *ListCompaniesResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListCompaniesResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *CompaniesListCall) Do(opts ...googleapi.CallOption) (*ListCompaniesResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListCompaniesResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists all companies associated with a Cloud Talent Solution account.",
	//   "flatPath": "v2/companies",
	//   "httpMethod": "GET",
	//   "id": "jobs.companies.list",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "mustHaveOpenJobs": {
	//       "description": "Optional.\n\nSet to true if the companies request must have open jobs.\n\nDefaults to false.\n\nIf true, at most page_size of companies are fetched, among which\nonly those with open jobs are returned.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "pageSize": {
	//       "description": "Optional.\n\nThe maximum number of companies to be returned, at most 100.\nDefault is 100 if a non-positive number is provided.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional.\n\nThe starting indicator from which to return results.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/companies",
	//   "response": {
	//     "$ref": "ListCompaniesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/jobs"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *CompaniesListCall) Pages(ctx context.Context, f func(*ListCompaniesResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "jobs.companies.patch":

type CompaniesPatchCall struct {
	s          *Service
	name       string
	company    *Company
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Patch: Updates the specified company. Company names can't be updated.
// To update a
// company name, delete the company and all jobs associated with it, and
// only
// then re-create them.
func (r *CompaniesService) Patch(name string, company *Company) *CompaniesPatchCall {
	c := &CompaniesPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.company = company
	return c
}

// UpdateCompanyFields sets the optional parameter
// "updateCompanyFields": Optional but strongly recommended to be
// provided for the best service
// experience.
//
// If update_company_fields is provided, only the specified fields
// in
// company are updated. Otherwise all the fields are updated.
//
// A field mask to specify the company fields to update. Valid values
// are:
//
// * displayName
// * website
// * imageUrl
// * companySize
// * distributorBillingCompanyId
// * companyInfoSources
// * careerPageLink
// * hiringAgency
// * hqLocation
// * eeoText
// * keywordSearchableCustomAttributes
// * title (deprecated)
// * keywordSearchableCustomFields (deprecated)
func (c *CompaniesPatchCall) UpdateCompanyFields(updateCompanyFields string) *CompaniesPatchCall {
	c.urlParams_.Set("updateCompanyFields", updateCompanyFields)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CompaniesPatchCall) Fields(s ...googleapi.Field) *CompaniesPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *CompaniesPatchCall) Context(ctx context.Context) *CompaniesPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *CompaniesPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *CompaniesPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.company)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "jobs.companies.patch" call.
// Exactly one of *Company or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Company.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *CompaniesPatchCall) Do(opts ...googleapi.CallOption) (*Company, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Company{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates the specified company. Company names can't be updated. To update a\ncompany name, delete the company and all jobs associated with it, and only\nthen re-create them.",
	//   "flatPath": "v2/companies/{companiesId}",
	//   "httpMethod": "PATCH",
	//   "id": "jobs.companies.patch",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required during company update.\n\nThe resource name for a company. This is generated by the service when a\ncompany is created, for example,\n\"companies/0000aaaa-1111-bbbb-2222-cccc3333dddd\".",
	//       "location": "path",
	//       "pattern": "^companies/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "updateCompanyFields": {
	//       "description": "Optional but strongly recommended to be provided for the best service\nexperience.\n\nIf update_company_fields is provided, only the specified fields in\ncompany are updated. Otherwise all the fields are updated.\n\nA field mask to specify the company fields to update. Valid values are:\n\n* displayName\n* website\n* imageUrl\n* companySize\n* distributorBillingCompanyId\n* companyInfoSources\n* careerPageLink\n* hiringAgency\n* hqLocation\n* eeoText\n* keywordSearchableCustomAttributes\n* title (deprecated)\n* keywordSearchableCustomFields (deprecated)",
	//       "format": "google-fieldmask",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/{+name}",
	//   "request": {
	//     "$ref": "Company"
	//   },
	//   "response": {
	//     "$ref": "Company"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/jobs"
	//   ]
	// }

}

// method id "jobs.companies.jobs.list":

type CompaniesJobsListCall struct {
	s            *Service
	companyName  string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Deprecated. Use ListJobs instead.
//
// Lists all jobs associated with a company.
func (r *CompaniesJobsService) List(companyName string) *CompaniesJobsListCall {
	c := &CompaniesJobsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.companyName = companyName
	return c
}

// IdsOnly sets the optional parameter "idsOnly": If set to `true`, only
// job ID, job requisition ID and language code will be
// returned.
//
// A typical use is to synchronize job repositories.
//
// Defaults to false.
func (c *CompaniesJobsListCall) IdsOnly(idsOnly bool) *CompaniesJobsListCall {
	c.urlParams_.Set("idsOnly", fmt.Sprint(idsOnly))
	return c
}

// IncludeJobsCount sets the optional parameter "includeJobsCount":
// Deprecated. Please DO NOT use this field except for small
// companies.
// Suggest counting jobs page by page instead.
//
//
//
// Set to true if the total number of open jobs is to be
// returned.
//
// Defaults to false.
func (c *CompaniesJobsListCall) IncludeJobsCount(includeJobsCount bool) *CompaniesJobsListCall {
	c.urlParams_.Set("includeJobsCount", fmt.Sprint(includeJobsCount))
	return c
}

// JobRequisitionId sets the optional parameter "jobRequisitionId": The
// requisition ID, also known as posting ID, assigned by the company
// to the job.
//
// The maximum number of allowable characters is 225.
func (c *CompaniesJobsListCall) JobRequisitionId(jobRequisitionId string) *CompaniesJobsListCall {
	c.urlParams_.Set("jobRequisitionId", jobRequisitionId)
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of jobs to be returned per page of results.
//
// If ids_only is set to true, the maximum allowed page size
// is 1000. Otherwise, the maximum allowed page size is 100.
//
// Default is 100 if empty or a number < 1 is specified.
func (c *CompaniesJobsListCall) PageSize(pageSize int64) *CompaniesJobsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The starting point
// of a query result.
func (c *CompaniesJobsListCall) PageToken(pageToken string) *CompaniesJobsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CompaniesJobsListCall) Fields(s ...googleapi.Field) *CompaniesJobsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *CompaniesJobsListCall) IfNoneMatch(entityTag string) *CompaniesJobsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *CompaniesJobsListCall) Context(ctx context.Context) *CompaniesJobsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *CompaniesJobsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *CompaniesJobsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/{+companyName}/jobs")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"companyName": c.companyName,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "jobs.companies.jobs.list" call.
// Exactly one of *ListCompanyJobsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListCompanyJobsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *CompaniesJobsListCall) Do(opts ...googleapi.CallOption) (*ListCompanyJobsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListCompanyJobsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deprecated. Use ListJobs instead.\n\nLists all jobs associated with a company.",
	//   "flatPath": "v2/companies/{companiesId}/jobs",
	//   "httpMethod": "GET",
	//   "id": "jobs.companies.jobs.list",
	//   "parameterOrder": [
	//     "companyName"
	//   ],
	//   "parameters": {
	//     "companyName": {
	//       "description": "Required.\n\nThe resource name of the company that owns the jobs to be listed,\nsuch as, \"companies/0000aaaa-1111-bbbb-2222-cccc3333dddd\".",
	//       "location": "path",
	//       "pattern": "^companies/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "idsOnly": {
	//       "description": "Optional.\n\nIf set to `true`, only job ID, job requisition ID and language code will be\nreturned.\n\nA typical use is to synchronize job repositories.\n\nDefaults to false.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "includeJobsCount": {
	//       "description": "Deprecated. Please DO NOT use this field except for small companies.\nSuggest counting jobs page by page instead.\n\nOptional.\n\nSet to true if the total number of open jobs is to be returned.\n\nDefaults to false.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "jobRequisitionId": {
	//       "description": "Optional.\n\nThe requisition ID, also known as posting ID, assigned by the company\nto the job.\n\nThe maximum number of allowable characters is 225.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Optional.\n\nThe maximum number of jobs to be returned per page of results.\n\nIf ids_only is set to true, the maximum allowed page size\nis 1000. Otherwise, the maximum allowed page size is 100.\n\nDefault is 100 if empty or a number \u003c 1 is specified.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional.\n\nThe starting point of a query result.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/{+companyName}/jobs",
	//   "response": {
	//     "$ref": "ListCompanyJobsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/jobs"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *CompaniesJobsListCall) Pages(ctx context.Context, f func(*ListCompanyJobsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "jobs.jobs.batchDelete":

type JobsBatchDeleteCall struct {
	s                      *Service
	batchdeletejobsrequest *BatchDeleteJobsRequest
	urlParams_             gensupport.URLParams
	ctx_                   context.Context
	header_                http.Header
}

// BatchDelete: Deletes a list of Job postings by filter.
func (r *JobsService) BatchDelete(batchdeletejobsrequest *BatchDeleteJobsRequest) *JobsBatchDeleteCall {
	c := &JobsBatchDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.batchdeletejobsrequest = batchdeletejobsrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *JobsBatchDeleteCall) Fields(s ...googleapi.Field) *JobsBatchDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *JobsBatchDeleteCall) Context(ctx context.Context) *JobsBatchDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *JobsBatchDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *JobsBatchDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.batchdeletejobsrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/jobs:batchDelete")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "jobs.jobs.batchDelete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *JobsBatchDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes a list of Job postings by filter.",
	//   "flatPath": "v2/jobs:batchDelete",
	//   "httpMethod": "POST",
	//   "id": "jobs.jobs.batchDelete",
	//   "parameterOrder": [],
	//   "parameters": {},
	//   "path": "v2/jobs:batchDelete",
	//   "request": {
	//     "$ref": "BatchDeleteJobsRequest"
	//   },
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/jobs"
	//   ]
	// }

}

// method id "jobs.jobs.create":

type JobsCreateCall struct {
	s                *Service
	createjobrequest *CreateJobRequest
	urlParams_       gensupport.URLParams
	ctx_             context.Context
	header_          http.Header
}

// Create: Creates a new job.
//
// Typically, the job becomes searchable within 10 seconds, but it may
// take
// up to 5 minutes.
func (r *JobsService) Create(createjobrequest *CreateJobRequest) *JobsCreateCall {
	c := &JobsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.createjobrequest = createjobrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *JobsCreateCall) Fields(s ...googleapi.Field) *JobsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *JobsCreateCall) Context(ctx context.Context) *JobsCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *JobsCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *JobsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.createjobrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/jobs")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "jobs.jobs.create" call.
// Exactly one of *Job or error will be non-nil. Any non-2xx status code
// is an error. Response headers are in either
// *Job.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *JobsCreateCall) Do(opts ...googleapi.CallOption) (*Job, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Job{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a new job.\n\nTypically, the job becomes searchable within 10 seconds, but it may take\nup to 5 minutes.",
	//   "flatPath": "v2/jobs",
	//   "httpMethod": "POST",
	//   "id": "jobs.jobs.create",
	//   "parameterOrder": [],
	//   "parameters": {},
	//   "path": "v2/jobs",
	//   "request": {
	//     "$ref": "CreateJobRequest"
	//   },
	//   "response": {
	//     "$ref": "Job"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/jobs"
	//   ]
	// }

}

// method id "jobs.jobs.delete":

type JobsDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes the specified job.
//
// Typically, the job becomes unsearchable within 10 seconds, but it may
// take
// up to 5 minutes.
func (r *JobsService) Delete(name string) *JobsDeleteCall {
	c := &JobsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// DisableFastProcess sets the optional parameter "disableFastProcess":
// Deprecated. This field is not working anymore.
//
//
//
// If set to true, this call waits for all processing steps to
// complete
// before the job is cleaned up. Otherwise, the call returns while
// some
// steps are still taking place asynchronously, hence faster.
func (c *JobsDeleteCall) DisableFastProcess(disableFastProcess bool) *JobsDeleteCall {
	c.urlParams_.Set("disableFastProcess", fmt.Sprint(disableFastProcess))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *JobsDeleteCall) Fields(s ...googleapi.Field) *JobsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *JobsDeleteCall) Context(ctx context.Context) *JobsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *JobsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *JobsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "jobs.jobs.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *JobsDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes the specified job.\n\nTypically, the job becomes unsearchable within 10 seconds, but it may take\nup to 5 minutes.",
	//   "flatPath": "v2/jobs/{jobsId}",
	//   "httpMethod": "DELETE",
	//   "id": "jobs.jobs.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "disableFastProcess": {
	//       "description": "Deprecated. This field is not working anymore.\n\nOptional.\n\nIf set to true, this call waits for all processing steps to complete\nbefore the job is cleaned up. Otherwise, the call returns while some\nsteps are still taking place asynchronously, hence faster.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "name": {
	//       "description": "Required.\n\nThe resource name of the job to be deleted, such as \"jobs/11111111\".",
	//       "location": "path",
	//       "pattern": "^jobs/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/{+name}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/jobs"
	//   ]
	// }

}

// method id "jobs.jobs.deleteByFilter":

type JobsDeleteByFilterCall struct {
	s                         *Service
	deletejobsbyfilterrequest *DeleteJobsByFilterRequest
	urlParams_                gensupport.URLParams
	ctx_                      context.Context
	header_                   http.Header
}

// DeleteByFilter: Deprecated. Use BatchDeleteJobs instead.
//
// Deletes the specified job by filter. You can specify whether
// to
// synchronously wait for validation, indexing, and general processing
// to be
// completed before the response is returned.
func (r *JobsService) DeleteByFilter(deletejobsbyfilterrequest *DeleteJobsByFilterRequest) *JobsDeleteByFilterCall {
	c := &JobsDeleteByFilterCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.deletejobsbyfilterrequest = deletejobsbyfilterrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *JobsDeleteByFilterCall) Fields(s ...googleapi.Field) *JobsDeleteByFilterCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *JobsDeleteByFilterCall) Context(ctx context.Context) *JobsDeleteByFilterCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *JobsDeleteByFilterCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *JobsDeleteByFilterCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.deletejobsbyfilterrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/jobs:deleteByFilter")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "jobs.jobs.deleteByFilter" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *JobsDeleteByFilterCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deprecated. Use BatchDeleteJobs instead.\n\nDeletes the specified job by filter. You can specify whether to\nsynchronously wait for validation, indexing, and general processing to be\ncompleted before the response is returned.",
	//   "flatPath": "v2/jobs:deleteByFilter",
	//   "httpMethod": "POST",
	//   "id": "jobs.jobs.deleteByFilter",
	//   "parameterOrder": [],
	//   "parameters": {},
	//   "path": "v2/jobs:deleteByFilter",
	//   "request": {
	//     "$ref": "DeleteJobsByFilterRequest"
	//   },
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/jobs"
	//   ]
	// }

}

// method id "jobs.jobs.get":

type JobsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Retrieves the specified job, whose status is OPEN or recently
// EXPIRED
// within the last 90 days.
func (r *JobsService) Get(name string) *JobsGetCall {
	c := &JobsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *JobsGetCall) Fields(s ...googleapi.Field) *JobsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *JobsGetCall) IfNoneMatch(entityTag string) *JobsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *JobsGetCall) Context(ctx context.Context) *JobsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *JobsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *JobsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "jobs.jobs.get" call.
// Exactly one of *Job or error will be non-nil. Any non-2xx status code
// is an error. Response headers are in either
// *Job.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *JobsGetCall) Do(opts ...googleapi.CallOption) (*Job, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Job{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the specified job, whose status is OPEN or recently EXPIRED\nwithin the last 90 days.",
	//   "flatPath": "v2/jobs/{jobsId}",
	//   "httpMethod": "GET",
	//   "id": "jobs.jobs.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required.\n\nThe resource name of the job to retrieve, such as \"jobs/11111111\".",
	//       "location": "path",
	//       "pattern": "^jobs/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/{+name}",
	//   "response": {
	//     "$ref": "Job"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/jobs"
	//   ]
	// }

}

// method id "jobs.jobs.histogram":

type JobsHistogramCall struct {
	s                   *Service
	gethistogramrequest *GetHistogramRequest
	urlParams_          gensupport.URLParams
	ctx_                context.Context
	header_             http.Header
}

// Histogram: Deprecated. Use SearchJobsRequest.histogram_facets instead
// to make
// a single call with both search and histogram.
//
// Retrieves a histogram for the given
// GetHistogramRequest. This call provides a structured
// count of jobs that match against the search query, grouped by
// specified
// facets.
//
// This call constrains the visibility of jobs
// present in the database, and only counts jobs the caller
// has
// permission to search against.
//
// For example, use this call to generate the
// number of jobs in the U.S. by state.
func (r *JobsService) Histogram(gethistogramrequest *GetHistogramRequest) *JobsHistogramCall {
	c := &JobsHistogramCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.gethistogramrequest = gethistogramrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *JobsHistogramCall) Fields(s ...googleapi.Field) *JobsHistogramCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *JobsHistogramCall) Context(ctx context.Context) *JobsHistogramCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *JobsHistogramCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *JobsHistogramCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.gethistogramrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/jobs:histogram")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "jobs.jobs.histogram" call.
// Exactly one of *GetHistogramResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *GetHistogramResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *JobsHistogramCall) Do(opts ...googleapi.CallOption) (*GetHistogramResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GetHistogramResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deprecated. Use SearchJobsRequest.histogram_facets instead to make\na single call with both search and histogram.\n\nRetrieves a histogram for the given\nGetHistogramRequest. This call provides a structured\ncount of jobs that match against the search query, grouped by specified\nfacets.\n\nThis call constrains the visibility of jobs\npresent in the database, and only counts jobs the caller has\npermission to search against.\n\nFor example, use this call to generate the\nnumber of jobs in the U.S. by state.",
	//   "flatPath": "v2/jobs:histogram",
	//   "httpMethod": "POST",
	//   "id": "jobs.jobs.histogram",
	//   "parameterOrder": [],
	//   "parameters": {},
	//   "path": "v2/jobs:histogram",
	//   "request": {
	//     "$ref": "GetHistogramRequest"
	//   },
	//   "response": {
	//     "$ref": "GetHistogramResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/jobs"
	//   ]
	// }

}

// method id "jobs.jobs.list":

type JobsListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists jobs by filter.
func (r *JobsService) List() *JobsListCall {
	c := &JobsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// Filter sets the optional parameter "filter": Required.
//
// The filter string specifies the jobs to be enumerated.
//
// Supported operator: =, AND
//
// The fields eligible for filtering are:
//
// * `companyName` (Required)
// * `requisitionId` (Optional)
//
// Sample Query:
//
// * companyName = "companies/123"
// * companyName = "companies/123" AND requisitionId = "req-1"
func (c *JobsListCall) Filter(filter string) *JobsListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// IdsOnly sets the optional parameter "idsOnly": If set to `true`, only
// Job.name, Job.requisition_id and
// Job.language_code will be returned.
//
// A typical use case is to synchronize job repositories.
//
// Defaults to false.
func (c *JobsListCall) IdsOnly(idsOnly bool) *JobsListCall {
	c.urlParams_.Set("idsOnly", fmt.Sprint(idsOnly))
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of jobs to be returned per page of results.
//
// If ids_only is set to true, the maximum allowed page size
// is 1000. Otherwise, the maximum allowed page size is 100.
//
// Default is 100 if empty or a number < 1 is specified.
func (c *JobsListCall) PageSize(pageSize int64) *JobsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The starting point
// of a query result.
func (c *JobsListCall) PageToken(pageToken string) *JobsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *JobsListCall) Fields(s ...googleapi.Field) *JobsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *JobsListCall) IfNoneMatch(entityTag string) *JobsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *JobsListCall) Context(ctx context.Context) *JobsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *JobsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *JobsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/jobs")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "jobs.jobs.list" call.
// Exactly one of *ListJobsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListJobsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *JobsListCall) Do(opts ...googleapi.CallOption) (*ListJobsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListJobsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists jobs by filter.",
	//   "flatPath": "v2/jobs",
	//   "httpMethod": "GET",
	//   "id": "jobs.jobs.list",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "filter": {
	//       "description": "Required.\n\nThe filter string specifies the jobs to be enumerated.\n\nSupported operator: =, AND\n\nThe fields eligible for filtering are:\n\n* `companyName` (Required)\n* `requisitionId` (Optional)\n\nSample Query:\n\n* companyName = \"companies/123\"\n* companyName = \"companies/123\" AND requisitionId = \"req-1\"",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "idsOnly": {
	//       "description": "Optional.\n\nIf set to `true`, only Job.name, Job.requisition_id and\nJob.language_code will be returned.\n\nA typical use case is to synchronize job repositories.\n\nDefaults to false.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "pageSize": {
	//       "description": "Optional.\n\nThe maximum number of jobs to be returned per page of results.\n\nIf ids_only is set to true, the maximum allowed page size\nis 1000. Otherwise, the maximum allowed page size is 100.\n\nDefault is 100 if empty or a number \u003c 1 is specified.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional.\n\nThe starting point of a query result.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/jobs",
	//   "response": {
	//     "$ref": "ListJobsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/jobs"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *JobsListCall) Pages(ctx context.Context, f func(*ListJobsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "jobs.jobs.patch":

type JobsPatchCall struct {
	s                *Service
	name             string
	updatejobrequest *UpdateJobRequest
	urlParams_       gensupport.URLParams
	ctx_             context.Context
	header_          http.Header
}

// Patch: Updates specified job.
//
// Typically, updated contents become visible in search results within
// 10
// seconds, but it may take up to 5 minutes.
func (r *JobsService) Patch(name string, updatejobrequest *UpdateJobRequest) *JobsPatchCall {
	c := &JobsPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.updatejobrequest = updatejobrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *JobsPatchCall) Fields(s ...googleapi.Field) *JobsPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *JobsPatchCall) Context(ctx context.Context) *JobsPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *JobsPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *JobsPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.updatejobrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "jobs.jobs.patch" call.
// Exactly one of *Job or error will be non-nil. Any non-2xx status code
// is an error. Response headers are in either
// *Job.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *JobsPatchCall) Do(opts ...googleapi.CallOption) (*Job, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Job{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates specified job.\n\nTypically, updated contents become visible in search results within 10\nseconds, but it may take up to 5 minutes.",
	//   "flatPath": "v2/jobs/{jobsId}",
	//   "httpMethod": "PATCH",
	//   "id": "jobs.jobs.patch",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required during job update.\n\nResource name assigned to a job by the API, for example, \"/jobs/foo\". Use\nof this field in job queries and API calls is preferred over the use of\nrequisition_id since this value is unique.",
	//       "location": "path",
	//       "pattern": "^jobs/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/{+name}",
	//   "request": {
	//     "$ref": "UpdateJobRequest"
	//   },
	//   "response": {
	//     "$ref": "Job"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/jobs"
	//   ]
	// }

}

// method id "jobs.jobs.search":

type JobsSearchCall struct {
	s                 *Service
	searchjobsrequest *SearchJobsRequest
	urlParams_        gensupport.URLParams
	ctx_              context.Context
	header_           http.Header
}

// Search: Searches for jobs using the provided SearchJobsRequest.
//
// This call constrains the visibility of jobs
// present in the database, and only returns jobs that the caller
// has
// permission to search against.
func (r *JobsService) Search(searchjobsrequest *SearchJobsRequest) *JobsSearchCall {
	c := &JobsSearchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.searchjobsrequest = searchjobsrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *JobsSearchCall) Fields(s ...googleapi.Field) *JobsSearchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *JobsSearchCall) Context(ctx context.Context) *JobsSearchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *JobsSearchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *JobsSearchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.searchjobsrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/jobs:search")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "jobs.jobs.search" call.
// Exactly one of *SearchJobsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *SearchJobsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *JobsSearchCall) Do(opts ...googleapi.CallOption) (*SearchJobsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &SearchJobsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Searches for jobs using the provided SearchJobsRequest.\n\nThis call constrains the visibility of jobs\npresent in the database, and only returns jobs that the caller has\npermission to search against.",
	//   "flatPath": "v2/jobs:search",
	//   "httpMethod": "POST",
	//   "id": "jobs.jobs.search",
	//   "parameterOrder": [],
	//   "parameters": {},
	//   "path": "v2/jobs:search",
	//   "request": {
	//     "$ref": "SearchJobsRequest"
	//   },
	//   "response": {
	//     "$ref": "SearchJobsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/jobs"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *JobsSearchCall) Pages(ctx context.Context, f func(*SearchJobsResponse) error) error {
	c.ctx_ = ctx
	defer func(pt string) { c.searchjobsrequest.PageToken = pt }(c.searchjobsrequest.PageToken) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.searchjobsrequest.PageToken = x.NextPageToken
	}
}

// method id "jobs.jobs.searchForAlert":

type JobsSearchForAlertCall struct {
	s                 *Service
	searchjobsrequest *SearchJobsRequest
	urlParams_        gensupport.URLParams
	ctx_              context.Context
	header_           http.Header
}

// SearchForAlert: Searches for jobs using the provided
// SearchJobsRequest.
//
// This API call is intended for the use case of targeting passive
// job
// seekers (for example, job seekers who have signed up to receive
// email
// alerts about potential job opportunities), and has different
// algorithmic
// adjustments that are targeted to passive job seekers.
//
// This call constrains the visibility of jobs
// present in the database, and only returns jobs the caller
// has
// permission to search against.
func (r *JobsService) SearchForAlert(searchjobsrequest *SearchJobsRequest) *JobsSearchForAlertCall {
	c := &JobsSearchForAlertCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.searchjobsrequest = searchjobsrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *JobsSearchForAlertCall) Fields(s ...googleapi.Field) *JobsSearchForAlertCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *JobsSearchForAlertCall) Context(ctx context.Context) *JobsSearchForAlertCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *JobsSearchForAlertCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *JobsSearchForAlertCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.searchjobsrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/jobs:searchForAlert")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "jobs.jobs.searchForAlert" call.
// Exactly one of *SearchJobsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *SearchJobsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *JobsSearchForAlertCall) Do(opts ...googleapi.CallOption) (*SearchJobsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &SearchJobsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Searches for jobs using the provided SearchJobsRequest.\n\nThis API call is intended for the use case of targeting passive job\nseekers (for example, job seekers who have signed up to receive email\nalerts about potential job opportunities), and has different algorithmic\nadjustments that are targeted to passive job seekers.\n\nThis call constrains the visibility of jobs\npresent in the database, and only returns jobs the caller has\npermission to search against.",
	//   "flatPath": "v2/jobs:searchForAlert",
	//   "httpMethod": "POST",
	//   "id": "jobs.jobs.searchForAlert",
	//   "parameterOrder": [],
	//   "parameters": {},
	//   "path": "v2/jobs:searchForAlert",
	//   "request": {
	//     "$ref": "SearchJobsRequest"
	//   },
	//   "response": {
	//     "$ref": "SearchJobsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/jobs"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *JobsSearchForAlertCall) Pages(ctx context.Context, f func(*SearchJobsResponse) error) error {
	c.ctx_ = ctx
	defer func(pt string) { c.searchjobsrequest.PageToken = pt }(c.searchjobsrequest.PageToken) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.searchjobsrequest.PageToken = x.NextPageToken
	}
}

// method id "jobs.complete":

type V2CompleteCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Complete: Completes the specified prefix with job keyword
// suggestions.
// Intended for use by a job search auto-complete search box.
func (r *V2Service) Complete() *V2CompleteCall {
	c := &V2CompleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// CompanyName sets the optional parameter "companyName": If provided,
// restricts completion to the specified company.
func (c *V2CompleteCall) CompanyName(companyName string) *V2CompleteCall {
	c.urlParams_.Set("companyName", companyName)
	return c
}

// LanguageCode sets the optional parameter "languageCode":
// Required.
//
// The language of the query. This is
// the BCP-47 language code, such as "en-US" or "sr-Latn".
// For more information, see
// [Tags for Identifying
// Languages](https://tools.ietf.org/html/bcp47).
//
// For CompletionType.JOB_TITLE type, only open jobs with
// same
// language_code are returned.
//
// For CompletionType.COMPANY_NAME type,
// only companies having open jobs with same language_code
// are
// returned.
//
// For CompletionType.COMBINED type, only open jobs with
// same
// language_code or companies having open jobs with same
// language_code are returned.
func (c *V2CompleteCall) LanguageCode(languageCode string) *V2CompleteCall {
	c.urlParams_.Set("languageCode", languageCode)
	return c
}

// PageSize sets the optional parameter "pageSize":
// Required.
//
// Completion result count.
// The maximum allowed page size is 10.
func (c *V2CompleteCall) PageSize(pageSize int64) *V2CompleteCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// Query sets the optional parameter "query": Required.
//
// The query used to generate suggestions.
func (c *V2CompleteCall) Query(query string) *V2CompleteCall {
	c.urlParams_.Set("query", query)
	return c
}

// Scope sets the optional parameter "scope": The scope of the
// completion. The defaults is CompletionScope.PUBLIC.
//
// Possible values:
//   "COMPLETION_SCOPE_UNSPECIFIED"
//   "TENANT"
//   "PUBLIC"
func (c *V2CompleteCall) Scope(scope string) *V2CompleteCall {
	c.urlParams_.Set("scope", scope)
	return c
}

// Type sets the optional parameter "type": The completion topic. The
// default is CompletionType.COMBINED.
//
// Possible values:
//   "COMPLETION_TYPE_UNSPECIFIED"
//   "JOB_TITLE"
//   "COMPANY_NAME"
//   "COMBINED"
func (c *V2CompleteCall) Type(type_ string) *V2CompleteCall {
	c.urlParams_.Set("type", type_)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *V2CompleteCall) Fields(s ...googleapi.Field) *V2CompleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *V2CompleteCall) IfNoneMatch(entityTag string) *V2CompleteCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *V2CompleteCall) Context(ctx context.Context) *V2CompleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *V2CompleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *V2CompleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2:complete")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "jobs.complete" call.
// Exactly one of *CompleteQueryResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *CompleteQueryResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *V2CompleteCall) Do(opts ...googleapi.CallOption) (*CompleteQueryResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &CompleteQueryResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Completes the specified prefix with job keyword suggestions.\nIntended for use by a job search auto-complete search box.",
	//   "flatPath": "v2:complete",
	//   "httpMethod": "GET",
	//   "id": "jobs.complete",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "companyName": {
	//       "description": "Optional.\n\nIf provided, restricts completion to the specified company.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "languageCode": {
	//       "description": "Required.\n\nThe language of the query. This is\nthe BCP-47 language code, such as \"en-US\" or \"sr-Latn\".\nFor more information, see\n[Tags for Identifying Languages](https://tools.ietf.org/html/bcp47).\n\nFor CompletionType.JOB_TITLE type, only open jobs with same\nlanguage_code are returned.\n\nFor CompletionType.COMPANY_NAME type,\nonly companies having open jobs with same language_code are\nreturned.\n\nFor CompletionType.COMBINED type, only open jobs with same\nlanguage_code or companies having open jobs with same\nlanguage_code are returned.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Required.\n\nCompletion result count.\nThe maximum allowed page size is 10.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "query": {
	//       "description": "Required.\n\nThe query used to generate suggestions.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "scope": {
	//       "description": "Optional.\n\nThe scope of the completion. The defaults is CompletionScope.PUBLIC.",
	//       "enum": [
	//         "COMPLETION_SCOPE_UNSPECIFIED",
	//         "TENANT",
	//         "PUBLIC"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "type": {
	//       "description": "Optional.\n\nThe completion topic. The default is CompletionType.COMBINED.",
	//       "enum": [
	//         "COMPLETION_TYPE_UNSPECIFIED",
	//         "JOB_TITLE",
	//         "COMPANY_NAME",
	//         "COMBINED"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2:complete",
	//   "response": {
	//     "$ref": "CompleteQueryResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/jobs"
	//   ]
	// }

}
