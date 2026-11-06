package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	result := tong(1,2,3)
	result1 := tong(3,4,5)
	fmt.Println(result)
	fmt.Println(result1)

	output, error := hang("bánh", 10000, 2);
	if error != nil {
		fmt.Println("Lỗi:", error)
	} else {
		fmt.Println(output)
	}

	diemResult := nhanXet(8)
	fmt.Println(diemResult)
	diemResult1 := nhanXet(5)
	fmt.Println(diemResult1)
	diemResult2 := nhanXet(3)
	fmt.Println(diemResult2)
	diemResult3 := nhanXet(-1)
	fmt.Println(diemResult3)

	for a := 2; a <= 9; a++ {
        for b := 1; b <= 9; b++ {
            fmt.Print(a*b, " ")
        }
        fmt.Println()
    }
}

func tong(cacSo ...int) int {
	sum := 0
	for _, so := range cacSo {
		sum += so
	}
	return sum
}

func hang(ten string, gia, soLuong int) (int, error){
	if ten == ""{
		return 0, fmt.Errorf("thiếu tên")
	}
	if gia < 0 || soLuong < 0 {
		return 0, fmt.Errorf("số không hợp lệ")
	}
	return gia * soLuong, nil
}

func nhanXet(diem int) string {
	switch {
	case diem >= 0 && diem <= 4:
		return "chưa đạt"
	case diem >= 5 && diem <= 7:
		return "đạt"
	case diem >=8 && diem <= 10:
		return "tốt"
	default: 
		return "điểm sai"
	}
}