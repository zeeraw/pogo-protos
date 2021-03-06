// Code generated by protoc-gen-go.
// source: settings_master.proto
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

// Ignoring public import of CameraAttributes from settings_master_pokemon.proto

// Ignoring public import of EncounterAttributes from settings_master_pokemon.proto

// Ignoring public import of StatsAttributes from settings_master_pokemon.proto

// Ignoring public import of ItemData from inventory_item.proto

// Ignoring public import of ItemAward from inventory_item.proto

// Ignoring public import of ItemId from inventory_item.proto

// Ignoring public import of ItemType from inventory_item.proto

// Ignoring public import of FortModifierAttributes from settings_master_item.proto

// Ignoring public import of ReviveAttributes from settings_master_item.proto

// Ignoring public import of PokeballAttributes from settings_master_item.proto

// Ignoring public import of PotionAttributes from settings_master_item.proto

// Ignoring public import of IncenseAttributes from settings_master_item.proto

// Ignoring public import of InventoryUpgradeAttributes from settings_master_item.proto

// Ignoring public import of ExperienceBoostAttributes from settings_master_item.proto

// Ignoring public import of FoodAttributes from settings_master_item.proto

// Ignoring public import of BattleAttributes from settings_master_item.proto

// Ignoring public import of EggIncubatorAttributes from settings_master_item.proto

type PokemonSettings_BuddySize int32

const (
	PokemonSettings_BUDDY_MEDIUM   PokemonSettings_BuddySize = 0
	PokemonSettings_BUDDY_SHOULDER PokemonSettings_BuddySize = 1
	PokemonSettings_BUDDY_BIG      PokemonSettings_BuddySize = 2
	PokemonSettings_BUDDY_FLYING   PokemonSettings_BuddySize = 3
)

var PokemonSettings_BuddySize_name = map[int32]string{
	0: "BUDDY_MEDIUM",
	1: "BUDDY_SHOULDER",
	2: "BUDDY_BIG",
	3: "BUDDY_FLYING",
}
var PokemonSettings_BuddySize_value = map[string]int32{
	"BUDDY_MEDIUM":   0,
	"BUDDY_SHOULDER": 1,
	"BUDDY_BIG":      2,
	"BUDDY_FLYING":   3,
}

func (x PokemonSettings_BuddySize) String() string {
	return proto.EnumName(PokemonSettings_BuddySize_name, int32(x))
}

type IapSettings struct {
	DailyBonusCoins                int32    `protobuf:"varint,1,opt,name=daily_bonus_coins" json:"daily_bonus_coins,omitempty"`
	DailyDefenderBonusPerPokemon   []int32  `protobuf:"varint,2,rep,name=daily_defender_bonus_per_pokemon" json:"daily_defender_bonus_per_pokemon,omitempty"`
	DailyDefenderBonusMaxDefenders int32    `protobuf:"varint,3,opt,name=daily_defender_bonus_max_defenders" json:"daily_defender_bonus_max_defenders,omitempty"`
	DailyDefenderBonusCurrency     []string `protobuf:"bytes,4,rep,name=daily_defender_bonus_currency" json:"daily_defender_bonus_currency,omitempty"`
	MinTimeBetweenClaimsMs         int64    `protobuf:"varint,5,opt,name=min_time_between_claims_ms" json:"min_time_between_claims_ms,omitempty"`
	DailyBonusEnabled              bool     `protobuf:"varint,6,opt,name=daily_bonus_enabled" json:"daily_bonus_enabled,omitempty"`
	DailyDefenderBonusEnabled      bool     `protobuf:"varint,7,opt,name=daily_defender_bonus_enabled" json:"daily_defender_bonus_enabled,omitempty"`
}

func (m *IapSettings) Reset()         { *m = IapSettings{} }
func (m *IapSettings) String() string { return proto.CompactTextString(m) }
func (*IapSettings) ProtoMessage()    {}

