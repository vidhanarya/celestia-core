package types

import (
	// it is ok to use math/rand here: we do not need a cryptographically secure random
	// number generator here and we can run the tests a bit faster
	stdbytes "bytes"
	"crypto/rand"
	"encoding/hex"
	"math"
	"os"
	"reflect"
	"sort"
	"testing"
	"time"

	gogotypes "github.com/gogo/protobuf/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/merkle"
	"github.com/tendermint/tendermint/crypto/tmhash"
	"github.com/tendermint/tendermint/libs/bits"
	"github.com/tendermint/tendermint/libs/bytes"
	tmrand "github.com/tendermint/tendermint/libs/rand"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmversion "github.com/tendermint/tendermint/proto/tendermint/version"
	tmtime "github.com/tendermint/tendermint/types/time"
	"github.com/tendermint/tendermint/version"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestBlockAddEvidence(t *testing.T) {
	txs := []Tx{Tx("foo"), Tx("bar")}
	lastID := makeBlockIDRandom()
	h := int64(3)

	voteSet, _, vals := randVoteSet(h-1, 1, tmproto.PrecommitType, 10, 1)
	commit, err := MakeCommit(lastID, h-1, 1, voteSet, vals, time.Now())
	require.NoError(t, err)

	ev := NewMockDuplicateVoteEvidenceWithValidator(h, time.Now(), vals[0], "block-test-chain")
	evList := []Evidence{ev}

	block := MakeBlock(h, makeData(txs, evList, nil), commit)
	require.NotNil(t, block)
	require.Equal(t, 1, len(block.Data.Evidence.Evidence))
	require.NotNil(t, block.EvidenceHash)
}

func TestBlockValidateBasic(t *testing.T) {
	require.Error(t, (*Block)(nil).ValidateBasic())

	txs := []Tx{Tx("foo"), Tx("bar")}
	lastID := makeBlockIDRandom()
	h := int64(3)

	voteSet, valSet, vals := randVoteSet(h-1, 1, tmproto.PrecommitType, 10, 1)
	commit, err := MakeCommit(lastID, h-1, 1, voteSet, vals, time.Now())
	require.NoError(t, err)

	ev := NewMockDuplicateVoteEvidenceWithValidator(h, time.Now(), vals[0], "block-test-chain")
	evList := []Evidence{ev}

	testCases := []struct {
		testName      string
		malleateBlock func(*Block)
		expErr        bool
	}{
		{"Make Block", func(blk *Block) {}, false},
		{"Make Block w/ proposer Addr", func(blk *Block) { blk.ProposerAddress = valSet.GetProposer().Address }, false},
		{"Negative Height", func(blk *Block) { blk.Height = -1 }, true},
		{"Remove 1/2 the commits", func(blk *Block) {
			blk.LastCommit.Signatures = commit.Signatures[:commit.Size()/2]
			blk.LastCommit.hash = nil // clear hash or change wont be noticed
		}, true},
		{"Remove LastCommitHash", func(blk *Block) { blk.LastCommitHash = []byte("something else") }, true},
		{"Tampered EvidenceHash", func(blk *Block) {
			blk.EvidenceHash = []byte("something else")
		}, true},
		{"Incorrect block protocol version", func(blk *Block) {
			blk.Version.Block = 1
		}, true},
	}
	for i, tc := range testCases {
		tc := tc
		i := i
		t.Run(tc.testName, func(t *testing.T) {
			block := MakeBlock(h, makeData(txs, evList, nil), commit)
			block.ProposerAddress = valSet.GetProposer().Address
			tc.malleateBlock(block)
			err = block.ValidateBasic()
			assert.Equal(t, tc.expErr, err != nil, "#%d: %v", i, err)
		})
	}
}

func TestBlockHash(t *testing.T) {
	assert.Nil(t, (*Block)(nil).Hash())
	assert.Nil(t, MakeBlock(int64(3), makeData([]Tx{Tx("Hello World")}, nil, nil), nil).Hash())
}

func TestBlockMakePartSet(t *testing.T) {
	assert.Nil(t, (*Block)(nil).MakePartSet(2))

	partSet := MakeBlock(int64(3), makeData([]Tx{Tx("Hello World")}, nil, nil), nil).MakePartSet(1024)
	assert.NotNil(t, partSet)
	assert.EqualValues(t, 1, partSet.Total())
}

func TestBlockMakePartSetWithEvidence(t *testing.T) {
	assert.Nil(t, (*Block)(nil).MakePartSet(2))

	lastID := makeBlockIDRandom()
	h := int64(3)

	voteSet, _, vals := randVoteSet(h-1, 1, tmproto.PrecommitType, 10, 1)
	commit, err := MakeCommit(lastID, h-1, 1, voteSet, vals, time.Now())
	require.NoError(t, err)

	ev := NewMockDuplicateVoteEvidenceWithValidator(h, time.Now(), vals[0], "block-test-chain")
	evList := []Evidence{ev}

	partSet := MakeBlock(h, makeData([]Tx{Tx("Hello World")}, evList, nil), commit).MakePartSet(512)
	assert.NotNil(t, partSet)
	assert.EqualValues(t, 4, partSet.Total())
}

func TestBlockHashesTo(t *testing.T) {
	assert.False(t, (*Block)(nil).HashesTo(nil))

	lastID := makeBlockIDRandom()
	h := int64(3)
	voteSet, valSet, vals := randVoteSet(h-1, 1, tmproto.PrecommitType, 10, 1)
	commit, err := MakeCommit(lastID, h-1, 1, voteSet, vals, time.Now())
	require.NoError(t, err)

	ev := NewMockDuplicateVoteEvidenceWithValidator(h, time.Now(), vals[0], "block-test-chain")
	evList := []Evidence{ev}

	block := MakeBlock(h, makeData([]Tx{Tx("Hello World")}, evList, nil), commit)
	block.ValidatorsHash = valSet.Hash()
	assert.False(t, block.HashesTo([]byte{}))
	assert.False(t, block.HashesTo([]byte("something else")))
	assert.True(t, block.HashesTo(block.Hash()))
}

func TestBlockSize(t *testing.T) {
	size := MakeBlock(int64(3), makeData([]Tx{Tx("Hello World")}, nil, nil), nil).Size()
	if size <= 0 {
		t.Fatal("Size of the block is zero or negative")
	}
}

func TestBlockString(t *testing.T) {
	assert.Equal(t, "nil-Block", (*Block)(nil).String())
	assert.Equal(t, "nil-Block", (*Block)(nil).StringIndented(""))
	assert.Equal(t, "nil-Block", (*Block)(nil).StringShort())

	block := MakeBlock(int64(3), makeData([]Tx{Tx("Hello World")}, nil, nil), nil)
	assert.NotEqual(t, "nil-Block", block.String())
	assert.NotEqual(t, "nil-Block", block.StringIndented(""))
	assert.NotEqual(t, "nil-Block", block.StringShort())
}

func makeBlockIDRandom() BlockID {
	var (
		blockHash   = make([]byte, tmhash.Size)
		partSetHash = make([]byte, tmhash.Size)
	)
	rand.Read(blockHash)   //nolint: errcheck // ignore errcheck for read
	rand.Read(partSetHash) //nolint: errcheck // ignore errcheck for read
	return BlockID{blockHash, PartSetHeader{123, partSetHash}}
}

func makeBlockID(hash []byte, partSetSize uint32, partSetHash []byte) BlockID {
	var (
		h   = make([]byte, tmhash.Size)
		psH = make([]byte, tmhash.Size)
	)
	copy(h, hash)
	copy(psH, partSetHash)
	return BlockID{
		Hash: h,
		PartSetHeader: PartSetHeader{
			Total: partSetSize,
			Hash:  psH,
		},
	}
}

var nilBytes []byte

func TestNilHeaderHashDoesntCrash(t *testing.T) {
	assert.Equal(t, nilBytes, []byte((*Header)(nil).Hash()))
	assert.Equal(t, nilBytes, []byte((new(Header)).Hash()))
}

func TestCommit(t *testing.T) {
	lastID := makeBlockIDRandom()
	h := int64(3)
	voteSet, _, vals := randVoteSet(h-1, 1, tmproto.PrecommitType, 10, 1)
	commit, err := MakeCommit(lastID, h-1, 1, voteSet, vals, time.Now())
	require.NoError(t, err)

	assert.Equal(t, h-1, commit.Height)
	assert.EqualValues(t, 1, commit.Round)
	assert.Equal(t, tmproto.PrecommitType, tmproto.SignedMsgType(commit.Type()))
	if commit.Size() <= 0 {
		t.Fatalf("commit %v has a zero or negative size: %d", commit, commit.Size())
	}

	require.NotNil(t, commit.BitArray())
	assert.Equal(t, bits.NewBitArray(10).Size(), commit.BitArray().Size())

	assert.Equal(t, voteSet.GetByIndex(0), commit.GetByIndex(0))
	assert.True(t, commit.IsCommit())
}

func TestCommitValidateBasic(t *testing.T) {
	testCases := []struct {
		testName       string
		malleateCommit func(*Commit)
		expectErr      bool
	}{
		{"Random Commit", func(com *Commit) {}, false},
		{"Incorrect signature", func(com *Commit) { com.Signatures[0].Signature = []byte{0} }, false},
		{"Incorrect height", func(com *Commit) { com.Height = int64(-100) }, true},
		{"Incorrect round", func(com *Commit) { com.Round = -100 }, true},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			com := randCommit(time.Now())
			tc.malleateCommit(com)
			assert.Equal(t, tc.expectErr, com.ValidateBasic() != nil, "Validate Basic had an unexpected result")
		})
	}
}

