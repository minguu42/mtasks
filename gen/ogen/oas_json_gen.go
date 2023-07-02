// Code generated by ogen, DO NOT EDIT.

package ogen

import (
	"math/bits"
	"strconv"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/validate"
)

// Encode implements json.Marshaler.
func (s *CreateProjectReq) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *CreateProjectReq) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("name")
		e.Str(s.Name)
	}
}

var jsonFieldsNameOfCreateProjectReq = [1]string{
	0: "name",
}

// Decode decodes CreateProjectReq from json.
func (s *CreateProjectReq) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode CreateProjectReq to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "name":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.Name = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"name\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode CreateProjectReq")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
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
				if fieldIdx < len(jsonFieldsNameOfCreateProjectReq) {
					name = jsonFieldsNameOfCreateProjectReq[fieldIdx]
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
func (s *CreateProjectReq) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *CreateProjectReq) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *CreateTaskReq) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *CreateTaskReq) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("title")
		e.Str(s.Title)
	}
}

var jsonFieldsNameOfCreateTaskReq = [1]string{
	0: "title",
}

// Decode decodes CreateTaskReq from json.
func (s *CreateTaskReq) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode CreateTaskReq to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "title":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.Title = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"title\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode CreateTaskReq")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
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
				if fieldIdx < len(jsonFieldsNameOfCreateTaskReq) {
					name = jsonFieldsNameOfCreateTaskReq[fieldIdx]
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
func (s *CreateTaskReq) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *CreateTaskReq) UnmarshalJSON(data []byte) error {
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
		e.FieldStart("message")
		e.Str(s.Message)
	}
	{
		e.FieldStart("debug")
		e.Str(s.Debug)
	}
}

var jsonFieldsNameOfError = [2]string{
	0: "message",
	1: "debug",
}

// Decode decodes Error from json.
func (s *Error) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Error to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "message":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.Message = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"message\"")
			}
		case "debug":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Debug = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"debug\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Error")
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
				if fieldIdx < len(jsonFieldsNameOfError) {
					name = jsonFieldsNameOfError[fieldIdx]
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
func (s *GetHealthOK) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *GetHealthOK) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("version")
		e.Str(s.Version)
	}
	{
		e.FieldStart("revision")
		e.Str(s.Revision)
	}
}

var jsonFieldsNameOfGetHealthOK = [2]string{
	0: "version",
	1: "revision",
}

// Decode decodes GetHealthOK from json.
func (s *GetHealthOK) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode GetHealthOK to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "version":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.Version = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"version\"")
			}
		case "revision":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Revision = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"revision\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode GetHealthOK")
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
				if fieldIdx < len(jsonFieldsNameOfGetHealthOK) {
					name = jsonFieldsNameOfGetHealthOK[fieldIdx]
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
func (s *GetHealthOK) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *GetHealthOK) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes bool as json.
func (o OptBool) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Bool(bool(o.Value))
}

// Decode decodes bool from json.
func (o *OptBool) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptBool to nil")
	}
	o.Set = true
	v, err := d.Bool()
	if err != nil {
		return err
	}
	o.Value = bool(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptBool) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptBool) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes time.Time as json.
func (o OptDateTime) Encode(e *jx.Encoder, format func(*jx.Encoder, time.Time)) {
	if !o.Set {
		return
	}
	format(e, o.Value)
}

// Decode decodes time.Time from json.
func (o *OptDateTime) Decode(d *jx.Decoder, format func(*jx.Decoder) (time.Time, error)) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptDateTime to nil")
	}
	o.Set = true
	v, err := format(d)
	if err != nil {
		return err
	}
	o.Value = v
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptDateTime) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e, json.EncodeDateTime)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptDateTime) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d, json.DecodeDateTime)
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
func (s *Project) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Project) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("id")
		e.Int64(s.ID)
	}
	{
		e.FieldStart("name")
		e.Str(s.Name)
	}
	{
		e.FieldStart("createdAt")
		json.EncodeDateTime(e, s.CreatedAt)
	}
	{
		e.FieldStart("updatedAt")
		json.EncodeDateTime(e, s.UpdatedAt)
	}
}

var jsonFieldsNameOfProject = [4]string{
	0: "id",
	1: "name",
	2: "createdAt",
	3: "updatedAt",
}

// Decode decodes Project from json.
func (s *Project) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Project to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Int64()
				s.ID = int64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "name":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Name = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"name\"")
			}
		case "createdAt":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := json.DecodeDateTime(d)
				s.CreatedAt = v
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"createdAt\"")
			}
		case "updatedAt":
			requiredBitSet[0] |= 1 << 3
			if err := func() error {
				v, err := json.DecodeDateTime(d)
				s.UpdatedAt = v
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"updatedAt\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Project")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00001111,
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
				if fieldIdx < len(jsonFieldsNameOfProject) {
					name = jsonFieldsNameOfProject[fieldIdx]
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
func (s *Project) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Project) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Projects) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Projects) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("projects")
		e.ArrStart()
		for _, elem := range s.Projects {
			elem.Encode(e)
		}
		e.ArrEnd()
	}
	{
		e.FieldStart("hasNext")
		e.Bool(s.HasNext)
	}
}

