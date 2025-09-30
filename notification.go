package goesi

type AcceptedAlly struct {
	AllyID   int64   `yaml:"allyID"`
	CharID   int64   `yaml:"charID"`
	EnemyID  int64   `yaml:"enemyID"`
	IskValue float64 `yaml:"iskValue"`
	Time     int64   `yaml:"time"`
}

type AcceptedSurrender struct {
	CharID     int64   `yaml:"charID"`
	EntityID   int64   `yaml:"entityID"`
	IskValue   float64 `yaml:"iskValue"`
	OfferingID int64   `yaml:"offeringID"`
}

type AllMaintenanceBillMsg struct {
	AllianceID int64 `yaml:"allianceID"`
	DueDate    int64 `yaml:"dueDate"`
}

type AllWarCorpJoinedAllianceMsg struct {
	AllianceID int64 `yaml:"allianceID"`
	CorpID     int64 `yaml:"corpID"`
}

type AllWarDeclaredMsg struct {
	AgainstID    int64   `yaml:"againstID"`
	Cost         float64 `yaml:"cost"`
	DeclaredByID int64   `yaml:"declaredByID"`
	DelayHours   int32   `yaml:"delayHours"`
	HostileState int32   `yaml:"hostileState"`
}

type AllWarInvalidatedMsg struct {
	AgainstID    int64   `yaml:"againstID"`
	Cost         float64 `yaml:"cost"`
	DeclaredByID int64   `yaml:"declaredByID"`
	DelayHours   int32   `yaml:"delayHours"`
	HostileState int32   `yaml:"hostileState"`
}

type AllWarRetractedMsg struct {
	AgainstID    int64   `yaml:"againstID"`
	Cost         float64 `yaml:"cost"`
	DeclaredByID int64   `yaml:"declaredByID"`
	DelayHours   int32   `yaml:"delayHours"`
	HostileState int32   `yaml:"hostileState"`
}

type AllWarSurrenderMsg struct {
	AgainstID    int64   `yaml:"againstID"`
	Cost         float64 `yaml:"cost"`
	DeclaredByID int64   `yaml:"declaredByID"`
	DelayHours   int32   `yaml:"delayHours"`
	HostileState int32   `yaml:"hostileState"`
}

type AllianceCapitalChanged struct {
	AllianceID    int64 `yaml:"allianceID"`
	SolarSystemID int64 `yaml:"solarSystemID"`
}

type AllyContractCancelled struct {
	AggressorID  int64 `yaml:"aggressorID"`
	DefenderID   int64 `yaml:"defenderID"`
	TimeFinished int64 `yaml:"timeFinished"`
}

type AllyJoinedWarAggressorMsg struct {
	AllyID     int64 `yaml:"allyID"`
	DefenderID int64 `yaml:"defenderID"`
	StartTime  int64 `yaml:"startTime"`
}

type AllyJoinedWarAllyMsg struct {
	AggressorID int64 `yaml:"aggressorID"`
	AllyID      int64 `yaml:"allyID"`
	DefenderID  int64 `yaml:"defenderID"`
	StartTime   int64 `yaml:"startTime"`
}

type AllyJoinedWarDefenderMsg struct {
	AggressorID int64 `yaml:"aggressorID"`
	AllyID      int64 `yaml:"allyID"`
	StartTime   int64 `yaml:"startTime"`
}

type BillOutOfMoneyMsg struct {
	BillTypeID int64 `yaml:"billTypeID"`
	DueDate    int64 `yaml:"dueDate"`
}

type BillPaidCorpAllMsg struct {
	Amount  int32 `yaml:"amount"`
	DueDate int64 `yaml:"dueDate"`
}

type BountyClaimMsg struct {
	Amount float64 `yaml:"amount"`
	CharID int64   `yaml:"charID"`
}

type BountyESSShared struct {
	CharID   int64   `yaml:"charID"`
	MyIsk    float64 `yaml:"myIsk"`
	TotalIsk float64 `yaml:"totalIsk"`
}

type BountyESSTaken struct {
	CharID   int64   `yaml:"charID"`
	MyIsk    float64 `yaml:"myIsk"`
	TotalIsk float64 `yaml:"totalIsk"`
}

type BountyPlacedAlliance struct {
	Bounty         float64 `yaml:"bounty"`
	BountyPlacerID int64   `yaml:"bountyPlacerID"`
}

type BountyPlacedChar struct {
	Bounty         float64 `yaml:"bounty"`
	BountyPlacerID int64   `yaml:"bountyPlacerID"`
}

type BountyPlacedCorp struct {
	Bounty         float64 `yaml:"bounty"`
	BountyPlacerID int64   `yaml:"bountyPlacerID"`
}

type BountyYourBountyClaimed struct {
	Bounty   float64 `yaml:"bounty"`
	VictimID int64   `yaml:"victimID"`
}

type BuddyConnectContactAdd struct {
	Level   int32  `yaml:"level"`
	Message string `yaml:"message"`
}

type CharAppAcceptMsg struct {
	ApplicationText string `yaml:"applicationText"`
	CharID          int64  `yaml:"charID"`
	CorpID          int64  `yaml:"corpID"`
}

type CharAppRejectMsg struct {
	ApplicationText string `yaml:"applicationText"`
	CharID          int64  `yaml:"charID"`
	CorpID          int64  `yaml:"corpID"`
}

type CharAppWithdrawMsg struct {
	ApplicationText string `yaml:"applicationText"`
	CharID          int64  `yaml:"charID"`
	CorpID          int64  `yaml:"corpID"`
}

type CharLeftCorpMsg struct {
	CharID int64 `yaml:"charID"`
	CorpID int64 `yaml:"corpID"`
}

type CharMedalMsg struct {
	CorpID  int64  `yaml:"corpID"`
	MedalID int64  `yaml:"medalID"`
	Reason  string `yaml:"reason"`
}

type CharTerminationMsg struct {
	CharID      int64   `yaml:"charID"`
	CorpID      int64   `yaml:"corpID"`
	RoleName    string  `yaml:"roleName"`
	RoleNameIDs []int64 `yaml:"roleNameIDs"`
	Security    float64 `yaml:"security"`
}

