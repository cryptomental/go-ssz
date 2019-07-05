// Code generated by yaml_to_go. DO NOT EDIT.
// source: ssz_minimal_one.yaml

package autogenerated

type MinimalFork struct {
	PreviousVersion []byte `json:"previous_version" ssz-size:"4"`
	CurrentVersion  []byte `json:"current_version" ssz-size:"4"`
	Epoch           uint64 `json:"epoch"`
}

type MinimalCheckpoint struct {
	Epoch uint64 `json:"epoch"`
	Root  []byte `json:"root" ssz-size:"32"`
}

type MinimalValidator struct {
	Pubkey                     []byte `json:"pubkey" ssz-size:"48"`
	WithdrawalCredentials      []byte `json:"withdrawal_credentials" ssz-size:"32"`
	EffectiveBalance           uint64 `json:"effective_balance"`
	Slashed                    bool   `json:"slashed"`
	ActivationEligibilityEpoch uint64 `json:"activation_eligibility_epoch"`
	ActivationEpoch            uint64 `json:"activation_epoch"`
	ExitEpoch                  uint64 `json:"exit_epoch"`
	WithdrawableEpoch          uint64 `json:"withdrawable_epoch"`
}

type MinimalCrosslink struct {
	Shard      uint64 `json:"shard"`
	ParentRoot []byte `json:"parent_root" ssz-size:"32"`
	StartEpoch uint64 `json:"start_epoch"`
	EndEpoch   uint64 `json:"end_epoch"`
	DataRoot   []byte `json:"data_root" ssz-size:"32"`
}

type MinimalAttestationData struct {
	BeaconBlockRoot []byte            `json:"beacon_block_root" ssz-size:"32"`
	Source          MinimalCheckpoint `json:"source"`
	Target          MinimalCheckpoint `json:"target"`
	Crosslink       MinimalCrosslink  `json:"crosslink"`
}

type MinimalAttestationAndCustodyBit struct {
	Data       MinimalAttestationData `json:"data"`
	CustodyBit bool                   `json:"custody_bit"`
}

type MinimalIndexedAttestation struct {
	CustodyBit0Indices []uint64               `json:"custody_bit_0_indices" ssz-max:"4096"`
	CustodyBit1Indices []uint64               `json:"custody_bit_1_indices" ssz-max:"4096"`
	Data               MinimalAttestationData `json:"data"`
	Signature          []byte                 `json:"signature" ssz-size:"96"`
}

type MinimalPendingAttestation struct {
	AggregationBits []byte                 `json:"aggregation_bits" ssz-max:"16" ssz-kind:"bitlist"`
	Data            MinimalAttestationData `json:"data"`
	InclusionDelay  uint64                 `json:"inclusion_delay"`
	ProposerIndex   uint64                 `json:"proposer_index"`
}

type MinimalEth1Data struct {
	DepositRoot  []byte `json:"deposit_root" ssz-size:"32"`
	DepositCount uint64 `json:"deposit_count"`
	BlockHash    []byte `json:"block_hash" ssz-size:"32"`
}

type MinimalHistoricalBatch struct {
	BlockRoots [][]byte `json:"block_roots" ssz-size:"64,32"`
	StateRoots [][]byte `json:"state_roots" ssz-size:"64,32"`
}

type MinimalDepositData struct {
	Pubkey                []byte `json:"pubkey" ssz-size:"48"`
	WithdrawalCredentials []byte `json:"withdrawal_credentials" ssz-size:"32"`
	Amount                uint64 `json:"amount"`
	Signature             []byte `json:"signature" ssz-size:"96"`
}

type MinimalCompactCommittee struct {
	Pubkeys           [][]byte `json:"pubkeys" ssz-size:"?,48" ssz-max:"4096"`
	CompactValidators []uint64 `json:"compact_validators" ssz-max:"4096"`
}

