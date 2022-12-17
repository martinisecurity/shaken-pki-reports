# STIR/SHAKEN CA Ecosystem Compliance

## Certificate intrado.com

Tested At: 17 Dec 22 00:03 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 274 day(s)\
Subject: C=US, ST=NE, L=Omaha, O=Intrado Corporation, OU=Intrado Communications, CN=intrado.com\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11125.10033

[View certificate details](https://understandingwebpki.com/?cert=MIIECjCCAvKgAwIBAgIUBy%2BzRv3Fg9%2FU54o271B0P76yIUYwDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0yMDA5MTUxMzE2MjVaFw0yMzA5MTYxMzE2MjVaMH8xFDASBgNVBAMMC2ludHJhZG8uY29tMR8wHQYDVQQLDBZJbnRyYWRvIENvbW11bmljYXRpb25zMRwwGgYDVQQKDBNJbnRyYWRvIENvcnBvcmF0aW9uMQ4wDAYDVQQHDAVPbWFoYTELMAkGA1UECAwCTkUxCzAJBgNVBAYTAlVTMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEEqChMF5v9WbqBa8kJ1YVT4FzSZfqD5V7o1ZIafETxfdamVy6QePW8VKVt2tOlvxlpUXoinWFisw%2FAW47wBTigKOCAUgwggFEMAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAUO7lcyzEXnEjP1Npm6422qNXhCfcwgYEGCCsGAQUFBwEBBHUwczBHBggrBgEFBQcwAoY7aHR0cDovL2NhY2VydHMuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcnQwKAYIKwYBBQUHMAGGHGh0dHA6Ly9vY3NwLWNhMS5jY2lkLm5ldXN0YXIwSAYDVR0fBEEwPzA9oDugOYY3aHR0cDovL2NybC5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNybDAdBgNVHQ4EFgQU7g2uMRFn%2B5JpAAyTb93eR2eWfT4wDgYDVR0PAQH%2FBAQDAgeAMBYGCCsGAQUFBwEaBAowCKAGFgQ5NTU1MA0GCSqGSIb3DQEBCwUAA4IBAQBQ2AO5ORH5gUo4w7FJJVN2%2FtZG4dWvORsEnf0exRPWSdyM8j8N6qm4DmSFT5mrChg0yOLNj%2BfwAoucM2olzZwdmzZ2YaqQbf6NZSnBqetyqlZKzwXtPBgZ%2F9mFRdMBSfRzozmL5YKwMHU4XJIzIkeZpL%2Bg%2FUFuruJg9m%2FRwejk4M0eq6jcpQ8pBFIyxhspmJUgM1VvcqyxdSIq7AxZAn3ExLYzgMBu02U411afVnlk%2FjrNYMhR1oBAZI1wWy5FTxv0MOju5%2Bu1mne1bVgWRNdTUO446qzPqyy6YOlhRE3oaCgFKOR7VgwJRlXkNIon5rgEs8Ux%2BPqCaTZHHDwhJeUy)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |

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


Generated: 17 Dec 22 00:12 UTC