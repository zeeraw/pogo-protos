// Code generated by protoc-gen-go.
// source: settings_master_pokemon.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of ActivityType from enums.proto

// Ignoring public import of BadgeType from enums.proto

// Ignoring public import of CameraInterpolation from enums.proto

// Ignoring public import of CameraTarget from enums.proto

// Ignoring public import of Gender from enums.proto

// Ignoring public import of HoloIapItemCategory from enums.proto

// Ignoring public import of ItemCategory from enums.proto

// Ignoring public import of ItemEffect from enums.proto

// Ignoring public import of Platform from enums.proto

// Ignoring public import of PokemonFamilyId from enums.proto

// Ignoring public import of PokemonId from enums.proto

// Ignoring public import of PokemonMove from enums.proto

// Ignoring public import of PokemonMovementType from enums.proto

// Ignoring public import of PokemonRarity from enums.proto

// Ignoring public import of PokemonType from enums.proto

// Ignoring public import of TeamColor from enums.proto

// Ignoring public import of TutorialState from enums.proto

type CameraAttributes struct {
	DiskRadiusM       float32 `protobuf:"fixed32,1,opt,name=disk_radius_m,json=diskRadiusM" json:"disk_radius_m,omitempty"`
	CylinderRadiusM   float32 `protobuf:"fixed32,2,opt,name=cylinder_radius_m,json=cylinderRadiusM" json:"cylinder_radius_m,omitempty"`
	CylinderHeightM   float32 `protobuf:"fixed32,3,opt,name=cylinder_height_m,json=cylinderHeightM" json:"cylinder_height_m,omitempty"`
	CylinderGroundM   float32 `protobuf:"fixed32,4,opt,name=cylinder_ground_m,json=cylinderGroundM" json:"cylinder_ground_m,omitempty"`
	ShoulderModeScale float32 `protobuf:"fixed32,5,opt,name=shoulder_mode_scale,json=shoulderModeScale" json:"shoulder_mode_scale,omitempty"`
}

func (m *CameraAttributes) Reset()                    { *m = CameraAttributes{} }
func (m *CameraAttributes) String() string            { return proto.CompactTextString(m) }
func (*CameraAttributes) ProtoMessage()               {}
func (*CameraAttributes) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{0} }

type EncounterAttributes struct {
	BaseCaptureRate      float32             `protobuf:"fixed32,1,opt,name=base_capture_rate,json=baseCaptureRate" json:"base_capture_rate,omitempty"`
	BaseFleeRate         float32             `protobuf:"fixed32,2,opt,name=base_flee_rate,json=baseFleeRate" json:"base_flee_rate,omitempty"`
	CollisionRadiusM     float32             `protobuf:"fixed32,3,opt,name=collision_radius_m,json=collisionRadiusM" json:"collision_radius_m,omitempty"`
	CollisionHeightM     float32             `protobuf:"fixed32,4,opt,name=collision_height_m,json=collisionHeightM" json:"collision_height_m,omitempty"`
	CollisionHeadRadiusM float32             `protobuf:"fixed32,5,opt,name=collision_head_radius_m,json=collisionHeadRadiusM" json:"collision_head_radius_m,omitempty"`
	MovementType         PokemonMovementType `protobuf:"varint,6,opt,name=movement_type,json=movementType,enum=POGOProtos.Enums.PokemonMovementType" json:"movement_type,omitempty"`
	MovementTimerS       float32             `protobuf:"fixed32,7,opt,name=movement_timer_s,json=movementTimerS" json:"movement_timer_s,omitempty"`
	JumpTimeS            float32             `protobuf:"fixed32,8,opt,name=jump_time_s,json=jumpTimeS" json:"jump_time_s,omitempty"`
	AttackTimerS         float32             `protobuf:"fixed32,9,opt,name=attack_timer_s,json=attackTimerS" json:"attack_timer_s,omitempty"`
}

func (m *EncounterAttributes) Reset()                    { *m = EncounterAttributes{} }
func (m *EncounterAttributes) String() string            { return proto.CompactTextString(m) }
func (*EncounterAttributes) ProtoMessage()               {}
func (*EncounterAttributes) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{1} }