func TestMaxCommitBytes(t *testing.T) {
	// time is varint encoded so need to pick the max.
	// year int, month Month, day, hour, min, sec, nsec int, loc *Location
	timestamp := time.Date(math.MaxInt64, 0, 0, 0, 0, 0, math.MaxInt64, time.UTC)

	cs := CommitSig{
		BlockIDFlag:      BlockIDFlagNil,
		ValidatorAddress: crypto.AddressHash([]byte("validator_address")),
		Timestamp:        timestamp,
		Signature:        crypto.CRandBytes(MaxSignatureSize),
	}

	pbSig := cs.ToProto()
	// test that a single commit sig doesn't exceed max commit sig bytes
	assert.EqualValues(t, MaxCommitSigBytes, pbSig.Size())

	// check size with a single commit
	commit := &Commit{
		Height: math.MaxInt64,
		Round:  math.MaxInt32,
		BlockID: BlockID{
			Hash: tmhash.Sum([]byte("blockID_hash")),
			PartSetHeader: PartSetHeader{
				Total: math.MaxInt32,
				Hash:  tmhash.Sum([]byte("blockID_part_set_header_hash")),
			},
		},
		Signatures: []CommitSig{cs},
	}

	pb := commit.ToProto()

	assert.EqualValues(t, MaxCommitBytes(1), int64(pb.Size()))

	// check the upper bound of the commit size
	for i := 1; i < MaxVotesCount; i++ {
		commit.Signatures = append(commit.Signatures, cs)
	}

	pb = commit.ToProto()

	assert.EqualValues(t, MaxCommitBytes(MaxVotesCount), int64(pb.Size()))

}

