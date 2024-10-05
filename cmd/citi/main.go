/*
We are developing a stock trading data management software that tracks the prices of different stocks over time and provides useful statistics.

The program includes three types: `Stock`, `PriceRecord`, and `StockCollection`.

Classes:
* The `Stock` type represents data about a specific stock.
* The `PriceRecord` type holds information about a single price record for a stock.
* The `StockCollection` type manages a collection of price records for a particular stock and provides methods to retrieve useful statistics about the stock's prices.
*/
/*
To begin with, we present you with two tasks:
1-1) Read through and understand the code below. Please take as much time as necessary, and feel free to run the code.
1-2) The test for StockCollection is not passing due to a bug in the code. Make the necessary changes to StockCollection to fix the bug.
*/
/*
2) We want to add a new function called "getBiggestChange" to the StockCollection type. This function calculates and returns the largest absolute change in stock price between any two consecutive days in the price records of a stock along with the dates of the change in a list. For example, let's consider the following price records of a stock:

Price Records:
Price:  110         112         90          105
Date:   2023-06-29  2023-07-01  2023-06-25  2023-07-06

Stock price changes (sorted based on date):
Date:     2023-06-25  ->  2023-06-29  ->  2023-07-01 ->  2023-07-06
Price:        90      ->      110     ->     112     ->     105
Change:              +20              +2             -7

In this case, the biggest absolute change in the stock price was +20, which occurred between 2023-06-25 and 2023-06-29. In this case, the function should return [20, "2023-06-25", "2023-06-29"]

Two days are considered consecutive if there are no other days' data in between them in the price records based on their dates.

To assist you in testing this new function, we have provided the testGetBiggestChange function.
*/
package main

import (
	"fmt"
	"math"
	"sort"
)

type Stock struct {
	Symbol string
	Name   string
}

func NewStock(symbol, name string) *Stock {
	return &Stock{
		Symbol: symbol,
		Name:   name,
	}
}

func (s *Stock) String() string {
	return s.Name
}

type PriceRecord struct {
	Stock *Stock
	Price int
	Date  string
}

func NewPriceRecord(stock *Stock, price int, date string) *PriceRecord {
	return &PriceRecord{
		Stock: stock,
		Price: price,
		Date:  date,
	}
}

type ByDate []*PriceRecord

func (d ByDate) Len() int           { return len(d) }
func (d ByDate) Less(i, j int) bool { return d[i].Date < d[j].Date }
func (d ByDate) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }

type StockCollection struct {
	PriceRecords []*PriceRecord
	Stock        *Stock
}

func NewStockCollection(stock *Stock) *StockCollection {
	return &StockCollection{
		PriceRecords: make([]*PriceRecord, 0),
		Stock:        stock,
	}
}

func (sc *StockCollection) GetNumPriceRecords() int {
	return len(sc.PriceRecords)
}

func (sc *StockCollection) AddPriceRecord(priceRecord *PriceRecord) error {
	if priceRecord.Stock != sc.Stock {
		return fmt.Errorf("PriceRecord's Stock is not the same as the StockCollection's")
	}
	sc.PriceRecords = append(sc.PriceRecords, priceRecord)
	return nil
}

func (sc *StockCollection) GetMaxPrice() int {
	maxPrice := -1
	for _, priceRecord := range sc.PriceRecords {
		if priceRecord.Price > maxPrice {
			maxPrice = priceRecord.Price
		}
	}
	return maxPrice
}

func (sc *StockCollection) GetMinPrice() int {
	minPrice := math.MaxInt64
	for _, priceRecord := range sc.PriceRecords {
		if priceRecord.Price < minPrice {
			minPrice = priceRecord.Price
		}
	}
	if minPrice == math.MaxInt64 {
		return -1
	}
	return minPrice
}

func (sc *StockCollection) GetAvgPrice() float64 {
	total := 0
	for _, priceRecord := range sc.PriceRecords {
		total += priceRecord.Price
	}
	if total == 0 {
		return -1.0
	}
	return float64(total) / float64(len(sc.PriceRecords))
}

func testPriceRecord() {
	priceRecord := NewPriceRecord(NewStock("AAPL", "Apple Inc."), 110, "2023-06-29")
	fmt.Println(priceRecord.Price == 110)
	fmt.Println(priceRecord.Date == "2023-06-29")
	fmt.Println(priceRecord.Stock.Symbol == "AAPL")
	fmt.Println(priceRecord.Stock.Name == "Apple Inc.")
}