type StatsAttributes struct {
	BaseStamina      int32 `protobuf:"varint,1,opt,name=base_stamina,json=baseStamina" json:"base_stamina,omitempty"`
	BaseAttack       int32 `protobuf:"varint,2,opt,name=base_attack,json=baseAttack" json:"base_attack,omitempty"`
	BaseDefense      int32 `protobuf:"varint,3,opt,name=base_defense,json=baseDefense" json:"base_defense,omitempty"`
	DodgeEnergyDelta int32 `protobuf:"varint,8,opt,name=dodge_energy_delta,json=dodgeEnergyDelta" json:"dodge_energy_delta,omitempty"`
}

func (m *StatsAttributes) Reset()                    { *m = StatsAttributes{} }
func (m *StatsAttributes) String() string            { return proto.CompactTextString(m) }
func (*StatsAttributes) ProtoMessage()               {}
func (*StatsAttributes) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{2} }

func init() {
	proto.RegisterType((*CameraAttributes)(nil), "POGOProtos.Settings.Master.Pokemon.CameraAttributes")
	proto.RegisterType((*EncounterAttributes)(nil), "POGOProtos.Settings.Master.Pokemon.EncounterAttributes")
	proto.RegisterType((*StatsAttributes)(nil), "POGOProtos.Settings.Master.Pokemon.StatsAttributes")
}

func init() { proto.RegisterFile("settings_master_pokemon.proto", fileDescriptor22) }