type CloneActivationMsg struct {
	CloneBought     int32 `yaml:"cloneBought"`
	CloneStationID  int64 `yaml:"cloneStationID"`
	CloneTypeID     int64 `yaml:"cloneTypeID"`
	CorpStationID   int64 `yaml:"corpStationID"`
	LastCloned      int64 `yaml:"lastCloned"`
	PodKillerID     int64 `yaml:"podKillerID"`
	SkillID         int64 `yaml:"skillID"`
	SkillPointsLost int32 `yaml:"skillPointsLost"`
}

type CloneActivationMsg2 struct {
	CloneStationID int64 `yaml:"cloneStationID"`
	CorpStationID  int64 `yaml:"corpStationID"`
	LastCloned     int64 `yaml:"lastCloned"`
	PodKillerID    int64 `yaml:"podKillerID"`
}

type CloneMovedMsg struct {
	CharsInCorpID int64 `yaml:"charsInCorpID"`
	CorpID        int64 `yaml:"corpID"`
	NewStationID  int64 `yaml:"newStationID"`
	StationID     int64 `yaml:"stationID"`
}

type CloneRevokedMsg2 struct {
	CorpID       int64 `yaml:"corpID"`
	NewStationID int64 `yaml:"newStationID"`
	StationID    int64 `yaml:"stationID"`
}

type ContactAdd struct {
	Level   int32  `yaml:"level"`
	Message string `yaml:"message"`
}

type ContactEdit struct {
	Level   float64 `yaml:"level"`
	Message string  `yaml:"message"`
}

type ContainerPasswordMsg struct {
	CharID        int64  `yaml:"charID"`
	Password      string `yaml:"password"`
	PasswordType  string `yaml:"passwordType"`
	SolarSystemID int64  `yaml:"solarSystemID"`
	StationID     int64  `yaml:"stationID"`
	TypeID        int64  `yaml:"typeID"`
}

type CorpAllBillMsg struct {
	Amount      float64 `yaml:"amount"`
	BillTypeID  int64   `yaml:"billTypeID"`
	CreditorID  int64   `yaml:"creditorID"`
	CurrentDate int64   `yaml:"currentDate"`
	DebtorID    int64   `yaml:"debtorID"`
	DueDate     int64   `yaml:"dueDate"`
	ExternalID  int64   `yaml:"externalID"`
	ExternalID2 int64   `yaml:"externalID2"`
}

// CorpAllBillMsgV2 represents the updated version for this notification type.
// Use [CorpAllBillMsg] to unmarshal older notifications.
type CorpAllBillMsgV2 struct {
	Amount      float64 `yaml:"amount"`
	BillTypeID  int64   `yaml:"billTypeID"`
	CreditorID  int64   `yaml:"creditorID"`
	CurrentDate int64   `yaml:"currentDate"`
	DebtorID    int64   `yaml:"debtorID"`
	DueDate     int64   `yaml:"dueDate"`
	ExternalID  int64   `yaml:"externalID"`
	ExternalID2 int64   `yaml:"externalID2"`
}

type CorpAppAcceptMsg struct {
	ApplicationText string `yaml:"applicationText"`
	CharID          int64  `yaml:"charID"`
	CorpID          int64  `yaml:"corpID"`
}

type CorpAppInvitedMsg struct {
	ApplicationText string `yaml:"applicationText"`
	CharID          int64  `yaml:"charID"`
	CorpID          int64  `yaml:"corpID"`
	InvokingCharID  int64  `yaml:"invokingCharID"`
}

type CorpAppNewMsg struct {
	ApplicationText string `yaml:"applicationText"`
	CharID          int64  `yaml:"charID"`
	CorpID          int64  `yaml:"corpID"`
}

type CorpAppRejectCustomMsg struct {
	ApplicationText string `yaml:"applicationText"`
	CharID          int64  `yaml:"charID"`
	CorpID          int64  `yaml:"corpID"`
	CustomMessage   string `yaml:"customMessage"`
}

type CorpAppRejectMsg struct {
	ApplicationText string `yaml:"applicationText"`
	CharID          int64  `yaml:"charID"`
	CorpID          int64  `yaml:"corpID"`
}

type CorpDividendMsg struct {
	Amount    float64 `yaml:"amount"`
	CorpID    int64   `yaml:"corpID"`
	IsMembers bool    `yaml:"isMembers"`
	Payout    float64 `yaml:"payout"`
}

type CorpFriendlyFireDisableTimerCompleted struct {
	CorpID int64 `yaml:"corpID"`
}

type CorpFriendlyFireDisableTimerStarted struct {
	CharID       int64 `yaml:"charID"`
	CorpID       int64 `yaml:"corpID"`
	TimeFinished int64 `yaml:"timeFinished"`
}

type CorpFriendlyFireEnableTimerCompleted struct {
	CorpID int64 `yaml:"corpID"`
}

type CorpFriendlyFireEnableTimerStarted struct {
	CharID       int64 `yaml:"charID"`
	CorpID       int64 `yaml:"corpID"`
	TimeFinished int64 `yaml:"timeFinished"`
}

type CorpKicked struct {
	CorpID int64 `yaml:"corpID"`
}

type CorpLiquidationMsg struct {
	Amount float64 `yaml:"amount"`
	CorpID int64   `yaml:"corpID"`
	Payout float64 `yaml:"payout"`
}

type CorpNewCEOMsg struct {
	CorpID   int64 `yaml:"corpID"`
	NewCeoID int64 `yaml:"newCeoID"`
	OldCeoID int64 `yaml:"oldCeoID"`
}

type CorpNewsMsg struct {
	Body      string `yaml:"body"`
	CorpID    int64  `yaml:"corpID"`
	InEffect  int32  `yaml:"inEffect"`
	Parameter int32  `yaml:"parameter"`
	VoteType  int32  `yaml:"voteType"`
}

type CorpTaxChangeMsg struct {
	CorpID     int64   `yaml:"corpID"`
	NewTaxRate float64 `yaml:"newTaxRate"`
	OldTaxRate float64 `yaml:"oldTaxRate"`
}

type CorpVoteMsg struct {
	Body    string `yaml:"body"`
	Subject string `yaml:"subject"`
}

type CorpWarDeclaredMsg struct {
	AgainstID    int64   `yaml:"againstID"`
	Cost         float64 `yaml:"cost"`
	DeclaredByID int64   `yaml:"declaredByID"`
}

