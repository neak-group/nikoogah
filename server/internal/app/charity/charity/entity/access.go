package entity

type CharityAccess string

const (
	ViewCharityAK = "ak.charity.view"
	EditCharityAK = "ak.charity.modify"

	AddNewRallyAK = "ak.rally.new"
	EditRallyAK   = "ak.rally.modify"

	ViewParticipationAK   = "ak.participation.view"
	ManageParticipationAK = "ak.participation.manage"
)

var roleAccess map[RepresentativeRole][]CharityAccess = map[RepresentativeRole][]CharityAccess{
	Manager: []CharityAccess{
		ViewCharityAK,
		EditCharityAK,
		AddNewRallyAK,
		EditRallyAK,
		ViewParticipationAK,
		ManageParticipationAK,
	},
	Employee: []CharityAccess{
		ViewCharityAK,
		AddNewRallyAK,
		ViewParticipationAK,
		ManageParticipationAK,
	},
}
