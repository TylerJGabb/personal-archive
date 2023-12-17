# Local Development with Docker and Google Application Credentials
This is a simple example of how to use Google Application Credentials with Docker for local development.

# What it Does
This example uses the [Google Cloud Storage Client](https://cloud.google.com/go/docs/reference/cloud.google.com/go/storage/latest) to list the contents of a bucket.

See [main.go](main.go) for the code.

# How it Works
The docker-compose file mounts the Google Application Credentials file into the container using a volume. The credentials file will be made available at `/tmp/keys/application_default_credentials.json` in the container.

```yaml
    volumes:
      # https://cloud.google.com/run/docs/testing/local
    - ~/.config/gcloud/application_default_credentials.json:/tmp/keys/application_default_credentials.json
```

The docker-compose file also sets the `GOOGLE_APPLICATION_CREDENTIALS` environment variable to the path of the credentials file.

```yaml
    environment:
      GOOGLE_APPLICATION_CREDENTIALS: /tmp/keys/application_default_credentials.json
```

The [storage client](https://cloud.google.com/go/docs/reference/cloud.google.com/go/storage/latest) in the application uses the file pointed to by the `GOOGLE_APPLICATION_CREDENTIALS` env var by default to authenticate with Google Cloud.


# How to Run the Demo
1. `export PROJECT_ID=<project-id>`
2. `gcloud auth login`
3. `gcloud config set project $PROJECT_ID`
4. `gcloud auth application-default login`
5. `docker-compose up`