type CorpWarFightingLegalMsg struct {
	AgainstID    int64   `yaml:"againstID"`
	Cost         float64 `yaml:"cost"`
	DeclaredByID int64   `yaml:"declaredByID"`
}

type CorpWarInvalidatedMsg struct {
	AgainstID    int64   `yaml:"againstID"`
	Cost         float64 `yaml:"cost"`
	DeclaredByID int64   `yaml:"declaredByID"`
}

type CorpWarRetractedMsg struct {
	AgainstID    int64   `yaml:"againstID"`
	Cost         float64 `yaml:"cost"`
	DeclaredByID int64   `yaml:"declaredByID"`
}

type CorpWarSurrenderMsg struct {
	AgainstID    int64   `yaml:"againstID"`
	Cost         float64 `yaml:"cost"`
	DeclaredByID int64   `yaml:"declaredByID"`
}

type CustomsMsg struct {
	FactionID int64 `yaml:"factionID"`
	LostList  []struct {
		Fine     float64 `yaml:"fine"`
		Penalty  float64 `yaml:"penalty"`
		Quantity int32   `yaml:"quantity"`
		TypeID   int64   `yaml:"typeID"`
	} `yaml:"lostList"`
	SecurityLevel    float64 `yaml:"securityLevel"`
	ShouldAttack     int32   `yaml:"shouldAttack"`
	ShouldConfiscate int32   `yaml:"shouldConfiscate"`
	SolarSystemID    int64   `yaml:"solarSystemID"`
	StandingDivision float64 `yaml:"standingDivision"`
}

type DeclareWar struct {
	CharID     int64 `yaml:"charID"`
	DefenderID int64 `yaml:"defenderID"`
	EntityID   int64 `yaml:"entityID"`
}

type EntosisCaptureStarted struct {
	SolarSystemID   int64 `yaml:"solarSystemID"`
	StructureTypeID int64 `yaml:"structureTypeID"`
}

type FWAllianceWarningMsg struct {
	AllianceID       int64   `yaml:"allianceID"`
	CorpList         string  `yaml:"corpList"`
	FactionID        int64   `yaml:"factionID"`
	RequiredStanding float64 `yaml:"requiredStanding"`
}

type FWCharRankGainMsg struct {
	FactionID int64 `yaml:"factionID"`
	NewRank   int32 `yaml:"newRank"`
}

type FWCharRankLossMsg struct {
	FactionID int64 `yaml:"factionID"`
	NewRank   int32 `yaml:"newRank"`
}

type FWCorpJoinMsg struct {
	CorpID    int64 `yaml:"corpID"`
	FactionID int64 `yaml:"factionID"`
}

type FWCorpKickMsg struct {
	CorpID           int64   `yaml:"corpID"`
	CurrentStanding  float64 `yaml:"currentStanding"`
	FactionID        int64   `yaml:"factionID"`
	RequiredStanding float64 `yaml:"requiredStanding"`
}

type FWCorpLeaveMsg struct {
	CorpID    int64 `yaml:"corpID"`
	FactionID int64 `yaml:"factionID"`
}

type FWCorpWarningMsg struct {
	CorpID           int64   `yaml:"corpID"`
	CurrentStanding  float64 `yaml:"currentStanding"`
	FactionID        int64   `yaml:"factionID"`
	RequiredStanding float64 `yaml:"requiredStanding"`
}

type FacWarCorpJoinRequestMsg struct {
	CorpID    int64 `yaml:"corpID"`
	FactionID int64 `yaml:"factionID"`
}

type FacWarCorpJoinWithdrawMsg struct {
	CorpID    int64 `yaml:"corpID"`
	FactionID int64 `yaml:"factionID"`
}

type FacWarCorpLeaveRequestMsg struct {
	CorpID    int64 `yaml:"corpID"`
	FactionID int64 `yaml:"factionID"`
}

type FacWarCorpLeaveWithdrawMsg struct {
	CorpID    int64 `yaml:"corpID"`
	FactionID int64 `yaml:"factionID"`
}

type FacWarLPDisqualifiedEvent struct {
	Amount               int32 `yaml:"amount"`
	CharRefID            int64 `yaml:"charRefID"`
	CorpID               int64 `yaml:"corpID"`
	DisqualificationType int32 `yaml:"disqualificationType"`
	Event                int32 `yaml:"event"`
	ItemRefID            int64 `yaml:"itemRefID"`
	LocationID           int64 `yaml:"locationID"`
}

type FacWarLPDisqualifiedKill struct {
	Amount               int32 `yaml:"amount"`
	CharRefID            int64 `yaml:"charRefID"`
	CorpID               int64 `yaml:"corpID"`
	DisqualificationType int32 `yaml:"disqualificationType"`
	Event                int32 `yaml:"event"`
	ItemRefID            int64 `yaml:"itemRefID"`
	LocationID           int64 `yaml:"locationID"`
}

type FacWarLPPayoutEvent struct {
	Amount     int32 `yaml:"amount"`
	CharRefID  int64 `yaml:"charRefID"`
	CorpID     int64 `yaml:"corpID"`
	Event      int32 `yaml:"event"`
	ItemRefID  int64 `yaml:"itemRefID"`
	LocationID int64 `yaml:"locationID"`
}

type FacWarLPPayoutKill struct {
	Amount     int32 `yaml:"amount"`
	CharRefID  int64 `yaml:"charRefID"`
	CorpID     int64 `yaml:"corpID"`
	Event      int32 `yaml:"event"`
	ItemRefID  int64 `yaml:"itemRefID"`
	LocationID int64 `yaml:"locationID"`
}

type GameTimeAdded struct {
}

type GameTimeReceived struct {
	Message      string `yaml:"message"`
	OfferID      int64  `yaml:"offerID"`
	Quantity     int32  `yaml:"quantity"`
	SenderCharID int64  `yaml:"senderCharID"`
}

type GameTimeSent struct {
	ReceiverCharID int64 `yaml:"receiverCharID"`
	SenderCharID   int64 `yaml:"senderCharID"`
}

type GiftReceived struct {
	Message      string `yaml:"message"`
	OfferID      int64  `yaml:"offerID"`
	Quantity     int32  `yaml:"quantity"`
	SenderCharID int64  `yaml:"senderCharID"`
}

