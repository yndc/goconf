package data

type Validation struct {
	// int with minimum/maximum
	Intmin10_ok_ int `minimum:"10"`
	Intmin10_ok2 int `minimum:"10"`
	Intmin10_ok3 int `minimum:"10"`
	Intmin10_x__ int `minimum:"10"`
	Intmin10_x2_ int `minimum:"10"`
	Intmin10_x3_ int `minimum:"10"`
	Intmax30_ok_ int `maximum:"30"`
	Intmax30_ok2 int `maximum:"30"`
	Intmax30_ok3 int `maximum:"30"`
	Intmax30_x__ int `exclusiveMaximum:"30"`

	// string with minimum/maximum length
	Strmin3_ok_      string `minLength:"3"`
	Strmin3_x__      string `minLength:"3"`
	Strmin0_ok_      string `minLength:"0"`
	Strmax3_ok_      string `maxLength:"3"`
	Strmax3_ok2      string `maxLength:"3"`
	Strmax3_x__      string `maxLength:"3"`
	Strmin3max10_ok_ string `minLength:"3" maxLength:"10"`
	Strmin3max10_x__ string `minLength:"3" maxLength:"10"`
	Strmin3max10_x2_ string `minLength:"3" maxLength:"10"`
	Strmin2max2_ok_  string `minLength:"2" maxLength:"2"`
	Strmin2max2_x__  string `minLength:"2" maxLength:"2"`

	// string with format
	Strfmtip_ok_ string `format:"ip"`
	Strfmtip_ok2 string `format:"ip"`
	Strfmtip_ok3 string `format:"ip"`
	Strfmtip_x__ string `format:"ip"`
	Strfmtip_x2_ string `format:"ip"`
	Strfmtip_x3_ string `format:"ip"`

	// email addresses
	Strfmt_email_ok_ string `format:"email"`
	Strfmt_email_ok2 string `format:"email"`
	Strfmt_email_x__ string `format:"email"`
	Strfmt_email_x2_ string `format:"email"`

	// hostnames
	Strhost_ok_ string `format:"hostname"`
	Strhost_ok2 string `format:"hostname"`
	Strhost_ok3 string `format:"hostname"`
	Strhost_x__ string `format:"hostname"`
	Strhost_x2_ string `format:"hostname"`

	// URL
	Strurl_ok_ string `format:"uri"`
	Strurl_ok2 string `format:"uri"`
	Strurl_ok3 string `format:"uri"`
	Strurl_ok4 string `format:"uri"`
	Strurl_ok5 string `format:"uri"`
	Strurl_ok6 string `format:"uri"`
	Strurl_ok7 string `format:"uri"`
	Strurl_x__ string `format:"uri"`
	Strurl_x2_ string `format:"uri"`
	Strurl_x3_ string `format:"uri"`

	// pattern
	Strpattern_ok_ string `pattern:"((\\d{3})(?:\\.|-))?(\\d{3})(?:\\.|-)(\\d{4})"`
}
