package main

/*1.transactionlar soni bo'yicha top branches
2.transactionlar summasi bo'yicha top branches
3.transactionda bo'lgan top productlar
4.transactionda bo'lgan top categorylar
5.har bir branchda har bir categorydan qancha transaction bo'lgani*/
import (
	"encoding/json"
	"fmt"
	"os"
	branch "project/branches"
	"project/product"
	"project/transaction"
	"sort"
	"time"
)

func main() {
	//read Product
	products := []product.Product{}
	fileProduct, err := os.ReadFile("data/products.json")
	if err != nil {
		fmt.Println("error is while reading products json file", err.Error())
	}
	err = json.Unmarshal(fileProduct, &products)
	//fmt.Println(products)
	//read Branches
	branches := []branch.Branch{}
	fileBranch, err := os.ReadFile("data/branches.json")
	if err != nil {
		fmt.Println("error is while reading branches json file", err.Error())
	}
	err = json.Unmarshal(fileBranch, &branches)
	//read Transaction
	transactions := []transaction.Transaction{}
	fileTransaction, err := os.ReadFile("data/branch_transaction.json")
	if err != nil {
		fmt.Println("error is while reading products json file", err.Error())
	}
	err = json.Unmarshal(fileTransaction, &transactions)
	fmt.Println(transactions)
	//read categories  - not read categories.json file
	categories := []product.Category{}
	fileCategory, err := os.ReadFile("data/categories.json")
	if err != nil {
		fmt.Println("error is while reading products json file", err.Error())
	}
	err = json.Unmarshal(fileCategory, &categories)
	fmt.Println(categories)

	//task-1 transactionlar boyica top branches
	branchMap := make(map[int]string)
	tMap := make(map[int]int)
	for _, branch := range branches { //id and name
		branchMap[branch.ID] = branch.Name
	}
	for _, t := range transactions { //branchID and count
		tMap[t.BranchID]++
	}
	type Count struct {
		Key   int
		Value int
	}
	var count = []Count{}
	for k, v := range tMap { //branchid , countini
		count = append(count, Count{k, v})
	}
	sort.Slice(count, func(i, j int) bool {
		return count[i].Value > count[j].Value
	})
	for _, c := range count {
		fmt.Println("Branch name: ", branchMap[c.Key], "->", c.Value)
	}
	//2. transactionlar summasi boyica top branches
	productMap := make(map[int]int)
	sumMap := make(map[int]int)
	for _, p := range products {
		productMap[p.ID] = p.Price
	}
	for _, t := range transactions {
		sumMap[t.BranchID] += productMap[t.ProductID] * t.Quantity
	}
	countP := []Count{}
	for k, v := range sumMap {
		countP = append(countP, Count{k, v})
	}
	sort.Slice(countP, func(i, j int) bool {
		return countP[i].Value > countP[j].Value
	})
	for _, c := range countP {
		fmt.Println("Branch name: ", branchMap[c.Key], "->", c.Value)
	}
	//3.transactionda bo'lgan top productlar
	topProduct := make(map[int]int)
	prodMap := make(map[int]string)
	for _, p := range products {
		prodMap[p.ID] = p.Name
	}
	for _, t := range transactions { //id va neci marta qatnashganligini oladi
		topProduct[t.ProductID]++
	}
	countTopProduct := []Count{}
	for k, c := range topProduct { //id - key, value - neci marta qatnashgani
		countTopProduct = append(countTopProduct, Count{k, c})
	}
	sort.Slice(countTopProduct, func(i, j int) bool {
		return countTopProduct[i].Value > countTopProduct[j].Value
	})
	for _, c := range countTopProduct {
		fmt.Println("Proudct name: ", prodMap[c.Key], c.Value)
	}
	//4.transactionda bo'lgan top categorylar
	categoryMap := make(map[int]string)
	ProductMap := make(map[int]int)
	for _, category := range categories {
		categoryMap[category.ID] = category.Name
	}
	for _, transaction := range transactions {
		for _, product := range products {
			if product.ID == transaction.ProductID {
				ProductMap[product.CategoryID]++
			}
		}
	}
	countCategory := []Count{}
	for k, v := range ProductMap {
		countCategory = append(countCategory, Count{k, v})
	}
	sort.Slice(countCategory, func(i, j int) bool {
		return countCategory[i].Value > countCategory[j].Value
	})
	for _, c := range countCategory {
		fmt.Println(categoryMap[c.Key], c.Value)
	}

	//5.har bir branchda har bir categorydan qancha transaction bolgani
	bMap := make(map[int]string)
	catMap := make(map[int]string)
	Cmap := make(map[int]map[int]int)
	for _, b := range branches {
		bMap[b.ID] = b.Name
		Cmap[b.ID] = make(map[int]int)
	}
	for _, c := range categories {
		catMap[c.ID] = c.Name
	}
	for _, p := range products {
		for _, t := range transactions {
			for _, b := range branches {
				if p.ID == t.ProductID && b.ID == t.BranchID {
					Cmap[b.ID][p.CategoryID]++
				}
				//fmt.Println(Cmap)
			}
		}
	}
	for k, v := range Cmap {
		fmt.Println(bMap[k])
		count5 := []Count{}
		for k, c := range v {
			count5 = append(count5, Count{k, c})
		}
		for _, c := range count5 {
			fmt.Printf("%s -> %d\n", catMap[c.Key], c.Value)
		}
	}
	fmt.Println("-----------------------------------------------------------------------")
	//6.har bir branch nechta plus/minus transactionlar soni, plus/minus transactionlar summasini
	/*
		chilonzor
		trans -> plus 2 minus 1
		total price -> plus 260000 minus 80000

		c1
		trans -> plus 1 minus 1
		total price -> plus 90000 minus 32000

		maksimgorkiy
		trans -> plus 1 minus 0
		total price -> plus 20000 minus 0

		uacademy
		trans -> plus 1 minus 1
		total price -> plus 80000 minus 120000
	*/

	plus := make(map[int]int)
	minus := make(map[int]int)
	brachmap := make(map[int]string)
	productmap := make(map[int]int)
	totalSumPlus := make(map[int]int)
	totalSumMinus := make(map[int]int)
	for _, b := range branches {
		brachmap[b.ID] = b.Name
	}
	for _, p := range products {
		productmap[p.ID] = p.Price
	}
	for _, t := range transactions {
		if t.Type == "plus" {
			plus[t.BranchID]++
			totalSumPlus[t.BranchID] += productmap[t.ProductID] * t.Quantity

		} else {
			minus[t.BranchID]++
			totalSumMinus[t.BranchID] += productmap[t.ProductID] * t.Quantity
		}
		//fmt.Printf("%s\n\nTransaction -> plus: %d  minus: %d\nTotalsum -> plus: %d  minus %d\n\n",bMap[t.BranchID],plus[t.BranchID],minus[t.ID],totalSumPlus[t.BranchID],totalSumMinus[t.BranchID])
	}
	for i := range brachmap {

		fmt.Printf("%s\n\nTransaction -> plus: %d  minus: %d\nTotalsum -> plus: %d  minus %d\n\n", brachmap[i], plus[i], minus[i], totalSumPlus[i], totalSumMinus[i])

	}
	/*
		2023-08-02 -> 10
		2023-08-16 -> 9
		2023-08-21 -> 6
		2023-08-09 -> 4
		2023-08-19 -> 4
	*/
	//7.har bir kunda kirgan productlar sonini kamayish tartibida chiqarish
	quantityMap := make(map[string]int)
	for _, t := range transactions {
		if t.Type == "plus" {
			times,err := time.Parse("2006-01-02 15:04:05",t.CreatedAt)
			if err != nil{
				fmt.Println(err)
				return
			}
			str := times.Format("2006-01-02")
			quantityMap[str] += t.Quantity
		}
	}
	for day,count := range quantityMap{
		fmt.Printf("%s: %d\n",day,count)
	}
	
	//8.product qanca ciqarilgan va kiritilgan jadval
	/*
	Coca cola -> plus 15 minus 0
	lavash -> plus 14 minus 10
	olma -> plus 4 minus 0
	*/
	mapProduct := make(map[int]string)
	for _,p := range products{
		mapProduct[p.ID] = p.Name
	}
	countMapPlus := make(map[int]int)
	countMapMinus := make(map[int]int)
	for _,t := range transactions{
				if t.Type == "plus"{
					countMapPlus[t.ProductID] += t.Quantity
				}else{
					countMapMinus[t.ProductID] += t.Quantity
				}
			}
		
	for p := range mapProduct{
		if countMapPlus[p] != 0 || countMapMinus[p] != 0{
			fmt.Println(mapProduct[p]," -> plus: ",countMapPlus[p],"  minus: ",countMapMinus[p])
		}
		
	}
	//9.filialda qanca summalik product qolganligi jadvali
	/*
	chilonzor 180000
	c1 58000
	maksimgorkiy 20000
	uacademy 40000
	*/
	mapBranch := make(map[int]string)
	mapProd := make(map[int]int)
	counterMap := make(map[int]int)
	//sum := make(map[int]int)
	for _,b := range branches{
		mapBranch[b.ID] = b.Name
	}
	for _,p := range products{
		mapProd[p.ID] = p.Price
	}
	for _,t := range transactions{
		if t.Type == "plus"{
			counterMap[t.BranchID] += t.Quantity * mapProd[t.ProductID]
		}else{
			counterMap[t.BranchID] -=t.Quantity * mapProd[t.ProductID]
		}
		//sum[t.BranchID] = mapProd[t.ID] * counterMap[t.BranchID]

	}
	for c := range counterMap{
		fmt.Println(mapBranch[c]," -> ",counterMap[c])
	}

}
