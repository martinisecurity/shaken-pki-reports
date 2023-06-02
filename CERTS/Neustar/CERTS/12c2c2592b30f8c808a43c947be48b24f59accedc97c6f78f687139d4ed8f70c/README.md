# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN

Tested At: 02 Jun 23 01:11 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 172 day(s)\
Subject: C=US, ST=Tennessee, L=Spring Hill, O=Hadlo Technologies LLC., CN=SHAKEN\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1\
Link: https://www.hadlotechnologies.com/ss_certs/hadlo_stirshaken.public.crt

[View certificate details](https://understandingwebpki.com/?cert=MIID9TCCAt2gAwIBAgIUa7AiMLnziJl2q7kf3qguPZxvim4wDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0yMDExMTkxMTE1NDBaFw0yMzExMjAxMTE1NDBaMGoxDzANBgNVBAMMBlNIQUtFTjEgMB4GA1UECgwXSGFkbG8gVGVjaG5vbG9naWVzIExMQy4xFDASBgNVBAcMC1NwcmluZyBIaWxsMRIwEAYDVQQIDAlUZW5uZXNzZWUxCzAJBgNVBAYTAlVTMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEiCyMvYVnDOSm61xfqj%2FzIe2JhixXa%2BsUKBgJrMdJzkj8GRadisc7QOB5VQLaMZkntccvthQWwHwKnRMxJEYH4KOCAUgwggFEMBYGCCsGAQUFBwEaBAowCKAGFgQ0MzZKMAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAUO7lcyzEXnEjP1Npm6422qNXhCfcwgYEGCCsGAQUFBwEBBHUwczBHBggrBgEFBQcwAoY7aHR0cDovL2NhY2VydHMuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcnQwKAYIKwYBBQUHMAGGHGh0dHA6Ly9vY3NwLWNhMS5jY2lkLm5ldXN0YXIwSAYDVR0fBEEwPzA9oDugOYY3aHR0cDovL2NybC5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNybDAdBgNVHQ4EFgQU7jB8GekEgggNvVkMH%2BH%2BsuK2AGMwDgYDVR0PAQH%2FBAQDAgeAMA0GCSqGSIb3DQEBCwUAA4IBAQBNgl6y2H8dRoNVGy09GJzRvzqWZFnQGPZOJWrXw8eLb7cG9wEMnfUnCt2HcjOQNK%2BQnwKNTqIbX%2B7pS39Gp%2BuwP6SgPXn5ULPB9kFVrnvIP5mbwrl9MkKwOVp%2BqKYwMeJhIgsqegcH9M69Y2Av5txMSJ%2Bi8n%2BVMnW%2Fi44ec2JS4zw6UyE7YL9LbKJjomrXV%2FlOaDLPvTydAmab4PXdSMfCMQmQsXepXzWt%2BDl3sWNZ7CrGOAJJ%2FcLEfXyyyth%2BprpW8nKJUGWMha3s4gph1WwMNyXPvbyNxPyOULAZ6%2Bxjga9cMcr5C7RprdN5Wd8Cs0emGgEX9KjD9OkHby6ZVHdv)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |

### Not Effective

- e_atis_authority_key_identifier
- e_atis_certificate_policies
- e_atis_crl_distribution
- e_atis_extension_unknown
- e_atis_issuer_dn
- e_atis_key_usage
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject
- e_atis_subject_cn
- e_atis_subject_key_identifier
- e_atis_subject_public_key
- e_atis_tn_auth_list
- e_atis_version

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 02 Jun 23 01:12 UTC