type TypeEffectiveSettings struct {
	AttackScalar []float32   `protobuf:"fixed32,1,rep,name=attack_scalar" json:"attack_scalar,omitempty"`
	AttackType   PokemonType `protobuf:"varint,2,opt,name=attack_type,enum=POGOProtos.Enums.PokemonType" json:"attack_type,omitempty"`
}

func (m *TypeEffectiveSettings) Reset()         { *m = TypeEffectiveSettings{} }
func (m *TypeEffectiveSettings) String() string { return proto.CompactTextString(m) }
func (*TypeEffectiveSettings) ProtoMessage()    {}

type CameraSettings struct {
	NextCamera        string                `protobuf:"bytes,1,opt,name=next_camera" json:"next_camera,omitempty"`
	Interpolation     []CameraInterpolation `protobuf:"varint,2,rep,name=interpolation,enum=POGOProtos.Enums.CameraInterpolation" json:"interpolation,omitempty"`
	TargetType        []CameraTarget        `protobuf:"varint,3,rep,name=target_type,enum=POGOProtos.Enums.CameraTarget" json:"target_type,omitempty"`
	EaseInSpeed       []float32             `protobuf:"fixed32,4,rep,name=ease_in_speed" json:"ease_in_speed,omitempty"`
	EastOutSpeed      []float32             `protobuf:"fixed32,5,rep,name=east_out_speed" json:"east_out_speed,omitempty"`
	DurationSeconds   []float32             `protobuf:"fixed32,6,rep,name=duration_seconds" json:"duration_seconds,omitempty"`
	WaitSeconds       []float32             `protobuf:"fixed32,7,rep,name=wait_seconds" json:"wait_seconds,omitempty"`
	TransitionSeconds []float32             `protobuf:"fixed32,8,rep,name=transition_seconds" json:"transition_seconds,omitempty"`
	AngleDegree       []float32             `protobuf:"fixed32,9,rep,name=angle_degree" json:"angle_degree,omitempty"`
	AngleOffsetDegree []float32             `protobuf:"fixed32,10,rep,name=angle_offset_degree" json:"angle_offset_degree,omitempty"`
	PitchDegree       []float32             `protobuf:"fixed32,11,rep,name=pitch_degree" json:"pitch_degree,omitempty"`
	PitchOffsetDegree []float32             `protobuf:"fixed32,12,rep,name=pitch_offset_degree" json:"pitch_offset_degree,omitempty"`
	RollDegree        []float32             `protobuf:"fixed32,13,rep,name=roll_degree" json:"roll_degree,omitempty"`
	DistanceMeters    []float32             `protobuf:"fixed32,14,rep,name=distance_meters" json:"distance_meters,omitempty"`
	HeightPercent     []float32             `protobuf:"fixed32,15,rep,name=height_percent" json:"height_percent,omitempty"`
	VertCtrRatio      []float32             `protobuf:"fixed32,16,rep,name=vert_ctr_ratio" json:"vert_ctr_ratio,omitempty"`
}

func (m *CameraSettings) Reset()         { *m = CameraSettings{} }
func (m *CameraSettings) String() string { return proto.CompactTextString(m) }
func (*CameraSettings) ProtoMessage()    {}

type EquippedBadgeSettings struct {
	EquipBadgeCooldownMs  int64     `protobuf:"varint,1,opt,name=equip_badge_cooldown_ms" json:"equip_badge_cooldown_ms,omitempty"`
	CatchProbabilityBonus []float32 `protobuf:"fixed32,2,rep,name=catch_probability_bonus" json:"catch_probability_bonus,omitempty"`
	FleeProbabilityBonus  []float32 `protobuf:"fixed32,3,rep,name=flee_probability_bonus" json:"flee_probability_bonus,omitempty"`
}

func (m *EquippedBadgeSettings) Reset()         { *m = EquippedBadgeSettings{} }
func (m *EquippedBadgeSettings) String() string { return proto.CompactTextString(m) }
func (*EquippedBadgeSettings) ProtoMessage()    {}

