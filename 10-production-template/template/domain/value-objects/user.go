package valueobjects

type UserLvl string

const (
	Regular           UserLvl = "REGULAR"
	SuperAdmin        UserLvl = "SUPER_ADMIN"
	OrganizationAdmin UserLvl = "ORGANIZATION_ADMIN"
)

type OwnUser struct {
	username       string
	firstName      string
	lastName       string
	id             string
	admin          bool
	userLvl        UserLvl
	organizationID string
}

func (o *OwnUser) Username() string {
	return o.username
}

func (o *OwnUser) FirstName() string {
	return o.firstName
}

func (o *OwnUser) LastName() string {
	return o.lastName
}

func (o *OwnUser) ID() string {
	return o.id
}

func (o *OwnUser) Admin() bool {
	return o.admin
}

func (o *OwnUser) UserLvl() UserLvl {
	return o.userLvl
}

func (o *OwnUser) OrganizationID() string {
	return o.organizationID
}

func NewOwnUser(
	username string,
	firstName string,
	lastName string,
	id string,
	admin bool,
	userLvl UserLvl,
	organizationID string,
) OwnUser {
	return OwnUser{
		username,
		firstName,
		lastName,
		id,
		admin,
		userLvl,
		organizationID,
	}
}
