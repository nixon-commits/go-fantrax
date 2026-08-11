package models

import "time"

// TransactionHistoryResponse represents the full response from getTransactionDetailsHistory
type TransactionHistoryResponse struct {
	Data struct {
		SDate int64  `json:"sDate"`
		Adrt  int    `json:"adrt"`
		Up    string `json:"up"`
	} `json:"data"`
	Roles     []string                  `json:"roles"`
	Responses []TransactionDataResponse `json:"responses"`
}

// TransactionDataResponse represents a single response in the responses array
type TransactionDataResponse struct {
	Data TransactionData `json:"data"`
}

// TransactionData contains the main transaction data
type TransactionData struct {
	PaginatedResultSet  PaginatedResultSet     `json:"paginatedResultSet"`
	FilterSettings      TransactionFilter      `json:"filterSettings"`
	DisplayedSelections TransactionFilter      `json:"displayedSelections"`
	MiscData            map[string]interface{} `json:"miscData"`
	DisplayedLists      DisplayedLists         `json:"displayedLists"`
	Table               TransactionTable       `json:"table"`
}

// PaginatedResultSet contains pagination information
type PaginatedResultSet struct {
	TotalNumPages     int `json:"totalNumPages"`
	PageNumber        int `json:"pageNumber"`
	MaxResultsPerPage int `json:"maxResultsPerPage"`
	TotalNumResults   int `json:"totalNumResults"`
}

// TransactionFilter represents filter settings
type TransactionFilter struct {
	PositionOrGroup string `json:"positionOrGroup"`
	View            string `json:"view"`
	AdminMode       bool   `json:"adminMode"`
	IncludeDeleted  bool   `json:"includeDeleted"`
	Team            string `json:"team"`
	ExecutedOnly    bool   `json:"executedOnly"`
}

// DisplayedLists contains lists of teams and other displayable data
type DisplayedLists struct {
	Teams []TeamOption `json:"teams"`
}

// TeamOption represents a team in the dropdown
type TeamOption struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

// TransactionTable contains the table structure with headers and rows
type TransactionTable struct {
	Caption string            `json:"caption"`
	Header  TransactionHeader `json:"header"`
	Rows    []TransactionRow  `json:"rows"`
}

// TransactionHeader contains the column headers
type TransactionHeader struct {
	Cells []HeaderCell `json:"cells"`
}

// HeaderCell represents a single header cell
type HeaderCell struct {
	Align         string `json:"align,omitempty"`
	SortDirection int    `json:"sortDirection,omitempty"`
	Name          string `json:"name"`
	ShortName     string `json:"shortName"`
	Key           string `json:"key"`
}

// TransactionRow represents a single transaction
type TransactionRow struct {
	Scorer          TransactionPlayer `json:"scorer"`
	ResultCode      string            `json:"resultCode"`
	Executed        bool              `json:"executed"`
	Result          CellContent       `json:"result"`
	ClaimType       string            `json:"claimType,omitempty"`
	NumInGroup      int               `json:"numInGroup,omitempty"`
	TxSetID         string            `json:"txSetId"`
	FeesUsed        bool              `json:"feesUsed"`
	TransactionCode string            `json:"transactionCode"`
	TransactionType string            `json:"transactionType"`
	Deleted         bool              `json:"deleted"`
	Disabled        bool              `json:"disabled,omitempty"`
	Cells           []TableCell       `json:"cells"`
	LinkedRows      []interface{}     `json:"linkedRows,omitempty"`
	// DraftPickDisplayParts is present instead of a populated Scorer when the
	// row represents a traded draft pick rather than a player. Fantrax gives
	// pick rows an empty scorer object -- this is the only place the pick's
	// year and round survive in the payload.
	DraftPickDisplayParts *DraftPickDisplayParts `json:"draftPickDisplayParts,omitempty"`
}

// DraftPickDisplayParts identifies a traded draft pick. Both fields carry
// inline HTML markup from Fantrax's own display rendering (e.g.
// "<b>2027</b> Draft Pick" / "Round <b>3</b> (Houston Swang and Bang)") and
// are plain display strings, not structured data -- callers that need the
// year/round/team as values must parse them (see ParseTransactions, which
// does this for the simplified Transaction).
type DraftPickDisplayParts struct {
	Year      string `json:"year"`
	RoundInfo string `json:"roundInfo"`
}

