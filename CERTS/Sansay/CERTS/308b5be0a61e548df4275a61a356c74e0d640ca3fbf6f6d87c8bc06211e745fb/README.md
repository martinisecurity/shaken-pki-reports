# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN WWT INC dba VoIP Networks 053K

Tested At: 31 Jan 23 21:42 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 287 day(s)\
Subject: CN=SHAKEN WWT INC dba VoIP Networks 053K, OU=Tier3, O=WWT INC dba VoIP Networks, ST=New Jersey, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/WWT_INC_dba_VoIP_Networks_053K

[View certificate details](https://understandingwebpki.com/?cert=MIIC6zCCApGgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkVVwwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTExNDE2NTI0M1oXDTIzMTExNDE2NTI0M1owgYYxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApOZXcgSmVyc2V5MSIwIAYDVQQKDBlXV1QgSU5DIGRiYSBWb0lQIE5ldHdvcmtzMQ4wDAYDVQQLDAVUaWVyMzEuMCwGA1UEAwwlU0hBS0VOIFdXVCBJTkMgZGJhIFZvSVAgTmV0d29ya3MgMDUzSzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABHqQUmpKdeH4pDq4Gor6kKgFEkIlU%2BmdBJ2SSMDjvO0yskQPfW%2BNrC60OaqOcGSwM4rnqQuDa%2FoGPJlN2NoOTWijgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDA1M0swFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBSA4DeNDRAWiEJtd14RPDLqdfsvTjAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIgYSVQhfY60diPQ5wS5yNxMmPU9X788hGNbQGjEGhXa64CIQCWqqoKQOGWAoJ4LMlSTHiemTEZc1Oz%2BCJEEo%2Flbn6eXA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 053K' |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 31 Jan 23 21:50 UTC