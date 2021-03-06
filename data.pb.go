// Code generated by protoc-gen-go.
// source: data.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

// Ignoring public import of BadgeType from enums.proto

// Ignoring public import of TutorialState from enums.proto

// Ignoring public import of PokemonId from enums.proto

// Ignoring public import of PokemonMovementType from enums.proto

// Ignoring public import of Gender from enums.proto

// Ignoring public import of ActivityType from enums.proto

// Ignoring public import of CameraTarget from enums.proto

// Ignoring public import of PokemonFamilyId from enums.proto

// Ignoring public import of ItemCategory from enums.proto

// Ignoring public import of CameraInterpolation from enums.proto

// Ignoring public import of Platform from enums.proto

// Ignoring public import of TeamColor from enums.proto

// Ignoring public import of PokemonMove from enums.proto

// Ignoring public import of ItemEffect from enums.proto

// Ignoring public import of PokemonType from enums.proto

// Ignoring public import of HoloIapItemCategory from enums.proto

// Ignoring public import of PokemonRarity from enums.proto

// Ignoring public import of ContactSettings from data_player.proto

// Ignoring public import of PlayerAvatar from data_player.proto

// Ignoring public import of Currency from data_player.proto

// Ignoring public import of EquippedBadge from data_player.proto

// Ignoring public import of PlayerStats from data_player.proto

// Ignoring public import of PlayerCamera from data_player.proto

// Ignoring public import of PlayerCurrency from data_player.proto

// Ignoring public import of PlayerPublicProfile from data_player.proto

// Ignoring public import of DailyBonus from data_player.proto

// Ignoring public import of ItemData from inventory_item.proto

// Ignoring public import of ItemAward from inventory_item.proto

// Ignoring public import of ItemId from inventory_item.proto

// Ignoring public import of ItemType from inventory_item.proto

type PokedexEntry struct {
	PokemonId            PokemonId `protobuf:"varint,1,opt,name=pokemon_id,enum=POGOProtos.Enums.PokemonId" json:"pokemon_id,omitempty"`
	TimesEncountered     int32     `protobuf:"varint,2,opt,name=times_encountered" json:"times_encountered,omitempty"`
	TimesCaptured        int32     `protobuf:"varint,3,opt,name=times_captured" json:"times_captured,omitempty"`
	EvolutionStonePieces int32     `protobuf:"varint,4,opt,name=evolution_stone_pieces" json:"evolution_stone_pieces,omitempty"`
	EvolutionStones      int32     `protobuf:"varint,5,opt,name=evolution_stones" json:"evolution_stones,omitempty"`
}

func (m *PokedexEntry) Reset()         { *m = PokedexEntry{} }
func (m *PokedexEntry) String() string { return proto.CompactTextString(m) }
func (*PokedexEntry) ProtoMessage()    {}

