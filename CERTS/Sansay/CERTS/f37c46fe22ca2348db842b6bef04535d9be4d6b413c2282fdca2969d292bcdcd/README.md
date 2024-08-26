# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Watchcomm  0590

Tested At: 26 Aug 24 18:15 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 178 day(s)\
Subject: CN=SHAKEN Watchcomm  0590, OU=Watchcomm\\ , O=Watchcomm\\ , ST=Ohio, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/0590/429C7C70711E3820F0B8E1DEAE6FF326226492AD.pem

[View certificate details](https://x509.io/?cert=MIICyzCCAnGgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkkq0wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTI0MDIyMTEzMzQ1OVoXDTI1MDIyMDEzMzQ1OVowZzELMAkGA1UEBhMCVVMxDTALBgNVBAgMBE9oaW8xEzARBgNVBAoMCldhdGNoY29tbSAxEzARBgNVBAsMCldhdGNoY29tbSAxHzAdBgNVBAMMFlNIQUtFTiBXYXRjaGNvbW0gIDA1OTAwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQad%2BOICqHr2w5yw9c%2F1w3OzZjFxQ2OmJB6l%2Bdr%2F7LEjdamvYNuIsAWRz1Q%2F4G10g4jE38WCMcXr5A6n6KdWQBIo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQwNTkwMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQUGvmovusZ%2BEsdLoBfdb3VR9GkQoAwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIGur3Jpud%2F%2Bb%2FB5voNOKoHKsqqppP%2BCPhLxiwvsLJbEWAiEA5ULIavijndjzREXPtcEhYUet4YXvLRdvq2b6vfP3LBY%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 0590', but common name is 'SHAKEN Watchcomm  0590' |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |


Generated: 26 Aug 24 18:49 UTC