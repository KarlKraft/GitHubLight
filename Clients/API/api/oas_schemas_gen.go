// Code generated by ogen, DO NOT EDIT.

package api

import (
	"fmt"

	"github.com/go-faster/errors"
)

func (s *ErrorStatusCode) Error() string {
	return fmt.Sprintf("code %d: %+v", s.StatusCode, s.Response)
}

// Ref: #/components/schemas/ClientReport
type ClientReport struct {
	Clientid string  `json:"clientid"`
	Reports  Reports `json:"reports"`
}

// GetClientid returns the value of Clientid.
func (s *ClientReport) GetClientid() string {
	return s.Clientid
}

// GetReports returns the value of Reports.
func (s *ClientReport) GetReports() Reports {
	return s.Reports
}

// SetClientid sets the value of Clientid.
func (s *ClientReport) SetClientid(val string) {
	s.Clientid = val
}

// SetReports sets the value of Reports.
func (s *ClientReport) SetReports(val Reports) {
	s.Reports = val
}

// Ref: #/components/schemas/Error
type Error struct {
	Summary OptString `json:"summary"`
}

// GetSummary returns the value of Summary.
func (s *Error) GetSummary() OptString {
	return s.Summary
}

// SetSummary sets the value of Summary.
func (s *Error) SetSummary(val OptString) {
	s.Summary = val
}

// ErrorStatusCode wraps Error with StatusCode.
type ErrorStatusCode struct {
	StatusCode int
	Response   Error
}

// GetStatusCode returns the value of StatusCode.
func (s *ErrorStatusCode) GetStatusCode() int {
	return s.StatusCode
}

// GetResponse returns the value of Response.
func (s *ErrorStatusCode) GetResponse() Error {
	return s.Response
}

// SetStatusCode sets the value of StatusCode.
func (s *ErrorStatusCode) SetStatusCode(val int) {
	s.StatusCode = val
}

// SetResponse sets the value of Response.
func (s *ErrorStatusCode) SetResponse(val Error) {
	s.Response = val
}

// Ref: #/components/schemas/LightColor
type LightColor struct {
	ReviewRGB string `json:"reviewRGB"`
	MergeRGB  string `json:"mergeRGB"`
	PullRGB   string `json:"pullRGB"`
}

// GetReviewRGB returns the value of ReviewRGB.
func (s *LightColor) GetReviewRGB() string {
	return s.ReviewRGB
}

// GetMergeRGB returns the value of MergeRGB.
func (s *LightColor) GetMergeRGB() string {
	return s.MergeRGB
}

// GetPullRGB returns the value of PullRGB.
func (s *LightColor) GetPullRGB() string {
	return s.PullRGB
}

// SetReviewRGB sets the value of ReviewRGB.
func (s *LightColor) SetReviewRGB(val string) {
	s.ReviewRGB = val
}

// SetMergeRGB sets the value of MergeRGB.
func (s *LightColor) SetMergeRGB(val string) {
	s.MergeRGB = val
}

// SetPullRGB sets the value of PullRGB.
func (s *LightColor) SetPullRGB(val string) {
	s.PullRGB = val
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Ref: #/components/schemas/ReportTuple
type ReportTuple struct {
	Repository string             `json:"repository"`
	Section    ReportTupleSection `json:"section"`
	Age        int                `json:"age"`
	URL        string             `json:"url"`
	Notes      string             `json:"notes"`
}

// GetRepository returns the value of Repository.
func (s *ReportTuple) GetRepository() string {
	return s.Repository
}

// GetSection returns the value of Section.
func (s *ReportTuple) GetSection() ReportTupleSection {
	return s.Section
}

// GetAge returns the value of Age.
func (s *ReportTuple) GetAge() int {
	return s.Age
}

// GetURL returns the value of URL.
func (s *ReportTuple) GetURL() string {
	return s.URL
}

// GetNotes returns the value of Notes.
func (s *ReportTuple) GetNotes() string {
	return s.Notes
}

// SetRepository sets the value of Repository.
func (s *ReportTuple) SetRepository(val string) {
	s.Repository = val
}

// SetSection sets the value of Section.
func (s *ReportTuple) SetSection(val ReportTupleSection) {
	s.Section = val
}

// SetAge sets the value of Age.
func (s *ReportTuple) SetAge(val int) {
	s.Age = val
}

// SetURL sets the value of URL.
func (s *ReportTuple) SetURL(val string) {
	s.URL = val
}

// SetNotes sets the value of Notes.
func (s *ReportTuple) SetNotes(val string) {
	s.Notes = val
}

type ReportTupleSection string

const (
	ReportTupleSectionReview ReportTupleSection = "review"
	ReportTupleSectionMerge  ReportTupleSection = "merge"
	ReportTupleSectionPull   ReportTupleSection = "pull"
)

// AllValues returns all ReportTupleSection values.
func (ReportTupleSection) AllValues() []ReportTupleSection {
	return []ReportTupleSection{
		ReportTupleSectionReview,
		ReportTupleSectionMerge,
		ReportTupleSectionPull,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s ReportTupleSection) MarshalText() ([]byte, error) {
	switch s {
	case ReportTupleSectionReview:
		return []byte(s), nil
	case ReportTupleSectionMerge:
		return []byte(s), nil
	case ReportTupleSectionPull:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *ReportTupleSection) UnmarshalText(data []byte) error {
	switch ReportTupleSection(data) {
	case ReportTupleSectionReview:
		*s = ReportTupleSectionReview
		return nil
	case ReportTupleSectionMerge:
		*s = ReportTupleSectionMerge
		return nil
	case ReportTupleSectionPull:
		*s = ReportTupleSectionPull
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

type Reports []ReportsItem

// ReportsItem represents sum type.
type ReportsItem struct {
	Type        ReportsItemType // switch on this field
	ReportTuple ReportTuple
}

// ReportsItemType is oneOf type of ReportsItem.
type ReportsItemType string

// Possible values for ReportsItemType.
const (
	ReportTupleReportsItem ReportsItemType = "ReportTuple"
)

// IsReportTuple reports whether ReportsItem is ReportTuple.
func (s ReportsItem) IsReportTuple() bool { return s.Type == ReportTupleReportsItem }

// SetReportTuple sets ReportsItem to ReportTuple.
func (s *ReportsItem) SetReportTuple(v ReportTuple) {
	s.Type = ReportTupleReportsItem
	s.ReportTuple = v
}

// GetReportTuple returns ReportTuple and true boolean if ReportsItem is ReportTuple.
func (s ReportsItem) GetReportTuple() (v ReportTuple, ok bool) {
	if !s.IsReportTuple() {
		return v, false
	}
	return s.ReportTuple, true
}

// NewReportTupleReportsItem returns new ReportsItem from ReportTuple.
func NewReportTupleReportsItem(v ReportTuple) ReportsItem {
	var s ReportsItem
	s.SetReportTuple(v)
	return s
}

// Ref: #/components/schemas/Result
type Result struct {
	Summary OptString `json:"summary"`
}

// GetSummary returns the value of Summary.
func (s *Result) GetSummary() OptString {
	return s.Summary
}

// SetSummary sets the value of Summary.
func (s *Result) SetSummary(val OptString) {
	s.Summary = val
}

// Ref: #/components/schemas/Status
type Status struct {
	Reports Reports `json:"reports"`
}

// GetReports returns the value of Reports.
func (s *Status) GetReports() Reports {
	return s.Reports
}

// SetReports sets the value of Reports.
func (s *Status) SetReports(val Reports) {
	s.Reports = val
}
