package test

import (
	"testing"
	"github.com/a-h/generate/test/additionalProperties2_gen"
	"encoding/json"
	"log"
	"reflect"
)

func TestMarshalUnmarshal(t *testing.T) {
	params := []struct {
		Name string
		Strct additionalProperties2.AdditionalProperties
		Validation func(t *testing.T, prop *additionalProperties2.AdditionalProperties)
	} {
		{
			Name: "Base Object",
			Strct: additionalProperties2.AdditionalProperties{
				Property1: "test",
			},
			Validation: func(t *testing.T, prop *additionalProperties2.AdditionalProperties) {
				if prop.Property1 != "test" {
					t.Fatal("property1 != test")
				}
			},
		},
		{
			Name: "Property7",
			Strct: additionalProperties2.AdditionalProperties{
				Property7: &additionalProperties2.Property7 {
					StreetNumber: 69,
					StreetName: "Elm St",
					PoBox: &additionalProperties2.PoBox{
						Suburb: "Smallville",
					},
					AdditionalProperties: map[string]*additionalProperties2.Hairy {
						"red": {
							AdditionalProperties: map[string]*additionalProperties2.Anonymous1 {
								"blue": {
									Color: "green",
									Conditions: []*additionalProperties2.ConditionsItem{
										{Name: "dry"},
									},
									Density: 42.42,
								},
							},
						},
						"orange": {

						},
					},
				},
			},
			Validation: func(t *testing.T, prop *additionalProperties2.AdditionalProperties) {

				if prop.Property7.StreetNumber != 69 {
					t.Fatal("wrong value")
				}

				if len(prop.Property7.AdditionalProperties) != 2 {
					t.Fatal("not enough additionalProperties")
				}

				if prop.Property7.AdditionalProperties["red"].AdditionalProperties["blue"].Color != "green" {
					t.Fatal("wrong nested value")
				}

				if prop.Property7.AdditionalProperties["red"].AdditionalProperties["blue"].Density != 42.42 {
					t.Fatal("wrong nested value")
				}
			},
		},
	}

	for _, p := range params {
		if str, err := json.MarshalIndent(&p.Strct, "", "  "); err != nil {
			t.Fatal(err)
		} else {
			log.Println(string(str))
			strct2 := &additionalProperties2.AdditionalProperties{}
			if err := json.Unmarshal(str, &strct2); err != nil {
				t.Fatal(err)
			}

			if reflect.DeepEqual(p.Strct, strct2) {
				log.Fatal("unmarshaled struct != given struct")
			}

			p.Validation(t, strct2)
		}
	}
}
