// Code generated by gen_entity.go DO NOT EDIT.
// Package entity stores information about entities in Minecraft.
package entity

// ID describes the numeric ID of an entity.
type ID uint32

// Entity describes information about a type of entity.
type Entity struct {
	ID          ID
	InternalID  uint32
	DisplayName string
	Name        string
	Width       float64
	Height      float64
	Type        string
}

var (
	Allay = Entity{
		ID:          0,
		InternalID:  0,
		DisplayName: "Allay",
		Name:        "allay",
		Width:       0.35,
		Height:      0.6,
		Type:        "mob",
	}
	AreaEffectCloud = Entity{
		ID:          1,
		InternalID:  1,
		DisplayName: "Area Effect Cloud",
		Name:        "area_effect_cloud",
		Width:       6,
		Height:      0.5,
		Type:        "other",
	}
	ArmorStand = Entity{
		ID:          2,
		InternalID:  2,
		DisplayName: "Armor Stand",
		Name:        "armor_stand",
		Width:       0.5,
		Height:      1.975,
		Type:        "living",
	}
	Arrow = Entity{
		ID:          3,
		InternalID:  3,
		DisplayName: "Arrow",
		Name:        "arrow",
		Width:       0.5,
		Height:      0.5,
		Type:        "projectile",
	}
	Axolotl = Entity{
		ID:          4,
		InternalID:  4,
		DisplayName: "Axolotl",
		Name:        "axolotl",
		Width:       0.75,
		Height:      0.42,
		Type:        "animal",
	}
	Bat = Entity{
		ID:          5,
		InternalID:  5,
		DisplayName: "Bat",
		Name:        "bat",
		Width:       0.5,
		Height:      0.9,
		Type:        "ambient",
	}
	Bee = Entity{
		ID:          6,
		InternalID:  6,
		DisplayName: "Bee",
		Name:        "bee",
		Width:       0.7,
		Height:      0.6,
		Type:        "animal",
	}
	Blaze = Entity{
		ID:          7,
		InternalID:  7,
		DisplayName: "Blaze",
		Name:        "blaze",
		Width:       0.6,
		Height:      1.8,
		Type:        "hostile",
	}
	Boat = Entity{
		ID:          8,
		InternalID:  8,
		DisplayName: "Boat",
		Name:        "boat",
		Width:       1.375,
		Height:      0.5625,
		Type:        "other",
	}
	ChestBoat = Entity{
		ID:          9,
		InternalID:  9,
		DisplayName: "Boat with Chest",
		Name:        "chest_boat",
		Width:       1.375,
		Height:      0.5625,
		Type:        "other",
	}
	Cat = Entity{
		ID:          10,
		InternalID:  10,
		DisplayName: "Cat",
		Name:        "cat",
		Width:       0.6,
		Height:      0.7,
		Type:        "animal",
	}
	CaveSpider = Entity{
		ID:          11,
		InternalID:  11,
		DisplayName: "Cave Spider",
		Name:        "cave_spider",
		Width:       0.7,
		Height:      0.5,
		Type:        "hostile",
	}
	Chicken = Entity{
		ID:          12,
		InternalID:  12,
		DisplayName: "Chicken",
		Name:        "chicken",
		Width:       0.4,
		Height:      0.7,
		Type:        "animal",
	}
	Cod = Entity{
		ID:          13,
		InternalID:  13,
		DisplayName: "Cod",
		Name:        "cod",
		Width:       0.5,
		Height:      0.3,
		Type:        "water_creature",
	}
	Cow = Entity{
		ID:          14,
		InternalID:  14,
		DisplayName: "Cow",
		Name:        "cow",
		Width:       0.9,
		Height:      1.4,
		Type:        "animal",
	}
	Creeper = Entity{
		ID:          15,
		InternalID:  15,
		DisplayName: "Creeper",
		Name:        "creeper",
		Width:       0.6,
		Height:      1.7,
		Type:        "hostile",
	}
	Dolphin = Entity{
		ID:          16,
		InternalID:  16,
		DisplayName: "Dolphin",
		Name:        "dolphin",
		Width:       0.9,
		Height:      0.6,
		Type:        "water_creature",
	}
	Donkey = Entity{
		ID:          17,
		InternalID:  17,
		DisplayName: "Donkey",
		Name:        "donkey",
		Width:       1.3964844,
		Height:      1.5,
		Type:        "animal",
	}
	DragonFireball = Entity{
		ID:          18,
		InternalID:  18,
		DisplayName: "Dragon Fireball",
		Name:        "dragon_fireball",
		Width:       1,
		Height:      1,
		Type:        "projectile",
	}
	Drowned = Entity{
		ID:          19,
		InternalID:  19,
		DisplayName: "Drowned",
		Name:        "drowned",
		Width:       0.6,
		Height:      1.95,
		Type:        "hostile",
	}
	ElderGuardian = Entity{
		ID:          20,
		InternalID:  20,
		DisplayName: "Elder Guardian",
		Name:        "elder_guardian",
		Width:       1.9975,
		Height:      1.9975,
		Type:        "hostile",
	}
	EndCrystal = Entity{
		ID:          21,
		InternalID:  21,
		DisplayName: "End Crystal",
		Name:        "end_crystal",
		Width:       2,
		Height:      2,
		Type:        "other",
	}
	EnderDragon = Entity{
		ID:          22,
		InternalID:  22,
		DisplayName: "Ender Dragon",
		Name:        "ender_dragon",
		Width:       16,
		Height:      8,
		Type:        "mob",
	}
	Enderman = Entity{
		ID:          23,
		InternalID:  23,
		DisplayName: "Enderman",
		Name:        "enderman",
		Width:       0.6,
		Height:      2.9,
		Type:        "hostile",
	}
	Endermite = Entity{
		ID:          24,
		InternalID:  24,
		DisplayName: "Endermite",
		Name:        "endermite",
		Width:       0.4,
		Height:      0.3,
		Type:        "hostile",
	}
	Evoker = Entity{
		ID:          25,
		InternalID:  25,
		DisplayName: "Evoker",
		Name:        "evoker",
		Width:       0.6,
		Height:      1.95,
		Type:        "hostile",
	}
	EvokerFangs = Entity{
		ID:          26,
		InternalID:  26,
		DisplayName: "Evoker Fangs",
		Name:        "evoker_fangs",
		Width:       0.5,
		Height:      0.8,
		Type:        "other",
	}
	ExperienceOrb = Entity{
		ID:          27,
		InternalID:  27,
		DisplayName: "Experience Orb",
		Name:        "experience_orb",
		Width:       0.5,
		Height:      0.5,
		Type:        "other",
	}
	EyeOfEnder = Entity{
		ID:          28,
		InternalID:  28,
		DisplayName: "Eye of Ender",
		Name:        "eye_of_ender",
		Width:       0.25,
		Height:      0.25,
		Type:        "other",
	}
	FallingBlock = Entity{
		ID:          29,
		InternalID:  29,
		DisplayName: "Falling Block",
		Name:        "falling_block",
		Width:       0.98,
		Height:      0.98,
		Type:        "other",
	}
	FireworkRocket = Entity{
		ID:          30,
		InternalID:  30,
		DisplayName: "Firework Rocket",
		Name:        "firework_rocket",
		Width:       0.25,
		Height:      0.25,
		Type:        "projectile",
	}
	Fox = Entity{
		ID:          31,
		InternalID:  31,
		DisplayName: "Fox",
		Name:        "fox",
		Width:       0.6,
		Height:      0.7,
		Type:        "animal",
	}
	Frog = Entity{
		ID:          32,
		InternalID:  32,
		DisplayName: "Frog",
		Name:        "frog",
		Width:       0.5,
		Height:      0.5,
		Type:        "animal",
	}
	Ghast = Entity{
		ID:          33,
		InternalID:  33,
		DisplayName: "Ghast",
		Name:        "ghast",
		Width:       4,
		Height:      4,
		Type:        "mob",
	}
	Giant = Entity{
		ID:          34,
		InternalID:  34,
		DisplayName: "Giant",
		Name:        "giant",
		Width:       3.6,
		Height:      12,
		Type:        "hostile",
	}
	GlowItemFrame = Entity{
		ID:          35,
		InternalID:  35,
		DisplayName: "Glow Item Frame",
		Name:        "glow_item_frame",
		Width:       0.5,
		Height:      0.5,
		Type:        "other",
	}
	GlowSquid = Entity{
		ID:          36,
		InternalID:  36,
		DisplayName: "Glow Squid",
		Name:        "glow_squid",
		Width:       0.8,
		Height:      0.8,
		Type:        "water_creature",
	}
	Goat = Entity{
		ID:          37,
		InternalID:  37,
		DisplayName: "Goat",
		Name:        "goat",
		Width:       0.9,
		Height:      1.3,
		Type:        "animal",
	}
	Guardian = Entity{
		ID:          38,
		InternalID:  38,
		DisplayName: "Guardian",
		Name:        "guardian",
		Width:       0.85,
		Height:      0.85,
		Type:        "hostile",
	}
	Hoglin = Entity{
		ID:          39,
		InternalID:  39,
		DisplayName: "Hoglin",
		Name:        "hoglin",
		Width:       1.3964844,
		Height:      1.4,
		Type:        "animal",
	}
	Horse = Entity{
		ID:          40,
		InternalID:  40,
		DisplayName: "Horse",
		Name:        "horse",
		Width:       1.3964844,
		Height:      1.6,
		Type:        "animal",
	}
	Husk = Entity{
		ID:          41,
		InternalID:  41,
		DisplayName: "Husk",
		Name:        "husk",
		Width:       0.6,
		Height:      1.95,
		Type:        "hostile",
	}
	Illusioner = Entity{
		ID:          42,
		InternalID:  42,
		DisplayName: "Illusioner",
		Name:        "illusioner",
		Width:       0.6,
		Height:      1.95,
		Type:        "hostile",
	}
	IronGolem = Entity{
		ID:          43,
		InternalID:  43,
		DisplayName: "Iron Golem",
		Name:        "iron_golem",
		Width:       1.4,
		Height:      2.7,
		Type:        "mob",
	}
	Item = Entity{
		ID:          44,
		InternalID:  44,
		DisplayName: "Item",
		Name:        "item",
		Width:       0.25,
		Height:      0.25,
		Type:        "other",
	}
	ItemFrame = Entity{
		ID:          45,
		InternalID:  45,
		DisplayName: "Item Frame",
		Name:        "item_frame",
		Width:       0.5,
		Height:      0.5,
		Type:        "other",
	}
	Fireball = Entity{
		ID:          46,
		InternalID:  46,
		DisplayName: "Fireball",
		Name:        "fireball",
		Width:       1,
		Height:      1,
		Type:        "projectile",
	}
	LeashKnot = Entity{
		ID:          47,
		InternalID:  47,
		DisplayName: "Leash Knot",
		Name:        "leash_knot",
		Width:       0.375,
		Height:      0.5,
		Type:        "other",
	}
	LightningBolt = Entity{
		ID:          48,
		InternalID:  48,
		DisplayName: "Lightning Bolt",
		Name:        "lightning_bolt",
		Width:       0,
		Height:      0,
		Type:        "other",
	}
	Llama = Entity{
		ID:          49,
		InternalID:  49,
		DisplayName: "Llama",
		Name:        "llama",
		Width:       0.9,
		Height:      1.87,
		Type:        "animal",
	}
	LlamaSpit = Entity{
		ID:          50,
		InternalID:  50,
		DisplayName: "Llama Spit",
		Name:        "llama_spit",
		Width:       0.25,
		Height:      0.25,
		Type:        "projectile",
	}
	MagmaCube = Entity{
		ID:          51,
		InternalID:  51,
		DisplayName: "Magma Cube",
		Name:        "magma_cube",
		Width:       2.04,
		Height:      2.04,
		Type:        "mob",
	}
	Marker = Entity{
		ID:          52,
		InternalID:  52,
		DisplayName: "Marker",
		Name:        "marker",
		Width:       0,
		Height:      0,
		Type:        "other",
	}
	Minecart = Entity{
		ID:          53,
		InternalID:  53,
		DisplayName: "Minecart",
		Name:        "minecart",
		Width:       0.98,
		Height:      0.7,
		Type:        "other",
	}
	ChestMinecart = Entity{
		ID:          54,
		InternalID:  54,
		DisplayName: "Minecart with Chest",
		Name:        "chest_minecart",
		Width:       0.98,
		Height:      0.7,
		Type:        "other",
	}
	CommandBlockMinecart = Entity{
		ID:          55,
		InternalID:  55,
		DisplayName: "Minecart with Command Block",
		Name:        "command_block_minecart",
		Width:       0.98,
		Height:      0.7,
		Type:        "other",
	}
	FurnaceMinecart = Entity{
		ID:          56,
		InternalID:  56,
		DisplayName: "Minecart with Furnace",
		Name:        "furnace_minecart",
		Width:       0.98,
		Height:      0.7,
		Type:        "other",
	}
	HopperMinecart = Entity{
		ID:          57,
		InternalID:  57,
		DisplayName: "Minecart with Hopper",
		Name:        "hopper_minecart",
		Width:       0.98,
		Height:      0.7,
		Type:        "other",
	}
	SpawnerMinecart = Entity{
		ID:          58,
		InternalID:  58,
		DisplayName: "Minecart with Spawner",
		Name:        "spawner_minecart",
		Width:       0.98,
		Height:      0.7,
		Type:        "other",
	}
	TntMinecart = Entity{
		ID:          59,
		InternalID:  59,
		DisplayName: "Minecart with TNT",
		Name:        "tnt_minecart",
		Width:       0.98,
		Height:      0.7,
		Type:        "other",
	}
	Mule = Entity{
		ID:          60,
		InternalID:  60,
		DisplayName: "Mule",
		Name:        "mule",
		Width:       1.3964844,
		Height:      1.6,
		Type:        "animal",
	}
	Mooshroom = Entity{
		ID:          61,
		InternalID:  61,
		DisplayName: "Mooshroom",
		Name:        "mooshroom",
		Width:       0.9,
		Height:      1.4,
		Type:        "animal",
	}
	Ocelot = Entity{
		ID:          62,
		InternalID:  62,
		DisplayName: "Ocelot",
		Name:        "ocelot",
		Width:       0.6,
		Height:      0.7,
		Type:        "animal",
	}
	Painting = Entity{
		ID:          63,
		InternalID:  63,
		DisplayName: "Painting",
		Name:        "painting",
		Width:       0.5,
		Height:      0.5,
		Type:        "other",
	}
	Panda = Entity{
		ID:          64,
		InternalID:  64,
		DisplayName: "Panda",
		Name:        "panda",
		Width:       1.3,
		Height:      1.25,
		Type:        "animal",
	}
	Parrot = Entity{
		ID:          65,
		InternalID:  65,
		DisplayName: "Parrot",
		Name:        "parrot",
		Width:       0.5,
		Height:      0.9,
		Type:        "animal",
	}
	Phantom = Entity{
		ID:          66,
		InternalID:  66,
		DisplayName: "Phantom",
		Name:        "phantom",
		Width:       0.9,
		Height:      0.5,
		Type:        "mob",
	}
	Pig = Entity{
		ID:          67,
		InternalID:  67,
		DisplayName: "Pig",
		Name:        "pig",
		Width:       0.9,
		Height:      0.9,
		Type:        "animal",
	}
	Piglin = Entity{
		ID:          68,
		InternalID:  68,
		DisplayName: "Piglin",
		Name:        "piglin",
		Width:       0.6,
		Height:      1.95,
		Type:        "hostile",
	}
	PiglinBrute = Entity{
		ID:          69,
		InternalID:  69,
		DisplayName: "Piglin Brute",
		Name:        "piglin_brute",
		Width:       0.6,
		Height:      1.95,
		Type:        "hostile",
	}
	Pillager = Entity{
		ID:          70,
		InternalID:  70,
		DisplayName: "Pillager",
		Name:        "pillager",
		Width:       0.6,
		Height:      1.95,
		Type:        "hostile",
	}
	PolarBear = Entity{
		ID:          71,
		InternalID:  71,
		DisplayName: "Polar Bear",
		Name:        "polar_bear",
		Width:       1.4,
		Height:      1.4,
		Type:        "animal",
	}
	Tnt = Entity{
		ID:          72,
		InternalID:  72,
		DisplayName: "Primed TNT",
		Name:        "tnt",
		Width:       0.98,
		Height:      0.98,
		Type:        "other",
	}
	Pufferfish = Entity{
		ID:          73,
		InternalID:  73,
		DisplayName: "Pufferfish",
		Name:        "pufferfish",
		Width:       0.7,
		Height:      0.7,
		Type:        "water_creature",
	}
	Rabbit = Entity{
		ID:          74,
		InternalID:  74,
		DisplayName: "Rabbit",
		Name:        "rabbit",
		Width:       0.4,
		Height:      0.5,
		Type:        "animal",
	}
	Ravager = Entity{
		ID:          75,
		InternalID:  75,
		DisplayName: "Ravager",
		Name:        "ravager",
		Width:       1.95,
		Height:      2.2,
		Type:        "hostile",
	}
	Salmon = Entity{
		ID:          76,
		InternalID:  76,
		DisplayName: "Salmon",
		Name:        "salmon",
		Width:       0.7,
		Height:      0.4,
		Type:        "water_creature",
	}
	Sheep = Entity{
		ID:          77,
		InternalID:  77,
		DisplayName: "Sheep",
		Name:        "sheep",
		Width:       0.9,
		Height:      1.3,
		Type:        "animal",
	}
	Shulker = Entity{
		ID:          78,
		InternalID:  78,
		DisplayName: "Shulker",
		Name:        "shulker",
		Width:       1,
		Height:      1,
		Type:        "mob",
	}
	ShulkerBullet = Entity{
		ID:          79,
		InternalID:  79,
		DisplayName: "Shulker Bullet",
		Name:        "shulker_bullet",
		Width:       0.3125,
		Height:      0.3125,
		Type:        "projectile",
	}
	Silverfish = Entity{
		ID:          80,
		InternalID:  80,
		DisplayName: "Silverfish",
		Name:        "silverfish",
		Width:       0.4,
		Height:      0.3,
		Type:        "hostile",
	}
	Skeleton = Entity{
		ID:          81,
		InternalID:  81,
		DisplayName: "Skeleton",
		Name:        "skeleton",
		Width:       0.6,
		Height:      1.99,
		Type:        "hostile",
	}
	SkeletonHorse = Entity{
		ID:          82,
		InternalID:  82,
		DisplayName: "Skeleton Horse",
		Name:        "skeleton_horse",
		Width:       1.3964844,
		Height:      1.6,
		Type:        "animal",
	}
	Slime = Entity{
		ID:          83,
		InternalID:  83,
		DisplayName: "Slime",
		Name:        "slime",
		Width:       2.04,
		Height:      2.04,
		Type:        "mob",
	}
	SmallFireball = Entity{
		ID:          84,
		InternalID:  84,
		DisplayName: "Small Fireball",
		Name:        "small_fireball",
		Width:       0.3125,
		Height:      0.3125,
		Type:        "projectile",
	}
	SnowGolem = Entity{
		ID:          85,
		InternalID:  85,
		DisplayName: "Snow Golem",
		Name:        "snow_golem",
		Width:       0.7,
		Height:      1.9,
		Type:        "mob",
	}
	Snowball = Entity{
		ID:          86,
		InternalID:  86,
		DisplayName: "Snowball",
		Name:        "snowball",
		Width:       0.25,
		Height:      0.25,
		Type:        "projectile",
	}
	SpectralArrow = Entity{
		ID:          87,
		InternalID:  87,
		DisplayName: "Spectral Arrow",
		Name:        "spectral_arrow",
		Width:       0.5,
		Height:      0.5,
		Type:        "projectile",
	}
	Spider = Entity{
		ID:          88,
		InternalID:  88,
		DisplayName: "Spider",
		Name:        "spider",
		Width:       1.4,
		Height:      0.9,
		Type:        "hostile",
	}
	Squid = Entity{
		ID:          89,
		InternalID:  89,
		DisplayName: "Squid",
		Name:        "squid",
		Width:       0.8,
		Height:      0.8,
		Type:        "water_creature",
	}
	Stray = Entity{
		ID:          90,
		InternalID:  90,
		DisplayName: "Stray",
		Name:        "stray",
		Width:       0.6,
		Height:      1.99,
		Type:        "hostile",
	}
	Strider = Entity{
		ID:          91,
		InternalID:  91,
		DisplayName: "Strider",
		Name:        "strider",
		Width:       0.9,
		Height:      1.7,
		Type:        "animal",
	}
	Tadpole = Entity{
		ID:          92,
		InternalID:  92,
		DisplayName: "Tadpole",
		Name:        "tadpole",
		Width:       0.4,
		Height:      0.3,
		Type:        "water_creature",
	}
	Egg = Entity{
		ID:          93,
		InternalID:  93,
		DisplayName: "Thrown Egg",
		Name:        "egg",
		Width:       0.25,
		Height:      0.25,
		Type:        "projectile",
	}
	EnderPearl = Entity{
		ID:          94,
		InternalID:  94,
		DisplayName: "Thrown Ender Pearl",
		Name:        "ender_pearl",
		Width:       0.25,
		Height:      0.25,
		Type:        "projectile",
	}
	ExperienceBottle = Entity{
		ID:          95,
		InternalID:  95,
		DisplayName: "Thrown Bottle o' Enchanting",
		Name:        "experience_bottle",
		Width:       0.25,
		Height:      0.25,
		Type:        "projectile",
	}
	Potion = Entity{
		ID:          96,
		InternalID:  96,
		DisplayName: "Potion",
		Name:        "potion",
		Width:       0.25,
		Height:      0.25,
		Type:        "projectile",
	}
	Trident = Entity{
		ID:          97,
		InternalID:  97,
		DisplayName: "Trident",
		Name:        "trident",
		Width:       0.5,
		Height:      0.5,
		Type:        "projectile",
	}
	TraderLlama = Entity{
		ID:          98,
		InternalID:  98,
		DisplayName: "Trader Llama",
		Name:        "trader_llama",
		Width:       0.9,
		Height:      1.87,
		Type:        "animal",
	}
	TropicalFish = Entity{
		ID:          99,
		InternalID:  99,
		DisplayName: "Tropical Fish",
		Name:        "tropical_fish",
		Width:       0.5,
		Height:      0.4,
		Type:        "water_creature",
	}
	Turtle = Entity{
		ID:          100,
		InternalID:  100,
		DisplayName: "Turtle",
		Name:        "turtle",
		Width:       1.2,
		Height:      0.4,
		Type:        "animal",
	}
	Vex = Entity{
		ID:          101,
		InternalID:  101,
		DisplayName: "Vex",
		Name:        "vex",
		Width:       0.4,
		Height:      0.8,
		Type:        "hostile",
	}
	Villager = Entity{
		ID:          102,
		InternalID:  102,
		DisplayName: "Villager",
		Name:        "villager",
		Width:       0.6,
		Height:      1.95,
		Type:        "passive",
	}
	Vindicator = Entity{
		ID:          103,
		InternalID:  103,
		DisplayName: "Vindicator",
		Name:        "vindicator",
		Width:       0.6,
		Height:      1.95,
		Type:        "hostile",
	}
	WanderingTrader = Entity{
		ID:          104,
		InternalID:  104,
		DisplayName: "Wandering Trader",
		Name:        "wandering_trader",
		Width:       0.6,
		Height:      1.95,
		Type:        "passive",
	}
	Warden = Entity{
		ID:          105,
		InternalID:  105,
		DisplayName: "Warden",
		Name:        "warden",
		Width:       0.9,
		Height:      2.9,
		Type:        "hostile",
	}
	Witch = Entity{
		ID:          106,
		InternalID:  106,
		DisplayName: "Witch",
		Name:        "witch",
		Width:       0.6,
		Height:      1.95,
		Type:        "hostile",
	}
	Wither = Entity{
		ID:          107,
		InternalID:  107,
		DisplayName: "Wither",
		Name:        "wither",
		Width:       0.9,
		Height:      3.5,
		Type:        "hostile",
	}
	WitherSkeleton = Entity{
		ID:          108,
		InternalID:  108,
		DisplayName: "Wither Skeleton",
		Name:        "wither_skeleton",
		Width:       0.7,
		Height:      2.4,
		Type:        "hostile",
	}
	WitherSkull = Entity{
		ID:          109,
		InternalID:  109,
		DisplayName: "Wither Skull",
		Name:        "wither_skull",
		Width:       0.3125,
		Height:      0.3125,
		Type:        "projectile",
	}
	Wolf = Entity{
		ID:          110,
		InternalID:  110,
		DisplayName: "Wolf",
		Name:        "wolf",
		Width:       0.6,
		Height:      0.85,
		Type:        "animal",
	}
	Zoglin = Entity{
		ID:          111,
		InternalID:  111,
		DisplayName: "Zoglin",
		Name:        "zoglin",
		Width:       1.3964844,
		Height:      1.4,
		Type:        "hostile",
	}
	Zombie = Entity{
		ID:          112,
		InternalID:  112,
		DisplayName: "Zombie",
		Name:        "zombie",
		Width:       0.6,
		Height:      1.95,
		Type:        "hostile",
	}
	ZombieHorse = Entity{
		ID:          113,
		InternalID:  113,
		DisplayName: "Zombie Horse",
		Name:        "zombie_horse",
		Width:       1.3964844,
		Height:      1.6,
		Type:        "animal",
	}
	ZombieVillager = Entity{
		ID:          114,
		InternalID:  114,
		DisplayName: "Zombie Villager",
		Name:        "zombie_villager",
		Width:       0.6,
		Height:      1.95,
		Type:        "hostile",
	}
	ZombifiedPiglin = Entity{
		ID:          115,
		InternalID:  115,
		DisplayName: "Zombified Piglin",
		Name:        "zombified_piglin",
		Width:       0.6,
		Height:      1.95,
		Type:        "hostile",
	}
	Player = Entity{
		ID:          116,
		InternalID:  116,
		DisplayName: "Player",
		Name:        "player",
		Width:       0.6,
		Height:      1.8,
		Type:        "player",
	}
	FishingBobber = Entity{
		ID:          117,
		InternalID:  117,
		DisplayName: "Fishing Bobber",
		Name:        "fishing_bobber",
		Width:       0.25,
		Height:      0.25,
		Type:        "projectile",
	}
)

