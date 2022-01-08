package Menus

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"

	"grocery.com/Grocery/Item"
)

var Inventory map[string]Item.Item

func MainMenu(){
	var choice int
	fmt.Println("*************************************")
	fmt.Println("*                                   *")
	fmt.Println("* Enter the number of the desired   *")
	fmt.Println("* menu choice or enter 0 to exit.   *")
	fmt.Println("*                                   *")
	fmt.Println("* Store Manager: manual operations  *")
	fmt.Println("*   1 -  List inventory             *")
	fmt.Println("*   2 -  Edit inventory             *")
	fmt.Println("*                                   *")
	fmt.Println("* Store Manager: batch operations   *")
	fmt.Println("*   3 - Load inventory file         *")
	fmt.Println("*   4 - Save inventory file         *")
	fmt.Println("*                                   *")
	fmt.Println("*************************************")
	fmt.Println("Choice must be '1', '2', '3', '4', or '0'")
	fmt.Printf("Choice: ")
	fmt.Scanln(&choice)

	for (choice != 1 && choice != 2 && choice != 3 && choice != 4 && choice != 0){
		fmt.Println("Choice must be '1', '2', '3', '4' or '0'")
		fmt.Printf("Choice: ")
		fmt.Scanln(&choice)
	}

	if choice == 1{
		List(Inventory)
		Cont()
	}
	if choice == 2{
		EditMenu(Inventory)
		Cont()
	}
	if choice == 3{
		Load()
		Cont()
	}
	if choice == 4{
		Write(Inventory)
		Cont()
	}
	if choice == 0{
		os.Exit(1)
	}

}

func Write(x map[string]Item.Item){
	var Filename string
	fmt.Println("What would you like to call the file?")
	fmt.Printf("Filename: ")
	fmt.Scanln(&Filename)

	out, err := os.Create(Filename)
	if err != nil{
		panic(err)
	}

	defer out.Close()

	for _, value := range x{
		fmt.Fprintln(out,"{",value.GTIN,",", value.Name,",", value.Description,",", value.Wholesale,",", value.Retail,",", value.Quantity, "}")
	}

}

