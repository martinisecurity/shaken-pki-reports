# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J

Tested At: 04 Oct 24 15:56 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -433 day(s)\
Subject: emailAddress=voiceops@clearfly.net, CN=SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J, OU=NOC, O=Greenfly Networks Inc dba Clearfly Communications, ST=Montana, C=US, emailAddress=voiceops@clearfly.net\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/210J/order/459_210J_63

[View certificate details](https://x509.io/?cert=MIIDPDCCAuKgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkdvUwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDYyOTE1MjAwN1oXDTIzMDcyOTE1MjAwN1owgdcxCzAJBgNVBAYTAlVTMRAwDgYDVQQIDAdNb250YW5hMTowOAYDVQQKDDFHcmVlbmZseSBOZXR3b3JrcyBJbmMgZGJhIENsZWFyZmx5IENvbW11bmljYXRpb25zMQwwCgYDVQQLDANOT0MxRjBEBgNVBAMMPVNIQUtFTiBHcmVlbmZseSBOZXR3b3JrcyBJbmMgZGJhIENsZWFyZmx5IENvbW11bmljYXRpb25zIDIxMEoxJDAiBgkqhkiG9w0BCQEWFXZvaWNlb3BzQGNsZWFyZmx5Lm5ldDBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABLd80UROkyOsAvdgaJlDDkYwWMba9%2FmHqLhUC6XioSi0IgY4SJ2NNx9lZWamYeNejzTkK%2FK%2Fhd6a7KVZ2krA5eKjgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDIxMEowFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBTP7%2FKdOTL%2BH67cbjLgPaI%2B1dnKFjAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIhAPsLB1vlT5QmdOIKZ5fnT4hFU10Y1csPMx6O%2BEDAwynNAiBMsJZJKUulMBwlAcCqRx2daoAgpqGPWpmZfoZQODYiaA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 210J', but common name is 'SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J' |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |


Generated: 04 Oct 24 16:29 UTC