package main

import ( 
"fmt" 
"time"
)

func main() {
i:=2
switch i{
case 1:
    fmt.Println("one");
case 2:
   fmt.Println("two");
case 3:
     fmt.Println("three");
default:
   fmt.Println("no case match here");
}

switch time.Now().Weekday(){
    case time.Sunday,time.Saturday:
     fmt.Println("today is weekend");
     default:
     fmt.Println("today is working day");
}

t := time.Now().UTC()

localZone,err := time.LoadLocation("Asia/Kolkata")

if err !=nil{
    fmt.Println("some issue into load location",err)
}

switch {
    case t.Hour() < 12:
    	localTime := t.In(localZone)
    fmt.Println("Now is morning",localTime);
    default:
    fmt.Println("afternoon")
}
}
