# STIR/SHAKEN CA Ecosystem Compliance

## Certificate DISH Wireless L.L.C.SHAKEN.490J

Tested At: 26 Aug 24 18:46 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 59 day(s)\
Subject: L=Denver, ST=CO, OU=Security Operations, O=Dish Wireless, C=US, CN=DISH Wireless L.L.C.SHAKEN.490J\
Issuer: L=Lowell, ST=Massachusettes, OU=US, O=NetNumber Inc, C=US, CN=NetNumber SHAKEN Root Intermediate CA 1\
Link: https://stsh.dish-wireless.com/pubsti.pem

[View certificate details](https://x509.io/?cert=MIIDETCCApegAwIBAgIJALzAJTVoHAF3MAoGCCqGSM49BAMCMIGOMTAwLgYDVQQDDCdOZXROdW1iZXIgU0hBS0VOIFJvb3QgSW50ZXJtZWRpYXRlIENBIDExCzAJBgNVBAYTAlVTMRYwFAYDVQQKDA1OZXROdW1iZXIgSW5jMQswCQYDVQQLDAJVUzEXMBUGA1UECAwOTWFzc2FjaHVzZXR0ZXMxDzANBgNVBAcMBkxvd2VsbDAeFw0yMzEwMjUxNTQ4MTZaFw0yNDEwMjQxNTQ4MTZaMIGLMSgwJgYDVQQDDB9ESVNIIFdpcmVsZXNzIEwuTC5DLlNIQUtFTi40OTBKMQswCQYDVQQGEwJVUzEWMBQGA1UECgwNRGlzaCBXaXJlbGVzczEcMBoGA1UECwwTU2VjdXJpdHkgT3BlcmF0aW9uczELMAkGA1UECAwCQ08xDzANBgNVBAcMBkRlbnZlcjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABKIxgxO2Un%2B4Nyrp%2BEkoOtXl7MgWh6q3faE7U6Rrwz8dFJ%2BjZ9ZlNe202C26RPTmAqSwdk5DmEE%2FYDHJrCWGpkqjgd4wgdswFgYIKwYBBQUHARoECjAIoAYWBDQ5MEowDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHwYDVR0jBBgwFoAUcS%2FIgtyo4CLj36Bo%2BfheXITe5b0wHQYDVR0OBBYEFL0aFp7wb495NIwZQgRjufV5K4V%2FMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAaBgNVHSABAf8EEDAOMAwGCmCGSAGG%2FwkBAQEwCgYIKoZIzj0EAwIDaAAwZQIwAcahm6hIg8MITcIyDOCcZxFHw2GNUKFTTLlhHwLzJbGk7fbu0qfF0TAwtV6ezdVfAjEApGMxWkkjmTZbo9u0NwCPq8cmB8dXfh6yQiqbXrjX78J0qEUX3lOs7BQbSeed4B4a)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |
| [e_atis_ext_certificate_policies](../../ISSUES/e_atis_ext_certificate_policies/README.md) | error | ATIS1000080 | The Certificate Policies extension is marked as critical |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 490J', but common name is 'DISH Wireless L.L.C.SHAKEN.490J' |


Generated: 26 Aug 24 18:49 UTC