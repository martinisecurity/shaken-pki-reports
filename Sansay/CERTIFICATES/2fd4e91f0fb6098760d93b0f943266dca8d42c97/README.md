# STIR/SHAKEN CA Ecosystem Compliance
## Sansay

### Certificate 2fd4e91f0fb6098760d93b0f943266dca8d42c97
Tested At: 2022-10-27 21:25:33 +0000 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 30 day(s)\
Subject: CN=SHAKEN Touchtone 683A, OU=NOC, O=Touchtone, ST=New Jersey, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/Touchtone_683A

View: [Click to view](https://understandingwebpki.com/?cert=MIICyDCCAm6gAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkU3gwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAyNzAwMzU0OFoXDTIyMTEyNjAwMzU0OFowZDELMAkGA1UEBhMCVVMxEzARBgNVBAgMCk5ldyBKZXJzZXkxEjAQBgNVBAoMCVRvdWNodG9uZTEMMAoGA1UECwwDTk9DMR4wHAYDVQQDDBVTSEFLRU4gVG91Y2h0b25lIDY4M0EwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATwPwhwR259HmHeCt9A84ohVtuEKJYTrlFMI5HXCls08mFEPsyKfoW%2BySkQmaGCaCmCEHCD6rRfhou6X8dWmxB4o4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ2ODNBMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQU4fP7CWiieHL3QK4YpzPbWW6r8hkwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIQDtF15YENMlhuGQajB1eSYt3ZpPnNvLbAjhzft79LTeIAIgTf0z8Rn%2FyJgq7QoayJjrIQZA%2FF77N40%2BGqdCglFg0Ds%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_sti_certificate_policies | error | ATIS-1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| e_cp1_3_ambiguous_identifier | error | United States SHAKEN CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_sti_subject_cn | error | ATIS-1000080 | Common name shall contain the text string 'SHAKEN 683A' |
| e_cp1_3_subject_sn | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| w_cp1_3_subject_rdn_unknown | warn | United States SHAKEN CP | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |

Generated: 27/10/2022 at 21:27:34