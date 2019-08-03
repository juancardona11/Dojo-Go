package main

import(
	"fmt"
	"errors"
	"os"
	"net"
	"net/rpc"
	"math"
)

type Args struct{
    A, B int
}

type Response struct{
	Quo, res int
}

type Math byte

func (m*Math)Add(args *Args, res *int) error {
	*res = args.A + args.B
	return nil 
}

func (m *Math) Divide(args *Args, res *Response) error {
	if args.B == 0 {
		return errors.New("You are trying divide by zero")
	}
	res.Quo = args.A / args.B
	res.Res = args.A % args.B
	return nil
}

func (m *Math) Major(slice *[]int, res *int) error {
	var major = math.MinInt32
	for _, v := range *slice {
		if v > major {
			major = v
		}
	}
	*res = major
	return nil
}

func (m *Math) Minor(slice *[]int, res *int) error {
	var minor = math.MaxInt32
	for _, v := range *slice {
		if v < minor {
			minor = v
		}
	}
	*res = minor
	return nil
}

func checkError(err error){
	if err != nil{
		fmt.Println("Erro! %v", err.Error())
		os.Exit(1)
	}
}
func main() {

	// DEFER PANIC RECOVER 
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("A Error has happen and I Survived, the error was: ", err)
			}
		}()
		if len(os.Args) != 2 {
			fmt.Println("Put a server port")
			os.Exit(1)
		}
		var service = os.Args[1]
		var reader = bufio.NewReader(os.Stdin)
	
		client, err := rpc.Dial("tcp", service)
		checkError(err, "Dialing")
	
		// Synchronous Call
		if err == nil {
			message := ""
		out:
			for {
				printMenu()
				option := setValue(reader, &message)
				switch option {
				case 1:
					var reply int
					var args = Args{}
					setArgs(reader, &args)
					if err = client.Call("Math.Add", args, &reply); err == nil {
						fmt.Printf("Math %d + %d = %d\n", args.A, args.B, reply)
					}
					checkError(err, "Math.Add Call")
				case 2:
					response := Response{}
					var args = Args{}
					setArgs(reader, &args)
					if err = func main() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("A Error has happen and I Survived, the error was: ", err)
			}
		}()
		if len(os.Args) != 2 {
			fmt.Println("Put a server port")
			os.Exit(1)
		}
		var service = os.Args[1]
		var reader = bufio.NewReader(os.Stdin)
	
		client, err := rpc.Dial("tcp", service)
		checkError(err, "Dialing")
	
		// Synchronous Call
		if err == nil {
			message := ""
		out:
			for {
				printMenu()
				option := setValue(reader, &message)
				switch option {
				case 1:
					var reply int
					var args = Args{}
					setArgs(reader, &args)
					if err = client.Call("Math.Add", args, &reply); err == nil {
						fmt.Printf("Math %d + %d = %d\n", args.A, args.B, reply)
					}
					checkError(err, "Math.Add Call")
				case 2:
					response := Response{}
					var args = Args{}
					setArgs(reader, &args)
					if err = client.Call("Math.Divide", args, &response); err == nil {
						fmt.Printf("Math %d/%d = %d it's residue is %d \n", args.A, args.B, response.Quo, response.Res)
					}
					checkError(err, "Math.Divide Call")
				case 3:
					operation := "Math.Major"
					majorOrMinor(&service, reader, &operation)
				case 4:
					operation := "Math.Minor"
					majorOrMinor(&service, reader, &operation)
				case 5:
					fmt.Printf("Good Bye!\n")
					break out
				default:
					fmt.Printf("Invalid Option\n")
				}
			}
			defer client.Close()
		}
	}
	; err == nil {
						fmt.Printf("Math %d/%d = %d it's residue is %d \n", args.A, args.B, response.Quo, response.Res)
					}
					checkError(err, "Math.Divide Call")
				case 3:
					operation := "Math.Major"
					majorOrMinor(&service, reader, &operation)
				case 4:
					operation := "Math.Minor"
					majorOrMinor(&service, reader, &operation)
				case 5:
					fmt.Printf("Good Bye!\n")
					break out
				default:
					fmt.Printf("Invalid Option\n")
				}
			}
			defer client.Close()
		}
	}