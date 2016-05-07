package response

type CommandResponseXml struct {
	Results []struct {
		Properties []struct {
			Name  string `xml:"NAME,attr"`
			Type  string `xml:"TYPE,attr"`
			Value string `xml:"VALUE"`
		} `xml:"PROPERTY"`
	} `xml:"RESULTS>CIM>INSTANCE"`
}
