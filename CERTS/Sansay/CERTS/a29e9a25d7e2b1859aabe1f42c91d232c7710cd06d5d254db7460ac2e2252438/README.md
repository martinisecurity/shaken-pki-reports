# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Apeiron Systems 012J

Tested At: 08 Feb 23 19:35 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 9 day(s)\
Subject: CN=SHAKEN Apeiron Systems 012J, OU=NOC, O=Apeiron Systems, ST=California, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Apeiron_012J

[View certificate details](https://understandingwebpki.com/?cert=MIIC1DCCAnqgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkXcUwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDExNzIxMzg1NFoXDTIzMDIxNjIxMzg1NFowcDELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExGDAWBgNVBAoMD0FwZWlyb24gU3lzdGVtczEMMAoGA1UECwwDTk9DMSQwIgYDVQQDDBtTSEFLRU4gQXBlaXJvbiBTeXN0ZW1zIDAxMkowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATvM7voHFBJOHVADYo1gKHK2wU0lgJXrz%2FJPCTX0AKFzXtzFfZF0xEGkCL70IlycCO4vTtcloIGOIJTSN8ZSPxko4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQwMTJKMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUfHBaMnt53HaA60SExkm4zJE0Wl4wHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIQDN1v3%2BEVwyoxEqxu0mIwP8k2IG6xhAnXA8%2BfGPPggcCgIgVZSi2Ul2pjTyzCZlTQ%2BszVmE6ZGo4bJajYeuD0XkdlU%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 012J' |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 08 Feb 23 19:45 UTC