type IncursionCompletedMsg struct {
	EmailMessageID float64   `yaml:"emailMessageId"`
	SolarSystemID  int64     `yaml:"solarSystemID"`
	TaleID         int64     `yaml:"taleID"`
	TopTen         [][]int64 `yaml:"topTen"`
}

type InfrastructureHubBillAboutToExpire struct {
	BillID        int64 `yaml:"billID"`
	CorpID        int64 `yaml:"corpID"`
	DueDate       int64 `yaml:"dueDate"`
	SolarSystemID int64 `yaml:"solarSystemID"`
}

type IndustryTeamAuctionLost struct {
	SolarSystemID int64             `yaml:"solarSystemID"`
	SystemBids    map[int64]float64 `yaml:"systemBids"`
	TeamNameInfo  []interface{}     `yaml:"teamNameInfo"`
	TotalIsk      float64           `yaml:"totalIsk"`
	YourAmount    float64           `yaml:"yourAmount"`
}

type IHubDestroyedByBillFailure struct {
	SolarSystemID   int64 `yaml:"solarSystemID"`
	StructureTypeID int64 `yaml:"structureTypeID"`
}

type InsuranceExpirationMsg struct {
	EndDate   int64  `yaml:"endDate"`
	ShipID    int64  `yaml:"shipID"`
	ShipName  string `yaml:"shipName"`
	StartDate int64  `yaml:"startDate"`
}

type InsuranceFirstShipMsg struct {
	IsHouseWarmingGift int32 `yaml:"isHouseWarmingGift"`
	ShipTypeID         int64 `yaml:"shipTypeID"`
}

type InsuranceInvalidatedMsg struct {
	EndDate   int64 `yaml:"endDate"`
	OwnerID   int64 `yaml:"ownerID"`
	Reason    int32 `yaml:"reason"`
	ShipID    int64 `yaml:"shipID"`
	StartDate int64 `yaml:"startDate"`
	TypeID    int64 `yaml:"typeID"`
}

type InsuranceIssuedMsg struct {
	EndDate   int64   `yaml:"endDate"`
	ItemID    int64   `yaml:"itemID"`
	Level     float64 `yaml:"level"`
	NumWeeks  int32   `yaml:"numWeeks"`
	ShipName  string  `yaml:"shipName"`
	StartDate int64   `yaml:"startDate"`
	TypeID    int64   `yaml:"typeID"`
}

type InsurancePayoutMsg struct {
	Amount float64 `yaml:"amount"`
	ItemID int64   `yaml:"itemID"`
	Payout int32   `yaml:"payout"`
}

type JumpCloneDeletedMsg1 struct {
	LocationID      int64   `yaml:"locationID"`
	LocationOwnerID int64   `yaml:"locationOwnerID"`
	OwnerID         int64   `yaml:"ownerID"`
	TypeIDs         []int64 `yaml:"typeIDs"`
}

type JumpCloneDeletedMsg2 struct {
	DestroyerID     int64   `yaml:"destroyerID"`
	LocationID      int64   `yaml:"locationID"`
	LocationOwnerID int64   `yaml:"locationOwnerID"`
	OwnerID         int64   `yaml:"ownerID"`
	TypeIDs         []int64 `yaml:"typeIDs"`
}

type KillReportFinalBlow struct {
	KillMailHash     string `yaml:"killMailHash"`
	KillMailID       int64  `yaml:"killMailID"`
	VictimID         int64  `yaml:"victimID"`
	VictimShipTypeID int64  `yaml:"victimShipTypeID"`
}

type KillReportVictim struct {
	KillMailHash     string `yaml:"killMailHash"`
	KillMailID       int64  `yaml:"killMailID"`
	VictimShipTypeID int64  `yaml:"victimShipTypeID"`
}

type KillRightAvailable struct {
	CharID     int64   `yaml:"charID"`
	Price      float64 `yaml:"price"`
	ToEntityID int64   `yaml:"toEntityID"`
}

type KillRightAvailableOpen struct {
	CharID int64   `yaml:"charID"`
	Price  float64 `yaml:"price"`
}

type KillRightEarned struct {
	CharID int64 `yaml:"charID"`
}

type KillRightUnavailable struct {
	CharID     int64 `yaml:"charID"`
	ToEntityID int64 `yaml:"toEntityID"`
}

type KillRightUnavailableOpen struct {
	CharID int64 `yaml:"charID"`
}

type KillRightUsed struct {
	CharID int64 `yaml:"charID"`
}

type LocateCharMsg struct {
	AgentLocation struct {
		Region        int64 `yaml:"3"`
		Constellation int64 `yaml:"4"`
		SolarSystem   int64 `yaml:"5"`
		Station       int64 `yaml:"15"`
	} `yaml:"agentLocation"`
	CharacterID    int64 `yaml:"characterID"`
	MessageIndex   int32 `yaml:"messageIndex"`
	TargetLocation struct {
		Region        int64 `yaml:"3"`
		Constellation int64 `yaml:"4"`
		SolarSystem   int64 `yaml:"5"`
		Station       int64 `yaml:"15"`
	} `yaml:"targetLocation"`
}

type MadeWarMutual struct {
	CharID  int64 `yaml:"charID"`
	EnemyID int64 `yaml:"enemyID"`
}

type MercOfferedNegotiationMsg struct {
	AggressorID int64   `yaml:"aggressorID"`
	DefenderID  int64   `yaml:"defenderID"`
	IskValue    float64 `yaml:"iskValue"`
	MercID      int64   `yaml:"mercID"`
}

type MissionOfferExpirationMsg struct {
	Body            []string `yaml:"body"`
	Header          []string `yaml:"header"`
	MissionKeywords struct {
		ObjectiveDestinationID       int64 `yaml:"objectiveDestinationID"`
		ObjectiveDestinationSystemID int64 `yaml:"objectiveDestinationSystemID"`
		ObjectiveLocationID          int64 `yaml:"objectiveLocationID"`
		ObjectiveLocationSystemID    int64 `yaml:"objectiveLocationSystemID"`
		ObjectiveQuantity            int32 `yaml:"objectiveQuantity"`
		ObjectiveTypeID              int64 `yaml:"objectiveTypeID"`
		RewardQuantity               int32 `yaml:"rewardQuantity"`
		RewardTypeID                 int64 `yaml:"rewardTypeID"`
	} `yaml:"missionKeywords"`
}

