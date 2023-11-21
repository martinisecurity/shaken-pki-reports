# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN_0377

Tested At: 21 Nov 23 16:41 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 168 day(s)\
Subject: C=US, ST=Georgia, L=Metter, O=Pineland Telephone Cooperative, OU=Service, CN=SHAKEN_0377\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1\
Link: http://prod001-cr.rbbnidhub.com/MU7B8iqMRz/sign-cert1

[View certificate details](https://understandingwebpki.com/?cert=MIIEDTCCAvWgAwIBAgIUapP3hDvWhGojoOkqiaOb5r2JjtcwDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0yMTA1MDcxNDA5MzRaFw0yNDA1MDcxNDA5MzRaMIGBMRQwEgYDVQQDDAtTSEFLRU5fMDM3NzEQMA4GA1UECwwHU2VydmljZTEnMCUGA1UECgweUGluZWxhbmQgVGVsZXBob25lIENvb3BlcmF0aXZlMQ8wDQYDVQQHDAZNZXR0ZXIxEDAOBgNVBAgMB0dlb3JnaWExCzAJBgNVBAYTAlVTMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE3wwaiDuMPIVcrGZ454uuzhwwti5QCIU7C1bzLF1AzMZVq2IRMZDZOw7JDwtuV3V%2BidIxNUAUlYXHJSMHLRw8XqOCAUgwggFEMBYGCCsGAQUFBwEaBAowCKAGFgQwMzc3MAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAUO7lcyzEXnEjP1Npm6422qNXhCfcwgYEGCCsGAQUFBwEBBHUwczBHBggrBgEFBQcwAoY7aHR0cDovL2NhY2VydHMuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcnQwKAYIKwYBBQUHMAGGHGh0dHA6Ly9vY3NwLWNhMS5jY2lkLm5ldXN0YXIwSAYDVR0fBEEwPzA9oDugOYY3aHR0cDovL2NybC5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNybDAdBgNVHQ4EFgQUvwnTrHqf%2BP7w9wWFo%2FlC10R962EwDgYDVR0PAQH%2FBAQDAgeAMA0GCSqGSIb3DQEBCwUAA4IBAQAP9W3YgMWin%2Bv6lIvmk01yEvB9ywt62I5ZdHtQY042IVq%2BtE0Kz4yHmi8oXDHs%2F2Cd9pqqjZ9rncyJoCA8i1H6zLyJ6ea4WfwcQxEI%2Fx%2BkWMu%2B1fyDHroYwkC%2FfCYKrqILC8aw7gAXIMn5lZf2aHwNMBgQ5rRtTKgOSXuYV1YPRhCTpk%2FwLw92lHyOdAnDBYWihyL4LFbbfSQxOhAzyC4N1QTxfbD03XpnE08VEx704KwWOFBbqYcygztHWXwuYgxZj1raoz3cT%2FXGy50E7pA5%2Bw%2BRSuuS6u3y6bWNJSVNR9mjN8rkGXy5%2FZrxnCQC0l6kwtUPAHcRtWBx44yUGZF2)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_certificate_policies](../../ISSUES/e_atis_ext_certificate_policies/README.md) | error | ATIS1000080 | the Certificate Policies extension is not present |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common Name attribute 'SHAKEN_0377' does not contain 'SHAKEN' |
| [e_atis_signature_algorithm](../../ISSUES/e_atis_signature_algorithm/README.md) | error | ATIS1000080 | SignatureAlgorithm field is not 'ecdsa-with-SHA256', got 1.2.840.113549.1.1.11 |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_cn_spc
- e_atis_subject_key_identifier_size
- e_atis_subject_o_required
- e_atis_tn_auth_list_spc_format
- e_shaken_certificate_policies_id

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 21 Nov 23 17:16 UTC