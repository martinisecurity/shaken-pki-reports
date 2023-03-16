# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN i3 Broadband 5800

Tested At: 16 Mar 23 19:07 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 28 day(s)\
Subject: CN=SHAKEN i3 Broadband 5800, OU=NOC, O=i3 Broadband, ST=Illinois, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/i3_Broadband_5800

[View certificate details](https://understandingwebpki.com/?cert=MIICzTCCAnKgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkZM0wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDMxMzE5MzYwM1oXDTIzMDQxMjE5MzYwM1owaDELMAkGA1UEBhMCVVMxETAPBgNVBAgMCElsbGlub2lzMRUwEwYDVQQKDAxpMyBCcm9hZGJhbmQxDDAKBgNVBAsMA05PQzEhMB8GA1UEAwwYU0hBS0VOIGkzIEJyb2FkYmFuZCA1ODAwMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEBvb8W6Lu8Pe4Rbs%2FKun%2BLTbMuBVA5sAroXehBf6tKtlA2bw6Ds4C6vjthj5cMDNPhUUjij9%2F5mvkKJG%2Fjev1bqOB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYENTgwMDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFD0%2BFBvX5%2Fcvl08LRCz3d4rB0TEhMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNJADBGAiEA5O4IIvukdA6m11mn19LVqOwe5wfbXVM9Y2gbg%2Bp385gCIQCusZ0n5sN%2BufSY7gX5dtcBwWjKZrCYQBVW8cZZCVaLNw%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 5800' |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 16 Mar 23 19:18 UTC