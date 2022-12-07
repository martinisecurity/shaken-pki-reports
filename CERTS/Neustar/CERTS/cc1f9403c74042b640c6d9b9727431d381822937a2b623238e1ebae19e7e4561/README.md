# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN

Tested At: 07 Dec 22 18:45 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 328 day(s)\
Subject: C=US, ST=Tennessee, L=Spring Hill, O=Simwood Inc., CN=SHAKEN\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1\
Link: https://cdn.simwood.com/shaken/simwood.crt

[View certificate details](https://understandingwebpki.com/?cert=MIID6jCCAtKgAwIBAgIUKEJQ2mnJVc8nv9TFUudgn5LuOxYwDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0yMDEwMzAxNDM5MDdaFw0yMzEwMzExNDM5MDdaMF8xDzANBgNVBAMMBlNIQUtFTjEVMBMGA1UECgwMU2ltd29vZCBJbmMuMRQwEgYDVQQHDAtTcHJpbmcgSGlsbDESMBAGA1UECAwJVGVubmVzc2VlMQswCQYDVQQGEwJVUzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABDYhFziaIymbUdC4Y1sYbHfAsPAFYtRZZOVjro68%2FOIQb2PDxQ%2BFkTAZXNMSt%2BoWmS%2FPePsAqKrMa878ZUnFE3mjggFIMIIBRDAWBggrBgEFBQcBGgQKMAigBhYENTEwSjAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFDu5XMsxF5xIz9TaZuuNtqjV4Qn3MIGBBggrBgEFBQcBAQR1MHMwRwYIKwYBBQUHMAKGO2h0dHA6Ly9jYWNlcnRzLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3J0MCgGCCsGAQUFBzABhhxodHRwOi8vb2NzcC1jYTEuY2NpZC5uZXVzdGFyMEgGA1UdHwRBMD8wPaA7oDmGN2h0dHA6Ly9jcmwuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcmwwHQYDVR0OBBYEFFhmzJWk3uspalTLrh8QXg%2BC8EYdMA4GA1UdDwEB%2FwQEAwIHgDANBgkqhkiG9w0BAQsFAAOCAQEAlnp2DHojxR%2F0rZnSZz3WPSVUKLA2c%2FGGUWkcWmO6QHlmXvu83gZcQ2mpPnVvDNNV8vYuxoFnsfkFRXQdxZy4J8TYY4dgSwnljzs6rcK8iFrwNDwUYLQ%2B6ylRPbQvp9tKNx3g6ht1tpu1CBn%2FqiVlXcadt06OwBry6rKEgKUQGSldi1gIm3SvK7RCM27oGQ6cCb%2FXxykxNqdu4hCd%2B4cX9izVwdjvn9rzTj%2BPsbQgZDtInPDAMSfArR8Sw5lpci9CCCM7J6gmrLpIE5qW73fwWnZ%2BaeBM8nlCTkmOweQV2JCr1iB0AlddP9d7Q6tmQBnT%2F1Y1M8joEI%2F%2Bx2YeGXda6A%3D%3D)

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


Generated: 07 Dec 22 18:54 UTC