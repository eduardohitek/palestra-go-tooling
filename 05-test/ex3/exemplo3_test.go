package main

import "testing"

func BenchmarkPalindromeFromMiddle_Banana(b *testing.B) {
	for n := 0; n < b.N; n++ {
		palindromeFromMiddle("banana")
	}
}
func BenchmarkPalindromeReverse_Banana(b *testing.B) {
	for n := 0; n < b.N; n++ {
		palindromeReverse("banana")
	}
}
func BenchmarkPalindromeFromEnd_Banana(b *testing.B) {
	for n := 0; n < b.N; n++ {
		palindromeFromEnd("banana")
	}
}

func BenchmarkPalindromeFromMiddle_civic(b *testing.B) {
	for n := 0; n < b.N; n++ {
		palindromeFromMiddle("civic")
	}
}
func BenchmarkPalindromeReverse_civic(b *testing.B) {
	for n := 0; n < b.N; n++ {
		palindromeReverse("civic")
	}
}
func BenchmarkPalindromeFromEnd_civic(b *testing.B) {
	for n := 0; n < b.N; n++ {
		palindromeFromEnd("civic")
	}
}

func BenchmarkPalindromeFromMiddle_rodador(b *testing.B) {
	for n := 0; n < b.N; n++ {
		palindromeFromMiddle("rodador")
	}
}
func BenchmarkPalindromeReverse_rodador(b *testing.B) {
	for n := 0; n < b.N; n++ {
		palindromeReverse("rodador")
	}
}
func BenchmarkPalindromeFromEnd_rodador(b *testing.B) {
	for n := 0; n < b.N; n++ {
		palindromeFromEnd("rodador")
	}
}
func BenchmarkPalindromeFromMiddle_nomelgibsonisacasinosbiglemon(b *testing.B) {
	for n := 0; n < b.N; n++ {
		palindromeFromMiddle("nomelgibsonisacasinosbiglemon")
	}
}
func BenchmarkPalindromeReverse_nomelgibsonisacasinosbiglemon(b *testing.B) {
	for n := 0; n < b.N; n++ {
		palindromeReverse("nomelgibsonisacasinosbiglemon")
	}
}
func BenchmarkPalindromeFromEnd_nomelgibsonisacasinosbiglemon(b *testing.B) {
	for n := 0; n < b.N; n++ {
		palindromeFromEnd("nomelgibsonisacasinosbiglemon")
	}
}

func BenchmarkPalindromeFromMiddle_longstring(b *testing.B) {
	for n := 0; n < b.N; n++ {
		palindromeFromMiddle("sdjfasdkflajkdslf;jaks;dlfjka;lsdjkf;lasjkdfl;ajksdlfjkalsdjkflajsdfnajsdfhquweurqjepfojiaopdsfjiaosdhufpiasdufiansdfuahsdfpuqpweufipadsnfasdhfuiaoshdufoaudfhauodfhuasdfnaiushdfuiaphfuipadnfaisudhfpadsufiahudfiauhdfiahsdfuiasuhdfpiadufnaiduhfiaduhfiahudfpiaudnciahdifphuadpfauhdpfauhpdfiaupfahusdfiphaupdifhuaipsdhufpaiudsfpihaupqiwehu9qwe8f9qwe8f9q8we9fhsdifhaidufiauhdfiahusdifuahidfuiasdufiaudfoiofduaifudsaiufdihaufidsuhaifdhuaifudiahfidshf9ew8q9f8ewq9f8ewq9uhewiqpuahipfsduiapfuhdspiauhfidpuahpifdsuhafpuaifdphuafpdhuafpdauhpfidhaicnduaipfduhaifhudaifhudianfudaipfdhusaiufdshaifdhuaifduhaifusdapfhdusiafndapiufhpaiufdhsuianfdsauhfdouahfduaofudhsoaiufhdsafnsdapifuewpqupfdshaufdsnaifudsaipfuhdsoaijfsdpoaijofpejqruewuqhfdsjanfdsjalfkjdslakjfldskja;lfdkjsal;fkjdsl;akjfld;skaj;flsdkjalfkdsafjds")
	}
}
func BenchmarkPalindromeReverse_longstring(b *testing.B) {
	for n := 0; n < b.N; n++ {
		palindromeReverse("sdjfasdkflajkdslf;jaks;dlfjka;lsdjkf;lasjkdfl;ajksdlfjkalsdjkflajsdfnajsdfhquweurqjepfojiaopdsfjiaosdhufpiasdufiansdfuahsdfpuqpweufipadsnfasdhfuiaoshdufoaudfhauodfhuasdfnaiushdfuiaphfuipadnfaisudhfpadsufiahudfiauhdfiahsdfuiasuhdfpiadufnaiduhfiaduhfiahudfpiaudnciahdifphuadpfauhdpfauhpdfiaupfahusdfiphaupdifhuaipsdhufpaiudsfpihaupqiwehu9qwe8f9qwe8f9q8we9fhsdifhaidufiauhdfiahusdifuahidfuiasdufiaudfoiofduaifudsaiufdihaufidsuhaifdhuaifudiahfidshf9ew8q9f8ewq9f8ewq9uhewiqpuahipfsduiapfuhdspiauhfidpuahpifdsuhafpuaifdphuafpdhuafpdauhpfidhaicnduaipfduhaifhudaifhudianfudaipfdhusaiufdshaifdhuaifduhaifusdapfhdusiafndapiufhpaiufdhsuianfdsauhfdouahfduaofudhsoaiufhdsafnsdapifuewpqupfdshaufdsnaifudsaipfuhdsoaijfsdpoaijofpejqruewuqhfdsjanfdsjalfkjdslakjfldskja;lfdkjsal;fkjdsl;akjfld;skaj;flsdkjalfkdsafjds")
	}
}
func BenchmarkPalindromeFromEnd_longstring(b *testing.B) {
	for n := 0; n < b.N; n++ {
		palindromeFromEnd("sdjfasdkflajkdslf;jaks;dlfjka;lsdjkf;lasjkdfl;ajksdlfjkalsdjkflajsdfnajsdfhquweurqjepfojiaopdsfjiaosdhufpiasdufiansdfuahsdfpuqpweufipadsnfasdhfuiaoshdufoaudfhauodfhuasdfnaiushdfuiaphfuipadnfaisudhfpadsufiahudfiauhdfiahsdfuiasuhdfpiadufnaiduhfiaduhfiahudfpiaudnciahdifphuadpfauhdpfauhpdfiaupfahusdfiphaupdifhuaipsdhufpaiudsfpihaupqiwehu9qwe8f9qwe8f9q8we9fhsdifhaidufiauhdfiahusdifuahidfuiasdufiaudfoiofduaifudsaiufdihaufidsuhaifdhuaifudiahfidshf9ew8q9f8ewq9f8ewq9uhewiqpuahipfsduiapfuhdspiauhfidpuahpifdsuhafpuaifdphuafpdhuafpdauhpfidhaicnduaipfduhaifhudaifhudianfudaipfdhusaiufdshaifdhuaifduhaifusdapfhdusiafndapiufhpaiufdhsuianfdsauhfdouahfduaofudhsoaiufhdsafnsdapifuewpqupfdshaufdsnaifudsaipfuhdsoaijfsdpoaijofpejqruewuqhfdsjanfdsjalfkjdslakjfldskja;lfdkjsal;fkjdsl;akjfld;skaj;flsdkjalfkdsafjds")
	}
}
