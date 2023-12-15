package fizzbuzz

import "strconv"

/*
โจทก์วันนี้ที่เราจะเขียนโปรแกรมกันก็คือ โจทก์ปัญหา FizzBuzz โดยโจทก์นี้จะให้เราแก้ปัญหาดังนี้

อะไรก็ตามที่หาร 3 ลงตัว ให้แสดงผลเป็น Fizz
อะไรก็ตามที่หาร 5 ลงตัว ให้แสดงผลเป็น Buzz
อะไรก็ตามที่หาร 15 ลงตัว ให้แสดงผลเป็ร FizzBuzz
ถ้าไม่เข้า 3 เงื่อนไขข้างบนให้แสดงตัวเลข ที่ใส่เข้าไปออกมา
*/

// Fizzbuzz : fizzbuzz function
func FizzBuzz(number int) string {

	if number%3 == 0 {
		return "Fizz"
	}
	if number%5 == 0 {
		return "Buzz"
	}
	if number%15 == 0 {
		return "FizzBuzz"
	}

	// defualt
	return strconv.Itoa(number)

}
