package main

import (
	"encoding/csv"
	"strings"
)

type BasicColumns struct {
	FirstName string
	LastName string 
	Phone string
	Email string
	PriceTotalLine float64
	PriceRefundTotal float64
	PriceTotalOutstanding float64
	PriceTotal float64
	PaymentStatus string
	OrderFulfillmentStatus string
	LineType string
}

type Address struct {
	FirstName string
	LastName string
	Company string
	Phone string
	Line1 string
	Line2 string
	City string
	ProvinceCode string
	Zip string
	Country string
}

type ProductLineItem struct {
	Name string
	Sku string
	Quantity string
	PricePer float64
	PriceTotal float64
	WeightGrams float64
	LineFulfillmentStatus string
}

type TransactionLineItem struct {
	id string
	Kind string
	ProccessedAt string
	Amount float64
	Status string
}

type FulfillmentLineItem struct {
	id string
	Status string
	CreatedAt string
	ShipmentStatus string
}

type RefundLineItem struct {
	RefundAmount float64
}

func main() {
	//load initial csv file into memory
	data, err := readCSVFile("bigcommerce-orders-export.csv")
	if err!= nil {
		panic(err)
    }
	//parse initial csv into usuable format
	reader, err := parseCSV(data)
    if err!= nil {
        panic(err)
    }
	
	//remove header line by reading it
	_, err = reader.Read()
	if err != nil {
		panic(err)
	}

	//read remaing record lines into array
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	//create csv writer to new file
	writer, file, err := createCSVWriter("matrixify-orders-import.csv")
	if err != nil {
        panic(err)
    }
	//defer file close to end of execution
	defer file.Close()

	//create file headers
	writeCSVRecord(writer, []string{})

	//for order products
	for _, record := range records {
		//record basic info used for all records
		basicInfo := BasicColumns {

		}
		billingAddress := Address {

		}
		shippingAddress := Address {

		}
		//split products into unique line items
		products := strings.Split(record[36], "|")
		if len(products) > 1 {
			for _, product := range products {
				//seperate product values
				productLineItem := ProductLineItem{
					Name: product,
				}
				writeProductRecord(writer, basicInfo, productLineItem)
			}
		} else {
			productLineItem := ProductLineItem{
				Name: "",
			}
			writeProductRecord(writer, basicInfo, productLineItem)
		}
		//create transaction line
		transaction := TransactionLineItem {
			Status: "",
		}
		writeTransactionRecord(writer, basicInfo, billingAddress, transaction)
		//create fulfillment line
		fulfillment := FulfillmentLineItem {
			Status: "",
		}
		writeFulfillmentRecord(writer, basicInfo, shippingAddress, fulfillment)

	}

	writer.Flush()
	if err := writer.Error(); err != nil {
        panic(err)
    }
}

func writeProductRecord(writer *csv.Writer, basicCols BasicColumns, product ProductLineItem){
	writeCSVRecord(writer, []string{})
}

func writeTransactionRecord(writer *csv.Writer, basicCols BasicColumns, billingAddress Address, transaction TransactionLineItem){
	writeCSVRecord(writer, []string{})
}

func writeFulfillmentRecord(writer *csv.Writer, basicCols BasicColumns, shippingAddress Address, fulfillment FulfillmentLineItem){
	writeCSVRecord(writer, []string{})
}