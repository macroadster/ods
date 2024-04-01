package sdk

import (
    "bytes"
    "encoding/json"
    "strconv"
    "net/http"
    "errors"
)

type Event struct {
	Data string `json:"data"`
}

var SDKEndpoint = "http://localhost:8080"

func SetEndpoint(endpoint string) {
  SDKEndpoint = endpoint
}

func Send(messages []string) error {
  for _, message := range(messages) {
    event := Event{ Data: message }
    body, err := json.Marshal(event);
    if err != nil {
      return err
    }

    r, err := http.NewRequest("POST", SDKEndpoint+"/api/v1/data", bytes.NewBuffer(body))
    if err != nil {
      return err
    }

    r.Header.Add("Content-Type", "application/json")
    client := &http.Client{}
    res, err := client.Do(r)
    if err != nil {
      return err
    }

    defer res.Body.Close()

  	if res.StatusCode != http.StatusOK {
      result := "Error sending message: "+strconv.Itoa(res.StatusCode)
      return errors.New(result)
    }
  }
  return nil
}

func Create(pipeline string) error {
  if pipeline == "" {
    return errors.New("pipeline name can not be null.")
  }
  r, err := http.NewRequest("POST", SDKEndpoint+"/api/v1/pipeline/"+pipeline, nil)
  if err != nil {
    return err
  }
  r.Header.Add("Content-Type", "application/json")
  client := &http.Client{}
  res, err := client.Do(r)
  if err != nil {
    return err
  }
  defer res.Body.Close()
  if res.StatusCode != http.StatusOK {
    result := "Error creating pipeline "+strconv.Itoa(res.StatusCode)
    return errors.New(result)
  }
  return nil
}

func Delete(pipeline string) error {
  if pipeline == "" {
    return errors.New("pipeline name can not be null.")
  }
  r, err := http.NewRequest("DELETE", SDKEndpoint+"/api/v1/pipeline/"+pipeline, nil)
  if err != nil {
    return err
  }
  r.Header.Add("Content-Type", "application/json")
  client := &http.Client{}
  res, err := client.Do(r)
  if err != nil {
    return err
  }
  defer res.Body.Close()
  if res.StatusCode != http.StatusOK {
    result := "Error deleting pipeline "+strconv.Itoa(res.StatusCode)
    return errors.New(result)
  }
  return nil
}

func Inspect(pipeline string) error {
  if pipeline == "" {
    return errors.New("pipeline name can not be null.")
  }
  r, err := http.NewRequest("GET", SDKEndpoint+"/api/v1/pipeline/"+pipeline, nil)
  if err != nil {
    return err
  }
  r.Header.Add("Accept", "application/json")
  client := &http.Client{}
  res, err := client.Do(r)
  if err != nil {
    return err
  }
  defer res.Body.Close()
  if res.StatusCode != http.StatusOK {
    result := "Error inspecting pipeline "+strconv.Itoa(res.StatusCode)
    return errors.New(result)
  }
  return nil
}
