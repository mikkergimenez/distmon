package proc

import (
  "github.com/akhenakh/statgo"
  "code.cloudfoundry.org/bytefmt"

  "strconv"
)

type Proc struct {
  Host *statgo.HostInfos
  CPU *statgo.CPUStats
  Mem *statgo.MemStats
}

func (p *Proc) Get() Proc {
  s := statgo.NewStat()

  return Proc{
    Host: s.HostInfos(),
    CPU: s.CPUStats(),
    Mem: s.MemStats(),
  }
}

func (p *Proc) GetHTML() string {
  str := ""

  s := statgo.NewStat()



  str += "<h2>" + s.HostInfos().HostName + "</h2>"
  str += "Operating System: " + s.HostInfos().OSName  + "<br>"
  str += "OS Release: " + s.HostInfos().OSRelease + "<br>"
  str += "OS Version: " + s.HostInfos().OSVersion + "<br>"
  str += "Platform: " + s.HostInfos().Platform + "<br>"
  str += "CPUs: " + strconv.Itoa(s.HostInfos().NCPUs) + "<br>"
  str += "Max CPUs: " + strconv.Itoa(s.HostInfos().MaxCPUs) + "<br>"
  str += "BitWidth: " + strconv.Itoa(s.HostInfos().BitWidth) + "<br>"
  str += "<br>"
  str += "User CPU: " + strconv.FormatFloat(s.CPUStats().User, 'g', 1, 64) + "<br>"
  str += "Kernel CPU: " + strconv.FormatFloat(s.CPUStats().Kernel, 'g', 1, 64) + "<br>"
  str += "Idle: " + strconv.FormatFloat(s.CPUStats().Idle, 'g', 1, 64) + "<br>"
  str += "IOWait: " + strconv.FormatFloat(s.CPUStats().IOWait, 'g', 1, 64) + "<br>"
  str += "Swap: " + strconv.FormatFloat(s.CPUStats().Swap, 'g', 1, 64) + "<br>"
  str += "Nice: " + strconv.FormatFloat(s.CPUStats().Nice, 'g', 1, 64) + "<br>"
  str += "LoadMin1: " + strconv.FormatFloat(s.CPUStats().LoadMin1, 'g', 1, 64) + "<br>"
  str += "LoadMin5: " + strconv.FormatFloat(s.CPUStats().LoadMin5, 'g', 1, 64) + "<br>"
  str += "LoadMin15: " + strconv.FormatFloat(s.CPUStats().LoadMin15, 'g', 1, 64) + "<br>"

  str += "<br>"
  str += "Total: " + bytefmt.ByteSize(uint64(s.MemStats().Total)) + "<br>"
  str += "Free: " + bytefmt.ByteSize(uint64(s.MemStats().Free)) + "<br>"
  str += "Used: " + bytefmt.ByteSize(uint64(s.MemStats().Used)) + "<br>"
  str += "Cache: " + strconv.Itoa(s.MemStats().Cache) + "<br>"
  str += "SwapTotal: " + strconv.Itoa(s.MemStats().SwapTotal) + "<br>"
  str += "SwapUsed: " + strconv.Itoa(s.MemStats().SwapUsed) + "<br>"
  str += "SwapFree: " + strconv.Itoa(s.MemStats().SwapFree) + "<br>"

  return str

  /*
  stat, err := linuxproc.ReadStat("/proc/stat")
  if err != nil {
    log.Print("stat read fail")
    return str
  }

  str += "<h1>CPU Stats</h1>"
  str += "<table>"
  str += "<tr>"
  str += "<th>User</th>"
  str += "<th>Nice</th>"
  str += "<th>System</th>"
  str += "<th>Idle</th>"
  str += "<th>IOWait</th>"
  str += "</tr>"

  for _, s := range stat.CPUStats {
    str += "<tr>"
    str += "<td>" + strconv.FormatUint(s.User, 10) + "</td>"
    str += "<td>" + strconv.FormatUint(s.Nice, 10) + "</td>"
    str += "<td>" + strconv.FormatUint(s.System, 10) + "</td>"
    str += "<td>" + strconv.FormatUint(s.Idle, 10) + "</td>"
    str += "<td>" + strconv.FormatUint(s.IOWait, 10) + "</td>"
    str += "</tr>"
  }

  str += "</table>"

  return str
  */
}
