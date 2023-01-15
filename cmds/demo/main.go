package main


import (
	"os"
	"fmt"
	"flag"
	"strings"
        //"encoding/json"
	//"math"
       "github.com/seldonsmule/simpleconffile"
       "github.com/seldonsmule/smartthings"
        "time"
        //"github.com/seldonsmule/restapi"
        "github.com/seldonsmule/logmsg"

)


type Configuration struct {

  ST_Token string             // Token from SmartThings to access your devices
  ConfFilename string         // Name of conffile
  Encrypted bool

  // non users input

  ST_Scenes smartthings.StScenes // holds the data from a screns querry to ST
  ST_Devices smartthings.StDevices // holds the data from a screns querry to ST

}


const COMPILE_IN_KEY = "example key 9999"

var gMyConf Configuration


func help(){

  fmt.Println("Demo of the smartthings package")

  fmt.Println("Usage smartthings -cmd [a command, see below]")
  fmt.Println()
  flag.PrintDefaults()
  fmt.Println()
  fmt.Println("cmds:")
  fmt.Println("       setconf - Setup Conf file")
  fmt.Println("             -token SmartThings API token")
  fmt.Println("             -conffile name of conffile (.smart.conf default)")
  fmt.Println("       readconf - Display conf info")
  fmt.Println("       get_st_structs - Saves off any neede SmartThings golang structs (chick/egg thing)")
  fmt.Println("       runscene - Runs a SmartThings Scene")
  fmt.Println("             -name Name of a predefined SmartThings scene")
  fmt.Println("       listscenes - List all the SmartThings Scense that are avaialble")
  fmt.Println("       switchon - Turns on a switch")
  fmt.Println("             -name Name of device (switch)")
  fmt.Println("       switchon - Turns off a switch")
  fmt.Println("             -name Name of device (switch)")
  fmt.Println("       switchstatus - status of a switch state")
  fmt.Println("             -name Name of device (switch)")
  fmt.Println()


}

func readconf(confFile string, printstd bool) bool{

  simple := simpleconffile.New(COMPILE_IN_KEY, confFile)

  if(!simple.ReadConf(&gMyConf)){
    msg := fmt.Sprintln("Error reading conf file: ", confFile)
    logmsg.Print(logmsg.Warning, msg)
    return false
  }

  if(gMyConf.Encrypted){
    gMyConf.ST_Token = simple.DecryptString(gMyConf.ST_Token)
  }

     
  if(printstd){

    fmt.Printf("Encrypted [%v]\n", gMyConf.Encrypted)
    fmt.Printf("ST_Token [%v]\n", gMyConf.ST_Token)
    fmt.Printf("ConfFilename [%v]\n", gMyConf.ConfFilename)

  }

  return true

}

func main() {


  cmdPtr := flag.String("cmd", "help", "Command to run")
  tokenPtr := flag.String("token", "notset", "SmartThings access Token")
  namePtr := flag.String("name", "notset", "SmartThings Device/Scene to call - used with cmd [switchon | switchoff]")
  confPtr := flag.String("conffile", ".smart.conf", "config file name")
  bdebugPtr := flag.Bool("debug", false, "If true, do debug magic")

  flag.Parse()

fmt.Printf("cmd=%s\n", *cmdPtr)

  logmsg.SetLogFile("smarthings.log");

  logmsg.Print(logmsg.Info, "cmdPtr = ", *cmdPtr)
  logmsg.Print(logmsg.Info, "confPtr = ", *confPtr)
  logmsg.Print(logmsg.Info, "tokenPtr = ", *tokenPtr)
  logmsg.Print(logmsg.Info, "namePtr = ", *namePtr)
  logmsg.Print(logmsg.Info, "bdebugPtr = ", *bdebugPtr)
  logmsg.Print(logmsg.Info, "tail = ", flag.Args())

  if(*cmdPtr == "help"){
    help()
    os.Exit(1)
  }

  readconf(*confPtr, false)

  st := smartthings.New()

  st.SetToken(gMyConf.ST_Token)

  st.Dump()

  switch *cmdPtr {

    case "readconf":
      fmt.Println("Reading conf file")
      readconf(*confPtr, true)

    case "setconf":

      readconf(*confPtr, false) // ignore errors

      fmt.Println("Setting conf file")

      simple := simpleconffile.New(COMPILE_IN_KEY, *confPtr)

      gMyConf.Encrypted = true

      if(strings.Compare(*tokenPtr, "notset") != 0){
        gMyConf.ST_Token = simple.EncryptString(*tokenPtr)
      }else{
        gMyConf.ST_Token = simple.EncryptString(gMyConf.ST_Token)
      }

      gMyConf.ConfFilename = *confPtr

      simple.SaveConf(gMyConf)

      readconf(*confPtr, true) // ignore errors


    case "get_st_structs":

      st.GetStructs()

    case "listscenes":
      st.PrintSceneList()

    case "listdevices":
      st.PrintDeviceList()

    case "switchon":
      st.DeviceSwitchOnOff(*namePtr, true)  

    case "switchoff":
      st.DeviceSwitchOnOff(*namePtr, false)  

    case "runscene":
      st.RunScene(*namePtr)

    case "switchstatus":
      success, status := st.GetDeviceSwitchStatus(*namePtr)

      if(!success){
        fmt.Println("get_st_device_switch_status failed")
      }else{
        fmt.Println(status.Switch.Value)
        fmt.Println(status.Switch.Timestamp)

        z, _ := status.Switch.Timestamp.Zone()
        fmt.Println("ZONE : ", z, " Time : ", status.Switch.Timestamp) // local time

        location, err := time.LoadLocation("EST")
        if err != nil {
            fmt.Println(err)
        }

        fmt.Println("ZONE : ", location, " Time : ", status.Switch.Timestamp.In(location)) // EST

      }

    default:
      help()
      os.Exit(2)

  }

  os.Exit(0)
     
}
