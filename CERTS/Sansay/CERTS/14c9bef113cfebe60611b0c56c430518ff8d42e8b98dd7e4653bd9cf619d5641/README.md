# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Ytel Inc. 703J

Tested At: 27 Nov 23 23:20 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 79 day(s)\
Subject: CN=SHAKEN Ytel Inc. 703J, OU=Connect, O=Ytel Inc., ST=California, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/YtelVoice20230214

[View certificate details](https://understandingwebpki.com/?cert=MIICyzCCAnKgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkYWswCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDIxNDE3MTIzOFoXDTI0MDIxNDE3MTIzOFowaDELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExEjAQBgNVBAoMCVl0ZWwgSW5jLjEQMA4GA1UECwwHQ29ubmVjdDEeMBwGA1UEAwwVU0hBS0VOIFl0ZWwgSW5jLiA3MDNKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEpyjhdlT3rpIHVGPi6cZ52IOv2y7dhv3ls7n%2FPO6TNuhp2DFuGw%2FSYYptOMaYbyw8Y8rudlGpq3cHa7rJ1%2FENcKOB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYENzAzSjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFDPhXure7G2QobiIsc4JnXnK%2FH%2FuMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNHADBEAiBUcDvCHWNnBctBV2DGYU%2B0uqlLGLD79MdVDVB00TTVpwIgAfAqw7z0E9EvoLiAp7J4X5vf8WHgwODpAmzh3sjR1FM%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 703J', but common name is 'SHAKEN Ytel Inc. 703J' |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 27 Nov 23 23:28 UTC