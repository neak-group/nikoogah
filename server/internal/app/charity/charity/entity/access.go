package entity

import "fmt"

type CharityAccess string

const (
	ViewCharityAK = "ak.charity.view"
	EditCharityAK = "ak.charity.modify"

	AddNewRallyAK = "ak.rally.new"
	EditRallyAK   = "ak.rally.modify"

	ViewParticipationAK   = "ak.participation.view"
	ManageParticipationAK = "ak.participation.manage"
)

func MapAccessKey(AKstr string) (CharityAccess, error) {
	switch AKstr {
	case "ak.charity.view":
		return ViewCharityAK, nil
	case "ak.charity.modify":
		return EditCharityAK, nil
	case "ak.rally.new":
		return AddNewRallyAK, nil
	case "ak.rally.modify":
		return EditRallyAK, nil
	case "ak.participation.view":
		return ViewParticipationAK, nil
	case "ak.participation.manage":
		return ManageParticipationAK, nil
	default:
		return "", fmt.Errorf("access key does not exist for this resource")
	}
}

var roleAccess map[RepresentativeRole][]CharityAccess = map[RepresentativeRole][]CharityAccess{
	Manager: {
		ViewCharityAK,
		EditCharityAK,
		AddNewRallyAK,
		EditRallyAK,
		ViewParticipationAK,
		ManageParticipationAK,
	},
	Employee: {
		ViewCharityAK,
		AddNewRallyAK,
		ViewParticipationAK,
		ManageParticipationAK,
	},
}

func GetRoleAccess(role RepresentativeRole) []CharityAccess {
	aks := roleAccess[role]
	aksCopy := make([]CharityAccess, len(aks))

	copy(aksCopy, aks)

	return aksCopy
}
