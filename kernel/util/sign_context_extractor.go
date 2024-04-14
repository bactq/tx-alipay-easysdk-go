package util

import (
	"errors"
	"strings"

	"github.com/tianxinzizhen/tx-alipay-easysdk-go/kernel/AlipayConstants"
)

const (
	/**
	 * 左大括号
	 */
	LEFT_BRACE = '{'

	/**
	 * 右大括号
	 */
	RIGHT_BRACE = '}'

	/**
	 * 双引号
	 */
	DOUBLE_QUOTES = '"'
)

func GetSignSourceData(body, method string) string {
	rootNode := strings.ReplaceAll(method, ".", "_") + AlipayConstants.RESPONSE_SUFFIX
	errorRootNode := AlipayConstants.ERROR_RESPONSE
	indexOfRootNode := strings.Index(body, rootNode)
	indexOfErrorRoot := strings.Index(body, errorRootNode)
	if indexOfRootNode > 0 {
		return parseSignSourceData(body, rootNode, indexOfRootNode)
	} else if indexOfErrorRoot > 0 {
		return parseSignSourceData(body, errorRootNode, indexOfErrorRoot)
	} else {
		return ""
	}
}

func parseSignSourceData(body, rootNode string, indexOfRootNode int) string {
	indexOfSign := strings.Index(body, "\""+AlipayConstants.SIGN_FIELD+"\"")
	if indexOfSign < 0 {
		return ""
	}
	signDataStartIndex := indexOfRootNode + len(rootNode) + 2
	signSourceData := extractSignContent(body, signDataStartIndex)
	if strings.LastIndex(body, rootNode) > signSourceData.endIndex {
		panic(errors.New("检测到响应报文中有重复的" + rootNode + "，验签失败。"))
	}
	return signSourceData.sourceData
}

func extractSignContent(str string, begin int) *signSourceData {
	if len(str) == 0 {
		return nil
	}
	beginIndex := extractBeginPosition(str, begin)
	if beginIndex >= len(str) {
		return nil
	}
	endIndex := extractEndPosition(str, beginIndex)

	return &signSourceData{
		sourceData: str[beginIndex:endIndex],
		beginIndex: beginIndex,
		endIndex:   endIndex,
	}
}
func extractBeginPosition(responseString string, begin int) int {
	beginPosition := begin
	for beginPosition < len(responseString) &&
		responseString[beginPosition] != LEFT_BRACE &&
		responseString[beginPosition] != DOUBLE_QUOTES {
		beginPosition++
	}
	return beginPosition
}

func extractEndPosition(responseString string, beginPosition int) int {
	//提取明文验签内容终点
	if responseString[beginPosition] == LEFT_BRACE {
		return extractJsonObjectEndPosition(responseString, beginPosition)
	} else { //提取密文验签内容终点
		return extractJsonBase64ValueEndPosition(responseString, beginPosition)
	}
}
func extractJsonBase64ValueEndPosition(responseString string, beginPosition int) int {
	for i := beginPosition; i < len(responseString); i++ {
		if responseString[i] == DOUBLE_QUOTES && i != beginPosition {
			return i + 1
		}
	}
	return len(responseString)
}

func extractJsonObjectEndPosition(responseString string, beginPosition int) int {
	bracesDeep := 0
	//记录当前字符是否在双引号中
	inQuotes := false
	//记录当前字符前面连续的转义字符个数
	consecutiveEscapeCount := 0
	for i := beginPosition; i < len(responseString); i++ {
		currentChar := responseString[i]
		//如果当前字符是"且前面有偶数个转义标记（0也是偶数）
		if currentChar == DOUBLE_QUOTES && consecutiveEscapeCount%2 == 0 {
			//是否在引号中的状态取反
			inQuotes = !inQuotes
			//如果当前字符是{且不在引号中
		} else if currentChar == LEFT_BRACE && !inQuotes {
			//将该{加入未闭合括号中
			bracesDeep++
			//如果当前字符是}且不在引号中
		} else if currentChar == RIGHT_BRACE && !inQuotes {
			//弹出一个未闭合括号
			bracesDeep--
			//如果弹出后，未闭合括号为空，说明已经找到终点
			if bracesDeep == 0 {
				return i + 1
			}
		}
		//如果当前字符是转义字符
		if currentChar == '\\' {
			//连续转义字符个数+1
			consecutiveEscapeCount++
		} else {
			//连续转义字符个数置0
			consecutiveEscapeCount = 0
		}
	}
	return len(responseString)
}

type signSourceData struct {
	/**
	 * 待验签原始内容
	 */
	sourceData string
	/**
	 * 待验签原始内容在响应字符串中的起始位置
	 */
	beginIndex int
	/**
	 * 待验签原始内容在响应字符串中的结束位置
	 */
	endIndex int
}
