package theory

import (
	"fmt"
	"log"

	_ "crypto/rand"
	"encoding/base64"
	"math/rand"
)

func GenerateRandomNumbersBySeed() {
	//Нужен "math/rand"
	rand1 := rand.New(rand.NewSource(10))
	rand2 := rand.New(rand.NewSource(10))

	for i := 0; i < 5; i++ {
		val1 := rand1.Int() // возращает рандомное псевдослучайное положительное число
		val2 := rand2.Int()

		fmt.Printf("Rand1 %d, rand2 %d", val1, val2) // На выходе получится два одинаковых числа
	}
}

func GenerateRundomNumbersByCrypto() {
	//тут нужен "crypto/rand"
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err.Error())
	}

	//Кодируем полученный рандомом слайс байт в кодировку base64
	encoded := base64.StdEncoding.EncodeToString(b)
	fmt.Print(encoded)
}