type MoonminingAutomaticFracture struct {
	MoonID          int64             `yaml:"moonID"`
	MoonLink        string            `yaml:"moonLink"`
	OreVolumeByType map[int64]float64 `yaml:"oreVolumeByType"`
	SolarSystemID   int64             `yaml:"solarSystemID"`
	SolarSystemLink string            `yaml:"solarSystemLink"`
	StructureID     int64             `yaml:"structureID"`
	StructureLink   string            `yaml:"structureLink"`
	StructureName   string            `yaml:"structureName"`
	StructureTypeID int64             `yaml:"structureTypeID"`
}

type MoonminingExtractionCancelled struct {
	CancelledBy     int64  `yaml:"cancelledBy"`
	CancelledByLink string `yaml:"cancelledByLink"`
	MoonID          int64  `yaml:"moonID"`
	MoonLink        string `yaml:"moonLink"`
	SolarSystemID   int64  `yaml:"solarSystemID"`
	SolarSystemLink string `yaml:"solarSystemLink"`
	StructureID     int64  `yaml:"structureID"`
	StructureLink   string `yaml:"structureLink"`
	StructureName   string `yaml:"structureName"`
	StructureTypeID int64  `yaml:"structureTypeID"`
}

type MoonminingExtractionFinished struct {
	AutoTime        int64             `yaml:"autoTime"`
	MoonID          int64             `yaml:"moonID"`
	MoonLink        string            `yaml:"moonLink"`
	OreVolumeByType map[int64]float64 `yaml:"oreVolumeByType"`
	SolarSystemID   int64             `yaml:"solarSystemID"`
	SolarSystemLink string            `yaml:"solarSystemLink"`
	StructureID     int64             `yaml:"structureID"`
	StructureLink   string            `yaml:"structureLink"`
	StructureName   string            `yaml:"structureName"`
	StructureTypeID int64             `yaml:"structureTypeID"`
}

type MoonminingExtractionStarted struct {
	AutoTime        int64             `yaml:"autoTime"`
	MoonID          int64             `yaml:"moonID"`
	MoonLink        string            `yaml:"moonLink"`
	OreVolumeByType map[int64]float64 `yaml:"oreVolumeByType"`
	ReadyTime       int64             `yaml:"readyTime"`
	SolarSystemID   int64             `yaml:"solarSystemID"`
	SolarSystemLink string            `yaml:"solarSystemLink"`
	StartedBy       int64             `yaml:"startedBy"`
	StartedByLink   string            `yaml:"startedByLink"`
	StructureID     int64             `yaml:"structureID"`
	StructureLink   string            `yaml:"structureLink"`
	StructureName   string            `yaml:"structureName"`
	StructureTypeID int64             `yaml:"structureTypeID"`
}

type MoonminingLaserFired struct {
	FiredBy         int64             `yaml:"firedBy"`
	FiredByLink     string            `yaml:"firedByLink"`
	MoonID          int64             `yaml:"moonID"`
	MoonLink        string            `yaml:"moonLink"`
	OreVolumeByType map[int64]float64 `yaml:"oreVolumeByType"`
	SolarSystemID   int64             `yaml:"solarSystemID"`
	SolarSystemLink string            `yaml:"solarSystemLink"`
	StructureID     int64             `yaml:"structureID"`
	StructureLink   string            `yaml:"structureLink"`
	StructureName   string            `yaml:"structureName"`
	StructureTypeID int64             `yaml:"structureTypeID"`
}

type NPCStandingsGained [][]float64

type NPCStandingsLost [][]float64

type OfferedSurrender struct {
	CharID    int64   `yaml:"charID"`
	EntityID  int64   `yaml:"entityID"`
	IskValue  float64 `yaml:"iskValue"`
	OfferedID int64   `yaml:"offeredID"`
}

type OfferedToAlly struct {
	CharID     int64   `yaml:"charID"`
	DefenderID int64   `yaml:"defenderID"`
	EnemyID    int64   `yaml:"enemyID"`
	IskValue   float64 `yaml:"iskValue"`
}

type OldLscMessages struct {
	Subject string `yaml:"subject"`
	Body    string `yaml:"body"`
}

type OperationFinished struct {
	OperationID int64 `yaml:"operationID"`
	Rewards     struct {
		Isk int32 `yaml:"isk"`
	} `yaml:"rewards"`
}

type OrbitalAttacked struct {
	AggressorAllianceID int64   `yaml:"aggressorAllianceID"`
	AggressorCorpID     int64   `yaml:"aggressorCorpID"`
	AggressorID         int64   `yaml:"aggressorID"`
	PlanetID            int64   `yaml:"planetID"`
	PlanetTypeID        int64   `yaml:"planetTypeID"`
	ShieldLevel         float64 `yaml:"shieldLevel"`
	SolarSystemID       int64   `yaml:"solarSystemID"`
	TypeID              int64   `yaml:"typeID"`
}

type OrbitalReinforced struct {
	AggressorAllianceID int64 `yaml:"aggressorAllianceID"`
	AggressorCorpID     int64 `yaml:"aggressorCorpID"`
	AggressorID         int64 `yaml:"aggressorID"`
	PlanetID            int64 `yaml:"planetID"`
	PlanetTypeID        int64 `yaml:"planetTypeID"`
	ReinforceExitTime   int64 `yaml:"reinforceExitTime"`
	SolarSystemID       int64 `yaml:"solarSystemID"`
	TypeID              int64 `yaml:"typeID"`
}

type OwnershipTransferred struct {
	CharacterLinkData       []interface{} `yaml:"characterLinkData"`
	CharacterName           string        `yaml:"characterName"`
	FromCorporationLinkData []interface{} `yaml:"fromCorporationLinkData"`
	FromCorporationName     string        `yaml:"fromCorporationName"`
	SolarSystemLinkData     []interface{} `yaml:"solarSystemLinkData"`
	SolarSystemName         string        `yaml:"solarSystemName"`
	StructureLinkData       []interface{} `yaml:"structureLinkData"`
	StructureName           string        `yaml:"structureName"`
	ToCorporationLinkData   []interface{} `yaml:"toCorporationLinkData"`
	ToCorporationName       string        `yaml:"toCorporationName"`
}

