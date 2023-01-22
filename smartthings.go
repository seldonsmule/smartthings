package smartthings


import (
	//"os"
	"fmt"
	//"flag"
	"strings"
        "encoding/json"
	//"math"
       //"github.com/seldonsmule/powerwall"
       //"github.com/seldonsmule/simpleconffile"
        //"time"
        "github.com/seldonsmule/restapi"
        "github.com/seldonsmule/logmsg"

)

type SmartThings struct {

  sToken string // User created personal token to ST account

  sBaseEndpoint string // Base ST endpoint

  oDevices StDevices // List of all devices from the hub
  oScenes StScenes // List of all scenes from the hub

}

func New() *SmartThings {

  st := new(SmartThings)

  st.sBaseEndpoint = "https://api.smartthings.com/v1"

  return(st)

}

func (pST *SmartThings) Dump(){

  fmt.Printf("SmartThings.sToken[%s]\n", pST.sToken)

}

func (pST *SmartThings) SetToken(sToken string) bool{

  pST.sToken = sToken

  if(!pST.GetDeviceList()){
    logmsg.Print(logmsg.Error,"Check Token - GetDeviceList failed")
    return false
  }

  if(!pST.GetSceneList()){
    logmsg.Print(logmsg.Error,"Check Token - GetSceneList failed")
    return false
  }

  return true

}

func (pST *SmartThings) GetStructs() bool{

  if(!pST.GetStructScenes()){
    logmsg.Print(logmsg.Error,"GetStructScenes() failed")
    return false
  }

  if(!pST.GetStructDevices()){
    logmsg.Print(logmsg.Error,"GetStructDevices() failed")
    return false
  }

  return true

}


func (pST *SmartThings) GetStructScenes() bool{

  endpointname := pST.sBaseEndpoint + "/scenes"

  r := restapi.NewGet("getscenes", endpointname)

  r.SetBearerAccessToken(pST.sToken)

  r.JsonOnly()

  if(!r.Send()){
    msg := fmt.Sprintf("Error getting [%s]\n", endpointname)
    //fmt.Printf("Error sending: %s\n", msg)
    logmsg.Print(logmsg.Error, msg)
    return false
  }

  r.SaveResponseBody("st_scenes", "StScenes", false)

  return true

}

func (pST *SmartThings) GetStructDevices() bool{

  endpointname := pST.sBaseEndpoint + "/devices"

  r := restapi.NewGet("getdevices", endpointname)

  r.SetBearerAccessToken(pST.sToken)

  r.JsonOnly()

  if(!r.Send()){
    msg := fmt.Sprintf("Error getting [%s]\n", endpointname)
    //fmt.Printf("Error sending: %s\n", msg)
    logmsg.Print(logmsg.Error, msg)
    return false
  }

  r.SaveResponseBody("st_devices", "StDevices", true)

  return true

}


func (pST *SmartThings) GetDeviceList() bool{

  endpointname := pST.sBaseEndpoint + "/devices"

  //bUpFound := false
  //bDownFound := false

  bRtnValue := true

  r := restapi.NewGet("getdevices", endpointname)

  r.SetBearerAccessToken(pST.sToken)

  r.JsonOnly()

  if(!r.Send()){
    msg := fmt.Sprintf("Error getting [%s]\n", endpointname)
  //  fmt.Printf("Error sending: %s\n", msg)
    logmsg.Print(logmsg.Error, msg)
    return false
  }

  //fmt.Println(r.GetResponseBody())

  json.Unmarshal(r.BodyBytes, &pST.oDevices)

  logmsg.Print(logmsg.Info, fmt.Sprintf("Number of ST Devices[%d]\n", len(pST.oDevices.Items)))

  return bRtnValue

}

func (pST *SmartThings) GetSceneList() bool{

  endpointname := pST.sBaseEndpoint + "/scenes"

  bRtnValue := true

  r := restapi.NewGet("getdevices", endpointname)

  r.SetBearerAccessToken(pST.sToken)

  r.JsonOnly()

  if(!r.Send()){
    msg := fmt.Sprintf("Error getting [%s]\n", endpointname)
  //  fmt.Printf("Error sending: %s\n", msg)
    logmsg.Print(logmsg.Error, msg)
    return false
  }

  //fmt.Println(r.GetResponseBody())

  json.Unmarshal(r.BodyBytes, &pST.oScenes)

  logmsg.Print(logmsg.Info, fmt.Sprintf("Number of ST Scenes[%d]\n", len(pST.oScenes.Items)))

  return bRtnValue

}

func (pST *SmartThings) PrintSceneList() bool{

  for i := 0; i < len(pST.oScenes.Items); i++ {


    fmt.Printf("Scene: [%s] ID is[%s]\n", pST.oScenes.Items[i].SceneName,
                                          pST.oScenes.Items[i].SceneID)

  }

  return true
}

func (pST *SmartThings) PrintDeviceList() bool{

  for i := 0; i < len(pST.oDevices.Items); i++ {


    fmt.Printf("Devices: Label[%s] ID[%s]\n", 
                                          pST.oDevices.Items[i].Label,
                                          pST.oDevices.Items[i].DeviceID)

    fmt.Printf("\tComponent[0] Label[%s]\n", 
                           pST.oDevices.Items[i].Components[0].Label)

    for j := 0; j < len(pST.oDevices.Items[i].Components[0].Capabilities); j++ {

      fmt.Printf("\t\tCap[%d] ID[%s]\n", j,
           pST.oDevices.Items[i].Components[0].Capabilities[j].ID)
      

    }
                                          

  }

  return true
}

