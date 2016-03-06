////////////////////////////////////////////////////////////////////////////
// Porgram: unmarshal02B.go
// Purpose: Go xml unmarshal demo, Unmarshal and retaining sub xml elements
// Authors: Tong Sun (c) 2015, All rights reserved
// Credits: https://golang.org/pkg/encoding/xml/
//          Mark Crook, https://play.golang.org/p/rSj5cs8yvR
////////////////////////////////////////////////////////////////////////////

/*

encoding/xml, how to get full xml string
https://groups.google.com/d/msg/golang-nuts/v9SDlW3kDeo/hR_YbON1DQAJ

Q: When parsing XML, say at the level of

   <QueryStringParameter Name="v" Value="SoftwareVersion" ... />

How I can get the full xml string ("<QueryStringParameter Name= ... />") at this level?

A:

That element is empty, it has no inner xml. If you want the raw xml, you have to go up one level.

You can create a type to get whatever attributes you need: https://play.golang.org/p/AyzmoUmVYS

If you want the raw data, or attribute that you don't know about ahead of time, you need to implement an xml.Unmarshaler

James Bardin

This demo shows how to go one level up.
To deal at the QueryStringParameter level, create a type with only attributes.

*/

package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Xml struct {
	Xml string `xml:",innerxml"`
}

const content = `
<Person>
	<Name>Lucie</Name>
	<Company>
		<Name>Example Inc.</Name>
		<Addr>Example work place</Addr>
	</Company>
    <Condition UniqueStringId="28d77bce-936c-4108-8a4d-94be2ab54af8">
      <ConditionalRule Classname="Microsoft.VisualStudio.TestTools.WebTesting.Rules.NumericalComparisonRule, Microsoft.VisualStudio.QualityTools.WebTestFramework, Version=10.0.0.0, Culture=neutral, PublicKeyToken=b03f5f7f11d50a3a" DisplayName="Number Comparison" Description="The condition is met when the value of the context parameter satisfies the comparison with the provided value.">
        <RuleParameters>
          <RuleParameter Name="ContextParameterName" Value="X" />
          <RuleParameter Name="ComparisonOperator" Value="==" />
          <RuleParameter Name="Value" Value="6" />
        </RuleParameters>
      </ConditionalRule>
    </Condition>
</Person>
`

type Person struct {
	Name      string
	Company   Xml
	Condition struct {
		ConditionalRule struct {
			RuleParameters Xml
		}
	}
}

func main() {
	var person Person

	err := xml.Unmarshal([]byte(content), &person)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s's company is %q\n", person.Name, person.Company.Xml)

	fmt.Printf("Condition is %q\n", person.Condition.ConditionalRule.RuleParameters.Xml)
}
