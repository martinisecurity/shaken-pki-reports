# STIR/SHAKEN CA Ecosystem Compliance
## Peeringhub

### Certificate 417122a7bf47730e788f52241d189a83b05cedc9
Tested At: 2022-10-26 22:30:10 +0000 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 305 day(s)\
Subject: CN=Teleinx SHAKEN 744J, O=Teleinx LLC, L=Chicago, ST=IL, C=US\
Issuer: CN=Peeringhub Inc SHAKEN Intermediate CA 2, OU=Certification Authorities, O=Peeringhub Inc, C=US

Link: https://certificates.peeringhub.io/744J/744J.crt

View: [Click to view](https://understandingwebpki.com/?cert=MIIDGTCCAr6gAwIBAgIQLj028KcpXmDFbWAuDEdOyDAKBggqhkjOPQQDAjB8MQswCQYDVQQGEwJVUzEXMBUGA1UECgwOUGVlcmluZ2h1YiBJbmMxIjAgBgNVBAsMGUNlcnRpZmljYXRpb24gQXV0aG9yaXRpZXMxMDAuBgNVBAMMJ1BlZXJpbmdodWIgSW5jIFNIQUtFTiBJbnRlcm1lZGlhdGUgQ0EgMjAeFw0yMjA4MjYyMzMxMzBaFw0yMzA4MjYyMzMxMzBaMGAxCzAJBgNVBAYTAlVTMQswCQYDVQQIDAJJTDEQMA4GA1UEBwwHQ2hpY2FnbzEUMBIGA1UECgwLVGVsZWlueCBMTEMxHDAaBgNVBAMME1RlbGVpbnggU0hBS0VOIDc0NEowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATG%2BFyfEfJrX4ic%2BCwvxuXBpGYz3Hg3jcdNONw3XyI3vENLsxFe%2FPA7sykjNKmSF2eH7p7XFF0JxZerFnFyIzz8o4IBPDCCATgwDgYDVR0PAQH%2FBAQDAgeAMAwGA1UdEwEB%2FwQCMAAwHQYDVR0OBBYEFCWCjv3puPEBWuzbdYxa%2FavObQ%2BCMB8GA1UdIwQYMBaAFK6hc1GIKVcRygyp9LEKbk64S00HMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAWBggrBgEFBQcBGgQKMAigBhYENzQ0SjCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwCgYIKoZIzj0EAwIDSQAwRgIhALdOTyAl%2FGe2JejEyp1tiNSv81kSbKRcjn20MJbfOQr9AiEAnayQzCoGtdNRt0fX0%2BV%2Bl6AYhzEKt1mWPBasFLeH5m8%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| w_cp1_3_subject_rdn_unknown | warn | CPv1.3 | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_sti_certificate_policies | error | ATIS-1000080v4 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| e_cp1_3_ambiguous_identifier | error | CPv1.3 | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |

Generated: 26/10/2022 at 22:31:35