// Code generated by protoc-gen-go.
// source: google.golang.org/cloud/bigtable/internal/service_proto/bigtable_service_messages.proto
// DO NOT EDIT!

/*
Package google_bigtable_v1 is a generated protocol buffer package.

It is generated from these files:
	google.golang.org/cloud/bigtable/internal/service_proto/bigtable_service_messages.proto
	google.golang.org/cloud/bigtable/internal/service_proto/bigtable_service.proto

It has these top-level messages:
	ReadRowsRequest
	ReadRowsResponse
	SampleRowKeysRequest
	SampleRowKeysResponse
	MutateRowRequest
	CheckAndMutateRowRequest
	CheckAndMutateRowResponse
	ReadModifyWriteRowRequest
*/
package google_bigtable_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_bigtable_v11 "google.golang.org/cloud/bigtable/internal/data_proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Request message for BigtableServer.ReadRows.
type ReadRowsRequest struct {
	// The unique name of the table from which to read.
	TableName string `protobuf:"bytes,1,opt,name=table_name" json:"table_name,omitempty"`
	// If neither row_key nor row_range is set, reads from all rows.
	//
	// Types that are valid to be assigned to Target:
	//	*ReadRowsRequest_RowKey
	//	*ReadRowsRequest_RowRange
	Target isReadRowsRequest_Target `protobuf_oneof:"target"`
	// The filter to apply to the contents of the specified row(s). If unset,
	// reads the entire table.
	Filter *google_bigtable_v11.RowFilter `protobuf:"bytes,5,opt,name=filter" json:"filter,omitempty"`
	// By default, rows are read sequentially, producing results which are
	// guaranteed to arrive in increasing row order. Setting
	// "allow_row_interleaving" to true allows multiple rows to be interleaved in
	// the response stream, which increases throughput but breaks this guarantee,
	// and may force the client to use more memory to buffer partially-received
	// rows. Cannot be set to true when specifying "num_rows_limit".
	AllowRowInterleaving bool `protobuf:"varint,6,opt,name=allow_row_interleaving" json:"allow_row_interleaving,omitempty"`
	// The read will terminate after committing to N rows' worth of results. The
	// default (zero) is to return all results.
	// Note that "allow_row_interleaving" cannot be set to true when this is set.
	NumRowsLimit int64 `protobuf:"varint,7,opt,name=num_rows_limit" json:"num_rows_limit,omitempty"`
}

func (m *ReadRowsRequest) Reset()         { *m = ReadRowsRequest{} }
func (m *ReadRowsRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRowsRequest) ProtoMessage()    {}

type isReadRowsRequest_Target interface {
	isReadRowsRequest_Target()
}

type ReadRowsRequest_RowKey struct {
	RowKey []byte `protobuf:"bytes,2,opt,name=row_key,proto3,oneof"`
}
type ReadRowsRequest_RowRange struct {
	RowRange *google_bigtable_v11.RowRange `protobuf:"bytes,3,opt,name=row_range,oneof"`
}

func (*ReadRowsRequest_RowKey) isReadRowsRequest_Target()   {}
func (*ReadRowsRequest_RowRange) isReadRowsRequest_Target() {}

func (m *ReadRowsRequest) GetTarget() isReadRowsRequest_Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *ReadRowsRequest) GetRowKey() []byte {
	if x, ok := m.GetTarget().(*ReadRowsRequest_RowKey); ok {
		return x.RowKey
	}
	return nil
}

func (m *ReadRowsRequest) GetRowRange() *google_bigtable_v11.RowRange {
	if x, ok := m.GetTarget().(*ReadRowsRequest_RowRange); ok {
		return x.RowRange
	}
	return nil
}

func (m *ReadRowsRequest) GetFilter() *google_bigtable_v11.RowFilter {
	if m != nil {
		return m.Filter
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ReadRowsRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), []interface{}) {
	return _ReadRowsRequest_OneofMarshaler, _ReadRowsRequest_OneofUnmarshaler, []interface{}{
		(*ReadRowsRequest_RowKey)(nil),
		(*ReadRowsRequest_RowRange)(nil),
	}
}

func _ReadRowsRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ReadRowsRequest)
	// target
	switch x := m.Target.(type) {
	case *ReadRowsRequest_RowKey:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.RowKey)
	case *ReadRowsRequest_RowRange:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RowRange); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ReadRowsRequest.Target has unexpected type %T", x)
	}
	return nil
}

