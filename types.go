package exactonline_xml

import (
	"encoding/xml"
)

// <Topics>
// 	<Topic code="MatchSets" ts_d="0x000000018A976F80" count="287" pagesize="1000" />
// </Topics>

type Topics []Topic

func (tt *Topics) UnmarshalXML(e *xml.Decoder, start xml.StartElement) error {
	payload := struct {
		Topics []Topic `xml:"Topic"`
	}{}

	err := e.DecodeElement(&payload, &start)
	if err != nil {
		return err
	}

	*tt = payload.Topics
	return nil
}

type Topic struct {
	Code     string `xml:"code,attr"`
	TSD      string `xml:"ts_d,attr"`
	Count    int    `xml:"count,attr"`
	PageSize int    `xml:"pagesize,attr"`
}

type Messages []Message

func (mm *Messages) UnmarshalXML(e *xml.Decoder, start xml.StartElement) error {
	payload := struct {
		Messages []Message `xml:"Message"`
	}{}

	err := e.DecodeElement(&payload, &start)
	if err != nil {
		return err
	}

	*mm = payload.Messages
	return nil
}

type Message struct {
}

// <MatchSets>
// 	<MatchSet>
// 		<GLAccount code="1300" />
// 		<Account code="1433" />
// 		<MatchLines>
// 			<MatchLine finyear="2019" finperiod="1" journal="60" entry="70174" amountdc="256.52" />
// 			<MatchLine finyear="2019" finperiod="2" journal="20" entry="19200012" amountdc="-256.52" />
// 		</MatchLines>
// 	</MatchSet>
// </MatchSets>

type MatchSets []MatchSet

func (ss *MatchSets) UnmarshalXML(e *xml.Decoder, start xml.StartElement) error {
	payload := struct {
		MatchSet []MatchSet `xml:"MatchSet"`
	}{}

	err := e.DecodeElement(&payload, &start)
	if err != nil {
		return err
	}

	*ss = payload.MatchSet
	return nil
}

type MatchSet struct {
	GLAccountCode string
	AccountCode   string
	MatchLines    MatchLines
}

func (ms *MatchSet) UnmarshalXML(e *xml.Decoder, start xml.StartElement) error {
	payload := struct {
		GLAccount struct {
			Code string `xml:"code,attr"`
		} `xml:"GLAccount"`
		Account struct {
			Code string `xml:"code,attr"`
		} `xml:"Account"`
		MatchLines MatchLines `xml:"MatchLines`
	}{}

	err := e.DecodeElement(&payload, &start)
	if err != nil {
		return err
	}

	ms.GLAccountCode = payload.GLAccount.Code
	ms.AccountCode = payload.Account.Code
	ms.MatchLines = payload.MatchLines
	return nil
}

type MatchLines []MatchLine

func (ll *MatchLines) UnmarshalXML(e *xml.Decoder, start xml.StartElement) error {
	payload := struct {
		MatchLines []MatchLine `xml:"MatchLine"`
	}{}

	err := e.DecodeElement(&payload, &start)
	if err != nil {
		return err
	}

	*ll = payload.MatchLines
	return nil
}

type MatchLine struct {
	FinYear   int     `xml:"finyear,attr"`
	FinPeriod int     `xml:"finperiod,attr"`
	Journal   string  `xml:"journal,attr"`
	Entry     int     `xml:"entry,attr"`
	AmountDC  float64 `xml:"amountdc,attr"`
}
