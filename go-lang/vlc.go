/*

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.

$ go build -o id.exe source.go
$ go run source.go

*/
package main
import "os/exec"
import "runtime"

func main() {
  myos := runtime.GOOS;
  myarch := runtime.GOARCH;
  var chrome = "";
  var tpt = "";
  var cmdopen *exec.Cmd;

  if myos == "windows" {
    if myarch == "386" {
      chrome = "C:/Program Files/VideoLan/VLC/vlc.exe";
      tpt =  "C:/Program Files/Google/Chrome/Application/chrome.exe";
    } else {
      chrome = "C:/Program Files (x86)/VideoLan/VLC/vlc.exe";
      tpt = "C:/Program Files (x86)/Google/Chrome/Application/chrome.exe";
    }   

    cmdopen = exec.Command(chrome, "vlc://quit");
    err := cmdopen.Start();
    if err != nil {
      println("Failed: ", err);
    }

    cmdopen = exec.Command(chrome, "rtsp://admin:9999@192.168.100.150:8557/h264", ":sout=#transcode{vcodec=theo,vb=1400,fps=30,scale=Auto,acodec=none}:http{mux=ogg,dst=:8081/test}", ":no-sout-keep", ":no-sout-audio", "--one-instance");
    err1 := cmdopen.Start();
    if err1 != nil {
      println("Failed: ", err1);
    } else {
      println("Success: ", err1);
      cmdopen = exec.Command(tpt, "--chrome-frame", "C:/Program Files (x86)/Attendedbyhumans/vlc.html");
      cmdopen.Start();
    }

  } else {
    chrome = "/root/Eid.jar";       
    cmdopen = exec.Command("java", "-cp", chrome, "eid.Eid");
    err := cmdopen.Start();
    if err != nil {
      println("Failed: ", err);
    } 

  } 


}