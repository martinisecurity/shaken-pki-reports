# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Google SHAKEN cert 969H

Tested At: 05 Apr 24 19:00 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -29 day(s)\
Subject: O=Google Voice Inc., C=US, CN=Google SHAKEN cert 969H\
Issuer: L=Lowell, ST=Massachusettes, OU=US, O=NetNumber Inc, C=US, CN=NetNumber SHAKEN Root Intermediate CA 1\
Link: https://www.gstatic.com/gtp/stir/lpzBPBX5gc1vz1KynYbDKA.pem

[View certificate details](https://x509.io/?cert=MIIC0DCCAlagAwIBAgIJAJWh7sa4ke%2BoMAoGCCqGSM49BAMCMIGOMTAwLgYDVQQDDCdOZXROdW1iZXIgU0hBS0VOIFJvb3QgSW50ZXJtZWRpYXRlIENBIDExCzAJBgNVBAYTAlVTMRYwFAYDVQQKDA1OZXROdW1iZXIgSW5jMQswCQYDVQQLDAJVUzEXMBUGA1UECAwOTWFzc2FjaHVzZXR0ZXMxDzANBgNVBAcMBkxvd2VsbDAeFw0yNDAyMDYxMjA1MjBaFw0yNDAzMDcxMjA1MjBaMEsxIDAeBgNVBAMMF0dvb2dsZSBTSEFLRU4gY2VydCA5NjlIMQswCQYDVQQGEwJVUzEaMBgGA1UECgwRR29vZ2xlIFZvaWNlIEluYy4wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAT56IQBY6QmPwIvlQtmfood9U2aOaUXbCu%2FENBBh4o0RRZwmc7hPzrVfwWEu1NwNhBXcQ5CD9BeTAAWFZz2ftN%2Bo4HeMIHbMBYGCCsGAQUFBwEaBAowCKAGFgQ5NjlIMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB8GA1UdIwQYMBaAFHEvyILcqOAi49%2BgaPn4XlyE3uW9MB0GA1UdDgQWBBT0OTcvDFSLG7ksvTJe8cNejcIkCjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwGgYDVR0gAQH%2FBBAwDjAMBgpghkgBhv8JAQEBMAoGCCqGSM49BAMCA2gAMGUCMQCO0Mj7ynGlwmcH8W8dy%2B1uYUcvAgbNI6xDoqWooFzehna6WPgU9IiL6y5qESuYVNQCMGHXD69zyzQHxp0k2Zi4LWOxQ9AZeaEJ3okPJh44VAz4NcgH1Eb%2FE4NTWWbMwbEc0w%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_certificate_policies](../../ISSUES/e_atis_ext_certificate_policies/README.md) | error | ATIS1000080 | The Certificate Policies extension is marked as critical |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 969H', but common name is 'Google SHAKEN cert 969H' |


Generated: 05 Apr 24 19:04 UTC