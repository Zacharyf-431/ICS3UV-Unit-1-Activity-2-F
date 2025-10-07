/*
 * @author: Zachary Fowler
 * @version:1.0.0
 * @date : 2025-10-6
 * @fileoverview: This Program uses a date of birth and combines it with the date of graduation with multiple forms of math equations
 */

package main

import "fmt"

func main() {

//Titles
	fmt.Println("\nYear Born     Year Graduated      Result")
//Addition
	fmt.Println("\n  2009     +     2027          =   " + fmt.Sprint(2009 + 2027))
//Subtraction
	fmt.Println("  2009     -     2027          =   " + fmt.Sprint(2009 - 2027))
//Multiplicaiton
	fmt.Println("  2009     *     2027          =   " + fmt.Sprint(2009 * 2027))
//Division
	fmt.Println("  2009     /     2027          =   " + fmt.Sprint(2009 / 2027))
//Modulo
	fmt.Println("  2009     %     2027          =   " + fmt.Sprint(2009 % 2027))


	fmt.Println("\nDone");

}