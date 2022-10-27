# STIR/SHAKEN CA Ecosystem Compliance
## Sansay

### Certificate 0a796b424529d808ab634e5294713b1739cbedea
Tested At: 2022-10-27 18:55:36 +0000 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 321 day(s)\
Subject: CN=SHAKEN Bulk Solutions\\, LLC 644J, OU=Voice Engineering, O=Bulk Solutions\\, LLC, ST=Delaware, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://bulkvs-cr.s3.amazonaws.com/644J_20210615001.crt

View: [Click to view](https://understandingwebpki.com/?cert=MIIDmDCCAz2gAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkSZAwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMDkxMjE5Mzc1N1oXDTIzMDkxMjE5Mzc1N1owgYQxCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhEZWxhd2FyZTEcMBoGA1UECgwTQnVsayBTb2x1dGlvbnMsIExMQzEaMBgGA1UECwwRVm9pY2UgRW5naW5lZXJpbmcxKDAmBgNVBAMMH1NIQUtFTiBCdWxrIFNvbHV0aW9ucywgTExDIDY0NEowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQJ%2BfKPjj1zc6BLNOXT3Te%2BjxZMizox8kcU5hoHmgDd9ZKvAPdIn2FpygHxROC0YUlbpWmyNyfv4QM2RmgceC8Jo4IBiDCCAYQwFgYIKwYBBQUHARoECjAIoAYWBDY0NEowFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBSEqLJYf%2BOMsVCAvZ3RzON79P2shjCBygYDVR0jBIHCMIG%2FgBSs05P1Q0PMCr5FWBcTfZJ83MMBRqGBkKSBjTCBijELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExEjAQBgNVBAcMCVNhbiBEaWVnbzEbMBkGA1UECgwSU2Fuc2F5IENvcnBvcmF0aW9uMRIwEAYDVQQLDAlTYW5zYXkgQ0ExITAfBgNVBAMMGFNIQUtFTiBTYW5zYXkgUm9vdCBDQSBVU4IUFLVfOAX18HsTtfiw3u0g8lFwPpowRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0kAMEYCIQCVLALwPgav0e7V44he6Z6TGNFPAtEfvQ6KYeK7GUllpwIhALBcdowOlk44wBjZ%2BvIC5coOAYeiTOMralcYF4DSdqEX)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_sti_certificate_policies | error | ATIS-1000080v4 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| w_cp1_3_subject_rdn_unknown | warn | CPv1.3 | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_sti_subject_cn | error | ATIS-1000080v4 | Common name shall contain the text string 'SHAKEN 644J' |
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |
| e_cp1_3_ambiguous_identifier | error | CPv1.3 | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |

Generated: 27/10/2022 at 18:57:26