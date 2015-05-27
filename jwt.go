package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"hash"
	"strings"
)

const (
	HS256 = "HS256"
	HS384 = "HS384"
	HS512 = "HS512"
)

var errAlgorithm = errors.New("Algorithm must be HS256, HS384 or HS512")

/*
Header:
{
  "alg": "HS256",
  "typ": "JWT"
}

Claims:

{
  "sub": "1234567890",
  "name": "John Doe",
  "admin": true
}

https://stormpath.com/blog/jwt-the-right-way

var headers = base64URLencode(Header);
var claims = base64URLencode(Claims);
var payload = header + "." + claims;
var signature = base64URLencode(HMACSHA256(payload, secret));

var encodedJWT = payload + "." + signature;

*/
/*
func setHeader(s string) ([]byte, error) {
	if s != HS256 || s != HS384 || s != HS512 {
		return []byte(""), errAlgorithm
	}

	header["alg"] = s
	b, _ := json.Marshal(header)
	return b, nil
}
*/

func ComputeHmac256(message, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func ComputeHmac384(message, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha512.New384, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func ComputeHmac512(message, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha512.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func ComputeHmac(alg, message, secret string) string {
	key := []byte(secret)

	var h hash.Hash

	if alg == HS256 {
		h = hmac.New(sha256.New, key)
	}

	if alg == HS384 {
		h = hmac.New(sha512.New384, key)
	}

	if alg == HS512 {
		h = hmac.New(sha512.New, key)
	}

	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func verify(alg, token, secret string) bool {
	//var payload = header + "." + claims;
	//var signature = base64URLencode(HMACSHA256(payload, secret));
	//var encodedJWT = payload + "." + signature;

	//token = header.claims.signature
	slcStr := strings.Split(token, ".")
	payload := slcStr[0] + "." + slcStr[1]

	hmacPayloadSecret := ComputeHmac(HS256, payload, "secret")

	signature := base64.URLEncoding.EncodeToString([]byte(hmacPayloadSecret))
	fmt.Println(signature)

	return token == payload+"."+signature
}

func main() {
	header := map[string]string{"alg": HS256, "typ": "JWT"}
	jsonHeader, _ := json.Marshal(header)

	claims := make(map[string]interface{})
	claims["sub"] = "1234567890"
	claims["name"] = "John Doe"
	claims["admin"] = true

	jsonClaims, _ := json.Marshal(claims)

	fmt.Println(string(jsonHeader))
	fmt.Println(string(jsonClaims))

	//var headers = base64URLencode(Header);
	b64header := base64.URLEncoding.EncodeToString(jsonHeader)

	//var claims = base64URLencode(Claims);
	b64claims := base64.URLEncoding.EncodeToString(jsonClaims)

	//var payload = header + "." + claims;
	payload := b64header + "." + b64claims

	//var signature = base64URLencode(HMACSHA256(payload, secret));
	hmacPayloadSecret := ComputeHmac(HS256, payload, "secret")

	fmt.Println("hmacPayloadSecret")
	fmt.Println(hmacPayloadSecret)

	signature := base64.URLEncoding.EncodeToString([]byte(hmacPayloadSecret))
	fmt.Println(signature)

	//var encodedJWT = payload + "." + signature;
	encodedJWT := payload + "." + signature

	fmt.Println(b64header)
	fmt.Println(b64claims)
	fmt.Println(payload)
	fmt.Println(signature)
	fmt.Println(encodedJWT)

	fmt.Println(verify(HS256, encodedJWT, "secret"))

	fmt.Println(ComputeHmac256("Message", "secret"))
	fmt.Println(ComputeHmac384("Message", "secret"))
	fmt.Println(ComputeHmac512("Message", "secret"))
}
