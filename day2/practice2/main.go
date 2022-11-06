package main

func main() {
	s := "寿限無、寿限無、 五劫の擦り切れ、 海砂利水魚の、 水行末・雲来末・風来末、 喰う寝る処に住む処、 藪ら柑子の藪柑子、 <改行コード> パイポ・パイポ・パイポのシューリンガン、シューリンガンのグーリンダイ、 グーリンダイのポンポコピーのポンポコナの、 長久命の長助"
	wordForNewLine := "<改行コード>"
	for i := 0; i < len(s); i++ {
		if s[i:i+1] == "<" {
			for j := 0; j < len(wordForNewLine); j++ {
				i++
			}
			print("\n")
		}
		print(s[i : i+1])
	}

	println()
	println("--------------")

	ss := "寿限無、寿限無、 五劫の擦り切れ、 海砂利水魚の、 水行末・雲来末・風来末、 喰う寝る処に住む処、 藪ら柑子の藪柑子、 \n パイポ・パイポ・パイポのシューリンガン、シューリンガンのグーリンダイ、 グーリンダイのポンポコピーのポンポコナの、 長久命の長助"
	wordForNewLine2 := "\\n"
	for i := 0; i < len(ss); i++ {
		if ss[i:i+1] == "\\" {
			for j := 0; j < len(wordForNewLine2); j++ {
				i++
			}
			print("\n")
		}
		print(ss[i : i+1])
	}
}
