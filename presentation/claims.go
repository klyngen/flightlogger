package presentation

/*
* We need some set of constants to represent our claims
 */

type Claim int

const (
	ClubLeader Claim = iota
	Admin
	NationalLeader
)

// ToString make the claim human readable
func (c Claim) ToString() string {
	return [...]string{"ClubLeader", "Admin", "NationalLeader"}[c]
}
