# Auth Service

This is a microservice which helps you authenticating different tenant users around your environment.
The idea is that the user can visit `itstenant.yourdomain.com` and it will load the correct configured IdP and start the oauth2 process. After that it will return to the visited domain with the access_token as GET parameter.

## Setup
1. Add the [config](config/README.md) you need. Currently only openid connect and google are possible as IdP's.
2. Inside the config you can use either the [Google Cloud Secret Manager](secretManager/manager/gsm/README.md) or just plain text
