# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Google SHAKEN cert 969H

Tested At: 31 Oct 22 16:43 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 24 day(s)\
Subject: O=Google, C=US, CN=Google SHAKEN cert 969H\
Issuer: L=Lowell, ST=Massachusettes, OU=US, O=NetNumber Inc, C=US, CN=NetNumber SHAKEN Root Intermediate CA 1\
Link: https://www.gstatic.com/gtp/stir/sFCvi_7snpTdPyz7VUUx8g.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIICxjCCAkugAwIBAgIJAKoe4K1JZ1j3MAoGCCqGSM49BAMCMIGOMTAwLgYDVQQDDCdOZXROdW1iZXIgU0hBS0VOIFJvb3QgSW50ZXJtZWRpYXRlIENBIDExCzAJBgNVBAYTAlVTMRYwFAYDVQQKDA1OZXROdW1iZXIgSW5jMQswCQYDVQQLDAJVUzEXMBUGA1UECAwOTWFzc2FjaHVzZXR0ZXMxDzANBgNVBAcMBkxvd2VsbDAeFw0yMjEwMjUwMTA4MzFaFw0yMjExMjQwMTA4MzFaMEAxIDAeBgNVBAMMF0dvb2dsZSBTSEFLRU4gY2VydCA5NjlIMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGR29vZ2xlMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEJBcUMLKgZucgu6mk60Rp5Vy7rtAT7LKt4%2BSpX0P4rDC8s%2Bva6PxQX1yecf1163%2Bd3mT4nvTi%2Fri%2FSbq35edRnKOB3jCB2zAWBggrBgEFBQcBGgQKMAigBhYEOTY5SDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAfBgNVHSMEGDAWgBRxL8iC3KjgIuPfoGj5%2BF5chN7lvTAdBgNVHQ4EFgQUV5SQnhscZUB8hgi6PX%2FVBOtirdcwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBoGA1UdIAEB%2FwQQMA4wDAYKYIZIAYb%2FCQEBATAKBggqhkjOPQQDAgNpADBmAjEAlsLIpfywDHGLW0Z06o6RDbvqAQhP6485S9vw%2BG0iAemkzuwA8oWoYJX6zI9AT8EqAjEA4koIb5jsGGRl%2FZBGXupQEN7QGoYcoaJjgbcsVhYWx4XR6PegUaLeW3RJhS3ZdQ9I)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [n_sti_certificate_policy_critical](../../ISSUES/n_sti_certificate_policy_critical/README.md) | notice | ATIS&#x2011;1000080 | STI certificates should contain a CertificatePolicies extension marked uncritical |
| [e_sti_subject_cn](../../ISSUES/e_sti_subject_cn/README.md) | error | ATIS&#x2011;1000080 | Common name shall contain the text string 'SHAKEN 969H' |
| [e_cp1_3_subject_sn](../../ISSUES/e_cp1_3_subject_sn/README.md) | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_sti_certificate_policies](../../ISSUES/e_sti_certificate_policies/README.md) | error | ATIS&#x2011;1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_cp1_3_ambiguous_identifier](../../ISSUES/e_cp1_3_ambiguous_identifier/README.md) | error | United States SHAKEN CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 31/10/2022 at 16:43:22