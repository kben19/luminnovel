package entity

type ProductTitle string

// Title representing keys in database and sheet name
const (
	Bofuri               ProductTitle = "Bofuri"
	ClassroomElite       ProductTitle = "Classroom Elite"
	DeathMarch           ProductTitle = "Death March"
	DevilPartTimer       ProductTitle = "Devil Part Timer"
	EightySix            ProductTitle = "EightySix"
	ImSpiderSoWhat       ProductTitle = "I'm a spider so what"
	InLoveVillainess     ProductTitle = "InLoveVillainess"
	KillingSlime         ProductTitle = "Killing Slime"
	Konosuba             ProductTitle = "Konosuba"
	MushokuTensei        ProductTitle = "Mushoku Tensei"
	ReincarnatedAsASlime ProductTitle = "Reincarnated as a slime"
	ReZero               ProductTitle = "ReZero"
	SkeletonKnight       ProductTitle = "Skeleton Knight"
	Smartphone           ProductTitle = "Smartphone"
	TrappedDatingSim     ProductTitle = "Dating Sim"
)

var ListAllTitles = []ProductTitle{
	Bofuri,
	ClassroomElite,
	DeathMarch,
	DevilPartTimer,
	EightySix,
	ImSpiderSoWhat,
	InLoveVillainess,
	KillingSlime,
	Konosuba,
	MushokuTensei,
	ReZero,
	ReincarnatedAsASlime,
	SkeletonKnight,
	Smartphone,
	TrappedDatingSim,
}
