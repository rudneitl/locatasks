package main

import (
  "log"
  // "os"
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

// type Credentials struct {
//     Type   string
//     Project_id   string
//     Private_key_id string
//     Private_key string
//     Client_email string
//     Client_id string
//     Auth_uri string
//     Token_uri string
//     Auth_provider_x509_cert_url string
//     Client_x509_cert_url string
// }

func saveTask(w http.ResponseWriter, r *http.Request) {
  godotenv.Load()

  ctx := context.Background()

  // var credentials Credentials
  // credentials.Type = os.Getenv("FIREBASE_TYPE")
  // credentials.Project_id = os.Getenv("FIREBASE_PROJECT_ID")
  // credentials.Private_key_id = os.Getenv("FIREBASE_PRIVATE_KEY_ID")
  // credentials.Private_key = os.Getenv("FIREBASE_PRIVATE_KEY")
  // credentials.Client_email = os.Getenv("FIREBASE_CLIENT_EMAIL")
  // credentials.Client_id = os.Getenv("FIREBASE_CLIENT_ID")
  // credentials.Auth_uri = os.Getenv("FIREBASE_AUTH_URI")
  // credentials.Token_uri = os.Getenv("FIREBASE_TOKEN_URI")
  // credentials.Auth_provider_x509_cert_url = os.Getenv("FIREBASE_AUTH_PROVIDER_X509_CERT_URL")
  // credentials.Client_x509_cert_url = os.Getenv("FIREBASE_CLIENT_X509_CERT_URL")

  // var credentials = &google.Credentials {
  //   Type: os.Getenv("FIREBASE_TYPE"),
  //   Project_id: os.Getenv("FIREBASE_PROJECT_ID"),
  //   Private_key_id: os.Getenv("FIREBASE_PRIVATE_KEY_ID"),
  //   Private_key: os.Getenv("FIREBASE_PRIVATE_KEY"),
  //   Client_email: os.Getenv("FIREBASE_CLIENT_EMAIL"),
  //   Client_id: os.Getenv("FIREBASE_CLIENT_ID"),
  //   Auth_uri: os.Getenv("FIREBASE_AUTH_URI"),
  //   Token_uri: os.Getenv("FIREBASE_TOKEN_URI"),
  //   Auth_provider_x509_cert_url: os.Getenv("FIREBASE_AUTH_PROVIDER_X509_CERT_URL"),
  //   Client_x509_cert_url: os.Getenv("FIREBASE_CLIENT_X509_CERT_URL"),
  // }

  // sa := option.WithCredentials(credentials)
  sa := option.WithCredentialsFile("private/key/locatasks.json")
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