// OwnershipTransferredV2 represents the updated version for this notification type.
// Use [OwnershipTransferred] to unmarshal older notifications.
type OwnershipTransferredV2 struct {
	CharID          int64  `yaml:"charID"`
	NewOwnerCorpID  int64  `yaml:"newOwnerCorpID"`
	OldOwnerCorpID  int64  `yaml:"oldOwnerCorpID"`
	SolarSystemID   int64  `yaml:"solarSystemID"`
	StructureID     int64  `yaml:"structureID"`
	StructureName   string `yaml:"structureName"`
	StructureTypeID int64  `yaml:"structureTypeID"`
}

type ReimbursementMsg struct {
	AddCloneInfo  int32 `yaml:"addCloneInfo"`
	ShipTypeID    int64 `yaml:"shipTypeID"`
	SolarSystemID int64 `yaml:"solarSystemID"`
	StationID     int64 `yaml:"stationID"`
}

type ResearchMissionAvailableMsg struct {
}

type RetractsWar struct {
	CharID  int64 `yaml:"charID"`
	EnemyID int64 `yaml:"enemyID"`
}

type SeasonalChallengeCompleted struct {
	ChallengeNameID int64 `yaml:"challenge_name_id"`
	MaxProgress     int32 `yaml:"max_progress"`
	PointsAwarded   int32 `yaml:"points_awarded"`
	ProgressText    int32 `yaml:"progress_text"`
}

type SkyhookLostShields struct {
	ItemID              int64         `yaml:"itemID"`
	PlanetID            int64         `yaml:"planetID"`
	PlanetShowInfoData  []interface{} `yaml:"planetShowInfoData"`
	SkyhookShowInfoData []interface{} `yaml:"skyhookShowInfoData"`
	SolarsystemID       int64         `yaml:"solarsystemID"`
	TimeLeft            int64         `yaml:"timeLeft"`
	Timestamp           int64         `yaml:"timestamp"`
	TypeID              int64         `yaml:"typeID"`
	VulnerableTime      int64         `yaml:"vulnerableTime"`
}

type SkyhookUnderAttack struct {
	AllianceID            int64         `yaml:"allianceID"`
	AllianceLinkData      []interface{} `yaml:"allianceLinkData"`
	AllianceName          string        `yaml:"allianceName"`
	ArmorPercentage       float64       `yaml:"armorPercentage"`
	CharID                int64         `yaml:"charID"`
	CorpLinkData          []interface{} `yaml:"corpLinkData"`
	CorpName              string        `yaml:"corpName"`
	IsActive              bool          `yaml:"isActive"`
	ItemID                string        `yaml:"itemID"`
	PlanetID              int64         `yaml:"planetID"`
	PlanetShowInfoData    []interface{} `yaml:"planetShowInfoData"`
	HullPercentage        float64       `yaml:"hullPercentage"`
	ShieldPercentage      float64       `yaml:"shieldPercentage"`
	SolarsystemID         int64         `yaml:"solarsystemID"`
	StructureID           int64         `yaml:"structureID"`
	StructureShowInfoData []interface{} `yaml:"structureShowInfoData"`
	SkyhookShowInfoData   []interface{} `yaml:"skyhookShowInfoData"`
	TypeID                int64         `yaml:"typeID"`
}

type SovAllClaimAquiredMsg struct {
	AllianceID    int64 `yaml:"allianceID"`
	CorpID        int64 `yaml:"corpID"`
	SolarSystemID int64 `yaml:"solarSystemID"`
}

type SovAllClaimLostMsg struct {
	AllianceID    int64 `yaml:"allianceID"`
	CorpID        int64 `yaml:"corpID"`
	SolarSystemID int64 `yaml:"solarSystemID"`
}

type SovCommandNodeEventStarted struct {
	CampaignEventType int32 `yaml:"campaignEventType"`
	ConstellationID   int64 `yaml:"constellationID"`
	SolarSystemID     int64 `yaml:"solarSystemID"`
}

type SovStationEnteredFreeport struct {
	Freeportexittime int64 `yaml:"freeportexittime"`
	SolarSystemID    int64 `yaml:"solarSystemID"`
	StructureTypeID  int64 `yaml:"structureTypeID"`
}

type SovStructureDestroyed struct {
	SolarSystemID   int64 `yaml:"solarSystemID"`
	StructureTypeID int64 `yaml:"structureTypeID"`
}

type SovStructureReinforced struct {
	CampaignEventType int32 `yaml:"campaignEventType"`
	DecloakTime       int64 `yaml:"decloakTime"`
	SolarSystemID     int64 `yaml:"solarSystemID"`
}

type SovStructureSelfDestructCancel struct {
	CharID          int64 `yaml:"charID"`
	SolarSystemID   int64 `yaml:"solarSystemID"`
	StructureTypeID int64 `yaml:"structureTypeID"`
}

type SovStructureSelfDestructFinished struct {
	SolarSystemID   int64 `yaml:"solarSystemID"`
	StructureTypeID int64 `yaml:"structureTypeID"`
}

type SovStructureSelfDestructRequested struct {
	CharID          int64  `yaml:"charID"`
	CorpName        string `yaml:"corpName"`
	DestructTime    int64  `yaml:"destructTime"`
	SolarSystemID   int64  `yaml:"solarSystemID"`
	StructureTypeID int64  `yaml:"structureTypeID"`
}

type SovereigntyIHDamageMsg struct {
	AggressorAllianceID int64   `yaml:"aggressorAllianceID"`
	AggressorCorpID     int64   `yaml:"aggressorCorpID"`
	AggressorID         int64   `yaml:"aggressorID"`
	ArmorValue          float64 `yaml:"armorValue"`
	HullValue           float64 `yaml:"hullValue"`
	ShieldValue         float64 `yaml:"shieldValue"`
	SolarSystemID       int64   `yaml:"solarSystemID"`
}

type SovereigntySBUDamageMsg struct {
	AggressorAllianceID int64   `yaml:"aggressorAllianceID"`
	AggressorCorpID     int64   `yaml:"aggressorCorpID"`
	AggressorID         int64   `yaml:"aggressorID"`
	ArmorValue          float64 `yaml:"armorValue"`
	HullValue           float64 `yaml:"hullValue"`
	ShieldValue         float64 `yaml:"shieldValue"`
	SolarSystemID       int64   `yaml:"solarSystemID"`
}