// ByID is an index of minecraft entities by their ID.
var ByID = map[ID]*Entity{
	0:   &Allay,
	1:   &AreaEffectCloud,
	2:   &ArmorStand,
	3:   &Arrow,
	4:   &Axolotl,
	5:   &Bat,
	6:   &Bee,
	7:   &Blaze,
	8:   &Boat,
	9:   &ChestBoat,
	10:  &Cat,
	11:  &CaveSpider,
	12:  &Chicken,
	13:  &Cod,
	14:  &Cow,
	15:  &Creeper,
	16:  &Dolphin,
	17:  &Donkey,
	18:  &DragonFireball,
	19:  &Drowned,
	20:  &ElderGuardian,
	21:  &EndCrystal,
	22:  &EnderDragon,
	23:  &Enderman,
	24:  &Endermite,
	25:  &Evoker,
	26:  &EvokerFangs,
	27:  &ExperienceOrb,
	28:  &EyeOfEnder,
	29:  &FallingBlock,
	30:  &FireworkRocket,
	31:  &Fox,
	32:  &Frog,
	33:  &Ghast,
	34:  &Giant,
	35:  &GlowItemFrame,
	36:  &GlowSquid,
	37:  &Goat,
	38:  &Guardian,
	39:  &Hoglin,
	40:  &Horse,
	41:  &Husk,
	42:  &Illusioner,
	43:  &IronGolem,
	44:  &Item,
	45:  &ItemFrame,
	46:  &Fireball,
	47:  &LeashKnot,
	48:  &LightningBolt,
	49:  &Llama,
	50:  &LlamaSpit,
	51:  &MagmaCube,
	52:  &Marker,
	53:  &Minecart,
	54:  &ChestMinecart,
	55:  &CommandBlockMinecart,
	56:  &FurnaceMinecart,
	57:  &HopperMinecart,
	58:  &SpawnerMinecart,
	59:  &TntMinecart,
	60:  &Mule,
	61:  &Mooshroom,
	62:  &Ocelot,
	63:  &Painting,
	64:  &Panda,
	65:  &Parrot,
	66:  &Phantom,
	67:  &Pig,
	68:  &Piglin,
	69:  &PiglinBrute,
	70:  &Pillager,
	71:  &PolarBear,
	72:  &Tnt,
	73:  &Pufferfish,
	74:  &Rabbit,
	75:  &Ravager,
	76:  &Salmon,
	77:  &Sheep,
	78:  &Shulker,
	79:  &ShulkerBullet,
	80:  &Silverfish,
	81:  &Skeleton,
	82:  &SkeletonHorse,
	83:  &Slime,
	84:  &SmallFireball,
	85:  &SnowGolem,
	86:  &Snowball,
	87:  &SpectralArrow,
	88:  &Spider,
	89:  &Squid,
	90:  &Stray,
	91:  &Strider,
	92:  &Tadpole,
	93:  &Egg,
	94:  &EnderPearl,
	95:  &ExperienceBottle,
	96:  &Potion,
	97:  &Trident,
	98:  &TraderLlama,
	99:  &TropicalFish,
	100: &Turtle,
	101: &Vex,
	102: &Villager,
	103: &Vindicator,
	104: &WanderingTrader,
	105: &Warden,
	106: &Witch,
	107: &Wither,
	108: &WitherSkeleton,
	109: &WitherSkull,
	110: &Wolf,
	111: &Zoglin,
	112: &Zombie,
	113: &ZombieHorse,
	114: &ZombieVillager,
	115: &ZombifiedPiglin,
	116: &Player,
	117: &FishingBobber,
}
