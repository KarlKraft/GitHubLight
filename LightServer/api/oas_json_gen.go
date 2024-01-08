// Code generated by ogen, DO NOT EDIT.

package api

import (
	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
)

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

// Encode encodes int as json.
func (o OptInt) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Int(int(o.Value))
}

// Decode decodes int from json.
func (o *OptInt) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptInt to nil")
	}
	o.Set = true
	v, err := d.Int()
	if err != nil {
		return err
	}
	o.Value = int(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptInt) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptInt) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes ReportTupleSection as json.
func (o OptReportTupleSection) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes ReportTupleSection from json.
func (o *OptReportTupleSection) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptReportTupleSection to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptReportTupleSection) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptReportTupleSection) UnmarshalJSON(data []byte) error {
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
		if s.Owner.Set {
			e.FieldStart("owner")
			s.Owner.Encode(e)
		}
	}
	{
		if s.Repository.Set {
			e.FieldStart("repository")
			s.Repository.Encode(e)
		}
	}
	{
		if s.Section.Set {
			e.FieldStart("section")
			s.Section.Encode(e)
		}
	}
	{
		if s.Age.Set {
			e.FieldStart("age")
			s.Age.Encode(e)
		}
	}
}

var jsonFieldsNameOfReportTuple = [4]string{
	0: "owner",
	1: "repository",
	2: "section",
	3: "age",
}

// Decode decodes ReportTuple from json.
func (s *ReportTuple) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ReportTuple to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "owner":
			if err := func() error {
				s.Owner.Reset()
				if err := s.Owner.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"owner\"")
			}
		case "repository":
			if err := func() error {
				s.Repository.Reset()
				if err := s.Repository.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"repository\"")
			}
		case "section":
			if err := func() error {
				s.Section.Reset()
				if err := s.Section.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"section\"")
			}
		case "age":
			if err := func() error {
				s.Age.Reset()
				if err := s.Age.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"age\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ReportTuple")
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

	e.ArrStart()
	for _, elem := range unwrapped {
		elem.Encode(e)
	}
	e.ArrEnd()
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