type SovereigntyTCUDamageMsg struct {
	AggressorAllianceID int64   `yaml:"aggressorAllianceID"`
	AggressorCorpID     int64   `yaml:"aggressorCorpID"`
	AggressorID         int64   `yaml:"aggressorID"`
	ArmorValue          float64 `yaml:"armorValue"`
	HullValue           float64 `yaml:"hullValue"`
	ShieldValue         float64 `yaml:"shieldValue"`
	SolarSystemID       int64   `yaml:"solarSystemID"`
}

type StationServiceDisabled struct {
	SolarSystemID   int64 `yaml:"solarSystemID"`
	StructureTypeID int64 `yaml:"structureTypeID"`
}

type StationServiceEnabled struct {
	SolarSystemID   int64 `yaml:"solarSystemID"`
	StructureTypeID int64 `yaml:"structureTypeID"`
}

type StructureAnchoring struct {
	OwnerCorpLinkData     []interface{} `yaml:"ownerCorpLinkData"`
	OwnerCorpName         string        `yaml:"ownerCorpName"`
	SolarsystemID         int64         `yaml:"solarsystemID"`
	StructureID           int64         `yaml:"structureID"`
	StructureShowInfoData []interface{} `yaml:"structureShowInfoData"`
	StructureTypeID       int64         `yaml:"structureTypeID"`
	TimeLeft              int64         `yaml:"timeLeft"`
	VulnerableTime        int64         `yaml:"vulnerableTime"`
}

type StructureDestroyed struct {
	OwnerCorpLinkData     []interface{} `yaml:"ownerCorpLinkData"`
	OwnerCorpName         string        `yaml:"ownerCorpName"`
	SolarsystemID         int64         `yaml:"solarsystemID"`
	StructureID           int64         `yaml:"structureID"`
	StructureShowInfoData []interface{} `yaml:"structureShowInfoData"`
	StructureTypeID       int64         `yaml:"structureTypeID"`
}

type StructureFuelAlert struct {
	ListOfTypesAndQty     [][]int64     `yaml:"listOfTypesAndQty"`
	SolarsystemID         int64         `yaml:"solarsystemID"`
	StructureID           int64         `yaml:"structureID"`
	StructureShowInfoData []interface{} `yaml:"structureShowInfoData"`
	StructureTypeID       int64         `yaml:"structureTypeID"`
}

type StructureImpendingAbandonmentAssetsAtRisk struct {
	DaysUntilAbandon      int32         `yaml:"daysUntilAbandon"`
	IsCorpOwned           bool          `yaml:"isCorpOwned"`
	SolarSystemID         int64         `yaml:"solarsystemID"`
	StructureID           int64         `yaml:"structureID"`
	StructureLink         string        `yaml:"structureLink"`
	StructureShowInfoData []interface{} `yaml:"structureShowInfoData"`
	StructureTypeID       int64         `yaml:"structureTypeID"`
}

type StructureItemsDelivered struct {
	CharID                int64         `yaml:"charID"`
	ListOfTypesAndQty     [][]int64     `yaml:"listOfTypesAndQty"`
	SolarsystemID         int64         `yaml:"solarsystemID"`
	StructureID           int64         `yaml:"structureID"`
	StructureShowInfoData []interface{} `yaml:"structureShowInfoData"`
	StructureTypeID       int64         `yaml:"structureTypeID"`
}

type StructureItemsMovedToSafety struct {
	AssetSafetyDurationFull     int64         `yaml:"assetSafetyDurationFull"`
	AssetSafetyDurationMinimum  int64         `yaml:"assetSafetyDurationMinimum"`
	AssetSafetyFullTimestamp    int64         `yaml:"assetSafetyFullTimestamp"`
	AssetSafetyMinimumTimestamp int64         `yaml:"assetSafetyMinimumTimestamp"`
	IsCorpOwned                 bool          `yaml:"isCorpOwned"`
	NewStationID                int64         `yaml:"newStationID"`
	SolarSystemID               int64         `yaml:"solarsystemID"`
	StructureID                 int64         `yaml:"structureID"`
	StructureLink               string        `yaml:"structureLink"`
	StructureShowInfoData       []interface{} `yaml:"structureShowInfoData"`
	StructureTypeID             int64         `yaml:"structureTypeID"`
}

type StructureLostArmor struct {
	SolarsystemID         int64         `yaml:"solarsystemID"`
	StructureID           int64         `yaml:"structureID"`
	StructureShowInfoData []interface{} `yaml:"structureShowInfoData"`
	StructureTypeID       int64         `yaml:"structureTypeID"`
	TimeLeft              int64         `yaml:"timeLeft"`
	Timestamp             int64         `yaml:"timestamp"`
	VulnerableTime        int64         `yaml:"vulnerableTime"`
}

type StructureLostShields struct {
	SolarsystemID         int64         `yaml:"solarsystemID"`
	StructureID           int64         `yaml:"structureID"`
	StructureShowInfoData []interface{} `yaml:"structureShowInfoData"`
	StructureTypeID       int64         `yaml:"structureTypeID"`
	TimeLeft              int64         `yaml:"timeLeft"`
	Timestamp             int64         `yaml:"timestamp"`
	VulnerableTime        int64         `yaml:"vulnerableTime"`
}

type StructureOnline struct {
	SolarsystemID         int64         `yaml:"solarsystemID"`
	StructureID           int64         `yaml:"structureID"`
	StructureShowInfoData []interface{} `yaml:"structureShowInfoData"`
	StructureTypeID       int64         `yaml:"structureTypeID"`
}

type StructureServicesOffline struct {
	ListOfServiceModuleIDs []int64       `yaml:"listOfServiceModuleIDs"`
	SolarsystemID          int64         `yaml:"solarsystemID"`
	StructureID            int64         `yaml:"structureID"`
	StructureShowInfoData  []interface{} `yaml:"structureShowInfoData"`
	StructureTypeID        int64         `yaml:"structureTypeID"`
}

type StructureUnanchoring struct {
	OwnerCorpLinkData     []interface{} `yaml:"ownerCorpLinkData"`
	OwnerCorpName         string        `yaml:"ownerCorpName"`
	SolarsystemID         int64         `yaml:"solarsystemID"`
	StructureID           int64         `yaml:"structureID"`
	StructureShowInfoData []interface{} `yaml:"structureShowInfoData"`
	StructureTypeID       int64         `yaml:"structureTypeID"`
	TimeLeft              int64         `yaml:"timeLeft"`
}

