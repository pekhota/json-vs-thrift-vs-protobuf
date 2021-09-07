// Code generated by Thrift Compiler (0.14.2). DO NOT EDIT.

package thrifttimestamp

import(
	"bytes"
	"context"
	"fmt"
	"time"
	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = time.Now
var _ = bytes.Equal

// Attributes:
//  - Seconds
//  - Nanos
type Timestamp struct {
  Seconds int64 `thrift:"seconds,1" db:"seconds" json:"seconds"`
  Nanos int32 `thrift:"nanos,2" db:"nanos" json:"nanos"`
}

func NewTimestamp() *Timestamp {
  return &Timestamp{}
}


func (p *Timestamp) GetSeconds() int64 {
  return p.Seconds
}

func (p *Timestamp) GetNanos() int32 {
  return p.Nanos
}
func (p *Timestamp) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.I64 {
        if err := p.ReadField1(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField2(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *Timestamp)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI64(ctx); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Seconds = v
}
  return nil
}

func (p *Timestamp)  ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(ctx); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.Nanos = v
}
  return nil
}

func (p *Timestamp) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "Timestamp"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(ctx, oprot); err != nil { return err }
    if err := p.writeField2(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *Timestamp) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "seconds", thrift.I64, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:seconds: ", p), err) }
  if err := oprot.WriteI64(ctx, int64(p.Seconds)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.seconds (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:seconds: ", p), err) }
  return err
}

func (p *Timestamp) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "nanos", thrift.I32, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:nanos: ", p), err) }
  if err := oprot.WriteI32(ctx, int32(p.Nanos)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.nanos (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:nanos: ", p), err) }
  return err
}

func (p *Timestamp) Equals(other *Timestamp) bool {
  if p == other {
    return true
  } else if p == nil || other == nil {
    return false
  }
  if p.Seconds != other.Seconds { return false }
  if p.Nanos != other.Nanos { return false }
  return true
}

func (p *Timestamp) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("Timestamp(%+v)", *p)
}
