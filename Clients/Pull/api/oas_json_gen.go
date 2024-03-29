// Code generated by ogen, DO NOT EDIT.

package api

import (
	"math/bits"
	"strconv"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/validate"
)

// Encode implements json.Marshaler.
func (s *ClientReport) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ClientReport) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("clientid")
		e.Str(s.Clientid)
	}
	{
		if s.Reports != nil {
			e.FieldStart("reports")
			s.Reports.Encode(e)
		}
	}
}

var jsonFieldsNameOfClientReport = [2]string{
	0: "clientid",
	1: "reports",
}

// Decode decodes ClientReport from json.
func (s *ClientReport) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ClientReport to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "clientid":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.Clientid = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"clientid\"")
			}
		case "reports":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				if err := s.Reports.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"reports\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ClientReport")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfClientReport) {
					name = jsonFieldsNameOfClientReport[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ClientReport) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ClientReport) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Error) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Error) encodeFields(e *jx.Encoder) {
	{
		if s.Summary.Set {
			e.FieldStart("summary")
			s.Summary.Encode(e)
		}
	}
}

var jsonFieldsNameOfError = [1]string{
	0: "summary",
}

// Decode decodes Error from json.
func (s *Error) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Error to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "summary":
			if err := func() error {
				s.Summary.Reset()
				if err := s.Summary.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"summary\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Error")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Error) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Error) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *LightColor) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *LightColor) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("reviewRGB")
		e.Str(s.ReviewRGB)
	}
	{
		e.FieldStart("mergeRGB")
		e.Str(s.MergeRGB)
	}
	{
		e.FieldStart("pullRGB")
		e.Str(s.PullRGB)
	}
}

var jsonFieldsNameOfLightColor = [3]string{
	0: "reviewRGB",
	1: "mergeRGB",
	2: "pullRGB",
}

// Decode decodes LightColor from json.
func (s *LightColor) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode LightColor to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "reviewRGB":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.ReviewRGB = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"reviewRGB\"")
			}
		case "mergeRGB":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.MergeRGB = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"mergeRGB\"")
			}
		case "pullRGB":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Str()
				s.PullRGB = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"pullRGB\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode LightColor")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfLightColor) {
					name = jsonFieldsNameOfLightColor[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *LightColor) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *LightColor) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes string as json.
func (o OptString) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes string from json.
func (o *OptString) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptString to nil")
	}
	o.Set = true
	v, err := d.Str()
	if err != nil {
		return err
	}
	o.Value = string(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptString) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptString) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *ReportTuple) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ReportTuple) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("repository")
		e.Str(s.Repository)
	}
	{
		e.FieldStart("section")
		s.Section.Encode(e)
	}
	{
		e.FieldStart("age")
		e.Int(s.Age)
	}
	{
		e.FieldStart("url")
		e.Str(s.URL)
	}
	{
		e.FieldStart("notes")
		e.Str(s.Notes)
	}
}

var jsonFieldsNameOfReportTuple = [5]string{
	0: "repository",
	1: "section",
	2: "age",
	3: "url",
	4: "notes",
}

// Decode decodes ReportTuple from json.
func (s *ReportTuple) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ReportTuple to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "repository":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.Repository = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"repository\"")
			}
		case "section":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				if err := s.Section.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"section\"")
			}
		case "age":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Int()
				s.Age = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"age\"")
			}
		case "url":
			requiredBitSet[0] |= 1 << 3
			if err := func() error {
				v, err := d.Str()
				s.URL = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"url\"")
			}
		case "notes":
			requiredBitSet[0] |= 1 << 4
			if err := func() error {
				v, err := d.Str()
				s.Notes = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"notes\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ReportTuple")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00011111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfReportTuple) {
					name = jsonFieldsNameOfReportTuple[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ReportTuple) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ReportTuple) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes ReportTupleSection as json.
func (s ReportTupleSection) Encode(e *jx.Encoder) {
	e.Str(string(s))
}

// Decode decodes ReportTupleSection from json.
func (s *ReportTupleSection) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ReportTupleSection to nil")
	}
	v, err := d.StrBytes()
	if err != nil {
		return err
	}
	// Try to use constant string.
	switch ReportTupleSection(v) {
	case ReportTupleSectionReview:
		*s = ReportTupleSectionReview
	case ReportTupleSectionMerge:
		*s = ReportTupleSectionMerge
	case ReportTupleSectionPull:
		*s = ReportTupleSectionPull
	default:
		*s = ReportTupleSection(v)
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s ReportTupleSection) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ReportTupleSection) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes Reports as json.
func (s Reports) Encode(e *jx.Encoder) {
	unwrapped := []ReportsItem(s)
	if unwrapped == nil {
		e.ArrEmpty()
		return
	}
	if unwrapped != nil {
		e.ArrStart()
		for _, elem := range unwrapped {
			elem.Encode(e)
		}
		e.ArrEnd()
	}
}

// Decode decodes Reports from json.
func (s *Reports) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Reports to nil")
	}
	var unwrapped []ReportsItem
	if err := func() error {
		unwrapped = make([]ReportsItem, 0)
		if err := d.Arr(func(d *jx.Decoder) error {
			var elem ReportsItem
			if err := elem.Decode(d); err != nil {
				return err
			}
			unwrapped = append(unwrapped, elem)
			return nil
		}); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = Reports(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s Reports) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Reports) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes ReportsItem as json.
func (s ReportsItem) Encode(e *jx.Encoder) {
	switch s.Type {
	case ReportTupleReportsItem:
		s.ReportTuple.Encode(e)
	}
}

// Decode decodes ReportsItem from json.
func (s *ReportsItem) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ReportsItem to nil")
	}
	// Sum type type_discriminator.
	switch t := d.Next(); t {
	case jx.Object:
		if err := s.ReportTuple.Decode(d); err != nil {
			return err
		}
		s.Type = ReportTupleReportsItem
	default:
		return errors.Errorf("unexpected json type %q", t)
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s ReportsItem) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ReportsItem) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Result) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Result) encodeFields(e *jx.Encoder) {
	{
		if s.Summary.Set {
			e.FieldStart("summary")
			s.Summary.Encode(e)
		}
	}
}

var jsonFieldsNameOfResult = [1]string{
	0: "summary",
}

// Decode decodes Result from json.
func (s *Result) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Result to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "summary":
			if err := func() error {
				s.Summary.Reset()
				if err := s.Summary.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"summary\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Result")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Result) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Result) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Status) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Status) encodeFields(e *jx.Encoder) {
	{
		if s.Reports != nil {
			e.FieldStart("reports")
			s.Reports.Encode(e)
		}
	}
}

var jsonFieldsNameOfStatus = [1]string{
	0: "reports",
}

// Decode decodes Status from json.
func (s *Status) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Status to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "reports":
			if err := func() error {
				if err := s.Reports.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"reports\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Status")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Status) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Status) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}
