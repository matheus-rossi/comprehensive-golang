package main

import "fmt"

func main() { /* Comments in GOLANG */
	/* Numbers */
	/* Integers */
	/* Unsigned Integer (Just Positive Numbers) */
	/* uint8 <byte> */
	/* uint16 */
	/* uint32 */
	/* uint64 */
	/* Signed Integer (Negative and Positive Numbers) */
	/* int8 */
	/* int16 */
	/* int32 <rune> */
	/* int64 */
	fmt.Println("1+1=", 1+1)

	/* Float */
	/* float32 */
	/* float64 */
	fmt.Println("2.5+1=", 2.5+1)

	/* String */
	/* "" or `` */
	fmt.Println("lenght:", len("Hello, World!"))

	hello := "Hello "
	world := "World!"
	phrase := hello + world
	fmt.Println(phrase)

	fmt.Println("first letter:", hello[1], "in binary")

	/* Boolean */
	/* Logical Operators */
	/* &&  and */
	/* ||  or  */
	/* !   not */

	fmt.Println("true && false =", true && false)
}