func TestHeaderHash(t *testing.T) {
	testCases := []struct {
		desc       string
		header     *Header
		expectHash bytes.HexBytes
	}{
		{"Generates expected hash", &Header{
			Version:            tmversion.Consensus{Block: 1, App: 2},
			ChainID:            "chainId",
			Height:             3,
			Time:               time.Date(2019, 10, 13, 16, 14, 44, 0, time.UTC),
			LastBlockID:        makeBlockID(make([]byte, tmhash.Size), 6, make([]byte, tmhash.Size)),
			LastCommitHash:     tmhash.Sum([]byte("last_commit_hash")),
			DataHash:           tmhash.Sum([]byte("data_hash")),
			ValidatorsHash:     tmhash.Sum([]byte("validators_hash")),
			NextValidatorsHash: tmhash.Sum([]byte("next_validators_hash")),
			ConsensusHash:      tmhash.Sum([]byte("consensus_hash")),
			AppHash:            tmhash.Sum([]byte("app_hash")),
			LastResultsHash:    tmhash.Sum([]byte("last_results_hash")),
			EvidenceHash:       tmhash.Sum([]byte("evidence_hash")),
			ProposerAddress:    crypto.AddressHash([]byte("proposer_address")),
		}, hexBytesFromString("F740121F553B5418C3EFBD343C2DBFE9E007BB67B0D020A0741374BAB65242A4")},
		{"nil header yields nil", nil, nil},
		{"nil ValidatorsHash yields nil", &Header{
			Version:            tmversion.Consensus{Block: 1, App: 2},
			ChainID:            "chainId",
			Height:             3,
			Time:               time.Date(2019, 10, 13, 16, 14, 44, 0, time.UTC),
			LastBlockID:        makeBlockID(make([]byte, tmhash.Size), 6, make([]byte, tmhash.Size)),
			LastCommitHash:     tmhash.Sum([]byte("last_commit_hash")),
			DataHash:           tmhash.Sum([]byte("data_hash")),
			ValidatorsHash:     nil,
			NextValidatorsHash: tmhash.Sum([]byte("next_validators_hash")),
			ConsensusHash:      tmhash.Sum([]byte("consensus_hash")),
			AppHash:            tmhash.Sum([]byte("app_hash")),
			LastResultsHash:    tmhash.Sum([]byte("last_results_hash")),
			EvidenceHash:       tmhash.Sum([]byte("evidence_hash")),
			ProposerAddress:    crypto.AddressHash([]byte("proposer_address")),
		}, nil},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			assert.Equal(t, tc.expectHash, tc.header.Hash())

			// We also make sure that all fields are hashed in struct order, and that all
			// fields in the test struct are non-zero.
			if tc.header != nil && tc.expectHash != nil {
				byteSlices := [][]byte{}

				s := reflect.ValueOf(*tc.header)
				for i := 0; i < s.NumField(); i++ {
					f := s.Field(i)

					assert.False(t, f.IsZero(), "Found zero-valued field %v",
						s.Type().Field(i).Name)

					switch f := f.Interface().(type) {
					case int64, bytes.HexBytes, string:
						byteSlices = append(byteSlices, cdcEncode(f))
					case time.Time:
						bz, err := gogotypes.StdTimeMarshal(f)
						require.NoError(t, err)
						byteSlices = append(byteSlices, bz)
					case tmversion.Consensus:
						bz, err := f.Marshal()
						require.NoError(t, err)
						byteSlices = append(byteSlices, bz)
					case BlockID:
						pbbi := f.ToProto()
						bz, err := pbbi.Marshal()
						require.NoError(t, err)
						byteSlices = append(byteSlices, bz)
					default:
						t.Errorf("unknown type %T", f)
					}
				}
				assert.Equal(t,
					bytes.HexBytes(merkle.HashFromByteSlices(byteSlices)), tc.header.Hash())
			}
		})
	}
}

