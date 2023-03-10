# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Google SHAKEN cert 969H

Tested At: 10 Mar 23 02:25 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -28 day(s)\
Subject: O=Google Voice Inc., C=US, CN=Google SHAKEN cert 969H\
Issuer: L=Lowell, ST=Massachusettes, OU=US, O=NetNumber Inc, C=US, CN=NetNumber SHAKEN Root Intermediate CA 1\
Link: https://www.gstatic.com/gtp/stir/94BX3GNPTFDrpdrUY-ndIw.pem

[View certificate details](https://understandingwebpki.com/?cert=MIICzzCCAlWgAwIBAgIIFxcu8dc9x7kwCgYIKoZIzj0EAwIwgY4xMDAuBgNVBAMMJ05ldE51bWJlciBTSEFLRU4gUm9vdCBJbnRlcm1lZGlhdGUgQ0EgMTELMAkGA1UEBhMCVVMxFjAUBgNVBAoMDU5ldE51bWJlciBJbmMxCzAJBgNVBAsMAlVTMRcwFQYDVQQIDA5NYXNzYWNodXNldHRlczEPMA0GA1UEBwwGTG93ZWxsMB4XDTIzMDExMDE1MDgzM1oXDTIzMDIwOTE1MDgzM1owSzEgMB4GA1UEAwwXR29vZ2xlIFNIQUtFTiBjZXJ0IDk2OUgxCzAJBgNVBAYTAlVTMRowGAYDVQQKDBFHb29nbGUgVm9pY2UgSW5jLjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABPnohAFjpCY%2FAi%2BVC2Z%2Bih31TZo5pRdsK78Q0EGHijRFFnCZzuE%2FOtV%2FBYS7U3A2EFdxDkIP0F5MABYVnPZ%2B036jgd4wgdswFgYIKwYBBQUHARoECjAIoAYWBDk2OUgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHwYDVR0jBBgwFoAUcS%2FIgtyo4CLj36Bo%2BfheXITe5b0wHQYDVR0OBBYEFPQ5Ny8MVIsbuSy9Ml7xw16NwiQKMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAaBgNVHSABAf8EEDAOMAwGCmCGSAGG%2FwkBAQEwCgYIKoZIzj0EAwIDaAAwZQIwMeHIUQUqFqgbE7euM7D95q04gskbEBKsVqX%2FXFnlsqTaXmW3TRDTP1vPfdlmhEvnAjEAsju4EHzIgcbf56kR%2BGnFLj3lmH%2FEi7kDk6iewmMxMjpiiaBcdgfCi4tYExqcq4Mq)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 969H' |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [n_atis_certificate_policy_critical](../../ISSUES/n_atis_certificate_policy_critical/README.md) | notice | ATIS1000080 | STI certificates should contain a CertificatePolicies extension marked uncritical |


Generated: 10 Mar 23 02:25 UTC