type PlayerData struct {
	CreationTimestampMs     int64            `protobuf:"varint,1,opt,name=creation_timestamp_ms" json:"creation_timestamp_ms,omitempty"`
	Username                string           `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Team                    TeamColor        `protobuf:"varint,5,opt,name=team,enum=POGOProtos.Enums.TeamColor" json:"team,omitempty"`
	TutorialState           []TutorialState  `protobuf:"varint,7,rep,packed,name=tutorial_state,enum=POGOProtos.Enums.TutorialState" json:"tutorial_state,omitempty"`
	Avatar                  *PlayerAvatar    `protobuf:"bytes,8,opt,name=avatar" json:"avatar,omitempty"`
	MaxPokemonStorage       int32            `protobuf:"varint,9,opt,name=max_pokemon_storage" json:"max_pokemon_storage,omitempty"`
	MaxItemStorage          int32            `protobuf:"varint,10,opt,name=max_item_storage" json:"max_item_storage,omitempty"`
	DailyBonus              *DailyBonus      `protobuf:"bytes,11,opt,name=daily_bonus" json:"daily_bonus,omitempty"`
	EquippedBadge           *EquippedBadge   `protobuf:"bytes,12,opt,name=equipped_badge" json:"equipped_badge,omitempty"`
	ContactSettings         *ContactSettings `protobuf:"bytes,13,opt,name=contact_settings" json:"contact_settings,omitempty"`
	Currencies              []*Currency      `protobuf:"bytes,14,rep,name=currencies" json:"currencies,omitempty"`
	RemainingCodenameClaims int32            `protobuf:"varint,15,opt,name=remaining_codename_claims" json:"remaining_codename_claims,omitempty"`
	BuddyPokemon            *BuddyPokemon    `protobuf:"bytes,16,opt,name=buddy_pokemon" json:"buddy_pokemon,omitempty"`
}

func (m *PlayerData) Reset()         { *m = PlayerData{} }
func (m *PlayerData) String() string { return proto.CompactTextString(m) }
func (*PlayerData) ProtoMessage()    {}

func (m *PlayerData) GetAvatar() *PlayerAvatar {
	if m != nil {
		return m.Avatar
	}
	return nil
}

func (m *PlayerData) GetDailyBonus() *DailyBonus {
	if m != nil {
		return m.DailyBonus
	}
	return nil
}

func (m *PlayerData) GetEquippedBadge() *EquippedBadge {
	if m != nil {
		return m.EquippedBadge
	}
	return nil
}

func (m *PlayerData) GetContactSettings() *ContactSettings {
	if m != nil {
		return m.ContactSettings
	}
	return nil
}

func (m *PlayerData) GetCurrencies() []*Currency {
	if m != nil {
		return m.Currencies
	}
	return nil
}

func (m *PlayerData) GetBuddyPokemon() *BuddyPokemon {
	if m != nil {
		return m.BuddyPokemon
	}
	return nil
}

type DownloadUrlEntry struct {
	AssetId  string `protobuf:"bytes,1,opt,name=asset_id" json:"asset_id,omitempty"`
	Url      string `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
	Size     int32  `protobuf:"varint,3,opt,name=size" json:"size,omitempty"`
	Checksum uint32 `protobuf:"fixed32,4,opt,name=checksum" json:"checksum,omitempty"`
}

func (m *DownloadUrlEntry) Reset()         { *m = DownloadUrlEntry{} }
func (m *DownloadUrlEntry) String() string { return proto.CompactTextString(m) }
func (*DownloadUrlEntry) ProtoMessage()    {}