type PlayerLevelSettings struct {
	RankNum                 []int32   `protobuf:"varint,1,rep,name=rank_num" json:"rank_num,omitempty"`
	RequiredExperience      []int32   `protobuf:"varint,2,rep,name=required_experience" json:"required_experience,omitempty"`
	CpMultiplier            []float32 `protobuf:"fixed32,3,rep,name=cp_multiplier" json:"cp_multiplier,omitempty"`
	MaxEggPlayerLevel       int32     `protobuf:"varint,4,opt,name=max_egg_player_level" json:"max_egg_player_level,omitempty"`
	MaxEncounterPlayerLevel int32     `protobuf:"varint,5,opt,name=max_encounter_player_level" json:"max_encounter_player_level,omitempty"`
}

func (m *PlayerLevelSettings) Reset()         { *m = PlayerLevelSettings{} }
func (m *PlayerLevelSettings) String() string { return proto.CompactTextString(m) }
func (*PlayerLevelSettings) ProtoMessage()    {}

type GymBattleSettings struct {
	EnergyPerSec                  float32 `protobuf:"fixed32,1,opt,name=energy_per_sec" json:"energy_per_sec,omitempty"`
	DodgeEnergyCost               float32 `protobuf:"fixed32,2,opt,name=dodge_energy_cost" json:"dodge_energy_cost,omitempty"`
	RetargetSeconds               float32 `protobuf:"fixed32,3,opt,name=retarget_seconds" json:"retarget_seconds,omitempty"`
	EnemyAttackInterval           float32 `protobuf:"fixed32,4,opt,name=enemy_attack_interval" json:"enemy_attack_interval,omitempty"`
	AttackServerInterval          float32 `protobuf:"fixed32,5,opt,name=attack_server_interval" json:"attack_server_interval,omitempty"`
	RoundDurationSeconds          float32 `protobuf:"fixed32,6,opt,name=round_duration_seconds" json:"round_duration_seconds,omitempty"`
	BonusTimePerAllySeconds       float32 `protobuf:"fixed32,7,opt,name=bonus_time_per_ally_seconds" json:"bonus_time_per_ally_seconds,omitempty"`
	MaximumAttackersPerBattle     int32   `protobuf:"varint,8,opt,name=maximum_attackers_per_battle" json:"maximum_attackers_per_battle,omitempty"`
	SameTypeAttackBonusMultiplier float32 `protobuf:"fixed32,9,opt,name=same_type_attack_bonus_multiplier" json:"same_type_attack_bonus_multiplier,omitempty"`
	MaximumEnergy                 int32   `protobuf:"varint,10,opt,name=maximum_energy" json:"maximum_energy,omitempty"`
	EnergyDeltaPerHealthLost      float32 `protobuf:"fixed32,11,opt,name=energy_delta_per_health_lost" json:"energy_delta_per_health_lost,omitempty"`
	DodgeDurationMs               int32   `protobuf:"varint,12,opt,name=dodge_duration_ms" json:"dodge_duration_ms,omitempty"`
	MinimumPlayerLevel            int32   `protobuf:"varint,13,opt,name=minimum_player_level" json:"minimum_player_level,omitempty"`
	SwapDurationMs                int32   `protobuf:"varint,14,opt,name=swap_duration_ms" json:"swap_duration_ms,omitempty"`
	DodgeDamageReductionPercent   float32 `protobuf:"fixed32,15,opt,name=dodge_damage_reduction_percent" json:"dodge_damage_reduction_percent,omitempty"`
}

func (m *GymBattleSettings) Reset()         { *m = GymBattleSettings{} }
func (m *GymBattleSettings) String() string { return proto.CompactTextString(m) }
func (*GymBattleSettings) ProtoMessage()    {}

