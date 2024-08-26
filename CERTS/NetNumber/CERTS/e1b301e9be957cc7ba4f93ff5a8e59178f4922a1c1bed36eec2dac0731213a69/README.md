# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Google SHAKEN cert 969H

Tested At: 26 Aug 24 18:01 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -241 day(s)\
Subject: O=Google Voice Inc., C=US, CN=Google SHAKEN cert 969H\
Issuer: L=Lowell, ST=Massachusettes, OU=US, O=NetNumber Inc, C=US, CN=NetNumber SHAKEN Root Intermediate CA 1\
Link: https://www.gstatic.com/gtp/stir/09khsjCYnY7ePU2oOLObNQ.pem

[View certificate details](https://x509.io/?cert=MIICzjCCAlWgAwIBAgIITDGrZYy4iI4wCgYIKoZIzj0EAwIwgY4xMDAuBgNVBAMMJ05ldE51bWJlciBTSEFLRU4gUm9vdCBJbnRlcm1lZGlhdGUgQ0EgMTELMAkGA1UEBhMCVVMxFjAUBgNVBAoMDU5ldE51bWJlciBJbmMxCzAJBgNVBAsMAlVTMRcwFQYDVQQIDA5NYXNzYWNodXNldHRlczEPMA0GA1UEBwwGTG93ZWxsMB4XDTIzMTEyOTA0MDEyOFoXDTIzMTIyOTA0MDEyOFowSzEgMB4GA1UEAwwXR29vZ2xlIFNIQUtFTiBjZXJ0IDk2OUgxCzAJBgNVBAYTAlVTMRowGAYDVQQKDBFHb29nbGUgVm9pY2UgSW5jLjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABPnohAFjpCY%2FAi%2BVC2Z%2Bih31TZo5pRdsK78Q0EGHijRFFnCZzuE%2FOtV%2FBYS7U3A2EFdxDkIP0F5MABYVnPZ%2B036jgd4wgdswFgYIKwYBBQUHARoECjAIoAYWBDk2OUgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHwYDVR0jBBgwFoAUcS%2FIgtyo4CLj36Bo%2BfheXITe5b0wHQYDVR0OBBYEFPQ5Ny8MVIsbuSy9Ml7xw16NwiQKMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAaBgNVHSABAf8EEDAOMAwGCmCGSAGG%2FwkBAQEwCgYIKoZIzj0EAwIDZwAwZAIwLFlf53Jbb3Dfsk0%2FfpHTVlpfcfarAlkPDkGkgH8hpUZ%2B5bhSVwou4QaGlWraIDVsAjBIIGvWK%2FLANBnXJCdvZU4T7M1%2FAoqFSs94OEybLS%2FQHWbKO3AdXvTEtfpHOe1iE%2FY%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 969H', but common name is 'Google SHAKEN cert 969H' |
| [e_atis_serial_number_size](../../ISSUES/e_atis_serial_number_size/README.md) | error | ATIS1000080 | STI certificates shall have a serial number that contains at least 64 bits, got 63 |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |
| [e_atis_ext_certificate_policies](../../ISSUES/e_atis_ext_certificate_policies/README.md) | error | ATIS1000080 | The Certificate Policies extension is marked as critical |


Generated: 26 Aug 24 18:03 UTC