# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Convoso 758J

Tested At: 06 Nov 22 17:28 UTC\
Initial Validity Period: 35 day(s)\
Remaining Validity Period: 34 day(s)\
Subject: emailAddress=stir-shaken@convoso.com, CN=SHAKEN Convoso 758J, OU=Infrastructure, O=Convoso, ST=California, C=US, emailAddress=stir-shaken@convoso.com\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://stirshaken.s3.us-west-1.amazonaws.com/stirshaken.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC9zCCAp6gAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkVJgwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTEwNTA2MDAzMFoXDTIyMTIxMDA2MDAzMFowgZMxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRAwDgYDVQQKDAdDb252b3NvMRcwFQYDVQQLDA5JbmZyYXN0cnVjdHVyZTEcMBoGA1UEAwwTU0hBS0VOIENvbnZvc28gNzU4SjEmMCQGCSqGSIb3DQEJARYXc3Rpci1zaGFrZW5AY29udm9zby5jb20wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQr0l5ZxKiYcoEQtP8DiZgZ22gPBvqCK41AW85shlZGPWjj6HD%2Ffq0GCeb9vaOdVm4VW%2FoTE03pP2jgHWjNgSfco4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ3NThKMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQU%2FUmADFIbdJOik3WEEsHeQZ5jFlkwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0cAMEQCIAaF3zRzqBRtdob2tZuch4%2F33V%2BFnp%2BLAHkas9YBpNxjAiA%2Fsi5cdu5AyxZnz72VNeCgogPj5eTluu0jDi0eO1kPHA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_cp_1_3_subject_email](../../ISSUES/w_cp_1_3_subject_email/README.md) | warn | US_SHAKEN_CP | Email addresses are not allowed as the CP does not specify how to validate them |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 758J' |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 06 Nov 22 17:28 UTC