type MinimalBlockHeader struct {
	Slot       uint64 `json:"slot"`
	ParentRoot []byte `json:"parent_root" ssz-size:"32"`
	StateRoot  []byte `json:"state_root" ssz-size:"32"`
	BodyRoot   []byte `json:"body_root" ssz-size:"32"`
	Signature  []byte `json:"signature" ssz-size:"96"`
}

type MinimalProposerSlashing struct {
	ProposerIndex uint64             `json:"proposer_index"`
	Header1       MinimalBlockHeader `json:"header_1"`
	Header2       MinimalBlockHeader `json:"header_2"`
}

type MinimalAttesterSlashing struct {
	Attestation1 MinimalIndexedAttestation `json:"attestation_1"`
	Attestation2 MinimalIndexedAttestation `json:"attestation_2"`
}

type MinimalAttestation struct {
	AggregationBits []byte                 `json:"aggregation_bits" ssz-max:"16" ssz-kind:"bitlist"`
	Data            MinimalAttestationData `json:"data"`
	CustodyBits     []byte                 `json:"custody_bits" ssz-max:"16" ssz-kind:"bitlist"`
	Signature       []byte                 `json:"signature" ssz-size:"96"`
}

type MinimalDeposit struct {
	Proof [][]byte           `json:"proof" ssz-size:"33,32"`
	Data  MinimalDepositData `json:"data"`
}

type MinimalVoluntaryExit struct {
	Epoch          uint64 `json:"epoch"`
	ValidatorIndex uint64 `json:"validator_index"`
	Signature      []byte `json:"signature" ssz-size:"96"`
}

type MinimalTransfer struct {
	Sender    uint64 `json:"sender"`
	Recipient uint64 `json:"recipient"`
	Amount    uint64 `json:"amount"`
	Fee       uint64 `json:"fee"`
	Slot      uint64 `json:"slot"`
	Pubkey    []byte `json:"pubkey" ssz-size:"48"`
	Signature []byte `json:"signature" ssz-size:"96"`
}

type MinimalBlockBody struct {
	RandaoReveal      []byte                    `json:"randao_reveal" ssz-size:"96"`
	Eth1Data          MinimalEth1Data           `json:"eth1_data"`
	Graffiti          []byte                    `json:"graffiti" ssz-size:"32"`
	ProposerSlashings []MinimalProposerSlashing `json:"proposer_slashings" ssz-max:"16"`
	AttesterSlashings []MinimalAttesterSlashing `json:"attester_slashings" ssz-max:"1"`
	Attestations      []MinimalAttestation      `json:"attestations" ssz-max:"128"`
	Deposits          []MinimalDeposit          `json:"deposits" ssz-max:"16"`
	VoluntaryExits    []MinimalVoluntaryExit    `json:"voluntary_exits" ssz-max:"16"`
	Transfers         []MinimalTransfer         `json:"transfers" ssz-max:"0"`
}

type MinimalBlock struct {
	Slot       uint64           `json:"slot"`
	ParentRoot []byte           `json:"parent_root" ssz-size:"32"`
	StateRoot  []byte           `json:"state_root" ssz-size:"32"`
	Body       MinimalBlockBody `json:"body"`
	Signature  []byte           `json:"signature" ssz-size:"96"`
}

