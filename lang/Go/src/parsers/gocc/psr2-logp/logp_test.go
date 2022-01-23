package calc

import (
	"testing"

	"github.com/suntong/lang/lang/Go/src/parsers/gocc/psr2-logp/lexer"
	"github.com/suntong/lang/lang/Go/src/parsers/gocc/psr2-logp/parser"
)

type TI struct {
	src    string
	expect string
}

var testData = []*TI{
	{`2022-01-18 11:19:41.6007 HRImportJob Running job
2022-01-18 11:21:24.8027 Start of Post Processing 1/18/2022 11
2022-01-18 11:21:24.8027 Start auto generation of data
2022-01-18 11:21:24.8027 End auto generation of data
2022-01-18 11:21:24.8199 Start HR core import SQL
2022-01-18 11:33:21.9885 End HR core import SQL
2022-01-18 11:33:21.9885 Start Process Non Active Employees
2022-01-18 11:33:21.9885 End Process Non Active Employees
2022-01-18 11:33:23.9087 Start Updating Employee Management Structure
2022-01-18 11:33:40.8774 End Updating Employee Management Structure
2022-01-18 11:33:40.8774 Start Salt Passwords
2022-01-18 11:35:13.4284 End Salt Passwords
2022-01-18 11:35:13.4445 Start Pay Card Processing
2022-01-18 11:35:13.5237 End Pay Card Processing
2022-01-18 11:35:13.5237 Start User Account Check
2022-01-18 11:35:13.6597 End User Account Check
2022-01-18 11:35:13.6597 Start Legal Entity Tax Processing
2022-01-18 11:36:24.4468 End Legal Entity Tax Processing
2022-01-18 11:36:24.4468 Start Newly Active Employee Processing
2022-01-18 11:36:24.4554 End Newly Active Employee Processing
2022-01-18 11:36:24.7238 End of Post Processing 1/18/2022 11
2022-01-18 11:36:24.9746 HRImportJob Job completed
`, `2022-01-18 11:19:41.6007 HRImportJob Running job
 ==> 2022-01-18 11:36:24.9746 
2022-01-18 11:21:24.8027 Start of Post Processing 1/18/2022 11
 ==> 2022-01-18 11:36:24.7238 
2022-01-18 11:21:24.8027 Start auto generation of data
 ==> 2022-01-18 11:21:24.8027 `},
}

func Test1(t *testing.T) {
	p := parser.NewParser()
	for _, ts := range testData {
		s := lexer.NewLexer([]byte(ts.src))
		ret, err := p.Parse(s)
		if err != nil {
			t.Error(err)
		}
		if ret != ts.expect {
			t.Errorf("Got: %s\nExpected: %s\n", ret, ts.expect)
		}
	}
}
