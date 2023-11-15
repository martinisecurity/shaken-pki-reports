# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN_0377

Tested At: 15 Nov 23 15:46 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 174 day(s)\
Subject: C=US, ST=Georgia, L=Metter, O=Pineland Telephone Cooperative, OU=Service, CN=SHAKEN_0377\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1\
Link: http://prod001-cr.rbbnidhub.com/MU7B8iqMRz/sign-cert1

[View certificate details](https://understandingwebpki.com/?cert=MIIEDTCCAvWgAwIBAgIUapP3hDvWhGojoOkqiaOb5r2JjtcwDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0yMTA1MDcxNDA5MzRaFw0yNDA1MDcxNDA5MzRaMIGBMRQwEgYDVQQDDAtTSEFLRU5fMDM3NzEQMA4GA1UECwwHU2VydmljZTEnMCUGA1UECgweUGluZWxhbmQgVGVsZXBob25lIENvb3BlcmF0aXZlMQ8wDQYDVQQHDAZNZXR0ZXIxEDAOBgNVBAgMB0dlb3JnaWExCzAJBgNVBAYTAlVTMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE3wwaiDuMPIVcrGZ454uuzhwwti5QCIU7C1bzLF1AzMZVq2IRMZDZOw7JDwtuV3V%2BidIxNUAUlYXHJSMHLRw8XqOCAUgwggFEMBYGCCsGAQUFBwEaBAowCKAGFgQwMzc3MAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAUO7lcyzEXnEjP1Npm6422qNXhCfcwgYEGCCsGAQUFBwEBBHUwczBHBggrBgEFBQcwAoY7aHR0cDovL2NhY2VydHMuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcnQwKAYIKwYBBQUHMAGGHGh0dHA6Ly9vY3NwLWNhMS5jY2lkLm5ldXN0YXIwSAYDVR0fBEEwPzA9oDugOYY3aHR0cDovL2NybC5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNybDAdBgNVHQ4EFgQUvwnTrHqf%2BP7w9wWFo%2FlC10R962EwDgYDVR0PAQH%2FBAQDAgeAMA0GCSqGSIb3DQEBCwUAA4IBAQAP9W3YgMWin%2Bv6lIvmk01yEvB9ywt62I5ZdHtQY042IVq%2BtE0Kz4yHmi8oXDHs%2F2Cd9pqqjZ9rncyJoCA8i1H6zLyJ6ea4WfwcQxEI%2Fx%2BkWMu%2B1fyDHroYwkC%2FfCYKrqILC8aw7gAXIMn5lZf2aHwNMBgQ5rRtTKgOSXuYV1YPRhCTpk%2FwLw92lHyOdAnDBYWihyL4LFbbfSQxOhAzyC4N1QTxfbD03XpnE08VEx704KwWOFBbqYcygztHWXwuYgxZj1raoz3cT%2FXGy50E7pA5%2Bw%2BRSuuS6u3y6bWNJSVNR9mjN8rkGXy5%2FZrxnCQC0l6kwtUPAHcRtWBx44yUGZF2)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |

### Not Effective

- e_atis_extension_unknown
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject_cn

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 15 Nov 23 16:51 UTC