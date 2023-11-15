# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Vumber LLC 225K

Tested At: 15 Nov 23 18:03 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 364 day(s)\
Subject: CN=SHAKEN Vumber LLC 225K, OU=VEHMP, O=Vumber LLC, ST=Pennsylvania, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Vumber_LLC_225K

[View certificate details](https://understandingwebpki.com/?cert=MIICzjCCAnSgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkhkcwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMTExNDEzNDUxNFoXDTI0MTExMzEzNDUxNFowajELMAkGA1UEBhMCVVMxFTATBgNVBAgMDFBlbm5zeWx2YW5pYTETMBEGA1UECgwKVnVtYmVyIExMQzEOMAwGA1UECwwFVkVITVAxHzAdBgNVBAMMFlNIQUtFTiBWdW1iZXIgTExDIDIyNUswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASN8XK9iv2A9tlv73jGyIGZqUOR7XImz4NrsH5wgeY7CUW5fljc%2B9tKkrJNRm%2Fv4I52pEL1JnTMoh71C48PCNBPo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQyMjVLMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUYNgIwAw%2By2ELf0EY84CtcxefjNQwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIAgtQU5SFbhcx22oehH55AfQhRqmKXMelLuKN0%2BrBUhRAiEAhOytmE%2BMkFCZMEQNzyZvhR4hg5WnN16ArSxZs4vgBg0%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 225K' |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 15 Nov 23 18:10 UTC