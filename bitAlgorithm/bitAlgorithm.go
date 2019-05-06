package bitAlgorithm

import "strings"

 	var bit62 []rune = []rune{'a','b','c','d','e','f','g','h','i','j','k','l','m','n','o',
	'p','q','r','s','t','u','v','w','x','y','z','A','B','C','D','E','F','G','H','I','J','K',
	'L','M','N','O','P','Q','R','S','T','U','V','W','X','Y','Z',0,1,2,3,4,5,6,7,8,9} 

func IntToString(seq uint64) (string) {
	var charSeq  []rune
	if seq!=0{
		for seq!=0{
			mod:=seq%62
			div:=seq/62
			charSeq=append(charSeq,bit62[mod])
			seq=div
		}
	}else{
		charSeq=append(charSeq,bit62[0])
	}
	return reverse(charSeq)	
}

func StringToInt(s string) uint64{
	bit62String:=string(bit62)
	var result uint64 = 0
	for _,char:=range s{
		result=result*62+uint64(strings.Index(bit62String,string(char)))
	}
	return result
}

func reverse(r []rune) string{
	for i,j:=0,len(r)-1;i<j;i,j=i+1,j-1{
		r[i],r[j]=r[j],r[i]
	}
	return string(r)
}