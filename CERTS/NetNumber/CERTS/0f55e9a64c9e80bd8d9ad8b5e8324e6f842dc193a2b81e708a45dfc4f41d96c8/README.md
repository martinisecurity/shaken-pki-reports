# STIR/SHAKEN CA Ecosystem Compliance

## Certificate DISH Wireless L.L.C.SHAKEN.490J

Tested At: 18 Aug 25 21:09 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 45 day(s)\
Subject: L=Denver, ST=CO, OU=Security Operations, O=Dish Wireless, C=US, CN=DISH Wireless L.L.C.SHAKEN.490J\
Issuer: L=Lowell, ST=Massachusettes, OU=US, O=NetNumber Inc, C=US, CN=NetNumber SHAKEN Root Intermediate CA 1\
Link: https://stsh.dish-wireless.com/pubsti.pem

[View certificate details](https://x509.io/?cert=MIIDETCCApegAwIBAgIJAJBOo05uNgrGMAoGCCqGSM49BAMCMIGOMTAwLgYDVQQDDCdOZXROdW1iZXIgU0hBS0VOIFJvb3QgSW50ZXJtZWRpYXRlIENBIDExCzAJBgNVBAYTAlVTMRYwFAYDVQQKDA1OZXROdW1iZXIgSW5jMQswCQYDVQQLDAJVUzEXMBUGA1UECAwOTWFzc2FjaHVzZXR0ZXMxDzANBgNVBAcMBkxvd2VsbDAeFw0yNDEwMDIxNTU2MjZaFw0yNTEwMDIxNTU2MjZaMIGLMSgwJgYDVQQDDB9ESVNIIFdpcmVsZXNzIEwuTC5DLlNIQUtFTi40OTBKMQswCQYDVQQGEwJVUzEWMBQGA1UECgwNRGlzaCBXaXJlbGVzczEcMBoGA1UECwwTU2VjdXJpdHkgT3BlcmF0aW9uczELMAkGA1UECAwCQ08xDzANBgNVBAcMBkRlbnZlcjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABKIxgxO2Un%2B4Nyrp%2BEkoOtXl7MgWh6q3faE7U6Rrwz8dFJ%2BjZ9ZlNe202C26RPTmAqSwdk5DmEE%2FYDHJrCWGpkqjgd4wgdswFgYIKwYBBQUHARoECjAIoAYWBDQ5MEowDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHwYDVR0jBBgwFoAUcS%2FIgtyo4CLj36Bo%2BfheXITe5b0wHQYDVR0OBBYEFL0aFp7wb495NIwZQgRjufV5K4V%2FMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAaBgNVHSABAf8EEDAOMAwGCmCGSAGG%2FwkBAQEwCgYIKoZIzj0EAwIDaAAwZQIwBla1d5eOTQgNxkqOlVm9iMFZm3i0Yy7KxuR4Wcn6gT6S5smP%2FAreA75rjg4Bwx8FAjEAtp6z78dX4C%2Fc9%2B2lLqD8T%2FIdrXxlEeVwpu4rkIkB19pUNJl6pAKK%2FMZjOUt%2BbnZ4)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 490J', but common name is 'DISH Wireless L.L.C.SHAKEN.490J' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |
| [e_atis_ext_certificate_policies](../../ISSUES/e_atis_ext_certificate_policies/README.md) | error | ATIS1000080 | The Certificate Policies extension is marked as critical |


Generated: 18 Aug 25 21:13 UTC