func EditMenu(x map[string]Item.Item){
	fmt.Println("*-----*-----*-----*-----*-----*----*")
	fmt.Println("*                                  *")
	fmt.Println("* * Inventory Editing Operations * *")
	fmt.Println("*                                  *")
	fmt.Println("* Enter menu choice or 0 to exit.  *")
	fmt.Println("*                                  *")
	fmt.Println("*   1 - Add new item               *")
	fmt.Println("*   2 - Edit item content          *")
	fmt.Println("*   3 - Delete item                *")
	fmt.Println("*                                  *")
	fmt.Println("************************************")
	fmt.Printf("Choice: ")

	var choice int
	fmt.Scanln(&choice)

	for (choice != 1 && choice != 2 && choice != 3 && choice != 0){
		fmt.Println("Choice must be '1', '2', '3', '4' or '0'")
		fmt.Printf("Choice: ")
		fmt.Scanln(&choice)
	}

	if choice == 1{
		//add
		in := bufio.NewReader(os.Stdin)


		fmt.Printf("Enter the GTIN of the new Item: ")
		GTIN, _ := in.ReadString('\n')
		GTIN = GTIN[:len(GTIN)-1]

		fmt.Printf("Enter the name of the new Item: ")
		name, _ := in.ReadString('\n')
		name = name[:len(name)-1]

		fmt.Printf("Enter the description of the new Item: ")
		description, _ := in.ReadString('\n')
		description = description[:len(description)-1]

		fmt.Printf("Enter the retail of the new Item: ")
		retailString, _ := in.ReadString('\n')
		retailString = retailString[:len(retailString)-1]
		retail, _ := strconv.ParseFloat(retailString, 64)

		fmt.Printf("Enter the wholesale of the new Item: ")
		wholesaleString, _ := in.ReadString('\n')
		wholesaleString = wholesaleString[:len(wholesaleString)-1]
		wholesale, _ := strconv.ParseFloat(wholesaleString, 64)

		fmt.Printf("Enter the quantity of the new Item: ")
		quantityString, _ := in.ReadString('\n')
		quantityString = quantityString[:len(quantityString)-1]
		quantity, _ := strconv.ParseInt(quantityString, 10, 64)

		if x == nil{
			x = make(map[string]Item.Item)
		}

		x[GTIN] = Item.Item{GTIN: GTIN, Name: name, Description: description, Retail: retail, Wholesale: wholesale, Quantity:  int64(quantity)}
		fmt.Println("New Inventory:")
		Inventory = x
		List(x)
	}

	if choice == 2{
		//edit
		var GTIN string
		fmt.Printf("Enter the GTIN of item to edit: ")
		fmt.Scanln(&GTIN)

		_, there := Inventory[GTIN] 

		if there{

			var choice int

			fmt.Println("Here is the selected Item: ")
			fmt.Println(Inventory[GTIN])

			fmt.Println("What would you like to edit?")
			fmt.Println("1: GTIN")
			fmt.Println("2: Name")
			fmt.Println("3: Description")
			fmt.Println("4: Retail")
			fmt.Println("5: Wholesale")
			fmt.Println("6: Quantity")

			fmt.Printf("Choice: ")
			fmt.Scanln(&choice)

			for choice != 1 && choice != 2 && choice != 3 && choice != 4 && choice != 5 && choice != 6{
				fmt.Println("Please enter 1-6")
				fmt.Println("What would you like to edit?")
				fmt.Println("1: GTIN")
				fmt.Println("2: Name")
				fmt.Println("3: Description")
				fmt.Println("4: Retail")
				fmt.Println("5: Wholesale")
				fmt.Println("6: Quantity")

				fmt.Printf("Choice: ")
				fmt.Scanln(&choice)
			}

			in := bufio.NewReader(os.Stdin)

			if choice == 1{
				tempGTIN, _ := in.ReadString('\n')
				tempGTIN = tempGTIN[:len(tempGTIN)-1]

				newItem := Inventory[GTIN]
				newItem.GTIN = tempGTIN
				Inventory[GTIN] = newItem
			}
			if choice == 2{
				tempName, _ := in.ReadString('\n')
				tempName = tempName[:len(tempName)-1]

				newItem := Inventory[GTIN]
				newItem.Name = tempName
				Inventory[GTIN] = newItem

			}
			if choice == 3{
				//desc
				tempDesc, _ := in.ReadString('\n')
				tempDesc = tempDesc[:len(tempDesc)-1]

				newItem := Inventory[GTIN]
				newItem.Description = tempDesc
				Inventory[GTIN] = newItem

			}
			if choice == 4{
				//retail
				tempRet, _ := in.ReadString('\n')
				tempRet = tempRet[:len(tempRet)-1]
				tempRetFl, _ := strconv.ParseFloat(tempRet, 64)

				newItem := Inventory[GTIN]
				newItem.Retail = tempRetFl
				Inventory[GTIN] = newItem

			}
			if choice == 5{
				//whole
				tempWhole, _ := in.ReadString('\n')
				tempWhole = tempWhole[:len(tempWhole)-1]
				tempWholeFl, _ := strconv.ParseFloat(tempWhole, 64)

				newItem := Inventory[GTIN]
				newItem.Wholesale = tempWholeFl
				Inventory[GTIN] = newItem

			}
			if choice == 6{
				//quant
				tempQuant, _ := in.ReadString('\n')
				tempQuant = tempQuant[:len(tempQuant)-1]
				tempQuantInt, _ := strconv.ParseInt(tempQuant, 10, 64)

				newItem := Inventory[GTIN]
				newItem.Quantity = tempQuantInt
				Inventory[GTIN] = newItem

			}

		}else{
			fmt.Println("Item not found.")
		}
	}

	if choice == 3{
		//delete
		var GTIN string
		fmt.Printf("Enter the GTIN of item to edit: ")
		fmt.Scanln(&GTIN)

		_, there := Inventory[GTIN] 


		if there{
			delete(Inventory, GTIN)
			fmt.Println("Item deleted.")
		}else{
			fmt.Println("Item not found.")
		}
		fmt.Println("New inventory: ")
		List(Inventory)

	}

	if choice == 0{
		//exit
		MainMenu()
	}

}

func List(x map[string]Item.Item){
	key := make([]string, 0, len(x))
	for _ , test := range x{
		key = append(key, test.GTIN)
	}

	sort.Strings(key)

	fmt.Println("{GTIN----NAME----DESCRIPTION----RETAIL----WHOLESALE----QUANTITY}")

	for _, value := range key{
		fmt.Println(x[value])
	}

}

func Load(){ 
	inventory := make(map[string]Item.Item)
	var file string
	fmt.Printf("Enter <inventory.csv> file to load: ")
	fmt.Scanln(&file)

	OpenedFile, err := os.Open(file)
	
	if err != nil{
		log.Fatal(err)
	}
	defer OpenedFile.Close()

	scanner := csv.NewReader(OpenedFile)

	for err != io.EOF{
		lines, err := scanner.Read()
		if err == io.EOF{
			break
		}
		if err != nil{
			log.Fatal(err)
		}
		GTIN := lines[0]
		Name := lines[1]
		Wholesale, _:= strconv.ParseFloat(lines[2],64)
		Retail, _:= strconv.ParseFloat(lines[3],64)
		Description := lines[4]
		Quantity, _ := strconv.ParseInt(lines[5], 10, 64)
		inventory[GTIN] = Item.Item{GTIN: GTIN, Name: Name, Wholesale: Wholesale, Retail: Retail, Description: Description, Quantity: Quantity}
	}
	if err != nil{
		log.Fatal(err)
	}
	Inventory = inventory
}

func Cont(){
	var choice string
	fmt.Println()
	fmt.Println("*************************************")
	fmt.Println("Would you like to Continue")
	fmt.Printf("y/n: ")
	fmt.Scanln(&choice)
	fmt.Println("*************************************")

	for choice != "y" && choice != "n"{
		fmt.Println("y/n")
		fmt.Scanln(&choice)
	}

	if choice == "y"{
		MainMenu()
	}else{
		fmt.Println("Exiting grocery inventory.")
		os.Exit(1)
	}
}