func TestMaxHeaderBytes(t *testing.T) {
	// Construct a UTF-8 string of MaxChainIDLen length using the supplementary
	// characters.
	// Each supplementary character takes 4 bytes.
	// http://www.i18nguy.com/unicode/supplementary-test.html
	maxChainID := ""
	for i := 0; i < MaxChainIDLen; i++ {
		maxChainID += "𠜎"
	}

	// time is varint encoded so need to pick the max.
	// year int, month Month, day, hour, min, sec, nsec int, loc *Location
	timestamp := time.Date(math.MaxInt64, 0, 0, 0, 0, 0, math.MaxInt64, time.UTC)

	h := Header{
		Version:            tmversion.Consensus{Block: math.MaxInt64, App: math.MaxInt64},
		ChainID:            maxChainID,
		Height:             math.MaxInt64,
		Time:               timestamp,
		LastBlockID:        makeBlockID(make([]byte, tmhash.Size), math.MaxInt32, make([]byte, tmhash.Size)),
		LastCommitHash:     tmhash.Sum([]byte("last_commit_hash")),
		DataHash:           tmhash.Sum([]byte("data_hash")),
		ValidatorsHash:     tmhash.Sum([]byte("validators_hash")),
		NextValidatorsHash: tmhash.Sum([]byte("next_validators_hash")),
		ConsensusHash:      tmhash.Sum([]byte("consensus_hash")),
		AppHash:            tmhash.Sum([]byte("app_hash")),
		LastResultsHash:    tmhash.Sum([]byte("last_results_hash")),
		EvidenceHash:       tmhash.Sum([]byte("evidence_hash")),
		ProposerAddress:    crypto.AddressHash([]byte("proposer_address")),
	}

	bz, err := h.ToProto().Marshal()
	require.NoError(t, err)

	assert.EqualValues(t, MaxHeaderBytes, int64(len(bz)))
}

func randCommit(now time.Time) *Commit {
	lastID := makeBlockIDRandom()
	h := int64(3)
	voteSet, _, vals := randVoteSet(h-1, 1, tmproto.PrecommitType, 10, 1)
	commit, err := MakeCommit(lastID, h-1, 1, voteSet, vals, now)
	if err != nil {
		panic(err)
	}
	return commit
}

func hexBytesFromString(s string) bytes.HexBytes {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return bytes.HexBytes(b)
}

func TestBlockMaxDataBytes(t *testing.T) {
	testCases := []struct {
		maxBytes      int64
		valsCount     int
		evidenceBytes int64
		panics        bool
		result        int64
	}{
		0: {-10, 1, 0, true, 0},
		1: {10, 1, 0, true, 0},
		2: {841, 1, 0, true, 0},
		3: {842, 1, 0, false, 0},
		4: {843, 1, 0, false, 1},
		5: {954, 2, 0, false, 1},
		6: {1053, 2, 100, false, 0},
	}

	for i, tc := range testCases {
		tc := tc
		if tc.panics {
			assert.Panics(t, func() {
				MaxDataBytes(tc.maxBytes, tc.evidenceBytes, tc.valsCount)
			}, "#%v", i)
		} else {
			assert.Equal(t,
				tc.result,
				MaxDataBytes(tc.maxBytes, tc.evidenceBytes, tc.valsCount),
				"#%v", i)
		}
	}
}

