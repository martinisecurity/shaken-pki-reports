# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN i3 Broadband 5800

Tested At: 24 Nov 23 11:09 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -193 day(s)\
Subject: CN=SHAKEN i3 Broadband 5800, OU=NOC, O=i3 Broadband, ST=Illinois, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/i3_Broadband_5800

[View certificate details](https://understandingwebpki.com/?cert=MIICzDCCAnKgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkarwwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDQxNDIyNTUzNVoXDTIzMDUxNDIyNTUzNVowaDELMAkGA1UEBhMCVVMxETAPBgNVBAgMCElsbGlub2lzMRUwEwYDVQQKDAxpMyBCcm9hZGJhbmQxDDAKBgNVBAsMA05PQzEhMB8GA1UEAwwYU0hBS0VOIGkzIEJyb2FkYmFuZCA1ODAwMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEBvb8W6Lu8Pe4Rbs%2FKun%2BLTbMuBVA5sAroXehBf6tKtlA2bw6Ds4C6vjthj5cMDNPhUUjij9%2F5mvkKJG%2Fjev1bqOB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYENTgwMDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFD0%2BFBvX5%2Fcvl08LRCz3d4rB0TEhMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiAWdl09jacGNwvKDDl60%2B2XWaG4dK0QuRLo%2FCH3Yt6flwIhAPuqCDI99ucBjBC5A42DbSZJK9%2F%2FXOWYbec1PSeT8Emw)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_crl_distribution_struct](../../ISSUES/e_atis_ext_crl_distribution_struct/README.md) | error | ATIS1000080 | CRL Distribution Point shall contain a CRLIssuer field |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 5800', but common name is 'SHAKEN i3 Broadband 5800' |


Generated: 24 Nov 23 11:17 UTC