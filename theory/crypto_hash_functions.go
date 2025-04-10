package theory

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
)

func CryptoToSHA256() {
	data := []byte("Какой то текст который нужно зашифровать")

	sh := sha256.New()
	sh.Write(data)
	cryptedData := sh.Sum(nil)

	fmt.Printf("data %x", cryptedData)

	//но можно так
	cryptedData2 := sha256.Sum256(data)
	fmt.Printf("data %x", cryptedData2)
}

func CryptoToMD5() {
	var (
		data  []byte         //слайс случайных байт
		hash1 []byte         //хеш с использованием hash.Hash
		hash2 [md5.Size]byte // хеш, возвращаемый функцией md5.Sum
	)

	data = make([]byte, 512)
	_, err := rand.Read(data)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	h := md5.New()
	h.Write(data)
	hash1 = h.Sum(nil)
	hash2 = md5.Sum(data)

	if bytes.Equal(hash1, hash2[:]) {
		fmt.Println("Все ок, хеши равны")
	} else {
		fmt.Println("Что то пошло не так")
	}
}

func generateRandomByteSlice(size int) ([]byte, error) {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func GenerateHashBasedMessageAuthenticationCode() {
	src := []byte("Тут какое то сообщение, которое нужно подписать")

	// генерируем случайную последовательность байт
	key, err := generateRandomByteSlice(16)
	if err != nil {
		log.Fatal(err)
		return
	}

	h := hmac.New(sha256.New, key)
	h.Write(src)
	dst := h.Sum(nil)

	fmt.Printf("%x", dst)
}

func DecodeWithHMAC() {
	secretKey := []byte("secret key")

	var (
		data []byte // декодированное сообщение с подписью
		id   uint32 // значение индификатора
		err  error
		sign []byte //HMAC-подпись от индификатора
	)

	msg := "048ff4ea240a9fdeac8f1422733e9f3b8b0291c969652225e25c5f0f9f8da654139c9e21"
	data, err = hex.DecodeString(msg)
	if err != nil {
		panic(err)
	}
	id = binary.BigEndian.Uint32(data[:4])
	h := hmac.New(sha256.New, secretKey)
	h.Write(data[:4])
	sign = h.Sum(nil)

	if hmac.Equal(sign, data[4:]) {
		fmt.Println("Подпись подлинная. ID:", id)
	} else {
		fmt.Println("Подпись неверна. Где-то ошибка")
	}
}

const (
	password = "x35k9f"
	msg      = `0ba7cd8c624345451df4710b81d1a349ce401e61bc7eb704ca` +
		`a84a8cde9f9959699f75d0d1075d676f1fe2eb475cf81f62ef` +
		`f701fee6a433cfd289d231440cf549e40b6c13d8843197a95f` +
		`8639911b7ed39a3aec4dfa9d286095c705e1a825b10a9104c6` +
		`be55d1079e6c6167118ac91318fe`
)

func DecodeGCM() {
	var key [32]byte = sha256.Sum256([]byte(password))
	aesblock, err := aes.NewCipher(key[:])
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	nonce := key[len(key)-aesgcm.NonceSize():]
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	encrypted, err := hex.DecodeString(msg)

	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	src2, err := aesgcm.Open(nil, nonce, encrypted, nil) // расшифровываем
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Printf("decrypted: %s\n", src2)
}
