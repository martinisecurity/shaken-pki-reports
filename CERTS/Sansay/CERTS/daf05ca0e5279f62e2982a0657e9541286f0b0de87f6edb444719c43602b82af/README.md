# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN ConvergeTel LLC 388K

Tested At: 21 Aug 23 20:14 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -164 day(s)\
Subject: CN=SHAKEN ConvergeTel LLC 388K, OU=IT Depart, O=ConvergeTel LLC, ST=Delaware, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/ConvergeTel_LLC_388K

[View certificate details](https://understandingwebpki.com/?cert=MIIC2TCCAn6gAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkYM8wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDIwODE5MzMyM1oXDTIzMDMxMDE5MzMyM1owdDELMAkGA1UEBhMCVVMxETAPBgNVBAgMCERlbGF3YXJlMRgwFgYDVQQKDA9Db252ZXJnZVRlbCBMTEMxEjAQBgNVBAsMCUlUIERlcGFydDEkMCIGA1UEAwwbU0hBS0VOIENvbnZlcmdlVGVsIExMQyAzODhLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEFkAq%2BvVxVgHj7C2jq4jrR3WQIPbKRm884zCY4XKO%2FApZufpFHtdqOo0lYIo2HcMSR8rnPzI1itlikNEAUJmam6OB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYEMzg4SzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFNElzzoPf9TVTMStsSlGKjgTUXkHMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNJADBGAiEAus0UtNWrw5p6ex6OwNQ81Is5Ig%2B5xd9ot75kBb2WagICIQDZJYriM7sfrY9%2FrSAnALizzbIzKzoccGhrCM5WbFwZvQ%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 388K' |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 21 Aug 23 20:18 UTC