// TransactionPlayer represents player information in a transaction
type TransactionPlayer struct {
	TeamName       string   `json:"teamName"`
	URLName        string   `json:"urlName"`
	HeadshotURL    string   `json:"headshotUrl"`
	ScorerID       string   `json:"scorerId"`
	PosIDsNoFlex   []string `json:"posIdsNoFlex"`
	DefaultPosID   string   `json:"defaultPosId"`
	PosShortNames  string   `json:"posShortNames"`
	Team           bool     `json:"team"`
	PrimaryPosID   string   `json:"primaryPosId"`
	Rookie         bool     `json:"rookie"`
	MinorsEligible bool     `json:"minorsEligible"`
	PosIDs         []string `json:"posIds"`
	TeamID         string   `json:"teamId"`
	Name           string   `json:"name"`
	TeamShortName  string   `json:"teamShortName"`
	ShortName      string   `json:"shortName"`
}

// CellContent represents content within a cell
type CellContent struct {
	Content string `json:"content"`
}

// TableCell represents a cell in the transaction table
type TableCell struct {
	Align       string `json:"align,omitempty"`
	Content     string `json:"content"`
	LeagueID    string `json:"leagueId,omitempty"`
	Rowspan     int    `json:"rowspan,omitempty"`
	Key         string `json:"key"`
	TeamID      string `json:"teamId,omitempty"`
	Icon        string `json:"icon,omitempty"`
	IconToolTip string `json:"iconToolTip,omitempty"`
	ToolTip     string `json:"toolTip,omitempty"`
}

// Transaction represents a simplified transaction for easier use
type Transaction struct {
	ID             string    `json:"id"`
	Type           string    `json:"type"`                   // "CLAIM", "DROP", "TRADE"
	ClaimType      string    `json:"claimType,omitempty"`    // "FA" (Free Agent) or "WW" (Waiver Wire) for CLAIM transactions
	TeamName       string    `json:"teamName"`               // For CLAIM/DROP transactions
	TeamID         string    `json:"teamId"`                 // For CLAIM/DROP transactions
	FromTeamName   string    `json:"fromTeamName,omitempty"` // For TRADE transactions
	FromTeamID     string    `json:"fromTeamId,omitempty"`   // For TRADE transactions
	ToTeamName     string    `json:"toTeamName,omitempty"`   // For TRADE transactions
	ToTeamID       string    `json:"toTeamId,omitempty"`     // For TRADE transactions
	PlayerName     string    `json:"playerName"`
	PlayerID       string    `json:"playerId"`
	PlayerTeam     string    `json:"playerTeam"`
	PlayerPosition string    `json:"playerPosition"`
	BidAmount      string    `json:"bidAmount,omitempty"`
	Priority       string    `json:"priority,omitempty"`
	ProcessedDate  time.Time `json:"processedDate"`
	Period         int       `json:"period"`
	Executed       bool      `json:"executed"`
	ExecutedBy     string    `json:"executedBy,omitempty"`     // "COMMISSIONER" if commissioner executed
	TradeGroupID   string    `json:"tradeGroupId,omitempty"`   // txSetId for grouping trade players
	TradeGroupSize int       `json:"tradeGroupSize,omitempty"` // numInGroup for trades

	// IsDraftPick and the DraftPick* fields below are populated instead of
	// PlayerName/PlayerID/PlayerPosition when the row moves a future draft
	// pick rather than a player -- parsed from the row's
	// DraftPickDisplayParts. PlayerName stays "" for a pick row, which is
	// the pre-existing signal callers already key off of.
	IsDraftPick bool `json:"isDraftPick,omitempty"`
	// DraftPickYear is the draft the pick belongs to (e.g. 2027).
	DraftPickYear int `json:"draftPickYear,omitempty"`
	// DraftPickRound is the round number (1st = 1, 2nd = 2, ...).
	DraftPickRound int `json:"draftPickRound,omitempty"`
	// DraftPickNumber is the specific slot within the round (e.g. 6), set
	// only once the draft order is known. Zero means the slot is still a
	// projection tied to DraftPickOriginalTeam's eventual standing.
	DraftPickNumber int `json:"draftPickNumber,omitempty"`
	// DraftPickOriginalTeam is the team whose future draft slot this is --
	// distinct from FromTeamName/ToTeamName, which describe this specific
	// trade's current holder and receiver. A pick can be re-traded before
	// it is ever drafted, so the original team can differ from both. Empty
	// once DraftPickNumber is known, since the slot no longer depends on
	// which team's future record it was.
	DraftPickOriginalTeam string `json:"draftPickOriginalTeam,omitempty"`
}
