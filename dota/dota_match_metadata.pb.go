// Code generated by protoc-gen-go.
// source: dota_match_metadata.proto
// DO NOT EDIT!

package dota

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CDOTAMatchMetadataFile struct {
	Version          *int32              `protobuf:"varint,1,req,name=version" json:"version,omitempty"`
	MatchId          *uint64             `protobuf:"varint,2,req,name=match_id" json:"match_id,omitempty"`
	Metadata         *CDOTAMatchMetadata `protobuf:"bytes,3,opt,name=metadata" json:"metadata,omitempty"`
	PrivateMetadata  []byte              `protobuf:"bytes,5,opt,name=private_metadata" json:"private_metadata,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *CDOTAMatchMetadataFile) Reset()                    { *m = CDOTAMatchMetadataFile{} }
func (m *CDOTAMatchMetadataFile) String() string            { return proto.CompactTextString(m) }
func (*CDOTAMatchMetadataFile) ProtoMessage()               {}
func (*CDOTAMatchMetadataFile) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{0} }

func (m *CDOTAMatchMetadataFile) GetVersion() int32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *CDOTAMatchMetadataFile) GetMatchId() uint64 {
	if m != nil && m.MatchId != nil {
		return *m.MatchId
	}
	return 0
}

func (m *CDOTAMatchMetadataFile) GetMetadata() *CDOTAMatchMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *CDOTAMatchMetadataFile) GetPrivateMetadata() []byte {
	if m != nil {
		return m.PrivateMetadata
	}
	return nil
}

type CDOTAMatchMetadata struct {
	Teams            []*CDOTAMatchMetadata_Team  `protobuf:"bytes,1,rep,name=teams" json:"teams,omitempty"`
	ItemRewards      []*CLobbyTimedRewardDetails `protobuf:"bytes,2,rep,name=item_rewards" json:"item_rewards,omitempty"`
	LobbyId          *uint64                     `protobuf:"fixed64,3,opt,name=lobby_id" json:"lobby_id,omitempty"`
	ReportUntilTime  *uint64                     `protobuf:"fixed64,4,opt,name=report_until_time" json:"report_until_time,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (m *CDOTAMatchMetadata) Reset()                    { *m = CDOTAMatchMetadata{} }
func (m *CDOTAMatchMetadata) String() string            { return proto.CompactTextString(m) }
func (*CDOTAMatchMetadata) ProtoMessage()               {}
func (*CDOTAMatchMetadata) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{1} }

func (m *CDOTAMatchMetadata) GetTeams() []*CDOTAMatchMetadata_Team {
	if m != nil {
		return m.Teams
	}
	return nil
}

func (m *CDOTAMatchMetadata) GetItemRewards() []*CLobbyTimedRewardDetails {
	if m != nil {
		return m.ItemRewards
	}
	return nil
}

func (m *CDOTAMatchMetadata) GetLobbyId() uint64 {
	if m != nil && m.LobbyId != nil {
		return *m.LobbyId
	}
	return 0
}

func (m *CDOTAMatchMetadata) GetReportUntilTime() uint64 {
	if m != nil && m.ReportUntilTime != nil {
		return *m.ReportUntilTime
	}
	return 0
}

type CDOTAMatchMetadata_Team struct {
	DotaTeam          *uint32                           `protobuf:"varint,1,opt,name=dota_team" json:"dota_team,omitempty"`
	Players           []*CDOTAMatchMetadata_Team_Player `protobuf:"bytes,2,rep,name=players" json:"players,omitempty"`
	GraphExperience   []float32                         `protobuf:"fixed32,3,rep,name=graph_experience" json:"graph_experience,omitempty"`
	GraphGoldEarned   []float32                         `protobuf:"fixed32,4,rep,name=graph_gold_earned" json:"graph_gold_earned,omitempty"`
	GraphNetWorth     []float32                         `protobuf:"fixed32,5,rep,name=graph_net_worth" json:"graph_net_worth,omitempty"`
	CmFirstPick       *bool                             `protobuf:"varint,6,opt,name=cm_first_pick" json:"cm_first_pick,omitempty"`
	CmCaptainPlayerId *uint32                           `protobuf:"varint,7,opt,name=cm_captain_player_id" json:"cm_captain_player_id,omitempty"`
	CmBans            []uint32                          `protobuf:"varint,8,rep,name=cm_bans" json:"cm_bans,omitempty"`
	CmPicks           []uint32                          `protobuf:"varint,9,rep,name=cm_picks" json:"cm_picks,omitempty"`
	CmPenalty         *uint32                           `protobuf:"varint,10,opt,name=cm_penalty" json:"cm_penalty,omitempty"`
	XXX_unrecognized  []byte                            `json:"-"`
}

