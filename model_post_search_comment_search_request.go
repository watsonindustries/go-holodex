/*
Holodex/HoloAPI V2

Holodex Public API. Successor to the HoloAPI v1

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package holodex

import (
	"encoding/json"
)

// PostSearchCommentSearchRequest 
type PostSearchCommentSearchRequest struct {
	Sort string `json:"sort"`
	// If set, will filter clips to only show clips with these langs + all vtuber streams (provided `target` is not set to filter out streams)
	Lang []string `json:"lang,omitempty"`
	// Target types of videos
	Target []string `json:"target,omitempty"`
	// Find videos with comments containing specified string (case insensitive)
	Comment string `json:"comment"`
	// Return videos that match one of the provided topics
	Topic []string `json:"topic,omitempty"`
	// Videos with all of the specified channel  ids. If two or more channel IDs are specified, will only return their collabs, or if one channel is a clipper, it will only show clips of the other vtubers made by this clipper.
	Vch []string `json:"vch,omitempty"`
	// Videos of channels in any of the specified orgs, or clips that involve a channel in the specified org.
	Org []string `json:"org,omitempty"`
	// If set at all, responds with total and items wrapping the array of objects
	Paginated *bool `json:"paginated,omitempty"`
	Offset float32 `json:"offset"`
	Limit float32 `json:"limit"`
}

// NewPostSearchCommentSearchRequest instantiates a new PostSearchCommentSearchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostSearchCommentSearchRequest(sort string, comment string, offset float32, limit float32) *PostSearchCommentSearchRequest {
	this := PostSearchCommentSearchRequest{}
	this.Sort = sort
	this.Comment = comment
	this.Offset = offset
	this.Limit = limit
	return &this
}

// NewPostSearchCommentSearchRequestWithDefaults instantiates a new PostSearchCommentSearchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostSearchCommentSearchRequestWithDefaults() *PostSearchCommentSearchRequest {
	this := PostSearchCommentSearchRequest{}
	var sort string = "newest"
	this.Sort = sort
	return &this
}

// GetSort returns the Sort field value
func (o *PostSearchCommentSearchRequest) GetSort() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sort
}

// GetSortOk returns a tuple with the Sort field value
// and a boolean to check if the value has been set.
func (o *PostSearchCommentSearchRequest) GetSortOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sort, true
}

// SetSort sets field value
func (o *PostSearchCommentSearchRequest) SetSort(v string) {
	o.Sort = v
}

// GetLang returns the Lang field value if set, zero value otherwise.
func (o *PostSearchCommentSearchRequest) GetLang() []string {
	if o == nil || o.Lang == nil {
		var ret []string
		return ret
	}
	return o.Lang
}

// GetLangOk returns a tuple with the Lang field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostSearchCommentSearchRequest) GetLangOk() ([]string, bool) {
	if o == nil || o.Lang == nil {
		return nil, false
	}
	return o.Lang, true
}

// HasLang returns a boolean if a field has been set.
func (o *PostSearchCommentSearchRequest) HasLang() bool {
	if o != nil && o.Lang != nil {
		return true
	}

	return false
}

// SetLang gets a reference to the given []string and assigns it to the Lang field.
func (o *PostSearchCommentSearchRequest) SetLang(v []string) {
	o.Lang = v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *PostSearchCommentSearchRequest) GetTarget() []string {
	if o == nil || o.Target == nil {
		var ret []string
		return ret
	}
	return o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostSearchCommentSearchRequest) GetTargetOk() ([]string, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *PostSearchCommentSearchRequest) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given []string and assigns it to the Target field.
func (o *PostSearchCommentSearchRequest) SetTarget(v []string) {
	o.Target = v
}

// GetComment returns the Comment field value
func (o *PostSearchCommentSearchRequest) GetComment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Comment
}

// GetCommentOk returns a tuple with the Comment field value
// and a boolean to check if the value has been set.
func (o *PostSearchCommentSearchRequest) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Comment, true
}

// SetComment sets field value
func (o *PostSearchCommentSearchRequest) SetComment(v string) {
	o.Comment = v
}

// GetTopic returns the Topic field value if set, zero value otherwise.
func (o *PostSearchCommentSearchRequest) GetTopic() []string {
	if o == nil || o.Topic == nil {
		var ret []string
		return ret
	}
	return o.Topic
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostSearchCommentSearchRequest) GetTopicOk() ([]string, bool) {
	if o == nil || o.Topic == nil {
		return nil, false
	}
	return o.Topic, true
}

// HasTopic returns a boolean if a field has been set.
func (o *PostSearchCommentSearchRequest) HasTopic() bool {
	if o != nil && o.Topic != nil {
		return true
	}

	return false
}

// SetTopic gets a reference to the given []string and assigns it to the Topic field.
func (o *PostSearchCommentSearchRequest) SetTopic(v []string) {
	o.Topic = v
}

// GetVch returns the Vch field value if set, zero value otherwise.
func (o *PostSearchCommentSearchRequest) GetVch() []string {
	if o == nil || o.Vch == nil {
		var ret []string
		return ret
	}
	return o.Vch
}

// GetVchOk returns a tuple with the Vch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostSearchCommentSearchRequest) GetVchOk() ([]string, bool) {
	if o == nil || o.Vch == nil {
		return nil, false
	}
	return o.Vch, true
}

// HasVch returns a boolean if a field has been set.
func (o *PostSearchCommentSearchRequest) HasVch() bool {
	if o != nil && o.Vch != nil {
		return true
	}

	return false
}

// SetVch gets a reference to the given []string and assigns it to the Vch field.
func (o *PostSearchCommentSearchRequest) SetVch(v []string) {
	o.Vch = v
}

// GetOrg returns the Org field value if set, zero value otherwise.
func (o *PostSearchCommentSearchRequest) GetOrg() []string {
	if o == nil || o.Org == nil {
		var ret []string
		return ret
	}
	return o.Org
}

// GetOrgOk returns a tuple with the Org field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostSearchCommentSearchRequest) GetOrgOk() ([]string, bool) {
	if o == nil || o.Org == nil {
		return nil, false
	}
	return o.Org, true
}

// HasOrg returns a boolean if a field has been set.
func (o *PostSearchCommentSearchRequest) HasOrg() bool {
	if o != nil && o.Org != nil {
		return true
	}

	return false
}

// SetOrg gets a reference to the given []string and assigns it to the Org field.
func (o *PostSearchCommentSearchRequest) SetOrg(v []string) {
	o.Org = v
}

// GetPaginated returns the Paginated field value if set, zero value otherwise.
func (o *PostSearchCommentSearchRequest) GetPaginated() bool {
	if o == nil || o.Paginated == nil {
		var ret bool
		return ret
	}
	return *o.Paginated
}

// GetPaginatedOk returns a tuple with the Paginated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostSearchCommentSearchRequest) GetPaginatedOk() (*bool, bool) {
	if o == nil || o.Paginated == nil {
		return nil, false
	}
	return o.Paginated, true
}

// HasPaginated returns a boolean if a field has been set.
func (o *PostSearchCommentSearchRequest) HasPaginated() bool {
	if o != nil && o.Paginated != nil {
		return true
	}

	return false
}

// SetPaginated gets a reference to the given bool and assigns it to the Paginated field.
func (o *PostSearchCommentSearchRequest) SetPaginated(v bool) {
	o.Paginated = &v
}

// GetOffset returns the Offset field value
func (o *PostSearchCommentSearchRequest) GetOffset() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *PostSearchCommentSearchRequest) GetOffsetOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *PostSearchCommentSearchRequest) SetOffset(v float32) {
	o.Offset = v
}

// GetLimit returns the Limit field value
func (o *PostSearchCommentSearchRequest) GetLimit() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *PostSearchCommentSearchRequest) GetLimitOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *PostSearchCommentSearchRequest) SetLimit(v float32) {
	o.Limit = v
}

func (o PostSearchCommentSearchRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sort"] = o.Sort
	}
	if o.Lang != nil {
		toSerialize["lang"] = o.Lang
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	if true {
		toSerialize["comment"] = o.Comment
	}
	if o.Topic != nil {
		toSerialize["topic"] = o.Topic
	}
	if o.Vch != nil {
		toSerialize["vch"] = o.Vch
	}
	if o.Org != nil {
		toSerialize["org"] = o.Org
	}
	if o.Paginated != nil {
		toSerialize["paginated"] = o.Paginated
	}
	if true {
		toSerialize["offset"] = o.Offset
	}
	if true {
		toSerialize["limit"] = o.Limit
	}
	return json.Marshal(toSerialize)
}

type NullablePostSearchCommentSearchRequest struct {
	value *PostSearchCommentSearchRequest
	isSet bool
}

func (v NullablePostSearchCommentSearchRequest) Get() *PostSearchCommentSearchRequest {
	return v.value
}

func (v *NullablePostSearchCommentSearchRequest) Set(val *PostSearchCommentSearchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostSearchCommentSearchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostSearchCommentSearchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostSearchCommentSearchRequest(val *PostSearchCommentSearchRequest) *NullablePostSearchCommentSearchRequest {
	return &NullablePostSearchCommentSearchRequest{value: val, isSet: true}
}

func (v NullablePostSearchCommentSearchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostSearchCommentSearchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


