package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"time"
)

func basicEncodedDecoded() {
	var data string = "Verdi"

	var encoded string = base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("Encoded:", encoded)

	var decoded, _ = base64.StdEncoding.DecodeString(encoded)
	fmt.Println("Decoded:", string(decoded))
}

func encodedDecodedLen() {
	var data string = "Verdiansyah"

	var encoded = make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(encoded, []byte(data))
	var encodedStr = string(encoded)
	fmt.Println("Encoded:", encodedStr)
	//
	var decoded = make([]byte, base64.StdEncoding.DecodedLen(len(encoded)))
	_, err := base64.StdEncoding.Decode(decoded, encoded)
	if err != nil {
		fmt.Println("Error Decoded:", err)
	}

	var decodedStr = string(decoded)
	fmt.Println("Decoded:", decodedStr)
}

func encodedNDecodedUrl() {
	var url string = "http://localhost:4545/api/user/login"

	var encoded = base64.URLEncoding.EncodeToString([]byte(url))
	fmt.Println("Encoded:", encoded)

	var decoded, err = base64.URLEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error Decoded:", err)
	}

	var decodedStr = string(decoded)
	fmt.Println("Decoded:", decodedStr)
}

func useSha() {
	var data string = "Secret key: ABCDE"

	var sha = sha1.New()
	sha.Write([]byte(data))
	var encrypted = sha.Sum([]byte("123")) // random for key, normally fill nil value
	var encryptedStr = fmt.Sprintf("%x", encrypted)

	fmt.Println("Original:", data)
	fmt.Println("Encrypted:", encryptedStr)

}

func doSalting(text string) (string, string) {
	var salt = fmt.Sprintf("%d", time.Now().UnixNano())
	var saltedTxt = fmt.Sprintf("text: %s, salt: %s", text, salt)
	fmt.Println(saltedTxt)

	var sha = sha1.New()
	sha.Write([]byte(saltedTxt))
	var encrypted = sha.Sum(nil)
	var encryptedStr = fmt.Sprintf("%x", encrypted)

	return encryptedStr, salt
}

func useSalting() {
	var data string = "Verdiansyah is developer"
	fmt.Println("Original :", data)

	hashed1, salt1 := doSalting(data)
	fmt.Println("Hashed 1 :", hashed1)

	hashed2, salt2 := doSalting(data)
	fmt.Println("Hashed 2 :", hashed2)

	hashed3, salt3 := doSalting(data)
	fmt.Println("Hashed 3 :", hashed3)

	_, _, _ = salt1, salt2, salt3
}

func main() {
	//basicEncodedDecoded()
	//encodedDecodedLen()
	//encodedNDecodedUrl()
	//useSha()
	useSalting()
}