func (m *CDOTAMatchMetadata_Team) Reset()                    { *m = CDOTAMatchMetadata_Team{} }
func (m *CDOTAMatchMetadata_Team) String() string            { return proto.CompactTextString(m) }
func (*CDOTAMatchMetadata_Team) ProtoMessage()               {}
func (*CDOTAMatchMetadata_Team) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{1, 0} }

func (m *CDOTAMatchMetadata_Team) GetDotaTeam() uint32 {
	if m != nil && m.DotaTeam != nil {
		return *m.DotaTeam
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team) GetPlayers() []*CDOTAMatchMetadata_Team_Player {
	if m != nil {
		return m.Players
	}
	return nil
}

func (m *CDOTAMatchMetadata_Team) GetGraphExperience() []float32 {
	if m != nil {
		return m.GraphExperience
	}
	return nil
}

func (m *CDOTAMatchMetadata_Team) GetGraphGoldEarned() []float32 {
	if m != nil {
		return m.GraphGoldEarned
	}
	return nil
}

func (m *CDOTAMatchMetadata_Team) GetGraphNetWorth() []float32 {
	if m != nil {
		return m.GraphNetWorth
	}
	return nil
}

func (m *CDOTAMatchMetadata_Team) GetCmFirstPick() bool {
	if m != nil && m.CmFirstPick != nil {
		return *m.CmFirstPick
	}
	return false
}

func (m *CDOTAMatchMetadata_Team) GetCmCaptainPlayerId() uint32 {
	if m != nil && m.CmCaptainPlayerId != nil {
		return *m.CmCaptainPlayerId
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team) GetCmBans() []uint32 {
	if m != nil {
		return m.CmBans
	}
	return nil
}

func (m *CDOTAMatchMetadata_Team) GetCmPicks() []uint32 {
	if m != nil {
		return m.CmPicks
	}
	return nil
}

func (m *CDOTAMatchMetadata_Team) GetCmPenalty() uint32 {
	if m != nil && m.CmPenalty != nil {
		return *m.CmPenalty
	}
	return 0
}

type CDOTAMatchMetadata_Team_PlayerKill struct {
	VictimSlot       *uint32 `protobuf:"varint,1,opt,name=victim_slot" json:"victim_slot,omitempty"`
	Count            *uint32 `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CDOTAMatchMetadata_Team_PlayerKill) Reset()         { *m = CDOTAMatchMetadata_Team_PlayerKill{} }
func (m *CDOTAMatchMetadata_Team_PlayerKill) String() string { return proto.CompactTextString(m) }
func (*CDOTAMatchMetadata_Team_PlayerKill) ProtoMessage()    {}
func (*CDOTAMatchMetadata_Team_PlayerKill) Descriptor() ([]byte, []int) {
	return fileDescriptor21, []int{1, 0, 0}
}

func (m *CDOTAMatchMetadata_Team_PlayerKill) GetVictimSlot() uint32 {
	if m != nil && m.VictimSlot != nil {
		return *m.VictimSlot
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_PlayerKill) GetCount() uint32 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return 0
}

type CDOTAMatchMetadata_Team_ItemPurchase struct {
	ItemId           *uint32 `protobuf:"varint,1,opt,name=item_id" json:"item_id,omitempty"`
	PurchaseTime     *int32  `protobuf:"varint,2,opt,name=purchase_time" json:"purchase_time,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CDOTAMatchMetadata_Team_ItemPurchase) Reset()         { *m = CDOTAMatchMetadata_Team_ItemPurchase{} }
func (m *CDOTAMatchMetadata_Team_ItemPurchase) String() string { return proto.CompactTextString(m) }
func (*CDOTAMatchMetadata_Team_ItemPurchase) ProtoMessage()    {}
func (*CDOTAMatchMetadata_Team_ItemPurchase) Descriptor() ([]byte, []int) {
	return fileDescriptor21, []int{1, 0, 1}
}

func (m *CDOTAMatchMetadata_Team_ItemPurchase) GetItemId() uint32 {
	if m != nil && m.ItemId != nil {
		return *m.ItemId
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_ItemPurchase) GetPurchaseTime() int32 {
	if m != nil && m.PurchaseTime != nil {
		return *m.PurchaseTime
	}
	return 0
}

type CDOTAMatchMetadata_Team_Player struct {
	AccountId         *uint32                                 `protobuf:"varint,1,opt,name=account_id" json:"account_id,omitempty"`
	AbilityUpgrades   []uint32                                `protobuf:"varint,2,rep,name=ability_upgrades" json:"ability_upgrades,omitempty"`
	PlayerSlot        *uint32                                 `protobuf:"varint,3,opt,name=player_slot" json:"player_slot,omitempty"`
	EquippedEconItems []*CSOEconItem                          `protobuf:"bytes,4,rep,name=equipped_econ_items" json:"equipped_econ_items,omitempty"`
	Kills             []*CDOTAMatchMetadata_Team_PlayerKill   `protobuf:"bytes,5,rep,name=kills" json:"kills,omitempty"`
	Items             []*CDOTAMatchMetadata_Team_ItemPurchase `protobuf:"bytes,6,rep,name=items" json:"items,omitempty"`
	AvgKillsX16       *uint32                                 `protobuf:"varint,7,opt,name=avg_kills_x16" json:"avg_kills_x16,omitempty"`
	AvgDeathsX16      *uint32                                 `protobuf:"varint,8,opt,name=avg_deaths_x16" json:"avg_deaths_x16,omitempty"`
	AvgAssistsX16     *uint32                                 `protobuf:"varint,9,opt,name=avg_assists_x16" json:"avg_assists_x16,omitempty"`
	AvgGpmX16         *uint32                                 `protobuf:"varint,10,opt,name=avg_gpm_x16" json:"avg_gpm_x16,omitempty"`
	AvgXpmX16         *uint32                                 `protobuf:"varint,11,opt,name=avg_xpm_x16" json:"avg_xpm_x16,omitempty"`
	BestKillsX16      *uint32                                 `protobuf:"varint,12,opt,name=best_kills_x16" json:"best_kills_x16,omitempty"`
	BestAssistsX16    *uint32                                 `protobuf:"varint,13,opt,name=best_assists_x16" json:"best_assists_x16,omitempty"`
	BestGpmX16        *uint32                                 `protobuf:"varint,14,opt,name=best_gpm_x16" json:"best_gpm_x16,omitempty"`
	BestXpmX16        *uint32                                 `protobuf:"varint,15,opt,name=best_xpm_x16" json:"best_xpm_x16,omitempty"`
	WinStreak         *uint32                                 `protobuf:"varint,16,opt,name=win_streak" json:"win_streak,omitempty"`
	BestWinStreak     *uint32                                 `protobuf:"varint,17,opt,name=best_win_streak" json:"best_win_streak,omitempty"`
	FightScore        *float32                                `protobuf:"fixed32,18,opt,name=fight_score" json:"fight_score,omitempty"`
	FarmScore         *float32                                `protobuf:"fixed32,19,opt,name=farm_score" json:"farm_score,omitempty"`
	SupportScore      *float32                                `protobuf:"fixed32,20,opt,name=support_score" json:"support_score,omitempty"`
	PushScore         *float32                                `protobuf:"fixed32,21,opt,name=push_score" json:"push_score,omitempty"`
	LevelUpTimes      []uint32                                `protobuf:"varint,22,rep,name=level_up_times" json:"level_up_times,omitempty"`
	GraphNetWorth     []float32                               `protobuf:"fixed32,23,rep,name=graph_net_worth" json:"graph_net_worth,omitempty"`
	XXX_unrecognized  []byte                                  `json:"-"`
}

func (m *CDOTAMatchMetadata_Team_Player) Reset()         { *m = CDOTAMatchMetadata_Team_Player{} }
func (m *CDOTAMatchMetadata_Team_Player) String() string { return proto.CompactTextString(m) }
func (*CDOTAMatchMetadata_Team_Player) ProtoMessage()    {}
func (*CDOTAMatchMetadata_Team_Player) Descriptor() ([]byte, []int) {
	return fileDescriptor21, []int{1, 0, 2}
}

func (m *CDOTAMatchMetadata_Team_Player) GetAccountId() uint32 {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetAbilityUpgrades() []uint32 {
	if m != nil {
		return m.AbilityUpgrades
	}
	return nil
}

func (m *CDOTAMatchMetadata_Team_Player) GetPlayerSlot() uint32 {
	if m != nil && m.PlayerSlot != nil {
		return *m.PlayerSlot
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetEquippedEconItems() []*CSOEconItem {
	if m != nil {
		return m.EquippedEconItems
	}
	return nil
}

func (m *CDOTAMatchMetadata_Team_Player) GetKills() []*CDOTAMatchMetadata_Team_PlayerKill {
	if m != nil {
		return m.Kills
	}
	return nil
}

func (m *CDOTAMatchMetadata_Team_Player) GetItems() []*CDOTAMatchMetadata_Team_ItemPurchase {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *CDOTAMatchMetadata_Team_Player) GetAvgKillsX16() uint32 {
	if m != nil && m.AvgKillsX16 != nil {
		return *m.AvgKillsX16
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetAvgDeathsX16() uint32 {
	if m != nil && m.AvgDeathsX16 != nil {
		return *m.AvgDeathsX16
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetAvgAssistsX16() uint32 {
	if m != nil && m.AvgAssistsX16 != nil {
		return *m.AvgAssistsX16
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetAvgGpmX16() uint32 {
	if m != nil && m.AvgGpmX16 != nil {
		return *m.AvgGpmX16
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetAvgXpmX16() uint32 {
	if m != nil && m.AvgXpmX16 != nil {
		return *m.AvgXpmX16
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetBestKillsX16() uint32 {
	if m != nil && m.BestKillsX16 != nil {
		return *m.BestKillsX16
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetBestAssistsX16() uint32 {
	if m != nil && m.BestAssistsX16 != nil {
		return *m.BestAssistsX16
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetBestGpmX16() uint32 {
	if m != nil && m.BestGpmX16 != nil {
		return *m.BestGpmX16
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetBestXpmX16() uint32 {
	if m != nil && m.BestXpmX16 != nil {
		return *m.BestXpmX16
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetWinStreak() uint32 {
	if m != nil && m.WinStreak != nil {
		return *m.WinStreak
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetBestWinStreak() uint32 {
	if m != nil && m.BestWinStreak != nil {
		return *m.BestWinStreak
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetFightScore() float32 {
	if m != nil && m.FightScore != nil {
		return *m.FightScore
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetFarmScore() float32 {
	if m != nil && m.FarmScore != nil {
		return *m.FarmScore
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetSupportScore() float32 {
	if m != nil && m.SupportScore != nil {
		return *m.SupportScore
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetPushScore() float32 {
	if m != nil && m.PushScore != nil {
		return *m.PushScore
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetLevelUpTimes() []uint32 {
	if m != nil {
		return m.LevelUpTimes
	}
	return nil
}

func (m *CDOTAMatchMetadata_Team_Player) GetGraphNetWorth() []float32 {
	if m != nil {
		return m.GraphNetWorth
	}
	return nil
}

type CDOTAMatchPrivateMetadata struct {
	Teams            []*CDOTAMatchPrivateMetadata_Team `protobuf:"bytes,1,rep,name=teams" json:"teams,omitempty"`
	XXX_unrecognized []byte                            `json:"-"`
}

func (m *CDOTAMatchPrivateMetadata) Reset()                    { *m = CDOTAMatchPrivateMetadata{} }
func (m *CDOTAMatchPrivateMetadata) String() string            { return proto.CompactTextString(m) }
func (*CDOTAMatchPrivateMetadata) ProtoMessage()               {}
func (*CDOTAMatchPrivateMetadata) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{2} }

func (m *CDOTAMatchPrivateMetadata) GetTeams() []*CDOTAMatchPrivateMetadata_Team {
	if m != nil {
		return m.Teams
	}
	return nil
}

type CDOTAMatchPrivateMetadata_Team struct {
	DotaTeam         *uint32                                    `protobuf:"varint,1,opt,name=dota_team" json:"dota_team,omitempty"`
	Players          []*CDOTAMatchPrivateMetadata_Team_Player   `protobuf:"bytes,2,rep,name=players" json:"players,omitempty"`
	Buildings        []*CDOTAMatchPrivateMetadata_Team_Building `protobuf:"bytes,3,rep,name=buildings" json:"buildings,omitempty"`
	XXX_unrecognized []byte                                     `json:"-"`
}

func (m *CDOTAMatchPrivateMetadata_Team) Reset()         { *m = CDOTAMatchPrivateMetadata_Team{} }
func (m *CDOTAMatchPrivateMetadata_Team) String() string { return proto.CompactTextString(m) }
func (*CDOTAMatchPrivateMetadata_Team) ProtoMessage()    {}
func (*CDOTAMatchPrivateMetadata_Team) Descriptor() ([]byte, []int) {
	return fileDescriptor21, []int{2, 0}
}

func (m *CDOTAMatchPrivateMetadata_Team) GetDotaTeam() uint32 {
	if m != nil && m.DotaTeam != nil {
		return *m.DotaTeam
	}
	return 0
}

func (m *CDOTAMatchPrivateMetadata_Team) GetPlayers() []*CDOTAMatchPrivateMetadata_Team_Player {
	if m != nil {
		return m.Players
	}
	return nil
}

func (m *CDOTAMatchPrivateMetadata_Team) GetBuildings() []*CDOTAMatchPrivateMetadata_Team_Building {
	if m != nil {
		return m.Buildings
	}
	return nil
}

type CDOTAMatchPrivateMetadata_Team_Player struct {
	AccountId        *uint32 `protobuf:"varint,1,opt,name=account_id" json:"account_id,omitempty"`
	PlayerSlot       *uint32 `protobuf:"varint,2,opt,name=player_slot" json:"player_slot,omitempty"`
	PositionStream   []byte  `protobuf:"bytes,3,opt,name=position_stream" json:"position_stream,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CDOTAMatchPrivateMetadata_Team_Player) Reset()         { *m = CDOTAMatchPrivateMetadata_Team_Player{} }
func (m *CDOTAMatchPrivateMetadata_Team_Player) String() string { return proto.CompactTextString(m) }
func (*CDOTAMatchPrivateMetadata_Team_Player) ProtoMessage()    {}
func (*CDOTAMatchPrivateMetadata_Team_Player) Descriptor() ([]byte, []int) {
	return fileDescriptor21, []int{2, 0, 0}
}

func (m *CDOTAMatchPrivateMetadata_Team_Player) GetAccountId() uint32 {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return 0
}

func (m *CDOTAMatchPrivateMetadata_Team_Player) GetPlayerSlot() uint32 {
	if m != nil && m.PlayerSlot != nil {
		return *m.PlayerSlot
	}
	return 0
}

func (m *CDOTAMatchPrivateMetadata_Team_Player) GetPositionStream() []byte {
	if m != nil {
		return m.PositionStream
	}
	return nil
}

type CDOTAMatchPrivateMetadata_Team_Building struct {
	UnitName         *string  `protobuf:"bytes,1,opt,name=unit_name" json:"unit_name,omitempty"`
	PositionQuantX   *uint32  `protobuf:"varint,2,opt,name=position_quant_x" json:"position_quant_x,omitempty"`
	PositionQuantY   *uint32  `protobuf:"varint,3,opt,name=position_quant_y" json:"position_quant_y,omitempty"`
	DeathTime        *float32 `protobuf:"fixed32,4,opt,name=death_time" json:"death_time,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *CDOTAMatchPrivateMetadata_Team_Building) Reset() {
	*m = CDOTAMatchPrivateMetadata_Team_Building{}
}
func (m *CDOTAMatchPrivateMetadata_Team_Building) String() string { return proto.CompactTextString(m) }
func (*CDOTAMatchPrivateMetadata_Team_Building) ProtoMessage()    {}
func (*CDOTAMatchPrivateMetadata_Team_Building) Descriptor() ([]byte, []int) {
	return fileDescriptor21, []int{2, 0, 1}
}

func (m *CDOTAMatchPrivateMetadata_Team_Building) GetUnitName() string {
	if m != nil && m.UnitName != nil {
		return *m.UnitName
	}
	return ""
}

func (m *CDOTAMatchPrivateMetadata_Team_Building) GetPositionQuantX() uint32 {
	if m != nil && m.PositionQuantX != nil {
		return *m.PositionQuantX
	}
	return 0
}

func (m *CDOTAMatchPrivateMetadata_Team_Building) GetPositionQuantY() uint32 {
	if m != nil && m.PositionQuantY != nil {
		return *m.PositionQuantY
	}
	return 0
}

func (m *CDOTAMatchPrivateMetadata_Team_Building) GetDeathTime() float32 {
	if m != nil && m.DeathTime != nil {
		return *m.DeathTime
	}
	return 0
}

func init() {
	proto.RegisterType((*CDOTAMatchMetadataFile)(nil), "dota.CDOTAMatchMetadataFile")
	proto.RegisterType((*CDOTAMatchMetadata)(nil), "dota.CDOTAMatchMetadata")
	proto.RegisterType((*CDOTAMatchMetadata_Team)(nil), "dota.CDOTAMatchMetadata.Team")
	proto.RegisterType((*CDOTAMatchMetadata_Team_PlayerKill)(nil), "dota.CDOTAMatchMetadata.Team.PlayerKill")
	proto.RegisterType((*CDOTAMatchMetadata_Team_ItemPurchase)(nil), "dota.CDOTAMatchMetadata.Team.ItemPurchase")
	proto.RegisterType((*CDOTAMatchMetadata_Team_Player)(nil), "dota.CDOTAMatchMetadata.Team.Player")
	proto.RegisterType((*CDOTAMatchPrivateMetadata)(nil), "dota.CDOTAMatchPrivateMetadata")
	proto.RegisterType((*CDOTAMatchPrivateMetadata_Team)(nil), "dota.CDOTAMatchPrivateMetadata.Team")
	proto.RegisterType((*CDOTAMatchPrivateMetadata_Team_Player)(nil), "dota.CDOTAMatchPrivateMetadata.Team.Player")
	proto.RegisterType((*CDOTAMatchPrivateMetadata_Team_Building)(nil), "dota.CDOTAMatchPrivateMetadata.Team.Building")
}

var fileDescriptor21 = []byte{
	// 873 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x55, 0xdb, 0x6e, 0x23, 0x45,
	0x10, 0xc5, 0xb7, 0xc4, 0x2e, 0xdb, 0x9b, 0x64, 0x72, 0xd9, 0x89, 0x05, 0x68, 0x15, 0xf1, 0x60,
	0x2d, 0xac, 0x05, 0x01, 0x16, 0x21, 0xf1, 0x00, 0xcb, 0x82, 0xc4, 0x65, 0xb5, 0x11, 0xe4, 0xbd,
	0xd5, 0x9e, 0xe9, 0x8c, 0x5b, 0x99, 0xdb, 0x4e, 0xf7, 0x38, 0xf1, 0x1b, 0xfc, 0x01, 0x7f, 0xc0,
	0x0f, 0xf1, 0x21, 0xbc, 0xf2, 0x07, 0x54, 0x55, 0xf7, 0xac, 0x73, 0xd9, 0x44, 0x79, 0x3d, 0xa7,
	0xab, 0xea, 0xd4, 0xa9, 0xae, 0x6e, 0x38, 0x8c, 0x0b, 0x2b, 0x45, 0x26, 0x6d, 0xb4, 0x10, 0x99,
	0xb2, 0x32, 0x96, 0x56, 0xce, 0xca, 0xaa, 0xb0, 0x45, 0xd0, 0x25, 0x6a, 0xb2, 0x3f, 0x97, 0x46,
	0x89, 0x24, 0xca, 0x94, 0x31, 0x32, 0x51, 0xc6, 0x91, 0x93, 0x67, 0x1c, 0xb7, 0x86, 0x45, 0x54,
	0x64, 0x59, 0x91, 0x37, 0x99, 0x64, 0x8e, 0x60, 0xa6, 0x72, 0xeb, 0x8e, 0x1f, 0xfd, 0xd9, 0x82,
	0x83, 0xef, 0x5f, 0xbe, 0x3e, 0xfd, 0xee, 0x15, 0xf1, 0xaf, 0x7c, 0xa1, 0x1f, 0x75, 0xaa, 0x82,
	0x2d, 0xd8, 0x5c, 0xaa, 0xca, 0xe8, 0x22, 0x0f, 0x5b, 0x4f, 0xda, 0xd3, 0x5e, 0xb0, 0x0d, 0x7d,
	0x97, 0x45, 0xc7, 0x61, 0x1b, 0x91, 0x6e, 0xf0, 0x14, 0x11, 0x1f, 0x12, 0x76, 0x9e, 0xb4, 0xa6,
	0xc3, 0xe3, 0x70, 0x46, 0xf5, 0x67, 0xb7, 0x53, 0x06, 0x21, 0x6c, 0x97, 0x95, 0x5e, 0x4a, 0xab,
	0xde, 0xf6, 0x13, 0xf6, 0x30, 0x66, 0x74, 0xf4, 0xd7, 0x00, 0x82, 0x77, 0x04, 0x7c, 0x02, 0x3d,
	0xab, 0x64, 0x66, 0xb0, 0x7a, 0x07, 0x33, 0x7f, 0x70, 0x57, 0xe6, 0xd9, 0x29, 0x9e, 0x0a, 0xbe,
	0x80, 0x91, 0xb6, 0x2a, 0x13, 0x95, 0xba, 0x90, 0x55, 0x6c, 0x50, 0x20, 0x05, 0x7d, 0xe8, 0x83,
	0x7e, 0x2d, 0xe6, 0xf3, 0xd5, 0xa9, 0xce, 0x54, 0xfc, 0x1b, 0xf3, 0x2f, 0x31, 0x56, 0xa7, 0x86,
	0x5a, 0x4a, 0x89, 0xa2, 0x96, 0xa8, 0x81, 0x8d, 0xe0, 0x10, 0x76, 0x2a, 0x55, 0x16, 0x95, 0x15,
	0x75, 0x6e, 0x75, 0x2a, 0x2c, 0x06, 0x85, 0x5d, 0xa2, 0x26, 0xff, 0x6e, 0x42, 0x97, 0x6b, 0xed,
	0xc0, 0x80, 0x5d, 0x26, 0x79, 0xa8, 0xae, 0x35, 0x1d, 0x07, 0x5f, 0xc2, 0x66, 0x99, 0xca, 0x15,
	0xfa, 0xe5, 0x2b, 0x7f, 0x74, 0xaf, 0xdc, 0xd9, 0x09, 0x1f, 0x26, 0x53, 0x92, 0x4a, 0x96, 0x0b,
	0xa1, 0x2e, 0x4b, 0x55, 0x69, 0x95, 0x47, 0x0a, 0x75, 0x74, 0xa6, 0x6d, 0xd2, 0xe1, 0x98, 0xa4,
	0x48, 0x63, 0xa1, 0x64, 0x95, 0xab, 0x18, 0x75, 0x10, 0xf5, 0x18, 0xb6, 0x1c, 0x95, 0x2b, 0x2b,
	0x2e, 0x50, 0xea, 0x02, 0x8d, 0x24, 0x62, 0x1f, 0xc6, 0x51, 0x26, 0xce, 0x74, 0x65, 0xac, 0x28,
	0x75, 0x74, 0x1e, 0x6e, 0xa0, 0xb6, 0x7e, 0xf0, 0x3e, 0xec, 0x21, 0x1c, 0xc9, 0x12, 0x7b, 0xce,
	0x85, 0x93, 0x49, 0x0d, 0x6f, 0xb2, 0x72, 0x1c, 0x33, 0xb2, 0x73, 0x99, 0x9b, 0xb0, 0x8f, 0x59,
	0xc6, 0xe4, 0x09, 0x02, 0x14, 0x6f, 0xc2, 0x01, 0x23, 0x01, 0x00, 0x21, 0x2a, 0x97, 0xa9, 0x5d,
	0x85, 0x40, 0x61, 0x93, 0x4f, 0x01, 0x5c, 0x0f, 0xbf, 0xe8, 0x34, 0x0d, 0x76, 0x61, 0xb8, 0xd4,
	0x11, 0x7a, 0x25, 0x4c, 0x5a, 0x58, 0xef, 0xc9, 0x18, 0x7a, 0x51, 0x81, 0x2e, 0xa2, 0x23, 0x14,
	0xf1, 0x1c, 0x46, 0x3f, 0xe1, 0x84, 0x4e, 0xea, 0x2a, 0x5a, 0xe0, 0xd5, 0xa5, 0xc2, 0x3c, 0x31,
	0x54, 0xe2, 0xce, 0xa3, 0xfc, 0xd2, 0x93, 0xce, 0x76, 0x8a, 0xeb, 0x4d, 0xfe, 0xeb, 0xc2, 0x86,
	0xb7, 0x0b, 0x85, 0xc8, 0x88, 0x73, 0xae, 0xa3, 0xd0, 0x42, 0x39, 0xd7, 0xa9, 0xb6, 0x2b, 0x51,
	0x97, 0xe8, 0x4b, 0xac, 0xdc, 0x08, 0xc6, 0x24, 0xca, 0x37, 0xcb, 0xa2, 0x3a, 0x7c, 0x7c, 0x06,
	0xbb, 0xea, 0x4d, 0xad, 0xcb, 0x52, 0xa1, 0xab, 0x11, 0x2e, 0x06, 0x69, 0x30, 0xec, 0xec, 0xf0,
	0x78, 0xc7, 0x0f, 0xed, 0xf7, 0xd7, 0x3f, 0x20, 0x45, 0x6a, 0x83, 0xaf, 0xa0, 0x77, 0x8e, 0x1d,
	0x1a, 0xb6, 0x78, 0x78, 0x3c, 0x7d, 0xc8, 0x58, 0xd9, 0x92, 0xaf, 0xa1, 0xe7, 0x52, 0x6f, 0x70,
	0xe0, 0xd3, 0xfb, 0x03, 0xaf, 0x39, 0x83, 0x46, 0xc8, 0x65, 0x22, 0xb8, 0xae, 0xb8, 0xfc, 0xec,
	0xb9, 0x9f, 0xd4, 0x01, 0x3c, 0x22, 0x38, 0x56, 0xd2, 0x2e, 0x1c, 0xde, 0x67, 0x1c, 0xef, 0x03,
	0xe1, 0xd2, 0x18, 0x6d, 0xac, 0x23, 0x06, 0x4c, 0xa0, 0x01, 0x44, 0x24, 0x65, 0xc6, 0x20, 0x5c,
	0x05, 0x2f, 0x3d, 0x38, 0x6c, 0x52, 0xcf, 0x15, 0xde, 0x9a, 0x75, 0xc9, 0x51, 0x63, 0x2e, 0xe3,
	0x57, 0x73, 0x8f, 0x99, 0xd9, 0x83, 0x11, 0x33, 0x4d, 0xf2, 0x47, 0xd7, 0xd0, 0x26, 0xfb, 0x16,
	0xa3, 0x38, 0xb6, 0x0b, 0xbc, 0x79, 0xc6, 0x56, 0x4a, 0x9e, 0x87, 0xdb, 0x8d, 0x68, 0x3e, 0x79,
	0x85, 0xd8, 0x69, 0xf4, 0x9d, 0xe9, 0x64, 0x61, 0x85, 0x89, 0x8a, 0x4a, 0x85, 0x01, 0x82, 0x6d,
	0xca, 0x70, 0x26, 0xab, 0xcc, 0x63, 0xbb, 0x8c, 0xa1, 0x4b, 0xa6, 0x2e, 0x79, 0x55, 0x1d, 0xbc,
	0xd7, 0x1c, 0x2d, 0x6b, 0xb3, 0xf0, 0xd8, 0x3e, 0x63, 0xd8, 0x5e, 0xaa, 0x96, 0x2a, 0xc5, 0x1b,
	0xc2, 0x37, 0xcb, 0x84, 0x07, 0x7c, 0x43, 0xde, 0xb1, 0x49, 0x8f, 0x69, 0x93, 0x8e, 0xfe, 0xee,
	0xc0, 0xe1, 0x7a, 0x54, 0x27, 0xee, 0xdd, 0x7a, 0xfb, 0x32, 0x7d, 0x7e, 0xfd, 0x65, 0xba, 0xb5,
	0xea, 0x37, 0xce, 0xf3, 0x84, 0x27, 0xff, 0xb4, 0xef, 0x7e, 0x3d, 0xbe, 0xb9, 0xf9, 0x7a, 0x7c,
	0xfc, 0x90, 0x94, 0xcd, 0x23, 0xf2, 0x2d, 0x0c, 0xe6, 0xb5, 0x4e, 0x63, 0x9d, 0x27, 0x86, 0x5f,
	0x8f, 0xe1, 0xf1, 0xb3, 0x07, 0xc5, 0xbf, 0xf0, 0x51, 0x93, 0x9f, 0xef, 0xdd, 0xb0, 0x1b, 0x7b,
	0xd4, 0x6e, 0xe6, 0x57, 0x16, 0x46, 0x5b, 0xfc, 0x1e, 0xdc, 0xfc, 0x32, 0x5e, 0xb0, 0xd1, 0x44,
	0x41, 0xbf, 0xc9, 0x4b, 0xad, 0xd6, 0xb9, 0xb6, 0x22, 0x97, 0xb8, 0xcd, 0x94, 0x6c, 0xc0, 0xdf,
	0x40, 0x13, 0xf7, 0xa6, 0x96, 0x58, 0xe7, 0xd2, 0x67, 0xbc, 0xcd, 0xac, 0xfc, 0xce, 0xa2, 0x28,
	0xbe, 0xf4, 0xeb, 0xc7, 0xb8, 0xfd, 0xa2, 0xf3, 0x47, 0xeb, 0xbd, 0xff, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x94, 0x86, 0xa9, 0xe7, 0x25, 0x07, 0x00, 0x00,
}
