# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Google SHAKEN cert 969H

Tested At: 21 Aug 23 20:17 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -32 day(s)\
Subject: O=Google Voice Inc., C=US, CN=Google SHAKEN cert 969H\
Issuer: L=Lowell, ST=Massachusettes, OU=US, O=NetNumber Inc, C=US, CN=NetNumber SHAKEN Root Intermediate CA 1\
Link: https://www.gstatic.com/gtp/stir/KgjlPz-S4sP2mOhn8gbHeQ.pem

[View certificate details](https://understandingwebpki.com/?cert=MIICzzCCAlagAwIBAgIJAO7Yg99V8Ck7MAoGCCqGSM49BAMCMIGOMTAwLgYDVQQDDCdOZXROdW1iZXIgU0hBS0VOIFJvb3QgSW50ZXJtZWRpYXRlIENBIDExCzAJBgNVBAYTAlVTMRYwFAYDVQQKDA1OZXROdW1iZXIgSW5jMQswCQYDVQQLDAJVUzEXMBUGA1UECAwOTWFzc2FjaHVzZXR0ZXMxDzANBgNVBAcMBkxvd2VsbDAeFw0yMzA2MjAxOTU4NTVaFw0yMzA3MjAxOTU4NTVaMEsxIDAeBgNVBAMMF0dvb2dsZSBTSEFLRU4gY2VydCA5NjlIMQswCQYDVQQGEwJVUzEaMBgGA1UECgwRR29vZ2xlIFZvaWNlIEluYy4wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAT56IQBY6QmPwIvlQtmfood9U2aOaUXbCu%2FENBBh4o0RRZwmc7hPzrVfwWEu1NwNhBXcQ5CD9BeTAAWFZz2ftN%2Bo4HeMIHbMBYGCCsGAQUFBwEaBAowCKAGFgQ5NjlIMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB8GA1UdIwQYMBaAFHEvyILcqOAi49%2BgaPn4XlyE3uW9MB0GA1UdDgQWBBT0OTcvDFSLG7ksvTJe8cNejcIkCjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwGgYDVR0gAQH%2FBBAwDjAMBgpghkgBhv8JAQEBMAoGCCqGSM49BAMCA2cAMGQCMBzXMT6AxzV%2B2UQSU%2FuxpvVPaZv%2ByIfNITYlXuDd9qj%2FW810l5N%2FsspCGyrvSChOBwIwXX6Iiu3NM2P%2FamYR2mSI0K%2BESi3aW1u5MIgIfciVubcpIDKMf4nm%2Frt%2FRu29MGjN)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [n_atis_certificate_policy_critical](../../ISSUES/n_atis_certificate_policy_critical/README.md) | notice | ATIS1000080 | STI certificates should contain a CertificatePolicies extension marked uncritical |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 969H' |


Generated: 21 Aug 23 20:18 UTC