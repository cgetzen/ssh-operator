package main
import (
  "io/ioutil"
  "github.com/ghodss/yaml"
  "k8s.io/api/apps/v1"
  "fmt"
)

func main() {
  bytes, err := ioutil.ReadFile("pod.yml")
  if err != nil {
    panic(err.Error())
  }

  var spec v1.Deployment
  err = yaml.Unmarshal(bytes, &spec)
  if err != nil {
    panic(err.Error())
  }

  fmt.Println(spec)
}