func (pST *SmartThings) ValidateDevice(sDeviceName string) bool {

  if(pST.FindDevice(sDeviceName) == -1){
    logmsg.Print(logmsg.Warning, "ValidateDevice: Device Name not found: " + sDeviceName)
    return false
  }

  return true
}

func (pST *SmartThings) FindDevice(sDeviceName string) int {

  for i := 0; i < len(pST.oDevices.Items); i++ {

    if(strings.Compare(pST.oDevices.Items[i].Label, sDeviceName) == 0){
      return(i)
    }

  }

  logmsg.Print(logmsg.Warning, "FindDevice: Device Name not found: " + sDeviceName)

  return -1
}


func (pST *SmartThings) DeviceSwitchOnOff(sDeviceName string, bOn bool) bool{


  sOnOrOff := "on"

  if(!bOn){
    sOnOrOff = "off"
  }


  sCmdStr := fmt.Sprintf("{\"commands\": [ { \"component\": \"main\", \"capability\": \"switch\", \"command\": \"%s\", \"arguments\": [] } ] }", sOnOrOff)


  index := pST.FindDevice(sDeviceName)

  if(index == -1){

    msg := fmt.Sprintf("Unable to locate device[%s]", sDeviceName)
    logmsg.Print(logmsg.Error, msg)
    return false
  }

  endpointname := fmt.Sprintf("%s/devices/%s/commands", pST.sBaseEndpoint, 
                               pST.oDevices.Items[index].DeviceID)

  logmsg.Print(logmsg.Info, endpointname)
  logmsg.Print(logmsg.Info, sCmdStr)

  r := restapi.NewPost("deviceswitch_onoff", endpointname)

  r.SetBearerAccessToken(pST.sToken)

  r.JsonOnly()

  r.SetPostJson(sCmdStr)

  if(!r.Send()){
    msg := fmt.Sprintf("Error getting [%s]\n", endpointname)
    logmsg.Print(logmsg.Error, msg)
    return false
  }

  
  return true

}

func (pST *SmartThings) GetDeviceSwitchStatus(sDeviceName string) (err bool, status StDeviceStatus) {

  var stStatus StDeviceStatus

  index := pST.FindDevice(sDeviceName)

  if(index == -1){

    msg := fmt.Sprintf("Unable to locate device[%s]", sDeviceName)
    logmsg.Print(logmsg.Error, msg)
    return false, stStatus
  }

  endpointname := fmt.Sprintf("%s/devices/%s/components/main/capabilities/switch/status", pST.sBaseEndpoint, pST.oDevices.Items[index].DeviceID)

  logmsg.Print(logmsg.Info, endpointname)

  r := restapi.NewGet("switchstatus", endpointname)

  r.SetBearerAccessToken(pST.sToken)

  //r.HasInnerMap("switch")

  r.JsonOnly()

  if(!r.Send()){
    msg := fmt.Sprintf("Error getting [%s]\n", endpointname)
    logmsg.Print(logmsg.Error, msg)
    return false, stStatus
  }

  //r.Dump()

  //r.SaveResponseBody("st_device_status", "StDeviceStatus", true)

  json.Unmarshal(r.BodyBytes, &stStatus)

  //fmt.Println(stStatus.Switch.Value)
  //fmt.Println(stStatus.Switch.Timestamp)

 // sOnOff := r.GetValueString("value")

  return true, stStatus
}


func (pST *SmartThings) ValidateScene(sName string) bool {

  if(pST.FindScene(sName) == -1){
    logmsg.Print(logmsg.Warning, "ValidateScene: Scene Name not found: " + sName)
    return false
  }

  return true
}

func (pST *SmartThings) FindScene(sName string) int {

  for i := 0; i < len(pST.oScenes.Items); i++ {

    if(strings.Compare(pST.oScenes.Items[i].SceneName, sName) == 0){
      return(i)
    }

  }

  logmsg.Print(logmsg.Warning, "FindScene: Scene Name not found: " + sName)

  return -1
}

func (pST *SmartThings) RunScene(sSceneName string) bool{

  var msg string

  msg = fmt.Sprintf("running SmartThings scene[%s]\n", sSceneName)
  logmsg.Print(logmsg.Info, msg)

  index := pST.FindScene(sSceneName)

  if(index == -1){

    logmsg.Print(logmsg.Error, "RunScene failed - invalid scene name: "+ sSceneName)
    return false
  }
 

  msg = fmt.Sprintf("Executing [%s] \n", sSceneName)
  logmsg.Print(logmsg.Info, msg)

  endpointname := fmt.Sprintf("%s/scenes/%s/execute", 
                          pST.sBaseEndpoint, pST.oScenes.Items[index].SceneID)

  logmsg.Print(logmsg.Info, endpointname)

  r := restapi.NewPost("execute_scenes", endpointname)

  r.SetBearerAccessToken(pST.sToken)

  r.JsonOnly()

  if(!r.Send()){
     msg = fmt.Sprintf("Error getting [%s]\n", endpointname)
     logmsg.Print(logmsg.Error, msg)
     return false
  }

  return true
}