type StructureUnderAttack struct {
	AllianceID            int64         `yaml:"allianceID"`
	AllianceLinkData      []interface{} `yaml:"allianceLinkData"`
	AllianceName          string        `yaml:"allianceName"`
	ArmorPercentage       float64       `yaml:"armorPercentage"`
	CharID                int64         `yaml:"charID"`
	CorpLinkData          []interface{} `yaml:"corpLinkData"`
	CorpName              string        `yaml:"corpName"`
	HullPercentage        float64       `yaml:"hullPercentage"`
	ShieldPercentage      float64       `yaml:"shieldPercentage"`
	SolarsystemID         int64         `yaml:"solarsystemID"`
	StructureID           int64         `yaml:"structureID"`
	StructureShowInfoData []interface{} `yaml:"structureShowInfoData"`
	StructureTypeID       int64         `yaml:"structureTypeID"`
}

type StructureWentHighPower struct {
	SolarsystemID         int64         `yaml:"solarsystemID"`
	StructureID           int64         `yaml:"structureID"`
	StructureShowInfoData []interface{} `yaml:"structureShowInfoData"`
	StructureTypeID       int64         `yaml:"structureTypeID"`
}

type StructureWentLowPower struct {
	SolarsystemID         int64         `yaml:"solarsystemID"`
	StructureID           int64         `yaml:"structureID"`
	StructureShowInfoData []interface{} `yaml:"structureShowInfoData"`
	StructureTypeID       int64         `yaml:"structureTypeID"`
}

type StructuresReinforcementChanged struct {
	AllStructureInfo [][]interface{} `yaml:"allStructureInfo"`
	Hour             int32           `yaml:"hour"`
	NumStructures    int32           `yaml:"numStructures"`
	Timestamp        int64           `yaml:"timestamp"`
	Weekday          int32           `yaml:"weekday"`
}

type TowerAlertMsg struct {
	AggressorAllianceID int64   `yaml:"aggressorAllianceID"`
	AggressorCorpID     int64   `yaml:"aggressorCorpID"`
	AggressorID         int64   `yaml:"aggressorID"`
	ArmorValue          float64 `yaml:"armorValue"`
	HullValue           float64 `yaml:"hullValue"`
	MoonID              int64   `yaml:"moonID"`
	ShieldValue         float64 `yaml:"shieldValue"`
	SolarSystemID       int64   `yaml:"solarSystemID"`
	TypeID              int64   `yaml:"typeID"`
}

type TowerResourceAlertMsg struct {
	AllianceID    int64 `yaml:"allianceID"`
	CorpID        int64 `yaml:"corpID"`
	MoonID        int64 `yaml:"moonID"`
	SolarSystemID int64 `yaml:"solarSystemID"`
	TypeID        int64 `yaml:"typeID"`
	Wants         []struct {
		Quantity int32 `yaml:"quantity"`
		TypeID   int64 `yaml:"typeID"`
	} `yaml:"wants"`
}

type WarAdopted struct {
	AgainstID    int64 `yaml:"againstID"`
	AllianceID   int64 `yaml:"allianceID"`
	DeclaredByID int64 `yaml:"declaredByID"`
}

type WarAllyOfferDeclinedMsg struct {
	AggressorID int64 `yaml:"aggressorID"`
	AllyID      int64 `yaml:"allyID"`
	CharID      int64 `yaml:"charID"`
	DefenderID  int64 `yaml:"defenderID"`
}

type WarDeclared struct {
	AgainstID    int64         `yaml:"againstID"`
	Cost         float64       `yaml:"cost"`
	DeclaredByID int64         `yaml:"declaredByID"`
	DelayHours   int32         `yaml:"delayHours"`
	HostileState bool          `yaml:"hostileState"`
	TimeStarted  int64         `yaml:"timeStarted"`
	WarHQ        string        `yaml:"warHQ"`
	WarHQIdType  []interface{} `yaml:"warHQ_IdType"`
}

type WarHQRemovedFromSpace struct {
	AgainstID    int64  `yaml:"againstID"`
	DeclaredByID int64  `yaml:"declaredByID"`
	TimeDeclared int64  `yaml:"timeDeclared"`
	WarHQ        string `yaml:"warHQ"`
}

type WarInherited struct {
	AgainstID    int64 `yaml:"againstID"`
	AllianceID   int64 `yaml:"allianceID"`
	DeclaredByID int64 `yaml:"declaredByID"`
	OpponentID   int64 `yaml:"opponentID"`
	QuitterID    int64 `yaml:"quitterID"`
}

type WarInvalid struct {
	AgainstID    int64 `yaml:"againstID"`
	DeclaredByID int64 `yaml:"declaredByID"`
	EndDate      int64 `yaml:"endDate"`
}

type WarRetractedByConcord struct {
	AgainstID    int64 `yaml:"againstID"`
	DeclaredByID int64 `yaml:"declaredByID"`
	EndDate      int64 `yaml:"endDate"`
}

type WarSurrenderDeclinedMsg struct {
	IskValue float64 `yaml:"iskValue"`
	OwnerID  int64   `yaml:"ownerID"`
}

type WarSurrenderOfferMsg struct {
	IskValue         float64 `yaml:"iskValue"`
	OwnerID1         int64   `yaml:"ownerID1"`
	OwnerID2         int64   `yaml:"ownerID2"`
	WarNegotiationID int64   `yaml:"warNegotiationID"`
}

// Deprecated: Use [MoonminingExtractionStarted] instead.
type NotificationTypeMoonminingExtractionStarted struct {
	AutoTime        int64             `yaml:"autoTime"`
	MoonID          int64             `yaml:"moonID"`
	MoonLink        string            `yaml:"moonLink"`
	OreVolumeByType map[int64]float64 `yaml:"oreVolumeByType"`
	ReadyTime       int64             `yaml:"readyTime"`
	SolarSystemID   int64             `yaml:"solarSystemID"`
	SolarSystemLink string            `yaml:"solarSystemLink"`
	StartedBy       int64             `yaml:"startedBy"`
	StartedByLink   string            `yaml:"startedByLink"`
	StructureID     int64             `yaml:"structureID"`
	StructureLink   string            `yaml:"structureLink"`
	StructureName   string            `yaml:"structureName"`
	StructureTypeID int64             `yaml:"structureTypeID"`
}
