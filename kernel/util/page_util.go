package util

import (
	"bytes"
	"strings"
)

const (
	beginFormElment = "<form name=\"punchout_form\" method=\"post\" action=\""
	endElment       = "\">\n"
	endForm         = "</form>\n"
	payInput        = "<input type=\"submit\" value=\"立即支付\" style=\"display:none\" >\n"
	payJs           = "<script>document.forms[0].submit();</script>"
	formMinLen      = len(beginFormElment) + len(endElment) + len(endForm) + len(payInput) + len(payJs)
)

func BuildForm(actionUrl string, parameters map[string]string) string {
	// refreshing value
	for k, v := range parameters {
		parameters[k] = strings.ReplaceAll(v, "\"", "&quot;")
	}
	ss := &strings.Builder{}
	ss.Grow(12)
	bb := &bytes.Buffer{}
	bb.Grow(buildHiddenFieldsLen(actionUrl, parameters))
	// form
	bb.WriteString(beginFormElment)
	bb.WriteString(actionUrl)
	bb.WriteString(endElment)
	// hidden input
	buildHiddenFields(bb, parameters)
	//pay input
	bb.WriteString(payInput)
	// end form
	bb.WriteString(endForm)
	// submit js
	bb.WriteString(payJs)
	return bb.String()
}

func buildHiddenFieldsLen(actionUrl string, parameters map[string]string) int {
	blen := 0
	// 39 hidden input len
	for k, v := range parameters {
		blen += (len(k) + len(v) + hiddenInputMinLen)
	}
	return formMinLen + len(actionUrl) + blen
}

const (
	beginHiddenInput  = "<input type=\"hidden\" name=\""
	valAttr           = "\" value=\""
	hiddenInputMinLen = len(beginHiddenInput) + len(valAttr) + len(endElment)
)

func buildHiddenFields(bb *bytes.Buffer, parameters map[string]string) {
	for k, v := range parameters {
		bb.WriteString(beginHiddenInput)
		bb.WriteString(k)
		bb.WriteString(valAttr)
		bb.WriteString(v)
		bb.WriteString(endElment)
	}
}
