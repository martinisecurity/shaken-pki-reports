# STIR/SHAKEN CA Ecosystem Compliance

## Certificate NovoLink SHAKEN cert

Tested At: 18 Aug 25 20:04 UTC\
Initial Validity Period: 364 day(s)\
Remaining Validity Period: 234 day(s)\
Subject: CN=NovoLink SHAKEN cert, O=NovoLink, ST=Texas, C=US\
Issuer: CN=SHAKEN 1RouteGroup Intermediate CA, O=1RouteGroup, ST=Texas, C=US\
Link: http://sti.novolink.net/crypt/sti-shaken.crt

[View certificate details](https://x509.io/?cert=MIICuTCCAmGgAwIBAgICESEwCgYIKoZIzj0EAwIwYDELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMRQwEgYDVQQKDAsxUm91dGVHcm91cDErMCkGA1UEAwwiU0hBS0VOIDFSb3V0ZUdyb3VwIEludGVybWVkaWF0ZSBDQTAeFw0yNTA0MDkyMjEzMDhaFw0yNjA0MDgyMjEzMDhaME8xCzAJBgNVBAYTAlVTMQ4wDAYDVQQIDAVUZXhhczERMA8GA1UECgwITm92b0xpbmsxHTAbBgNVBAMMFE5vdm9MaW5rIFNIQUtFTiBjZXJ0MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEQ6%2FEnVPtHb9nFRMWwyDey5WCkyerYu4OatuhPmJaZUrFovMgDtC6PtKMHpD7lNBGAiS2l7tahZjckAKocsOqQ6OCARowggEWMBYGCCsGAQUFBwEaBAowCKAGFgQzMzJHMAwGA1UdEwEB%2FwQCMAAwHQYDVR0OBBYEFAfDd2liLKZEfgXS%2B%2F4Ky%2FI%2BvmVLMB8GA1UdIwQYMBaAFKdsSAmToL9B4BNhceYD7TWHHe6BMA4GA1UdDwEB%2FwQEAwIHgDCBhAYDVR0fBH0wezB5oD6gPIY6aHR0cHM6Ly9hdXRoZW50aWNhdGUtZXh0LWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKI3pDUwMzELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTETMBEGA1UEAwwKU1RJLVBBIENSTDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQQwCgYIKoZIzj0EAwIDRgAwQwIgeS9DypA3gIOzXPPfvKERRVtPFQCNbeOmF4MuOHqjWvQCH1CGjysNEwHVKXVpsBgZYkx5yBvJp9V9VLj7jKXT%2Fa8%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_serial_number_size](../../ISSUES/e_atis_serial_number_size/README.md) | error | ATIS1000080 | STI certificates shall have a serial number that contains at least 64 bits, got 13 |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 332G', but common name is 'NovoLink SHAKEN cert' |


Generated: 18 Aug 25 21:13 UTC