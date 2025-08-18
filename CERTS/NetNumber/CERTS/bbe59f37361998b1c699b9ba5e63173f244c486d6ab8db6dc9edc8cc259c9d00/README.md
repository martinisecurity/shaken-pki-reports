# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Plivo Inc

Tested At: 18 Aug 25 20:05 UTC\
Initial Validity Period: 364 day(s)\
Remaining Validity Period: 233 day(s)\
Subject: L=Austin, ST=Texas, O=Plivo Inc, C=US, CN=Plivo Inc\
Issuer: L=Lowell, ST=Massachusettes, OU=US, O=NetNumber Inc, C=US, CN=NetNumber SHAKEN Root Intermediate CA 1\
Link: https://certificate.zt.plivo.com/cert20260408.crt

[View certificate details](https://x509.io/?cert=MIIC2zCCAmGgAwIBAgIJAMU7tWDnuKr%2BMAoGCCqGSM49BAMCMIGOMTAwLgYDVQQDDCdOZXROdW1iZXIgU0hBS0VOIFJvb3QgSW50ZXJtZWRpYXRlIENBIDExCzAJBgNVBAYTAlVTMRYwFAYDVQQKDA1OZXROdW1iZXIgSW5jMQswCQYDVQQLDAJVUzEXMBUGA1UECAwOTWFzc2FjaHVzZXR0ZXMxDzANBgNVBAcMBkxvd2VsbDAeFw0yNTA0MDkwMDAwMDBaFw0yNjA0MDgwMDAwMDBaMFYxEjAQBgNVBAMMCVBsaXZvIEluYzELMAkGA1UEBhMCVVMxEjAQBgNVBAoMCVBsaXZvIEluYzEOMAwGA1UECAwFVGV4YXMxDzANBgNVBAcMBkF1c3RpbjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABIm4JrOIcWeyvHjcYRbFSmpv1Zcpk05AAS0OZbDv4IwJe8TkFtSqd6%2FqXCxHrfZKtHYZUVvZ6zkKu%2FQR9Y0I%2Bwujgd4wgdswFgYIKwYBBQUHARoECjAIoAYWBDgwMEowDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHwYDVR0jBBgwFoAUcS%2FIgtyo4CLj36Bo%2BfheXITe5b0wHQYDVR0OBBYEFMHyXWhSlyDSvqNWR3zP2FM0J94iMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAaBgNVHSABAf8EEDAOMAwGCmCGSAGG%2FwkBAQEwCgYIKoZIzj0EAwIDaAAwZQIxAOE%2F7sCJfmnz85Fxkh8JM9y2qQOTSXwLnumcsjxscO52s9pZxXM5FuXJnd1g21essQIwQ8zD4ABQwcsBLsSy61skcf8DsmySjoQY5VkaSCNYgs4IaOSqxHTYC%2BZ9O2S3zerI)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_certificate_policies](../../ISSUES/e_atis_ext_certificate_policies/README.md) | error | ATIS1000080 | The Certificate Policies extension is marked as critical |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 800J', but common name is 'Plivo Inc' |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common Name attribute 'Plivo Inc' does not contain 'SHAKEN' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |


Generated: 18 Aug 25 21:13 UTC