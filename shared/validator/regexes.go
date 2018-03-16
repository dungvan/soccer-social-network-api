package validator

import "regexp"

const (
	// imageNameRegexString validate request im_name
	imageNameRegexString    = "^[a-zA-Z0-9._:-]+$"
	hashtagRegexString      = "^[Ａ-Ｚａ-ｚA-Za-z一-鿆0-9０-９ぁ-ゖゝ-ゟァ-ヺー-ヿㇰ-ㇿｦ-ﾝー_]+$"
	sourceIMNameRegexString = "^[a-zA-Z0-9_\\.]{1,256}(\\?w=[0-9]{1,5}\\&h=[0-9]{1,5})?$"
)

var (
	imageNameRegex    = regexp.MustCompile(imageNameRegexString)
	hashtagRegex      = regexp.MustCompile(hashtagRegexString)
	sourceIMNameRegex = regexp.MustCompile(sourceIMNameRegexString)
)