type MoveSettings struct {
	MovementId          PokemonMove `protobuf:"varint,1,opt,name=movement_id,enum=POGOProtos.Enums.PokemonMove" json:"movement_id,omitempty"`
	AnimationId         int32       `protobuf:"varint,2,opt,name=animation_id" json:"animation_id,omitempty"`
	PokemonType         PokemonType `protobuf:"varint,3,opt,name=pokemon_type,enum=POGOProtos.Enums.PokemonType" json:"pokemon_type,omitempty"`
	Power               float32     `protobuf:"fixed32,4,opt,name=power" json:"power,omitempty"`
	AccuracyChance      float32     `protobuf:"fixed32,5,opt,name=accuracy_chance" json:"accuracy_chance,omitempty"`
	CriticalChance      float32     `protobuf:"fixed32,6,opt,name=critical_chance" json:"critical_chance,omitempty"`
	HealScalar          float32     `protobuf:"fixed32,7,opt,name=heal_scalar" json:"heal_scalar,omitempty"`
	StaminaLossScalar   float32     `protobuf:"fixed32,8,opt,name=stamina_loss_scalar" json:"stamina_loss_scalar,omitempty"`
	TrainerLevelMin     int32       `protobuf:"varint,9,opt,name=trainer_level_min" json:"trainer_level_min,omitempty"`
	TrainerLevelMax     int32       `protobuf:"varint,10,opt,name=trainer_level_max" json:"trainer_level_max,omitempty"`
	VfxName             string      `protobuf:"bytes,11,opt,name=vfx_name" json:"vfx_name,omitempty"`
	DurationMs          int32       `protobuf:"varint,12,opt,name=duration_ms" json:"duration_ms,omitempty"`
	DamageWindowStartMs int32       `protobuf:"varint,13,opt,name=damage_window_start_ms" json:"damage_window_start_ms,omitempty"`
	DamageWindowEndMs   int32       `protobuf:"varint,14,opt,name=damage_window_end_ms" json:"damage_window_end_ms,omitempty"`
	EnergyDelta         int32       `protobuf:"varint,15,opt,name=energy_delta" json:"energy_delta,omitempty"`
}

func (m *MoveSettings) Reset()         { *m = MoveSettings{} }
func (m *MoveSettings) String() string { return proto.CompactTextString(m) }
func (*MoveSettings) ProtoMessage()    {}