func _ReadRowsRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ReadRowsRequest)
	switch tag {
	case 2: // target.row_key
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Target = &ReadRowsRequest_RowKey{x}
		return true, err
	case 3: // target.row_range
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_bigtable_v11.RowRange)
		err := b.DecodeMessage(msg)
		m.Target = &ReadRowsRequest_RowRange{msg}
		return true, err
	default:
		return false, nil
	}
}

// Response message for BigtableService.ReadRows.
type ReadRowsResponse struct {
	// The key of the row for which we're receiving data.
	// Results will be received in increasing row key order, unless
	// "allow_row_interleaving" was specified in the request.
	RowKey []byte `protobuf:"bytes,1,opt,name=row_key,proto3" json:"row_key,omitempty"`
	// One or more chunks of the row specified by "row_key".
	Chunks []*ReadRowsResponse_Chunk `protobuf:"bytes,2,rep,name=chunks" json:"chunks,omitempty"`
}

func (m *ReadRowsResponse) Reset()         { *m = ReadRowsResponse{} }
func (m *ReadRowsResponse) String() string { return proto.CompactTextString(m) }
func (*ReadRowsResponse) ProtoMessage()    {}

func (m *ReadRowsResponse) GetChunks() []*ReadRowsResponse_Chunk {
	if m != nil {
		return m.Chunks
	}
	return nil
}

// Specifies a piece of a row's contents returned as part of the read
// response stream.
type ReadRowsResponse_Chunk struct {
	// Types that are valid to be assigned to Chunk:
	//	*ReadRowsResponse_Chunk_RowContents
	//	*ReadRowsResponse_Chunk_ResetRow
	//	*ReadRowsResponse_Chunk_CommitRow
	Chunk isReadRowsResponse_Chunk_Chunk `protobuf_oneof:"chunk"`
}

func (m *ReadRowsResponse_Chunk) Reset()         { *m = ReadRowsResponse_Chunk{} }
func (m *ReadRowsResponse_Chunk) String() string { return proto.CompactTextString(m) }
func (*ReadRowsResponse_Chunk) ProtoMessage()    {}

type isReadRowsResponse_Chunk_Chunk interface {
	isReadRowsResponse_Chunk_Chunk()
}

type ReadRowsResponse_Chunk_RowContents struct {
	RowContents *google_bigtable_v11.Family `protobuf:"bytes,1,opt,name=row_contents,oneof"`
}
type ReadRowsResponse_Chunk_ResetRow struct {
	ResetRow bool `protobuf:"varint,2,opt,name=reset_row,oneof"`
}
type ReadRowsResponse_Chunk_CommitRow struct {
	CommitRow bool `protobuf:"varint,3,opt,name=commit_row,oneof"`
}

func (*ReadRowsResponse_Chunk_RowContents) isReadRowsResponse_Chunk_Chunk() {}
func (*ReadRowsResponse_Chunk_ResetRow) isReadRowsResponse_Chunk_Chunk()    {}
func (*ReadRowsResponse_Chunk_CommitRow) isReadRowsResponse_Chunk_Chunk()   {}

func (m *ReadRowsResponse_Chunk) GetChunk() isReadRowsResponse_Chunk_Chunk {
	if m != nil {
		return m.Chunk
	}
	return nil
}

func (m *ReadRowsResponse_Chunk) GetRowContents() *google_bigtable_v11.Family {
	if x, ok := m.GetChunk().(*ReadRowsResponse_Chunk_RowContents); ok {
		return x.RowContents
	}
	return nil
}

func (m *ReadRowsResponse_Chunk) GetResetRow() bool {
	if x, ok := m.GetChunk().(*ReadRowsResponse_Chunk_ResetRow); ok {
		return x.ResetRow
	}
	return false
}

func (m *ReadRowsResponse_Chunk) GetCommitRow() bool {
	if x, ok := m.GetChunk().(*ReadRowsResponse_Chunk_CommitRow); ok {
		return x.CommitRow
	}
	return false
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ReadRowsResponse_Chunk) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), []interface{}) {
	return _ReadRowsResponse_Chunk_OneofMarshaler, _ReadRowsResponse_Chunk_OneofUnmarshaler, []interface{}{
		(*ReadRowsResponse_Chunk_RowContents)(nil),
		(*ReadRowsResponse_Chunk_ResetRow)(nil),
		(*ReadRowsResponse_Chunk_CommitRow)(nil),
	}
}

