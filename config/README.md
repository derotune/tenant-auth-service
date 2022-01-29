# Configuration
This package is responsible for the config loading. You can define your personal configuration inside the `config.yml` file. 

The idea is that you can define multiple hosts here, so really use this service around your whole cloud and for every tenant you might have. 

Every host starts with a regex. The auth service tries to match the given `returUrl` over the regex to find the matching configuration. If found it tries to match the tenant. So e.g. tenant1 can have a different configuration for the service X than for the service Y.

## Example Configuration:
```yaml
hosts:
  - regex: "(?s)^(?:https?:\\/\\/)?(?P<Client>[a-z0-9_-]+)\\.mydomain\\.com([\\s\\S]*)$"
    tenants:
      - name: "mytenant"
        clientId: "gsm:clientId"
        clientSecret: "gsm:clientSecret"
        provider: "google"
      - name: "othertenant"
        clientId: "plaintext_client_id"
        clientSecret: "plaintext_client_secret"
        provider: 
          name: "oidc"
          issuer: "OIDC issuer url"
```