type MinimalBeaconState struct {
	Slot                   uint64             `json:"slot"`
	GenesisTime            uint64             `json:"genesis_time"`
	Fork                   MinimalFork        `json:"fork"`
	LatestBlockHeader      MinimalBlockHeader `json:"latest_block_header"`
	BlockRoots             [][]byte           `json:"block_roots" ssz-size:"64,32"`
	StateRoots             [][]byte           `json:"state_roots" ssz-size:"64,32"`
	HistoricalRoots        [][]byte           `json:"historical_roots" ssz-size:"?,32" ssz-max:"16777216"`
	Eth1Data               MinimalEth1Data    `json:"eth1_data"`
	Eth1DataVotes          []MinimalEth1Data  `json:"eth1_data_votes" ssz-max:"16"`
	Eth1DepositIndex       uint64             `json:"eth1_deposit_index"`
	Validators             []MinimalValidator `json:"validators" ssz-max:"1099511627776"`
	Balances               []uint64           `json:"balances" ssz-max:"1099511627776"`
	StartShard             uint64             `json:"start_shard"`
	RandaoMixes            [][]byte           `json:"randao_mixes" ssz-size:"64,32"`
	ActiveIndexRoots       [][]byte           `json:"active_index_roots" ssz-size:"64,32"`
	CompactCommitteesRoots [][]byte           `json:"compact_committees_roots" ssz-size:"64,32"`
	Slashings              []uint64           `json:"slashings" ssz-size:"64"`

	PreviousEpochAttestations []MinimalPendingAttestation `json:"previous_epoch_attestations" ssz-max:"1024"`
	CurrentEpochAttestations  []MinimalPendingAttestation `json:"current_epoch_attestations" ssz-max:"1024"`
	PreviousCrosslinks        []MinimalCrosslink          `json:"previous_crosslinks" ssz-size:"8"`
	CurrentCrosslinks         []MinimalCrosslink          `json:"current_crosslinks" ssz-size:"8"`
	JustificationBits         []byte                      `json:"justification_bits" ssz-size:"4" ssz-kind:"bitvector"`

	PreviousJustifiedCheckpoint MinimalCheckpoint `json:"previous_justified_checkpoint"`
	CurrentJustifiedCheckpoint  MinimalCheckpoint `json:"current_justified_checkpoint"`
	FinalizedCheckpoint         MinimalCheckpoint `json:"finalized_checkpoint"`
}
type SszMinimalTest struct {
	Title         string   `json:"title"`
	Summary       string   `json:"summary"`
	ForksTimeline string   `json:"forks_timeline"`
	Forks         []string `json:"forks"`
	Config        string   `json:"config"`
	Runner        string   `json:"runner"`
	Handler       string   `json:"handler"`
	TestCases     []struct {
		Attestation struct {
			Value       MinimalAttestation `json:"value"`
			Serialized  []byte             `json:"serialized"`
			Root        []byte             `json:"root" ssz:"size=32"`
			SigningRoot []byte             `json:"signing_root" ssz:"size=32"`
		} `json:"Attestation,omitempty"`
		AttestationData struct {
			Value      MinimalAttestationData `json:"value"`
			Serialized []byte                 `json:"serialized"`
			Root       []byte                 `json:"root" ssz:"size=32"`
		} `json:"AttestationData,omitempty"`
		AttestationDataAndCustodyBit struct {
			Value      MinimalAttestationAndCustodyBit `json:"value"`
			Serialized []byte                          `json:"serialized"`
			Root       []byte                          `json:"root" ssz:"size=32"`
		} `json:"AttestationDataAndCustodyBit,omitempty"`
		AttesterSlashing struct {
			Value      MinimalAttesterSlashing `json:"value"`
			Serialized []byte                  `json:"serialized"`
			Root       []byte                  `json:"root" ssz:"size=32"`
		} `json:"AttesterSlashing,omitempty"`
		BeaconBlock struct {
			Value       MinimalBlock `json:"value"`
			Serialized  []byte       `json:"serialized"`
			Root        []byte       `json:"root" ssz:"size=32"`
			SigningRoot []byte       `json:"signing_root" ssz:"size=32"`
		} `json:"BeaconBlock,omitempty"`
		BeaconBlockBody struct {
			Value      MinimalBlockBody `json:"value"`
			Serialized []byte           `json:"serialized"`
			Root       []byte           `json:"root" ssz:"size=32"`
		} `json:"BeaconBlockBody,omitempty"`
		BeaconBlockHeader struct {
			Value       MinimalBlockHeader `json:"value"`
			Serialized  []byte             `json:"serialized"`
			Root        []byte             `json:"root" ssz:"size=32"`
			SigningRoot []byte             `json:"signing_root" ssz:"size=32"`
		} `json:"BeaconBlockHeader,omitempty"`
		BeaconState struct {
			Value      MinimalBeaconState `json:"value"`
			Serialized []byte             `json:"serialized"`
			Root       []byte             `json:"root" ssz:"size=32"`
		} `json:"BeaconState,omitempty"`
		Checkpoint struct {
			Value      MinimalCheckpoint `json:"value"`
			Serialized []byte            `json:"serialized"`
			Root       []byte            `json:"root" ssz:"size=32"`
		} `json:"Checkpoint,omitempty"`
		CompactCommittee struct {
			Value      MinimalCompactCommittee `json:"value"`
			Serialized []byte                  `json:"serialized"`
			Root       []byte                  `json:"root" ssz:"size=32"`
		} `json:"CompactCommittee,omitempty"`
		Crosslink struct {
			Value      MinimalCrosslink `json:"value"`
			Serialized []byte           `json:"serialized"`
			Root       []byte           `json:"root" ssz:"size=32"`
		} `json:"Crosslink,omitempty"`
		Deposit struct {
			Value      MinimalDeposit `json:"value"`
			Serialized []byte         `json:"serialized"`
			Root       []byte         `json:"root" ssz:"size=32"`
		} `json:"Deposit,omitempty"`
		DepositData struct {
			Value       MinimalDepositData `json:"value"`
			Serialized  []byte             `json:"serialized"`
			Root        []byte             `json:"root" ssz:"size=32"`
			SigningRoot []byte             `json:"signing_root" ssz:"size=32"`
		} `json:"DepositData,omitempty"`
		Eth1Data struct {
			Value      MinimalEth1Data `json:"value"`
			Serialized []byte          `json:"serialized"`
			Root       []byte          `json:"root" ssz:"size=32"`
		} `json:"Eth1Data,omitempty"`
		Fork struct {
			Value      MinimalFork `json:"value"`
			Serialized []byte      `json:"serialized"`
			Root       []byte      `json:"root" ssz:"size=32"`
		} `json:"Fork,omitempty"`
		HistoricalBatch struct {
			Value      MinimalHistoricalBatch `json:"value"`
			Serialized []byte                 `json:"serialized"`
			Root       []byte                 `json:"root" ssz:"size=32"`
		} `json:"HistoricalBatch,omitempty"`
		IndexedAttestation struct {
			Value       MinimalIndexedAttestation `json:"value"`
			Serialized  []byte                    `json:"serialized"`
			Root        []byte                    `json:"root" ssz:"size=32"`
			SigningRoot []byte                    `json:"signing_root" ssz:"size=32"`
		} `json:"IndexedAttestation,omitempty"`
		PendingAttestation struct {
			Value      MinimalPendingAttestation `json:"value"`
			Serialized []byte                    `json:"serialized"`
			Root       []byte                    `json:"root" ssz:"size=32"`
		} `json:"PendingAttestation,omitempty"`
		ProposerSlashing struct {
			Value      MinimalProposerSlashing `json:"value"`
			Serialized []byte                  `json:"serialized"`
			Root       []byte                  `json:"root" ssz:"size=32"`
		} `json:"ProposerSlashing,omitempty"`
		Transfer struct {
			Value       MinimalTransfer `json:"value"`
			Serialized  []byte          `json:"serialized"`
			Root        []byte          `json:"root" ssz:"size=32"`
			SigningRoot []byte          `json:"signing_root" ssz:"size=32"`
		} `json:"Transfer,omitempty"`
		Validator struct {
			Value      MinimalValidator `json:"value"`
			Serialized []byte           `json:"serialized"`
			Root       []byte           `json:"root" ssz:"size=32"`
		} `json:"Validator,omitempty"`
		VoluntaryExit struct {
			Value       MinimalVoluntaryExit `json:"value"`
			Serialized  []byte               `json:"serialized"`
			Root        []byte               `json:"root" ssz:"size=32"`
			SigningRoot []byte               `json:"signing_root" ssz:"size=32"`
		} `json:"VoluntaryExit,omitempty"`
	} `json:"test_cases"`
}