type AssetDigestEntry struct {
	AssetId    string `protobuf:"bytes,1,opt,name=asset_id" json:"asset_id,omitempty"`
	BundleName string `protobuf:"bytes,2,opt,name=bundle_name" json:"bundle_name,omitempty"`
	Version    int64  `protobuf:"varint,3,opt,name=version" json:"version,omitempty"`
	Checksum   uint32 `protobuf:"fixed32,4,opt,name=checksum" json:"checksum,omitempty"`
	Size       int32  `protobuf:"varint,5,opt,name=size" json:"size,omitempty"`
	Key        []byte `protobuf:"bytes,6,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *AssetDigestEntry) Reset()         { *m = AssetDigestEntry{} }
func (m *AssetDigestEntry) String() string { return proto.CompactTextString(m) }
func (*AssetDigestEntry) ProtoMessage()    {}

type PlayerBadge struct {
	BadgeType    BadgeType `protobuf:"varint,1,opt,name=badge_type,enum=POGOProtos.Enums.BadgeType" json:"badge_type,omitempty"`
	Rank         int32     `protobuf:"varint,2,opt,name=rank" json:"rank,omitempty"`
	StartValue   int32     `protobuf:"varint,3,opt,name=start_value" json:"start_value,omitempty"`
	EndValue     int32     `protobuf:"varint,4,opt,name=end_value" json:"end_value,omitempty"`
	CurrentValue float64   `protobuf:"fixed64,5,opt,name=current_value" json:"current_value,omitempty"`
}

func (m *PlayerBadge) Reset()         { *m = PlayerBadge{} }
func (m *PlayerBadge) String() string { return proto.CompactTextString(m) }
func (*PlayerBadge) ProtoMessage()    {}

type PokemonData struct {
	Id                     uint64      `protobuf:"fixed64,1,opt,name=id" json:"id,omitempty"`
	PokemonId              PokemonId   `protobuf:"varint,2,opt,name=pokemon_id,enum=POGOProtos.Enums.PokemonId" json:"pokemon_id,omitempty"`
	Cp                     int32       `protobuf:"varint,3,opt,name=cp" json:"cp,omitempty"`
	Stamina                int32       `protobuf:"varint,4,opt,name=stamina" json:"stamina,omitempty"`
	StaminaMax             int32       `protobuf:"varint,5,opt,name=stamina_max" json:"stamina_max,omitempty"`
	Move_1                 PokemonMove `protobuf:"varint,6,opt,name=move_1,enum=POGOProtos.Enums.PokemonMove" json:"move_1,omitempty"`
	Move_2                 PokemonMove `protobuf:"varint,7,opt,name=move_2,enum=POGOProtos.Enums.PokemonMove" json:"move_2,omitempty"`
	DeployedFortId         string      `protobuf:"bytes,8,opt,name=deployed_fort_id" json:"deployed_fort_id,omitempty"`
	OwnerName              string      `protobuf:"bytes,9,opt,name=owner_name" json:"owner_name,omitempty"`
	IsEgg                  bool        `protobuf:"varint,10,opt,name=is_egg" json:"is_egg,omitempty"`
	EggKmWalkedTarget      float64     `protobuf:"fixed64,11,opt,name=egg_km_walked_target" json:"egg_km_walked_target,omitempty"`
	EggKmWalkedStart       float64     `protobuf:"fixed64,12,opt,name=egg_km_walked_start" json:"egg_km_walked_start,omitempty"`
	Origin                 int32       `protobuf:"varint,14,opt,name=origin" json:"origin,omitempty"`
	HeightM                float32     `protobuf:"fixed32,15,opt,name=height_m" json:"height_m,omitempty"`
	WeightKg               float32     `protobuf:"fixed32,16,opt,name=weight_kg" json:"weight_kg,omitempty"`
	IndividualAttack       int32       `protobuf:"varint,17,opt,name=individual_attack" json:"individual_attack,omitempty"`
	IndividualDefense      int32       `protobuf:"varint,18,opt,name=individual_defense" json:"individual_defense,omitempty"`
	IndividualStamina      int32       `protobuf:"varint,19,opt,name=individual_stamina" json:"individual_stamina,omitempty"`
	CpMultiplier           float32     `protobuf:"fixed32,20,opt,name=cp_multiplier" json:"cp_multiplier,omitempty"`
	Pokeball               ItemId      `protobuf:"varint,21,opt,name=pokeball,enum=POGOProtos.Inventory.Item.ItemId" json:"pokeball,omitempty"`
	CapturedCellId         uint64      `protobuf:"varint,22,opt,name=captured_cell_id" json:"captured_cell_id,omitempty"`
	BattlesAttacked        int32       `protobuf:"varint,23,opt,name=battles_attacked" json:"battles_attacked,omitempty"`
	BattlesDefended        int32       `protobuf:"varint,24,opt,name=battles_defended" json:"battles_defended,omitempty"`
	EggIncubatorId         string      `protobuf:"bytes,25,opt,name=egg_incubator_id" json:"egg_incubator_id,omitempty"`
	CreationTimeMs         uint64      `protobuf:"varint,26,opt,name=creation_time_ms" json:"creation_time_ms,omitempty"`
	NumUpgrades            int32       `protobuf:"varint,27,opt,name=num_upgrades" json:"num_upgrades,omitempty"`
	AdditionalCpMultiplier float32     `protobuf:"fixed32,28,opt,name=additional_cp_multiplier" json:"additional_cp_multiplier,omitempty"`
	Favorite               int32       `protobuf:"varint,29,opt,name=favorite" json:"favorite,omitempty"`
	Nickname               string      `protobuf:"bytes,30,opt,name=nickname" json:"nickname,omitempty"`
	FromFort               int32       `protobuf:"varint,31,opt,name=from_fort" json:"from_fort,omitempty"`
	BuddyCandyAwarded      int32       `protobuf:"varint,32,opt,name=buddy_candy_awarded" json:"buddy_candy_awarded,omitempty"`
}

func (m *PokemonData) Reset()         { *m = PokemonData{} }
func (m *PokemonData) String() string { return proto.CompactTextString(m) }
func (*PokemonData) ProtoMessage()    {}

type BuddyPokemon struct {
	Id            uint64  `protobuf:"fixed64,1,opt,name=id" json:"id,omitempty"`
	StartKmWalked float64 `protobuf:"fixed64,2,opt,name=start_km_walked" json:"start_km_walked,omitempty"`
	LastKmAwarded float64 `protobuf:"fixed64,3,opt,name=last_km_awarded" json:"last_km_awarded,omitempty"`
}

func (m *BuddyPokemon) Reset()         { *m = BuddyPokemon{} }
func (m *BuddyPokemon) String() string { return proto.CompactTextString(m) }
func (*BuddyPokemon) ProtoMessage()    {}

func init() {
}