type PokemonSettings struct {
	PokemonId         PokemonId                 `protobuf:"varint,1,opt,name=pokemon_id,enum=POGOProtos.Enums.PokemonId" json:"pokemon_id,omitempty"`
	ModelScale        float32                   `protobuf:"fixed32,3,opt,name=model_scale" json:"model_scale,omitempty"`
	Type              PokemonType               `protobuf:"varint,4,opt,name=type,enum=POGOProtos.Enums.PokemonType" json:"type,omitempty"`
	Type_2            PokemonType               `protobuf:"varint,5,opt,name=type_2,enum=POGOProtos.Enums.PokemonType" json:"type_2,omitempty"`
	Camera            *CameraAttributes         `protobuf:"bytes,6,opt,name=camera" json:"camera,omitempty"`
	Encounter         *EncounterAttributes      `protobuf:"bytes,7,opt,name=encounter" json:"encounter,omitempty"`
	Stats             *StatsAttributes          `protobuf:"bytes,8,opt,name=stats" json:"stats,omitempty"`
	QuickMoves        []PokemonMove             `protobuf:"varint,9,rep,name=quick_moves,enum=POGOProtos.Enums.PokemonMove" json:"quick_moves,omitempty"`
	CinematicMoves    []PokemonMove             `protobuf:"varint,10,rep,name=cinematic_moves,enum=POGOProtos.Enums.PokemonMove" json:"cinematic_moves,omitempty"`
	AnimationTime     []float32                 `protobuf:"fixed32,11,rep,name=animation_time" json:"animation_time,omitempty"`
	EvolutionIds      []PokemonId               `protobuf:"varint,12,rep,name=evolution_ids,enum=POGOProtos.Enums.PokemonId" json:"evolution_ids,omitempty"`
	EvolutionPips     int32                     `protobuf:"varint,13,opt,name=evolution_pips" json:"evolution_pips,omitempty"`
	Rarity            PokemonRarity             `protobuf:"varint,14,opt,name=rarity,enum=POGOProtos.Enums.PokemonRarity" json:"rarity,omitempty"`
	PokedexHeightM    float32                   `protobuf:"fixed32,15,opt,name=pokedex_height_m" json:"pokedex_height_m,omitempty"`
	PokedexWeightKg   float32                   `protobuf:"fixed32,16,opt,name=pokedex_weight_kg" json:"pokedex_weight_kg,omitempty"`
	ParentPokemonId   PokemonId                 `protobuf:"varint,17,opt,name=parent_pokemon_id,enum=POGOProtos.Enums.PokemonId" json:"parent_pokemon_id,omitempty"`
	HeightStdDev      float32                   `protobuf:"fixed32,18,opt,name=height_std_dev" json:"height_std_dev,omitempty"`
	WeightStdDev      float32                   `protobuf:"fixed32,19,opt,name=weight_std_dev" json:"weight_std_dev,omitempty"`
	KmDistanceToHatch float32                   `protobuf:"fixed32,20,opt,name=km_distance_to_hatch" json:"km_distance_to_hatch,omitempty"`
	FamilyId          PokemonFamilyId           `protobuf:"varint,21,opt,name=family_id,enum=POGOProtos.Enums.PokemonFamilyId" json:"family_id,omitempty"`
	CandyToEvolve     int32                     `protobuf:"varint,22,opt,name=candy_to_evolve" json:"candy_to_evolve,omitempty"`
	KmBuddyDistance   float32                   `protobuf:"fixed32,23,opt,name=km_buddy_distance" json:"km_buddy_distance,omitempty"`
	BuddySize         PokemonSettings_BuddySize `protobuf:"varint,24,opt,name=buddy_size,enum=POGOProtos.Settings.Master.PokemonSettings_BuddySize" json:"buddy_size,omitempty"`
}

func (m *PokemonSettings) Reset()         { *m = PokemonSettings{} }
func (m *PokemonSettings) String() string { return proto.CompactTextString(m) }
func (*PokemonSettings) ProtoMessage()    {}

func (m *PokemonSettings) GetCamera() *CameraAttributes {
	if m != nil {
		return m.Camera
	}
	return nil
}

func (m *PokemonSettings) GetEncounter() *EncounterAttributes {
	if m != nil {
		return m.Encounter
	}
	return nil
}

func (m *PokemonSettings) GetStats() *StatsAttributes {
	if m != nil {
		return m.Stats
	}
	return nil
}

type IapItemDisplay struct {
	Sku       string              `protobuf:"bytes,1,opt,name=sku" json:"sku,omitempty"`
	Category  HoloIapItemCategory `protobuf:"varint,2,opt,name=category,enum=POGOProtos.Enums.HoloIapItemCategory" json:"category,omitempty"`
	SortOrder int32               `protobuf:"varint,3,opt,name=sort_order" json:"sort_order,omitempty"`
	ItemIds   []ItemId            `protobuf:"varint,4,rep,name=item_ids,enum=POGOProtos.Inventory.Item.ItemId" json:"item_ids,omitempty"`
	Counts    []int32             `protobuf:"varint,5,rep,name=counts" json:"counts,omitempty"`
}

func (m *IapItemDisplay) Reset()         { *m = IapItemDisplay{} }
func (m *IapItemDisplay) String() string { return proto.CompactTextString(m) }
func (*IapItemDisplay) ProtoMessage()    {}

