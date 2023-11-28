# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN ConvergeTel LLC 388K

Tested At: 28 Nov 23 20:11 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 72 day(s)\
Subject: CN=SHAKEN ConvergeTel LLC 388K, OU=IT Depart, O=ConvergeTel LLC, ST=Delaware, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/ConvergeStirShaken

[View certificate details](https://understandingwebpki.com/?cert=MIIC2DCCAn6gAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkYM0wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDIwODE5MDc1MVoXDTI0MDIwODE5MDc1MVowdDELMAkGA1UEBhMCVVMxETAPBgNVBAgMCERlbGF3YXJlMRgwFgYDVQQKDA9Db252ZXJnZVRlbCBMTEMxEjAQBgNVBAsMCUlUIERlcGFydDEkMCIGA1UEAwwbU0hBS0VOIENvbnZlcmdlVGVsIExMQyAzODhLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEFkAq%2BvVxVgHj7C2jq4jrR3WQIPbKRm884zCY4XKO%2FApZufpFHtdqOo0lYIo2HcMSR8rnPzI1itlikNEAUJmam6OB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYEMzg4SzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFNElzzoPf9TVTMStsSlGKjgTUXkHMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiBGpE71pt9M5spmCYTkbXiJwlSqtAM5ojUcafhicX7v4QIhALQJsLkrgSt9i0PK%2FjcSM5BERsMlPO7e0G4HJdFAerYW)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 388K', but common name is 'SHAKEN ConvergeTel LLC 388K' |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 28 Nov 23 20:21 UTC