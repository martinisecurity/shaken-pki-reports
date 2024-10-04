# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN ONE OWL TELECOM INC 412K

Tested At: 04 Oct 24 16:15 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -565 day(s)\
Subject: CN=SHAKEN ONE OWL TELECOM INC 412K, OU=ONEOWLTELECOM, O=ONE OWL TELECOM INC, ST=Maryland, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/ONE_OWL_TELECOM_INC_412K

[View certificate details](https://x509.io/?cert=MIIC5TCCAougAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkYb4wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDIxNzEwMDA0NVoXDTIzMDMxOTEwMDA0NVowgYAxCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhNYXJ5bGFuZDEcMBoGA1UECgwTT05FIE9XTCBURUxFQ09NIElOQzEWMBQGA1UECwwNT05FT1dMVEVMRUNPTTEoMCYGA1UEAwwfU0hBS0VOIE9ORSBPV0wgVEVMRUNPTSBJTkMgNDEySzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABILltt1rRucTYd1VEjaGA689%2Bhvtph8O496CsQ2m7YmUbAGOG07vJDwBHhtG68Ql6I33Gs4fXjY4aN1MQyKX01CjgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDQxMkswFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBQ53VJLBv%2FgIdSzL42m4k3O%2F0kxfjAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIhAOytvtYyqbV0nTrKiAM9b%2BBauB41q%2BvDO67XPyhy7abcAiAFv2tBEY6TL4YB8Cg6Fk12YZx4UN8D3FIhhaL6geWMSg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 412K', but common name is 'SHAKEN ONE OWL TELECOM INC 412K' |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 04 Oct 24 16:29 UTC