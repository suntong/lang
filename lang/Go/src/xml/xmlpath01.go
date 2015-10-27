package main

import (
	"fmt"
	"strings"
)

import "gopkg.in/xmlpath.v1"

func main() {

	data := `<soapenv:Envelope>
  <soapenv:Body>
    <getPosHistoryByGroupResponse>
      <getPosHistoryByGroupReturn>
        <item>
          <licencePlate>
            <licencePlateNumber>AA 3135</licencePlateNumber>
            <nationality>DK</nationality>
          </licencePlate>
          <posHistories>
            <item>
              <dateTime>2013-12-22T04:47:24.000Z</dateTime>
              <ebs>
                <bogieLoad xsi:nil="true"/>
                <totalDistance>176915750</totalDistance>
                <velocity>0</velocity>
              </ebs>
              <position>
                <direction>182</direction>
                <distance>11.4</distance>
                <latitude>44.545295</latitude>
                <location>M?dena, Italy</location>
                <longitude>10.921779</longitude>
              </position>
              <powerSupply>
                <isCoupled>true</isCoupled>
                <isDoorOpen>false</isDoorOpen>
                <isIgnitionOn>false</isIgnitionOn>
                <powerSupplyVoltage>14000</powerSupplyVoltage>
              </powerSupply>
            </item>
            <item>
              <dateTime>2013-12-22T04:32:22.000Z</dateTime>
              <ebs>
                <bogieLoad xsi:nil="true"/>
                <totalDistance>176915750</totalDistance>
                <velocity>0</velocity>
              </ebs>
              <position>
                <direction>182</direction>
                <distance>11.4</distance>
                <latitude>44.545308</latitude>
                <location>M?dena, Italy</location>
                <longitude>10.921697</longitude>
              </position>
              <powerSupply>
                <isCoupled>true</isCoupled>
                <isDoorOpen>false</isDoorOpen>
                <isIgnitionOn>false</isIgnitionOn>
                <powerSupplyVoltage>12600</powerSupplyVoltage>
              </powerSupply>
            </item>
          </posHistories>
        </item>
      </getPosHistoryByGroupReturn>
    </getPosHistoryByGroupResponse>
  </soapenv:Body>
</soapenv:Envelope>`

	root, err := xmlpath.Parse(strings.NewReader(data))
	if err != nil {
		fmt.Println(err)
		return
	}

	path := xmlpath.MustCompile("//licencePlateNumber")
	if value, ok := path.String(root); ok {
		fmt.Println("Found:", value)
	}

}
