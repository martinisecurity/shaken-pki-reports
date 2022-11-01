# STIR/SHAKEN CA Ecosystem Compliance

## Certificate PRD

Tested At: 01 Nov 22 22:19 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 553 day(s)\
Subject: C=US, ST=CA, L=San Francisco, O=LiveVox, OU=VOIP, CN=PRD\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1\
Link: https://stir.na3.livevox.com/cert/2B6FU4qN

View: [Click to view](https://understandingwebpki.com/?cert=MIID7DCCAtSgAwIBAgIUQKXL95t9CNlQghuNrFqYqgMzCqAwDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0yMTA1MDcxNjU1NTFaFw0yNDA1MDcxNjU1NTFaMGExDDAKBgNVBAMMA1BSRDENMAsGA1UECwwEVk9JUDEQMA4GA1UECgwHTGl2ZVZveDEWMBQGA1UEBwwNU2FuIEZyYW5jaXNjbzELMAkGA1UECAwCQ0ExCzAJBgNVBAYTAlVTMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE63qlUfjX9AUW8MGzxZsQG5wsNvklVUVZjUEzMOHi88nuBt1jIn72aILqqdl61dxQsEnz9mBjPN2hoaOReAHTxaOCAUgwggFEMBYGCCsGAQUFBwEaBAowCKAGFgQ0NjVKMAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAUO7lcyzEXnEjP1Npm6422qNXhCfcwgYEGCCsGAQUFBwEBBHUwczBHBggrBgEFBQcwAoY7aHR0cDovL2NhY2VydHMuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcnQwKAYIKwYBBQUHMAGGHGh0dHA6Ly9vY3NwLWNhMS5jY2lkLm5ldXN0YXIwSAYDVR0fBEEwPzA9oDugOYY3aHR0cDovL2NybC5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNybDAdBgNVHQ4EFgQUES63SDle%2BABEgwupIfgc%2BLLtXT4wDgYDVR0PAQH%2FBAQDAgeAMA0GCSqGSIb3DQEBCwUAA4IBAQBl5wIkslkGg2nL7SSNKlMIl4XivVm8GMAj4k3MWSAjX%2FjKI%2F7bTjizVXQl5wp5f1whblZjmSb37XZsVI9Xn4OfySTMg3NiHfU6YEoMgcLZ201COf0AdCcU1nsYGHQs7H1i5VQUSmdotJe%2BcI8GbDBL79NTyXWVojM449KIscQc600sR4EWNY1U0SQaVHEBSAFhg7XujEZRF9cjaQTzy4SnJBkDrXzpbDYgw%2F0FpyaWmeTCczI3c0doSt648DmoCQf6jepBbbiuCQlS%2FHXFCJSM5O7BDuxqAwQWfjIzp%2B%2FaD3TrwH2f%2FZj6XiFExLvI684oiTGGTHLgrhEt6kMneseP)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |


### Not Effective

- e_atis_extension_unknown
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject_cn

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 01/11/2022 at 22:19:34