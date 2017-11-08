package main



import (
"fmt"
"bufio"
"strconv"
"io/ioutil"
"os"
"strings"
//"io"
)



func printTriangularStars(numberOfRows int){
	for i:=1; i<=numberOfRows; i++ {
		for j:=0; j<i; j++{
			fmt.Print("*")
		}
		fmt.Println("");
	}
}

func printRectangleStars(numberOfRows int){
	for i:=1 ; i<= numberOfRows; i++{
		for j := 0; j<numberOfRows*2;j++{
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

func printInverseTriangle(numberOfRows int) {

	for i:=0;i<=numberOfRows; i++{
		for j:=1;j<=(numberOfRows-i);j++{
			fmt.Print("*")
		}
		fmt.Println("")
	}
}


func printSquareStars(numberOfRows int) {
	for i := 1; i <= numberOfRows; i++ {
		for j := 0; j < numberOfRows; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}





func readUserEnteredValueAsInt(preEnterMessage string) int{
	fmt.Println(preEnterMessage);

	reader := bufio.NewReader(os.Stdin)

	numberOfRowsString, err := reader.ReadString('\n')
	if(err != nil){
		fmt.Println(err)
		return -1;
	}
	fmt.Println("Thank you for providing the preEnterMessage - "+numberOfRowsString)
	numberOfRowsString = strings.TrimRight(numberOfRowsString, "\n");
	numberOfRowsInt,stringToIntError := strconv.Atoi(numberOfRowsString)
	if(stringToIntError != nil){
		fmt.Println(stringToIntError)
		return -1;
	}
	return numberOfRowsInt;
}

func readFile1(fileName string){
	data,error := ioutil.ReadFile(fileName);
	if(error != nil){
		fmt.Println("ERROR")
		fmt.Println(error);
	}
	fmt.Println(string(data))
}

func myWriteFile(fileName string, data string){
	dataInBytes := []byte(data)
	fileHandle,error := os.OpenFile(fileName,os.O_APPEND|os.O_CREATE|os.O_WRONLY,0600);
	if(error != nil){
		fmt.Println(error)
		return;
	}
	defer fileHandle.Close()
	fileHandle.WriteAt(dataInBytes,19);
	//ioutil.WriteFile(fileName,dataInBytes,0644)

}

func main() {
	//readFile("/Users/Naseer/Downloads/Afzaal/dev/src/hello/yusuf.pngf ")
	//readFile("/Users/Naseer/Downloads/Afzaal/dev/src/hello/data5.txt")
	myWriteFile("/Users/Naseer/Downloads/Afzaal/dev/src/hello/yusuf.png","TTTTTTTT");
	//readFile("/Users/Naseer/Downloads/Afzaal/dev/src/hello/data5.txt")
}

func enterPatternSize() int{
	rows := readUserEnteredValueAsInt("Please type the number of rows you need in the star pattern")
	for rows > 30 {
		fmt.Println("Allowed value should always be less than 30")
		return readUserEnteredValueAsInt("Please type the number of rows you need in the star pattern: ")
	}
	return  rows;
}


func selectedPatternWithIfStatements (){
	fmt.Println("enter the desired triangle pattern")
	fmt.Println("enter 1 for inverse triangle")
	fmt.Println("enter 2 for square triangle")
	fmt.Println("enter 3 for triangular star triangle")

	option := readUserEnteredValueAsInt("pick your option")
	rows := 1

	if option == 1 {
		rows = enterPatternSize()
		printInverseTriangle(rows)
	}
	if option == 2 {
		rows = enterPatternSize()
		printSquareStars(rows)
	}
	if option == 3 {
		rows = enterPatternSize()
		printTriangularStars(rows)
	}
}



func pickYourPattern() int{

	fmt.Println("enter the desired triangle pattern")
	fmt.Println("enter 1 for inverse triangle")
	fmt.Println("enter 2 for square triangle")
	fmt.Println("enter 3 for triangular star triangle")
	fmt.Println("enter -1 to Quit")

	selectedOption := readUserEnteredValueAsInt("pick your option")
	if(selectedOption == -1){
		os.Exit(1);
	}
	return selectedOption
}

func selectedPatternWithSwitchStatement (){
	for {
		option := pickYourPattern();
		rows := 1

		switch option {
		case 1:
			{
				rows = enterPatternSize()
				printInverseTriangle(rows)
			}
		case 2:
			{
				rows = enterPatternSize()
				printSquareStars(rows)
			}
		case 3:
			{
				rows = enterPatternSize()
				printTriangularStars(rows)
			}
		default:
			{
				fmt.Println("Please read carefully and select the correct options..")
				selectedPatternWithSwitchStatement()
			}
		}
	}

}


func printCustomPattern(){
	for {

		rows := readUserEnteredValueAsInt("Please type the number of rows you need in the star pattern, enter -1 to exits: ")
		if rows == -1{
			fmt.Println("Thankyou for using my functionality")
			return
		}
		for rows > 30 {
			fmt.Println("Allowed value should always be less than 30")
			rows = readUserEnteredValueAsInt("Please type the number of rows you need in the star pattern: ")
		}
		printInverseTriangle(rows);
		printSquareStars(rows);
		printTriangularStars(rows);
	}

}



