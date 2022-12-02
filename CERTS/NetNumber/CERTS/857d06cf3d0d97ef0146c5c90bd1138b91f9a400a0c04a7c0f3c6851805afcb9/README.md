# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Google SHAKEN cert 969H

Tested At: 02 Dec 22 07:46 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -31 day(s)\
Subject: O=Google, C=US, CN=Google SHAKEN cert 969H\
Issuer: L=Lowell, ST=Massachusettes, OU=US, O=NetNumber Inc, C=US, CN=NetNumber SHAKEN Root Intermediate CA 1\
Link: https://www.gstatic.com/gtp/stir/yWm5JCvzCtTKhZihGKFIFA.pem

[View certificate details](https://understandingwebpki.com/?cert=MIICxTCCAkugAwIBAgIJAIYWc8mnJm%2FzMAoGCCqGSM49BAMCMIGOMTAwLgYDVQQDDCdOZXROdW1iZXIgU0hBS0VOIFJvb3QgSW50ZXJtZWRpYXRlIENBIDExCzAJBgNVBAYTAlVTMRYwFAYDVQQKDA1OZXROdW1iZXIgSW5jMQswCQYDVQQLDAJVUzEXMBUGA1UECAwOTWFzc2FjaHVzZXR0ZXMxDzANBgNVBAcMBkxvd2VsbDAeFw0yMjEwMDIwMTAxMzJaFw0yMjExMDEwMTAxMzJaMEAxIDAeBgNVBAMMF0dvb2dsZSBTSEFLRU4gY2VydCA5NjlIMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGR29vZ2xlMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEJBcUMLKgZucgu6mk60Rp5Vy7rtAT7LKt4%2BSpX0P4rDC8s%2Bva6PxQX1yecf1163%2Bd3mT4nvTi%2Fri%2FSbq35edRnKOB3jCB2zAWBggrBgEFBQcBGgQKMAigBhYEOTY5SDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAfBgNVHSMEGDAWgBRxL8iC3KjgIuPfoGj5%2BF5chN7lvTAdBgNVHQ4EFgQUV5SQnhscZUB8hgi6PX%2FVBOtirdcwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBoGA1UdIAEB%2FwQQMA4wDAYKYIZIAYb%2FCQEBATAKBggqhkjOPQQDAgNoADBlAjEA33GEcgAiSK5H4TqazXF6YuOxqu983EueyAS9hFcOwzBL3A3yNdTzoQ%2BQ9A%2Bwp601AjBfJcolaAd2iFZdoZ1FGGiv%2B%2BnS%2BjkWPlfeP7Cjd28YtYNKd6OM0%2FIXc5ta5y0PNEo%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 969H' |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [n_atis_certificate_policy_critical](../../ISSUES/n_atis_certificate_policy_critical/README.md) | notice | ATIS1000080 | STI certificates should contain a CertificatePolicies extension marked uncritical |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 02 Dec 22 07:46 UTC