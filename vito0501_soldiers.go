package main

import "fmt"
import "os"
import "io"

type PriceType struct{
	close float64
	open float64
}
var price[29734] PriceType
var date, time, x, i, countbuy int
var openprc, highprc, lowprc float64
var err error
var file *os.File
var fname string


func main (){
    fname = "input1.dat"
	file,err := os.Open(fname)
	countbuy = 0
	_,err =fmt.Fscanf(file,"%d %d;%v;%v;%v;%v;0\n",&date,&time,&price[0].open,&highprc,&lowprc,&price[0].close)
	i = 1
	for err != io.EOF && i < 29734	 {
		_,err =fmt.Fscanf(file,"%d %d;%v;%v;%v;%v;0\n",&date,&time,&price[i].open,&highprc,&lowprc,&price[i].close)
		Formula()
		PrintData()
		i = i + 1
	}
}

func Formula (){
	/*
	I.S. : From the input on the file we get date, time, openprice, highprice, lowprice, closeprice. We
	       compare the first closseprice and second closeprice,
	F.S. : We get when we must buy in variable countbuy or in chart it's three bar increasing
	*/
	if price[i].close > price[i-1].close && price[i].open > price[i-1].open && price[i].close > price[i].open && price[i-1].close > price[i-1].open {
		countbuy = 1 + countbuy
	}else {
		countbuy = 0
	}
}

func PrintData(){
    /*
	I.S. : The data is printed if countbuy equals to 3 or in chart it's three bar increasing
	F.S. : Time to Buy is printed
	*/
	if countbuy == 3{
		fmt.Println(date,time,price[i].close, "***buy")
	}
}