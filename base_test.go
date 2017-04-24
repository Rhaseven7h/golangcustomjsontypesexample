package golangcustomjsontypesexample_test

import (
	"encoding/json"
	"fmt"
	"testing"

	gcjte "github.com/Rhaseven7h/golangcustomjsontypesexample"
	. "github.com/smartystreets/goconvey/convey"
)

func TestJSONSpike(t *testing.T) {
	Convey("Given a Profession Type", t, func() {
		Convey("When converted to string", func() {
			doctor := gcjte.ProfessionDoctor
			Convey("Then it should return a string representation", func() {
				s := doctor.String()
				So(s, ShouldEqual, "A Fine Doctor")
			})
		})
		Convey("When marshalling a profession", func() {
			strReps := map[string]gcjte.Profession{
				`"A Fine Doctor"`:         gcjte.ProfessionDoctor,
				`"A Sweet Engineer"`:      gcjte.ProfessionEngineer,
				`"An Awful Lawyer"`:       gcjte.ProfessionLawyer,
				`"A Smart Mathematician"`: gcjte.ProfessionMathematician,
				`"A Brilliant Physicist"`: gcjte.ProfessionPhysicist,
			}
			for str, prof := range strReps {
				s, err := json.Marshal(prof)
				Convey(fmt.Sprintf("Then we should get no error, and a valid string representation of a %s", str), func() {
					So(err, ShouldBeNil)
					So(string(s), ShouldEqual, str)
				})
			}
		})
		Convey("When unmarshalling a valid profession", func() {
			strReps := map[string]gcjte.Profession{
				`"A Fine Doctor"`:         gcjte.ProfessionDoctor,
				`"A Sweet Engineer"`:      gcjte.ProfessionEngineer,
				`"An Awful Lawyer"`:       gcjte.ProfessionLawyer,
				`"A Smart Mathematician"`: gcjte.ProfessionMathematician,
				`"A Brilliant Physicist"`: gcjte.ProfessionPhysicist,
			}
			for str, prof := range strReps {
				var v gcjte.Profession
				err := json.Unmarshal([]byte(str), &v)
				Convey(fmt.Sprintf("Then it should recognize a %s", str), func() {
					So(err, ShouldBeNil)
					So(v, ShouldEqual, prof)
				})
			}
		})
		Convey("When unmarshalling an invalid profession string", func() {
			var v gcjte.Profession
			err := json.Unmarshal([]byte(`"something invalid"`), &v)
			Convey("Then it should give an error, and announce unknown profession", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "unknown profession")
			})
		})
		Convey("When unmarshalling an invalid profession number", func() {
			var v gcjte.Profession
			err := json.Unmarshal([]byte(`45`), &v)
			Convey("Then it should give an error, and announce unknown profession", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "unknown profession")
			})
		})
		Convey("When unmarshalling an invalid profession boolean value", func() {
			var v gcjte.Profession
			err := json.Unmarshal([]byte(`true`), &v)
			Convey("Then it should give an error, and announce unknown profession", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "unknown profession")
			})
		})
		Convey("And given a struct using a Profession type", func() {
			type SampleUser struct {
				Name     string           `json:"name"`
				Activity gcjte.Profession `json:"activity"`
			}
			sampleUser1 := &SampleUser{
				Name:     "John Doe",
				Activity: gcjte.ProfessionMathematician,
			}
			Convey("When struct is marshalled", func() {
				b, err := json.Marshal(sampleUser1)
				Convey("Then we should get a correctly marshalled string", func() {
					So(err, ShouldBeNil)
					So(string(b), ShouldEqual, `{"name":"John Doe","activity":"A Smart Mathematician"}`)
				})
			})
			Convey("When a json using it is unmarshalled", func() {
				var sampleUser2 SampleUser
				err := json.Unmarshal([]byte(`{"name":"Jane Doe","activity":"An Awful Lawyer"}`), &sampleUser2)
				So(err, ShouldBeNil)
				So(sampleUser2.Name, ShouldEqual, "Jane Doe")
				So(sampleUser2.Activity, ShouldEqual, gcjte.ProfessionLawyer)
			})
		})
	})
}