var jsonFieldsNameOfProjects = [2]string{
	0: "projects",
	1: "hasNext",
}

// Decode decodes Projects from json.
func (s *Projects) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Projects to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "projects":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				s.Projects = make([]Project, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem Project
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.Projects = append(s.Projects, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"projects\"")
			}
		case "hasNext":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Bool()
				s.HasNext = bool(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"hasNext\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Projects")
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
				if fieldIdx < len(jsonFieldsNameOfProjects) {
					name = jsonFieldsNameOfProjects[fieldIdx]
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
func (s *Projects) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Projects) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Task) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Task) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("id")
		e.Int64(s.ID)
	}
	{
		e.FieldStart("project_id")
		e.Int64(s.ProjectID)
	}
	{
		e.FieldStart("title")
		e.Str(s.Title)
	}
	{
		if s.CompletedAt.Set {
			e.FieldStart("completedAt")
			s.CompletedAt.Encode(e, json.EncodeDateTime)
		}
	}
	{
		e.FieldStart("createdAt")
		json.EncodeDateTime(e, s.CreatedAt)
	}
	{
		e.FieldStart("updatedAt")
		json.EncodeDateTime(e, s.UpdatedAt)
	}
}

var jsonFieldsNameOfTask = [6]string{
	0: "id",
	1: "project_id",
	2: "title",
	3: "completedAt",
	4: "createdAt",
	5: "updatedAt",
}

// Decode decodes Task from json.
func (s *Task) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Task to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Int64()
				s.ID = int64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "project_id":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Int64()
				s.ProjectID = int64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"project_id\"")
			}
		case "title":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Str()
				s.Title = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"title\"")
			}
		case "completedAt":
			if err := func() error {
				s.CompletedAt.Reset()
				if err := s.CompletedAt.Decode(d, json.DecodeDateTime); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"completedAt\"")
			}
		case "createdAt":
			requiredBitSet[0] |= 1 << 4
			if err := func() error {
				v, err := json.DecodeDateTime(d)
				s.CreatedAt = v
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"createdAt\"")
			}
		case "updatedAt":
			requiredBitSet[0] |= 1 << 5
			if err := func() error {
				v, err := json.DecodeDateTime(d)
				s.UpdatedAt = v
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"updatedAt\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Task")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00110111,
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
				if fieldIdx < len(jsonFieldsNameOfTask) {
					name = jsonFieldsNameOfTask[fieldIdx]
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
func (s *Task) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Task) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Tasks) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Tasks) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("tasks")
		e.ArrStart()
		for _, elem := range s.Tasks {
			elem.Encode(e)
		}
		e.ArrEnd()
	}
	{
		e.FieldStart("hasNext")
		e.Bool(s.HasNext)
	}
}

var jsonFieldsNameOfTasks = [2]string{
	0: "tasks",
	1: "hasNext",
}

// Decode decodes Tasks from json.
func (s *Tasks) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Tasks to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "tasks":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				s.Tasks = make([]Task, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem Task
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.Tasks = append(s.Tasks, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"tasks\"")
			}
		case "hasNext":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Bool()
				s.HasNext = bool(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"hasNext\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Tasks")
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
				if fieldIdx < len(jsonFieldsNameOfTasks) {
					name = jsonFieldsNameOfTasks[fieldIdx]
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
func (s *Tasks) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Tasks) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *UpdateProjectReq) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *UpdateProjectReq) encodeFields(e *jx.Encoder) {
	{
		if s.Name.Set {
			e.FieldStart("name")
			s.Name.Encode(e)
		}
	}
}

var jsonFieldsNameOfUpdateProjectReq = [1]string{
	0: "name",
}

// Decode decodes UpdateProjectReq from json.
func (s *UpdateProjectReq) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode UpdateProjectReq to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "name":
			if err := func() error {
				s.Name.Reset()
				if err := s.Name.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"name\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode UpdateProjectReq")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *UpdateProjectReq) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *UpdateProjectReq) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *UpdateTaskReq) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *UpdateTaskReq) encodeFields(e *jx.Encoder) {
	{
		if s.IsCompleted.Set {
			e.FieldStart("isCompleted")
			s.IsCompleted.Encode(e)
		}
	}
}

var jsonFieldsNameOfUpdateTaskReq = [1]string{
	0: "isCompleted",
}

// Decode decodes UpdateTaskReq from json.
func (s *UpdateTaskReq) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode UpdateTaskReq to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "isCompleted":
			if err := func() error {
				s.IsCompleted.Reset()
				if err := s.IsCompleted.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"isCompleted\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode UpdateTaskReq")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *UpdateTaskReq) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *UpdateTaskReq) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}