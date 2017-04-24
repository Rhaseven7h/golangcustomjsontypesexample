package golangcustomjsontypesexample

import "fmt"

// Profession type represnts a profession
type Profession string

// Profession constants
const (
	ProfessionDoctor        Profession = "A Fine Doctor"
	ProfessionEngineer      Profession = "A Sweet Engineer"
	ProfessionLawyer        Profession = "An Awful Lawyer"
	ProfessionMathematician Profession = "A Smart Mathematician"
	ProfessionPhysicist     Profession = "A Brilliant Physicist"
)

func (p Profession) String() string {
	return string(p)
}

// UnmarshalJSON adds support for JSON unmarshalling
func (p *Profession) UnmarshalJSON(b []byte) (err error) {
	strProfession := string(b)
	switch strProfession {
	case `"A Fine Doctor"`:
		*p = ProfessionDoctor
	case `"A Sweet Engineer"`:
		*p = ProfessionEngineer
	case `"An Awful Lawyer"`:
		*p = ProfessionLawyer
	case `"A Smart Mathematician"`:
		*p = ProfessionMathematician
	case `"A Brilliant Physicist"`:
		*p = ProfessionPhysicist
	default:
		err = fmt.Errorf("unknown profession '%s'", strProfession)
	}
	return
}

// MarshalJSON adds support for JSON marshalling
func (p Profession) MarshalJSON() (rep []byte, err error) {
	return []byte(fmt.Sprintf(`"%s"`, p)), nil
}