func TestBlockMaxDataBytesNoEvidence(t *testing.T) {
	testCases := []struct {
		maxBytes  int64
		valsCount int
		panics    bool
		result    int64
	}{
		0: {-10, 1, true, 0},
		1: {10, 1, true, 0},
		2: {841, 1, true, 0},
		3: {842, 1, false, 0},
		4: {843, 1, false, 1},
	}

	for i, tc := range testCases {
		tc := tc
		if tc.panics {
			assert.Panics(t, func() {
				MaxDataBytesNoEvidence(tc.maxBytes, tc.valsCount)
			}, "#%v", i)
		} else {
			assert.Equal(t,
				tc.result,
				MaxDataBytesNoEvidence(tc.maxBytes, tc.valsCount),
				"#%v", i)
		}
	}
}

func TestCommitToVoteSet(t *testing.T) {
	lastID := makeBlockIDRandom()
	h := int64(3)

	voteSet, valSet, vals := randVoteSet(h-1, 1, tmproto.PrecommitType, 10, 1)
	commit, err := MakeCommit(lastID, h-1, 1, voteSet, vals, time.Now())
	assert.NoError(t, err)

	chainID := voteSet.ChainID()
	voteSet2 := CommitToVoteSet(chainID, commit, valSet)

	for i := int32(0); int(i) < len(vals); i++ {
		vote1 := voteSet.GetByIndex(i)
		vote2 := voteSet2.GetByIndex(i)
		vote3 := commit.GetVote(i)

		vote1bz, err := vote1.ToProto().Marshal()
		require.NoError(t, err)
		vote2bz, err := vote2.ToProto().Marshal()
		require.NoError(t, err)
		vote3bz, err := vote3.ToProto().Marshal()
		require.NoError(t, err)
		assert.Equal(t, vote1bz, vote2bz)
		assert.Equal(t, vote1bz, vote3bz)
	}
}

func TestCommitToVoteSetWithVotesForNilBlock(t *testing.T) {
	blockID := makeBlockID([]byte("blockhash"), 1000, []byte("partshash"))

	const (
		height = int64(3)
		round  = 0
	)

	type commitVoteTest struct {
		blockIDs      []BlockID
		numVotes      []int // must sum to numValidators
		numValidators int
		valid         bool
	}

	testCases := []commitVoteTest{
		{[]BlockID{blockID, {}}, []int{67, 33}, 100, true},
	}

	for _, tc := range testCases {
		voteSet, valSet, vals := randVoteSet(height-1, round, tmproto.PrecommitType, tc.numValidators, 1)

		vi := int32(0)
		for n := range tc.blockIDs {
			for i := 0; i < tc.numVotes[n]; i++ {
				pubKey, err := vals[vi].GetPubKey()
				require.NoError(t, err)
				vote := &Vote{
					ValidatorAddress: pubKey.Address(),
					ValidatorIndex:   vi,
					Height:           height - 1,
					Round:            round,
					Type:             tmproto.PrecommitType,
					BlockID:          tc.blockIDs[n],
					Timestamp:        tmtime.Now(),
				}

				added, err := signAddVote(vals[vi], vote, voteSet)
				assert.NoError(t, err)
				assert.True(t, added)

				vi++
			}
		}

		if tc.valid {
			commit := voteSet.MakeCommit() // panics without > 2/3 valid votes
			assert.NotNil(t, commit)
			err := valSet.VerifyCommit(voteSet.ChainID(), blockID, height-1, commit)
			assert.Nil(t, err)
		} else {
			assert.Panics(t, func() { voteSet.MakeCommit() })
		}
	}
}

func TestBlockIDValidateBasic(t *testing.T) {
	validBlockID := BlockID{
		Hash: bytes.HexBytes{},
		PartSetHeader: PartSetHeader{
			Total: 1,
			Hash:  bytes.HexBytes{},
		},
	}

	invalidBlockID := BlockID{
		Hash: []byte{0},
		PartSetHeader: PartSetHeader{
			Total: 1,
			Hash:  []byte{0},
		},
	}

	testCases := []struct {
		testName             string
		blockIDHash          bytes.HexBytes
		blockIDPartSetHeader PartSetHeader
		expectErr            bool
	}{
		{"Valid BlockID", validBlockID.Hash, validBlockID.PartSetHeader, false},
		{"Invalid BlockID", invalidBlockID.Hash, validBlockID.PartSetHeader, true},
		{"Invalid BlockID", validBlockID.Hash, invalidBlockID.PartSetHeader, true},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			blockID := BlockID{
				Hash:          tc.blockIDHash,
				PartSetHeader: tc.blockIDPartSetHeader,
			}
			assert.Equal(t, tc.expectErr, blockID.ValidateBasic() != nil, "Validate Basic had an unexpected result")
		})
	}
}

