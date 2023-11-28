# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Google SHAKEN cert 969H

Tested At: 28 Nov 23 20:17 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -15 day(s)\
Subject: O=Google Voice Inc., C=US, CN=Google SHAKEN cert 969H\
Issuer: L=Lowell, ST=Massachusettes, OU=US, O=NetNumber Inc, C=US, CN=NetNumber SHAKEN Root Intermediate CA 1\
Link: https://www.gstatic.com/gtp/stir/w2ZFOTTUJ-cMe4ziMjJj-A.pem

[View certificate details](https://understandingwebpki.com/?cert=MIICzzCCAlWgAwIBAgIIKfzo5GNaZ%2BgwCgYIKoZIzj0EAwIwgY4xMDAuBgNVBAMMJ05ldE51bWJlciBTSEFLRU4gUm9vdCBJbnRlcm1lZGlhdGUgQ0EgMTELMAkGA1UEBhMCVVMxFjAUBgNVBAoMDU5ldE51bWJlciBJbmMxCzAJBgNVBAsMAlVTMRcwFQYDVQQIDA5NYXNzYWNodXNldHRlczEPMA0GA1UEBwwGTG93ZWxsMB4XDTIzMTAxMzIyNTkxOVoXDTIzMTExMjIyNTkxOVowSzEgMB4GA1UEAwwXR29vZ2xlIFNIQUtFTiBjZXJ0IDk2OUgxCzAJBgNVBAYTAlVTMRowGAYDVQQKDBFHb29nbGUgVm9pY2UgSW5jLjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABPnohAFjpCY%2FAi%2BVC2Z%2Bih31TZo5pRdsK78Q0EGHijRFFnCZzuE%2FOtV%2FBYS7U3A2EFdxDkIP0F5MABYVnPZ%2B036jgd4wgdswFgYIKwYBBQUHARoECjAIoAYWBDk2OUgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHwYDVR0jBBgwFoAUcS%2FIgtyo4CLj36Bo%2BfheXITe5b0wHQYDVR0OBBYEFPQ5Ny8MVIsbuSy9Ml7xw16NwiQKMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAaBgNVHSABAf8EEDAOMAwGCmCGSAGG%2FwkBAQEwCgYIKoZIzj0EAwIDaAAwZQIwfVt20op5d6EVwK5sP0wYJ%2B32Isw%2F0AUjb%2FBMeZkA00v0DntpfPo3jmP2YB8EgDriAjEA8%2BB3RXsOpE8NNwijXIrLTa8wKoakLGeYfbT4tRl2%2Bd4LPTTTU84K8vx43rZjzKxs)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_certificate_policies](../../ISSUES/e_atis_ext_certificate_policies/README.md) | error | ATIS1000080 | The Certificate Policies extension is marked as critical |
| [e_atis_serial_number_size](../../ISSUES/e_atis_serial_number_size/README.md) | error | ATIS1000080 | STI certificates shall have a serial number that contains at least 64 bits, got 62 |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 969H', but common name is 'Google SHAKEN cert 969H' |


Generated: 28 Nov 23 20:21 UTC