# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN i3 Broadband 5800

Tested At: 21 Aug 23 20:12 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -87 day(s)\
Subject: emailAddress=doug@i3broadband.com, CN=SHAKEN i3 Broadband 5800, OU=NOC, O=i3 Broadband, ST=Illinois, C=US, emailAddress=doug@i3broadband.com\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/5800/order/2_5800_165

[View certificate details](https://understandingwebpki.com/?cert=MIIC8TCCApigAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkbMMwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDQyNjE1MDM0M1oXDTIzMDUyNjE1MDM0M1owgY0xCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhJbGxpbm9pczEVMBMGA1UECgwMaTMgQnJvYWRiYW5kMQwwCgYDVQQLDANOT0MxITAfBgNVBAMMGFNIQUtFTiBpMyBCcm9hZGJhbmQgNTgwMDEjMCEGCSqGSIb3DQEJARYUZG91Z0BpM2Jyb2FkYmFuZC5jb20wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQG9vxbou7w97hFuz8q6f4tNsy4FUDmwCuhd6EF%2Fq0q2UDZvDoOzgLq%2BO2GPlwwM0%2BFRSOKP3%2Fma%2BQokb%2BN6%2FVuo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ1ODAwMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQUPT4UG9fn9y%2BXTwtELPd3isHRMSEwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0cAMEQCIFzNEvTI4mWsZQ95AUwTOLCmDN9roRV3PbgB7%2B3QPLPmAiBuzNAr%2F68%2FsfGFMV%2Fsj7HcX5N7Xqy5I58RWKUNEM5OVg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 5800' |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_cp_1_3_subject_email](../../ISSUES/w_cp_1_3_subject_email/README.md) | warn | US_SHAKEN_CP | Email addresses are not allowed as the CP does not specify how to validate them |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 21 Aug 23 20:18 UTC