func TestBlockProtoBuf(t *testing.T) {
	h := tmrand.Int63()
	c1 := randCommit(time.Now())
	b1 := MakeBlock(h, makeData([]Tx{Tx([]byte{1})}, []Evidence{}, nil), &Commit{Signatures: []CommitSig{}})
	b1.ProposerAddress = tmrand.Bytes(crypto.AddressSize)

	evidenceTime := time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
	evi := NewMockDuplicateVoteEvidence(h, evidenceTime, "block-test-chain")
	b2 := MakeBlock(h, makeData([]Tx{Tx([]byte{1})}, []Evidence{evi}, nil), c1)
	b2.ProposerAddress = tmrand.Bytes(crypto.AddressSize)
	b2.Data.Evidence.ByteSize()

	b3 := MakeBlock(h, makeData([]Tx{}, []Evidence{}, nil), c1)
	b3.ProposerAddress = tmrand.Bytes(crypto.AddressSize)
	testCases := []struct {
		msg      string
		b1       *Block
		expPass  bool
		expPass2 bool
	}{
		{"nil block", nil, false, false},
		{"b1", b1, true, true},
		{"b2", b2, true, true},
		{"b3", b3, true, true},
	}
	for _, tc := range testCases {
		pb, err := tc.b1.ToProto()
		if tc.expPass {
			require.NoError(t, err, tc.msg)
		} else {
			require.Error(t, err, tc.msg)
		}

		block, err := BlockFromProto(pb)
		if tc.expPass2 {
			require.NoError(t, err, tc.msg)
			require.EqualValues(t, tc.b1.Header, block.Header, tc.msg)
			require.EqualValues(t, tc.b1.Data, block.Data, tc.msg) // todo
			require.EqualValues(t, tc.b1.Evidence.Evidence, block.Evidence.Evidence, tc.msg)
			require.EqualValues(t, *tc.b1.LastCommit, *block.LastCommit, tc.msg)
		} else {
			require.Error(t, err, tc.msg)
		}
	}
}

func TestBlockDataProtobuf(t *testing.T) {
	type test struct {
		name  string
		txs   Txs
		evd   EvidenceData
		blobs []Blob
	}
	tests := []test{
		{
			name: "only txs", txs: Txs([]Tx{stdbytes.Repeat([]byte{1}, 200)}),
		},
		{
			name: "everything",
			txs:  Txs([]Tx{stdbytes.Repeat([]byte{1}, 200)}),
			evd:  EvidenceData{Evidence: EvidenceList([]Evidence{})},
			blobs: []Blob{
				{
					NamespaceID: []byte{8, 7, 6, 5, 4, 3, 2, 1},
					Data:        stdbytes.Repeat([]byte{3, 2, 1, 0}, 100),
				},
				{
					NamespaceID: []byte{1, 2, 3, 4, 5, 6, 7, 8},
					Data:        stdbytes.Repeat([]byte{1, 2, 3}, 100),
				},
			},
		},
	}

	for _, tt := range tests {
		d := Data{Txs: tt.txs, Evidence: tt.evd, Blobs: tt.blobs}
		firstHash := d.Hash()
		pd := d.ToProto()
		d2, err := DataFromProto(&pd)
		require.NoError(t, err)
		secondHash := d2.Hash()
		assert.Equal(t, firstHash, secondHash, tt.name)
	}
}

// TestEvidenceDataProtoBuf ensures parity in converting to and from proto.
func TestEvidenceDataProtoBuf(t *testing.T) {
	const chainID = "mychain"
	ev := NewMockDuplicateVoteEvidence(math.MaxInt64, time.Now(), chainID)
	data := &EvidenceData{Evidence: EvidenceList{ev}}
	_ = data.ByteSize()
	testCases := []struct {
		msg      string
		data1    *EvidenceData
		expPass1 bool
		expPass2 bool
	}{
		{"success", data, true, true},
		{"empty evidenceData", &EvidenceData{Evidence: EvidenceList{}}, true, true},
		{"fail nil Data", nil, false, false},
	}

	for _, tc := range testCases {
		protoData, err := tc.data1.ToProto()
		if tc.expPass1 {
			require.NoError(t, err, tc.msg)
		} else {
			require.Error(t, err, tc.msg)
		}

		eviD := new(EvidenceData)
		err = eviD.FromProto(protoData)
		if tc.expPass2 {
			require.NoError(t, err, tc.msg)
			require.Equal(t, tc.data1, eviD, tc.msg)
		} else {
			require.Error(t, err, tc.msg)
		}
	}
}