type EncounterSettings struct {
	SpinBonusThreshold      float32 `protobuf:"fixed32,1,opt,name=spin_bonus_threshold" json:"spin_bonus_threshold,omitempty"`
	ExcellentThrowThreshold float32 `protobuf:"fixed32,2,opt,name=excellent_throw_threshold" json:"excellent_throw_threshold,omitempty"`
	GreatThrowThreshold     float32 `protobuf:"fixed32,3,opt,name=great_throw_threshold" json:"great_throw_threshold,omitempty"`
	NiceThrowThreshold      float32 `protobuf:"fixed32,4,opt,name=nice_throw_threshold" json:"nice_throw_threshold,omitempty"`
	MilestoneThreshold      int32   `protobuf:"varint,5,opt,name=milestone_threshold" json:"milestone_threshold,omitempty"`
}

func (m *EncounterSettings) Reset()         { *m = EncounterSettings{} }
func (m *EncounterSettings) String() string { return proto.CompactTextString(m) }
func (*EncounterSettings) ProtoMessage()    {}

type PokemonUpgradeSettings struct {
	UpgradesPerLevel         int32   `protobuf:"varint,1,opt,name=upgrades_per_level" json:"upgrades_per_level,omitempty"`
	AllowedLevelsAbovePlayer int32   `protobuf:"varint,2,opt,name=allowed_levels_above_player" json:"allowed_levels_above_player,omitempty"`
	CandyCost                []int32 `protobuf:"varint,3,rep,name=candy_cost" json:"candy_cost,omitempty"`
	StardustCost             []int32 `protobuf:"varint,4,rep,name=stardust_cost" json:"stardust_cost,omitempty"`
}

func (m *PokemonUpgradeSettings) Reset()         { *m = PokemonUpgradeSettings{} }
func (m *PokemonUpgradeSettings) String() string { return proto.CompactTextString(m) }
func (*PokemonUpgradeSettings) ProtoMessage()    {}

type ItemSettings struct {
	ItemId           ItemId                      `protobuf:"varint,1,opt,name=item_id,enum=POGOProtos.Inventory.Item.ItemId" json:"item_id,omitempty"`
	ItemType         ItemType                    `protobuf:"varint,2,opt,name=item_type,enum=POGOProtos.Inventory.Item.ItemType" json:"item_type,omitempty"`
	Category         ItemCategory                `protobuf:"varint,3,opt,name=category,enum=POGOProtos.Enums.ItemCategory" json:"category,omitempty"`
	DropFreq         float32                     `protobuf:"fixed32,4,opt,name=drop_freq" json:"drop_freq,omitempty"`
	DropTrainerLevel int32                       `protobuf:"varint,5,opt,name=drop_trainer_level" json:"drop_trainer_level,omitempty"`
	Pokeball         *PokeballAttributes         `protobuf:"bytes,6,opt,name=pokeball" json:"pokeball,omitempty"`
	Potion           *PotionAttributes           `protobuf:"bytes,7,opt,name=potion" json:"potion,omitempty"`
	Revive           *ReviveAttributes           `protobuf:"bytes,8,opt,name=revive" json:"revive,omitempty"`
	Battle           *BattleAttributes           `protobuf:"bytes,9,opt,name=battle" json:"battle,omitempty"`
	Food             *FoodAttributes             `protobuf:"bytes,10,opt,name=food" json:"food,omitempty"`
	InventoryUpgrade *InventoryUpgradeAttributes `protobuf:"bytes,11,opt,name=inventory_upgrade" json:"inventory_upgrade,omitempty"`
	XpBoost          *ExperienceBoostAttributes  `protobuf:"bytes,12,opt,name=xp_boost" json:"xp_boost,omitempty"`
	Incense          *IncenseAttributes          `protobuf:"bytes,13,opt,name=incense" json:"incense,omitempty"`
	EggIncubator     *EggIncubatorAttributes     `protobuf:"bytes,14,opt,name=egg_incubator" json:"egg_incubator,omitempty"`
	FortModifier     *FortModifierAttributes     `protobuf:"bytes,15,opt,name=fort_modifier" json:"fort_modifier,omitempty"`
}

