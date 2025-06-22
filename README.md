
// go run main.go
// go mod init my-app

//create launch.json for go app 

// các tp trong golang cần package 
// các hàm muốn export thì phải ghi hoa chữ đầu

// func add(x int , y int) int {
    return x+y;
}

Rut gon
// func add(x ,y int) int {
return x+y;
}

//Co the tra ve nhieu loai bien 

func add (x,y string ) (string ,string) {
    return x,y 
}

//bien tra ve co the duoc khoi tao truoc


func add(sum int) (x,y int){
    x = sum + 2 
    y = sum /3 
    return  -> nó tự trả về x y
}

ở main gọi  -> add(20)

// khai báo -> có var thì khong có := mà có =

x:= 20  -> tự suy kiểu

var x int =1 
var x int    -> mặc định là  0 
c , python := true , false -> bool mặc định là false , chuỗi không gán mặc định là  ""



// For thì giống

for i:=0 ; i< 10 ; i++

if a>0 {

}

hoặc 
if a:=b+c ; if a >10 {
    return "A lớn hơn 10"
}
return "A bé hơn 10"

//switch case -> không cần break;



//Defer   -> STACK LIFO (dung for de xem chi tiet) 

    defer fmt.Println("") // dòng này in ra sau cùng -> dùng cho việc đóng file neu lỗi 


//

var p *int  -> khai bao con tro
i :=42
p &i -> p đang tham chieu tới i -> nêu i thay doi p thay doi theo

*p = 21  -> i = 21 bi thay doi theo

*p -> tuong duong i 


//struct
-> viết hoa chu đầu cho package khác nhìn
type Test struct {
    X int 
    Y int

}

    v:= Test{1,2}
    p:= &v  -> ref tới Test
    p.X = 3
fmt.Println(v)


//make([]int , 5) // len = 5  -> như slice



//Method -> chỉ la func với 1 reciver thôi 

func bình thường -> không dùng được potiner trong bien truyen vao 
func +ten hàm + kieu dl tra ve

func + reciver + ten ham + kieu dl tra ve

-> reciver (có thể là pointer hoặc value)

func (v Test) add() int {

}
func (v *Test) add() int {

}

//Interface

type Abser interface {
    Abs() float64  
}

chỉ cần method nào có 

Abs() float64 -. thì nó tự implement 

testing
 tesst