func _ReadRowsResponse_Chunk_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ReadRowsResponse_Chunk)
	// chunk
	switch x := m.Chunk.(type) {
	case *ReadRowsResponse_Chunk_RowContents:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RowContents); err != nil {
			return err
		}
	case *ReadRowsResponse_Chunk_ResetRow:
		t := uint64(0)
		if x.ResetRow {
			t = 1
		}
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *ReadRowsResponse_Chunk_CommitRow:
		t := uint64(0)
		if x.CommitRow {
			t = 1
		}
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case nil:
	default:
		return fmt.Errorf("ReadRowsResponse_Chunk.Chunk has unexpected type %T", x)
	}
	return nil
}

func _ReadRowsResponse_Chunk_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ReadRowsResponse_Chunk)
	switch tag {
	case 1: // chunk.row_contents
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_bigtable_v11.Family)
		err := b.DecodeMessage(msg)
		m.Chunk = &ReadRowsResponse_Chunk_RowContents{msg}
		return true, err
	case 2: // chunk.reset_row
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Chunk = &ReadRowsResponse_Chunk_ResetRow{x != 0}
		return true, err
	case 3: // chunk.commit_row
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Chunk = &ReadRowsResponse_Chunk_CommitRow{x != 0}
		return true, err
	default:
		return false, nil
	}
}

// Request message for BigtableService.SampleRowKeys.
type SampleRowKeysRequest struct {
	// The unique name of the table from which to sample row keys.
	TableName string `protobuf:"bytes,1,opt,name=table_name" json:"table_name,omitempty"`
}

func (m *SampleRowKeysRequest) Reset()         { *m = SampleRowKeysRequest{} }
func (m *SampleRowKeysRequest) String() string { return proto.CompactTextString(m) }
func (*SampleRowKeysRequest) ProtoMessage()    {}

// Response message for BigtableService.SampleRowKeys.
type SampleRowKeysResponse struct {
	// Sorted streamed sequence of sample row keys in the table. The table might
	// have contents before the first row key in the list and after the last one,
	// but a key containing the empty string indicates "end of table" and will be
	// the last response given, if present.
	// Note that row keys in this list may not have ever been written to or read
	// from, and users should therefore not make any assumptions about the row key
	// structure that are specific to their use case.
	RowKey []byte `protobuf:"bytes,1,opt,name=row_key,proto3" json:"row_key,omitempty"`
	// Approximate total storage space used by all rows in the table which precede
	// "row_key". Buffering the contents of all rows between two subsequent
	// samples would require space roughly equal to the difference in their
	// "offset_bytes" fields.
	OffsetBytes int64 `protobuf:"varint,2,opt,name=offset_bytes" json:"offset_bytes,omitempty"`
}

func (m *SampleRowKeysResponse) Reset()         { *m = SampleRowKeysResponse{} }
func (m *SampleRowKeysResponse) String() string { return proto.CompactTextString(m) }
func (*SampleRowKeysResponse) ProtoMessage()    {}

// Request message for BigtableService.MutateRow.
type MutateRowRequest struct {
	// The unique name of the table to which the mutation should be applied.
	TableName string `protobuf:"bytes,1,opt,name=table_name" json:"table_name,omitempty"`
	// The key of the row to which the mutation should be applied.
	RowKey []byte `protobuf:"bytes,2,opt,name=row_key,proto3" json:"row_key,omitempty"`
	// Changes to be atomically applied to the specified row. Entries are applied
	// in order, meaning that earlier mutations can be masked by later ones.
	// Must contain at least one entry and at most 100000.
	Mutations []*google_bigtable_v11.Mutation `protobuf:"bytes,3,rep,name=mutations" json:"mutations,omitempty"`
}

func (m *MutateRowRequest) Reset()         { *m = MutateRowRequest{} }
func (m *MutateRowRequest) String() string { return proto.CompactTextString(m) }
func (*MutateRowRequest) ProtoMessage()    {}

func (m *MutateRowRequest) GetMutations() []*google_bigtable_v11.Mutation {
	if m != nil {
		return m.Mutations
	}
	return nil
}