func makeRandHeader() Header {
	chainID := "test"
	t := time.Now()
	height := tmrand.Int63()
	randBytes := tmrand.Bytes(tmhash.Size)
	randAddress := tmrand.Bytes(crypto.AddressSize)
	h := Header{
		Version:            tmversion.Consensus{Block: version.BlockProtocol, App: 1},
		ChainID:            chainID,
		Height:             height,
		Time:               t,
		LastBlockID:        BlockID{},
		LastCommitHash:     randBytes,
		DataHash:           randBytes,
		ValidatorsHash:     randBytes,
		NextValidatorsHash: randBytes,
		ConsensusHash:      randBytes,
		AppHash:            randBytes,

		LastResultsHash: randBytes,

		EvidenceHash:    randBytes,
		ProposerAddress: randAddress,
	}

	return h
}

func TestHeaderProto(t *testing.T) {
	h1 := makeRandHeader()
	tc := []struct {
		msg     string
		h1      *Header
		expPass bool
	}{
		{"success", &h1, true},
		{"failure empty Header", &Header{}, false},
	}

	for _, tt := range tc {
		tt := tt
		t.Run(tt.msg, func(t *testing.T) {
			pb := tt.h1.ToProto()
			h, err := HeaderFromProto(pb)
			if tt.expPass {
				require.NoError(t, err, tt.msg)
				require.Equal(t, tt.h1, &h, tt.msg)
			} else {
				require.Error(t, err, tt.msg)
			}

		})
	}
}

func TestBlockIDProtoBuf(t *testing.T) {
	blockID := makeBlockID([]byte("hash"), 2, []byte("part_set_hash"))
	testCases := []struct {
		msg     string
		bid1    *BlockID
		expPass bool
	}{
		{"success", &blockID, true},
		{"success empty", &BlockID{}, true},
		{"failure BlockID nil", nil, false},
	}
	for _, tc := range testCases {
		protoBlockID := tc.bid1.ToProto()

		bi, err := BlockIDFromProto(&protoBlockID)
		if tc.expPass {
			require.NoError(t, err)
			require.Equal(t, tc.bid1, bi, tc.msg)
		} else {
			require.NotEqual(t, tc.bid1, bi, tc.msg)
		}
	}
}

func TestSignedHeaderProtoBuf(t *testing.T) {
	commit := randCommit(time.Now())
	h := makeRandHeader()

	sh := SignedHeader{Header: &h, Commit: commit}

	testCases := []struct {
		msg     string
		sh1     *SignedHeader
		expPass bool
	}{
		{"empty SignedHeader 2", &SignedHeader{}, true},
		{"success", &sh, true},
		{"failure nil", nil, false},
	}
	for _, tc := range testCases {
		protoSignedHeader := tc.sh1.ToProto()

		sh, err := SignedHeaderFromProto(protoSignedHeader)

		if tc.expPass {
			require.NoError(t, err, tc.msg)
			require.Equal(t, tc.sh1, sh, tc.msg)
		} else {
			require.Error(t, err, tc.msg)
		}
	}
}

func TestBlockIDEquals(t *testing.T) {
	var (
		blockID          = makeBlockID([]byte("hash"), 2, []byte("part_set_hash"))
		blockIDDuplicate = makeBlockID([]byte("hash"), 2, []byte("part_set_hash"))
		blockIDDifferent = makeBlockID([]byte("different_hash"), 2, []byte("part_set_hash"))
		blockIDEmpty     = BlockID{}
	)

	assert.True(t, blockID.Equals(blockIDDuplicate))
	assert.False(t, blockID.Equals(blockIDDifferent))
	assert.False(t, blockID.Equals(blockIDEmpty))
	assert.True(t, blockIDEmpty.Equals(blockIDEmpty))
	assert.False(t, blockIDEmpty.Equals(blockIDDifferent))
}

