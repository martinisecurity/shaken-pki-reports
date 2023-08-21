# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 633K

Tested At: 21 Aug 23 20:07 UTC\
Initial Validity Period: 367 day(s)\
Remaining Validity Period: 311 day(s)\
Subject: serialNumber=64f123e0-7881-4313-ad74-c3417406ac80, CN=SHAKEN 633K, O=Telkart LLC, C=US\
Issuer: CN=Telonium STI-CA Intermediate CA, O=Telonium STI-CA\
Link: https://certs.telonium.net/23/1c0d4923.crt

[View certificate details](https://understandingwebpki.com/?cert=MIIC5jCCAoygAwIBAgIRAOcO1Mc0te2P%2BUgICVIKkOkwCgYIKoZIzj0EAwIwRDEYMBYGA1UEChMPVGVsb25pdW0gU1RJLUNBMSgwJgYDVQQDEx9UZWxvbml1bSBTVEktQ0EgSW50ZXJtZWRpYXRlIENBMB4XDTIzMDYyNzE1MzA1NFoXDTI0MDYyNzE1MzE1NFowaDELMAkGA1UEBhMCVVMxFDASBgNVBAoTC1RlbGthcnQgTExDMRQwEgYDVQQDEwtTSEFLRU4gNjMzSzEtMCsGA1UEBRMkNjRmMTIzZTAtNzg4MS00MzEzLWFkNzQtYzM0MTc0MDZhYzgwMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEJxvAAzlc1A%2BBnQ8rDrJXJZIpiFknpEFuWEZHwxmlloEw4KRIRVTn%2FOrKsoCSlvs%2B8%2FMxo4FJgZmeW%2FSe2u6l%2FKOCATkwggE1MA4GA1UdDwEB%2FwQEAwIHgDAMBgNVHRMBAf8EAjAAMB0GA1UdDgQWBBQjOM69ucZlyLQUXGdXTgOgQFfVCzAfBgNVHSMEGDAWgBSqJLv%2FFHVAeS2Hb%2BgNQXfKu82IsDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBQGCCsGAQUFBwEaBAgwBhMENjMzSzBdBgwrBgEEAYKkZMYoQAEETTBLAgEBBBlUZWxvbml1bSBTVEktQ0EgU1BDIFRva2VuBCtYVkNWYnNYNmlrd0psaXZCdnlZaGV4SmFiMjNiVHFmcjJhdGxkM2JVLXZFMAoGCCqGSM49BAMCA0gAMEUCIQCHsHXTZj1eTa3avRGLbw1fr11I4EGJ%2FD7Y%2FF%2BYrhwlHQIgbytPNRa7Ph8X0ycOaRYdmPHfeerYjl3takTEI2a71Q0%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Cannot get SPC value from the TNAuthList extension, TNAuthorizationList shall have only one TN Entry |
| [e_atis_tn_auth_list](../../ISSUES/e_atis_tn_auth_list/README.md) | error | ATIS1000080 | TNAuthorizationList shall have only one TN Entry |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | TNAuthorizationList shall have only one TN Entry |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |


Generated: 21 Aug 23 20:18 UTC