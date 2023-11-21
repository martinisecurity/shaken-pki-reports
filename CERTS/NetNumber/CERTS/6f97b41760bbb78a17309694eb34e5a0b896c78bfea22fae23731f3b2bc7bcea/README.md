# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Google SHAKEN cert 969H

Tested At: 21 Nov 23 17:12 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -308 day(s)\
Subject: O=Google Voice Inc., C=US, CN=Google SHAKEN cert 969H\
Issuer: L=Lowell, ST=Massachusettes, OU=US, O=NetNumber Inc, C=US, CN=NetNumber SHAKEN Root Intermediate CA 1\
Link: https://www.gstatic.com/gtp/stir/yExHv85COF4_CVlvUiqb8A.pem

[View certificate details](https://understandingwebpki.com/?cert=MIICzzCCAlWgAwIBAgIICCKQagNsIWYwCgYIKoZIzj0EAwIwgY4xMDAuBgNVBAMMJ05ldE51bWJlciBTSEFLRU4gUm9vdCBJbnRlcm1lZGlhdGUgQ0EgMTELMAkGA1UEBhMCVVMxFjAUBgNVBAoMDU5ldE51bWJlciBJbmMxCzAJBgNVBAsMAlVTMRcwFQYDVQQIDA5NYXNzYWNodXNldHRlczEPMA0GA1UEBwwGTG93ZWxsMB4XDTIyMTIxODE0NTk1OVoXDTIzMDExNzE0NTk1OVowSzEgMB4GA1UEAwwXR29vZ2xlIFNIQUtFTiBjZXJ0IDk2OUgxCzAJBgNVBAYTAlVTMRowGAYDVQQKDBFHb29nbGUgVm9pY2UgSW5jLjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABPnohAFjpCY%2FAi%2BVC2Z%2Bih31TZo5pRdsK78Q0EGHijRFFnCZzuE%2FOtV%2FBYS7U3A2EFdxDkIP0F5MABYVnPZ%2B036jgd4wgdswFgYIKwYBBQUHARoECjAIoAYWBDk2OUgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHwYDVR0jBBgwFoAUcS%2FIgtyo4CLj36Bo%2BfheXITe5b0wHQYDVR0OBBYEFPQ5Ny8MVIsbuSy9Ml7xw16NwiQKMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAaBgNVHSABAf8EEDAOMAwGCmCGSAGG%2FwkBAQEwCgYIKoZIzj0EAwIDaAAwZQIxAJGUfIoP6dNdUm9Gp6dOh9cUIIZE1UwY8hekVp0Om08g88gxuVE7tFY12JDDqI2ZCQIwN69Chz0oROiJvchYbKdPNQFF018%2BYjKRSPurhzcF95uMxcxDbn6%2BEV%2FrI9F6bV4m)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | the Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: [2.16.840.1.114569.1.1.3 2.16.840.1.114569.1.1.4] |
| [e_atis_ext_certificate_policies](../../ISSUES/e_atis_ext_certificate_policies/README.md) | error | ATIS1000080 | the Certificate Policies extension is marked as critical |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 969H' |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 21 Nov 23 17:16 UTC