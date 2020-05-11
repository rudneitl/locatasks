package main

import (
  "log"
  "os"
  "net/http"
  "encoding/json"
  "golang.org/x/net/context"
  firebase "firebase.google.com/go"
  "github.com/joho/godotenv"
  "google.golang.org/api/option"
)

type Task struct {
  Trigger_word string `firestore:"trigger_word,omitempty"`
  Channel_name string `firestore:"channel_name,omitempty"`
  Timestamp    string `firestore:"timestamp,omitempty"`
  User_name    string `firestore:"user_name,omitempty"`
  Text         string `firestore:"text,omitempty"`
}

func saveTask(w http.ResponseWriter, r *http.Request) {
  godotenv.Load()

  ctx := context.Background()

  sa := option.WithCredentialsJSON([]byte(os.Getenv("JSON_CREDS")))
  //sa := option.WithCredentialsFile("private/key/locatasks.json")

  app, err := firebase.NewApp(ctx, nil, sa)
  if err != nil {
    log.Fatalln(err)
  }

  client, err := app.Firestore(ctx)
  if err != nil {
    log.Fatalln(err)
  }

  decoder := json.NewDecoder(r.Body)

  var task Task
  err = decoder.Decode(&task)
  if err != nil {
    panic(err)
  }

  _, _, err = client.Collection("tasks").Add(ctx, task)

  if err != nil {
    log.Fatalf("Failed adding alovelace: %v", err)
  }

  defer client.Close()
}
