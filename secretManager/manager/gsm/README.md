# GSM Google Secret Manager
This is the manager implementation for the Google Cloud Secret Manager.

## Setup
1. Place service account json with at least read permission for the Google Cloud Secret Manager under `./key.json`
2. Make use of the secret manager by typing `gsm:` in front of your config. So e.g. `gsm:my_tenant_client_id` where `my_tenant_client_id` is the name of the secret in your Google Cloud

## Env Variables
* `GCP_PROJECT_ID` -> Your Google Cloud project id
* `GOOGLE_APPLICATION_CREDENTIALS` -> The path to your key file. Default is `./key.json`