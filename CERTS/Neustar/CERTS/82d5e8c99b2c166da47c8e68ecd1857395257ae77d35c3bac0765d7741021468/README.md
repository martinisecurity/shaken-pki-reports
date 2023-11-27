# STIR/SHAKEN CA Ecosystem Compliance

## Certificate ATMC

Tested At: 27 Nov 23 22:28 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 167 day(s)\
Subject: C=US, ST=NC, L=Shallotte, O=ATMC, OU=ATMC, CN=ATMC\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11278.10190.pem

[View certificate details](https://understandingwebpki.com/?cert=MIID5jCCAs6gAwIBAgIUVT9IRL%2FpNkhYtYVDtSYxfGlFKXQwDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0yMTA1MTIyMjM3MDRaFw0yNDA1MTIyMjM3MDRaMFsxDTALBgNVBAMMBEFUTUMxDTALBgNVBAsMBEFUTUMxDTALBgNVBAoMBEFUTUMxEjAQBgNVBAcMCVNoYWxsb3R0ZTELMAkGA1UECAwCTkMxCzAJBgNVBAYTAlVTMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAErKWBW919wNmHUFIb9%2BBLJfUDIoWQhneT9urlyM7guhxCAj%2Fjmi2TsddXLBsQSM2KHiDUkQkNVGjzsbCkRd0lRaOCAUgwggFEMAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAUO7lcyzEXnEjP1Npm6422qNXhCfcwgYEGCCsGAQUFBwEBBHUwczBHBggrBgEFBQcwAoY7aHR0cDovL2NhY2VydHMuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcnQwKAYIKwYBBQUHMAGGHGh0dHA6Ly9vY3NwLWNhMS5jY2lkLm5ldXN0YXIwSAYDVR0fBEEwPzA9oDugOYY3aHR0cDovL2NybC5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNybDAdBgNVHQ4EFgQU1z9sCCqIuFr1gFNg%2FgTFqqKPkJEwDgYDVR0PAQH%2FBAQDAgeAMBYGCCsGAQUFBwEaBAowCKAGFgQwNDY4MA0GCSqGSIb3DQEBCwUAA4IBAQC0VRATd8bNMXJsI2l3vGXk9XOYJOwW9S6jXjwBpWloS1V6bGrg8G6bbnNN4hn4x6OqtJuwN534kiVTzGX7ELI4tydr8GqpzDcZdD9uyt3kKnsDBWgdG4Bfbm5%2Bhd4mWrqYTfqa2zPu2iU9uzlsLYypyvKBQDm2MkrECdos%2BoPfHCSkTw2s2WqJx%2Byb8JmQA%2FSXfHIU9F%2BYGKrjDv2k4zxsCohTPsbQmPqfML4IKLlo3yd%2B7Hef%2BhrtpEJAt4CNaiGNlW%2BwK5yjY9xP4N8h6A1FmI3T0dfGPzLOR1eZvt7OANVdBBzkV5xziel3NjfQ%2FBXGoAejTGyfdzzC4Bu36NfG)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_certificate_policies](../../ISSUES/e_atis_ext_certificate_policies/README.md) | error | ATIS1000080 | The Certificate Policies extension is not present |
| [e_atis_signature_algorithm](../../ISSUES/e_atis_signature_algorithm/README.md) | error | ATIS1000080 | SignatureAlgorithm field is not 'ecdsa-with-SHA256', got 1.2.840.113549.1.1.11 |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common Name attribute 'ATMC' does not contain 'SHAKEN' |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_subject_cn_spc
- e_atis_subject_key_identifier_size
- e_atis_subject_o_required
- e_atis_tn_auth_list_spc_format
- e_shaken_certificate_policies_id

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 27 Nov 23 22:56 UTC