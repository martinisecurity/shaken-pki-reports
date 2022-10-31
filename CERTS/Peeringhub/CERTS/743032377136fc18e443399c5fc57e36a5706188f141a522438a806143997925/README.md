# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Teleinx SHAKEN 744J

Tested At: 31 Oct 22 16:40 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 300 day(s)\
Subject: CN=Teleinx SHAKEN 744J, O=Teleinx LLC, L=Chicago, ST=IL, C=US\
Issuer: CN=Peeringhub Inc SHAKEN Intermediate CA 2, OU=Certification Authorities, O=Peeringhub Inc, C=US\
Link: https://certificates.peeringhub.io/744J/744J.crt

View: [Click to view](https://understandingwebpki.com/?cert=MIIDGTCCAr6gAwIBAgIQLj028KcpXmDFbWAuDEdOyDAKBggqhkjOPQQDAjB8MQswCQYDVQQGEwJVUzEXMBUGA1UECgwOUGVlcmluZ2h1YiBJbmMxIjAgBgNVBAsMGUNlcnRpZmljYXRpb24gQXV0aG9yaXRpZXMxMDAuBgNVBAMMJ1BlZXJpbmdodWIgSW5jIFNIQUtFTiBJbnRlcm1lZGlhdGUgQ0EgMjAeFw0yMjA4MjYyMzMxMzBaFw0yMzA4MjYyMzMxMzBaMGAxCzAJBgNVBAYTAlVTMQswCQYDVQQIDAJJTDEQMA4GA1UEBwwHQ2hpY2FnbzEUMBIGA1UECgwLVGVsZWlueCBMTEMxHDAaBgNVBAMME1RlbGVpbnggU0hBS0VOIDc0NEowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATG%2BFyfEfJrX4ic%2BCwvxuXBpGYz3Hg3jcdNONw3XyI3vENLsxFe%2FPA7sykjNKmSF2eH7p7XFF0JxZerFnFyIzz8o4IBPDCCATgwDgYDVR0PAQH%2FBAQDAgeAMAwGA1UdEwEB%2FwQCMAAwHQYDVR0OBBYEFCWCjv3puPEBWuzbdYxa%2FavObQ%2BCMB8GA1UdIwQYMBaAFK6hc1GIKVcRygyp9LEKbk64S00HMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAWBggrBgEFBQcBGgQKMAigBhYENzQ0SjCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwCgYIKoZIzj0EAwIDSQAwRgIhALdOTyAl%2FGe2JejEyp1tiNSv81kSbKRcjn20MJbfOQr9AiEAnayQzCoGtdNRt0fX0%2BV%2Bl6AYhzEKt1mWPBasFLeH5m8%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_sti_certificate_policies](../../ISSUES/e_sti_certificate_policies/README.md) | error | ATIS&#x2011;1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [w_pki_subject_rdn_unknown](../../ISSUES/w_pki_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_cp1_3_ambiguous_identifier](../../ISSUES/e_cp1_3_ambiguous_identifier/README.md) | error | United States SHAKEN CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_cp1_3_subject_sn](../../ISSUES/e_cp1_3_subject_sn/README.md) | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 31/10/2022 at 16:43:22