func (m *ItemSettings) Reset()         { *m = ItemSettings{} }
func (m *ItemSettings) String() string { return proto.CompactTextString(m) }
func (*ItemSettings) ProtoMessage()    {}

func (m *ItemSettings) GetPokeball() *PokeballAttributes {
	if m != nil {
		return m.Pokeball
	}
	return nil
}

func (m *ItemSettings) GetPotion() *PotionAttributes {
	if m != nil {
		return m.Potion
	}
	return nil
}

func (m *ItemSettings) GetRevive() *ReviveAttributes {
	if m != nil {
		return m.Revive
	}
	return nil
}

func (m *ItemSettings) GetBattle() *BattleAttributes {
	if m != nil {
		return m.Battle
	}
	return nil
}

func (m *ItemSettings) GetFood() *FoodAttributes {
	if m != nil {
		return m.Food
	}
	return nil
}

func (m *ItemSettings) GetInventoryUpgrade() *InventoryUpgradeAttributes {
	if m != nil {
		return m.InventoryUpgrade
	}
	return nil
}

func (m *ItemSettings) GetXpBoost() *ExperienceBoostAttributes {
	if m != nil {
		return m.XpBoost
	}
	return nil
}

func (m *ItemSettings) GetIncense() *IncenseAttributes {
	if m != nil {
		return m.Incense
	}
	return nil
}

func (m *ItemSettings) GetEggIncubator() *EggIncubatorAttributes {
	if m != nil {
		return m.EggIncubator
	}
	return nil
}

func (m *ItemSettings) GetFortModifier() *FortModifierAttributes {
	if m != nil {
		return m.FortModifier
	}
	return nil
}

type MoveSequenceSettings struct {
	Sequence []string `protobuf:"bytes,1,rep,name=sequence" json:"sequence,omitempty"`
}

func (m *MoveSequenceSettings) Reset()         { *m = MoveSequenceSettings{} }
func (m *MoveSequenceSettings) String() string { return proto.CompactTextString(m) }
func (*MoveSequenceSettings) ProtoMessage()    {}

type BadgeSettings struct {
	BadgeType BadgeType `protobuf:"varint,1,opt,name=badge_type,enum=POGOProtos.Enums.BadgeType" json:"badge_type,omitempty"`
	BadgeRank int32     `protobuf:"varint,2,opt,name=badge_rank" json:"badge_rank,omitempty"`
	Targets   []int32   `protobuf:"varint,3,rep,name=targets" json:"targets,omitempty"`
}

func (m *BadgeSettings) Reset()         { *m = BadgeSettings{} }
func (m *BadgeSettings) String() string { return proto.CompactTextString(m) }
func (*BadgeSettings) ProtoMessage()    {}

type GymLevelSettings struct {
	RequiredExperience []int32 `protobuf:"varint,1,rep,name=required_experience" json:"required_experience,omitempty"`
	LeaderSlots        []int32 `protobuf:"varint,2,rep,name=leader_slots" json:"leader_slots,omitempty"`
	TrainerSlots       []int32 `protobuf:"varint,3,rep,name=trainer_slots" json:"trainer_slots,omitempty"`
	SearchRollBonus    []int32 `protobuf:"varint,4,rep,name=search_roll_bonus" json:"search_roll_bonus,omitempty"`
}

func (m *GymLevelSettings) Reset()         { *m = GymLevelSettings{} }
func (m *GymLevelSettings) String() string { return proto.CompactTextString(m) }
func (*GymLevelSettings) ProtoMessage()    {}

func init() {
	proto.RegisterEnum("POGOProtos.Settings.Master.PokemonSettings_BuddySize", PokemonSettings_BuddySize_name, PokemonSettings_BuddySize_value)
}