func TestMessagesIsSorted(t *testing.T) {
	sortedBlobs := []Blob{
		{
			NamespaceID: []byte{1, 2, 3, 4, 5, 6, 7, 8},
			Data:        stdbytes.Repeat([]byte{1}, 100),
		},
		{
			NamespaceID: []byte{8, 7, 6, 5, 4, 3, 2, 1},
			Data:        stdbytes.Repeat([]byte{2}, 100),
		},
	}
	sameNamespacedBlobs := []Blob{
		{
			NamespaceID: []byte{1, 2, 3, 4, 5, 6, 7, 8},
			Data:        stdbytes.Repeat([]byte{1}, 100),
		},
		{
			NamespaceID: []byte{1, 2, 3, 4, 5, 6, 7, 8},
			Data:        stdbytes.Repeat([]byte{2}, 100),
		},
	}
	unsortedBlobs := []Blob{
		{
			NamespaceID: []byte{8, 7, 6, 5, 4, 3, 2, 1},
			Data:        stdbytes.Repeat([]byte{1}, 100),
		},
		{
			NamespaceID: []byte{1, 2, 3, 4, 5, 6, 7, 8},
			Data:        stdbytes.Repeat([]byte{2}, 100),
		},
	}

	type testCase struct {
		descripton string
		blobs      []Blob
		want       bool
	}

	tests := []testCase{
		{"sorted blobs", sortedBlobs, true},
		{"same namespace blobs", sameNamespacedBlobs, true},
		{"unsorted blobs", unsortedBlobs, false},
	}

	for _, tc := range tests {
		t.Run(tc.descripton, func(t *testing.T) {
			bs := tc.blobs
			assert.Equal(t, tc.want, sort.IsSorted(BlobsByNamespace(bs)))
		})
	}
}

// TestDataProto tests DataFromProto and Data.ToProto
func TestDataProto(t *testing.T) {
	type testCase struct {
		name    string
		proto   *tmproto.Data
		data    Data
		wantErr bool
	}
	testCases := []testCase{
		{
			name:    "nil proto",
			proto:   nil,
			data:    Data{},
			wantErr: true,
		},
		{
			name: "empty data",
			proto: &tmproto.Data{
				Txs:        [][]uint8(nil),
				Evidence:   tmproto.EvidenceList{Evidence: []tmproto.Evidence{}},
				Blobs:      []tmproto.Blob{},
				SquareSize: 0x0,
				Hash:       []uint8(nil),
			},
			data: Data{
				Txs:        []Tx{},
				Evidence:   EvidenceData{Evidence: EvidenceList{}},
				Blobs:      []Blob{},
				SquareSize: 0,
			},
		},
		{
			name: "one blob",
			proto: &tmproto.Data{
				Txs:      [][]uint8(nil),
				Evidence: tmproto.EvidenceList{Evidence: []tmproto.Evidence{}},
				Blobs: []tmproto.Blob{
					{
						NamespaceId:  []uint8{1, 2, 3, 4, 5, 6, 7, 8},
						Data:         []uint8{1},
						ShareVersion: 0x0,
					},
				},
				SquareSize: 0x0,
				Hash:       []uint8(nil),
			},
			data: Data{
				Txs:      []Tx{},
				Evidence: EvidenceData{Evidence: EvidenceList{}},
				Blobs: []Blob{
					{
						NamespaceID:  []byte{1, 2, 3, 4, 5, 6, 7, 8},
						Data:         []byte{1},
						ShareVersion: 0,
					},
				},
				SquareSize: 0,
			},
		},
		{
			name: "two blobs",
			proto: &tmproto.Data{
				Txs:      [][]uint8(nil),
				Evidence: tmproto.EvidenceList{Evidence: []tmproto.Evidence{}},
				Blobs: []tmproto.Blob{
					{
						NamespaceId:  []uint8{1, 1, 1, 1, 1, 1, 1, 1},
						Data:         []uint8{1},
						ShareVersion: 0x1,
					},
					{
						NamespaceId:  []uint8{2, 2, 2, 2, 2, 2, 2, 2},
						Data:         []uint8{2},
						ShareVersion: 0x2,
					},
				},
				SquareSize: 0x0,
				Hash:       []uint8(nil),
			},
			data: Data{
				Txs:      []Tx{},
				Evidence: EvidenceData{Evidence: EvidenceList{}},
				Blobs: []Blob{
					{
						NamespaceID:  []byte{1, 1, 1, 1, 1, 1, 1, 1},
						Data:         []byte{1},
						ShareVersion: 1,
					},
					{
						NamespaceID:  []byte{2, 2, 2, 2, 2, 2, 2, 2},
						Data:         []byte{2},
						ShareVersion: 2,
					},
				},
				SquareSize: 0,
			},
		},
		{
			name: "one blob with too large of a share version",
			proto: &tmproto.Data{
				Txs:      [][]uint8(nil),
				Evidence: tmproto.EvidenceList{Evidence: []tmproto.Evidence{}},
				Blobs: []tmproto.Blob{
					{
						NamespaceId:  []uint8{1, 2, 3, 4, 5, 6, 7, 8},
						Data:         []uint8{1},
						ShareVersion: 257, // does not fit in a uint8
					},
				},
				SquareSize: 0x0,
				Hash:       []uint8(nil),
			},
			data:    Data{},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := DataFromProto(tc.proto)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}
			assert.Equal(t, tc.data, got)

			proto := tc.data.ToProto()
			assert.Equal(t, tc.proto, &proto)
		})
	}
}