// Request message for BigtableService.CheckAndMutateRowRequest
type CheckAndMutateRowRequest struct {
	// The unique name of the table to which the conditional mutation should be
	// applied.
	TableName string `protobuf:"bytes,1,opt,name=table_name" json:"table_name,omitempty"`
	// The key of the row to which the conditional mutation should be applied.
	RowKey []byte `protobuf:"bytes,2,opt,name=row_key,proto3" json:"row_key,omitempty"`
	// The filter to be applied to the contents of the specified row. Depending
	// on whether or not any results are yielded, either "true_mutations" or
	// "false_mutations" will be executed. If unset, checks that the row contains
	// any values at all.
	PredicateFilter *google_bigtable_v11.RowFilter `protobuf:"bytes,6,opt,name=predicate_filter" json:"predicate_filter,omitempty"`
	// Changes to be atomically applied to the specified row if "predicate_filter"
	// yields at least one cell when applied to "row_key". Entries are applied in
	// order, meaning that earlier mutations can be masked by later ones.
	// Must contain at least one entry if "false_mutations" is empty, and at most
	// 100000.
	TrueMutations []*google_bigtable_v11.Mutation `protobuf:"bytes,4,rep,name=true_mutations" json:"true_mutations,omitempty"`
	// Changes to be atomically applied to the specified row if "predicate_filter"
	// does not yield any cells when applied to "row_key". Entries are applied in
	// order, meaning that earlier mutations can be masked by later ones.
	// Must contain at least one entry if "true_mutations" is empty, and at most
	// 100000.
	FalseMutations []*google_bigtable_v11.Mutation `protobuf:"bytes,5,rep,name=false_mutations" json:"false_mutations,omitempty"`
}

func (m *CheckAndMutateRowRequest) Reset()         { *m = CheckAndMutateRowRequest{} }
func (m *CheckAndMutateRowRequest) String() string { return proto.CompactTextString(m) }
func (*CheckAndMutateRowRequest) ProtoMessage()    {}

func (m *CheckAndMutateRowRequest) GetPredicateFilter() *google_bigtable_v11.RowFilter {
	if m != nil {
		return m.PredicateFilter
	}
	return nil
}

func (m *CheckAndMutateRowRequest) GetTrueMutations() []*google_bigtable_v11.Mutation {
	if m != nil {
		return m.TrueMutations
	}
	return nil
}

func (m *CheckAndMutateRowRequest) GetFalseMutations() []*google_bigtable_v11.Mutation {
	if m != nil {
		return m.FalseMutations
	}
	return nil
}

// Response message for BigtableService.CheckAndMutateRowRequest.
type CheckAndMutateRowResponse struct {
	// Whether or not the request's "predicate_filter" yielded any results for
	// the specified row.
	PredicateMatched bool `protobuf:"varint,1,opt,name=predicate_matched" json:"predicate_matched,omitempty"`
}

func (m *CheckAndMutateRowResponse) Reset()         { *m = CheckAndMutateRowResponse{} }
func (m *CheckAndMutateRowResponse) String() string { return proto.CompactTextString(m) }
func (*CheckAndMutateRowResponse) ProtoMessage()    {}

// Request message for BigtableService.ReadModifyWriteRowRequest.
type ReadModifyWriteRowRequest struct {
	// The unique name of the table to which the read/modify/write rules should be
	// applied.
	TableName string `protobuf:"bytes,1,opt,name=table_name" json:"table_name,omitempty"`
	// The key of the row to which the read/modify/write rules should be applied.
	RowKey []byte `protobuf:"bytes,2,opt,name=row_key,proto3" json:"row_key,omitempty"`
	// Rules specifying how the specified row's contents are to be transformed
	// into writes. Entries are applied in order, meaning that earlier rules will
	// affect the results of later ones.
	Rules []*google_bigtable_v11.ReadModifyWriteRule `protobuf:"bytes,3,rep,name=rules" json:"rules,omitempty"`
}

func (m *ReadModifyWriteRowRequest) Reset()         { *m = ReadModifyWriteRowRequest{} }
func (m *ReadModifyWriteRowRequest) String() string { return proto.CompactTextString(m) }
func (*ReadModifyWriteRowRequest) ProtoMessage()    {}

func (m *ReadModifyWriteRowRequest) GetRules() []*google_bigtable_v11.ReadModifyWriteRule {
	if m != nil {
		return m.Rules
	}
	return nil
}
