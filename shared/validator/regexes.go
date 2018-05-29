package validator

import "regexp"

const (
	hashtagRegexString   = "^[Ａ-Ｚａ-ｚA-Za-z一-鿆0-9０-９ぁ-ゖゝ-ゟァ-ヺー-ヿㇰ-ㇿｦ-ﾝー_]+$"
	imageNameRegexString = "^[a-zA-Z0-9_\\/._:-]{1,256}(\\?w=[0-9]{1,5}\\&h=[0-9]{1,5})?$"
	videoNameRegexString = "^[a-zA-Z0-9_\\/._:-]{1,256}(\\?w=[0-9]{1,5}\\&h=[0-9]{1,5})?$"
)

var (
	hashtagRegex   = regexp.MustCompile(hashtagRegexString)
	imageNameRegex = regexp.MustCompile(imageNameRegexString)
	videoNameRegex = regexp.MustCompile(videoNameRegexString)
)