var fileDescriptor22 = []byte{
	// 507 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x93, 0xdf, 0x6e, 0xd3, 0x3c,
	0x14, 0xc0, 0xbf, 0x6e, 0xeb, 0xbe, 0xcd, 0xdd, 0xda, 0x2e, 0x43, 0xa2, 0x42, 0xe2, 0x5f, 0x05,
	0xd2, 0x84, 0x50, 0x2e, 0x40, 0x3c, 0xc0, 0xd8, 0xca, 0x10, 0x52, 0xb5, 0x2a, 0xe1, 0x8a, 0x1b,
	0xcb, 0x8d, 0xcf, 0x5a, 0xd3, 0xd8, 0x8e, 0x62, 0x07, 0xa9, 0x2f, 0xc4, 0x23, 0xf0, 0x50, 0x3c,
	0x05, 0xf6, 0x71, 0x92, 0xa6, 0x88, 0xbb, 0xf8, 0x77, 0x7e, 0x3e, 0xa7, 0xe7, 0x1c, 0x97, 0x3c,
	0x35, 0x60, 0xad, 0x50, 0x2b, 0x43, 0x25, 0x33, 0x16, 0x4a, 0x5a, 0xe8, 0x0d, 0x48, 0xad, 0xe2,
	0xa2, 0xd4, 0x56, 0x47, 0xd3, 0xc5, 0xfd, 0xdd, 0xfd, 0xc2, 0x7f, 0x9a, 0x38, 0xad, 0xcd, 0x78,
	0x8e, 0x66, 0xbc, 0x08, 0xe6, 0x93, 0x01, 0xa8, 0x4a, 0x9a, 0x70, 0x61, 0xfa, 0xbb, 0x47, 0xc6,
	0x37, 0x4c, 0x42, 0xc9, 0xae, 0xad, 0x2d, 0xc5, 0xb2, 0xb2, 0x60, 0xa2, 0x29, 0x39, 0xe7, 0xc2,
	0x6c, 0x68, 0xc9, 0xb8, 0xa8, 0x5c, 0xa5, 0x49, 0xef, 0x45, 0xef, 0xea, 0x20, 0x19, 0x78, 0x98,
	0x20, 0x9b, 0x47, 0x6f, 0xc8, 0x45, 0xb6, 0xcd, 0x85, 0xe2, 0xee, 0x37, 0xb4, 0xde, 0x01, 0x7a,
	0xa3, 0x26, 0xf0, 0x2f, 0x77, 0x0d, 0x62, 0xb5, 0xb6, 0xce, 0x3d, 0xdc, 0x77, 0x3f, 0x23, 0xdf,
	0x77, 0x57, 0xa5, 0xae, 0x14, 0x77, 0xee, 0xd1, 0xbe, 0x7b, 0x87, 0x7c, 0x1e, 0xc5, 0xe4, 0xd2,
	0xac, 0x75, 0x95, 0x7b, 0x57, 0x6a, 0x0e, 0xd4, 0x64, 0x2c, 0x87, 0x49, 0x1f, 0xed, 0x8b, 0x26,
	0x34, 0x77, 0x91, 0xd4, 0x07, 0xa6, 0xbf, 0x0e, 0xc9, 0xe5, 0x4c, 0x65, 0xee, 0xb2, 0x9b, 0x47,
	0xa7, 0x5f, 0x57, 0x73, 0xc9, 0x0c, 0xd0, 0x8c, 0x15, 0xb6, 0x2a, 0xc1, 0xf5, 0x63, 0xa1, 0xee,
	0x79, 0xe4, 0x03, 0x37, 0x81, 0x27, 0x0e, 0x47, 0xaf, 0xc8, 0x10, 0xdd, 0x87, 0x1c, 0x6a, 0x31,
	0x34, 0x7d, 0xe6, 0xe9, 0x27, 0x07, 0xd1, 0x7a, 0x4b, 0xa2, 0x4c, 0xe7, 0xb9, 0x30, 0x42, 0xab,
	0xdd, 0x78, 0x42, 0xcb, 0xe3, 0x36, 0xd2, 0xcc, 0x67, 0xcf, 0x6e, 0x07, 0x74, 0xf4, 0x97, 0xdd,
	0x4c, 0xe8, 0x03, 0x79, 0xdc, 0xb5, 0x19, 0xdf, 0x15, 0x08, 0x9d, 0x3f, 0xea, 0x5c, 0x61, 0xbc,
	0x29, 0xf2, 0x85, 0x9c, 0x4b, 0xfd, 0x03, 0x24, 0x28, 0x4b, 0xed, 0xb6, 0x80, 0xc9, 0xb1, 0x93,
	0x87, 0xef, 0x5e, 0xc7, 0x9d, 0x27, 0x33, 0xc3, 0x97, 0x51, 0x3f, 0x94, 0x79, 0x6d, 0x7f, 0x75,
	0x72, 0x72, 0x26, 0x3b, 0xa7, 0xe8, 0x8a, 0x8c, 0x77, 0xb9, 0x84, 0x7b, 0x3d, 0xd4, 0x4c, 0xfe,
	0xc7, 0xda, 0xc3, 0xd6, 0xf3, 0x38, 0x8d, 0x9e, 0x91, 0xc1, 0xf7, 0x4a, 0x16, 0x68, 0x39, 0xe9,
	0x04, 0xa5, 0x53, 0x8f, 0xbc, 0x90, 0xfa, 0x71, 0x32, 0x6b, 0x59, 0xb6, 0x69, 0xf3, 0x9c, 0x86,
	0x71, 0x06, 0x1a, 0xb2, 0x4c, 0x7f, 0xf6, 0xc8, 0x28, 0xb5, 0xcc, 0x9a, 0xce, 0xd2, 0x5e, 0x12,
	0x1c, 0x39, 0x35, 0x96, 0x49, 0xa1, 0x18, 0xee, 0xab, 0x9f, 0x0c, 0x3c, 0x4b, 0x03, 0x8a, 0x9e,
	0x13, 0x3c, 0xd2, 0x90, 0x0b, 0x17, 0xd5, 0x4f, 0x88, 0x47, 0xd7, 0x48, 0xda, 0x1c, 0x1c, 0x1e,
	0x40, 0x19, 0xc0, 0x05, 0xd5, 0x39, 0x6e, 0x03, 0xf2, 0xbb, 0xe1, 0x9a, 0xaf, 0x80, 0x82, 0x82,
	0x72, 0xb5, 0x75, 0x6a, 0x6e, 0x19, 0xf6, 0xd1, 0x4f, 0xc6, 0x18, 0x99, 0x61, 0xe0, 0xd6, 0xf3,
	0x8f, 0x27, 0xdf, 0x8e, 0xf1, 0x7f, 0x65, 0x16, 0xff, 0x2d, 0xc3, 0xd7, 0xfb, 0x3f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x7a, 0x3d, 0x39, 0x16, 0xb3, 0x03, 0x00, 0x00,
}