func testStockCollection() {
	testStock := NewStock("AAPL", "Apple Inc.")
	stockCollection := NewStockCollection(testStock)
	fmt.Println(stockCollection.GetNumPriceRecords() == 0)
	fmt.Println(stockCollection.GetMaxPrice() == -1)
	fmt.Println(stockCollection.GetMinPrice() == -1)
	fmt.Println(stockCollection.GetAvgPrice() == -1.0)

	priceData := []struct {
		price int
		date  string
	}{
		{110, "2023-06-29"},
		{112, "2023-07-01"},
		{90, "2023-06-28"},
		{105, "2023-07-06"},
	}
	for _, data := range priceData {
		priceRecord := NewPriceRecord(testStock, data.price, data.date)
		err := stockCollection.AddPriceRecord(priceRecord)
		if err != nil {
			fmt.Println(false)
		}
	}

	fmt.Println(stockCollection.GetNumPriceRecords() == len(priceData))
	fmt.Println(stockCollection.GetMaxPrice() == 112)
	fmt.Println(stockCollection.GetMinPrice() == 90)
	fmt.Println(math.Abs(stockCollection.GetAvgPrice()-104.25) <= 0.1)
}

func testGetBiggestChange() {
	testStock := NewStock("AAPL", "Apple Inc.")
	stockCollection := NewStockCollection(testStock)
	fmt.Println(stockCollection.GetNumPriceRecords() == 0)
	biggestChange, beforeDate, afterDate := stockCollection.GetBiggestChange()
	fmt.Println(biggestChange == 0 && beforeDate == "" && afterDate == "")

	priceData := []struct {
		price int
		date  string
	}{
		{110, "2023-06-29"},
		{112, "2023-07-01"},
		{90, "2023-06-25"},
		{105, "2023-07-06"},
	}
	for _, data := range priceData {
		priceRecord := NewPriceRecord(testStock, data.price, data.date)
		stockCollection.AddPriceRecord(priceRecord)
	}

	biggestChange, beforeDate, afterDate = stockCollection.GetBiggestChange()
	fmt.Println(biggestChange == 20 && beforeDate == "2023-06-25" && afterDate == "2023-06-29")

	priceData2 := []struct {
		price int
		date  string
	}{
		{200, "2000-01-04"},
		{210, "1999-12-30"},
		{190, "2000-01-03"},
		{180, "2000-01-01"},
	}
	stockCollection2 := NewStockCollection(testStock)
	for _, data := range priceData2 {
		priceRecord := NewPriceRecord(testStock, data.price, data.date)
		stockCollection2.AddPriceRecord(priceRecord)
	}

	biggestChange2, beforeDate2, afterDate2 := stockCollection2.GetBiggestChange()
	fmt.Println(biggestChange2 == -30 && beforeDate2 == "1999-12-30" && afterDate2 == "2000-01-01")
}

func main() {
	testPriceRecord()
	testStockCollection()
	testGetBiggestChange()
}

func (sc *StockCollection) GetBiggestChange() (int, string, string) {

	type Stock1 struct {
		Date  string
		Price int
	}

	st := []Stock1{}

	for _, priceRecord := range sc.PriceRecords {
		st = append(st, Stock1{priceRecord.Date, priceRecord.Price})
	}

	if len(st) == 0 {
		return 0, "", ""
	}

	// s := []string{"2000-02-04","1999-01-04"}
	// sort.Strings(stock.Date)
	sort.Slice(st, func(i, j int) bool {
		return st[i].Date < st[j].Date // sort the string rune
	})

	fmt.Printf("st is %v", st)

	maxAbsoluteIdx := 0
	maxAbsolute := 0
	for idx, s := range st {
		if idx < len(st)-1 {
			if math.Abs(float64(s.Price)-float64(st[idx+1].Price)) > math.Abs(float64(maxAbsolute)) {
				maxAbsoluteIdx = idx
				maxAbsolute = st[idx+1].Price - s.Price
			}
		}
	}

	fmt.Println(maxAbsolute)

	return maxAbsolute, st[maxAbsoluteIdx].Date, st[maxAbsoluteIdx+1].Date
}

// func (sc *StockCollection) GetBiggestChange() (int, string, string){
//
// 	type struct price {
// 		date string
// 		price int
// 	}
//
// 		price := make([]price, 10)
//
// 		for _, priceRecord := range sc.PriceRecords {
// 		price = append(price, price(priceRecord.Date, priceRecord.Price);
// 	}
//
// 		// s := []string{"2000-02-04","1999-01-04"}
// 		sort.Strings(price.Date)
//
// 		maxAbsoluteIdx := 0;
// 		for idx, p := range price {
// 		if p[idx].price - p[idx+1].price > maxAbsolute {
// 		maxAbsoluteIdx =
// 	}
// 	}
//
// 		fmt.Println(date)
//
// 		return 0 , "", "" ;
// 	}
