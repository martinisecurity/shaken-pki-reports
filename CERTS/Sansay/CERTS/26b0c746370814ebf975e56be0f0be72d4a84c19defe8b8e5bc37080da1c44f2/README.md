# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Convoso 758J

Tested At: 02 Nov 22 07:52 UTC\
Initial Validity Period: 35 day(s)\
Remaining Validity Period: 8 day(s)\
Subject: emailAddress=stir-shaken@convoso.com, CN=SHAKEN Convoso 758J, OU=Infrastructure, O=Convoso, ST=California, C=US, emailAddress=stir-shaken@convoso.com\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://stirshaken.s3.us-west-1.amazonaws.com/stirshaken.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIDpjCCA0ygAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkTPMwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAwNjA2MDAyNVoXDTIyMTExMDA2MDAyNVowgZMxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRAwDgYDVQQKDAdDb252b3NvMRcwFQYDVQQLDA5JbmZyYXN0cnVjdHVyZTEcMBoGA1UEAwwTU0hBS0VOIENvbnZvc28gNzU4SjEmMCQGCSqGSIb3DQEJARYXc3Rpci1zaGFrZW5AY29udm9zby5jb20wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQr0l5ZxKiYcoEQtP8DiZgZ22gPBvqCK41AW85shlZGPWjj6HD%2Ffq0GCeb9vaOdVm4VW%2FoTE03pP2jgHWjNgSfco4IBiDCCAYQwFgYIKwYBBQUHARoECjAIoAYWBDc1OEowFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBT9SYAMUht0k6KTdYQSwd5BnmMWWTCBygYDVR0jBIHCMIG%2FgBSs05P1Q0PMCr5FWBcTfZJ83MMBRqGBkKSBjTCBijELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExEjAQBgNVBAcMCVNhbiBEaWVnbzEbMBkGA1UECgwSU2Fuc2F5IENvcnBvcmF0aW9uMRIwEAYDVQQLDAlTYW5zYXkgQ0ExITAfBgNVBAMMGFNIQUtFTiBTYW5zYXkgUm9vdCBDQSBVU4IUFLVfOAX18HsTtfiw3u0g8lFwPpwwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIQCLM8SA0npNWq540uYEG%2BBo72boLwayj%2FdCZCdJ6slMQwIgEOBf9fNtJFj20ycBour0sHVQ7b7GitKfJoepi0MWuro%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_cp_1_3_subject_email](../../ISSUES/w_cp_1_3_subject_email/README.md) | warn | US_SHAKEN_CP | Email addresses are not allowed as the CP does not specify how to validate them |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 758J' |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 02 Nov 22 07:52 UTC