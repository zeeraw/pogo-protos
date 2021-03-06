// Code generated by protoc-gen-go.
// source: data_gym.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

// Ignoring public import of FortSummary from map_fort.proto

// Ignoring public import of FortLureInfo from map_fort.proto

// Ignoring public import of FortModifier from map_fort.proto

// Ignoring public import of FortData from map_fort.proto

// Ignoring public import of FortRenderingType from map_fort.proto

// Ignoring public import of FortSponsor from map_fort.proto

// Ignoring public import of FortType from map_fort.proto

// Ignoring public import of PokedexEntry from data.proto

// Ignoring public import of PlayerData from data.proto

// Ignoring public import of DownloadUrlEntry from data.proto

// Ignoring public import of AssetDigestEntry from data.proto

// Ignoring public import of PlayerBadge from data.proto

// Ignoring public import of PokemonData from data.proto

// Ignoring public import of BuddyPokemon from data.proto

// Ignoring public import of ContactSettings from data_player.proto

// Ignoring public import of PlayerAvatar from data_player.proto

// Ignoring public import of Currency from data_player.proto

// Ignoring public import of EquippedBadge from data_player.proto

// Ignoring public import of PlayerStats from data_player.proto

// Ignoring public import of PlayerCamera from data_player.proto

// Ignoring public import of PlayerCurrency from data_player.proto

// Ignoring public import of PlayerPublicProfile from data_player.proto

// Ignoring public import of DailyBonus from data_player.proto

type GymState struct {
	FortData    *FortData        `protobuf:"bytes,1,opt,name=fort_data" json:"fort_data,omitempty"`
	Memberships []*GymMembership `protobuf:"bytes,2,rep,name=memberships" json:"memberships,omitempty"`
}

func (m *GymState) Reset()         { *m = GymState{} }
func (m *GymState) String() string { return proto.CompactTextString(m) }
func (*GymState) ProtoMessage()    {}

func (m *GymState) GetFortData() *FortData {
	if m != nil {
		return m.FortData
	}
	return nil
}

func (m *GymState) GetMemberships() []*GymMembership {
	if m != nil {
		return m.Memberships
	}
	return nil
}

type GymMembership struct {
	PokemonData          *PokemonData         `protobuf:"bytes,1,opt,name=pokemon_data" json:"pokemon_data,omitempty"`
	TrainerPublicProfile *PlayerPublicProfile `protobuf:"bytes,2,opt,name=trainer_public_profile" json:"trainer_public_profile,omitempty"`
}

func (m *GymMembership) Reset()         { *m = GymMembership{} }
func (m *GymMembership) String() string { return proto.CompactTextString(m) }
func (*GymMembership) ProtoMessage()    {}

func (m *GymMembership) GetPokemonData() *PokemonData {
	if m != nil {
		return m.PokemonData
	}
	return nil
}

func (m *GymMembership) GetTrainerPublicProfile() *PlayerPublicProfile {
	if m != nil {
		return m.TrainerPublicProfile
	}
	return nil
}

func init() {
}
