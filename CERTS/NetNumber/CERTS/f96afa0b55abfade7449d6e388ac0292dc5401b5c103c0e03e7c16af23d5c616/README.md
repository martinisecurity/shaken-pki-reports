# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Google SHAKEN cert 969H

Tested At: 06 Jan 23 03:03 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -89 day(s)\
Subject: O=Google, C=US, CN=Google SHAKEN cert 969H\
Issuer: L=Lowell, ST=Massachusettes, OU=US, O=NetNumber Inc, C=US, CN=NetNumber SHAKEN Root Intermediate CA 1\
Link: https://www.gstatic.com/gtp/stir/Pkm8HyNNMEiDhJ67_a-tWw.pem

[View certificate details](https://understandingwebpki.com/?cert=MIICxDCCAkugAwIBAgIJAK3pmj%2FQANqyMAoGCCqGSM49BAMCMIGOMTAwLgYDVQQDDCdOZXROdW1iZXIgU0hBS0VOIFJvb3QgSW50ZXJtZWRpYXRlIENBIDExCzAJBgNVBAYTAlVTMRYwFAYDVQQKDA1OZXROdW1iZXIgSW5jMQswCQYDVQQLDAJVUzEXMBUGA1UECAwOTWFzc2FjaHVzZXR0ZXMxDzANBgNVBAcMBkxvd2VsbDAeFw0yMjA5MDkwMDEwMzlaFw0yMjEwMDkwMDEwMzlaMEAxIDAeBgNVBAMMF0dvb2dsZSBTSEFLRU4gY2VydCA5NjlIMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGR29vZ2xlMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEJBcUMLKgZucgu6mk60Rp5Vy7rtAT7LKt4%2BSpX0P4rDC8s%2Bva6PxQX1yecf1163%2Bd3mT4nvTi%2Fri%2FSbq35edRnKOB3jCB2zAWBggrBgEFBQcBGgQKMAigBhYEOTY5SDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAfBgNVHSMEGDAWgBRxL8iC3KjgIuPfoGj5%2BF5chN7lvTAdBgNVHQ4EFgQUV5SQnhscZUB8hgi6PX%2FVBOtirdcwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBoGA1UdIAEB%2FwQQMA4wDAYKYIZIAYb%2FCQEBATAKBggqhkjOPQQDAgNnADBkAjB4pJixDIXI15wG6mytJ9onqd1I2%2FkKfQtaucBzWKrgpbaIcElrS6i9NPE5Zz%2FFVpICMGjlWBoey3KgyqSIs%2Bn%2FtgPaLoZNKCs68lGNNz3cA5gC4OVnQuxMIv%2B3sP23kxLqEw%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 969H' |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [n_atis_certificate_policy_critical](../../ISSUES/n_atis_certificate_policy_critical/README.md) | notice | ATIS1000080 | STI certificates should contain a CertificatePolicies extension marked uncritical |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 06 Jan 23 03:03 UTC