package sparql

import (
	//"fmt"
	//"github.com/kr/pretty"
	"testing"
)

func TestQueryParse(t *testing.T) {
	for _, test := range []struct {
		str string
	}{
		{
			"SELECT ?ax WHERE { ?ax rdf:type brick:Room }",
		},
		{
			"SELECT ?x WHERE { ?x rdf:type <http://buildsys.org/ontologies/Brick#Room> }",
		},
		{
			"SELECT ?x WHERE { ?x ?y brick:Room }",
		},
		{
			"SELECT ?x ?y WHERE { ?x ?y brick:Room }",
		},
		{
			"SELECT ?x ?y WHERE { ?y rdf:type rdf:type . ?x ?y brick:Room }",
		},
		{
			"SELECT ?x WHERE { ?x rdf:type+ brick:Room }",
		},
		{
			"SELECT ?x ?y WHERE { ?y rdf:type/rdfs:subClassOf ?x }",
		},
		{
			"SELECT ?x ?y ?z WHERE { { ?y bf:isFedBy ?x } }",
		},
		{
			"SELECT ?x ?y ?z WHERE { { ?y bf:isFedBy ?x } UNION { ?y bf:feeds ?z } }",
		},
		{
			"SELECT ?x ?y ?z WHERE { { ?y bf:isFedBy ?x } UNION { ?y bf:feeds ?z } UNION { ?y bf:isPointOf ?x } }",
		},
		{
			"SELECT ?x ?y ?z WHERE { { ?y bf:isFedBy ?x . ?y bf:hasPoint ?z } UNION { ?y bf:feeds ?x } }",
		},
		{
			"SELECT ?x ?y ?z WHERE { ?y rdf:type brick:VAV { ?y bf:isFedBy ?x . ?y bf:hasPoint ?z } UNION { ?y bf:feeds ?x } }",
		},
	} {
		q, err := Parse(test.str)
		if err != nil {
			t.Error(err)
		}
		//fmt.Printf("%# v", pretty.Formatter(q))
		_ = q
	}
}

func TestInsertQueryParse(t *testing.T) {
	for _, test := range []struct {
		str string
	}{
		{
			"INSERT { ?x rdf:type brick:Location } WHERE { ?ax rdf:type brick:Room }",
		},
		{
			"INSERT { ?x rdf:type brick:Location . ?x bf:isLocatedIn brick:Floor } WHERE { ?ax rdf:type brick:Room }",
		},
		{
			"INSERT { ?x rdf:type brick:Location . ?x bf:isLocatedIn brick:Floor } TO ciee WHERE { ?ax rdf:type brick:Room }",
		},
		{
			"INSERT { ?x rdf:type ?y } WHERE { ?x rdf:type/rdfs:subClassOf* ?y }",
		},
	} {
		q, err := Parse(test.str)
		if err != nil {
			t.Error(err)
		}
		//fmt.Printf("%# v", pretty.Formatter(q))
		_ = q
	}
}
