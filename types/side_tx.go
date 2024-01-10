package types

import (
	"encoding/binary"
	"errors"
	"fmt"
	cmn "github.com/dojimanetwork/dojimamint/libs/bytes"
)

// SideTxResult side tx result for vote
type SideTxResult struct {
	TxHash []byte `json:"tx_hash"`
	Result int32  `json:"result"`
	Sig    []byte `json:"sig"`
}

func (sp *SideTxResult) String() string {
	if sp == nil {
		return ""
	}

	return fmt.Sprintf("SideTxResult{%X (%v) %X}",
		cmn.Fingerprint(sp.TxHash),
		sp.Result,
		cmn.Fingerprint(sp.Sig),
	)
}

func (m *SideTxResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SideTxResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	var idx int
	if len(dAtA) < m.Size() {
		return 0, errors.New("buffer too small")
	}

	// Encoding Result (int32)
	binary.BigEndian.PutUint32(dAtA[idx:], uint32(m.Result))
	idx += 4

	// Encoding TxHash
	copy(dAtA[idx:], m.TxHash)
	idx += len(m.TxHash)

	// Encoding Sig
	copy(dAtA[idx:], m.Sig)
	idx += len(m.Sig)

	return idx, nil // idx is the number of bytes written
}

func (m *SideTxResult) Size() int {
	return len(m.TxHash) + len(m.Sig) + 4 // 4 bytes for int32 Result
}

// SideTxResultWithData side tx result with data for vote
type SideTxResultWithData struct {
	SideTxResult

	Data []byte `json:"data"`
}

// GetBytes returns data bytes for sign
func (sp *SideTxResultWithData) GetBytes() []byte {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, uint32(sp.Result))

	data := make([]byte, 0)
	data = append(data, bs[3]) // use last byte as result
	if len(sp.Data) > 0 {
		data = append(data, sp.Data...)
	}
	return data
}

func (sp *SideTxResultWithData) String() string {
	if sp == nil {
		return ""
	}

	return fmt.Sprintf("SideTxResultWithData {%s %X}",
		sp.SideTxResult.String(),
		cmn.Fingerprint(sp.Data),
	)
}
