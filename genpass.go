package genpass

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	"io"
	"math/big"
	"strconv"
)

var (
	LCGm int64 = 4294967296
	LCGa int64 = 1103515245
	LCGb int64 = 12345
)

const (
	CharUpper  = 1 << iota //A B C D E F G H I J K L M N O P Q R S T U V W X Y Z
	CharLower              //a b c d e f g h i j k l m n o p q r s t u v w x y z
	CharDigit              //0 1 2 3 4 5 6 7 8 9
	CharPunct              //! " # $ % & ' ( ) * + , - . / : ; < = > ? @ [ \ ] ^ _ ` { | } ~.
	CharXdigit             //0 1 2 3 4 5 6 7 8 9 A B C D E F a b c d e f

	CharAlpha = CharLower | CharUpper
	CharAlnum = CharAlpha | CharDigit
	CharGraph = CharAlnum | CharPunct

	_DEFAULT_PRIMARY     = "passwd"
	_DEFAULT_DESCRIPTION = "www.google.com"
)

func genHash(method string) (checksum hash.Hash) {
	if method == "md5" {
		checksum = md5.New()
	}
	if method == "sha1" {
		checksum = sha1.New()
	}
	if method == "sha256" {
		checksum = sha256.New()
	}
	if method == "sha512" {
		checksum = sha512.New()
	}
	return
}

func genSeed(method, primary, desc string, rev int) *big.Int {
	i := new(big.Int)
	t := genHash(method)
	//t := sha256.New()
	io.WriteString(t, primary+desc+strconv.Itoa(rev))
	s := fmt.Sprintf("%x", t.Sum(nil))
	i.SetString(s, 16)
	return i
}

//linear congruential generator
func lcg(seed *big.Int) *big.Int {
	i := seed
	i.Mul(i, big.NewInt(LCGa))
	i.Add(i, big.NewInt(LCGb))
	i.Mod(i, big.NewInt(LCGm))
	return i
}

func genPwd(seed *big.Int, size int) []int64 {
	var val []int64

	now := lcg(seed)
	for i := 0; i < size; i++ {
		val = append(val, now.Int64())
		now = lcg(now)
	}
	return val
}

func fmtPwd(pwd []int64, cflag int, extra string) string {
	var chars, npwd string
	if cflag&(CharUpper|CharLower|CharDigit|CharPunct) != 0 {
		if cflag&CharUpper != 0 {
			chars += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		}
		if cflag&(CharLower) != 0 {
			chars += "abcdefghijklmnopqrstuvwxyz"
		}
		if cflag&(CharDigit) != 0 {
			chars += "0123456789"
		}
		if cflag&(CharPunct) != 0 {
			chars += `!"#$%&'()*+,-./:;<=>?@[\]^_`
			chars += "`{|}~."
		}
	}
	if cflag&CharXdigit != 0 {
		chars += "0123456789ABCDEFabcdef"
	}

	if len(extra) > 0 {
		chars += extra
	}

	length := int64(len(chars))
	for _, v := range pwd {
		npwd += string(chars[v%length])
	}
	return npwd
}

func Gen(options ...Options) string {
	opt := prepareOptions(options)
	seed := genSeed(opt.Method, opt.Primary, opt.Description, opt.Reversion)
	pwd := genPwd(seed, opt.Length)
	result := fmtPwd(pwd, opt.Flag, opt.ExtraChars)
	return result
}

type Options struct {
	Primary     string
	Description string
	ExtraChars  string
	Method      string
	Flag        int
	Reversion   int
	Length      int
}

func prepareOptions(options []Options) Options {
	var opt Options
	if len(options) > 0 {
		opt = options[0]
	}

	if len(opt.Primary) == 0 {
		opt.Primary = _DEFAULT_PRIMARY
	}

	if len(opt.Description) == 0 {
		opt.Description = _DEFAULT_DESCRIPTION
	}
	if len(opt.Method) == 0 {
		opt.Method = "sha256"
	}
	if !(opt.Method == "md5" || opt.Method == "sha1" || opt.Method == "sha256" || opt.Method == "sha512") {
		opt.Method = "sha256"
	}
	if opt.Flag == 0 {
		opt.Flag = CharAlnum
	}
	if opt.Length == 0 {
		opt.Length = 8
	}
	return opt
}
